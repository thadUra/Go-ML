package dataframe

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func (df *Dataframe) printWrapper(start, end int) {
	// Get max len for each column
	lengths := make(map[string]int)

	// Start with id column
	id_len := 1
	n := end
	for n > 9 {
		n /= 10
		id_len++
	}
	if id_len < 2 {
		id_len = 2
	}
	lengths["id"] = id_len

	// Get the lengths of other columns
	for i := range df.order {
		length := len(df.order[i])
		maxLen := func(a int) {
			if a > length {
				length = a
			}
		}
		for j := start; j < end; j++ {
			for k := 0; k < len(df.order); k++ {
				switch v := df.data[df.order[i]][j].(type) {
				case int:
					len := 1
					for v > 9 {
						v /= 10
						len++
					}
					maxLen(len)
				case string:
					maxLen(len(v))
				case float64:
					if math.IsNaN(v) {
						maxLen(3)
					} else {
						len := 1
						for v > 9 {
							v /= 10
							len++
						}
						maxLen(len + 7)
					}
				case bool:
					if v {
						maxLen(4)
					} else {
						maxLen(5)
					}
				}
			}
		}
		lengths[df.order[i]] = length
	}

	// Get line length
	var lineLength int
	for _, c := range lengths {
		lineLength += c + 3
	}
	lineLength += 1

	// Print the headers
	fmt.Printf("+%s+\n", strings.Repeat("-", lineLength-2)) // lineLength-2 because of "+" as first and last character
	fmt.Printf("| %-*s ", lengths["id"], "id")
	for i := 0; i < len(df.order); i++ {
		fmt.Printf("| %-*s ", lengths[df.order[i]], df.order[i])
	}
	fmt.Printf("|\n+%s+\n", strings.Repeat("-", lineLength-2))

	// Print the data
	for r := start; r < end; r++ {
		fmt.Printf("| %-*d ", lengths["id"], r+1)
		for i := 0; i < len(df.order); i++ {
			switch v := df.data[df.order[i]][r].(type) {
			case int:
				fmt.Printf("| %-*d ", lengths[df.order[i]], v)
			case string:
				fmt.Printf("| %-*s ", lengths[df.order[i]], v)
			case float64:
				fmt.Printf("| %-*.5f ", lengths[df.order[i]], v)
			case bool:
				fmt.Printf("| %-*t ", lengths[df.order[i]], v)
			}
		}
		fmt.Printf("|\n")
	}

	// Print last line
	fmt.Printf("+%s+\n", strings.Repeat("-", lineLength-2))
}

// Print prints all data from the dataframe type.
func (df *Dataframe) Print() {
	df.printWrapper(0, df.rows)
	fmt.Printf("\n[%d rows x %d columns]\n", df.rows, df.cols)
}

// Head prints the first n rows from the dataframe type.
func (df *Dataframe) Head(n int) {
	max_rows := n
	if df.rows < n {
		max_rows = df.rows
	}
	df.printWrapper(0, max_rows)
	fmt.Printf("\n[%d rows x %d columns]\n", max_rows, df.cols)
}

// Head prints the last n rows from the dataframe type.
func (df *Dataframe) Tail(n int) {
	max_rows := df.rows - n
	if max_rows < 0 {
		max_rows = 0
	}
	df.printWrapper(max_rows, df.rows)
	fmt.Printf("\n[%d rows x %d columns]\n", df.rows-max_rows, df.cols)
}

// At returns a value given a label for the row and column.
func (df *Dataframe) At(label string, row int) (interface{}, error) {
	if row > df.rows || row < 0 {
		return nil, errors.New("dataframe.At(): row out of range")
	}
	found := false
	for _, col := range df.order {
		if label == col {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("dataframe.At(): col does not exist")
	}
	return df.data[label][row], nil
}

// Iat returns a value given an index for the row and column.
func (df *Dataframe) Iat(row, col int) (interface{}, error) {
	if row > df.rows || row < 0 {
		return nil, errors.New("dataframe.Iat(): row out of range")
	}
	if col > df.cols || row < 0 {
		return nil, errors.New("dataframe.Iat(): col out of range")
	}
	return df.data[df.order[col]][row], nil
}

// Loc returns a row of data given a value.
func (df *Dataframe) Loc(value interface{}) {

}

// Iloc returns a row of data given an index value.
func (df *Dataframe) Iloc(idx int) {

}

// Iloc returns rows of data given a range of indices.
func (df *Dataframe) IlocRange(start, end int) {

}

// IsNull prints booleans to see if a value in the dataframe is NaN or not.
func (df *Dataframe) IsNull() {

}

// Iterrows returns something to allow iterating over the dataframe
func (df *Dataframe) Iterrows() {

}

// Shape returns the dimensionality (rows and columns) of the dataframe.
func (df *Dataframe) Shape() (int, int) {
	return df.rows, df.cols
}

// Count returns the number of non-null values in each column.
func (df *Dataframe) Count() {

}

// Unique returns the number of non-null values in each column.
// If 'axis' is '1', it grabs unique values at each row instead.
func (df *Dataframe) Nunique(axis int) {

}
