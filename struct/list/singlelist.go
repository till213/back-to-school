package list

// Element of a single-linked list
type Element struct {
	val  int
	next *Element
}

// Insert inserts a new list element with value v after this element,
// and returns the newly inserted element
func (e *Element) Insert(val int) *Element {
	q := new(Element)
	q.val = val
	q.next = e.next
	e.next = q
	return q
}

// RemoveNext removes the successor element of this element.
func (e *Element) RemoveNext() {
	if e.next != nil {
		e.next = e.next.next
	}
}

// Remove removes this element, by copying the value from
// the next element and removing the next element
func (e *Element) Remove() {
	if e.next != nil {
		e.val = e.next.val
		e.RemoveNext()
	}
}
