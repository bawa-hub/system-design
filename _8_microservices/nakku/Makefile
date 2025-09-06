# Nakku Quick Commerce - Makefile

.PHONY: help build run test clean docker-build docker-up docker-down migrate

# Default target
help: ## Show this help message
	@echo "Nakku Quick Commerce - Available Commands:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Build commands
build: ## Build all services
	@echo "Building all services..."
	@go build -o bin/api-gateway ./api-gateway
	@go build -o bin/user-service ./services/user
	@go build -o bin/product-service ./services/product
	@go build -o bin/inventory-service ./services/inventory
	@go build -o bin/order-service ./services/order
	@go build -o bin/cart-service ./services/cart
	@go build -o bin/payment-service ./services/payment
	@go build -o bin/delivery-service ./services/delivery
	@go build -o bin/notification-service ./services/notification
	@go build -o bin/location-service ./services/location
	@go build -o bin/analytics-service ./services/analytics

build-gateway: ## Build API Gateway
	@go build -o bin/api-gateway ./api-gateway

build-user: ## Build User Service
	@go build -o bin/user-service ./services/user

build-product: ## Build Product Service
	@go build -o bin/product-service ./services/product

# Run commands
run-all: ## Run all services
	@echo "Starting all services..."
	@make run-gateway &
	@make run-user &
	@make run-product &
	@make run-inventory &
	@make run-order &
	@make run-cart &
	@make run-payment &
	@make run-delivery &
	@make run-notification &
	@make run-location &
	@make run-analytics &

run-gateway: ## Run API Gateway
	@echo "Starting API Gateway..."
	@go run ./api-gateway

run-user: ## Run User Service
	@echo "Starting User Service..."
	@go run ./services/user

run-product: ## Run Product Service
	@echo "Starting Product Service..."
	@go run ./services/product

run-inventory: ## Run Inventory Service
	@echo "Starting Inventory Service..."
	@go run ./services/inventory

run-order: ## Run Order Service
	@echo "Starting Order Service..."
	@go run ./services/order

run-cart: ## Run Cart Service
	@echo "Starting Cart Service..."
	@go run ./services/cart

run-payment: ## Run Payment Service
	@echo "Starting Payment Service..."
	@go run ./services/payment

run-delivery: ## Run Delivery Service
	@echo "Starting Delivery Service..."
	@go run ./services/delivery

run-notification: ## Run Notification Service
	@echo "Starting Notification Service..."
	@go run ./services/notification

run-location: ## Run Location Service
	@echo "Starting Location Service..."
	@go run ./services/location

run-analytics: ## Run Analytics Service
	@echo "Starting Analytics Service..."
	@go run ./services/analytics

# Test commands
test: ## Run all tests
	@echo "Running all tests..."
	@go test ./...

test-user: ## Run User Service tests
	@go test ./services/user/...

test-product: ## Run Product Service tests
	@go test ./services/product/...

test-gateway: ## Run API Gateway tests
	@go test ./api-gateway/...

# Database commands
migrate-all: ## Run all database migrations
	@echo "Running all migrations..."
	@make migrate-user
	@make migrate-product
	@make migrate-inventory
	@make migrate-order
	@make migrate-cart
	@make migrate-payment
	@make migrate-delivery
	@make migrate-notification
	@make migrate-location
	@make migrate-analytics

migrate-user: ## Run User Service migrations
	@go run ./services/user/cmd/migrate

migrate-product: ## Run Product Service migrations
	@go run ./services/product/cmd/migrate

migrate-inventory: ## Run Inventory Service migrations
	@go run ./services/inventory/cmd/migrate

migrate-order: ## Run Order Service migrations
	@go run ./services/order/cmd/migrate

migrate-cart: ## Run Cart Service migrations
	@go run ./services/cart/cmd/migrate

migrate-payment: ## Run Payment Service migrations
	@go run ./services/payment/cmd/migrate

migrate-delivery: ## Run Delivery Service migrations
	@go run ./services/delivery/cmd/migrate

migrate-notification: ## Run Notification Service migrations
	@go run ./services/notification/cmd/migrate

migrate-location: ## Run Location Service migrations
	@go run ./services/location/cmd/migrate

migrate-analytics: ## Run Analytics Service migrations
	@go run ./services/analytics/cmd/migrate

# Docker commands
docker-build: ## Build all Docker images
	@echo "Building Docker images..."
	@docker-compose build

docker-up: ## Start all services with Docker Compose
	@echo "Starting services with Docker Compose..."
	@docker-compose up -d

docker-down: ## Stop all Docker Compose services
	@echo "Stopping Docker Compose services..."
	@docker-compose down

docker-logs: ## Show Docker Compose logs
	@docker-compose logs -f

# Infrastructure commands
infra-up: ## Start infrastructure services (MySQL, Kafka, Redis)
	@echo "Starting infrastructure services..."
	@docker-compose up -d mysql kafka zookeeper redis

infra-down: ## Stop infrastructure services
	@echo "Stopping infrastructure services..."
	@docker-compose stop mysql kafka zookeeper redis

# Development commands
dev-setup: ## Setup development environment
	@echo "Setting up development environment..."
	@make infra-up
	@sleep 10
	@make migrate-all
	@echo "Development environment ready!"

dev-clean: ## Clean development environment
	@echo "Cleaning development environment..."
	@make docker-down
	@docker system prune -f

# Utility commands
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@go clean

fmt: ## Format Go code
	@echo "Formatting Go code..."
	@go fmt ./...

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

# Health check
health: ## Check health of all services
	@echo "Checking service health..."
	@curl -f http://localhost:8080/health || echo "API Gateway: DOWN"
	@curl -f http://localhost:8081/health || echo "User Service: DOWN"
	@curl -f http://localhost:8082/health || echo "Product Service: DOWN"
