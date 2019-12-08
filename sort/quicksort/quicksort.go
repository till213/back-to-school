package main

import (
	"fmt"
)

func guessMedian(s []int) int {
	// The most simple way to guesstimate the median
	m := len(s) / 2
	return s[m]
}

func partition(s []int) (int, int) {
	var m int

	m = guessMedian(s)
	i := 0
	j := len(s) - 1
	for i <= j {

		for s[i] < m {
			i++
		}
		for s[j] > m {
			j--
		}
		if i <= j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}

	}

	return i, j

}

// QuickSort sorts slice s in ascending order. Unstable sort.
func QuickSort(s []int) {
	i, j := partition(s)
	if j > 0 {
		QuickSort(s[0 : j+1])
	}
	if i < len(s) {
		QuickSort(s[i:])
	}

}

func main() {
	array := [10]int{4, 2, 3, 1, 6, 9, 8, 7, 0, 5}
	fmt.Println("\n--- Unsorted --- \n", array)
	QuickSort(array[:])
	fmt.Println("\n--- Sorted --- \n", array)
}
