package main

import (
	"fmt"

	"github.com/till213/back-to-school/struct/heap"
)

// HeapSort sorts the array h in decreasing order by
// first establishing a "min-heap" order and then
// taking the minimum heap element, swapping it with
// the "end" element, restoring the heap order of the
// previous elements each time
// Stable sort, O(n log n)
func HeapSort(h *heap.Heap) {
	// Create heap
	for i := heap.M / 2; i > 0; i-- {
		h.Restore(i, heap.M)
	}
	// Shift-up: elements are extracted from heap in
	// increasing order
	for i := heap.M; i > 1; i-- {
		h.H[i], h.H[1] = h.H[1], h.H[i]
		h.Restore(1, i-1)
	}
}

func main() {
	var h heap.Heap

	h.H = [heap.M + 1]int{0, 19, 13, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("\n--- Unsorted --- \n", h.H)
	HeapSort(&h)
	fmt.Println("\n--- Sorted --- \n", h.H)
}
