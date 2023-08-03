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
	LAYERS     []*Layer
	LOSSFUNC   func(y_true, y_pred *mat.Dense, params []float64) float64
	LOSSDERIV  func(y_true, y_pred *mat.Dense, params []float64) *mat.Dense
	LOSSPARAMS []float64
}

/**
 * AddLayer() WIP
 * Adds specific layer to neural network given activation type and input/output neuron dimensions
 */
func (net *Network) AddLayer(layerType, activationType string, input_nodes, output_nodes int) {
	var layer Layer
	var activation Layer

	// Get layerType
	switch layerType {
	case "DENSE":
		layer = Layer(InitFCLayer(input_nodes, output_nodes))
	case "FLATTEN":
		layer = Layer(InitFCLayer(input_nodes, output_nodes)) // CHANGE ONCE IMPLEMENTED
	case "CONVOLUTIONAL":
		layer = Layer(InitFCLayer(input_nodes, output_nodes)) // CHANGE ONCE IMPLEMENTED
	default:
		layer = Layer(InitFCLayer(input_nodes, output_nodes))
	}

	// Get activationType
	switch activationType {
	case "TANH":
		activation = Layer(InitActivationLayer(Tanh, TanhPrime))
	case "SIGMOID":
		activation = Layer(InitActivationLayer(Tanh, TanhPrime)) // CHANGE ONCE IMPLEMENTED
	case "RELU":
		activation = Layer(InitActivationLayer(Tanh, TanhPrime)) // CHANGE ONCE IMPLEMENTED
	case "ARCTAN":
		activation = Layer(InitActivationLayer(Tanh, TanhPrime)) // CHANGE ONCE IMPLEMENTED
	case "SOFTPLUS":
		activation = Layer(InitActivationLayer(Tanh, TanhPrime)) // CHANGE ONCE IMPLEMENTED
	case "GAUSSIAN":
		activation = Layer(InitActivationLayer(Tanh, TanhPrime)) // CHANGE ONCE IMPLEMENTED
	default:
		activation = Layer(InitActivationLayer(Tanh, TanhPrime)) // CHANGE ONCE IMPLEMENTED (linear)
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
 *
 */
func (net *Network) Fit(x_train, y_train [][]float64, epochs int, learning_rate float64) {

	// record training time
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
				output = (*net.LAYERS[l]).ForwardPropagation(output)
			}
			reference := mat.NewDense(1, len(y_train[j]), y_train[j])

			// compute loss
			err += net.LOSSFUNC(reference, output, net.LOSSPARAMS)

			// backwards propagation
			error := net.LOSSDERIV(reference, output, net.LOSSPARAMS)
			for l := len(net.LAYERS) - 1; l >= 0; l-- {
				error = (*net.LAYERS[l]).BackPropagation(error, learning_rate)
			}
		}
		err /= float64(samples)
		if i < 3 || i >= epochs-3 {
			fmt.Printf("epoch %d/%d  error=%f\n", i+1, epochs, err)
		} else if i >= 3 && i < epochs-3 && i == 4 {
			fmt.Println("...")
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Training Time: %s\n", elapsed)
}
