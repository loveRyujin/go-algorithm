package main

import (
	"fmt"

	"github.com/loveRyujin/go-algorithm/cache/lru"
)

func main() {
	// Create an LRU cache with capacity 3
	cache := lru.New(3)

	fmt.Println("=== LRU Cache Demo ===")

	// Add some key-value pairs
	fmt.Println("\n1. Adding key-value pairs:")
	cache.Put("apple", "苹果")
	cache.Put("banana", "香蕉")
	cache.Put("cherry", "樱桃")

	fmt.Printf("Cache size: %d/%d\n", cache.Len(), cache.Cap())
	fmt.Printf("Cache keys: %v\n", cache.Keys())

	// Get values
	fmt.Println("\n2. Getting values:")
	if value, ok := cache.Get("apple"); ok {
		fmt.Printf("apple: %s\n", value)
	}

	fmt.Printf("Key order after accessing apple: %v\n", cache.Keys())

	// Add new key, triggering eviction
	fmt.Println("\n3. Adding new key to trigger eviction:")
	cache.Put("date", "枣子")
	fmt.Printf("Keys after adding date: %v\n", cache.Keys())

	// Check evicted key
	if _, ok := cache.Get("banana"); !ok {
		fmt.Println("banana has been evicted")
	}

	// Update existing key
	fmt.Println("\n4. Updating existing key:")
	cache.Put("apple", "红苹果")
	if value, ok := cache.Get("apple"); ok {
		fmt.Printf("Updated apple: %s\n", value)
	}

	// Use Peek method (doesn't affect access order)
	fmt.Println("\n5. Using Peek method:")
	fmt.Printf("Key order before Peek: %v\n", cache.Keys())
	if value, ok := cache.Peek("cherry"); ok {
		fmt.Printf("Peek cherry: %s\n", value)
	}
	fmt.Printf("Key order after Peek: %v\n", cache.Keys())

	// Check if key exists
	fmt.Println("\n6. Checking key existence:")
	fmt.Printf("Contains apple: %v\n", cache.Contains("apple"))
	fmt.Printf("Contains banana: %v\n", cache.Contains("banana"))

	// Remove key
	fmt.Println("\n7. Removing key:")
	if cache.Remove("date") {
		fmt.Println("date has been removed")
	}
	fmt.Printf("Keys after removal: %v\n", cache.Keys())
	fmt.Printf("Cache size: %d/%d\n", cache.Len(), cache.Cap())

	// Clear cache
	fmt.Println("\n8. Clearing cache:")
	cache.Clear()
	fmt.Printf("Cache size after clear: %d/%d\n", cache.Len(), cache.Cap())
	fmt.Printf("Cache keys: %v\n", cache.Keys())

	// Demonstrate core LRU algorithm features
	fmt.Println("\n=== LRU Algorithm Features Demo ===")
	demoLRUBehavior()
}

func demoLRUBehavior() {
	cache := lru.New(3)

	fmt.Println("Create cache with capacity 3")

	// Add 3 elements
	cache.Put(1, "one")
	cache.Put(2, "two")
	cache.Put(3, "three")
	fmt.Printf("After adding 1,2,3: %v\n", cache.Keys())

	// Access element 1 to make it most recently used
	cache.Get(1)
	fmt.Printf("After accessing 1: %v\n", cache.Keys())

	// Add element 4, should evict element 2 (least recently used)
	cache.Put(4, "four")
	fmt.Printf("After adding 4: %v\n", cache.Keys())

	// Verify element 2 was evicted
	if !cache.Contains(2) {
		fmt.Println("Element 2 has been evicted (LRU policy)")
	}

	// Access element 3
	cache.Get(3)
	fmt.Printf("After accessing 3: %v\n", cache.Keys())

	// Add element 5, should evict element 1
	cache.Put(5, "five")
	fmt.Printf("After adding 5: %v\n", cache.Keys())

	if !cache.Contains(1) {
		fmt.Println("Element 1 has been evicted (LRU policy)")
	}
}
