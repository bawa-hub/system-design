package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("📊 Kafka Monitoring and Observability Demo")
	fmt.Println(strings.Repeat("=", 55))
	fmt.Println("Comprehensive monitoring for production Kafka applications...")
	fmt.Println()

	// Run different monitoring demonstrations
	demonstrateMetricsCollection()
	demonstrateStructuredLogging()
	demonstrateAlertingSystem()
	demonstrateHealthChecks()
}

func demonstrateMetricsCollection() {
	fmt.Println("📈 Metrics Collection Demonstration")
	fmt.Println(strings.Repeat("-", 40))

	// Create metrics collector
	collector := NewMetricsCollector()
	
	// Simulate producer metrics
	fmt.Println("🔧 Simulating producer metrics...")
	for i := 0; i < 10; i++ {
		collector.RecordProducerSend(50*time.Millisecond, 1024, true)
		time.Sleep(100 * time.Millisecond)
	}
	
	// Simulate consumer metrics
	fmt.Println("🔧 Simulating consumer metrics...")
	for i := 0; i < 10; i++ {
		collector.RecordConsumerMessage(25*time.Millisecond, 1024)
		time.Sleep(100 * time.Millisecond)
	}
	
	// Display metrics
	producerMetrics := collector.GetProducerMetrics()
	consumerMetrics := collector.GetConsumerMetrics()
	
	fmt.Printf("📊 Producer Metrics:\n")
	fmt.Printf("   Messages Sent: %d\n", producerMetrics.MessagesSent)
	fmt.Printf("   Bytes Sent: %d\n", producerMetrics.BytesSent)
	fmt.Printf("   Avg Send Latency: %.2fms\n", producerMetrics.AvgSendLatency.Milliseconds())
	fmt.Printf("   Error Rate: %.2f%%\n", producerMetrics.ErrorRate)
	
	fmt.Printf("\n📊 Consumer Metrics:\n")
	fmt.Printf("   Messages Consumed: %d\n", consumerMetrics.MessagesConsumed)
	fmt.Printf("   Bytes Consumed: %d\n", consumerMetrics.BytesConsumed)
	fmt.Printf("   Avg Processing Latency: %.2fms\n", consumerMetrics.AvgProcessingLatency.Milliseconds())
	fmt.Printf("   Consumer Lag: %d\n", consumerMetrics.ConsumerLag)
	fmt.Println()
}

func demonstrateStructuredLogging() {
	fmt.Println("📝 Structured Logging Demonstration")
	fmt.Println(strings.Repeat("-", 40))

	// Create structured logger
	logger := NewStructuredLogger()
	
	// Simulate different log events
	fmt.Println("🔧 Simulating structured log events...")
	
	// Producer events
	logger.LogProducerEvent("message_sent", ProducerMetrics{
		MessagesSent: 100,
		BytesSent:    102400,
		AvgSendLatency: 50 * time.Millisecond,
		ErrorRate:    0.5,
	})
	
	// Consumer events
	logger.LogConsumerEvent("message_processed", ConsumerMetrics{
		MessagesConsumed: 95,
		BytesConsumed:    97280,
		AvgProcessingLatency: 25 * time.Millisecond,
		ConsumerLag:      5,
	})
	
	// Error events
	logger.LogError("producer_error", "Failed to send message", map[string]interface{}{
		"topic":     "user-events",
		"partition": 0,
		"offset":    12345,
		"error":     "timeout",
	})
	
	fmt.Println("✅ Structured logs generated (check console output)")
	fmt.Println()
}

