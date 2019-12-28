package main

import "fmt"

type elem struct {
	val  int
	next *elem
}

// Returns the intersecting element (same reference) of lists
// l1 and l2, if exists; otherwise nil
// Example:
// l1: a -> b -> c -> d
// l2: d-> e -> ...
// intersecting element: d
func intersect(l1, l2 *elem) *elem {

	var x *elem
	if l1 == nil || l2 == nil {
		// no intersection with nil lists
		return nil
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

func main() {
	var h1, h2 *elem
	var t1 *elem
	for i := 0; i < 5; i++ {
		e1 := new(elem)
		e1.val = i
		e1.next = h1
		h1 = e1
		if i == 0 {
			t1 = e1
		}

		e2 := new(elem)
		e2.val = 5 + i
		e2.next = h2
		h2 = e2
	}

	x := intersect(h1, h2)
	fmt.Println("Intersecting element", x)

	// Create intersection (by simply appending l2 to l1)
	t1.next = h2

	x = intersect(h1, h2)
	fmt.Println("Intersecting element", x)
}
