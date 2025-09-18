package infrastructure

import (
	"log"
	"product-service/domain"
)

// EventPublisher implements event publishing
type EventPublisher struct{}

// NewEventPublisher creates a new event publisher
func NewEventPublisher() *EventPublisher {
	return &EventPublisher{}
}

// Publish publishes domain events
func (p *EventPublisher) Publish(events []domain.DomainEvent) error {
	for _, event := range events {
		log.Printf("Publishing event: %s for aggregate: %s", event.EventType(), event.AggregateID())
		// In a real implementation, this would publish to a message broker
		// For now, we just log the events
	}
	return nil
}
