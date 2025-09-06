# Production Deployment Guide

## Overview

This document explains how to deploy the Nakku Quick Commerce microservices architecture in production, including multi-server deployment strategies, scaling considerations, and production best practices.

## Production Architecture

### Multi-Server Deployment Strategy

#### 1. **Infrastructure Layer** (Dedicated Servers)
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Load Balancer │    │   Load Balancer │    │   Load Balancer │
│   (NGINX/HAProxy)│    │   (NGINX/HAProxy)│    │   (NGINX/HAProxy)│
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   API Gateway   │    │   API Gateway   │    │   API Gateway   │
│   Server 1      │    │   Server 2      │    │   Server 3      │
│   Port: 8080    │    │   Port: 8080    │    │   Port: 8080    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

#### 2. **Application Services Layer** (Multiple Servers)
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   User Service  │    │ Product Service │    │Inventory Service│
│   Server 1      │    │   Server 2      │    │   Server 3      │
│   Port: 8081    │    │   Port: 8082    │    │   Port: 8083    │
└─────────────────┘    └─────────────────┘    └─────────────────┘

┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  Order Service  │    │  Cart Service   │    │Payment Service  │
│   Server 4      │    │   Server 5      │    │   Server 6      │
│   Port: 8084    │    │   Port: 8085    │    │   Port: 8086    │
└─────────────────┘    └─────────────────┘    └─────────────────┘

┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│Delivery Service │    │Notification Svc │    │ Location Service│
│   Server 7      │    │   Server 8      │    │   Server 9      │
│   Port: 8087    │    │   Port: 8088    │    │   Port: 8089    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

#### 3. **Data Layer** (Dedicated Database Servers)
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   MySQL Master  │    │   MySQL Slave   │    │   Redis Cluster │
│   Server 1      │    │   Server 2      │    │   Server 3      │
│   Port: 3306    │    │   Port: 3306    │    │   Port: 6379    │
└─────────────────┘    └─────────────────┘    └─────────────────┘

┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Kafka Cluster │    │   Kafka Cluster │    │   Kafka Cluster │
│   Server 4      │    │   Server 5      │    │   Server 6      │
│   Port: 9092    │    │   Port: 9092    │    │   Port: 9092    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## Deployment Strategies

### 1. **Single Server Deployment** (Development/Small Scale)
```yaml
# All services on one server
Server 1:
  - API Gateway
  - All Microservices
  - MySQL
  - Kafka
  - Redis
```

### 2. **Service-per-Server Deployment** (Medium Scale)
```yaml
# Each service on dedicated server
Server 1: API Gateway
Server 2: User Service + MySQL (nakku_user)
Server 3: Product Service + MySQL (nakku_product)
Server 4: Order Service + MySQL (nakku_order)
Server 5: Shared Infrastructure (Kafka, Redis)
```

### 3. **Multi-Server with Load Balancing** (Production Scale)
```yaml
# Multiple instances of each service
Load Balancer:
  - 3x API Gateway instances
  - 2x User Service instances
  - 2x Product Service instances
  - 2x Order Service instances
  - 3x Kafka brokers
  - 2x MySQL servers (Master-Slave)
  - 3x Redis nodes (Cluster)
```

### 4. **Kubernetes Deployment** (Enterprise Scale)
```yaml
# Container orchestration
Kubernetes Cluster:
  - API Gateway: 3 replicas
  - User Service: 2 replicas
  - Product Service: 2 replicas
  - Order Service: 2 replicas
  - Database: StatefulSets
  - Kafka: StatefulSets
  - Redis: StatefulSets
```

## Production Configuration

### Environment Variables for Production

```bash
# Production Environment Variables
# Database Configuration
DB_HOST=mysql-cluster.internal
DB_PORT=3306
DB_USER=nakku_prod
DB_PASSWORD=secure_production_password
DB_NAME=nakku_production

# Kafka Configuration
KAFKA_BROKERS=kafka-1.internal:9092,kafka-2.internal:9092,kafka-3.internal:9092
KAFKA_GROUP_ID=nakku-production-group

# Redis Configuration
REDIS_URL=redis://redis-cluster.internal:6379/0

# JWT Configuration
JWT_SECRET=super_secure_jwt_secret_for_production

# Service Discovery
USER_SERVICE_URL=http://user-service.internal:8081
PRODUCT_SERVICE_URL=http://product-service.internal:8082
ORDER_SERVICE_URL=http://order-service.internal:8084

# Logging
LOG_LEVEL=info
LOG_FORMAT=json
```

### Production Docker Compose

