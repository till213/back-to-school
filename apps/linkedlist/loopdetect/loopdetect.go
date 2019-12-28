package main

import "fmt"

type elem struct {
	val  int
	next *elem
}

// Returns 0 if there is no loop in list l; otherwise
// the length (number of elements) of the loop
func loopLength(l *elem) int {
	var len int
	if l == nil {
		return 0
	} else if l.next == l {
		// Loop to itself
		return 1
	} else if l.next != nil {
		s := l
		f := l.next
		// This loop terminates if
		// f == nil: the "fast" pointer reached the end
		//           of the list -> no loop
		// s == f: the "slow" pointer eventually caught
		//         up with the "fast" pointer -> loop
		for f != nil && s != f {
			s = s.next
			f = f.next
			if f != nil {
				f = f.next
			}
		}
		if f == nil {
			return 0
		}
		// we are in a loop, determine length by counting
		// the time until the "fast pointer" catches up
		// with the "slow" pointer
		s = s.next
		f = f.next.next
		len = 1
		for s != f {
			s = s.next
			f = f.next.next
			len++
		}
	}
	// List with one element, no loop
	return len
}

// Returns the element which forms a loop
// Example:
// a -> b -> c* -> d -> e -> c* (-> d -> e -> c* -> ...)
// (c* is the "loop element")
func loopDetect(l *elem) *elem {
	var e *elem
	len := loopLength(l)
	if len > 0 {
		// Start at beginning and "peek ahead" len elements (with a "fast" pointer)
		e = l
		f := l
		// Move the "fast pointer" from the first to the len-th position
		for i := 0; i < len; i++ {
			f = f.next
		}
		// Move the "slow" and "fast" pointers at the same speed: once they
		// are equal we have found the "looping element"
		for e != f {
			e = e.next
			f = f.next
		}
	}
	return e
}

func main() {
	var h, e3, t *elem
	for i := 0; i < 5; i++ {
		e := new(elem)
		if i == 0 {
			t = e
		} else if i == 3 {
			e3 = e
		}
		e.val = i
		e.next = h
		h = e
	}

	// Create a vicious loop
	t.next = e3
	loopElem := loopDetect(h)
	fmt.Println("Has loop elem", loopElem)
}
