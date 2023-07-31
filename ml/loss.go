package ml

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func Mse(y_true, y_pred *mat.Dense) float64 {
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

func Mse_prime(y_true, y_pred *mat.Dense) *mat.Dense {
	ret := mat.DenseCopyOf(y_pred)
	ret.Sub(y_pred, y_true)
	rows, cols := y_true.Dims()
	mult := float64(2 / float64(rows*cols))
	ret.Scale(mult, ret)
	return ret
}
