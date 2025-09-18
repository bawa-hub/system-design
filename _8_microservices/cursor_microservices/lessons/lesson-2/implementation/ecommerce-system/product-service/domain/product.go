package domain

import (
	"errors"
	"time"
)

// Product represents a product in the catalog
type Product struct {
	ID          ProductID
	Name        string
	Description string
	Price       Money
	Category    string
	Status      ProductStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	events      []DomainEvent
}

// ProductStatus represents the status of a product
type ProductStatus string

const (
	ProductStatusActive      ProductStatus = "active"
	ProductStatusInactive    ProductStatus = "inactive"
	ProductStatusDiscontinued ProductStatus = "discontinued"
)

// NewProduct creates a new product
func NewProduct(name, description string, price Money, category string) (*Product, error) {
	if name == "" {
		return nil, errors.New("product name cannot be empty")
	}
	
	if description == "" {
		return nil, errors.New("product description cannot be empty")
	}
	
	if price.IsZero() {
		return nil, errors.New("product price cannot be zero")
	}
	
	if category == "" {
		return nil, errors.New("product category cannot be empty")
	}
	
	now := time.Now()
	product := &Product{
		ID:          NewProductID(),
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
		Status:      ProductStatusActive,
		CreatedAt:   now,
		UpdatedAt:   now,
		events:      []DomainEvent{},
	}
	
	// Publish domain event
	event := NewProductCreated(product.ID, product.Name, product.Description, product.Price, product.Category)
	product.events = append(product.events, event)
	
	return product, nil
}

// UpdateProduct updates product information
func (p *Product) UpdateProduct(name, description string, price Money, category string) error {
	if p.Status == ProductStatusDiscontinued {
		return errors.New("cannot update discontinued product")
	}
	
	if name == "" {
		return errors.New("product name cannot be empty")
	}
	
	if description == "" {
		return errors.New("product description cannot be empty")
	}
	
	if price.IsZero() {
		return errors.New("product price cannot be zero")
	}
	
	if category == "" {
		return errors.New("product category cannot be empty")
	}
	
	p.Name = name
	p.Description = description
	p.Price = price
	p.Category = category
	p.UpdatedAt = time.Now()
	
	// Publish domain event
	event := NewProductUpdated(p.ID, p.Name, p.Description, p.Price, p.Category)
	p.events = append(p.events, event)
	
	return nil
}

// Discontinue discontinues the product
func (p *Product) Discontinue() error {
	if p.Status == ProductStatusDiscontinued {
		return errors.New("product is already discontinued")
	}
	
	p.Status = ProductStatusDiscontinued
	p.UpdatedAt = time.Now()
	
	// Publish domain event
	event := NewProductDiscontinued(p.ID, p.Name)
	p.events = append(p.events, event)
	
	return nil
}

// Activate activates the product
func (p *Product) Activate() error {
	if p.Status == ProductStatusActive {
		return errors.New("product is already active")
	}
	
	if p.Status == ProductStatusDiscontinued {
		return errors.New("cannot activate discontinued product")
	}
	
	p.Status = ProductStatusActive
	p.UpdatedAt = time.Now()
	
	return nil
}

// Deactivate deactivates the product
func (p *Product) Deactivate() error {
	if p.Status == ProductStatusInactive {
		return errors.New("product is already inactive")
	}
	
	if p.Status == ProductStatusDiscontinued {
		return errors.New("cannot deactivate discontinued product")
	}
	
	p.Status = ProductStatusInactive
	p.UpdatedAt = time.Now()
	
	return nil
}

// IsActive checks if the product is active
func (p *Product) IsActive() bool {
	return p.Status == ProductStatusActive
}

// IsDiscontinued checks if the product is discontinued
func (p *Product) IsDiscontinued() bool {
	return p.Status == ProductStatusDiscontinued
}

// GetEvents returns domain events
func (p *Product) GetEvents() []DomainEvent {
	return p.events
}

// ClearEvents clears domain events
func (p *Product) ClearEvents() {
	p.events = []DomainEvent{}
}

// ApplyEvent applies a domain event (for event sourcing)
func (p *Product) ApplyEvent(event DomainEvent) error {
	switch e := event.(type) {
	case ProductCreated:
		p.ID = e.BaseEvent.AggregateID
		p.Name = e.Name
		p.Description = e.Description
		p.Price = e.Price
		p.Category = e.Category
		p.Status = ProductStatusActive
		p.CreatedAt = e.OccurredAt
		p.UpdatedAt = e.OccurredAt
		
	case ProductUpdated:
		p.Name = e.Name
		p.Description = e.Description
		p.Price = e.Price
		p.Category = e.Category
		p.UpdatedAt = e.OccurredAt
		
	case ProductDiscontinued:
		p.Status = ProductStatusDiscontinued
		p.UpdatedAt = e.OccurredAt
		
	default:
		return errors.New("unknown event type")
	}
	
	return nil
}
