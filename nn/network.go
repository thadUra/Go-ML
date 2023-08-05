package nn

import (
	"fmt"
	"time"

	"gonum.org/v1/gonum/mat"
)

/**
 * Network Struct
 * Contains neural network specifics and helper functions
 * Defines layers and loss functions
 */
type Network struct {
	LAYERS      []*Layer
	LOSSFUNC    func(y_true, y_pred *mat.Dense, params []float64) float64
	LOSSDERIV   func(y_true, y_pred *mat.Dense, params []float64) *mat.Dense
	LOSSPARAMS  []float64
	OUTPUT_SIZE int
	LAYER_COUNT int
}

/**
 * AddLayer()
 * Adds specific layer to neural network given activation type and input/output neuron dimensions
 */
func (net *Network) AddLayer(layerType, activationType string, output_nodes int) {
	var layer Layer
	var activation Layer

	// Get layerType
	switch layerType {
	case "INPUT":
		net.OUTPUT_SIZE = output_nodes
		return
	case "DENSE":
		layer = Layer(InitDenseLayer(net.OUTPUT_SIZE, output_nodes))
	case "FLATTEN":
		if net.LAYER_COUNT == 0 {
			net.OUTPUT_SIZE = output_nodes
			return
		}
		flat := InitFlattenLayer(net.OUTPUT_SIZE, output_nodes)
		layer = Layer(flat)
		net.LAYERS = append(net.LAYERS, &layer)
		_, net.OUTPUT_SIZE = flat.GetShape()
		return
	case "CONVOLUTIONAL":
		layer = Layer(InitConvolutionalLayer(net.OUTPUT_SIZE, output_nodes))
	default:
		layer = Layer(InitDenseLayer(net.OUTPUT_SIZE, output_nodes))
	}
	net.OUTPUT_SIZE = output_nodes

	// Get activationType
	switch activationType {
	case "TANH":
		activation = Layer(InitActivationLayer(Tanh, TanhDeriv))
	case "SIGMOID":
		activation = Layer(InitActivationLayer(Sigmoid, SigmoidDeriv))
	case "RELU":
		activation = Layer(InitActivationLayer(ReLu, ReLuDeriv))
	case "ARCTAN":
		activation = Layer(InitActivationLayer(Arctan, ArctanDeriv))
	case "GAUSSIAN":
		activation = Layer(InitActivationLayer(Gaussian, GaussianDeriv))
	default:
		activation = Layer(InitActivationLayer(Linear, LinearDeriv))
	}

	// Add layers to network
	net.LAYERS = append(net.LAYERS, &layer)
	net.LAYERS = append(net.LAYERS, &activation)
}

/**
 * SetLoss()
 * Contains neural network specifics and helper functions
 * Defines layers and loss functions
 */
func (net *Network) SetLoss(lossType string, params []float64) {
	net.LOSSPARAMS = params
	switch lossType {
	case "MSE":
		net.LOSSFUNC = Mse
		net.LOSSDERIV = MseDerivative
	case "HMSE":
		net.LOSSFUNC = Hmse
		net.LOSSDERIV = HmseDerivative
	case "RMSE":
		net.LOSSFUNC = Rmse
		net.LOSSDERIV = RmseDerivative
	case "MAE":
		net.LOSSFUNC = Mae
		net.LOSSDERIV = MaeDerivative
	case "HUBER":
		net.LOSSFUNC = Huber
		net.LOSSDERIV = HuberDerivative
	default:
		net.LOSSFUNC = Mse
		net.LOSSDERIV = MseDerivative
	}
}

/**
 * Predict()
 * Given sample data, return predicted values ran through the neural network
 */
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

/**
 * Fit()
 * Fit the neural network given training data, number of epochs, and a learning rate
 * Debug option provided for printing data from each epoch
 */
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
			err += net.LOSSFUNC(reference, output, net.LOSSPARAMS)
			// Backwards propagation
			error := net.LOSSDERIV(reference, output, net.LOSSPARAMS)
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
	if debug {
		fmt.Printf("Training Time: %s\n", elapsed)
	}
}
