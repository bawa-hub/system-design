# 🏗️ MICROSERVICES ARCHITECTURE PATTERNS

## 🎯 Overview
Architecture patterns are reusable solutions to common problems in microservices design. Understanding these patterns is crucial for building robust, scalable systems.

## 🔄 Service Communication Patterns

### 1. Synchronous Communication
**Definition**: Direct request-response communication between services.

**Characteristics**:
- Immediate response required
- Tight coupling
- Simple to implement
- Can cause cascade failures

**Examples**:
- REST APIs
- GraphQL
- gRPC
- HTTP/2

**Use Cases**:
- Real-time data requirements
- Simple request-response scenarios
- When immediate consistency is needed

**Pros**:
- Simple to understand
- Easy to debug
- Immediate feedback
- Familiar patterns

**Cons**:
- Tight coupling
- Cascade failure risk
- Network dependency
- Latency issues

### 2. Asynchronous Communication
**Definition**: Communication through message queues or event streams.

**Characteristics**:
- Loose coupling
- Eventual consistency
- Better resilience
- More complex to implement

**Examples**:
- Message Queues (RabbitMQ, Apache Kafka)
- Event Streaming
- Pub/Sub patterns
- Event-driven architecture

**Use Cases**:
- Event processing
- Background tasks
- Loose coupling requirements
- High throughput scenarios

**Pros**:
- Loose coupling
- Better resilience
- Scalability
- Decoupling

**Cons**:
- Eventual consistency
- Complexity
- Debugging challenges
- Message ordering issues

## 📊 Data Management Patterns

### 1. Database per Service
**Definition**: Each service has its own database, no shared data.

**Benefits**:
- Data isolation
- Independent scaling
- Technology diversity
- Team autonomy

**Challenges**:
- Data consistency
- Cross-service queries
- Transaction management

**Implementation**:
```go
// User Service - PostgreSQL
type UserService struct {
    db *sql.DB // PostgreSQL
}

// Order Service - MongoDB
type OrderService struct {
    db *mongo.Database // MongoDB
}
```

### 2. Saga Pattern
**Definition**: Managing distributed transactions across multiple services.

**Types**:
- **Choreography**: Each service knows what to do next
- **Orchestration**: Central coordinator manages the flow

**Example - E-commerce Order**:
1. Create order
2. Reserve inventory
3. Process payment
4. Send confirmation

**Choreography Example**:
```go
// Order Service
func (s *OrderService) CreateOrder(order Order) {
    // Create order
    s.db.CreateOrder(order)
    
    // Publish event
    s.eventBus.Publish("OrderCreated", order)
}

// Inventory Service
func (s *InventoryService) HandleOrderCreated(event OrderCreated) {
    // Reserve inventory
    s.reserveInventory(event.OrderID)
    
    // Publish event
    s.eventBus.Publish("InventoryReserved", event)
}
```

### 3. CQRS (Command Query Responsibility Segregation)
**Definition**: Separating read and write operations into different models.

**Benefits**:
- Independent scaling
- Different data models
- Performance optimization
- Simplified complex domains

**Implementation**:
```go
// Command Side (Write)
type CreateUserCommand struct {
    Email     string
    FirstName string
    LastName  string
}

type UserCommandHandler struct {
    db *sql.DB
}

func (h *UserCommandHandler) Handle(cmd CreateUserCommand) {
    // Write to database
    h.db.CreateUser(cmd)
    
    // Publish event
    h.eventBus.Publish("UserCreated", cmd)
}

// Query Side (Read)
type UserQuery struct {
    ID int
}

type UserQueryHandler struct {
    readDB *sql.DB // Read-optimized database
}

func (h *UserQueryHandler) Handle(query UserQuery) User {
    // Read from optimized database
    return h.readDB.GetUser(query.ID)
}
```

### 4. Event Sourcing
**Definition**: Storing state changes as events rather than current state.

**Benefits**:
- Complete audit trail
- Time travel debugging
- Event replay capability
- Better scalability

**Implementation**:
```go
type Event struct {
    ID        string
    Type      string
    Data      interface{}
    Timestamp time.Time
}

type EventStore struct {
    events []Event
}

func (es *EventStore) AppendEvent(event Event) {
    es.events = append(es.events, event)
}

func (es *EventStore) GetEvents(aggregateID string) []Event {
    var result []Event
    for _, event := range es.events {
        if event.AggregateID == aggregateID {
            result = append(result, event)
        }
    }
    return result
}

// User Aggregate
type User struct {
    ID        string
    Email     string
    FirstName string
    LastName  string
}

func (u *User) ApplyEvent(event Event) {
    switch event.Type {
    case "UserCreated":
        data := event.Data.(UserCreatedData)
        u.ID = data.ID
        u.Email = data.Email
        u.FirstName = data.FirstName
        u.LastName = data.LastName
    case "UserUpdated":
        data := event.Data.(UserUpdatedData)
        u.Email = data.Email
        u.FirstName = data.FirstName
        u.LastName = data.LastName
    }
}
```

## 🛡️ Resilience Patterns

### 1. Circuit Breaker
**Definition**: Prevents cascade failures by opening the circuit when a service is failing.

**States**:
- **Closed**: Normal operation
- **Open**: Service is failing
- **Half-open**: Testing recovery

