# Lesson 11: Kafka Streams Fundamentals

## 🎯 Learning Objectives
By the end of this lesson, you will understand:
- Kafka Streams architecture and concepts
- Stream processing fundamentals
- KStreams, KTables, and GlobalKTables
- Windowing and time-based operations
- Stateful and stateless transformations
- Stream processing patterns and best practices
- Real-time analytics and event processing

## 📚 Theory Section

### 1. **Kafka Streams Overview**

#### **What is Kafka Streams?**
Kafka Streams is a client library for building applications and microservices that process data in real-time from Kafka topics. It provides:

- **Stream Processing**: Real-time data processing capabilities
- **Fault Tolerance**: Built-in fault tolerance and recovery
- **Scalability**: Automatic scaling and load balancing
- **Integration**: Seamless integration with Kafka ecosystem
- **Exactly-Once Semantics**: Guaranteed message processing

#### **Key Concepts:**
- **Stream**: An unbounded, continuously updating data set
- **Table**: A view of a stream at a point in time
- **Processor**: A node in the topology that processes data
- **Topology**: A directed acyclic graph (DAG) of processors
- **State Store**: Local storage for stateful operations

### 2. **Stream Processing Fundamentals**

#### **Stream vs Batch Processing:**
```go
// Batch Processing (Traditional)
func processBatch(data []Record) {
    for _, record := range data {
        processRecord(record)
    }
    // Process all at once
}

// Stream Processing (Real-time)
func processStream(record <-chan Record) {
    for record := range record {
        processRecord(record) // Process immediately
    }
}
```

#### **Stream Processing Characteristics:**
- **Unbounded**: Data streams have no defined end
- **Real-time**: Low latency processing
- **Continuous**: Always running and processing
- **Stateful**: Can maintain state across events
- **Fault-tolerant**: Handles failures gracefully

### 3. **Kafka Streams Architecture**

#### **Topology Structure:**
```
Input Topic → Source → Processor → Processor → Sink → Output Topic
                ↓
            State Store
```

#### **Core Components:**
- **KStream**: Represents a stream of records
- **KTable**: Represents a changelog stream
- **GlobalKTable**: Read-only table replicated across all instances
- **Processor**: Custom processing logic
- **State Store**: Local storage for stateful operations

### 4. **KStreams - Stream Processing**

#### **Basic KStream Operations:**
```go
type KStreamProcessor struct {
    inputTopic  string
    outputTopic string
    config      *sarama.Config
}

func (ksp *KStreamProcessor) ProcessStream() error {
    // Create Kafka Streams configuration
    config := sarama.NewConfig()
    config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
    
    // Create consumer group
    consumer, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "stream-processor", config)
    if err != nil {
        return err
    }
    defer consumer.Close()
    
    // Create handler
    handler := &StreamHandler{
        processor: ksp,
    }
    
    // Start consuming
    return consumer.Consume(context.Background(), []string{ksp.inputTopic}, handler)
}

type StreamHandler struct {
    processor *KStreamProcessor
}

func (sh *StreamHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
    for message := <-claim.Messages() {
        if message == nil {
            return nil
        }
        
        // Process message
        result := sh.processor.TransformMessage(message)
        
        // Send to output topic
        if err := sh.processor.SendToOutput(result); err != nil {
            log.Printf("Error sending to output: %v", err)
        }
        
        session.MarkMessage(message, "")
    }
    return nil
}
```

#### **Stream Transformations:**
```go
// Map transformation
func (ksp *KStreamProcessor) Map(transform func([]byte) []byte) *KStreamProcessor {
    ksp.transformations = append(ksp.transformations, Transform{
        Type: "map",
        Func: transform,
    })
    return ksp
}

// Filter transformation
func (ksp *KStreamProcessor) Filter(predicate func([]byte) bool) *KStreamProcessor {
    ksp.transformations = append(ksp.transformations, Transform{
        Type: "filter",
        Func: func(data []byte) []byte {
            if predicate(data) {
                return data
            }
            return nil
        },
    })
    return ksp
}

// FlatMap transformation
func (ksp *KStreamProcessor) FlatMap(transform func([]byte) [][]byte) *KStreamProcessor {
    ksp.transformations = append(ksp.transformations, Transform{
        Type: "flatmap",
        Func: func(data []byte) []byte {
            results := transform(data)
            // For simplicity, return first result
            if len(results) > 0 {
                return results[0]
            }
            return nil
        },
    })
    return ksp
}
```

### 5. **KTables - Table Processing**

#### **KTable Concepts:**
- **Changelog Stream**: Represents updates to a table
- **Key-Value Pairs**: Each record has a key and value
- **Upsert Semantics**: New records update existing ones
- **Compaction**: Old records are eventually removed

