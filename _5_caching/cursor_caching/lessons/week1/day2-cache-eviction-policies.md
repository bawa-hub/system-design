# Day 2: Cache Eviction Policies

## 🎯 Learning Objectives
- Understand different cache eviction policies
- Implement LRU (Least Recently Used) cache
- Learn about LFU, FIFO, and other policies
- Build a configurable cache with multiple eviction strategies

## 📚 Theory: Cache Eviction Policies

### What is Cache Eviction?
When a cache reaches its capacity limit, we need to decide which items to remove to make space for new ones. This decision is made by **eviction policies**.

### Why Eviction is Important
- **Memory Management**: Prevents unlimited memory growth
- **Performance**: Keeps most valuable data in cache
- **Efficiency**: Maximizes cache hit rates

## 🔄 Common Eviction Policies

### 1. LRU (Least Recently Used)
- **Rule**: Remove the item that hasn't been accessed for the longest time
- **Use Case**: Web browsers, database buffers
- **Implementation**: Doubly linked list + hash map

### 2. LFU (Least Frequently Used)
- **Rule**: Remove the item with the lowest access count
- **Use Case**: Content delivery networks
- **Implementation**: Min-heap + hash map

### 3. FIFO (First In, First Out)
- **Rule**: Remove the oldest item by insertion time
- **Use Case**: Simple scenarios, streaming data
- **Implementation**: Queue + hash map

### 4. LIFO (Last In, First Out)
- **Rule**: Remove the most recently inserted item
- **Use Case**: Stack-like behavior
- **Implementation**: Stack + hash map

### 5. TTL (Time To Live)
- **Rule**: Remove items after a fixed time period
- **Use Case**: Session data, temporary data
- **Implementation**: Timestamp + cleanup routine

## 🛠️ Hands-on: Implementing LRU Cache

### Project Structure
```
projects/
├── lru-cache/
│   ├── main.go
│   ├── lru.go
│   ├── cache_interface.go
│   └── README.md
```

### Step 1: Cache Interface

```go
// cache_interface.go
package main

import "time"

// EvictionPolicy defines different eviction strategies
type EvictionPolicy int

const (
    LRU EvictionPolicy = iota
    LFU
    FIFO
    LIFO
    TTL
)

// CacheItem represents a single item in the cache
type CacheItem struct {
    Key       string
    Value     interface{}
    CreatedAt time.Time
    LastUsed  time.Time
    UseCount  int
    ExpiresAt time.Time
}

// Cache defines the interface for all cache implementations
type Cache interface {
    Get(key string) (interface{}, bool)
    Set(key string, value interface{}, ttl time.Duration)
    Delete(key string)
    Clear()
    Size() int
    Capacity() int
    SetCapacity(capacity int)
    GetStats() CacheStats
}

// CacheStats holds cache performance statistics
type CacheStats struct {
    Hits       int64
    Misses     int64
    Evictions  int64
    Size       int
    Capacity   int
    HitRate    float64
    MissRate   float64
}

// CalculateRates updates hit and miss rates
func (s *CacheStats) CalculateRates() {
    total := s.Hits + s.Misses
    if total > 0 {
        s.HitRate = float64(s.Hits) / float64(total) * 100
        s.MissRate = float64(s.Misses) / float64(total) * 100
    }
}
```

### Step 2: LRU Cache Implementation

