// The dataframe package contains code for the dataframe datatype, its functionality,
// and manipulation.
//
// Inspiration for the datatype comes from the pandas.Dataframe library.
package dataframe

// Dataframe represents the a two-dimensional tabular data structure
// inspired by the pandas dataframe.
type Dataframe struct {
	rows int
	cols int
	data map[string][]interface{}
}

// DataframeFromMap returns a dataframe given a map of tabular data.
func DataframeFromMap(data map[interface{}][]interface{}) Dataframe {
	var df Dataframe
	return df
}

// DataframeFromSlice returns a dataframe given a slice of data.
func DataframeFromSlice(data []interface{}) Dataframe {
	var df Dataframe
	return df
}

// DataframeFrom2DSlice returns a dataframe given a two dimensional
// slice of data.
func DataframeFrom2DSlice(data [][]interface{}) Dataframe {
	var df Dataframe
	return df
}

// DataframeFromCSV returns a dataframe given a filepath name for a
// csv file. If 'header' is 'true', then it is assumed that the first
// row in the csv file is a header row and not a data row.
func DataframeFromCSV(filepath string, header bool) Dataframe {
	var df Dataframe
	return df
}
