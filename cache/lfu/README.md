# LFU Cache Implementation

这是一个用Go语言实现的LFU（Least Frequently Used）缓存算法。

## 算法原理

LFU缓存算法的核心思想是：**优先淘汰访问频率最低的数据**。当多个数据具有相同的最低访问频率时，使用LRU策略（淘汰最久未使用的数据）。

### 数据结构设计

LFU算法使用以下数据结构的组合：

```
cache: map[key] -> *list.Element     // 快速定位元素
frequencies: map[frequency] -> *list.List  // 频率分组的双向链表
minFreq: int                         // 当前最小频率
```

每个缓存条目包含：
- `key`: 键
- `value`: 值  
- `freq`: 访问频率

### 核心操作

#### Get操作 - O(1)
1. 在哈希表中查找key
2. 如果找到，将元素从当前频率链表移除
3. 频率+1，加入新频率链表的头部
4. 更新minFreq（如果需要）

#### Put操作 - O(1) 
1. 如果key已存在，更新值并增加频率
2. 如果key不存在且缓存未满，直接插入（频率=1）
3. 如果缓存已满，淘汰最小频率链表的尾部元素，然后插入新元素

#### 淘汰策略
- 选择频率最小的元素
- 如果有多个相同最小频率的元素，选择最久未使用的（链表尾部）

## 时间复杂度
- Get: O(1)
- Put: O(1)
- Remove: O(1)
- 所有操作都是常数时间复杂度

## API文档

### 构造函数
```go
func New(capacity int) *Cache
```
创建一个指定容量的LFU缓存。

### 核心方法

#### Get
```go
func (c *Cache) Get(key any) (any, bool)
```
获取缓存中的值。**注意：Get操作会增加key的访问频率。**

#### Put
```go
func (c *Cache) Put(key, value any)
```
向缓存中添加或更新键值对。如果缓存已满，会自动淘汰频率最低的数据。

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
查看缓存中的值，但**不会增加访问频率**。

#### Contains
```go
func (c *Cache) Contains(key any) bool
```
检查缓存中是否包含指定的key。

#### GetFrequency
```go
func (c *Cache) GetFrequency(key any) int
```
获取指定key的访问频率，如果key不存在返回0。

#### Keys
```go
func (c *Cache) Keys() map[int][]any
```
返回缓存中所有的key，按频率分组。

#### MinFrequency
```go
func (c *Cache) MinFrequency() int
```
返回缓存中的最小访问频率。

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
    "github.com/loveRyujin/go-algorithm/cache/lfu"
)

func main() {
    // 创建容量为3的LFU缓存
    cache := lfu.New(3)
    
    // 添加数据
    cache.Put("a", 1)
    cache.Put("b", 2)
    cache.Put("c", 3)
    
    // 访问数据创建不同频率
    cache.Get("a") // a: freq 2
    cache.Get("a") // a: freq 3
    cache.Get("b") // b: freq 2
    // c: freq 1
    
    // 添加新数据，会淘汰频率最低的c
    cache.Put("d", 4)
    
    fmt.Printf("Contains c: %v\n", cache.Contains("c")) // false
    fmt.Printf("Frequency of a: %d\n", cache.GetFrequency("a")) // 3
}
```

## 运行测试

```bash
# 运行测试
go test -v

# 运行基准测试
go test -bench=.

# 运行示例
go run example/main.go
```

## 并发安全

这个实现是**线程安全的**，使用`sync.RWMutex`来保护并发访问：
- 读操作（Get、Peek、Contains等）使用读锁
- 写操作（Put、Remove、Clear等）使用写锁

## LFU vs LRU 对比

| 特性 | LFU | LRU |
|------|-----|-----|
| 淘汰策略 | 基于访问频率 | 基于访问时间 |
| 适用场景 | 访问模式相对稳定 | 访问模式时间局部性强 |
| 实现复杂度 | 较复杂（需要维护频率） | 较简单 |
| 内存开销 | 稍高（存储频率信息） | 较低 |
| 性能 | O(1)所有操作 | O(1)所有操作 |

## 应用场景

1. **数据库缓冲池**：频繁访问的数据页应该保留在内存中
2. **Web缓存**：热门页面和资源的缓存
3. **CDN系统**：内容分发网络中的热点内容缓存
4. **编译器优化**：代码块的访问频率统计和优化
5. **操作系统**：某些页面置换算法的实现

## 注意事项

1. **访问频率重置**：长时间运行后，一些旧数据的高频率可能不再反映当前访问模式
2. **冷启动问题**：新加入的数据初始频率较低，可能被过早淘汰
3. **频率溢出**：理论上频率可能会溢出，但在实际应用中很少发生
4. **零容量处理**：当capacity为0时，缓存不会存储任何数据
