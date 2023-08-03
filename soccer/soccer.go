package soccer

import (
	"Soccer-Penalty-Kick-ML-Threading/nn"
	"errors"
	"math"
	"math/rand"
)

/**
 *  Soccer Struct
 *  Implements environment interface
 *  Contains position and field information
 */
type Soccer struct {
	pos               Position
	field             Field
	ACTION_SIZE       int
	OBSERVATION_SIZE  int
	ACTION_SPACE      [][]float64
	OBSERVATION_SPACE [][]float64
	shotModel         *nn.Network
}

/**
 *  InitSoccer() Func
 *  Generates soccer env
 *  Action space conist of 9 actions: dribble in any of the 8 directions or shoot the ball
 *  Observation space consist of just the location on field
 */
func InitSoccer() Soccer {
	var env Soccer
	env.pos = GeneratePos(0, 0, true)
	env.field = GenerateField(0, 0, 0, 0, 0, 0, 0, true)
	env.ACTION_SIZE = 1
	env.OBSERVATION_SIZE = 2
	env.ACTION_SPACE = append(env.ACTION_SPACE, []float64{0, 8})
	env.OBSERVATION_SPACE = append(env.OBSERVATION_SPACE, []float64{0, float64(env.field.FIELD_WIDTH)})
	env.OBSERVATION_SPACE = append(env.OBSERVATION_SPACE, []float64{0, float64(env.field.FIELD_HEIGHT)})
	env.shotModel = nil
	return env
}

/**
 *  Step() Func
 *  Performs one step in soccer environment
 *  Can either dribble or shoot ball from current position
 *  Shot parameters picked by other ml model
 */
func (scr *Soccer) Step(
	action []float64,
) (float64, bool, error) {
	// Check dimensions
	if len(action) > 1 {
		return -1, true, errors.New("soccer.Step: action dimensions are incorrect")
	}

	// Perform action (WIP WITH MANUAL SHOT PARAMS)
	if action[0] == 0 {
		// Get shot params from ml model if initialized
		var action []float64
		if scr.shotModel != nil {
			return -1, true, errors.New("soccer.Step: shot model not initialized")
		} else {
			action = []float64{-15.0 * math.Pi / 180.0, 25.0 * math.Pi / 180.0, 25.0} // Manual pararms temporarily
			// action = []float64{
			// 	((rand.Float64() * 160) - 80) * math.Pi / 180.0,
			// 	(rand.Float64() * 90) * math.Pi / 180.0,
			// 	rand.Float64() * 150.0} // Random params temporarily
		}
		// Check errors with any model prediction

		result, err := scr.field.Shoot(scr.pos, action, true)
		if err != nil {
			return -1, true, err
		} else {
			if result == "GOAL" {
				return 500, true, nil
			} else if result == "SAVED" {
				return 50, true, nil
			} else if result == "POST" || result == "CROSSBAR" {
				return 25, true, nil
			} else {
				return -500, true, nil
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
			return -500, true, nil
		}
		return 1, false, nil
	}
}

/**
 *  Reset() Func
 *  Resets position to random spot on field
 */
func (scr *Soccer) Reset() {
	scr.pos = GeneratePos(rand.Float64()*scr.field.FIELD_HEIGHT, rand.Float64()*scr.field.FIELD_WIDTH, false)
}

/**
 *  GetNumActions() Func
 *  Accessor for action space size
 */
func (scr *Soccer) GetNumActions() int {
	return len(scr.ACTION_SPACE)
}

/**
 *  GetNumObservations() Func
 *  Accessor for observation space size
 */
func (scr *Soccer) GetNumObservations() int {
	return len(scr.OBSERVATION_SPACE)
}
