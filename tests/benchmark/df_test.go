package bench

import (
	"fmt"
	"testing"
	"time"

	"github.com/thadUra/Go-ML/dataframe"
)

func TestDF(t *testing.T) {
	// Start timer
	start := time.Now()

	// Import data from a CSV file
	df := dataframe.DataframeFromCSV("../misc/iris_data.csv", false)
	labels := []string{"Sepal Length", "Sepal Width", "Petal Length", "Petal Width", "Species"}
	df.Relabel(labels)

	// Example Accessor functions
	df.Shape()
	df.At("Sepal Width", 67)

	// Sort dataframe based on sepal length
	df.Sort_values("Sepal Length", true)

	// Remove a column and get its data
	df.Pop("Petal Width")

	// End timer
	duration := time.Since(start)
	fmt.Printf("Dataframe: %v seconds\n", duration.Seconds())
}
