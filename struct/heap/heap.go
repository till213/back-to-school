// Package heap provides a fixed-length heap.
package heap

// M is the lenght of heap
const M = 32

// Heap structure
type Heap struct {
	// Fist element shall have index 1 (not 0)
	H [M + 1]int // Fixed-length heap
}

// Restore restores the heap properties from L to R, were
// L <= R and L, R in [1, M]
// H[i] <= H[2 * i] AND H[i] <= H[2 * i + 1]
func (heap *Heap) Restore(L, R int) {
	var i, j int

	i = L
	for i <= (R / 2) {
		if (2*i < R) && (heap.H[2*i+1] < heap.H[2*i]) {
			j = 2*i + 1
		} else {
			j = 2 * i
		}
		if heap.H[j] < heap.H[i] {
			heap.H[i], heap.H[j] = heap.H[j], heap.H[i]
			i = j
		} else {
			i = R
		}
	}
}

// Create will create the heap properties for the entire array Heap.H
func (heap *Heap) Create() {
	for i := M / 2; i > 0; i-- {
		heap.Restore(i, M)
	}
}
