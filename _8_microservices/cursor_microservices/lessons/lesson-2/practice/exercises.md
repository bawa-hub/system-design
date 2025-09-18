# 🧪 LESSON 2 PRACTICE EXERCISES - DOMAIN-DRIVEN DESIGN

## 🎯 Learning Objectives
By completing these exercises, you will:
- Understand Domain-Driven Design principles
- Practice identifying bounded contexts
- Learn to design aggregates and entities
- Master value object design
- Implement domain events
- Apply DDD to microservices

## 📚 Exercise 1: Theory Understanding

### Question 1.1: Bounded Contexts
**Scenario**: You're designing an e-commerce system with the following capabilities:
- User registration and authentication
- Product catalog management
- Shopping cart functionality
- Order processing and fulfillment
- Payment processing
- Inventory management
- Customer support and returns
- Analytics and reporting

**Question**: Identify the bounded contexts and explain why you chose these boundaries.

**Answer Space**:
```
Bounded Contexts:

1. User Management Context:
   - Responsibilities:
   - Data owned:
   - Why separate:

2. Product Catalog Context:
   - Responsibilities:
   - Data owned:
   - Why separate:

3. Order Processing Context:
   - Responsibilities:
   - Data owned:
   - Why separate:

4. Payment Context:
   - Responsibilities:
   - Data owned:
   - Why separate:

5. Inventory Context:
   - Responsibilities:
   - Data owned:
   - Why separate:

6. Customer Support Context:
   - Responsibilities:
   - Data owned:
   - Why separate:

7. Analytics Context:
   - Responsibilities:
   - Data owned:
   - Why separate:
```

### Question 1.2: Aggregates and Entities
**Scenario**: In the Order Processing context, you need to model:
- Orders with multiple items
- Order status and lifecycle
- Customer information
- Shipping details
- Payment information

**Question**: Design the aggregates, entities, and value objects for this context.

**Answer Space**:
```
Aggregates:

1. Order Aggregate:
   - Root Entity: Order
   - Entities: OrderItem
   - Value Objects: OrderID, Money, Address
   - Invariants:
   - Business Rules:

2. Customer Aggregate:
   - Root Entity: Customer
   - Value Objects: CustomerID, Email, Name
   - Invariants:
   - Business Rules:

Entities:

1. Order Entity:
   - Identity: OrderID
   - Attributes:
   - Behaviors:
   - Lifecycle:

2. OrderItem Entity:
   - Identity: OrderItemID
   - Attributes:
   - Behaviors:
   - Lifecycle:

Value Objects:

1. OrderID:
   - Properties:
   - Validation:
   - Immutability:

2. Money:
   - Properties:
   - Validation:
   - Behaviors:

3. Address:
   - Properties:
   - Validation:
   - Behaviors:
```

### Question 1.3: Domain Events
**Scenario**: When a customer places an order, several things happen:
- Order is created
- Inventory is reserved
- Payment is processed
- Confirmation email is sent
- Analytics are updated

**Question**: Design the domain events and event flow for this scenario.

**Answer Space**:
```
Domain Events:

1. OrderCreated:
   - Triggered by:
   - Data included:
   - Subscribers:

2. InventoryReserved:
   - Triggered by:
   - Data included:
   - Subscribers:

3. PaymentProcessed:
   - Triggered by:
   - Data included:
   - Subscribers:

4. OrderConfirmed:
   - Triggered by:
   - Data included:
   - Subscribers:

Event Flow:
1. Customer places order
2. OrderCreated event published
3. Inventory service reserves items
4. Payment service processes payment
5. OrderConfirmed event published
6. Email service sends confirmation
7. Analytics service updates metrics
```

## 🛠️ Exercise 2: Value Object Design

### Task 2.1: Design Email Value Object
Create a robust Email value object with the following requirements:
- Valid email format validation
- Case insensitive storage
- Domain extraction
- Equality comparison
- Immutability

