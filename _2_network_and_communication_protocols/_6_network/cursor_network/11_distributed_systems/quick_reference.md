# Distributed Systems Networking Quick Reference Guide

## 🚀 Essential Distributed Systems Concepts

### CAP Theorem
- **Consistency (C)**: All nodes see the same data at the same time
- **Availability (A)**: System remains operational and responsive
- **Partition Tolerance (P)**: System continues operating despite network partitions
- **Trade-offs**: CP Systems, AP Systems, CA Systems

### Distributed System Challenges
- **Network Partitions**: Network splits can isolate components
- **Partial Failures**: Some components fail while others continue
- **Consistency**: Maintaining data consistency across nodes
- **Availability**: Ensuring system remains available during failures
- **Clock Synchronization**: Coordinating time across nodes

## 🔧 Go Implementation Patterns

### Raft Consensus Algorithm
```go
type RaftNode struct {
    Node
    ElectionTimeout  time.Duration
    HeartbeatTimeout time.Duration
    VoteCount        int32
    LeaderID         string
    StopChan         chan struct{}
    RequestVoteChan  chan RequestVoteRequest
    AppendEntriesChan chan AppendEntriesRequest
    ClientRequestChan chan ClientRequest
}

type Node struct {
    ID          string
    Address     string
    Port        int
    State       string // "leader", "follower", "candidate"
    Term        int64
    VotedFor    string
    Log         []LogEntry
    CommitIndex int64
    LastApplied int64
    NextIndex   map[string]int64
    MatchIndex  map[string]int64
    Peers       map[string]*Node
    IsHealthy   bool
    LastHeartbeat time.Time
    mutex       sync.RWMutex
}
```

### Service Discovery
```go
type ServiceRegistry struct {
    Services map[string]*Service
    mutex    sync.RWMutex
}

type Service struct {
    ID          string
    Name        string
    Address     string
    Port        int
    HealthCheck string
    Tags        []string
    LastSeen    time.Time
    IsHealthy   bool
}

func (sr *ServiceRegistry) Register(service *Service) {
    sr.mutex.Lock()
    defer sr.mutex.Unlock()
    
    service.LastSeen = time.Now()
    service.IsHealthy = true
    sr.Services[service.ID] = service
}

func (sr *ServiceRegistry) GetServicesByName(name string) []*Service {
    sr.mutex.RLock()
    defer sr.mutex.RUnlock()
    
    var services []*Service
    for _, service := range sr.Services {
        if service.Name == name && service.IsHealthy {
            services = append(services, service)
        }
    }
    return services
}
```

### Load Balancing
```go
type LoadBalancer struct {
    Services    map[string][]*Service
    Algorithm   string
    Current     map[string]int32
    mutex       sync.RWMutex
}

func (lb *LoadBalancer) GetService(serviceName string) *Service {
    lb.mutex.RLock()
    defer lb.mutex.RUnlock()
    
    services, exists := lb.Services[serviceName]
    if !exists || len(services) == 0 {
        return nil
    }
    
    // Filter healthy services
    var healthyServices []*Service
    for _, service := range services {
        if service.IsHealthy {
            healthyServices = append(healthyServices, service)
        }
    }
    
    if len(healthyServices) == 0 {
        return nil
    }
    
    switch lb.Algorithm {
    case "round_robin":
        return lb.getRoundRobinService(serviceName, healthyServices)
    case "random":
        return lb.getRandomService(healthyServices)
    case "least_connections":
        return lb.getLeastConnectionsService(healthyServices)
    default:
        return lb.getRoundRobinService(serviceName, healthyServices)
    }
}
```

## 🔄 Distributed Patterns

