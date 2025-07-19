package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/loveRyujin/go-algorithm/cache/lru"
)

func main() {
	fmt.Println("=== LRU Cache Concurrent Demo ===")

	// Create a cache with capacity 100
	cache := lru.New(100)

	// Demo 1: Concurrent Put operations
	fmt.Println("\n1. Concurrent Put operations:")
	concurrentPutDemo(cache)

	// Demo 2: Concurrent Get operations
	fmt.Println("\n2. Concurrent Get operations:")
	concurrentGetDemo(cache)

	// Demo 3: Mixed concurrent operations
	fmt.Println("\n3. Mixed concurrent operations:")
	mixedConcurrentDemo(cache)
}

func concurrentPutDemo(cache *lru.Cache) {
	var wg sync.WaitGroup
	numWorkers := 10
	itemsPerWorker := 20

	start := time.Now()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < itemsPerWorker; j++ {
				key := fmt.Sprintf("worker%d_item%d", workerID, j)
				value := workerID*1000 + j
				cache.Put(key, value)
			}
		}(i)
	}

	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("Added %d items concurrently in %v\n", numWorkers*itemsPerWorker, duration)
	fmt.Printf("Final cache size: %d\n", cache.Len())
}

func concurrentGetDemo(cache *lru.Cache) {
	// First, populate the cache
	for i := 0; i < 50; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i*10)
	}

	var wg sync.WaitGroup
	numReaders := 20
	readsPerReader := 50

	start := time.Now()

	for i := 0; i < numReaders; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			for j := 0; j < readsPerReader; j++ {
				key := fmt.Sprintf("key%d", j%50)
				if value, ok := cache.Get(key); ok {
					_ = value // Use the value to avoid optimization
				}
			}
		}(i)
	}

	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("Performed %d reads concurrently in %v\n", numReaders*readsPerReader, duration)
}

func mixedConcurrentDemo(cache *lru.Cache) {
	var wg sync.WaitGroup
	numOperations := 1000

	start := time.Now()

	// Writer goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(writerID int) {
			defer wg.Done()
			for j := 0; j < numOperations/5; j++ {
				key := fmt.Sprintf("mixed_key_%d_%d", writerID, j)
				cache.Put(key, writerID*1000+j)
			}
		}(i)
	}

	// Reader goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			for j := 0; j < numOperations/10; j++ {
				// Try to read keys that might exist
				key := fmt.Sprintf("mixed_key_%d_%d", readerID%5, j)
				cache.Get(key)

				// Also try Peek operations
				cache.Peek(key)

				// Check if key exists
				cache.Contains(key)
			}
		}(i)
	}

	// Metadata reader goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < numOperations/10; i++ {
			cache.Len()
			cache.Keys()
			time.Sleep(time.Microsecond) // Small delay
		}
	}()

	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("Mixed operations completed in %v\n", duration)
	fmt.Printf("Final cache size: %d\n", cache.Len())
}
