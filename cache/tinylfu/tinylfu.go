package tinylfu

import (
	"sync"
)

// TinyLFU cache implementation
type Cache struct {
	capacity      int
	windowSize    int // Size of the admission window (LRU)
	protectedSize int // Size of the protected segment
	probationSize int // Size of the probation segment

	window    *lruCache // Admission window (LRU)
	probation *lruCache // Probation segment (LRU)
	protected *lruCache // Protected segment (LRU)

	sketch     *countMinSketch // Frequency estimation
	doorkeeper *bloomFilter    // Bloom filter for doorkeeper

	size       int // Current total size
	sampleSize int // Sample size for reset
	mutex      sync.RWMutex
}

// New creates a new TinyLFU cache
func New(capacity int) *Cache {
	if capacity <= 0 {
		capacity = 1
	}

	var windowSize, probationSize, protectedSize int

	if capacity == 1 {
		// Special case for capacity 1
		windowSize = 1
		probationSize = 0
		protectedSize = 0
	} else if capacity == 2 {
		// Special case for capacity 2
		windowSize = 1
		probationSize = 1
		protectedSize = 0
	} else {
		// Normal distribution: 1% window, but at least 1
		windowSize = max(1, capacity/100)
		remaining := capacity - windowSize
		protectedSize = max(1, remaining*80/100)
		probationSize = capacity - windowSize - protectedSize

		// Ensure probationSize is not negative
		if probationSize < 0 {
			probationSize = 0
			protectedSize = capacity - windowSize
		}
	}

	cache := &Cache{
		capacity:      capacity,
		windowSize:    windowSize,
		protectedSize: protectedSize,
		probationSize: probationSize,
		window:        newLRUCache(windowSize),
		probation:     newLRUCache(probationSize),
		protected:     newLRUCache(protectedSize),
		sketch:        newCountMinSketch(capacity * 10), // 10x capacity for better accuracy
		doorkeeper:    newBloomFilter(capacity * 4),     // 4x capacity for bloom filter
		sampleSize:    capacity * 10,                    // Reset sketch every 10x accesses
	}

	return cache
}

// Get retrieves a value from the cache
func (c *Cache) Get(key any) (any, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.sketch.increment(key)

	// Check window first
	if value, ok := c.window.get(key); ok {
		return value, true
	}

	// Check probation
	if value, ok := c.probation.get(key); ok {
		// Move to protected
		c.probation.remove(key)
		evicted := c.protected.put(key, value)
		if evicted != nil {
			// Move evicted from protected to probation
			c.probation.put(evicted.key, evicted.value)
		}
		return value, true
	}

	// Check protected
	if value, ok := c.protected.get(key); ok {
		return value, true
	}

	return nil, false
}

// Put adds or updates a key-value pair in the cache
func (c *Cache) Put(key, value any) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Check if key exists in any segment
	if c.window.contains(key) {
		c.window.put(key, value)
		c.sketch.increment(key)
		return
	}
	if c.probation.contains(key) {
		c.probation.put(key, value)
		c.sketch.increment(key)
		return
	}
	if c.protected.contains(key) {
		c.protected.put(key, value)
		c.sketch.increment(key)
		return
	}

	// New key
	c.sketch.increment(key)

	// Add to window (new keys always go to window first)
	c.addToWindow(key, value)
}

// addToWindow adds a new key to the admission window
func (c *Cache) addToWindow(key, value any) {
	// For very small caches, handle specially
	if c.capacity == 1 {
		c.window.put(key, value)
		c.size++
		if c.size > c.sampleSize {
			c.sketch.reset()
			c.doorkeeper.reset()
			c.size = 0
		}
		return
	}

	evicted := c.window.put(key, value)
	if evicted != nil {
		// Window is full, try to admit evicted key to main cache
		if c.admit(evicted.key) {
			// Admit to probation
			victim := c.probation.put(evicted.key, evicted.value)
			if victim != nil {
				// Probation is full, try to move victim to protected
				protectedVictim := c.protected.put(victim.key, victim.value)
				if protectedVictim != nil {
					// Protected is full, add the final victim to doorkeeper
					c.doorkeeper.add(protectedVictim.key)
				}
			}
		} else {
			// Not admitted, add to doorkeeper
			c.doorkeeper.add(evicted.key)
		}
	}

	c.size++
	if c.size > c.sampleSize {
		c.sketch.reset()
		c.doorkeeper.reset()
		c.size = 0
	}
}

// admit decides whether a key should be admitted to the main cache
func (c *Cache) admit(candidateKey any) bool {
	// If probation segment doesn't exist (capacity too small), don't admit
	if c.probationSize == 0 {
		return false
	}

	// Always admit if there's space in probation
	if c.probation.len() < c.probationSize {
		return true
	}

	// Use doorkeeper bloom filter for recently evicted items
	if c.doorkeeper.mightContain(candidateKey) {
		return false
	}

	// Compare with LRU item in probation
	victimKey := c.probation.getLRUKey()
	if victimKey == nil {
		return true
	}

	candidateFreq := c.sketch.estimate(candidateKey)
	victimFreq := c.sketch.estimate(victimKey)

	// Admit if candidate has higher or equal frequency
	if candidateFreq >= victimFreq {
		c.doorkeeper.add(victimKey)
		return true
	}

	return false
}

// Remove removes a key from the cache
func (c *Cache) Remove(key any) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.window.remove(key) {
		return true
	}
	if c.probation.remove(key) {
		return true
	}
	if c.protected.remove(key) {
		return true
	}
	return false
}

// Contains checks if a key exists in the cache
func (c *Cache) Contains(key any) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.window.contains(key) ||
		c.probation.contains(key) ||
		c.protected.contains(key)
}

// Len returns the current size of the cache
func (c *Cache) Len() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.window.len() + c.probation.len() + c.protected.len()
}

// Cap returns the capacity of the cache
func (c *Cache) Cap() int {
	return c.capacity
}

// Clear removes all entries from the cache
func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.window.clear()
	c.probation.clear()
	c.protected.clear()
	c.sketch.reset()
	c.doorkeeper.reset()
	c.size = 0
}

// Stats returns cache statistics
func (c *Cache) Stats() map[string]int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return map[string]int{
		"capacity":       c.capacity,
		"size":           c.Len(),
		"window_size":    c.window.len(),
		"probation_size": c.probation.len(),
		"protected_size": c.protected.len(),
		"window_cap":     c.windowSize,
		"probation_cap":  c.probationSize,
		"protected_cap":  c.protectedSize,
	}
}

// GetFrequency returns the estimated frequency of a key
func (c *Cache) GetFrequency(key any) int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.sketch.estimate(key)
}

// Helper function
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
