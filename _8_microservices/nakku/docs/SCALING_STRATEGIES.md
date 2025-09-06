# Scaling Strategies for Microservices

## Overview

This document outlines various scaling strategies for the Nakku Quick Commerce microservices architecture, including horizontal scaling, vertical scaling, database scaling, and performance optimization techniques.

## Scaling Dimensions

### 1. **Horizontal Scaling (Scale Out)**
Adding more instances of services to handle increased load.

### 2. **Vertical Scaling (Scale Up)**
Increasing resources (CPU, RAM) of existing service instances.

### 3. **Database Scaling**
Scaling database infrastructure to handle increased data and queries.

### 4. **Caching Scaling**
Implementing multi-level caching strategies.

## Service-Level Scaling

### 1. **API Gateway Scaling**

#### Load Balancer Configuration
```nginx
# nginx.conf
upstream api_gateway {
    least_conn;
    server api-gateway-1:8080 max_fails=3 fail_timeout=30s;
    server api-gateway-2:8080 max_fails=3 fail_timeout=30s;
    server api-gateway-3:8080 max_fails=3 fail_timeout=30s;
    server api-gateway-4:8080 max_fails=3 fail_timeout=30s;
}

server {
    listen 80;
    location / {
        proxy_pass http://api_gateway;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

#### Auto Scaling Configuration
```yaml
# Kubernetes HPA
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: api-gateway-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: api-gateway
  minReplicas: 3
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
```

### 2. **User Service Scaling**

#### Scaling Based on User Load
```yaml
# user-service-deployment.yml
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
        env:
        - name: DB_HOST
          value: "mysql-cluster"
        - name: KAFKA_BROKERS
          value: "kafka-1:9092,kafka-2:9092,kafka-3:9092"
```

#### Custom Metrics Scaling
```yaml
# Custom metrics HPA
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
  maxReplicas: 8
  metrics:
  - type: Pods
    pods:
      metric:
        name: active_users_per_pod
      target:
        type: AverageValue
        averageValue: "100"
```

### 3. **Product Service Scaling**

#### Read Replica Scaling
```go
// Product service with read replicas
type ProductService struct {
    masterDB *gorm.DB
    readDB   *gorm.DB
    cache    *redis.Client
}

func (s *ProductService) GetProduct(ctx context.Context, productID string) (*Product, error) {
    // Try cache first
    cached, err := s.cache.Get(ctx, fmt.Sprintf("product:%s", productID)).Result()
    if err == nil {
        var product Product
        if err := json.Unmarshal([]byte(cached), &product); err == nil {
            return &product, nil
        }
    }
    
    // Try read replica
    var product Product
    err = s.readDB.WithContext(ctx).First(&product, productID).Error
    if err != nil {
        return nil, err
    }
    
    // Cache the result
    productJSON, _ := json.Marshal(product)
    s.cache.Set(ctx, fmt.Sprintf("product:%s", productID), productJSON, 5*time.Minute)
    
    return &product, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, req *CreateProductRequest) (*Product, error) {
    // Write to master
    product, err := s.masterDB.WithContext(ctx).Create(product)
    if err != nil {
        return nil, err
    }
    
    // Invalidate cache
    s.cache.Del(ctx, fmt.Sprintf("product:%s", product.ID))
    
    return product, nil
}
```

### 4. **Order Service Scaling**

#### Database Sharding
```go
// Order service with database sharding
type OrderService struct {
    shards map[string]*gorm.DB
    cache  *redis.Client
}

func (s *OrderService) getShard(userID string) *gorm.DB {
    // Simple hash-based sharding
    hash := fnv.New32a()
    hash.Write([]byte(userID))
    shardIndex := hash.Sum32() % uint32(len(s.shards))
    
    shardKeys := make([]string, 0, len(s.shards))
    for k := range s.shards {
        shardKeys = append(shardKeys, k)
    }
    
    return s.shards[shardKeys[shardIndex]]
}

func (s *OrderService) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*Order, error) {
    shard := s.getShard(req.UserID)
    
    order := &Order{
        UserID: req.UserID,
        Items:  req.Items,
        Status: "pending",
    }
    
    return s.orderRepo.Create(ctx, shard, order)
}
```

## Database Scaling Strategies

### 1. **Read Replicas**

#### MySQL Master-Slave Configuration
```yaml
# mysql-master.yml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mysql-master-config
data:
  my.cnf: |
    [mysqld]
    server-id = 1
    log-bin = mysql-bin
    binlog-format = ROW
    gtid-mode = ON
    enforce-gtid-consistency = ON

---
# mysql-slave.yml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mysql-slave-config
data:
  my.cnf: |
    [mysqld]
    server-id = 2
    relay-log = mysql-relay-bin
    read-only = 1
    gtid-mode = ON
    enforce-gtid-consistency = ON
