package ml

type Layer struct {
	INPUT  int
	OUTPUT int
}

func (node Layer) ForwardPropagation(input int) float32 {
	return 0
}

func (node Layer) BackPropagation(output_error float32, learning_rate float32) float32 {
	return 0
}
