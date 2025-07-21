package tinylfu

import (
	"hash/fnv"
	"math"
)

// countMinSketch implements a Count-Min Sketch for frequency estimation
type countMinSketch struct {
	width   int     // Number of buckets per row
	depth   int     // Number of hash functions
	table   [][]int // The sketch table
	size    int     // Total number of items added
	maxSize int     // Maximum size before reset
}

// newCountMinSketch creates a new Count-Min Sketch
func newCountMinSketch(maxSize int) *countMinSketch {
	// Use standard parameters: width = maxSize, depth = 4
	width := maxSize
	depth := 4

	table := make([][]int, depth)
	for i := range table {
		table[i] = make([]int, width)
	}

	return &countMinSketch{
		width:   width,
		depth:   depth,
		table:   table,
		maxSize: maxSize,
	}
}

// increment increases the count for a key
func (c *countMinSketch) increment(key any) {
	c.size++

	keyStr := c.keyToString(key)
	hashes := c.hash(keyStr)

	for i, hash := range hashes {
		bucket := int(hash) % c.width
		if bucket < 0 {
			bucket = -bucket
		}
		c.table[i][bucket]++
	}
}

// estimate returns the estimated frequency of a key
func (c *countMinSketch) estimate(key any) int {
	keyStr := c.keyToString(key)
	hashes := c.hash(keyStr)

	minCount := math.MaxInt32
	for i, hash := range hashes {
		bucket := int(hash) % c.width
		if bucket < 0 {
			bucket = -bucket
		}
		count := c.table[i][bucket]
		if count < minCount {
			minCount = count
		}
	}

	if minCount == math.MaxInt32 {
		return 0
	}
	return minCount
}

// reset clears the sketch with aging (divide all counts by 2)
func (c *countMinSketch) reset() {
	for i := range c.table {
		for j := range c.table[i] {
			c.table[i][j] = c.table[i][j] / 2
		}
	}
	c.size = c.size / 2
}

// hash generates multiple hash values for a key
func (c *countMinSketch) hash(key string) []uint32 {
	hashes := make([]uint32, c.depth)

	h := fnv.New32a()
	h.Write([]byte(key))
	hash1 := h.Sum32()

	h.Reset()
	h.Write([]byte(key + "salt"))
	hash2 := h.Sum32()

	// Generate multiple hashes using double hashing
	for i := 0; i < c.depth; i++ {
		hashes[i] = hash1 + uint32(i)*hash2
	}

	return hashes
}

// keyToString converts a key to string for hashing
func (c *countMinSketch) keyToString(key any) string {
	switch k := key.(type) {
	case string:
		return k
	case int:
		return string(rune(k))
	case int64:
		return string(rune(k))
	default:
		return ""
	}
}
