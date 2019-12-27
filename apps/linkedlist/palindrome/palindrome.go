package main

import (
	"fmt"
)

type elem struct {
	val  int
	next *elem
}

func isPalindrome(head *elem) bool {
	var p, m1 *elem
	if head == nil {
		return false
	}

	f, m := head, head
	
	odd := true
	// runner pointer is f
	for f.next != nil {
		t := m
		m = m.next
		// m1 points "one ahead" of m
		m1 = m.next
		f = f.next
		// reverse pointer to previous p
		t.next = p
		p = t
		if f.next == nil {
			odd = false
		} else {
			// fast pointer advances one more (if possible)
			f = f.next
		}
	}
	// Post: m is in the "middle"
	// - the exact middle, if odd == true
	// - just one element before, if odd == false

	// in an "odd" palindrome the middle element can be anything
	// -> simply "advance" the "middle" pointer (which is in fact
	// "backwards" towards head, due to the reversed pointers)
	if odd {
		// s may already be nil afterwards (list with just 1 element)
		m = m.next
	}
	f = m1
	palin := true
	for m != nil && palin {
		if m.val != f.val {
			palin = false
		} else {
			// f will move towards tail while...
			f = f.next
			// m moves "backwards" (reversed pointers) towards head
			m = m.next
		}
	}

	return palin
}

func main() {
	const half = 3
	head := new(elem)
	head.val = half
	for i := half; i > 0; i-- {
		e := new(elem)
		e.val = i
		e.next = head.next
		head.next = e
	}
	for i := 0; i < half; i++ {
		e := new(elem)
		e.val = i
		e.next = head.next
		head.next = e
	}
	fmt.Println("--- Palindrome list ---")
	for e := head; e != nil; e = e.next {
		fmt.Println("Value: ", e.val)
	}
	fmt.Println("Is palindrome: ", isPalindrome(head))

	for i := 0; i < half; i++ {
		e := new(elem)
		e.val = i * i
		e.next = head.next
		head.next = e
	}
	fmt.Println("--- Non-Palindrome list ---")
	for e := head; e != nil; e = e.next {
		fmt.Println("Value: ", e.val)
	}
	fmt.Println("Is palindrome: ", isPalindrome(head))
}
