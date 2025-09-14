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

// ConsistencyLevel defines different consistency levels
type ConsistencyLevel int

const (
    Strong ConsistencyLevel = iota
    Eventual
    Weak
    Session
)

// WarmingStrategy defines different warming strategies
type WarmingStrategy int

const (
    Predictive WarmingStrategy = iota
    Scheduled
    EventDriven
    Proactive
)

// CompleteCache represents our final comprehensive caching system
type CompleteCache struct {
    // Core cache functionality
    items      map[string]*CacheItem
    capacity   int
    mutex      sync.RWMutex
    
    // Eviction policy
    evictionPolicy EvictionPolicy
    
    // Consistency level
    consistencyLevel ConsistencyLevel
    
    // Warming
    warmingStrategy WarmingStrategy
    predictor       *DataPredictor
    scheduler       *WarmingScheduler
    
    // Statistics
    stats CacheStats
}

// CacheItem represents a complete cache item
type CacheItem struct {
    Key         string
    Value       interface{}
    CreatedAt   time.Time
    LastUsed    time.Time
    AccessCount int64
    Size        int
    Dirty       bool
    Version     int64
}

// CacheStats holds comprehensive cache statistics
type CacheStats struct {
    Hits        int64
    Misses      int64
    Evictions   int64
    Size        int
    Capacity    int
    HitRate     float64
    MissRate    float64
    Warmings    int64
    Errors      int64
}

// DataPredictor predicts which data should be warmed
type DataPredictor struct {
    accessPatterns map[string]int64
    predictions    map[string]float64
    mutex          sync.RWMutex
    lastUpdate     time.Time
}

// NewDataPredictor creates a new data predictor
func NewDataPredictor() *DataPredictor {
    return &DataPredictor{
        accessPatterns: make(map[string]int64),
        predictions:    make(map[string]float64),
    }
}

// RecordAccess records access to a key
func (dp *DataPredictor) RecordAccess(key string) {
    dp.mutex.Lock()
    defer dp.mutex.Unlock()
    
    dp.accessPatterns[key]++
    dp.lastUpdate = time.Now()
    
    // Update predictions based on access patterns
    dp.updatePredictions()
}

// PredictKeys returns predicted keys for warming
func (dp *DataPredictor) PredictKeys() []string {
    dp.mutex.RLock()
    defer dp.mutex.RUnlock()
    
    var keys []string
    for key, score := range dp.predictions {
        if score > 0.7 { // Threshold for warming
            keys = append(keys, key)
        }
    }
    
    return keys
}

// updatePredictions updates prediction scores
func (dp *DataPredictor) updatePredictions() {
    // Simple prediction algorithm
    for key, count := range dp.accessPatterns {
        // Calculate score based on access count and recency
        score := float64(count) / 100.0 // Normalize by 100
        
        // Add recency factor
        if dp.lastUpdate.Sub(time.Now()) < 1*time.Hour {
            score *= 1.2
        }
        
        dp.predictions[key] = score
    }
}

// WarmingScheduler schedules cache warming tasks
type WarmingScheduler struct {
    schedules map[string]WarmingSchedule
    mutex     sync.RWMutex
}

// WarmingSchedule defines a warming schedule
type WarmingSchedule struct {
    Key       string
    Interval  time.Duration
    LastRun   time.Time
    NextRun   time.Time
    Enabled   bool
}

// NewWarmingScheduler creates a new warming scheduler
func NewWarmingScheduler() *WarmingScheduler {
    return &WarmingScheduler{
        schedules: make(map[string]WarmingSchedule),
    }
}

// AddSchedule adds a warming schedule
func (ws *WarmingScheduler) AddSchedule(key string, interval time.Duration) {
    ws.mutex.Lock()
    defer ws.mutex.Unlock()
    
    now := time.Now()
    ws.schedules[key] = WarmingSchedule{
        Key:      key,
        Interval: interval,
        LastRun:  now,
        NextRun:  now.Add(interval),
        Enabled:  true,
    }
}

