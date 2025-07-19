package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/loveRyujin/go-algorithm/cache/lru"
)

func main() {
	fmt.Println("=== LRU Cache Implementation Comparison ===")
	fmt.Println()

	// Test both implementations
	testRWMutexImplementation()
	fmt.Println()
	testSyncMapImplementation()
	fmt.Println()
	comparePerformance()
}

func testRWMutexImplementation() {
	fmt.Println("1. RWMutex Implementation:")
	cache := lru.New(3)

	// Basic operations
	cache.Put("apple", "red")
	cache.Put("banana", "yellow")
	cache.Put("cherry", "red")

	fmt.Printf("   Keys: %v\n", cache.Keys())

	if value, ok := cache.Get("apple"); ok {
		fmt.Printf("   Get apple: %v\n", value)
	}

	fmt.Printf("   Keys after accessing apple: %v\n", cache.Keys())

	// Test eviction
	cache.Put("date", "brown")
	fmt.Printf("   Keys after adding date: %v\n", cache.Keys())
}

func testSyncMapImplementation() {
	fmt.Println("2. sync.Map Implementation:")
	cache := lru.NewSyncMap(3)

	// Basic operations
	cache.Put("apple", "red")
	cache.Put("banana", "yellow")
	cache.Put("cherry", "red")

	fmt.Printf("   Keys: %v\n", cache.Keys())

	if value, ok := cache.Get("apple"); ok {
		fmt.Printf("   Get apple: %v\n", value)
	}

	fmt.Printf("   Keys after accessing apple: %v\n", cache.Keys())

	// Test eviction
	cache.Put("date", "brown")
	fmt.Printf("   Keys after adding date: %v\n", cache.Keys())
}

func comparePerformance() {
	fmt.Println("3. Performance Comparison (Concurrent Access):")

	const numGoroutines = 50
	const numOperations = 1000

	// Test RWMutex implementation
	rwCache := lru.New(100)
	start := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				if j%3 == 0 {
					rwCache.Put(id*numOperations+j, j)
				} else {
					rwCache.Get(id*100 + j%100)
				}
			}
		}(i)
	}
	wg.Wait()
	rwDuration := time.Since(start)

	// Test sync.Map implementation
	syncCache := lru.NewSyncMap(100)
	start = time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				if j%3 == 0 {
					syncCache.Put(id*numOperations+j, j)
				} else {
					syncCache.Get(id*100 + j%100)
				}
			}
		}(i)
	}
	wg.Wait()
	syncDuration := time.Since(start)

	fmt.Printf("   RWMutex implementation: %v\n", rwDuration)
	fmt.Printf("   sync.Map implementation: %v\n", syncDuration)

	if rwDuration < syncDuration {
		fmt.Printf("   RWMutex is %.1f%% faster\n", float64(syncDuration-rwDuration)/float64(syncDuration)*100)
	} else {
		fmt.Printf("   sync.Map is %.1f%% faster\n", float64(rwDuration-syncDuration)/float64(rwDuration)*100)
	}

	fmt.Printf("\n   Final cache sizes - RWMutex: %d, sync.Map: %d\n", rwCache.Len(), syncCache.Len())
}
