# Go-ML dataframe

[![Documentation](https://img.shields.io/badge/documentation-GoDoc-blue.svg)](https://pkg.go.dev/github.com/thadUra/Go-ML/dataframe)

Package dataframe is a library inspired by the pandas.Dataframe library. It is very basic in comparison, but will include more functionality in the future.

## Example Usage

Below contains example usage of the dataframe package on the Iris dataset. Some more examples can also be found in `../tests/dataframe_test.go`.

### Dataframe Manipulation on Iris Dataset
```
    // Import data from a CSV file
	df := dataframe.DataframeFromCSV("../tests/misc/iris_data.csv", false)
	labels := []string{"Sepal Length", "Sepal Width", "Petal Length", "Petal Width", "Species"}
	df.Relabel(labels)
	df.Head(5)

	// Example Accessor functions
	rows, cols := df.Shape()
	fmt.Printf("Rows: %d, Cols: %d\n", rows, cols)
	value, _ := df.At("Sepal Width", 67)
	fmt.Printf("Value at column sepal width, row 67: %f\n", value)

	// Sort dataframe based on sepal length
	df.Sort_values("Sepal Length", true)
	df.Head(5)

	// Remove a column and get its data
	df.Pop("Petal Width")
```

#### Result
```
    +----------------------------------------------------------------------------+
    | id | Sepal Length | Sepal Width | Petal Length | Petal Width | Species     |
    +----------------------------------------------------------------------------+
    | 1  | 5.10000      | 3.50000     | 1.40000      | 0.20000     | Iris-setosa |
    | 2  | 4.90000      | 3.00000     | 1.40000      | 0.20000     | Iris-setosa |
    | 3  | 4.70000      | 3.20000     | 1.30000      | 0.20000     | Iris-setosa |
    | 4  | 4.60000      | 3.10000     | 1.50000      | 0.20000     | Iris-setosa |
    | 5  | 5.00000      | 3.60000     | 1.40000      | 0.20000     | Iris-setosa |
    +----------------------------------------------------------------------------+

    [5 rows x 5 columns]
    Rows: 150, Cols: 5
    Value at column sepal width, row 67: 2.700000
    +----------------------------------------------------------------------------+
    | id | Sepal Length | Sepal Width | Petal Length | Petal Width | Species     |
    +----------------------------------------------------------------------------+
    | 1  | 4.30000      | 3.00000     | 1.10000      | 0.10000     | Iris-setosa |
    | 2  | 4.40000      | 3.20000     | 1.30000      | 0.20000     | Iris-setosa |
    | 3  | 4.40000      | 3.00000     | 1.30000      | 0.20000     | Iris-setosa |
    | 4  | 4.40000      | 2.90000     | 1.40000      | 0.20000     | Iris-setosa |
    | 5  | 4.50000      | 2.30000     | 1.30000      | 0.30000     | Iris-setosa |
    +----------------------------------------------------------------------------+

    [5 rows x 5 columns]
```