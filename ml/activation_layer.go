package ml

import (
	"gonum.org/v1/gonum/mat"
)

type Activation func(input *mat.Dense) *mat.Dense

type ActivationLayer struct {
	INPUT           *mat.Dense
	OUTPUT          *mat.Dense
	ACTIVATION      Activation
	ACTIVATIONPRIME Activation
}

func InitActivationLayer(activation Activation, activationPrime Activation) *ActivationLayer {
	var layer ActivationLayer
	layer.ACTIVATION = activation
	layer.ACTIVATIONPRIME = activationPrime
	return &layer
}

func (layer *ActivationLayer) ForwardPropagation(input *mat.Dense) *mat.Dense {

	// fmt.Printf(" ACT_LAYER\n")

	layer.INPUT = input
	layer.OUTPUT = layer.ACTIVATION(layer.INPUT)
	return layer.OUTPUT
}

func (layer *ActivationLayer) BackPropagation(output_error *mat.Dense, learning_rate float64) *mat.Dense {
	// fmt.Printf(" ACT_LAYER\n")
	prime := layer.ACTIVATIONPRIME(layer.INPUT)
	rows, cols := prime.Dims()
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			temp := prime.At(r, c) * output_error.At(r, c)
			prime.Set(r, c, temp)
		}
	}
	return prime
}
