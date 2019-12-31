package main

import (
	"fmt"

	"github.com/till213/back-to-school/struct/tree/heap/binaryheap"
)

// HeapSort sorts the array h in decreasing order by
// first establishing a "min-heap" order and then
// taking the minimum heap element, swapping it with
// the "end" element, restoring the heap order of the
// previous elements each time
// Stable sort, O(n log n)
func HeapSort(h *binaryheap.BinaryHeap) {
	// Create heap
	for i := binaryheap.M / 2; i > 0; i-- {
		h.Restore(i, binaryheap.M)
	}
	// Shift-up: elements are extracted from heap in
	// increasing order
	for i := binaryheap.M; i > 1; i-- {
		h.H[i], h.H[1] = h.H[1], h.H[i]
		h.Restore(1, i-1)
	}
}

func main() {
	var h binaryheap.BinaryHeap

	h.H = [binaryheap.M + 1]int{99999, 5, 10, 13, 1, 8, 7, 19, 4, 3, 6, 2, 9}
	fmt.Println("\n--- Unsorted --- \n", h.H)
	HeapSort(&h)
	fmt.Println("\n--- Sorted --- \n", h.H)
}
