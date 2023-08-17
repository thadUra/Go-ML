// The dataframe package contains code for the dataframe datatype, its functionality,
// and manipulation. Currently, the only variable types accepted are ints, strings,
// float64s, and bools. The ints need to be in base10. However, all slices and data
// need to be formatted as an interface type.
//
// Inspiration for the datatype comes from the pandas.Dataframe library.
// However, this package is very limited in functionality in comparison. For example, row indexing is not currently supported unlike pandas. More functionality will be available with more support and time.
package dataframe

import (
	"encoding/csv"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

// Dataframe represents the a two-dimensional tabular data structure
// inspired by the pandas dataframe.
type Dataframe struct {
	rows  int
	cols  int
	data  map[string][]interface{}
	order []string
}

// DataframeFromMap returns a dataframe given a map of tabular data.
func DataframeFromMap(data map[string][]interface{}) Dataframe {
	var df Dataframe

	// Get number of rows and cols
	for label, vals := range data {
		df.cols++
		df.order = append(df.order, label)
		if len(vals) > df.rows {
			df.rows = len(vals)
		}
	}

	// Initialize table data
	df.data = make(map[string][]interface{})
	for i := 0; i < df.cols; i++ {
		df.data[df.order[i]] = make([]interface{}, df.rows)
		copy(df.data[df.order[i]], data[df.order[i]])
		for j := len(data[df.order[i]]); j < df.rows; j++ {
			df.data[df.order[i]][j] = math.NaN()
		}
	}
	return df
}

// DataframeFromSlice returns a dataframe given a slice of data.
func DataframeFromSlice(data []interface{}) Dataframe {
	var df Dataframe
	df.order = append(df.order, "0")
	df.rows = len(data)
	df.cols = 1
	df.data = make(map[string][]interface{})
	df.data["0"] = make([]interface{}, len(data))
	copy(df.data["0"], data)
	return df
}

// DataframeFrom2DSlice returns a dataframe given a two dimensional
// slice of data.
func DataframeFrom2DSlice(data [][]interface{}) Dataframe {
	var df Dataframe
	// Get number of rows and cols
	df.rows = len(data)
	for i := 0; i < df.rows; i++ {
		if len(data[i]) > df.cols {
			df.cols = len(data[i])
		}
	}

	// Initialize table data
	df.data = make(map[string][]interface{})
	for i := 0; i < df.cols; i++ {
		df.order = append(df.order, strconv.Itoa(i))
		df.data[df.order[i]] = make([]interface{}, df.rows)
	}

	// Fill table with data
	for i := 0; i < df.rows; i++ {
		for j := 0; j < len(data[i]); j++ {
			df.data[df.order[j]][i] = data[i][j]
		}
		// Fill empty cells with NaN
		for j := len(data[i]); j < df.cols; j++ {
			df.data[df.order[j]][i] = math.NaN()
		}
	}
	return df
}

// DataframeFromCSV returns a dataframe given a filepath name for a
// csv file. If 'header' is 'true', then it is assumed that the first
// row in the csv file is a header row and not a data row.
func DataframeFromCSV(filepath string, header bool) Dataframe {
	// Open csv file for dimension read
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf(`DataframeFromCSV(): failed to open file -> "%s"`, err)
	}
	r := csv.NewReader(file)

	// Get csv dimensions
	rows, cols := 0, 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf(`DataframeFromCSV(): failed in file reading -> "%s"`, err)
		}
		rows++
		if len(record) > cols {
			cols = len(record)
		}
	}
	if header {
		rows--
	}
	file.Close()

	// Create dataframe
	var df Dataframe
	df.rows = rows
	df.cols = cols
	df.data = make(map[string][]interface{})

	// Open csv file again for data read
	file, err = os.Open(filepath)
	if err != nil {
		log.Fatalf(`DataframeFromCSV(): failed to open file -> "%s"`, err)
	}
	r = csv.NewReader(file)
	added_header := false
	for i := 0; i < rows; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf(`DataframeFromCSV(): failed in file reading -> "%s"`, err)
		}
		// Handle any column headers if any
		if !added_header {
			if header {
				df.order = append(df.order, record...)
			}
			if len(df.order) < df.cols {
				for j := len(df.order); j < df.cols; j++ {
					df.order = append(df.order, strconv.Itoa(j))
				}
			}
			for j := range df.order {
				df.data[df.order[j]] = make([]interface{}, df.rows)
			}
			added_header = true
			if header {
				i--
				continue
			}
		}
		for j := 0; j < df.cols; j++ {
			if j < len(record) {
				if it, err := strconv.ParseInt(record[j], 10, 0); err == nil {
					df.data[df.order[j]][i] = int(it)
					continue
				} else if fl, err := strconv.ParseFloat(record[j], 64); err == nil {
					df.data[df.order[j]][i] = fl
					continue
				} else if bo, err := strconv.ParseBool(record[j]); err == nil {
					df.data[df.order[j]][i] = bo
					continue
				} else if record[j] == "Null" || record[j] == "N/A" || record[j] == "" || record[j] == "NaN" {
					df.data[df.order[j]][i] = math.NaN()
				} else {
					df.data[df.order[j]][i] = record[j]
				}

			} else {
				df.data[df.order[j]][i] = math.NaN()
			}
		}
	}
	file.Close()
	return df
}
