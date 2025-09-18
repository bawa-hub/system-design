package service

import (
	"product-service/domain"
	"product-service/repository"
)

// ProductService defines the interface for product business logic
type ProductService interface {
	CreateProduct(name, description string, price domain.Money, category string) (*domain.Product, error)
	GetProduct(id domain.ProductID) (*domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
	GetProductsByCategory(category string) ([]domain.Product, error)
	GetActiveProducts() ([]domain.Product, error)
	UpdateProduct(id domain.ProductID, name, description string, price domain.Money, category string) error
	DiscontinueProduct(id domain.ProductID) error
	ActivateProduct(id domain.ProductID) error
	DeactivateProduct(id domain.ProductID) error
	DeleteProduct(id domain.ProductID) error
}

// ProductServiceImpl implements ProductService
type ProductServiceImpl struct {
	productRepository repository.ProductRepository
	eventPublisher    EventPublisher
	logger            Logger
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

// NewProductService creates a new product service
func NewProductService(
	productRepository repository.ProductRepository,
	eventPublisher EventPublisher,
	logger Logger,
) ProductService {
	return &ProductServiceImpl{
		productRepository: productRepository,
		eventPublisher:    eventPublisher,
		logger:            logger,
	}
}

// CreateProduct creates a new product
func (s *ProductServiceImpl) CreateProduct(name, description string, price domain.Money, category string) (*domain.Product, error) {
	s.logger.WithField("name", name).Info("Creating product")
	
	// Create product aggregate
	product, err := domain.NewProduct(name, description, price, category)
	if err != nil {
		s.logger.WithError(err).Error("Failed to create product")
		return nil, err
	}
	
	// Save to repository
	if err := s.productRepository.Save(product); err != nil {
		s.logger.WithError(err).Error("Failed to save product")
		return nil, err
	}
	
	// Publish domain events
	events := product.GetEvents()
	if err := s.eventPublisher.Publish(events); err != nil {
		s.logger.WithError(err).Error("Failed to publish events")
		return nil, err
	}
	
	// Clear events
	product.ClearEvents()
	
	s.logger.WithField("product_id", product.ID.String()).Info("Product created successfully")
	return product, nil
}

// GetProduct gets a product by ID
func (s *ProductServiceImpl) GetProduct(id domain.ProductID) (*domain.Product, error) {
	s.logger.WithField("product_id", id.String()).Info("Getting product")
	
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get product")
		return nil, err
	}
	
	return product, nil
}

// GetAllProducts gets all products
func (s *ProductServiceImpl) GetAllProducts() ([]domain.Product, error) {
	s.logger.Info("Getting all products")
	
	products, err := s.productRepository.FindAll()
	if err != nil {
		s.logger.WithError(err).Error("Failed to get products")
		return nil, err
	}
	
	return products, nil
}

// GetProductsByCategory gets products by category
func (s *ProductServiceImpl) GetProductsByCategory(category string) ([]domain.Product, error) {
	s.logger.WithField("category", category).Info("Getting products by category")
	
	products, err := s.productRepository.FindByCategory(category)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get products by category")
		return nil, err
	}
	
	return products, nil
}

// GetActiveProducts gets all active products
func (s *ProductServiceImpl) GetActiveProducts() ([]domain.Product, error) {
	s.logger.Info("Getting active products")
	
	products, err := s.productRepository.FindActive()
	if err != nil {
		s.logger.WithError(err).Error("Failed to get active products")
		return nil, err
	}
	
	return products, nil
}

// UpdateProduct updates a product
func (s *ProductServiceImpl) UpdateProduct(id domain.ProductID, name, description string, price domain.Money, category string) error {
	s.logger.WithField("product_id", id.String()).Info("Updating product")
	
	// Get existing product
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get product for update")
		return err
	}
	
	// Update product
	if err := product.UpdateProduct(name, description, price, category); err != nil {
		s.logger.WithError(err).Error("Failed to update product")
		return err
	}
	
	// Save to repository
	if err := s.productRepository.Save(product); err != nil {
		s.logger.WithError(err).Error("Failed to save updated product")
		return err
	}
	
	// Publish domain events
	events := product.GetEvents()
	if err := s.eventPublisher.Publish(events); err != nil {
		s.logger.WithError(err).Error("Failed to publish events")
		return err
	}
	
	// Clear events
	product.ClearEvents()
	
	s.logger.WithField("product_id", id.String()).Info("Product updated successfully")
	return nil
}

// DiscontinueProduct discontinues a product
func (s *ProductServiceImpl) DiscontinueProduct(id domain.ProductID) error {
	s.logger.WithField("product_id", id.String()).Info("Discontinuing product")
	
	// Get existing product
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get product for discontinuation")
		return err
	}
	
	// Discontinue product
	if err := product.Discontinue(); err != nil {
		s.logger.WithError(err).Error("Failed to discontinue product")
		return err
	}
	
	// Save to repository
	if err := s.productRepository.Save(product); err != nil {
		s.logger.WithError(err).Error("Failed to save discontinued product")
		return err
	}
	
	// Publish domain events
	events := product.GetEvents()
	if err := s.eventPublisher.Publish(events); err != nil {
		s.logger.WithError(err).Error("Failed to publish events")
		return err
	}
	
	// Clear events
	product.ClearEvents()
	
	s.logger.WithField("product_id", id.String()).Info("Product discontinued successfully")
	return nil
}

// ActivateProduct activates a product
func (s *ProductServiceImpl) ActivateProduct(id domain.ProductID) error {
	s.logger.WithField("product_id", id.String()).Info("Activating product")
	
	// Get existing product
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get product for activation")
		return err
	}
	
	// Activate product
	if err := product.Activate(); err != nil {
		s.logger.WithError(err).Error("Failed to activate product")
		return err
	}
	
	// Save to repository
	if err := s.productRepository.Save(product); err != nil {
		s.logger.WithError(err).Error("Failed to save activated product")
		return err
	}
	
	s.logger.WithField("product_id", id.String()).Info("Product activated successfully")
	return nil
}

// DeactivateProduct deactivates a product
func (s *ProductServiceImpl) DeactivateProduct(id domain.ProductID) error {
	s.logger.WithField("product_id", id.String()).Info("Deactivating product")
	
	// Get existing product
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get product for deactivation")
		return err
	}
	
	// Deactivate product
	if err := product.Deactivate(); err != nil {
		s.logger.WithError(err).Error("Failed to deactivate product")
		return err
	}
	
	// Save to repository
	if err := s.productRepository.Save(product); err != nil {
		s.logger.WithError(err).Error("Failed to save deactivated product")
		return err
	}
	
	s.logger.WithField("product_id", id.String()).Info("Product deactivated successfully")
	return nil
}

// DeleteProduct deletes a product
func (s *ProductServiceImpl) DeleteProduct(id domain.ProductID) error {
	s.logger.WithField("product_id", id.String()).Info("Deleting product")
	
	// Check if product exists
	_, err := s.productRepository.FindByID(id)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get product for deletion")
		return err
	}
	
	// Delete product
	if err := s.productRepository.Delete(id); err != nil {
		s.logger.WithError(err).Error("Failed to delete product")
		return err
	}
	
	s.logger.WithField("product_id", id.String()).Info("Product deleted successfully")
	return nil
}
