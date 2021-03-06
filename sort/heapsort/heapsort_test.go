package main

import (
	"fmt"
	"testing"

	"github.com/till213/back-to-school/data/array"
	"github.com/till213/back-to-school/struct/tree/heap/binaryheap"
)

func TestHeapSortRandom(t *testing.T) {
	var h binaryheap.BinaryHeap

	h.H = [binaryheap.M + 1]int{99999, 5, 10, 13, 1, 8, 7, 19, 4, 3, 6, 2, 9}

	fmt.Println("Unsorted", h.H)
	HeapSort(&h)
	fmt.Println("Sorted", h.H)
	if !array.IsSorted(h.H[:], false) {
		t.Errorf("Not sorted")
	}
}
