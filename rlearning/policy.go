package rlearning

import (
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

// Policy is a representation for reinforcement learning policies with a function to
// select an action given values and a mode type. An update function is needed to update
// any policy parameters.
type Policy interface {
	SelectAction(mode string, values *mat.Dense, args []float64) float64
	Update()
}

// NewPolicy returns a policy defined by `policyType` acting as a wrapper initialization
// function.
func NewPolicy(policyType string, args []float64) *Policy {
	var policy Policy
	switch policyType {
	case "DecayExploration":
		test := DecayExplorationPolicy{
			EXPLORATION_RATE:  args[0],
			EXPLORATION_DECAY: args[1]}
		policy = Policy(&test)
	}
	return &policy
}

// DecayExplorationPolicy implements the Policy interface to represent a decayed
// exploration policy for reinforcement learning.
type DecayExplorationPolicy struct {
	EXPLORATION_RATE  float64
	EXPLORATION_DECAY float64
}

// SelectAction returns an action given a state from `args`.
func (policy *DecayExplorationPolicy) SelectAction(
	mode string,
	values *mat.Dense,
	args []float64,
) float64 {
	value := values.At(int(args[0]), 0.0)
	index := 0
	_, cols := values.Dims()
	if mode == "train" {
		if rand.Float64() > policy.EXPLORATION_RATE {
			for c := 1; c < cols; c++ {
				if values.At(int(args[0]), c) > value {
					value = values.At(int(args[0]), c)
					index = c
				}
			}
		} else {
			test := rand.Float64() * float64(cols)
			index = int(test)
		}
	} else if mode == "test" {
		for c := 1; c < cols; c++ {
			if values.At(int(args[0]), c) > value {
				value = values.At(int(args[0]), c)
				index = c
			}
		}
	}
	return float64(index)
}

// Update updates the `EXPLORATION_RATE` using the `EXPLORATION_DECAY` value.
func (policy *DecayExplorationPolicy) Update() {
	if policy.EXPLORATION_RATE > 0.001 {
		policy.EXPLORATION_RATE -= policy.EXPLORATION_DECAY
	}
}
