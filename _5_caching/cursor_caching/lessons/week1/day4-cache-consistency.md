# Day 4: Cache Consistency Models

## 🎯 Learning Objectives
- Understand different cache consistency models
- Learn about cache invalidation strategies
- Implement eventual consistency patterns
- Build a cache with configurable consistency levels

## 📚 Theory: Cache Consistency Models

### What is Cache Consistency?
Cache consistency refers to how up-to-date the data in the cache is compared to the source of truth (database). Different consistency models provide different guarantees about data freshness.

## 🔄 Consistency Models

### 1. Strong Consistency
- **Definition**: All reads return the most recent write
- **Guarantee**: Linearizability - operations appear to execute atomically
- **Use Case**: Financial systems, critical data
- **Trade-off**: Higher latency, lower availability

### 2. Eventual Consistency
- **Definition**: System will become consistent over time
- **Guarantee**: No guarantees about when consistency is achieved
- **Use Case**: Social media, content delivery
- **Trade-off**: May read stale data temporarily

### 3. Weak Consistency
- **Definition**: No guarantees about consistency
- **Guarantee**: System may never become consistent
- **Use Case**: Real-time systems, performance-critical
- **Trade-off**: Fastest performance, no consistency guarantees

### 4. Session Consistency
- **Definition**: Consistency within a user session
- **Guarantee**: User sees their own writes immediately
- **Use Case**: Web applications, user-specific data
- **Trade-off**: Good balance of performance and consistency

## 🛠️ Hands-on: Implementing Consistency Models

### Project Structure
```
projects/
├── cache-consistency/
│   ├── main.go
│   ├── consistency.go
│   ├── invalidation.go
│   ├── database.go
│   └── README.md
```

### Step 1: Consistency Models Implementation

