# Lesson 10: Monitoring and Observability

## 🎯 Learning Objectives
By the end of this lesson, you will understand:
- Comprehensive monitoring strategies for Kafka applications
- Metrics collection and analysis techniques
- Logging and distributed tracing
- Alerting and incident response
- Performance dashboards and visualization
- Health checks and status pages
- Production monitoring best practices

## 📚 Theory Section

### 1. **Monitoring Fundamentals**

#### **The Three Pillars of Observability:**
1. **Metrics**: Quantitative data about system behavior
2. **Logs**: Detailed records of events and activities
3. **Traces**: Request flow through distributed systems

#### **Monitoring Hierarchy:**
- **Infrastructure Monitoring**: CPU, memory, disk, network
- **Application Monitoring**: Business metrics, performance
- **User Experience Monitoring**: End-to-end user journey
- **Security Monitoring**: Threats, vulnerabilities, compliance

### 2. **Kafka Metrics Collection**

#### **Producer Metrics:**
```go
type ProducerMetrics struct {
    MessagesSent      int64
    BytesSent         int64
    SendLatency       time.Duration
    RetryCount        int64
    ErrorCount        int64
    BatchSize         int64
    CompressionRatio  float64
}

func (pm *ProducerMetrics) RecordSend(latency time.Duration, bytes int, compressed bool) {
    atomic.AddInt64(&pm.MessagesSent, 1)
    atomic.AddInt64(&pm.BytesSent, int64(bytes))
    
    // Update latency (simplified)
    pm.SendLatency = latency
    
    if compressed {
        pm.CompressionRatio = calculateCompressionRatio(bytes)
    }
}
```

#### **Consumer Metrics:**
```go
type ConsumerMetrics struct {
    MessagesConsumed  int64
    BytesConsumed     int64
    ProcessingLatency time.Duration
    ConsumerLag       int64
    RebalanceCount    int64
    ErrorCount        int64
    OffsetCommitted   int64
}

func (cm *ConsumerMetrics) RecordMessage(latency time.Duration, bytes int) {
    atomic.AddInt64(&cm.MessagesConsumed, 1)
    atomic.AddInt64(&cm.BytesConsumed, int64(bytes))
    cm.ProcessingLatency = latency
}
```

#### **Broker Metrics:**
```go
type BrokerMetrics struct {
    CPUUsage          float64
    MemoryUsage       float64
    DiskUsage         float64
    NetworkIO         int64
    ActiveConnections int64
    RequestLatency    time.Duration
    ErrorRate         float64
}
```

### 3. **Logging and Structured Logging**

#### **Structured Logging Implementation:**
```go
import (
    "log/slog"
    "os"
)

type Logger struct {
    logger *slog.Logger
}

func NewLogger() *Logger {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
    }))
    return &Logger{logger: logger}
}

func (l *Logger) LogProducerEvent(event string, metrics ProducerMetrics) {
    l.logger.Info("producer_event",
        "event", event,
        "messages_sent", metrics.MessagesSent,
        "bytes_sent", metrics.BytesSent,
        "send_latency_ms", metrics.SendLatency.Milliseconds(),
        "retry_count", metrics.RetryCount,
        "error_count", metrics.ErrorCount,
    )
}

func (l *Logger) LogConsumerEvent(event string, metrics ConsumerMetrics) {
    l.logger.Info("consumer_event",
        "event", event,
        "messages_consumed", metrics.MessagesConsumed,
        "bytes_consumed", metrics.BytesConsumed,
        "processing_latency_ms", metrics.ProcessingLatency.Milliseconds(),
        "consumer_lag", metrics.ConsumerLag,
        "rebalance_count", metrics.RebalanceCount,
    )
}
```

#### **Log Aggregation:**
```go
type LogAggregator struct {
    logs []LogEntry
    mutex sync.RWMutex
}

type LogEntry struct {
    Timestamp time.Time
    Level     string
    Message   string
    Fields    map[string]interface{}
}

func (la *LogAggregator) AddLog(level, message string, fields map[string]interface{}) {
    la.mutex.Lock()
    defer la.mutex.Unlock()
    
    entry := LogEntry{
        Timestamp: time.Now(),
        Level:     level,
        Message:   message,
        Fields:    fields,
    }
    
    la.logs = append(la.logs, entry)
    
    // Keep only last 1000 logs
    if len(la.logs) > 1000 {
        la.logs = la.logs[1:]
    }
}
```

### 4. **Distributed Tracing**

#### **Tracing Implementation:**
```go
type TraceContext struct {
    TraceID    string
    SpanID     string
    ParentID   string
    StartTime  time.Time
    EndTime    time.Time
    Tags       map[string]string
    Logs       []TraceLog
}

type TraceLog struct {
    Timestamp time.Time
    Event     string
    Fields    map[string]interface{}
}

func (tc *TraceContext) StartSpan(operation string) *TraceContext {
    return &TraceContext{
        TraceID:   tc.TraceID,
        SpanID:    generateSpanID(),
        ParentID:  tc.SpanID,
        StartTime: time.Now(),
        Tags:      make(map[string]string),
        Logs:      []TraceLog{},
    }
}

func (tc *TraceContext) AddTag(key, value string) {
    tc.Tags[key] = value
}

func (tc *TraceContext) AddLog(event string, fields map[string]interface{}) {
    tc.Logs = append(tc.Logs, TraceLog{
        Timestamp: time.Now(),
        Event:     event,
        Fields:    fields,
    })
}

func (tc *TraceContext) Finish() {
    tc.EndTime = time.Now()
}
```

