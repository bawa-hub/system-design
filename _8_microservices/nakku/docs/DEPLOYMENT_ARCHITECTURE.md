# Deployment Architecture Overview

## Production Deployment Scenarios

### 1. **Single Server Deployment** (Development/Small Scale)
```
┌─────────────────────────────────────────────────────────────┐
│                    Single Server                            │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │ API Gateway │  │ User Service │  │Product Svc  │        │
│  │   :8080     │  │   :8081     │  │   :8082     │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
│                                                             │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │Order Service│  │ Cart Service │  │Payment Svc  │        │
│  │   :8084     │  │   :8085     │  │   :8086     │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
│                                                             │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │   MySQL     │  │    Kafka    │  │    Redis    │        │
│  │   :3306     │  │   :9092     │  │   :6379     │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
└─────────────────────────────────────────────────────────────┘
```

### 2. **Service-per-Server Deployment** (Medium Scale)
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Server 1      │    │   Server 2      │    │   Server 3      │
│  ┌─────────────┐│    │  ┌─────────────┐│    │  ┌─────────────┐│
│  │ API Gateway ││    │  │User Service ││    │  │Product Svc  ││
│  │   :8080     ││    │  │   :8081     ││    │  │   :8082     ││
│  └─────────────┘│    │  └─────────────┘│    │  └─────────────┘│
│                 │    │  ┌─────────────┐│    │  ┌─────────────┐│
│                 │    │  │MySQL (User) ││    │  │MySQL (Prod) ││
│                 │    │  │   :3306     ││    │  │   :3306     ││
│                 │    │  └─────────────┘│    │  └─────────────┘│
└─────────────────┘    └─────────────────┘    └─────────────────┘

┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Server 4      │    │   Server 5      │    │   Server 6      │
│  ┌─────────────┐│    │  ┌─────────────┐│    │  ┌─────────────┐│
│  │Order Service││    │  │ Cart Service ││    │  │Payment Svc  ││
│  │   :8084     ││    │  │   :8085     ││    │  │   :8086     ││
│  └─────────────┘│    │  └─────────────┘│    │  └─────────────┘│
│  ┌─────────────┐│    │  ┌─────────────┐│    │  ┌─────────────┐│
│  │MySQL (Order)││    │  │MySQL (Cart) ││    │  │MySQL (Pay)  ││
│  │   :3306     ││    │  │   :3306     ││    │  │   :3306     ││
│  └─────────────┘│    │  └─────────────┘│    │  └─────────────┘│
└─────────────────┘    └─────────────────┘    └─────────────────┘

┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Server 7      │    │   Server 8      │    │   Server 9      │
│  ┌─────────────┐│    │  ┌─────────────┐│    │  ┌─────────────┐│
│  │Delivery Svc ││    │  │Notification ││    │  │Location Svc ││
│  │   :8087     ││    │  │   :8088     ││    │  │   :8089     ││
│  └─────────────┘│    │  └─────────────┘│    │  └─────────────┘│
│  ┌─────────────┐│    │  ┌─────────────┐│    │  ┌─────────────┐│
│  │MySQL (Del)  ││    │  │MySQL (Notif)││    │  │MySQL (Loc)  ││
│  │   :3306     ││    │  │   :3306     ││    │  │   :3306     ││
│  └─────────────┘│    │  └─────────────┘│    │  └─────────────┘│
└─────────────────┘    └─────────────────┘    └─────────────────┘

┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Server 10     │    │   Server 11     │    │   Server 12     │
│  ┌─────────────┐│    │  ┌─────────────┐│    │  ┌─────────────┐│
│  │Analytics Svc││    │  │    Kafka    ││    │  │    Redis    ││
│  │   :8090     ││    │  │   :9092     ││    │  │   :6379     ││
│  └─────────────┘│    │  └─────────────┘│    │  └─────────────┘│
│  ┌─────────────┐│    │  ┌─────────────┐│    │  ┌─────────────┐│
│  │MySQL (Anal) ││    │  │   Zookeeper ││    │  │Redis Cluster││
│  │   :3306     ││    │  │   :2181     ││    │  │   :6379     ││
│  └─────────────┘│    │  └─────────────┘│    │  └─────────────┘│
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

