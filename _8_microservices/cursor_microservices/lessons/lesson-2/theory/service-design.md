# 🏗️ SERVICE DESIGN - APPLYING DDD TO MICROSERVICES

## 🎯 From DDD to Microservices

### The Connection
Domain-Driven Design provides the perfect foundation for microservices architecture:
- **Bounded contexts** become microservices
- **Aggregates** become service boundaries
- **Domain events** enable service communication
- **Ubiquitous language** guides API design

### Service Design Principles

#### 1. One Bounded Context = One Microservice
**Each microservice represents a single bounded context**:
- Clear ownership and responsibility
- Independent domain model
- Own data and business rules
- Clear team boundaries

**Example**:
```
User Management Context → User Service
Product Catalog Context → Product Service
Order Processing Context → Order Service
Payment Context → Payment Service
```

#### 2. Aggregate Root = Service Boundary
**Each service manages one or more aggregates**:
- Single responsibility per service
- Clear data ownership
- Consistent business rules
- Transaction boundaries

**Example**:
```
User Service:
- User aggregate
- UserPreferences aggregate
- UserSession aggregate

Order Service:
- Order aggregate
- OrderItem aggregate
- OrderStatus aggregate
```

#### 3. Domain Events = Service Communication
**Services communicate through domain events**:
- Loose coupling
- Eventual consistency
- Scalable architecture
- Independent evolution

**Example**:
```
User Service publishes: UserRegistered
Order Service subscribes: CreateOrderForNewUser
Product Service publishes: ProductUpdated
Order Service subscribes: UpdateOrderItems
```

## 🏗️ Service Architecture Patterns

### 1. Layered Architecture
**Organize code in layers** within each service:

```
┌─────────────────────────────────────┐
│           Presentation Layer        │
│         (HTTP Handlers, APIs)       │
└─────────────────────────────────────┘
┌─────────────────────────────────────┐
│           Application Layer         │
│        (Use Cases, Services)        │
└─────────────────────────────────────┘
┌─────────────────────────────────────┐
│            Domain Layer             │
│    (Aggregates, Entities, Events)   │
└─────────────────────────────────────┘
┌─────────────────────────────────────┐
│          Infrastructure Layer       │
│    (Database, External APIs)        │
└─────────────────────────────────────┘
```