**Your Implementation**:
```go
type Email struct {
    // Your implementation here
}

func NewEmail(value string) (Email, error) {
    // Your implementation here
}

func (e Email) String() string {
    // Your implementation here
}

func (e Email) Equals(other Email) bool {
    // Your implementation here
}

func (e Email) Domain() string {
    // Your implementation here
}

// Additional methods as needed
```

### Task 2.2: Design Money Value Object
Create a Money value object with the following requirements:
- Amount and currency validation
- Arithmetic operations (add, subtract, multiply)
- Currency conversion support
- Equality comparison
- Immutability

**Your Implementation**:
```go
type Money struct {
    // Your implementation here
}

func NewMoney(amount int64, currency string) (Money, error) {
    // Your implementation here
}

func (m Money) Add(other Money) (Money, error) {
    // Your implementation here
}

func (m Money) Subtract(other Money) (Money, error) {
    // Your implementation here
}

func (m Money) Multiply(factor float64) (Money, error) {
    // Your implementation here
}

func (m Money) Equals(other Money) bool {
    // Your implementation here
}

// Additional methods as needed
```

### Task 2.3: Design Address Value Object
Create an Address value object with the following requirements:
- Street, city, state, postal code, country
- Validation for required fields
- Format validation for postal code
- Equality comparison
- String representation

**Your Implementation**:
```go
type Address struct {
    // Your implementation here
}

func NewAddress(street, city, state, postalCode, country string) (Address, error) {
    // Your implementation here
}

func (a Address) String() string {
    // Your implementation here
}

func (a Address) Equals(other Address) bool {
    // Your implementation here
}

// Additional methods as needed
```

## 🏗️ Exercise 3: Entity Design

### Task 3.1: Design User Entity
Create a User entity with the following requirements:
- Unique identity (UserID)
- Email and name (value objects)
- Status (pending, active, inactive)
- Creation and update timestamps
- Rich behavior methods
- Domain event publishing

**Your Implementation**:
```go
type User struct {
    // Your implementation here
}

func NewUser(email Email, name Name) *User {
    // Your implementation here
}

func (u *User) Activate() error {
    // Your implementation here
}

func (u *User) Deactivate() error {
    // Your implementation here
}

func (u *User) ChangeEmail(newEmail Email) error {
    // Your implementation here
}

func (u *User) UpdateName(newName Name) error {
    // Your implementation here
}

func (u *User) IsActive() bool {
    // Your implementation here
}

// Additional methods as needed
```

### Task 3.2: Design Order Entity
Create an Order entity with the following requirements:
- Unique identity (OrderID)
- Customer ID
- Order items (collection)
- Status (draft, confirmed, shipped, delivered, cancelled)
- Total amount calculation
- Rich behavior methods
- Domain event publishing

**Your Implementation**:
```go
type Order struct {
    // Your implementation here
}

func NewOrder(customerID CustomerID) *Order {
    // Your implementation here
}

func (o *Order) AddItem(productID ProductID, quantity int, price Money) error {
    // Your implementation here
}

func (o *Order) RemoveItem(productID ProductID) error {
    // Your implementation here
}

func (o *Order) Confirm() error {
    // Your implementation here
}

func (o *Order) Ship(trackingNumber string) error {
    // Your implementation here
}

func (o *Order) Deliver() error {
    // Your implementation here
}

func (o *Order) Cancel(reason string) error {
    // Your implementation here
}

// Additional methods as needed
```

## 🎯 Exercise 4: Aggregate Design

### Task 4.1: Design Order Aggregate
Create an Order aggregate with the following requirements:
- Order as aggregate root
- OrderItem as entity within aggregate
- Invariant enforcement
- Controlled access to internal objects
- Domain event publishing

