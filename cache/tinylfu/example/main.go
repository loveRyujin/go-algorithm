package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/loveRyujin/go-algorithm/cache/tinylfu"
)

func main() {
	fmt.Println("=== TinyLFU Cache Demo ===")

	// Create a TinyLFU cache with capacity 10
	cache := tinylfu.New(10)

	// Basic operations demo
	basicOperationsDemo(cache)
	fmt.Println()

	// Admission policy demo
	admissionPolicyDemo()
	fmt.Println()

	// Performance comparison demo
	performanceDemo()
}

func basicOperationsDemo(cache *tinylfu.Cache) {
	fmt.Println("1. Basic Operations:")

	// Add some key-value pairs
	cache.Put("apple", "red fruit")
	cache.Put("banana", "yellow fruit")
	cache.Put("cherry", "red fruit")

	fmt.Printf("   Cache stats after adding 3 items: %v\n", cache.Stats())

	// Get values
	if value, ok := cache.Get("apple"); ok {
		fmt.Printf("   apple: %s\n", value)
	}

	// Access some keys multiple times to build frequency
	for i := 0; i < 5; i++ {
		cache.Get("apple")
		cache.Get("banana")
	}
	cache.Get("cherry") // Less frequent access

	fmt.Printf("   Frequencies - apple: %d, banana: %d, cherry: %d\n",
		cache.GetFrequency("apple"),
		cache.GetFrequency("banana"),
		cache.GetFrequency("cherry"))

	// Add more items to trigger eviction
	for i := 1; i <= 10; i++ {
		cache.Put(fmt.Sprintf("item%d", i), fmt.Sprintf("value%d", i))
	}

	fmt.Printf("   Final cache stats: %v\n", cache.Stats())
	fmt.Printf("   Contains apple: %v, cherry: %v\n",
		cache.Contains("apple"), cache.Contains("cherry"))
}

func admissionPolicyDemo() {
	fmt.Println("2. Admission Policy Demo:")

	cache := tinylfu.New(8)

	// Phase 1: Populate cache with initial data
	fmt.Println("   Phase 1: Adding initial data...")
	for i := 1; i <= 8; i++ {
		cache.Put(fmt.Sprintf("init%d", i), i)
	}

	// Phase 2: Create frequency differences
	fmt.Println("   Phase 2: Creating frequency patterns...")

	// Hot data - accessed frequently
	hotKeys := []string{"init1", "init2", "init3"}
	for _, key := range hotKeys {
		for i := 0; i < 20; i++ {
			cache.Get(key)
		}
	}

	// Warm data - accessed moderately
	warmKeys := []string{"init4", "init5"}
	for _, key := range warmKeys {
		for i := 0; i < 5; i++ {
			cache.Get(key)
		}
	}

	// Cold data - rarely accessed (init6, init7, init8)

	fmt.Printf("   Frequencies before new arrivals:\n")
	for i := 1; i <= 8; i++ {
		key := fmt.Sprintf("init%d", i)
		fmt.Printf("     %s: %d\n", key, cache.GetFrequency(key))
	}

	// Phase 3: Add new data to test admission policy
	fmt.Println("   Phase 3: Adding new data to test admission...")
	newKeys := []string{"new1", "new2", "new3", "new4", "new5"}

	for _, key := range newKeys {
		cache.Put(key, "new_value")
		cache.Get(key) // Access once to build some frequency
	}

	fmt.Printf("   Final cache stats: %v\n", cache.Stats())

	// Check which keys survived
	fmt.Println("   Survival check:")
	allKeys := append([]string{"init1", "init2", "init3", "init4", "init5", "init6", "init7", "init8"}, newKeys...)
	survivors := 0
	for _, key := range allKeys {
		if cache.Contains(key) {
			fmt.Printf("     ✓ %s (freq: %d)\n", key, cache.GetFrequency(key))
			survivors++
		} else {
			fmt.Printf("     ✗ %s (evicted)\n", key)
		}
	}
	fmt.Printf("   Total survivors: %d/%d\n", survivors, len(allKeys))
}

func performanceDemo() {
	fmt.Println("3. Performance Characteristics Demo:")

	capacities := []int{100, 1000, 10000}

	for _, capacity := range capacities {
		cache := tinylfu.New(capacity)

		// Warm up phase
		for i := 0; i < capacity; i++ {
			cache.Put(i, i*2)
		}

		// Create access pattern: 80% hot data, 20% new data
		numOperations := capacity * 10
		hitCount := 0

		start := time.Now()

		for i := 0; i < numOperations; i++ {
			if rand.Float32() < 0.8 {
				// Access hot data (first 20% of keys)
				key := rand.Intn(capacity / 5)
				if _, ok := cache.Get(key); ok {
					hitCount++
				}
			} else {
				// Add new data
				newKey := capacity + rand.Intn(capacity)
				cache.Put(newKey, newKey*2)
			}
		}

		duration := time.Since(start)
		readOps := float64(numOperations) * 0.8
		hitRate := float64(hitCount) / readOps * 100 // Only consider get operations

		fmt.Printf("   Capacity %d: %d operations in %v (%.2f%% hit rate)\n",
			capacity, numOperations, duration, hitRate)

		stats := cache.Stats()
		fmt.Printf("     Segment distribution - Window: %d, Probation: %d, Protected: %d\n",
			stats["window_size"], stats["probation_size"], stats["protected_size"])
	}
}

// Additional demo: Zipfian distribution simulation
func zipfianDemo() {
	fmt.Println("4. Zipfian Distribution Simulation:")

	cache := tinylfu.New(100)
	numKeys := 1000
	numOperations := 10000

	// Initialize with data
	for i := 0; i < numKeys; i++ {
		cache.Put(i, fmt.Sprintf("value_%d", i))
	}

	hitCount := 0

	// Simulate Zipfian access pattern (popular items accessed more frequently)
	start := time.Now()
	for i := 0; i < numOperations; i++ {
		// Zipfian: first few items are accessed much more frequently
		var key int
		r := rand.Float64()
		if r < 0.5 {
			key = rand.Intn(10) // Top 10 items (50% of accesses)
		} else if r < 0.8 {
			key = 10 + rand.Intn(40) // Next 40 items (30% of accesses)
		} else {
			key = 50 + rand.Intn(950) // Remaining items (20% of accesses)
		}

		if _, ok := cache.Get(key); ok {
			hitCount++
		}
	}
	duration := time.Since(start)

	hitRate := float64(hitCount) / float64(numOperations) * 100
	fmt.Printf("   Zipfian simulation: %.2f%% hit rate over %d operations in %v\n",
		hitRate, numOperations, duration)

	// Show frequency distribution of top items
	fmt.Println("   Top 10 key frequencies:")
	for i := 0; i < 10; i++ {
		fmt.Printf("     key_%d: %d\n", i, cache.GetFrequency(i))
	}
}
