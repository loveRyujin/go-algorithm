package tinylfu

import "container/list"

// lruCache is a simple LRU cache implementation
type lruCache struct {
	capacity int
	cache    map[any]*list.Element
	list     *list.List
}

// lruEntry represents an entry in the LRU cache
type lruEntry struct {
	key   any
	value any
}

// newLRUCache creates a new LRU cache
func newLRUCache(capacity int) *lruCache {
	return &lruCache{
		capacity: capacity,
		cache:    make(map[any]*list.Element),
		list:     list.New(),
	}
}

// get retrieves a value and moves it to front
func (l *lruCache) get(key any) (any, bool) {
	if element, ok := l.cache[key]; ok {
		l.list.MoveToFront(element)
		return element.Value.(*lruEntry).value, true
	}
	return nil, false
}

// put adds or updates a key-value pair, returns evicted entry if any
func (l *lruCache) put(key, value any) *lruEntry {
	if element, ok := l.cache[key]; ok {
		// Update existing key
		l.list.MoveToFront(element)
		entry := element.Value.(*lruEntry)
		entry.value = value
		return nil
	}

	// Add new key
	var evicted *lruEntry
	if l.list.Len() >= l.capacity {
		// Evict LRU
		oldest := l.list.Back()
		if oldest != nil {
			entry := oldest.Value.(*lruEntry)
			evicted = &lruEntry{key: entry.key, value: entry.value}
			delete(l.cache, entry.key)
			l.list.Remove(oldest)
		}
	}

	newEntry := &lruEntry{key: key, value: value}
	element := l.list.PushFront(newEntry)
	l.cache[key] = element
	return evicted
}

// remove removes a key from the cache
func (l *lruCache) remove(key any) bool {
	if element, ok := l.cache[key]; ok {
		delete(l.cache, key)
		l.list.Remove(element)
		return true
	}
	return false
}

// contains checks if a key exists in the cache
func (l *lruCache) contains(key any) bool {
	_, ok := l.cache[key]
	return ok
}

// len returns the current size of the cache
func (l *lruCache) len() int {
	return l.list.Len()
}

// getLRUKey returns the least recently used key
func (l *lruCache) getLRUKey() any {
	if l.list.Len() == 0 {
		return nil
	}
	oldest := l.list.Back()
	if oldest != nil {
		return oldest.Value.(*lruEntry).key
	}
	return nil
}

// clear removes all entries from the cache
func (l *lruCache) clear() {
	l.cache = make(map[any]*list.Element)
	l.list = list.New()
}