### Circuit Breaker Pattern
```go
type CircuitBreaker struct {
    Name          string
    MaxRequests   int32
    Interval      time.Duration
    Timeout       time.Duration
    FailureCount  int32
    SuccessCount  int32
    LastFailTime  time.Time
    State         string // "closed", "open", "half-open"
    mutex         sync.RWMutex
}

func (cb *CircuitBreaker) Call(fn func() (interface{}, error)) (interface{}, error) {
    cb.mutex.Lock()
    defer cb.mutex.Unlock()
    
    // Check if circuit should be opened
    if cb.State == "closed" && cb.FailureCount >= cb.MaxRequests {
        if time.Since(cb.LastFailTime) > cb.Interval {
            cb.State = "open"
            cb.FailureCount = 0
        }
    }
    
    // Check if circuit should be half-opened
    if cb.State == "open" && time.Since(cb.LastFailTime) > cb.Timeout {
        cb.State = "half-open"
        cb.SuccessCount = 0
    }
    
    // Execute function based on state
    switch cb.State {
    case "closed", "half-open":
        result, err := fn()
        if err != nil {
            cb.FailureCount++
            cb.LastFailTime = time.Now()
            if cb.State == "half-open" {
                cb.State = "open"
            }
            return nil, err
        }
        
        cb.SuccessCount++
        if cb.State == "half-open" && cb.SuccessCount >= cb.MaxRequests {
            cb.State = "closed"
            cb.FailureCount = 0
        }
        
        return result, nil
    case "open":
        return nil, fmt.Errorf("circuit breaker is open")
    default:
        return nil, fmt.Errorf("unknown circuit breaker state")
    }
}
```

### Bulkhead Pattern
```go
type Bulkhead struct {
    Name        string
    MaxWorkers  int
    MaxQueue    int
    Workers     chan struct{}
    Queue       chan func()
    IsRunning   bool
    StopChan    chan struct{}
    mutex       sync.RWMutex
}

func (b *Bulkhead) Execute(fn func()) error {
    select {
    case b.Queue <- fn:
        return nil
    default:
        return fmt.Errorf("bulkhead queue is full")
    }
}

func (b *Bulkhead) worker() {
    for {
        select {
        case <-b.StopChan:
            return
        case fn := <-b.Queue:
            b.Workers <- struct{}{} // Acquire worker
            fn()
            <-b.Workers // Release worker
        }
    }
}
```

### Saga Pattern
```go
type Saga struct {
    ID          string
    Steps       []SagaStep
    CurrentStep int
    State       string // "running", "completed", "failed", "compensating"
    Result      interface{}
    Error       error
    mutex       sync.RWMutex
}

type SagaStep struct {
    ID           string
    Action       func() (interface{}, error)
    Compensation func() error
    Timeout      time.Duration
}

func (s *Saga) Execute() error {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    
    s.State = "running"
    
    for i, step := range s.Steps {
        s.CurrentStep = i
        
        // Execute step with timeout
        ctx, cancel := context.WithTimeout(context.Background(), step.Timeout)
        done := make(chan struct{})
        var result interface{}
        var err error
        
        go func() {
            defer close(done)
            result, err = step.Action()
        }()
        
        select {
        case <-done:
            cancel()
            if err != nil {
                s.State = "failed"
                s.Error = err
                // Compensate for completed steps
                s.compensate(i - 1)
                return err
            }
        case <-ctx.Done():
            cancel()
            s.State = "failed"
            s.Error = fmt.Errorf("step %s timed out", step.ID)
            // Compensate for completed steps
            s.compensate(i - 1)
            return s.Error
        }
    }
    
    s.State = "completed"
    s.Result = result
    return nil
}
```

## 📊 Event Sourcing & CQRS

### Event Store
```go
type EventStore struct {
    Events map[string][]Event
    mutex  sync.RWMutex
}

type Event struct {
    ID        string
    Type      string
    Data      interface{}
    Timestamp time.Time
    Version   int64
}

func (es *EventStore) Append(streamID string, eventType string, data interface{}) error {
    es.mutex.Lock()
    defer es.mutex.Unlock()
    
    event := Event{
        ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
        Type:      eventType,
        Data:      data,
        Timestamp: time.Now(),
        Version:   int64(len(es.Events[streamID])),
    }
    
    es.Events[streamID] = append(es.Events[streamID], event)
    return nil
}

func (es *EventStore) GetEvents(streamID string) ([]Event, error) {
    es.mutex.RLock()
    defer es.mutex.RUnlock()
    
    events, exists := es.Events[streamID]
    if !exists {
        return nil, fmt.Errorf("stream not found: %s", streamID)
    }
    
    // Return a copy
    result := make([]Event, len(events))
    copy(result, events)
    return result, nil
}
```

