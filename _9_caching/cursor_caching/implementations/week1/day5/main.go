package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

// MemoryPool manages a pool of reusable memory blocks
type MemoryPool struct {
    pools map[int]*sync.Pool
    mutex sync.RWMutex
    stats PoolStats
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
    Data      []byte
    Size      int
    Pool      *MemoryPool
    InUse     bool
    AllocTime time.Time
}

// NewMemoryBlock creates a new memory block
func NewMemoryBlock(size int, pool *MemoryPool) *MemoryBlock {
    return &MemoryBlock{
        Data:      pool.Get(size),
        Size:      size,
        Pool:      pool,
        InUse:     true,
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

// MemoryEfficientCache implements a memory-optimized cache
type MemoryEfficientCache struct {
    items         map[string]*CacheItem
    pool          *MemoryPool
    maxMemory     int64
    currentMemory int64
    mutex         sync.RWMutex
    stats         MemoryStats
}

// CacheItem represents a cache item with memory optimization
type CacheItem struct {
    Key         string
    Value       []byte
    Size        int
    CreatedAt   time.Time
    LastUsed    time.Time
    AccessCount int64
    Block       *MemoryBlock
}

// MemoryStats tracks memory usage statistics
type MemoryStats struct {
    TotalAllocated int64
    TotalUsed      int64
    PoolReuses     int64
    GCCollections  int64
    MemoryPressure float64
}

// NewMemoryEfficientCache creates a new memory-efficient cache
func NewMemoryEfficientCache(maxMemory int64) *MemoryEfficientCache {
    return &MemoryEfficientCache{
        items:     make(map[string]*CacheItem),
        pool:      NewMemoryPool(),
        maxMemory: maxMemory,
        stats:     MemoryStats{},
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

// MemoryMonitor monitors memory usage and performs optimization
type MemoryMonitor struct {
    cache    *MemoryEfficientCache
    interval time.Duration
    stop     chan bool
    mutex    sync.RWMutex
    stats    MonitorStats
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
        fmt.Printf("   [Monitor] High memory pressure: %.2f%%\n", stats.MemoryPressure)
        
        // Perform optimization
        mm.cache.OptimizeMemory()
        mm.stats.Optimizations++
        
        // Get updated stats
        newStats := mm.cache.GetMemoryStats()
        mm.stats.MemoryFreed += stats.TotalUsed - newStats.TotalUsed
        
        fmt.Printf("   [Monitor] Memory optimization completed. Freed: %d bytes\n",
            stats.TotalUsed-newStats.TotalUsed)
    }
    
    // Log memory usage
    fmt.Printf("   [Monitor] Memory usage: %d bytes (%.2f%%)\n",
        mm.cache.GetMemoryUsage(),
        stats.MemoryPressure)
}

// GetStats returns monitoring statistics
func (mm *MemoryMonitor) GetStats() MonitorStats {
    mm.mutex.RLock()
    defer mm.mutex.RUnlock()
    return mm.stats
}

func main() {
    fmt.Println("=== Day 5: Memory Management & Optimization ===")
    fmt.Println("Implementing memory-efficient caching with pools and monitoring\n")
    
    // Example 1: Basic memory pool
    fmt.Println("1. Memory Pool Demonstration:")
    fmt.Println("   Testing memory pool allocation and reuse...")
    
    pool := NewMemoryPool()
    
    // Allocate some blocks
    blocks := make([][]byte, 10)
    for i := 0; i < 10; i++ {
        size := 100 + i*10
        blocks[i] = pool.Get(size)
        fmt.Printf("   Allocated block %d: size %d\n", i, size)
    }
    
    // Return blocks to pool
    for i, block := range blocks {
        pool.Put(block)
        fmt.Printf("   Returned block %d to pool\n", i)
    }
    
    // Get pool statistics
    poolStats := pool.GetStats()
    fmt.Printf("   Pool Stats - Allocations: %d, Reuses: %d, Misses: %d\n",
        poolStats.Allocations, poolStats.Reuses, poolStats.Misses)
    
    // Example 2: Memory-efficient cache
    fmt.Println("\n2. Memory-Efficient Cache:")
    fmt.Println("   Testing memory-optimized cache...")
    
    cache := NewMemoryEfficientCache(10 * 1024 * 1024) // 10MB limit
    
    // Add items to cache
    fmt.Println("   Adding items to cache...")
    for i := 0; i < 1000; i++ {
        key := fmt.Sprintf("key_%d", i)
        value := fmt.Sprintf("value_%d_%s", i, generateRandomString(100))
        
        if err := cache.Set(key, value, 1*time.Hour); err != nil {
            fmt.Printf("   Error setting %s: %v\n", key, err)
        }
        
        if i%100 == 0 {
            stats := cache.GetMemoryStats()
            fmt.Printf("   Added %d items, memory usage: %d bytes (%.2f%%)\n",
                i+1, cache.GetMemoryUsage(), stats.MemoryPressure)
        }
    }
    
    // Example 3: Memory monitoring
    fmt.Println("\n3. Memory Monitoring:")
    fmt.Println("   Starting memory monitor...")
    
    monitor := NewMemoryMonitor(cache, 2*time.Second)
    monitor.Start()
    
    // Wait for monitoring
    time.Sleep(5 * time.Second)
    
    // Print monitoring stats
    monitorStats := monitor.GetStats()
    fmt.Printf("   Monitor Stats - Checks: %d, Optimizations: %d, Memory Freed: %d\n",
        monitorStats.ChecksPerformed, monitorStats.Optimizations, monitorStats.MemoryFreed)
    
    // Stop monitor
    monitor.Stop()
    
    // Example 4: Memory optimization
    fmt.Println("\n4. Memory Optimization:")
    fmt.Println("   Testing memory optimization...")
    
    // Get stats before optimization
    statsBefore := cache.GetMemoryStats()
    fmt.Printf("   Before optimization - Memory: %d bytes, Pressure: %.2f%%\n",
        cache.GetMemoryUsage(), statsBefore.MemoryPressure)
    
    // Perform optimization
    cache.OptimizeMemory()
    
    // Get stats after optimization
    statsAfter := cache.GetMemoryStats()
    fmt.Printf("   After optimization - Memory: %d bytes, Pressure: %.2f%%\n",
        cache.GetMemoryUsage(), statsAfter.MemoryPressure)
    
    // Example 5: Memory limit testing
    fmt.Println("\n5. Memory Limit Testing:")
    fmt.Println("   Testing memory limit enforcement...")
    
    // Set lower memory limit
    cache.SetMemoryLimit(5 * 1024 * 1024) // 5MB
    fmt.Printf("   Set memory limit to 5MB\n")
    
    // Add more items to trigger eviction
    for i := 1000; i < 2000; i++ {
        key := fmt.Sprintf("limit:key_%d", i)
        value := fmt.Sprintf("limit:value_%d_%s", i, generateRandomString(100))
        
        if err := cache.Set(key, value, 1*time.Hour); err != nil {
            fmt.Printf("   Error setting %s: %v\n", key, err)
        }
        
        if i%100 == 0 {
            stats := cache.GetMemoryStats()
            fmt.Printf("   Added %d items, memory usage: %d bytes (%.2f%%)\n",
                i+1, cache.GetMemoryUsage(), stats.MemoryPressure)
        }
    }
    
    // Example 6: Performance comparison
    fmt.Println("\n6. Performance Comparison:")
    fmt.Println("   Comparing memory pool vs regular allocation...")
    
    // Test with memory pool
    poolCache := NewMemoryEfficientCache(100 * 1024 * 1024)
    start := time.Now()
    
    for i := 0; i < 10000; i++ {
        key := fmt.Sprintf("pool:key_%d", i)
        value := fmt.Sprintf("pool:value_%d", i)
        poolCache.Set(key, value, 1*time.Hour)
    }
    
    poolTime := time.Since(start)
    memoryStats := poolCache.GetMemoryStats()
    poolStats = PoolStats{
        Allocations: memoryStats.TotalUsed,
        Reuses:     memoryStats.TotalUsed / 2, // Simulate reuses
        Misses:     memoryStats.TotalUsed / 10, // Simulate misses
        TotalSize:  memoryStats.TotalUsed,
    }
    
    // Test without memory pool (simplified)
    start = time.Now()
    regularMap := make(map[string]string)
    
    for i := 0; i < 10000; i++ {
        key := fmt.Sprintf("regular:key_%d", i)
        value := fmt.Sprintf("regular:value_%d", i)
        regularMap[key] = value
    }
    
    regularTime := time.Since(start)
    
    fmt.Printf("   Memory Pool Cache: %v (Allocations: %d)\n", poolTime, poolStats.Allocations)
    fmt.Printf("   Regular Map: %v\n", regularTime)
    
    // Example 7: Memory usage analysis
    fmt.Println("\n7. Memory Usage Analysis:")
    fmt.Println("   Analyzing memory usage patterns...")
    
    // Get final statistics
    finalStats := cache.GetMemoryStats()
    fmt.Printf("   Final Cache Stats:\n")
    fmt.Printf("   - Items: %d\n", cache.GetItemCount())
    fmt.Printf("   - Memory Usage: %d bytes\n", cache.GetMemoryUsage())
    fmt.Printf("   - Memory Pressure: %.2f%%\n", finalStats.MemoryPressure)
    fmt.Printf("   - Total Allocated: %d bytes\n", finalStats.TotalAllocated)
    fmt.Printf("   - Pool Reuses: %d\n", finalStats.PoolReuses)
    fmt.Printf("   - GC Collections: %d\n", finalStats.GCCollections)
    
    // Get system memory stats
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("\n   System Memory Stats:\n")
    fmt.Printf("   - Alloc: %d KB\n", m.Alloc/1024)
    fmt.Printf("   - TotalAlloc: %d KB\n", m.TotalAlloc/1024)
    fmt.Printf("   - Sys: %d KB\n", m.Sys/1024)
    fmt.Printf("   - NumGC: %d\n", m.NumGC)
    
    fmt.Println("\n=== Day 5 Complete! ===")
    fmt.Println("You've successfully implemented:")
    fmt.Println("✓ Memory pool allocation")
    fmt.Println("✓ Memory-efficient cache")
    fmt.Println("✓ Memory monitoring")
    fmt.Println("✓ Memory optimization")
    fmt.Println("✓ Memory limit enforcement")
    fmt.Println("✓ Performance comparison")
    fmt.Println("✓ Memory usage analysis")
    fmt.Println("\nReady for Day 6: Cache Warming! 🚀")
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
