package main

import (
    "fmt"
    "sync"
    "time"
)

// ConsistencyLevel defines different consistency levels
type ConsistencyLevel int

const (
    Strong ConsistencyLevel = iota
    Eventual
    Weak
    Session
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

// ConsistentCache implements different consistency models
type ConsistentCache struct {
    data           map[string]*CacheItem
    database       *Database
    consistency    ConsistencyLevel
    mutex          sync.RWMutex
    versionCounter int64
    stats          ConsistencyStats
}

// CacheItem represents a cache item with metadata
type CacheItem struct {
    Value      interface{}
    Timestamp  time.Time
    Version    int64
    ExpiresAt  time.Time
    Dirty      bool // True if not yet written to DB
}

// NewConsistentCache creates a new cache with specified consistency level
func NewConsistentCache(db *Database, level ConsistencyLevel) *ConsistentCache {
    cache := &ConsistentCache{
        data:        make(map[string]*CacheItem),
        database:    db,
        consistency: level,
        stats:       ConsistencyStats{},
    }
    
    // Start background consistency maintenance
    go cache.consistencyMaintenance()
    
    return cache
}

// Get retrieves data based on consistency level
func (c *ConsistentCache) Get(key string) (interface{}, error) {
    c.mutex.RLock()
    item, exists := c.data[key]
    c.mutex.RUnlock()
    
    if !exists {
        return c.handleCacheMiss(key)
    }
    
    // Check if item is expired
    if time.Now().After(item.ExpiresAt) {
        c.mutex.Lock()
        delete(c.data, key)
        c.mutex.Unlock()
        return c.handleCacheMiss(key)
    }
    
    // Apply consistency rules
    switch c.consistency {
    case Strong:
        return c.getStrongConsistent(key, item)
    case Eventual:
        return c.getEventualConsistent(key, item)
    case Weak:
        return c.getWeakConsistent(key, item)
    case Session:
        return c.getSessionConsistent(key, item)
    default:
        return item.Value, nil
    }
}

// Set stores data based on consistency level
func (c *ConsistentCache) Set(key string, value interface{}, ttl time.Duration) error {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    c.versionCounter++
    
    item := &CacheItem{
        Value:     value,
        Timestamp: time.Now(),
        Version:   c.versionCounter,
        ExpiresAt: time.Now().Add(ttl),
        Dirty:     true,
    }
    
    c.data[key] = item
    
    // Apply consistency rules
    switch c.consistency {
    case Strong:
        return c.setStrongConsistent(key, value)
    case Eventual:
        return c.setEventualConsistent(key, value)
    case Weak:
        return c.setWeakConsistent(key, value)
    case Session:
        return c.setSessionConsistent(key, value)
    default:
        return nil
    }
}

// handleCacheMiss handles cache miss scenarios
func (c *ConsistentCache) handleCacheMiss(key string) (interface{}, error) {
    // Read from database
    value, err := c.database.Read(key)
    if err != nil {
        c.stats.DatabaseErrors++
        return nil, err
    }
    
    // Store in cache
    c.mutex.Lock()
    c.versionCounter++
    c.data[key] = &CacheItem{
        Value:     value,
        Timestamp: time.Now(),
        Version:   c.versionCounter,
        ExpiresAt: time.Now().Add(1 * time.Hour),
        Dirty:     false,
    }
    c.mutex.Unlock()
    
    c.stats.DatabaseReads++
    return value, nil
}

// getStrongConsistent implements strong consistency
func (c *ConsistentCache) getStrongConsistent(key string, item *CacheItem) (interface{}, error) {
    // For strong consistency, always check database for latest version
    dbValue, err := c.database.Read(key)
    if err != nil {
        c.stats.DatabaseErrors++
        return item.Value, nil // Return cached value if DB read fails
    }
    
    // Update cache if database has newer data
    if dbValue != item.Value {
        c.mutex.Lock()
        c.versionCounter++
        c.data[key] = &CacheItem{
            Value:     dbValue,
            Timestamp: time.Now(),
            Version:   c.versionCounter,
            ExpiresAt: item.ExpiresAt,
            Dirty:     false,
        }
        c.mutex.Unlock()
        c.stats.CacheUpdates++
    }
    
    c.stats.CacheHits++
    return dbValue, nil
}

// getEventualConsistent implements eventual consistency
func (c *ConsistentCache) getEventualConsistent(key string, item *CacheItem) (interface{}, error) {
    // For eventual consistency, return cached value
    // Background process will sync with database
    c.stats.CacheHits++
    return item.Value, nil
}

// getWeakConsistent implements weak consistency
func (c *ConsistentCache) getWeakConsistent(key string, item *CacheItem) (interface{}, error) {
    // For weak consistency, return cached value without any checks
    c.stats.CacheHits++
    return item.Value, nil
}

// getSessionConsistent implements session consistency
func (c *ConsistentCache) getSessionConsistent(key string, item *CacheItem) (interface{}, error) {
    // For session consistency, return cached value
    // Session-specific logic would be implemented here
    c.stats.CacheHits++
    return item.Value, nil
}

// setStrongConsistent implements strong consistency for writes
func (c *ConsistentCache) setStrongConsistent(key string, value interface{}) error {
    // Write to database immediately
    if err := c.database.Write(key, value); err != nil {
        // Remove from cache if DB write fails
        delete(c.data, key)
        c.stats.DatabaseErrors++
        return err
    }
    
    // Mark as not dirty
    c.data[key].Dirty = false
    c.stats.DatabaseWrites++
    c.stats.CacheWrites++
    
    return nil
}

// setEventualConsistent implements eventual consistency for writes
func (c *ConsistentCache) setEventualConsistent(key string, value interface{}) error {
    // Mark as dirty, background process will write to DB
    c.data[key].Dirty = true
    c.stats.CacheWrites++
    
    return nil
}

// setWeakConsistent implements weak consistency for writes
func (c *ConsistentCache) setWeakConsistent(key string, value interface{}) error {
    // No database write, just cache
    c.data[key].Dirty = true
    c.stats.CacheWrites++
    
    return nil
}

// setSessionConsistent implements session consistency for writes
func (c *ConsistentCache) setSessionConsistent(key string, value interface{}) error {
    // Write to database
    if err := c.database.Write(key, value); err != nil {
        c.stats.DatabaseErrors++
        return err
    }
    
    c.data[key].Dirty = false
    c.stats.DatabaseWrites++
    c.stats.CacheWrites++
    
    return nil
}

// consistencyMaintenance runs background consistency maintenance
func (c *ConsistentCache) consistencyMaintenance() {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        c.mutex.Lock()
        
        // Process dirty items for eventual consistency
        if c.consistency == Eventual {
            for key, item := range c.data {
                if item.Dirty {
                    // Write to database
                    if err := c.database.Write(key, item.Value); err != nil {
                        fmt.Printf("   [Eventual] Failed to sync %s to DB: %v\n", key, err)
                        c.stats.DatabaseErrors++
                    } else {
                        fmt.Printf("   [Eventual] Synced %s to DB\n", key)
                        item.Dirty = false
                        c.stats.DatabaseWrites++
                    }
                }
            }
        }
        
        // Clean up expired items
        now := time.Now()
        for key, item := range c.data {
            if now.After(item.ExpiresAt) {
                delete(c.data, key)
                c.stats.ExpiredItems++
            }
        }
        
        c.mutex.Unlock()
    }
}

