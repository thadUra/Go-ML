package cluster

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/mat"
)

// Euclidean returns all the euclidean distances from point to each coordinate in data
func Euclidean(point []float64, data [][]float64) []float64 {
	result := make([]float64, len(data))
	for i := range data {
		result[i] = math.Sqrt(math.Pow(point[0]-data[i][0], 2) + math.Pow(point[1]-data[i][1], 2))
	}
	return result
}

// Covariance returns the covariance matrix of data. If `rowvar` is true, the data is
// treated with the variables being arranged via rows. It returns an error if data
// dimensions are inconsistent/empty.
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

// EigenSym returns the eigenvalues and eigenvectors of a real symmetric matrix. It returns
// an error if data dimensions are inconsistent/empty or if it fails to factorize
func EigenSym(data [][]float64) ([]float64, [][]float64, error) {
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
	ok := dense.Factorize(float64TwoDimenToSymDense(data), true)
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

// Float64TwoDimenToDense returns a matrix (mat.Dense) of `data`.
func float64TwoDimenToDense(data [][]float64) *mat.Dense {
	rows, cols := len(data), len(data[0])
	data_1d := make([]float64, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data_1d[i*cols+j] = data[i][j]
		}
	}
	return mat.NewDense(rows, cols, data_1d)
}

// Float64TwoDimenToSymDense returns a symmetrical matrix (mat.SymDense) of `data`.
func float64TwoDimenToSymDense(data [][]float64) *mat.SymDense {
	rows, cols := len(data), len(data[0])
	data_1d := make([]float64, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data_1d[i*cols+j] = data[i][j]
		}
	}
	return mat.NewSymDense(rows, data_1d)
}
