package main

import "fmt"

const tail = 0
const front = 1

// StackQueue implements a queue using two stacks
type StackQueue struct {
	stacks [2][]int
}

// New creates a new StackQueue
func New() *StackQueue {
	s := new(StackQueue)
	s.stacks[tail] = make([]int, 0, 32)
	s.stacks[front] = make([]int, 0, 32)
	return s
}

// Reorders the elements in the tail stack by
// pushing them onto the front stack (popping from tail)
// Only to be called when the front stack is
// empty, in order to preserve the correct FIFO order.
func (q *StackQueue) tailToFront() {
	for n := len(q.stacks[tail]) - 1; n >= 0; n-- {
		// Pop from tail stack
		value := q.stacks[tail][n]
		q.stacks[tail] = q.stacks[tail][:n]
		// Push onto front stack
		q.stacks[front] = append(q.stacks[front], value)
	}
}

// Enqueue adds the value at the end of the queue
func (q *StackQueue) Enqueue(value int) {
	q.stacks[tail] = append(q.stacks[tail], value)
}

// Dequeue removes and returns the value at the front
// of the queue
func (q *StackQueue) Dequeue() int {
	if len(q.stacks[front]) == 0 {
		q.tailToFront()
	}
	// Pop from front stack
	n := len(q.stacks[front]) - 1
	value := q.stacks[front][n]
	q.stacks[front] = q.stacks[front][:n]
	return value
}

func main() {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println("Dequeue", q.Dequeue())
	q.Enqueue(3)
	fmt.Println("Dequeue", q.Dequeue())
	fmt.Println("Dequeue", q.Dequeue())
}
