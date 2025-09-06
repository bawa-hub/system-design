# Nakku - Quick Commerce Microservices Backend

A microservices-based backend for a quick commerce application built with Go and MySQL.

## Architecture Overview

This project implements a microservices architecture for a quick commerce platform with the following services:

### Core Services
- **API Gateway** - Entry point for all client requests
- **User Service** - User management and authentication
- **Product Catalog Service** - Product information and categories
- **Inventory Service** - Real-time stock management
- **Order Service** - Order processing and management
- **Cart Service** - Shopping cart functionality
- **Payment Service** - Payment processing
- **Delivery Service** - Delivery scheduling and tracking
- **Notification Service** - Push notifications and alerts
- **Location Service** - Geocoding and delivery zones
- **Analytics Service** - Business intelligence and reporting

### Technology Stack
- **Language**: Go 1.21+
- **Database**: MySQL 8.0
- **Message Queue**: Apache Kafka
- **Containerization**: Docker & Docker Compose
- **API Gateway**: Custom Go implementation
- **Monitoring**: Prometheus + Grafana
- **Logging**: Structured logging with JSON

## Project Structure

```
nakku/
├── api-gateway/          # API Gateway service
├── services/             # Individual microservices
│   ├── user/            # User service
│   ├── product/         # Product catalog service
│   ├── inventory/       # Inventory service
│   ├── order/           # Order service
│   ├── cart/            # Cart service
│   ├── payment/         # Payment service
│   ├── delivery/        # Delivery service
│   ├── notification/    # Notification service
│   ├── location/        # Location service
│   └── analytics/       # Analytics service
├── shared/              # Shared libraries and utilities
│   ├── config/          # Configuration management
│   ├── database/        # Database utilities
│   ├── middleware/      # Common middleware
│   ├── models/          # Shared data models
│   └── utils/           # Utility functions
├── docker/              # Docker configurations
├── scripts/             # Deployment and utility scripts
├── docs/                # Documentation
└── docker-compose.yml   # Docker Compose configuration
```

## Quick Start

1. **Prerequisites**
   - Go 1.21+
   - Docker & Docker Compose
   - Make (optional, for using Makefile commands)

2. **Clone and Setup**
   ```bash
   git clone <repository-url>
   cd nakku
   ```

3. **Start Infrastructure**
   ```bash
   docker-compose up -d mysql kafka zookeeper redis
   ```

4. **Run Services**
   ```bash
   # Start all services
   make run-all
   
   # Or start individual services
   make run-gateway
   make run-user
   ```

5. **Access Services**
   - API Gateway: http://localhost:8080
   - MySQL: localhost:3306
   - Kafka: localhost:9092
   - Kafka UI: http://localhost:8080/ui
   - Redis: localhost:6379

## Development

### Adding a New Service

1. Create service directory in `services/`
2. Implement service following the standard structure
3. Add to `docker-compose.yml`
4. Update API Gateway routes
5. Add service to Makefile

### Database Migrations

```bash
# Run migrations for all services
make migrate-all

# Run migrations for specific service
make migrate-user
```

### Testing

```bash
# Run all tests
make test

# Run tests for specific service
make test-user
```

## Environment Variables

Each service uses environment variables for configuration. See individual service README files for specific variables.

Common variables:
- `DB_HOST` - Database host
- `DB_PORT` - Database port
- `DB_USER` - Database username
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name
- `KAFKA_BROKERS` - Kafka broker addresses
- `REDIS_URL` - Redis connection URL

## Contributing

1. Follow Go coding standards
2. Write tests for new features
3. Update documentation
4. Use conventional commits

## License

MIT License
