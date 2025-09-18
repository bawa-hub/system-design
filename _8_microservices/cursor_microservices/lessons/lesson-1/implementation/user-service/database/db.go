package database

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
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(36) PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Connection.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	// Create indexes for better performance
	indexSQL := `
	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	CREATE INDEX IF NOT EXISTS idx_users_created_at ON users(created_at);
	`

	_, err = db.Connection.Exec(indexSQL)
	if err != nil {
		return fmt.Errorf("failed to create indexes: %w", err)
	}

	log.Println("Users table and indexes created successfully!")
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
