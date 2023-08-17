# Go-ML
[![License: MIT](https://img.shields.io/badge/license-MIT-orange.svg)](http://www.gnu.org/licenses/gpl-3.0)
[![Documentation](https://img.shields.io/badge/documentation-GoDoc-blue.svg)](https://pkg.go.dev/github.com/thadUra/Go-ML)
[![stability-unstable](https://img.shields.io/badge/stability-unstable-yellow.svg)](https://github.com/emersion/stability-badges#unstable)

Go Machine Learning is a library aimed at providing ML functionality and capability to Go. Several packages provide usage with neural networks, clustering algorithms, reinforcement learning, and a custom dataframe type.

## Installation
The library currently contains only pure Go. Installation can be done using `go get`.
```
go get -u github.com/thadUra/Go-ML
```

Importing the library into your Go files just requires an import statement.
```
import "github.com/thadUra/Go-ML/"
```

## Roadmap
| Tasks To Do                                 | Current Status | Finished | 
|---------------------------------------------|----------------|----------|
| Initial Documentation for all packages      | Completed      | &check;  |
| Add benchmark tests against Python          | Completed      | &check;  |
| Implement custom dataframe type             | Completed      | &check;  |
| Implement better error handling (log.fatal) | Completed      | &check;  |
| Change visibility of struct vars and funcs  | Completed      | &check;  |
| Optimize memory in implementation           | In Progress    | &cross;  |
| Utilize goroutines to optimize runtime      | In Progress    | &cross;  |
| Make custom soccer env more deterministic   | Not Started    | &cross;  |
| Complete convolutional and flat layer       | Not Started    | &cross;  |
| Add more activation layer functions         | Not Started    | &cross;  |
| ...                                         | ...            | ...      |

## Benchmarks
To view the benefits of using Go over a popular language such as Python for machine learning, benchmark tests were made comparing the GoML package to its Python counterpart. The benchmark script can be found at `./tests/benchmark/benchmark.sh`. Runtime comparisons are currently made for Dataframe operations on the Iris Dataset, Principal Component Analysis, K-Means clustering, and a simple Neural Network.

This benchmark run was made on a MacBook Pro wiht a 2Ghz Quad-Core Intel Core i5 processor.
```
    systemDir benchmark % ./benchmark.sh
    === GO BENCHMARK ===
    Dataframe: 0.000479472 seconds
    K-Means: 0.000426072 seconds
    ...
    NN: 0.025185177 seconds
    PCA: 9.6088e-05 seconds
    PASS
    ok      github.com/thadUra/Go-ML/tests/benchmark        0.220s
    === END GO BENCHMARK ===

    === PY BENCHMARK ===
    Dataframe: 0.004310131072998047 seconds
    K-Means: 0.00932002067565918 seconds
    ...
    NN: 1.6983931064605713 seconds
    PCA: 0.002899169921875 seconds
```

As seen, Golang performs each function at much faster speed compared to Python, with some descrepancies with the Neural Network test. The runtime differences should make Go worth considering for machine learning computation in the future, as long as stable packages and support are present.

### Notes
This library is actively being developed. If you would like to contribute to this library, please feel free to make any requests.