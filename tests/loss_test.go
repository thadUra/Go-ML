package tests

import (
	"Soccer-Penalty-Kick-ML-Threading/nn"
	"fmt"
	"math"
	"testing"

	"gonum.org/v1/gonum/mat"
)

/**
 * TestLossFunctions()
 * Creates various tests for the different loss functions defined in loss.go
 */
func TestLossFunctions(t *testing.T) {

	// Utilize random data set
	length := 5
	observed := mat.NewDense(1, length, []float64{34, 37, 44, 47, 48})
	predicted := mat.NewDense(1, length, []float64{37, 40, 46, 44, 46})

	// Test MSE
	err := nn.Mse(observed, predicted, []float64{-1})
	deriv := nn.MseDerivative(observed, predicted, []float64{-1})

	fmt.Printf("MSE: %f\n", err)
	format_deriv := mat.Formatted(deriv, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("MSE_DERIV = %v\n", format_deriv)

	actual_err := 7.0
	if int(err*10000.0) != int(actual_err*10000.0) {
		t.Fatalf(`nn.Mse() gave "%f", want "%f"`, err, actual_err)
	}
	actual_deriv := mat.NewDense(1, length, []float64{3.0 * 2.0 / 5.0,
		3.0 * 2.0 / 5.0,
		2.0 * 2.0 / 5.0,
		-3.0 * 2.0 / 5.0,
		-2.0 * 2.0 / 5.0})
	for i := 0; i < length; i++ {
		if int(deriv.At(0, i)*10000.0) != int(actual_deriv.At(0, i)*10000.0) {
			t.Fatalf(`nn.MseDerivative() at (0,%d) gave "%f", want "%f"`, i, deriv.At(0, i), actual_deriv.At(0, i))
		}
	}
	fmt.Println("")

	// Test HMSE
	err = nn.Hmse(observed, predicted, []float64{-1})
	deriv = nn.HmseDerivative(observed, predicted, []float64{-1})

	fmt.Printf("HMSE: %f\n", err)
	format_deriv = mat.Formatted(deriv, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("HMSE_DERIV = %v\n", format_deriv)

	actual_err = 3.5
	if int(err*10000.0) != int(actual_err*10000.0) {
		t.Fatalf(`nn.Hmse() gave "%f", want "%f"`, err, actual_err)
	}
	actual_deriv = mat.NewDense(1, length, []float64{3.0 / 5.0,
		3.0 / 5.0,
		2.0 / 5.0,
		-3.0 / 5.0,
		-2.0 / 5.0})
	for i := 0; i < length; i++ {
		if int(deriv.At(0, i)*10000.0) != int(actual_deriv.At(0, i)*10000.0) {
			t.Fatalf(`nn.HmseDerivative() at (0,%d) gave "%f", want "%f"`, i, deriv.At(0, i), actual_deriv.At(0, i))
		}
	}
	fmt.Println("")

	// Test RMSE
	err = nn.Rmse(observed, predicted, []float64{-1})
	deriv = nn.RmseDerivative(observed, predicted, []float64{-1})

	fmt.Printf("RMSE: %f\n", err)
	format_deriv = mat.Formatted(deriv, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("RMSE_DERIV = %v\n", format_deriv)

	actual_err = 2.64575
	if int(err*10000.0) != int(actual_err*10000.0) {
		t.Fatalf(`nn.Rmse() gave "%f", want "%f"`, err, actual_err)
	}
	actual_deriv = mat.NewDense(1, length, []float64{3.0 * 2.0 / 5.0 / 2.0 / math.Sqrt(7.0),
		3.0 * 2.0 / 5.0 / 2.0 / math.Sqrt(7.0),
		2.0 * 2.0 / 5.0 / 2.0 / math.Sqrt(7.0),
		-3.0 * 2.0 / 5.0 / 2.0 / math.Sqrt(7.0),
		-2.0 * 2.0 / 5.0 / 2.0 / math.Sqrt(7.0)})
	for i := 0; i < length; i++ {
		if int(deriv.At(0, i)*10000.0) != int(actual_deriv.At(0, i)*10000.0) {
			t.Fatalf(`nn.RmseDerivative() at (0,%d) gave "%f", want "%f"`, i, deriv.At(0, i), actual_deriv.At(0, i))
		}
	}
	fmt.Println("")

	// Test MAE
	err = nn.Mae(observed, predicted, []float64{-1})
	deriv = nn.MaeDerivative(observed, predicted, []float64{-1})

	fmt.Printf("MAE: %f\n", err)
	format_deriv = mat.Formatted(deriv, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("MAE_DERIV = %v\n", format_deriv)

	actual_err = 2.6
	if int(err*10000.0) != int(actual_err*10000.0) {
		t.Fatalf(`nn.Mae() gave "%f", want "%f"`, err, actual_err)
	}
	actual_deriv = mat.NewDense(1, length, []float64{1.0, 1.0, 1.0, -1.0, -1.0})
	for i := 0; i < length; i++ {
		if int(deriv.At(0, i)*10000.0) != int(actual_deriv.At(0, i)*10000.0) {
			t.Fatalf(`nn.MaeDerivative() at (0,%d) gave "%f", want "%f"`, i, deriv.At(0, i), actual_deriv.At(0, i))
		}
	}
	fmt.Println("")

	// Test Huber
	err = nn.Huber(observed, predicted, []float64{1.35})
	deriv = nn.HuberDerivative(observed, predicted, []float64{1.35})

	fmt.Printf("HUBER: %f\n", err)
	format_deriv = mat.Formatted(deriv, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("HUBER_DERIV = %v\n", format_deriv)

	actual_err = (2.619 + 2.619 + 1.435 + 2.619 + 1.435)
	if int(err*10000.0)/10000 != int(actual_err*10000.0)/10000 {
		t.Fatalf(`nn.Huber() gave "%f", want "%f"`, err, actual_err)
	}
	actual_deriv = mat.NewDense(1, length, []float64{1.231, 1.231, 1.119, -1.231, -1.119})
	for i := 0; i < length; i++ {
		if int(deriv.At(0, i)*10000.0)/10000 != int(actual_deriv.At(0, i)*10000.0)/10000 {
			t.Fatalf(`nn.HuberDerivative() at (0,%d) gave "%f", want "%f"`, i, deriv.At(0, i), actual_deriv.At(0, i))
		}
	}
	fmt.Println("")
}
