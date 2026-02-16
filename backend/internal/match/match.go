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

type Match struct {
	ID      string
	Players map[string]*player.Player
	Status  string
	Game    game.GameState
	mu      sync.Mutex
}

func New(id string) *Match {
	return &Match{
		ID:      id,
		Players: make(map[string]*player.Player),
		Status:  "waiting",
		Game:    game.InitGameState(),
	}
}

func (m *Match) Join(p *player.Player) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.Players) >= 2 {
		return fmt.Errorf("match full")
	}

	m.Players[p.ID] = p
	fmt.Println("🟢 Player joined:", p.ID)

	if len(m.Players) == 2 {
		m.Status = "playing"
		fmt.Println("🏁 Match", m.ID, "starts!")
		m.StartGameLoop()
	}

	return nil
}

func (m *Match) Leave(playerID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.Players, playerID)
	fmt.Println("🔴 Player left:", playerID)

	if len(m.Players) < 2 {
		m.Status = "waiting"
		fmt.Println("⏸️ Match", m.ID, "waiting for players...")
	}
}

func (m *Match) StartGameLoop() {
	go func() {
		tick := 100
		for {
			m.mu.Lock()
			if m.Status != "playing" {
				m.mu.Unlock()
				time.Sleep(time.Millisecond * time.Duration(tick))
				continue
			}

			// Update puck
			game.Update(&m.Game)

			// Appliquer inputs joueurs
			game.ApplyPlayerInputs(&m.Game, m.Players)

			// Broadcast à tous les joueurs
			stateJSON, _ := json.Marshal(m.Game)
			for _, p := range m.Players {
				if p.Conn != nil {
					p.Conn.WriteMessage(websocket.TextMessage, stateJSON)
				}
			}

			m.mu.Unlock()
			time.Sleep(time.Millisecond * time.Duration(tick))
		}
	}()
}
