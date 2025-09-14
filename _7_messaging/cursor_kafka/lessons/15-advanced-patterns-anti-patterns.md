# Lesson 15: Advanced Patterns and Anti-patterns

## 🎯 Learning Objectives
By the end of this lesson, you will understand:
- Advanced Kafka design patterns and their applications
- Common anti-patterns and how to avoid them
- Microservices integration patterns with Kafka
- Event-driven architecture patterns
- Data consistency patterns and trade-offs
- Performance optimization patterns
- Production-ready pattern implementations

## 📚 Theory Section

### 1. **Advanced Design Patterns**

#### **Saga Pattern for Distributed Transactions:**
```go
type SagaOrchestrator struct {
    steps    []SagaStep
    state    SagaState
    mutex    sync.RWMutex
    logger   *Logger
}

type SagaStep struct {
    ID          string
    Name        string
    Execute     func(context.Context) error
    Compensate  func(context.Context) error
    Status      StepStatus
    RetryCount  int
    MaxRetries  int
}

type SagaState struct {
    SagaID      string
    Status      SagaStatus
    CurrentStep int
    Data        map[string]interface{}
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

func (so *SagaOrchestrator) ExecuteSaga(ctx context.Context) error {
    so.mutex.Lock()
    defer so.mutex.Unlock()
    
    so.state.Status = SAGA_RUNNING
    so.state.UpdatedAt = time.Now()
    
    for i, step := range so.steps {
        so.state.CurrentStep = i
        
        // Execute step with retry logic
        if err := so.executeStepWithRetry(ctx, step); err != nil {
            // Compensate for completed steps
            so.compensateCompletedSteps(ctx, i)
            return fmt.Errorf("saga execution failed at step %d: %w", i, err)
        }
        
        step.Status = STEP_COMPLETED
        so.steps[i] = step
    }
    
    so.state.Status = SAGA_COMPLETED
    so.state.UpdatedAt = time.Now()
    return nil
}

func (so *SagaOrchestrator) executeStepWithRetry(ctx context.Context, step SagaStep) error {
    for step.RetryCount <= step.MaxRetries {
        if err := step.Execute(ctx); err != nil {
            step.RetryCount++
            so.logger.Warn("Step execution failed, retrying", 
                "step", step.Name, 
                "retry", step.RetryCount,
                "error", err)
            
            if step.RetryCount <= step.MaxRetries {
                time.Sleep(time.Duration(step.RetryCount) * time.Second)
                continue
            }
            return err
        }
        return nil
    }
    return fmt.Errorf("max retries exceeded for step %s", step.Name)
}
```

#### **Outbox Pattern for Event Sourcing:**
```go
type OutboxEvent struct {
    ID          string    `json:"id"`
    AggregateID  string    `json:"aggregate_id"`
    EventType   string    `json:"event_type"`
    EventData   []byte    `json:"event_data"`
    CreatedAt   time.Time `json:"created_at"`
    Processed   bool      `json:"processed"`
    ProcessedAt *time.Time `json:"processed_at,omitempty"`
}

type OutboxProcessor struct {
    db          *sql.DB
    producer    sarama.SyncProducer
    topic       string
    batchSize   int
    interval    time.Duration
    logger      *Logger
}

func (op *OutboxProcessor) ProcessOutboxEvents(ctx context.Context) error {
    ticker := time.NewTicker(op.interval)
    defer ticker.Stop()
    
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        case <-ticker.C:
            if err := op.processBatch(ctx); err != nil {
                op.logger.Error("Failed to process outbox batch", "error", err)
            }
        }
    }
}

func (op *OutboxProcessor) processBatch(ctx context.Context) error {
    // Get unprocessed events
    events, err := op.getUnprocessedEvents(op.batchSize)
    if err != nil {
        return err
    }
    
    if len(events) == 0 {
        return nil
    }
    
    // Process events in transaction
    tx, err := op.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    for _, event := range events {
        // Publish to Kafka
        if err := op.publishEvent(event); err != nil {
            return err
        }
        
        // Mark as processed
        if err := op.markAsProcessed(tx, event.ID); err != nil {
            return err
        }
    }
    
    return tx.Commit()
}
```

