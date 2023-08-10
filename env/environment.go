// The env package contains code for generating a custom gym environment or to
// replicate a similar environment from the gym library.
package env

// Environment is a representation of a gym environment, which needs to implement the
// Step, Reset, GetNumActions, and GetNumObservations functions.
type Environment interface {
	Step(action []float64) (float64, float64, bool, error)
	Reset() float64
	GetNumActions() int
	GetNumObservations() int
}
