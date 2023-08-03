package nn

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// https://medium.com/nerd-for-tech/what-loss-function-to-use-for-machine-learning-project-b5c5bd4a151e

/**
 * Mean Squared Error Loss Functions
 */
func Mse(y_true, y_pred *mat.Dense, params []float64) float64 {
	ret := mat.DenseCopyOf(y_true)
	ret.Sub(y_true, y_pred)
	rows, cols := ret.Dims()
	var sum float64
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			sum += math.Pow(ret.At(r, c), 2)
		}
	}
	mean := float64(sum / float64(rows*cols))
	return mean
}
func MseDerivative(y_true, y_pred *mat.Dense, params []float64) *mat.Dense {
	ret := mat.DenseCopyOf(y_pred)
	ret.Sub(y_pred, y_true)
	rows, cols := y_true.Dims()
	mult := float64(2 / float64(rows*cols))
	ret.Scale(mult, ret)
	return ret
}

/**
 * Half Mean Squared Error Loss Functions
 */
func Hmse(y_true, y_pred *mat.Dense, params []float64) float64 {
	return Mse(y_true, y_pred, params) / 2.0
}
func HmseDerivative(y_true, y_pred *mat.Dense, params []float64) *mat.Dense {
	ret := MseDerivative(y_true, y_pred, params)
	mult := 0.5
	ret.Scale(mult, ret)
	return ret
}

/**
 * Root Mean Squared Error Loss Functions
 */
func Rmse(y_true, y_pred *mat.Dense, params []float64) float64 {
	return math.Sqrt(Mse(y_true, y_pred, params))
}
func RmseDerivative(y_true, y_pred *mat.Dense, params []float64) *mat.Dense {
	ret := MseDerivative(y_true, y_pred, params)
	mult := 1.0 / (2.0 * Rmse(y_true, y_pred, params))
	ret.Scale(mult, ret)
	return ret
}

/**
 * Mean Absolute Error Loss Functions
 * Derivative defaults to 1 when y_pred = y_true due to not being differentiable
 */
func Mae(y_true, y_pred *mat.Dense, params []float64) float64 {
	ret := mat.DenseCopyOf(y_true)
	ret.Sub(y_true, y_pred)
	rows, cols := ret.Dims()
	var sum float64
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			sum += math.Abs(ret.At(r, c))
		}
	}
	mean := float64(sum / float64(rows*cols))
	return mean
}
func MaeDerivative(y_true, y_pred *mat.Dense, params []float64) *mat.Dense {
	ret := mat.DenseCopyOf(y_pred)
	rows, cols := y_true.Dims()
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if y_pred.At(r, c) >= y_true.At(r, c) {
				ret.Set(r, c, 1.0)
			} else {
				ret.Set(r, c, -1.0)
			}
		}
	}
	return ret
}

/**
 * Huber Loss Functions
 * Implements a pseudo-huber loss type to approximate
 * Takes in a delta value as a parameter
 */
func Huber(y_true, y_pred *mat.Dense, params []float64) float64 {
	ret := mat.DenseCopyOf(y_true)
	ret.Sub(y_true, y_pred)
	rows, cols := ret.Dims()
	var sum float64
	delta := params[0]
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			sum += math.Pow(delta, 2) * (math.Sqrt(1+math.Pow((ret.At(r, c)/delta), 2)) - 1)
		}
	}
	return sum
}
func HuberDerivative(y_true, y_pred *mat.Dense, params []float64) *mat.Dense {
	ret := mat.DenseCopyOf(y_true)
	ret.Sub(y_pred, y_true)
	rows, cols := ret.Dims()
	delta := params[0]
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			deriv_val := ret.At(r, c) / (math.Sqrt(1 + math.Pow(ret.At(r, c)/delta, 2)))
			ret.Set(r, c, deriv_val)
		}
	}
	return ret
}
