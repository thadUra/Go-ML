package soccer

/**
 *  Environment Interface
 *  Outlines necessary accessors and modifiers for environment
 */
type EnvironmentTest interface {
	Step(action []float64) (float64, bool, error)
	Reset()
	GetActionSpace() [][]float64
	GetObservationSpace() [][]float64
	GetNumActions() int
	GetNumObservations() int
}
