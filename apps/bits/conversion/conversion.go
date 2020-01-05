package main

import "fmt"

// Returns the population count for w, using Skip Zeros.
func skipZeros(x int) int {
	c := 0
	for x != 0 {
		c++
		// Good with sparse count
		// x : 11000
		// -1: 10111
		//  &: 10000
		x = x & (x - 1)
	}
	return c
}

func nofBitsToFlip(a, b int) int {
	// XOR sets the bits which are different in a and b
	diffBits := a ^ b
	// Now we simply need to count those bits
	return skipZeros(diffBits)
}

func main() {
	a := 0b01010101
	b := 0b10101010
	n := nofBitsToFlip(a, b)
	fmt.Printf("Number of bits to flip for conversion of %016b to %016b: %d\n", a, b, n)
}
