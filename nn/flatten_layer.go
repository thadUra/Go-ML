package nn

import (
	"gonum.org/v1/gonum/mat"
)

/**
 * FlattenLayer Struct
 *
 */
type FlattenLayer struct {
	INPUT  *mat.Dense
	OUTPUT *mat.Dense
}

/**
 * InitFlattenLayer()
 *
 */
func InitFlattenLayer(input_size int, output_size int) *FlattenLayer {
	// Initialize

	// Return layer
	var layer FlattenLayer
	return &layer
}

/**
 * ForwardPropagation()
 * Performs forward propagation for a flat layer required by Layer interface
 *
 */
func (layer *FlattenLayer) ForwardPropagation(input *mat.Dense) *mat.Dense {

	return layer.OUTPUT
}

/**
 * BackPropagation()
 * Performs back propagation for a flat layer required by Layer interface
 *
 */
func (layer *FlattenLayer) BackPropagation(output_error *mat.Dense, learning_rate float64) *mat.Dense {

	return layer.INPUT
}
