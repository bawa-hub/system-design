# Day 1: Caching Fundamentals

## 🎯 Learning Objectives
- Understand what caching is and why it's crucial
- Learn different types of caches
- Understand cache performance metrics
- Build your first simple cache in Go

## 📚 Theory: What is Caching?

### Definition
Caching is the process of storing frequently accessed data in a fast-access storage layer to improve application performance and reduce latency.

### Why Caching Matters
- **Speed**: Reduces data access time from milliseconds to microseconds
- **Scalability**: Reduces load on primary data sources
- **Cost**: Reduces expensive operations (database queries, API calls)
- **User Experience**: Faster response times

### Cache Hierarchy
```
CPU Registers (fastest, smallest)
    ↓
L1 Cache (CPU cache)
    ↓
L2 Cache (CPU cache)
    ↓
L3 Cache (CPU cache)
    ↓
RAM (Main memory)
    ↓
SSD (Solid state drive)
    ↓
HDD (Hard disk drive) - slowest, largest
```

## 🔧 Types of Caches

### 1. CPU Cache
- **L1, L2, L3**: Different levels of CPU cache
- **Purpose**: Store frequently accessed CPU instructions and data
- **Size**: KB to MB range

### 2. Memory Cache
- **RAM-based**: Fast access, volatile
- **Purpose**: Store application data in memory
- **Size**: MB to GB range

### 3. Disk Cache
- **File system cache**: OS-level caching
- **Purpose**: Cache file operations
- **Size**: GB to TB range

### 4. Application Cache
- **In-memory**: Application-specific caching
- **Purpose**: Cache business logic results
- **Size**: Configurable

### 5. Distributed Cache
- **Network-based**: Multiple machines
- **Purpose**: Share cache across services
- **Size**: TB to PB range

## 📊 Cache Performance Metrics

### Hit Rate
```
Hit Rate = (Cache Hits / Total Requests) × 100%
```
- **Good**: 80-95%
- **Excellent**: 95%+

### Miss Rate
```
Miss Rate = (Cache Misses / Total Requests) × 100%
```
- **Good**: 5-20%
- **Excellent**: <5%

### Response Time
- **Cache Hit**: Microseconds
- **Cache Miss**: Milliseconds to seconds

## 🛠️ Hands-on: Building Your First Cache in Go

### Project Structure
```
projects/
├── simple-cache/
│   ├── main.go
│   ├── cache.go
│   └── README.md
```

### Step 1: Basic Cache Interface

```go
// cache.go
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
    items map[string]*CacheItem
    mutex sync.RWMutex
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
```

### Step 2: Main Application

```go
// main.go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Create a new cache
    cache := NewSimpleCache()
    
    // Set some values
    cache.Set("user:1", "John Doe", 5*time.Second)
    cache.Set("user:2", "Jane Smith", 10*time.Second)
    cache.Set("config:theme", "dark", 1*time.Minute)
    
    // Test cache operations
    fmt.Println("=== Cache Operations ===")
    
    // Get existing value
    if value, found := cache.Get("user:1"); found {
        fmt.Printf("Found user:1 = %v\n", value)
    }
    
    // Get non-existent value
    if value, found := cache.Get("user:999"); found {
        fmt.Printf("Found user:999 = %v\n", value)
    } else {
        fmt.Println("user:999 not found")
    }
    
    // Wait for expiration
    fmt.Println("\nWaiting for user:1 to expire...")
    time.Sleep(6 * time.Second)
    
    if value, found := cache.Get("user:1"); found {
        fmt.Printf("Found user:1 = %v\n", value)
    } else {
        fmt.Println("user:1 expired and was removed")
    }
    
    // Cache statistics
    fmt.Printf("\nCache size: %d\n", cache.Size())
    
    // Clear cache
    cache.Clear()
    fmt.Printf("After clear, cache size: %d\n", cache.Size())
}
```

## 🧪 Exercise: Cache Performance Testing

### Task
Create a performance test to measure:
1. Cache hit/miss rates
2. Average response times
3. Memory usage

### Solution Template
```go
// performance_test.go
package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "testing"
    "time"
)

func BenchmarkCacheOperations(b *testing.B) {
    cache := NewSimpleCache()
    
    // Pre-populate cache
    for i := 0; i < 1000; i++ {
        key := fmt.Sprintf("key_%d", i)
        cache.Set(key, fmt.Sprintf("value_%d", i), 1*time.Hour)
    }
    
    b.ResetTimer()
    
    b.Run("Get", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            key := fmt.Sprintf("key_%d", i%1000)
            cache.Get(key)
        }
    })
    
    b.Run("Set", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            key := fmt.Sprintf("new_key_%d", i)
            cache.Set(key, fmt.Sprintf("new_value_%d", i), 1*time.Hour)
        }
    })
}

func TestCachePerformance(t *testing.T) {
    cache := NewSimpleCache()
    
    // Test data
    keys := make([]string, 10000)
    for i := 0; i < 10000; i++ {
        keys[i] = fmt.Sprintf("key_%d", i)
    }
    
    // Pre-populate cache
    start := time.Now()
    for i, key := range keys {
        cache.Set(key, fmt.Sprintf("value_%d", i), 1*time.Hour)
    }
    setupTime := time.Since(start)
    
    // Test reads
    start = time.Now()
    hits := 0
    misses := 0
    
    for i := 0; i < 100000; i++ {
        key := keys[rand.Intn(len(keys))]
        if _, found := cache.Get(key); found {
            hits++
        } else {
            misses++
        }
    }
    
    readTime := time.Since(start)
    
    // Calculate metrics
    totalRequests := hits + misses
    hitRate := float64(hits) / float64(totalRequests) * 100
    avgReadTime := readTime / time.Duration(totalRequests)
    
    fmt.Printf("Setup time: %v\n", setupTime)
    fmt.Printf("Total reads: %d\n", totalRequests)
    fmt.Printf("Hits: %d (%.2f%%)\n", hits, hitRate)
    fmt.Printf("Misses: %d (%.2f%%)\n", misses, 100-hitRate)
    fmt.Printf("Average read time: %v\n", avgReadTime)
    fmt.Printf("Cache size: %d\n", cache.Size())
    
    // Memory usage
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Memory usage: %d KB\n", m.Alloc/1024)
}
```

## 📝 Key Takeaways

1. **Caching is about speed**: Store frequently accessed data in fast storage
2. **Hit rate matters**: Higher hit rates mean better performance
3. **TTL is important**: Expire data to prevent stale information
4. **Thread safety**: Use mutexes for concurrent access
5. **Memory management**: Monitor memory usage and implement cleanup

## 🎯 Practice Questions

1. What's the difference between a cache hit and a cache miss?
2. Why is thread safety important in caching?
3. How does TTL (Time To Live) help with cache management?
4. What are the trade-offs of using more memory for caching?
5. How would you measure cache effectiveness?

## 🚀 Next Steps

Tomorrow we'll dive deeper into:
- Cache eviction policies (LRU, LFU, FIFO)
- Memory management techniques
- Advanced caching patterns

---

**Your Notes:**
<!-- Add your notes, questions, and insights here -->

**Code Experiments:**
<!-- Document any code experiments or modifications you try -->

**Questions:**
<!-- Write down any questions you have -->
