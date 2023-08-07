package tests

import (
	"Soccer-Penalty-Kick-ML-Threading/environment/soccer"
	"Soccer-Penalty-Kick-ML-Threading/rl"
	"testing"
)

/**
 * TestSoccer()
 * Tests the Soccer environment with Q learning agent
 */
func TestSoccer(t *testing.T) {

	// Initialize env and parameters
	env := soccer.InitSoccer()
	max_episodes := 1000000
	max_actions := 150
	learning_rate := 0.93
	discount := 0.95
	exploration_rate := 1.0
	exploration_decay := 1.0 / float64(max_episodes)

	// Initialize agent and set policy
	agent := rl.InitQAgent(&env, max_episodes, max_actions, learning_rate, discount)
	agent.SetPolicy("", []float64{exploration_rate, exploration_decay})

	// Train the agent
	agent.Train(false)

	// Test the agent
	agent.Test(true)
}
