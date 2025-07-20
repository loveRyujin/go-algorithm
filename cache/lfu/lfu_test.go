package lfu

import (
	"sync"
	"testing"
)

func TestLFUCache(t *testing.T) {
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

	// Access key1 multiple times to increase its frequency
	cache.Get("key1")
	cache.Get("key1")

	// Now key1 has frequency 3, key2 has frequency 1
	// Adding key3 should evict key2 (lower frequency)
	cache.Put("key3", "value3")

	if _, ok := cache.Get("key2"); ok {
		t.Error("key2 should have been evicted (lower frequency)")
	}

	if value, ok := cache.Get("key1"); !ok || value != "value1" {
		t.Errorf("key1 should still be in cache, got %v", value)
	}

	if value, ok := cache.Get("key3"); !ok || value != "value3" {
		t.Errorf("key3 should be in cache, got %v", value)
	}
}

func TestLFUCacheUpdate(t *testing.T) {
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

func TestLFUCacheEvictionPolicy(t *testing.T) {
	cache := New(3)

	// Fill cache
	cache.Put(1, "one")
	cache.Put(2, "two")
	cache.Put(3, "three")

	// Access pattern: 1 accessed 3 times, 2 accessed 2 times, 3 accessed 1 time
	cache.Get(1)
	cache.Get(1)
	cache.Get(2)

	// Frequencies: 1->3, 2->2, 3->1
	if freq := cache.GetFrequency(1); freq != 3 {
		t.Errorf("Expected frequency 3 for key 1, got %d", freq)
	}
	if freq := cache.GetFrequency(2); freq != 2 {
		t.Errorf("Expected frequency 2 for key 2, got %d", freq)
	}
	if freq := cache.GetFrequency(3); freq != 1 {
		t.Errorf("Expected frequency 1 for key 3, got %d", freq)
	}

	// Add new key, should evict key 3 (lowest frequency)
	cache.Put(4, "four")

	if cache.Contains(3) {
		t.Error("key 3 should have been evicted (lowest frequency)")
	}

	// All other keys should still be present
	if !cache.Contains(1) {
		t.Error("key 1 should still be in cache")
	}
	if !cache.Contains(2) {
		t.Error("key 2 should still be in cache")
	}
	if !cache.Contains(4) {
		t.Error("key 4 should be in cache")
	}
}

func TestLFUCacheSameFrequencyEviction(t *testing.T) {
	cache := New(2)

	// Add two items with same frequency
	cache.Put("a", 1)
	cache.Put("b", 2)

	// Both have frequency 1, add third item
	// Should evict "a" (LRU among items with same frequency)
	cache.Put("c", 3)

	if cache.Contains("a") {
		t.Error("key 'a' should have been evicted (LRU among same frequency)")
	}

	if !cache.Contains("b") {
		t.Error("key 'b' should still be in cache")
	}

	if !cache.Contains("c") {
		t.Error("key 'c' should be in cache")
	}
}

func TestLFUCacheRemove(t *testing.T) {
	cache := New(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// Test removing existing key
	if !cache.Remove("b") {
		t.Error("Remove should return true for existing key")
	}

	if cache.Contains("b") {
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

func TestLFUCachePeek(t *testing.T) {
	cache := New(2)

	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	// Peek should not affect frequency
	if value, ok := cache.Peek("key1"); !ok || value != "value1" {
		t.Errorf("Peek should return value1, got %v", value)
	}

	// Frequency should still be 1
	if freq := cache.GetFrequency("key1"); freq != 1 {
		t.Errorf("Peek should not affect frequency, got %d", freq)
	}
}

func TestLFUCacheKeys(t *testing.T) {
	cache := New(4)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// Create different frequencies
	cache.Get("a") // freq 2
	cache.Get("b") // freq 2
	cache.Get("a") // freq 3

	keys := cache.Keys()

	// Check that we have the expected frequency groups
	if len(keys[1]) != 1 || keys[1][0] != "c" {
		t.Errorf("Expected [c] for frequency 1, got %v", keys[1])
	}

	if len(keys[2]) != 1 || keys[2][0] != "b" {
		t.Errorf("Expected [b] for frequency 2, got %v", keys[2])
	}

	if len(keys[3]) != 1 || keys[3][0] != "a" {
		t.Errorf("Expected [a] for frequency 3, got %v", keys[3])
	}
}

func TestLFUCacheClear(t *testing.T) {
	cache := New(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	cache.Clear()

	if cache.Len() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", cache.Len())
	}

	if cache.Contains("a") {
		t.Error("Cache should be empty after clear")
	}
}

func TestLFUCacheCapacity(t *testing.T) {
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

func TestLFUCacheZeroCapacity(t *testing.T) {
	cache := New(0)

	cache.Put("key", "value")
	
	if cache.Len() != 0 {
		t.Errorf("Expected length 0 for zero capacity cache, got %d", cache.Len())
	}
}

func TestLFUCacheMinFrequency(t *testing.T) {
	cache := New(3)

	cache.Put("a", 1)
	if cache.MinFrequency() != 1 {
		t.Errorf("Expected min frequency 1, got %d", cache.MinFrequency())
	}

	cache.Get("a") // frequency becomes 2
	cache.Put("b", 2) // new item with frequency 1

	if cache.MinFrequency() != 1 {
		t.Errorf("Expected min frequency 1, got %d", cache.MinFrequency())
	}

	cache.Get("b") // both items now have frequency 2
	if cache.MinFrequency() != 2 {
		t.Errorf("Expected min frequency 2, got %d", cache.MinFrequency())
	}
}

// Test concurrent access to ensure thread safety
func TestLFUCacheConcurrency(t *testing.T) {
	cache := New(100)
	var wg sync.WaitGroup
	numGoroutines := 10
	numOperations := 100

	// Start multiple goroutines to perform concurrent operations
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				key := id*numOperations + j

				// Mix of operations
				switch j % 5 {
				case 0:
					cache.Put(key, key*2)
				case 1:
					cache.Get(key)
				case 2:
					cache.Contains(key)
				case 3:
					cache.Peek(key)
				case 4:
					cache.GetFrequency(key)
				}
			}
		}(i)
	}

	wg.Wait()

	// Verify cache is in a consistent state
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache length %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

// Benchmark tests
func BenchmarkLFUCachePut(b *testing.B) {
	cache := New(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Put(i, i)
	}
}

func BenchmarkLFUCacheGet(b *testing.B) {
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

func BenchmarkLFUCacheMixed(b *testing.B) {
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