```go
// consistency.go
package main

import (
    "fmt"
    "sync"
    "time"
)

// ConsistencyLevel defines different consistency levels
type ConsistencyLevel int

const (
    Strong Consistency ConsistencyLevel = iota
    Eventual Consistency
    Weak Consistency
    Session Consistency
)

// CacheItem represents a cache item with metadata
type CacheItem struct {
    Value      interface{}
    Timestamp  time.Time
    Version    int64
    ExpiresAt  time.Time
    Dirty      bool // True if not yet written to DB
}

// ConsistentCache implements different consistency models
type ConsistentCache struct {
    data           map[string]*CacheItem
    database       *Database
    consistency    ConsistencyLevel
    mutex          sync.RWMutex
    versionCounter int64
    invalidation   *InvalidationManager
}

// NewConsistentCache creates a new cache with specified consistency level
func NewConsistentCache(db *Database, level ConsistencyLevel) *ConsistentCache {
    cache := &ConsistentCache{
        data:         make(map[string]*CacheItem),
        database:     db,
        consistency:  level,
        invalidation: NewInvalidationManager(),
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
    case Strong Consistency:
        return c.getStrongConsistent(key, item)
    case Eventual Consistency:
        return c.getEventualConsistent(key, item)
    case Weak Consistency:
        return c.getWeakConsistent(key, item)
    case Session Consistency:
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
    case Strong Consistency:
        return c.setStrongConsistent(key, value)
    case Eventual Consistency:
        return c.setEventualConsistent(key, value)
    case Weak Consistency:
        return c.setWeakConsistent(key, value)
    case Session Consistency:
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
    
    return value, nil
}

// getStrongConsistent implements strong consistency
func (c *ConsistentCache) getStrongConsistent(key string, item *CacheItem) (interface{}, error) {
    // For strong consistency, always check database for latest version
    dbValue, err := c.database.Read(key)
    if err != nil {
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
    }
    
    return dbValue, nil
}

// getEventualConsistent implements eventual consistency
func (c *ConsistentCache) getEventualConsistent(key string, item *CacheItem) (interface{}, error) {
    // For eventual consistency, return cached value
    // Background process will sync with database
    return item.Value, nil
}

// getWeakConsistent implements weak consistency
func (c *ConsistentCache) getWeakConsistent(key string, item *CacheItem) (interface{}, error) {
    // For weak consistency, return cached value without any checks
    return item.Value, nil
}

// getSessionConsistent implements session consistency
func (c *ConsistentCache) getSessionConsistent(key string, item *CacheItem) (interface{}, error) {
    // For session consistency, return cached value
    // Session-specific logic would be implemented here
    return item.Value, nil
}

// setStrongConsistent implements strong consistency for writes
func (c *ConsistentCache) setStrongConsistent(key string, value interface{}) error {
    // Write to database immediately
    if err := c.database.Write(key, value); err != nil {
        // Remove from cache if DB write fails
        delete(c.data, key)
        return err
    }
    
    // Mark as not dirty
    c.data[key].Dirty = false
    
    // Invalidate other caches
    c.invalidation.Invalidate(key)
    
    return nil
}

// setEventualConsistent implements eventual consistency for writes
func (c *ConsistentCache) setEventualConsistent(key string, value interface{}) error {
    // Mark as dirty, background process will write to DB
    c.data[key].Dirty = true
    
    // Invalidate other caches
    c.invalidation.Invalidate(key)
    
    return nil
}

// setWeakConsistent implements weak consistency for writes
func (c *ConsistentCache) setWeakConsistent(key string, value interface{}) error {
    // No database write, just cache
    c.data[key].Dirty = true
    return nil
}

// setSessionConsistent implements session consistency for writes
func (c *ConsistentCache) setSessionConsistent(key string, value interface{}) error {
    // Write to database
    if err := c.database.Write(key, value); err != nil {
        return err
    }
    
    c.data[key].Dirty = false
    return nil
}

// consistencyMaintenance runs background consistency maintenance
func (c *ConsistentCache) consistencyMaintenance() {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        c.mutex.Lock()
        
        // Process dirty items for eventual consistency
        if c.consistency == Eventual Consistency {
            for key, item := range c.data {
                if item.Dirty {
                    // Write to database
                    if err := c.database.Write(key, item.Value); err != nil {
                        fmt.Printf("Failed to sync %s to DB: %v\n", key, err)
                    } else {
                        item.Dirty = false
                        fmt.Printf("Synced %s to DB\n", key)
                    }
                }
            }
        }
        
        // Clean up expired items
        now := time.Now()
        for key, item := range c.data {
            if now.After(item.ExpiresAt) {
                delete(c.data, key)
            }
        }
        
        c.mutex.Unlock()
    }
}

// GetStats returns cache statistics
func (c *ConsistentCache) GetStats() map[string]interface{} {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    dirtyCount := 0
    for _, item := range c.data {
        if item.Dirty {
            dirtyCount++
        }
    }
    
    return map[string]interface{}{
        "size":        len(c.data),
        "dirty_items": dirtyCount,
        "consistency": c.consistency,
    }
}
```

### Step 2: Cache Invalidation Implementation

