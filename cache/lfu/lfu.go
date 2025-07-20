package lfu

import (
	"container/list"
	"sync"
)

// Cache LFU cache structure
type Cache struct {
	capacity    int
	cache       map[any]*list.Element
	frequencies map[int]*list.List // frequency -> doubly linked list
	minFreq     int                // minimum frequency
	mutex       sync.RWMutex
}

// entry cache entry with frequency information
type entry struct {
	key   any
	value any
	freq  int
}

// New creates a new LFU cache
func New(capacity int) *Cache {
	return &Cache{
		capacity:    capacity,
		cache:       make(map[any]*list.Element),
		frequencies: make(map[int]*list.List),
		minFreq:     1,
	}
}

// Get retrieves a value from the cache
func (c *Cache) Get(key any) (any, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, ok := c.cache[key]; ok {
		entry := element.Value.(*entry)
		c.updateFrequency(element, entry)
		return entry.value, true
	}
	return nil, false
}

// Put adds a key-value pair to the cache
func (c *Cache) Put(key, value any) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.capacity <= 0 {
		return
	}

	if element, ok := c.cache[key]; ok {
		// Update existing key
		entry := element.Value.(*entry)
		entry.value = value
		c.updateFrequency(element, entry)
		return
	}

	// If cache is full, evict least frequently used item
	if len(c.cache) >= c.capacity {
		c.evictLFU()
	}

	// Add new entry with frequency 1
	newEntry := &entry{key: key, value: value, freq: 1}

	// Initialize frequency list if it doesn't exist
	if c.frequencies[1] == nil {
		c.frequencies[1] = list.New()
	}

	element := c.frequencies[1].PushFront(newEntry)
	c.cache[key] = element
	c.minFreq = 1
}

// updateFrequency moves an entry to the next frequency level
func (c *Cache) updateFrequency(element *list.Element, entry *entry) {
	oldFreq := entry.freq
	newFreq := oldFreq + 1

	// Remove from old frequency list
	c.frequencies[oldFreq].Remove(element)

	// If old frequency list is empty and it's the minimum frequency, update minFreq
	if c.frequencies[oldFreq].Len() == 0 && oldFreq == c.minFreq {
		c.minFreq = newFreq
	}

	// Update entry frequency
	entry.freq = newFreq

	// Initialize new frequency list if it doesn't exist
	if c.frequencies[newFreq] == nil {
		c.frequencies[newFreq] = list.New()
	}

	// Add to new frequency list
	newElement := c.frequencies[newFreq].PushFront(entry)
	c.cache[entry.key] = newElement
}

// evictLFU removes the least frequently used item
func (c *Cache) evictLFU() {
	if c.frequencies[c.minFreq] == nil || c.frequencies[c.minFreq].Len() == 0 {
		return
	}

	// Get the least recently used item among items with minimum frequency
	oldest := c.frequencies[c.minFreq].Back()
	if oldest != nil {
		entry := oldest.Value.(*entry)
		c.frequencies[c.minFreq].Remove(oldest)
		delete(c.cache, entry.key)
	}
}

// Remove removes a key from the cache
func (c *Cache) Remove(key any) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, ok := c.cache[key]; ok {
		entry := element.Value.(*entry)
		c.frequencies[entry.freq].Remove(element)
		delete(c.cache, key)

		// Update minFreq if necessary
		if c.frequencies[c.minFreq].Len() == 0 {
			c.updateMinFreq()
		}

		return true
	}
	return false
}

// updateMinFreq finds the new minimum frequency
func (c *Cache) updateMinFreq() {
	for freq := c.minFreq; freq <= len(c.cache)+c.minFreq; freq++ {
		if c.frequencies[freq] != nil && c.frequencies[freq].Len() > 0 {
			c.minFreq = freq
			return
		}
	}
	c.minFreq = 1
}

// Len returns the number of elements in the cache
func (c *Cache) Len() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return len(c.cache)
}

// Cap returns the capacity of the cache
func (c *Cache) Cap() int {
	return c.capacity
}

// Clear removes all elements from the cache
func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache = make(map[any]*list.Element)
	c.frequencies = make(map[int]*list.List)
	c.minFreq = 1
}

// Contains checks if the cache contains a specific key
func (c *Cache) Contains(key any) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	_, ok := c.cache[key]
	return ok
}

// Peek looks up a value without updating the frequency
func (c *Cache) Peek(key any) (any, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if element, ok := c.cache[key]; ok {
		entry := element.Value.(*entry)
		return entry.value, true
	}
	return nil, false
}

// Keys returns all keys in the cache grouped by frequency (highest frequency first)
func (c *Cache) Keys() map[int][]any {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	result := make(map[int][]any)

	for freq, freqList := range c.frequencies {
		if freqList.Len() > 0 {
			keys := make([]any, 0, freqList.Len())
			for element := freqList.Front(); element != nil; element = element.Next() {
				entry := element.Value.(*entry)
				keys = append(keys, entry.key)
			}
			result[freq] = keys
		}
	}

	return result
}

// GetFrequency returns the frequency of a key
func (c *Cache) GetFrequency(key any) int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if element, ok := c.cache[key]; ok {
		entry := element.Value.(*entry)
		return entry.freq
	}
	return 0
}

// MinFrequency returns the minimum frequency in the cache
func (c *Cache) MinFrequency() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.minFreq
}
