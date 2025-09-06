# Service Communication Patterns

## Overview

This document explains how different microservices communicate with each other in the Nakku Quick Commerce architecture, including synchronous and asynchronous communication patterns, service discovery, and data consistency strategies.

## Communication Patterns

### 1. **Synchronous Communication (HTTP/REST)**

#### API Gateway to Microservices
```
Client Request → API Gateway → Microservice → Response
```

**Example Flow:**
```
1. Client: POST /api/v1/users/register
2. API Gateway: Routes to User Service
3. User Service: Processes registration
4. User Service: Returns response
5. API Gateway: Returns response to client
```

#### Service-to-Service Communication
```
Service A → HTTP Request → Service B → Response
```

**Example: Order Service calling User Service**
```go
// Order Service calling User Service
func (s *OrderService) CreateOrder(ctx context.Context, userID string, items []OrderItem) (*Order, error) {
    // Validate user exists
    user, err := s.userClient.GetUser(ctx, userID)
    if err != nil {
        return nil, fmt.Errorf("user not found: %w", err)
    }
    
    // Create order
    order := &Order{
        UserID: user.ID,
        Items:  items,
        Status: "pending",
    }
    
    return s.orderRepo.Create(ctx, order)
}
```

### 2. **Asynchronous Communication (Kafka Events)**

#### Event-Driven Architecture
```
Service A → Publish Event → Kafka → Service B (Consumer)
```

**Example: Order Creation Flow**
```
1. Order Service: Creates order
2. Order Service: Publishes "order.created" event
3. Kafka: Distributes event to subscribers
4. Inventory Service: Consumes event, reserves items
5. Payment Service: Consumes event, processes payment
6. Notification Service: Consumes event, sends confirmation
```

#### Event Publishing
```go
// Order Service publishing event
func (s *OrderService) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*Order, error) {
    // Create order
    order, err := s.orderRepo.Create(ctx, order)
    if err != nil {
        return nil, err
    }
    
    // Publish event
    event := OrderCreatedEvent{
        OrderID:    order.ID,
        UserID:     order.UserID,
        Items:      order.Items,
        TotalAmount: order.TotalAmount,
        Timestamp:  time.Now(),
    }
    
    err = s.producer.PublishMessage(ctx, "order-events", fmt.Sprintf("order-%d", order.ID), event)
    if err != nil {
        s.logger.WithError(err).Error("Failed to publish order created event")
    }
    
    return order, nil
}
```

#### Event Consumption
```go
// Inventory Service consuming events
func (s *InventoryService) StartEventConsumer(ctx context.Context) error {
    consumer := utils.NewKafkaConsumer(s.kafkaBrokers, "order-events", "inventory-group", s.logger)
    
    return consumer.ConsumeMessages(ctx, func(msg utils.KafkaMessage) error {
        switch msg.Topic {
        case "order-events":
            return s.handleOrderEvent(ctx, msg)
        case "inventory-events":
            return s.handleInventoryEvent(ctx, msg)
        default:
            s.logger.WithField("topic", msg.Topic).Warn("Unknown event topic")
            return nil
        }
    })
}

func (s *InventoryService) handleOrderEvent(ctx context.Context, msg utils.KafkaMessage) error {
    var event OrderCreatedEvent
    if err := json.Unmarshal([]byte(msg.Value.(string)), &event); err != nil {
        return err
    }
    
    // Reserve inventory
    for _, item := range event.Items {
        if err := s.ReserveInventory(ctx, item.ProductID, item.Quantity); err != nil {
            return err
        }
    }
    
    return nil
}
```

## Service Discovery

### 1. **Static Configuration**
```yaml
# API Gateway configuration
services:
  user-service:
    url: "http://user-service:8081"
  product-service:
    url: "http://product-service:8082"
  order-service:
    url: "http://order-service:8084"
```

### 2. **Environment Variables**
```bash
# Service URLs
USER_SERVICE_URL=http://user-service:8081
PRODUCT_SERVICE_URL=http://product-service:8082
ORDER_SERVICE_URL=http://order-service:8084
```

### 3. **Service Registry (Consul)**
```go
// Service discovery with Consul
type ServiceRegistry struct {
    consul *consul.Client
}

func (sr *ServiceRegistry) GetServiceURL(serviceName string) (string, error) {
    services, _, err := sr.consul.Health().Service(serviceName, "", true, nil)
    if err != nil {
        return "", err
    }
    
    if len(services) == 0 {
        return "", fmt.Errorf("no healthy instances of %s", serviceName)
    }
    
    service := services[0]
    return fmt.Sprintf("http://%s:%d", service.Service.Address, service.Service.Port), nil
}
```

