package main

import (
	"fmt"
	"reflect"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Returns the length of the longest connected 1 bit sequence
// when just one bit in m can be flipped to 1
func bitFlipToWin(m int) int {

	if m == ^0 {
		// All 1s, which is already the longest sequence
		return int(reflect.TypeOf(m).Size()) * 8
	}

	currentLength := 0
	previousLength := 0
	// We can always flip one bit
	maxLength := 1
	// we insect on bit after another, starting with the LSB
	for m > 0 {
		if m&1 == 1 {
			currentLength++
		} else {
			// Check if there are one or (at least) two 0s
			if m&2 == 2 {
				// we have a 01 (just one zero, which we could flip),
				// so remember the current length
				previousLength = currentLength
			} else {
				// we have a 00 (plus possibly more 0s), discard
				// current length (we cannot connect those 1 segments)
				previousLength = 0
			}
			currentLength = 0
		}

		// Also count the bit that we can flip -> +1
		maxLength = max(maxLength, currentLength+previousLength+1)

		// inspect next bit
		m >>= 1
	}
	return maxLength
}

func main() {

	m := ^0
	size := bitFlipToWin(m)
	fmt.Println("Length:", size)

	// 8 1s and 7 1s, 6 1s and 5 1s
	m = 0b111111110111111100111111011111
	size = bitFlipToWin(m)
	fmt.Println("Length:", size)

	m = 0
	size = bitFlipToWin(m)
	fmt.Println("Length:", size)

}
