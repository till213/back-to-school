package fifo

type cell struct {
	val int
	next *cell
}

// ListFIFO implements a FIFO using a single-linked list
type ListFIFO struct {
	head, tail *cell
}

// Create initialises a new FIFO
func (f *ListFIFO) Create() {
	// allocate sentinel cell
	f.head = new(cell)
	f.tail = f.head
}

// Empty returns true if the FIFO is empty
func (f *ListFIFO) Empty() bool {
	return f.head == f.tail
}

// Enqueue adds the value val to the FIFO
func (f *ListFIFO) Enqueue(val int) {
	f.tail.val = val
	f.tail.next = new(cell)
	f.tail = f.tail.next
}

// Front returns the front element
func (f *ListFIFO) Front() int {
	return f.head.val
}

// Dequeue removes the value val from the FIFO
func (f *ListFIFO) Dequeue() {
	f.head = f.head.next
}