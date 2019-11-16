package main

import (
	"fmt"
)

func guessMedian(array []int) int {
	// The most simple way to guesstimate the median
	m := len(array) / 2
	return array[m]
}

func partition(array []int) (int, int) {
	var m int

	m = guessMedian(array)
	i := 0
	j := len(array) - 1
	for i <= j {

		for array[i] < m {
			i++
		}
		for array[j] > m {
			j--
		}
		if i <= j {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		}

	}

	return i, j

}

func quickSort(array []int) {
	i, j := partition(array)
	if j > 0 {
		quickSort(array[0 : j+1])
	}
	if i < len(array) {
		quickSort(array[i:])
	}

}

func main() {
	array := [10]int{4, 2, 3, 1, 6, 9, 8, 7, 0, 5}
	quickSort(array[:])
	for _, n := range array {
		fmt.Printf("%d ", n)
	}
	fmt.Println("")
}
