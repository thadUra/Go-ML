package main

import (
	"Soccer-Penalty-Kick-ML-Threading/game"
	"fmt"
	"math"
	"sync"
)

var wg sync.WaitGroup

// var count int
var reward int
var goals int

// func addFull(start int, end int, thread bool) {
// 	// var count int
// 	for i := start; i < end; i++ {
// 		count += 1
// 	}
// 	if thread {
// 		defer wg.Done()
// 	}
// }

func main() {

	// Count Test No Threading
	// start := time.Now()
	// addFull(0, 1000000000, false)
	// elapsed := time.Since(start)
	// fmt.Printf("No Threading:%s\n", elapsed)

	// Count Test Multithreading
	// start = time.Now()
	// wg.Add(2)
	// go addFull(0, 500000000, true)
	// go addFull(500000000, 1000000000, true)
	// wg.Wait()
	// elapsed = time.Since(start)
	// fmt.Printf("Multithreading:%s\n", elapsed)

	// Initialize game env
	params := game.InitSoccer(0, 0, 0, 0, 0, 0, 0, true)
	env := game.InitEnvironment(params)
	// location := game.InitShot(0, 0, true)

	// Environment Space Tests

	// Manual Test
	location := game.InitShot(145, 50, false)
	action := []float64{-25.0 * math.Pi / 180.0, 50.0 * math.Pi / 180.0, 35.0}
	fmt.Printf("Reward: %d\n\n", env.Step(action, location, true))

	// 1000000 Random Episodes Timed No Threading
	// test := func(action []float64) {
	// 	result := env.Step(action, location, false)
	// 	reward += result
	// 	if result == 10 {
	// 		goals++
	// 	}
	// }
	// start := time.Now()
	// for episode := 1; episode <= 100000000; episode++ {
	// 	action := []float64{
	// 		((rand.Float64() * 180) - 90) * math.Pi / 180.0,
	// 		(rand.Float64() * 90) * math.Pi / 180.0,
	// 		rand.Float64() * 145.0}
	// 	test(action)
	// }
	// elapsed := time.Since(start)
	// fmt.Printf("Multithreading: %s with reward: %d and %d goals\n", elapsed, reward, goals)

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
