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

// SimpleCache represents our basic cache
type SimpleCache struct {
    items map[string]interface{}
    mutex sync.RWMutex
}

// NewSimpleCache creates a new cache instance
func NewSimpleCache() *SimpleCache {
    return &SimpleCache{
        items: make(map[string]interface{}),
    }
}

// Get retrieves a value from the cache
func (c *SimpleCache) Get(key string) (interface{}, bool) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    if value, exists := c.items[key]; exists {
        return value, true
    }
    
    return nil, false
}

// Set stores a value in the cache
func (c *SimpleCache) Set(key string, value interface{}) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.items[key] = value
}

// Delete removes an item from the cache
func (c *SimpleCache) Delete(key string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    delete(c.items, key)
}

// CacheStrategy defines the interface for cache strategies
type CacheStrategy interface {
    Get(key string) (interface{}, error)
    Set(key string, value interface{}) error
    Delete(key string) error
    GetName() string
    GetStats() StrategyStats
}

// CacheAside implements the cache-aside pattern
type CacheAside struct {
    cache    *SimpleCache
    database *Database
    mutex    sync.RWMutex
    stats    StrategyStats
}

// NewCacheAside creates a new cache-aside strategy
func NewCacheAside(cache *SimpleCache, db *Database) *CacheAside {
    return &CacheAside{
        cache:    cache,
        database: db,
        stats:    StrategyStats{},
    }
}

// Get retrieves data using cache-aside pattern
func (ca *CacheAside) Get(key string) (interface{}, error) {
    // Try cache first
    if value, found := ca.cache.Get(key); found {
        fmt.Printf("   [Cache-Aside] Cache HIT for key: %s\n", key)
        ca.stats.CacheHits++
        return value, nil
    }
    
    fmt.Printf("   [Cache-Aside] Cache MISS for key: %s, reading from DB\n", key)
    ca.stats.CacheMisses++
    
    // Read from database
    value, err := ca.database.Read(key)
    if err != nil {
        ca.stats.DatabaseErrors++
        return nil, err
    }
    
    // Store in cache for next time
    ca.cache.Set(key, value)
    ca.stats.DatabaseReads++
    
    return value, nil
}

// Set stores data using cache-aside pattern
func (ca *CacheAside) Set(key string, value interface{}) error {
    // Write to database first
    if err := ca.database.Write(key, value); err != nil {
        ca.stats.DatabaseErrors++
        return err
    }
    
    // Update cache
    ca.cache.Set(key, value)
    ca.stats.DatabaseWrites++
    
    fmt.Printf("   [Cache-Aside] Updated DB and cache for key: %s\n", key)
    return nil
}

// Delete removes data using cache-aside pattern
func (ca *CacheAside) Delete(key string) error {
    // Delete from database
    if err := ca.database.Delete(key); err != nil {
        ca.stats.DatabaseErrors++
        return err
    }
    
    // Remove from cache
    ca.cache.Delete(key)
    ca.stats.DatabaseWrites++
    
    fmt.Printf("   [Cache-Aside] Deleted from DB and cache for key: %s\n", key)
    return nil
}

// GetName returns the strategy name
func (ca *CacheAside) GetName() string {
    return "Cache-Aside"
}

// GetStats returns strategy statistics
func (ca *CacheAside) GetStats() StrategyStats {
    return ca.stats
}

// ReadThrough implements the read-through pattern
type ReadThrough struct {
    cache    *SimpleCache
    database *Database
    mutex    sync.RWMutex
    stats    StrategyStats
}

// NewReadThrough creates a new read-through strategy
func NewReadThrough(cache *SimpleCache, db *Database) *ReadThrough {
    return &ReadThrough{
        cache:    cache,
        database: db,
        stats:    StrategyStats{},
    }
}

