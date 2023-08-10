package soccer

// Position represents current X and Y coordinate position on the field.
type Position struct {
	DISTANCE_X float64
	DISTANCE_Y float64
}

// GeneratePos returns a new position. If `def` is true, it returns a default
// position at the free kick spot.
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

// OutOfBounds returns true if position is out of bounds on the field.
func (pos *Position) OutOfBounds(f Field) bool {
	return pos.DISTANCE_X < 0 || pos.DISTANCE_Y < 0 || pos.DISTANCE_X > f.FIELD_WIDTH || pos.DISTANCE_Y > f.FIELD_HEIGHT
}

// Modify the position to dribble up.
func (pos *Position) DribbleUp() {
	pos.DISTANCE_Y++
}

// Modify the position to dribble up right.
func (pos *Position) DribbleUpRight() {
	pos.DISTANCE_Y++
	pos.DISTANCE_X--
}

// Modify the position to dribble right.
func (pos *Position) DribbleRight() {
	pos.DISTANCE_X--
}

// Modify the position to dribble down right.
func (pos *Position) DribbleDownRight() {
	pos.DISTANCE_Y--
	pos.DISTANCE_X--
}

// Modify the position to dribble down.
func (pos *Position) DribbleDown() {
	pos.DISTANCE_Y--
}

// Modify the position to dribble down left.
func (pos *Position) DribbleDownLeft() {
	pos.DISTANCE_Y--
	pos.DISTANCE_X++
}

// Modify the position to dribble left.
func (pos *Position) DribbleLeft() {
	pos.DISTANCE_X++
}

// Modify the position to dribble up left.
func (pos *Position) DribbleUpLeft() {
	pos.DISTANCE_Y++
	pos.DISTANCE_X++
}