## Data Consistency Patterns

### 1. **Eventual Consistency**
```
Service A → Event → Service B (Eventually Consistent)
```

**Example: User Profile Update**
```
1. User Service: Updates user profile
2. User Service: Publishes "user.updated" event
3. Analytics Service: Consumes event, updates analytics data
4. Notification Service: Consumes event, updates notification preferences
```

### 2. **Saga Pattern**
```
Order Service → Payment Service → Inventory Service → Delivery Service
```

**Example: Order Processing Saga**
```go
// Saga coordinator
type OrderSaga struct {
    orderService    *OrderService
    paymentService  *PaymentService
    inventoryService *InventoryService
    deliveryService *DeliveryService
}

func (s *OrderSaga) ProcessOrder(ctx context.Context, orderID string) error {
    // Step 1: Reserve inventory
    if err := s.inventoryService.ReserveInventory(ctx, orderID); err != nil {
        return s.compensateReserveInventory(ctx, orderID)
    }
    
    // Step 2: Process payment
    if err := s.paymentService.ProcessPayment(ctx, orderID); err != nil {
        return s.compensatePayment(ctx, orderID)
    }
    
    // Step 3: Schedule delivery
    if err := s.deliveryService.ScheduleDelivery(ctx, orderID); err != nil {
        return s.compensateDelivery(ctx, orderID)
    }
    
    // Step 4: Complete order
    return s.orderService.CompleteOrder(ctx, orderID)
}

func (s *OrderSaga) compensateReserveInventory(ctx context.Context, orderID string) error {
    // Release reserved inventory
    return s.inventoryService.ReleaseInventory(ctx, orderID)
}
```

### 3. **CQRS (Command Query Responsibility Segregation)**
```
Write Model → Event → Read Model
```

**Example: Product Catalog**
```go
// Write Model (Product Service)
func (s *ProductService) CreateProduct(ctx context.Context, req *CreateProductRequest) (*Product, error) {
    product, err := s.productRepo.Create(ctx, product)
    if err != nil {
        return nil, err
    }
    
    // Publish event
    event := ProductCreatedEvent{
        ProductID: product.ID,
        Name:      product.Name,
        Price:     product.Price,
        Category:  product.Category,
    }
    
    s.producer.PublishMessage(ctx, "product-events", fmt.Sprintf("product-%d", product.ID), event)
    return product, nil
}

// Read Model (Search Service)
func (s *SearchService) handleProductEvent(ctx context.Context, msg utils.KafkaMessage) error {
    var event ProductCreatedEvent
    if err := json.Unmarshal([]byte(msg.Value.(string)), &event); err != nil {
        return err
    }
    
    // Update search index
    return s.searchIndex.UpdateProduct(event)
}
```

## Error Handling and Resilience

### 1. **Circuit Breaker Pattern**
```go
type CircuitBreaker struct {
    maxFailures int
    timeout     time.Duration
    failures    int
    lastFailure time.Time
    state       string // "closed", "open", "half-open"
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    if cb.state == "open" {
        if time.Since(cb.lastFailure) > cb.timeout {
            cb.state = "half-open"
        } else {
            return fmt.Errorf("circuit breaker is open")
        }
    }
    
    err := fn()
    if err != nil {
        cb.failures++
        cb.lastFailure = time.Now()
        if cb.failures >= cb.maxFailures {
            cb.state = "open"
        }
        return err
    }
    
    cb.failures = 0
    cb.state = "closed"
    return nil
}
```

### 2. **Retry Pattern**
```go
func (s *OrderService) callPaymentService(ctx context.Context, paymentReq *PaymentRequest) error {
    maxRetries := 3
    retryDelay := time.Second
    
    for i := 0; i < maxRetries; i++ {
        err := s.paymentClient.ProcessPayment(ctx, paymentReq)
        if err == nil {
            return nil
        }
        
        if i < maxRetries-1 {
            time.Sleep(retryDelay)
            retryDelay *= 2 // Exponential backoff
        }
    }
    
    return fmt.Errorf("payment service failed after %d retries", maxRetries)
}
```

