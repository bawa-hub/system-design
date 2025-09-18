package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

// BenchmarkResult represents the result of a performance benchmark
type BenchmarkResult struct {
	Name         string        `json:"name"`
	Duration     time.Duration `json:"duration"`
	Operations   int64         `json:"operations"`
	Throughput   float64       `json:"throughput"`   // operations per second
	Latency      time.Duration `json:"latency"`      // average latency
	MinLatency   time.Duration `json:"min_latency"`
	MaxLatency   time.Duration `json:"max_latency"`
	P95Latency   time.Duration `json:"p95_latency"`
	P99Latency   time.Duration `json:"p99_latency"`
	Errors       int64         `json:"errors"`
	ErrorRate    float64       `json:"error_rate"`
	Bytes        int64         `json:"bytes"`
	Bandwidth    float64       `json:"bandwidth"`    // bytes per second
	CPUUsage     float64       `json:"cpu_usage"`
	MemoryUsage  uint64        `json:"memory_usage"`
	Timestamp    time.Time     `json:"timestamp"`
}

// NewBenchmarkResult creates a new benchmark result
func NewBenchmarkResult(name string) *BenchmarkResult {
	return &BenchmarkResult{
		Name:      name,
		Timestamp: time.Now(),
	}
}

// CalculateStatistics calculates performance statistics
func (br *BenchmarkResult) CalculateStatistics() {
	if br.Operations > 0 {
		br.Throughput = float64(br.Operations) / br.Duration.Seconds()
		br.ErrorRate = float64(br.Errors) / float64(br.Operations) * 100
	}
	
	if br.Bytes > 0 {
		br.Bandwidth = float64(br.Bytes) / br.Duration.Seconds()
	}
}

// BenchmarkRunner runs performance benchmarks
type BenchmarkRunner struct {
	Results []*BenchmarkResult
	mutex   sync.RWMutex
}

// NewBenchmarkRunner creates a new benchmark runner
func NewBenchmarkRunner() *BenchmarkRunner {
	return &BenchmarkRunner{
		Results: make([]*BenchmarkResult, 0),
	}
}

// RunBenchmark runs a benchmark with the given configuration
func (br *BenchmarkRunner) RunBenchmark(name string, duration time.Duration, concurrency int, fn func() error) *BenchmarkResult {
	result := NewBenchmarkResult(name)
	
	// Start monitoring
	startTime := time.Now()
	startCPU := br.getCPUUsage()
	startMemory := br.getMemoryUsage()
	
	// Run benchmark
	br.runConcurrent(duration, concurrency, fn, result)
	
	// Calculate final statistics
	endTime := time.Now()
	endCPU := br.getCPUUsage()
	endMemory := br.getMemoryUsage()
	
	result.Duration = endTime.Sub(startTime)
	result.CPUUsage = endCPU - startCPU
	result.MemoryUsage = endMemory - startMemory
	result.CalculateStatistics()
	
	// Store result
	br.mutex.Lock()
	br.Results = append(br.Results, result)
	br.mutex.Unlock()
	
	return result
}

// runConcurrent runs the benchmark function concurrently
func (br *BenchmarkRunner) runConcurrent(duration time.Duration, concurrency int, fn func() error, result *BenchmarkResult) {
	var wg sync.WaitGroup
	var operations int64
	var errors int64
	var bytes int64
	var latencies []time.Duration
	
	// Channel to collect latencies
	latencyChan := make(chan time.Duration, 1000)
	
	// Start latency collector
	go func() {
		for latency := range latencyChan {
			latencies = append(latencies, latency)
		}
	}()
	
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	
	// Start workers
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			
			for {
				select {
				case <-ctx.Done():
					return
				default:
					start := time.Now()
					err := fn()
					latency := time.Since(start)
					
					atomic.AddInt64(&operations, 1)
					latencyChan <- latency
					
					if err != nil {
						atomic.AddInt64(&errors, 1)
					}
				}
			}
		}()
	}
	
	// Wait for completion
	wg.Wait()
	close(latencyChan)
	
	// Wait for latency collection to complete
	time.Sleep(100 * time.Millisecond)
	
	// Update result
	result.Operations = operations
	result.Errors = errors
	
	// Calculate latency statistics
	if len(latencies) > 0 {
		br.calculateLatencyStats(latencies, result)
	}
}