```go
// invalidation.go
package main

import (
    "fmt"
    "sync"
    "time"
)

// InvalidationManager manages cache invalidation
type InvalidationManager struct {
    subscribers map[string][]chan string
    mutex       sync.RWMutex
}

// NewInvalidationManager creates a new invalidation manager
func NewInvalidationManager() *InvalidationManager {
    return &InvalidationManager{
        subscribers: make(map[string][]chan string),
    }
}

// Subscribe subscribes to invalidation events for a key pattern
func (im *InvalidationManager) Subscribe(pattern string) <-chan string {
    im.mutex.Lock()
    defer im.mutex.Unlock()
    
    ch := make(chan string, 100)
    im.subscribers[pattern] = append(im.subscribers[pattern], ch)
    
    return ch
}

// Invalidate invalidates caches for a specific key
func (im *InvalidationManager) Invalidate(key string) {
    im.mutex.RLock()
    defer im.mutex.RUnlock()
    
    // Notify all subscribers
    for pattern, channels := range im.subscribers {
        if im.matchesPattern(key, pattern) {
            for _, ch := range channels {
                select {
                case ch <- key:
                default:
                    // Channel is full, skip
                }
            }
        }
    }
}

// matchesPattern checks if a key matches a pattern
func (im *InvalidationManager) matchesPattern(key, pattern string) bool {
    // Simple pattern matching - can be enhanced
    if pattern == "*" {
        return true
    }
    if pattern == key {
        return true
    }
    // Add more sophisticated pattern matching here
    return false
}

// InvalidationStrategy defines different invalidation strategies
type InvalidationStrategy int

const (
    Immediate InvalidationStrategy = iota
    Lazy
    TimeBased
    EventBased
)

// SmartInvalidation implements intelligent invalidation
type SmartInvalidation struct {
    strategy InvalidationStrategy
    manager  *InvalidationManager
    timers   map[string]*time.Timer
    mutex    sync.RWMutex
}

// NewSmartInvalidation creates a new smart invalidation system
func NewSmartInvalidation(strategy InvalidationStrategy) *SmartInvalidation {
    return &SmartInvalidation{
        strategy: strategy,
        manager:  NewInvalidationManager(),
        timers:   make(map[string]*time.Timer),
    }
}

// Invalidate invalidates a key using the specified strategy
func (si *SmartInvalidation) Invalidate(key string) {
    switch si.strategy {
    case Immediate:
        si.immediateInvalidation(key)
    case Lazy:
        si.lazyInvalidation(key)
    case TimeBased:
        si.timeBasedInvalidation(key)
    case EventBased:
        si.eventBasedInvalidation(key)
    }
}

// immediateInvalidation invalidates immediately
func (si *SmartInvalidation) immediateInvalidation(key string) {
    si.manager.Invalidate(key)
    fmt.Printf("Immediately invalidated: %s\n", key)
}

// lazyInvalidation invalidates on next access
func (si *SmartInvalidation) lazyInvalidation(key string) {
    // Mark for lazy invalidation
    fmt.Printf("Marked for lazy invalidation: %s\n", key)
}

// timeBasedInvalidation invalidates after a delay
func (si *SmartInvalidation) timeBasedInvalidation(key string) {
    si.mutex.Lock()
    defer si.mutex.Unlock()
    
    // Cancel existing timer if any
    if timer, exists := si.timers[key]; exists {
        timer.Stop()
    }
    
    // Set new timer
    si.timers[key] = time.AfterFunc(5*time.Second, func() {
        si.manager.Invalidate(key)
        fmt.Printf("Time-based invalidation: %s\n", key)
        
        si.mutex.Lock()
        delete(si.timers, key)
        si.mutex.Unlock()
    })
}

// eventBasedInvalidation invalidates based on events
func (si *SmartInvalidation) eventBasedInvalidation(key string) {
    // Wait for specific events before invalidating
    fmt.Printf("Event-based invalidation queued: %s\n", key)
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
    // Create database
    db := NewDatabase()
    
    // Test different consistency levels
    consistencyLevels := []ConsistencyLevel{
        Strong Consistency,
        Eventual Consistency,
        Weak Consistency,
        Session Consistency,
    }
    
    for _, level := range consistencyLevels {
        fmt.Printf("\n=== Testing %s Consistency ===\n", getConsistencyName(level))
        
        cache := NewConsistentCache(db, level)
        
        // Test write
        if err := cache.Set("user:1", "John Doe", 1*time.Hour); err != nil {
            fmt.Printf("Error setting user:1: %v\n", err)
            continue
        }
        
        // Test read
        if value, err := cache.Get("user:1"); err != nil {
            fmt.Printf("Error getting user:1: %v\n", err)
        } else {
            fmt.Printf("Retrieved user:1: %v\n", value)
        }
        
        // Test read again
        if value, err := cache.Get("user:1"); err != nil {
            fmt.Printf("Error getting user:1: %v\n", err)
        } else {
            fmt.Printf("Retrieved user:1 (cached): %v\n", value)
        }
        
        // Print stats
        stats := cache.GetStats()
        fmt.Printf("Cache stats: %+v\n", stats)
        
        // Wait for eventual consistency to sync
        if level == Eventual Consistency {
            time.Sleep(2 * time.Second)
        }
    }
    
    // Test invalidation strategies
    fmt.Println("\n=== Testing Invalidation Strategies ===")
    
    strategies := []InvalidationStrategy{
        Immediate,
        Lazy,
        TimeBased,
        EventBased,
    }
    
    for _, strategy := range strategies {
        fmt.Printf("\n--- %s Invalidation ---\n", getInvalidationName(strategy))
        
        invalidation := NewSmartInvalidation(strategy)
        
        // Test invalidation
        invalidation.Invalidate("user:1")
        invalidation.Invalidate("user:2")
        
        // Wait for time-based invalidation
        if strategy == TimeBased {
            time.Sleep(6 * time.Second)
        }
    }
}

// Helper functions
func getConsistencyName(level ConsistencyLevel) string {
    switch level {
    case Strong Consistency:
        return "Strong"
    case Eventual Consistency:
        return "Eventual"
    case Weak Consistency:
        return "Weak"
    case Session Consistency:
        return "Session"
    default:
        return "Unknown"
    }
}

func getInvalidationName(strategy InvalidationStrategy) string {
    switch strategy {
    case Immediate:
        return "Immediate"
    case Lazy:
        return "Lazy"
    case TimeBased:
        return "Time-Based"
    case EventBased:
        return "Event-Based"
    default:
        return "Unknown"
    }
}
```

