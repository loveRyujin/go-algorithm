package main

import (
	"fmt"

	"github.com/loveRyujin/go-algorithm/cache/lru"
)

func main() {
	// 创建一个容量为3的LRU缓存
	cache := lru.New(3)

	fmt.Println("=== LRU Cache Demo ===")

	// 添加一些键值对
	fmt.Println("\n1. 添加键值对:")
	cache.Put("apple", "苹果")
	cache.Put("banana", "香蕉")
	cache.Put("cherry", "樱桃")

	fmt.Printf("缓存大小: %d/%d\n", cache.Len(), cache.Cap())
	fmt.Printf("缓存键: %v\n", cache.Keys())

	// 获取值
	fmt.Println("\n2. 获取值:")
	if value, ok := cache.Get("apple"); ok {
		fmt.Printf("apple: %s\n", value)
	}

	fmt.Printf("访问apple后的键顺序: %v\n", cache.Keys())

	// 添加新键，触发淘汰
	fmt.Println("\n3. 添加新键触发淘汰:")
	cache.Put("date", "枣子")
	fmt.Printf("添加date后的键: %v\n", cache.Keys())

	// 检查被淘汰的键
	if _, ok := cache.Get("banana"); !ok {
		fmt.Println("banana 已被淘汰")
	}

	// 更新已存在的键
	fmt.Println("\n4. 更新已存在的键:")
	cache.Put("apple", "红苹果")
	if value, ok := cache.Get("apple"); ok {
		fmt.Printf("更新后的apple: %s\n", value)
	}

	// 使用Peek方法（不影响访问顺序）
	fmt.Println("\n5. 使用Peek方法:")
	fmt.Printf("Peek前的键顺序: %v\n", cache.Keys())
	if value, ok := cache.Peek("cherry"); ok {
		fmt.Printf("Peek cherry: %s\n", value)
	}
	fmt.Printf("Peek后的键顺序: %v\n", cache.Keys())

	// 检查是否包含某个键
	fmt.Println("\n6. 检查键是否存在:")
	fmt.Printf("包含apple: %v\n", cache.Contains("apple"))
	fmt.Printf("包含banana: %v\n", cache.Contains("banana"))

	// 删除键
	fmt.Println("\n7. 删除键:")
	if cache.Remove("date") {
		fmt.Println("date 已删除")
	}
	fmt.Printf("删除后的键: %v\n", cache.Keys())
	fmt.Printf("缓存大小: %d/%d\n", cache.Len(), cache.Cap())

	// 清空缓存
	fmt.Println("\n8. 清空缓存:")
	cache.Clear()
	fmt.Printf("清空后的缓存大小: %d/%d\n", cache.Len(), cache.Cap())
	fmt.Printf("缓存键: %v\n", cache.Keys())

	// 演示LRU算法的核心特性
	fmt.Println("\n=== LRU算法特性演示 ===")
	demoLRUBehavior()
}

func demoLRUBehavior() {
	cache := lru.New(3)

	fmt.Println("创建容量为3的缓存")

	// 添加3个元素
	cache.Put(1, "one")
	cache.Put(2, "two")
	cache.Put(3, "three")
	fmt.Printf("添加1,2,3后: %v\n", cache.Keys())

	// 访问元素1，使其成为最近使用的
	cache.Get(1)
	fmt.Printf("访问1后: %v\n", cache.Keys())

	// 添加元素4，应该淘汰元素2（最久未使用）
	cache.Put(4, "four")
	fmt.Printf("添加4后: %v\n", cache.Keys())

	// 验证元素2被淘汰
	if !cache.Contains(2) {
		fmt.Println("元素2已被淘汰（LRU策略）")
	}

	// 访问元素3
	cache.Get(3)
	fmt.Printf("访问3后: %v\n", cache.Keys())

	// 添加元素5，应该淘汰元素1
	cache.Put(5, "five")
	fmt.Printf("添加5后: %v\n", cache.Keys())

	if !cache.Contains(1) {
		fmt.Println("元素1已被淘汰（LRU策略）")
	}
}