#### **CQRS with Event Sourcing:**
```go
type CQRSManager struct {
    commandStore  CommandStore
    queryStore    QueryStore
    eventStore    EventStore
    projections   map[string]Projection
    mutex         sync.RWMutex
}

type Command struct {
    ID        string
    Type      string
    Data      map[string]interface{}
    UserID    string
    Timestamp time.Time
}

type Query struct {
    ID       string
    Type     string
    Criteria map[string]interface{}
    UserID   string
}

type Projection struct {
    Name     string
    Handler  func(Event) error
    Position int64
}

func (cqrs *CQRSManager) HandleCommand(ctx context.Context, cmd Command) error {
    // Validate command
    if err := cqrs.validateCommand(cmd); err != nil {
        return err
    }
    
    // Execute command
    events, err := cqrs.executeCommand(ctx, cmd)
    if err != nil {
        return err
    }
    
    // Store events
    for _, event := range events {
        if err := cqrs.eventStore.AppendEvent(event); err != nil {
            return err
        }
    }
    
    // Update projections
    for _, event := range events {
        if err := cqrs.updateProjections(event); err != nil {
            return err
        }
    }
    
    return nil
}

func (cqrs *CQRSManager) HandleQuery(ctx context.Context, query Query) (interface{}, error) {
    // Get projection
    projection, exists := cqrs.projections[query.Type]
    if !exists {
        return nil, fmt.Errorf("projection not found: %s", query.Type)
    }
    
    // Query projection
    return cqrs.queryStore.Query(projection.Name, query.Criteria)
}
```

### 2. **Microservices Integration Patterns**

#### **Event-Driven Microservices:**
```go
type Microservice struct {
    name        string
    consumer    sarama.ConsumerGroup
    producer    sarama.SyncProducer
    handlers    map[string]EventHandler
    logger      *Logger
    metrics     *Metrics
}

type EventHandler interface {
    Handle(ctx context.Context, event Event) error
    CanHandle(eventType string) bool
}

type Event struct {
    ID        string                 `json:"id"`
    Type      string                 `json:"type"`
    Source    string                 `json:"source"`
    Data      map[string]interface{} `json:"data"`
    Timestamp time.Time             `json:"timestamp"`
    Version   int                   `json:"version"`
}

func (ms *Microservice) Start(ctx context.Context) error {
    // Start consuming events
    go func() {
        for {
            select {
            case <-ctx.Done():
                return
            default:
                if err := ms.consumeEvents(ctx); err != nil {
                    ms.logger.Error("Error consuming events", "error", err)
                }
            }
        }
    }()
    
    return nil
}

func (ms *Microservice) consumeEvents(ctx context.Context) error {
    handler := &EventConsumerHandler{
        microservice: ms,
    }
    
    return ms.consumer.Consume(ctx, []string{"microservice-events"}, handler)
}

func (ms *Microservice) PublishEvent(ctx context.Context, event Event) error {
    event.Source = ms.name
    event.Timestamp = time.Now()
    
    data, err := json.Marshal(event)
    if err != nil {
        return err
    }
    
    message := &sarama.ProducerMessage{
        Topic: "microservice-events",
        Key:   sarama.StringEncoder(event.ID),
        Value: sarama.ByteEncoder(data),
        Headers: []sarama.RecordHeader{
            {Key: []byte("event-type"), Value: []byte(event.Type)},
            {Key: []byte("source"), Value: []byte(event.Source)},
        },
    }
    
    _, _, err = ms.producer.SendMessage(message)
    return err
}
```

#### **API Gateway Pattern:**
```go
type APIGateway struct {
    routes      map[string]Route
    kafkaClient sarama.Client
    logger      *Logger
    metrics     *Metrics
}

type Route struct {
    Path        string
    Method      string
    Handler     func(*http.Request) (*http.Response, error)
    Topics      []string
    Permissions []string
}

func (gw *APIGateway) HandleRequest(req *http.Request) (*http.Response, error) {
    route, exists := gw.routes[req.URL.Path]
    if !exists {
        return &http.Response{
            StatusCode: http.StatusNotFound,
            Body:       io.NopCloser(strings.NewReader("Not Found")),
        }, nil
    }
    
    // Check permissions
    if err := gw.checkPermissions(req, route.Permissions); err != nil {
        return &http.Response{
            StatusCode: http.StatusForbidden,
            Body:       io.NopCloser(strings.NewReader("Forbidden")),
        }, nil
    }
    
    // Route to appropriate handler
    return route.Handler(req)
}

func (gw *APIGateway) PublishToKafka(topic string, data []byte) error {
    producer, err := sarama.NewSyncProducerFromClient(gw.kafkaClient)
    if err != nil {
        return err
    }
    defer producer.Close()
    
    message := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.ByteEncoder(data),
    }
    
    _, _, err = producer.SendMessage(message)
    return err
}
```

