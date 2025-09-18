package service

import (
	"order-service/domain"
	"order-service/repository"
)

// OrderService defines the interface for order business logic
type OrderService interface {
	CreateOrder(customerID domain.CustomerID) (*domain.Order, error)
	GetOrder(id domain.OrderID) (*domain.Order, error)
	GetAllOrders() ([]domain.Order, error)
	GetOrdersByCustomerID(customerID domain.CustomerID) ([]domain.Order, error)
	AddItemToOrder(orderID domain.OrderID, productID domain.ProductID, quantity int, price domain.Money) error
	RemoveItemFromOrder(orderID domain.OrderID, productID domain.ProductID) error
	UpdateItemQuantity(orderID domain.OrderID, productID domain.ProductID, quantity int) error
	ConfirmOrder(orderID domain.OrderID) error
	ShipOrder(orderID domain.OrderID, trackingNumber string) error
	DeliverOrder(orderID domain.OrderID) error
	CancelOrder(orderID domain.OrderID, reason string) error
	DeleteOrder(orderID domain.OrderID) error
}

// OrderServiceImpl implements OrderService
type OrderServiceImpl struct {
	orderRepository repository.OrderRepository
	eventPublisher  EventPublisher
	logger          Logger
}

// EventPublisher defines the interface for publishing domain events
type EventPublisher interface {
	Publish(events []domain.DomainEvent) error
}

// Logger defines the interface for logging
type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	WithField(key string, value interface{}) Logger
	WithError(err error) Logger
}

// NewOrderService creates a new order service
func NewOrderService(
	orderRepository repository.OrderRepository,
	eventPublisher EventPublisher,
	logger Logger,
) OrderService {
	return &OrderServiceImpl{
		orderRepository: orderRepository,
		eventPublisher:  eventPublisher,
		logger:          logger,
	}
}

// CreateOrder creates a new order
func (s *OrderServiceImpl) CreateOrder(customerID domain.CustomerID) (*domain.Order, error) {
	s.logger.WithField("customer_id", customerID.String()).Info("Creating order")
	
	// Create order aggregate
	order, err := domain.NewOrder(customerID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to create order")
		return nil, err
	}
	
	// Save to repository
	if err := s.orderRepository.Save(order); err != nil {
		s.logger.WithError(err).Error("Failed to save order")
		return nil, err
	}
	
	// Publish domain events
	events := order.GetEvents()
	if err := s.eventPublisher.Publish(events); err != nil {
		s.logger.WithError(err).Error("Failed to publish events")
		return nil, err
	}
	
	// Clear events
	order.ClearEvents()
	
	s.logger.WithField("order_id", order.ID.String()).Info("Order created successfully")
	return order, nil
}

// GetOrder gets an order by ID
func (s *OrderServiceImpl) GetOrder(id domain.OrderID) (*domain.Order, error) {
	s.logger.WithField("order_id", id.String()).Info("Getting order")
	
	order, err := s.orderRepository.FindByID(id)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get order")
		return nil, err
	}
	
	return order, nil
}

// GetAllOrders gets all orders
func (s *OrderServiceImpl) GetAllOrders() ([]domain.Order, error) {
	s.logger.Info("Getting all orders")
	
	orders, err := s.orderRepository.FindAll()
	if err != nil {
		s.logger.WithError(err).Error("Failed to get orders")
		return nil, err
	}
	
	return orders, nil
}

// GetOrdersByCustomerID gets orders by customer ID
func (s *OrderServiceImpl) GetOrdersByCustomerID(customerID domain.CustomerID) ([]domain.Order, error) {
	s.logger.WithField("customer_id", customerID.String()).Info("Getting orders by customer ID")
	
	orders, err := s.orderRepository.FindByCustomerID(customerID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get orders by customer ID")
		return nil, err
	}
	
	return orders, nil
}

// AddItemToOrder adds an item to an order
func (s *OrderServiceImpl) AddItemToOrder(orderID domain.OrderID, productID domain.ProductID, quantity int, price domain.Money) error {
	s.logger.WithFields(map[string]interface{}{
		"order_id":   orderID.String(),
		"product_id": productID.String(),
		"quantity":   quantity,
	}).Info("Adding item to order")
	
	// Get existing order
	order, err := s.orderRepository.FindByID(orderID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get order for adding item")
		return err
	}
	
	// Add item to order
	if err := order.AddItem(productID, quantity, price); err != nil {
		s.logger.WithError(err).Error("Failed to add item to order")
		return err
	}
	
	// Save to repository
	if err := s.orderRepository.Save(order); err != nil {
		s.logger.WithError(err).Error("Failed to save order with new item")
		return err
	}
	
	s.logger.WithField("order_id", orderID.String()).Info("Item added to order successfully")
	return nil
}

// RemoveItemFromOrder removes an item from an order
func (s *OrderServiceImpl) RemoveItemFromOrder(orderID domain.OrderID, productID domain.ProductID) error {
	s.logger.WithFields(map[string]interface{}{
		"order_id":   orderID.String(),
		"product_id": productID.String(),
	}).Info("Removing item from order")
	
	// Get existing order
	order, err := s.orderRepository.FindByID(orderID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get order for removing item")
		return err
	}
	
	// Remove item from order
	if err := order.RemoveItem(productID); err != nil {
		s.logger.WithError(err).Error("Failed to remove item from order")
		return err
	}
	
	// Save to repository
	if err := s.orderRepository.Save(order); err != nil {
		s.logger.WithError(err).Error("Failed to save order after removing item")
		return err
	}
	
	s.logger.WithField("order_id", orderID.String()).Info("Item removed from order successfully")
	return nil
}