```yaml
# docker-compose.prod.yml
version: '3.8'

services:
  # Load Balancer
  nginx:
    image: nginx:alpine
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx/nginx.conf:/etc/nginx/nginx.conf
      - ./ssl:/etc/nginx/ssl
    depends_on:
      - api-gateway-1
      - api-gateway-2
      - api-gateway-3

  # API Gateway Instances
  api-gateway-1:
    build: .
    dockerfile: ./api-gateway/Dockerfile
    environment:
      - PORT=8080
      - USER_SERVICE_URL=http://user-service:8081
      - PRODUCT_SERVICE_URL=http://product-service:8082
    networks:
      - nakku-network

  api-gateway-2:
    build: .
    dockerfile: ./api-gateway/Dockerfile
    environment:
      - PORT=8080
      - USER_SERVICE_URL=http://user-service:8081
      - PRODUCT_SERVICE_URL=http://product-service:8082
    networks:
      - nakku-network

  api-gateway-3:
    build: .
    dockerfile: ./api-gateway/Dockerfile
    environment:
      - PORT=8080
      - USER_SERVICE_URL=http://user-service:8081
      - PRODUCT_SERVICE_URL=http://product-service:8082
    networks:
      - nakku-network

  # Microservices
  user-service:
    build: .
    dockerfile: ./services/user/Dockerfile
    environment:
      - PORT=8081
      - DB_HOST=mysql-master
      - KAFKA_BROKERS=kafka-1:9092,kafka-2:9092,kafka-3:9092
    networks:
      - nakku-network

  product-service:
    build: .
    dockerfile: ./services/product/Dockerfile
    environment:
      - PORT=8082
      - DB_HOST=mysql-master
      - KAFKA_BROKERS=kafka-1:9092,kafka-2:9092,kafka-3:9092
    networks:
      - nakku-network

  # Database Cluster
  mysql-master:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: root_password
      MYSQL_DATABASE: nakku_production
    volumes:
      - mysql_master_data:/var/lib/mysql
    networks:
      - nakku-network

  mysql-slave:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: root_password
      MYSQL_DATABASE: nakku_production
    volumes:
      - mysql_slave_data:/var/lib/mysql
    networks:
      - nakku-network

  # Kafka Cluster
  kafka-1:
    image: confluentinc/cp-kafka:7.4.0
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper-1:2181,zookeeper-2:2181,zookeeper-3:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka-1:9092
    networks:
      - nakku-network

  kafka-2:
    image: confluentinc/cp-kafka:7.4.0
    environment:
      KAFKA_BROKER_ID: 2
      KAFKA_ZOOKEEPER_CONNECT: zookeeper-1:2181,zookeeper-2:2181,zookeeper-3:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka-2:9092
    networks:
      - nakku-network

  kafka-3:
    image: confluentinc/cp-kafka:7.4.0
    environment:
      KAFKA_BROKER_ID: 3
      KAFKA_ZOOKEEPER_CONNECT: zookeeper-1:2181,zookeeper-2:2181,zookeeper-3:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka-3:9092
    networks:
      - nakku-network

volumes:
  mysql_master_data:
  mysql_slave_data:

networks:
  nakku-network:
    driver: bridge
```

## Service Discovery and Load Balancing

### 1. **NGINX Load Balancer Configuration**

```nginx
# nginx/nginx.conf
upstream api_gateway {
    server api-gateway-1:8080;
    server api-gateway-2:8080;
    server api-gateway-3:8080;
}

upstream user_service {
    server user-service-1:8081;
    server user-service-2:8081;
}

upstream product_service {
    server product-service-1:8082;
    server product-service-2:8082;
}

server {
    listen 80;
    server_name api.nakku.com;

    # API Gateway Load Balancing
    location / {
        proxy_pass http://api_gateway;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Health Check
    location /health {
        proxy_pass http://api_gateway/health;
    }
}
```

### 2. **Service Discovery with Consul** (Optional)

```yaml
# consul-config.yml
services:
  - name: user-service
    port: 8081
    check:
      http: http://localhost:8081/health
      interval: 10s
  - name: product-service
    port: 8082
    check:
      http: http://localhost:8082/health
      interval: 10s
```

## Database Deployment Strategies

### 1. **Database per Service on Separate Servers**

```yaml
# Each service has its own database server
User Service Server:
  - User Service Application
  - MySQL (nakku_user database)
  - Redis (user session cache)

Product Service Server:
  - Product Service Application
  - MySQL (nakku_product database)
  - Redis (product cache)

Order Service Server:
  - Order Service Application
  - MySQL (nakku_order database)
  - Redis (order cache)
```

### 2. **Shared Database Cluster**

```yaml
# Shared database infrastructure
Database Cluster:
  - MySQL Master (all databases)
  - MySQL Slave 1 (read replicas)
  - MySQL Slave 2 (read replicas)
  - Redis Cluster (shared cache)
```

### 3. **Database Sharding**

```yaml
# Sharded databases
Shard 1:
  - MySQL (nakku_user, nakku_product)
  - User Service, Product Service

Shard 2:
  - MySQL (nakku_order, nakku_cart)
  - Order Service, Cart Service

Shard 3:
  - MySQL (nakku_payment, nakku_delivery)
  - Payment Service, Delivery Service
```

