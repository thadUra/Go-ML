package nn

import (
	"fmt"
	"time"

	"gonum.org/v1/gonum/mat"
)

type LossFunc func(y_true, y_pred *mat.Dense) float64
type LossPrimeFunc func(y_true, y_pred *mat.Dense) *mat.Dense

type Network struct {
	LAYERS    []*Layer
	LOSS      LossFunc
	LOSSPRIME LossPrimeFunc
}

func (net *Network) AddLayer(layer Layer) {
	net.LAYERS = append(net.LAYERS, &layer)
}

func (net *Network) SetLoss(loss LossFunc, lossPrime LossPrimeFunc) {
	net.LOSS = loss
	net.LOSSPRIME = lossPrime
}

// func (net *Network) SetPolicy() {}

func (net *Network) Predict(input []float64) []float64 {
	output := mat.NewDense(1, len(input), input)
	for i := range net.LAYERS {
		output = (*net.LAYERS[i]).ForwardPropagation(output)
	}
	rows, cols := output.Dims()
	result := make([]float64, rows*cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			result[(cols*r)+c] = output.At(r, c)
		}
	}
	return result
}

// func (net *Network) Fit(env game.Environment, epochs int, learning_rate float64) {

// }

func (net *Network) Fit(x_train, y_train [][]float64, epochs int, learning_rate float64) {

	start := time.Now()

	// sample dimensions
	samples := len(x_train)

	// training loop
	for i := 0; i < epochs; i++ {
		err := 0.0
		for j := range x_train {

			// forward propagation
			output := mat.NewDense(1, len(x_train[j]), x_train[j])
			for l := range net.LAYERS {
				// fmt.Printf("	Forward for layer %d:", l)
				output = (*net.LAYERS[l]).ForwardPropagation(output)
			}
			reference := mat.NewDense(1, len(y_train[j]), y_train[j])

			// compute loss
			err += net.LOSS(reference, output)
			// fmt.Printf("	Loss after epoch %d: %f\n", i+1, net.LOSS(reference, output))

			// backwards propagation
			error := net.LOSSPRIME(reference, output)
			for l := len(net.LAYERS) - 1; l >= 0; l-- {
				// fmt.Printf("	Backwards for layer %d:", l)
				error = (*net.LAYERS[l]).BackPropagation(error, learning_rate)
			}
		}
		err /= float64(samples)
		// if i < 3 || i >= epochs-3 {
		// 	fmt.Printf("epoch %d/%d  error=%f\n", i+1, epochs, err)
		// } else if i >= 3 && i < epochs-3 && i == 4 {
		// 	fmt.Println("...")
		// }
	}
	elapsed := time.Since(start)
	fmt.Printf("Training Time: %s\n", elapsed)
}
