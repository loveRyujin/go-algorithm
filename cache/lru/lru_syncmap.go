package lru

import (
	"container/list"
	"sync"
)

// SyncMapCache LRU cache using sync.Map
type SyncMapCache struct {
	capacity int
	cache    sync.Map // sync.Map for concurrent access
	list     *list.List
	mutex    sync.Mutex // Still need mutex for list operations
}

// NewSyncMap creates a new LRU cache using sync.Map
func NewSyncMap(capacity int) *SyncMapCache {
	return &SyncMapCache{
		capacity: capacity,
		list:     list.New(),
	}
}

// Get retrieves a value from the cache using sync.Map
func (c *SyncMapCache) Get(key any) (any, bool) {
	if value, ok := c.cache.Load(key); ok {
		element := value.(*list.Element)

		// Still need to protect list operations
		c.mutex.Lock()
		c.list.MoveToFront(element)
		c.mutex.Unlock()

		return element.Value.(*entry).value, true
	}
	return nil, false
}

// Put adds a key-value pair to the cache using sync.Map
func (c *SyncMapCache) Put(key, value any) {
	if existingValue, ok := c.cache.Load(key); ok {
		element := existingValue.(*list.Element)

		c.mutex.Lock()
		element.Value.(*entry).value = value
		c.list.MoveToFront(element)
		c.mutex.Unlock()
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	// If cache is full, remove the least recently used element
	if c.list.Len() >= c.capacity {
		c.removeOldest()
	}

	// Add new element to the front of the list
	newEntry := &entry{key: key, value: value}
	element := c.list.PushFront(newEntry)
	c.cache.Store(key, element)
}

// Remove removes a key from the cache
func (c *SyncMapCache) Remove(key any) bool {
	if value, ok := c.cache.LoadAndDelete(key); ok {
		element := value.(*list.Element)

		c.mutex.Lock()
		c.list.Remove(element)
		c.mutex.Unlock()

		return true
	}
	return false
}

// removeOldest removes the least recently used element (tail of the list)
func (c *SyncMapCache) removeOldest() {
	if c.list.Len() == 0 {
		return
	}
	oldest := c.list.Back()
	if oldest != nil {
		c.list.Remove(oldest)
		c.cache.Delete(oldest.Value.(*entry).key)
	}
}

// Len returns the number of elements in the cache
func (c *SyncMapCache) Len() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.list.Len()
}

// Cap returns the capacity of the cache
func (c *SyncMapCache) Cap() int {
	return c.capacity
}

// Clear removes all elements from the cache
func (c *SyncMapCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache = sync.Map{}
	c.list = list.New()
}

// Keys returns all keys in the cache (in access order, most recent first)
func (c *SyncMapCache) Keys() []any {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	keys := make([]any, 0, c.list.Len())
	for element := c.list.Front(); element != nil; element = element.Next() {
		keys = append(keys, element.Value.(*entry).key)
	}
	return keys
}

// Contains checks if the cache contains a specific key
func (c *SyncMapCache) Contains(key any) bool {
	_, ok := c.cache.Load(key)
	return ok
}

// Peek looks up a value without updating the access order
func (c *SyncMapCache) Peek(key any) (any, bool) {
	if value, ok := c.cache.Load(key); ok {
		element := value.(*list.Element)
		return element.Value.(*entry).value, true
	}
	return nil, false
}
