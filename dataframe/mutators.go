package dataframe

import (
	"errors"
	"math"
	"sort"
	"strconv"
)

// Relabel takes a string array and relabels the columns with the given strings.
// A new label cannot contain the name of an existing label.
func (df *Dataframe) Relabel(labels []string) {
	length := df.cols
	if len(labels) < df.cols {
		length = len(labels)
	}
	for i := 0; i < length; i++ {
		if df.data[labels[i]] != nil {
			continue
		}
		df.data[labels[i]] = make([]interface{}, df.rows)
		copy(df.data[labels[i]], df.data[df.order[i]])
		df.data[df.order[i]] = nil
		df.order[i] = labels[i]
	}
}

// Insert creates a new column and inserts it into the dataframe.
func (df *Dataframe) InsertCol(data []interface{}, label string) {
	df.cols++
	df.order = append(df.order, label)
	if len(data) > df.rows {
		df.rows = len(data)
		df.data[label] = make([]interface{}, df.rows)
		copy(df.data[label], data)
		for i := 0; i < len(df.order); i++ {
			if len(df.data[df.order[i]]) < df.rows {
				add := make([]interface{}, df.rows-len(df.data[df.order[i]]))
				for j := 0; j < len(add); j++ {
					add[j] = math.NaN()
				}
				df.data[df.order[i]] = append(df.data[df.order[i]], add...)
			}
		}
	} else {
		df.data[label] = make([]interface{}, df.rows)
		copy(df.data[label], data)
		for i := len(data); i < df.rows; i++ {
			df.data[label][i] = math.NaN()
		}
	}
}

// Insert creates a new row and inserts it into the dataframe.
func (df *Dataframe) InsertRow(data []interface{}) {
	df.rows++
	if len(data) > df.cols {
		// Determine the labels for the new columns and their order
		iter := 0
		new_cols := 0
		for len(df.order) < len(data) {
			exists := false
			for _, label := range df.order {
				if label == strconv.Itoa(iter) {
					exists = true
					break
				}
			}
			if !exists {
				df.order = append(df.order, strconv.Itoa(iter))
				df.data[strconv.Itoa(iter)] = make([]interface{}, df.rows)
				// Initialize values to be Nan
				for i := 0; i < df.rows; i++ {
					df.data[strconv.Itoa(iter)][i] = math.NaN()
				}
				new_cols++
			}
			iter++
		}
		df.cols = len(data)
		// Add row
		for i := 0; i < df.cols; i++ {
			if i < df.cols-new_cols {
				df.data[df.order[i]] = append(df.data[df.order[i]], data[i])
			} else {
				df.data[df.order[i]][df.rows-1] = data[i]
			}
		}
	} else {
		// Add row
		for i := 0; i < len(data); i++ {
			df.data[df.order[i]] = append(df.data[df.order[i]], data[i])
		}
		// Add NaN for empty cells in row
		for i := len(data); i < df.cols; i++ {
			df.data[df.order[i]] = append(df.data[df.order[i]], math.NaN())
		}
	}
}

// DropNull drops any rows if a single NaN value exists in its row.
func (df *Dataframe) DropNull() {
	// Grab indices of rows with null values
	indices := make(map[int]int)
	for i := 0; i < len(df.order); i++ {
		for j := 0; j < df.rows; j++ {
			switch v := df.data[df.order[i]][j].(type) {
			case float64:
				if math.IsNaN(v) {
					indices[j]++
				}
			}
		}
	}

	// Delete each row in indices
	for i := 0; i < len(df.order); i++ {
		new_col := make([]interface{}, 0)
		for j := 0; j < df.rows; j++ {
			if indices[j] == 0 {
				new_col = append(new_col, df.data[df.order[i]][j])
			}
		}
		df.data[df.order[i]] = new_col
	}
	df.rows -= len(indices)
}

// Pop removes and returns a column from the dataframe.
func (df *Dataframe) Pop(col string) ([]interface{}, error) {
	if df.data[col] == nil {
		return nil, errors.New("dataframe.Pop(): column doesnt exist")
	}
	df.cols--
	for i := 0; i < len(df.order); i++ {
		if df.order[i] == col {
			df.order = append(df.order[:i], df.order[i+1:]...)
			break
		}
	}
	ret := make([]interface{}, df.rows)
	copy(ret, df.data[col])
	df.data[col] = nil
	return ret, nil
}

// Sort_values sorts the dataframe according to the column `label`. If
// `ascending` is false, it sorts in descending order. If different types
// exist in a column, precedence is ordered by int or float64, bool,
// string, and then null values.
func (df *Dataframe) Sort_values(label string, ascending bool) {
	// Create array of indices for keeping track of sort across columns
	sorted_indices := make([]int, df.rows)
	for i := range sorted_indices {
		sorted_indices[i] = i
	}

	// Get indices after sort
	if df.data[label] == nil {
		return
	}
	if !ascending {
		sort.Sort(sort.Reverse(sortWrapper{df.data[label], sorted_indices}))
	} else {
		sort.Sort(sortWrapper{df.data[label], sorted_indices})
	}

	// Sort all the data according to indices
	for i := 0; i < len(df.order); i++ {
		if df.order[i] == label {
			continue
		}
		col_data := make([]interface{}, df.rows)
		for j := 0; j < df.rows; j++ {
			col_data[j] = df.data[df.order[i]][sorted_indices[j]]
		}
		df.data[df.order[i]] = col_data
	}
}

// sortWrapper defines a struct to help sort values and track their indices
// given a slice of interface values.
type sortWrapper struct {
	vals []interface{}
	idxs []int
}

// Functions here implement the sort interface.
func (s sortWrapper) Len() int {
	return len(s.vals)
}
func (s sortWrapper) Less(i, j int) bool {
	switch v1 := s.vals[i].(type) {
	case int:
		switch v2 := s.vals[j].(type) {
		case int:
			return v1 < v2
		case float64:
			if math.IsNaN(v2) {
				return true
			}
			return float64(v1) < v2
		default:
			return true
		}
	case float64:
		if math.IsNaN(v1) {
			return false
		}
		switch v2 := s.vals[j].(type) {
		case int:
			return v1 < float64(v2)
		case float64:
			if math.IsNaN(v2) {
				return true
			}
			return v1 < v2
		default:
			return true
		}
	case bool:
		switch v2 := s.vals[j].(type) {
		case int:
			return false
		case float64:
			if math.IsNaN(v2) {
				return true
			}
			return false
		case bool:
			if v1 && !v2 {
				return true
			} else {
				return false
			}
		default:
			return true
		}
	case string:
		switch v2 := s.vals[j].(type) {
		case int:
			return false
		case float64:
			if math.IsNaN(v2) {
				return true
			}
			return false
		case bool:
			return false
		case string:
			return v1 < v2
		}
	}
	return false
}
func (s sortWrapper) Swap(i, j int) {
	s.vals[i], s.vals[j] = s.vals[j], s.vals[i]
	s.idxs[i], s.idxs[j] = s.idxs[j], s.idxs[i]
}
