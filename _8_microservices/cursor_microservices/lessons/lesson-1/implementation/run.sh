#!/bin/bash

# Lesson 1: Microservices Implementation Script
# This script helps you run the microservices

echo "🚀 Starting Microservices Learning Journey - Lesson 1"
echo "=================================================="

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
ports=(5432 8080 8081 6379)
for port in "${ports[@]}"; do
    if lsof -Pi :$port -sTCP:LISTEN -t >/dev/null 2>&1; then
        print_warning "Port $port is already in use. This might cause issues."
    else
        print_success "Port $port is available"
    fi
done

# Build and start services
print_status "Building and starting services..."
docker-compose up --build -d

# Wait for services to be ready
print_status "Waiting for services to be ready..."
sleep 15

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

# Check Redis
print_status "Checking Redis..."
if docker-compose exec redis redis-cli ping > /dev/null 2>&1; then
    print_success "Redis is healthy"
else
    print_error "Redis is not responding"
fi

echo ""
echo "🎉 Services are running!"
echo ""
echo "📋 Available Services:"
echo "  - User Service: http://localhost:8080"
echo "  - Product Service: http://localhost:8081"
echo "  - PostgreSQL: localhost:5432"
echo "  - Redis: localhost:6379"
echo ""
echo "📚 API Endpoints:"
echo "  User Service:"
echo "    GET    http://localhost:8080/api/v1/users"
echo "    POST   http://localhost:8080/api/v1/users"
echo "    GET    http://localhost:8080/api/v1/users/{id}"
echo "    PUT    http://localhost:8080/api/v1/users/{id}"
echo "    DELETE http://localhost:8080/api/v1/users/{id}"
echo ""
echo "  Product Service:"
echo "    GET    http://localhost:8081/api/v1/products"
echo "    POST   http://localhost:8081/api/v1/products"
echo "    GET    http://localhost:8081/api/v1/products/{id}"
echo "    PUT    http://localhost:8081/api/v1/products/{id}"
echo "    DELETE http://localhost:8081/api/v1/products/{id}"
echo ""
echo "🧪 Test the APIs:"
echo "  # Create a user"
echo "  curl -X POST http://localhost:8080/api/v1/users \\"
echo "    -H 'Content-Type: application/json' \\"
echo "    -d '{\"email\":\"john@example.com\",\"first_name\":\"John\",\"last_name\":\"Doe\"}'"
echo ""
echo "  # Get all users"
echo "  curl http://localhost:8080/api/v1/users"
echo ""
echo "  # Create a product"
echo "  curl -X POST http://localhost:8081/api/v1/products \\"
echo "    -H 'Content-Type: application/json' \\"
echo "    -d '{\"name\":\"Laptop\",\"description\":\"Gaming Laptop\",\"price\":999.99,\"stock\":10}'"
echo ""
echo "  # Get all products"
echo "  curl http://localhost:8081/api/v1/products"
echo ""
echo "🔧 Management Commands:"
echo "  🛑 To stop services: docker-compose down"
echo "  📊 To view logs: docker-compose logs -f"
echo "  🔄 To restart: docker-compose restart"
echo "  🧹 To clean up: docker-compose down -v"
echo ""
echo "📖 Next Steps:"
echo "  1. Test the APIs using the curl commands above"
echo "  2. Check the health endpoints: http://localhost:8080/health"
echo "  3. View the metrics: http://localhost:8080/metrics"
echo "  4. Read the lesson theory in the theory/ folder"
echo "  5. Complete the practice exercises in the practice/ folder"
echo ""
echo "🎯 Ready to learn microservices! Good luck! 🚀"