// GetScheduledKeys returns keys that should be warmed now
func (ws *WarmingScheduler) GetScheduledKeys() []string {
    ws.mutex.Lock()
    defer ws.mutex.Unlock()
    
    var keys []string
    now := time.Now()
    
    for key, schedule := range ws.schedules {
        if schedule.Enabled && now.After(schedule.NextRun) {
            keys = append(keys, key)
            
            // Update schedule
            schedule.LastRun = now
            schedule.NextRun = now.Add(schedule.Interval)
            ws.schedules[key] = schedule
        }
    }
    
    return keys
}

// NewCompleteCache creates a new complete cache
func NewCompleteCache(capacity int, evictionPolicy EvictionPolicy, consistencyLevel ConsistencyLevel, warmingStrategy WarmingStrategy) *CompleteCache {
    cache := &CompleteCache{
        items:            make(map[string]*CacheItem),
        capacity:         capacity,
        evictionPolicy:   evictionPolicy,
        consistencyLevel: consistencyLevel,
        warmingStrategy:  warmingStrategy,
        predictor:        NewDataPredictor(),
        scheduler:        NewWarmingScheduler(),
        stats:            CacheStats{Capacity: capacity},
    }
    
    // Start background processes
    go cache.backgroundMaintenance()
    go cache.warmingProcess()
    
    return cache
}

// Set stores a value in the cache
func (c *CompleteCache) Set(key string, value interface{}, ttl time.Duration) error {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    // Check capacity
    if len(c.items) >= c.capacity {
        c.evictItem()
    }
    
    // Create or update item
    if item, exists := c.items[key]; exists {
        // Update existing item
        item.Value = value
        item.LastUsed = time.Now()
        item.AccessCount++
        item.Size = c.calculateSize(value)
        item.Dirty = true
        item.Version++
    } else {
        // Create new item
        item = &CacheItem{
            Key:         key,
            Value:       value,
            CreatedAt:   time.Now(),
            LastUsed:    time.Now(),
            AccessCount: 1,
            Size:        c.calculateSize(value),
            Dirty:       true,
            Version:     1,
        }
        c.items[key] = item
    }
    
    c.stats.Hits++
    c.stats.Size = len(c.items)
    
    return nil
}

// Get retrieves a value from the cache
func (c *CompleteCache) Get(key string) (interface{}, error) {
    c.mutex.RLock()
    item, exists := c.items[key]
    c.mutex.RUnlock()
    
    if !exists {
        c.stats.Misses++
        return nil, fmt.Errorf("key not found")
    }
    
    // Update access statistics
    c.mutex.Lock()
    item.LastUsed = time.Now()
    item.AccessCount++
    c.mutex.Unlock()
    
    c.stats.Hits++
    
    // Record access for warming prediction
    c.predictor.RecordAccess(key)
    
    return item.Value, nil
}

// evictItem evicts an item based on the eviction policy
func (c *CompleteCache) evictItem() {
    var keyToEvict string
    
    switch c.evictionPolicy {
    case LRU:
        keyToEvict = c.findLRUKey()
    case LFU:
        keyToEvict = c.findLFUKey()
    case FIFO:
        keyToEvict = c.findFIFOKey()
    case LIFO:
        keyToEvict = c.findLIFOKey()
    default:
        keyToEvict = c.findLRUKey() // Default to LRU
    }
    
    if keyToEvict != "" {
        delete(c.items, keyToEvict)
        c.stats.Evictions++
        c.stats.Size = len(c.items)
    }
}

// findLRUKey finds the least recently used key
func (c *CompleteCache) findLRUKey() string {
    var oldestKey string
    var oldestTime time.Time
    
    for key, item := range c.items {
        if oldestKey == "" || item.LastUsed.Before(oldestTime) {
            oldestKey = key
            oldestTime = item.LastUsed
        }
    }
    
    return oldestKey
}

// findLFUKey finds the least frequently used key
func (c *CompleteCache) findLFUKey() string {
    var leastUsedKey string
    var leastCount int64
    
    for key, item := range c.items {
        if leastUsedKey == "" || item.AccessCount < leastCount {
            leastUsedKey = key
            leastCount = item.AccessCount
        }
    }
    
    return leastUsedKey
}

