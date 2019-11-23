// Package prioqueue provides a fixed-length priority queue using a heap.
package prioqueue

import "github.com/till213/back-to-school/struct/heap"

// PrioQueue structure
type PrioQueue struct {
	heap heap.Heap // Fixed-length heap
	n    int       // Current number of elements
}

// Init initialises a new priority queue
func (p *PrioQueue) Init() {
	p.n = 0
}

// Empty returns true if the priority queue is empty.
func (p *PrioQueue) Empty() bool {
	return p.n == 0
}

// Full returns true if the priority queue is full.
func (p *PrioQueue) Full() bool {
	return p.n == heap.M
}

// Insert inserts element elt
func (p *PrioQueue) Insert(elt int) {
	p.n++
	p.heap.H[p.n] = elt
	i := p.n
	for i > 1 && p.heap.H[i] < p.heap.H[i/2] {
		p.heap.H[i], p.heap.H[i/2] = p.heap.H[i/2], p.heap.H[i]
		i /= 2
	}
}

// Min returns the element with the highest priority. Not to be called
// when the priority queue is empty.
func (p *PrioQueue) Min() int {
	return p.heap.H[1]
}

// Delete deletes the element with the highest priority. Not to be called
// when the priority queue is empty.
func (p *PrioQueue) Delete() {
	p.heap.H[1] = p.heap.H[p.n]
	p.n--
	p.heap.Restore(1, p.n)
}
