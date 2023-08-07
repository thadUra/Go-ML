package soccer

/**
 *  Shot Struct
 *  Location on the field where the soccer shot will be taken
 *  Distances are in feet
 */
type Position struct {
	DISTANCE_X float64
	DISTANCE_Y float64
}

/**
 *  GeneratePos()
 *  Generate a position for the shot
 *  Default to penalty shot location
 */
func GeneratePos(
	x float64,
	y float64,
	def bool,
) Position {
	var pos Position
	if def {
		pos.DISTANCE_X = 112.0
		pos.DISTANCE_Y = 36.0
	} else {
		pos.DISTANCE_X = x
		pos.DISTANCE_Y = y
	}
	return pos
}

/**
 *  OutOfBounds()
 *  Determines if position is out of bounds on the field
 */
func (pos *Position) OutOfBounds(f Field) bool {
	return pos.DISTANCE_X < 0 || pos.DISTANCE_Y < 0 || pos.DISTANCE_X > f.FIELD_WIDTH || pos.DISTANCE_Y > f.FIELD_HEIGHT
}

/**
 *  DribbleDir()
 *  Modifier functions to assist in dribbling action
 */
func (pos *Position) DribbleUp() {
	pos.DISTANCE_Y++
}
func (pos *Position) DribbleUpRight() {
	pos.DISTANCE_Y++
	pos.DISTANCE_X--
}
func (pos *Position) DribbleRight() {
	pos.DISTANCE_X--
}
func (pos *Position) DribbleDownRight() {
	pos.DISTANCE_Y--
	pos.DISTANCE_X--
}
func (pos *Position) DribbleDown() {
	pos.DISTANCE_Y--
}
func (pos *Position) DribbleDownLeft() {
	pos.DISTANCE_Y--
	pos.DISTANCE_X++
}
func (pos *Position) DribbleLeft() {
	pos.DISTANCE_X++
}
func (pos *Position) DribbleUpLeft() {
	pos.DISTANCE_Y++
	pos.DISTANCE_X++
}