### 3. **Data Consistency Patterns**

#### **Eventual Consistency with Conflict Resolution:**
```go
type ConflictResolver struct {
    strategies map[string]ConflictResolutionStrategy
    mutex      sync.RWMutex
}

type ConflictResolutionStrategy interface {
    Resolve(conflict Conflict) (interface{}, error)
    CanResolve(conflictType string) bool
}

type Conflict struct {
    Type        string
    Key         string
    Values      []interface{}
    Timestamps  []time.Time
    Sources     []string
}

func (cr *ConflictResolver) ResolveConflict(conflict Conflict) (interface{}, error) {
    cr.mutex.RLock()
    strategy, exists := cr.strategies[conflict.Type]
    cr.mutex.RUnlock()
    
    if !exists {
        return nil, fmt.Errorf("no resolution strategy for conflict type: %s", conflict.Type)
    }
    
    return strategy.Resolve(conflict)
}

// Last-Write-Wins Strategy
type LastWriteWinsStrategy struct{}

func (lww *LastWriteWinsStrategy) Resolve(conflict Conflict) (interface{}, error) {
    if len(conflict.Values) == 0 {
        return nil, fmt.Errorf("no values to resolve")
    }
    
    // Find the value with the latest timestamp
    latestIndex := 0
    for i, timestamp := range conflict.Timestamps {
        if timestamp.After(conflict.Timestamps[latestIndex]) {
            latestIndex = i
        }
    }
    
    return conflict.Values[latestIndex], nil
}

// Custom Business Logic Strategy
type BusinessLogicStrategy struct {
    rules map[string]func(Conflict) (interface{}, error)
}

func (bls *BusinessLogicStrategy) Resolve(conflict Conflict) (interface{}, error) {
    rule, exists := bls.rules[conflict.Type]
    if !exists {
        return nil, fmt.Errorf("no business rule for conflict type: %s", conflict.Type)
    }
    
    return rule(conflict)
}
```

#### **Two-Phase Commit Pattern:**
```go
type TwoPhaseCommitCoordinator struct {
    participants map[string]Participant
    logger       *Logger
    mutex        sync.RWMutex
}

type Participant struct {
    ID       string
    Prepare  func() error
    Commit   func() error
    Rollback func() error
}

func (tpc *TwoPhaseCommitCoordinator) ExecuteTransaction(participantIDs []string) error {
    tpc.mutex.Lock()
    defer tpc.mutex.Unlock()
    
    // Phase 1: Prepare
    prepared := make([]string, 0)
    for _, id := range participantIDs {
        participant, exists := tpc.participants[id]
        if !exists {
            tpc.rollbackPrepared(prepared)
            return fmt.Errorf("participant not found: %s", id)
        }
        
        if err := participant.Prepare(); err != nil {
            tpc.logger.Error("Prepare failed", "participant", id, "error", err)
            tpc.rollbackPrepared(prepared)
            return err
        }
        
        prepared = append(prepared, id)
    }
    
    // Phase 2: Commit
    for _, id := range prepared {
        participant := tpc.participants[id]
        if err := participant.Commit(); err != nil {
            tpc.logger.Error("Commit failed", "participant", id, "error", err)
            // Try to rollback remaining participants
            tpc.rollbackRemaining(prepared, id)
            return err
        }
    }
    
    return nil
}
```

### 4. **Performance Optimization Patterns**

