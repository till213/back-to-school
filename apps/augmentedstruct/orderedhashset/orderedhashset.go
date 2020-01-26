package orderedhashset

import (
	"container/list"
)

// OrderedHashSet implements a set which keeps track
// of the insertion order, that is the iterator
// returns the elements in the order in which they
// were inserted. Inserting the same element twice does
// not change the insertion order, unless the element
// is deleted first.
type OrderedHashSet struct {
	values map[interface{}]*list.Element
	order  *list.List
}

// New creates and returns a new OrderedHashSet.
func New() *OrderedHashSet {
	return &OrderedHashSet{
		values: make(map[interface{}]*list.Element),
		order:  list.New(),
	}
}

// Add adds the value to the ordered set.
func (s *OrderedHashSet) Add(value interface{}) {
	if _, exists := s.values[value]; !exists {
		element := s.order.PushBack(value)
		s.values[value] = element
	}
}

// Contains returns whether the ordered set contains
// the value or not.
func (s *OrderedHashSet) Contains(value interface{}) bool {
	_, exists := s.values[value]
	return exists
}

// Remove removes the value from the ordered set, if it exists.
func (s *OrderedHashSet) Remove(value interface{}) {
	if element, exists := s.values[value]; exists {
		s.order.Remove(element)
		delete(s.values, value)
	}
}

// Keys returns the keys of the set in the order of insertion.
func (s *OrderedHashSet) Keys() []interface{} {
	keys := make([]interface{}, s.order.Len())
	i := 0
	for element := s.order.Front(); element != nil; element = element.Next() {
		keys[i] = element.Value
		i++
	}
	return keys
}
