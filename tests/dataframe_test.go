package tests

import (
	"testing"

	"github.com/thadUra/Go-ML/dataframe"
)

/**
 * TestDataframe()
 * Tests the dataframe package
 */
func TestDataframe(t *testing.T) {
	// Test data
	data := make(map[interface{}][]interface{})
	floats := []float64{1.0, 2.0}
	strs := []string{"first", "second"}
	names := make([]interface{}, len(strs))
	nums := make([]interface{}, len(floats))
	for i, s := range strs {
		names[i] = s
	}
	for i, f := range floats {
		nums[i] = f
	}
	data["names"] = names
	data["floats"] = nums

	// Test dataframe creation
	df := dataframe.DataframeFromMap(data)
	df.Print()
}