// Get retrieves data using read-through pattern
func (rt *ReadThrough) Get(key string) (interface{}, error) {
    // Try cache first
    if value, found := rt.cache.Get(key); found {
        fmt.Printf("   [Read-Through] Cache HIT for key: %s\n", key)
        rt.stats.CacheHits++
        return value, nil
    }
    
    fmt.Printf("   [Read-Through] Cache MISS for key: %s, loading from DB\n", key)
    rt.stats.CacheMisses++
    
    // Cache automatically loads from database
    value, err := rt.database.Read(key)
    if err != nil {
        rt.stats.DatabaseErrors++
        return nil, err
    }
    
    // Store in cache
    rt.cache.Set(key, value)
    rt.stats.DatabaseReads++
    
    return value, nil
}

// Set stores data using read-through pattern
func (rt *ReadThrough) Set(key string, value interface{}) error {
    // Write to cache (cache handles database write)
    rt.cache.Set(key, value)
    
    // Write to database
    if err := rt.database.Write(key, value); err != nil {
        // If DB write fails, remove from cache
        rt.cache.Delete(key)
        rt.stats.DatabaseErrors++
        return err
    }
    
    rt.stats.DatabaseWrites++
    fmt.Printf("   [Read-Through] Updated cache and DB for key: %s\n", key)
    return nil
}

// Delete removes data using read-through pattern
func (rt *ReadThrough) Delete(key string) error {
    // Remove from cache
    rt.cache.Delete(key)
    
    // Delete from database
    if err := rt.database.Delete(key); err != nil {
        rt.stats.DatabaseErrors++
        return err
    }
    
    rt.stats.DatabaseWrites++
    fmt.Printf("   [Read-Through] Deleted from cache and DB for key: %s\n", key)
    return nil
}

// GetName returns the strategy name
func (rt *ReadThrough) GetName() string {
    return "Read-Through"
}

// GetStats returns strategy statistics
func (rt *ReadThrough) GetStats() StrategyStats {
    return rt.stats
}

// WriteThrough implements the write-through pattern
type WriteThrough struct {
    cache    *SimpleCache
    database *Database
    mutex    sync.RWMutex
    stats    StrategyStats
}

// NewWriteThrough creates a new write-through strategy
func NewWriteThrough(cache *SimpleCache, db *Database) *WriteThrough {
    return &WriteThrough{
        cache:    cache,
        database: db,
        stats:    StrategyStats{},
    }
}

// Get retrieves data using write-through pattern
func (wt *WriteThrough) Get(key string) (interface{}, error) {
    // Try cache first
    if value, found := wt.cache.Get(key); found {
        fmt.Printf("   [Write-Through] Cache HIT for key: %s\n", key)
        wt.stats.CacheHits++
        return value, nil
    }
    
    fmt.Printf("   [Write-Through] Cache MISS for key: %s, reading from DB\n", key)
    wt.stats.CacheMisses++
    
    // Read from database
    value, err := wt.database.Read(key)
    if err != nil {
        wt.stats.DatabaseErrors++
        return nil, err
    }
    
    // Store in cache
    wt.cache.Set(key, value)
    wt.stats.DatabaseReads++
    
    return value, nil
}

// Set stores data using write-through pattern
func (wt *WriteThrough) Set(key string, value interface{}) error {
    // Write to both cache and database simultaneously
    wt.cache.Set(key, value)
    
    if err := wt.database.Write(key, value); err != nil {
        // If DB write fails, remove from cache
        wt.cache.Delete(key)
        wt.stats.DatabaseErrors++
        return err
    }
    
    wt.stats.DatabaseWrites++
    fmt.Printf("   [Write-Through] Updated cache and DB simultaneously for key: %s\n", key)
    return nil
}

// Delete removes data using write-through pattern
func (wt *WriteThrough) Delete(key string) error {
    // Remove from both cache and database
    wt.cache.Delete(key)
    
    if err := wt.database.Delete(key); err != nil {
        wt.stats.DatabaseErrors++
        return err
    }
    
    wt.stats.DatabaseWrites++
    fmt.Printf("   [Write-Through] Deleted from cache and DB for key: %s\n", key)
    return nil
}

// GetName returns the strategy name
func (wt *WriteThrough) GetName() string {
    return "Write-Through"
}

