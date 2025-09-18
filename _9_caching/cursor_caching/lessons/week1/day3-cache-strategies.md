# Day 3: Cache Strategies & Patterns

## 🎯 Learning Objectives
- Understand different cache strategies (write-through, write-behind, etc.)
- Learn cache consistency models
- Implement cache-aside pattern
- Build a configurable caching system

## 📚 Theory: Cache Strategies

### What are Cache Strategies?
Cache strategies define **when** and **how** data is written to and read from the cache in relation to the primary data store.

## 🔄 Common Cache Strategies

### 1. Cache-Aside (Lazy Loading)
- **Read**: Application checks cache first, then database
- **Write**: Application writes to database, then updates cache
- **Use Case**: Most common pattern, flexible control

```
Read Flow:
1. Check cache
2. If hit: return data
3. If miss: read from DB
4. Store in cache
5. Return data

Write Flow:
1. Write to database
2. Update cache (or invalidate)
```

### 2. Read-Through
- **Read**: Cache automatically loads from database on miss
- **Write**: Application writes to cache, cache writes to database
- **Use Case**: When cache can handle database operations

```
Read Flow:
1. Check cache
2. If hit: return data
3. If miss: cache loads from DB
4. Return data

Write Flow:
1. Write to cache
2. Cache writes to database
```

### 3. Write-Through
- **Read**: Application reads from cache
- **Write**: Application writes to both cache and database simultaneously
- **Use Case**: When data consistency is critical

```
Write Flow:
1. Write to cache
2. Write to database
3. Both must succeed
```

### 4. Write-Behind (Write-Back)
- **Read**: Application reads from cache
- **Write**: Application writes to cache, database write is deferred
- **Use Case**: High write performance, eventual consistency

```
Write Flow:
1. Write to cache
2. Return immediately
3. Cache writes to DB asynchronously
```

### 5. Write-Around
- **Read**: Application reads from cache
- **Write**: Application writes directly to database, bypassing cache
- **Use Case**: When written data is rarely read

```
Write Flow:
1. Write to database
2. Skip cache
3. Cache will load on next read
```

## 🛠️ Hands-on: Implementing Cache Strategies

### Project Structure
```
projects/
├── cache-strategies/
│   ├── main.go
│   ├── strategies.go
│   ├── database.go
│   ├── cache.go
│   └── README.md
```

### Step 1: Database Interface

```go
// database.go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Database represents a mock database
type Database struct {
    data map[string]interface{}
    mutex sync.RWMutex
}

// NewDatabase creates a new database instance
func NewDatabase() *Database {
    return &Database{
        data: make(map[string]interface{}),
    }
}

// Read retrieves data from database
func (db *Database) Read(key string) (interface{}, error) {
    db.mutex.RLock()
    defer db.mutex.RUnlock()
    
    // Simulate database latency
    time.Sleep(10 * time.Millisecond)
    
    if value, exists := db.data[key]; exists {
        return value, nil
    }
    
    return nil, fmt.Errorf("key %s not found", key)
}

// Write stores data in database
func (db *Database) Write(key string, value interface{}) error {
    db.mutex.Lock()
    defer db.mutex.Unlock()
    
    // Simulate database latency
    time.Sleep(15 * time.Millisecond)
    
    db.data[key] = value
    return nil
}

// Delete removes data from database
func (db *Database) Delete(key string) error {
    db.mutex.Lock()
    defer db.mutex.Unlock()
    
    // Simulate database latency
    time.Sleep(5 * time.Millisecond)
    
    delete(db.data, key)
    return nil
}

// List returns all keys in database
func (db *Database) List() []string {
    db.mutex.RLock()
    defer db.mutex.RUnlock()
    
    keys := make([]string, 0, len(db.data))
    for key := range db.data {
        keys = append(keys, key)
    }
    return keys
}
```

### Step 2: Cache Strategies Implementation

