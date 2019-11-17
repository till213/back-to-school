package main

import (
	"fmt"
	"time"
)

func main() {
	// Give the CPU something to bite on
	const N = 1e11
	x := 42
	y := 1001

	start := time.Now()
	for i := 0; i < N; i++ {
		// Bitswap using XOR
		x = x ^ y
		y = x ^ y
		x = x ^ y
	}
	end := time.Now()
	fmt.Printf("Bitswap: elapsed: %v\n", end.Sub(start))

	start = time.Now()
	for i := 0; i < N; i++ {
		// The usual swapping
		x, y = y, x
	}
	end = time.Now()
	fmt.Printf("Goswap: elapsed: %v\n", end.Sub(start))
}
