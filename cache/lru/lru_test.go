package lru

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := New(2)

	// Test basic Put and Get operations
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	if value, ok := cache.Get("key1"); !ok || value != "value1" {
		t.Errorf("Expected value1, got %v", value)
	}

	if value, ok := cache.Get("key2"); !ok || value != "value2" {
		t.Errorf("Expected value2, got %v", value)
	}

	// Test capacity limit
	cache.Put("key3", "value3") // This should evict key1 since key2 was accessed recently

	if _, ok := cache.Get("key1"); ok {
		t.Error("key1 should have been evicted")
	}

	if value, ok := cache.Get("key2"); !ok || value != "value2" {
		t.Errorf("key2 should still be in cache, got %v", value)
	}

	if value, ok := cache.Get("key3"); !ok || value != "value3" {
		t.Errorf("key3 should be in cache, got %v", value)
	}
}

func TestLRUCacheUpdate(t *testing.T) {
	cache := New(2)

	// Test updating existing key
	cache.Put("key1", "value1")
	cache.Put("key1", "updated_value1")

	if value, ok := cache.Get("key1"); !ok || value != "updated_value1" {
		t.Errorf("Expected updated_value1, got %v", value)
	}

	// Cache should still have space
	if cache.Len() != 1 {
		t.Errorf("Expected length 1, got %d", cache.Len())
	}
}

func TestLRUCacheEviction(t *testing.T) {
	cache := New(3)

	// Fill up the cache
	cache.Put(1, "one")
	cache.Put(2, "two")
	cache.Put(3, "three")

	// Access key 1 to make it most recently used
	cache.Get(1)

	// Add new key, should evict key 2
	cache.Put(4, "four")

	if _, ok := cache.Get(2); ok {
		t.Error("key 2 should have been evicted")
	}

	// Check that other keys are still in cache
	if _, ok := cache.Get(1); !ok {
		t.Error("key 1 should still be in cache")
	}

	if _, ok := cache.Get(3); !ok {
		t.Error("key 3 should still be in cache")
	}

	if _, ok := cache.Get(4); !ok {
		t.Error("key 4 should be in cache")
	}
}

func TestLRUCacheRemove(t *testing.T) {
	cache := New(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// Test removing existing key
	if !cache.Remove("b") {
		t.Error("Remove should return true for existing key")
	}

	if _, ok := cache.Get("b"); ok {
		t.Error("key b should have been removed")
	}

	if cache.Len() != 2 {
		t.Errorf("Expected length 2, got %d", cache.Len())
	}

	// Test removing non-existing key
	if cache.Remove("d") {
		t.Error("Remove should return false for non-existing key")
	}
}

func TestLRUCacheContains(t *testing.T) {
	cache := New(2)

	cache.Put("key1", "value1")

	if !cache.Contains("key1") {
		t.Error("Cache should contain key1")
	}

	if cache.Contains("key2") {
		t.Error("Cache should not contain key2")
	}
}

func TestLRUCachePeek(t *testing.T) {
	cache := New(2)

	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	// Peek should not affect access order
	if value, ok := cache.Peek("key1"); !ok || value != "value1" {
		t.Errorf("Peek should return value1, got %v", value)
	}

	// Add new key, key1 should be evicted (since Peek didn't update access order)
	cache.Put("key3", "value3")

	if _, ok := cache.Get("key1"); ok {
		t.Error("key1 should have been evicted")
	}
}

func TestLRUCacheKeys(t *testing.T) {
	cache := New(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// Access a to make it most recently used
	cache.Get("a")

	keys := cache.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	// Check key order (most recently used first)
	if keys[0] != "a" {
		t.Errorf("Expected first key to be 'a', got %v", keys[0])
	}
}

func TestLRUCacheClear(t *testing.T) {
	cache := New(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	cache.Clear()

	if cache.Len() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", cache.Len())
	}

	if _, ok := cache.Get("a"); ok {
		t.Error("Cache should be empty after clear")
	}
}

func TestLRUCacheCapacity(t *testing.T) {
	cache := New(5)

	if cache.Cap() != 5 {
		t.Errorf("Expected capacity 5, got %d", cache.Cap())
	}

	// Adding elements should not change capacity
	cache.Put("key", "value")
	if cache.Cap() != 5 {
		t.Errorf("Capacity should remain 5, got %d", cache.Cap())
	}
}

// Benchmark tests
func BenchmarkLRUCachePut(b *testing.B) {
	cache := New(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Put(i, i)
	}
}

func BenchmarkLRUCacheGet(b *testing.B) {
	cache := New(1000)

	// Pre-populate cache
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Get(i % 1000)
	}
}

func BenchmarkLRUCacheMixed(b *testing.B) {
	cache := New(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			cache.Put(i, i)
		} else {
			cache.Get(i % 1000)
		}
	}
}
