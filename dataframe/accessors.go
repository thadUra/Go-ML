package dataframe

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Wrapper print function given a range of row indices.
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
	if row >= df.rows || row < 0 {
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
	if row >= df.rows || row < 0 {
		return nil, errors.New("dataframe.Iat(): row out of range")
	}
	if col >= df.cols || row < 0 {
		return nil, errors.New("dataframe.Iat(): col out of range")
	}
	return df.data[df.order[col]][row], nil
}

// IlocRow returns a row of data given an index in the form of a map.
func (df *Dataframe) LocRow(row int) (map[string]interface{}, error) {
	if row >= df.rows || row < 0 {
		return nil, errors.New("dataframe.Loc(): row out of range")
	}
	ret := make(map[string]interface{})
	for i := 0; i < len(df.order); i++ {
		ret[df.order[i]] = df.data[df.order[i]][row]
	}
	return ret, nil
}

// LocCol returns a column of data given a label.
func (df *Dataframe) LocCol(col string) ([]interface{}, error) {
	if df.data[col] == nil {
		return nil, errors.New("dataframe.Pop(): column doesnt exist")
	}
	ret := make([]interface{}, df.rows)
	copy(ret, df.data[col])
	return ret, nil
}

// IocCol returns a column of data given an array of labels.
func (df *Dataframe) IlocCol(cols []string) ([][]interface{}, error) {
	for _, i := range cols {
		if df.data[i] == nil {
			return nil, errors.New("dataframe.Pop(): one of the columns dont exist")
		}
	}
	ret := make([][]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		ret[i] = make([]interface{}, df.rows)
		copy(ret[i], df.data[cols[i]])
	}
	return ret, nil
}

// IlocRow returns rows of data given a range of indices inclusive.
func (df *Dataframe) IlocRow(start, end int) (map[string][]interface{}, error) {
	if start >= df.rows || start < 0 {
		return nil, errors.New("dataframe.Iloc(): out of range")
	}
	if end >= df.rows || end < 0 {
		return nil, errors.New("dataframe.Iloc(): out of range")
	}
	ret := make(map[string][]interface{})
	for i := 0; i < len(df.order); i++ {
		ret[df.order[i]] = make([]interface{}, end-start+1)
		for j := start; j <= end; j++ {
			ret[df.order[i]][j-start] = df.data[df.order[i]][j]
		}
	}
	return ret, nil
}

// IsNull returns dataframe of booleans to see if a value is NaN or not.
func (df *Dataframe) IsNull() Dataframe {
	ret := DataframeFromMap(df.data)
	copy(ret.order, df.order)
	for i := 0; i < len(ret.order); i++ {
		for j := 0; j < ret.rows; j++ {
			switch v := ret.data[ret.order[i]][j].(type) {
			case float64:
				if math.IsNaN(v) {
					ret.data[ret.order[i]][j] = true
				} else {
					ret.data[ret.order[i]][j] = false
				}
			default:
				ret.data[ret.order[i]][j] = false
			}
		}
	}
	return ret
}

// Shape returns the dimensionality (rows and columns) of the dataframe.
func (df *Dataframe) Shape() (int, int) {
	return df.rows, df.cols
}

// Count returns the number of non-null values in each column.
func (df *Dataframe) Count() map[string]int {
	ret := make(map[string]int)
	for i := 0; i < len(df.order); i++ {
		count := 0
		for j := 0; j < df.rows; j++ {
			switch v := df.data[df.order[i]][j].(type) {
			case float64:
				if !math.IsNaN(v) {
					count++
				}
			default:
				count++
			}
		}
		ret[df.order[i]] = count
	}
	return ret
}

// Unique returns the number of unique non-null values in each column.
// If 'axis' is '1', it grabs unique values at each row instead.
func (df *Dataframe) Nunique(axis int) map[string]int {
	ret := make(map[string]int)
	if axis == 1 {
		for i := 0; i < df.rows; i++ {
			count := make(map[interface{}]int)
			for j := 0; j < len(df.order); j++ {
				switch v := df.data[df.order[j]][i].(type) {
				case float64:
					if !math.IsNaN(v) {
						count[v]++
					}
				default:
					count[v]++
				}
			}
			ret[strconv.Itoa(i)] = len(count)
		}
	} else {
		for i := 0; i < len(df.order); i++ {
			count := make(map[interface{}]int)
			for j := 0; j < df.rows; j++ {
				switch v := df.data[df.order[i]][j].(type) {
				case float64:
					if !math.IsNaN(v) {
						count[v]++
					}
				default:
					count[v]++
				}
			}
			ret[df.order[i]] = len(count)
		}
	}
	return ret
}