// calculateLatencyStats calculates latency statistics
func (br *BenchmarkRunner) calculateLatencyStats(latencies []time.Duration, result *BenchmarkResult) {
	if len(latencies) == 0 {
		return
	}
	
	// Sort latencies
	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i] < latencies[j]
	})
	
	// Calculate statistics
	total := time.Duration(0)
	for _, latency := range latencies {
		total += latency
	}
	
	result.Latency = total / time.Duration(len(latencies))
	result.MinLatency = latencies[0]
	result.MaxLatency = latencies[len(latencies)-1]
	
	// Calculate percentiles
	p95Index := int(float64(len(latencies)) * 0.95)
	p99Index := int(float64(len(latencies)) * 0.99)
	
	if p95Index < len(latencies) {
		result.P95Latency = latencies[p95Index]
	}
	if p99Index < len(latencies) {
		result.P99Latency = latencies[p99Index]
	}
}

// getCPUUsage gets current CPU usage
func (br *BenchmarkRunner) getCPUUsage() float64 {
	// Simplified CPU usage calculation
	// In a real implementation, you would use system calls
	return float64(runtime.NumGoroutine()) * 10.0
}

// getMemoryUsage gets current memory usage
func (br *BenchmarkRunner) getMemoryUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

// GetResults returns all benchmark results
func (br *BenchmarkRunner) GetResults() []*BenchmarkResult {
	br.mutex.RLock()
	defer br.mutex.RUnlock()
	
	results := make([]*BenchmarkResult, len(br.Results))
	copy(results, br.Results)
	return results
}

// NetworkBenchmark provides network-specific benchmarks
type NetworkBenchmark struct {
	Runner *BenchmarkRunner
}

// NewNetworkBenchmark creates a new network benchmark
func NewNetworkBenchmark() *NetworkBenchmark {
	return &NetworkBenchmark{
		Runner: NewBenchmarkRunner(),
	}
}

// BenchmarkTCPConnection tests TCP connection performance
func (nb *NetworkBenchmark) BenchmarkTCPConnection(host string, port string, duration time.Duration, concurrency int) *BenchmarkResult {
	address := fmt.Sprintf("%s:%s", host, port)
	
	return nb.Runner.RunBenchmark("TCP Connection", duration, concurrency, func() error {
		conn, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err != nil {
			return err
		}
		conn.Close()
		return nil
	})
}

// BenchmarkTCPThroughput tests TCP throughput
func (nb *NetworkBenchmark) BenchmarkTCPThroughput(host string, port string, duration time.Duration, concurrency int, messageSize int) *BenchmarkResult {
	address := fmt.Sprintf("%s:%s", host, port)
	message := make([]byte, messageSize)
	
	return nb.Runner.RunBenchmark("TCP Throughput", duration, concurrency, func() error {
		conn, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err != nil {
			return err
		}
		defer conn.Close()
		
		_, err = conn.Write(message)
		if err != nil {
			return err
		}
		
		// Read response
		buffer := make([]byte, messageSize)
		_, err = conn.Read(buffer)
		return err
	})
}

// BenchmarkHTTPRequest tests HTTP request performance
func (nb *NetworkBenchmark) BenchmarkHTTPRequest(url string, duration time.Duration, concurrency int) *BenchmarkResult {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	return nb.Runner.RunBenchmark("HTTP Request", duration, concurrency, func() error {
		resp, err := client.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		
		_, err = io.Copy(io.Discard, resp.Body)
		return err
	})
}

// BenchmarkUDPPacket tests UDP packet performance
func (nb *NetworkBenchmark) BenchmarkUDPPacket(host string, port string, duration time.Duration, concurrency int, packetSize int) *BenchmarkResult {
	address := fmt.Sprintf("%s:%s", host, port)
	packet := make([]byte, packetSize)
	
	return nb.Runner.RunBenchmark("UDP Packet", duration, concurrency, func() error {
		conn, err := net.Dial("udp", address)
		if err != nil {
			return err
		}
		defer conn.Close()
		
		_, err = conn.Write(packet)
		return err
	})
}

