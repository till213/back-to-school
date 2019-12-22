package list

// Element of a single-linked list
type Element struct {
	Val  int
	Next *Element
}

// Insert inserts a new list element with Value v after this element,
// and returns the newly inserted element
func (e *Element) Insert(Val int) *Element {
	q := new(Element)
	q.Val = Val
	q.Next = e.Next
	e.Next = q
	return q
}

// RemoveNext removes the successor element of this element.
func (e *Element) RemoveNext() {
	if e.Next != nil {
		e.Next = e.Next.Next
	}
}

// Remove removes this element, by copying the Value from
// the Next element and removing the Next element
func (e *Element) Remove() {
	if e.Next != nil {
		e.Val = e.Next.Val
		e.RemoveNext()
	}
}
