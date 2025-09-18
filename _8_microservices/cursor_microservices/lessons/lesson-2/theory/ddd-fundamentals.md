# 📖 DOMAIN-DRIVEN DESIGN FUNDAMENTALS

## 🎯 What is Domain-Driven Design (DDD)?

### Definition
Domain-Driven Design is a software development approach that focuses on creating software that reflects a deep understanding of the business domain. It emphasizes:

- **Domain expertise** over technical implementation
- **Ubiquitous language** shared between developers and domain experts
- **Rich domain models** that capture business logic
- **Strategic design** for large, complex systems

### Why DDD Matters for Microservices

1. **Service Boundaries**: DDD helps identify where to split monoliths into microservices
2. **Domain Modeling**: Each service represents a bounded context with its own domain model
3. **Communication**: Services communicate through well-defined domain interfaces
4. **Evolution**: Domain models can evolve independently within each service

## 🏗️ Core DDD Concepts

### 1. Domain
**Definition**: The sphere of knowledge and activity around which the application logic revolves.

**Example**: E-commerce domain includes:
- Product catalog management
- Order processing
- Customer management
- Payment processing
- Inventory management

**Characteristics**:
- Contains business rules and logic
- Represents real-world concepts
- Changes based on business needs
- Independent of technical implementation

### 2. Ubiquitous Language
**Definition**: A common vocabulary used by all team members to connect all activities of the team with the software.

**Benefits**:
- Reduces miscommunication
- Aligns technical and business teams
- Makes code self-documenting
- Improves domain understanding

**Example**:
```
Business Term: "Order"
Technical Term: "OrderEntity"
Code: type Order struct { ... }
Database: orders table
API: /api/orders
```

### 3. Bounded Context
**Definition**: A logical boundary within which a particular domain model is defined and applicable.

**Characteristics**:
- Has its own ubiquitous language
- Contains its own domain model
- Has clear boundaries
- Can be implemented as a microservice

**Example**: E-commerce bounded contexts:
- **User Management Context**: User registration, authentication, profiles
- **Product Catalog Context**: Product information, categories, inventory
- **Order Processing Context**: Order creation, fulfillment, tracking
- **Payment Context**: Payment processing, billing, refunds

### 4. Context Map
**Definition**: A diagram that shows the relationships between bounded contexts.

**Types of Relationships**:
- **Shared Kernel**: Shared code between contexts
- **Customer-Supplier**: One context depends on another
- **Conformist**: One context conforms to another's model
- **Anti-Corruption Layer**: Translation layer between contexts
- **Open Host Service**: Published interface for other contexts
- **Published Language**: Shared data format between contexts

## 🎯 Strategic Design Patterns

### 1. Bounded Context Identification
**Process**:
1. **Identify domain experts** and stakeholders
2. **Map business processes** and workflows
3. **Identify data ownership** and responsibilities
4. **Find natural boundaries** in the domain
5. **Define context interfaces** and contracts

**Questions to Ask**:
- What are the main business capabilities?
- Who are the different user types?
- What data does each team own?
- Where are the natural boundaries?
- How do different parts communicate?

### 2. Context Mapping
**Process**:
1. **List all bounded contexts**
2. **Identify relationships** between contexts
3. **Define integration patterns**
4. **Plan communication strategies**
5. **Document dependencies**

**Example Context Map**:
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   User Mgmt     │    │  Product Cat    │    │  Order Proc     │
│   Context       │    │  Context        │    │  Context        │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  Auth Service   │    │ Product Service │    │  Order Service  │
│  (Microservice) │    │ (Microservice)  │    │ (Microservice)  │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

### 3. Anti-Corruption Layer
**Definition**: A layer that translates between different domain models to prevent corruption.

**When to Use**:
- Integrating with legacy systems
- Working with third-party APIs
- Connecting different bounded contexts

**Implementation**:
```go
// Anti-corruption layer for external payment service
type PaymentServiceAdapter struct {
    externalAPI ExternalPaymentAPI
    domainMapper PaymentDomainMapper
}

func (a *PaymentServiceAdapter) ProcessPayment(payment Payment) error {
    // Convert domain model to external API format
    externalPayment := a.domainMapper.ToExternal(payment)
    
    // Call external service
    result, err := a.externalAPI.Process(externalPayment)
    if err != nil {
        return err
    }
    
    // Convert external result back to domain model
    return a.domainMapper.ToDomain(result)
}
```

## 🏛️ Tactical Design Patterns

### 1. Entities
**Definition**: Objects that have a distinct identity that runs through time and different representations.

**Characteristics**:
- Have a unique identifier
- Can change over time
- Identity remains constant
- Can be compared by identity

**Example**:
```go
type User struct {
    ID        UserID
    Email     Email
    Name      Name
    CreatedAt time.Time
    UpdatedAt time.Time
}

// User can change email, name, but ID remains the same
func (u *User) ChangeEmail(newEmail Email) error {
    if !newEmail.IsValid() {
        return errors.New("invalid email")
    }
    u.Email = newEmail
    u.UpdatedAt = time.Now()
    return nil
}
```

