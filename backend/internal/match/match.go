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
	delete(m.Game.Players, playerID)
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
			stateJSON, err := json.Marshal(m.Game)
			if err != nil {
				fmt.Println("❌ Failed to serialize game state for match", m.ID, ":", err)
				m.mu.Unlock()
				time.Sleep(time.Millisecond * time.Duration(tick))
				continue
			}
			for _, p := range m.Players {
				if p.Conn != nil {
					if err := p.Conn.WriteMessage(websocket.TextMessage, stateJSON); err != nil {
						fmt.Println("❌ Failed to write message to player", p.ID, "in match", m.ID, ":", err)
						_ = p.Conn.Close()
						delete(m.Players, p.ID)
					}
				}
			}

			m.mu.Unlock()
			time.Sleep(time.Millisecond * time.Duration(tick))
		}
	}()
}
