package nn

import "gonum.org/v1/gonum/mat"

/**
 * Layer Interface
 * Defines required functions for any layer
 */
type Layer interface {
	ForwardPropagation(input *mat.Dense) *mat.Dense
	BackPropagation(output_error *mat.Dense, learning_rate float64) *mat.Dense
}
