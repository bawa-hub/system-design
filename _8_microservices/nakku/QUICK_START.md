# Quick Start Guide - Single Server Deployment

## 🚀 How to Run Nakku Quick Commerce on a Single Server

This guide will help you get the Nakku Quick Commerce microservices running on a single server in just a few minutes.

## Prerequisites

Make sure you have these installed:
- **Docker & Docker Compose** (for infrastructure)
- **Go 1.21+** (for running services)
- **Make** (for easy commands)

## Step-by-Step Setup

### 1. **Start Infrastructure Services**
```bash
# This starts MySQL, Kafka, and Redis in Docker containers
make infra-up
```

**What this does:**
- Starts MySQL database on port 3306
- Starts Kafka message queue on port 9092
- Starts Redis cache on port 6379
- Creates separate databases for each service

### 2. **Wait for Services to be Ready**
```bash
# Wait about 30 seconds for all services to start
sleep 30

# Check if services are running
docker ps
```

### 3. **Setup Kafka Topics**
```bash
# Create Kafka topics for all microservices
./scripts/setup-kafka-topics.sh
```

### 4. **Run Database Migrations**
```bash
# Create database tables for all services
make migrate-all
```

### 5. **Start All Microservices**
```bash
# Start all services (this will start multiple Go processes)
make run-all
```

**What this starts:**
- API Gateway (port 8080) - Entry point for all requests
- User Service (port 8081) - User management and authentication
- Product Service (port 8082) - Product catalog
- Inventory Service (port 8083) - Stock management
- Order Service (port 8084) - Order processing
- Cart Service (port 8085) - Shopping cart
- Payment Service (port 8086) - Payment processing
- Delivery Service (port 8087) - Delivery management
- Notification Service (port 8088) - Notifications
- Location Service (port 8089) - Location services
- Analytics Service (port 8090) - Analytics

### 6. **Verify Everything is Working**
```bash
# Check health of all services
make health
```

## 🌐 Access Your Services

### **API Gateway (Main Entry Point)**
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

### 1. **Register a User**
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

### 2. **Login**
```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

### 3. **Get User Profile (with token)**
```bash
# Use the token from login response
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
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

## 🛑 Stop Everything

```bash
# Stop all services
make docker-down

# Or stop infrastructure only
make infra-down
```

## 📁 Project Structure

```
nakku/
├── api-gateway/          # API Gateway service
├── services/             # Individual microservices
│   ├── user/            # User service (COMPLETE)
│   ├── product/         # Product service (TO BE IMPLEMENTED)
│   ├── inventory/       # Inventory service (TO BE IMPLEMENTED)
│   └── ...              # Other services
├── shared/              # Shared libraries
├── docker-compose.yml   # Docker configuration
├── Makefile            # Build and run commands
└── scripts/            # Utility scripts
```

## 🎯 What's Working Right Now

✅ **Infrastructure**: MySQL, Kafka, Redis  
✅ **API Gateway**: Request routing and authentication  
✅ **User Service**: Complete with registration, login, profile, addresses  
✅ **Shared Libraries**: Database, Kafka, middleware, models  
✅ **Docker Setup**: All services containerized  
✅ **Health Checks**: All services have health endpoints  

## 🔄 What Happens When You Run `make run-all`

1. **API Gateway** starts on port 8080
2. **User Service** starts on port 8081 (fully functional)
3. **Other services** start on their respective ports (basic structure)
4. **All services** connect to MySQL, Kafka, and Redis
5. **Health endpoints** become available
6. **API Gateway** routes requests to appropriate services

## 🚀 Next Steps

1. **Test the User Service** - Register, login, manage profile
2. **Implement Product Service** - Add product catalog functionality
3. **Implement Order Service** - Add order processing
4. **Add more services** - Cart, Payment, Delivery, etc.
5. **Deploy to production** - Use the production deployment guide

## 📚 More Documentation

- **Architecture**: `docs/ARCHITECTURE.md`
- **Getting Started**: `docs/GETTING_STARTED.md`
- **Production Deployment**: `docs/PRODUCTION_DEPLOYMENT.md`
- **Service Communication**: `docs/SERVICE_COMMUNICATION.md`

---

**Ready to start? Run `make infra-up` and follow the steps above!** 🚀
