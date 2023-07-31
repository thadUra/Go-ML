package ml

import "gonum.org/v1/gonum/mat"

type Activation func(input *mat.Dense) *mat.Dense

type ActivationLayer struct {
	INPUT           *mat.Dense
	OUTPUT          *mat.Dense
	ACTIVATION      Activation
	ACTIVATIONPRIME Activation
}

func InitActivationLayer(activation Activation, activationPrime Activation) ActivationLayer {
	var layer ActivationLayer
	layer.ACTIVATION = activation
	layer.ACTIVATIONPRIME = activationPrime
	return layer
}

func (layer *ActivationLayer) ForwardPropagation(input *mat.Dense) *mat.Dense {
	layer.INPUT = input
	layer.OUTPUT = layer.ACTIVATION(layer.INPUT)
	return layer.OUTPUT
}

func (layer *ActivationLayer) BackPropagation(output_error *mat.Dense) *mat.Dense {
	var prime *mat.Dense
	prime.Mul(layer.ACTIVATIONPRIME(layer.INPUT), output_error)
	return prime
}