// BenchmarkPing tests ping performance
func (nb *NetworkBenchmark) BenchmarkPing(host string, duration time.Duration, concurrency int) *BenchmarkResult {
	return nb.Runner.RunBenchmark("Ping", duration, concurrency, func() error {
		conn, err := net.DialTimeout("tcp", host+":80", 5*time.Second)
		if err != nil {
			return err
		}
		conn.Close()
		return nil
	})
}

// LoadTester provides load testing functionality
type LoadTester struct {
	Targets    map[string]*LoadTarget
	Results    map[string]*LoadTestResult
	mutex      sync.RWMutex
}

// LoadTarget represents a load test target
type LoadTarget struct {
	Name        string
	URL         string
	Method      string
	Headers     map[string]string
	Body        []byte
	Concurrency int
	Duration    time.Duration
	Rate        int // requests per second
}

// LoadTestResult represents the result of a load test
type LoadTestResult struct {
	Target        string
	TotalRequests int64
	SuccessfulRequests int64
	FailedRequests int64
	AverageLatency time.Duration
	MinLatency     time.Duration
	MaxLatency     time.Duration
	P95Latency     time.Duration
	P99Latency     time.Duration
	Throughput     float64
	ErrorRate      float64
	StartTime      time.Time
	EndTime        time.Time
	Duration       time.Duration
}

// NewLoadTester creates a new load tester
func NewLoadTester() *LoadTester {
	return &LoadTester{
		Targets: make(map[string]*LoadTarget),
		Results: make(map[string]*LoadTestResult),
	}
}

// AddTarget adds a load test target
func (lt *LoadTester) AddTarget(target *LoadTarget) {
	lt.mutex.Lock()
	defer lt.mutex.Unlock()
	
	lt.Targets[target.Name] = target
}

// RunLoadTest runs a load test for a target
func (lt *LoadTester) RunLoadTest(targetName string) (*LoadTestResult, error) {
	lt.mutex.RLock()
	target, exists := lt.Targets[targetName]
	lt.mutex.RUnlock()
	
	if !exists {
		return nil, fmt.Errorf("target not found: %s", targetName)
	}
	
	result := &LoadTestResult{
		Target:   targetName,
		StartTime: time.Now(),
	}
	
	// Run load test
	lt.runLoadTest(target, result)
	
	// Calculate final statistics
	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.Throughput = float64(result.TotalRequests) / result.Duration.Seconds()
	result.ErrorRate = float64(result.FailedRequests) / float64(result.TotalRequests) * 100
	
	// Store result
	lt.mutex.Lock()
	lt.Results[targetName] = result
	lt.mutex.Unlock()
	
	return result, nil
}

// runLoadTest runs the actual load test
func (lt *LoadTester) runLoadTest(target *LoadTarget, result *LoadTestResult) {
	var wg sync.WaitGroup
	var totalRequests int64
	var successfulRequests int64
	var failedRequests int64
	var latencies []time.Duration
	
	// Channel to collect latencies
	latencyChan := make(chan time.Duration, 1000)
	
	// Start latency collector
	go func() {
		for latency := range latencyChan {
			latencies = append(latencies, latency)
		}
	}()
	
	// Create HTTP client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), target.Duration)
	defer cancel()
	
	// Start workers
	for i := 0; i < target.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			
			for {
				select {
				case <-ctx.Done():
					return
				default:
					start := time.Now()
					
					// Create request
					req, err := http.NewRequest(target.Method, target.URL, nil)
					if err != nil {
						atomic.AddInt64(&failedRequests, 1)
						continue
					}
					
					// Add headers
					for key, value := range target.Headers {
						req.Header.Set(key, value)
					}
					
					// Send request
					resp, err := client.Do(req)
					if err != nil {
						atomic.AddInt64(&failedRequests, 1)
					} else {
						atomic.AddInt64(&successfulRequests, 1)
						resp.Body.Close()
					}
					
					latency := time.Since(start)
					latencyChan <- latency
					atomic.AddInt64(&totalRequests, 1)
					
					// Rate limiting
					if target.Rate > 0 {
						time.Sleep(time.Second / time.Duration(target.Rate))
					}
				}
			}
		}()
	}
	
	// Wait for completion
	wg.Wait()
	close(latencyChan)
	
	// Wait for latency collection to complete
	time.Sleep(100 * time.Millisecond)
	
	// Update result
	result.TotalRequests = totalRequests
	result.SuccessfulRequests = successfulRequests
	result.FailedRequests = failedRequests
	
	// Calculate latency statistics
	if len(latencies) > 0 {
		lt.calculateLatencyStats(latencies, result)
	}
}

