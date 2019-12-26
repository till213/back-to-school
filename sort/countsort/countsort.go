package main

import "fmt"

// CountSort sorts arr - each member is assumed to be in
// range 1..k, for some integer k. If k=O(n) then sort
// runs in O(n). Stable.
// Additional array storage is required, to hold (intermediate)
// results B and C
func CountSort(arr, sorted []int, k int) {
	c := make([]int, k)

	// This is not required, as Go int has already a zero
	// value of 0 - but it is explicitly mentioned in the
	// algorithm given in Introduction to Algorithms, MIT Press
	for i := 0; i < k; i++ {
		c[i] = 0
	}
	for j := 0; j < len(arr); j++ {
		c[arr[j]] = c[arr[j]] + 1
	}
	// c[i] now contains the number of elements equal to i
	for i := 1; i < k; i++ {
		c[i] = c[i] + c[i-1]
	}
	// c[i] now contains the number of elements less than or equal to i
	for j := len(arr) - 1; j >= 0; j-- {
		sorted[c[arr[j]]-1] = arr[j]
		c[arr[j]] = c[arr[j]] - 1
	}
}

func main() {
	array := []int{3, 6, 4, 1, 3, 4, 1, 4}
	fmt.Println("\n--- Unsorted --- \n", array)
	sorted := make([]int, len(array))
	CountSort(array, sorted, 7)
	fmt.Println("\n--- Sorted --- \n", sorted)
}
