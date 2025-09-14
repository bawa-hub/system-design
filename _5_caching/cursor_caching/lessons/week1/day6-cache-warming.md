# Day 6: Cache Warming & Preloading

## 🎯 Learning Objectives
- Understand cache warming strategies
- Learn preloading techniques
- Implement intelligent cache warming
- Build a cache with automatic warming

## 📚 Theory: Cache Warming & Preloading

### What is Cache Warming?
Cache warming is the process of pre-populating a cache with frequently accessed data before it's actually requested by users.

### Why Cache Warming Matters
- **Performance**: Reduces cold start latency
- **User Experience**: Faster response times from the start
- **Scalability**: Handles traffic spikes better
- **Reliability**: Reduces database load during peak times

### Cache Warming Strategies

#### 1. Predictive Warming
- **Definition**: Predict what data will be needed based on patterns
- **Use Case**: E-commerce, content delivery
- **Implementation**: Machine learning, historical data analysis

#### 2. Scheduled Warming
- **Definition**: Preload data at specific times
- **Use Case**: Daily reports, batch processing
- **Implementation**: Cron jobs, scheduled tasks

#### 3. Event-Driven Warming
- **Definition**: Warm cache based on events
- **Use Case**: User login, system startup
- **Implementation**: Event listeners, triggers

#### 4. Proactive Warming
- **Definition**: Warm cache based on user behavior
- **Use Case**: Recommendation systems, personalization
- **Implementation**: User tracking, behavior analysis

## 🛠️ Hands-on: Implementing Cache Warming

### Project Structure
```
projects/
├── cache-warming/
│   ├── main.go
│   ├── warming.go
│   ├── predictor.go
│   ├── scheduler.go
│   └── README.md
```

### Step 1: Cache Warming Implementation

```go
// warming.go
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

// CacheWarmer implements intelligent cache warming
type CacheWarmer struct {
    cache      *MemoryEfficientCache
    database   *Database
    config     WarmingConfig
    predictor  *DataPredictor
    scheduler  *WarmingScheduler
    mutex      sync.RWMutex
    stats      WarmingStats
}

// WarmingStats tracks warming statistics
type WarmingStats struct {
    ItemsWarmed    int64
    WarmingTime    time.Duration
    SuccessRate    float64
    LastWarming    time.Time
    Errors         int64
}

// NewCacheWarmer creates a new cache warmer
func NewCacheWarmer(cache *MemoryEfficientCache, db *Database, config WarmingConfig) *CacheWarmer {
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
    
    fmt.Printf("Predictive warming completed: %d items in %v\n", 
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
    
    fmt.Printf("Scheduled warming completed: %d items in %v\n", 
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
    
    fmt.Printf("Event-driven warming completed: %d items in %v\n", 
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
    
    fmt.Printf("Proactive warming completed: %d items in %v\n", 
        successCount, time.Since(start))
}

// warmKey warms a specific key
func (cw *CacheWarmer) warmKey(key string) error {
    // Check if key already exists in cache
    if _, err := cw.cache.Get(key); err == nil {
        return nil // Already in cache
    }
    
    // Read from database
    value, err := cw.database.Read(key)
    if err != nil {
        return err
    }
    
    // Store in cache
    return cw.cache.Set(key, value, 1*time.Hour)
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
```

### Step 2: Data Predictor Implementation

```go
// predictor.go
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

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
```

### Step 3: Warming Scheduler Implementation

```go
// scheduler.go
package main

import (
    "fmt"
    "sync"
    "time"
)

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

// UpdateSchedule updates a warming schedule
func (ws *WarmingScheduler) UpdateSchedule(key string, interval time.Duration) {
    ws.mutex.Lock()
    defer ws.mutex.Unlock()
    
    if schedule, exists := ws.schedules[key]; exists {
        schedule.Interval = interval
        schedule.NextRun = schedule.LastRun.Add(interval)
        ws.schedules[key] = schedule
    }
}

// EnableSchedule enables a warming schedule
func (ws *WarmingScheduler) EnableSchedule(key string) {
    ws.mutex.Lock()
    defer ws.mutex.Unlock()
    
    if schedule, exists := ws.schedules[key]; exists {
        schedule.Enabled = true
        ws.schedules[key] = schedule
    }
}

// DisableSchedule disables a warming schedule
func (ws *WarmingScheduler) DisableSchedule(key string) {
    ws.mutex.Lock()
    defer ws.mutex.Unlock()
    
    if schedule, exists := ws.schedules[key]; exists {
        schedule.Enabled = false
        ws.schedules[key] = schedule
    }
}

// GetSchedules returns all warming schedules
func (ws *WarmingScheduler) GetSchedules() map[string]WarmingSchedule {
    ws.mutex.RLock()
    defer ws.mutex.RUnlock()
    
    schedules := make(map[string]WarmingSchedule)
    for key, schedule := range ws.schedules {
        schedules[key] = schedule
    }
    
    return schedules
}
```

