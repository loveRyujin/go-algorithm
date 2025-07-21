package tinylfu

import (
	"hash/fnv"
	"math"
)

// bloomFilter implements a simple bloom filter
type bloomFilter struct {
	bitArray []bool
	size     int
	numHash  int
}

// newBloomFilter creates a new bloom filter
func newBloomFilter(expectedElements int) *bloomFilter {
	// Calculate optimal size and number of hash functions
	// For 1% false positive rate: m = -n*ln(p) / (ln(2)^2)
	size := int(math.Ceil(-float64(expectedElements) * math.Log(0.01) / (math.Ln2 * math.Ln2)))
	numHash := int(math.Ceil(float64(size) / float64(expectedElements) * math.Ln2))

	// Limit number of hash functions
	if numHash > 10 {
		numHash = 10
	}

	return &bloomFilter{
		bitArray: make([]bool, size),
		size:     size,
		numHash:  numHash,
	}
}

// add adds a key to the bloom filter
func (b *bloomFilter) add(key any) {
	keyStr := b.keyToString(key)
	hashes := b.hash(keyStr)

	for _, hash := range hashes {
		index := int(hash) % b.size
		if index < 0 {
			index = -index
		}
		b.bitArray[index] = true
	}
}

// mightContain checks if a key might be in the set
func (b *bloomFilter) mightContain(key any) bool {
	keyStr := b.keyToString(key)
	hashes := b.hash(keyStr)

	for _, hash := range hashes {
		index := int(hash) % b.size
		if index < 0 {
			index = -index
		}
		if !b.bitArray[index] {
			return false
		}
	}
	return true
}

// reset clears the bloom filter
func (b *bloomFilter) reset() {
	for i := range b.bitArray {
		b.bitArray[i] = false
	}
}

// hash generates multiple hash values for a key
func (b *bloomFilter) hash(key string) []uint32 {
	hashes := make([]uint32, b.numHash)

	h := fnv.New32a()
	h.Write([]byte(key))
	hash1 := h.Sum32()

	h.Reset()
	h.Write([]byte(key + "bloom"))
	hash2 := h.Sum32()

	// Generate multiple hashes using double hashing
	for i := 0; i < b.numHash; i++ {
		hashes[i] = hash1 + uint32(i)*hash2
	}

	return hashes
}

// keyToString converts a key to string for hashing
func (b *bloomFilter) keyToString(key any) string {
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
