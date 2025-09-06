# How to Run Nakku Quick Commerce - Complete Guide

## 🎯 Overview

This guide will walk you through exactly how to run the Nakku Quick Commerce microservices on a single server. We'll start from scratch and get everything working.

## 📋 Prerequisites

### 1. **Install Docker Desktop**
```bash
# Download and install Docker Desktop from:
# https://www.docker.com/products/docker-desktop/

# Verify Docker is running
docker --version
docker-compose --version
```

### 2. **Install Go 1.21+**
```bash
# Download from: https://golang.org/dl/
# Or use package manager:

# macOS
brew install go

# Ubuntu/Debian
sudo apt install golang-go

# Verify installation
go version
```

### 3. **Install Make** (usually pre-installed)
```bash
# macOS
brew install make

# Ubuntu/Debian
sudo apt install make

# Verify
make --version
```

## 🚀 Step-by-Step Setup

### Step 1: **Start Docker Desktop**
- Open Docker Desktop application
- Wait for it to start (you'll see a green icon in the system tray)
- Verify it's running:
```bash
docker ps
```

### Step 2: **Start Infrastructure Services**
```bash
# This starts MySQL, Kafka, and Redis
make infra-up
```

**Expected Output:**
```
Starting infrastructure services...
Creating network "nakku_nakku-network" with driver "bridge"
Creating volume "nakku_mysql_data" with default driver
Creating volume "nakku_kafka_data" with default driver
Creating volume "nakku_redis_data" with default driver
Creating nakku-zookeeper ... done
Creating nakku-mysql ... done
Creating nakku-kafka ... done
Creating nakku-kafka-ui ... done
Creating nakku-redis ... done
```

### Step 3: **Wait for Services to be Ready**
```bash
# Wait 30 seconds for all services to start
sleep 30

# Check if services are running
docker ps
```

**Expected Output:**
```
CONTAINER ID   IMAGE                              COMMAND                  CREATED         STATUS         PORTS                    NAMES
abc123def456   confluentinc/cp-kafka:7.4.0        "/etc/confluent/dock…"   2 minutes ago   Up 2 minutes   0.0.0.0:9092->9092/tcp   nakku-kafka
def456ghi789   mysql:8.0                          "docker-entrypoint.s…"   2 minutes ago   Up 2 minutes   0.0.0.0:3306->3306/tcp   nakku-mysql
ghi789jkl012   redis:7-alpine                     "docker-entrypoint.s…"   2 minutes ago   Up 2 minutes   0.0.0.0:6379->6379/tcp   nakku-redis
jkl012mno345   confluentinc/cp-zookeeper:7.4.0    "/etc/confluent/dock…"   2 minutes ago   Up 2 minutes   0.0.0.0:2181->2181/tcp   nakku-zookeeper
mno345pqr678   provectuslabs/kafka-ui:latest      "/bin/sh -c 'java -j…"   2 minutes ago   Up 2 minutes   0.0.0.0:8080->8080/tcp   nakku-kafka-ui
```

### Step 4: **Setup Kafka Topics**
```bash
# Create Kafka topics for all services
./scripts/setup-kafka-topics.sh
```

**Expected Output:**
```
Setting up Kafka topics for Nakku Quick Commerce...
Creating topic: user-events
Creating topic: product-events
Creating topic: inventory-events
Creating topic: order-events
Creating topic: cart-events
Creating topic: payment-events
Creating topic: delivery-events
Creating topic: notification-events
Creating topic: location-events
Creating topic: analytics-events
Kafka topics setup completed!
Listing all topics:
__consumer_offsets
analytics-events
cart-events
delivery-events
inventory-events
location-events
notification-events
order-events
payment-events
product-events
user-events
```

### Step 5: **Run Database Migrations**
```bash
# Create database tables for all services
make migrate-all
```

**Note:** This will run migrations for all services. Since we only have the User Service implemented, other services will show "no migrations" which is expected.

### Step 6: **Start All Microservices**
```bash
# Start all services
make run-all
```

**What happens:**
- API Gateway starts on port 8080
- User Service starts on port 8081
- Other services start on their ports (basic structure)
- All services connect to MySQL, Kafka, and Redis

**Expected Output:**
```
Starting all services...
Starting API Gateway...
Starting User Service...
Starting Product Service...
Starting Inventory Service...
Starting Order Service...
Starting Cart Service...
Starting Payment Service...
Starting Delivery Service...
Starting Notification Service...
Starting Location Service...
Starting Analytics Service...
```

### Step 7: **Verify Everything is Working**
```bash
# Check health of all services
make health
```

**Expected Output:**
```
Checking service health...
API Gateway: OK
User Service: OK
Product Service: OK
Inventory Service: OK
Order Service: OK
Cart Service: OK
Payment Service: OK
Delivery Service: OK
Notification Service: OK
Location Service: OK
Analytics Service: OK
```

## 🌐 Access Your Services

### **Main Entry Point (API Gateway)**
- **URL**: http://localhost:8080
- **Health Check**: http://localhost:8080/health

### **Individual Services**
- User Service: http://localhost:8081/health
- Product Service: http://localhost:8082/health
- Order Service: http://localhost:8084/health
- (and so on...)

### **Infrastructure Services**
- **Kafka UI**: http://localhost:8080/ui (Kafka management interface)
- **MySQL**: localhost:3306
- **Redis**: localhost:6379

## 🧪 Test the API

### 1. **Test Health Endpoint**
```bash
curl http://localhost:8080/health
```

**Expected Response:**
```json
{
  "status": "healthy",
  "timestamp": 1703123456,
  "service": "api-gateway"
}
```

### 2. **Register a User**
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "username": "testuser",
    "password": "password123",
    "first_name": "Test",
    "last_name": "User"
  }'
