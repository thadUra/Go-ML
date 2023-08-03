package tests

// import (
// 	"Soccer-Penalty-Kick-ML-Threading/nn"
// 	"Soccer-Penalty-Kick-ML-Threading/soccer"
// 	"fmt"
// 	"math"
// )

// /*
// POLICY INTERFACE AND HELPER FUNCS
// */
// type Policy interface {
// 	SelectAction(values []float64) float64
// }

// type BoltzmannQPolicy struct {
// 	tau  float64
// 	clip [2]float64
// }

// func (policy BoltzmannQPolicy) SelectAction(values []float64) float64 {
// 	/*
// 				Return the selected action
// 		        # Arguments
// 		            q_values (np.ndarray): List of the estimations of Q for each action
// 		        # Returns
// 		            Selection action

// 		        assert q_values.ndim == 1
// 		        q_values = q_values.astype('float64')
// 		        nb_actions = q_values.shape[0]

// 		        exp_values = np.exp(np.clip(q_values / self.tau, self.clip[0], self.clip[1]))
// 		        probs = exp_values / np.sum(exp_values)
// 		        action = np.random.choice(range(nb_actions), p=probs)
// 		        return action
// 	*/
// 	exp_values := values
// 	sum := 0.0
// 	for i := range exp_values {
// 		clip := (exp_values[i] / policy.tau)
// 		if clip < policy.clip[0] {
// 			clip = policy.clip[0]
// 		}
// 		if clip > policy.clip[1] {
// 			clip = policy.clip[1]
// 		}
// 		exp_values[i] = math.Exp(clip)
// 		sum += exp_values[i]
// 	}
// 	for i := range exp_values {
// 		exp_values[i] /= sum
// 	}

// 	return 0.0
// }

// func BuildPolicy() Policy {
// 	return BoltzmannQPolicy{1.0, [2]float64{-500, 500}}
// }

// /*
// MODEL GENERATION AND HELPER FUNCS
// */
// func BuildModel(input_nodes, output_nodes int) *nn.Network {
// 	var net nn.Network
// 	net.AddLayer(nn.InitFCLayer(input_nodes, 10))
// 	net.AddLayer(nn.InitActivationLayer(nn.Tanh, nn.TanhPrime))
// 	net.AddLayer(nn.InitFCLayer(10, 5))
// 	net.AddLayer(nn.InitActivationLayer(nn.Tanh, nn.TanhPrime))
// 	net.AddLayer(nn.InitFCLayer(5, output_nodes))
// 	net.SetLoss(nn.Mse, nn.Mse_prime)
// 	return &net
// }

// /*
// AGENT STRUCT AND HELPER FUNCS
// */
// type Agent struct {
// 	max_episodes int
// 	env          soccer.Environment
// 	model        *nn.Network
// 	policy       Policy
// }

// func (agt Agent) Train() {

// 	for i := 0; i < agt.max_episodes; i++ {

// 		// Get action and predicted Q-values from model + policy

// 		// Perform the action and get reward

// 		// Back Propagate to set weights and values

// 		// Repeat for another action

// 	}

// }

// /**
//  * RunDeepQTest()
//  * Creates a Deep-Q neural network for the soccer game env
//  * Extends nn package for deep-q learning
//  */
// func RunDeepQTest() {

// 	fmt.Printf("=====RUNNING DEEP Q TEST=====\n")

// 	// Initalize game environment
// 	params := game.InitSoccer(0, 0, 0, 0, 0, 0, 0, true)
// 	env := game.InitEnvironment(params)

// 	// Initialize network parameters
// 	num_input := len(env.OBSERVATION_SPACE)
// 	num_output := len(env.ACTION_SPACE)

// 	// Build NN model
// 	model := BuildModel(num_input, num_output)

// 	// Build Exploration Policy
// 	policy := BuildPolicy()

// 	// Build Agent
// 	agent := Agent{1000, env, model, policy}
// 	agent.Train()

// 	fmt.Printf("=====ENDING DEEP Q TEST=====\n")

// }
