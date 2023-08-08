package rl

import (
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

/**
 *  Policy Interface
 *  Outlines necessary accessors and modifiers for RL policies
 */
type Policy interface {
	SelectAction(mode string, values *mat.Dense, args []float64) float64
	Update()
}

/**
 *  InitPolicy()
 *  Wrapper function that initializes policy given string and args
 */
func InitPolicy(policyType string, args []float64) *Policy {
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

/**
 *  DecayExplorationPolicy struct
 *  Contains exploration rate and decay for selecting action over time
 */
type DecayExplorationPolicy struct {
	EXPLORATION_RATE  float64
	EXPLORATION_DECAY float64
}

/**
 *  DecayExplorationPolicy interface functions
 *  State represented by row -> args[0]
 */
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
func (policy *DecayExplorationPolicy) Update() {
	if policy.EXPLORATION_RATE > 0.001 {
		policy.EXPLORATION_RATE -= policy.EXPLORATION_DECAY
	}
}
