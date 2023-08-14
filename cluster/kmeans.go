package cluster

import (
	"errors"
	"math/rand"
	"strconv"
)

// KMeans represents a classification model utilizing K-Means clustering
type KMeans struct {
	CENTROIDS    [][]float64
	clusters     int
	max_episodes int
}

// NewKMeans creates a new K-Means model with `n` clusters. `Max_eps` represents the maximum
// amount of episodes or iterations the model will train on until convergence.
func NewKMeans(n, max_eps int) *KMeans {
	var model KMeans
	model.clusters = n
	model.max_episodes = max_eps
	return &model
}

// Train performs the K-Means clustering algorithm on `data` and generates a classification model.
// An error is returned if dimensions are empty or not consistent. As well, it returns an
// error if centroid calculations fail at any point.
func (model *KMeans) Train(data [][]float64) error {
	// Check data dimensions
	if len(data) < 1 {
		return errors.New("KMeans.Train(): data cannnot be empty")
	}

	// Initialize first centroid with random point
	init_idx := int(rand.Float64() * float64(len(data)))
	model.CENTROIDS = append(model.CENTROIDS, data[init_idx])

	// Initialize other centroids with probabilities proportional to their distances to the first centroid
	for i := 1; i < model.clusters; i++ {
		// Calculate dist from points to centroids
		dist := make([]float64, len(data))
		sum := 0.0
		for j := range model.CENTROIDS {
			tmp := Euclidean(model.CENTROIDS[j], data)
			for k := range tmp {
				dist[k] += tmp[k]
				sum += tmp[k]
			}
		}
		// Normalize distances
		for j := 0; j < len(dist); j++ {
			dist[j] /= sum
		}
		// Choose next centroid given probabilities
		p := rand.Float64()
		sum = 0.0
		found := false
		for j := 0; j < len(dist)-1; j++ {
			sum += dist[j]
			if p < sum {
				model.CENTROIDS = append(model.CENTROIDS, data[j])
				found = true
				break
			}
		}
		if !found {
			return errors.New("KMeans.Train(): failed to add centroid")
		}
	}

	// Adjust centroids over iterations until convergence or max_episodes
	for iter := 0; iter < model.max_episodes; iter++ {
		// Assign point to nearest centroid
		points := make(map[int][][]float64)
		for j := range data {
			// Get centroid idx
			dist := Euclidean(data[j], model.CENTROIDS)
			min := dist[0]
			min_idx := 0
			for k := range dist {
				if dist[k] < min {
					min = dist[k]
					min_idx = k
				}
			}
			// Store point in map
			points[min_idx] = append(points[min_idx], data[j])
		}

		// Store previous centroids
		previous := make([][]float64, len(model.CENTROIDS))
		for j := range model.CENTROIDS {
			previous[j] = make([]float64, len(model.CENTROIDS[j]))
			copy(previous[j], model.CENTROIDS[j])
		}

		// Reassign centroids as mean of its own points
		for j := range model.CENTROIDS {
			mean_x, mean_y := 0.0, 0.0
			for k := range points[j] {
				mean_x += points[j][k][0] / float64(len(points[j]))
				mean_y += points[j][k][1] / float64(len(points[j]))
			}
			if len(points[j]) == 0 {
				model.CENTROIDS[j] = []float64{previous[j][0], previous[j][1]}
			} else {
				model.CENTROIDS[j] = []float64{mean_x, mean_y}
			}
		}

		// Check if centroids converged
		converged := true
		for j := range model.CENTROIDS {
			if previous[j][0] != model.CENTROIDS[j][0] && previous[j][1] != model.CENTROIDS[j][1] {
				converged = false
				break
			}
		}
		if converged {
			break
		}
	}
	return nil
}

// Evaluate takes in `data` and classifies it according to the trained K-Means model. It
// returns the data along with the centroid points with its correlated labels. An error
// is returned if the model was not yet trained.
func (model *KMeans) Evaluate(data [][]float64) ([][]float64, []string, error) {
	var result [][]float64
	var label []string

	// Check if model has been fit
	if len(model.CENTROIDS) < 1 {
		return result, label, errors.New("KMeans.Evaluate(): model not yet trained")
	}

	// Add centroid points
	for i := range model.CENTROIDS {
		// Add label and point
		label = append(label, "Centroid Center")
		result = append(result, model.CENTROIDS[i])
	}

	// Add evalulated data points
	for i := range data {
		// Get centroid idx
		dist := Euclidean(data[i], model.CENTROIDS)
		min := dist[0]
		min_idx := 0
		for j := range dist {
			if dist[j] < min {
				min = dist[j]
				min_idx = j
			}
		}
		// Add label and point
		name := "Cluster " + strconv.Itoa(min_idx)
		label = append(label, name)
		result = append(result, data[i])
	}
	return result, label, nil
}
