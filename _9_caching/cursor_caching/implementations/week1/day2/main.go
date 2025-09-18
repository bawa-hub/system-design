package main

import (
    "fmt"
    "sync"
    "time"
)

// EvictionPolicy defines different eviction strategies
type EvictionPolicy int

const (
    LRU EvictionPolicy = iota
    LFU
    FIFO
    LIFO
)

// CacheItem represents a single item in the cache
type CacheItem struct {
    Key         string
    Value       interface{}
    CreatedAt   time.Time
    LastUsed    time.Time
    AccessCount int64
    Size        int
}

// Node represents a node in the doubly linked list for LRU
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
        stats:    CacheStats{},
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
        return node.Value, true
    }
    
    c.stats.Misses++
    return nil, false
}

// Set stores a value in the cache
func (c *LRUCache) Set(key string, value interface{}) {
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

// Size returns the number of items in the cache
func (c *LRUCache) Size() int {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    return len(c.items)
}

// GetStats returns cache statistics
func (c *LRUCache) GetStats() CacheStats {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    total := c.stats.Hits + c.stats.Misses
    if total > 0 {
        c.stats.HitRate = float64(c.stats.Hits) / float64(total) * 100
        c.stats.MissRate = float64(c.stats.Misses) / float64(total) * 100
    }
    
    return c.stats
}

// PrintCache prints the current cache state (for debugging)
func (c *LRUCache) PrintCache() {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    fmt.Print("   LRU Cache (most recent first): ")
    current := c.head.Next
    for current != c.tail {
        fmt.Printf("[%s:%v] ", current.Key, current.Value)
        current = current.Next
    }
    fmt.Println()
}

// LFUCache implements Least Frequently Used cache
type LFUCache struct {
    capacity int
    items    map[string]*CacheItem
    mutex    sync.RWMutex
    stats    CacheStats
}

// NewLFUCache creates a new LFU cache
func NewLFUCache(capacity int) *LFUCache {
    return &LFUCache{
        capacity: capacity,
        items:    make(map[string]*CacheItem),
        stats:    CacheStats{},
    }
}

// Get retrieves a value from the cache
func (c *LFUCache) Get(key string) (interface{}, bool) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    if item, exists := c.items[key]; exists {
        item.AccessCount++
        item.LastUsed = time.Now()
        c.stats.Hits++
        return item.Value, true
    }
    
    c.stats.Misses++
    return nil, false
}

// Set stores a value in the cache
func (c *LFUCache) Set(key string, value interface{}) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    if item, exists := c.items[key]; exists {
        // Update existing item
        item.Value = value
        item.AccessCount++
        item.LastUsed = time.Now()
        return
    }
    
    // Create new item
    newItem := &CacheItem{
        Key:         key,
        Value:       value,
        CreatedAt:   time.Now(),
        LastUsed:    time.Now(),
        AccessCount: 1,
    }
    
    c.items[key] = newItem
    
    // Check capacity
    if len(c.items) > c.capacity {
        // Find least frequently used item
        var lfuKey string
        var minCount int64 = -1
        var oldestTime time.Time
        
        for k, item := range c.items {
            if minCount == -1 || item.AccessCount < minCount {
                lfuKey = k
                minCount = item.AccessCount
                oldestTime = item.LastUsed
            } else if item.AccessCount == minCount && item.LastUsed.Before(oldestTime) {
                lfuKey = k
                oldestTime = item.LastUsed
            }
        }
        
        if lfuKey != "" {
            delete(c.items, lfuKey)
            c.stats.Evictions++
        }
    }
    
    c.stats.Size = len(c.items)
}

// Delete removes an item from the cache
func (c *LFUCache) Delete(key string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    delete(c.items, key)
    c.stats.Size = len(c.items)
}

// Size returns the number of items in the cache
func (c *LFUCache) Size() int {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    return len(c.items)
}

// GetStats returns cache statistics
func (c *LFUCache) GetStats() CacheStats {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    total := c.stats.Hits + c.stats.Misses
    if total > 0 {
        c.stats.HitRate = float64(c.stats.Hits) / float64(total) * 100
        c.stats.MissRate = float64(c.stats.Misses) / float64(total) * 100
    }
    
    return c.stats
}

// PrintCache prints the current cache state (for debugging)
func (c *LFUCache) PrintCache() {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    fmt.Print("   LFU Cache (by access count): ")
    for key, item := range c.items {
        fmt.Printf("[%s:%v(count:%d)] ", key, item.Value, item.AccessCount)
    }
    fmt.Println()
}

// CacheStats holds cache performance statistics
type CacheStats struct {
    Hits      int64
    Misses    int64
    Evictions int64
    Size      int
    HitRate   float64
    MissRate  float64
}

