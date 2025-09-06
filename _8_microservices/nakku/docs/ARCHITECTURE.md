# Nakku Quick Commerce - Architecture Documentation

## Overview

Nakku is a microservices-based backend for a quick commerce application built with Go and MySQL. The architecture follows domain-driven design principles and implements event-driven communication patterns.

## Architecture Principles

### 1. Microservices Architecture
- **Single Responsibility**: Each service handles one business domain
- **Independent Deployment**: Services can be deployed independently
- **Technology Agnostic**: Services can use different technologies if needed
- **Fault Isolation**: Failure in one service doesn't affect others

### 2. Domain-Driven Design (DDD)
- **Bounded Contexts**: Each service represents a bounded context
- **Domain Models**: Rich domain models with business logic
- **Ubiquitous Language**: Consistent terminology across services

### 3. Event-Driven Architecture
- **Asynchronous Communication**: Services communicate via events
- **Loose Coupling**: Services don't directly depend on each other
- **Event Sourcing**: State changes are captured as events
- **CQRS**: Separate read and write models where appropriate

## Service Architecture

### Core Services

#### 1. API Gateway
- **Port**: 8080
- **Responsibilities**:
  - Request routing and load balancing
  - Authentication and authorization
  - Rate limiting and throttling
  - Request/response transformation
  - Circuit breaker pattern implementation

#### 2. User Service
- **Port**: 8081
- **Database**: `nakku_user`
- **Responsibilities**:
  - User registration and authentication
  - Profile management
  - Address management
  - JWT token generation
- **Events Published**:
  - `user.registered`
  - `user.login`
  - `user.updated`
  - `user.address.added`

#### 3. Product Catalog Service
- **Port**: 8082
- **Database**: `nakku_product`
- **Responsibilities**:
  - Product information management
  - Category management
  - Product search and filtering
  - Product recommendations
- **Events Published**:
  - `product.created`
  - `product.updated`
  - `product.deleted`

#### 4. Inventory Service
- **Port**: 8083
- **Database**: `nakku_inventory`
- **Responsibilities**:
  - Real-time stock management
  - Inventory tracking
  - Low stock alerts
  - Warehouse management
- **Events Published**:
  - `inventory.updated`
  - `inventory.low_stock`
  - `inventory.out_of_stock`

#### 5. Order Service
- **Port**: 8084
- **Database**: `nakku_order`
- **Responsibilities**:
  - Order creation and management
  - Order status tracking
  - Order history
  - Order modifications
- **Events Published**:
  - `order.created`
  - `order.updated`
  - `order.cancelled`
  - `order.completed`

#### 6. Cart Service
- **Port**: 8085
- **Database**: `nakku_cart`
- **Responsibilities**:
  - Shopping cart management
  - Cart persistence
  - Cart sharing
  - Cart validation
- **Events Published**:
  - `cart.updated`
  - `cart.cleared`

#### 7. Payment Service
- **Port**: 8086
- **Database**: `nakku_payment`
- **Responsibilities**:
  - Payment processing
  - Payment method management
  - Refund processing
  - Payment status tracking
- **Events Published**:
  - `payment.initiated`
  - `payment.completed`
  - `payment.failed`
  - `payment.refunded`

#### 8. Delivery Service
- **Port**: 8087
- **Database**: `nakku_delivery`
- **Responsibilities**:
  - Delivery scheduling
  - Delivery partner management
  - Route optimization
  - Real-time tracking
- **Events Published**:
  - `delivery.scheduled`
  - `delivery.started`
  - `delivery.completed`
  - `delivery.failed`

#### 9. Notification Service
- **Port**: 8088
- **Database**: `nakku_notification`
- **Responsibilities**:
  - Push notifications
  - SMS notifications
  - Email notifications
  - Notification preferences
- **Events Consumed**:
  - All events from other services
- **Events Published**:
  - `notification.sent`

#### 10. Location Service
- **Port**: 8089
- **Database**: `nakku_location`
- **Responsibilities**:
  - Geocoding and reverse geocoding
  - Distance calculations
  - Delivery zone validation
  - Address validation
- **Events Published**:
  - `location.validated`

#### 11. Analytics Service
- **Port**: 8090
- **Database**: `nakku_analytics`
- **Responsibilities**:
  - Business analytics
  - Reporting
  - Metrics collection
  - Dashboard data
- **Events Consumed**:
  - All events from other services

## Data Flow

### Order Placement Flow
1. **User adds items to cart** → Cart Service
2. **User places order** → Order Service
3. **Order Service publishes** `order.created` event
4. **Inventory Service** reserves items
5. **Payment Service** processes payment
6. **Delivery Service** schedules delivery
7. **Notification Service** sends confirmations

