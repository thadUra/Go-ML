package tests

import (
	"testing"

	"github.com/thadUra/Go-ML/env/frozenlake"
	"github.com/thadUra/Go-ML/rlearning"
)

/**
 * TestFrozenLake()
 * Tests the FrozenLake environment with Q learning agent
 */
func TestFrozenLake(t *testing.T) {
	// Initialize frozen lake env with 4x4 map
	env := frozenlake.NewFrozenLake(4, 4, 1.25, false)
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
	agent := rlearning.NewQAgent(&env, max_episodes, max_actions, learning_rate, discount)
	agent.SetPolicy("DecayExploration", []float64{exploration_rate, exploration_decay})

	// Train the agent
	agent.Train(false)

	// Test the agent
	agent.Test(true)
}
