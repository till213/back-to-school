package orderedhashset

import "container/list"

// Iterator implements an iterator over the ordered set.
type Iterator struct {
	current *list.Element
}

// Iterator creates and returns a new iterator for the ordered set.
func (s *OrderedHashSet) Iterator() *Iterator {
	iter := new(Iterator)
	iter.current = s.order.Front()
	return iter
}

// Next returns the next key of the ordered set, or nil if
// the end of the element list has been reached.
func (i *Iterator) Next() *Iterator {
	i.current = i.current.Next()
	if i.current == nil {
		return nil
	}
	return i
}

// Value returns the value of the iterator.
func (i *Iterator) Value() interface{} {
	if i == nil {
		return nil
	}
	return i.current.Value
}