// UpdateItemQuantity updates the quantity of an item in an order
func (s *OrderServiceImpl) UpdateItemQuantity(orderID domain.OrderID, productID domain.ProductID, quantity int) error {
	s.logger.WithFields(map[string]interface{}{
		"order_id":   orderID.String(),
		"product_id": productID.String(),
		"quantity":   quantity,
	}).Info("Updating item quantity in order")
	
	// Get existing order
	order, err := s.orderRepository.FindByID(orderID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get order for updating item quantity")
		return err
	}
	
	// Update item quantity
	if err := order.UpdateItemQuantity(productID, quantity); err != nil {
		s.logger.WithError(err).Error("Failed to update item quantity in order")
		return err
	}
	
	// Save to repository
	if err := s.orderRepository.Save(order); err != nil {
		s.logger.WithError(err).Error("Failed to save order after updating item quantity")
		return err
	}
	
	s.logger.WithField("order_id", orderID.String()).Info("Item quantity updated in order successfully")
	return nil
}

// ConfirmOrder confirms an order
func (s *OrderServiceImpl) ConfirmOrder(orderID domain.OrderID) error {
	s.logger.WithField("order_id", orderID.String()).Info("Confirming order")
	
	// Get existing order
	order, err := s.orderRepository.FindByID(orderID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get order for confirmation")
		return err
	}
	
	// Confirm order
	if err := order.Confirm(); err != nil {
		s.logger.WithError(err).Error("Failed to confirm order")
		return err
	}
	
	// Save to repository
	if err := s.orderRepository.Save(order); err != nil {
		s.logger.WithError(err).Error("Failed to save confirmed order")
		return err
	}
	
	// Publish domain events
	events := order.GetEvents()
	if err := s.eventPublisher.Publish(events); err != nil {
		s.logger.WithError(err).Error("Failed to publish events")
		return err
	}
	
	// Clear events
	order.ClearEvents()
	
	s.logger.WithField("order_id", orderID.String()).Info("Order confirmed successfully")
	return nil
}

// ShipOrder ships an order
func (s *OrderServiceImpl) ShipOrder(orderID domain.OrderID, trackingNumber string) error {
	s.logger.WithFields(map[string]interface{}{
		"order_id":        orderID.String(),
		"tracking_number": trackingNumber,
	}).Info("Shipping order")
	
	// Get existing order
	order, err := s.orderRepository.FindByID(orderID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get order for shipping")
		return err
	}
	
	// Ship order
	if err := order.Ship(trackingNumber); err != nil {
		s.logger.WithError(err).Error("Failed to ship order")
		return err
	}
	
	// Save to repository
	if err := s.orderRepository.Save(order); err != nil {
		s.logger.WithError(err).Error("Failed to save shipped order")
		return err
	}
	
	// Publish domain events
	events := order.GetEvents()
	if err := s.eventPublisher.Publish(events); err != nil {
		s.logger.WithError(err).Error("Failed to publish events")
		return err
	}
	
	// Clear events
	order.ClearEvents()
	
	s.logger.WithField("order_id", orderID.String()).Info("Order shipped successfully")
	return nil
}

// DeliverOrder delivers an order
func (s *OrderServiceImpl) DeliverOrder(orderID domain.OrderID) error {
	s.logger.WithField("order_id", orderID.String()).Info("Delivering order")
	
	// Get existing order
	order, err := s.orderRepository.FindByID(orderID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get order for delivery")
		return err
	}
	
	// Deliver order
	if err := order.Deliver(); err != nil {
		s.logger.WithError(err).Error("Failed to deliver order")
		return err
	}
	
	// Save to repository
	if err := s.orderRepository.Save(order); err != nil {
		s.logger.WithError(err).Error("Failed to save delivered order")
		return err
	}
	
	// Publish domain events
	events := order.GetEvents()
	if err := s.eventPublisher.Publish(events); err != nil {
		s.logger.WithError(err).Error("Failed to publish events")
		return err
	}
	
	// Clear events
	order.ClearEvents()
	
	s.logger.WithField("order_id", orderID.String()).Info("Order delivered successfully")
	return nil
}

// CancelOrder cancels an order
func (s *OrderServiceImpl) CancelOrder(orderID domain.OrderID, reason string) error {
	s.logger.WithFields(map[string]interface{}{
		"order_id": orderID.String(),
		"reason":   reason,
	}).Info("Cancelling order")
	
	// Get existing order
	order, err := s.orderRepository.FindByID(orderID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get order for cancellation")
		return err
	}
	
	// Cancel order
	if err := order.Cancel(reason); err != nil {
		s.logger.WithError(err).Error("Failed to cancel order")
		return err
	}
	
	// Save to repository
	if err := s.orderRepository.Save(order); err != nil {
		s.logger.WithError(err).Error("Failed to save cancelled order")
		return err
	}
	
	// Publish domain events
	events := order.GetEvents()
	if err := s.eventPublisher.Publish(events); err != nil {
		s.logger.WithError(err).Error("Failed to publish events")
		return err
	}
	
	// Clear events
	order.ClearEvents()
	
	s.logger.WithField("order_id", orderID.String()).Info("Order cancelled successfully")
	return nil
}

// DeleteOrder deletes an order
func (s *OrderServiceImpl) DeleteOrder(orderID domain.OrderID) error {
	s.logger.WithField("order_id", orderID.String()).Info("Deleting order")
	
	// Check if order exists
	_, err := s.orderRepository.FindByID(orderID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get order for deletion")
		return err
	}
	
	// Delete order
	if err := s.orderRepository.Delete(orderID); err != nil {
		s.logger.WithError(err).Error("Failed to delete order")
		return err
	}
	
	s.logger.WithField("order_id", orderID.String()).Info("Order deleted successfully")
	return nil
}
