package game

type Shot struct {
	DISTANCE_X float64 // width  in feet
	DISTANCE_Y float64 // height in feet
}

/**
 *  Initialize shot location
*   Default to penalty shot location
*/
func InitShot(
	x float64,
	y float64,
	def bool,
) Shot {
	var shot Shot
	if def {
		// penalty shot distance
		shot.DISTANCE_X = 112.0
		shot.DISTANCE_Y = 36.0
	} else {
		shot.DISTANCE_X = x
		shot.DISTANCE_Y = y
	}
	return shot
}
