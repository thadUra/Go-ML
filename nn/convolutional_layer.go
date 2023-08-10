package nn

import (
	"gonum.org/v1/gonum/mat"
)

// ConvolutionalLayer implements the Layer interface for a convolutional layer (CURRENT WIP).
type ConvolutionalLayer struct {
	INPUT  *mat.Dense
	OUTPUT *mat.Dense
}

// NewConvolutionalLayer returns a new instance of a convolutional layer (CURRENT WIP).
func NewConvolutionalLayer(input_size int, output_size int) *ConvolutionalLayer {
	// Initialize and return
	var layer ConvolutionalLayer
	return &layer
}

// ForwardPropagation implements the Layer interface and returns a matrix after
// performing forward propagation (CURRENT WIP).
func (layer *ConvolutionalLayer) ForwardPropagation(input *mat.Dense) *mat.Dense {

	return layer.OUTPUT
}

// BackPropagation implements the Layer interface and returns the error matrix after
// performing back propagation (CURRENT WIP).
func (layer *ConvolutionalLayer) BackPropagation(output_error *mat.Dense, learning_rate float64) *mat.Dense {

	return layer.INPUT
}