## Scaling Strategies

### 1. **Horizontal Scaling**

```bash
# Scale API Gateway
docker-compose up --scale api-gateway=3

# Scale User Service
docker-compose up --scale user-service=2

# Scale Product Service
docker-compose up --scale product-service=2
```

### 2. **Auto Scaling with Kubernetes**

```yaml
# k8s-deployment.yml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: user-service
spec:
  replicas: 2
  selector:
    matchLabels:
      app: user-service
  template:
    metadata:
      labels:
        app: user-service
    spec:
      containers:
      - name: user-service
        image: nakku/user-service:latest
        ports:
        - containerPort: 8081
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
---
apiVersion: v1
kind: Service
metadata:
  name: user-service
spec:
  selector:
    app: user-service
  ports:
  - port: 8081
    targetPort: 8081
  type: ClusterIP
```

### 3. **Load Testing and Capacity Planning**

```bash
# Load testing with Apache Bench
ab -n 1000 -c 10 http://api.nakku.com/api/v1/users/profile

# Load testing with wrk
wrk -t12 -c400 -d30s http://api.nakku.com/api/v1/products
```

## Monitoring and Observability

### 1. **Application Monitoring**

```yaml
# prometheus.yml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'api-gateway'
    static_configs:
      - targets: ['api-gateway-1:8080', 'api-gateway-2:8080', 'api-gateway-3:8080']
  
  - job_name: 'user-service'
    static_configs:
      - targets: ['user-service-1:8081', 'user-service-2:8081']
  
  - job_name: 'product-service'
    static_configs:
      - targets: ['product-service-1:8082', 'product-service-2:8082']
```

### 2. **Log Aggregation**

```yaml
# ELK Stack
elasticsearch:
  image: docker.elastic.co/elasticsearch/elasticsearch:7.15.0
  environment:
    - discovery.type=single-node
  ports:
    - "9200:9200"

logstash:
  image: docker.elastic.co/logstash/logstash:7.15.0
  volumes:
    - ./logstash.conf:/usr/share/logstash/pipeline/logstash.conf

kibana:
  image: docker.elastic.co/kibana/kibana:7.15.0
  ports:
    - "5601:5601"
```

## Security Considerations

### 1. **Network Security**

```yaml
# Firewall rules
# Only allow necessary ports
Port 80/443: Load Balancer (Public)
Port 8080: API Gateway (Internal)
Port 8081-8090: Microservices (Internal)
Port 3306: MySQL (Internal)
Port 9092: Kafka (Internal)
Port 6379: Redis (Internal)
```

### 2. **SSL/TLS Configuration**

```nginx
# SSL configuration
server {
    listen 443 ssl;
    server_name api.nakku.com;
    
    ssl_certificate /etc/nginx/ssl/nakku.crt;
    ssl_certificate_key /etc/nginx/ssl/nakku.key;
    
    location / {
        proxy_pass http://api_gateway;
    }
}
```

### 3. **Database Security**

```sql
-- Create production users with limited privileges
CREATE USER 'nakku_user'@'%' IDENTIFIED BY 'secure_password';
GRANT SELECT, INSERT, UPDATE, DELETE ON nakku_user.* TO 'nakku_user'@'%';

CREATE USER 'nakku_product'@'%' IDENTIFIED BY 'secure_password';
GRANT SELECT, INSERT, UPDATE, DELETE ON nakku_product.* TO 'nakku_product'@'%';
```

## Deployment Automation

### 1. **CI/CD Pipeline**

```yaml
# .github/workflows/deploy.yml
name: Deploy to Production

on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Build Docker images
        run: |
          docker build -t nakku/api-gateway:latest ./api-gateway
          docker build -t nakku/user-service:latest ./services/user
      
      - name: Deploy to production
        run: |
          docker-compose -f docker-compose.prod.yml up -d
```

### 2. **Blue-Green Deployment**

```bash
# Blue-Green deployment script
#!/bin/bash

# Deploy to green environment
docker-compose -f docker-compose.green.yml up -d

# Health check
curl -f http://green.nakku.com/health

# Switch traffic
nginx -s reload

# Deploy to blue environment
docker-compose -f docker-compose.blue.yml up -d
```

## Cost Optimization

### 1. **Resource Allocation**

```yaml
# Resource limits per service
API Gateway: 512MB RAM, 0.5 CPU
User Service: 256MB RAM, 0.25 CPU
Product Service: 256MB RAM, 0.25 CPU
Order Service: 512MB RAM, 0.5 CPU
Cart Service: 128MB RAM, 0.1 CPU
```

### 2. **Auto Scaling Policies**

```yaml
# Kubernetes HPA
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: user-service-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: user-service
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
```

This production deployment strategy ensures high availability, scalability, and maintainability of your microservices architecture. Each service can be deployed independently, scaled based on demand, and maintained without affecting other services.