## 🧪 Exercise: Consistency Testing

### Task
Create a test to verify consistency guarantees:

```go
// consistency_test.go
package main

import (
    "testing"
    "time"
)

func TestConsistencyGuarantees(t *testing.T) {
    db := NewDatabase()
    
    // Test strong consistency
    cache := NewConsistentCache(db, Strong Consistency)
    
    // Write value
    cache.Set("test", "value1", 1*time.Hour)
    
    // Read should return latest value
    value, err := cache.Get("test")
    if err != nil {
        t.Fatalf("Error getting value: %v", err)
    }
    
    if value != "value1" {
        t.Errorf("Expected 'value1', got '%v'", value)
    }
    
    // Update in database directly
    db.Write("test", "value2")
    
    // Read should return updated value (strong consistency)
    value, err = cache.Get("test")
    if err != nil {
        t.Fatalf("Error getting value: %v", err)
    }
    
    if value != "value2" {
        t.Errorf("Expected 'value2', got '%v'", value)
    }
}
```

## 📝 Key Takeaways

1. **Consistency vs Performance**: Trade-off between data freshness and performance
2. **Choose based on requirements**: Different use cases need different consistency levels
3. **Invalidation is complex**: Smart invalidation can improve performance
4. **Background processes**: Essential for eventual consistency
5. **Monitoring**: Track consistency metrics in production

## 🎯 Practice Questions

1. When would you choose eventual consistency over strong consistency?
2. How do you handle invalidation in a distributed system?
3. What are the trade-offs of different invalidation strategies?
4. How do you measure consistency in production?
5. What happens if background sync fails in eventual consistency?

## 🚀 Next Steps

Tomorrow we'll cover:
- Memory management and optimization
- Cache warming and preloading
- Advanced caching patterns

---

**Your Notes:**
<!-- Add your notes, questions, and insights here -->

**Code Experiments:**
<!-- Document any code experiments or modifications you try -->

**Questions:**
<!-- Write down any questions you have -->
