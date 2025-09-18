# 🏗️ AGGREGATES, ENTITIES & VALUE OBJECTS

## 🎯 Understanding Domain Objects

### The Three Types of Domain Objects

In Domain-Driven Design, there are three main types of domain objects:
1. **Entities** - Objects with identity
2. **Value Objects** - Objects defined by their attributes
3. **Aggregates** - Clusters of related objects

## 🆔 Entities

### Definition
Entities are objects that have a distinct identity that runs through time and different representations. They are defined by their identity, not by their attributes.

### Key Characteristics
- **Have a unique identifier** that remains constant
- **Can change over time** while maintaining identity
- **Identity is immutable** - never changes
- **Can be compared by identity** (not by attributes)
- **Have a lifecycle** - created, modified, deleted

### Entity Design Principles

#### 1. Strong Identity
```go
type UserID struct {
    value string
}

func NewUserID(value string) (UserID, error) {
    if value == "" {
        return UserID{}, errors.New("user ID cannot be empty")
    }
    return UserID{value: value}, nil
}

func (id UserID) String() string {
    return id.value
}

func (id UserID) Equals(other UserID) bool {
    return id.value == other.value
}
```

#### 2. Immutable Identity
```go
type User struct {
    ID        UserID    // Identity - never changes
    Email     Email     // Can change
    Name      Name      // Can change
    CreatedAt time.Time // Can change
    UpdatedAt time.Time // Can change
}

// Identity remains constant even when other attributes change
func (u *User) ChangeEmail(newEmail Email) error {
    if !newEmail.IsValid() {
        return errors.New("invalid email")
    }
    u.Email = newEmail
    u.UpdatedAt = time.Now()
    return nil
}
```

#### 3. Rich Behavior
```go
type User struct {
    ID        UserID
    Email     Email
    Name      Name
    Status    UserStatus
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Business logic in the entity
func (u *User) Activate() error {
    if u.Status == UserStatusActive {
        return errors.New("user is already active")
    }
    u.Status = UserStatusActive
    u.UpdatedAt = time.Now()
    return nil
}

func (u *User) Deactivate() error {
    if u.Status == UserStatusInactive {
        return errors.New("user is already inactive")
    }
    u.Status = UserStatusInactive
    u.UpdatedAt = time.Now()
    return nil
}

func (u *User) IsActive() bool {
    return u.Status == UserStatusActive
}
```

### Entity Examples

#### User Entity
```go
type User struct {
    ID          UserID
    Email       Email
    Name        Name
    Status      UserStatus
    Preferences UserPreferences
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

func NewUser(email Email, name Name) *User {
    now := time.Now()
    return &User{
        ID:        NewUserID(),
        Email:     email,
        Name:      name,
        Status:    UserStatusPending,
        CreatedAt: now,
        UpdatedAt: now,
    }
}

func (u *User) ChangeEmail(newEmail Email) error {
    if !newEmail.IsValid() {
        return errors.New("invalid email")
    }
    u.Email = newEmail
    u.UpdatedAt = time.Now()
    return nil
}

func (u *User) UpdatePreferences(prefs UserPreferences) error {
    if err := prefs.Validate(); err != nil {
        return err
    }
    u.Preferences = prefs
    u.UpdatedAt = time.Now()
    return nil
}
```

#### Order Entity
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

func NewOrder(customerID CustomerID) *Order {
    now := time.Now()
    return &Order{
        ID:         NewOrderID(),
        CustomerID: customerID,
        Items:      []OrderItem{},
        Status:     OrderStatusDraft,
        TotalAmount: Money{},
        CreatedAt:  now,
        UpdatedAt:  now,
    }
}

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

func (o *Order) Confirm() error {
    if o.Status != OrderStatusDraft {
        return errors.New("only draft orders can be confirmed")
    }
    
    if len(o.Items) == 0 {
        return errors.New("cannot confirm empty order")
    }
    
    o.Status = OrderStatusConfirmed
    o.UpdatedAt = time.Now()
    
    return nil
}
```

## 💎 Value Objects

### Definition
Value objects are objects that are defined entirely by their attributes and have no conceptual identity. They are immutable and can be compared by value.

### Key Characteristics
- **No identity** - defined by attributes
- **Immutable** - cannot be changed after creation
- **Can be compared by value** (not by identity)
- **Can be replaced** when values change
- **Encapsulate validation** and business rules

### Value Object Design Principles

#### 1. Immutability
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

// No setter methods - value objects are immutable
```