#### **Batching and Buffering:**
```go
type BatchedProducer struct {
    producer    sarama.AsyncProducer
    buffer      chan *sarama.ProducerMessage
    batchSize   int
    flushInterval time.Duration
    logger      *Logger
    metrics     *Metrics
}

func (bp *BatchedProducer) Start(ctx context.Context) error {
    go bp.batchProcessor(ctx)
    return nil
}

func (bp *BatchedProducer) batchProcessor(ctx context.Context) {
    ticker := time.NewTicker(bp.flushInterval)
    defer ticker.Stop()
    
    batch := make([]*sarama.ProducerMessage, 0, bp.batchSize)
    
    for {
        select {
        case <-ctx.Done():
            bp.flushBatch(batch)
            return
        case <-ticker.C:
            bp.flushBatch(batch)
            batch = batch[:0]
        case message := <-bp.buffer:
            batch = append(batch, message)
            if len(batch) >= bp.batchSize {
                bp.flushBatch(batch)
                batch = batch[:0]
            }
        }
    }
}

func (bp *BatchedProducer) flushBatch(batch []*sarama.ProducerMessage) {
    if len(batch) == 0 {
        return
    }
    
    for _, message := range batch {
        bp.producer.Input() <- message
    }
    
    bp.metrics.RecordBatchSize(len(batch))
}
```

#### **Connection Pooling:**
```go
type ConnectionPool struct {
    connections chan sarama.Client
    factory     func() (sarama.Client, error)
    maxSize     int
    mutex       sync.RWMutex
    logger      *Logger
}

func (cp *ConnectionPool) Get() (sarama.Client, error) {
    select {
    case conn := <-cp.connections:
        return conn, nil
    default:
        return cp.factory()
    }
}

func (cp *ConnectionPool) Put(conn sarama.Client) {
    select {
    case cp.connections <- conn:
    default:
        // Pool is full, close connection
        conn.Close()
    }
}

func (cp *ConnectionPool) Close() {
    close(cp.connections)
    for conn := range cp.connections {
        conn.Close()
    }
}
```

### 5. **Common Anti-patterns and How to Avoid Them**

#### **Anti-pattern: Tight Coupling**
```go
// ❌ BAD: Tight coupling
type OrderService struct {
    paymentService *PaymentService
    inventoryService *InventoryService
    notificationService *NotificationService
}

func (os *OrderService) ProcessOrder(order Order) error {
    // Direct calls to other services
    if err := os.paymentService.ProcessPayment(order.Payment); err != nil {
        return err
    }
    
    if err := os.inventoryService.ReserveItems(order.Items); err != nil {
        return err
    }
    
    if err := os.notificationService.SendConfirmation(order.CustomerID); err != nil {
        return err
    }
    
    return nil
}

// ✅ GOOD: Loose coupling with events
type OrderService struct {
    producer sarama.SyncProducer
    logger   *Logger
}

func (os *OrderService) ProcessOrder(order Order) error {
    // Publish events instead of direct calls
    events := []Event{
        {Type: "PaymentRequested", Data: order.Payment},
        {Type: "InventoryReservationRequested", Data: order.Items},
        {Type: "OrderConfirmed", Data: order},
    }
    
    for _, event := range events {
        if err := os.publishEvent(event); err != nil {
            return err
        }
    }
    
    return nil
}
```

#### **Anti-pattern: Synchronous Processing**
```go
// ❌ BAD: Synchronous processing
func (ps *PaymentService) ProcessPayment(payment Payment) error {
    // Synchronous validation
    if err := ps.validatePayment(payment); err != nil {
        return err
    }
    
    // Synchronous external API call
    if err := ps.callPaymentGateway(payment); err != nil {
        return err
    }
    
    // Synchronous database update
    if err := ps.updatePaymentStatus(payment.ID, "completed"); err != nil {
        return err
    }
    
    return nil
}

// ✅ GOOD: Asynchronous processing
func (ps *PaymentService) ProcessPayment(payment Payment) error {
    // Publish event for asynchronous processing
    event := Event{
        Type: "PaymentProcessingRequested",
        Data: payment,
    }
    
    return ps.publishEvent(event)
}

func (ps *PaymentService) handlePaymentProcessing(event Event) error {
    payment := event.Data.(Payment)
    
    // Asynchronous validation
    if err := ps.validatePaymentAsync(payment); err != nil {
        return err
    }
    
    // Asynchronous external API call
    go ps.callPaymentGatewayAsync(payment)
    
    return nil
}
```