func demonstrateAlertingSystem() {
	fmt.Println("🚨 Alerting System Demonstration")
	fmt.Println(strings.Repeat("-", 40))

	// Create alert manager
	alertManager := NewAlertManager()
	
	// Add alert rules
	alertManager.AddRule(AlertRule{
		Name:     "High Error Rate",
		Severity: "CRITICAL",
		Message:  "Error rate exceeds 5%",
		Condition: func(metrics interface{}) bool {
			if pm, ok := metrics.(ProducerMetrics); ok {
				return pm.ErrorRate > 5.0
			}
			return false
		},
	})
	
	alertManager.AddRule(AlertRule{
		Name:     "High Consumer Lag",
		Severity: "WARNING",
		Message:  "Consumer lag exceeds 1000 messages",
		Condition: func(metrics interface{}) bool {
			if cm, ok := metrics.(ConsumerMetrics); ok {
				return cm.ConsumerLag > 1000
			}
			return false
		},
	})
	
	// Simulate different metric scenarios
	fmt.Println("🔧 Testing alert conditions...")
	
	// Normal metrics (no alerts)
	fmt.Println("   Testing normal metrics...")
	normalProducer := ProducerMetrics{ErrorRate: 1.0}
	normalConsumer := ConsumerMetrics{ConsumerLag: 100}
	alertManager.CheckAlerts(normalProducer)
	alertManager.CheckAlerts(normalConsumer)
	
	// High error rate (should trigger alert)
	fmt.Println("   Testing high error rate...")
	highErrorProducer := ProducerMetrics{ErrorRate: 8.0}
	alertManager.CheckAlerts(highErrorProducer)
	
	// High consumer lag (should trigger alert)
	fmt.Println("   Testing high consumer lag...")
	highLagConsumer := ConsumerMetrics{ConsumerLag: 1500}
	alertManager.CheckAlerts(highLagConsumer)
	
	fmt.Println()
}

func demonstrateHealthChecks() {
	fmt.Println("🏥 Health Checks Demonstration")
	fmt.Println(strings.Repeat("-", 40))

	// Create health checker
	healthChecker := NewHealthChecker()
	
	// Add health checks
	healthChecker.AddCheck(HealthCheck{
		Name: "Kafka Connectivity",
		Check: func() HealthStatus {
			// Simulate Kafka connectivity check
			if rand.Float64() > 0.1 { // 90% success rate
				return HealthStatus{
					Healthy:   true,
					Message:   "Kafka broker is reachable",
					Timestamp: time.Now(),
				}
			}
			return HealthStatus{
				Healthy:   false,
				Message:   "Cannot connect to Kafka broker",
				Timestamp: time.Now(),
			}
		},
		Critical: true,
	})
	
	healthChecker.AddCheck(HealthCheck{
		Name: "Consumer Group Health",
		Check: func() HealthStatus {
			// Simulate consumer group health check
			if rand.Float64() > 0.2 { // 80% success rate
				return HealthStatus{
					Healthy:   true,
					Message:   "Consumer group is healthy",
					Timestamp: time.Now(),
				}
			}
			return HealthStatus{
				Healthy:   false,
				Message:   "Consumer group has issues",
				Timestamp: time.Now(),
			}
		},
		Critical: false,
	})
	
	// Run health checks
	fmt.Println("🔧 Running health checks...")
	results := healthChecker.RunChecks()
	
	for name, status := range results {
		if status.Healthy {
			fmt.Printf("   ✅ %s: %s\n", name, status.Message)
		} else {
			fmt.Printf("   ❌ %s: %s\n", name, status.Message)
		}
	}
	
	// Generate health summary
	summary := healthChecker.GenerateSummary()
	fmt.Printf("\n📊 Health Summary:\n")
	fmt.Printf("   Overall Health: %s\n", summary.OverallHealth)
	fmt.Printf("   Critical Issues: %d\n", summary.CriticalIssues)
	fmt.Printf("   Warning Issues: %d\n", summary.WarningIssues)
	fmt.Printf("   Total Checks: %d\n", summary.TotalChecks)
	fmt.Println()
}

// Metrics collection implementation
type MetricsCollector struct {
	producerMetrics ProducerMetrics
	consumerMetrics ConsumerMetrics
	brokerMetrics   BrokerMetrics
	mutex           sync.RWMutex
}

type ProducerMetrics struct {
	MessagesSent      int64
	BytesSent         int64
	TotalSendLatency  time.Duration
	SendCount         int64
	ErrorCount        int64
	RetryCount        int64
	CompressionRatio  float64
	AvgSendLatency    time.Duration
	ErrorRate         float64
}