### 3. **Multi-Server with Load Balancing** (Production Scale)
```
┌─────────────────────────────────────────────────────────────┐
│                    Load Balancer                            │
│                  (NGINX/HAProxy)                            │
└─────────────────┬─────────────────┬─────────────────────────┘
                  │                 │
                  ▼                 ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   API Gateway   │    │   API Gateway   │    │   API Gateway   │
│   Server 1      │    │   Server 2      │    │   Server 3      │
│   :8080         │    │   :8080         │    │   :8080         │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   User Service  │    │ Product Service │    │Inventory Service│
│   Server 4      │    │   Server 5      │    │   Server 6      │
│   :8081         │    │   :8082         │    │   :8083         │
└─────────────────┘    └─────────────────┘    └─────────────────┘

┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  Order Service  │    │  Cart Service   │    │Payment Service  │
│   Server 7      │    │   Server 8      │    │   Server 9      │
│   :8084         │    │   :8085         │    │   :8086         │
└─────────────────┘    └─────────────────┘    └─────────────────┘

┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│Delivery Service │    │Notification Svc │    │ Location Service│
│   Server 10     │    │   Server 11     │    │   Server 12     │
│   :8087         │    │   :8088         │    │   :8089         │
└─────────────────┘    └─────────────────┘    └─────────────────┘

┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│Analytics Service│    │   MySQL Master  │    │   MySQL Slave   │
│   Server 13     │    │   Server 14     │    │   Server 15     │
│   :8090         │    │   :3306         │    │   :3306         │
└─────────────────┘    └─────────────────┘    └─────────────────┘

┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Kafka Broker  │    │   Kafka Broker  │    │   Kafka Broker  │
│   Server 16     │    │   Server 17     │    │   Server 18     │
│   :9092         │    │   :9092         │    │   :9092         │
└─────────────────┘    └─────────────────┘    └─────────────────┘

┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Redis Node    │    │   Redis Node    │    │   Redis Node    │
│   Server 19     │    │   Server 20     │    │   Server 21     │
│   :6379         │    │   :6379         │    │   :6379         │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

### 4. **Kubernetes Deployment** (Enterprise Scale)
```
┌─────────────────────────────────────────────────────────────┐
│                    Kubernetes Cluster                       │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │   Node 1    │  │   Node 2    │  │   Node 3    │        │
│  │             │  │             │  │             │        │
│  │ ┌─────────┐ │  │ ┌─────────┐ │  │ ┌─────────┐ │        │
│  │ │API GW   │ │  │ │API GW   │ │  │ │API GW   │ │        │
│  │ │Pod      │ │  │ │Pod      │ │  │ │Pod      │ │        │
│  │ └─────────┘ │  │ └─────────┘ │  │ └─────────┘ │        │
│  │             │  │             │  │             │        │
│  │ ┌─────────┐ │  │ ┌─────────┐ │  │ ┌─────────┐ │        │
│  │ │User Svc │ │  │ │Product  │ │  │ │Order    │ │        │
│  │ │Pod      │ │  │ │Svc Pod  │ │  │ │Svc Pod  │ │        │
│  │ └─────────┘ │  │ └─────────┘ │  │ └─────────┘ │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
│                                                             │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │   Node 4    │  │   Node 5    │  │   Node 6    │        │
│  │             │  │             │  │             │        │
│  │ ┌─────────┐ │  │ ┌─────────┐ │  │ ┌─────────┐ │        │
│  │ │MySQL    │ │  │ │Kafka    │ │  │ │Redis    │ │        │
│  │ │Stateful │ │  │ │Stateful │ │  │ │Stateful │ │        │
│  │ │Set      │ │  │ │Set      │ │  │ │Set      │ │        │
│  │ └─────────┘ │  │ └─────────┘ │  │ └─────────┘ │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
└─────────────────────────────────────────────────────────────┘
```

## Service Communication Flow

### 1. **Request Flow**
```
Client → Load Balancer → API Gateway → Microservice → Database
   ↓           ↓              ↓            ↓           ↓
  HTTPS      HTTP          HTTP         HTTP        MySQL
```

### 2. **Event Flow**
```
Microservice A → Kafka → Microservice B → Kafka → Microservice C
      ↓           ↓            ↓           ↓            ↓
   Publish    Message      Consume     Publish      Consume
```

### 3. **Data Flow**
```
API Gateway → User Service → MySQL (User DB)
     ↓              ↓
   Route        Validate
     ↓              ↓
Product Service → MySQL (Product DB)
     ↓              ↓
   Process        Store
     ↓              ↓
Order Service → MySQL (Order DB)
```

## Deployment Commands

### 1. **Single Server Deployment**
```bash
# Start all services on one server
docker-compose up -d

# Check status
docker-compose ps

# View logs
docker-compose logs -f
```

### 2. **Multi-Server Deployment**
```bash
# Deploy to different servers
# Server 1: API Gateway
docker-compose -f docker-compose.gateway.yml up -d

# Server 2: User Service
docker-compose -f docker-compose.user.yml up -d

# Server 3: Product Service
docker-compose -f docker-compose.product.yml up -d

# Server 4: Database
docker-compose -f docker-compose.database.yml up -d

# Server 5: Kafka
docker-compose -f docker-compose.kafka.yml up -d
```

### 3. **Kubernetes Deployment**
```bash
# Deploy to Kubernetes
kubectl apply -f k8s/namespace.yml
kubectl apply -f k8s/configmap.yml
kubectl apply -f k8s/secrets.yml
kubectl apply -f k8s/deployments/
kubectl apply -f k8s/services/
kubectl apply -f k8s/ingress.yml

# Check status
kubectl get pods -n nakku
kubectl get services -n nakku
kubectl get ingress -n nakku
```

## Scaling Commands

### 1. **Horizontal Scaling**
```bash
# Scale API Gateway
kubectl scale deployment api-gateway --replicas=3 -n nakku

# Scale User Service
kubectl scale deployment user-service --replicas=2 -n nakku

# Scale Product Service
kubectl scale deployment product-service --replicas=2 -n nakku
```

### 2. **Auto Scaling**
```bash
# Create HPA for API Gateway
kubectl apply -f k8s/hpa/api-gateway-hpa.yml

# Create HPA for User Service
kubectl apply -f k8s/hpa/user-service-hpa.yml

# Check HPA status
kubectl get hpa -n nakku
```

## Monitoring and Health Checks

### 1. **Service Health**
```bash
# Check API Gateway health
curl http://api-gateway:8080/health

# Check User Service health
curl http://user-service:8081/health

# Check Product Service health
curl http://product-service:8082/health
```

### 2. **Database Health**
```bash
# Check MySQL health
mysql -h mysql-server -u nakku -p -e "SELECT 1"

# Check Redis health
redis-cli -h redis-server ping
```

### 3. **Kafka Health**
```bash
# Check Kafka health
kafka-topics --bootstrap-server kafka-server:9092 --list

# Check Kafka consumer groups
kafka-consumer-groups --bootstrap-server kafka-server:9092 --list
```

This deployment architecture provides flexibility to scale from a single server development environment to a multi-server production environment with proper load balancing, high availability, and monitoring.