```go
// lru.go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Node represents a node in the doubly linked list
type Node struct {
    Key   string
    Value interface{}
    Prev  *Node
    Next  *Node
}

// LRUCache implements Least Recently Used cache
type LRUCache struct {
    capacity int
    items    map[string]*Node
    head     *Node
    tail     *Node
    mutex    sync.RWMutex
    stats    CacheStats
}

// NewLRUCache creates a new LRU cache with given capacity
func NewLRUCache(capacity int) *LRUCache {
    cache := &LRUCache{
        capacity: capacity,
        items:    make(map[string]*Node),
        stats:    CacheStats{Capacity: capacity},
    }
    
    // Initialize dummy head and tail nodes
    cache.head = &Node{}
    cache.tail = &Node{}
    cache.head.Next = cache.tail
    cache.tail.Prev = cache.head
    
    return cache
}

// addToHead adds a node to the head of the list
func (c *LRUCache) addToHead(node *Node) {
    node.Prev = c.head
    node.Next = c.head.Next
    c.head.Next.Prev = node
    c.head.Next = node
}

// removeNode removes a node from the list
func (c *LRUCache) removeNode(node *Node) {
    node.Prev.Next = node.Next
    node.Next.Prev = node.Prev
}

// moveToHead moves a node to the head of the list
func (c *LRUCache) moveToHead(node *Node) {
    c.removeNode(node)
    c.addToHead(node)
}

// removeTail removes the tail node from the list
func (c *LRUCache) removeTail() *Node {
    lastNode := c.tail.Prev
    c.removeNode(lastNode)
    return lastNode
}

// Get retrieves a value from the cache
func (c *LRUCache) Get(key string) (interface{}, bool) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    if node, exists := c.items[key]; exists {
        // Move to head (most recently used)
        c.moveToHead(node)
        c.stats.Hits++
        c.stats.CalculateRates()
        return node.Value, true
    }
    
    c.stats.Misses++
    c.stats.CalculateRates()
    return nil, false
}

// Set stores a value in the cache
func (c *LRUCache) Set(key string, value interface{}, ttl time.Duration) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    if node, exists := c.items[key]; exists {
        // Update existing node
        node.Value = value
        c.moveToHead(node)
        return
    }
    
    // Create new node
    newNode := &Node{
        Key:   key,
        Value: value,
    }
    
    // Add to head
    c.addToHead(newNode)
    c.items[key] = newNode
    
    // Check capacity
    if len(c.items) > c.capacity {
        // Remove least recently used (tail)
        tail := c.removeTail()
        delete(c.items, tail.Key)
        c.stats.Evictions++
    }
    
    c.stats.Size = len(c.items)
}

// Delete removes an item from the cache
func (c *LRUCache) Delete(key string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    if node, exists := c.items[key]; exists {
        c.removeNode(node)
        delete(c.items, key)
        c.stats.Size = len(c.items)
    }
}

// Clear removes all items from the cache
func (c *LRUCache) Clear() {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    c.items = make(map[string]*Node)
    c.head.Next = c.tail
    c.tail.Prev = c.head
    c.stats = CacheStats{Capacity: c.capacity}
}

// Size returns the number of items in the cache
func (c *LRUCache) Size() int {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    return len(c.items)
}

// Capacity returns the maximum capacity of the cache
func (c *LRUCache) Capacity() int {
    return c.capacity
}

// SetCapacity changes the cache capacity
func (c *LRUCache) SetCapacity(capacity int) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    c.capacity = capacity
    
    // Evict items if necessary
    for len(c.items) > capacity {
        tail := c.removeTail()
        delete(c.items, tail.Key)
        c.stats.Evictions++
    }
    
    c.stats.Capacity = capacity
    c.stats.Size = len(c.items)
}

// GetStats returns cache performance statistics
func (c *LRUCache) GetStats() CacheStats {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    stats := c.stats
    stats.Size = len(c.items)
    stats.CalculateRates()
    return stats
}

// PrintCache prints the current cache state (for debugging)
func (c *LRUCache) PrintCache() {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    fmt.Print("Cache (LRU order): ")
    current := c.head.Next
    for current != c.tail {
        fmt.Printf("[%s:%v] ", current.Key, current.Value)
        current = current.Next
    }
    fmt.Println()
}
```

### Step 3: Main Application with Testing

```go
// main.go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Create LRU cache with capacity 3
    cache := NewLRUCache(3)
    
    fmt.Println("=== LRU Cache Demo ===")
    
    // Test basic operations
    cache.Set("a", 1, 0)
    cache.Set("b", 2, 0)
    cache.Set("c", 3, 0)
    
    fmt.Println("After adding a, b, c:")
    cache.PrintCache()
    
    // Access 'a' to make it most recently used
    if value, found := cache.Get("a"); found {
        fmt.Printf("Found 'a': %v\n", value)
    }
    
    fmt.Println("After accessing 'a':")
    cache.PrintCache()
    
    // Add 'd' - should evict 'b' (least recently used)
    cache.Set("d", 4, 0)
    fmt.Println("After adding 'd' (should evict 'b'):")
    cache.PrintCache()
    
    // Try to get 'b' - should be a miss
    if value, found := cache.Get("b"); found {
        fmt.Printf("Found 'b': %v\n", value)
    } else {
        fmt.Println("'b' was evicted (cache miss)")
    }
    
    // Print statistics
    stats := cache.GetStats()
    fmt.Printf("\nCache Statistics:\n")
    fmt.Printf("Size: %d/%d\n", stats.Size, stats.Capacity)
    fmt.Printf("Hits: %d\n", stats.Hits)
    fmt.Printf("Misses: %d\n", stats.Misses)
    fmt.Printf("Hit Rate: %.2f%%\n", stats.HitRate)
    fmt.Printf("Evictions: %d\n", stats.Evictions)
    
    // Test capacity change
    fmt.Println("\n=== Testing Capacity Change ===")
    cache.SetCapacity(2)
    fmt.Println("After reducing capacity to 2:")
    cache.PrintCache()
    
    // Test clear
    fmt.Println("\n=== Testing Clear ===")
    cache.Clear()
    fmt.Println("After clear:")
    cache.PrintCache()
}
```

