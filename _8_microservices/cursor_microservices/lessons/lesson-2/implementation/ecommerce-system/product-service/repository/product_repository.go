package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"product-service/domain"
)

// ProductRepository defines the interface for product data access
type ProductRepository interface {
	Save(product *domain.Product) error
	FindByID(id domain.ProductID) (*domain.Product, error)
	FindAll() ([]domain.Product, error)
	FindByCategory(category string) ([]domain.Product, error)
	FindActive() ([]domain.Product, error)
	Delete(id domain.ProductID) error
}

// ProductRepositoryImpl implements ProductRepository
type ProductRepositoryImpl struct {
	db *sql.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

// Save saves a product to the database
func (r *ProductRepositoryImpl) Save(product *domain.Product) error {
	query := `
		INSERT INTO products (id, name, description, price_amount, price_currency, category, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (id) DO UPDATE SET
		name = $2, description = $3, price_amount = $4, price_currency = $5, 
		category = $6, status = $7, updated_at = $9
	`
	
	_, err := r.db.Exec(query,
		product.ID.String(),
		product.Name,
		product.Description,
		product.Price.Amount(),
		product.Price.Currency(),
		product.Category,
		string(product.Status),
		product.CreatedAt,
		product.UpdatedAt,
	)
	
	return err
}

// FindByID finds a product by ID
func (r *ProductRepositoryImpl) FindByID(id domain.ProductID) (*domain.Product, error) {
	query := `
		SELECT id, name, description, price_amount, price_currency, category, status, created_at, updated_at
		FROM products
		WHERE id = $1
	`
	
	var product domain.Product
	var priceAmount int64
	var priceCurrency string
	var status string
	
	err := r.db.QueryRow(query, id.String()).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&priceAmount,
		&priceCurrency,
		&product.Category,
		&status,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	
	// Create Money value object
	price, err := domain.NewMoney(priceAmount, priceCurrency)
	if err != nil {
		return nil, err
	}
	product.Price = price
	
	// Set status
	product.Status = domain.ProductStatus(status)
	
	return &product, nil
}

// FindAll finds all products
func (r *ProductRepositoryImpl) FindAll() ([]domain.Product, error) {
	query := `
		SELECT id, name, description, price_amount, price_currency, category, status, created_at, updated_at
		FROM products
		ORDER BY created_at DESC
	`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		var priceAmount int64
		var priceCurrency string
		var status string
		
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&priceAmount,
			&priceCurrency,
			&product.Category,
			&status,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		
		// Create Money value object
		price, err := domain.NewMoney(priceAmount, priceCurrency)
		if err != nil {
			return nil, err
		}
		product.Price = price
		
		// Set status
		product.Status = domain.ProductStatus(status)
		
		products = append(products, product)
	}
	
	return products, nil
}

// FindByCategory finds products by category
func (r *ProductRepositoryImpl) FindByCategory(category string) ([]domain.Product, error) {
	query := `
		SELECT id, name, description, price_amount, price_currency, category, status, created_at, updated_at
		FROM products
		WHERE category = $1
		ORDER BY created_at DESC
	`
	
	rows, err := r.db.Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		var priceAmount int64
		var priceCurrency string
		var status string
		
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&priceAmount,
			&priceCurrency,
			&product.Category,
			&status,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		
		// Create Money value object
		price, err := domain.NewMoney(priceAmount, priceCurrency)
		if err != nil {
			return nil, err
		}
		product.Price = price
		
		// Set status
		product.Status = domain.ProductStatus(status)
		
		products = append(products, product)
	}
	
	return products, nil
}

// FindActive finds all active products
func (r *ProductRepositoryImpl) FindActive() ([]domain.Product, error) {
	query := `
		SELECT id, name, description, price_amount, price_currency, category, status, created_at, updated_at
		FROM products
		WHERE status = 'active'
		ORDER BY created_at DESC
	`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		var priceAmount int64
		var priceCurrency string
		var status string
		
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&priceAmount,
			&priceCurrency,
			&product.Category,
			&status,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		
		// Create Money value object
		price, err := domain.NewMoney(priceAmount, priceCurrency)
		if err != nil {
			return nil, err
		}
		product.Price = price
		
		// Set status
		product.Status = domain.ProductStatus(status)
		
		products = append(products, product)
	}
	
	return products, nil
}

// Delete deletes a product
func (r *ProductRepositoryImpl) Delete(id domain.ProductID) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id.String())
	return err
}
