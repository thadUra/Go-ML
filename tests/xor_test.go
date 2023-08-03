package tests

import (
	"Soccer-Penalty-Kick-ML-Threading/nn"
	"fmt"
	"math"
	"testing"
)

/**
 * TestXorNeuralNetwork()
 * Creates a simple neural network to train for XOR operations
 * Tests the nn package on a baseline level
 */
func TestXorNeuralNetwork(t *testing.T) {

	// Create training data
	var x_train = [][]float64{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	var y_train = [][]float64{{0}, {1}, {1}, {0}}

	// Initialize neural network with one hidden layer and MSE loss
	var net nn.Network
	net.AddLayer("DENSE", "TANH", 2, 3)
	net.AddLayer("DENSE", "TANH", 3, 1)
	net.SetLoss("HUBER", []float64{1.35})

	// Train the model and display results
	net.Fit(x_train, y_train, 1000, 0.1)
	result := net.Predict(x_train)
	for i := range result {
		fmt.Printf("%d: %f, %f\n", i, result[i][0], y_train[i][0])
		if math.Round(result[i][0]) != y_train[i][0] {
			t.Fatalf(`net.Predict() gave "%f", want "%f"`, math.Round(result[i][0]), y_train[i][0])
		}
	}
}
