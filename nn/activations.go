package nn

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

/**
 * Apply()
 * Helper that takes a function as a parameter to perform on the layer input
 */
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

/**
 * Tanh Activation Funcs
 */
func Tanh(input *mat.Dense) *mat.Dense {
	return Apply(math.Tanh, input)
}
func TanhDeriv(input *mat.Dense) *mat.Dense {
	tanh_deriv := func(x float64) float64 {
		return 1 - (math.Pow(math.Tanh(x), 2))
	}
	return Apply(tanh_deriv, input)
}

/**
 * Sigmoid Activation Funcs
 */
func Sigmoid(input *mat.Dense) *mat.Dense {
	sigmoid := func(x float64) float64 {
		return 1 / (1 + math.Pow(math.E, -x))
	}
	return Apply(sigmoid, input)
}
func SigmoidDeriv(input *mat.Dense) *mat.Dense {
	sigmoid_deriv := func(x float64) float64 {
		return math.Pow(math.E, x) / math.Pow(math.Pow(math.E, x)+1, 2)
	}
	return Apply(sigmoid_deriv, input)
}

/**
 * ReLu Activation Funcs
 * Derivative defaults to 0 if x = 0
 */
func ReLu(input *mat.Dense) *mat.Dense {
	relu := func(x float64) float64 {
		return math.Max(0, x)
	}
	return Apply(relu, input)
}
func ReLuDeriv(input *mat.Dense) *mat.Dense {
	relu_deriv := func(x float64) float64 {
		if x <= 0 {
			return 0
		} else {
			return 1
		}
	}
	return Apply(relu_deriv, input)
}

/**
 * Arctan Activation Funcs
 */
func Arctan(input *mat.Dense) *mat.Dense {
	return Apply(math.Atan, input)
}
func ArctanDeriv(input *mat.Dense) *mat.Dense {
	arctan_deriv := func(x float64) float64 {
		return 1 / (math.Pow(x, 2) + 1)
	}
	return Apply(arctan_deriv, input)
}

/**
 * Gaussian Activation Funcs
 */
func Gaussian(input *mat.Dense) *mat.Dense {
	gaussian := func(x float64) float64 {
		return math.Pow(math.E, -1*x*x)
	}
	return Apply(gaussian, input)
}
func GaussianDeriv(input *mat.Dense) *mat.Dense {
	gaussian_deriv := func(x float64) float64 {
		return -2 * x * math.Pow(math.E, -1*x*x)
	}
	return Apply(gaussian_deriv, input)
}

/**
 * Linear Activation Funcs
 */
func Linear(input *mat.Dense) *mat.Dense {
	linear := func(x float64) float64 {
		return x
	}
	return Apply(linear, input)
}

func LinearDeriv(input *mat.Dense) *mat.Dense {
	linear_deriv := func(x float64) float64 {
		return 1
	}
	return Apply(linear_deriv, input)
}
