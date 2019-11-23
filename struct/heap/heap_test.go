package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {

	var heap Heap

	heap.H = [M + 1]int{0, 19, 13, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Println("Array: ", heap.H)
	heap.Create()
	fmt.Println("Heap: ", heap.H)

	for i := 1; i <= M/2; i++ {
		if !(heap.H[i] <= heap.H[2*i]) {
			t.Errorf("Heap order violated: H[%d]=%d > H[%d]=%d", i, heap.H[i], 2*i, heap.H[2*i])
		}
	}

}
