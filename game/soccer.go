package game

import "math"

type Soccer struct {
	FIELD_WIDTH   int // feet
	FIELD_HEIGHT  int // feet
	NET_WIDTH     int // feet
	NET_HEIGHT    int // feet
	BALL_WEIGHT   int // grams
	BALL_DIAMETER int // inches
}

/**
 *  Initialize soccer field
*   Default to fifa regulations
*/
func InitSoccer(field_width int, field_height int, net_width int, net_height int, weight int, diameter int, def bool) Soccer {
	var env Soccer
	if def {
		env.FIELD_HEIGHT = 345 // fifa regulations roughly
		env.FIELD_WIDTH = 224  // fifa regulations roughly
		env.NET_WIDTH = 24
		env.NET_HEIGHT = 8
		env.BALL_WEIGHT = 430
		env.BALL_DIAMETER = 9
	} else {
		env.FIELD_HEIGHT = field_height
		env.FIELD_WIDTH = field_width
		env.NET_WIDTH = net_width
		env.NET_HEIGHT = net_height
		env.BALL_WEIGHT = weight
		env.BALL_DIAMETER = diameter
	}
	return env
}

func (env Soccer) GetLeftPostAngle(shot Shot) float64 {
	width := (env.FIELD_WIDTH / 2) - (env.NET_WIDTH / 2) - shot.DISTANCE_X
	height := shot.DISTANCE_Y
	return math.Atan(float64(width / height))
}

func (env Soccer) GetRightPostAngle(shot Shot) float64 {
	width := (env.FIELD_WIDTH / 2) + (env.NET_WIDTH / 2) - shot.DISTANCE_X
	height := shot.DISTANCE_Y
	return math.Atan(float64(width / height))
}

func (env Soccer) GetCrossbarAngle(shot Shot) float64 {
	width := (env.FIELD_WIDTH / 2) - shot.DISTANCE_X
	length := math.Sqrt(float64((shot.DISTANCE_Y * shot.DISTANCE_Y) + (width * width)))
	return math.Atan(float64(env.NET_HEIGHT) / length)
}
