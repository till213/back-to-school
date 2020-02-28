package randomkeymap

import (
	"math/rand"
)

type element struct {
	value interface{}
	keyIndex int
}

// RandomKeyMap stores key/values and provides the following
// functions in O(1):
// * Insert
// * Delete
// * Get
// * GetRandomKey
type RandomKeyMap struct {
	kv map[interface{}]element
	keys []interface{}
}

// This is the TRICK! We can remove elements from an array (slice) in O(1) for as
// long as the element order is not important: we simply place the last element
// at the position which is to be deleted and truncate the length of the array by 1.
func (m *RandomKeyMap) removeKey(index int) {
	// https://github.com/golang/go/wiki/SliceTricks#delete-without-preserving-order
	last := len(m.keys)-1
	m.keys[index] = m.keys[last]
	// Make sure the "internal slice pointers" are pointing to nil, for
	// garbage collection to take effect
	m.keys[last] = nil
	m.keys = m.keys[:last]
}

// New creates and returns a new RandomKeyMap.
func New() *RandomKeyMap {
	return &RandomKeyMap{
		kv: make(map[interface{}]element, 0),
		keys: make([]interface{}, 0),
	}
}

// Insert inserts the given value and key into the map.
func (m *RandomKeyMap) Insert(key, value interface{}) {
	elem, ok := m.kv[key]
	if !ok {
		// new element
		m.keys = append(m.keys, key)
		newElem := element{value, len(m.keys) - 1}
		m.kv[key] = newElem
	} else {
		// update value of existing element
		elem.value = value
		m.kv[key] = elem
	}
}

// Get returns the value for the given key. The boolean ok
// indicates whether the value was found in the map.
func (m *RandomKeyMap) Get(key interface{}) (value interface{}, ok bool) {
	var elem element
	elem, ok = m.kv[key]
	if ok {
		value = elem.value
	}
	return
}

// Delete removes the value with the given key from the map.
func (m *RandomKeyMap) Delete(key interface{}) {
	if elem, ok := m.kv[key]; ok {
		delete(m.kv, key)
		keyIndex := elem.keyIndex
		m.removeKey(keyIndex)
		// update the key index: the previous last element was
		// moved to keyIndex
		prevKey := m.keys[keyIndex]
		prevElem := m.kv[prevKey]
		prevElem.keyIndex = keyIndex
		m.kv[prevKey] = prevElem
	}
}

// GetRandomKey returns a random key of the map.
func (m *RandomKeyMap) GetRandomKey() interface{} {
	len := len(m.keys)
	r := rand.Intn(len)
	return m.keys[r]
}