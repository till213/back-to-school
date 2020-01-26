// Example usage:
// cache := Constructor(capacity);
// value := cache.Get(key);
// cache.Put(key, value);
package lrucache

import ( "container/list")

const NotFound = -1

// Item stores the value and a reference to the entry
// in the LRU list
type Item struct {
	value int
	lruElement *list.Element
}

// LRUCache implements a cache with a least recently used (LRU)
// eviction strategy
// https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU)
type LRUCache struct {
	capacity int
	cache map[int]Item
	lruList *list.List
}

// Constructor creates a new LRUCache
func Constructor(capacity int) LRUCache {
	var lruCache LRUCache
	lruCache.cache = make(map[int]Item)
	lruCache.lruList = list.New()
	lruCache.capacity = capacity
	return lruCache
}

// Get returns the value for key and makes it the least
// recently used
func (c *LRUCache) Get(key int) int {
	v, ok := c.cache[key]
	if ok {
		c.lruList.MoveToFront(v.lruElement)
		return v.value
	} else {
		return NotFound
	}
}

// Put inserts a new element with value and key and makes 
// it the least recently used
func (c *LRUCache) Put(key int, value int)  {
	item, ok := c.cache[key]
	if ok {
		// update existing cache item
		c.lruList.MoveToFront(item.lruElement)
		item.value = value
	} else {
		// new item - capacity?
		if len(c.cache) == c.capacity {
			// cache eviction
			oldkey := c.lruList.Back()
			delete(c.cache, oldkey.Value.(int))
			c.lruList.Remove(oldkey)		
		}
		lruElement := c.lruList.PushFront(key)
		item.value = value
		item.lruElement = lruElement
		c.cache[key] = item
	}
}
