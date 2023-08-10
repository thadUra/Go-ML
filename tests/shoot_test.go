package tests

// import (
// 	"testing"

// 	"github.com/thadUra/Golang-Machine-Learning/env/soccer"
// 	"github.com/thadUra/Golang-Machine-Learning/rlearning"
// )

// /**
//  * TestSoccerShoot()
//  * Tests the Soccer shot function
//  */
// func TestSoccerShoot(t *testing.T) {
// 	// Initialize env and parameters
// 	field := soccer.GenerateField(0, 0, 0, 0, 0, 0, 0, true)
// 	pos := soccer.GeneratePos(0, 0, true)

// 	max_episodes := 1000000
// 	max_actions := 500
// 	learning_rate := 0.93
// 	discount := 0.95
// 	exploration_rate := 1.0
// 	exploration_decay := 1.0 / float64(max_episodes)
// 	if env.GetNumObservations() != 77850 {
// 		t.Fatalf(`env.GetNumObservations() is incorrect: want "%d"`, 77850)
// 	}
// 	if env.GetNumActions() != 9 {
// 		t.Fatalf(`env.GetNumActions() is incorrect: want "%d"`, 9)
// 	}

// 	// Initialize agent and set policy
// 	agent := rlearning.NewQAgent(&env, max_episodes, max_actions, learning_rate, discount)
// 	agent.SetPolicy("", []float64{exploration_rate, exploration_decay})

// 	// Train the agent
// 	agent.Train(false)

// 	// Test the agent
// 	agent.Test(true)
// }
