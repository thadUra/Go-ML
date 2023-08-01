package testing

import (
	"Soccer-Penalty-Kick-ML-Threading/game"
	"fmt"
	"math"
	"math/rand"
	"time"
)

/**
 * RunEnvSpaceTest()
 * Creates a simple environment to perform steps on soccer game
 * Tests the game package with random actions
 */
func RunEnvSpaceTest() {

	fmt.Printf("=====RUNNING ENV SPACE TEST=====\n")

	// Initialize Game Env
	var reward int
	var goals int
	params := game.InitSoccer(0, 0, 0, 0, 0, 0, 0, true)
	env := game.InitEnvironment(params)

	// Chosen Manual Test
	location := game.InitShot(145, 50, false)
	action := []float64{-25.0 * math.Pi / 180.0, 10.0 * math.Pi / 180.0, 35.0}
	fmt.Printf("Reward: %d\n\n", env.Step(action, location, true))

	// 1000 Random Steps Timed Single Threaded for Penalty Shot
	location = game.InitShot(0, 0, true)
	test := func(action []float64) {
		result := env.Step(action, location, false)
		reward += result
		if result == 10 {
			goals++
		}
	}

	start := time.Now()
	for episode := 1; episode <= 1000; episode++ {
		action := []float64{
			((rand.Float64() * 180) - 90) * math.Pi / 180.0,
			(rand.Float64() * 90) * math.Pi / 180.0,
			rand.Float64() * 145.0}
		test(action)
	}
	elapsed := time.Since(start)
	fmt.Printf("Time: %s with reward: %d and %d goals\n", elapsed, reward, goals)

	// 1000000 Random Episodes Timed Multithreading
	// var wg sync.WaitGroup
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

	fmt.Printf("=====ENDING ENV SPACE TEST=====\n")

}
