package bench

import (
	"fmt"
	"testing"
	"time"

	"github.com/thadUra/Go-ML/nn"
)

func TestNN(t *testing.T) {
	// Start timer
	start := time.Now()

	// Create training data
	var x_train = [][]float64{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	var y_train = [][]float64{{0}, {1}, {1}, {0}}

	// Initialize neural network with two hidden layers and MSE loss
	var net nn.Network
	net.AddLayer("INPUT", "", 2)
	net.AddLayer("DENSE", "TANH", 16)
	net.AddLayer("DENSE", "SIGMOID", 1)
	net.SetLoss("MSE", []float64{})

	// Train the model and display results
	net.Fit(x_train, y_train, 1000, 0.001, false)

	// End timer
	duration := time.Since(start)
	fmt.Printf("NN: %v seconds\n", duration.Seconds())
}
