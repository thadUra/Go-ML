# Golang-Machine-Learning rlearning

[![Documentation](https://img.shields.io/badge/documentation-GoDoc-blue.svg)](https://pkg.go.dev/github.com/thadUra/Golang-Machine-Learning/rlearning)

Package rlearning is a ml reinforcement learning package for Go.

## Example Usage

Below contains example usage of the rlearning package on the frozen lake environment. This example can be found in `../tests/frozenlake_test.go`.

### Q-Learning on Frozen Lake
```
    // Initialize frozen lake env with 4x4 map
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
	agent := rlearning.NewQAgent(&env, max_episodes, max_actions, learning_rate, discount)
	agent.SetPolicy("DecayExploration", []float64{exploration_rate, exploration_decay})

	// Train the agent
	agent.Train(false)

	// Test the agent
	agent.Test(true)
```

#### Result
```
    ===MAP LAYOUT===
    S F F F 
    H F H F 
    F F F F 
    H F F G 
    ===END MAP LAYOUT===
    ===QAGENT TEST===
            Action     1: 2
            Action     2: 1
            Action     3: 1
            Action     4: 2
            Action     5: 2
            Action     6: 1
            Test Total Reward: 1.000000
    ===END QAGENT TEST===
```