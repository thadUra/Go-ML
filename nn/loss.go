package nn

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// LossFwd defines a type for any loss function for forward propagation.
type LossFwd func(y_true, y_pred *mat.Dense, params []float64) float64

// LossBwd defines a type for any loss function for back propagation.
type LossBwd func(y_true, y_pred *mat.Dense, params []float64) *mat.Dense

// Loss defines the essential loss functions for the neural network.
type Loss struct {
	LOSSFUNC   LossFwd
	LOSSDERIV  LossBwd
	LOSSPARAMS []float64
}

// Mse returns the mean squared error given `y_true` and `y_pred` matrices.
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

// MseDerivative returns the derivative matrix of the mean squared error given `y_true` and `y_pred` matrices.
func MseDerivative(y_true, y_pred *mat.Dense, params []float64) *mat.Dense {
	ret := mat.DenseCopyOf(y_pred)
	ret.Sub(y_pred, y_true)
	rows, cols := y_true.Dims()
	mult := float64(2 / float64(rows*cols))
	ret.Scale(mult, ret)
	return ret
}

// Hmse returns the half mean squared error given `y_true` and `y_pred` matrices.
func Hmse(y_true, y_pred *mat.Dense, params []float64) float64 {
	return Mse(y_true, y_pred, params) / 2.0
}

// HmseDerivative returns the derivative matrix of the half mean squared error given `y_true` and `y_pred` matrices.
func HmseDerivative(y_true, y_pred *mat.Dense, params []float64) *mat.Dense {
	ret := MseDerivative(y_true, y_pred, params)
	mult := 0.5
	ret.Scale(mult, ret)
	return ret
}

// Rmse returns the root mean squared error given `y_true` and `y_pred` matrices.
func Rmse(y_true, y_pred *mat.Dense, params []float64) float64 {
	return math.Sqrt(Mse(y_true, y_pred, params))
}

// RmseDerivative returns the derivative matrix of the root mean squared error given `y_true` and `y_pred` matrices.
func RmseDerivative(y_true, y_pred *mat.Dense, params []float64) *mat.Dense {
	ret := MseDerivative(y_true, y_pred, params)
	mult := 1.0 / (2.0 * Rmse(y_true, y_pred, params))
	ret.Scale(mult, ret)
	return ret
}

// Mae returns the mean absolute error given `y_true` and `y_pred` matrices.
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

// MaeDerivative returns the derivative matrix of the mean absolute error given `y_true` and `y_pred` matrices.
// The derivative defaults to 1 when y_pred = y_true due to not being differentiable.
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

// Huber returns the pseudo-huber loss given `y_true` and `y_pred` matrices. Params[0]
// represents the provided delta value.
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

// Huber returns the derivative matrix of the pseudo-huber loss given `y_true` and `y_pred`
// matrices. Params[0] represents the provided delta value.
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