**Implementation**:
```go
type CircuitBreaker struct {
    failureCount    int
    failureThreshold int
    timeout         time.Duration
    state          string
    lastFailureTime time.Time
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    if cb.state == "open" {
        if time.Since(cb.lastFailureTime) > cb.timeout {
            cb.state = "half-open"
        } else {
            return errors.New("circuit breaker is open")
        }
    }
    
    err := fn()
    if err != nil {
        cb.failureCount++
        cb.lastFailureTime = time.Now()
        
        if cb.failureCount >= cb.failureThreshold {
            cb.state = "open"
        }
        return err
    }
    
    cb.failureCount = 0
    cb.state = "closed"
    return nil
}
```

### 2. Retry Pattern
**Definition**: Automatically retry failed operations with exponential backoff.

**Implementation**:
```go
func Retry(fn func() error, maxRetries int, backoff time.Duration) error {
    var err error
    
    for i := 0; i < maxRetries; i++ {
        err = fn()
        if err == nil {
            return nil
        }
        
        if i < maxRetries-1 {
            time.Sleep(backoff * time.Duration(math.Pow(2, float64(i))))
        }
    }
    
    return err
}
```

### 3. Timeout Pattern
**Definition**: Setting timeouts to prevent indefinite waits.

**Implementation**:
```go
func CallWithTimeout(fn func() error, timeout time.Duration) error {
    done := make(chan error, 1)
    
    go func() {
        done <- fn()
    }()
    
    select {
    case err := <-done:
        return err
    case <-time.After(timeout):
        return errors.New("operation timed out")
    }
}
```

### 4. Bulkhead Pattern
**Definition**: Isolating resources to prevent failures from spreading.

**Implementation**:
```go
type Bulkhead struct {
    userPool    chan struct{}
    orderPool   chan struct{}
    paymentPool chan struct{}
}

func NewBulkhead() *Bulkhead {
    return &Bulkhead{
        userPool:    make(chan struct{}, 10),
        orderPool:   make(chan struct{}, 5),
        paymentPool: make(chan struct{}, 3),
    }
}

func (b *Bulkhead) ExecuteUserOperation(fn func() error) error {
    b.userPool <- struct{}{}
    defer func() { <-b.userPool }()
    
    return fn()
}
```

## 🔍 Service Discovery Patterns

### 1. Client-Side Discovery
**Definition**: Client queries service registry to find service instances.

**Benefits**:
- Simple implementation
- No additional network hop
- Client has full control

**Challenges**:
- Client complexity
- Service registry coupling
- Language-specific implementation

### 2. Server-Side Discovery
**Definition**: Load balancer queries service registry and routes requests.

**Benefits**:
- Client simplicity
- Centralized routing
- Language agnostic

**Challenges**:
- Additional network hop
- Load balancer as single point of failure
- More complex infrastructure

## 🚪 API Gateway Patterns

### 1. API Gateway
**Definition**: Single entry point for all client requests.

**Responsibilities**:
- Request routing
- Authentication and authorization
- Rate limiting
- Request/response transformation
- Logging and monitoring

**Implementation**:
```go
type APIGateway struct {
    routes map[string]http.Handler
    auth   AuthService
    rate   RateLimiter
}

func (gw *APIGateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Authentication
    if !gw.auth.Authenticate(r) {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    
    // Rate limiting
    if !gw.rate.Allow(r.RemoteAddr) {
        http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
        return
    }
    
    // Route request
    handler := gw.routes[r.URL.Path]
    if handler == nil {
        http.Error(w, "Not found", http.StatusNotFound)
        return
    }
    
    handler.ServeHTTP(w, r)
}
```

### 2. Backend for Frontend (BFF)
**Definition**: Separate API gateway for each client type.

**Benefits**:
- Client-specific optimization
- Reduced client complexity
- Better performance
- Independent evolution

**Examples**:
- Mobile BFF
- Web BFF
- Admin BFF

## 🔐 Security Patterns

### 1. Zero Trust Architecture
**Definition**: Never trust, always verify.

**Principles**:
- Verify explicitly
- Use least privilege access
- Assume breach

### 2. mTLS (Mutual TLS)
**Definition**: Both client and server authenticate each other.

**Benefits**:
- Strong authentication
- Encrypted communication
- Service-to-service security

### 3. JWT (JSON Web Token)
**Definition**: Stateless authentication token.

**Structure**:
- Header: Algorithm and token type
- Payload: Claims and data
- Signature: Verification

## 📊 Monitoring Patterns

### 1. Three Pillars of Observability
- **Metrics**: Numerical data over time
- **Logs**: Discrete events
- **Traces**: Request flow through services

### 2. Distributed Tracing
**Definition**: Tracking requests across multiple services.

**Benefits**:
- Request flow visualization
- Performance bottleneck identification
- Error root cause analysis

## 🎯 Pattern Selection Guidelines

### Choose Synchronous When:
- Immediate response required
- Simple request-response
- Real-time data needs
- Tight coupling acceptable

### Choose Asynchronous When:
- Loose coupling desired
- Eventual consistency acceptable
- High throughput needed
- Background processing

### Choose Database per Service When:
- Data isolation needed
- Independent scaling required
- Technology diversity desired
- Team autonomy important

### Choose Saga When:
- Distributed transactions needed
- ACID not possible
- Eventual consistency acceptable
- Complex business workflows

## 🔗 Next Steps

Now that you understand architecture patterns, we'll move to:
1. **When to Use What** - Decision-making framework
2. **Implementation** - Building your first microservice
3. **Practice** - Hands-on exercises

**Ready to continue? Let's learn when to use each pattern!** 🚀
