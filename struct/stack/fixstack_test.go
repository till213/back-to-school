package stack

import (
	"testing"
)

func TestStack_Empty(t *testing.T) {
	var s FixedLengthStack

	s.Create()
	if !s.Empty() {
		t.Error("Heap is not empty")
	}
	s.Push(42)
	if s.Empty() {
		t.Error("Heap is empty")
	}
	s.Pop()
	if !s.Empty() {
		t.Error("Heap is not empty")
	}
}

func TestStack_Full(t *testing.T) {
	var s FixedLengthStack

	s.Create()
	if s.Full() {
		t.Error("Heap is full")
	}
	for i := 0; i < N; i++ {
		s.Push(i)
	}
	if !s.Full() {
		t.Error("Heap is not full")
	}
	for i := 0; i < N; i++ {
		s.Pop()
	}
	if s.Full() {
		t.Error("Heap is full")
	}
}

func TestStack_PushPop(t *testing.T) {
	const val1, val2 = 42, 1001
	var s FixedLengthStack

	s.Create()
	s.Push(val1)
	if s.Top() != val1 {
		t.Errorf("Top is not %v", val1)
	}
	s.Push(val2)
	if s.Top() != val2 {
		t.Errorf("Top is not %v", val2)
	}
	s.Pop()
	if s.Top() != val1 {
		t.Errorf("Top is not %v", val1)
	}
	s.Pop()
	if !s.Empty() {
		t.Error("Heap is not empty")
	}
}
