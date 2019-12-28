package main

import "fmt"

type stackID int

const (
	left   stackID = iota
	middle         = iota
	right          = iota
)

const empty = 0

// Stack implements three stacks with
// a singe array
type Stack struct {
	data []int
	// positions of next insert
	lnext, rnext int
	// position of "middle stack"
	mpos int
	// extend of "middle stack"
	mleft, mright int
}

func (s *Stack) init(size int) {
	s.data = make([]int, size)
	s.lnext = 0
	s.rnext = size - 1
	s.mpos = size / 2
	s.mleft, s.mright = s.mpos-1, s.mpos
}

func (s *Stack) emptyLeft() bool {
	return s.lnext == 0
}

func (s *Stack) emptyMiddle() bool {
	return s.mleft == s.mright-1
}

func (s *Stack) emptyRight() bool {
	return s.rnext == len(s.data)-1
}

func (s *Stack) fullLeft() bool {
	return s.lnext > s.mleft
}

func (s *Stack) fullMiddle() bool {
	return s.fullLeft() || s.fullRight()
}

func (s *Stack) fullRight() bool {
	return s.rnext < s.mright
}

func (s *Stack) pushLeft(v int) {
	if !s.fullLeft() {
		s.data[s.lnext] = v
		s.lnext++
	}
}

func (s *Stack) popLeft() int {
	var v int
	if !s.emptyLeft() {
		s.lnext--
		v = s.data[s.lnext]
		// Setting the data to "empty" is for "visualisation purposes" only
		// (when debugging s.data)
		s.data[s.lnext] = empty
	} else {
		// error
		v = empty
	}
	return v
}

func (s *Stack) pushMiddle(v int) {
	// First grow right, then left, then right again... ("fair growing")
	if s.mpos-1-s.mleft < s.mright-s.mpos {
		if !s.fullLeft() {
			// grow towards left
			s.data[s.mleft] = v
			s.mleft--
		}
	} else {
		if !s.fullRight() {
			// grow towards right
			s.data[s.mright] = v
			s.mright++
		}
	}
}

func (s *Stack) popMiddle() int {
	var v int
	if !s.emptyMiddle() {
		if s.mpos-1-s.mleft >= s.mright-s.mpos {
			// pop from left
			s.mleft++
			v = s.data[s.mleft]
			s.data[s.mleft] = empty
		} else {
			// pop from right
			s.mright--
			v = s.data[s.mright]
			s.data[s.mright] = empty
		}
	} else {
		v = empty
	}
	return v
}

func (s *Stack) pushRight(v int) {
	if !s.fullRight() {
		s.data[s.rnext] = v
		s.rnext--
	}
}

func (s *Stack) popRight() int {
	var v int
	if !s.emptyRight() {
		s.rnext--
		v = s.data[s.rnext]
		s.data[s.rnext] = empty
	} else {
		// error
		v = empty
	}
	return v
}

func (s *Stack) empty(id stackID) bool {
	switch id {
	case left:
		return s.emptyLeft()
	case middle:
		return s.emptyMiddle()
	case right:
		return s.emptyRight()
	default:
		// error
		return false
	}
}

func (s *Stack) full(id stackID) bool {
	switch id {
	case left:
		return s.fullLeft()
	case middle:
		return s.fullMiddle()
	case right:
		return s.fullRight()
	default:
		// error
		return true
	}
}

func (s *Stack) push(id stackID, v int) {
	switch id {
	case left:
		s.pushLeft(v)
	case middle:
		s.pushMiddle(v)
	case right:
		s.pushRight(v)
	default:
		// error
	}
}

func (s *Stack) pop(id stackID) int {
	switch id {
	case left:
		return s.popLeft()
	case middle:
		return s.popMiddle()
	case right:
		return s.popRight()
	default:
		// error
		return empty
	}
}

func main() {
	var s Stack

	s.init(7)
	fmt.Printf("Empty left: %v middle: %v right: %v\n", s.empty(left), s.empty(middle), s.empty(right))
	s.push(left, 1)
	s.push(middle, 2)
	s.push(right, 3)
	fmt.Printf("Empty left: %v middle: %v right: %v\n", s.empty(left), s.empty(middle), s.empty(right))
	fmt.Println("State", s.data)
	s.push(left, 4)
	s.push(middle, 5)
	s.push(right, 6)
	fmt.Printf("Full left: %v middle: %v right: %v\n", s.full(left), s.full(middle), s.full(right))
	fmt.Println("State", s.data)
	fmt.Println("Pop middle", s.pop(middle))
	fmt.Println("Pop middle", s.pop(middle))
	fmt.Println("State", s.data)
}
