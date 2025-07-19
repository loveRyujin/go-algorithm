# LRU Cache Implementation

这是一个用Go语言实现的LRU（Least Recently Used）缓存算法。

## 算法原理

LRU缓存算法的核心思想是：当缓存满时，优先淘汰最久未使用的数据。这个算法使用了两个数据结构的组合：

1. **双向链表（Double Linked List）**：维护数据的访问顺序
2. **哈希表（Hash Map）**：提供O(1)的查找时间复杂度
3. **读写锁（RWMutex）**：保证并发安全，读操作使用读锁，写操作使用写锁

## 数据结构设计

```
哈希表: key -> 链表节点
双向链表: [最新] <-> [node] <-> [node] <-> [最旧]
读写锁: 保护并发访问
```

- 链表头部：最近使用的数据
- 链表尾部：最久未使用的数据
- 哈希表：快速定位链表中的节点
- 读写锁：确保多线程环境下的数据一致性

## 核心操作

### Get操作
1. 在哈希表中查找key
2. 如果找到，将对应节点移到链表头部
3. 返回值

### Put操作
1. 如果key已存在，更新值并移到链表头部
2. 如果key不存在：
   - 缓存未满：在链表头部插入新节点
   - 缓存已满：删除链表尾部节点，在头部插入新节点

## 时间复杂度
- Get: O(1) 
- Put: O(1)
- Remove: O(1)
- 并发安全：使用读写锁保护，读操作可以并发执行

## API文档

### 构造函数
```go
func New(capacity int) *Cache
```
创建一个指定容量的LRU缓存。

### 核心方法

#### Get
```go
func (c *Cache) Get(key any) (any, bool)
```
获取缓存中的值，如果key存在则返回对应的值和true，否则返回nil和false。
**注意：Get操作会更新key的访问顺序。**

#### Put
```go
func (c *Cache) Put(key, value any)
```
向缓存中添加或更新键值对。如果缓存已满，会自动淘汰最久未使用的数据。

#### Remove
```go
func (c *Cache) Remove(key any) bool
```
从缓存中删除指定的key，返回删除是否成功。

### 辅助方法

#### Peek
```go
func (c *Cache) Peek(key any) (any, bool)
```
查看缓存中的值，但**不会更新访问顺序**。

#### Contains
```go
func (c *Cache) Contains(key any) bool
```
检查缓存中是否包含指定的key。

#### Keys
```go
func (c *Cache) Keys() []any
```
返回缓存中所有的key，按访问顺序排列（最近使用的在前）。

#### Len
```go
func (c *Cache) Len() int
```
返回缓存中当前的元素数量。

#### Cap
```go
func (c *Cache) Cap() int
```
返回缓存的容量。

#### Clear
```go
func (c *Cache) Clear()
```
清空缓存中的所有数据。

## 使用示例

```go
package main

import (
    "fmt"
    "github.com/loveRyujin/go-algorithm/cache/lru"
)

func main() {
    // 创建容量为3的LRU缓存
    cache := lru.New(3)
    
    // 添加数据
    cache.Put("a", 1)
    cache.Put("b", 2)
    cache.Put("c", 3)
    
    // 获取数据
    if value, ok := cache.Get("a"); ok {
        fmt.Printf("a: %v\n", value)
    }
    
    // 添加新数据，触发淘汰
    cache.Put("d", 4) // 会淘汰"b"，因为"a"刚被访问过
    
    // 检查是否被淘汰
    if !cache.Contains("b") {
        fmt.Println("b has been evicted")
    }
}
```

## 运行测试

```bash
# 运行测试
go test -v

# 运行并发安全测试
go test -v -race

# 运行基准测试
go test -bench=.

# 运行示例
go run example/main.go
```

## 应用场景

1. **操作系统页面置换**：操作系统使用LRU算法来决定哪些页面应该被换出内存
2. **CPU缓存**：处理器缓存使用LRU算法来决定缓存行的替换
3. **数据库缓冲池**：数据库系统使用LRU来管理内存中的数据页
4. **Web缓存**：HTTP缓存服务器使用LRU来管理缓存的网页
5. **应用程序缓存**：各种应用程序使用LRU来缓存计算结果或数据

## 优缺点

### 优点
- **时间复杂度低**：所有操作都是O(1)
- **空间效率高**：只需要额外的指针空间
- **符合局部性原理**：最近使用的数据很可能再次被使用
- **并发安全**：使用读写锁保护，支持多线程环境

### 缺点
- **实现复杂**：需要维护双向链表和哈希表
- **内存开销**：每个节点需要额外的前驱和后继指针，以及锁的开销
- **锁竞争**：在高并发场景下可能存在锁竞争

## 注意事项

1. 这个实现**是线程安全的**，使用读写锁保护并发访问
2. `Get`操作会更新访问顺序，如果只是查看而不想影响顺序，请使用`Peek`方法
3. 缓存的容量必须大于0，否则可能导致未定义行为
4. 在高并发场景下，锁可能成为性能瓶颈，可考虑分片缓存或无锁实现
