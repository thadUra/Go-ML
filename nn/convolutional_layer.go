package nn

import (
	"gonum.org/v1/gonum/mat"
)

/**
 * ConvolutionalLayer Struct
 * WIP
 */
type ConvolutionalLayer struct {
	INPUT  *mat.Dense
	OUTPUT *mat.Dense
}

/**
 * InitConvolutionalLayer()
 * WIP
 */
func InitConvolutionalLayer(input_size int, output_size int) *ConvolutionalLayer {
	// Initialize and return
	var layer ConvolutionalLayer
	return &layer
}

/**
 * ForwardPropagation()
 * Performs forward propagation for a convolutional layer required by Layer interface
 * WIP
 */
func (layer *ConvolutionalLayer) ForwardPropagation(input *mat.Dense) *mat.Dense {

	return layer.OUTPUT
}

/**
 * BackPropagation()
 * Performs back propagation for a convolutional layer required by Layer interface
 * WIP
 */
func (layer *ConvolutionalLayer) BackPropagation(output_error *mat.Dense, learning_rate float64) *mat.Dense {

	return layer.INPUT
}