#### 2. Validation
```go
type Money struct {
    amount   int64
    currency string
}

func NewMoney(amount int64, currency string) (Money, error) {
    if amount < 0 {
        return Money{}, errors.New("amount cannot be negative")
    }
    
    if !isValidCurrency(currency) {
        return Money{}, errors.New("invalid currency")
    }
    
    return Money{
        amount:   amount,
        currency: currency,
    }, nil
}

func (m Money) Amount() int64 {
    return m.amount
}

func (m Money) Currency() string {
    return m.currency
}

func (m Money) Equals(other Money) bool {
    return m.amount == other.amount && m.currency == other.currency
}
```

#### 3. Rich Behavior
```go
type Money struct {
    amount   int64
    currency string
}

func (m Money) Add(other Money) (Money, error) {
    if m.currency != other.currency {
        return Money{}, errors.New("cannot add different currencies")
    }
    
    return NewMoney(m.amount+other.amount, m.currency)
}

func (m Money) Multiply(factor float64) (Money, error) {
    if factor < 0 {
        return Money{}, errors.New("factor cannot be negative")
    }
    
    newAmount := int64(float64(m.amount) * factor)
    return NewMoney(newAmount, m.currency)
}

func (m Money) IsZero() bool {
    return m.amount == 0
}

func (m Money) IsPositive() bool {
    return m.amount > 0
}
```

### Value Object Examples

#### Email Value Object
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

func (e Email) Domain() string {
    parts := strings.Split(e.value, "@")
    if len(parts) != 2 {
        return ""
    }
    return parts[1]
}

func isValidEmail(email string) bool {
    // Email validation logic
    return strings.Contains(email, "@") && len(email) > 5
}
```

#### Name Value Object
```go
type Name struct {
    first string
    last  string
}

func NewName(first, last string) (Name, error) {
    if strings.TrimSpace(first) == "" {
        return Name{}, errors.New("first name cannot be empty")
    }
    
    if strings.TrimSpace(last) == "" {
        return Name{}, errors.New("last name cannot be empty")
    }
    
    return Name{
        first: strings.TrimSpace(first),
        last:  strings.TrimSpace(last),
    }, nil
}

func (n Name) First() string {
    return n.first
}

func (n Name) Last() string {
    return n.last
}

func (n Name) Full() string {
    return n.first + " " + n.last
}

func (n Name) Equals(other Name) bool {
    return n.first == other.first && n.last == other.last
}
```

#### Address Value Object
```go
type Address struct {
    street     string
    city       string
    state      string
    postalCode string
    country    string
}

func NewAddress(street, city, state, postalCode, country string) (Address, error) {
    if strings.TrimSpace(street) == "" {
        return Address{}, errors.New("street cannot be empty")
    }
    
    if strings.TrimSpace(city) == "" {
        return Address{}, errors.New("city cannot be empty")
    }
    
    if strings.TrimSpace(country) == "" {
        return Address{}, errors.New("country cannot be empty")
    }
    
    return Address{
        street:     strings.TrimSpace(street),
        city:       strings.TrimSpace(city),
        state:      strings.TrimSpace(state),
        postalCode: strings.TrimSpace(postalCode),
        country:    strings.TrimSpace(country),
    }, nil
}

func (a Address) Street() string {
    return a.street
}

func (a Address) City() string {
    return a.city
}

func (a Address) State() string {
    return a.state
}

func (a Address) PostalCode() string {
    return a.postalCode
}

func (a Address) Country() string {
    return a.country
}

func (a Address) Equals(other Address) bool {
    return a.street == other.street &&
           a.city == other.city &&
           a.state == other.state &&
           a.postalCode == other.postalCode &&
           a.country == other.country
}

func (a Address) String() string {
    return fmt.Sprintf("%s, %s, %s %s, %s", 
        a.street, a.city, a.state, a.postalCode, a.country)
}
```

## 🏗️ Aggregates

### Definition
Aggregates are clusters of domain objects that can be treated as a single unit for data changes. They have a root entity (aggregate root) that controls access to all objects within the aggregate.

### Key Characteristics
- **Has a root entity** (aggregate root)
- **Enforces invariants** within the aggregate
- **Controls access** to internal objects
- **Maintains consistency** across the aggregate
- **Transaction boundary** for data changes

### Aggregate Design Principles

#### 1. Single Aggregate Root
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

// Order is the aggregate root
// OrderItem is an entity within the aggregate
type OrderItem struct {
    ProductID ProductID
    Quantity  int
    Price     Money
}

// Only Order can create, modify, or delete OrderItems
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
```

