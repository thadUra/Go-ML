package nn

import (
	"fmt"
	"time"

	"gonum.org/v1/gonum/mat"
)

// Network represents the neural network model specifics and related helper functions for loss.
type Network struct {
	LAYERS      []*Layer
	LOSS        Loss
	OUTPUT_SIZE int
	LAYER_COUNT int
}

// AddLayer appends to `LAYERS` in the network given `layerType`, `activationType`, and output neurons.
func (net *Network) AddLayer(layerType, activationType string, output_nodes int) {
	var layer Layer
	var activation Layer

	// Get layerType
	switch layerType {
	case "INPUT":
		net.OUTPUT_SIZE = output_nodes
		return
	case "DENSE":
		layer = Layer(NewDenseLayer(net.OUTPUT_SIZE, output_nodes))
	case "FLATTEN":
		if net.LAYER_COUNT == 0 {
			net.OUTPUT_SIZE = output_nodes
			return
		}
		flat := NewFlattenLayer(net.OUTPUT_SIZE, output_nodes)
		layer = Layer(flat)
		net.LAYERS = append(net.LAYERS, &layer)
		_, net.OUTPUT_SIZE = flat.GetShape()
		return
	case "CONVOLUTIONAL":
		layer = Layer(NewConvolutionalLayer(net.OUTPUT_SIZE, output_nodes))
	default:
		layer = Layer(NewDenseLayer(net.OUTPUT_SIZE, output_nodes))
	}
	net.OUTPUT_SIZE = output_nodes

	// Get activationType
	switch activationType {
	case "TANH":
		activation = Layer(NewActivationLayer(Tanh, TanhDeriv))
	case "SIGMOID":
		activation = Layer(NewActivationLayer(Sigmoid, SigmoidDeriv))
	case "RELU":
		activation = Layer(NewActivationLayer(ReLu, ReLuDeriv))
	case "ARCTAN":
		activation = Layer(NewActivationLayer(Arctan, ArctanDeriv))
	case "GAUSSIAN":
		activation = Layer(NewActivationLayer(Gaussian, GaussianDeriv))
	default:
		activation = Layer(NewActivationLayer(Linear, LinearDeriv))
	}

	// Add layers to network
	net.LAYERS = append(net.LAYERS, &layer)
	net.LAYERS = append(net.LAYERS, &activation)
}

// SetLoss initializes the loss function for the network given `lossType`. If the string
// is empty, it defaults to MSE loss.
func (net *Network) SetLoss(lossType string, params []float64) {
	net.LOSS.LOSSPARAMS = params
	switch lossType {
	case "MSE":
		net.LOSS.LOSSFUNC = Mse
		net.LOSS.LOSSDERIV = MseDerivative
	case "HMSE":
		net.LOSS.LOSSFUNC = Hmse
		net.LOSS.LOSSDERIV = HmseDerivative
	case "RMSE":
		net.LOSS.LOSSFUNC = Rmse
		net.LOSS.LOSSDERIV = RmseDerivative
	case "MAE":
		net.LOSS.LOSSFUNC = Mae
		net.LOSS.LOSSDERIV = MaeDerivative
	case "HUBER":
		net.LOSS.LOSSFUNC = Huber
		net.LOSS.LOSSDERIV = HuberDerivative
	default:
		net.LOSS.LOSSFUNC = Mse
		net.LOSS.LOSSDERIV = MseDerivative
	}
}

// Predict returns predicted values given `input` through running the values through the neural network.
func (net *Network) Predict(input [][]float64) [][]float64 {
	// Init result and sample size
	var result [][]float64
	size := len(input)

	// Run model on each input sample
	for i := 0; i < size; i++ {
		// Propagate forward
		output := mat.NewDense(1, len(input[i]), input[i])
		for j := range net.LAYERS {
			output = (*net.LAYERS[j]).ForwardPropagation(output)
		}
		// Append result for iteration
		rows, cols := output.Dims()
		temp := make([]float64, rows*cols)
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				temp[(cols*r)+c] = output.At(r, c)
			}
		}
		result = append(result, temp)
	}
	return result
}

// Fit trains the neural network provided training data, the number of epochs, and learning rate.
// If `debug` is true, then the error at each epoch is printed in terminal.
func (net *Network) Fit(x_train, y_train [][]float64, epochs int, learning_rate float64, debug bool) {
	// Record training duration
	start := time.Now()

	// Training loop
	for i := 0; i < epochs; i++ {
		err := 0.0
		for j := range x_train {
			// Forward propagation
			output := mat.NewDense(1, len(x_train[j]), x_train[j])
			for l := range net.LAYERS {
				output = (*net.LAYERS[l]).ForwardPropagation(output)
			}
			// Convert output to matrix for loss computation
			reference := mat.NewDense(1, len(y_train[j]), y_train[j])
			// Add error
			err += net.LOSS.LOSSFUNC(reference, output, net.LOSS.LOSSPARAMS)
			// Backwards propagation
			error := net.LOSS.LOSSDERIV(reference, output, net.LOSS.LOSSPARAMS)
			for l := len(net.LAYERS) - 1; l >= 0; l-- {
				error = (*net.LAYERS[l]).BackPropagation(error, learning_rate)
			}
		}
		// Adjust error
		err /= float64(len(x_train))

		// Debug statements (print first and last three epochs if not true)
		if debug {
			fmt.Printf("epoch %d/%d  error=%f\n", i+1, epochs, err)
		} else {
			if i < 3 || i >= epochs-3 {
				fmt.Printf("epoch %d/%d  error=%f\n", i+1, epochs, err)
			} else if i >= 3 && i < epochs-3 && i == 4 {
				fmt.Println("...")
			}
		}
	}

	// Print training time
	elapsed := time.Since(start)
	fmt.Printf("Training Time: %s\n\n", elapsed)
}
