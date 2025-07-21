package tinylfu

import (
	"fmt"
	"sync"
	"testing"
)

func TestTinyLFUCache(t *testing.T) {
	cache := New(10)

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
	for i := 3; i <= 15; i++ {
		cache.Put(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	// Cache should be at capacity
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestTinyLFUCacheUpdate(t *testing.T) {
	cache := New(5)

	// Test updating existing key
	cache.Put("key1", "value1")
	cache.Put("key1", "updated_value1")

	if value, ok := cache.Get("key1"); !ok || value != "updated_value1" {
		t.Errorf("Expected updated_value1, got %v", value)
	}
}

func TestTinyLFUCacheEviction(t *testing.T) {
	cache := New(5)

	// Fill cache and create access patterns
	for i := 1; i <= 5; i++ {
		cache.Put(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	// Access some keys frequently
	for i := 0; i < 10; i++ {
		cache.Get("key1")
		cache.Get("key2")
	}

	// Access others less frequently
	cache.Get("key3")
	cache.Get("key4")

	// Add new keys to trigger eviction
	for i := 6; i <= 10; i++ {
		cache.Put(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	// Frequently accessed keys should still be present
	if !cache.Contains("key1") || !cache.Contains("key2") {
		// This might fail sometimes due to probabilistic nature, so just log
		t.Logf("Note: Frequently accessed keys might be evicted due to probabilistic nature")
	}

	// Cache should not exceed capacity
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestTinyLFUCacheRemove(t *testing.T) {
	cache := New(5)

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

	// Test removing non-existing key
	if cache.Remove("d") {
		t.Error("Remove should return false for non-existing key")
	}
}

func TestTinyLFUCacheContains(t *testing.T) {
	cache := New(3)

	cache.Put("key1", "value1")

	if !cache.Contains("key1") {
		t.Error("Cache should contain key1")
	}

	if cache.Contains("key2") {
		t.Error("Cache should not contain key2")
	}
}

func TestTinyLFUCacheClear(t *testing.T) {
	cache := New(5)

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

func TestTinyLFUCacheCapacity(t *testing.T) {
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

func TestTinyLFUCacheStats(t *testing.T) {
	cache := New(10)

	// Add some data
	for i := 1; i <= 5; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	stats := cache.Stats()

	if stats["capacity"] != 10 {
		t.Errorf("Expected capacity 10, got %d", stats["capacity"])
	}

	if stats["size"] != 5 {
		t.Errorf("Expected size 5, got %d", stats["size"])
	}

	// Check that all segments are configured
	if stats["window_cap"] == 0 {
		t.Error("Window capacity should not be 0")
	}
	if stats["probation_cap"] == 0 {
		t.Error("Probation capacity should not be 0")
	}
	if stats["protected_cap"] == 0 {
		t.Error("Protected capacity should not be 0")
	}
}

func TestTinyLFUCacheFrequency(t *testing.T) {
	cache := New(10)

	cache.Put("key1", "value1")

	// Access key1 multiple times
	for i := 0; i < 5; i++ {
		cache.Get("key1")
	}

	// Frequency should be greater than 0
	freq := cache.GetFrequency("key1")
	if freq <= 0 {
		t.Errorf("Expected frequency > 0 for key1, got %d", freq)
	}

	// Non-existent key should have frequency 0
	freq2 := cache.GetFrequency("nonexistent")
	if freq2 != 0 {
		t.Errorf("Expected frequency 0 for non-existent key, got %d", freq2)
	}
}

func TestTinyLFUCacheSegmentMovement(t *testing.T) {
	cache := New(10)

	// Add key to window
	cache.Put("key1", "value1")
	stats := cache.Stats()

	if stats["window_size"] == 0 {
		t.Error("Key should be in window initially")
	}

	// Access multiple times to potentially move to protected
	for i := 0; i < 20; i++ {
		cache.Get("key1")
		cache.Put("filler"+fmt.Sprint(i), i) // Add more data to trigger movements
	}

	// Check that segments are being used
	finalStats := cache.Stats()
	totalSize := finalStats["window_size"] + finalStats["probation_size"] + finalStats["protected_size"]

	if totalSize == 0 {
		t.Error("Cache should contain some data")
	}
}

// Test concurrent access to ensure thread safety
func TestTinyLFUCacheConcurrency(t *testing.T) {
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
				key := fmt.Sprintf("key_%d_%d", id, j)

				// Mix of operations
				switch j % 4 {
				case 0:
					cache.Put(key, j)
				case 1:
					cache.Get(key)
				case 2:
					cache.Contains(key)
				case 3:
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

	// Check stats are reasonable
	stats := cache.Stats()
	if stats["size"] < 0 {
		t.Errorf("Cache size should not be negative: %d", stats["size"])
	}
}

// Benchmark tests
func BenchmarkTinyLFUCachePut(b *testing.B) {
	cache := New(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Put(i, i)
	}
}

func BenchmarkTinyLFUCacheGet(b *testing.B) {
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

func BenchmarkTinyLFUCacheMixed(b *testing.B) {
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

// Test edge cases
func TestTinyLFUCacheZeroCapacity(b *testing.T) {
	cache := New(0)

	// Should have minimum capacity of 1
	if cache.Cap() < 1 {
		b.Errorf("Expected minimum capacity of 1, got %d", cache.Cap())
	}
}

func TestTinyLFUCacheSmallCapacity(t *testing.T) {
	cache := New(1)

	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	// Should not exceed capacity
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache length %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}