### Event Flow
```
User Action → Service → Event → Kafka → Other Services
```

## Technology Stack

### Backend
- **Language**: Go 1.21+
- **Framework**: Gin (HTTP framework)
- **ORM**: GORM
- **Database**: MySQL 8.0
- **Message Queue**: Apache Kafka
- **Cache**: Redis
- **Authentication**: JWT

### Infrastructure
- **Containerization**: Docker & Docker Compose
- **API Gateway**: Custom Go implementation
- **Monitoring**: Prometheus + Grafana (planned)
- **Logging**: Structured logging with JSON

### Development Tools
- **Build Tool**: Make
- **Testing**: Go testing framework
- **Linting**: golangci-lint
- **Code Formatting**: gofmt

## Database Design

### Database per Service
Each microservice has its own database:
- `nakku_user` - User and authentication data
- `nakku_product` - Product catalog data
- `nakku_inventory` - Inventory and stock data
- `nakku_order` - Order and transaction data
- `nakku_cart` - Shopping cart data
- `nakku_payment` - Payment and transaction data
- `nakku_delivery` - Delivery and logistics data
- `nakku_notification` - Notification preferences and history
- `nakku_location` - Location and address data
- `nakku_analytics` - Analytics and reporting data

### Common Patterns
- **Base Model**: All entities extend BaseModel with ID, timestamps, and soft delete
- **Audit Trail**: Created/Updated timestamps on all entities
- **Soft Delete**: DeletedAt field for soft deletion
- **Indexes**: Proper indexing for performance

## Security

### Authentication & Authorization
- **JWT Tokens**: Stateless authentication
- **Role-Based Access**: User roles (user, admin, delivery_partner)
- **API Gateway**: Centralized authentication
- **Password Hashing**: bcrypt for password security

### Data Protection
- **Input Validation**: Request validation using Gin binding
- **SQL Injection Prevention**: GORM ORM with parameterized queries
- **CORS**: Cross-origin resource sharing configuration
- **Rate Limiting**: Request rate limiting (planned)

## Monitoring & Observability

### Logging
- **Structured Logging**: JSON format for easy parsing
- **Log Levels**: Debug, Info, Warn, Error
- **Request Tracing**: Request ID for tracing across services
- **Context Logging**: User ID, service name in logs

### Health Checks
- **Health Endpoints**: `/health` endpoint for each service
- **Dependency Checks**: Database and Kafka connectivity
- **Graceful Shutdown**: Proper shutdown handling

### Metrics (Planned)
- **Prometheus**: Metrics collection
- **Grafana**: Metrics visualization
- **Custom Metrics**: Business metrics and KPIs

## Deployment

### Development
```bash
# Start infrastructure
make infra-up

# Run services locally
make run-all

# Run with Docker
make docker-up
```

### Production (Planned)
- **Kubernetes**: Container orchestration
- **Helm Charts**: Kubernetes deployment
- **CI/CD**: GitHub Actions or GitLab CI
- **Load Balancing**: NGINX or cloud load balancer

## Scalability Considerations

### Horizontal Scaling
- **Stateless Services**: All services are stateless
- **Database Sharding**: Can be implemented per service
- **Load Balancing**: Multiple instances of each service
- **Caching**: Redis for frequently accessed data

### Performance Optimization
- **Connection Pooling**: Database connection pooling
- **Caching Strategy**: Redis caching for hot data
- **Async Processing**: Event-driven architecture
- **Database Indexing**: Proper indexing strategy

## Error Handling

### Error Patterns
- **Graceful Degradation**: Services continue working when dependencies fail
- **Circuit Breaker**: Prevent cascade failures
- **Retry Logic**: Automatic retry for transient failures
- **Dead Letter Queue**: Handle failed messages

### Error Response Format
```json
{
  "success": false,
  "error": "Error message",
  "code": 400
}
```

## Future Enhancements

### Planned Features
1. **Service Mesh**: Istio for service-to-service communication
2. **Distributed Tracing**: Jaeger for request tracing
3. **API Versioning**: Proper API versioning strategy
4. **GraphQL**: GraphQL API layer
5. **Real-time Updates**: WebSocket support
6. **Machine Learning**: Recommendation engine
7. **Multi-tenancy**: Support for multiple tenants

### Performance Improvements
1. **Database Optimization**: Query optimization and indexing
2. **Caching Strategy**: Multi-level caching
3. **CDN Integration**: Static content delivery
4. **Message Batching**: Batch processing for events
5. **Connection Pooling**: Optimize database connections
