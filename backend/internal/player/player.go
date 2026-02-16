package player

import "github.com/gorilla/websocket"

type Vector struct {
	X, Y float64
}

type InputState struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

type Player struct {
	ID       string
	Position Vector
	Conn     *websocket.Conn
	Input    InputState
}

func New(id string, conn *websocket.Conn) *Player {
	return &Player{
		ID:       id,
		Conn:     conn,
		Position: Vector{X: 10, Y: 10},
		Input:    InputState{},
	}
}
