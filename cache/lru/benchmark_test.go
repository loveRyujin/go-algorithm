package lru

import (
	"sync"
	"testing"
)

// Benchmark tests for sync.Map implementation
func BenchmarkSyncMapCachePut(b *testing.B) {
	cache := NewSyncMap(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Put(i, i)
	}
}

func BenchmarkSyncMapCacheGet(b *testing.B) {
	cache := NewSyncMap(1000)

	// Pre-populate cache
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Get(i % 1000)
	}
}

func BenchmarkSyncMapCacheMixed(b *testing.B) {
	cache := NewSyncMap(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			cache.Put(i, i)
		} else {
			cache.Get(i % 1000)
		}
	}
}

// Concurrent benchmarks comparing RWMutex vs sync.Map
func BenchmarkRWMutexCacheConcurrentRead(b *testing.B) {
	cache := New(1000)

	// Pre-populate cache
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.Get(42) // Read same key to test concurrent reads
		}
	})
}

func BenchmarkSyncMapCacheConcurrentRead(b *testing.B) {
	cache := NewSyncMap(1000)

	// Pre-populate cache
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.Get(42) // Read same key to test concurrent reads
		}
	})
}

func BenchmarkRWMutexCacheConcurrentWrite(b *testing.B) {
	cache := New(1000)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			cache.Put(i, i)
			i++
		}
	})
}

func BenchmarkSyncMapCacheConcurrentWrite(b *testing.B) {
	cache := NewSyncMap(1000)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			cache.Put(i, i)
			i++
		}
	})
}

func BenchmarkRWMutexCacheConcurrentReadWrite(b *testing.B) {
	cache := New(1000)

	// Pre-populate cache
	for i := 0; i < 500; i++ {
		cache.Put(i, i)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%3 == 0 {
				cache.Put(i, i)
			} else {
				cache.Get(i % 500)
			}
			i++
		}
	})
}

func BenchmarkSyncMapCacheConcurrentReadWrite(b *testing.B) {
	cache := NewSyncMap(1000)

	// Pre-populate cache
	for i := 0; i < 500; i++ {
		cache.Put(i, i)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%3 == 0 {
				cache.Put(i, i)
			} else {
				cache.Get(i % 500)
			}
			i++
		}
	})
}

// Test sync.Map implementation for correctness
func TestSyncMapCache(t *testing.T) {
	cache := NewSyncMap(2)

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

// Test sync.Map concurrent access
func TestSyncMapCacheConcurrency(t *testing.T) {
	cache := NewSyncMap(100)
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
				switch j % 4 {
				case 0:
					cache.Put(key, key*2)
				case 1:
					cache.Get(key)
				case 2:
					cache.Contains(key)
				case 3:
					cache.Peek(key)
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
