package tests

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/thadUra/Golang-Machine-Learning/rlearning"
	"gonum.org/v1/gonum/mat"
)

/**
 * TestPolicies()
 * Tests the different policies of the rl package
 */
func TestPolicies(t *testing.T) {
	// Init table
	rows, cols := 4, 5
	values := make([]float64, rows*cols)
	for i := range values {
		values[i] = rand.Float64()
	}
	table := mat.NewDense(rows, cols, values)
	ftable := mat.Formatted(table, mat.Prefix("        "), mat.Squeeze())
	fmt.Printf("table = %v\n", ftable)

	// Test DecayExplorationPolicy
	policy := rlearning.InitPolicy("DecayExploration", []float64{0.5, 1.0 / 1000.0})
	for i := 0; i < 25; i++ {
		ret := (*policy).SelectAction("train", table, []float64{float64(i % rows)})
		fmt.Printf("Training Action %d: %f\n", i+1, ret)
	}
	ret := (*policy).SelectAction("test", table, []float64{1.0})
	fmt.Printf("Testing Action: %f\n", ret)
}