```

**Expected Response:**
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "id": 1,
    "email": "test@example.com",
    "username": "testuser",
    "first_name": "Test",
    "last_name": "User",
    "role": "user",
    "is_active": true,
    "created_at": "2023-12-21T10:30:00Z",
    "updated_at": "2023-12-21T10:30:00Z"
  }
}
```

### 3. **Login**
```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

**Expected Response:**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "user": {
      "id": 1,
      "email": "test@example.com",
      "username": "testuser",
      "first_name": "Test",
      "last_name": "User",
      "role": "user",
      "is_active": true
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### 4. **Get User Profile (with token)**
```bash
# Replace YOUR_JWT_TOKEN with the token from login response
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "message": "Profile retrieved successfully",
  "data": {
    "id": 1,
    "email": "test@example.com",
    "username": "testuser",
    "first_name": "Test",
    "last_name": "User",
    "role": "user",
    "is_active": true
  }
}
```

## 📊 Monitor Your Services

### **View Logs**
```bash
# View all service logs
make docker-logs

# View specific service logs
docker logs nakku-user-service
docker logs nakku-api-gateway
```

### **Check Service Status**
```bash
# Check which services are running
docker ps

# Check service health
curl http://localhost:8080/health
curl http://localhost:8081/health
```

### **View Kafka Topics**
```bash
# List all Kafka topics
docker exec nakku-kafka kafka-topics --list --bootstrap-server localhost:9092
```

### **Connect to Database**
```bash
# Connect to MySQL
docker exec -it nakku-mysql mysql -u nakku -p

# Password: nakkupassword

# View databases
SHOW DATABASES;

# View tables in user database
USE nakku_user;
SHOW TABLES;

# View users table
SELECT * FROM users;
```

## 🛠️ Development Commands

### **Run Individual Services**
```bash
# Run only API Gateway
make run-gateway

# Run only User Service
make run-user

# Run only Product Service
make run-product
```

### **Build Services**
```bash
# Build all services
make build

# Build specific service
make build-user
make build-gateway
```

### **Database Operations**
```bash
# Connect to MySQL
docker exec -it nakku-mysql mysql -u nakku -p

# View databases
docker exec -it nakku-mysql mysql -u nakku -p -e "SHOW DATABASES;"

# View tables in user database
docker exec -it nakku-mysql mysql -u nakku -p -e "USE nakku_user; SHOW TABLES;"
```

## 🚨 Troubleshooting

### **Docker Not Running**
```bash
# Start Docker Desktop
# Or restart Docker service
sudo systemctl restart docker  # Linux
```

### **Port Already in Use**
```bash
# Find what's using port 8080
lsof -i :8080

# Kill the process
kill -9 <PID>
```

### **Services Not Starting**
```bash
# Check Docker logs
docker logs nakku-mysql
docker logs nakku-kafka
docker logs nakku-redis

# Restart infrastructure
make infra-down
make infra-up
```

### **Database Connection Issues**
```bash
# Check MySQL status
docker exec nakku-mysql mysqladmin ping -h localhost

# Check MySQL logs
docker logs nakku-mysql
```

### **Kafka Issues**
```bash
# Check Kafka status
docker exec nakku-kafka kafka-broker-api-versions --bootstrap-server localhost:9092

# Check Kafka logs
docker logs nakku-kafka
```

### **Go Module Issues**
```bash
# Clean and download modules
go clean -modcache
go mod download
go mod tidy
```

## 🛑 Stop Everything

```bash
# Stop all services
make docker-down

# Or stop infrastructure only
make infra-down

# Or stop individual services
docker stop nakku-mysql nakku-kafka nakku-redis
```

## 📁 What's Working Right Now

✅ **Infrastructure**: MySQL, Kafka, Redis  
✅ **API Gateway**: Request routing and authentication  
✅ **User Service**: Complete with registration, login, profile, addresses  
✅ **Shared Libraries**: Database, Kafka, middleware, models  
✅ **Docker Setup**: All services containerized  
✅ **Health Checks**: All services have health endpoints  

## 🔄 What Happens When You Run the Project

1. **Infrastructure starts** (MySQL, Kafka, Redis)
2. **API Gateway** starts on port 8080
3. **User Service** starts on port 8081 (fully functional)
4. **Other services** start on their ports (basic structure)
5. **All services** connect to MySQL, Kafka, and Redis
6. **Health endpoints** become available
7. **API Gateway** routes requests to appropriate services

## 🚀 Next Steps

1. **Test the User Service** - Register, login, manage profile
2. **Implement Product Service** - Add product catalog functionality
3. **Implement Order Service** - Add order processing
4. **Add more services** - Cart, Payment, Delivery, etc.
5. **Deploy to production** - Use the production deployment guide

## 📚 More Documentation

- **Quick Start**: `QUICK_START.md`
- **Architecture**: `docs/ARCHITECTURE.md`
- **Getting Started**: `docs/GETTING_STARTED.md`
- **Production Deployment**: `docs/PRODUCTION_DEPLOYMENT.md`
- **Service Communication**: `docs/SERVICE_COMMUNICATION.md`

---

**Ready to start? Follow the steps above and you'll have a working microservices architecture in minutes!** 🚀
