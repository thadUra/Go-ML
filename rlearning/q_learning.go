package rlearning

import (
	"fmt"

	env "github.com/thadUra/Golang-Machine-Learning/env"
	"gonum.org/v1/gonum/mat"
)

// QAgent implements the Agent interface for a Q-Learning agent. It contains the max
// iterations for training, max actions in each iteration, learning rate, discount rate,
// the learning policy, and table of Q values.
type QAgent struct {
	ENV           *env.Environment
	MAX_EPISODES  int
	MAX_ACTIONS   int
	LEARNING_RATE float64
	DISCOUNT      float64
	POLICY        *Policy
	Q_TABLE       *mat.Dense
}

// NewQAgent returns a new QAgent instance given the environment and other parameters.
func NewQAgent(
	env *env.Environment,
	max_eps int,
	max_acts int,
	learning_rate float64,
	discount float64,
) QAgent {
	var agent QAgent
	agent.ENV = env
	agent.MAX_EPISODES = max_eps
	agent.MAX_ACTIONS = max_acts
	agent.LEARNING_RATE = learning_rate
	agent.DISCOUNT = discount
	agent.Q_TABLE = mat.NewDense((*env).GetNumObservations(), (*env).GetNumActions(), nil)
	agent.Q_TABLE.Zero()
	return agent
}

// SetPolicy sets the learning policy for the QAgent.
func (agt *QAgent) SetPolicy(policyType string, args []float64) {
	agt.POLICY = NewPolicy("DecayExploration", args)
}

// Train performs Q-learning on its environment. It returns `true` if it successfully
// trains. It returns `false` if an error is thrown. If `info` is `true`, then print
// statements are given for the reward at each episode.
func (agt *QAgent) Train(info bool) (bool, error) {
	// Info print statement
	if info {
		fmt.Println("=====")
	}

	// Training loop
	for i := 0; i < agt.MAX_EPISODES; i++ {
		state := (*agt.ENV).Reset()
		total_reward := 0.0
		for j := 0; j < agt.MAX_ACTIONS; j++ {
			// Get next action
			action := (*agt.POLICY).SelectAction("train", agt.Q_TABLE, []float64{state})
			if info {
				fmt.Printf("	action: %d\n", int(action))
			}
			new_state, reward, done, err := (*agt.ENV).Step([]float64{action})
			// Check error from step
			if err != nil {
				// Info print statement
				if info {
					fmt.Println("=====")
				}
				return false, err
			}
			// Get max q value for state
			max_val_in_new_state := agt.Q_TABLE.At(int(new_state), 0)
			_, cols := agt.Q_TABLE.Dims()
			for c := 1; c < cols; c++ {
				if agt.Q_TABLE.At(int(state), c) > max_val_in_new_state {
					max_val_in_new_state = agt.Q_TABLE.At(int(state), c)
				}
			}
			// Update q value in table
			next_state_val := agt.Q_TABLE.At(int(state), int(action)) + agt.LEARNING_RATE*(reward+(agt.DISCOUNT*max_val_in_new_state)-agt.Q_TABLE.At(int(state), int(action)))
			agt.Q_TABLE.Set(int(state), int(action), next_state_val)
			// Update reward
			total_reward += reward
			// Update next state
			state = new_state
			// Kill episode if environment is terminated
			if done {
				if info {
					fmt.Printf("Total reward for episode %d: %f\n", i+1, total_reward)
				}
				break
			}
		}
		// Update exploration policy
		(*agt.POLICY).Update()
	}
	// Info print statement
	if info {
		fmt.Println("=====")
	}
	return true, nil
}

// Test performs one episode on the environment. It returns `true` if it successfully
// performs the episode. If `info` is `true`, then print statements are given for the
// episode.
func (agt *QAgent) Test(info bool) (bool, error) {
	// Info print statement
	if info {
		fmt.Println("=====")
	}

	// Test from reset environment
	state := (*agt.ENV).Reset()
	total_reward := 0.0
	for j := 0; j < agt.MAX_ACTIONS; j++ {
		// Get next action
		action := (*agt.POLICY).SelectAction("test", agt.Q_TABLE, []float64{state})
		// fmt.Printf("	Prev State %d: %d\n", j+1, int(state))
		fmt.Printf("	Action     %d: %d\n", j+1, int(action))
		new_state, reward, done, err := (*agt.ENV).Step([]float64{action})
		// fmt.Printf("	Next State %d: %d\n", j+1, int(new_state))
		// Check error from step
		if err != nil {
			if info {
				fmt.Println("=====")
			}
			return false, err
		}
		// Update reward
		total_reward += reward
		// Update state
		state = new_state
		// Kill episode once done
		if done {
			break
		}
	}
	if info {
		fmt.Printf("	Testing Total Reward: %f\n", total_reward)
		fmt.Println("=====")
	}
	return true, nil
}
