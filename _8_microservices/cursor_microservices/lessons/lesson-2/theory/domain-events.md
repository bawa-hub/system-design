# 📡 DOMAIN EVENTS - EVENT-DRIVEN ARCHITECTURE

## 🎯 What are Domain Events?

### Definition
Domain events represent something significant that happened in the domain. They are immutable objects that capture what occurred, when it occurred, and what was involved.

### Key Characteristics
- **Immutable** - cannot be changed after creation
- **Represent past occurrences** - something that already happened
- **Can trigger side effects** - other parts of the system can react
- **Enable loose coupling** - services don't need to know about each other
- **Support eventual consistency** - different parts can update at different times

## 🏗️ Domain Event Design

### 1. Event Structure
```go
type DomainEvent interface {
    EventType() string
    OccurredAt() time.Time
    AggregateID() string
    EventVersion() int
}

type UserRegistered struct {
    UserID    UserID
    Email     Email
    Name      Name
    OccurredAt time.Time
    Version   int
}

func (e UserRegistered) EventType() string {
    return "user.registered"
}

func (e UserRegistered) OccurredAt() time.Time {
    return e.OccurredAt
}

func (e UserRegistered) AggregateID() string {
    return string(e.UserID)
}

func (e UserRegistered) EventVersion() int {
    return e.Version
}
```

### 2. Event Naming Conventions
**Past tense verbs** - represent something that already happened:
- `UserRegistered`
- `OrderPlaced`
- `PaymentProcessed`
- `InventoryReserved`
- `ProductCreated`

**Domain context** - include the domain context:
- `user.registered`
- `order.placed`
- `payment.processed`
- `inventory.reserved`
- `product.created`

### 3. Event Data
**Include relevant data** - what other services need to know:
```go
type OrderPlaced struct {
    OrderID     OrderID
    CustomerID  CustomerID
    Items       []OrderItem
    TotalAmount Money
    OccurredAt  time.Time
    Version     int
}

type PaymentProcessed struct {
    PaymentID   PaymentID
    OrderID     OrderID
    Amount      Money
    Status      PaymentStatus
    OccurredAt  time.Time
    Version     int
}
```

## 🔄 Event Publishing

### 1. Aggregate Root Responsibility
**Aggregate roots publish events** when significant things happen:

```go
type Order struct {
    ID          OrderID
    CustomerID  CustomerID
    Items       []OrderItem
    Status      OrderStatus
    TotalAmount Money
    CreatedAt   time.Time
    UpdatedAt   time.Time
    
    // Events that occurred but not yet published
    events []DomainEvent
}

func (o *Order) Confirm() error {
    if o.Status != OrderStatusDraft {
        return errors.New("only draft orders can be confirmed")
    }
    
    if len(o.Items) == 0 {
        return errors.New("cannot confirm empty order")
    }
    
    o.Status = OrderStatusConfirmed
    o.UpdatedAt = time.Now()
    
    // Publish domain event
    event := OrderConfirmed{
        OrderID:     o.ID,
        CustomerID:  o.CustomerID,
        Items:       o.Items,
        TotalAmount: o.TotalAmount,
        OccurredAt:  time.Now(),
        Version:     1,
    }
    
    o.events = append(o.events, event)
    
    return nil
}

func (o *Order) GetEvents() []DomainEvent {
    return o.events
}

func (o *Order) ClearEvents() {
    o.events = []DomainEvent{}
}
```

### 2. Event Publisher
**Service layer publishes events** after saving aggregates:

```go
type EventPublisher interface {
    Publish(events []DomainEvent) error
}

type OrderService struct {
    orderRepository OrderRepository
    eventPublisher  EventPublisher
}

func (s *OrderService) ConfirmOrder(orderID OrderID) error {
    // Load aggregate
    order, err := s.orderRepository.FindByID(orderID)
    if err != nil {
        return err
    }
    
    // Execute business logic
    if err := order.Confirm(); err != nil {
        return err
    }
    
    // Save aggregate
    if err := s.orderRepository.Save(order); err != nil {
        return err
    }
    
    // Publish events
    events := order.GetEvents()
    if err := s.eventPublisher.Publish(events); err != nil {
        return err
    }
    
    // Clear events
    order.ClearEvents()
    
    return nil
}
```

## 🎯 Event Handlers

### 1. Event Handler Interface
```go
type EventHandler interface {
    Handle(event DomainEvent) error
    CanHandle(eventType string) bool
}

type UserRegisteredHandler struct {
    emailService EmailService
    analyticsService AnalyticsService
}

func (h *UserRegisteredHandler) CanHandle(eventType string) bool {
    return eventType == "user.registered"
}

func (h *UserRegisteredHandler) Handle(event DomainEvent) error {
    userEvent, ok := event.(UserRegistered)
    if !ok {
        return errors.New("invalid event type")
    }
    
    // Send welcome email
    if err := h.emailService.SendWelcomeEmail(userEvent.Email, userEvent.Name); err != nil {
        return err
    }
    
    // Track user registration
    if err := h.analyticsService.TrackUserRegistration(userEvent.UserID); err != nil {
        return err
    }
    
    return nil
}
```

