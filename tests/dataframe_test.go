package tests

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/thadUra/Go-ML/dataframe"
)

/**
 * TestDataframe()
 * Tests the dataframe package
 */
func TestDataframe(t *testing.T) {
	// Sample creation data
	data := make(map[string][]interface{})

	floats := []float64{1.0, 2.0}
	strs := []string{"first", "second", "LONG THIRD"}
	ints := []int{1, 6, 432423, 2}
	bools := []bool{true, false, true, false, false}

	data["floats"] = make([]interface{}, len(floats))
	for i, x := range floats {
		data["floats"][i] = x
	}

	data["strs"] = make([]interface{}, len(strs))
	for i, x := range strs {
		data["strs"][i] = x
	}

	data["ints"] = make([]interface{}, len(ints))
	for i, x := range ints {
		data["ints"][i] = x
	}

	data["bools"] = make([]interface{}, len(bools))
	for i, x := range bools {
		data["bools"][i] = x
	}

	data2D := make([][]interface{}, int(rand.Float64()*9)+1)
	for i := 0; i < len(data2D); i++ {
		data2D[i] = make([]interface{}, 5)
		for j := 0; j < len(data2D[i]); j++ {
			data2D[i][j] = rand.Float64() * math.Pow(10, float64(i+1))
		}
	}

	// Test dataframe creation functions
	df := dataframe.DataframeFromMap(data)
	df = dataframe.DataframeFromSlice(data["bools"])
	df = dataframe.DataframeFrom2DSlice(data2D)
	df = dataframe.DataframeFromCSV("../tests/misc/iris_data.csv", false)

	// Test accessor functions
	df.Head(5)
	df.Tail(5)
	val, _ := df.At("0", 2)
	fmt.Println(val)
	val, _ = df.At("0", -1)
	fmt.Println(val)
	val, _ = df.Iat(2, 2)
	fmt.Println(val)
	val, _ = df.Iat(-1, 1)
	fmt.Println(val)
	val, _ = df.Loc(122)
	fmt.Println(val)
	val, _ = df.Iloc(100, 106)
	fmt.Println(val)
	df = dataframe.DataframeFromMap(data)
	null := df.IsNull()
	df.Head(5)
	null.Head(5)
	count := df.Count()
	fmt.Println(count)
	unique := df.Nunique(0)
	fmt.Println(unique)
	unique = df.Nunique(1)
	fmt.Println(unique)

	// Test mutator functions
	new_col := make([]interface{}, 7)
	for j := 0; j < len(new_col); j++ {
		new_col[j] = rand.Float64() * 10
	}
	new_col[2] = math.NaN()
	df.InsertCol(new_col, "testNewCol")
	df.Head(8)
	new_row := make([]interface{}, 7)
	for j := 0; j < len(new_row); j++ {
		new_row[j] = rand.Float64() * 10
	}
	new_row[1] = true
	df.InsertRow(new_row[0:4])
	df.Tail(5)
	df.DropNull()
	df.Head(5)
	df.InsertRow(new_row[0:6])
	new_row[1] = 523
	df.Tail(5)
	df.InsertRow(new_row[0:7])
	df.Tail(5)
	new_row[1] = rand.Float64()
	df.InsertRow(new_row[0:7])
	fmt.Println(df.Pop("-1"))
	fmt.Println(df.Pop("testNewCol"))
	df.Head(5)
	df.Sort_values("bools", true)
	df.Head(5)
	new_labels := []string{"some", "label", "here", "i", "guess", "lol", "test"}
	df.Relabel(new_labels)
	df.Head(5)
	new_labels[0] = "REPLACE HERE?"
	df.Relabel(new_labels)
	df.Head(5)
}
