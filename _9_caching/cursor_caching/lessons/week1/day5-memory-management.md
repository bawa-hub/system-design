# Day 5: Memory Management & Optimization

## 🎯 Learning Objectives
- Understand memory management in caching systems
- Learn memory optimization techniques
- Implement memory-efficient cache structures
- Build a cache with automatic memory management

## 📚 Theory: Memory Management in Caching

### Why Memory Management Matters
- **Performance**: Efficient memory usage improves cache performance
- **Scalability**: Better memory management allows larger caches
- **Cost**: Memory is expensive, especially in cloud environments
- **Stability**: Prevents out-of-memory errors

### Memory Allocation Patterns
1. **Stack Allocation**: Fast, limited size, automatic cleanup
2. **Heap Allocation**: Flexible, larger size, manual management
3. **Pool Allocation**: Reuse objects, reduce GC pressure
4. **Arena Allocation**: Allocate in blocks, bulk cleanup

## 🛠️ Hands-on: Memory-Efficient Cache Implementation

### Project Structure
```
projects/
├── memory-cache/
│   ├── main.go
│   ├── memory_cache.go
│   ├── memory_pool.go
│   ├── memory_monitor.go
│   └── README.md
```

### Step 1: Memory Pool Implementation

```go
// memory_pool.go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
    "unsafe"
)

// MemoryPool manages a pool of reusable memory blocks
type MemoryPool struct {
    pools    map[int]*sync.Pool
    mutex    sync.RWMutex
    stats    PoolStats
}

// PoolStats tracks memory pool statistics
type PoolStats struct {
    Allocations int64
    Reuses      int64
    Misses      int64
    TotalSize   int64
}

// NewMemoryPool creates a new memory pool
func NewMemoryPool() *MemoryPool {
    return &MemoryPool{
        pools: make(map[int]*sync.Pool),
    }
}

// Get retrieves a memory block of specified size
func (mp *MemoryPool) Get(size int) []byte {
    mp.mutex.RLock()
    pool, exists := mp.pools[size]
    mp.mutex.RUnlock()
    
    if !exists {
        mp.mutex.Lock()
        pool = &sync.Pool{
            New: func() interface{} {
                return make([]byte, size)
            },
        }
        mp.pools[size] = pool
        mp.mutex.Unlock()
    }
    
    block := pool.Get().([]byte)
    mp.stats.Allocations++
    
    // Check if we're reusing memory
    if len(block) == size {
        mp.stats.Reuses++
    } else {
        mp.stats.Misses++
        // Create new block if size doesn't match
        block = make([]byte, size)
    }
    
    mp.stats.TotalSize += int64(size)
    return block
}

// Put returns a memory block to the pool
func (mp *MemoryPool) Put(block []byte) {
    size := len(block)
    if size == 0 {
        return
    }
    
    mp.mutex.RLock()
    pool, exists := mp.pools[size]
    mp.mutex.RUnlock()
    
    if exists {
        // Reset the block
        for i := range block {
            block[i] = 0
        }
        pool.Put(block)
    }
}

// GetStats returns pool statistics
func (mp *MemoryPool) GetStats() PoolStats {
    mp.mutex.RLock()
    defer mp.mutex.RUnlock()
    return mp.stats
}

// MemoryBlock represents a memory block with metadata
type MemoryBlock struct {
    Data   []byte
    Size   int
    Pool   *MemoryPool
    InUse  bool
    AllocTime time.Time
}

// NewMemoryBlock creates a new memory block
func NewMemoryBlock(size int, pool *MemoryPool) *MemoryBlock {
    return &MemoryBlock{
        Data:   pool.Get(size),
        Size:   size,
        Pool:   pool,
        InUse:  true,
        AllocTime: time.Now(),
    }
}

// Release releases the memory block back to the pool
func (mb *MemoryBlock) Release() {
    if mb.InUse {
        mb.Pool.Put(mb.Data)
        mb.InUse = false
    }
}

// MemoryArena manages memory allocation in arenas
type MemoryArena struct {
    blocks    [][]byte
    current   []byte
    offset    int
    blockSize int
    mutex     sync.Mutex
}

// NewMemoryArena creates a new memory arena
func NewMemoryArena(blockSize int) *MemoryArena {
    return &MemoryArena{
        blockSize: blockSize,
        current:   make([]byte, blockSize),
        blocks:    [][]byte{},
    }
}

// Allocate allocates memory from the arena
func (ma *MemoryArena) Allocate(size int) []byte {
    ma.mutex.Lock()
    defer ma.mutex.Unlock()
    
    // Check if current block has enough space
    if ma.offset+size > len(ma.current) {
        // Save current block and allocate new one
        ma.blocks = append(ma.blocks, ma.current)
        ma.current = make([]byte, ma.blockSize)
        ma.offset = 0
    }
    
    // Allocate from current block
    start := ma.offset
    ma.offset += size
    
    return ma.current[start:ma.offset]
}

// Reset resets the arena for reuse
func (ma *MemoryArena) Reset() {
    ma.mutex.Lock()
    defer ma.mutex.Unlock()
    
    ma.blocks = ma.blocks[:0]
    ma.current = make([]byte, ma.blockSize)
    ma.offset = 0
}

// GetMemoryUsage returns current memory usage
func (ma *MemoryArena) GetMemoryUsage() int {
    ma.mutex.Lock()
    defer ma.mutex.Unlock()
    
    total := len(ma.current)
    for _, block := range ma.blocks {
        total += len(block)
    }
    return total
}
```

