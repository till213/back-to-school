package main

import (
	"fmt"
)

// InsertionSort sorts the array arr. Stable sort. O(n^2)
// https://en.wikipedia.org/wiki/Insertion_sort
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for ; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

func main() {
	array := [10]int{4, 2, 3, 1, 6, 9, 8, 7, 0, 5}
	fmt.Println("\n--- Unsorted --- \n", array)
	InsertionSort(array[:])
	fmt.Println("\n--- Sorted --- \n", array)
}
