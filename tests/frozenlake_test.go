package tests

import (
	"Soccer-Penalty-Kick-ML-Threading/environment/frozenlake"
	"Soccer-Penalty-Kick-ML-Threading/rl"
	"testing"
)

/**
 * TestFrozenLake()
 * Tests the FrozenLake environment with Q learning agent
 */
func TestFrozenLake(t *testing.T) {

	// Initialize env
	env := frozenlake.InitFrozenLake(4, 4, 1.25, false)
	_, _, _, err := env.Step([]float64{})
	if err == nil {
		t.Fatalf(`Failed to get error from step with no action\n`)
	}

	// Initialize parameters
	max_episodes := 1000
	max_actions := 99
	learning_rate := 0.83
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
