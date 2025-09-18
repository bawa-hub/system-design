package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"order-service/domain"
)

// OrderRepository defines the interface for order data access
type OrderRepository interface {
	Save(order *domain.Order) error
	FindByID(id domain.OrderID) (*domain.Order, error)
	FindAll() ([]domain.Order, error)
	FindByCustomerID(customerID domain.CustomerID) ([]domain.Order, error)
	Delete(id domain.OrderID) error
}

// OrderRepositoryImpl implements OrderRepository
type OrderRepositoryImpl struct {
	db *sql.DB
}

// NewOrderRepository creates a new order repository
func NewOrderRepository(db *sql.DB) OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

// Save saves an order to the database
func (r *OrderRepositoryImpl) Save(order *domain.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Save order
	orderQuery := `
		INSERT INTO orders (id, customer_id, total_amount, total_currency, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE SET
		customer_id = $2, total_amount = $3, total_currency = $4, status = $5, updated_at = $7
	`
	
	_, err = tx.Exec(orderQuery,
		order.ID.String(),
		order.CustomerID.String(),
		order.TotalAmount.Amount(),
		order.TotalAmount.Currency(),
		string(order.Status),
		order.CreatedAt,
		order.UpdatedAt,
	)
	
	if err != nil {
		return err
	}

	// Delete existing order items
	_, err = tx.Exec("DELETE FROM order_items WHERE order_id = $1", order.ID.String())
	if err != nil {
		return err
	}

	// Insert order items
	for _, item := range order.Items {
		itemQuery := `
			INSERT INTO order_items (order_id, product_id, quantity, price_amount, price_currency, created_at)
			VALUES ($1, $2, $3, $4, $5, $6)
		`
		
		_, err = tx.Exec(itemQuery,
			order.ID.String(),
			item.ProductID.String(),
			item.Quantity,
			item.Price.Amount(),
			item.Price.Currency(),
			order.CreatedAt,
		)
		
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// FindByID finds an order by ID
func (r *OrderRepositoryImpl) FindByID(id domain.OrderID) (*domain.Order, error) {
	// Get order
	orderQuery := `
		SELECT id, customer_id, total_amount, total_currency, status, created_at, updated_at
		FROM orders
		WHERE id = $1
	`
	
	var order domain.Order
	var totalAmount int64
	var totalCurrency string
	var status string
	
	err := r.db.QueryRow(orderQuery, id.String()).Scan(
		&order.ID,
		&order.CustomerID,
		&totalAmount,
		&totalCurrency,
		&status,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("order not found")
		}
		return nil, err
	}
	
	// Create Money value object
	total, err := domain.NewMoney(totalAmount, totalCurrency)
	if err != nil {
		return nil, err
	}
	order.TotalAmount = total
	
	// Set status
	order.Status = domain.OrderStatus(status)
	
	// Get order items
	itemsQuery := `
		SELECT product_id, quantity, price_amount, price_currency
		FROM order_items
		WHERE order_id = $1
		ORDER BY created_at
	`
	
	rows, err := r.db.Query(itemsQuery, id.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var items []domain.OrderItem
	for rows.Next() {
		var item domain.OrderItem
		var priceAmount int64
		var priceCurrency string
		
		err := rows.Scan(
			&item.ProductID,
			&item.Quantity,
			&priceAmount,
			&priceCurrency,
		)
		if err != nil {
			return nil, err
		}
		
		// Create Money value object
		price, err := domain.NewMoney(priceAmount, priceCurrency)
		if err != nil {
			return nil, err
		}
		item.Price = price
		
		items = append(items, item)
	}
	
	order.Items = items
	
	return &order, nil
}

// FindAll finds all orders
func (r *OrderRepositoryImpl) FindAll() ([]domain.Order, error) {
	query := `
		SELECT id, customer_id, total_amount, total_currency, status, created_at, updated_at
		FROM orders
		ORDER BY created_at DESC
	`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		var totalAmount int64
		var totalCurrency string
		var status string
		
		err := rows.Scan(
			&order.ID,
			&order.CustomerID,
			&totalAmount,
			&totalCurrency,
			&status,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		
		// Create Money value object
		total, err := domain.NewMoney(totalAmount, totalCurrency)
		if err != nil {
			return nil, err
		}
		order.TotalAmount = total
		
		// Set status
		order.Status = domain.OrderStatus(status)
		
		// Get order items
		items, err := r.getOrderItems(order.ID)
		if err != nil {
			return nil, err
		}
		order.Items = items
		
		orders = append(orders, order)
	}
	
	return orders, nil
}

// FindByCustomerID finds orders by customer ID
func (r *OrderRepositoryImpl) FindByCustomerID(customerID domain.CustomerID) ([]domain.Order, error) {
	query := `
		SELECT id, customer_id, total_amount, total_currency, status, created_at, updated_at
		FROM orders
		WHERE customer_id = $1
		ORDER BY created_at DESC
	`
	
	rows, err := r.db.Query(query, customerID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		var totalAmount int64
		var totalCurrency string
		var status string
		
		err := rows.Scan(
			&order.ID,
			&order.CustomerID,
			&totalAmount,
			&totalCurrency,
			&status,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		
		// Create Money value object
		total, err := domain.NewMoney(totalAmount, totalCurrency)
		if err != nil {
			return nil, err
		}
		order.TotalAmount = total
		
		// Set status
		order.Status = domain.OrderStatus(status)
		
		// Get order items
		items, err := r.getOrderItems(order.ID)
		if err != nil {
			return nil, err
		}
		order.Items = items
		
		orders = append(orders, order)
	}
	
	return orders, nil
}

// Delete deletes an order
func (r *OrderRepositoryImpl) Delete(id domain.OrderID) error {
	query := `DELETE FROM orders WHERE id = $1`
	_, err := r.db.Exec(query, id.String())
	return err
}

// getOrderItems gets order items for an order
func (r *OrderRepositoryImpl) getOrderItems(orderID domain.OrderID) ([]domain.OrderItem, error) {
	query := `
		SELECT product_id, quantity, price_amount, price_currency
		FROM order_items
		WHERE order_id = $1
		ORDER BY created_at
	`
	
	rows, err := r.db.Query(query, orderID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var items []domain.OrderItem
	for rows.Next() {
		var item domain.OrderItem
		var priceAmount int64
		var priceCurrency string
		
		err := rows.Scan(
			&item.ProductID,
			&item.Quantity,
			&priceAmount,
			&priceCurrency,
		)
		if err != nil {
			return nil, err
		}
		
		// Create Money value object
		price, err := domain.NewMoney(priceAmount, priceCurrency)
		if err != nil {
			return nil, err
		}
		item.Price = price
		
		items = append(items, item)
	}
	
	return items, nil
}
