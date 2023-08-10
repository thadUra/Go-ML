package nn

import "gonum.org/v1/gonum/mat"

// Layer is a representation for different neural network layers, which require the functions
// ForwardPropagation and BackPropagation.
type Layer interface {
	ForwardPropagation(input *mat.Dense) *mat.Dense
	BackPropagation(output_error *mat.Dense, learning_rate float64) *mat.Dense
}