### 3. **Timeout Pattern**
```go
func (s *OrderService) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*Order, error) {
    // Set timeout for the entire operation
    ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
    defer cancel()
    
    // Create order with timeout
    order, err := s.orderRepo.Create(ctx, order)
    if err != nil {
        return nil, err
    }
    
    // Process payment with timeout
    paymentCtx, paymentCancel := context.WithTimeout(ctx, 10*time.Second)
    defer paymentCancel()
    
    err = s.paymentService.ProcessPayment(paymentCtx, order.ID)
    if err != nil {
        return nil, err
    }
    
    return order, nil
}
```

## Message Queue Patterns

### 1. **Request-Reply Pattern**
```
Service A → Request Message → Kafka → Service B → Reply Message → Kafka → Service A
```

### 2. **Publish-Subscribe Pattern**
```
Service A → Event → Kafka → Multiple Services (Subscribers)
```

### 3. **Message Routing**
```go
// Message router
type MessageRouter struct {
    handlers map[string]func(utils.KafkaMessage) error
}

func (mr *MessageRouter) Route(msg utils.KafkaMessage) error {
    handler, exists := mr.handlers[msg.Topic]
    if !exists {
        return fmt.Errorf("no handler for topic %s", msg.Topic)
    }
    
    return handler(msg)
}

// Register handlers
router.RegisterHandler("user-events", userService.HandleUserEvent)
router.RegisterHandler("order-events", orderService.HandleOrderEvent)
router.RegisterHandler("payment-events", paymentService.HandlePaymentEvent)
```

## Performance Optimization

### 1. **Connection Pooling**
```go
// HTTP client with connection pooling
type ServiceClient struct {
    client *http.Client
}

func NewServiceClient() *ServiceClient {
    transport := &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    }
    
    return &ServiceClient{
        client: &http.Client{
            Transport: transport,
            Timeout:   30 * time.Second,
        },
    }
}
```

### 2. **Caching**
```go
// Redis caching
type CachedUserService struct {
    userService *UserService
    redis       *redis.Client
    cacheTTL    time.Duration
}

func (s *CachedUserService) GetUser(ctx context.Context, userID string) (*User, error) {
    // Try cache first
    cached, err := s.redis.Get(ctx, fmt.Sprintf("user:%s", userID)).Result()
    if err == nil {
        var user User
        if err := json.Unmarshal([]byte(cached), &user); err == nil {
            return &user, nil
        }
    }
    
    // Get from service
    user, err := s.userService.GetUser(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    // Cache the result
    userJSON, _ := json.Marshal(user)
    s.redis.Set(ctx, fmt.Sprintf("user:%s", userID), userJSON, s.cacheTTL)
    
    return user, nil
}
```

### 3. **Batch Processing**
```go
// Batch event processing
type BatchProcessor struct {
    batchSize    int
    batchTimeout time.Duration
    events       []Event
    mutex        sync.Mutex
}

func (bp *BatchProcessor) AddEvent(event Event) {
    bp.mutex.Lock()
    defer bp.mutex.Unlock()
    
    bp.events = append(bp.events, event)
    
    if len(bp.events) >= bp.batchSize {
        go bp.processBatch()
    }
}

func (bp *BatchProcessor) processBatch() {
    bp.mutex.Lock()
    events := bp.events
    bp.events = nil
    bp.mutex.Unlock()
    
    // Process batch
    for _, event := range events {
        bp.processEvent(event)
    }
}
```

## Monitoring and Observability

### 1. **Distributed Tracing**
```go
// OpenTelemetry tracing
func (s *OrderService) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*Order, error) {
    tracer := otel.Tracer("order-service")
    ctx, span := tracer.Start(ctx, "CreateOrder")
    defer span.End()
    
    // Add attributes
    span.SetAttributes(
        attribute.String("user.id", req.UserID),
        attribute.Int("items.count", len(req.Items)),
    )
    
    // Create order
    order, err := s.orderRepo.Create(ctx, order)
    if err != nil {
        span.RecordError(err)
        return nil, err
    }
    
    span.SetAttributes(attribute.String("order.id", order.ID))
    return order, nil
}
```

### 2. **Metrics Collection**
```go
// Prometheus metrics
var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint", "status"},
    )
    
    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "http_request_duration_seconds",
            Help: "HTTP request duration in seconds",
        },
        []string{"method", "endpoint"},
    )
)

func init() {
    prometheus.MustRegister(httpRequestsTotal)
    prometheus.MustRegister(httpRequestDuration)
}
```

This communication pattern ensures that your microservices can work together effectively while maintaining loose coupling, high availability, and data consistency across the distributed system.
