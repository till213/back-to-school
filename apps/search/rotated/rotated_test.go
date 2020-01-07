package main

import (
	"testing"
)

func rotate(arr []int, rightShift int) []int {
	left := arr[0:len(arr)-rightShift]
	right := arr[len(arr)-rightShift:len(arr)]
	return append(right, left...)
}

func TestRotatedRecursive(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for rightShift := 0; rightShift < len(arr); rightShift++ {
		rotated := rotate(arr, rightShift)
		for i, v := range arr {
			index := recursiveSearch(rotated, v)
			exp := (i + rightShift) % len(arr)
			if index != exp {
				t.Errorf("Expected index %d of value %d of rotated array with right shift %d, received: %d (recursive)", exp, v, rightShift, index)
			}
		}
	}
}

func TestRotatedIterative(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for rightShift := 0; rightShift < len(arr); rightShift++ {
		rotated := rotate(arr, rightShift)
		for i, v := range arr {
			index := iterativeSearch(rotated, v)
			exp := (i + rightShift) % len(arr)
			if index != exp {
				t.Errorf("Expected index %d of value %d of rotated array with right shift %d, received: %d (iterative)", exp, v, rightShift, index)
			}
		}
	}
}
