package lru

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := New(2)

	// 测试基本的Put和Get操作
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	if value, ok := cache.Get("key1"); !ok || value != "value1" {
		t.Errorf("Expected value1, got %v", value)
	}

	if value, ok := cache.Get("key2"); !ok || value != "value2" {
		t.Errorf("Expected value2, got %v", value)
	}

	// 测试容量限制
	cache.Put("key3", "value3") // 这应该会淘汰key1，因为key2最近被访问过

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

func TestLRUCacheUpdate(t *testing.T) {
	cache := New(2)

	// 测试更新已存在的键
	cache.Put("key1", "value1")
	cache.Put("key1", "updated_value1")

	if value, ok := cache.Get("key1"); !ok || value != "updated_value1" {
		t.Errorf("Expected updated_value1, got %v", value)
	}

	// 缓存应该还有空间
	if cache.Len() != 1 {
		t.Errorf("Expected length 1, got %d", cache.Len())
	}
}

func TestLRUCacheEviction(t *testing.T) {
	cache := New(3)

	// 填满缓存
	cache.Put(1, "one")
	cache.Put(2, "two")
	cache.Put(3, "three")

	// 访问key 1，使其成为最近使用的
	cache.Get(1)

	// 添加新的键，应该淘汰key 2
	cache.Put(4, "four")

	if _, ok := cache.Get(2); ok {
		t.Error("key 2 should have been evicted")
	}

	// 检查其他键是否还在
	if _, ok := cache.Get(1); !ok {
		t.Error("key 1 should still be in cache")
	}

	if _, ok := cache.Get(3); !ok {
		t.Error("key 3 should still be in cache")
	}

	if _, ok := cache.Get(4); !ok {
		t.Error("key 4 should be in cache")
	}
}

func TestLRUCacheRemove(t *testing.T) {
	cache := New(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// 测试删除存在的键
	if !cache.Remove("b") {
		t.Error("Remove should return true for existing key")
	}

	if _, ok := cache.Get("b"); ok {
		t.Error("key b should have been removed")
	}

	if cache.Len() != 2 {
		t.Errorf("Expected length 2, got %d", cache.Len())
	}

	// 测试删除不存在的键
	if cache.Remove("d") {
		t.Error("Remove should return false for non-existing key")
	}
}

func TestLRUCacheContains(t *testing.T) {
	cache := New(2)

	cache.Put("key1", "value1")

	if !cache.Contains("key1") {
		t.Error("Cache should contain key1")
	}

	if cache.Contains("key2") {
		t.Error("Cache should not contain key2")
	}
}

func TestLRUCachePeek(t *testing.T) {
	cache := New(2)

	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	// Peek不应该影响访问顺序
	if value, ok := cache.Peek("key1"); !ok || value != "value1" {
		t.Errorf("Peek should return value1, got %v", value)
	}

	// 添加新键，key1应该被淘汰（因为Peek没有更新访问顺序）
	cache.Put("key3", "value3")

	if _, ok := cache.Get("key1"); ok {
		t.Error("key1 should have been evicted")
	}
}

func TestLRUCacheKeys(t *testing.T) {
	cache := New(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// 访问a，使其成为最近使用的
	cache.Get("a")

	keys := cache.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	// 检查键的顺序（最近使用的在前）
	if keys[0] != "a" {
		t.Errorf("Expected first key to be 'a', got %v", keys[0])
	}
}

func TestLRUCacheClear(t *testing.T) {
	cache := New(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	cache.Clear()

	if cache.Len() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", cache.Len())
	}

	if _, ok := cache.Get("a"); ok {
		t.Error("Cache should be empty after clear")
	}
}

func TestLRUCacheCapacity(t *testing.T) {
	cache := New(5)

	if cache.Cap() != 5 {
		t.Errorf("Expected capacity 5, got %d", cache.Cap())
	}

	// 添加元素不应该改变容量
	cache.Put("key", "value")
	if cache.Cap() != 5 {
		t.Errorf("Capacity should remain 5, got %d", cache.Cap())
	}
}

// 基准测试
func BenchmarkLRUCachePut(b *testing.B) {
	cache := New(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Put(i, i)
	}
}

func BenchmarkLRUCacheGet(b *testing.B) {
	cache := New(1000)

	// 预填充缓存
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Get(i % 1000)
	}
}

func BenchmarkLRUCacheMixed(b *testing.B) {
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
