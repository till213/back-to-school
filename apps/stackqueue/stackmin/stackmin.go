package main

import (
	"fmt"
)

// Interface defines the expected functions for a stack value
type Interface interface {
	// Less reports whether this value is less than the other
	Less(other Interface) bool
}

// MinStack implements a stack with a minimum operation
type MinStack struct {
	data []Interface
	// Stores the index of the minimum element for each
	// stack depth
	mins []int
}

// New creates a new Stack
func New() *MinStack {
	s := new(MinStack)
	s.data = make([]Interface, 0, 32)
	s.mins = make([]int, 0, 32)
	return s
}

// Push pushes the value onto the MinStack
func (s *MinStack) Push(value Interface) {
	l := len(s.data)
	if l > 0 {
		oldMin := s.Min()
		if value.Less(oldMin) {
			s.mins = append(s.mins, len(s.mins))
		} else {
			s.mins = append(s.mins, s.mins[len(s.mins)-1])
		}
	} else {
		s.mins = append(s.mins, 0)
	}
	s.data = append(s.data, value)
}

// Pop removes the topmost value from MinStack
func (s *MinStack) Pop() Interface {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.mins = s.mins[:len(s.mins)-1]
	return val
}

// Min returns the minimum value in the stack
func (s *MinStack) Min() Interface {
	val := s.data[s.mins[len(s.mins)-1]]
	return val
}

type myint int

func (i myint) Less(other Interface) bool {
	return i < other.(myint)
}

func main() {
	s := New()
	s.Push(myint(42))
	fmt.Println("Min", s.Min().(myint))
	s.Push(myint(1001))
	s.Push(myint(1))
	fmt.Println("Min", s.Min().(myint))
	fmt.Println("Pop", s.Pop().(myint))
	fmt.Println("Min", s.Min().(myint))
	fmt.Println("Pop", s.Pop().(myint))
	fmt.Println("Min", s.Min().(myint))
	fmt.Println("Pop", s.Pop().(myint))
}
