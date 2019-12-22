package fifo

import (
	"testing"
)

func TestCircularFIFOEmpty(t *testing.T) {
	var s CircularFIFO

	s.Create()
	if !s.Empty() {
		t.Error("FIFO is not empty")
	}
	s.Enqueue(42)
	if s.Empty() {
		t.Error("FIFO is empty")
	}
	s.Dequeue()
	if !s.Empty() {
		t.Error("FIFO is not empty")
	}
}

func TestCircularFIFOFull(t *testing.T) {
	var s CircularFIFO

	s.Create()
	if s.Full() {
		t.Error("FIFO is full")
	}
	for i := 0; i < N-1; i++ {
		s.Enqueue(i)
	}
	if !s.Full() {
		t.Error("FIFO is not full")
	}
	for i := 0; i < N-1; i++ {
		s.Dequeue()
	}
	if s.Full() {
		t.Error("FIFO is full")
	}
}

func TestCircularFIFOEnqueueDequeue(t *testing.T) {
	const val1, val2 = 42, 1001
	var s CircularFIFO

	s.Create()
	s.Enqueue(val1)
	if s.Front() != val1 {
		t.Errorf("Top is not %v", val1)
	}
	s.Enqueue(val2)
	if s.Front() != val1 {
		t.Errorf("Top is not %v", val1)
	}
	s.Dequeue()
	if s.Front() != val2 {
		t.Errorf("Top is not %v", val2)
	}
	s.Dequeue()
	if !s.Empty() {
		t.Error("FIFO is not empty")
	}
}
