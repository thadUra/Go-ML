package game

import "math"

type Environment struct {
	args              Soccer
	ACTION_SPACE      [][]float64
	OBSERVATION_SPACE [][]float64
}

func InitEnvironment(args Soccer) Environment {

	horizontal_angle := 180 * math.Pi / 180 // radians
	vertical_angle := 90 * math.Pi / 180    // radians
	power := 145.0                          // feet per second

	var env Environment
	env.args = args
	env.ACTION_SPACE = append(env.ACTION_SPACE, []float64{0, vertical_angle})
	env.ACTION_SPACE = append(env.ACTION_SPACE, []float64{-horizontal_angle, horizontal_angle})
	env.ACTION_SPACE = append(env.ACTION_SPACE, []float64{0, power})
	env.OBSERVATION_SPACE = append(env.OBSERVATION_SPACE, []float64{0, float64(env.args.FIELD_WIDTH)})
	env.OBSERVATION_SPACE = append(env.OBSERVATION_SPACE, []float64{0, float64(env.args.FIELD_HEIGHT)})
	return env
}

func (env Environment) step(action []float64, location Shot) int {

	// 0: miss, 1: hit post/crossbar, 2: goal
	goal := 2

	horizontal_angle := action[0]
	vertical_angle := action[1]
	power := action[2]
	left_post := env.args.GetLeftPostAngle(location)
	right_post := env.args.GetRightPostAngle(location)
	crossbar := env.args.GetCrossbarAngle(location)

	// Check horizontal direction
	// if horizontal_angle > left_post && horizontal_angle < right_post {
	// 	goal = 0
	// }

	// Check height

	if goal == 2 {
		return 10
	} else if goal == 1 {
		return -1
	} else {
		return -20
	}
}
