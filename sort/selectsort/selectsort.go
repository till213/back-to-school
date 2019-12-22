package main

import "fmt"

// SelectionSort sorts the array arr. Unstable sort. O(n^2)
// https://en.wikipedia.org/wiki/Selection_sort
func SelectionSort(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		minKey := arr[i]
		minIndex := i
		// select minimum key in unsorted partition
		for j := i + 1; j < l; j++ {
			if arr[j] < minKey {
				minKey = arr[j]
				minIndex = j
			}
		}
		// swap (unstable) - we can make this a stable
		// sort when instead of swapping we *insert* the
		// minimum to arr[i] and *move* all subsequent
		// entries "one up". Leads to erforming O(n^2)
		// writes, or use linked list which supports O(1)
		// insertion
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	array := [10]int{4, 2, 3, 1, 6, 9, 8, 7, 0, 5}
	fmt.Println("\n--- Unsorted --- \n", array)
	SelectionSort(array[:])
	fmt.Println("\n--- Sorted --- \n", array)
}
