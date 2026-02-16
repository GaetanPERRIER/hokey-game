package game

import "hockey-game/internal/player"

type Vector struct {
	X, Y float64
}

type GameState struct {
	Players map[string]Vector
	Puck    Vector
}

// Créer un état initial
func InitGameState() GameState {
	return GameState{
		Players: make(map[string]Vector),
		Puck:    Vector{X: 50, Y: 50},
	}
}

// Update : update du puck uniquement
func Update(state *GameState) {
	state.Puck.Y += 1
	if state.Puck.Y > 100 {
		state.Puck.Y = 0
	}
}

// Applique les inputs des joueurs
func ApplyPlayerInputs(state *GameState, players map[string]*player.Player) {
	speed := 1.0
	for _, p := range players {
		if p.Input.Up {
			p.Position.Y -= speed
		}
		if p.Input.Down {
			p.Position.Y += speed
		}
		if p.Input.Left {
			p.Position.X -= speed
		}
		if p.Input.Right {
			p.Position.X += speed
		}

		// Reset inputs pour next tick
		p.Input = player.InputState{}
		state.Players[p.ID] = Vector{X: p.Position.X, Y: p.Position.Y}
	}
}
