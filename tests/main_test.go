package tests

import (
	"Soccer-Penalty-Kick-ML-Threading/soccer"
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	env := soccer.InitSoccer()
	action := []float64{0}
	reward, done, _ := env.Step(action)
	fmt.Printf("%f, %t\n", reward, done)

	// Test 1000000 Shots
	// env := soccer.InitSoccer()
	// action := []float64{0}
	// results := [4]int{0, 0, 0, 0}
	// for i := 0; i < 1000000; i++ {
	// 	reward, _, _ := env.Step(action)
	// 	if reward == 500 {
	// 		results[0]++
	// 	} else if reward == 50 {
	// 		results[1]++
	// 	} else if reward == 25 {
	// 		results[2]++
	// 	} else {
	// 		results[3]++
	// 	}
	// }
	// fmt.Printf("Goals  : %d\n", results[0])
	// fmt.Printf("Saves  : %d\n", results[1])
	// fmt.Printf("Dingers: %d\n", results[2])
	// fmt.Printf("Misses : %d\n", results[3])
	// fmt.Printf("Goals  : %f\n", float64(results[0])/1000000.0)
	// fmt.Printf("Saves  : %f\n", float64(results[1])/1000000.0)
	// fmt.Printf("Dingers: %f\n", float64(results[2])/1000000.0)
	// fmt.Printf("Misses : %f\n", float64(results[3])/1000000.0)
}
