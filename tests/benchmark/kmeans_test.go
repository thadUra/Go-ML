package bench

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/thadUra/Go-ML/cluster"
)

func TestKMeans(t *testing.T) {
	// Open iris data file
	file, err := os.Open("../misc/iris_data.csv")
	if err != nil {
		t.Fatalf(`TestKMeans(): failed to open file -> "%s"`, err)
	}
	r := csv.NewReader(file)

	// Parse data and store into slices
	var data [][]float64
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf(`TestKMeans(): failed in file parsing -> "%s"`, err)
		}
		next := make([]float64, 4)
		for i := 0; i < 4; i++ {
			next[i], err = strconv.ParseFloat(record[i], 64)
			if err != nil {
				t.Fatalf(`TestKMeans(): failed in file parsing -> "%s"`, err)
			}
		}
		data = append(data, next)
	}
	file.Close()

	// Run PCA
	result, err := cluster.PCA(data, 2)
	if err != nil {
		t.Fatalf(`TestKMeans(): failed at building PCA -> "%s"`, err)
	}

	// Start timer
	start := time.Now()

	// Test KMeans
	km := cluster.NewKMeans(3, 500)
	err = km.Train(result)
	if err != nil {
		t.Fatalf(`TestKMeans(): failed at kmeans training -> "%s"`, err)
	}

	// End timer
	duration := time.Since(start)
	fmt.Printf("K-Means: %v seconds\n", duration.Seconds())
}
