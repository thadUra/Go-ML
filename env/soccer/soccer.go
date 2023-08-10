// The soccer package contains code for generating the soccer environment.
package soccer

import (
	"errors"
	"math"

	"github.com/thadUra/Golang-Machine-Learning/env"
	"github.com/thadUra/Golang-Machine-Learning/nn"
)

// Soccer represents the soccer environment and its parameters like Position and Field.
type Soccer struct {
	pos               Position
	field             Field
	ACTION_SIZE       int
	OBSERVATION_SIZE  int
	ACTION_SPACE      [][]float64
	OBSERVATION_SPACE [][]float64
	shotModel         *nn.Network
}

// NewSoccer generates a new environment instance given default parameters for the position and field dimensions.
func NewSoccer() env.Environment {
	var scr Soccer
	scr.pos = GeneratePos(0, 0, true)
	scr.field = GenerateField(0, 0, 0, 0, 0, 0, 0, true)
	scr.ACTION_SIZE = 1
	scr.OBSERVATION_SIZE = 2
	scr.ACTION_SPACE = append(scr.ACTION_SPACE, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8})
	scr.OBSERVATION_SPACE = append(scr.OBSERVATION_SPACE, []float64{0, float64(scr.field.FIELD_WIDTH)})
	scr.OBSERVATION_SPACE = append(scr.OBSERVATION_SPACE, []float64{0, float64(scr.field.FIELD_HEIGHT)})
	scr.shotModel = nil
	return env.Environment(&scr)
}

// Step performs one action inside the environment.
func (scr *Soccer) Step(
	action []float64,
) (float64, float64, bool, error) {
	// Current state
	state := scr.OBSERVATION_SPACE[0][1]*scr.pos.DISTANCE_Y + scr.pos.DISTANCE_X

	// Check dimensions
	if len(action) > 1 {
		return state, -1, true, errors.New("soccer.Step: action dimensions are incorrect")
	}

	// Perform action (WIP WITH MANUAL SHOT PARAMS)
	if action[0] == 0 {
		// Get shot params from ml model if initialized
		var action []float64
		if scr.shotModel != nil {
			return state, -1, true, errors.New("soccer.Step: shot model not initialized")
		} else {
			// Manually shoot between posts at half power
			width_angle := math.Atan(((scr.GetField().FIELD_WIDTH / 2) - scr.GetPos().DISTANCE_X) / scr.GetPos().DISTANCE_Y)
			action = []float64{width_angle, 10.0 * math.Pi / 180.0, 30}
		}
		// Check errors with any model prediction

		result, err := scr.field.Shoot(scr.pos, action, false)
		if err != nil {
			return state, -1, true, err
		} else {
			if result == "GOAL" {
				return state, 1000, true, nil
			} else if result == "SAVED" {
				return state, 50, true, nil
			} else if result == "POST" || result == "CROSSBAR" {
				return state, 0, true, nil
			} else {
				return state, -500, true, nil
			}
		}
	} else {
		// Dribble Actions
		if action[0] == 1 {
			scr.pos.DribbleUp()
		} else if action[0] == 2 {
			scr.pos.DribbleUpRight()
		} else if action[0] == 3 {
			scr.pos.DribbleRight()
		} else if action[0] == 4 {
			scr.pos.DribbleDownRight()
		} else if action[0] == 5 {
			scr.pos.DribbleDown()
		} else if action[0] == 6 {
			scr.pos.DribbleDownLeft()
		} else if action[0] == 7 {
			scr.pos.DribbleLeft()
		} else {
			scr.pos.DribbleUpLeft()
		}
		// Check if position is out of bounds
		if scr.pos.OutOfBounds(scr.field) {
			// fmt.Printf("	OUT OF BOUNDS AT (%f,%f)\n", scr.pos.DISTANCE_X, scr.pos.DISTANCE_Y)
			return state, -500, true, nil
		}
		state = scr.OBSERVATION_SPACE[0][1]*scr.pos.DISTANCE_Y + scr.pos.DISTANCE_X
		return state, -1, false, nil
	}
}

// Reset sets the current state to the mid backfield.
func (scr *Soccer) Reset() float64 {
	scr.pos = GeneratePos(112, 300, false)
	return scr.OBSERVATION_SPACE[0][1]*scr.pos.DISTANCE_Y + scr.pos.DISTANCE_X
}

// GetNumActions returns the size of the action space.
func (scr *Soccer) GetNumActions() int {
	return len(scr.ACTION_SPACE[0])
}

// GetNumObservations returns the size of the observation space.
func (scr *Soccer) GetNumObservations() int {
	return int((scr.OBSERVATION_SPACE[0][1] + 1) * (scr.OBSERVATION_SPACE[1][1] + 1))
}

// GetPos returns the current Position object of the environment.
func (scr *Soccer) GetPos() Position {
	return scr.pos
}

// GetField returns the current Field object of the environment.
func (scr *Soccer) GetField() Field {
	return scr.field
}
