# 🎯 BOUNDED CONTEXTS - IDENTIFYING SERVICE BOUNDARIES

## 🎯 What are Bounded Contexts?

### Definition
A bounded context is a logical boundary within which a particular domain model is defined and applicable. It's the context in which a specific domain model is valid and meaningful.

### Key Characteristics
- **Clear boundaries** that define what's inside and outside
- **Own ubiquitous language** specific to that context
- **Independent domain model** that can evolve separately
- **Own data** and business rules
- **Clear ownership** by a specific team

## 🔍 Identifying Bounded Contexts

### 1. Domain Analysis Process

#### Step 1: Map Business Capabilities
**Questions to Ask**:
- What are the main business functions?
- What are the different user types?
- What are the core business processes?
- What are the different business areas?

**Example - E-commerce Domain**:
```
Business Capabilities:
- User Management (registration, authentication, profiles)
- Product Catalog (products, categories, inventory)
- Order Processing (orders, fulfillment, tracking)
- Payment Processing (payments, billing, refunds)
- Shipping (delivery, tracking, logistics)
- Customer Service (support, returns, complaints)
```

#### Step 2: Identify Data Ownership
**Questions to Ask**:
- Who owns what data?
- Who can modify what data?
- What data is shared vs. owned?
- What are the data dependencies?

**Example - Data Ownership**:
```
User Management Context:
- Owns: User profiles, authentication data, preferences
- Needs: Product information (for recommendations)
- Provides: User information to other contexts

Product Catalog Context:
- Owns: Product information, categories, inventory
- Needs: User preferences (for personalization)
- Provides: Product data to other contexts
```

#### Step 3: Find Natural Boundaries
**Questions to Ask**:
- Where are the natural breaks in the domain?
- What can change independently?
- What are the different user perspectives?
- Where are the different business rules?

**Example - Natural Boundaries**:
```
User Management: Authentication, profiles, preferences
Product Catalog: Product information, categories, inventory
Order Processing: Orders, fulfillment, tracking
Payment: Payment processing, billing, refunds
```

### 2. Context Mapping Techniques

#### Event Storming
**Process**:
1. **Gather domain experts** and stakeholders
2. **Map business events** on sticky notes
3. **Identify aggregates** and entities
4. **Find context boundaries** naturally
5. **Document relationships** between contexts

**Example Event Storming Output**:
```
User Management Context:
- UserRegistered
- UserProfileUpdated
- UserPreferencesChanged
- UserAuthenticated

Product Catalog Context:
- ProductCreated
- ProductUpdated
- InventoryUpdated
- CategoryAdded

Order Processing Context:
- OrderCreated
- OrderConfirmed
- OrderShipped
- OrderDelivered
```

#### Domain Storytelling
**Process**:
1. **Tell stories** about business processes
2. **Identify actors** and their roles
3. **Map interactions** between actors
4. **Find context boundaries** in stories
5. **Validate with stakeholders**

**Example Domain Story**:
```
Story: Customer places an order
Actors: Customer, Order System, Payment System, Inventory System
Interactions:
1. Customer views products (Product Catalog Context)
2. Customer adds items to cart (Order Context)
3. Customer provides payment info (Payment Context)
4. System checks inventory (Product Catalog Context)
5. Order is confirmed (Order Context)
6. Payment is processed (Payment Context)
```

## 🗺️ Context Mapping Patterns

### 1. Shared Kernel
**Definition**: Shared code between two contexts.

**When to Use**:
- Common domain concepts
- Shared business rules
- Tightly coupled contexts

**Example**:
```go
// Shared kernel for common domain concepts
package shared

type Money struct {
    Amount   int64
    Currency string
}

type Email struct {
    value string
}

type UserID struct {
    value string
}
```

**Pros**:
- Reduces duplication
- Ensures consistency
- Shared understanding

**Cons**:
- Tight coupling
- Coordination overhead
- Change impact