#### **Anti-pattern: Ignoring Error Handling**
```go
// ❌ BAD: Ignoring errors
func (cs *ConsumerService) ConsumeMessages() {
    for message := range cs.messages {
        cs.processMessage(message) // Error ignored
    }
}

// ✅ GOOD: Proper error handling
func (cs *ConsumerService) ConsumeMessages() {
    for message := range cs.messages {
        if err := cs.processMessage(message); err != nil {
            cs.logger.Error("Failed to process message", "error", err)
            
            // Send to dead letter queue
            if err := cs.sendToDLQ(message, err); err != nil {
                cs.logger.Error("Failed to send to DLQ", "error", err)
            }
        }
    }
}
```

## 🧪 Hands-on Experiments

### Experiment 1: Saga Pattern Implementation

**Goal**: Implement distributed transaction handling with Saga pattern

**Steps**:
1. Create saga orchestrator
2. Define saga steps with compensation
3. Implement retry logic
4. Test failure scenarios

**Expected Results**:
- Distributed transactions handled correctly
- Compensation logic working
- Retry mechanism functioning

### Experiment 2: Outbox Pattern Setup

**Goal**: Implement reliable event publishing with outbox pattern

**Steps**:
1. Create outbox table
2. Implement outbox processor
3. Test event publishing
4. Verify reliability

**Expected Results**:
- Events published reliably
- No data loss
- Proper error handling

### Experiment 3: CQRS Implementation

**Goal**: Implement Command Query Responsibility Segregation

**Steps**:
1. Create command and query stores
2. Implement event sourcing
3. Create projections
4. Test CQRS operations

**Expected Results**:
- Commands and queries separated
- Event sourcing working
- Projections updating correctly

### Experiment 4: Anti-pattern Detection

**Goal**: Identify and fix common anti-patterns

**Steps**:
1. Analyze existing code for anti-patterns
2. Refactor tight coupling
3. Implement proper error handling
4. Test improvements

**Expected Results**:
- Anti-patterns identified and fixed
- Code quality improved
- Better maintainability

## 📊 Pattern Performance Metrics

### **Key Performance Indicators:**
- **Transaction Success Rate**: Percentage of successful distributed transactions
- **Event Processing Latency**: Time to process events
- **Error Recovery Time**: Time to recover from failures
- **Throughput**: Messages processed per second
- **Resource Utilization**: CPU, memory, network usage

### **Pattern Monitoring:**
```go
type PatternMetrics struct {
    SagaSuccessRate      float64
    EventProcessingTime  time.Duration
    ErrorRecoveryTime    time.Duration
    Throughput           int64
    ResourceUsage        ResourceMetrics
}

type ResourceMetrics struct {
    CPUUsage    float64
    MemoryUsage float64
    NetworkIO   int64
}
```

## 🎯 Key Takeaways

1. **Advanced patterns** solve complex distributed systems problems
2. **Anti-patterns** should be identified and avoided
3. **Microservices integration** requires careful design
4. **Data consistency** has trade-offs and requires strategies
5. **Performance optimization** patterns improve system efficiency
6. **Error handling** is crucial for reliability
7. **Monitoring** is essential for pattern effectiveness

## 📝 Lesson 15 Assessment Questions

1. **What is the Saga pattern and when should you use it?**
2. **How does the Outbox pattern ensure reliable event publishing?**
3. **What are the benefits of CQRS with Event Sourcing?**
4. **What are common anti-patterns in Kafka applications?**
5. **How do you handle data consistency in distributed systems?**
6. **What performance optimization patterns can you use?**
7. **How do you monitor pattern effectiveness?**
8. **What are the trade-offs of different consistency patterns?**

---

## 🔄 Next Lesson Preview: Production Architecture Design

**What we'll learn**:
- Production-ready architecture design
- Scalability and performance planning
- Monitoring and observability strategies
- Disaster recovery and business continuity
- Security architecture and compliance
- Cost optimization and resource planning

**Hands-on experiments**:
- Design production architecture
- Implement monitoring strategies
- Test disaster recovery procedures
- Optimize for cost and performance
