package game

import (
	"math"
	"hockey-game/internal/player"
)

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

// Applique les inputs des joueurs en utilisant un modèle simple d'accélération + friction
// tickMs : durée d'une frame en millisecondes (ex: 16 pour ~60 FPS)
func ApplyPlayerInputs(state *GameState, players map[string]*player.Player, tickMs int) {
	// Paramètres de référence (tuned pour tick = 100ms)
	const accelBase = 10     // augmentation pour des changements de direction plus rapides
	const maxSpeedBase = 60 // vitesse maximale plus élevée
	const frictionBase = 0.98 // friction par tick de référence (100ms) — proche de 1 pour glisse

	// Adapter les paramètres à la durée réelle du tick
	scale := float64(tickMs) / 100.0
	accel := accelBase * scale      // accel max changée par frame
	maxSpeed := maxSpeedBase * scale
	friction := math.Pow(frictionBase, scale)

	for _, p := range players {
		// calculer la direction d'accélération en fonction des inputs
		dx := 0.0
		dy := 0.0
		if p.Input.Up {
			dy -= 1
		}
		if p.Input.Down {
			dy += 1
		}
		if p.Input.Left {
			dx -= 1
		}
		if p.Input.Right {
			dx += 1
		}

		if dx == 0 && dy == 0 {
			// pas d'input : appliquer friction (glisse) pour ralentir progressivement
			p.Velocity.X *= friction
			p.Velocity.Y *= friction
		} else {
			// normaliser la direction pour garder la même magnitude en diagonale
			lenDir := math.Hypot(dx, dy)
			dx /= lenDir
			dy /= lenDir

			// desired velocity (direction * maxSpeed)
			dxDesired := dx * maxSpeed
			dyDesired := dy * maxSpeed

			// steering = desired - current velocity
			steerX := dxDesired - p.Velocity.X
			steerY := dyDesired - p.Velocity.Y
			steerMag := math.Hypot(steerX, steerY)

			// limiter la magnitude du steer pour n'appliquer qu'une accélération max par frame
			if steerMag > accel {
				steerX = steerX / steerMag * accel
				steerY = steerY / steerMag * accel
			}

			// appliquer le steering
			p.Velocity.X += steerX
			p.Velocity.Y += steerY

			// appliquer quand même une légère friction pour simuler la glisse même en mouvement
			p.Velocity.X *= friction
			p.Velocity.Y *= friction
		}

		// limiter la vitesse après friction
		speed := math.Hypot(p.Velocity.X, p.Velocity.Y)
		if speed > maxSpeed {
			p.Velocity.X = (p.Velocity.X / speed) * maxSpeed
			p.Velocity.Y = (p.Velocity.Y / speed) * maxSpeed
		}

		// mettre à jour la position
		p.Position.X += p.Velocity.X
		p.Position.Y += p.Velocity.Y

		// Reset inputs pour next tick
		p.Input = player.InputState{}

		// Exporter l'état du joueur
		state.Players[p.ID] = Vector{X: p.Position.X, Y: p.Position.Y}
	}
}
