package domain

import (
	"time"
)

// DomainEvent represents a domain event
type DomainEvent interface {
	EventType() string
	OccurredAt() time.Time
	AggregateID() string
	EventVersion() int
}

// BaseEvent provides common event functionality
type BaseEvent struct {
	EventType   string
	OccurredAt  time.Time
	AggregateID string
	Version     int
}

func (e BaseEvent) EventType() string {
	return e.EventType
}

func (e BaseEvent) OccurredAt() time.Time {
	return e.OccurredAt
}

func (e BaseEvent) AggregateID() string {
	return e.AggregateID
}

func (e BaseEvent) EventVersion() int {
	return e.Version
}

// User Events
type UserCreated struct {
	BaseEvent
	Email Email
	Name  Name
}

func NewUserCreated(userID UserID, email Email, name Name) UserCreated {
	return UserCreated{
		BaseEvent: BaseEvent{
			EventType:   "user.created",
			OccurredAt:  time.Now(),
			AggregateID: string(userID),
			Version:     1,
		},
		Email: email,
		Name:  name,
	}
}

type UserActivated struct {
	BaseEvent
	Email Email
}

func NewUserActivated(userID UserID, email Email) UserActivated {
	return UserActivated{
		BaseEvent: BaseEvent{
			EventType:   "user.activated",
			OccurredAt:  time.Now(),
			AggregateID: string(userID),
			Version:     1,
		},
		Email: email,
	}
}

type UserDeactivated struct {
	BaseEvent
	Email Email
}

func NewUserDeactivated(userID UserID, email Email) UserDeactivated {
	return UserDeactivated{
		BaseEvent: BaseEvent{
			EventType:   "user.deactivated",
			OccurredAt:  time.Now(),
			AggregateID: string(userID),
			Version:     1,
		},
		Email: email,
	}
}

// Product Events
type ProductCreated struct {
	BaseEvent
	Name        string
	Description string
	Price       Money
	Category    string
}

func NewProductCreated(productID ProductID, name, description string, price Money, category string) ProductCreated {
	return ProductCreated{
		BaseEvent: BaseEvent{
			EventType:   "product.created",
			OccurredAt:  time.Now(),
			AggregateID: string(productID),
			Version:     1,
		},
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
	}
}

type ProductUpdated struct {
	BaseEvent
	Name        string
	Description string
	Price       Money
	Category    string
}

func NewProductUpdated(productID ProductID, name, description string, price Money, category string) ProductUpdated {
	return ProductUpdated{
		BaseEvent: BaseEvent{
			EventType:   "product.updated",
			OccurredAt:  time.Now(),
			AggregateID: string(productID),
			Version:     1,
		},
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
	}
}

type ProductDiscontinued struct {
	BaseEvent
	Name string
}

func NewProductDiscontinued(productID ProductID, name string) ProductDiscontinued {
	return ProductDiscontinued{
		BaseEvent: BaseEvent{
			EventType:   "product.discontinued",
			OccurredAt:  time.Now(),
			AggregateID: string(productID),
			Version:     1,
		},
		Name: name,
	}
}

// Order Events
type OrderCreated struct {
	BaseEvent
	CustomerID  CustomerID
	Items       []OrderItem
	TotalAmount Money
}

func NewOrderCreated(orderID OrderID, customerID CustomerID, items []OrderItem, totalAmount Money) OrderCreated {
	return OrderCreated{
		BaseEvent: BaseEvent{
			EventType:   "order.created",
			OccurredAt:  time.Now(),
			AggregateID: string(orderID),
			Version:     1,
		},
		CustomerID:  customerID,
		Items:       items,
		TotalAmount: totalAmount,
	}
}

type OrderConfirmed struct {
	BaseEvent
	CustomerID  CustomerID
	TotalAmount Money
}

func NewOrderConfirmed(orderID OrderID, customerID CustomerID, totalAmount Money) OrderConfirmed {
	return OrderConfirmed{
		BaseEvent: BaseEvent{
			EventType:   "order.confirmed",
			OccurredAt:  time.Now(),
			AggregateID: string(orderID),
			Version:     1,
		},
		CustomerID:  customerID,
		TotalAmount: totalAmount,
	}
}

type OrderShipped struct {
	BaseEvent
	CustomerID CustomerID
	TrackingNumber string
}

func NewOrderShipped(orderID OrderID, customerID CustomerID, trackingNumber string) OrderShipped {
	return OrderShipped{
		BaseEvent: BaseEvent{
			EventType:   "order.shipped",
			OccurredAt:  time.Now(),
			AggregateID: string(orderID),
			Version:     1,
		},
		CustomerID:    customerID,
		TrackingNumber: trackingNumber,
	}
}

type OrderDelivered struct {
	BaseEvent
	CustomerID CustomerID
}

func NewOrderDelivered(orderID OrderID, customerID CustomerID) OrderDelivered {
	return OrderDelivered{
		BaseEvent: BaseEvent{
			EventType:   "order.delivered",
			OccurredAt:  time.Now(),
			AggregateID: string(orderID),
			Version:     1,
		},
		CustomerID: customerID,
	}
}

type OrderCancelled struct {
	BaseEvent
	CustomerID CustomerID
	Reason     string
}

func NewOrderCancelled(orderID OrderID, customerID CustomerID, reason string) OrderCancelled {
	return OrderCancelled{
		BaseEvent: BaseEvent{
			EventType:   "order.cancelled",
			OccurredAt:  time.Now(),
			AggregateID: string(orderID),
			Version:     1,
		},
		CustomerID: customerID,
		Reason:     reason,
	}
}