**Your Implementation**:
```go
type Order struct {
    // Your implementation here
}

type OrderItem struct {
    // Your implementation here
}

func (o *Order) AddItem(productID ProductID, quantity int, price Money) error {
    // Your implementation here
    // Enforce invariants
    // Update total
    // Publish events
}

func (o *Order) RemoveItem(productID ProductID) error {
    // Your implementation here
    // Enforce invariants
    // Update total
    // Publish events
}

func (o *Order) GetItems() []OrderItem {
    // Your implementation here
    // Return copy to prevent external modification
}

func (o *Order) GetTotalAmount() Money {
    // Your implementation here
}

// Private methods
func (o *Order) recalculateTotal() {
    // Your implementation here
}

func (o *Order) enforceInvariants() error {
    // Your implementation here
    // Check business rules
}
```

### Task 4.2: Design User Aggregate
Create a User aggregate with the following requirements:
- User as aggregate root
- UserPreferences as entity within aggregate
- Invariant enforcement
- Controlled access to internal objects
- Domain event publishing

**Your Implementation**:
```go
type User struct {
    // Your implementation here
}

type UserPreferences struct {
    // Your implementation here
}

func (u *User) UpdatePreferences(prefs UserPreferences) error {
    // Your implementation here
    // Enforce invariants
    // Publish events
}

func (u *User) GetPreferences() UserPreferences {
    // Your implementation here
    // Return copy to prevent external modification
}

// Additional methods as needed
```

## 📡 Exercise 5: Domain Events

### Task 5.1: Design Domain Events
Create domain events for the User aggregate:
- UserCreated
- UserActivated
- UserDeactivated
- UserEmailChanged
- UserNameChanged

**Your Implementation**:
```go
type UserCreated struct {
    // Your implementation here
}

func NewUserCreated(userID UserID, email Email, name Name) UserCreated {
    // Your implementation here
}

type UserActivated struct {
    // Your implementation here
}

func NewUserActivated(userID UserID, email Email) UserActivated {
    // Your implementation here
}

// Additional events as needed
```

### Task 5.2: Implement Event Publishing
Implement event publishing in the User aggregate:
- Collect events during operations
- Provide method to get events
- Clear events after publishing

**Your Implementation**:
```go
type User struct {
    // Your implementation here
    events []DomainEvent
}

func (u *User) GetEvents() []DomainEvent {
    // Your implementation here
}

func (u *User) ClearEvents() {
    // Your implementation here
}

func (u *User) addEvent(event DomainEvent) {
    // Your implementation here
}
```

## 🔄 Exercise 6: Service Design

### Task 6.1: Design User Service
Create a User service that:
- Encapsulates business logic
- Uses repository for data access
- Publishes domain events
- Handles external service integration

**Your Implementation**:
```go
type UserService struct {
    // Your implementation here
}

func NewUserService(repository UserRepository, eventPublisher EventPublisher, emailService EmailService) *UserService {
    // Your implementation here
}

func (s *UserService) CreateUser(email Email, name Name) (*User, error) {
    // Your implementation here
    // 1. Create user aggregate
    // 2. Save to repository
    // 3. Publish events
    // 4. Send welcome email
}

func (s *UserService) ActivateUser(userID UserID) error {
    // Your implementation here
    // 1. Load user aggregate
    // 2. Execute business logic
    // 3. Save to repository
    // 4. Publish events
}

// Additional methods as needed
```

### Task 6.2: Design Repository Interface
Create repository interfaces for the User aggregate:
- Abstract data access
- Domain-oriented interface
- Easy to mock for testing

**Your Implementation**:
```go
type UserRepository interface {
    Save(user *User) error
    FindByID(id UserID) (*User, error)
    FindByEmail(email Email) (*User, error)
    FindAll() ([]User, error)
    Delete(id UserID) error
}

type UserRepositoryImpl struct {
    // Your implementation here
}

func (r *UserRepositoryImpl) Save(user *User) error {
    // Your implementation here
}

func (r *UserRepositoryImpl) FindByID(id UserID) (*User, error) {
    // Your implementation here
}

// Additional methods as needed
```

## 🧪 Exercise 7: Testing

### Task 7.1: Unit Testing
Create unit tests for the User entity:
- Test creation
- Test activation/deactivation
- Test email changes
- Test name changes
- Test domain events