```go
// strategies.go
package main

import (
    "fmt"
    "sync"
    "time"
)

// CacheStrategy defines the interface for cache strategies
type CacheStrategy interface {
    Get(key string) (interface{}, error)
    Set(key string, value interface{}) error
    Delete(key string) error
    GetName() string
}

// CacheAside implements the cache-aside pattern
type CacheAside struct {
    cache    *SimpleCache
    database *Database
    mutex    sync.RWMutex
}

// NewCacheAside creates a new cache-aside strategy
func NewCacheAside(cache *SimpleCache, db *Database) *CacheAside {
    return &CacheAside{
        cache:    cache,
        database: db,
    }
}

// Get retrieves data using cache-aside pattern
func (ca *CacheAside) Get(key string) (interface{}, error) {
    // Try cache first
    if value, found := ca.cache.Get(key); found {
        fmt.Printf("[Cache-Aside] Cache HIT for key: %s\n", key)
        return value, nil
    }
    
    fmt.Printf("[Cache-Aside] Cache MISS for key: %s, reading from DB\n", key)
    
    // Read from database
    value, err := ca.database.Read(key)
    if err != nil {
        return nil, err
    }
    
    // Store in cache for next time
    ca.cache.Set(key, value, 1*time.Hour)
    
    return value, nil
}

// Set stores data using cache-aside pattern
func (ca *CacheAside) Set(key string, value interface{}) error {
    // Write to database first
    if err := ca.database.Write(key, value); err != nil {
        return err
    }
    
    // Update cache
    ca.cache.Set(key, value, 1*time.Hour)
    
    fmt.Printf("[Cache-Aside] Updated DB and cache for key: %s\n", key)
    return nil
}

// Delete removes data using cache-aside pattern
func (ca *CacheAside) Delete(key string) error {
    // Delete from database
    if err := ca.database.Delete(key); err != nil {
        return err
    }
    
    // Remove from cache
    ca.cache.Delete(key)
    
    fmt.Printf("[Cache-Aside] Deleted from DB and cache for key: %s\n", key)
    return nil
}

// GetName returns the strategy name
func (ca *CacheAside) GetName() string {
    return "Cache-Aside"
}

// ReadThrough implements the read-through pattern
type ReadThrough struct {
    cache    *SimpleCache
    database *Database
    mutex    sync.RWMutex
}

// NewReadThrough creates a new read-through strategy
func NewReadThrough(cache *SimpleCache, db *Database) *ReadThrough {
    return &ReadThrough{
        cache:    cache,
        database: db,
    }
}

// Get retrieves data using read-through pattern
func (rt *ReadThrough) Get(key string) (interface{}, error) {
    // Try cache first
    if value, found := rt.cache.Get(key); found {
        fmt.Printf("[Read-Through] Cache HIT for key: %s\n", key)
        return value, nil
    }
    
    fmt.Printf("[Read-Through] Cache MISS for key: %s, loading from DB\n", key)
    
    // Cache automatically loads from database
    value, err := rt.database.Read(key)
    if err != nil {
        return nil, err
    }
    
    // Store in cache
    rt.cache.Set(key, value, 1*time.Hour)
    
    return value, nil
}

// Set stores data using read-through pattern
func (rt *ReadThrough) Set(key string, value interface{}) error {
    // Write to cache (cache handles database write)
    rt.cache.Set(key, value, 1*time.Hour)
    
    // Write to database
    if err := rt.database.Write(key, value); err != nil {
        // If DB write fails, remove from cache
        rt.cache.Delete(key)
        return err
    }
    
    fmt.Printf("[Read-Through] Updated cache and DB for key: %s\n", key)
    return nil
}

// Delete removes data using read-through pattern
func (rt *ReadThrough) Delete(key string) error {
    // Remove from cache
    rt.cache.Delete(key)
    
    // Delete from database
    if err := rt.database.Delete(key); err != nil {
        return err
    }
    
    fmt.Printf("[Read-Through] Deleted from cache and DB for key: %s\n", key)
    return nil
}

// GetName returns the strategy name
func (rt *ReadThrough) GetName() string {
    return "Read-Through"
}

// WriteThrough implements the write-through pattern
type WriteThrough struct {
    cache    *SimpleCache
    database *Database
    mutex    sync.RWMutex
}

// NewWriteThrough creates a new write-through strategy
func NewWriteThrough(cache *SimpleCache, db *Database) *WriteThrough {
    return &WriteThrough{
        cache:    cache,
        database: db,
    }
}

// Get retrieves data using write-through pattern
func (wt *WriteThrough) Get(key string) (interface{}, error) {
    // Try cache first
    if value, found := wt.cache.Get(key); found {
        fmt.Printf("[Write-Through] Cache HIT for key: %s\n", key)
        return value, nil
    }
    
    fmt.Printf("[Write-Through] Cache MISS for key: %s, reading from DB\n", key)
    
    // Read from database
    value, err := wt.database.Read(key)
    if err != nil {
        return nil, err
    }
    
    // Store in cache
    wt.cache.Set(key, value, 1*time.Hour)
    
    return value, nil
}

// Set stores data using write-through pattern
func (wt *WriteThrough) Set(key string, value interface{}) error {
    // Write to both cache and database simultaneously
    wt.cache.Set(key, value, 1*time.Hour)
    
    if err := wt.database.Write(key, value); err != nil {
        // If DB write fails, remove from cache
        wt.cache.Delete(key)
        return err
    }
    
    fmt.Printf("[Write-Through] Updated cache and DB simultaneously for key: %s\n", key)
    return nil
}

// Delete removes data using write-through pattern
func (wt *WriteThrough) Delete(key string) error {
    // Remove from both cache and database
    wt.cache.Delete(key)
    
    if err := wt.database.Delete(key); err != nil {
        return err
    }
    
    fmt.Printf("[Write-Through] Deleted from cache and DB for key: %s\n", key)
    return nil
}

// GetName returns the strategy name
func (wt *WriteThrough) GetName() string {
    return "Write-Through"
}

// WriteBehind implements the write-behind pattern
type WriteBehind struct {
    cache    *SimpleCache
    database *Database
    mutex    sync.RWMutex
    queue    chan WriteOperation
    stop     chan bool
}

// WriteOperation represents a write operation to be processed
type WriteOperation struct {
    Key   string
    Value interface{}
}

// NewWriteBehind creates a new write-behind strategy
func NewWriteBehind(cache *SimpleCache, db *Database) *WriteBehind {
    wb := &WriteBehind{
        cache:    cache,
        database: db,
        queue:    make(chan WriteOperation, 1000),
        stop:     make(chan bool),
    }
    
    // Start background writer
    go wb.backgroundWriter()
    
    return wb
}

// backgroundWriter processes write operations asynchronously
func (wb *WriteBehind) backgroundWriter() {
    for {
        select {
        case op := <-wb.queue:
            // Write to database asynchronously
            if err := wb.database.Write(op.Key, op.Value); err != nil {
                fmt.Printf("[Write-Behind] Failed to write %s to DB: %v\n", op.Key, err)
            } else {
                fmt.Printf("[Write-Behind] Successfully wrote %s to DB\n", op.Key)
            }
        case <-wb.stop:
            return
        }
    }
}

// Get retrieves data using write-behind pattern
func (wb *WriteBehind) Get(key string) (interface{}, error) {
    // Try cache first
    if value, found := wb.cache.Get(key); found {
        fmt.Printf("[Write-Behind] Cache HIT for key: %s\n", key)
        return value, nil
    }
    
    fmt.Printf("[Write-Behind] Cache MISS for key: %s, reading from DB\n", key)
    
    // Read from database
    value, err := wb.database.Read(key)
    if err != nil {
        return nil, err
    }
    
    // Store in cache
    wb.cache.Set(key, value, 1*time.Hour)
    
    return value, nil
}

// Set stores data using write-behind pattern
func (wb *WriteBehind) Set(key string, value interface{}) error {
    // Write to cache immediately
    wb.cache.Set(key, value, 1*time.Hour)
    
    // Queue database write
    select {
    case wb.queue <- WriteOperation{Key: key, Value: value}:
        fmt.Printf("[Write-Behind] Queued write for key: %s\n", key)
    default:
        fmt.Printf("[Write-Behind] Queue full, dropping write for key: %s\n", key)
    }
    
    return nil
}

// Delete removes data using write-behind pattern
func (wb *WriteBehind) Delete(key string) error {
    // Remove from cache
    wb.cache.Delete(key)
    
    // Delete from database
    if err := wb.database.Delete(key); err != nil {
        return err
    }
    
    fmt.Printf("[Write-Behind] Deleted from cache and DB for key: %s\n", key)
    return nil
}

// GetName returns the strategy name
func (wb *WriteBehind) GetName() string {
    return "Write-Behind"
}

// Stop stops the background writer
func (wb *WriteBehind) Stop() {
    close(wb.stop)
}
```

