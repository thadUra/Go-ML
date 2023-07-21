package game

type Shot struct {
	DISTANCE_X int // width
	DISTANCE_Y int // height
}

/**
 *  Initialize shot location
*   Default to penalty shot location
*/
func InitShot(x int, y int, def bool) Shot {
	var shot Shot
	if def {
		shot.DISTANCE_X = 112 // penalty
		shot.DISTANCE_Y = 36  // penalty
	} else {
		shot.DISTANCE_X = x
		shot.DISTANCE_Y = y
	}
	return shot
}
