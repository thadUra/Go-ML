package nn

import (
	"gonum.org/v1/gonum/mat"
)

// Activation defines a type for any activation function.
type Activation func(input *mat.Dense) *mat.Dense

// ActivationLayer implements the Layer interface for an Activation layer in a neural
// network.
type ActivationLayer struct {
	INPUT           *mat.Dense
	OUTPUT          *mat.Dense
	ACTIVATION      Activation
	ACTIVATIONDERIV Activation
}

// NewActivationLayer returns a new instance of an activation layer.
func NewActivationLayer(activation Activation, activationDeriv Activation) *ActivationLayer {
	var layer ActivationLayer
	layer.ACTIVATION = activation
	layer.ACTIVATIONDERIV = activationDeriv
	return &layer
}

// ForwardPropagation implements the Layer interface and returns a matrix after
// performing the activation function on the layer values.
func (layer *ActivationLayer) ForwardPropagation(input *mat.Dense) *mat.Dense {
	layer.INPUT = input
	layer.OUTPUT = layer.ACTIVATION(layer.INPUT)
	return layer.OUTPUT
}

// BackPropagation implements the Layer interface and returns a matrix after
// performing the derivative activation function on the layer values.
func (layer *ActivationLayer) BackPropagation(output_error *mat.Dense, learning_rate float64) *mat.Dense {
	deriv := layer.ACTIVATIONDERIV(layer.INPUT)
	rows, cols := deriv.Dims()
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			deriv.Set(r, c, deriv.At(r, c)*output_error.At(r, c))
		}
	}
	return deriv
}
