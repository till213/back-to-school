package main

import "fmt"

// inserts the bits of m into n, starting at position
// i to j (counting from LSB)
func insert(n, m uint16, i, j int) uint16 {
	var mask, leftMask, rightMask uint16
	var res uint16

	allOnes := ^0
	leftMask = uint16(allOnes) << (j + 1)
	rightMask = (1 << i) - 1
	mask = leftMask | rightMask

	// clear values
	n &= mask

	// set values
	res = n | m<<i
	fmt.Printf("n: %016b, m: %016b, mask: %016b\n", n, m, mask)
	return res
}

func main() {
	n := uint16(0b1010101010101010)
	m := uint16(0b11111111)
	res := insert(n, m, 4, 12)
	fmt.Printf("Result: %b\n", res)
}
