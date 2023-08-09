package cluster

import (
	"errors"
	"math"
	"sort"

	"gonum.org/v1/gonum/mat"
)

/**
 * PCA()
 * Performs principal component analysis on given data
 * Reduces dimensionality to n components
 * Assumes data is organized via column
 */
func PCA(data [][]float64, n int) ([][]float64, error) {
	// Check data dimensions
	if len(data) == 0 {
		return data, errors.New("PCA(): Data cannot have zero cols")
	}
	if len(data[0]) == 0 {
		return data, errors.New("PCA(): Data cannot have zero rows")
	}
	rows, cols := len(data), len(data[0])
	for i := range data {
		if len(data[i]) != cols {
			return data, errors.New("PCA(): Data cannot have mismatched column dimens")
		}
	}
	if cols < n {
		return data, errors.New("PCA(): Cannot reduce dimensions as cols is less than n")
	}

	// Subtract mean from each col
	x_mean := data
	for i := 0; i < cols; i++ {
		mean := 0.0
		for j := 0; j < rows; j++ {
			mean += (x_mean[j][i] / float64(rows))
		}
		for j := 0; j < rows; j++ {
			x_mean[j][i] -= mean
		}
	}

	// Calculate covariance matrix
	cov_matrix, err := Covariance(x_mean, false)
	if err != nil {
		return data, err
	}

	// Compute eigenvalues and eigenvectors
	eigen_vals, eigen_vects, err := Eigen(cov_matrix)
	if err != nil {
		return data, err
	}

	// Sort eigenvalues/vectors in descending order
	sorted_eigenvals := make([]float64, len(eigen_vals))
	copy(sorted_eigenvals, eigen_vals)
	sort.Sort(sort.Reverse(sort.Float64Slice(sorted_eigenvals)))
	sorted_indices := make([]int, len(eigen_vals))
	for i := 0; i < len(eigen_vals); i++ {
		for j := 0; j < len(eigen_vals); j++ {
			if sorted_eigenvals[j] == eigen_vals[i] {
				sorted_indices[i] = j
				break
			}
		}
	}
	var sorted_eigenvects [][]float64
	for i := 0; i < len(eigen_vects); i++ {
		temp := make([]float64, len(eigen_vects[i]))
		for j := 0; j < len(eigen_vects[i]); j++ {
			temp[sorted_indices[j]] = eigen_vects[i][j]
		}
		sorted_eigenvects = append(sorted_eigenvects, temp)
	}

	// Select subset from eigenvalue matrix
	var subset [][]float64
	for i := 0; i < len(sorted_eigenvects); i++ {
		temp := make([]float64, n)
		for j := 0; j < n; j++ {
			temp[j] = sorted_eigenvects[i][j]
		}
		subset = append(subset, temp)
	}

	// Transform data
	var reduced mat.Dense
	mean_matrix := Float64TwoDimenToDense(x_mean)
	subset_matrix := Float64TwoDimenToDense(subset)
	reduced.Mul(subset_matrix.T(), mean_matrix.T())

	// Convert back to 2D float64 slice to return while transposing
	r, c := reduced.Dims()
	var pca [][]float64
	for i := 0; i < c; i++ {
		temp := make([]float64, r)
		for j := 0; j < r; j++ {
			temp[j] = reduced.At(j, i)
		}
		pca = append(pca, temp)
	}
	return pca, nil
}

/**
 * Covariance()
 * Helper function utilized for PCA
 * Returns covariance matrix of data
 * Rowvar utilized if variable data is arranged by row
 */
func Covariance(data [][]float64, rowvar bool) ([][]float64, error) {
	// Check data dimensions
	if len(data) == 0 {
		return data, errors.New("Covariance(): Data cannot have zero cols")
	}
	if len(data[0]) == 0 {
		return data, errors.New("Covariance(): Data cannot have zero rows")
	}
	rows, cols := len(data), len(data[0])
	for i := range data {
		if len(data[i]) != cols {
			return data, errors.New("Covariance(): Data cannot have mismatched column dimens")
		}
	}

	// Flip row and column if variable data is arranged via row
	if rowvar {
		temp := rows
		rows = cols
		cols = temp
	}

	// Get mean of each variable
	mean := make([]float64, cols)
	for i := 0; i < cols; i++ {
		curr_mean := 0.0
		for j := 0; j < rows; j++ {
			curr_mean += (data[j][i] / float64(rows))
		}
		mean[i] = curr_mean
	}

	// Build covariance matrix starting with variance diagonal
	var cov [][]float64
	for i := 0; i < cols; i++ {
		temp := make([]float64, cols)
		for j := 0; j < cols; j++ {
			if j != i {
				temp[j] = 0
			} else {
				for k := 0; k < rows; k++ {
					temp[j] += math.Pow(data[k][j]-mean[j], 2)
				}
				temp[j] /= float64(rows - 1)
			}
		}
		cov = append(cov, temp)
	}

	// Calculate covariance values in matrix
	for i := 0; i < cols; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < rows; k++ {
				cov[i][j] += (data[k][j] - mean[j]) * (data[k][i] - mean[i])
			}
			cov[i][j] /= float64(rows - 1)
			cov[j][i] = cov[i][j]
		}
	}
	return cov, nil
}

/**
 * Eigen()
 * Return the eigenvalues and eigenvectors of a real symmetric matrix
 */
func Eigen(data [][]float64) ([]float64, [][]float64, error) {
	// Check data dimensions
	if len(data) == 0 {
		return nil, nil, errors.New("Eigen(): Data cannot have zero cols")
	}
	if len(data[0]) == 0 {
		return nil, nil, errors.New("Eigen(): Data cannot have zero rows")
	}
	rows, cols := len(data), len(data[0])
	for i := range data {
		if len(data[i]) != cols {
			return nil, nil, errors.New("Eigen(): Data cannot have mismatched column dimens")
		}
	}

	// Init return vars
	var values []float64
	var vectors [][]float64

	// Set matrix vars for calculating eigenvalues/vectors
	var evectors mat.Dense
	var dense mat.EigenSym
	ok := dense.Factorize(Float64TwoDimenToSymDense(data), true)
	if !ok {
		return values, vectors, errors.New("Eigen(): Failed to factorize")
	}
	dense.VectorsTo(&evectors)

	// Set and return vals
	values = dense.Values(nil)
	for i := 0; i < rows; i++ {
		temp := make([]float64, cols)
		for j := 0; j < cols; j++ {
			temp[j] = evectors.At(i, j)
		}
		vectors = append(vectors, temp)
	}
	return values, vectors, nil
}

/**
 * Float64TwoDimenToDense()
 * Helper function utilized for PCA
 * Return mat.Dense instance of 2D float64 slice
 */
func Float64TwoDimenToDense(data [][]float64) *mat.Dense {
	rows, cols := len(data), len(data[0])
	data_1d := make([]float64, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data_1d[i*cols+j] = data[i][j]
		}
	}
	return mat.NewDense(rows, cols, data_1d)
}

/**
 * Float64TwoDimenToDense()
 * Helper function utilized for PCA
 * Return mat.SymDense instance of 2D float64 slice
 */
func Float64TwoDimenToSymDense(data [][]float64) *mat.SymDense {
	rows, cols := len(data), len(data[0])
	data_1d := make([]float64, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data_1d[i*cols+j] = data[i][j]
		}
	}
	return mat.NewSymDense(rows, data_1d)
}