## 🧪 Exercise: Implement LFU Cache

### Task
Implement a Least Frequently Used (LFU) cache that:
1. Tracks access frequency for each item
2. Evicts the least frequently used item when full
3. In case of tie, evicts the least recently used among tied items

### Solution Template
```go
// lfu.go
package main

import (
    "container/heap"
    "sync"
    "time"
)

// LFUNode represents a node in the LFU cache
type LFUNode struct {
    Key       string
    Value     interface{}
    Frequency int
    LastUsed  time.Time
}

// LFUHeap implements a min-heap for LFU eviction
type LFUHeap []*LFUNode

func (h LFUHeap) Len() int { return len(h) }

func (h LFUHeap) Less(i, j int) bool {
    if h[i].Frequency == h[j].Frequency {
        return h[i].LastUsed.Before(h[j].LastUsed)
    }
    return h[i].Frequency < h[j].Frequency
}

func (h LFUHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *LFUHeap) Push(x interface{}) {
    *h = append(*h, x.(*LFUNode))
}

func (h *LFUHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// LFUCache implements Least Frequently Used cache
type LFUCache struct {
    capacity int
    items    map[string]*LFUNode
    heap     LFUHeap
    mutex    sync.RWMutex
    stats    CacheStats
}

// NewLFUCache creates a new LFU cache
func NewLFUCache(capacity int) *LFUCache {
    cache := &LFUCache{
        capacity: capacity,
        items:    make(map[string]*LFUNode),
        heap:     make(LFUHeap, 0),
        stats:    CacheStats{Capacity: capacity},
    }
    heap.Init(&cache.heap)
    return cache
}

// Get retrieves a value from the cache
func (c *LFUCache) Get(key string) (interface{}, bool) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    if node, exists := c.items[key]; exists {
        node.Frequency++
        node.LastUsed = time.Now()
        heap.Fix(&c.heap, 0) // Re-heapify
        c.stats.Hits++
        c.stats.CalculateRates()
        return node.Value, true
    }
    
    c.stats.Misses++
    c.stats.CalculateRates()
    return nil, false
}

// Set stores a value in the cache
func (c *LFUCache) Set(key string, value interface{}, ttl time.Duration) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    if node, exists := c.items[key]; exists {
        // Update existing node
        node.Value = value
        node.Frequency++
        node.LastUsed = time.Now()
        heap.Fix(&c.heap, 0)
        return
    }
    
    // Create new node
    newNode := &LFUNode{
        Key:       key,
        Value:     value,
        Frequency: 1,
        LastUsed:  time.Now(),
    }
    
    // Add to heap and map
    heap.Push(&c.heap, newNode)
    c.items[key] = newNode
    
    // Check capacity
    if len(c.items) > c.capacity {
        // Remove least frequently used
        lfu := heap.Pop(&c.heap).(*LFUNode)
        delete(c.items, lfu.Key)
        c.stats.Evictions++
    }
    
    c.stats.Size = len(c.items)
}

// ... (implement other methods similar to LRU)
```

## 📝 Key Takeaways

1. **LRU is most common**: Good for temporal locality
2. **LFU is good for frequency**: Better for access patterns
3. **Implementation matters**: Use appropriate data structures
4. **Thread safety**: Always protect shared data
5. **Statistics**: Track performance metrics

## 🎯 Practice Questions

1. When would you choose LRU over LFU?
2. How does the doubly linked list help in LRU implementation?
3. What's the time complexity of LRU operations?
4. How would you implement a cache with both LRU and TTL?
5. What are the trade-offs of different eviction policies?

## 🚀 Next Steps

Tomorrow we'll cover:
- Cache strategies (write-through, write-behind, etc.)
- Cache consistency models
- Advanced caching patterns

---

**Your Notes:**
<!-- Add your notes, questions, and insights here -->

**Code Experiments:**
<!-- Document any code experiments or modifications you try -->

**Questions:**
<!-- Write down any questions you have -->
