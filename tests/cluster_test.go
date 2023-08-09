package tests

import (
	cluster "Soccer-Penalty-Kick-ML-Threading/clustering"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"testing"
)

/**
 * TestCluster()
 * Tests the various methods and algorithms in the cluster package with the Iris dataset
 * Uses scatterplot to view results
 */
func TestCluster(t *testing.T) {
	// Open data file
	file, err := os.Open("./misc/iris_data.csv")
	if err != nil {
		t.Fatalf(`TestCluster(): failed to open file -> "%s"`, err)
	}
	r := csv.NewReader(file)

	// Parse data and store into slices
	var data [][]float64
	var label []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf(`TestCluster(): failed in file parsing -> "%s"`, err)
		}
		next := make([]float64, 4)
		for i := 0; i < 4; i++ {
			next[i], err = strconv.ParseFloat(record[i], 64)
			if err != nil {
				t.Fatalf(`TestCluster(): failed in file parsing -> "%s"`, err)
			}
		}
		data = append(data, next)
		label = append(label, record[4])
	}

	// Test PCA
	result, err := cluster.PCA(data, 2)
	if err != nil {
		t.Fatalf(`TestCluster(): failed at building PCA -> "%s"`, err)
	}

	// Generate scatterplot for PCA results
	plot_params := []string{"../tests/misc/pcaTest.png", "PCA Test", "PCA 1", "PCA 2"}
	err = cluster.ScatterPlot2DimenData(result, label, plot_params)
	if err != nil {
		t.Fatalf(`TestCluster(): failed at plotting -> "%s"`, err)
	}

	// Test KMeans
}
