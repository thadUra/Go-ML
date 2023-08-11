package cluster

import (
	"errors"
	"sort"

	"gonum.org/v1/gonum/mat"
)

// PCA performs principal component analysis on `data` and reduces dimensionality to `n`
// dimensions. It assumes that the variable data is arranged via columns. It returns the
// reduced data. It returns an error if data dimensions are inconsistent/empty or if
// there are more variables than n.
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

	// Subtract mean from each col to standardize
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
	eigen_vals, eigen_vects, err := EigenSym(cov_matrix)
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
	mean_matrix := float64TwoDimenToDense(x_mean)
	subset_matrix := float64TwoDimenToDense(subset)
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
