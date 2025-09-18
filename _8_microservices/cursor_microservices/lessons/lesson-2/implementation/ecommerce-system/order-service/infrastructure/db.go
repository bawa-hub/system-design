package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// DB holds the database connection and configuration
type DB struct {
	Connection *sql.DB
	Host       string
	Port       string
	User       string
	Password   string
	DBName     string
}

// NewDB creates a new database connection
func NewDB(host, port, user, password, dbname string) (*DB, error) {
	// Construct the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to database!")

	return &DB{
		Connection: db,
		Host:       host,
		Port:       port,
		User:       user,
		Password:   password,
		DBName:     dbname,
	}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.Connection.Close()
}

// InitSchema creates the necessary tables
func (db *DB) InitSchema() error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS orders (
		id VARCHAR(36) PRIMARY KEY,
		customer_id VARCHAR(36) NOT NULL,
		total_amount BIGINT NOT NULL,
		total_currency VARCHAR(3) NOT NULL,
		status VARCHAR(20) NOT NULL DEFAULT 'draft',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE TABLE IF NOT EXISTS order_items (
		id SERIAL PRIMARY KEY,
		order_id VARCHAR(36) NOT NULL,
		product_id VARCHAR(36) NOT NULL,
		quantity INTEGER NOT NULL,
		price_amount BIGINT NOT NULL,
		price_currency VARCHAR(3) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
	);
	`

	_, err := db.Connection.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create orders table: %w", err)
	}

	// Create indexes for better performance
	indexSQL := `
	CREATE INDEX IF NOT EXISTS idx_orders_customer_id ON orders(customer_id);
	CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);
	CREATE INDEX IF NOT EXISTS idx_orders_created_at ON orders(created_at);
	CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON order_items(order_id);
	CREATE INDEX IF NOT EXISTS idx_order_items_product_id ON order_items(product_id);
	`

	_, err = db.Connection.Exec(indexSQL)
	if err != nil {
		return fmt.Errorf("failed to create indexes: %w", err)
	}

	log.Println("Orders table and indexes created successfully!")
	return nil
}

// HealthCheck checks if the database is healthy
func (db *DB) HealthCheck() error {
	return db.Connection.Ping()
}

// GetStats returns database statistics
func (db *DB) GetStats() map[string]interface{} {
	stats := db.Connection.Stats()
	return map[string]interface{}{
		"max_open_connections":     stats.MaxOpenConnections,
		"open_connections":         stats.OpenConnections,
		"in_use":                  stats.InUse,
		"idle":                    stats.Idle,
		"wait_count":              stats.WaitCount,
		"wait_duration":           stats.WaitDuration,
		"max_idle_closed":         stats.MaxIdleClosed,
		"max_idle_time_closed":    stats.MaxIdleTimeClosed,
		"max_lifetime_closed":     stats.MaxLifetimeClosed,
	}
}
