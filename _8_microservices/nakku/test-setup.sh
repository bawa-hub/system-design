#!/bin/bash

# Test Setup Script for Nakku Quick Commerce
# This script tests if the Go services can run without Docker

echo "🚀 Testing Nakku Quick Commerce Setup"
echo "====================================="

# Check if Go is installed
echo "1. Checking Go installation..."
if command -v go &> /dev/null; then
    echo "✅ Go is installed: $(go version)"
else
    echo "❌ Go is not installed. Please install Go 1.21+"
    exit 1
fi

# Check if Make is installed
echo "2. Checking Make installation..."
if command -v make &> /dev/null; then
    echo "✅ Make is installed: $(make --version | head -n1)"
else
    echo "❌ Make is not installed. Please install Make"
    exit 1
fi

# Check Go modules
echo "3. Checking Go modules..."
if [ -f "go.mod" ]; then
    echo "✅ go.mod found"
    go mod tidy
    echo "✅ Go modules updated"
else
    echo "❌ go.mod not found"
    exit 1
fi

# Test building API Gateway
echo "4. Testing API Gateway build..."
if go build -o /tmp/api-gateway ./api-gateway; then
    echo "✅ API Gateway builds successfully"
    rm -f /tmp/api-gateway
else
    echo "❌ API Gateway build failed"
    exit 1
fi

# Test building User Service
echo "5. Testing User Service build..."
if go build -o /tmp/user-service ./services/user; then
    echo "✅ User Service builds successfully"
    rm -f /tmp/user-service
else
    echo "❌ User Service build failed"
    exit 1
fi

# Check if configuration files exist
echo "6. Checking configuration files..."
if [ -f "config.yaml" ]; then
    echo "✅ config.yaml found"
else
    echo "❌ config.yaml not found"
fi

if [ -f "env.example" ]; then
    echo "✅ env.example found"
else
    echo "❌ env.example not found"
fi

# Check if shared libraries exist
echo "7. Checking shared libraries..."
if [ -d "shared" ]; then
    echo "✅ shared directory found"
    if [ -f "shared/config/config.go" ]; then
        echo "✅ shared/config/config.go found"
    fi
    if [ -f "shared/database/database.go" ]; then
        echo "✅ shared/database/database.go found"
    fi
    if [ -f "shared/utils/kafka.go" ]; then
        echo "✅ shared/utils/kafka.go found"
    fi
else
    echo "❌ shared directory not found"
fi

# Check if services exist
echo "8. Checking services..."
services=("api-gateway" "services/user" "services/product" "services/inventory" "services/order")
for service in "${services[@]}"; do
    if [ -d "$service" ]; then
        echo "✅ $service directory found"
    else
        echo "❌ $service directory not found"
    fi
done

# Check if Makefile exists
echo "9. Checking Makefile..."
if [ -f "Makefile" ]; then
    echo "✅ Makefile found"
    echo "Available commands:"
    make help | head -n 20
else
    echo "❌ Makefile not found"
fi

echo ""
echo "🎉 Setup Test Complete!"
echo ""
echo "Next steps:"
echo "1. If Docker is working: make infra-up"
echo "2. If Docker has issues: Follow SETUP_WITHOUT_DOCKER.md"
echo "3. Start services: make run-gateway (in one terminal) and make run-user (in another)"
echo ""
echo "For detailed instructions, see:"
echo "- HOW_TO_RUN.md (with Docker)"
echo "- SETUP_WITHOUT_DOCKER.md (without Docker)"
echo "- QUICK_START.md (quick overview)"
