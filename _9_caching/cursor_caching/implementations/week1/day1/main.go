package main

import (
    "fmt"
    "sync"
    "time"
)

// CacheItem represents a single item in the cache
type CacheItem struct {
    Value     interface{}
    ExpiresAt time.Time
}

// IsExpired checks if the cache item has expired
func (item *CacheItem) IsExpired() bool {
    return time.Now().After(item.ExpiresAt)
}

// SimpleCache represents our basic cache
type SimpleCache struct {
    items  map[string]*CacheItem
    mutex  sync.RWMutex
    hits   int64
    misses int64
}

// NewSimpleCache creates a new cache instance
func NewSimpleCache() *SimpleCache {
    return &SimpleCache{
        items: make(map[string]*CacheItem),
    }
}

// Set stores a value in the cache with TTL
func (c *SimpleCache) Set(key string, value interface{}, ttl time.Duration) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    c.items[key] = &CacheItem{
        Value:     value,
        ExpiresAt: time.Now().Add(ttl),
    }
}

// Get retrieves a value from the cache
func (c *SimpleCache) Get(key string) (interface{}, bool) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    item, exists := c.items[key]
    if !exists {
        return nil, false
    }
    
    if item.IsExpired() {
        // Clean up expired item
        c.mutex.RUnlock()
        c.mutex.Lock()
        delete(c.items, key)
        c.mutex.Unlock()
        c.mutex.RLock()
        return nil, false
    }
    
    return item.Value, true
}

// Delete removes an item from the cache
func (c *SimpleCache) Delete(key string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    delete(c.items, key)
}

// Clear removes all items from the cache
func (c *SimpleCache) Clear() {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.items = make(map[string]*CacheItem)
}

// Size returns the number of items in the cache
func (c *SimpleCache) Size() int {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    return len(c.items)
}

// CacheStats holds cache performance statistics
type CacheStats struct {
    Hits   int64
    Misses int64
    Size   int
}

// GetStats returns cache statistics
func (c *SimpleCache) GetStats() CacheStats {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    return CacheStats{
        Hits:   c.hits,
        Misses: c.misses,
        Size:   len(c.items),
    }
}

// recordHit records a cache hit
func (c *SimpleCache) recordHit() {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.hits++
}

// recordMiss records a cache miss
func (c *SimpleCache) recordMiss() {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.misses++
}

func main() {
    fmt.Println("=== Day 1: Caching Fundamentals ===")
    fmt.Println("Building our first cache implementation in Go\n")
    
    // Create a new cache
    cache := NewSimpleCache()
    
    // Example 1: Basic operations
    fmt.Println("1. Basic Cache Operations:")
    fmt.Println("   Setting values with different TTLs...")
    
    cache.Set("user:1", "John Doe", 5*time.Second)
    cache.Set("user:2", "Jane Smith", 10*time.Second)
    cache.Set("config:theme", "dark", 1*time.Minute)
    cache.Set("temp:data", "temporary", 2*time.Second)
    
    fmt.Printf("   Cache size: %d\n", cache.Size())
    
    // Example 2: Cache hits and misses
    fmt.Println("\n2. Cache Hits and Misses:")
    
    // Test cache hit
    if value, found := cache.Get("user:1"); found {
        fmt.Printf("   ✓ Cache HIT: user:1 = %v\n", value)
        cache.recordHit()
    }
    
    // Test cache miss
    if value, found := cache.Get("user:999"); found {
        fmt.Printf("   ✓ Cache HIT: user:999 = %v\n", value)
        cache.recordHit()
    } else {
        fmt.Println("   ✗ Cache MISS: user:999 not found")
        cache.recordMiss()
    }
    
    // Example 3: TTL expiration
    fmt.Println("\n3. TTL Expiration Test:")
    fmt.Println("   Waiting for temp:data to expire (2 seconds)...")
    
    // Check before expiration
    if value, found := cache.Get("temp:data"); found {
        fmt.Printf("   Before expiration: temp:data = %v\n", value)
    }
    
    // Wait for expiration
    time.Sleep(3 * time.Second)
    
    // Check after expiration
    if value, found := cache.Get("temp:data"); found {
        fmt.Printf("   After expiration: temp:data = %v\n", value)
    } else {
        fmt.Println("   After expiration: temp:data expired and was removed")
    }
    
    // Example 4: Cache statistics
    fmt.Println("\n4. Cache Statistics:")
    stats := cache.GetStats()
    total := stats.Hits + stats.Misses
    hitRate := float64(stats.Hits) / float64(total) * 100
    
    fmt.Printf("   Hits: %d\n", stats.Hits)
    fmt.Printf("   Misses: %d\n", stats.Misses)
    fmt.Printf("   Hit Rate: %.2f%%\n", hitRate)
    fmt.Printf("   Cache Size: %d\n", stats.Size)
    
    // Example 5: Concurrent access
    fmt.Println("\n5. Concurrent Access Test:")
    fmt.Println("   Testing thread safety with goroutines...")
    
    var wg sync.WaitGroup
    numGoroutines := 10
    
    // Start multiple goroutines
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            // Write operations
            key := fmt.Sprintf("concurrent:key:%d", id)
            value := fmt.Sprintf("value:%d", id)
            cache.Set(key, value, 1*time.Minute)
            
            // Read operations
            if val, found := cache.Get(key); found {
                fmt.Printf("   Goroutine %d: Found %s = %v\n", id, key, val)
            }
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("   Final cache size: %d\n", cache.Size())
    
    // Example 6: Cache operations
    fmt.Println("\n6. Cache Management:")
    fmt.Println("   Testing delete and clear operations...")
    
    // Delete specific item
    cache.Delete("user:1")
    fmt.Printf("   After deleting user:1, cache size: %d\n", cache.Size())
    
    // Clear all items
    cache.Clear()
    fmt.Printf("   After clearing all items, cache size: %d\n", cache.Size())
    
    // Example 7: Performance demonstration
    fmt.Println("\n7. Performance Demonstration:")
    fmt.Println("   Measuring cache performance...")
    
    // Reset cache
    cache = NewSimpleCache()
    
    // Measure set performance
    start := time.Now()
    for i := 0; i < 10000; i++ {
        key := fmt.Sprintf("perf:key:%d", i)
        value := fmt.Sprintf("perf:value:%d", i)
        cache.Set(key, value, 1*time.Hour)
    }
    setTime := time.Since(start)
    
    // Measure get performance
    start = time.Now()
    for i := 0; i < 10000; i++ {
        key := fmt.Sprintf("perf:key:%d", i)
        cache.Get(key)
    }
    getTime := time.Since(start)
    
    fmt.Printf("   Set 10,000 items in: %v\n", setTime)
    fmt.Printf("   Get 10,000 items in: %v\n", getTime)
    fmt.Printf("   Average set time: %v\n", setTime/10000)
    fmt.Printf("   Average get time: %v\n", getTime/10000)
    
    fmt.Println("\n=== Day 1 Complete! ===")
    fmt.Println("You've successfully implemented a basic cache with:")
    fmt.Println("✓ Thread-safe operations")
    fmt.Println("✓ TTL support")
    fmt.Println("✓ Basic statistics")
    fmt.Println("✓ Performance testing")
    fmt.Println("\nReady for Day 2: Cache Eviction Policies! 🚀")
}
