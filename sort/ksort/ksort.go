package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
// https://golang.org/pkg/container/heap/
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// https://www.techiedelight.com/sort-k-sorted-array/
func sortKSortedArray(arr []int, k int) {
	minHeap := &IntHeap{}

	n := len(arr)
	// Insert the first k+1 elements into the heap
	for i := 0; i <= k; i++ {
		heap.Push(minHeap, arr[i])
	}

	// index into the original array 'arr'
	index := 0

	// do for remaining elements of the array
	for i := k + 1; i < n; i++ {
		// pop top element from min-heap and assign it to
		// next available array index
		arr[index] = heap.Pop(minHeap).(int)
		index++
		// push next array element into min-heap
		heap.Push(minHeap, arr[i])
	}

	// pop all remaining elements from the min heap and assign
    // it to next available array index
    for len(*minHeap) > 0 {
		arr[index] = heap.Pop(minHeap).(int);
		index++
    }
}

func main() {
	const k = 2
	arr := []int{1, 4, 5, 2, 3, 7, 8, 6, 10, 9}
	sortKSortedArray(arr, k)
	fmt.Println(arr)
}
