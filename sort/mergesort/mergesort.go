package main

import "fmt"

func combine(arr []int, l, r, m int) {
	var arrB []int
	arrB = make([]int, len(arr))
	i, j := l, m+1
	for k := l; k <= r; k++ {
		if i > m || j <= r && arr[j] < arr[i] {
			arrB[k] = arr[j]
			j++
		} else {
			arrB[k] = arr[i]
			i++
		}
	}
	for k := l; k <= r; k++ {
		arr[k] = arrB[k]
	}
}

// MergeSort sorts the array arr. Stable sort. O(n log(n)). Auxiliary space O(n)
// Suitable for sorting "external memory" (where "disk IO" is expensive)
// https://en.wikipedia.org/wiki/Merge_sort
func MergeSort(arr []int, l, r int) {
	if l < r {
		m := (l + r) / 2
		MergeSort(arr, l, m)
		MergeSort(arr, m+1, r)
		combine(arr, l, r, m)
	}
}

func main() {
	array := [10]int{4, 2, 3, 1, 6, 9, 8, 7, 0, 5}
	fmt.Println("\n--- Unsorted --- \n", array)
	MergeSort(array[:], 0, len(array)-1)
	fmt.Println("\n--- Sorted --- \n", array)
}