### 2. Value Objects
**Definition**: Objects that are defined entirely by their attributes and have no conceptual identity.

**Characteristics**:
- Immutable
- No identity
- Can be compared by value
- Can be replaced

**Example**:
```go
type Email struct {
    value string
}

func NewEmail(value string) (Email, error) {
    if !isValidEmail(value) {
        return Email{}, errors.New("invalid email format")
    }
    return Email{value: value}, nil
}

func (e Email) String() string {
    return e.value
}

func (e Email) Equals(other Email) bool {
    return e.value == other.value
}
```

### 3. Aggregates
**Definition**: A cluster of domain objects that can be treated as a single unit for data changes.

**Characteristics**:
- Has a root entity (aggregate root)
- Enforces invariants
- Controls access to internal objects
- Maintains consistency

**Example**:
```go
type Order struct {
    ID          OrderID
    CustomerID  CustomerID
    Items       []OrderItem
    Status      OrderStatus
    TotalAmount Money
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// Aggregate root controls access to order items
func (o *Order) AddItem(productID ProductID, quantity int, price Money) error {
    if o.Status != OrderStatusDraft {
        return errors.New("cannot add items to non-draft order")
    }
    
    item := OrderItem{
        ProductID: productID,
        Quantity:  quantity,
        Price:     price,
    }
    
    o.Items = append(o.Items, item)
    o.recalculateTotal()
    o.UpdatedAt = time.Now()
    
    return nil
}

func (o *Order) recalculateTotal() {
    total := Money(0)
    for _, item := range o.Items {
        total += item.Price * Money(item.Quantity)
    }
    o.TotalAmount = total
}
```

### 4. Domain Services
**Definition**: Services that contain domain logic that doesn't naturally fit in entities or value objects.

**Characteristics**:
- Stateless
- Contain domain logic
- Not tied to specific entities
- Can be used across aggregates

**Example**:
```go
type PricingService struct {
    discountRules []DiscountRule
}

func (s *PricingService) CalculatePrice(product Product, customer Customer) Money {
    basePrice := product.Price
    
    // Apply customer-specific discounts
    for _, rule := range s.discountRules {
        if rule.AppliesTo(customer) {
            basePrice = rule.Apply(basePrice)
        }
    }
    
    return basePrice
}
```

### 5. Repositories
**Definition**: Objects that encapsulate the logic needed to access data sources.

**Characteristics**:
- Abstract data access
- Provide domain-oriented interface
- Hide implementation details
- Can be easily mocked for testing

**Example**:
```go
type UserRepository interface {
    Save(user User) error
    FindByID(id UserID) (User, error)
    FindByEmail(email Email) (User, error)
    FindAll() ([]User, error)
    Delete(id UserID) error
}

type UserRepositoryImpl struct {
    db *sql.DB
}

func (r *UserRepositoryImpl) Save(user User) error {
    // Implementation details hidden
    query := `INSERT INTO users (id, email, name, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) 
              ON CONFLICT (id) DO UPDATE SET 
              email = $2, name = $3, updated_at = $5`
    
    _, err := r.db.Exec(query, user.ID, user.Email, user.Name, user.CreatedAt, user.UpdatedAt)
    return err
}
```

## 🔄 Domain Events

### Definition
Domain events represent something significant that happened in the domain.

**Characteristics**:
- Immutable
- Represent past occurrences
- Can trigger side effects
- Enable loose coupling

**Example**:
```go
type DomainEvent interface {
    EventType() string
    OccurredAt() time.Time
    AggregateID() string
}

type UserRegistered struct {
    UserID    UserID
    Email     Email
    Name      Name
    OccurredAt time.Time
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
```

## 🎯 DDD Benefits for Microservices

### 1. Clear Service Boundaries
- Each bounded context becomes a microservice
- Clear ownership and responsibilities
- Reduced coupling between services

### 2. Domain-Focused Teams
- Teams organized around business capabilities
- Deep domain expertise
- Better communication with stakeholders

### 3. Independent Evolution
- Services can evolve independently
- Technology diversity
- Independent deployment

### 4. Better Testing
- Domain logic is isolated
- Easier to test business rules
- Clear test boundaries

## ⚠️ Common DDD Pitfalls

### 1. Over-Engineering
- Don't apply DDD to simple domains
- Start with simple models
- Add complexity when needed

### 2. Anemic Domain Models
- Avoid getter/setter only models
- Put business logic in domain objects
- Make models behavior-rich

### 3. Ignoring Context Boundaries
- Don't share domain models across contexts
- Use anti-corruption layers
- Maintain context independence

### 4. Premature Abstraction
- Don't abstract too early
- Let patterns emerge naturally
- Refactor when needed

## 🔗 Next Steps

Now that you understand DDD fundamentals, we'll move to:
1. **Bounded Contexts** - How to identify service boundaries
2. **Aggregates and Entities** - Detailed modeling patterns
3. **Domain Events** - Event-driven architecture
4. **Service Design** - Applying DDD to microservices

**Ready to continue? Let's learn about bounded contexts!** 🚀
