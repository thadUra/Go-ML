package ml

import "gonum.org/v1/gonum/mat"

type Layer struct {
	INPUT  *mat.Dense
	OUTPUT *mat.Dense
}

func (node Layer) ForwardPropagation(input int) float32 {
	return 0
}

func (node Layer) BackPropagation(output_error float32, learning_rate float32) float32 {
	return 0
}
