package lru

import "container/list"

// Cache LRU缓存结构
type Cache struct {
	capacity int
	cache    map[any]*list.Element
	list     *list.List
}

// entry 缓存条目
type entry struct {
	key   any
	value any
}

// New 创建一个新的LRU缓存
func New(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		cache:    make(map[any]*list.Element),
		list:     list.New(),
	}
}

// Get 获取缓存中的值
func (c *Cache) Get(key any) (any, bool) {
	if element, ok := c.cache[key]; ok {
		// 将访问的元素移动到链表头部
		c.list.MoveToFront(element)
		return element.Value.(*entry).value, true
	}
	return nil, false
}

// Put 向缓存中添加键值对
func (c *Cache) Put(key, value any) {
	if element, ok := c.cache[key]; ok {
		// 如果键已存在，更新值并移动到链表头部
		element.Value.(*entry).value = value
		c.list.MoveToFront(element)
		return
	}

	// 如果缓存已满，删除最久未使用的元素
	if c.list.Len() >= c.capacity {
		c.removeOldest()
	}

	// 添加新元素到链表头部
	newEntry := &entry{key: key, value: value}
	element := c.list.PushFront(newEntry)
	c.cache[key] = element
}

// Remove 从缓存中删除指定键
func (c *Cache) Remove(key any) bool {
	if element, ok := c.cache[key]; ok {
		c.removeElement(element)
		return true
	}
	return false
}

// removeOldest 删除最久未使用的元素（链表尾部）
func (c *Cache) removeOldest() {
	if c.list.Len() == 0 {
		return
	}
	oldest := c.list.Back()
	if oldest != nil {
		c.removeElement(oldest)
	}
}

// removeElement 删除指定元素
func (c *Cache) removeElement(element *list.Element) {
	c.list.Remove(element)
	delete(c.cache, element.Value.(*entry).key)
}

// Len 返回缓存中的元素数量
func (c *Cache) Len() int {
	return c.list.Len()
}

// Cap 返回缓存的容量
func (c *Cache) Cap() int {
	return c.capacity
}

// Clear 清空缓存
func (c *Cache) Clear() {
	c.cache = make(map[any]*list.Element)
	c.list = list.New()
}

// Keys 返回缓存中所有的键（按访问顺序，最新的在前）
func (c *Cache) Keys() []any {
	keys := make([]any, 0, c.list.Len())
	for element := c.list.Front(); element != nil; element = element.Next() {
		keys = append(keys, element.Value.(*entry).key)
	}
	return keys
}

// Contains 检查缓存中是否包含指定键
func (c *Cache) Contains(key any) bool {
	_, ok := c.cache[key]
	return ok
}

// Peek 查看缓存中的值但不更新访问顺序
func (c *Cache) Peek(key any) (any, bool) {
	if element, ok := c.cache[key]; ok {
		return element.Value.(*entry).value, true
	}
	return nil, false
}