### CQRS Implementation
```go
type CQRS struct {
    CommandStore *EventStore
    QueryStore   map[string]interface{}
    mutex        sync.RWMutex
}

func (cqrs *CQRS) ExecuteCommand(streamID string, commandType string, data interface{}) error {
    // Append command to event store
    if err := cqrs.CommandStore.Append(streamID, commandType, data); err != nil {
        return err
    }
    
    // Update query store based on command
    cqrs.mutex.Lock()
    defer cqrs.mutex.Unlock()
    
    switch commandType {
    case "CreateUser":
        userData := data.(map[string]interface{})
        cqrs.QueryStore[streamID] = map[string]interface{}{
            "id":    streamID,
            "name":  userData["name"],
            "email": userData["email"],
            "state": "active",
        }
    case "UpdateUser":
        userData := data.(map[string]interface{})
        if existing, exists := cqrs.QueryStore[streamID]; exists {
            existingUser := existing.(map[string]interface{})
            if name, ok := userData["name"]; ok {
                existingUser["name"] = name
            }
            if email, ok := userData["email"]; ok {
                existingUser["email"] = email
            }
        }
    }
    
    return nil
}

func (cqrs *CQRS) Query(streamID string) (interface{}, error) {
    cqrs.mutex.RLock()
    defer cqrs.mutex.RUnlock()
    
    result, exists := cqrs.QueryStore[streamID]
    if !exists {
        return nil, fmt.Errorf("stream not found: %s", streamID)
    }
    
    return result, nil
}
```

## 🔄 Message Queuing

### Message Queue
```go
type MessageQueue struct {
    Messages chan Message
    Subscribers map[string][]chan Message
    mutex    sync.RWMutex
}

type Message struct {
    ID        string
    Topic     string
    Data      interface{}
    Timestamp time.Time
    TTL       time.Duration
}

func (mq *MessageQueue) Publish(topic string, data interface{}) error {
    message := Message{
        ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
        Topic:     topic,
        Data:      data,
        Timestamp: time.Now(),
        TTL:       24 * time.Hour,
    }
    
    select {
    case mq.Messages <- message:
        return nil
    default:
        return fmt.Errorf("queue full")
    }
}

func (mq *MessageQueue) Subscribe(topic string) <-chan Message {
    mq.mutex.Lock()
    defer mq.mutex.Unlock()
    
    ch := make(chan Message, 100)
    mq.Subscribers[topic] = append(mq.Subscribers[topic], ch)
    return ch
}
```

### Distributed Cache
```go
type DistributedCache struct {
    Data  map[string]CacheItem
    mutex sync.RWMutex
}

type CacheItem struct {
    Value     interface{}
    ExpiresAt time.Time
    Version   int64
}

func (dc *DistributedCache) Set(key string, value interface{}, ttl time.Duration) {
    dc.mutex.Lock()
    defer dc.mutex.Unlock()
    
    dc.Data[key] = CacheItem{
        Value:     value,
        ExpiresAt: time.Now().Add(ttl),
        Version:   time.Now().UnixNano(),
    }
}

func (dc *DistributedCache) Get(key string) (interface{}, bool) {
    dc.mutex.RLock()
    defer dc.mutex.RUnlock()
    
    item, exists := dc.Data[key]
    if !exists {
        return nil, false
    }
    
    if time.Now().After(item.ExpiresAt) {
        // Item expired
        dc.mutex.RUnlock()
        dc.mutex.Lock()
        delete(dc.Data, key)
        dc.mutex.Unlock()
        dc.mutex.RLock()
        return nil, false
    }
    
    return item.Value, true
}
```

## 🛡️ Resilience Patterns