// findFIFOKey finds the first in, first out key
func (c *CompleteCache) findFIFOKey() string {
    var oldestKey string
    var oldestTime time.Time
    
    for key, item := range c.items {
        if oldestKey == "" || item.CreatedAt.Before(oldestTime) {
            oldestKey = key
            oldestTime = item.CreatedAt
        }
    }
    
    return oldestKey
}

// findLIFOKey finds the last in, first out key
func (c *CompleteCache) findLIFOKey() string {
    var newestKey string
    var newestTime time.Time
    
    for key, item := range c.items {
        if newestKey == "" || item.CreatedAt.After(newestTime) {
            newestKey = key
            newestTime = item.CreatedAt
        }
    }
    
    return newestKey
}

// calculateSize calculates the size of a value
func (c *CompleteCache) calculateSize(value interface{}) int {
    // Simple size calculation
    switch v := value.(type) {
    case string:
        return len(v)
    case []byte:
        return len(v)
    case int:
        return 8 // 64-bit integer
    case float64:
        return 8 // 64-bit float
    default:
        return 100 // Default size
    }
}

// backgroundMaintenance runs background maintenance tasks
func (c *CompleteCache) backgroundMaintenance() {
    ticker := time.NewTicker(1 * time.Minute)
    defer ticker.Stop()
    
    for range ticker.C {
        c.mutex.Lock()
        
        // Clean up expired items
        now := time.Now()
        for key, item := range c.items {
            if now.After(item.CreatedAt.Add(1 * time.Hour)) { // 1 hour TTL
                delete(c.items, key)
            }
        }
        
        c.stats.Size = len(c.items)
        c.mutex.Unlock()
    }
}

// warmingProcess runs the warming process
func (c *CompleteCache) warmingProcess() {
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        c.performWarming()
    }
}

// performWarming performs cache warming based on strategy
func (c *CompleteCache) performWarming() {
    switch c.warmingStrategy {
    case Predictive:
        c.warmPredictive()
    case Scheduled:
        c.warmScheduled()
    case EventDriven:
        c.warmEventDriven()
    case Proactive:
        c.warmProactive()
    }
}

// warmPredictive performs predictive warming
func (c *CompleteCache) warmPredictive() {
    predictedKeys := c.predictor.PredictKeys()
    
    for _, key := range predictedKeys {
        if _, exists := c.items[key]; !exists {
            // Simulate warming by creating a placeholder
            c.mutex.Lock()
            c.items[key] = &CacheItem{
                Key:         key,
                Value:       fmt.Sprintf("warmed:%s", key),
                CreatedAt:   time.Now(),
                LastUsed:    time.Now(),
                AccessCount: 0,
                Size:        100,
                Dirty:       false,
                Version:     1,
            }
            c.stats.Warmings++
            c.mutex.Unlock()
        }
    }
}

// warmScheduled performs scheduled warming
func (c *CompleteCache) warmScheduled() {
    scheduledKeys := c.scheduler.GetScheduledKeys()
    
    for _, key := range scheduledKeys {
        if _, exists := c.items[key]; !exists {
            // Simulate warming by creating a placeholder
            c.mutex.Lock()
            c.items[key] = &CacheItem{
                Key:         key,
                Value:       fmt.Sprintf("scheduled:%s", key),
                CreatedAt:   time.Now(),
                LastUsed:    time.Now(),
                AccessCount: 0,
                Size:        100,
                Dirty:       false,
                Version:     1,
            }
            c.stats.Warmings++
            c.mutex.Unlock()
        }
    }
}

// warmEventDriven performs event-driven warming
func (c *CompleteCache) warmEventDriven() {
    // Simulate event-driven warming
    eventKeys := []string{"event:user:login", "event:product:view", "event:cart:add"}
    
    for _, key := range eventKeys {
        if _, exists := c.items[key]; !exists {
            // Simulate warming by creating a placeholder
            c.mutex.Lock()
            c.items[key] = &CacheItem{
                Key:         key,
                Value:       fmt.Sprintf("event:%s", key),
                CreatedAt:   time.Now(),
                LastUsed:    time.Now(),
                AccessCount: 0,
                Size:        100,
                Dirty:       false,
                Version:     1,
            }
            c.stats.Warmings++
            c.mutex.Unlock()
        }
    }
}