#### 2. Enforce Invariants
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

// Enforce business rules within the aggregate
func (o *Order) Confirm() error {
    if o.Status != OrderStatusDraft {
        return errors.New("only draft orders can be confirmed")
    }
    
    if len(o.Items) == 0 {
        return errors.New("cannot confirm empty order")
    }
    
    if o.TotalAmount.IsZero() {
        return errors.New("cannot confirm order with zero total")
    }
    
    o.Status = OrderStatusConfirmed
    o.UpdatedAt = time.Now()
    
    return nil
}

func (o *Order) AddItem(productID ProductID, quantity int, price Money) error {
    if o.Status != OrderStatusDraft {
        return errors.New("cannot add items to non-draft order")
    }
    
    if quantity <= 0 {
        return errors.New("quantity must be positive")
    }
    
    if price.IsZero() {
        return errors.New("price cannot be zero")
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
```

#### 3. Control Access
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

// Provide controlled access to internal objects
func (o *Order) GetItems() []OrderItem {
    // Return a copy to prevent external modification
    items := make([]OrderItem, len(o.Items))
    copy(items, o.Items)
    return items
}

func (o *Order) GetItemCount() int {
    return len(o.Items)
}

func (o *Order) HasItem(productID ProductID) bool {
    for _, item := range o.Items {
        if item.ProductID == productID {
            return true
        }
    }
    return false
}

// Private method - only accessible within the aggregate
func (o *Order) recalculateTotal() {
    total := Money{Amount: 0, Currency: "USD"}
    for _, item := range o.Items {
        itemTotal := item.Price.Multiply(float64(item.Quantity))
        total = total.Add(itemTotal)
    }
    o.TotalAmount = total
}
```

### Aggregate Examples

#### Order Aggregate
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

func NewOrder(customerID CustomerID) *Order {
    now := time.Now()
    return &Order{
        ID:         NewOrderID(),
        CustomerID: customerID,
        Items:      []OrderItem{},
        Status:     OrderStatusDraft,
        TotalAmount: Money{Amount: 0, Currency: "USD"},
        CreatedAt:  now,
        UpdatedAt:  now,
    }
}

func (o *Order) AddItem(productID ProductID, quantity int, price Money) error {
    if o.Status != OrderStatusDraft {
        return errors.New("cannot add items to non-draft order")
    }
    
    if quantity <= 0 {
        return errors.New("quantity must be positive")
    }
    
    if price.IsZero() {
        return errors.New("price cannot be zero")
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

func (o *Order) RemoveItem(productID ProductID) error {
    if o.Status != OrderStatusDraft {
        return errors.New("cannot remove items from non-draft order")
    }
    
    for i, item := range o.Items {
        if item.ProductID == productID {
            o.Items = append(o.Items[:i], o.Items[i+1:]...)
            o.recalculateTotal()
            o.UpdatedAt = time.Now()
            return nil
        }
    }
    
    return errors.New("item not found")
}

func (o *Order) Confirm() error {
    if o.Status != OrderStatusDraft {
        return errors.New("only draft orders can be confirmed")
    }
    
    if len(o.Items) == 0 {
        return errors.New("cannot confirm empty order")
    }
    
    if o.TotalAmount.IsZero() {
        return errors.New("cannot confirm order with zero total")
    }
    
    o.Status = OrderStatusConfirmed
    o.UpdatedAt = time.Now()
    
    return nil
}

func (o *Order) Ship() error {
    if o.Status != OrderStatusConfirmed {
        return errors.New("only confirmed orders can be shipped")
    }
    
    o.Status = OrderStatusShipped
    o.UpdatedAt = time.Now()
    
    return nil
}

func (o *Order) Deliver() error {
    if o.Status != OrderStatusShipped {
        return errors.New("only shipped orders can be delivered")
    }
    
    o.Status = OrderStatusDelivered
    o.UpdatedAt = time.Now()
    
    return nil
}

func (o *Order) Cancel() error {
    if o.Status == OrderStatusDelivered {
        return errors.New("cannot cancel delivered order")
    }
    
    o.Status = OrderStatusCancelled
    o.UpdatedAt = time.Now()
    
    return nil
}

func (o *Order) GetItems() []OrderItem {
    items := make([]OrderItem, len(o.Items))
    copy(items, o.Items)
    return items
}

func (o *Order) GetItemCount() int {
    return len(o.Items)
}

func (o *Order) HasItem(productID ProductID) bool {
    for _, item := range o.Items {
        if item.ProductID == productID {
            return true
        }
    }
    return false
}

func (o *Order) GetTotalAmount() Money {
    return o.TotalAmount
}

func (o *Order) IsDraft() bool {
    return o.Status == OrderStatusDraft
}

func (o *Order) IsConfirmed() bool {
    return o.Status == OrderStatusConfirmed
}

func (o *Order) IsShipped() bool {
    return o.Status == OrderStatusShipped
}

func (o *Order) IsDelivered() bool {
    return o.Status == OrderStatusDelivered
}

func (o *Order) IsCancelled() bool {
    return o.Status == OrderStatusCancelled
}

func (o *Order) recalculateTotal() {
    total := Money{Amount: 0, Currency: "USD"}
    for _, item := range o.Items {
        itemTotal := item.Price.Multiply(float64(item.Quantity))
        total = total.Add(itemTotal)
    }
    o.TotalAmount = total
}
```

#### User Aggregate
```go
type User struct {
    ID          UserID
    Email       Email
    Name        Name
    Status      UserStatus
    Preferences UserPreferences
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

func NewUser(email Email, name Name) *User {
    now := time.Now()
    return &User{
        ID:          NewUserID(),
        Email:       email,
        Name:        name,
        Status:      UserStatusPending,
        Preferences: UserPreferences{},
        CreatedAt:   now,
        UpdatedAt:   now,
    }
}

func (u *User) Activate() error {
    if u.Status == UserStatusActive {
        return errors.New("user is already active")
    }
    
    u.Status = UserStatusActive
    u.UpdatedAt = time.Now()
    
    return nil
}

func (u *User) Deactivate() error {
    if u.Status == UserStatusInactive {
        return errors.New("user is already inactive")
    }
    
    u.Status = UserStatusInactive
    u.UpdatedAt = time.Now()
    
    return nil
}

func (u *User) ChangeEmail(newEmail Email) error {
    if !newEmail.IsValid() {
        return errors.New("invalid email")
    }
    
    u.Email = newEmail
    u.UpdatedAt = time.Now()
    
    return nil
}

func (u *User) UpdateName(newName Name) error {
    if err := newName.Validate(); err != nil {
        return err
    }
    
    u.Name = newName
    u.UpdatedAt = time.Now()
    
    return nil
}

func (u *User) UpdatePreferences(prefs UserPreferences) error {
    if err := prefs.Validate(); err != nil {
        return err
    }
    
    u.Preferences = prefs
    u.UpdatedAt = time.Now()
    
    return nil
}

func (u *User) IsActive() bool {
    return u.Status == UserStatusActive
}

func (u *User) IsPending() bool {
    return u.Status == UserStatusPending
}

func (u *User) IsInactive() bool {
    return u.Status == UserStatusInactive
}
```

## 🎯 Design Guidelines

### 1. Entity Guidelines
- **Strong identity** - unique, immutable identifier
- **Rich behavior** - business logic in the entity
- **Lifecycle management** - creation, modification, deletion
- **Validation** - enforce business rules
- **Encapsulation** - control access to internal state

### 2. Value Object Guidelines
- **Immutability** - cannot be changed after creation
- **Validation** - validate on creation
- **Equality** - compare by value, not identity
- **Rich behavior** - methods for common operations
- **Encapsulation** - hide internal representation

### 3. Aggregate Guidelines
- **Single root** - one aggregate root per aggregate
- **Invariant enforcement** - business rules within aggregate
- **Controlled access** - only through aggregate root
- **Transaction boundary** - consistency within aggregate
- **Size limits** - keep aggregates small and focused

## 🔗 Next Steps

Now that you understand aggregates, entities, and value objects, we'll move to:
1. **Domain Events** - Event-driven architecture
2. **Service Design** - Applying DDD to microservices
3. **Implementation** - Building DDD-based services
4. **Practice** - Hands-on exercises

**Ready to continue? Let's learn about domain events!** 🚀