// GetStats returns strategy statistics
func (wt *WriteThrough) GetStats() StrategyStats {
    return wt.stats
}

// WriteBehind implements the write-behind pattern
type WriteBehind struct {
    cache    *SimpleCache
    database *Database
    mutex    sync.RWMutex
    queue    chan WriteOperation
    stop     chan bool
    stats    StrategyStats
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
        stats:    StrategyStats{},
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
                fmt.Printf("   [Write-Behind] Failed to write %s to DB: %v\n", op.Key, err)
                wb.stats.DatabaseErrors++
            } else {
                fmt.Printf("   [Write-Behind] Successfully wrote %s to DB\n", op.Key)
                wb.stats.DatabaseWrites++
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
        fmt.Printf("   [Write-Behind] Cache HIT for key: %s\n", key)
        wb.stats.CacheHits++
        return value, nil
    }
    
    fmt.Printf("   [Write-Behind] Cache MISS for key: %s, reading from DB\n", key)
    wb.stats.CacheMisses++
    
    // Read from database
    value, err := wb.database.Read(key)
    if err != nil {
        wb.stats.DatabaseErrors++
        return nil, err
    }
    
    // Store in cache
    wb.cache.Set(key, value)
    wb.stats.DatabaseReads++
    
    return value, nil
}

// Set stores data using write-behind pattern
func (wb *WriteBehind) Set(key string, value interface{}) error {
    // Write to cache immediately
    wb.cache.Set(key, value)
    
    // Queue database write
    select {
    case wb.queue <- WriteOperation{Key: key, Value: value}:
        fmt.Printf("   [Write-Behind] Queued write for key: %s\n", key)
    default:
        fmt.Printf("   [Write-Behind] Queue full, dropping write for key: %s\n", key)
    }
    
    return nil
}

// Delete removes data using write-behind pattern
func (wb *WriteBehind) Delete(key string) error {
    // Remove from cache
    wb.cache.Delete(key)
    
    // Delete from database
    if err := wb.database.Delete(key); err != nil {
        wb.stats.DatabaseErrors++
        return err
    }
    
    wb.stats.DatabaseWrites++
    fmt.Printf("   [Write-Behind] Deleted from cache and DB for key: %s\n", key)
    return nil
}

// GetName returns the strategy name
func (wb *WriteBehind) GetName() string {
    return "Write-Behind"
}

// GetStats returns strategy statistics
func (wb *WriteBehind) GetStats() StrategyStats {
    return wb.stats
}

// Stop stops the background writer
func (wb *WriteBehind) Stop() {
    close(wb.stop)
}

// StrategyStats holds statistics for cache strategies
type StrategyStats struct {
    CacheHits     int64
    CacheMisses   int64
    DatabaseReads int64
    DatabaseWrites int64
    DatabaseErrors int64
}