**Implementation**:
```go
// Presentation Layer
type UserHandler struct {
    userService UserService
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    user, err := h.userService.CreateUser(req.Email, req.Name)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

// Application Layer
type UserService struct {
    userRepository UserRepository
    eventPublisher EventPublisher
}

func (s *UserService) CreateUser(email Email, name Name) (*User, error) {
    // Business logic
    user := NewUser(email, name)
    
    // Save to repository
    if err := s.userRepository.Save(user); err != nil {
        return nil, err
    }
    
    // Publish events
    events := user.GetEvents()
    if err := s.eventPublisher.Publish(events); err != nil {
        return nil, err
    }
    
    user.ClearEvents()
    return user, nil
}

// Domain Layer
type User struct {
    ID        UserID
    Email     Email
    Name      Name
    Status    UserStatus
    CreatedAt time.Time
    UpdatedAt time.Time
    events    []DomainEvent
}

func NewUser(email Email, name Name) *User {
    now := time.Now()
    user := &User{
        ID:        NewUserID(),
        Email:     email,
        Name:      name,
        Status:    UserStatusPending,
        CreatedAt: now,
        UpdatedAt: now,
        events:    []DomainEvent{},
    }
    
    // Publish domain event
    event := UserCreated{
        UserID:    user.ID,
        Email:     user.Email,
        Name:      user.Name,
        OccurredAt: now,
        Version:   1,
    }
    user.events = append(user.events, event)
    
    return user
}

// Infrastructure Layer
type UserRepositoryImpl struct {
    db *sql.DB
}

func (r *UserRepositoryImpl) Save(user *User) error {
    query := `
        INSERT INTO users (id, email, name, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        ON CONFLICT (id) DO UPDATE SET
        email = $2, name = $3, status = $4, updated_at = $6
    `
    
    _, err := r.db.Exec(query, user.ID, user.Email, user.Name, user.Status, user.CreatedAt, user.UpdatedAt)
    return err
}
```

### 2. Hexagonal Architecture (Ports and Adapters)
**Decouple business logic from external concerns**:

```
┌─────────────────────────────────────┐
│           Business Logic            │
│    (Domain + Application Layer)     │
└─────────────────────────────────────┘
         ▲                    ▲
         │                    │
    ┌─────────┐          ┌─────────┐
    │  Ports  │          │ Adapters│
    │ (Interfaces)       │ (Impls) │
    └─────────┘          └─────────┘
         ▲                    ▲
         │                    │
┌─────────────────────────────────────┐
│        External Systems             │
│    (Database, APIs, UI)             │
└─────────────────────────────────────┘
```

**Implementation**:
```go
// Ports (Interfaces)
type UserRepository interface {
    Save(user *User) error
    FindByID(id UserID) (*User, error)
    FindByEmail(email Email) (*User, error)
}

type EventPublisher interface {
    Publish(events []DomainEvent) error
}

type EmailService interface {
    SendWelcomeEmail(email Email, name Name) error
}

// Adapters (Implementations)
type UserRepositoryImpl struct {
    db *sql.DB
}

func (r *UserRepositoryImpl) Save(user *User) error {
    // Database implementation
}

type EventPublisherImpl struct {
    messageBroker MessageBroker
}

func (p *EventPublisherImpl) Publish(events []DomainEvent) error {
    // Message broker implementation
}

type EmailServiceImpl struct {
    smtpClient SMTPClient
}

func (s *EmailServiceImpl) SendWelcomeEmail(email Email, name Name) error {
    // SMTP implementation
}
```

### 3. CQRS (Command Query Responsibility Segregation)
**Separate read and write operations**:

```
┌─────────────────────────────────────┐
│           Commands                  │
│    (Write Operations)               │
└─────────────────────────────────────┘
         ▲                    ▲
         │                    │
    ┌─────────┐          ┌─────────┐
    │ Command │          │  Event  │
    │ Handler │          │ Publisher│
    └─────────┘          └─────────┘
         ▲                    ▲
         │                    │
┌─────────────────────────────────────┐
│           Queries                   │
│    (Read Operations)                │
└─────────────────────────────────────┘
```

**Implementation**:
```go
// Command Side
type CreateUserCommand struct {
    Email Email
    Name  Name
}

type CreateUserHandler struct {
    userRepository UserRepository
    eventPublisher EventPublisher
}

func (h *CreateUserHandler) Handle(cmd CreateUserCommand) error {
    user := NewUser(cmd.Email, cmd.Name)
    
    if err := h.userRepository.Save(user); err != nil {
        return err
    }
    
    events := user.GetEvents()
    return h.eventPublisher.Publish(events)
}

// Query Side
type UserQuery struct {
    ID UserID
}

type UserQueryHandler struct {
    userReadModel UserReadModel
}

func (h *UserQueryHandler) Handle(query UserQuery) (*UserView, error) {
    return h.userReadModel.FindByID(query.ID)
}

// Read Model
type UserView struct {
    ID        string
    Email     string
    Name      string
    Status    string
    CreatedAt time.Time
}

type UserReadModel interface {
    FindByID(id UserID) (*UserView, error)
    FindByEmail(email Email) (*UserView, error)
    FindAll() ([]UserView, error)
}
```

## 🔄 Service Communication Patterns

### 1. Synchronous Communication
**Direct request-response** between services:

```go
// Order Service calls User Service
type OrderService struct {
    userService UserService
}

func (s *OrderService) CreateOrder(customerID CustomerID, items []OrderItem) (*Order, error) {
    // Validate customer exists
    user, err := s.userService.GetUser(customerID)
    if err != nil {
        return nil, err
    }
    
    if !user.IsActive() {
        return nil, errors.New("customer is not active")
    }
    
    // Create order
    order := NewOrder(customerID, items)
    return order, nil
}

// User Service provides interface
type UserService interface {
    GetUser(id UserID) (*User, error)
    ValidateUser(id UserID) (bool, error)
}
```

### 2. Asynchronous Communication
**Event-driven communication** between services:

```go
// Order Service publishes events
type OrderService struct {
    eventPublisher EventPublisher
}

func (s *OrderService) CreateOrder(customerID CustomerID, items []OrderItem) (*Order, error) {
    order := NewOrder(customerID, items)
    
    // Publish event
    event := OrderCreated{
        OrderID:     order.ID,
        CustomerID:  customerID,
        Items:       items,
        OccurredAt:  time.Now(),
        Version:     1,
    }
    
    if err := s.eventPublisher.Publish([]DomainEvent{event}); err != nil {
        return nil, err
    }
    
    return order, nil
}

// Inventory Service subscribes to events
type InventoryService struct {
    inventoryRepository InventoryRepository
}

func (s *InventoryService) HandleOrderCreated(event OrderCreated) error {
    for _, item := range event.Items {
        if err := s.inventoryRepository.Reserve(item.ProductID, item.Quantity); err != nil {
            return err
        }
    }
    
    return nil
}
```

### 3. Hybrid Communication
**Combine synchronous and asynchronous** patterns:

```go
// Order Service uses both patterns
type OrderService struct {
    userService      UserService        // Synchronous
    eventPublisher   EventPublisher     // Asynchronous
    inventoryService InventoryService   // Asynchronous
}

func (s *OrderService) CreateOrder(customerID CustomerID, items []OrderItem) (*Order, error) {
    // Synchronous: Validate customer
    user, err := s.userService.GetUser(customerID)
    if err != nil {
        return nil, err
    }
    
    if !user.IsActive() {
        return nil, errors.New("customer is not active")
    }
    
    // Create order
    order := NewOrder(customerID, items)
    
    // Asynchronous: Reserve inventory
    event := InventoryReservationRequested{
        OrderID:    order.ID,
        Items:      items,
        OccurredAt: time.Now(),
        Version:    1,
    }
    
    if err := s.eventPublisher.Publish([]DomainEvent{event}); err != nil {
        return nil, err
    }
    
    return order, nil
}
```

## 🎯 Service Design Guidelines

### 1. Single Responsibility
**Each service has one clear responsibility**:
- User Service: User management
- Product Service: Product catalog
- Order Service: Order processing
- Payment Service: Payment processing

### 2. Loose Coupling
**Minimize dependencies between services**:
- Use events for communication
- Avoid direct database access
- Use anti-corruption layers
- Implement circuit breakers

### 3. High Cohesion
**Keep related functionality together**:
- User Service: User, UserPreferences, UserSession
- Order Service: Order, OrderItem, OrderStatus
- Product Service: Product, Category, Inventory

### 4. Clear Interfaces
**Well-defined APIs and contracts**:
- RESTful APIs
- Event schemas
- Data contracts
- Versioning strategy

### 5. Independent Deployment
**Services can be deployed independently**:
- Own database
- Own configuration
- Own deployment pipeline
- Own monitoring

## 🔧 Service Implementation Patterns

### 1. Repository Pattern
**Abstract data access**:

```go
type UserRepository interface {
    Save(user *User) error
    FindByID(id UserID) (*User, error)
    FindByEmail(email Email) (*User, error)
    FindAll() ([]User, error)
    Delete(id UserID) error
}

type UserRepositoryImpl struct {
    db *sql.DB
}

func (r *UserRepositoryImpl) Save(user *User) error {
    // Implementation details
}

func (r *UserRepositoryImpl) FindByID(id UserID) (*User, error) {
    // Implementation details
}
```

### 2. Service Layer Pattern
**Encapsulate business logic**:

```go
type UserService struct {
    userRepository UserRepository
    eventPublisher EventPublisher
    emailService   EmailService
}

func (s *UserService) CreateUser(email Email, name Name) (*User, error) {
    // Business logic
    user := NewUser(email, name)
    
    // Save to repository
    if err := s.userRepository.Save(user); err != nil {
        return nil, err
    }
    
    // Send welcome email
    if err := s.emailService.SendWelcomeEmail(user.Email, user.Name); err != nil {
        // Log error but don't fail the operation
        log.Printf("Failed to send welcome email: %v", err)
    }
    
    // Publish events
    events := user.GetEvents()
    if err := s.eventPublisher.Publish(events); err != nil {
        return nil, err
    }
    
    user.ClearEvents()
    return user, nil
}
```

### 3. Factory Pattern
**Create complex objects**:

```go
type UserFactory struct {
    idGenerator IDGenerator
    validator   UserValidator
}

func (f *UserFactory) CreateUser(email Email, name Name) (*User, error) {
    // Validate input
    if err := f.validator.ValidateEmail(email); err != nil {
        return nil, err
    }
    
    if err := f.validator.ValidateName(name); err != nil {
        return nil, err
    }
    
    // Generate ID
    id, err := f.idGenerator.Generate()
    if err != nil {
        return nil, err
    }
    
    // Create user
    user := &User{
        ID:        UserID(id),
        Email:     email,
        Name:      name,
        Status:    UserStatusPending,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        events:    []DomainEvent{},
    }
    
    // Add domain event
    event := UserCreated{
        UserID:    user.ID,
        Email:     user.Email,
        Name:      user.Name,
        OccurredAt: time.Now(),
        Version:   1,
    }
    user.events = append(user.events, event)
    
    return user, nil
}
```

## 🎯 Service Testing Strategies

### 1. Unit Testing
**Test individual components**:

```go
func TestUserService_CreateUser(t *testing.T) {
    // Arrange
    mockRepo := &MockUserRepository{}
    mockPublisher := &MockEventPublisher{}
    mockEmailService := &MockEmailService{}
    
    service := &UserService{
        userRepository: mockRepo,
        eventPublisher: mockPublisher,
        emailService:   mockEmailService,
    }
    
    email := Email("test@example.com")
    name := Name{First: "John", Last: "Doe"}
    
    // Act
    user, err := service.CreateUser(email, name)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, email, user.Email)
    assert.Equal(t, name, user.Name)
    
    // Verify interactions
    mockRepo.AssertExpectations(t)
    mockPublisher.AssertExpectations(t)
    mockEmailService.AssertExpectations(t)
}
```

### 2. Integration Testing
**Test service interactions**:

```go
func TestUserService_Integration(t *testing.T) {
    // Setup test database
    db := setupTestDB(t)
    defer cleanupTestDB(t, db)
    
    // Setup real dependencies
    repo := NewUserRepositoryImpl(db)
    publisher := NewEventPublisherImpl()
    emailService := NewEmailServiceImpl()
    
    service := &UserService{
        userRepository: repo,
        eventPublisher: publisher,
        emailService:   emailService,
    }
    
    // Test
    email := Email("test@example.com")
    name := Name{First: "John", Last: "Doe"}
    
    user, err := service.CreateUser(email, name)
    assert.NoError(t, err)
    assert.NotNil(t, user)
    
    // Verify in database
    savedUser, err := repo.FindByID(user.ID)
    assert.NoError(t, err)
    assert.Equal(t, user.ID, savedUser.ID)
}
```

### 3. Contract Testing
**Test service contracts**:

```go
func TestUserService_Contract(t *testing.T) {
    // Test API contract
    server := setupTestServer(t)
    defer server.Close()
    
    // Test create user endpoint
    req := CreateUserRequest{
        Email:     "test@example.com",
        FirstName: "John",
        LastName:  "Doe",
    }
    
    resp, err := http.Post(server.URL+"/api/v1/users", "application/json", 
        strings.NewReader(toJSON(req)))
    assert.NoError(t, err)
    assert.Equal(t, http.StatusCreated, resp.StatusCode)
    
    // Test get user endpoint
    var user UserResponse
    err = json.NewDecoder(resp.Body).Decode(&user)
    assert.NoError(t, err)
    
    getResp, err := http.Get(server.URL + "/api/v1/users/" + user.Data.ID)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, getResp.StatusCode)
}
```

## 🔗 Next Steps

Now that you understand service design, we'll move to:
1. **Implementation** - Building DDD-based services
2. **Practice** - Hands-on exercises
3. **Real-world Examples** - E-commerce system
4. **Testing** - Comprehensive testing strategies

**Ready to continue? Let's build some services!** 🚀
