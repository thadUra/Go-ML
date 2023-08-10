package rlearning

// Agent is a representation for reinforcement learning agents with setting a learning
// policy, training function, and testing function.
type Agent interface {
	SetPolicy(policyType string, args []float64)
	Train(info bool) (bool, error)
	Test(info bool)
}