### Step 4: Main Application

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
    cache := NewMemoryEfficientCache(100 * 1024 * 1024) // 100MB
    
    // Test different warming strategies
    strategies := []WarmingStrategy{
        Predictive,
        Scheduled,
        EventDriven,
        Proactive,
    }
    
    for _, strategy := range strategies {
        fmt.Printf("\n=== Testing %s Warming Strategy ===\n", getWarmingStrategyName(strategy))
        
        // Create warming config
        config := WarmingConfig{
            Strategy:      strategy,
            Interval:      10 * time.Second,
            BatchSize:     100,
            MaxConcurrency: 10,
            Enabled:       true,
        }
        
        // Create cache warmer
        warmer := NewCacheWarmer(cache, db, config)
        
        // Wait for warming to complete
        time.Sleep(15 * time.Second)
        
        // Print statistics
        stats := warmer.GetStats()
        fmt.Printf("Warming Statistics:\n")
        fmt.Printf("  Items warmed: %d\n", stats.ItemsWarmed)
        fmt.Printf("  Warming time: %v\n", stats.WarmingTime)
        fmt.Printf("  Success rate: %.2f%%\n", stats.SuccessRate)
        fmt.Printf("  Errors: %d\n", stats.Errors)
        fmt.Printf("  Last warming: %v\n", stats.LastWarming)
        
        // Disable warming
        warmer.SetEnabled(false)
    }
    
    // Test manual warming
    fmt.Println("\n=== Testing Manual Warming ===")
    
    config := WarmingConfig{
        Strategy:      Predictive,
        Interval:      1 * time.Minute,
        BatchSize:     50,
        MaxConcurrency: 5,
        Enabled:       false,
    }
    
    warmer := NewCacheWarmer(cache, db, config)
    
    // Trigger manual warming
    fmt.Println("Triggering manual warming...")
    warmer.WarmNow()
    
    // Print final statistics
    stats := warmer.GetStats()
    fmt.Printf("Manual warming completed: %d items in %v\n", 
        stats.ItemsWarmed, stats.WarmingTime)
    
    // Test warming with different intervals
    fmt.Println("\n=== Testing Different Intervals ===")
    
    intervals := []time.Duration{
        5 * time.Second,
        10 * time.Second,
        30 * time.Second,
        1 * time.Minute,
    }
    
    for _, interval := range intervals {
        fmt.Printf("Testing interval: %v\n", interval)
        
        config.Interval = interval
        warmer = NewCacheWarmer(cache, db, config)
        
        // Wait for warming
        time.Sleep(interval + 2*time.Second)
        
        stats := warmer.GetStats()
        fmt.Printf("  Items warmed: %d\n", stats.ItemsWarmed)
        fmt.Printf("  Success rate: %.2f%%\n", stats.SuccessRate)
        
        warmer.SetEnabled(false)
    }
}

// Helper function
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
```

## 🧪 Exercise: Warming Performance Testing

### Task
Create a test to measure warming performance:

```go
// warming_test.go
package main

import (
    "testing"
    "time"
)

func TestWarmingPerformance(t *testing.T) {
    db := NewDatabase()
    cache := NewMemoryEfficientCache(100 * 1024 * 1024)
    
    config := WarmingConfig{
        Strategy:      Predictive,
        Interval:      1 * time.Second,
        BatchSize:     100,
        MaxConcurrency: 10,
        Enabled:       true,
    }
    
    warmer := NewCacheWarmer(cache, db, config)
    
    // Wait for warming
    time.Sleep(5 * time.Second)
    
    // Check if items were warmed
    stats := warmer.GetStats()
    if stats.ItemsWarmed == 0 {
        t.Error("No items were warmed")
    }
    
    if stats.SuccessRate < 50.0 {
        t.Errorf("Success rate too low: %.2f%%", stats.SuccessRate)
    }
}

func BenchmarkWarming(b *testing.B) {
    db := NewDatabase()
    cache := NewMemoryEfficientCache(100 * 1024 * 1024)
    
    config := WarmingConfig{
        Strategy:      Predictive,
        Interval:      1 * time.Millisecond,
        BatchSize:     1000,
        MaxConcurrency: 100,
        Enabled:       false,
    }
    
    warmer := NewCacheWarmer(cache, db, config)
    
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        warmer.WarmNow()
    }
}
```

## 📝 Key Takeaways

1. **Warming Strategies**: Choose based on your use case
2. **Performance**: Warming should not impact normal operations
3. **Prediction**: Use data patterns to predict what to warm
4. **Scheduling**: Plan warming during low-traffic periods
5. **Monitoring**: Track warming effectiveness

## 🎯 Practice Questions

1. When would you use predictive vs scheduled warming?
2. How do you measure warming effectiveness?
3. What are the trade-offs of different warming strategies?
4. How do you prevent warming from impacting performance?
5. How do you handle warming failures?

## 🚀 Next Steps

Tomorrow we'll cover:
- Advanced caching patterns
- Performance optimization techniques
- Week 1 review and assessment

---

**Your Notes:**
<!-- Add your notes, questions, and insights here -->

**Code Experiments:**
<!-- Document any code experiments or modifications you try -->

**Questions:**
<!-- Write down any questions you have -->
