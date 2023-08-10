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
	NUM_STEPS         float64
	POS               Position
	FIELD             Field
	ACTION_SIZE       int
	OBSERVATION_SIZE  int
	ACTION_SPACE      [][]float64
	OBSERVATION_SPACE [][]float64
	shotModel         *nn.Network
}

// NewSoccer generates a new environment instance given default parameters for the position and field dimensions.
func NewSoccer() env.Environment {
	var scr Soccer
	scr.NUM_STEPS = 0
	scr.POS = GeneratePos(0, 0, true)
	scr.FIELD = GenerateField(0, 0, 0, 0, 0, 0, 0, true)
	scr.ACTION_SIZE = 1
	scr.OBSERVATION_SIZE = 2
	scr.ACTION_SPACE = append(scr.ACTION_SPACE, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8})
	scr.OBSERVATION_SPACE = append(scr.OBSERVATION_SPACE, []float64{0, float64(scr.FIELD.FIELD_WIDTH)})
	scr.OBSERVATION_SPACE = append(scr.OBSERVATION_SPACE, []float64{0, float64(scr.FIELD.FIELD_HEIGHT)})
	scr.shotModel = nil
	return env.Environment(&scr)
}

// Step performs one action inside the environment.
func (scr *Soccer) Step(
	action []float64,
) (float64, float64, bool, error) {
	// Current state
	state := scr.OBSERVATION_SPACE[0][1]*scr.POS.DISTANCE_Y + scr.POS.DISTANCE_X
	scr.NUM_STEPS++

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

		result, time, err := scr.FIELD.Shoot(scr.POS, action, false)
		if err != nil {
			return state, -1, true, err
		} else {
			if result == "GOAL" {
				return state, (10.0 / time), true, nil
			} else if result == "SAVED" {
				return state, (5.0 / time), true, nil
			} else if result == "POST" || result == "CROSSBAR" {
				return state, 0, true, nil
			} else {
				return state, 0, true, nil
			}
		}
	} else {
		// Dribble Actions
		if action[0] == 1 {
			scr.POS.DribbleUp()
		} else if action[0] == 2 {
			scr.POS.DribbleUpRight()
		} else if action[0] == 3 {
			scr.POS.DribbleRight()
		} else if action[0] == 4 {
			scr.POS.DribbleDownRight()
		} else if action[0] == 5 {
			scr.POS.DribbleDown()
		} else if action[0] == 6 {
			scr.POS.DribbleDownLeft()
		} else if action[0] == 7 {
			scr.POS.DribbleLeft()
		} else {
			scr.POS.DribbleUpLeft()
		}
		// Check if position is out of bounds
		if scr.POS.OutOfBounds(scr.FIELD) {
			return state, 0, true, nil
		}
		state = scr.OBSERVATION_SPACE[0][1]*scr.POS.DISTANCE_Y + scr.POS.DISTANCE_X
		return state, 0, false, nil
	}
}

// Reset sets the current state to the mid backfield.
func (scr *Soccer) Reset() float64 {
	scr.POS = GeneratePos(112, 100, false)
	scr.NUM_STEPS = 0
	return scr.OBSERVATION_SPACE[0][1]*scr.POS.DISTANCE_Y + scr.POS.DISTANCE_X
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
	return scr.POS
}

// GetField returns the current Field object of the environment.
func (scr *Soccer) GetField() Field {
	return scr.FIELD
}
