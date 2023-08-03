package soccer

/**
 *  Environment Interface
 *  Outlines necessary accessors and modifiers for environment
 */
type EnvironmentTest interface {
	Step(action []float64) (float64, bool, error)
	Reset()
	GetNumActions() int
	GetNumObservations() int
}