### Retry Policy
```go
type RetryPolicy struct {
    MaxAttempts int
    BaseDelay   time.Duration
    MaxDelay    time.Duration
    Multiplier  float64
    Jitter      bool
}

func (rp *RetryPolicy) Execute(fn func() (interface{}, error)) (interface{}, error) {
    var lastErr error
    delay := rp.BaseDelay
    
    for attempt := 0; attempt < rp.MaxAttempts; attempt++ {
        result, err := fn()
        if err == nil {
            return result, nil
        }
        
        lastErr = err
        
        if attempt == rp.MaxAttempts-1 {
            break
        }
        
        // Calculate delay
        actualDelay := delay
        if rp.Jitter {
            actualDelay = time.Duration(float64(delay) * (0.5 + rand.Float64()*0.5))
        }
        
        time.Sleep(actualDelay)
        
        // Calculate next delay
        delay = time.Duration(float64(delay) * rp.Multiplier)
        if delay > rp.MaxDelay {
            delay = rp.MaxDelay
        }
    }
    
    return nil, fmt.Errorf("max attempts reached: %v", lastErr)
}
```

### Timeout Policy
```go
type TimeoutPolicy struct {
    Timeout time.Duration
}

func (tp *TimeoutPolicy) Execute(fn func() (interface{}, error)) (interface{}, error) {
    ctx, cancel := context.WithTimeout(context.Background(), tp.Timeout)
    defer cancel()
    
    done := make(chan struct{})
    var result interface{}
    var err error
    
    go func() {
        defer close(done)
        result, err = fn()
    }()
    
    select {
    case <-done:
        return result, err
    case <-ctx.Done():
        return nil, fmt.Errorf("operation timed out after %v", tp.Timeout)
    }
}
```

## 🎯 Best Practices

### 1. Service Design
- **Single Responsibility**: Each service has one responsibility
- **Loose Coupling**: Services are loosely coupled
- **High Cohesion**: Related functionality is grouped together
- **Stateless**: Services should be stateless when possible

### 2. Communication
- **Synchronous**: Use for request-response patterns
- **Asynchronous**: Use for event-driven patterns
- **Message Queues**: Use for reliable message delivery
- **Event Streaming**: Use for real-time event processing

### 3. Data Management
- **Event Sourcing**: Store events instead of state
- **CQRS**: Separate command and query responsibilities
- **Replication**: Replicate data for availability
- **Partitioning**: Partition data for scalability

### 4. Fault Tolerance
- **Circuit Breakers**: Prevent cascading failures
- **Bulkheads**: Isolate resources
- **Retries**: Handle transient failures
- **Timeouts**: Prevent hanging operations

### 5. Monitoring
- **Health Checks**: Monitor service health
- **Metrics**: Collect performance metrics
- **Logging**: Log important events
- **Tracing**: Track request flows

## 🚨 Common Issues

### Consistency Issues
- **Eventual Consistency**: Data becomes consistent over time
- **Strong Consistency**: All nodes see same data immediately
- **Causal Consistency**: Preserve causal relationships
- **Weak Consistency**: No consistency guarantees

### Availability Issues
- **Single Points of Failure**: Eliminate SPOFs
- **Redundancy**: Multiple copies of critical components
- **Failover**: Automatic switching to backup systems
- **Load Balancing**: Distribute load across multiple instances

### Partition Tolerance Issues
- **Network Partitions**: Handle network splits
- **Split Brain**: Prevent multiple leaders
- **Consensus**: Agree on decisions despite partitions
- **Quorum**: Majority-based decision making

### Performance Issues
- **Latency**: Minimize communication latency
- **Throughput**: Maximize request processing rate
- **Scalability**: Handle increased load
- **Resource Usage**: Optimize resource consumption

## 📊 Monitoring and Observability

### Key Metrics
- **Request Rate**: Requests per second
- **Response Time**: Average response time
- **Error Rate**: Percentage of failed requests
- **Availability**: Uptime percentage
- **Throughput**: Data processed per second
- **Latency**: P50, P95, P99 latencies

### Health Checks
- **Liveness**: Is the service running?
- **Readiness**: Is the service ready to serve?
- **Startup**: Is the service starting up?
- **Custom**: Application-specific checks

### Distributed Tracing
- **Trace**: Complete request journey
- **Span**: Individual operation within trace
- **Context Propagation**: Pass trace context
- **Sampling**: Reduce trace volume

---

**Remember**: Distributed systems are complex but powerful. Use these patterns and practices to build reliable, scalable systems! 🚀