func main() {
    fmt.Println("=== Day 3: Cache Strategies & Patterns ===")
    fmt.Println("Implementing different caching strategies\n")
    
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
    
    // Example 1: Basic strategy testing
    fmt.Println("1. Testing Cache Strategies:")
    fmt.Println("   Setting up test data...")
    
    // Pre-populate database
    db.Write("user:1", "John Doe")
    db.Write("user:2", "Jane Smith")
    db.Write("config:theme", "dark")
    
    for _, strategy := range strategies {
        fmt.Printf("\n   Testing %s Strategy:\n", strategy.GetName())
        
        // Test write
        if err := strategy.Set("user:3", "Bob Johnson"); err != nil {
            fmt.Printf("   Error setting user:3: %v\n", err)
            continue
        }
        
        // Test read
        if value, err := strategy.Get("user:1"); err != nil {
            fmt.Printf("   Error getting user:1: %v\n", err)
        } else {
            fmt.Printf("   Retrieved user:1: %v\n", value)
        }
        
        // Test read again (should hit cache)
        if value, err := strategy.Get("user:1"); err != nil {
            fmt.Printf("   Error getting user:1: %v\n", err)
        } else {
            fmt.Printf("   Retrieved user:1 (cached): %v\n", value)
        }
        
        // Test non-existent key
        if value, err := strategy.Get("user:999"); err != nil {
            fmt.Printf("   Error getting user:999: %v\n", err)
        } else {
            fmt.Printf("   Retrieved user:999: %v\n", value)
        }
        
        // Wait a bit for write-behind to process
        if strategy.GetName() == "Write-Behind" {
            time.Sleep(100 * time.Millisecond)
        }
    }
    
    // Example 2: Performance comparison
    fmt.Println("\n2. Performance Comparison:")
    fmt.Println("   Comparing different strategies...")
    
    for _, strategy := range strategies {
        fmt.Printf("\n   %s Performance Test:\n", strategy.GetName())
        
        start := time.Now()
        
        // Perform multiple operations
        for i := 0; i < 100; i++ {
            key := fmt.Sprintf("perf:key:%d", i%10)
            value := fmt.Sprintf("perf:value:%d", i)
            
            strategy.Set(key, value)
            strategy.Get(key)
        }
        
        duration := time.Since(start)
        stats := strategy.GetStats()
        
        fmt.Printf("   Time: %v\n", duration)
        fmt.Printf("   Cache Hits: %d\n", stats.CacheHits)
        fmt.Printf("   Cache Misses: %d\n", stats.CacheMisses)
        fmt.Printf("   DB Reads: %d\n", stats.DatabaseReads)
        fmt.Printf("   DB Writes: %d\n", stats.DatabaseWrites)
        fmt.Printf("   DB Errors: %d\n", stats.DatabaseErrors)
        
        if stats.CacheHits+stats.CacheMisses > 0 {
            hitRate := float64(stats.CacheHits) / float64(stats.CacheHits+stats.CacheMisses) * 100
            fmt.Printf("   Hit Rate: %.2f%%\n", hitRate)
        }
    }
    
    // Example 3: Strategy-specific behavior
    fmt.Println("\n3. Strategy-Specific Behavior:")
    
    // Test write-behind queue behavior
    fmt.Println("   Testing Write-Behind queue behavior...")
    wbStrategy := NewWriteBehind(NewSimpleCache(), db)
    
    // Rapid writes to test queue
    for i := 0; i < 5; i++ {
        key := fmt.Sprintf("queue:test:%d", i)
        value := fmt.Sprintf("queue:value:%d", i)
        wbStrategy.Set(key, value)
    }
    
    // Wait for queue processing
    time.Sleep(200 * time.Millisecond)
    
    // Test read-after-write consistency
    fmt.Println("   Testing read-after-write consistency...")
    for i := 0; i < 5; i++ {
        key := fmt.Sprintf("queue:test:%d", i)
        if value, err := wbStrategy.Get(key); err != nil {
            fmt.Printf("   Error reading %s: %v\n", key, err)
        } else {
            fmt.Printf("   Read %s: %v\n", key, value)
        }
    }
    
    // Stop write-behind strategy
    wbStrategy.Stop()
    
    // Example 4: Error handling
    fmt.Println("\n4. Error Handling:")
    fmt.Println("   Testing error scenarios...")
    
    // Test with non-existent key
    strategy := NewCacheAside(NewSimpleCache(), db)
    if _, err := strategy.Get("nonexistent:key"); err != nil {
        fmt.Printf("   Expected error for non-existent key: %v\n", err)
    }
    
    // Test delete non-existent key
    if err := strategy.Delete("nonexistent:key"); err != nil {
        fmt.Printf("   Expected error deleting non-existent key: %v\n", err)
    }
    
    fmt.Println("\n=== Day 3 Complete! ===")
    fmt.Println("You've successfully implemented:")
    fmt.Println("✓ Cache-Aside pattern")
    fmt.Println("✓ Read-Through pattern")
    fmt.Println("✓ Write-Through pattern")
    fmt.Println("✓ Write-Behind pattern")
    fmt.Println("✓ Performance comparison")
    fmt.Println("✓ Error handling")
    fmt.Println("\nReady for Day 4: Cache Consistency Models! 🚀")
}
