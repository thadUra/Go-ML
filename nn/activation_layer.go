package nn

import (
	"gonum.org/v1/gonum/mat"
)

/**
 * Activation Function
 * Defines format of any activation function
 */
type Activation func(input *mat.Dense) *mat.Dense

/**
 * ActivationLayer Struct
 * Contains the activation function and its derivative for the layer
 */
type ActivationLayer struct {
	INPUT           *mat.Dense
	OUTPUT          *mat.Dense
	ACTIVATION      Activation
	ACTIVATIONDERIV Activation
}

/**
 * InitActivationLayer()
 * Initializes an activation the activation function and its derivative
 */
func InitActivationLayer(activation Activation, activationDeriv Activation) *ActivationLayer {
	var layer ActivationLayer
	layer.ACTIVATION = activation
	layer.ACTIVATIONDERIV = activationDeriv
	return &layer
}

/**
 * ForwardPropagation()
 * Performs forward propagation for an activation layer required by Layer interface
 * Returns output matrix with the function performed
 */
func (layer *ActivationLayer) ForwardPropagation(input *mat.Dense) *mat.Dense {
	layer.INPUT = input
	layer.OUTPUT = layer.ACTIVATION(layer.INPUT)
	return layer.OUTPUT
}

/**
 * BackPropagation()
 * Performs back propagation for an activation layer required by Layer interface
 * Returns matrix of error generated from derivative of activation function
 */
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
