package main

import (
	"Soccer-Penalty-Kick-ML-Threading/ml"
	"fmt"
)

// var wg sync.WaitGroup

func main() {
	// Neural Network Package Tests
	// full := ml.InitFCLayer(2, 3)
	// fc := mat.Formatted(full.WEIGHTS, mat.Prefix("    "), mat.Squeeze())
	// fmt.Printf("c = %v\n", fc)
	// act := ml.InitActivationLayer(ml.Tanh, ml.TanhPrime)
	// c := act.ForwardPropagation(full.WEIGHTS)
	// fc = mat.Formatted(c, mat.Prefix("    "), mat.Squeeze())
	// fmt.Printf("\nf = %v\n", fc)
	// b := act.ACTIVATIONPRIME(full.WEIGHTS)
	// fc = mat.Formatted(b, mat.Prefix("    "), mat.Squeeze())
	// fmt.Printf("\nb = %v", fc)
	// mse := ml.Mse(full.WEIGHTS, c)
	// fmt.Printf("\nMSE: %f\n", mse)
	// deriv := ml.Mse_prime(full.WEIGHTS, c)
	// fc = mat.Formatted(deriv, mat.Prefix("    "), mat.Squeeze())
	// fmt.Printf("\nd = %v\n", fc)

	// Solve XOR with Baseline NN
	var x_train = [][]float64{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	var y_train = [][]float64{{0}, {1}, {1}, {0}}

	var net ml.Network
	net.AddLayer(ml.InitFCLayer(2, 3))
	net.AddLayer(ml.InitActivationLayer(ml.Tanh, ml.TanhPrime))
	net.AddLayer(ml.InitFCLayer(3, 1))
	net.AddLayer(ml.InitActivationLayer(ml.Tanh, ml.TanhPrime))
	net.SetLoss(ml.Mse, ml.Mse_prime)
	net.Fit(x_train, y_train, 1000, 0.1)
	out := net.Predict(x_train[0])
	fmt.Printf("Result for 0: %f\n", out[0])
	out = net.Predict(x_train[1])
	fmt.Printf("Result for 1: %f\n", out[0])
	out = net.Predict(x_train[2])
	fmt.Printf("Result for 2: %f\n", out[0])
	out = net.Predict(x_train[3])
	fmt.Printf("Result for 3: %f\n", out[0])

	// Initialize game env
	// var reward int
	// var goals int
	// params := game.InitSoccer(0, 0, 0, 0, 0, 0, 0, true)
	// env := game.InitEnvironment(params)
	// location := game.InitShot(0, 0, true)

	// Environment Space Tests

	// Manual Test
	// location := game.InitShot(145, 50, false)
	// action := []float64{-25.0 * math.Pi / 180.0, 10.0 * math.Pi / 180.0, 35.0}
	// fmt.Printf("Reward: %d\n\n", env.Step(action, location, true))

	// 1000 Random Episodes Timed No Threading
	// test := func(action []float64) {
	// 	result := env.Step(action, location, false)
	// 	reward += result
	// 	if result == 10 {
	// 		goals++
	// 	}
	// }
	// start := time.Now()
	// for episode := 1; episode <= 1000; episode++ {
	// 	action := []float64{
	// 		((rand.Float64() * 180) - 90) * math.Pi / 180.0,
	// 		(rand.Float64() * 90) * math.Pi / 180.0,
	// 		rand.Float64() * 145.0}
	// 	test(action)
	// }
	// elapsed := time.Since(start)
	// fmt.Printf("Single Threaded: %s with reward: %d and %d goals\n", elapsed, reward, goals)

	// 1000000 Random Episodes Timed Multithreading
	// reward = 0
	// goals = 0
	// testThread := func(action []float64) {
	// 	result := env.Step(action, location, false)
	// 	reward += result
	// 	if result == 10 {
	// 		goals++
	// 	}
	// 	defer wg.Done()
	// }
	// start = time.Now()
	// wg.Add(100000000)
	// for episode := 1; episode <= 100000000; episode++ {
	// 	action := []float64{
	// 		((rand.Float64() * 180) - 90) * math.Pi / 180.0,
	// 		(rand.Float64() * 90) * math.Pi / 180.0,
	// 		rand.Float64() * 145.0}
	// 	go testThread(action)
	// }
	// wg.Wait()
	// elapsed = time.Since(start)
	// fmt.Printf("Multithreading: %s with reward: %d and %d goals\n", elapsed, reward, goals)
}
