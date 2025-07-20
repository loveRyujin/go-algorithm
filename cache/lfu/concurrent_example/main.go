package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/loveRyujin/go-algorithm/cache/lfu"
)

func main() {
	fmt.Println("=== LFU Cache Concurrent Access Demo ===")

	cache := lfu.New(50)
	var wg sync.WaitGroup

	// Statistics
	var readCount, writeCount, hitCount, missCount int64
	var mu sync.Mutex

	// Number of concurrent workers
	numWorkers := 10
	operationsPerWorker := 1000

	fmt.Printf("Starting %d workers, each performing %d operations...\n", numWorkers, operationsPerWorker)

	start := time.Now()

	// Start multiple workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, cache, operationsPerWorker, &wg, &readCount, &writeCount, &hitCount, &missCount, &mu)
	}

	// Wait for all workers to complete
	wg.Wait()
	duration := time.Since(start)

	// Print results
	fmt.Printf("\n=== Results ===\n")
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("Total operations: %d\n", readCount+writeCount)
	fmt.Printf("Read operations: %d\n", readCount)
	fmt.Printf("Write operations: %d\n", writeCount)
	fmt.Printf("Cache hits: %d\n", hitCount)
	fmt.Printf("Cache misses: %d\n", missCount)
	fmt.Printf("Hit ratio: %.2f%%\n", float64(hitCount)/float64(readCount)*100)
	fmt.Printf("Final cache size: %d/%d\n", cache.Len(), cache.Cap())

	// Show frequency distribution
	fmt.Printf("\n=== Frequency Distribution ===\n")
	keys := cache.Keys()
	for freq, keyList := range keys {
		fmt.Printf("Frequency %d: %d items\n", freq, len(keyList))
	}

	fmt.Printf("Minimum frequency: %d\n", cache.MinFrequency())
}

func worker(id int, cache *lfu.Cache, operations int, wg *sync.WaitGroup, 
	readCount, writeCount, hitCount, missCount *int64, mu *sync.Mutex) {
	
	defer wg.Done()

	r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(id)))

	for i := 0; i < operations; i++ {
		// 70% reads, 30% writes
		if r.Float32() < 0.7 {
			// Read operation
			key := fmt.Sprintf("key_%d", r.Intn(100))
			_, found := cache.Get(key)
			
			mu.Lock()
			(*readCount)++
			if found {
				(*hitCount)++
			} else {
				(*missCount)++
			}
			mu.Unlock()
		} else {
			// Write operation
			key := fmt.Sprintf("key_%d", r.Intn(100))
			value := fmt.Sprintf("value_%d_%d", id, i)
			cache.Put(key, value)
			
			mu.Lock()
			(*writeCount)++
			mu.Unlock()
		}

		// Occasionally perform other operations
		if i%50 == 0 {
			switch r.Intn(4) {
			case 0:
				key := fmt.Sprintf("key_%d", r.Intn(100))
				cache.Contains(key)
			case 1:
				key := fmt.Sprintf("key_%d", r.Intn(100))
				cache.Peek(key)
			case 2:
				key := fmt.Sprintf("key_%d", r.Intn(100))
				cache.GetFrequency(key)
			case 3:
				cache.MinFrequency()
			}
		}
	}
}
