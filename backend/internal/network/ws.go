package network

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"hockey-game/internal/match"
	"hockey-game/internal/player"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var activeMatch = match.New("match-1")

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	playerID := activeMatch.GeneratePlayerID()
	p := player.New(playerID, conn)

	if err := activeMatch.Join(p); err != nil {
		log.Println("Join error:", err)
		_ = conn.WriteMessage(websocket.TextMessage, []byte("Match full"))
		return
	}
	defer activeMatch.Leave(p.ID)

	log.Println("🟢 Player connected:", p.ID)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("🔴 Player disconnected:", p.ID)
			return
		}

		// Decode input JSON
		var inputMsg struct {
			Type string `json:"type"`
			Key  string `json:"key"`
		}

		if err := json.Unmarshal(msg, &inputMsg); err != nil {
			log.Printf("Invalid JSON from player %s: %v (raw: %s)", p.ID, err, string(msg))
			continue
		}

		if inputMsg.Type == "input" {
			switch inputMsg.Key {
			case "ArrowUp":
				p.Input.Up = true
			case "ArrowDown":
				p.Input.Down = true
			case "ArrowLeft":
				p.Input.Left = true
			case "ArrowRight":
				p.Input.Right = true
			}
		}
	}
}
