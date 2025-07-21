# TinyLFU Cache Implementation

这是一个用Go语言实现的TinyLFU（Tiny Least Frequently Used）缓存算法。TinyLFU是一个高效的缓存淘汰算法，它结合了多种数据结构来提供接近LFU的性能，但内存开销更小。

## 算法原理

TinyLFU算法由以下几个核心组件组成：

### 1. 三段式架构
- **Admission Window (1%)**: 新数据的缓冲区，使用LRU策略
- **Probation Segment (~19%)**: 试用区，新晋升的数据在此接受考验
- **Protected Segment (~80%)**: 保护区，频繁访问的热数据

### 2. 频率估计组件
- **Count-Min Sketch**: 用于估计访问频率，内存高效
- **Bloom Filter (Doorkeeper)**: 追踪最近被淘汰的数据，避免缓存污染

### 3. 工作流程
```
新数据 → Admission Window → (淘汰时) → 准入测试 → Probation Segment
                                  ↓                    ↓
                              拒绝/添加到Doorkeeper    访问频繁时
                                                       ↓
                                                Protected Segment
```

## 核心特性

1. **低内存开销**: 使用概率数据结构，避免存储完整的频率信息
2. **高准入精度**: 通过Doorkeeper避免重复添加最近淘汰的数据
3. **适应性强**: 通过定期重置适应访问模式变化
4. **线程安全**: 支持并发访问

## 时间复杂度
- Get: O(1)
- Put: O(1)
- Remove: O(1)

## API文档

### 构造函数
```go
func New(capacity int) *Cache
```
创建一个指定容量的TinyLFU缓存。

### 核心方法

#### Get
```go
func (c *Cache) Get(key any) (any, bool)
```
获取缓存中的值。每次访问会更新频率估计和段间晋升。

#### Put
```go
func (c *Cache) Put(key, value any)
```
向缓存中添加或更新键值对。新数据首先进入Admission Window。

#### Remove
```go
func (c *Cache) Remove(key any) bool
```
从缓存中删除指定的key，返回删除是否成功。

### 辅助方法

#### Contains
```go
func (c *Cache) Contains(key any) bool
```
检查缓存中是否包含指定的key。

#### GetFrequency
```go
func (c *Cache) GetFrequency(key any) int
```
获取指定key的估计访问频率。

#### Stats
```go
func (c *Cache) Stats() map[string]int
```
返回缓存的详细统计信息，包括各段的大小和容量。

#### Len
```go
func (c *Cache) Len() int
```
返回缓存中当前的元素总数。

#### Cap
```go
func (c *Cache) Cap() int
```
返回缓存的总容量。

#### Clear
```go
func (c *Cache) Clear()
```
清空缓存中的所有数据和统计信息。

## 使用示例

```go
package main

import (
    "fmt"
    "github.com/loveRyujin/go-algorithm/cache/tinylfu"
)

func main() {
    // 创建容量为100的TinyLFU缓存
    cache := tinylfu.New(100)
    
    // 添加数据
    cache.Put("user:123", map[string]interface{}{
        "name": "Alice",
        "age":  25,
    })
    
    // 访问数据
    if data, ok := cache.Get("user:123"); ok {
        fmt.Printf("User data: %v\n", data)
    }
    
    // 创建访问模式
    for i := 0; i < 10; i++ {
        cache.Get("user:123") // 增加访问频率
    }
    
    fmt.Printf("Access frequency: %d\n", cache.GetFrequency("user:123"))
    fmt.Printf("Cache stats: %v\n", cache.Stats())
}
```

## 段间晋升机制

### Window → Probation
- 当Window满时，被淘汰的数据需要通过准入测试
- 比较候选数据与Probation中LRU数据的频率
- 频率更高的数据被准入，频率低的被拒绝

### Probation → Protected  
- 当Probation中的数据被访问时，晋升到Protected
- 如果Protected满，被淘汰的数据降级到Probation

## 准入策略

TinyLFU使用以下策略决定是否准入新数据：

1. **空间检查**: 如果有空间，直接准入
2. **Doorkeeper检查**: 如果最近被淘汰过，拒绝准入
3. **频率比较**: 比较新数据与待淘汰数据的访问频率
4. **频率重置**: 定期重置避免频率饱和

## 性能特征

### 优势
- **内存高效**: 使用O(1)额外内存而非O(n)
- **高命中率**: 在大多数工作负载下接近最优LFU性能
- **快速响应**: 能快速适应访问模式变化
- **避免污染**: Doorkeeper机制防止缓存污染

### 适用场景
- **Web缓存**: HTTP缓存、CDN
- **数据库缓存**: 查询结果缓存
- **应用缓存**: 对象缓存、计算结果缓存
- **大数据处理**: 热点数据缓存

## 运行测试

```bash
# 运行测试
go test -v

# 运行基准测试
go test -bench=.

# 运行示例
go run example/main.go
```

## 算法比较

| 特性 | TinyLFU | LFU | LRU |
|------|---------|-----|-----|
| 内存开销 | O(1) | O(n) | O(1) |
| 准入精度 | 高 | 最高 | 无 |
| 适应性 | 强 | 弱 | 强 |
| 实现复杂度 | 高 | 中 | 低 |
| 污染抵抗 | 强 | 中 | 弱 |

## 配置参数

当前实现使用以下默认配置：

- **Window Size**: 1% of capacity (最小为1)
- **Protected Size**: 80% of remaining capacity  
- **Probation Size**: 剩余容量
- **Sketch Size**: 10x capacity (更好的准确性)
- **Bloom Filter Size**: 4x capacity
- **Reset Threshold**: 10x capacity (重置频率)

## 并发安全

这个实现是**线程安全的**，使用`sync.RWMutex`来保护并发访问：
- 读操作（Get、Contains等）使用读锁
- 写操作（Put、Remove、Clear等）使用写锁

## 注意事项

1. **概率性质**: 频率估计是近似的，可能存在少量误差
2. **内存使用**: Sketch和Bloom Filter需要额外内存
3. **重置开销**: 定期重置会有短暂的性能影响
4. **小容量限制**: 容量太小时段间分布可能不够理想

## 实现细节

### Count-Min Sketch
- 4行哈希表，每行宽度等于容量
- 使用最小值估计频率
- 定期将所有计数减半（老化）

### Bloom Filter
- 动态计算最优大小和哈希函数数量
- 目标假阳率：1%
- 用于Doorkeeper功能

### 段间管理
- 每个段使用独立的LRU缓存
- 自动处理段间数据移动
- 维护全局一致性
