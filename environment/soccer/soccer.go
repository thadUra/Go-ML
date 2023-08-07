package soccer

import (
	"Soccer-Penalty-Kick-ML-Threading/environment"
	"Soccer-Penalty-Kick-ML-Threading/nn"
	"errors"
	"math"
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
 *  InitSoccer()
 *  Generates soccer env
 *  Action space conist of 9 actions: dribble in any of the 8 directions or shoot the ball
 *  Observation space consist of just the location on field
 */
func InitSoccer() environment.Environment {
	var env Soccer
	env.pos = GeneratePos(0, 0, true)
	env.field = GenerateField(0, 0, 0, 0, 0, 0, 0, true)
	env.ACTION_SIZE = 1
	env.OBSERVATION_SIZE = 2
	env.ACTION_SPACE = append(env.ACTION_SPACE, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8})
	env.OBSERVATION_SPACE = append(env.OBSERVATION_SPACE, []float64{0, float64(env.field.FIELD_WIDTH)})
	env.OBSERVATION_SPACE = append(env.OBSERVATION_SPACE, []float64{0, float64(env.field.FIELD_HEIGHT)})
	env.shotModel = nil
	return environment.Environment(&env)
}

/**
 *  Step() WIP
 *  Performs one step in soccer environment
 *  Can either dribble or shoot ball from current position
 *  Shot parameters picked by other ml model
 */
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
				return state, 0, true, nil
			} else if result == "POST" || result == "CROSSBAR" {
				return state, 0, true, nil
			} else {
				return state, 0, true, nil
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
			return state, 0, true, nil
		}
		state = scr.OBSERVATION_SPACE[0][1]*scr.pos.DISTANCE_Y + scr.pos.DISTANCE_X
		return state, 0, false, nil
	}
}

/**
 *  Reset()
 *  Resets position to random spot on field
 */
func (scr *Soccer) Reset() float64 {
	scr.pos = GeneratePos(112, 250, false)
	// scr.pos = GeneratePos(rand.Float64()*scr.field.FIELD_HEIGHT, rand.Float64()*scr.field.FIELD_WIDTH, false)
	return scr.OBSERVATION_SPACE[0][1]*scr.pos.DISTANCE_Y + scr.pos.DISTANCE_X
}

/**
 *  GetNumActions()
 *  Accessor for action space size
 */
func (scr *Soccer) GetNumActions() int {
	return len(scr.ACTION_SPACE[0])
}

/**
 *  GetNumObservations()
 *  Accessor for observation space size
 */
func (scr *Soccer) GetNumObservations() int {
	return int(scr.OBSERVATION_SPACE[0][1] * scr.OBSERVATION_SPACE[1][1])
}

/**
 *  GetPos()
 *  Accessor for position
 */
func (scr *Soccer) GetPos() Position {
	return scr.pos
}

/**
 *  GetField()
 *  Accessor for field
 */
func (scr *Soccer) GetField() Field {
	return scr.field
}