### 2. Event Dispatcher
```go
type EventDispatcher struct {
    handlers map[string][]EventHandler
}

func NewEventDispatcher() *EventDispatcher {
    return &EventDispatcher{
        handlers: make(map[string][]EventHandler),
    }
}

func (d *EventDispatcher) RegisterHandler(eventType string, handler EventHandler) {
    d.handlers[eventType] = append(d.handlers[eventType], handler)
}

func (d *EventDispatcher) Dispatch(event DomainEvent) error {
    eventType := event.EventType()
    handlers := d.handlers[eventType]
    
    for _, handler := range handlers {
        if handler.CanHandle(eventType) {
            if err := handler.Handle(event); err != nil {
                return err
            }
        }
    }
    
    return nil
}
```

## 🏗️ Event Sourcing

### 1. Event Store
**Store events** instead of current state:

```go
type EventStore interface {
    SaveEvents(aggregateID string, events []DomainEvent, expectedVersion int) error
    GetEvents(aggregateID string) ([]DomainEvent, error)
    GetEventsFromVersion(aggregateID string, fromVersion int) ([]DomainEvent, error)
}

type EventStoreImpl struct {
    db *sql.DB
}

func (s *EventStoreImpl) SaveEvents(aggregateID string, events []DomainEvent, expectedVersion int) error {
    tx, err := s.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    // Check current version
    var currentVersion int
    err = tx.QueryRow("SELECT version FROM aggregates WHERE id = $1", aggregateID).Scan(&currentVersion)
    if err != nil {
        if err == sql.ErrNoRows {
            currentVersion = 0
        } else {
            return err
        }
    }
    
    if currentVersion != expectedVersion {
        return errors.New("concurrent modification detected")
    }
    
    // Save events
    for i, event := range events {
        version := expectedVersion + i + 1
        eventData, err := json.Marshal(event)
        if err != nil {
            return err
        }
        
        _, err = tx.Exec(`
            INSERT INTO events (aggregate_id, event_type, event_data, version, occurred_at)
            VALUES ($1, $2, $3, $4, $5)
        `, aggregateID, event.EventType(), eventData, version, event.OccurredAt())
        if err != nil {
            return err
        }
    }
    
    // Update aggregate version
    _, err = tx.Exec(`
        INSERT INTO aggregates (id, version) VALUES ($1, $2)
        ON CONFLICT (id) DO UPDATE SET version = $2
    `, aggregateID, expectedVersion+len(events))
    if err != nil {
        return err
    }
    
    return tx.Commit()
}

func (s *EventStoreImpl) GetEvents(aggregateID string) ([]DomainEvent, error) {
    rows, err := s.db.Query(`
        SELECT event_type, event_data, version, occurred_at
        FROM events
        WHERE aggregate_id = $1
        ORDER BY version
    `, aggregateID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var events []DomainEvent
    for rows.Next() {
        var eventType string
        var eventData []byte
        var version int
        var occurredAt time.Time
        
        err := rows.Scan(&eventType, &eventData, &version, &occurredAt)
        if err != nil {
            return nil, err
        }
        
        event, err := s.deserializeEvent(eventType, eventData, version, occurredAt)
        if err != nil {
            return nil, err
        }
        
        events = append(events, event)
    }
    
    return events, nil
}
```

### 2. Aggregate Reconstruction
**Rebuild aggregates** from events:

```go
type Order struct {
    ID          OrderID
    CustomerID  CustomerID
    Items       []OrderItem
    Status      OrderStatus
    TotalAmount Money
    CreatedAt   time.Time
    UpdatedAt   time.Time
    Version     int
}

func (o *Order) ApplyEvent(event DomainEvent) error {
    switch e := event.(type) {
    case OrderCreated:
        o.ID = e.OrderID
        o.CustomerID = e.CustomerID
        o.Status = OrderStatusDraft
        o.CreatedAt = e.OccurredAt
        o.UpdatedAt = e.OccurredAt
        o.Version = e.Version
        
    case OrderItemAdded:
        item := OrderItem{
            ProductID: e.ProductID,
            Quantity:  e.Quantity,
            Price:     e.Price,
        }
        o.Items = append(o.Items, item)
        o.recalculateTotal()
        o.UpdatedAt = e.OccurredAt
        o.Version = e.Version
        
    case OrderConfirmed:
        o.Status = OrderStatusConfirmed
        o.UpdatedAt = e.OccurredAt
        o.Version = e.Version
        
    case OrderShipped:
        o.Status = OrderStatusShipped
        o.UpdatedAt = e.OccurredAt
        o.Version = e.Version
        
    case OrderDelivered:
        o.Status = OrderStatusDelivered
        o.UpdatedAt = e.OccurredAt
        o.Version = e.Version
        
    default:
        return errors.New("unknown event type")
    }
    
    return nil
}

func RebuildOrderFromEvents(events []DomainEvent) (*Order, error) {
    order := &Order{}
    
    for _, event := range events {
        if err := order.ApplyEvent(event); err != nil {
            return nil, err
        }
    }
    
    return order, nil
}
```