### 5. **Alerting and Incident Response**

#### **Alert Manager:**
```go
type AlertManager struct {
    rules       []AlertRule
    channels    []AlertChannel
    mutex       sync.RWMutex
}

type AlertRule struct {
    Name        string
    Condition   func(metrics interface{}) bool
    Severity    string
    Message     string
    Cooldown    time.Duration
    LastFired   time.Time
}

type AlertChannel interface {
    Send(alert Alert) error
}

type Alert struct {
    Rule        string
    Severity    string
    Message     string
    Timestamp   time.Time
    Metrics     interface{}
}

func (am *AlertManager) CheckAlerts(metrics interface{}) {
    am.mutex.RLock()
    defer am.mutex.RUnlock()
    
    for _, rule := range am.rules {
        if rule.Condition(metrics) {
            if time.Since(rule.LastFired) > rule.Cooldown {
                alert := Alert{
                    Rule:      rule.Name,
                    Severity:  rule.Severity,
                    Message:   rule.Message,
                    Timestamp: time.Now(),
                    Metrics:   metrics,
                }
                
                am.sendAlert(alert)
                rule.LastFired = time.Now()
            }
        }
    }
}
```

#### **Alert Channels:**
```go
type EmailChannel struct {
    SMTPHost string
    SMTPPort int
    Username string
    Password string
    To       []string
}

func (ec *EmailChannel) Send(alert Alert) error {
    // Implement email sending
    return nil
}

type SlackChannel struct {
    WebhookURL string
    Channel    string
}

func (sc *SlackChannel) Send(alert Alert) error {
    // Implement Slack notification
    return nil
}

type PagerDutyChannel struct {
    IntegrationKey string
}

func (pc *PagerDutyChannel) Send(alert Alert) error {
    // Implement PagerDuty alert
    return nil
}
```

### 6. **Performance Dashboards**

#### **Dashboard Data Structure:**
```go
type DashboardData struct {
    Timestamp    time.Time
    Producer     ProducerMetrics
    Consumer     ConsumerMetrics
    Broker       BrokerMetrics
    System       SystemMetrics
    Alerts       []Alert
}

type SystemMetrics struct {
    CPUUsage     float64
    MemoryUsage  float64
    DiskUsage    float64
    NetworkIO    int64
    Uptime       time.Duration
}

func (dd *DashboardData) GenerateReport() DashboardReport {
    return DashboardReport{
        Timestamp:     dd.Timestamp,
        Health:        dd.calculateHealth(),
        Performance:   dd.calculatePerformance(),
        Alerts:        dd.Alerts,
        Recommendations: dd.generateRecommendations(),
    }
}
```

#### **Health Check Implementation:**
```go
type HealthChecker struct {
    checks []HealthCheck
}

type HealthCheck struct {
    Name        string
    Check       func() HealthStatus
    Critical    bool
    LastCheck   time.Time
    LastStatus  HealthStatus
}

type HealthStatus struct {
    Healthy   bool
    Message   string
    Timestamp time.Time
    Duration  time.Duration
}

func (hc *HealthChecker) RunChecks() map[string]HealthStatus {
    results := make(map[string]HealthStatus)
    
    for _, check := range hc.checks {
        start := time.Now()
        status := check.Check()
        status.Duration = time.Since(start)
        status.Timestamp = time.Now()
        
        results[check.Name] = status
        check.LastCheck = time.Now()
        check.LastStatus = status
    }
    
    return results
}
```

### 7. **Monitoring Best Practices**

#### **Metric Naming Conventions:**
```go
// Use consistent naming patterns
const (
    // Producer metrics
    PRODUCER_MESSAGES_SENT = "kafka.producer.messages.sent"
    PRODUCER_BYTES_SENT    = "kafka.producer.bytes.sent"
    PRODUCER_SEND_LATENCY  = "kafka.producer.send.latency"
    PRODUCER_ERROR_RATE    = "kafka.producer.error.rate"
    
    // Consumer metrics
    CONSUMER_MESSAGES_CONSUMED = "kafka.consumer.messages.consumed"
    CONSUMER_LAG              = "kafka.consumer.lag"
    CONSUMER_PROCESSING_TIME  = "kafka.consumer.processing.time"
    
    // System metrics
    SYSTEM_CPU_USAGE    = "system.cpu.usage"
    SYSTEM_MEMORY_USAGE = "system.memory.usage"
    SYSTEM_DISK_USAGE   = "system.disk.usage"
)
```

