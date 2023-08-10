package nn

import (
	"gonum.org/v1/gonum/mat"
)

// ConvolutionalLayer implements the Layer interface for a flat layer (CURRENT WIP).
type FlattenLayer struct {
	INPUT        *mat.Dense
	OUTPUT       *mat.Dense
	INPUT_SHAPE  int
	OUTPUT_SHAPE int
}

// NewFlattenLayer returns a new instance of a flat layer (CURRENT WIP).
func NewFlattenLayer(input_size int, output_size int) *FlattenLayer {
	// Initialize and return
	var layer FlattenLayer
	return &layer
}

// ForwardPropagation implements the Layer interface and returns a matrix after
// performing forward propagation (CURRENT WIP).
func (layer *FlattenLayer) ForwardPropagation(input *mat.Dense) *mat.Dense {
	layer.INPUT = input
	rows, cols := input.Dims()
	flat := make([]float64, rows*cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			flat[(cols*r)+c] = input.At(r, c)
		}
	}
	layer.OUTPUT = mat.NewDense(1, rows*cols, flat)
	return layer.OUTPUT
}

// BackPropagation implements the Layer interface and returns the error matrix after
// performing back propagation (CURRENT WIP).
func (layer *FlattenLayer) BackPropagation(output_error *mat.Dense, learning_rate float64) *mat.Dense {
	rows, cols := output_error.Dims()
	arr := make([]float64, rows*cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			arr[(cols*r)+c] = output_error.At(r, c)
		}
	}
	rows, cols = layer.INPUT.Dims()
	reshape := mat.NewDense(rows, cols, arr)
	return reshape
}

// GetShape returns the flat layers dimensions
func (layer *FlattenLayer) GetShape() (int, int) {
	return layer.INPUT_SHAPE, layer.OUTPUT_SHAPE
}
