package nn

import (
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

/**
 * DenseLayer Struct
 * Contains weights and bias for a dense / fully connected layer in a neural network
 */
type DenseLayer struct {
	INPUT   *mat.Dense
	OUTPUT  *mat.Dense
	WEIGHTS *mat.Dense
	BIAS    *mat.Dense
}

/**
 * InitDenseLayer()
 * Initializes a dense layer given input and output shape
 */
func InitDenseLayer(input_size int, output_size int) *DenseLayer {
	// Initialize weights and bias matrices with random values
	weight := make([]float64, input_size*output_size)
	bias := make([]float64, output_size)
	for i := range weight {
		weight[i] = 2*rand.Float64() - 1.0
	}
	for i := range bias {
		bias[i] = 2*rand.Float64() - 1.0
	}

	// Return layer
	var layer DenseLayer
	layer.WEIGHTS = mat.NewDense(input_size, output_size, weight)
	layer.BIAS = mat.NewDense(1, output_size, bias)
	return &layer
}

/**
 * ForwardPropagation()
 * Performs forward propagation for a dense layer required by Layer interface
 * Returns dot product of input and weights plus the bias
 */
func (layer *DenseLayer) ForwardPropagation(input *mat.Dense) *mat.Dense {
	layer.INPUT = input
	var dot mat.Dense
	dot.Mul(layer.INPUT, layer.WEIGHTS)
	layer.OUTPUT = &dot
	layer.OUTPUT.Add(&dot, layer.BIAS)
	return layer.OUTPUT
}

/**
 * BackPropagation()
 * Performs back propagation for a dense layer required by Layer interface
 * Returns matrix of the input error from the current layer
 */
func (layer *DenseLayer) BackPropagation(output_error *mat.Dense, learning_rate float64) *mat.Dense {
	// Initalize matrices for manipulation
	var input_error, weights_error, new_weight, new_bias mat.Dense
	input_error.Mul(output_error, layer.WEIGHTS.T())
	weights_error.Mul(layer.INPUT.T(), output_error)
	new_weight.Scale(learning_rate, &weights_error)
	new_bias.Scale(learning_rate, output_error)

	// Update weights and bias and return input error
	layer.WEIGHTS.Sub(layer.WEIGHTS, &new_weight)
	layer.BIAS.Sub(layer.BIAS, &new_bias)
	return &input_error
}