```

#### Connection Pooling
```go
// Database connection pool configuration
type DatabaseConfig struct {
    MasterDSN string
    SlaveDSN  string
    MaxConns  int
    MinConns  int
}

func NewDatabasePool(config DatabaseConfig) (*gorm.DB, *gorm.DB, error) {
    // Master connection
    masterDB, err := gorm.Open(mysql.Open(config.MasterDSN), &gorm.Config{})
    if err != nil {
        return nil, nil, err
    }
    
    masterSQL, _ := masterDB.DB()
    masterSQL.SetMaxOpenConns(config.MaxConns)
    masterSQL.SetMaxIdleConns(config.MinConns)
    
    // Slave connection
    slaveDB, err := gorm.Open(mysql.Open(config.SlaveDSN), &gorm.Config{})
    if err != nil {
        return nil, nil, err
    }
    
    slaveSQL, _ := slaveDB.DB()
    slaveSQL.SetMaxOpenConns(config.MaxConns)
    slaveSQL.SetMaxIdleConns(config.MinConns)
    
    return masterDB, slaveDB, nil
}
```

### 2. **Database Sharding**

#### Horizontal Sharding
```go
// Shard configuration
type ShardConfig struct {
    Shards map[string]string // shard_name -> dsn
}

type ShardedRepository struct {
    shards map[string]*gorm.DB
    config ShardConfig
}

func (r *ShardedRepository) getShard(key string) *gorm.DB {
    hash := fnv.New32a()
    hash.Write([]byte(key))
    shardIndex := hash.Sum32() % uint32(len(r.shards))
    
    shardKeys := make([]string, 0, len(r.shards))
    for k := range r.shards {
        shardKeys = append(shardKeys, k)
    }
    
    return r.shards[shardKeys[shardIndex]]
}

func (r *ShardedRepository) Create(ctx context.Context, entity interface{}, shardKey string) error {
    shard := r.getShard(shardKey)
    return shard.WithContext(ctx).Create(entity).Error
}
```

### 3. **Caching Strategies**

#### Multi-Level Caching
```go
// Multi-level cache implementation
type CacheManager struct {
    l1Cache *sync.Map          // In-memory cache
    l2Cache *redis.Client      // Redis cache
    l3Cache *gorm.DB          // Database
}

func (cm *CacheManager) Get(ctx context.Context, key string, dest interface{}) error {
    // L1 Cache (In-memory)
    if cached, ok := cm.l1Cache.Load(key); ok {
        return json.Unmarshal(cached.([]byte), dest)
    }
    
    // L2 Cache (Redis)
    cached, err := cm.l2Cache.Get(ctx, key).Result()
    if err == nil {
        // Store in L1 cache
        cm.l1Cache.Store(key, []byte(cached))
        return json.Unmarshal([]byte(cached), dest)
    }
    
    // L3 Cache (Database)
    // This would be implemented by the specific service
    return fmt.Errorf("not found in cache")
}

func (cm *CacheManager) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
    data, err := json.Marshal(value)
    if err != nil {
        return err
    }
    
    // Set in L1 cache
    cm.l1Cache.Store(key, data)
    
    // Set in L2 cache
    return cm.l2Cache.Set(ctx, key, data, ttl).Err()
}
```

## Performance Optimization

### 1. **Connection Pooling**

#### HTTP Client Pooling
```go
// HTTP client with connection pooling
type ServiceClient struct {
    client *http.Client
}

func NewServiceClient() *ServiceClient {
    transport := &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
        DisableKeepAlives:   false,
    }
    
    return &ServiceClient{
        client: &http.Client{
            Transport: transport,
            Timeout:   30 * time.Second,
        },
    }
}
```

#### Database Connection Pooling
```go
// Database connection pool
func setupDatabasePool(dsn string) *gorm.DB {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    
    sqlDB, err := db.DB()
    if err != nil {
        panic(err)
    }
    
    // Connection pool settings
    sqlDB.SetMaxOpenConns(100)    // Maximum number of open connections
    sqlDB.SetMaxIdleConns(10)     // Maximum number of idle connections
    sqlDB.SetConnMaxLifetime(time.Hour) // Maximum connection lifetime
    
    return db
}
```

### 2. **Async Processing**

#### Background Job Processing
```go
// Background job processor
type JobProcessor struct {
    jobQueue chan Job
    workers  int
}

type Job struct {
    ID   string
    Type string
    Data interface{}
}

func (jp *JobProcessor) Start(ctx context.Context) {
    for i := 0; i < jp.workers; i++ {
        go jp.worker(ctx, i)
    }
}

