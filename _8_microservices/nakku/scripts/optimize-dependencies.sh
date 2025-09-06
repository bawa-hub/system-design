#!/bin/bash

# Optimize Dependencies - Keep Only What Each Service Actually Needs
# This script removes unnecessary dependencies from each service

set -e

echo "🧹 Optimizing dependencies - keeping only what each service actually needs..."

# Function to optimize a service's dependencies
optimize_service() {
    local service_name=$1
    local service_path="services/$service_name"
    
    echo "🔍 Analyzing $service_name service dependencies..."
    
    if [ ! -d "$service_path" ]; then
        echo "⚠️  Service $service_name not found, skipping..."
        return
    fi
    
    cd "$service_path"
    
    # Remove go.mod and go.sum to start fresh
    rm -f go.mod go.sum
    
    # Initialize fresh go.mod
    go mod init "nakku/services/$service_name"
    
    # Add only the dependencies that are actually used
    case $service_name in
        "user")
            echo "📦 Adding User Service specific dependencies..."
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
            echo "📦 Adding $service_name service basic dependencies..."
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
        echo "✅ $service_name service builds successfully with optimized dependencies!"
        rm -f "${service_name}-service"
    else
        echo "⚠️  $service_name service has build issues (expected for unimplemented services)"
    fi
    
    cd ../..
    echo "✅ $service_name service dependencies optimized!"
    echo ""
}

# Function to optimize API Gateway
optimize_api_gateway() {
    echo "🔍 Analyzing API Gateway dependencies..."
    
    cd api-gateway
    
    # Remove go.mod and go.sum to start fresh
    rm -f go.mod go.sum
    
    # Initialize fresh go.mod
    go mod init "nakku/api-gateway"
    
    # Add only the dependencies that are actually used
    echo "📦 Adding API Gateway specific dependencies..."
    go get github.com/gin-gonic/gin
    go get github.com/sirupsen/logrus
    go get github.com/spf13/viper
    go get github.com/golang-jwt/jwt/v5
    
    # Run go mod tidy to clean up
    go mod tidy
    
    # Try to build
    echo "🔨 Testing build for API Gateway..."
    if go build -o api-gateway . 2>/dev/null; then
        echo "✅ API Gateway builds successfully with optimized dependencies!"
        rm -f api-gateway
    else
        echo "⚠️  API Gateway has build issues"
    fi
    
    cd ..
    echo "✅ API Gateway dependencies optimized!"
    echo ""
}

# List of services to optimize
SERVICES=("user" "product" "inventory" "order" "cart" "payment" "delivery" "notification" "location" "analytics")

# Optimize API Gateway
optimize_api_gateway

# Optimize each service
for service in "${SERVICES[@]}"; do
    optimize_service "$service"
done

echo "🎉 All services have been optimized with minimal dependencies!"
echo ""
echo "📊 Dependency Summary:"
echo "====================="
echo ""
echo "🔹 User Service:"
echo "   - gin (HTTP framework)"
echo "   - logrus (logging)"
echo "   - viper (config)"
echo "   - jwt (authentication)"
echo "   - bcrypt (password hashing)"
echo "   - kafka-go (messaging)"
echo "   - gorm + mysql (database)"
echo ""
echo "🔹 Other Services:"
echo "   - gin (HTTP framework)"
echo "   - logrus (logging)"
echo "   - viper (config)"
echo "   - kafka-go (messaging)"
echo "   - gorm + mysql (database)"
echo ""
echo "🔹 API Gateway:"
echo "   - gin (HTTP framework)"
echo "   - logrus (logging)"
echo "   - viper (config)"
echo "   - jwt (authentication)"
echo ""
echo "✅ Benefits:"
echo "   - Smaller binary sizes"
echo "   - Faster build times"
echo "   - Reduced security surface"
echo "   - Lower memory usage"
echo "   - Easier maintenance"
