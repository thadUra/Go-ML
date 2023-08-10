# Golang-Machine-Learning nn

[![Documentation](https://img.shields.io/badge/documentation-GoDoc-blue.svg)](https://pkg.go.dev/github.com/thadUra/Golang-Machine-Learning/nn)

Package nn is a machine learning neural network package for Go.

## Example Usage

Below contains example usage of the nn package on XOR training data. This example can be found in `../tests/xor_test.go`.

### Neural Network for Predicting XOR Operations
```
    // Create training data
	var x_train = [][]float64{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	var y_train = [][]float64{{0}, {1}, {1}, {0}}

	// Initialize neural network with two hidden layers and MSE loss
	var net nn.Network
	net.AddLayer("INPUT", "", 2)
	net.AddLayer("DENSE", "TANH", 3)
	net.AddLayer("DENSE", "SIGMOID", 3)
	net.AddLayer("DENSE", "ARCTAN", 1)
	net.SetLoss("HUBER", []float64{1.35}) // delta = 1.35

	// Train the model and display results
	net.Fit(x_train, y_train, 1000, 0.1, false)
	result := net.Predict(x_train)
	for i := range result {
		fmt.Printf("index %d: got %f, want %f\n", i, result[i][0], y_train[i][0])
		if math.Round(result[i][0]) != y_train[i][0] {
			t.Fatalf(`net.Predict() gave "%f", want "%f"`, math.Round(result[i][0]), y_train[i][0])
		}
	}
```

#### Result
```
    epoch 1/1000  error=0.190287
    epoch 2/1000  error=0.178683
    epoch 3/1000  error=0.167760
    ...
    epoch 998/1000  error=0.000001
    epoch 999/1000  error=0.000001
    epoch 1000/1000  error=0.000001
    Training Time: 26.300815ms

    index 0: got 0.000250, want 0.000000
    index 1: got 0.997680, want 1.000000
    index 2: got 0.998551, want 1.000000
    index 3: got 0.000402, want 0.000000
```