// calculateLatencyStats calculates latency statistics
func (lt *LoadTester) calculateLatencyStats(latencies []time.Duration, result *LoadTestResult) {
	if len(latencies) == 0 {
		return
	}
	
	// Sort latencies
	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i] < latencies[j]
	})
	
	// Calculate statistics
	total := time.Duration(0)
	for _, latency := range latencies {
		total += latency
	}
	
	result.AverageLatency = total / time.Duration(len(latencies))
	result.MinLatency = latencies[0]
	result.MaxLatency = latencies[len(latencies)-1]
	
	// Calculate percentiles
	p95Index := int(float64(len(latencies)) * 0.95)
	p99Index := int(float64(len(latencies)) * 0.99)
	
	if p95Index < len(latencies) {
		result.P95Latency = latencies[p95Index]
	}
	if p99Index < len(latencies) {
		result.P99Latency = latencies[p99Index]
	}
}

// GetResults returns all load test results
func (lt *LoadTester) GetResults() map[string]*LoadTestResult {
	lt.mutex.RLock()
	defer lt.mutex.RUnlock()
	
	results := make(map[string]*LoadTestResult)
	for name, result := range lt.Results {
		results[name] = result
	}
	return results
}

// PerformanceProfiler provides performance profiling functionality
type PerformanceProfiler struct {
	StartTime time.Time
	EndTime   time.Time
	Metrics   map[string]interface{}
	mutex     sync.RWMutex
}

// NewPerformanceProfiler creates a new performance profiler
func NewPerformanceProfiler() *PerformanceProfiler {
	return &PerformanceProfiler{
		Metrics: make(map[string]interface{}),
	}
}

// Start starts the profiler
func (pp *PerformanceProfiler) Start() {
	pp.mutex.Lock()
	defer pp.mutex.Unlock()
	
	pp.StartTime = time.Now()
	pp.Metrics = make(map[string]interface{})
}

// Stop stops the profiler
func (pp *PerformanceProfiler) Stop() {
	pp.mutex.Lock()
	defer pp.mutex.Unlock()
	
	pp.EndTime = time.Now()
}

// AddMetric adds a metric to the profiler
func (pp *PerformanceProfiler) AddMetric(name string, value interface{}) {
	pp.mutex.Lock()
	defer pp.mutex.Unlock()
	
	pp.Metrics[name] = value
}

// GetMetrics returns all profiler metrics
func (pp *PerformanceProfiler) GetMetrics() map[string]interface{} {
	pp.mutex.RLock()
	defer pp.mutex.RUnlock()
	
	metrics := make(map[string]interface{})
	for name, value := range pp.Metrics {
		metrics[name] = value
	}
	
	// Add duration
	if !pp.StartTime.IsZero() && !pp.EndTime.IsZero() {
		metrics["duration"] = pp.EndTime.Sub(pp.StartTime)
	}
	
	return metrics
}

// Demonstrate performance testing
func demonstratePerformanceTesting() {
	fmt.Println("=== Performance Testing Demo ===\n")
	
	// Create network benchmark
	benchmark := NewNetworkBenchmark()
	
	// Run TCP connection benchmark
	fmt.Println("Running TCP Connection Benchmark...")
	result := benchmark.BenchmarkTCPConnection("localhost", "8080", 5*time.Second, 10)
	
	fmt.Printf("TCP Connection Results:\n")
	fmt.Printf("  Operations: %d\n", result.Operations)
	fmt.Printf("  Duration: %v\n", result.Duration)
	fmt.Printf("  Throughput: %.2f ops/sec\n", result.Throughput)
	fmt.Printf("  Average Latency: %v\n", result.Latency)
	fmt.Printf("  Min Latency: %v\n", result.MinLatency)
	fmt.Printf("  Max Latency: %v\n", result.MaxLatency)
	fmt.Printf("  P95 Latency: %v\n", result.P95Latency)
	fmt.Printf("  P99 Latency: %v\n", result.P99Latency)
	fmt.Printf("  Errors: %d\n", result.Errors)
	fmt.Printf("  Error Rate: %.2f%%\n", result.ErrorRate)
	fmt.Printf("  CPU Usage: %.2f%%\n", result.CPUUsage)
	fmt.Printf("  Memory Usage: %d bytes\n", result.MemoryUsage)
}