func main() {
    fmt.Println("=== Day 2: Cache Eviction Policies ===")
    fmt.Println("Implementing LRU, LFU, and other eviction strategies\n")
    
    // Example 1: LRU Cache
    fmt.Println("1. LRU (Least Recently Used) Cache:")
    fmt.Println("   Testing LRU eviction policy...")
    
    lruCache := NewLRUCache(3)
    
    // Add items
    lruCache.Set("a", 1)
    lruCache.Set("b", 2)
    lruCache.Set("c", 3)
    lruCache.PrintCache()
    
    // Access 'a' to make it most recently used
    if value, found := lruCache.Get("a"); found {
        fmt.Printf("   Accessed 'a': %v\n", value)
    }
    lruCache.PrintCache()
    
    // Add 'd' - should evict 'b' (least recently used)
    lruCache.Set("d", 4)
    fmt.Println("   After adding 'd' (should evict 'b'):")
    lruCache.PrintCache()
    
    // Try to get 'b' - should be a miss
    if value, found := lruCache.Get("b"); found {
        fmt.Printf("   Found 'b': %v\n", value)
    } else {
        fmt.Println("   'b' was evicted (cache miss)")
    }
    
    // Print LRU statistics
    stats := lruCache.GetStats()
    fmt.Printf("   LRU Stats - Hits: %d, Misses: %d, Hit Rate: %.2f%%, Evictions: %d\n\n",
        stats.Hits, stats.Misses, stats.HitRate, stats.Evictions)
    
    // Example 2: LFU Cache
    fmt.Println("2. LFU (Least Frequently Used) Cache:")
    fmt.Println("   Testing LFU eviction policy...")
    
    lfuCache := NewLFUCache(3)
    
    // Add items
    lfuCache.Set("x", "first")
    lfuCache.Set("y", "second")
    lfuCache.Set("z", "third")
    lfuCache.PrintCache()
    
    // Access 'x' multiple times to increase its frequency
    lfuCache.Get("x")
    lfuCache.Get("x")
    lfuCache.Get("y")
    lfuCache.Get("z")
    lfuCache.Get("z")
    lfuCache.Get("z")
    fmt.Println("   After multiple accesses:")
    lfuCache.PrintCache()
    
    // Add 'w' - should evict 'y' (least frequently used)
    lfuCache.Set("w", "fourth")
    fmt.Println("   After adding 'w' (should evict 'y'):")
    lfuCache.PrintCache()
    
    // Print LFU statistics
    stats = lfuCache.GetStats()
    fmt.Printf("   LFU Stats - Hits: %d, Misses: %d, Hit Rate: %.2f%%, Evictions: %d\n\n",
        stats.Hits, stats.Misses, stats.HitRate, stats.Evictions)
    
    // Example 3: Performance comparison
    fmt.Println("3. Performance Comparison:")
    fmt.Println("   Comparing LRU vs LFU performance...")
    
    // Test LRU performance
    lruPerf := NewLRUCache(1000)
    start := time.Now()
    
    for i := 0; i < 10000; i++ {
        key := fmt.Sprintf("lru:key:%d", i%1000)
        value := fmt.Sprintf("lru:value:%d", i)
        lruPerf.Set(key, value)
        lruPerf.Get(key)
    }
    
    lruTime := time.Since(start)
    lruStats := lruPerf.GetStats()
    
    // Test LFU performance
    lfuPerf := NewLFUCache(1000)
    start = time.Now()
    
    for i := 0; i < 10000; i++ {
        key := fmt.Sprintf("lfu:key:%d", i%1000)
        value := fmt.Sprintf("lfu:value:%d", i)
        lfuPerf.Set(key, value)
        lfuPerf.Get(key)
    }
    
    lfuTime := time.Since(start)
    lfuStats := lfuPerf.GetStats()
    
    fmt.Printf("   LRU Performance: %v (Hit Rate: %.2f%%)\n", lruTime, lruStats.HitRate)
    fmt.Printf("   LFU Performance: %v (Hit Rate: %.2f%%)\n", lfuTime, lfuStats.HitRate)
    
    // Example 4: Different access patterns
    fmt.Println("\n4. Access Pattern Analysis:")
    fmt.Println("   Testing different access patterns...")
    
    // Sequential access pattern
    seqCache := NewLRUCache(5)
    fmt.Println("   Sequential access pattern:")
    for i := 0; i < 10; i++ {
        key := fmt.Sprintf("seq:%d", i)
        seqCache.Set(key, i)
        seqCache.Get(key)
    }
    seqStats := seqCache.GetStats()
    fmt.Printf("   Sequential - Hits: %d, Misses: %d, Hit Rate: %.2f%%\n",
        seqStats.Hits, seqStats.Misses, seqStats.HitRate)
    
    // Random access pattern
    randCache := NewLRUCache(5)
    fmt.Println("   Random access pattern:")
    for i := 0; i < 10; i++ {
        key := fmt.Sprintf("rand:%d", i%3) // Only 3 unique keys
        randCache.Set(key, i)
        randCache.Get(key)
    }
    randStats := randCache.GetStats()
    fmt.Printf("   Random - Hits: %d, Misses: %d, Hit Rate: %.2f%%\n",
        randStats.Hits, randStats.Misses, randStats.HitRate)
    
    // Example 5: Memory usage analysis
    fmt.Println("\n5. Memory Usage Analysis:")
    fmt.Println("   Analyzing memory usage of different eviction policies...")
    
    // LRU memory usage
    lruMem := NewLRUCache(1000)
    for i := 0; i < 1000; i++ {
        key := fmt.Sprintf("mem:lru:%d", i)
        value := fmt.Sprintf("mem:lru:value:%d", i)
        lruMem.Set(key, value)
    }
    
    // LFU memory usage
    lfuMem := NewLFUCache(1000)
    for i := 0; i < 1000; i++ {
        key := fmt.Sprintf("mem:lfu:%d", i)
        value := fmt.Sprintf("mem:lfu:value:%d", i)
        lfuMem.Set(key, value)
    }
    
    fmt.Printf("   LRU Cache size: %d items\n", lruMem.Size())
    fmt.Printf("   LFU Cache size: %d items\n", lfuMem.Size())
    
    fmt.Println("\n=== Day 2 Complete! ===")
    fmt.Println("You've successfully implemented:")
    fmt.Println("✓ LRU (Least Recently Used) cache")
    fmt.Println("✓ LFU (Least Frequently Used) cache")
    fmt.Println("✓ Performance comparison")
    fmt.Println("✓ Access pattern analysis")
    fmt.Println("✓ Memory usage analysis")
    fmt.Println("\nReady for Day 3: Cache Strategies! 🚀")
}