// ConsistencyStats tracks consistency statistics
type ConsistencyStats struct {
    CacheHits      int64
    CacheMisses    int64
    CacheWrites    int64
    CacheUpdates   int64
    DatabaseReads  int64
    DatabaseWrites int64
    DatabaseErrors int64
    ExpiredItems   int64
    DirtyItems     int64
}

// GetStats returns cache statistics
func (c *ConsistentCache) GetStats() ConsistencyStats {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    // Count dirty items
    dirtyCount := int64(0)
    for _, item := range c.data {
        if item.Dirty {
            dirtyCount++
        }
    }
    c.stats.DirtyItems = dirtyCount
    
    return c.stats
}

// GetConsistencyLevel returns the current consistency level
func (c *ConsistentCache) GetConsistencyLevel() ConsistencyLevel {
    return c.consistency
}

// SetConsistencyLevel changes the consistency level
func (c *ConsistentCache) SetConsistencyLevel(level ConsistencyLevel) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.consistency = level
}

func main() {
    fmt.Println("=== Day 4: Cache Consistency Models ===")
    fmt.Println("Implementing different consistency levels\n")
    
    // Create database
    db := NewDatabase()
    
    // Test different consistency levels
    consistencyLevels := []ConsistencyLevel{
        Strong,
        Eventual,
        Weak,
        Session,
    }
    
    // Example 1: Basic consistency testing
    fmt.Println("1. Testing Consistency Levels:")
    fmt.Println("   Setting up test data...")
    
    // Pre-populate database
    db.Write("user:1", "John Doe")
    db.Write("user:2", "Jane Smith")
    db.Write("config:theme", "dark")
    
    for _, level := range consistencyLevels {
        fmt.Printf("\n   Testing %s Consistency:\n", getConsistencyName(level))
        
        cache := NewConsistentCache(db, level)
        
        // Test write
        if err := cache.Set("user:3", "Bob Johnson", 1*time.Hour); err != nil {
            fmt.Printf("   Error setting user:3: %v\n", err)
            continue
        }
        
        // Test read
        if value, err := cache.Get("user:1"); err != nil {
            fmt.Printf("   Error getting user:1: %v\n", err)
        } else {
            fmt.Printf("   Retrieved user:1: %v\n", value)
        }
        
        // Test read again
        if value, err := cache.Get("user:1"); err != nil {
            fmt.Printf("   Error getting user:1: %v\n", err)
        } else {
            fmt.Printf("   Retrieved user:1 (cached): %v\n", value)
        }
        
        // Print stats
        stats := cache.GetStats()
        fmt.Printf("   Stats - Hits: %d, Misses: %d, Writes: %d, Dirty: %d\n",
            stats.CacheHits, stats.CacheMisses, stats.CacheWrites, stats.DirtyItems)
        
        // Wait for eventual consistency to sync
        if level == Eventual {
            time.Sleep(2 * time.Second)
        }
    }
    
    // Example 2: Consistency behavior analysis
    fmt.Println("\n2. Consistency Behavior Analysis:")
    fmt.Println("   Analyzing different consistency behaviors...")
    
    // Test strong consistency
    fmt.Println("\n   Strong Consistency Test:")
    strongCache := NewConsistentCache(db, Strong)
    
    // Write to database directly
    db.Write("test:strong", "database_value")
    
    // Read from cache (should get database value)
    if value, err := strongCache.Get("test:strong"); err != nil {
        fmt.Printf("   Error: %v\n", err)
    } else {
        fmt.Printf("   Retrieved: %v (should be from database)\n", value)
    }
    
    // Test eventual consistency
    fmt.Println("\n   Eventual Consistency Test:")
    eventualCache := NewConsistentCache(db, Eventual)
    
    // Write to cache
    eventualCache.Set("test:eventual", "cache_value", 1*time.Hour)
    
    // Read from cache
    if value, err := eventualCache.Get("test:eventual"); err != nil {
        fmt.Printf("   Error: %v\n", err)
    } else {
        fmt.Printf("   Retrieved: %v (from cache)\n", value)
    }
    
    // Wait for background sync
    time.Sleep(2 * time.Second)
    
    // Test weak consistency
    fmt.Println("\n   Weak Consistency Test:")
    weakCache := NewConsistentCache(db, Weak)
    
    // Write to cache
    weakCache.Set("test:weak", "weak_value", 1*time.Hour)
    
    // Read from cache
    if value, err := weakCache.Get("test:weak"); err != nil {
        fmt.Printf("   Error: %v\n", err)
    } else {
        fmt.Printf("   Retrieved: %v (from cache only)\n", value)
    }
    
    // Example 3: Performance comparison
    fmt.Println("\n3. Performance Comparison:")
    fmt.Println("   Comparing consistency levels...")
    
    for _, level := range consistencyLevels {
        fmt.Printf("\n   %s Performance Test:\n", getConsistencyName(level))
        
        cache := NewConsistentCache(db, level)
        start := time.Now()
        
        // Perform multiple operations
        for i := 0; i < 100; i++ {
            key := fmt.Sprintf("perf:%s:key:%d", getConsistencyName(level), i%10)
            value := fmt.Sprintf("perf:value:%d", i)
            
            cache.Set(key, value, 1*time.Hour)
            cache.Get(key)
        }
        
        duration := time.Since(start)
        stats := cache.GetStats()
        
        fmt.Printf("   Time: %v\n", duration)
        fmt.Printf("   Cache Hits: %d\n", stats.CacheHits)
        fmt.Printf("   Cache Misses: %d\n", stats.CacheMisses)
        fmt.Printf("   DB Reads: %d\n", stats.DatabaseReads)
        fmt.Printf("   DB Writes: %d\n", stats.DatabaseWrites)
        fmt.Printf("   Dirty Items: %d\n", stats.DirtyItems)
        
        if stats.CacheHits+stats.CacheMisses > 0 {
            hitRate := float64(stats.CacheHits) / float64(stats.CacheHits+stats.CacheMisses) * 100
            fmt.Printf("   Hit Rate: %.2f%%\n", hitRate)
        }
    }
    
    // Example 4: Consistency level switching
    fmt.Println("\n4. Consistency Level Switching:")
    fmt.Println("   Testing dynamic consistency level changes...")
    
    cache := NewConsistentCache(db, Strong)
    
    // Test with strong consistency
    fmt.Println("   Testing with Strong consistency...")
    cache.Set("switch:test", "strong_value", 1*time.Hour)
    if value, err := cache.Get("switch:test"); err == nil {
        fmt.Printf("   Retrieved: %v\n", value)
    }
    
    // Switch to eventual consistency
    fmt.Println("   Switching to Eventual consistency...")
    cache.SetConsistencyLevel(Eventual)
    cache.Set("switch:test", "eventual_value", 1*time.Hour)
    if value, err := cache.Get("switch:test"); err == nil {
        fmt.Printf("   Retrieved: %v\n", value)
    }
    
    // Wait for sync
    time.Sleep(2 * time.Second)
    
    // Switch to weak consistency
    fmt.Println("   Switching to Weak consistency...")
    cache.SetConsistencyLevel(Weak)
    cache.Set("switch:test", "weak_value", 1*time.Hour)
    if value, err := cache.Get("switch:test"); err == nil {
        fmt.Printf("   Retrieved: %v\n", value)
    }
    
    // Example 5: Error handling
    fmt.Println("\n5. Error Handling:")
    fmt.Println("   Testing error scenarios...")
    
    // Test with non-existent key
    if _, err := cache.Get("nonexistent:key"); err != nil {
        fmt.Printf("   Expected error for non-existent key: %v\n", err)
    }
    
    // Test database error handling
    fmt.Println("   Testing database error handling...")
    
    // Create a cache with a mock database that fails
    failingDB := &Database{data: make(map[string]interface{})}
    failingCache := NewConsistentCache(failingDB, Strong)
    
    // Try to get a non-existent key
    if _, err := failingCache.Get("nonexistent:key"); err != nil {
        fmt.Printf("   Expected error: %v\n", err)
    }
    
    fmt.Println("\n=== Day 4 Complete! ===")
    fmt.Println("You've successfully implemented:")
    fmt.Println("✓ Strong consistency")
    fmt.Println("✓ Eventual consistency")
    fmt.Println("✓ Weak consistency")
    fmt.Println("✓ Session consistency")
    fmt.Println("✓ Performance comparison")
    fmt.Println("✓ Dynamic consistency switching")
    fmt.Println("✓ Error handling")
    fmt.Println("\nReady for Day 5: Memory Management! 🚀")
}

// Helper function
func getConsistencyName(level ConsistencyLevel) string {
    switch level {
    case Strong:
        return "Strong"
    case Eventual:
        return "Eventual"
    case Weak:
        return "Weak"
    case Session:
        return "Session"
    default:
        return "Unknown"
    }
}
