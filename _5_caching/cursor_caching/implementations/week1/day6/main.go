package main

import (
    "fmt"
    "sync"
    "time"
)

// WarmingStrategy defines different warming strategies
type WarmingStrategy int

const (
    Predictive WarmingStrategy = iota
    Scheduled
    EventDriven
    Proactive
)

// WarmingConfig holds configuration for cache warming
type WarmingConfig struct {
    Strategy      WarmingStrategy
    Interval      time.Duration
    BatchSize     int
    MaxConcurrency int
    Enabled       bool
}

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

// Size returns the number of items in the cache
func (c *SimpleCache) Size() int {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    return len(c.items)
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

// GetProactiveKeys returns keys for proactive warming
func (dp *DataPredictor) GetProactiveKeys() []string {
    dp.mutex.RLock()
    defer dp.mutex.RUnlock()
    
    var keys []string
    for key, score := range dp.predictions {
        if score > 0.5 { // Lower threshold for proactive warming
            keys = append(keys, key)
        }
    }
    
    return keys
}

// updatePredictions updates prediction scores
func (dp *DataPredictor) updatePredictions() {
    // Simple prediction algorithm
    // In real implementation, use machine learning
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

// GetAccessPatterns returns current access patterns
func (dp *DataPredictor) GetAccessPatterns() map[string]int64 {
    dp.mutex.RLock()
    defer dp.mutex.RUnlock()
    
    patterns := make(map[string]int64)
    for key, count := range dp.accessPatterns {
        patterns[key] = count
    }
    
    return patterns
}

// GetPredictions returns current predictions
func (dp *DataPredictor) GetPredictions() map[string]float64 {
    dp.mutex.RLock()
    defer dp.mutex.RUnlock()
    
    predictions := make(map[string]float64)
    for key, score := range dp.predictions {
        predictions[key] = score
    }
    
    return predictions
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

// CacheWarmer implements intelligent cache warming
type CacheWarmer struct {
    cache      *SimpleCache
    database   *Database
    config     WarmingConfig
    predictor  *DataPredictor
    scheduler  *WarmingScheduler
    mutex      sync.RWMutex
    stats      WarmingStats
}

// WarmingStats tracks warming statistics
type WarmingStats struct {
    ItemsWarmed int64
    WarmingTime time.Duration
    SuccessRate float64
    LastWarming time.Time
    Errors      int64
}

// NewCacheWarmer creates a new cache warmer
func NewCacheWarmer(cache *SimpleCache, db *Database, config WarmingConfig) *CacheWarmer {
    warmer := &CacheWarmer{
        cache:     cache,
        database:  db,
        config:    config,
        predictor: NewDataPredictor(),
        scheduler: NewWarmingScheduler(),
        stats:     WarmingStats{},
    }
    
    // Start warming based on strategy
    switch config.Strategy {
    case Predictive:
        go warmer.startPredictiveWarming()
    case Scheduled:
        go warmer.startScheduledWarming()
    case EventDriven:
        go warmer.startEventDrivenWarming()
    case Proactive:
        go warmer.startProactiveWarming()
    }
    
    return warmer
}

// startPredictiveWarming starts predictive warming
func (cw *CacheWarmer) startPredictiveWarming() {
    ticker := time.NewTicker(cw.config.Interval)
    defer ticker.Stop()
    
    for range ticker.C {
        if !cw.config.Enabled {
            continue
        }
        
        cw.warmPredictiveData()
    }
}

// startScheduledWarming starts scheduled warming
func (cw *CacheWarmer) startScheduledWarming() {
    ticker := time.NewTicker(cw.config.Interval)
    defer ticker.Stop()
    
    for range ticker.C {
        if !cw.config.Enabled {
            continue
        }
        
        cw.warmScheduledData()
    }
}

// startEventDrivenWarming starts event-driven warming
func (cw *CacheWarmer) startEventDrivenWarming() {
    // This would typically listen to events
    // For demo purposes, we'll simulate with a ticker
    ticker := time.NewTicker(cw.config.Interval)
    defer ticker.Stop()
    
    for range ticker.C {
        if !cw.config.Enabled {
            continue
        }
        
        cw.warmEventDrivenData()
    }
}

// startProactiveWarming starts proactive warming
func (cw *CacheWarmer) startProactiveWarming() {
    ticker := time.NewTicker(cw.config.Interval)
    defer ticker.Stop()
    
    for range ticker.C {
        if !cw.config.Enabled {
            continue
        }
        
        cw.warmProactiveData()
    }
}

// warmPredictiveData warms cache with predicted data
func (cw *CacheWarmer) warmPredictiveData() {
    start := time.Now()
    
    // Get predicted keys
    predictedKeys := cw.predictor.PredictKeys()
    
    // Warm cache with predicted data
    successCount := 0
    for _, key := range predictedKeys {
        if err := cw.warmKey(key); err != nil {
            cw.stats.Errors++
            continue
        }
        successCount++
    }
    
    // Update statistics
    cw.mutex.Lock()
    cw.stats.ItemsWarmed += int64(successCount)
    cw.stats.WarmingTime = time.Since(start)
    cw.stats.LastWarming = time.Now()
    cw.stats.SuccessRate = float64(successCount) / float64(len(predictedKeys)) * 100
    cw.mutex.Unlock()
    
    fmt.Printf("   [Predictive] Warming completed: %d items in %v\n",
        successCount, time.Since(start))
}

// warmScheduledData warms cache with scheduled data
func (cw *CacheWarmer) warmScheduledData() {
    start := time.Now()
    
    // Get scheduled keys
    scheduledKeys := cw.scheduler.GetScheduledKeys()
    
    // Warm cache with scheduled data
    successCount := 0
    for _, key := range scheduledKeys {
        if err := cw.warmKey(key); err != nil {
            cw.stats.Errors++
            continue
        }
        successCount++
    }
    
    // Update statistics
    cw.mutex.Lock()
    cw.stats.ItemsWarmed += int64(successCount)
    cw.stats.WarmingTime = time.Since(start)
    cw.stats.LastWarming = time.Now()
    cw.stats.SuccessRate = float64(successCount) / float64(len(scheduledKeys)) * 100
    cw.mutex.Unlock()
    
    fmt.Printf("   [Scheduled] Warming completed: %d items in %v\n",
        successCount, time.Since(start))
}

// warmEventDrivenData warms cache with event-driven data
func (cw *CacheWarmer) warmEventDrivenData() {
    start := time.Now()
    
    // Simulate event-driven warming
    // In real implementation, this would be triggered by events
    eventKeys := []string{"event:user:login", "event:product:view", "event:cart:add"}
    
    successCount := 0
    for _, key := range eventKeys {
        if err := cw.warmKey(key); err != nil {
            cw.stats.Errors++
            continue
        }
        successCount++
    }
    
    // Update statistics
    cw.mutex.Lock()
    cw.stats.ItemsWarmed += int64(successCount)
    cw.stats.WarmingTime = time.Since(start)
    cw.stats.LastWarming = time.Now()
    cw.stats.SuccessRate = float64(successCount) / float64(len(eventKeys)) * 100
    cw.mutex.Unlock()
    
    fmt.Printf("   [Event-Driven] Warming completed: %d items in %v\n",
        successCount, time.Since(start))
}

// warmProactiveData warms cache with proactive data
func (cw *CacheWarmer) warmProactiveData() {
    start := time.Now()
    
    // Get proactive keys based on user behavior
    proactiveKeys := cw.predictor.GetProactiveKeys()
    
    successCount := 0
    for _, key := range proactiveKeys {
        if err := cw.warmKey(key); err != nil {
            cw.stats.Errors++
            continue
        }
        successCount++
    }
    
    // Update statistics
    cw.mutex.Lock()
    cw.stats.ItemsWarmed += int64(successCount)
    cw.stats.WarmingTime = time.Since(start)
    cw.stats.LastWarming = time.Now()
    cw.stats.SuccessRate = float64(successCount) / float64(len(proactiveKeys)) * 100
    cw.mutex.Unlock()
    
    fmt.Printf("   [Proactive] Warming completed: %d items in %v\n",
        successCount, time.Since(start))
}

// warmKey warms a specific key
func (cw *CacheWarmer) warmKey(key string) error {
    // Check if key already exists in cache
    if _, found := cw.cache.Get(key); found {
        return nil // Already in cache
    }
    
    // Read from database
    value, err := cw.database.Read(key)
    if err != nil {
        return err
    }
    
    // Store in cache
    cw.cache.Set(key, value)
    return nil
}

// GetStats returns warming statistics
func (cw *CacheWarmer) GetStats() WarmingStats {
    cw.mutex.RLock()
    defer cw.mutex.RUnlock()
    return cw.stats
}

// SetEnabled enables or disables warming
func (cw *CacheWarmer) SetEnabled(enabled bool) {
    cw.mutex.Lock()
    defer cw.mutex.Unlock()
    cw.config.Enabled = enabled
}

// WarmNow triggers immediate warming
func (cw *CacheWarmer) WarmNow() {
    switch cw.config.Strategy {
    case Predictive:
        cw.warmPredictiveData()
    case Scheduled:
        cw.warmScheduledData()
    case EventDriven:
        cw.warmEventDrivenData()
    case Proactive:
        cw.warmProactiveData()
    }
}

func main() {
    fmt.Println("=== Day 6: Cache Warming & Preloading ===")
    fmt.Println("Implementing intelligent cache warming strategies\n")
    
    // Create database and cache
    db := NewDatabase()
    cache := NewSimpleCache()
    
    // Pre-populate database with test data
    fmt.Println("1. Setting up test data...")
    db.Write("user:1", "John Doe")
    db.Write("user:2", "Jane Smith")
    db.Write("user:3", "Bob Johnson")
    db.Write("product:1", "Laptop")
    db.Write("product:2", "Phone")
    db.Write("product:3", "Tablet")
    db.Write("config:theme", "dark")
    db.Write("config:language", "en")
    db.Write("event:user:login", "login_event")
    db.Write("event:product:view", "view_event")
    db.Write("event:cart:add", "add_event")
    
    // Example 1: Predictive warming
    fmt.Println("\n2. Predictive Warming Strategy:")
    fmt.Println("   Testing predictive warming based on access patterns...")
    
    config := WarmingConfig{
        Strategy:      Predictive,
        Interval:      3 * time.Second,
        BatchSize:     10,
        MaxConcurrency: 5,
        Enabled:       true,
    }
    
    warmer := NewCacheWarmer(cache, db, config)
    
    // Simulate some access patterns
    predictor := warmer.predictor
    predictor.RecordAccess("user:1")
    predictor.RecordAccess("user:1")
    predictor.RecordAccess("user:2")
    predictor.RecordAccess("product:1")
    predictor.RecordAccess("product:1")
    predictor.RecordAccess("product:1")
    
    fmt.Println("   Access patterns recorded, waiting for warming...")
    time.Sleep(5 * time.Second)
    
    // Print statistics
    stats := warmer.GetStats()
    fmt.Printf("   Predictive Warming Stats:\n")
    fmt.Printf("   - Items warmed: %d\n", stats.ItemsWarmed)
    fmt.Printf("   - Success rate: %.2f%%\n", stats.SuccessRate)
    fmt.Printf("   - Errors: %d\n", stats.Errors)
    
    // Disable warming
    warmer.SetEnabled(false)
    
    // Example 2: Scheduled warming
    fmt.Println("\n3. Scheduled Warming Strategy:")
    fmt.Println("   Testing scheduled warming...")
    
    config.Strategy = Scheduled
    config.Interval = 2 * time.Second
    warmer = NewCacheWarmer(cache, db, config)
    
    // Add some scheduled keys
    scheduler := warmer.scheduler
    scheduler.AddSchedule("user:1", 5*time.Second)
    scheduler.AddSchedule("product:1", 3*time.Second)
    scheduler.AddSchedule("config:theme", 10*time.Second)
    
    fmt.Println("   Scheduled keys added, waiting for warming...")
    time.Sleep(8 * time.Second)
    
    // Print statistics
    stats = warmer.GetStats()
    fmt.Printf("   Scheduled Warming Stats:\n")
    fmt.Printf("   - Items warmed: %d\n", stats.ItemsWarmed)
    fmt.Printf("   - Success rate: %.2f%%\n", stats.SuccessRate)
    fmt.Printf("   - Errors: %d\n", stats.Errors)
    
    // Disable warming
    warmer.SetEnabled(false)
    
    // Example 3: Event-driven warming
    fmt.Println("\n4. Event-Driven Warming Strategy:")
    fmt.Println("   Testing event-driven warming...")
    
    config.Strategy = EventDriven
    config.Interval = 2 * time.Second
    warmer = NewCacheWarmer(cache, db, config)
    
    fmt.Println("   Event-driven warming started, waiting for events...")
    time.Sleep(6 * time.Second)
    
    // Print statistics
    stats = warmer.GetStats()
    fmt.Printf("   Event-Driven Warming Stats:\n")
    fmt.Printf("   - Items warmed: %d\n", stats.ItemsWarmed)
    fmt.Printf("   - Success rate: %.2f%%\n", stats.SuccessRate)
    fmt.Printf("   - Errors: %d\n", stats.Errors)
    
    // Disable warming
    warmer.SetEnabled(false)
    
    // Example 4: Proactive warming
    fmt.Println("\n5. Proactive Warming Strategy:")
    fmt.Println("   Testing proactive warming...")
    
    config.Strategy = Proactive
    config.Interval = 2 * time.Second
    warmer = NewCacheWarmer(cache, db, config)
    
    // Simulate proactive access patterns
    predictor = warmer.predictor
    predictor.RecordAccess("user:2")
    predictor.RecordAccess("user:2")
    predictor.RecordAccess("product:2")
    predictor.RecordAccess("product:2")
    predictor.RecordAccess("config:language")
    
    fmt.Println("   Proactive patterns recorded, waiting for warming...")
    time.Sleep(5 * time.Second)
    
    // Print statistics
    stats = warmer.GetStats()
    fmt.Printf("   Proactive Warming Stats:\n")
    fmt.Printf("   - Items warmed: %d\n", stats.ItemsWarmed)
    fmt.Printf("   - Success rate: %.2f%%\n", stats.SuccessRate)
    fmt.Printf("   - Errors: %d\n", stats.Errors)
    
    // Disable warming
    warmer.SetEnabled(false)
    
    // Example 5: Manual warming
    fmt.Println("\n6. Manual Warming:")
    fmt.Println("   Testing manual warming trigger...")
    
    config.Enabled = false
    warmer = NewCacheWarmer(cache, db, config)
    
    // Trigger manual warming
    fmt.Println("   Triggering manual warming...")
    warmer.WarmNow()
    
    // Print statistics
    stats = warmer.GetStats()
    fmt.Printf("   Manual Warming Stats:\n")
    fmt.Printf("   - Items warmed: %d\n", stats.ItemsWarmed)
    fmt.Printf("   - Success rate: %.2f%%\n", stats.SuccessRate)
    fmt.Printf("   - Errors: %d\n", stats.Errors)
    
    // Example 6: Warming effectiveness analysis
    fmt.Println("\n7. Warming Effectiveness Analysis:")
    fmt.Println("   Analyzing warming effectiveness...")
    
    // Test different warming intervals
    intervals := []time.Duration{
        1 * time.Second,
        2 * time.Second,
        5 * time.Second,
        10 * time.Second,
    }
    
    for _, interval := range intervals {
        fmt.Printf("\n   Testing interval: %v\n", interval)
        
        config.Interval = interval
        config.Enabled = true
        warmer = NewCacheWarmer(cache, db, config)
        
        // Wait for warming
        time.Sleep(interval + 2*time.Second)
        
        stats := warmer.GetStats()
        fmt.Printf("   - Items warmed: %d\n", stats.ItemsWarmed)
        fmt.Printf("   - Success rate: %.2f%%\n", stats.SuccessRate)
        fmt.Printf("   - Warming time: %v\n", stats.WarmingTime)
        
        // Disable warming
        warmer.SetEnabled(false)
    }
    
    // Example 7: Cache state after warming
    fmt.Println("\n8. Cache State After Warming:")
    fmt.Println("   Checking cache contents...")
    
    fmt.Printf("   Cache size: %d items\n", cache.Size())
    
    // Check if warmed items are in cache
    testKeys := []string{"user:1", "user:2", "product:1", "product:2", "config:theme"}
    for _, key := range testKeys {
        if value, found := cache.Get(key); found {
            fmt.Printf("   ✓ %s: %v\n", key, value)
        } else {
            fmt.Printf("   ✗ %s: not found\n", key)
        }
    }
    
    // Example 8: Performance impact
    fmt.Println("\n9. Performance Impact Analysis:")
    fmt.Println("   Measuring warming performance impact...")
    
    // Test warming performance
    config.Enabled = false
    warmer = NewCacheWarmer(cache, db, config)
    
    start := time.Now()
    for i := 0; i < 100; i++ {
        warmer.WarmNow()
    }
    warmingTime := time.Since(start)
    
    fmt.Printf("   Warming 100 items took: %v\n", warmingTime)
    fmt.Printf("   Average warming time: %v\n", warmingTime/100)
    
    // Test cache access performance
    start = time.Now()
    for i := 0; i < 1000; i++ {
        key := fmt.Sprintf("perf:key:%d", i%10)
        cache.Get(key)
    }
    accessTime := time.Since(start)
    
    fmt.Printf("   Accessing 1000 items took: %v\n", accessTime)
    fmt.Printf("   Average access time: %v\n", accessTime/1000)
    
    fmt.Println("\n=== Day 6 Complete! ===")
    fmt.Println("You've successfully implemented:")
    fmt.Println("✓ Predictive warming")
    fmt.Println("✓ Scheduled warming")
    fmt.Println("✓ Event-driven warming")
    fmt.Println("✓ Proactive warming")
    fmt.Println("✓ Manual warming triggers")
    fmt.Println("✓ Warming effectiveness analysis")
    fmt.Println("✓ Performance impact measurement")
    fmt.Println("\nReady for Day 7: Week 1 Review! 🚀")
}