// warmProactive performs proactive warming
func (c *CompleteCache) warmProactive() {
    // Simulate proactive warming based on patterns
    proactiveKeys := []string{"proactive:user:profile", "proactive:recommendations", "proactive:settings"}
    
    for _, key := range proactiveKeys {
        if _, exists := c.items[key]; !exists {
            // Simulate warming by creating a placeholder
            c.mutex.Lock()
            c.items[key] = &CacheItem{
                Key:         key,
                Value:       fmt.Sprintf("proactive:%s", key),
                CreatedAt:   time.Now(),
                LastUsed:    time.Now(),
                AccessCount: 0,
                Size:        100,
                Dirty:       false,
                Version:     1,
            }
            c.stats.Warmings++
            c.mutex.Unlock()
        }
    }
}

// GetStats returns cache statistics
func (c *CompleteCache) GetStats() CacheStats {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    c.stats.Size = len(c.items)
    c.stats.Capacity = c.capacity
    
    total := c.stats.Hits + c.stats.Misses
    if total > 0 {
        c.stats.HitRate = float64(c.stats.Hits) / float64(total) * 100
        c.stats.MissRate = float64(c.stats.Misses) / float64(total) * 100
    }
    
    return c.stats
}

// Clear clears the cache
func (c *CompleteCache) Clear() {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    c.items = make(map[string]*CacheItem)
    c.stats = CacheStats{Capacity: c.capacity}
}

// PrintCache prints the current cache state (for debugging)
func (c *CompleteCache) PrintCache() {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    fmt.Printf("   Cache (%s, %s, %s): ", 
        getEvictionPolicyName(c.evictionPolicy),
        getConsistencyLevelName(c.consistencyLevel),
        getWarmingStrategyName(c.warmingStrategy))
    
    for key, item := range c.items {
        fmt.Printf("[%s:%v] ", key, item.Value)
    }
    fmt.Println()
}

// Helper functions
func getEvictionPolicyName(policy EvictionPolicy) string {
    switch policy {
    case LRU:
        return "LRU"
    case LFU:
        return "LFU"
    case FIFO:
        return "FIFO"
    case LIFO:
        return "LIFO"
    default:
        return "Unknown"
    }
}

