# Running Nakku Without Docker (Alternative Setup)

Since Docker is having network connectivity issues, here's how to run the project using local installations of the required services.

## Prerequisites

### 1. **Install MySQL Locally**
```bash
# macOS
brew install mysql
brew services start mysql

# Ubuntu/Debian
sudo apt update
sudo apt install mysql-server
sudo systemctl start mysql
sudo systemctl enable mysql

# Create databases
mysql -u root -p
```

### 2. **Install Redis Locally**
```bash
# macOS
brew install redis
brew services start redis

# Ubuntu/Debian
sudo apt install redis-server
sudo systemctl start redis
sudo systemctl enable redis
```

### 3. **Install Kafka Locally**
```bash
# Download Kafka
wget https://downloads.apache.org/kafka/2.13-3.6.0/kafka_2.13-3.6.0.tgz
tar -xzf kafka_2.13-3.6.0.tgz
cd kafka_2.13-3.6.0

# Start Zookeeper
bin/zookeeper-server-start.sh config/zookeeper.properties

# In another terminal, start Kafka
bin/kafka-server-start.sh config/server.properties
```

## Database Setup

### 1. **Create Databases**
```sql
-- Connect to MySQL
mysql -u root -p

-- Create databases
CREATE DATABASE nakku_user;
CREATE DATABASE nakku_product;
CREATE DATABASE nakku_inventory;
CREATE DATABASE nakku_order;
CREATE DATABASE nakku_cart;
CREATE DATABASE nakku_payment;
CREATE DATABASE nakku_delivery;
CREATE DATABASE nakku_notification;
CREATE DATABASE nakku_location;
CREATE DATABASE nakku_analytics;

-- Create user
CREATE USER 'nakku'@'localhost' IDENTIFIED BY 'nakkupassword';
GRANT ALL PRIVILEGES ON nakku_user.* TO 'nakku'@'localhost';
GRANT ALL PRIVILEGES ON nakku_product.* TO 'nakku'@'localhost';
GRANT ALL PRIVILEGES ON nakku_inventory.* TO 'nakku'@'localhost';
GRANT ALL PRIVILEGES ON nakku_order.* TO 'nakku'@'localhost';
GRANT ALL PRIVILEGES ON nakku_cart.* TO 'nakku'@'localhost';
GRANT ALL PRIVILEGES ON nakku_payment.* TO 'nakku'@'localhost';
GRANT ALL PRIVILEGES ON nakku_delivery.* TO 'nakku'@'localhost';
GRANT ALL PRIVILEGES ON nakku_notification.* TO 'nakku'@'localhost';
GRANT ALL PRIVILEGES ON nakku_location.* TO 'nakku'@'localhost';
GRANT ALL PRIVILEGES ON nakku_analytics.* TO 'nakku'@'localhost';
FLUSH PRIVILEGES;
```

## Environment Configuration

### 1. **Create .env file**
```bash
# Copy the example environment file
cp env.example .env
```

### 2. **Update .env for local setup**
```bash
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=nakku
DB_PASSWORD=nakkupassword
DB_NAME=nakku_user

# Kafka Configuration
KAFKA_BROKERS=localhost:9092
KAFKA_GROUP_ID=nakku-group

# Redis Configuration
REDIS_URL=redis://localhost:6379/0

# JWT Configuration
JWT_SECRET=nakku-secret-key-change-in-production

# Logging Configuration
LOG_LEVEL=info
LOG_FORMAT=json
```

## Running the Services

### 1. **Start Infrastructure Services**
```bash
# Start MySQL (if not already running)
brew services start mysql  # macOS
sudo systemctl start mysql  # Linux

# Start Redis (if not already running)
brew services start redis  # macOS
sudo systemctl start redis  # Linux

# Start Kafka (in separate terminals)
# Terminal 1: Zookeeper
cd kafka_2.13-3.6.0
bin/zookeeper-server-start.sh config/zookeeper.properties

# Terminal 2: Kafka
cd kafka_2.13-3.6.0
bin/kafka-server-start.sh config/server.properties
```

### 2. **Create Kafka Topics**
```bash
# Create topics
cd kafka_2.13-3.6.0
bin/kafka-topics.sh --create --topic user-events --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
bin/kafka-topics.sh --create --topic product-events --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
bin/kafka-topics.sh --create --topic order-events --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
bin/kafka-topics.sh --create --topic cart-events --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
bin/kafka-topics.sh --create --topic payment-events --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
bin/kafka-topics.sh --create --topic delivery-events --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
bin/kafka-topics.sh --create --topic notification-events --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
bin/kafka-topics.sh --create --topic location-events --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
bin/kafka-topics.sh --create --topic analytics-events --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
```

### 3. **Run Database Migrations**
```bash
# Run migrations for User Service
make migrate-user
```

### 4. **Start Services**
```bash
# Start API Gateway
make run-gateway

# In another terminal, start User Service
make run-user
```

## Testing the Setup

### 1. **Test Health Endpoints**
```bash
# Test API Gateway
curl http://localhost:8080/health

# Test User Service
curl http://localhost:8081/health
```

### 2. **Test User Registration**
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

### 3. **Test User Login**
```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

## Troubleshooting

### **MySQL Connection Issues**
```bash
# Check MySQL status
brew services list | grep mysql  # macOS
sudo systemctl status mysql      # Linux

# Test connection
mysql -u nakku -p -h localhost
```

### **Redis Connection Issues**
```bash
# Check Redis status
brew services list | grep redis  # macOS
sudo systemctl status redis      # Linux

# Test connection
redis-cli ping
```

### **Kafka Connection Issues**
```bash
# Check if Kafka is running
netstat -an | grep 9092

# List topics
cd kafka_2.13-3.6.0
bin/kafka-topics.sh --list --bootstrap-server localhost:9092
```

### **Go Module Issues**
```bash
# Clean and download modules
go clean -modcache
go mod download
go mod tidy
```

## Alternative: Use Docker When Network is Fixed

Once your Docker network connectivity is resolved, you can use the Docker setup:

```bash
# Start infrastructure with Docker
make infra-up

# Setup Kafka topics
./scripts/setup-kafka-topics.sh

# Run migrations
make migrate-all

# Start all services
make run-all
```

## Benefits of This Setup

1. **No Docker dependency** - Works even with Docker issues
2. **Direct service control** - Start/stop services individually
3. **Easier debugging** - Direct access to service logs
4. **Faster development** - No container overhead
5. **Local development** - Services run natively on your machine

This setup gives you the same functionality as the Docker setup but runs everything locally on your machine.