### Step 3: Main Application

```go
// main.go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Create database and cache
    db := NewDatabase()
    cache := NewSimpleCache()
    
    // Test different strategies
    strategies := []CacheStrategy{
        NewCacheAside(cache, db),
        NewReadThrough(cache, db),
        NewWriteThrough(cache, db),
        NewWriteBehind(cache, db),
    }
    
    for _, strategy := range strategies {
        fmt.Printf("\n=== Testing %s Strategy ===\n", strategy.GetName())
        
        // Test write
        if err := strategy.Set("user:1", "John Doe"); err != nil {
            fmt.Printf("Error setting user:1: %v\n", err)
            continue
        }
        
        // Test read
        if value, err := strategy.Get("user:1"); err != nil {
            fmt.Printf("Error getting user:1: %v\n", err)
        } else {
            fmt.Printf("Retrieved user:1: %v\n", value)
        }
        
        // Test read again (should hit cache)
        if value, err := strategy.Get("user:1"); err != nil {
            fmt.Printf("Error getting user:1: %v\n", err)
        } else {
            fmt.Printf("Retrieved user:1 (cached): %v\n", value)
        }
        
        // Test non-existent key
        if value, err := strategy.Get("user:999"); err != nil {
            fmt.Printf("Error getting user:999: %v\n", err)
        } else {
            fmt.Printf("Retrieved user:999: %v\n", value)
        }
        
        // Wait a bit for write-behind to process
        if strategy.GetName() == "Write-Behind" {
            time.Sleep(100 * time.Millisecond)
        }
    }
    
    // Test write-behind cleanup
    if wb, ok := strategies[3].(*WriteBehind); ok {
        wb.Stop()
    }
}
```