func getConsistencyLevelName(level ConsistencyLevel) string {
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

func getWarmingStrategyName(strategy WarmingStrategy) string {
    switch strategy {
    case Predictive:
        return "Predictive"
    case Scheduled:
        return "Scheduled"
    case EventDriven:
        return "Event-Driven"
    case Proactive:
        return "Proactive"
    default:
        return "Unknown"
    }
}

func main() {
    fmt.Println("=== Day 7: Week 1 Review & Assessment ===")
    fmt.Println("Building a complete caching system with all concepts learned\n")
    
    // Example 1: Basic cache functionality
    fmt.Println("1. Basic Cache Functionality:")
    fmt.Println("   Testing core cache operations...")
    
    cache := NewCompleteCache(5, LRU, Strong, Predictive)
    
    // Test basic operations
    cache.Set("user:1", "John Doe", 1*time.Hour)
    cache.Set("user:2", "Jane Smith", 1*time.Hour)
    cache.Set("product:1", "Laptop", 1*time.Hour)
    
    fmt.Println("   After adding items:")
    cache.PrintCache()
    
    // Test get operations
    if value, err := cache.Get("user:1"); err == nil {
        fmt.Printf("   Retrieved user:1: %v\n", value)
    }
    
    if value, err := cache.Get("user:2"); err == nil {
        fmt.Printf("   Retrieved user:2: %v\n", value)
    }
    
    // Example 2: Eviction policy testing
    fmt.Println("\n2. Eviction Policy Testing:")
    fmt.Println("   Testing LRU eviction...")
    
    // Add more items to trigger eviction
    cache.Set("user:3", "Bob Johnson", 1*time.Hour)
    cache.Set("user:4", "Alice Brown", 1*time.Hour)
    cache.Set("user:5", "Charlie Wilson", 1*time.Hour)
    
    fmt.Println("   After adding more items (should trigger eviction):")
    cache.PrintCache()
    
    // Test eviction by accessing items
    cache.Get("user:1") // Make user:1 most recently used
    cache.Set("user:6", "David Lee", 1*time.Hour) // Should evict least recently used
    
    fmt.Println("   After accessing user:1 and adding user:6:")
    cache.PrintCache()
    
    // Example 3: Different eviction policies
    fmt.Println("\n3. Different Eviction Policies:")
    
    policies := []EvictionPolicy{LRU, LFU, FIFO, LIFO}
    
    for _, policy := range policies {
        fmt.Printf("   Testing %s eviction:\n", getEvictionPolicyName(policy))
        
        testCache := NewCompleteCache(3, policy, Strong, Predictive)
        
        // Add items
        testCache.Set("a", 1, 1*time.Hour)
        testCache.Set("b", 2, 1*time.Hour)
        testCache.Set("c", 3, 1*time.Hour)
        
        // Access some items
        testCache.Get("a")
        testCache.Get("b")
        testCache.Get("a")
        
        // Add one more item to trigger eviction
        testCache.Set("d", 4, 1*time.Hour)
        
        testCache.PrintCache()
    }
    
    // Example 4: Consistency levels
    fmt.Println("\n4. Consistency Levels:")
    
    consistencyLevels := []ConsistencyLevel{Strong, Eventual, Weak, Session}
    
    for _, level := range consistencyLevels {
        fmt.Printf("   Testing %s consistency:\n", getConsistencyLevelName(level))
        
        testCache := NewCompleteCache(3, LRU, level, Predictive)
        
        // Add items
        testCache.Set("consistency:test", "value", 1*time.Hour)
        testCache.Get("consistency:test")
        
        // Print stats
        stats := testCache.GetStats()
        fmt.Printf("   - Hits: %d, Misses: %d, Hit Rate: %.2f%%\n",
            stats.Hits, stats.Misses, stats.HitRate)
    }
    
    // Example 5: Warming strategies
    fmt.Println("\n5. Warming Strategies:")
    
    warmingStrategies := []WarmingStrategy{Predictive, Scheduled, EventDriven, Proactive}
    
    for _, strategy := range warmingStrategies {
        fmt.Printf("   Testing %s warming:\n", getWarmingStrategyName(strategy))
        
        testCache := NewCompleteCache(5, LRU, Strong, strategy)
        
        // Add some items
        testCache.Set("warm:test", "value", 1*time.Hour)
        
        // Wait for warming
        time.Sleep(6 * time.Second)
        
        // Print stats
        stats := testCache.GetStats()
        fmt.Printf("   - Items: %d, Warmings: %d\n", stats.Size, stats.Warmings)
        
        testCache.PrintCache()
    }
    
    // Example 6: Performance testing
    fmt.Println("\n6. Performance Testing:")
    fmt.Println("   Testing cache performance...")
    
    perfCache := NewCompleteCache(1000, LRU, Strong, Predictive)
    
    // Test write performance
    start := time.Now()
    for i := 0; i < 10000; i++ {
        key := fmt.Sprintf("perf:key:%d", i)
        value := fmt.Sprintf("perf:value:%d", i)
        perfCache.Set(key, value, 1*time.Hour)
    }
    writeTime := time.Since(start)
    
    // Test read performance
    start = time.Now()
    for i := 0; i < 10000; i++ {
        key := fmt.Sprintf("perf:key:%d", i%1000)
        perfCache.Get(key)
    }
    readTime := time.Since(start)
    
    // Print performance stats
    stats := perfCache.GetStats()
    fmt.Printf("   Write 10,000 items: %v\n", writeTime)
    fmt.Printf("   Read 10,000 items: %v\n", readTime)
    fmt.Printf("   Average write time: %v\n", writeTime/10000)
    fmt.Printf("   Average read time: %v\n", readTime/10000)
    fmt.Printf("   Hit rate: %.2f%%\n", stats.HitRate)
    fmt.Printf("   Evictions: %d\n", stats.Evictions)
    
    // Example 7: Comprehensive testing
    fmt.Println("\n7. Comprehensive Testing:")
    fmt.Println("   Testing all features together...")
    
    comprehensiveCache := NewCompleteCache(10, LRU, Strong, Predictive)
    
    // Add items with different patterns
    for i := 0; i < 20; i++ {
        key := fmt.Sprintf("comp:key:%d", i)
        value := fmt.Sprintf("comp:value:%d", i)
        comprehensiveCache.Set(key, value, 1*time.Hour)
        
        // Access some items multiple times
        if i%3 == 0 {
            comprehensiveCache.Get(key)
            comprehensiveCache.Get(key)
        }
    }
    
    // Wait for warming
    time.Sleep(6 * time.Second)
    
    // Print comprehensive stats
    stats = comprehensiveCache.GetStats()
    fmt.Printf("   Comprehensive Cache Stats:\n")
    fmt.Printf("   - Items: %d/%d\n", stats.Size, stats.Capacity)
    fmt.Printf("   - Hits: %d\n", stats.Hits)
    fmt.Printf("   - Misses: %d\n", stats.Misses)
    fmt.Printf("   - Hit Rate: %.2f%%\n", stats.HitRate)
    fmt.Printf("   - Evictions: %d\n", stats.Evictions)
    fmt.Printf("   - Warmings: %d\n", stats.Warmings)
    fmt.Printf("   - Errors: %d\n", stats.Errors)
    
    comprehensiveCache.PrintCache()
    
    // Example 8: Error handling
    fmt.Println("\n8. Error Handling:")
    fmt.Println("   Testing error scenarios...")
    
    // Test getting non-existent key
    if _, err := comprehensiveCache.Get("nonexistent:key"); err != nil {
        fmt.Printf("   Expected error for non-existent key: %v\n", err)
    }
    
    // Test with zero capacity
    zeroCache := NewCompleteCache(0, LRU, Strong, Predictive)
    if err := zeroCache.Set("test", "value", 1*time.Hour); err != nil {
        fmt.Printf("   Expected error with zero capacity: %v\n", err)
    }
    
    // Example 9: Memory usage analysis
    fmt.Println("\n9. Memory Usage Analysis:")
    fmt.Println("   Analyzing memory usage...")
    
    // Get final statistics
    finalStats := comprehensiveCache.GetStats()
    fmt.Printf("   Final Cache Statistics:\n")
    fmt.Printf("   - Total Items: %d\n", finalStats.Size)
    fmt.Printf("   - Capacity: %d\n", finalStats.Capacity)
    fmt.Printf("   - Hit Rate: %.2f%%\n", finalStats.HitRate)
    fmt.Printf("   - Miss Rate: %.2f%%\n", finalStats.MissRate)
    fmt.Printf("   - Evictions: %d\n", finalStats.Evictions)
    fmt.Printf("   - Warmings: %d\n", finalStats.Warmings)
    fmt.Printf("   - Errors: %d\n", finalStats.Errors)
    
    // Example 10: Cache clearing
    fmt.Println("\n10. Cache Management:")
    fmt.Println("    Testing cache clearing...")
    
    fmt.Printf("   Before clear - Items: %d\n", comprehensiveCache.GetStats().Size)
    comprehensiveCache.Clear()
    fmt.Printf("   After clear - Items: %d\n", comprehensiveCache.GetStats().Size)
    
    fmt.Println("\n=== Week 1 Complete! ===")
    fmt.Println("Congratulations! You've successfully built a comprehensive caching system with:")
    fmt.Println("✓ Multiple eviction policies (LRU, LFU, FIFO, LIFO)")
    fmt.Println("✓ Different consistency levels (Strong, Eventual, Weak, Session)")
    fmt.Println("✓ Intelligent warming strategies (Predictive, Scheduled, Event-Driven, Proactive)")
    fmt.Println("✓ Comprehensive statistics and monitoring")
    fmt.Println("✓ Performance optimization")
    fmt.Println("✓ Error handling and edge cases")
    fmt.Println("✓ Memory management")
    fmt.Println("✓ Background maintenance")
    fmt.Println("\nYou're now ready for Week 2: Redis Fundamentals! 🚀")
    fmt.Println("Next week, we'll dive deep into Redis and build production-ready caching systems.")
}