#### **Monitoring Configuration:**
```go
type MonitoringConfig struct {
    MetricsEnabled    bool
    LoggingEnabled    bool
    TracingEnabled    bool
    AlertingEnabled   bool
    
    MetricsInterval   time.Duration
    LogLevel          string
    TraceSampleRate   float64
    
    AlertRules        []AlertRule
    AlertChannels     []AlertChannel
}

func LoadMonitoringConfig(env string) MonitoringConfig {
    switch env {
    case "development":
        return getDevelopmentMonitoringConfig()
    case "staging":
        return getStagingMonitoringConfig()
    case "production":
        return getProductionMonitoringConfig()
    default:
        return getDefaultMonitoringConfig()
    }
}
```

## 🧪 Hands-on Experiments

### Experiment 1: Metrics Collection

**Goal**: Implement comprehensive metrics collection

**Steps**:
1. Create metrics structures for producer, consumer, and broker
2. Implement metrics collection in producer and consumer
3. Set up metrics aggregation and storage
4. Create metrics visualization

**Expected Results**:
- Real-time metrics collection
- Historical metrics storage
- Performance trend analysis

### Experiment 2: Structured Logging

**Goal**: Implement structured logging with correlation IDs

**Steps**:
1. Set up structured logging framework
2. Add correlation IDs to all log entries
3. Implement log aggregation
4. Create log analysis tools

**Expected Results**:
- Structured, searchable logs
- Request tracing across services
- Error correlation and analysis

### Experiment 3: Alerting System

**Goal**: Set up comprehensive alerting

**Steps**:
1. Define alert rules for different scenarios
2. Implement alert channels (email, Slack, PagerDuty)
3. Set up alert escalation policies
4. Test alert delivery

**Expected Results**:
- Proactive alerting for issues
- Multiple alert channels
- Escalation policies

### Experiment 4: Health Checks

**Goal**: Implement comprehensive health checks

**Steps**:
1. Create health check endpoints
2. Implement dependency health checks
3. Set up health check monitoring
4. Create status page

**Expected Results**:
- Real-time health status
- Dependency monitoring
- Public status page

## 📊 Monitoring Dashboard

### **Key Performance Indicators (KPIs):**

#### **Throughput Metrics:**
- **Messages per second**: Producer and consumer throughput
- **Bytes per second**: Data processing rate
- **Partitions per second**: Parallel processing rate

#### **Latency Metrics:**
- **P50, P95, P99 Latency**: Response time percentiles
- **End-to-end latency**: Complete message journey time
- **Processing latency**: Message processing time

#### **Error Metrics:**
- **Error rate**: Percentage of failed operations
- **Retry rate**: Percentage of operations requiring retry
- **Circuit breaker trips**: Fault tolerance indicators

#### **Resource Metrics:**
- **CPU utilization**: System resource usage
- **Memory usage**: Memory consumption patterns
- **Disk I/O**: Storage performance
- **Network I/O**: Network performance

### **Dashboard Implementation:**
```go
type MonitoringDashboard struct {
    metrics    *MetricsCollector
    alerts     *AlertManager
    health     *HealthChecker
    logger     *Logger
}

func (md *MonitoringDashboard) GenerateDashboard() DashboardData {
    return DashboardData{
        Timestamp: time.Now(),
        Producer:  md.metrics.GetProducerMetrics(),
        Consumer:  md.metrics.GetConsumerMetrics(),
        Broker:    md.metrics.GetBrokerMetrics(),
        System:    md.metrics.GetSystemMetrics(),
        Alerts:    md.alerts.GetActiveAlerts(),
    }
}
```

## 🎯 Key Takeaways

1. **Monitoring is essential** for production systems
2. **Metrics, logs, and traces** provide complete observability
3. **Structured logging** enables better debugging and analysis
4. **Alerting** enables proactive issue detection
5. **Health checks** provide real-time system status
6. **Dashboards** provide visibility into system behavior
7. **Monitoring configuration** should be environment-specific

## 📝 Lesson 10 Assessment Questions

1. **What are the three pillars of observability?**
2. **How do you implement structured logging?**
3. **What metrics should you collect for Kafka applications?**
4. **How do you set up effective alerting?**
5. **What is distributed tracing and why is it important?**
6. **How do you implement health checks?**
7. **What are the best practices for monitoring dashboards?**
8. **How do you handle monitoring in different environments?**

---

## 🎉 Phase 2 Completion

**Congratulations!** You have completed Phase 2: Intermediate Level - Advanced Configuration and Performance!

### **What You've Mastered:**
- ✅ Advanced Producer Configuration
- ✅ Advanced Consumer Configuration  
- ✅ Performance Tuning and Optimization
- ✅ Error Handling and Resilience
- ✅ Monitoring and Observability

### **Ready for Phase 3: Advanced Level**
- Stream Processing with Kafka Streams
- Schema Registry and Data Evolution
- Security and Authentication
- Multi-cluster and Cross-Datacenter

**You're now ready to build production-ready Kafka applications!** 🚀