#### **KTable Operations:**
```go
type KTableProcessor struct {
    stateStore map[string]interface{}
    mutex      sync.RWMutex
}

func (ktp *KTableProcessor) ProcessTable() error {
    // Create consumer for changelog topic
    consumer, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "table-processor", nil)
    if err != nil {
        return err
    }
    defer consumer.Close()
    
    handler := &TableHandler{
        processor: ktp,
    }
    
    return consumer.Consume(context.Background(), []string{"user-table-changelog"}, handler)
}

func (ktp *KTableProcessor) Upsert(key string, value interface{}) {
    ktp.mutex.Lock()
    defer ktp.mutex.Unlock()
    
    ktp.stateStore[key] = value
}

func (ktp *KTableProcessor) Get(key string) (interface{}, bool) {
    ktp.mutex.RLock()
    defer ktp.mutex.RUnlock()
    
    value, exists := ktp.stateStore[key]
    return value, exists
}
```

### 6. **Windowing Operations**

#### **Time-based Windows:**
```go
type WindowProcessor struct {
    windowSize    time.Duration
    windowStores  map[string]*WindowStore
    mutex         sync.RWMutex
}

type WindowStore struct {
    Windows map[time.Time]map[string]interface{}
    mutex   sync.RWMutex
}

func (wp *WindowProcessor) ProcessWindowed(record *sarama.ConsumerMessage) {
    timestamp := time.Unix(record.Timestamp.Unix(), 0)
    windowStart := wp.getWindowStart(timestamp)
    
    wp.mutex.Lock()
    defer wp.mutex.Unlock()
    
    if wp.windowStores[record.Topic] == nil {
        wp.windowStores[record.Topic] = &WindowStore{
            Windows: make(map[time.Time]map[string]interface{}),
        }
    }
    
    store := wp.windowStores[record.Topic]
    store.mutex.Lock()
    defer store.mutex.Unlock()
    
    if store.Windows[windowStart] == nil {
        store.Windows[windowStart] = make(map[string]interface{})
    }
    
    store.Windows[windowStart][string(record.Key)] = record.Value
}

func (wp *WindowProcessor) getWindowStart(timestamp time.Time) time.Time {
    return timestamp.Truncate(wp.windowSize)
}
```

#### **Window Types:**
- **Tumbling Windows**: Fixed-size, non-overlapping windows
- **Hopping Windows**: Fixed-size, overlapping windows
- **Sliding Windows**: Variable-size windows based on data
- **Session Windows**: Windows based on activity gaps

### 7. **Stateful Operations**

#### **State Store Implementation:**
```go
type StateStore struct {
    data map[string]interface{}
    mutex sync.RWMutex
}

func (ss *StateStore) Put(key string, value interface{}) {
    ss.mutex.Lock()
    defer ss.mutex.Unlock()
    
    ss.data[key] = value
}

func (ss *StateStore) Get(key string) (interface{}, bool) {
    ss.mutex.RLock()
    defer ss.mutex.RUnlock()
    
    value, exists := ss.data[key]
    return value, exists
}

func (ss *StateStore) Delete(key string) {
    ss.mutex.Lock()
    defer ss.mutex.Unlock()
    
    delete(ss.data, key)
}
```

#### **Aggregations:**
```go
type AggregationProcessor struct {
    stateStore *StateStore
    windowSize time.Duration
}

func (ap *AggregationProcessor) Aggregate(key string, value interface{}, timestamp time.Time) {
    windowStart := ap.getWindowStart(timestamp)
    windowKey := fmt.Sprintf("%s:%d", key, windowStart.Unix())
    
    // Get existing aggregation
    existing, exists := ap.stateStore.Get(windowKey)
    if !exists {
        existing = make(map[string]interface{})
    }
    
    agg := existing.(map[string]interface{})
    
    // Update aggregation
    if count, ok := agg["count"]; ok {
        agg["count"] = count.(int) + 1
    } else {
        agg["count"] = 1
    }
    
    if sum, ok := agg["sum"]; ok {
        agg["sum"] = sum.(float64) + value.(float64)
    } else {
        agg["sum"] = value.(float64)
    }
    
    agg["avg"] = agg["sum"].(float64) / float64(agg["count"].(int))
    
    ap.stateStore.Put(windowKey, agg)
}
```

### 8. **Stream Processing Patterns**

#### **Event Sourcing Pattern:**
```go
type EventSourcingProcessor struct {
    eventStore map[string][]Event
    mutex      sync.RWMutex
}

type Event struct {
    ID        string
    Type      string
    Data      interface{}
    Timestamp time.Time
    Version   int
}

func (esp *EventSourcingProcessor) AppendEvent(streamID string, event Event) {
    esp.mutex.Lock()
    defer esp.mutex.Unlock()
    
    if esp.eventStore[streamID] == nil {
        esp.eventStore[streamID] = make([]Event, 0)
    }
    
    event.Version = len(esp.eventStore[streamID]) + 1
    esp.eventStore[streamID] = append(esp.eventStore[streamID], event)
}

func (esp *EventSourcingProcessor) GetStream(streamID string) []Event {
    esp.mutex.RLock()
    defer esp.mutex.RUnlock()
    
    return esp.eventStore[streamID]
}
```

