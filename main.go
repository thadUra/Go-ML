package main

import (
	"Soccer-Penalty-Kick-ML-Threading/ml"
)

// var wg sync.WaitGroup
// var count int

// func addFull(start int, end int) {
// 	for i := start; i < end; i++ {
// 		count += 1
// 	}
// 	defer wg.Done()
// }

func main() {
	// count = 0
	// start := time.Now()
	ml.InitFCLayer(10, 10)

	// addFull(0, 10000000000)
	// wg.Add(2)
	// go addFull(0, 5000000000)
	// go addFull(5000000000, 10000000000)
	// wg.Wait()

	// elapsed := time.Since(start)
	// fmt.Printf("\n\n%s", elapsed)
}
