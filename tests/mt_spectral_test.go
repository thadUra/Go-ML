package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/thadUra/Go-ML/cluster"
	"github.com/thadUra/Go-ML/dataframe"
)

/**
 * TestMultithreadSpectral()
 * Tests multithreading on spectral clustering with make_moons data
 */
func TestMultithreadSpectral(t *testing.T) {
	// Make dataframe of make_moons_ext data
	df := dataframe.DataframeFromCSV("./misc/make_moons_ext.csv", true)
	df.Head(5)
	rows, _ := df.Shape()

	data_intfc, err := df.IlocCol([]string{"X", "Y"})
	if err != nil {
		t.Fatalf(`TestMultithreadSpectral(): failed at getting data -> "%s"`, err)
	}

	data := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]float64, 2)
		data[i][0] = data_intfc[0][i].(float64)
		data[i][1] = data_intfc[1][i].(float64)
	}

	// Test single thread
	start := time.Now()
	_, err = cluster.Spectral(data, 0.4)
	if err != nil {
		t.Fatalf(`TestMultithreadSpectral(): failed at spectral clustering -> "%s"`, err)
	}
	duration := time.Since(start)
	fmt.Printf("Single thread: %v\n\n", duration)

	// Test multithread
	start = time.Now()
	new_label, err := spectralMt(data, 0.4)
	if err != nil {
		t.Fatalf(`TestMultithreadSpectral(): failed at spectral clustering -> "%s"`, err)
	}
	duration = time.Since(start)
	fmt.Printf("Multithread: %v\n\n", duration)

	// Print
	// Generate scatterplot for Spectral Clustering results
	plot_params := []string{"../tests/misc/mtSpectral.png", "Spectral Clustering Test", "X", "Y"}
	err = cluster.ScatterPlot2DimenData(data, new_label, plot_params)
	if err != nil {
		t.Fatalf(`TestCluster(): failed at spectral plotting -> "%s"`, err)
	}
}

func spectralMt(data [][]float64, radius float64) ([]string, error) {
	var label []string

	// Generate Laplacian matrix from a Radius Neighbors Graph
	laplacian := make([][]float64, len(data))
	var wg sync.WaitGroup
	groups := 4
	wg.Add(groups)

	for i := 0; i < groups; i++ {
		go func(start int) {
			for j := start; j < (len(data)/groups)+start; j++ {
				laplacian[j] = cluster.Euclidean(data[j], data)
				sum := 0.0
				for k := range laplacian[j] {
					if laplacian[j][k] > radius {
						laplacian[j][k] = 0.0
					}
					sum += laplacian[j][k]
					laplacian[j][k] *= -1
				}
				laplacian[j][j] = sum
			}
			wg.Done()
		}(i * len(data) / groups)
	}
	wg.Wait()

	// Generate eigenvectors from laplacian matrix
	_, eigen_vectors, err := cluster.EigenSym(laplacian)
	if err != nil {
		return label, err
	}

	// Determine the cluster for each data point
	// NOTE: No need to sort via eigenvalues since EigenSym() returns them sorted in ascending order
	for i := range data {
		if eigen_vectors[i][1] >= 0 {
			label = append(label, "Cluster One")
		} else {
			label = append(label, "Cluster Two")
		}
	}
	return label, nil
}
