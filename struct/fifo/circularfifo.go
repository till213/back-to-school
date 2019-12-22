package fifo

// N defines the capacity of the FIFO, which is in fact N-1 elements
// (N minus one sentinel element)
const N = 32

// CircularFIFO implements a circular first-in, first-out
// structure, using a sentinel value: one cell is always
// left unused ("free"): we do not need an additional
// field which keeps track of the n stored elements, and
// hence two concurrent processes may queue and dequeue
// concurrently, without the need for additional locks
type CircularFIFO struct {
	fifo    [N]int
	in, out int
}

// Create initialises a new FIFO
func (f *CircularFIFO) Create() {
	f.in, f.out = 0, 0
}

// Empty returns true if the FIFO is empty
func (f *CircularFIFO) Empty() bool {
	return f.in == f.out
}

// Full returns true if the FIFO is full
func (f *CircularFIFO) Full() bool {
	return (f.in+1)%N == f.out
}

// Enqueue adds the value val to the FIFO
func (f *CircularFIFO) Enqueue(val int) {
	f.fifo[f.in] = val
	f.in = (f.in + 1) % N
}

// Front returns the front element
func (f *CircularFIFO) Front() int {
	return f.fifo[f.out]
}

// Dequeue removes the value val from the FIFO
func (f *CircularFIFO) Dequeue() {
	f.out = (f.out + 1) % N
}
