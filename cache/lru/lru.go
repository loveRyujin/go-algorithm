package lru

import (
	"container/list"
	"sync"
)

// Cache LRU cache structure
type Cache struct {
	capacity int
	cache    map[any]*list.Element
	list     *list.List
	mutex    sync.RWMutex
}

// entry cache entry
type entry struct {
	key   any
	value any
}

// New creates a new LRU cache
func New(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		cache:    make(map[any]*list.Element),
		list:     list.New(),
	}
}

// Get retrieves a value from the cache
func (c *Cache) Get(key any) (any, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, ok := c.cache[key]; ok {
		// Move the accessed element to the front of the list
		c.list.MoveToFront(element)
		return element.Value.(*entry).value, true
	}
	return nil, false
}

// Put adds a key-value pair to the cache
func (c *Cache) Put(key, value any) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, ok := c.cache[key]; ok {
		// If the key already exists, update the value and move to front
		element.Value.(*entry).value = value
		c.list.MoveToFront(element)
		return
	}

	// If the cache is full, remove the least recently used element
	if c.list.Len() >= c.capacity {
		c.removeOldest()
	}

	// Add new element to the front of the list
	newEntry := &entry{key: key, value: value}
	element := c.list.PushFront(newEntry)
	c.cache[key] = element
}

// Remove removes a key from the cache
func (c *Cache) Remove(key any) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, ok := c.cache[key]; ok {
		c.removeElement(element)
		return true
	}
	return false
}

// removeOldest removes the least recently used element (tail of the list)
func (c *Cache) removeOldest() {
	if c.list.Len() == 0 {
		return
	}
	oldest := c.list.Back()
	if oldest != nil {
		c.removeElement(oldest)
	}
}

// removeElement removes a specific element
func (c *Cache) removeElement(element *list.Element) {
	c.list.Remove(element)
	delete(c.cache, element.Value.(*entry).key)
}

// Len returns the number of elements in the cache
func (c *Cache) Len() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.list.Len()
}

// Cap returns the capacity of the cache
func (c *Cache) Cap() int {
	// Capacity doesn't change, no lock needed
	return c.capacity
}

// Clear removes all elements from the cache
func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache = make(map[any]*list.Element)
	c.list = list.New()
}

// Keys returns all keys in the cache (in access order, most recent first)
func (c *Cache) Keys() []any {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	keys := make([]any, 0, c.list.Len())
	for element := c.list.Front(); element != nil; element = element.Next() {
		keys = append(keys, element.Value.(*entry).key)
	}
	return keys
}

// Contains checks if the cache contains a specific key
func (c *Cache) Contains(key any) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	_, ok := c.cache[key]
	return ok
}

// Peek looks up a value without updating the access order
func (c *Cache) Peek(key any) (any, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if element, ok := c.cache[key]; ok {
		return element.Value.(*entry).value, true
	}
	return nil, false
}