## 🧪 Exercise: Performance Comparison

### Task
Create a benchmark to compare the performance of different cache strategies:

```go
// benchmark.go
package main

import (
    "fmt"
    "testing"
    "time"
)

func BenchmarkCacheStrategies(b *testing.B) {
    db := NewDatabase()
    cache := NewSimpleCache()
    
    strategies := []CacheStrategy{
        NewCacheAside(cache, db),
        NewReadThrough(cache, db),
        NewWriteThrough(cache, db),
        NewWriteBehind(cache, db),
    }
    
    for _, strategy := range strategies {
        b.Run(strategy.GetName(), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                key := fmt.Sprintf("key_%d", i%100)
                value := fmt.Sprintf("value_%d", i)
                
                // Test write
                strategy.Set(key, value)
                
                // Test read
                strategy.Get(key)
            }
        })
    }
}
```

## 📝 Key Takeaways

1. **Cache-Aside**: Most flexible, application controls everything
2. **Read-Through**: Good when cache can handle DB operations
3. **Write-Through**: Best for data consistency
4. **Write-Behind**: Best for write performance
5. **Choose based on requirements**: Consistency vs Performance

## 🎯 Practice Questions

1. When would you use write-behind vs write-through?
2. What are the trade-offs of each strategy?
3. How do you handle failures in write-through?
4. What happens if write-behind queue is full?
5. How do you ensure consistency across strategies?

## 🚀 Next Steps

Tomorrow we'll cover:
- Cache consistency models
- Cache invalidation strategies
- Advanced caching patterns

---

**Your Notes:**
<!-- Add your notes, questions, and insights here -->

**Code Experiments:**
<!-- Document any code experiments or modifications you try -->

**Questions:**
<!-- Write down any questions you have -->
