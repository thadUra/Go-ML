package nn

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// Application Function
func Apply(fn func(x float64) float64, input *mat.Dense) *mat.Dense {
	rows, cols := input.Dims()
	apply := make([]float64, rows*cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			apply[(cols*r)+c] = fn(input.At(r, c))
		}
	}
	return mat.NewDense(rows, cols, apply)
}

// Individual Activation Functions
func Tanh(input *mat.Dense) *mat.Dense {
	return Apply(math.Tanh, input)
}

func TanhPrime(input *mat.Dense) *mat.Dense {
	tanh_prime := func(x float64) float64 {
		return 1 - (math.Pow(math.Tanh(x), 2))
	}
	return Apply(tanh_prime, input)
}