type ConsumerMetrics struct {
	MessagesConsumed      int64
	BytesConsumed         int64
	TotalProcessingLatency time.Duration
	ProcessingCount       int64
	ConsumerLag           int64
	RebalanceCount        int64
	ErrorCount            int64
	AvgProcessingLatency  time.Duration
}

type BrokerMetrics struct {
	CPUUsage          float64
	MemoryUsage       float64
	DiskUsage         float64
	NetworkIO         int64
	ActiveConnections int64
	RequestLatency    time.Duration
	ErrorRate         float64
}

func NewMetricsCollector() *MetricsCollector {
	return &MetricsCollector{
		producerMetrics: ProducerMetrics{},
		consumerMetrics: ConsumerMetrics{},
		brokerMetrics:   BrokerMetrics{},
	}
}

func (mc *MetricsCollector) RecordProducerSend(latency time.Duration, bytes int, compressed bool) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	
	atomic.AddInt64(&mc.producerMetrics.MessagesSent, 1)
	atomic.AddInt64(&mc.producerMetrics.BytesSent, int64(bytes))
	atomic.AddInt64(&mc.producerMetrics.SendCount, 1)
	
	mc.producerMetrics.TotalSendLatency += latency
	
	if compressed {
		mc.producerMetrics.CompressionRatio = 0.6 // Simulated compression ratio
	}
}

func (mc *MetricsCollector) RecordConsumerMessage(latency time.Duration, bytes int) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	
	atomic.AddInt64(&mc.consumerMetrics.MessagesConsumed, 1)
	atomic.AddInt64(&mc.consumerMetrics.BytesConsumed, int64(bytes))
	atomic.AddInt64(&mc.consumerMetrics.ProcessingCount, 1)
	
	mc.consumerMetrics.TotalProcessingLatency += latency
}

func (mc *MetricsCollector) GetProducerMetrics() ProducerMetrics {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()
	
	metrics := mc.producerMetrics
	if metrics.SendCount > 0 {
		metrics.AvgSendLatency = time.Duration(int64(metrics.TotalSendLatency) / metrics.SendCount)
	}
	if metrics.MessagesSent > 0 {
		metrics.ErrorRate = float64(metrics.ErrorCount) / float64(metrics.MessagesSent) * 100
	}
	
	return metrics
}

func (mc *MetricsCollector) GetConsumerMetrics() ConsumerMetrics {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()
	
	metrics := mc.consumerMetrics
	if metrics.ProcessingCount > 0 {
		metrics.AvgProcessingLatency = time.Duration(int64(metrics.TotalProcessingLatency) / metrics.ProcessingCount)
	}
	
	return metrics
}

// Structured logging implementation
type StructuredLogger struct {
	logs []LogEntry
	mutex sync.RWMutex
}

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
	Fields    map[string]interface{}
}

func NewStructuredLogger() *StructuredLogger {
	return &StructuredLogger{
		logs: make([]LogEntry, 0),
	}
}

func (sl *StructuredLogger) LogProducerEvent(event string, metrics ProducerMetrics) {
	fields := map[string]interface{}{
		"event":           event,
		"messages_sent":   metrics.MessagesSent,
		"bytes_sent":      metrics.BytesSent,
		"send_latency_ms": metrics.AvgSendLatency.Milliseconds(),
		"error_rate":      metrics.ErrorRate,
	}
	
	sl.addLog("INFO", fmt.Sprintf("Producer %s", event), fields)
}

func (sl *StructuredLogger) LogConsumerEvent(event string, metrics ConsumerMetrics) {
	fields := map[string]interface{}{
		"event":                event,
		"messages_consumed":    metrics.MessagesConsumed,
		"bytes_consumed":       metrics.BytesConsumed,
		"processing_latency_ms": metrics.AvgProcessingLatency.Milliseconds(),
		"consumer_lag":        metrics.ConsumerLag,
	}
	
	sl.addLog("INFO", fmt.Sprintf("Consumer %s", event), fields)
}

func (sl *StructuredLogger) LogError(event string, message string, fields map[string]interface{}) {
	fields["event"] = event
	sl.addLog("ERROR", message, fields)
}

