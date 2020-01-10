package main

import (
	"container/heap"
	"fmt"
)

// Heap implements basic heap operations: Len, Swap, Push, Pop
type Heap struct {
	heap *[]int
}

func (h Heap) Len() int           { return len(*(h.heap)) }
func (h Heap) Swap(i, j int)      { (*h.heap)[i], (*h.heap)[j] = (*h.heap)[j], (*h.heap)[i] }
func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*(h.heap) = append(*(h.heap), x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *(h.heap)
	n := len(old)
	x := old[n-1]
	*(h.heap) = old[0 : n-1]
	return x
}

// A MaxHeap is a max-heap of ints.
type MaxHeap struct {
	Heap
}

// A MinHeap is a max-heap of ints.
type MinHeap struct {
	Heap
}

func (h MaxHeap) Less(i, j int) bool { return (*h.heap)[i] > (*h.heap)[j] }
func (h MinHeap) Less(i, j int) bool { return (*h.heap)[i] < (*h.heap)[j] }

var minHeap *MinHeap
var maxHeap *MaxHeap

func addNumber(randomNumber int) {
	// Invariant: len(maxHeap) >= len(minHeap), maximum one element larger
	if len(*maxHeap.heap) == len(*minHeap.heap) {
		// existing heaps have same size (which may be 0)
		if len(*minHeap.heap) > 0 && randomNumber > (*minHeap.heap)[0] {
			// Number is larger than the minimum
			// First move the minimum element from the min-heap to the max-heap...
			heap.Push(maxHeap, heap.Pop(minHeap))
			// ... and then add the new number into the min-heap
			heap.Push(minHeap, randomNumber)
		} else {
			// Number is smaller than the minimum
			// Add it to the max heap
			heap.Push(maxHeap, randomNumber)
		}
	} else {
		if randomNumber < (*maxHeap.heap)[0] {
			heap.Push(minHeap, heap.Pop(maxHeap))
			heap.Push(maxHeap, randomNumber)
		} else {
			heap.Push(minHeap, randomNumber)
		}
	}
}

func median() float64 {
	if len(*maxHeap.heap) == 0 {
		return 0.0
	}
	if len(*maxHeap.heap) == len(*minHeap.heap) {
		min := (*minHeap.heap)[0]
		max := (*maxHeap.heap)[0]
		return float64(min + max) / 2.0
	} else {
		return float64((*maxHeap.heap)[0])
	}
}

func main() {
	minHeap = new(MinHeap)
	maxHeap = new(MaxHeap)
	minHeap.heap =  &[]int{}
	maxHeap.heap =  &[]int{}

	addNumber(2)
	fmt.Printf("Median: %.2f\n", median())
	addNumber(4)
	fmt.Printf("Median: %.2f\n", median())
	addNumber(6)
	fmt.Printf("Median: %.2f\n", median())
	addNumber(8)
	fmt.Printf("Median: %.2f\n", median())
	addNumber(1)
	fmt.Printf("Median: %.2f\n", median())
	addNumber(3)
	fmt.Printf("Median: %.2f\n", median())
}
