package rl

/**
 *  Agent Interface
 *  Outlines necessary accessors and modifiers for RL Agent
 */
type Agent interface {
	SetPolicy(policyType string, args []float64)
	Train(info bool) (bool, error)
	Test(info bool)
}
