package main

import "fmt"

type elem struct {
	val  int
	next *elem
}

// Returns the intersecting element (same reference) of lists
// l1 and l2, if exists; otherwise nil, using a hashmap
// Example:
// l1: a -> b -> c -> d
// l2: d -> e -> ...
// intersecting element: d
func intersect(l1, l2 *elem) *elem {
	var x *elem
	if l1 == nil || l2 == nil {
		// no intersection with nil lists
		return nil
	} else if l1 == l2 {
		return l1
	}
	em := make(map[*elem]bool)
	// assumption: we do not know the lengths
	// of the lists; otherwise we could save
	// auxiliary space, by iterating over the
	// smaller list first
	for l1 != nil {
		em[l1] = true
		l1 = l1.next
	}
	for l2 != nil && x == nil {
		if em[l2] {
			x = l2
		} else {
			l2 = l2.next
		}
	}
	return x
}

func tail(e *elem) (t *elem, len int) {
	for e != nil {
		t = e
		e = e.next
		len++
	}
	return
}

func hasIntersection(e1, e2 *elem) (intersect bool, len1, len2 int) {
	t1, len1 := tail(e1)
	t2, len2 := tail(e2)
	intersect = t1 == t2
	return
}

// Returns the intersecting element (same reference) of lists
// l1 and l2, if exists; otherwise nil, using O(1) space
// Observation: intersecting lists are the same "as seen from
// the end" (in reverse order), until some "branching" element
// -> that's the "intersecting" element that we are looking for
// Example:
// l1: a -> b -> c -> d -> e -> f
// l2: x -> y -> e -> f
// intersecting element: e
// Note: "destructive iteration" (e.g. reversing pointers) does
//       not work in case there is indeed an intersection, as
//       pointer reversal for l1 would also affect l2
func intersect2(l1, l2 *elem) *elem {
	var x *elem
	if l1 == nil || l2 == nil {
		// no intersection with nil lists
		return nil
	} else if l1 == l2 {
		return l1
	}
	intersect, len1, len2 := hasIntersection(l1, l2)
	if intersect {
		var e1, e2 *elem
		var n int
		if len1 < len2 {
			e1 = l2
			e2 = l1
			n = len2 - len1
		} else {
			e1 = l1
			e2 = l2
			n = len1 - len2
		}
		// advance e1 (which points to the longer list)
		// by n positions
		for i := 0; i < n; i++ {
			e1 = e1.next
		}
		for e1 != e2 {

			e1 = e1.next
			e2 = e2.next
		}
		x = e1
	}
	return x
}

func main() {
	var h1, h2 *elem
	var t1 *elem
	for i := 0; i < 2; i++ {
		e1 := new(elem)
		e1.val = i
		e1.next = h1
		h1 = e1
		if i == 0 {
			t1 = e1
		}

		e2 := new(elem)
		e2.val = 2 + i
		e2.next = h2
		h2 = e2
	}

	for e := h1; e != nil; e = e.next {
		fmt.Printf("#1 Element %p, value: %v\n", e, e.val)
	}
	for e := h2; e != nil; e = e.next {
		fmt.Printf("#2 Element %p, value: %v\n", e, e.val)
	}

	x := intersect2(h1, h2)
	fmt.Println("Intersecting element", x)

	// Create intersection (by simply appending l2 to l1)
	t1.next = h2
	for e := h1; e != nil; e = e.next {
		fmt.Printf("#1 Element %p, value: %v\n", e, e.val)
	}
	for e := h2; e != nil; e = e.next {
		fmt.Printf("#2 Element %p, value: %v\n", e, e.val)
	}

	x = intersect2(h1, h2)
	fmt.Println("Intersecting element", x)
}
