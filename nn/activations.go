package nn

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// apply performs the function on a matrix and returns the new matrix.
func apply(fn func(x float64) float64, input *mat.Dense) *mat.Dense {
	rows, cols := input.Dims()
	apply := make([]float64, rows*cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			apply[(cols*r)+c] = fn(input.At(r, c))
		}
	}
	return mat.NewDense(rows, cols, apply)
}

// Tanh Activation function
func Tanh(input *mat.Dense) *mat.Dense {
	return apply(math.Tanh, input)
}

// Tanh Derivative function
func TanhDeriv(input *mat.Dense) *mat.Dense {
	tanh_deriv := func(x float64) float64 {
		return 1 - (math.Pow(math.Tanh(x), 2))
	}
	return apply(tanh_deriv, input)
}

// Sigmoid Activation function
func Sigmoid(input *mat.Dense) *mat.Dense {
	sigmoid := func(x float64) float64 {
		return 1 / (1 + math.Pow(math.E, -x))
	}
	return apply(sigmoid, input)
}

// Sigmoid Derivative function
func SigmoidDeriv(input *mat.Dense) *mat.Dense {
	sigmoid_deriv := func(x float64) float64 {
		return math.Pow(math.E, x) / math.Pow(math.Pow(math.E, x)+1, 2)
	}
	return apply(sigmoid_deriv, input)
}

// ReLu Activation function: it defaults to 0 if x = 0
func ReLu(input *mat.Dense) *mat.Dense {
	relu := func(x float64) float64 {
		return math.Max(0, x)
	}
	return apply(relu, input)
}

// ReLu Derivative function
func ReLuDeriv(input *mat.Dense) *mat.Dense {
	relu_deriv := func(x float64) float64 {
		if x <= 0 {
			return 0
		} else {
			return 1
		}
	}
	return apply(relu_deriv, input)
}

// Arctan Activation function
func Arctan(input *mat.Dense) *mat.Dense {
	return apply(math.Atan, input)
}

// Arctan Derivative function
func ArctanDeriv(input *mat.Dense) *mat.Dense {
	arctan_deriv := func(x float64) float64 {
		return 1 / (math.Pow(x, 2) + 1)
	}
	return apply(arctan_deriv, input)
}

// Gaussian Activation function
func Gaussian(input *mat.Dense) *mat.Dense {
	gaussian := func(x float64) float64 {
		return math.Pow(math.E, -1*x*x)
	}
	return apply(gaussian, input)
}

// Gaussian Derivative function
func GaussianDeriv(input *mat.Dense) *mat.Dense {
	gaussian_deriv := func(x float64) float64 {
		return -2 * x * math.Pow(math.E, -1*x*x)
	}
	return apply(gaussian_deriv, input)
}

// Linear Activation function
func Linear(input *mat.Dense) *mat.Dense {
	linear := func(x float64) float64 {
		return x
	}
	return apply(linear, input)
}

// Linear Derivative function
func LinearDeriv(input *mat.Dense) *mat.Dense {
	linear_deriv := func(x float64) float64 {
		return 1
	}
	return apply(linear_deriv, input)
}
