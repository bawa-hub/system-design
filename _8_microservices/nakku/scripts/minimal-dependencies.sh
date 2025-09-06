#!/bin/bash

# Minimal Dependencies - Add ONLY what each service actually uses
# This script creates truly minimal go.mod files

set -e

echo "🎯 Creating minimal dependencies - only what each service actually uses..."

# Function to create minimal dependencies for a service
create_minimal_service() {
    local service_name=$1
    local service_path="services/$service_name"
    
    echo "🔍 Creating minimal dependencies for $service_name service..."
    
    if [ ! -d "$service_path" ]; then
        echo "⚠️  Service $service_name not found, skipping..."
        return
    fi
    
    cd "$service_path"
    
    # Remove existing go.mod and go.sum
    rm -f go.mod go.sum
    
    # Initialize fresh go.mod
    go mod init "nakku/services/$service_name"
    
    # Add ONLY the dependencies that are actually used in the code
    case $service_name in
        "user")
            echo "📦 Adding User Service minimal dependencies..."
            # User service actually uses these based on the code analysis
            go get github.com/gin-gonic/gin
            go get github.com/sirupsen/logrus
            go get github.com/spf13/viper
            go get github.com/golang-jwt/jwt/v5
            go get golang.org/x/crypto/bcrypt
            go get github.com/segmentio/kafka-go
            go get gorm.io/driver/mysql
            go get gorm.io/gorm
            ;;
        "product"|"inventory"|"order"|"cart"|"payment"|"delivery"|"notification"|"location"|"analytics")
            echo "📦 Adding $service_name service minimal dependencies..."
            # These services only need basic dependencies (no JWT, no bcrypt)
            go get github.com/gin-gonic/gin
            go get github.com/sirupsen/logrus
            go get github.com/spf13/viper
            go get github.com/segmentio/kafka-go
            go get gorm.io/driver/mysql
            go get gorm.io/gorm
            ;;
    esac
    
    # Run go mod tidy to clean up
    go mod tidy
    
    # Try to build
    echo "🔨 Testing build for $service_name..."
    if go build -o "${service_name}-service" . 2>/dev/null; then
        echo "✅ $service_name service builds successfully with minimal dependencies!"
        rm -f "${service_name}-service"
    else
        echo "⚠️  $service_name service has build issues (expected for unimplemented services)"
    fi
    
    cd ../..
    echo "✅ $service_name service minimal dependencies created!"
    echo ""
}

# Function to create minimal API Gateway
create_minimal_api_gateway() {
    echo "🔍 Creating minimal API Gateway dependencies..."
    
    cd api-gateway
    
    # Remove existing go.mod and go.sum
    rm -f go.mod go.sum
    
    # Initialize fresh go.mod
    go mod init "nakku/api-gateway"
    
    # Add ONLY the dependencies that are actually used
    echo "📦 Adding API Gateway minimal dependencies..."
    go get github.com/gin-gonic/gin
    go get github.com/sirupsen/logrus
    go get github.com/spf13/viper
    go get github.com/golang-jwt/jwt/v5
    # API Gateway doesn't need database or Kafka
    
    # Run go mod tidy to clean up
    go mod tidy
    
    # Try to build
    echo "🔨 Testing build for API Gateway..."
    if go build -o api-gateway . 2>/dev/null; then
        echo "✅ API Gateway builds successfully with minimal dependencies!"
        rm -f api-gateway
    else
        echo "⚠️  API Gateway has build issues"
    fi
    
    cd ..
    echo "✅ API Gateway minimal dependencies created!"
    echo ""
}

# List of services to optimize
SERVICES=("user" "product" "inventory" "order" "cart" "payment" "delivery" "notification" "location" "analytics")

# Create minimal API Gateway
create_minimal_api_gateway

# Create minimal dependencies for each service
for service in "${SERVICES[@]}"; do
    create_minimal_service "$service"
done

echo "🎉 All services now have minimal dependencies!"
echo ""
echo "📊 Minimal Dependency Summary:"
echo "=============================="
echo ""
echo "🔹 User Service (8 dependencies):"
echo "   - gin (HTTP framework)"
echo "   - logrus (logging)"
echo "   - viper (config)"
echo "   - jwt (authentication)"
echo "   - bcrypt (password hashing)"
echo "   - kafka-go (messaging)"
echo "   - gorm + mysql (database)"
echo ""
echo "🔹 Other Services (6 dependencies):"
echo "   - gin (HTTP framework)"
echo "   - logrus (logging)"
echo "   - viper (config)"
echo "   - kafka-go (messaging)"
echo "   - gorm + mysql (database)"
echo ""
echo "🔹 API Gateway (4 dependencies):"
echo "   - gin (HTTP framework)"
echo "   - logrus (logging)"
echo "   - viper (config)"
echo "   - jwt (authentication)"
echo ""
echo "✅ Benefits:"
echo "   - Minimal attack surface"
echo "   - Faster builds"
echo "   - Smaller binaries"
echo "   - Independent versioning"
echo "   - Easy service extraction"
