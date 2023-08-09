package tests

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"testing"

	cluster "github.com/thadUra/Golang-Machine-Learning/clustering"
)

/**
 * TestCluster()
 * Tests the various methods and algorithms in the cluster package with the Iris dataset
 * Uses scatterplot to view results
 */
func TestCluster(t *testing.T) {
	// Open iris data file (WIP CLOSE FILE)
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
		t.Fatalf(`TestCluster(): failed at pca plotting -> "%s"`, err)
	}

	// Test KMeans
	km := cluster.InitKMeans(3, 500)
	err = km.Train(result)
	if err != nil {
		t.Fatalf(`TestCluster(): failed at kmeans training -> "%s"`, err)
	}
	result, new_label, err := km.Evaluate(result)
	if err != nil {
		t.Fatalf(`TestCluster(): failed at kmeans evaluation -> "%s"`, err)
	}

	// Generate scatterplot for KMeans results
	plot_params = []string{"../tests/misc/kmeansIrisTest.png", "KMeans Test", "X", "Y"}
	err = cluster.ScatterPlot2DimenData(result, new_label, plot_params)
	if err != nil {
		t.Fatalf(`TestCluster(): failed at kmeans plotting -> "%s"`, err)
	}

	// Open make_moons data file
	file, err = os.Open("./misc/make_moons.csv")
	if err != nil {
		t.Fatalf(`TestCluster(): failed to open file -> "%s"`, err)
	}
	r = csv.NewReader(file)

	// Parse data and store into slices
	data = [][]float64{}
	label = []string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf(`TestCluster(): failed in file parsing -> "%s"`, err)
		}
		next := make([]float64, 2)
		for i := 0; i < 2; i++ {
			next[i], err = strconv.ParseFloat(record[i], 64)
			if err != nil {
				t.Fatalf(`TestCluster(): failed in file parsing -> "%s"`, err)
			}
		}
		data = append(data, next)
		label = append(label, record[2])
	}

	// Test Kmeans
	km = cluster.InitKMeans(2, 500)
	err = km.Train(data)
	if err != nil {
		t.Fatalf(`TestCluster(): failed at kmeans training -> "%s"`, err)
	}
	result, new_label, err = km.Evaluate(data)
	if err != nil {
		t.Fatalf(`TestCluster(): failed at kmeans evaluation -> "%s"`, err)
	}

	// Generate scatterplot for KMeans results
	plot_params = []string{"../tests/misc/kmeansMoonsTest.png", "KMeans Test", "X", "Y"}
	err = cluster.ScatterPlot2DimenData(result, new_label, plot_params)
	if err != nil {
		t.Fatalf(`TestCluster(): failed at kmeans plotting -> "%s"`, err)
	}

	// Test Spectral Clustering
	new_label, err = cluster.Spectral(data, 0.4)
	if err != nil {
		t.Fatalf(`TestCluster(): failed at spectral clustering -> "%s"`, err)
	}

	// Generate scatterplot for Spectral Clustering results
	plot_params = []string{"../tests/misc/spectralMoonsTest.png", "Spectral Clustering Test", "X", "Y"}
	err = cluster.ScatterPlot2DimenData(data, new_label, plot_params)
	if err != nil {
		t.Fatalf(`TestCluster(): failed at spectral plotting -> "%s"`, err)
	}
}