#### **CQRS Pattern:**
```go
type CQRSProcessor struct {
    commandStore map[string]Command
    queryStore   map[string]Query
    mutex        sync.RWMutex
}

type Command struct {
    ID      string
    Type    string
    Data    interface{}
    Status  string
}

type Query struct {
    ID      string
    Type    string
    Data    interface{}
    Result  interface{}
}

func (cp *CQRSProcessor) HandleCommand(command Command) error {
    cp.mutex.Lock()
    defer cp.mutex.Unlock()
    
    command.Status = "processing"
    cp.commandStore[command.ID] = command
    
    // Process command and update query store
    result := cp.processCommand(command)
    
    query := Query{
        ID:     command.ID,
        Type:   command.Type,
        Data:   command.Data,
        Result: result,
    }
    
    cp.queryStore[command.ID] = query
    command.Status = "completed"
    cp.commandStore[command.ID] = command
    
    return nil
}
```

## 🧪 Hands-on Experiments

### Experiment 1: Basic Stream Processing

**Goal**: Create a simple stream processor that transforms messages

**Steps**:
1. Create a KStream processor
2. Implement basic transformations (map, filter)
3. Process messages from input topic
4. Send transformed messages to output topic

**Expected Results**:
- Real-time message transformation
- Fault-tolerant processing
- Scalable stream processing

### Experiment 2: Stateful Aggregations

**Goal**: Implement stateful aggregations with windowing

**Steps**:
1. Create a state store for aggregations
2. Implement windowing logic
3. Aggregate data within time windows
4. Output aggregated results

**Expected Results**:
- Time-based aggregations
- Stateful processing
- Windowed results

### Experiment 3: Event Sourcing

**Goal**: Implement event sourcing pattern

**Steps**:
1. Create event store
2. Implement event appending
3. Create event replay functionality
4. Test event sourcing

**Expected Results**:
- Event store implementation
- Event replay capability
- Audit trail maintenance

### Experiment 4: CQRS Implementation

**Goal**: Implement Command Query Responsibility Segregation

**Steps**:
1. Create command and query stores
2. Implement command handling
3. Implement query processing
4. Test CQRS pattern

**Expected Results**:
- Command processing
- Query optimization
- Separation of concerns

## 📊 Stream Processing Metrics

### **Key Performance Indicators:**
- **Throughput**: Messages processed per second
- **Latency**: End-to-end processing time
- **Backpressure**: Queue depth and processing lag
- **State Size**: Size of state stores
- **Window Processing**: Window completion rate

### **Monitoring Stream Processing:**
```go
type StreamMetrics struct {
    MessagesProcessed int64
    ProcessingLatency time.Duration
    StateStoreSize    int64
    WindowCount       int64
    ErrorCount        int64
}

func (sm *StreamMetrics) RecordMessage(latency time.Duration) {
    atomic.AddInt64(&sm.MessagesProcessed, 1)
    sm.ProcessingLatency = latency
}

func (sm *StreamMetrics) RecordError() {
    atomic.AddInt64(&sm.ErrorCount, 1)
}
```

## 🎯 Key Takeaways

1. **Kafka Streams** provides powerful stream processing capabilities
2. **KStreams** are for continuous stream processing
3. **KTables** are for table-like operations with upsert semantics
4. **Windowing** enables time-based aggregations
5. **Stateful operations** require state stores
6. **Stream processing patterns** solve common use cases
7. **Monitoring** is crucial for stream processing applications

## 📝 Lesson 11 Assessment Questions

1. **What is Kafka Streams and how does it differ from batch processing?**
2. **What are the key components of a Kafka Streams topology?**
3. **How do KStreams and KTables differ?**
4. **What are the different types of windows in stream processing?**
5. **How do you implement stateful operations in Kafka Streams?**
6. **What is the event sourcing pattern?**
7. **How do you implement CQRS with Kafka Streams?**
8. **What metrics should you monitor for stream processing?**

---

## 🔄 Next Lesson Preview: Advanced Stream Processing

**What we'll learn**:
- Advanced windowing operations
- Stream joins and merges
- Complex event processing
- Stream processing optimization
- Fault tolerance and recovery
- Performance tuning for streams

**Hands-on experiments**:
- Implement complex stream topologies
- Build real-time analytics pipelines
- Optimize stream processing performance
- Handle stream processing failures
