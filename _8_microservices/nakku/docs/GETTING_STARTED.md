# Getting Started with Nakku Quick Commerce

This guide will help you set up and run the Nakku Quick Commerce microservices backend locally.

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.21+**: [Download Go](https://golang.org/dl/)
- **Docker & Docker Compose**: [Download Docker](https://www.docker.com/get-started)
- **Make**: Usually pre-installed on macOS/Linux
- **Git**: [Download Git](https://git-scm.com/downloads)

## Quick Start

### 1. Clone the Repository
```bash
git clone <repository-url>
cd nakku
```

### 2. Start Infrastructure Services
```bash
# Start MySQL, Kafka, and Redis
make infra-up

# Wait for services to be ready (about 30 seconds)
```

### 3. Setup Kafka Topics
```bash
# Create Kafka topics for all services
./scripts/setup-kafka-topics.sh
```

### 4. Run Database Migrations
```bash
# Run migrations for all services
make migrate-all
```

### 5. Start All Services
```bash
# Start all microservices
make run-all
```

### 6. Verify Setup
```bash
# Check health of all services
make health
```

## Detailed Setup

### Environment Configuration

1. **Copy environment file**:
   ```bash
   cp env.example .env
   ```

2. **Edit configuration** (optional):
   ```bash
   # Edit .env file with your preferred settings
   nano .env
   ```

### Infrastructure Services

#### MySQL Database
- **Host**: localhost:3306
- **Username**: nakku
- **Password**: nakkupassword
- **Databases**: Separate database for each service

#### Apache Kafka
- **Broker**: localhost:9092
- **UI**: http://localhost:8080/ui
- **Topics**: Auto-created by setup script

#### Redis Cache
- **Host**: localhost:6379
- **Database**: 0

### Service Endpoints

| Service | Port | Health Check | Description |
|---------|------|--------------|-------------|
| API Gateway | 8080 | http://localhost:8080/health | Entry point for all requests |
| User Service | 8081 | http://localhost:8081/health | User management and auth |
| Product Service | 8082 | http://localhost:8082/health | Product catalog |
| Inventory Service | 8083 | http://localhost:8083/health | Stock management |
| Order Service | 8084 | http://localhost:8084/health | Order processing |
| Cart Service | 8085 | http://localhost:8085/health | Shopping cart |
| Payment Service | 8086 | http://localhost:8086/health | Payment processing |
| Delivery Service | 8087 | http://localhost:8087/health | Delivery management |
| Notification Service | 8088 | http://localhost:8088/health | Notifications |
| Location Service | 8089 | http://localhost:8089/health | Location services |
| Analytics Service | 8090 | http://localhost:8090/health | Analytics and reporting |

## API Usage

### Authentication

Most endpoints require authentication. First, register a user and login:

#### 1. Register a User
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "username": "testuser",
    "password": "password123",
    "first_name": "Test",
    "last_name": "User"
  }'
```

#### 2. Login
```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

#### 3. Use the Token
```bash
# Use the token from login response in Authorization header
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Example API Calls

#### Get User Profile
```bash
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

#### Add Address
```bash
curl -X POST http://localhost:8080/api/v1/users/addresses \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "home",
    "address_line1": "123 Main St",
    "city": "New York",
    "state": "NY",
    "postal_code": "10001",
    "country": "USA",
    "is_default": true
  }'
```

## Development

### Running Individual Services

```bash
# Run only the API Gateway
make run-gateway

# Run only the User Service
make run-user

# Run only the Product Service
make run-product
```

### Building Services

```bash
# Build all services
make build

# Build specific service
make build-user
make build-gateway
```

### Testing

```bash
# Run all tests
make test

# Run tests for specific service
make test-user
make test-gateway
```

### Database Operations

```bash
# Run migrations for all services
make migrate-all

# Run migrations for specific service
make migrate-user
make migrate-product
```

### Docker Operations

```bash
# Build all Docker images
make docker-build

# Start all services with Docker
make docker-up

# Stop all Docker services
make docker-down

# View logs
make docker-logs
```

## Troubleshooting

### Common Issues

#### 1. Port Already in Use
```bash
# Check what's using the port
lsof -i :8080

# Kill the process
kill -9 <PID>
```

#### 2. Database Connection Issues
```bash
# Check if MySQL is running
docker ps | grep mysql

# Check MySQL logs
docker logs nakku-mysql
```

#### 3. Kafka Connection Issues
```bash
# Check if Kafka is running
docker ps | grep kafka

# Check Kafka logs
docker logs nakku-kafka

# List Kafka topics
docker exec nakku-kafka kafka-topics --list --bootstrap-server localhost:9092
```

#### 4. Service Health Check Fails
```bash
# Check service logs
make docker-logs

# Check individual service
curl http://localhost:8081/health
```

### Logs and Debugging

#### View Service Logs
```bash
# View all logs
make docker-logs

# View specific service logs
docker logs nakku-user-service
docker logs nakku-api-gateway
```

#### Enable Debug Logging
```bash
# Set log level to debug
export LOG_LEVEL=debug

# Restart services
make docker-down
make docker-up
```

### Database Access

#### Connect to MySQL
```bash
# Connect to MySQL container
docker exec -it nakku-mysql mysql -u nakku -p

# Use specific database
USE nakku_user;
SHOW TABLES;
```

#### Reset Database
```bash
# Stop services
make docker-down

# Remove volumes
docker volume rm nakku_mysql_data

# Start services again
make docker-up
make migrate-all
```

## Production Deployment

### Environment Variables

For production deployment, set these environment variables:

```bash
# Database
DB_HOST=your-db-host
DB_PASSWORD=your-secure-password

# Kafka
KAFKA_BROKERS=your-kafka-brokers

# JWT
JWT_SECRET=your-secure-jwt-secret

# Redis
REDIS_URL=your-redis-url
```

### Docker Compose for Production

```bash
# Use production docker-compose
docker-compose -f docker-compose.prod.yml up -d
```

### Health Monitoring

```bash
# Check all services health
make health

# Monitor logs
make docker-logs
```

## Next Steps

1. **Explore the API**: Use the API endpoints to understand the functionality
2. **Read Documentation**: Check `docs/ARCHITECTURE.md` for detailed architecture
3. **Add Features**: Implement additional microservices
4. **Customize**: Modify configuration and add business logic
5. **Deploy**: Deploy to your preferred cloud platform

## Support

If you encounter any issues:

1. Check the troubleshooting section above
2. Review the logs for error messages
3. Ensure all prerequisites are installed
4. Verify that all infrastructure services are running
5. Check the GitHub issues for known problems

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

For more information, see the main README.md file.
