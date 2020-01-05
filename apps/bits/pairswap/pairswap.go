package main

import "fmt"

// Swaps the odd/even bits of m pairwise
// and returns the new value
func pairwiseSwap(m uint16) uint16 {
	var oddMask, evenMask uint16
	oddMask = 0b0101010101010101
	evenMask = 0b1010101010101010
	swapped := (m&evenMask)>>1 | (m&oddMask)<<1
	return swapped
}

func main() {
	var m uint16

	m = 0b1010101010101010
	n := pairwiseSwap(m)

	fmt.Printf("m: %016b\nn: %016b\n", m, n)
}
