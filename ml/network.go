package ml

import "gonum.org/v1/gonum/mat"

type LossFunc func(y_true, y_pred *mat.Dense) float64
type LossPrimeFunc func(y_true, y_pred *mat.Dense) *mat.Dense

type Network struct {
	LAYERS    []Layer
	LOSS      LossFunc
	LOSSPRIME LossPrimeFunc
}

func (net *Network) AddLayer(layer Layer) {
	net.LAYERS = append(net.LAYERS, layer)
}

func (net *Network) SetLoss(loss LossFunc, lossPrime LossPrimeFunc) {
	net.LOSS = loss
	net.LOSSPRIME = lossPrime
}