### Step 2: Memory-Efficient Cache

```go
// memory_cache.go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
    "unsafe"
)

// MemoryEfficientCache implements a memory-optimized cache
type MemoryEfficientCache struct {
    items      map[string]*CacheItem
    pool       *MemoryPool
    arena      *MemoryArena
    maxMemory  int64
    currentMemory int64
    mutex      sync.RWMutex
    stats      MemoryStats
}

// CacheItem represents a cache item with memory optimization
type CacheItem struct {
    Key        string
    Value      []byte
    Size       int
    CreatedAt  time.Time
    LastUsed   time.Time
    AccessCount int64
    Block      *MemoryBlock
}

// MemoryStats tracks memory usage statistics
type MemoryStats struct {
    TotalAllocated int64
    TotalUsed      int64
    PoolReuses     int64
    ArenaAllocs    int64
    GCCollections  int64
    MemoryPressure float64
}

// NewMemoryEfficientCache creates a new memory-efficient cache
func NewMemoryEfficientCache(maxMemory int64) *MemoryEfficientCache {
    return &MemoryEfficientCache{
        items:      make(map[string]*CacheItem),
        pool:       NewMemoryPool(),
        arena:      NewMemoryArena(1024 * 1024), // 1MB blocks
        maxMemory:  maxMemory,
        stats:      MemoryStats{},
    }
}

// Set stores a value in the cache with memory optimization
func (c *MemoryEfficientCache) Set(key string, value interface{}, ttl time.Duration) error {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    // Convert value to bytes
    valueBytes := c.valueToBytes(value)
    size := len(valueBytes)
    
    // Check memory limit
    if c.currentMemory+int64(size) > c.maxMemory {
        c.evictItems(size)
    }
    
    // Try to reuse existing item
    if item, exists := c.items[key]; exists {
        // Release old memory
        item.Block.Release()
        
        // Allocate new memory
        item.Block = NewMemoryBlock(size, c.pool)
        copy(item.Block.Data, valueBytes)
        item.Value = item.Block.Data[:size]
        item.Size = size
        item.LastUsed = time.Now()
        item.AccessCount++
        
        c.stats.TotalUsed += int64(size)
        return nil
    }
    
    // Create new item
    block := NewMemoryBlock(size, c.pool)
    copy(block.Data, valueBytes)
    
    item := &CacheItem{
        Key:         key,
        Value:       block.Data[:size],
        Size:        size,
        CreatedAt:   time.Now(),
        LastUsed:    time.Now(),
        AccessCount: 1,
        Block:       block,
    }
    
    c.items[key] = item
    c.currentMemory += int64(size)
    c.stats.TotalAllocated += int64(size)
    c.stats.TotalUsed += int64(size)
    
    return nil
}

// Get retrieves a value from the cache
func (c *MemoryEfficientCache) Get(key string) (interface{}, error) {
    c.mutex.RLock()
    item, exists := c.items[key]
    c.mutex.RUnlock()
    
    if !exists {
        c.stats.TotalUsed++ // Count as miss
        return nil, fmt.Errorf("key not found")
    }
    
    // Update access statistics
    c.mutex.Lock()
    item.LastUsed = time.Now()
    item.AccessCount++
    c.mutex.Unlock()
    
    // Convert bytes back to value
    return c.bytesToValue(item.Value), nil
}

// evictItems evicts items to make space
func (c *MemoryEfficientCache) evictItems(requiredSize int) {
    // Simple LRU eviction
    var oldestKey string
    var oldestTime time.Time
    
    for key, item := range c.items {
        if oldestKey == "" || item.LastUsed.Before(oldestTime) {
            oldestKey = key
            oldestTime = item.LastUsed
        }
    }
    
    if oldestKey != "" {
        c.deleteItem(oldestKey)
        
        // Recursively evict if still not enough space
        if c.currentMemory+int64(requiredSize) > c.maxMemory {
            c.evictItems(requiredSize)
        }
    }
}

// deleteItem removes an item from the cache
func (c *MemoryEfficientCache) deleteItem(key string) {
    if item, exists := c.items[key]; exists {
        item.Block.Release()
        c.currentMemory -= int64(item.Size)
        delete(c.items, key)
    }
}

// valueToBytes converts a value to bytes
func (c *MemoryEfficientCache) valueToBytes(value interface{}) []byte {
    // Simple conversion - in real implementation, use proper serialization
    switch v := value.(type) {
    case string:
        return []byte(v)
    case []byte:
        return v
    case int:
        return []byte(fmt.Sprintf("%d", v))
    case float64:
        return []byte(fmt.Sprintf("%f", v))
    default:
        return []byte(fmt.Sprintf("%v", v))
    }
}

// bytesToValue converts bytes back to a value
func (c *MemoryEfficientCache) bytesToValue(data []byte) interface{} {
    // Simple conversion - in real implementation, use proper deserialization
    return string(data)
}

// GetMemoryStats returns memory usage statistics
func (c *MemoryEfficientCache) GetMemoryStats() MemoryStats {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    // Update GC stats
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    c.stats.GCCollections = int64(m.NumGC)
    
    // Calculate memory pressure
    if c.maxMemory > 0 {
        c.stats.MemoryPressure = float64(c.currentMemory) / float64(c.maxMemory) * 100
    }
    
    return c.stats
}

// OptimizeMemory performs memory optimization
func (c *MemoryEfficientCache) OptimizeMemory() {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    // Force garbage collection
    runtime.GC()
    
    // Reset arena if it's getting too large
    if c.arena.GetMemoryUsage() > 10*1024*1024 { // 10MB
        c.arena.Reset()
    }
    
    // Update stats
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    c.stats.GCCollections = int64(m.NumGC)
}

// SetMemoryLimit sets the maximum memory limit
func (c *MemoryEfficientCache) SetMemoryLimit(limit int64) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    c.maxMemory = limit
    
    // Evict items if over limit
    if c.currentMemory > limit {
        c.evictItems(0)
    }
}

// GetMemoryUsage returns current memory usage
func (c *MemoryEfficientCache) GetMemoryUsage() int64 {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    return c.currentMemory
}

// GetItemCount returns the number of items in cache
func (c *MemoryEfficientCache) GetItemCount() int {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    return len(c.items)
}
```

