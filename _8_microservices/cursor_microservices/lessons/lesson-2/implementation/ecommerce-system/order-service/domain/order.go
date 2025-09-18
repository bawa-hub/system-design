package domain

import (
	"errors"
	"time"
)

// Order represents an order in the system
type Order struct {
	ID          OrderID
	CustomerID  CustomerID
	Items       []OrderItem
	Status      OrderStatus
	TotalAmount Money
	CreatedAt   time.Time
	UpdatedAt   time.Time
	events      []DomainEvent
}

// OrderItem represents an item in an order
type OrderItem struct {
	ProductID ProductID
	Quantity  int
	Price     Money
}

// OrderStatus represents the status of an order
type OrderStatus string

const (
	OrderStatusDraft      OrderStatus = "draft"
	OrderStatusConfirmed  OrderStatus = "confirmed"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

// NewOrder creates a new order
func NewOrder(customerID CustomerID) (*Order, error) {
	if customerID.String() == "" {
		return nil, errors.New("customer ID cannot be empty")
	}
	
	now := time.Now()
	order := &Order{
		ID:          NewOrderID(),
		CustomerID:  customerID,
		Items:       []OrderItem{},
		Status:      OrderStatusDraft,
		TotalAmount: Money{Amount: 0, Currency: "USD"},
		CreatedAt:   now,
		UpdatedAt:   now,
		events:      []DomainEvent{},
	}
	
	// Publish domain event
	event := NewOrderCreated(order.ID, order.CustomerID, order.Items, order.TotalAmount)
	order.events = append(order.events, event)
	
	return order, nil
}

// AddItem adds an item to the order
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
	
	// Check if item already exists
	for i, item := range o.Items {
		if item.ProductID == productID {
			// Update existing item
			o.Items[i].Quantity += quantity
			o.recalculateTotal()
			o.UpdatedAt = time.Now()
			return nil
		}
	}
	
	// Add new item
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

// RemoveItem removes an item from the order
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

// UpdateItemQuantity updates the quantity of an item
func (o *Order) UpdateItemQuantity(productID ProductID, quantity int) error {
	if o.Status != OrderStatusDraft {
		return errors.New("cannot update items in non-draft order")
	}
	
	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}
	
	for i, item := range o.Items {
		if item.ProductID == productID {
			o.Items[i].Quantity = quantity
			o.recalculateTotal()
			o.UpdatedAt = time.Now()
			return nil
		}
	}
	
	return errors.New("item not found")
}

// Confirm confirms the order
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
	
	// Publish domain event
	event := NewOrderConfirmed(o.ID, o.CustomerID, o.TotalAmount)
	o.events = append(o.events, event)
	
	return nil
}

// Ship ships the order
func (o *Order) Ship(trackingNumber string) error {
	if o.Status != OrderStatusConfirmed {
		return errors.New("only confirmed orders can be shipped")
	}
	
	if trackingNumber == "" {
		return errors.New("tracking number is required")
	}
	
	o.Status = OrderStatusShipped
	o.UpdatedAt = time.Now()
	
	// Publish domain event
	event := NewOrderShipped(o.ID, o.CustomerID, trackingNumber)
	o.events = append(o.events, event)
	
	return nil
}

// Deliver delivers the order
func (o *Order) Deliver() error {
	if o.Status != OrderStatusShipped {
		return errors.New("only shipped orders can be delivered")
	}
	
	o.Status = OrderStatusDelivered
	o.UpdatedAt = time.Now()
	
	// Publish domain event
	event := NewOrderDelivered(o.ID, o.CustomerID)
	o.events = append(o.events, event)
	
	return nil
}

// Cancel cancels the order
func (o *Order) Cancel(reason string) error {
	if o.Status == OrderStatusDelivered {
		return errors.New("cannot cancel delivered order")
	}
	
	if o.Status == OrderStatusCancelled {
		return errors.New("order is already cancelled")
	}
	
	o.Status = OrderStatusCancelled
	o.UpdatedAt = time.Now()
	
	// Publish domain event
	event := NewOrderCancelled(o.ID, o.CustomerID, reason)
	o.events = append(o.events, event)
	
	return nil
}

// GetItems returns a copy of order items
func (o *Order) GetItems() []OrderItem {
	items := make([]OrderItem, len(o.Items))
	copy(items, o.Items)
	return items
}

// GetItemCount returns the number of items in the order
func (o *Order) GetItemCount() int {
	return len(o.Items)
}

// HasItem checks if the order has a specific product
func (o *Order) HasItem(productID ProductID) bool {
	for _, item := range o.Items {
		if item.ProductID == productID {
			return true
		}
	}
	return false
}

// GetTotalAmount returns the total amount of the order
func (o *Order) GetTotalAmount() Money {
	return o.TotalAmount
}

// IsDraft checks if the order is in draft status
func (o *Order) IsDraft() bool {
	return o.Status == OrderStatusDraft
}

// IsConfirmed checks if the order is confirmed
func (o *Order) IsConfirmed() bool {
	return o.Status == OrderStatusConfirmed
}

// IsShipped checks if the order is shipped
func (o *Order) IsShipped() bool {
	return o.Status == OrderStatusShipped
}

// IsDelivered checks if the order is delivered
func (o *Order) IsDelivered() bool {
	return o.Status == OrderStatusDelivered
}

// IsCancelled checks if the order is cancelled
func (o *Order) IsCancelled() bool {
	return o.Status == OrderStatusCancelled
}

// GetEvents returns domain events
func (o *Order) GetEvents() []DomainEvent {
	return o.events
}

// ClearEvents clears domain events
func (o *Order) ClearEvents() {
	o.events = []DomainEvent{}
}

// recalculateTotal recalculates the total amount
func (o *Order) recalculateTotal() {
	total := Money{Amount: 0, Currency: "USD"}
	for _, item := range o.Items {
		itemTotal, err := item.Price.Multiply(float64(item.Quantity))
		if err == nil {
			total, _ = total.Add(itemTotal)
		}
	}
	o.TotalAmount = total
}

// ApplyEvent applies a domain event (for event sourcing)
func (o *Order) ApplyEvent(event DomainEvent) error {
	switch e := event.(type) {
	case OrderCreated:
		o.ID = e.BaseEvent.AggregateID
		o.CustomerID = e.CustomerID
		o.Items = e.Items
		o.TotalAmount = e.TotalAmount
		o.Status = OrderStatusDraft
		o.CreatedAt = e.OccurredAt
		o.UpdatedAt = e.OccurredAt
		
	case OrderConfirmed:
		o.Status = OrderStatusConfirmed
		o.UpdatedAt = e.OccurredAt
		
	case OrderShipped:
		o.Status = OrderStatusShipped
		o.UpdatedAt = e.OccurredAt
		
	case OrderDelivered:
		o.Status = OrderStatusDelivered
		o.UpdatedAt = e.OccurredAt
		
	case OrderCancelled:
		o.Status = OrderStatusCancelled
		o.UpdatedAt = e.OccurredAt
		
	default:
		return errors.New("unknown event type")
	}
	
	return nil
}
