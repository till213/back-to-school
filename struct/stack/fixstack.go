package stack

// N defines the maximum depth of fixed-size stack
const N = 32

// Stack of maximum depth N
type FixedLengthStack struct {
	a [N]int
	d int
}

// Create initialises a new stack
func (s *FixedLengthStack) Create() {
	s.d = 0
}

// Empty returns true if the stack is empty
func (s *FixedLengthStack) Empty() bool {
	return s.d == 0
}

// Full returns true if the stack is full
func (s *FixedLengthStack) Full() bool {
	return s.d == N
}

// Push pushes the value val onto the stack, which
// must be non-full
func (s *FixedLengthStack) Push(val int) {
	s.a[s.d] = val
	s.d++
}

// Top returns the top value on the stack
// must be non-empty
func (s *FixedLengthStack) Top() int {
	return s.a[s.d-1]
}

// Pop removes the top value from the stack
func (s *FixedLengthStack) Pop() {
	s.d--
}
