package main

import (
	"fmt"

	"github.com/loveRyujin/go-algorithm/cache/lfu"
)

func main() {
	fmt.Println("=== LFU Cache Demo ===")

	// Create an LFU cache with capacity 3
	cache := lfu.New(3)

	// Add some key-value pairs
	fmt.Println("\n1. Adding key-value pairs:")
	cache.Put("apple", "red")
	cache.Put("banana", "yellow")
	cache.Put("cherry", "red")

	fmt.Printf("Cache size: %d/%d\n", cache.Len(), cache.Cap())
	printCacheState(cache)

	// Access keys to create different frequencies
	fmt.Println("\n2. Creating different access frequencies:")
	cache.Get("apple")  // apple: freq 2
	cache.Get("apple")  // apple: freq 3
	cache.Get("banana") // banana: freq 2

	fmt.Println("After accessing apple twice and banana once:")
	printCacheState(cache)
	printFrequencies(cache)

	// Add new key, should evict cherry (lowest frequency)
	fmt.Println("\n3. Adding new key to trigger eviction:")
	cache.Put("date", "brown")

	fmt.Printf("After adding 'date':\n")
	printCacheState(cache)

	if !cache.Contains("cherry") {
		fmt.Println("'cherry' was evicted (lowest frequency)")
	}

	// Update existing key
	fmt.Println("\n4. Updating existing key:")
	cache.Put("apple", "green apple")
	if value, ok := cache.Get("apple"); ok {
		fmt.Printf("Updated apple: %s (frequency: %d)\n", value, cache.GetFrequency("apple"))
	}

	// Use Peek method (doesn't affect frequency)
	fmt.Println("\n5. Using Peek method:")
	fmt.Printf("Frequencies before peek: ")
	printFrequencies(cache)

	if value, ok := cache.Peek("banana"); ok {
		fmt.Printf("Peek banana: %s\n", value)
	}

	fmt.Printf("Frequencies after peek (unchanged): ")
	printFrequencies(cache)

	// Remove key
	fmt.Println("\n6. Removing key:")
	if cache.Remove("date") {
		fmt.Println("'date' has been removed")
	}

	fmt.Printf("Cache size after removal: %d/%d\n", cache.Len(), cache.Cap())
	printCacheState(cache)

	// Demonstrate LFU eviction policy
	fmt.Println("\n=== LFU Eviction Policy Demo ===")
	demoLFUEviction()
}

func printCacheState(cache *lfu.Cache) {
	keys := cache.Keys()
	fmt.Printf("Cache contents by frequency:\n")
	for freq, keyList := range keys {
		fmt.Printf("  Frequency %d: %v\n", freq, keyList)
	}
}

func printFrequencies(cache *lfu.Cache) {
	keys := cache.Keys()
	fmt.Printf("{ ")
	for freq, keyList := range keys {
		for _, key := range keyList {
			fmt.Printf("%v:%d ", key, freq)
		}
	}
	fmt.Printf("}\n")
}

func demoLFUEviction() {
	cache := lfu.New(3)

	fmt.Println("Create cache with capacity 3")

	// Add 3 elements
	cache.Put("A", "value_A")
	cache.Put("B", "value_B")
	cache.Put("C", "value_C")
	fmt.Printf("Added A, B, C: ")
	printFrequencies(cache)

	// Create different access patterns
	cache.Get("A") // A: freq 2
	cache.Get("A") // A: freq 3
	cache.Get("B") // B: freq 2
	fmt.Printf("After accessing A twice, B once: ")
	printFrequencies(cache)
	fmt.Printf("Min frequency: %d\n", cache.MinFrequency())

	// Add new element, should evict C (lowest frequency)
	cache.Put("D", "value_D")
	fmt.Printf("After adding D: ")
	printFrequencies(cache)

	if !cache.Contains("C") {
		fmt.Println("C was evicted (lowest frequency)")
	}

	// Add another element
	cache.Get("D")            // D: freq 2
	cache.Put("E", "value_E") // Should evict E based on LFU+LRU
	fmt.Printf("After accessing D and adding E: ")
	printFrequencies(cache)

	// Show final state
	fmt.Printf("Final min frequency: %d\n", cache.MinFrequency())
	fmt.Println("Final cache state:")
	printCacheState(cache)
}