// Demonstrate load testing
func demonstrateLoadTesting() {
	fmt.Println("=== Load Testing Demo ===\n")
	
	// Create load tester
	loadTester := NewLoadTester()
	
	// Add load test target
	target := &LoadTarget{
		Name:        "HTTP Server",
		URL:         "http://localhost:8080",
		Method:      "GET",
		Headers:     map[string]string{"User-Agent": "LoadTester/1.0"},
		Concurrency: 10,
		Duration:    10 * time.Second,
		Rate:        100, // 100 requests per second
	}
	
	loadTester.AddTarget(target)
	
	// Run load test
	fmt.Println("Running Load Test...")
	result, err := loadTester.RunLoadTest("HTTP Server")
	if err != nil {
		fmt.Printf("Error running load test: %v\n", err)
		return
	}
	
	fmt.Printf("Load Test Results:\n")
	fmt.Printf("  Target: %s\n", result.Target)
	fmt.Printf("  Total Requests: %d\n", result.TotalRequests)
	fmt.Printf("  Successful Requests: %d\n", result.SuccessfulRequests)
	fmt.Printf("  Failed Requests: %d\n", result.FailedRequests)
	fmt.Printf("  Duration: %v\n", result.Duration)
	fmt.Printf("  Throughput: %.2f req/sec\n", result.Throughput)
	fmt.Printf("  Average Latency: %v\n", result.AverageLatency)
	fmt.Printf("  Min Latency: %v\n", result.MinLatency)
	fmt.Printf("  Max Latency: %v\n", result.MaxLatency)
	fmt.Printf("  P95 Latency: %v\n", result.P95Latency)
	fmt.Printf("  P99 Latency: %v\n", result.P99Latency)
	fmt.Printf("  Error Rate: %.2f%%\n", result.ErrorRate)
}

// Demonstrate performance profiling
func demonstratePerformanceProfiling() {
	fmt.Println("=== Performance Profiling Demo ===\n")
	
	// Create profiler
	profiler := NewPerformanceProfiler()
	
	// Start profiling
	profiler.Start()
	
	// Simulate some work
	time.Sleep(2 * time.Second)
	
	// Add some metrics
	profiler.AddMetric("goroutines", runtime.NumGoroutine())
	profiler.AddMetric("memory_alloc", getMemoryAlloc())
	profiler.AddMetric("gc_runs", getGCRuns())
	
	// Stop profiling
	profiler.Stop()
	
	// Get metrics
	metrics := profiler.GetMetrics()
	
	fmt.Printf("Profiler Results:\n")
	for name, value := range metrics {
		fmt.Printf("  %s: %v\n", name, value)
	}
}

// getMemoryAlloc gets current memory allocation
func getMemoryAlloc() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

// getGCRuns gets number of GC runs
func getGCRuns() uint32 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.NumGC
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 10: Performance Testing")
	fmt.Println("================================================================\n")
	
	// Run all demonstrations
	demonstratePerformanceTesting()
	fmt.Println()
	demonstrateLoadTesting()
	fmt.Println()
	demonstratePerformanceProfiling()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Performance testing is essential for network optimization")
	fmt.Println("2. Benchmarking helps identify performance bottlenecks")
	fmt.Println("3. Load testing simulates real-world traffic patterns")
	fmt.Println("4. Profiling provides detailed performance insights")
	fmt.Println("5. Metrics collection enables data-driven optimization")
	fmt.Println("6. Go provides excellent tools for performance testing")
	fmt.Println("7. Understanding performance characteristics is crucial")
	fmt.Println("8. Continuous monitoring ensures optimal performance")
	
	fmt.Println("\n📚 Next Topic: Distributed Systems Networking")
}
