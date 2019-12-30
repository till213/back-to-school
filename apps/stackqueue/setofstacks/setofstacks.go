package main

import "fmt"

const maxSize = 2

// SetOfStacks implements a stack with multiple internal
// stacks with a maximum size of maxSize
type SetOfStacks struct {
	stacks      [][]int
	activeStack int
}

// New creates a new SetOfStacks
func New() *SetOfStacks {
	s := new(SetOfStacks)
	s.stacks = make([][]int, 1, 8)
	s.activeStack = 0
	s.stacks[s.activeStack] = make([]int, 0, maxSize)
	return s
}

// Push pushes the value onto the active stack
func (s *SetOfStacks) Push(value int) {
	if len(s.stacks[s.activeStack]) == maxSize {
		s.activeStack++
		s.stacks = append(s.stacks, make([]int, 0, maxSize))
	}
	s.stacks[s.activeStack] = append(s.stacks[s.activeStack], value)
}

// Pop removes and return the topmost value of the stack
func (s *SetOfStacks) Pop() int {
	if len(s.stacks[s.activeStack]) == 0 {
		s.activeStack--
	}
	stack := s.stacks[s.activeStack]
	value := stack[len(stack)-1]
	s.stacks[s.activeStack] = stack[:len(stack)-1]
	return value
}

func main() {
	s := New()
	s.Push(42)
	s.Push(1001)
	s.Push(23)
	s.Push(13)
	fmt.Println("Pop", s.Pop())
	fmt.Println("Pop", s.Pop())
	fmt.Println("Pop", s.Pop())
	fmt.Println("Pop", s.Pop())

}
