package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func sortStack(s *stack.Stack) {

	aux := stack.New()
	l := s.Len()
	for i := 0; i < l; i++ {

		// Move all the upper i values onto the
		// aux stack, determine maximum value
		max := s.Peek().(int)
		for j := i; j < l; j++ {
			val := s.Pop().(int)
			if val > max {
				max = val
			}
			aux.Push(val)
		}
		// post: all len-i upper values are now in
		//       aux stack, their max is known,
		//       the bottom i values are sorted
		//       (max is at bottom of stack)

		// First move max value back onto original
		// stack s
		s.Push(max)
		// Push the remaining values from aux onto s
		// except max
		for aux.Len() > 0 {
			val := aux.Pop()
			if val != max {
				s.Push(val)
			}
		}
	}
}

func main() {
	s := stack.New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for s.Len() > 0 {
		fmt.Println("Unsorted pop", s.Pop())
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	sortStack(s)
	for s.Len() > 0 {
		fmt.Println("Sorted pop", s.Pop())
	}
}