### Step 3: Memory Monitor

```go
// memory_monitor.go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

// MemoryMonitor monitors memory usage and performs optimization
type MemoryMonitor struct {
    cache      *MemoryEfficientCache
    interval   time.Duration
    stop       chan bool
    mutex      sync.RWMutex
    stats      MonitorStats
}

// MonitorStats tracks monitoring statistics
type MonitorStats struct {
    ChecksPerformed int64
    Optimizations   int64
    MemoryFreed     int64
    LastCheck       time.Time
}

// NewMemoryMonitor creates a new memory monitor
func NewMemoryMonitor(cache *MemoryEfficientCache, interval time.Duration) *MemoryMonitor {
    return &MemoryMonitor{
        cache:    cache,
        interval: interval,
        stop:     make(chan bool),
        stats:    MonitorStats{},
    }
}

// Start starts the memory monitor
func (mm *MemoryMonitor) Start() {
    go mm.monitor()
}

// Stop stops the memory monitor
func (mm *MemoryMonitor) Stop() {
    close(mm.stop)
}

// monitor runs the memory monitoring loop
func (mm *MemoryMonitor) monitor() {
    ticker := time.NewTicker(mm.interval)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            mm.checkMemory()
        case <-mm.stop:
            return
        }
    }
}

// checkMemory checks memory usage and performs optimization
func (mm *MemoryMonitor) checkMemory() {
    mm.mutex.Lock()
    defer mm.mutex.Unlock()
    
    mm.stats.ChecksPerformed++
    mm.stats.LastCheck = time.Now()
    
    // Get current memory stats
    stats := mm.cache.GetMemoryStats()
    
    // Check if memory pressure is high
    if stats.MemoryPressure > 80.0 {
        fmt.Printf("High memory pressure: %.2f%%\n", stats.MemoryPressure)
        
        // Perform optimization
        mm.cache.OptimizeMemory()
        mm.stats.Optimizations++
        
        // Get updated stats
        newStats := mm.cache.GetMemoryStats()
        mm.stats.MemoryFreed += stats.TotalUsed - newStats.TotalUsed
        
        fmt.Printf("Memory optimization completed. Freed: %d bytes\n", 
            stats.TotalUsed - newStats.TotalUsed)
    }
    
    // Log memory usage
    fmt.Printf("Memory usage: %d/%d bytes (%.2f%%)\n", 
        mm.cache.GetMemoryUsage(), 
        mm.cache.GetMemoryLimit(),
        stats.MemoryPressure)
}

// GetStats returns monitoring statistics
func (mm *MemoryMonitor) GetStats() MonitorStats {
    mm.mutex.RLock()
    defer mm.mutex.RUnlock()
    return mm.stats
}

// SetInterval sets the monitoring interval
func (mm *MemoryMonitor) SetInterval(interval time.Duration) {
    mm.mutex.Lock()
    defer mm.mutex.Unlock()
    mm.interval = interval
}
```

