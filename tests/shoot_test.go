package tests

import (
	"testing"
)

/**
 * TestShootingNNModel()
 * Attempting to train a machine learning model that can score goals
 * Observation is the location on field, actions are the shot param values
 */
func TestShootingNNModel(t *testing.T) {

	// // Generate soccer env
	// env := soccer.InitSoccer()

	// // Generate training data
	// var horizontal_x, horizontal_y, vertical_x, vertical_y, power_x, power_y [][]float64
	// location_max := 100
	// goal_max := 1000
	// shot_max := 100000
	// param_range := env.GetField().GetShotParameterLimits()

	// for i := 0; i < location_max; i++ {
	// 	goals := 0
	// 	for j := 0; j < shot_max && goals < goal_max; j++ {
	// 		action := []float64{
	// 			((rand.Float64() * 160) - 80) * math.Pi / 180.0,
	// 			(rand.Float64() * 90) * math.Pi / 180.0,
	// 			rand.Float64() * 100.0}
	// 		result, err := env.GetField().Shoot(env.GetPos(), action, false)
	// 		if err != nil {
	// 			t.Fatalf(`env.Shotot() gave error:"%s"`, err)
	// 		}
	// 		if result == "GOAL" {
	// 			goals++

	// 			// Set horizontal
	// 			// horizontal_data := make([]float64, int(param_range[0][1]-param_range[0][0]))
	// 			// for i := range horizontal_data {
	// 			// 	if int(action[0]) == i {
	// 			// 		horizontal_data[i] = 10
	// 			// 	} else {
	// 			// 		horizontal_data[i] = 0
	// 			// 	}
	// 			// }
	// 			horizontal_y = append(horizontal_y, []float64{action[0]})
	// 			horizontal_x = append(horizontal_x, []float64{env.GetPos().DISTANCE_X, env.GetPos().DISTANCE_Y})

	// 			// Set vertical
	// 			vertical_data := make([]float64, int(param_range[1][1]-param_range[1][0]))
	// 			for i := range vertical_data {
	// 				if int(action[1]) == i {
	// 					vertical_data[i] = 1
	// 				} else {
	// 					vertical_data[i] = 0
	// 				}
	// 			}
	// 			vertical_y = append(vertical_y, vertical_data)
	// 			vertical_x = append(vertical_x, []float64{env.GetPos().DISTANCE_X, env.GetPos().DISTANCE_Y})

	// 			// Set power
	// 			power_data := make([]float64, int(param_range[2][1]-param_range[2][0]))
	// 			for i := range power_data {
	// 				if int(action[2]) == i {
	// 					power_data[i] = 1
	// 				} else {
	// 					power_data[i] = 0
	// 				}
	// 			}
	// 			power_y = append(power_y, power_data)
	// 			power_x = append(power_x, []float64{env.GetPos().DISTANCE_X, env.GetPos().DISTANCE_Y})
	// 		}
	// 	}
	// 	env.Reset()
	// }

	// // Initalize model for horizontal
	// input_size := len(horizontal_x[0])
	// output_size := len(horizontal_y[0])
	// var net nn.Network
	// net.AddLayer("INPUT", "", input_size)
	// net.AddLayer("DENSE", "TANH", 10)
	// net.AddLayer("DENSE", "RELU", 10)
	// net.AddLayer("DENSE", "TANH", output_size)
	// net.SetLoss("MSE", []float64{1.35})

	// // Train model
	// net.Fit(horizontal_x, horizontal_y, 100, 0.1, true)

	// // Test model
	// env = soccer.InitSoccer() // penalty
	// test := [][]float64{{env.GetPos().DISTANCE_X, env.GetPos().DISTANCE_Y}}
	// result := net.Predict(test)
	// for i := range result {
	// 	fmt.Printf("%d at (%f,%f): %f\n", i, test[0][0], test[0][1], result[i][0])
	// 	action := []float64{result[i][0], 5.0, 75}
	// 	env.GetField().Shoot(env.GetPos(), action, true)
	// 	// for j := int(param_range[0][0]); j < int(param_range[0][1]); j++ {
	// 	// 	fmt.Printf("%d: %f\n", j, result[i][j+80])
	// 	// }
	// }
}