### 2. Customer-Supplier
**Definition**: One context depends on another, with the supplier providing services.

**When to Use**:
- Clear dependency relationship
- Supplier can meet customer needs
- Supplier has control over interface

**Example**:
```
Payment Context (Supplier) → Order Context (Customer)
- Order Context depends on Payment Context
- Payment Context provides payment services
- Order Context adapts to Payment Context interface
```

**Implementation**:
```go
// Order Context depends on Payment Context
type OrderService struct {
    paymentService PaymentService // Interface to Payment Context
}

func (s *OrderService) ProcessOrder(order Order) error {
    // Use payment service from Payment Context
    paymentResult, err := s.paymentService.ProcessPayment(order.PaymentInfo)
    if err != nil {
        return err
    }
    
    // Update order with payment result
    order.Status = OrderStatusPaid
    return s.orderRepository.Save(order)
}
```

### 3. Conformist
**Definition**: One context conforms to another's model without modification.

**When to Use**:
- Clear power imbalance
- Supplier controls the interface
- Customer has no choice

**Example**:
```
External Payment Provider → Payment Context
- Payment Context conforms to external API
- No modification of external model
- Direct integration
```

**Implementation**:
```go
// Payment Context conforms to external payment provider
type ExternalPaymentService struct {
    api ExternalPaymentAPI
}

func (s *ExternalPaymentService) ProcessPayment(payment Payment) error {
    // Convert to external format
    externalPayment := ExternalPayment{
        Amount:      payment.Amount,
        Currency:    payment.Currency,
        CardNumber:  payment.CardNumber,
        ExpiryDate:  payment.ExpiryDate,
    }
    
    // Call external API
    result, err := s.api.Process(externalPayment)
    if err != nil {
        return err
    }
    
    // Handle result
    return s.handlePaymentResult(result)
}
```

### 4. Anti-Corruption Layer
**Definition**: Translation layer between contexts to prevent corruption.

**When to Use**:
- Legacy system integration
- Third-party API integration
- Different domain models
- Need to protect domain model

**Example**:
```
Legacy User System → User Management Context
- Anti-corruption layer translates between models
- Protects User Management Context from legacy model
- Allows gradual migration
```

**Implementation**:
```go
// Anti-corruption layer for legacy user system
type LegacyUserAdapter struct {
    legacyAPI LegacyUserAPI
    mapper    UserMapper
}

func (a *LegacyUserAdapter) GetUser(id UserID) (User, error) {
    // Call legacy system
    legacyUser, err := a.legacyAPI.GetUser(string(id))
    if err != nil {
        return User{}, err
    }
    
    // Convert to domain model
    return a.mapper.ToDomain(legacyUser), nil
}

func (a *LegacyUserAdapter) CreateUser(user User) error {
    // Convert to legacy format
    legacyUser := a.mapper.ToLegacy(user)
    
    // Call legacy system
    return a.legacyAPI.CreateUser(legacyUser)
}
```

### 5. Open Host Service
**Definition**: Published interface for other contexts to use.

**When to Use**:
- Multiple contexts need the same service
- Service is stable and well-defined
- Need to reduce integration complexity

**Example**:
```
User Management Context → Open Host Service
- Provides user information to multiple contexts
- Stable, well-defined interface
- Reduces integration complexity
```

**Implementation**:
```go
// Open host service for user information
type UserService struct {
    userRepository UserRepository
}

// Published interface for other contexts
func (s *UserService) GetUserInfo(id UserID) (UserInfo, error) {
    user, err := s.userRepository.FindByID(id)
    if err != nil {
        return UserInfo{}, err
    }
    
    return UserInfo{
        ID:    user.ID,
        Email: user.Email,
        Name:  user.Name,
    }, nil
}

func (s *UserService) ValidateUser(id UserID) (bool, error) {
    user, err := s.userRepository.FindByID(id)
    if err != nil {
        return false, err
    }
    
    return user.IsActive(), nil
}
```