### Step 4: Main Application

```go
// main.go
package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    // Create memory-efficient cache with 10MB limit
    cache := NewMemoryEfficientCache(10 * 1024 * 1024)
    
    // Start memory monitor
    monitor := NewMemoryMonitor(cache, 5*time.Second)
    monitor.Start()
    defer monitor.Stop()
    
    fmt.Println("=== Memory-Efficient Cache Demo ===")
    
    // Test memory allocation
    fmt.Println("Adding items to cache...")
    for i := 0; i < 1000; i++ {
        key := fmt.Sprintf("key_%d", i)
        value := fmt.Sprintf("value_%d_%s", i, generateRandomString(100))
        
        if err := cache.Set(key, value, 1*time.Hour); err != nil {
            fmt.Printf("Error setting %s: %v\n", key, err)
        }
        
        if i%100 == 0 {
            stats := cache.GetMemoryStats()
            fmt.Printf("Added %d items, memory usage: %d bytes (%.2f%%)\n", 
                i+1, cache.GetMemoryUsage(), stats.MemoryPressure)
        }
    }
    
    // Test memory optimization
    fmt.Println("\nPerforming memory optimization...")
    cache.OptimizeMemory()
    
    // Print final statistics
    stats := cache.GetMemoryStats()
    monitorStats := monitor.GetStats()
    
    fmt.Printf("\n=== Final Statistics ===\n")
    fmt.Printf("Items in cache: %d\n", cache.GetItemCount())
    fmt.Printf("Memory usage: %d bytes\n", cache.GetMemoryUsage())
    fmt.Printf("Memory pressure: %.2f%%\n", stats.MemoryPressure)
    fmt.Printf("Pool reuses: %d\n", stats.PoolReuses)
    fmt.Printf("GC collections: %d\n", stats.GCCollections)
    fmt.Printf("Monitor checks: %d\n", monitorStats.ChecksPerformed)
    fmt.Printf("Optimizations: %d\n", monitorStats.Optimizations)
    
    // Test memory limit
    fmt.Println("\nTesting memory limit...")
    cache.SetMemoryLimit(5 * 1024 * 1024) // 5MB
    
    // Add more items to trigger eviction
    for i := 1000; i < 2000; i++ {
        key := fmt.Sprintf("key_%d", i)
        value := fmt.Sprintf("value_%d_%s", i, generateRandomString(100))
        
        if err := cache.Set(key, value, 1*time.Hour); err != nil {
            fmt.Printf("Error setting %s: %v\n", key, err)
        }
        
        if i%100 == 0 {
            stats := cache.GetMemoryStats()
            fmt.Printf("Added %d items, memory usage: %d bytes (%.2f%%)\n", 
                i+1, cache.GetMemoryUsage(), stats.MemoryPressure)
        }
    }
    
    // Wait for monitoring
    time.Sleep(10 * time.Second)
}

// generateRandomString generates a random string of specified length
func generateRandomString(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    result := make([]byte, length)
    for i := range result {
        result[i] = charset[i%len(charset)]
    }
    return string(result)
}
```

