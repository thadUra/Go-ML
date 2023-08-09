package cluster

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
)

/**
 * KMeans struct
 * Stores centroids and number of clusters in model
 */
type KMeans struct {
	CENTROIDS    [][]float64
	CLUSTERS     int
	MAX_EPISODES int
}

/**
 * InitKMeans()
 * Initializes KMeans model with n clusters and max_eps iterations
 */
func InitKMeans(n, max_eps int) *KMeans {
	var model KMeans
	model.CLUSTERS = n
	model.MAX_EPISODES = max_eps
	return &model
}

/**
 * Train()
 * Generates KMeans model with training data using KMeans++ algorithm
 * Throws error if data dimensions are empty/not consistent or fail in centroid calculations
 */
func (model *KMeans) Train(data [][]float64) error {
	// Check data dimensions
	if len(data) < 1 {
		return errors.New("KMeans.Train(): data cannnot be empty")
	}

	// Initialize first centroid with random point
	init_idx := int(rand.Float64() * float64(len(data)))
	model.CENTROIDS = append(model.CENTROIDS, data[init_idx])

	// Initialize other centroids with probabilities proportional to their distances to the first centroid
	for i := 1; i < model.CLUSTERS; i++ {
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
	for iter := 0; iter < model.MAX_EPISODES; iter++ {
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

/**
 * Evaluate()
 * Evaluates data with trained model to perform cluster analysis
 * Returns analysis with labels for plotting
 * Throws error if model not trained
 */
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

/**
 * Euclidean()
 * Grabs all the euclidean distances between a point and set of points
 */
func Euclidean(point []float64, data [][]float64) []float64 {
	result := make([]float64, len(data))
	for i := range data {
		result[i] = math.Sqrt(math.Pow(point[0]-data[i][0], 2) + math.Pow(point[1]-data[i][1], 2))
	}
	return result
}