### 6. Published Language
**Definition**: Shared data format between contexts.

**When to Use**:
- Multiple contexts need to exchange data
- Need for data consistency
- Complex data structures

**Example**:
```
Order Context ↔ Payment Context
- Shared order data format
- Consistent data exchange
- Well-defined contracts
```

**Implementation**:
```go
// Published language for order data
type OrderData struct {
    ID          string    `json:"id"`
    CustomerID  string    `json:"customer_id"`
    Items       []Item    `json:"items"`
    TotalAmount Money     `json:"total_amount"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
}

type Item struct {
    ProductID string `json:"product_id"`
    Quantity  int    `json:"quantity"`
    Price     Money  `json:"price"`
}

type Money struct {
    Amount   int64  `json:"amount"`
    Currency string `json:"currency"`
}
```

## 🎯 Context Design Principles

### 1. High Cohesion
**Principle**: Keep related concepts together within a context.

**Example**:
```
Good: User Management Context
- User registration
- User authentication
- User profiles
- User preferences

Bad: Mixed Context
- User registration
- Product catalog
- Order processing
- Payment processing
```

### 2. Loose Coupling
**Principle**: Minimize dependencies between contexts.

**Example**:
```
Good: Loose Coupling
User Context → Order Context (via events)
Order Context → Payment Context (via events)

Bad: Tight Coupling
User Context → Order Context (direct calls)
Order Context → Payment Context (direct calls)
```

### 3. Clear Ownership
**Principle**: Each context has a clear owner and responsibility.

**Example**:
```
User Management Context:
- Owner: User Management Team
- Responsibility: User lifecycle, authentication, profiles
- Data: User data, authentication data, preferences

Product Catalog Context:
- Owner: Product Team
- Responsibility: Product information, categories, inventory
- Data: Product data, category data, inventory data
```

### 4. Independent Evolution
**Principle**: Contexts can evolve independently.

**Example**:
```
User Management Context:
- Can change user model
- Can add new features
- Can change technology
- Independent deployment

Product Catalog Context:
- Can change product model
- Can add new features
- Can change technology
- Independent deployment
```

## 🔄 Context Evolution

### 1. Context Splitting
**When to Split**:
- Context becomes too large
- Different teams need different models
- Different business rules
- Different data ownership

**Process**:
1. Identify natural boundaries
2. Define new contexts
3. Create context maps
4. Plan migration strategy
5. Implement gradually

### 2. Context Merging
**When to Merge**:
- Contexts are too small
- Tightly coupled contexts
- Same team ownership
- Similar business rules

**Process**:
1. Identify merge candidates
2. Define merged context
3. Plan integration strategy
4. Implement gradually
5. Update context maps

### 3. Context Refactoring
**When to Refactor**:
- Context boundaries are unclear
- Dependencies are complex
- Business rules are scattered
- Data ownership is unclear

**Process**:
1. Analyze current context
2. Identify issues
3. Design new structure
4. Plan refactoring
5. Implement gradually

## 🎯 Context Design Checklist

### Before Creating a Context
- [ ] Clear business capability
- [ ] Clear data ownership
- [ ] Clear team ownership
- [ ] Clear boundaries
- [ ] Clear responsibilities

### Context Design
- [ ] High cohesion within context
- [ ] Loose coupling between contexts
- [ ] Clear interfaces
- [ ] Well-defined contracts
- [ ] Independent evolution

### Context Implementation
- [ ] Proper domain modeling
- [ ] Clear ubiquitous language
- [ ] Well-defined APIs
- [ ] Proper error handling
- [ ] Good documentation

## 🔗 Next Steps

Now that you understand bounded contexts, we'll move to:
1. **Aggregates and Entities** - Detailed modeling patterns
2. **Domain Events** - Event-driven architecture
3. **Service Design** - Applying DDD to microservices
4. **Implementation** - Building DDD-based services

**Ready to continue? Let's learn about aggregates and entities!** 🚀
