package lrucache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	const capacity = 2

	cache := Constructor(capacity)

	cache.Put(1, 1);
	cache.Put(2, 2);
	val := cache.Get(1);
	if val != 1 {
		t.Errorf("Expected 1, received %v", val)
	}

	// evicts key 2
	cache.Put(3, 3);
	val = cache.Get(2);
	if val != NotFound {
		t.Errorf("Expected NotFound, received %v", val)
	}

	// evicts key 1
	cache.Put(4, 4);  
	val = cache.Get(1);      
	if val != NotFound {
		t.Errorf("Expected NotFound, received %v", val)
	}
	val = cache.Get(3); 
	if val != 3 {
		t.Errorf("Expected 3, received %v", val)
	}
	val = cache.Get(4);
	if val != 4 {
		t.Errorf("Expected 4, received %v", val)
	}
}

