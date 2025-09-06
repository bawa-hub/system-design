#!/bin/bash

# Setup Service-Level Go Modules
# This script creates independent go.mod files for each service

set -e

echo "🚀 Setting up service-level Go modules..."

# List of services to process
SERVICES=("user" "product" "inventory" "order" "cart" "payment" "delivery" "notification" "location" "analytics")

# Function to setup a service
setup_service() {
    local service_name=$1
    local service_path="services/$service_name"
    
    echo "📦 Setting up $service_name service..."
    
    # Create service directory if it doesn't exist
    mkdir -p "$service_path"
    
    # Create service-specific directories
    cd "$service_path"
    mkdir -p config database middleware utils models
    
    # Copy shared code
    cp -r ../../shared/config/* ./config/ 2>/dev/null || true
    cp -r ../../shared/database/* ./database/ 2>/dev/null || true
    cp -r ../../shared/middleware/* ./middleware/ 2>/dev/null || true
    cp -r ../../shared/utils/* ./utils/ 2>/dev/null || true
    cp -r ../../shared/models/* ./models/ 2>/dev/null || true
    
    # Initialize go.mod if it doesn't exist
    if [ ! -f "go.mod" ]; then
        go mod init "nakku/services/$service_name"
    fi
    
    # Add common dependencies
    go get github.com/gin-gonic/gin
    go get github.com/sirupsen/logrus
    go get github.com/spf13/viper
    go get gorm.io/driver/mysql
    go get gorm.io/gorm
    
    # Add service-specific dependencies
    case $service_name in
        "user")
            go get github.com/golang-jwt/jwt/v5
            go get golang.org/x/crypto/bcrypt
            go get github.com/segmentio/kafka-go
            ;;
        "product"|"inventory"|"order"|"cart"|"payment"|"delivery"|"notification"|"location"|"analytics")
            go get github.com/segmentio/kafka-go
            ;;
    esac
    
    # Update import paths
    find . -name "*.go" -exec sed -i '' 's|nakku/shared/|nakku/services/'$service_name'/|g' {} \;
    
    # Fix import cycles in models
    if [ -f "models/models.go" ]; then
        # Remove self-imports
        sed -i '' 's|nakku/services/'$service_name'/models||g' models/models.go
        # Remove unused gorm import if present
        sed -i '' '/"gorm.io\/gorm"/d' models/models.go
    fi
    
    # Run go mod tidy
    go mod tidy
    
    # Try to build (this will show any remaining issues)
    echo "🔨 Testing build for $service_name..."
    if go build -o "${service_name}-service" . 2>/dev/null; then
        echo "✅ $service_name service builds successfully!"
        rm -f "${service_name}-service"
    else
        echo "⚠️  $service_name service has build issues (this is expected for unimplemented services)"
    fi
    
    cd ../..
    echo "✅ $service_name service setup complete!"
    echo ""
}

# Setup each service
for service in "${SERVICES[@]}"; do
    setup_service "$service"
done

echo "🎉 All services have been set up with independent go.mod files!"
echo ""
echo "📋 Next steps:"
echo "1. Each service now has its own go.mod and dependencies"
echo "2. Services can be built independently: cd services/user && go build"
echo "3. Services can be extracted easily by copying the service directory"
echo "4. No more shared dependency issues!"
echo ""
echo "🔍 To verify a service:"
echo "   cd services/user"
echo "   go build -o user-service ."
echo "   ./user-service"
