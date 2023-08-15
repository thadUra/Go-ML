package nn

import (
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

// DenseLayer implements the Layer interface for a Dense layer in a neural network.
type DenseLayer struct {
	input   *mat.Dense
	output  *mat.Dense
	weights *mat.Dense
	bias    *mat.Dense
}

// NewDenseLayer returns a new instance of a dense layer.
func NewDenseLayer(input_size int, output_size int) *DenseLayer {
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
	layer.weights = mat.NewDense(input_size, output_size, weight)
	layer.bias = mat.NewDense(1, output_size, bias)
	return &layer
}

// ForwardPropagation implements the Layer interface and returns a matrix after
// performing forward propagation.
func (layer *DenseLayer) ForwardPropagation(input *mat.Dense) *mat.Dense {
	layer.input = input
	var dot mat.Dense
	dot.Mul(layer.input, layer.weights)
	layer.output = &dot
	layer.output.Add(&dot, layer.bias)
	return layer.output
}

// BackPropagation implements the Layer interface and returns the error matrix after
// performing back propagation.
func (layer *DenseLayer) BackPropagation(output_error *mat.Dense, learning_rate float64) *mat.Dense {
	// Initalize matrices for manipulation
	var input_error, weights_error, new_weight, new_bias mat.Dense
	input_error.Mul(output_error, layer.weights.T())
	weights_error.Mul(layer.input.T(), output_error)
	new_weight.Scale(learning_rate, &weights_error)
	new_bias.Scale(learning_rate, output_error)

	// Update weights and bias and return input error
	layer.weights.Sub(layer.weights, &new_weight)
	layer.bias.Sub(layer.bias, &new_bias)
	return &input_error
}
