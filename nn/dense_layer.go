package nn

import (
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

type FCLayer struct {
	INPUT   *mat.Dense
	OUTPUT  *mat.Dense
	WEIGHTS *mat.Dense
	BIAS    *mat.Dense
}

func InitFCLayer(input_size int, output_size int) *FCLayer {
	// Initialize weights and bias matrices
	weight := make([]float64, input_size*output_size)
	bias := make([]float64, output_size)
	for i := range weight {
		weight[i] = 2*rand.Float64() - 1.0
	}
	for i := range bias {
		bias[i] = 2*rand.Float64() - 1.0
	}

	// Return FCLayer
	var layer FCLayer
	layer.WEIGHTS = mat.NewDense(input_size, output_size, weight)
	layer.BIAS = mat.NewDense(1, output_size, bias)
	return &layer
}

func (layer *FCLayer) ForwardPropagation(input *mat.Dense) *mat.Dense {
	layer.INPUT = input
	var dot mat.Dense
	dot.Mul(layer.INPUT, layer.WEIGHTS)
	layer.OUTPUT = &dot
	layer.OUTPUT.Add(&dot, layer.BIAS)
	return layer.OUTPUT
}

func (layer *FCLayer) BackPropagation(output_error *mat.Dense, learning_rate float64) *mat.Dense {

	var input_error, weights_error, new_weight, new_bias mat.Dense

	input_error.Mul(output_error, layer.WEIGHTS.T())
	weights_error.Mul(layer.INPUT.T(), output_error)

	new_weight.Scale(learning_rate, &weights_error)
	new_bias.Scale(learning_rate, output_error)

	// fc := mat.Formatted(layer.WEIGHTS, mat.Prefix("     "), mat.Squeeze())
	// fmt.Printf("\nOW = %v\n", fc)
	// fc = mat.Formatted(layer.BIAS, mat.Prefix("     "), mat.Squeeze())
	// fmt.Printf("\nOB = %v\n", fc)

	layer.WEIGHTS.Sub(layer.WEIGHTS, &new_weight)
	layer.BIAS.Sub(layer.BIAS, &new_bias)

	// fc = mat.Formatted(layer.WEIGHTS, mat.Prefix("     "), mat.Squeeze())
	// fmt.Printf("\nNW = %v\n", fc)
	// fc = mat.Formatted(layer.BIAS, mat.Prefix("     "), mat.Squeeze())
	// fmt.Printf("\nNB = %v\n", fc)

	return &input_error
}
