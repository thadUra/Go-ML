package environment

/**
 *  Environment Interface
 *  Outlines necessary accessors and modifiers for environment
 */
type Environment interface {
	Step(action []float64) (float64, float64, bool, error)
	Reset() float64
	GetNumActions() int
	GetNumObservations() int
}
