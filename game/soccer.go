package game

type Soccer struct {
	FIELD_WIDTH   float64 // feet
	FIELD_HEIGHT  float64 // feet
	NET_WIDTH     float64 // feet
	NET_HEIGHT    float64 // feet
	NET_DIAMETER  float64 // feet
	BALL_WEIGHT   float64 // grams
	BALL_DIAMETER float64 // feet
}

/**
 *  Initialize soccer field
*   Default to fifa regulations for field dimensions
*/
func InitSoccer(
	field_width float64,
	field_height float64,
	net_width float64,
	net_height float64,
	net_diameter float64,
	weight float64,
	ball_diameter float64,
	def bool,
) Soccer {
	var env Soccer
	if def {
		env.FIELD_HEIGHT = 345.0 // fifa regulations roughly
		env.FIELD_WIDTH = 224.0  // fifa regulations roughly
		env.NET_WIDTH = 24.0
		env.NET_HEIGHT = 8.0
		env.NET_DIAMETER = 2.0 / 3.0
		env.BALL_WEIGHT = 430.0
		env.BALL_DIAMETER = 0.75
	} else {
		env.FIELD_HEIGHT = field_height
		env.FIELD_WIDTH = field_width
		env.NET_WIDTH = net_width
		env.NET_HEIGHT = net_height
		env.NET_DIAMETER = net_diameter
		env.BALL_WEIGHT = weight
		env.BALL_DIAMETER = ball_diameter
	}
	return env
}