## 🔄 Event-Driven Architecture

### 1. Event Choreography
**Services react to events** without central coordination:

```go
// Order Service publishes OrderPlaced event
type OrderService struct {
    eventPublisher EventPublisher
}

func (s *OrderService) PlaceOrder(order Order) error {
    // Business logic
    if err := order.Place(); err != nil {
        return err
    }
    
    // Publish event
    event := OrderPlaced{
        OrderID:     order.ID,
        CustomerID:  order.CustomerID,
        Items:       order.Items,
        TotalAmount: order.TotalAmount,
        OccurredAt:  time.Now(),
        Version:     1,
    }
    
    return s.eventPublisher.Publish([]DomainEvent{event})
}

// Inventory Service reacts to OrderPlaced event
type InventoryService struct {
    inventoryRepository InventoryRepository
}

func (s *InventoryService) HandleOrderPlaced(event OrderPlaced) error {
    for _, item := range event.Items {
        if err := s.inventoryRepository.Reserve(item.ProductID, item.Quantity); err != nil {
            return err
        }
    }
    
    return nil
}

// Payment Service reacts to OrderPlaced event
type PaymentService struct {
    paymentGateway PaymentGateway
}

func (s *PaymentService) HandleOrderPlaced(event OrderPlaced) error {
    payment := Payment{
        OrderID: event.OrderID,
        Amount:  event.TotalAmount,
        Status:  PaymentStatusPending,
    }
    
    return s.paymentGateway.ProcessPayment(payment)
}
```

### 2. Event Orchestration
**Central coordinator** manages the flow:

```go
type OrderOrchestrator struct {
    orderService     OrderService
    inventoryService InventoryService
    paymentService   PaymentService
    eventPublisher   EventPublisher
}

func (o *OrderOrchestrator) ProcessOrder(orderID OrderID) error {
    // 1. Place order
    order, err := o.orderService.GetOrder(orderID)
    if err != nil {
        return err
    }
    
    if err := o.orderService.PlaceOrder(order); err != nil {
        return err
    }
    
    // 2. Reserve inventory
    if err := o.inventoryService.ReserveInventory(order.Items); err != nil {
        // Compensate: cancel order
        o.orderService.CancelOrder(orderID)
        return err
    }
    
    // 3. Process payment
    if err := o.paymentService.ProcessPayment(order); err != nil {
        // Compensate: release inventory and cancel order
        o.inventoryService.ReleaseInventory(order.Items)
        o.orderService.CancelOrder(orderID)
        return err
    }
    
    // 4. Confirm order
    return o.orderService.ConfirmOrder(orderID)
}
```

## 🎯 Event Design Patterns

### 1. Event Sourcing
**Store events** instead of current state:
- Complete audit trail
- Time travel debugging
- Event replay capability
- Better scalability

### 2. CQRS (Command Query Responsibility Segregation)
**Separate read and write models**:
- Commands update aggregates
- Queries read from projections
- Events update projections
- Independent scaling

### 3. Saga Pattern
**Manage distributed transactions**:
- Choreography: services react to events
- Orchestration: central coordinator
- Compensating actions for failures
- Eventual consistency

### 4. Event Streaming
**Stream events** for real-time processing:
- Apache Kafka
- Event ordering and partitioning
- Replay and replay capability
- High throughput

## ⚠️ Event Design Pitfalls

### 1. Event Explosion
**Too many events** can overwhelm the system:
- Focus on significant business events
- Batch related events
- Use event filtering
- Consider event aggregation

### 2. Event Ordering
**Events must be processed in order**:
- Use event versioning
- Implement idempotency
- Handle out-of-order events
- Use event sourcing

### 3. Event Schema Evolution
**Events change over time**:
- Use versioning
- Implement backward compatibility
- Plan migration strategies
- Use schema registry

### 4. Event Reliability
**Events must be delivered reliably**:
- Implement retry mechanisms
- Use dead letter queues
- Monitor event processing
- Handle failures gracefully

## 🔗 Next Steps

Now that you understand domain events, we'll move to:
1. **Service Design** - Applying DDD to microservices
2. **Implementation** - Building DDD-based services
3. **Practice** - Hands-on exercises
4. **Real-world Examples** - E-commerce system

**Ready to continue? Let's learn about service design!** 🚀