**Your Implementation**:
```go
func TestUser_Creation(t *testing.T) {
    // Your test implementation here
}

func TestUser_Activation(t *testing.T) {
    // Your test implementation here
}

func TestUser_EmailChange(t *testing.T) {
    // Your test implementation here
}

func TestUser_DomainEvents(t *testing.T) {
    // Your test implementation here
}

// Additional tests as needed
```

### Task 7.2: Integration Testing
Create integration tests for the User service:
- Test with real database
- Test event publishing
- Test external service integration

**Your Implementation**:
```go
func TestUserService_Integration(t *testing.T) {
    // Setup test database
    // Setup real dependencies
    // Test service operations
    // Verify database state
    // Verify events published
}

func TestUserService_EventPublishing(t *testing.T) {
    // Setup mock event publisher
    // Test service operations
    // Verify events published
}

// Additional tests as needed
```

## 🎯 Exercise 8: Real-World Application

### Task 8.1: E-commerce System Design
Design a complete e-commerce system using DDD principles:
- Identify all bounded contexts
- Design aggregates for each context
- Define domain events
- Plan service communication

**Your Design**:
```
Bounded Contexts:
1. User Management Context
2. Product Catalog Context
3. Order Processing Context
4. Payment Context
5. Inventory Context
6. Shipping Context
7. Customer Support Context

Aggregates:
1. User Management Context:
   - User aggregate
   - UserPreferences aggregate

2. Product Catalog Context:
   - Product aggregate
   - Category aggregate

3. Order Processing Context:
   - Order aggregate
   - OrderItem aggregate

4. Payment Context:
   - Payment aggregate
   - PaymentMethod aggregate

5. Inventory Context:
   - InventoryItem aggregate
   - StockReservation aggregate

Domain Events:
1. UserCreated, UserActivated, UserDeactivated
2. ProductCreated, ProductUpdated, ProductDiscontinued
3. OrderCreated, OrderConfirmed, OrderShipped, OrderDelivered
4. PaymentProcessed, PaymentFailed, PaymentRefunded
5. InventoryReserved, InventoryReleased, InventoryUpdated

Service Communication:
1. Synchronous: User validation, Product information
2. Asynchronous: Order processing, Payment processing
3. Event-driven: Inventory management, Shipping updates
```

### Task 8.2: Implementation Plan
Create an implementation plan for the e-commerce system:
- Phase 1: Core domains
- Phase 2: Service integration
- Phase 3: Event-driven architecture
- Phase 4: Advanced features

**Your Plan**:
```
Phase 1: Core Domains (Weeks 1-2)
- Implement User Management Context
- Implement Product Catalog Context
- Basic CRUD operations
- Unit testing

Phase 2: Service Integration (Weeks 3-4)
- Implement Order Processing Context
- Implement Payment Context
- Service-to-service communication
- Integration testing

Phase 3: Event-Driven Architecture (Weeks 5-6)
- Implement domain events
- Event publishing and handling
- Asynchronous communication
- Event sourcing

Phase 4: Advanced Features (Weeks 7-8)
- Implement remaining contexts
- Advanced patterns (CQRS, Saga)
- Performance optimization
- Production deployment
```

## ✅ Completion Checklist

- [ ] Read all theory notes
- [ ] Understand DDD concepts
- [ ] Complete all exercises
- [ ] Design value objects
- [ ] Design entities and aggregates
- [ ] Implement domain events
- [ ] Design services and repositories
- [ ] Write comprehensive tests
- [ ] Design complete system
- [ ] Ready for Lesson 3

## 🎯 Next Steps

After completing these exercises, you should:
1. Have a deep understanding of DDD principles
2. Know how to identify bounded contexts
3. Be able to design aggregates and entities
4. Understand value object design
5. Know how to implement domain events
6. Be able to apply DDD to microservices
7. Have hands-on experience with DDD patterns

**Ready for Lesson 3? Let me know when you've completed these exercises!** 🚀