func (sl *StructuredLogger) addLog(level, message string, fields map[string]interface{}) {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()
	
	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Fields:    fields,
	}
	
	sl.logs = append(sl.logs, entry)
	
	// Keep only last 1000 logs
	if len(sl.logs) > 1000 {
		sl.logs = sl.logs[1:]
	}
	
	// Print structured log
	fmt.Printf("📝 [%s] %s: %s\n", level, entry.Timestamp.Format("15:04:05"), message)
	for key, value := range fields {
		fmt.Printf("     %s: %v\n", key, value)
	}
}

// Alerting system implementation
type AlertManager struct {
	rules []AlertRule
	mutex sync.RWMutex
}

type AlertRule struct {
	Name      string
	Severity  string
	Message   string
	Condition func(metrics interface{}) bool
	LastFired time.Time
	Cooldown  time.Duration
}

type Alert struct {
	Rule      string
	Severity  string
	Message   string
	Timestamp time.Time
	Metrics   interface{}
}

func NewAlertManager() *AlertManager {
	return &AlertManager{
		rules: make([]AlertRule, 0),
	}
}

func (am *AlertManager) AddRule(rule AlertRule) {
	am.mutex.Lock()
	defer am.mutex.Unlock()
	
	rule.Cooldown = 5 * time.Minute // Default cooldown
	am.rules = append(am.rules, rule)
}

func (am *AlertManager) CheckAlerts(metrics interface{}) {
	am.mutex.RLock()
	defer am.mutex.RUnlock()
	
	for i, rule := range am.rules {
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
				am.rules[i].LastFired = time.Now()
			}
		}
	}
}

func (am *AlertManager) sendAlert(alert Alert) {
	fmt.Printf("🚨 ALERT [%s] %s: %s\n", alert.Severity, alert.Rule, alert.Message)
	fmt.Printf("   Timestamp: %s\n", alert.Timestamp.Format("2006-01-02 15:04:05"))
}

// Health check implementation
type HealthChecker struct {
	checks []HealthCheck
	mutex  sync.RWMutex
}

type HealthCheck struct {
	Name       string
	Check      func() HealthStatus
	Critical   bool
	LastCheck  time.Time
	LastStatus HealthStatus
}

type HealthStatus struct {
	Healthy   bool
	Message   string
	Timestamp time.Time
	Duration  time.Duration
}

type HealthSummary struct {
	OverallHealth   string
	CriticalIssues  int
	WarningIssues   int
	TotalChecks     int
}

func NewHealthChecker() *HealthChecker {
	return &HealthChecker{
		checks: make([]HealthCheck, 0),
	}
}

func (hc *HealthChecker) AddCheck(check HealthCheck) {
	hc.mutex.Lock()
	defer hc.mutex.Unlock()
	
	hc.checks = append(hc.checks, check)
}

func (hc *HealthChecker) RunChecks() map[string]HealthStatus {
	hc.mutex.RLock()
	defer hc.mutex.RUnlock()
	
	results := make(map[string]HealthStatus)
	
	for _, check := range hc.checks {
		start := time.Now()
		status := check.Check()
		status.Duration = time.Since(start)
		status.Timestamp = time.Now()
		
		results[check.Name] = status
	}
	
	return results
}

func (hc *HealthChecker) GenerateSummary() HealthSummary {
	hc.mutex.RLock()
	defer hc.mutex.RUnlock()
	
	summary := HealthSummary{
		TotalChecks: len(hc.checks),
	}
	
	criticalIssues := 0
	warningIssues := 0
	
	for _, check := range hc.checks {
		if !check.LastStatus.Healthy {
			if check.Critical {
				criticalIssues++
			} else {
				warningIssues++
			}
		}
	}
	
	summary.CriticalIssues = criticalIssues
	summary.WarningIssues = warningIssues
	
	if criticalIssues > 0 {
		summary.OverallHealth = "CRITICAL"
	} else if warningIssues > 0 {
		summary.OverallHealth = "WARNING"
	} else {
		summary.OverallHealth = "HEALTHY"
	}
	
	return summary
}
