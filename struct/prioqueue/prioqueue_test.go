package prioqueue

import (
	"fmt"
	"testing"

	"github.com/till213/back-to-school/data/array"
	"github.com/till213/back-to-school/struct/heap"
)

func TestPrioQueue(t *testing.T) {

	const first = 42
	const second = 21
	var pq PrioQueue

	pq.Init()
	if !pq.Empty() {
		t.Errorf("Prio Queue should be empty")
	}
	if pq.Full() {
		t.Errorf("Prio Queue should not be full")
	}

	pq.Insert(first)
	if pq.Empty() {
		t.Errorf("Prio Queue should not be empty")
	}
	elt := pq.Min()
	if elt != first {
		t.Errorf("First prio element was %d, should be: %d", elt, first)
	}

	pq.Insert(second)
	elt = pq.Min()
	if elt != second {
		t.Errorf("First prio element was %d, should be: %d", elt, second)
	}

	pq.Delete()
	elt = pq.Min()
	if elt != first {
		t.Errorf("First prio element was %d, should be: %d", elt, first)
	}

	pq.Delete()
	if !pq.Empty() {
		t.Errorf("Prio Queue should be empty")
	}

}

func TestPrioQueueReverse(t *testing.T) {
	var pq PrioQueue
	a := array.Sorted(heap.M, false)

	for _, elt := range a {
		pq.Insert(elt)
	}

	fmt.Println("Priority Queue heap: ", pq.heap.H)
	if !pq.Full() {
		t.Errorf("Prio Queue should be full")
	}

	for i := 0; i < heap.M; i++ {
		elt := pq.Min()
		if elt != i {
			t.Errorf("First prio element was %d, should be: %d", elt, i)
		}
		pq.Delete()
	}
}
