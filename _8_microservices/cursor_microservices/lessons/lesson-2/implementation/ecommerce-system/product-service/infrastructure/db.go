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
	CREATE TABLE IF NOT EXISTS products (
		id VARCHAR(36) PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		description TEXT NOT NULL,
		price_amount BIGINT NOT NULL,
		price_currency VARCHAR(3) NOT NULL,
		category VARCHAR(50) NOT NULL,
		status VARCHAR(20) NOT NULL DEFAULT 'active',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Connection.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create products table: %w", err)
	}

	// Create indexes for better performance
	indexSQL := `
	CREATE INDEX IF NOT EXISTS idx_products_category ON products(category);
	CREATE INDEX IF NOT EXISTS idx_products_status ON products(status);
	CREATE INDEX IF NOT EXISTS idx_products_created_at ON products(created_at);
	`

	_, err = db.Connection.Exec(indexSQL)
	if err != nil {
		return fmt.Errorf("failed to create indexes: %w", err)
	}

	log.Println("Products table and indexes created successfully!")
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