func (jp *JobProcessor) worker(ctx context.Context, workerID int) {
    for {
        select {
        case job := <-jp.jobQueue:
            jp.processJob(ctx, job)
        case <-ctx.Done():
            return
        }
    }
}

func (jp *JobProcessor) processJob(ctx context.Context, job Job) {
    switch job.Type {
    case "send_notification":
        // Process notification
    case "update_analytics":
        // Process analytics
    case "cleanup_expired_sessions":
        // Cleanup sessions
    }
}
```

### 3. **Batch Processing**

#### Batch Event Processing
```go
// Batch event processor
type BatchProcessor struct {
    batchSize    int
    batchTimeout time.Duration
    events       []Event
    mutex        sync.Mutex
    processor    func([]Event) error
}

func (bp *BatchProcessor) AddEvent(event Event) {
    bp.mutex.Lock()
    defer bp.mutex.Unlock()
    
    bp.events = append(bp.events, event)
    
    if len(bp.events) >= bp.batchSize {
        go bp.processBatch()
    }
}

func (bp *BatchProcessor) processBatch() {
    bp.mutex.Lock()
    events := bp.events
    bp.events = nil
    bp.mutex.Unlock()
    
    if err := bp.processor(events); err != nil {
        // Handle error
        log.Printf("Batch processing failed: %v", err)
    }
}
```

## Monitoring and Alerting

### 1. **Resource Monitoring**

#### Prometheus Metrics
```go
// Custom metrics
var (
    activeConnections = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "active_connections",
            Help: "Number of active connections",
        },
        []string{"service", "type"},
    )
    
    requestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "request_duration_seconds",
            Help: "Request duration in seconds",
        },
        []string{"service", "method", "endpoint"},
    )
)

func init() {
    prometheus.MustRegister(activeConnections)
    prometheus.MustRegister(requestDuration)
}
```

#### Health Checks
```go
// Health check endpoint
func (s *UserService) HealthCheck(c *gin.Context) {
    health := map[string]interface{}{
        "status":    "healthy",
        "timestamp": time.Now().Unix(),
        "service":   "user-service",
        "version":   "1.0.0",
    }
    
    // Check database
    if err := s.db.DB().Ping(); err != nil {
        health["status"] = "unhealthy"
        health["database"] = "unhealthy"
        c.JSON(http.StatusServiceUnavailable, health)
        return
    }
    
    // Check Redis
    if err := s.redis.Ping(c.Request.Context()).Err(); err != nil {
        health["status"] = "unhealthy"
        health["redis"] = "unhealthy"
        c.JSON(http.StatusServiceUnavailable, health)
        return
    }
    
    // Check Kafka
    if err := s.kafkaProducer.Ping(); err != nil {
        health["status"] = "unhealthy"
        health["kafka"] = "unhealthy"
        c.JSON(http.StatusServiceUnavailable, health)
        return
    }
    
    c.JSON(http.StatusOK, health)
}
```

### 2. **Auto Scaling Triggers**

#### CPU-Based Scaling
```yaml
# CPU-based HPA
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

#### Memory-Based Scaling
```yaml
# Memory-based HPA
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: product-service-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: product-service
  minReplicas: 2
  maxReplicas: 8
  metrics:
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
```

#### Custom Metrics Scaling
```yaml
# Custom metrics HPA
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: order-service-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: order-service
  minReplicas: 2
  maxReplicas: 15
  metrics:
  - type: Pods
    pods:
      metric:
        name: orders_per_minute
      target:
        type: AverageValue
        averageValue: "50"
```

## Cost Optimization

### 1. **Resource Optimization**

#### Resource Requests and Limits
```yaml
# Optimized resource configuration
apiVersion: apps/v1
kind: Deployment
metadata:
  name: user-service
spec:
  template:
    spec:
      containers:
      - name: user-service
        resources:
          requests:
            memory: "128Mi"
            cpu: "100m"
          limits:
            memory: "256Mi"
            cpu: "200m"
```

#### Spot Instances
```yaml
# Spot instance configuration
apiVersion: v1
kind: Node
metadata:
  labels:
    node.kubernetes.io/instance-type: spot
spec:
  taints:
  - key: spot-instance
    value: "true"
    effect: NoSchedule
```

### 2. **Efficient Scaling**

#### Predictive Scaling
```go
// Predictive scaling based on historical data
type PredictiveScaler struct {
    historicalData []MetricData
    currentLoad    float64
}

type MetricData struct {
    Timestamp time.Time
    CPU       float64
    Memory    float64
    Requests  int
}

func (ps *PredictiveScaler) PredictScaling() int {
    // Analyze historical data
    // Predict future load
    // Return recommended replica count
    return ps.calculateOptimalReplicas()
}
```

This comprehensive scaling strategy ensures that your microservices architecture can handle varying loads efficiently while maintaining performance and cost-effectiveness.