## 🧪 Exercise: Memory Profiling

### Task
Create a memory profiling test to analyze memory usage:

```go
// memory_profile_test.go
package main

import (
    "fmt"
    "runtime"
    "testing"
    "time"
)

func TestMemoryUsage(t *testing.T) {
    cache := NewMemoryEfficientCache(100 * 1024 * 1024) // 100MB
    
    // Test memory allocation
    for i := 0; i < 10000; i++ {
        key := fmt.Sprintf("key_%d", i)
        value := fmt.Sprintf("value_%d", i)
        
        if err := cache.Set(key, value, 1*time.Hour); err != nil {
            t.Fatalf("Error setting %s: %v", key, err)
        }
    }
    
    // Check memory usage
    stats := cache.GetMemoryStats()
    if stats.MemoryPressure > 100.0 {
        t.Errorf("Memory pressure too high: %.2f%%", stats.MemoryPressure)
    }
    
    // Test memory optimization
    cache.OptimizeMemory()
    
    // Check if memory was freed
    newStats := cache.GetMemoryStats()
    if newStats.TotalUsed >= stats.TotalUsed {
        t.Error("Memory optimization did not free memory")
    }
}

func BenchmarkMemoryAllocation(b *testing.B) {
    cache := NewMemoryEfficientCache(100 * 1024 * 1024)
    
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        key := fmt.Sprintf("key_%d", i)
        value := fmt.Sprintf("value_%d", i)
        
        cache.Set(key, value, 1*time.Hour)
    }
}

func BenchmarkMemoryPool(b *testing.B) {
    pool := NewMemoryPool()
    
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        size := 100 + (i % 1000)
        block := pool.Get(size)
        pool.Put(block)
    }
}
```

## 📝 Key Takeaways

1. **Memory Pools**: Reuse memory blocks to reduce allocation overhead
2. **Arena Allocation**: Allocate in blocks for better performance
3. **Memory Monitoring**: Track usage and optimize automatically
4. **Eviction Strategies**: Remove items when memory limit is reached
5. **Garbage Collection**: Understand GC impact on performance

## 🎯 Practice Questions

1. How do memory pools improve cache performance?
2. What are the trade-offs of arena allocation?
3. How do you choose the right memory limit for a cache?
4. What happens when memory pressure is high?
5. How do you measure memory efficiency in production?

## 🚀 Next Steps

Tomorrow we'll cover:
- Cache warming and preloading
- Advanced caching patterns
- Performance optimization techniques

---

**Your Notes:**
<!-- Add your notes, questions, and insights here -->

**Code Experiments:**
<!-- Document any code experiments or modifications you try -->

**Questions:**
<!-- Write down any questions you have -->
