package testing

import (
	"Soccer-Penalty-Kick-ML-Threading/ml"
	"fmt"
)

/**
 * RunXorTest()
 * Creates a simple neural network to train for XOR operations
 * Tests the ML package on a baseline level
 */
func RunXorTest() {

	fmt.Printf("=====RUNNING XOR TEST=====\n")

	// Create training data
	var x_train = [][]float64{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	var y_train = [][]float64{{0}, {1}, {1}, {0}}

	// Initialize neural network with one hidden layer and MSE loss
	var net ml.Network
	net.AddLayer(ml.InitFCLayer(2, 3))
	net.AddLayer(ml.InitActivationLayer(ml.Tanh, ml.TanhPrime))
	net.AddLayer(ml.InitFCLayer(3, 1))
	net.AddLayer(ml.InitActivationLayer(ml.Tanh, ml.TanhPrime))
	net.SetLoss(ml.Mse, ml.Mse_prime)

	// Train the model and display results
	net.Fit(x_train, y_train, 1000, 0.1)
	for i := range x_train {
		out := net.Predict(x_train[i])
		fmt.Printf("Result: %f -> Expected: %f\n", out[0], y_train[i])
	}

	fmt.Printf("=====ENDING XOR TEST=====\n")

}
