// The frozenlake package contains code for generating the frozen lake environment.
//
// Credit goes to the gym library for implementation.
package frozenlake

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	"github.com/thadUra/Go-ML/env"
)

// FrozenLake represents the frozen lake environment and its parameters.
type FrozenLake struct {
	CURRENT_STATE     int
	SLIPPERY          bool
	ROWS              int
	COLS              int
	ACTION_SIZE       int
	OBSERVATION_SIZE  int
	ACTION_SPACE      []int
	OBSERVATION_SPACE []string
}

// NewFrozenLake generates a new environment instance given the number of rows and columns for the map,
// the multiplier for how many holes there will be, and if the map is slippery.
func NewFrozenLake(rows, cols int, hole_multiplier float64, slippery bool) env.Environment {
	var frzn FrozenLake
	frzn.CURRENT_STATE = 0
	frzn.SLIPPERY = slippery
	frzn.ROWS = rows
	frzn.COLS = cols
	frzn.ACTION_SIZE = 1
	frzn.OBSERVATION_SIZE = 1
	frzn.ACTION_SPACE = []int{0, 1, 2, 3}
	frzn.OBSERVATION_SPACE = make([]string, rows*cols)

	// Generate map for S: start, G: goal, F: frozen, H: hole
	valid_map := false
	count, max_attempts := 0, 25
	for !valid_map && count < max_attempts {
		total_holes := 0
		max_holes := math.Sqrt(float64(rows * cols))
		for i := 0; i < rows*cols; i++ {
			if i == 0 {
				frzn.OBSERVATION_SPACE[i] = "S"
			} else if i == rows*cols-1 {
				frzn.OBSERVATION_SPACE[i] = "G"
			} else {
				if total_holes < int(max_holes) && rand.Float64() < hole_multiplier*max_holes/float64(rows*cols) {
					frzn.OBSERVATION_SPACE[i] = "H"
				} else {
					frzn.OBSERVATION_SPACE[i] = "F"
				}
			}
		}
		// Check if path exists using BFS
		queue := make([]int, 0)
		visited := make([]bool, rows*cols)
		queue = append(queue, 0)
		for len(queue) != 0 {
			top := queue[0]
			if visited[top] {
				queue = queue[1:]
			} else {
				r, c := int(float64(top)/float64(cols)), top%cols
				if r-1 >= 0 && frzn.OBSERVATION_SPACE[top-cols] != "H" {
					queue = append(queue, top-cols)
				}
				if r+1 < rows && frzn.OBSERVATION_SPACE[top+cols] != "H" {
					queue = append(queue, top+cols)
				}
				if c-1 >= 0 && frzn.OBSERVATION_SPACE[top-1] != "H" {
					queue = append(queue, top-1)
				}
				if c+1 < cols && frzn.OBSERVATION_SPACE[top+1] != "H" {
					queue = append(queue, top+1)
				}
				visited[top] = true
			}
		}
		if visited[rows*cols-1] {
			valid_map = true
		}
		count++
	}
	// Print map
	fmt.Println("===MAP LAYOUT===")
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			fmt.Printf("%s ", frzn.OBSERVATION_SPACE[r*cols+c])
		}
		fmt.Println("")
	}
	fmt.Println("===END MAP LAYOUT===")
	return env.Environment(&frzn)
}

// Step performs one action inside the environment.
func (frzn *FrozenLake) Step(
	action []float64,
) (float64, float64, bool, error) {
	// Check dimensions
	if len(action) > 1 || len(action) == 0 {
		return -1, -1, true, errors.New("FrozenLake.Step: action dimensions are incorrect")
	}
	value := int(action[0])

	// Perform action for left: 0, down: 1, right: 2, up: 3
	r, c := int(frzn.CURRENT_STATE/frzn.COLS), frzn.CURRENT_STATE%frzn.COLS
	if frzn.SLIPPERY {
		option := int(rand.Float64() * 3)
		if value == 0 {
			if option == 0 && c-1 >= 0 {
				frzn.CURRENT_STATE--
			} else if option == 1 && r+1 < frzn.ROWS {
				frzn.CURRENT_STATE += frzn.COLS
			} else if option == 2 && r-1 >= 0 {
				frzn.CURRENT_STATE -= frzn.COLS
			}
		} else if value == 1 {
			if option == 0 && r+1 < frzn.ROWS {
				frzn.CURRENT_STATE += frzn.COLS
			} else if option == 1 && c-1 >= 0 {
				frzn.CURRENT_STATE--
			} else if option == 2 && c+1 < frzn.COLS {
				frzn.CURRENT_STATE++
			}
		} else if value == 2 {
			if option == 0 && c+1 < frzn.COLS {
				frzn.CURRENT_STATE++
			} else if option == 1 && r+1 < frzn.ROWS {
				frzn.CURRENT_STATE += frzn.COLS
			} else if option == 2 && r-1 >= 0 {
				frzn.CURRENT_STATE -= frzn.COLS
			}
		} else if value == 3 {
			if option == 0 && r-1 >= 0 {
				frzn.CURRENT_STATE -= frzn.COLS
			} else if option == 1 && c-1 >= 0 {
				frzn.CURRENT_STATE--
			} else if option == 2 && c+1 < frzn.COLS {
				frzn.CURRENT_STATE++
			}
		}
	} else {
		if value == 0 && c-1 >= 0 {
			frzn.CURRENT_STATE--
		} else if value == 1 && r+1 < frzn.ROWS {
			frzn.CURRENT_STATE += frzn.COLS
		} else if value == 2 && c+1 < frzn.COLS {
			frzn.CURRENT_STATE++
		} else if value == 3 && r-1 >= 0 {
			frzn.CURRENT_STATE -= frzn.COLS
		}
	}
	// Get reward and check if done
	var reward float64
	var done bool
	if frzn.OBSERVATION_SPACE[frzn.CURRENT_STATE] == "G" {
		reward = 1.0
		done = true
	} else {
		reward = 0
		done = false
	}
	if frzn.OBSERVATION_SPACE[frzn.CURRENT_STATE] == "H" {
		reward = -1.0
		done = true
	}
	return float64(frzn.CURRENT_STATE), reward, done, nil
}

// Reset sets the current state back to the first square denoted as S.
func (frzn *FrozenLake) Reset() float64 {
	frzn.CURRENT_STATE = 0
	return 0.0
}

// GetNumActions returns the size of the action space.
func (frzn *FrozenLake) GetNumActions() int {
	return len(frzn.ACTION_SPACE)
}

// GetNumObservations returns the size of the observation space.
func (frzn *FrozenLake) GetNumObservations() int {
	return len(frzn.OBSERVATION_SPACE)
}
