package match

import (
	"encoding/json"
	"fmt"
	"hockey-game/internal/game"
	"hockey-game/internal/player"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Match représente une partie et contient l'état du jeu, les joueurs connectés
// et la synchronisation nécessaire.
type Match struct {
	ID         string
	Game       game.GameState
	Players    map[string]*player.Player
	Status     string
	mu         sync.Mutex
	nextPlayer int
}

// New crée une nouvelle Match, initialise l'état et démarre la boucle de jeu.
func New(id string) *Match {
	m := &Match{
		ID:         id,
		Game:       game.InitGameState(),
		Players:    make(map[string]*player.Player),
		Status:     "waiting",
		nextPlayer: 1,
	}
	go m.StartGameLoop()
	return m
}

// GeneratePlayerID renvoie un identifiant unique pour un joueur.
func (m *Match) GeneratePlayerID() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Attribuer le premier slot libre entre "1" et "2"
	if _, ok := m.Players["1"]; !ok {
		return "1"
	}
	if _, ok := m.Players["2"]; !ok {
		return "2"
	}
	// fallback (ne devrait pas arriver car Join vérifie la taille)
	m.nextPlayer++
	return fmt.Sprintf("%d", m.nextPlayer)
}

// Join ajoute un joueur à la partie si la partie n'est pas pleine.
func (m *Match) Join(p *player.Player) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.Players) >= 2 {
		return fmt.Errorf("match full")
	}
	// Attribuer un id de slot fixe ("1" ou "2") pour ce joueur
	var slot string
	if _, ok := m.Players["1"]; !ok {
		slot = "1"
	} else {
		slot = "2"
	}
	p.ID = slot
	m.Players[slot] = p
	fmt.Printf("[match] Player joined: assigned slot=%s total=%d\n", slot, len(m.Players))
	// Si la position initiale est connue dans l'état du jeu, l'appliquer au joueur
	if v, ok := m.Game.Players[slot]; ok {
		p.Position = player.Vector{X: v.X, Y: v.Y}
	}
	// Si on a 2 joueurs, démarrer la partie
	if len(m.Players) == 2 {
		m.Status = "playing"
		fmt.Printf("[match] Status -> playing\n")
	}
	return nil
}

// Leave enlève un joueur de la partie et ferme sa connexion si nécessaire.
func (m *Match) Leave(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if p, ok := m.Players[id]; ok {
		if p.Conn != nil {
			_ = p.Conn.Close()
		}
		delete(m.Players, id)
		fmt.Printf("[match] Player left: id=%s total=%d\n", id, len(m.Players))
	}
	// Supprimer également l'entrée dans l'état du jeu afin que
	// le joueur ne soit plus renvoyé dans les broadcasts au front.
	delete(m.Game.Players, id)
	// si moins de 2 joueurs, repasser en attente
	if len(m.Players) < 2 {
		m.Status = "waiting"
		fmt.Printf("[match] Status -> waiting\n")
	}
}

func (m *Match) StartGameLoop() {
	// Placer les joueurs à leurs positions de départ (1 -> gauche, 2 -> droite)
	m.Game.Players["1"] = game.Vector{X: 160, Y: 360}
	m.Game.Players["2"] = game.Vector{X: 1120, Y: 360}

	go func() {
		// tickMs en millisecondes (~60 FPS)
		tickMs := 16
		for {
			m.mu.Lock()
			if m.Status != "playing" {
				m.mu.Unlock()
				time.Sleep(time.Millisecond * time.Duration(tickMs))
				continue
			}

			// Update puck
			game.Update(&m.Game)

			// Appliquer inputs joueurs avec tick en ms
			game.ApplyPlayerInputs(&m.Game, m.Players, tickMs)

			// Copy game state and player connections under lock
			stateJSON, err := json.Marshal(m.Game)
			if err != nil {
				fmt.Printf("⚠️ Error marshaling game state: %v\n", err)
				m.mu.Unlock()
				time.Sleep(time.Millisecond * time.Duration(tickMs))
				continue
			}

			// Create a snapshot of player connections
			type connSnapshot struct {
				playerID string
				conn     *websocket.Conn
			}
			conns := make([]connSnapshot, 0, len(m.Players))
			for id, p := range m.Players {
				if p.Conn != nil {
					conns = append(conns, connSnapshot{
						playerID: id,
						conn:     p.Conn,
					})
				}
			}

			m.mu.Unlock()

			// Broadcast to all players without holding the lock
			// Collect failed player IDs for cleanup
			var failedPlayers []string
			for _, snapshot := range conns {
				err := snapshot.conn.WriteMessage(websocket.TextMessage, stateJSON)
				if err != nil {
					fmt.Printf("⚠️ Error writing to player %s: %v\n", snapshot.playerID, err)
					failedPlayers = append(failedPlayers, snapshot.playerID)
				}
			}

			// Cleanup disconnected players
			for _, playerID := range failedPlayers {
				m.Leave(playerID)
			}

			time.Sleep(time.Millisecond * time.Duration(tickMs))
		}
	}()
}
