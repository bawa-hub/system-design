#!/bin/bash

# Lesson 2: Domain-Driven Design Implementation Script
# This script helps you run the DDD-based microservices

echo "🚀 Starting Domain-Driven Design Learning Journey - Lesson 2"
echo "=========================================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Docker is running
print_status "Checking Docker status..."
if ! docker info > /dev/null 2>&1; then
    print_error "Docker is not running. Please start Docker first."
    exit 1
fi
print_success "Docker is running"

# Check if Docker Compose is available
print_status "Checking Docker Compose..."
if ! command -v docker-compose &> /dev/null; then
    print_error "Docker Compose is not installed. Please install Docker Compose."
    exit 1
fi
print_success "Docker Compose is available"

# Check if ports are available
print_status "Checking port availability..."
ports=(5432 8080 8081 8082 6379)
for port in "${ports[@]}"; do
    if lsof -Pi :$port -sTCP:LISTEN -t >/dev/null 2>&1; then
        print_warning "Port $port is already in use. This might cause issues."
    else
        print_success "Port $port is available"
    fi
done

# Build and start services
print_status "Building and starting DDD-based microservices..."
docker-compose up --build -d

# Wait for services to be ready
print_status "Waiting for services to be ready..."
sleep 20

# Check service health
print_status "Checking service health..."

# Check PostgreSQL
print_status "Checking PostgreSQL..."
if docker-compose exec postgres pg_isready -U postgres > /dev/null 2>&1; then
    print_success "PostgreSQL is healthy"
else
    print_error "PostgreSQL is not responding"
fi

# Check User Service
print_status "Checking User Service..."
if curl -s http://localhost:8080/health > /dev/null 2>&1; then
    print_success "User Service is healthy"
else
    print_error "User Service is not responding"
fi

# Check Product Service
print_status "Checking Product Service..."
if curl -s http://localhost:8081/health > /dev/null 2>&1; then
    print_success "Product Service is healthy"
else
    print_error "Product Service is not responding"
fi

# Check Order Service
print_status "Checking Order Service..."
if curl -s http://localhost:8082/health > /dev/null 2>&1; then
    print_success "Order Service is healthy"
else
    print_error "Order Service is not responding"
fi

# Check Redis
print_status "Checking Redis..."
if docker-compose exec redis redis-cli ping > /dev/null 2>&1; then
    print_success "Redis is healthy"
else
    print_error "Redis is not responding"
fi

echo ""
echo "🎉 DDD-based microservices are running!"
echo ""
echo "📋 Available Services:"
echo "  - User Service: http://localhost:8080"
echo "  - Product Service: http://localhost:8081"
echo "  - Order Service: http://localhost:8082"
echo "  - PostgreSQL: localhost:5432"
echo "  - Redis: localhost:6379"
echo ""
echo "📚 API Endpoints:"
echo "  User Service (DDD-based):"
echo "    GET    http://localhost:8080/api/v1/users"
echo "    POST   http://localhost:8080/api/v1/users"
echo "    GET    http://localhost:8080/api/v1/users/{id}"
echo "    PUT    http://localhost:8080/api/v1/users/{id}"
echo "    DELETE http://localhost:8080/api/v1/users/{id}"
echo "    POST   http://localhost:8080/api/v1/users/{id}/activate"
echo "    POST   http://localhost:8080/api/v1/users/{id}/deactivate"
echo ""
echo "  Product Service (DDD-based):"
echo "    GET    http://localhost:8081/api/v1/products"
echo "    POST   http://localhost:8081/api/v1/products"
echo "    GET    http://localhost:8081/api/v1/products/{id}"
echo "    PUT    http://localhost:8081/api/v1/products/{id}"
echo "    DELETE http://localhost:8081/api/v1/products/{id}"
echo ""
echo "  Order Service (DDD-based):"
echo "    GET    http://localhost:8082/api/v1/orders"
echo "    POST   http://localhost:8082/api/v1/orders"
echo "    GET    http://localhost:8082/api/v1/orders/{id}"
echo "    PUT    http://localhost:8082/api/v1/orders/{id}"
echo "    POST   http://localhost:8082/api/v1/orders/{id}/confirm"
echo "    POST   http://localhost:8082/api/v1/orders/{id}/ship"
echo ""
echo "🧪 Test the DDD APIs:"
echo "  # Create a user (with domain validation)"
echo "  curl -X POST http://localhost:8080/api/v1/users \\"
echo "    -H 'Content-Type: application/json' \\"
echo "    -d '{\"email\":\"john@example.com\",\"first_name\":\"John\",\"last_name\":\"Doe\"}'"
echo ""
echo "  # Activate user (domain behavior)"
echo "  curl -X POST http://localhost:8080/api/v1/users/{user_id}/activate"
echo ""
echo "  # Create a product (with domain validation)"
echo "  curl -X POST http://localhost:8081/api/v1/products \\"
echo "    -H 'Content-Type: application/json' \\"
echo "    -d '{\"name\":\"Laptop\",\"description\":\"Gaming Laptop\",\"price\":99999,\"currency\":\"USD\",\"category\":\"Electronics\"}'"
echo ""
echo "  # Create an order (with domain events)"
echo "  curl -X POST http://localhost:8082/api/v1/orders \\"
echo "    -H 'Content-Type: application/json' \\"
echo "    -d '{\"customer_id\":\"{user_id}\",\"items\":[{\"product_id\":\"{product_id}\",\"quantity\":1,\"price\":99999,\"currency\":\"USD\"}]}'"
echo ""
echo "🔧 Management Commands:"
echo "  🛑 To stop services: docker-compose down"
echo "  📊 To view logs: docker-compose logs -f"
echo "  🔄 To restart: docker-compose restart"
echo "  🧹 To clean up: docker-compose down -v"
echo ""
echo "📖 DDD Concepts Demonstrated:"
echo "  ✅ Bounded Contexts (User, Product, Order services)"
echo "  ✅ Aggregates (User, Product, Order aggregates)"
echo "  ✅ Entities (User, Product, Order entities)"
echo "  ✅ Value Objects (Email, Name, Money, IDs)"
echo "  ✅ Domain Events (UserCreated, ProductCreated, OrderCreated)"
echo "  ✅ Repositories (Data access abstraction)"
echo "  ✅ Services (Business logic encapsulation)"
echo "  ✅ Anti-Corruption Layer (External service integration)"
echo ""
echo "🎯 Next Steps:"
echo "  1. Test the APIs using the curl commands above"
echo "  2. Check the health endpoints: http://localhost:8080/health"
echo "  3. Read the lesson theory in the theory/ folder"
echo "  4. Complete the practice exercises in the practice/ folder"
echo "  5. Study the DDD patterns in the implementation"
echo ""
echo "🎯 Ready to master Domain-Driven Design! Good luck! 🚀"
