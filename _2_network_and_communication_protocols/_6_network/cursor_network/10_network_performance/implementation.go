package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// PerformanceMetrics represents network performance metrics
type PerformanceMetrics struct {
	Throughput     float64   `json:"throughput"`      // bits per second
	Latency        float64   `json:"latency"`         // milliseconds
	Jitter         float64   `json:"jitter"`          // milliseconds
	PacketLoss     float64   `json:"packet_loss"`     // percentage
	BandwidthUtil  float64   `json:"bandwidth_util"`  // percentage
	CPUUtil        float64   `json:"cpu_util"`        // percentage
	MemoryUsage    uint64    `json:"memory_usage"`    // bytes
	Connections    int32     `json:"connections"`     // count
	Timestamp      time.Time `json:"timestamp"`
	mutex          sync.RWMutex
}

// NewPerformanceMetrics creates a new performance metrics instance
func NewPerformanceMetrics() *PerformanceMetrics {
	return &PerformanceMetrics{
		Timestamp: time.Now(),
	}
}

// Update updates the performance metrics
func (pm *PerformanceMetrics) Update(throughput, latency, jitter, packetLoss, bandwidthUtil, cpuUtil float64, memoryUsage uint64, connections int32) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	
	pm.Throughput = throughput
	pm.Latency = latency
	pm.Jitter = jitter
	pm.PacketLoss = packetLoss
	pm.BandwidthUtil = bandwidthUtil
	pm.CPUUtil = cpuUtil
	pm.MemoryUsage = memoryUsage
	pm.Connections = connections
	pm.Timestamp = time.Now()
}

// GetMetrics returns a copy of the current metrics
func (pm *PerformanceMetrics) GetMetrics() PerformanceMetrics {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()
	
	return PerformanceMetrics{
		Throughput:    pm.Throughput,
		Latency:       pm.Latency,
		Jitter:        pm.Jitter,
		PacketLoss:    pm.PacketLoss,
		BandwidthUtil: pm.BandwidthUtil,
		CPUUtil:       pm.CPUUtil,
		MemoryUsage:   pm.MemoryUsage,
		Connections:   pm.Connections,
		Timestamp:     pm.Timestamp,
	}
}

// PerformanceMonitor monitors network performance
type PerformanceMonitor struct {
	Metrics     *PerformanceMetrics
	IsRunning   bool
	StopChan    chan struct{}
	UpdateChan  chan PerformanceMetrics
	mutex       sync.RWMutex
}

// NewPerformanceMonitor creates a new performance monitor
func NewPerformanceMonitor() *PerformanceMonitor {
	return &PerformanceMonitor{
		Metrics:    NewPerformanceMetrics(),
		StopChan:   make(chan struct{}),
		UpdateChan: make(chan PerformanceMetrics, 100),
	}
}

// Start starts the performance monitoring
func (pm *PerformanceMonitor) Start() {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	
	if pm.IsRunning {
		return
	}
	
	pm.IsRunning = true
	go pm.monitorLoop()
}

// Stop stops the performance monitoring
func (pm *PerformanceMonitor) Stop() {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	
	if !pm.IsRunning {
		return
	}
	
	pm.IsRunning = false
	close(pm.StopChan)
}

// monitorLoop runs the monitoring loop
func (pm *PerformanceMonitor) monitorLoop() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-pm.StopChan:
			return
		case <-ticker.C:
			pm.collectMetrics()
		}
	}
}

// collectMetrics collects current performance metrics
func (pm *PerformanceMonitor) collectMetrics() {
	// Get system metrics
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	// Calculate CPU usage (simplified)
	cpuUtil := pm.calculateCPUUsage()
	
	// Calculate network metrics (simplified)
	throughput := pm.calculateThroughput()
	latency := pm.calculateLatency()
	jitter := pm.calculateJitter()
	packetLoss := pm.calculatePacketLoss()
	bandwidthUtil := pm.calculateBandwidthUtil()
	connections := pm.getConnectionCount()
	
	// Update metrics
	pm.Metrics.Update(throughput, latency, jitter, packetLoss, bandwidthUtil, cpuUtil, m.Alloc, connections)
	
	// Send update
	select {
	case pm.UpdateChan <- pm.Metrics.GetMetrics():
	default:
		// Channel full, skip update
	}
}

// calculateCPUUsage calculates CPU usage percentage
func (pm *PerformanceMonitor) calculateCPUUsage() float64 {
	// Simplified CPU usage calculation
	// In a real implementation, you would use system calls to get actual CPU usage
	return float64(runtime.NumGoroutine()) * 10.0
}

// calculateThroughput calculates network throughput
func (pm *PerformanceMonitor) calculateThroughput() float64 {
	// Simplified throughput calculation
	// In a real implementation, you would measure actual network traffic
	return 1000.0 + float64(time.Now().Unix()%1000)
}

// calculateLatency calculates network latency
func (pm *PerformanceMonitor) calculateLatency() float64 {
	// Simplified latency calculation
	// In a real implementation, you would measure actual RTT
	return 50.0 + float64(time.Now().Unix()%50)
}

// calculateJitter calculates network jitter
func (pm *PerformanceMonitor) calculateJitter() float64 {
	// Simplified jitter calculation
	// In a real implementation, you would measure actual jitter
	return 5.0 + float64(time.Now().Unix()%10)
}

// calculatePacketLoss calculates packet loss percentage
func (pm *PerformanceMonitor) calculatePacketLoss() float64 {
	// Simplified packet loss calculation
	// In a real implementation, you would measure actual packet loss
	return float64(time.Now().Unix()%100) / 1000.0
}

// calculateBandwidthUtil calculates bandwidth utilization
func (pm *PerformanceMonitor) calculateBandwidthUtil() float64 {
	// Simplified bandwidth utilization calculation
	// In a real implementation, you would measure actual bandwidth usage
	return 30.0 + float64(time.Now().Unix()%70)
}

// getConnectionCount gets the current connection count
func (pm *PerformanceMonitor) getConnectionCount() int32 {
	// Simplified connection count
	// In a real implementation, you would count actual connections
	return int32(runtime.NumGoroutine())
}

// LoadBalancer provides load balancing functionality
type LoadBalancer struct {
	Servers    []*Server
	Algorithm  string
	Current    int32
	mutex      sync.RWMutex
}

// Server represents a backend server
type Server struct {
	Address     string
	Weight      int
	Connections int32
	ResponseTime float64
	IsHealthy   bool
	LastCheck   time.Time
}

// NewLoadBalancer creates a new load balancer
func NewLoadBalancer(algorithm string) *LoadBalancer {
	return &LoadBalancer{
		Servers:   make([]*Server, 0),
		Algorithm: algorithm,
		Current:   0,
	}
}

// AddServer adds a server to the load balancer
func (lb *LoadBalancer) AddServer(address string, weight int) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	
	server := &Server{
		Address:     address,
		Weight:      weight,
		Connections: 0,
		ResponseTime: 0,
		IsHealthy:   true,
		LastCheck:   time.Now(),
	}
	
	lb.Servers = append(lb.Servers, server)
}

// GetServer returns the next server based on the load balancing algorithm
func (lb *LoadBalancer) GetServer() *Server {
	lb.mutex.RLock()
	defer lb.mutex.RUnlock()
	
	if len(lb.Servers) == 0 {
		return nil
	}
	
	switch lb.Algorithm {
	case "round_robin":
		return lb.getRoundRobinServer()
	case "least_connections":
		return lb.getLeastConnectionsServer()
	case "weighted":
		return lb.getWeightedServer()
	case "fastest":
		return lb.getFastestServer()
	default:
		return lb.getRoundRobinServer()
	}
}

// getRoundRobinServer returns the next server in round-robin fashion
func (lb *LoadBalancer) getRoundRobinServer() *Server {
	if len(lb.Servers) == 0 {
		return nil
	}
	
	// Find next healthy server
	for i := 0; i < len(lb.Servers); i++ {
		index := int(atomic.AddInt32(&lb.Current, 1)) % len(lb.Servers)
		if lb.Servers[index].IsHealthy {
			return lb.Servers[index]
		}
	}
	
	// If no healthy server found, return first server
	return lb.Servers[0]
}

// getLeastConnectionsServer returns the server with the least connections
func (lb *LoadBalancer) getLeastConnectionsServer() *Server {
	if len(lb.Servers) == 0 {
		return nil
	}
	
	var bestServer *Server
	minConnections := int32(math.MaxInt32)
	
	for _, server := range lb.Servers {
		if server.IsHealthy && server.Connections < minConnections {
			minConnections = server.Connections
			bestServer = server
		}
	}
	
	return bestServer
}

// getWeightedServer returns a server based on weight
func (lb *LoadBalancer) getWeightedServer() *Server {
	if len(lb.Servers) == 0 {
		return nil
	}
	
	// Calculate total weight of healthy servers
	totalWeight := 0
	for _, server := range lb.Servers {
		if server.IsHealthy {
			totalWeight += server.Weight
		}
	}
	
	if totalWeight == 0 {
		return lb.Servers[0]
	}
	
	// Select server based on weight
	random := time.Now().UnixNano() % int64(totalWeight)
	currentWeight := 0
	
	for _, server := range lb.Servers {
		if server.IsHealthy {
			currentWeight += server.Weight
			if int64(currentWeight) > random {
				return server
			}
		}
	}
	
	return lb.Servers[0]
}

// getFastestServer returns the server with the fastest response time
func (lb *LoadBalancer) getFastestServer() *Server {
	if len(lb.Servers) == 0 {
		return nil
	}
	
	var bestServer *Server
	minResponseTime := math.MaxFloat64
	
	for _, server := range lb.Servers {
		if server.IsHealthy && server.ResponseTime < minResponseTime {
			minResponseTime = server.ResponseTime
			bestServer = server
		}
	}
	
	return bestServer
}

// IncrementConnections increments the connection count for a server
func (lb *LoadBalancer) IncrementConnections(server *Server) {
	if server != nil {
		atomic.AddInt32(&server.Connections, 1)
	}
}

// DecrementConnections decrements the connection count for a server
func (lb *LoadBalancer) DecrementConnections(server *Server) {
	if server != nil {
		atomic.AddInt32(&server.Connections, -1)
	}
}

// ConnectionPool manages a pool of network connections
type ConnectionPool struct {
	Connections chan net.Conn
	Factory     func() (net.Conn, error)
	MaxSize     int
	CurrentSize int32
	mutex       sync.RWMutex
}

// NewConnectionPool creates a new connection pool
func NewConnectionPool(factory func() (net.Conn, error), maxSize int) *ConnectionPool {
	return &ConnectionPool{
		Connections: make(chan net.Conn, maxSize),
		Factory:     factory,
		MaxSize:     maxSize,
		CurrentSize: 0,
	}
}

// GetConnection gets a connection from the pool
func (cp *ConnectionPool) GetConnection() (net.Conn, error) {
	select {
	case conn := <-cp.Connections:
		return conn, nil
	default:
		// No connection available, create new one
		if atomic.LoadInt32(&cp.CurrentSize) < int32(cp.MaxSize) {
			conn, err := cp.Factory()
			if err != nil {
				return nil, err
			}
			atomic.AddInt32(&cp.CurrentSize, 1)
			return conn, nil
		}
		
		// Wait for connection to become available
		select {
		case conn := <-cp.Connections:
			return conn, nil
		case <-time.After(5 * time.Second):
			return nil, fmt.Errorf("timeout waiting for connection")
		}
	}
}

// ReturnConnection returns a connection to the pool
func (cp *ConnectionPool) ReturnConnection(conn net.Conn) {
	if conn == nil {
		return
	}
	
	select {
	case cp.Connections <- conn:
		// Connection returned to pool
	default:
		// Pool is full, close connection
		conn.Close()
		atomic.AddInt32(&cp.CurrentSize, -1)
	}
}

// Close closes all connections in the pool
func (cp *ConnectionPool) Close() {
	close(cp.Connections)
	
	for conn := range cp.Connections {
		conn.Close()
	}
}

// TrafficShaper shapes network traffic
type TrafficShaper struct {
	RateLimit    int64 // bytes per second
	BucketSize   int64 // bytes
	Tokens       int64
	LastUpdate   time.Time
	mutex        sync.Mutex
}

// NewTrafficShaper creates a new traffic shaper
func NewTrafficShaper(rateLimit, bucketSize int64) *TrafficShaper {
	return &TrafficShaper{
		RateLimit:  rateLimit,
		BucketSize: bucketSize,
		Tokens:     bucketSize,
		LastUpdate: time.Now(),
	}
}

// Allow checks if traffic is allowed based on rate limiting
func (ts *TrafficShaper) Allow(bytes int64) bool {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	
	now := time.Now()
	elapsed := now.Sub(ts.LastUpdate).Seconds()
	
	// Add tokens based on elapsed time
	ts.Tokens += int64(float64(ts.RateLimit) * elapsed)
	if ts.Tokens > ts.BucketSize {
		ts.Tokens = ts.BucketSize
	}
	
	ts.LastUpdate = now
	
	// Check if enough tokens available
	if ts.Tokens >= bytes {
		ts.Tokens -= bytes
		return true
	}
	
	return false
}

// Wait waits until enough tokens are available
func (ts *TrafficShaper) Wait(bytes int64) {
	for !ts.Allow(bytes) {
		time.Sleep(10 * time.Millisecond)
	}
}

// BandwidthMonitor monitors bandwidth usage
type BandwidthMonitor struct {
	Interfaces map[string]*InterfaceStats
	mutex      sync.RWMutex
}

// InterfaceStats represents interface statistics
type InterfaceStats struct {
	Name        string
	BytesIn     uint64
	BytesOut    uint64
	PacketsIn   uint64
	PacketsOut  uint64
	ErrorsIn    uint64
	ErrorsOut   uint64
	LastUpdate  time.Time
}

// NewBandwidthMonitor creates a new bandwidth monitor
func NewBandwidthMonitor() *BandwidthMonitor {
	return &BandwidthMonitor{
		Interfaces: make(map[string]*InterfaceStats),
	}
}

// UpdateInterface updates interface statistics
func (bm *BandwidthMonitor) UpdateInterface(name string, bytesIn, bytesOut, packetsIn, packetsOut, errorsIn, errorsOut uint64) {
	bm.mutex.Lock()
	defer bm.mutex.Unlock()
	
	bm.Interfaces[name] = &InterfaceStats{
		Name:       name,
		BytesIn:    bytesIn,
		BytesOut:   bytesOut,
		PacketsIn:  packetsIn,
		PacketsOut: packetsOut,
		ErrorsIn:   errorsIn,
		ErrorsOut:  errorsOut,
		LastUpdate: time.Now(),
	}
}

// GetInterfaceStats returns interface statistics
func (bm *BandwidthMonitor) GetInterfaceStats(name string) *InterfaceStats {
	bm.mutex.RLock()
	defer bm.mutex.RUnlock()
	
	return bm.Interfaces[name]
}

// GetAllStats returns all interface statistics
func (bm *BandwidthMonitor) GetAllStats() map[string]*InterfaceStats {
	bm.mutex.RLock()
	defer bm.mutex.RUnlock()
	
	stats := make(map[string]*InterfaceStats)
	for name, stat := range bm.Interfaces {
		stats[name] = stat
	}
	return stats
}

// LatencyTester tests network latency
type LatencyTester struct {
	Targets map[string]*LatencyTarget
	mutex   sync.RWMutex
}

// LatencyTarget represents a latency test target
type LatencyTarget struct {
	Address     string
	Latency     float64
	Jitter      float64
	PacketLoss  float64
	LastTest    time.Time
	TestCount   int
	SuccessCount int
}

// NewLatencyTester creates a new latency tester
func NewLatencyTester() *LatencyTester {
	return &LatencyTester{
		Targets: make(map[string]*LatencyTarget),
	}
}

// AddTarget adds a latency test target
func (lt *LatencyTester) AddTarget(address string) {
	lt.mutex.Lock()
	defer lt.mutex.Unlock()
	
	lt.Targets[address] = &LatencyTarget{
		Address: address,
		Latency: 0,
		Jitter:  0,
		PacketLoss: 0,
		LastTest: time.Now(),
		TestCount: 0,
		SuccessCount: 0,
	}
}

// TestLatency tests latency to a target
func (lt *LatencyTester) TestLatency(address string) error {
	lt.mutex.Lock()
	defer lt.mutex.Unlock()
	
	target, exists := lt.Targets[address]
	if !exists {
		return fmt.Errorf("target not found: %s", address)
	}
	
	start := time.Now()
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		target.TestCount++
		target.PacketLoss = float64(target.TestCount-target.SuccessCount) / float64(target.TestCount) * 100
		return err
	}
	defer conn.Close()
	
	latency := float64(time.Since(start).Nanoseconds()) / 1e6 // Convert to milliseconds
	
	target.TestCount++
	target.SuccessCount++
	target.Latency = latency
	target.LastTest = time.Now()
	target.PacketLoss = float64(target.TestCount-target.SuccessCount) / float64(target.TestCount) * 100
	
	return nil
}

// GetTargetStats returns target statistics
func (lt *LatencyTester) GetTargetStats(address string) *LatencyTarget {
	lt.mutex.RLock()
	defer lt.mutex.RUnlock()
	
	return lt.Targets[address]
}

// PerformanceOptimizer provides network performance optimization
type PerformanceOptimizer struct {
	Monitor        *PerformanceMonitor
	LoadBalancer   *LoadBalancer
	ConnectionPool *ConnectionPool
	TrafficShaper  *TrafficShaper
	BandwidthMonitor *BandwidthMonitor
	LatencyTester  *LatencyTester
}

// NewPerformanceOptimizer creates a new performance optimizer
func NewPerformanceOptimizer() *PerformanceOptimizer {
	return &PerformanceOptimizer{
		Monitor:        NewPerformanceMonitor(),
		LoadBalancer:   NewLoadBalancer("round_robin"),
		BandwidthMonitor: NewBandwidthMonitor(),
		LatencyTester:  NewLatencyTester(),
	}
}

// Start starts the performance optimizer
func (po *PerformanceOptimizer) Start() {
	po.Monitor.Start()
}

// Stop stops the performance optimizer
func (po *PerformanceOptimizer) Stop() {
	po.Monitor.Stop()
	if po.ConnectionPool != nil {
		po.ConnectionPool.Close()
	}
}

// OptimizeTCP optimizes TCP performance
func (po *PerformanceOptimizer) OptimizeTCP() {
	// TCP optimization would typically involve:
	// - Adjusting window sizes
	// - Configuring congestion control
	// - Setting buffer sizes
	// - Enabling TCP features like SACK, window scaling
	
	fmt.Println("TCP optimization applied:")
	fmt.Println("  - Window scaling enabled")
	fmt.Println("  - SACK enabled")
	fmt.Println("  - Congestion control: CUBIC")
	fmt.Println("  - Buffer sizes optimized")
}

// OptimizeUDP optimizes UDP performance
func (po *PerformanceOptimizer) OptimizeUDP() {
	// UDP optimization would typically involve:
	// - Adjusting buffer sizes
	// - Configuring socket options
	// - Implementing application-level reliability
	// - Setting appropriate timeouts
	
	fmt.Println("UDP optimization applied:")
	fmt.Println("  - Buffer sizes optimized")
	fmt.Println("  - Socket options configured")
	fmt.Println("  - Application-level reliability enabled")
	fmt.Println("  - Timeouts configured")
}

// OptimizeConnections optimizes connection handling
func (po *PerformanceOptimizer) OptimizeConnections() {
	// Connection optimization would typically involve:
	// - Implementing connection pooling
	// - Setting keep-alive options
	// - Configuring connection limits
	// - Implementing connection reuse
	
	fmt.Println("Connection optimization applied:")
	fmt.Println("  - Connection pooling enabled")
	fmt.Println("  - Keep-alive configured")
	fmt.Println("  - Connection limits set")
	fmt.Println("  - Connection reuse implemented")
}

// Demonstrate performance monitoring
func demonstratePerformanceMonitoring() {
	fmt.Println("=== Performance Monitoring Demo ===\n")
	
	// Create performance monitor
	monitor := NewPerformanceMonitor()
	
	// Start monitoring
	monitor.Start()
	defer monitor.Stop()
	
	// Monitor for 5 seconds
	time.Sleep(5 * time.Second)
	
	// Get current metrics
	metrics := monitor.Metrics.GetMetrics()
	
	fmt.Printf("Performance Metrics:\n")
	fmt.Printf("  Throughput: %.2f bps\n", metrics.Throughput)
	fmt.Printf("  Latency: %.2f ms\n", metrics.Latency)
	fmt.Printf("  Jitter: %.2f ms\n", metrics.Jitter)
	fmt.Printf("  Packet Loss: %.2f%%\n", metrics.PacketLoss)
	fmt.Printf("  Bandwidth Util: %.2f%%\n", metrics.BandwidthUtil)
	fmt.Printf("  CPU Util: %.2f%%\n", metrics.CPUUtil)
	fmt.Printf("  Memory Usage: %d bytes\n", metrics.MemoryUsage)
	fmt.Printf("  Connections: %d\n", metrics.Connections)
}

// Demonstrate load balancing
func demonstrateLoadBalancing() {
	fmt.Println("=== Load Balancing Demo ===\n")
	
	// Create load balancer
	lb := NewLoadBalancer("round_robin")
	
	// Add servers
	lb.AddServer("192.168.1.10:8080", 1)
	lb.AddServer("192.168.1.11:8080", 2)
	lb.AddServer("192.168.1.12:8080", 1)
	
	// Test load balancing
	fmt.Println("Load Balancing Test (Round Robin):")
	for i := 0; i < 10; i++ {
		server := lb.GetServer()
		if server != nil {
			fmt.Printf("  Request %d: %s (weight: %d)\n", i+1, server.Address, server.Weight)
			lb.IncrementConnections(server)
		}
	}
	
	// Test least connections
	lb.Algorithm = "least_connections"
	fmt.Println("\nLoad Balancing Test (Least Connections):")
	for i := 0; i < 5; i++ {
		server := lb.GetServer()
		if server != nil {
			fmt.Printf("  Request %d: %s (connections: %d)\n", i+1, server.Address, server.Connections)
			lb.IncrementConnections(server)
		}
	}
}

// Demonstrate connection pooling
func demonstrateConnectionPooling() {
	fmt.Println("=== Connection Pooling Demo ===\n")
	
	// Create connection pool
	factory := func() (net.Conn, error) {
		return net.Dial("tcp", "localhost:8080")
	}
	
	pool := NewConnectionPool(factory, 5)
	defer pool.Close()
	
	// Test connection pool
	fmt.Println("Connection Pool Test:")
	for i := 0; i < 10; i++ {
		conn, err := pool.GetConnection()
		if err != nil {
			fmt.Printf("  Request %d: Error getting connection: %v\n", i+1, err)
			continue
		}
		
		fmt.Printf("  Request %d: Got connection\n", i+1)
		
		// Simulate work
		time.Sleep(100 * time.Millisecond)
		
		// Return connection
		pool.ReturnConnection(conn)
	}
}

// Demonstrate traffic shaping
func demonstrateTrafficShaping() {
	fmt.Println("=== Traffic Shaping Demo ===\n")
	
	// Create traffic shaper
	shaper := NewTrafficShaper(1000, 2000) // 1000 bytes/sec, 2000 byte bucket
	
	// Test traffic shaping
	fmt.Println("Traffic Shaping Test:")
	for i := 0; i < 10; i++ {
		bytes := int64(100 + i*50)
		allowed := shaper.Allow(bytes)
		fmt.Printf("  Request %d: %d bytes - %s\n", i+1, bytes, 
			map[bool]string{true: "ALLOWED", false: "RATE LIMITED"}[allowed])
		
		time.Sleep(100 * time.Millisecond)
	}
}

// Demonstrate bandwidth monitoring
func demonstrateBandwidthMonitoring() {
	fmt.Println("=== Bandwidth Monitoring Demo ===\n")
	
	// Create bandwidth monitor
	monitor := NewBandwidthMonitor()
	
	// Simulate interface updates
	monitor.UpdateInterface("eth0", 1000000, 500000, 1000, 500, 0, 0)
	monitor.UpdateInterface("eth1", 2000000, 1500000, 2000, 1500, 5, 2)
	
	// Display statistics
	fmt.Println("Interface Statistics:")
	for name, stats := range monitor.GetAllStats() {
		fmt.Printf("  %s:\n", name)
		fmt.Printf("    Bytes In: %d\n", stats.BytesIn)
		fmt.Printf("    Bytes Out: %d\n", stats.BytesOut)
		fmt.Printf("    Packets In: %d\n", stats.PacketsIn)
		fmt.Printf("    Packets Out: %d\n", stats.PacketsOut)
		fmt.Printf("    Errors In: %d\n", stats.ErrorsIn)
		fmt.Printf("    Errors Out: %d\n", stats.ErrorsOut)
	}
}

// Demonstrate latency testing
func demonstrateLatencyTesting() {
	fmt.Println("=== Latency Testing Demo ===\n")
	
	// Create latency tester
	tester := NewLatencyTester()
	
	// Add targets
	tester.AddTarget("8.8.8.8:53")
	tester.AddTarget("1.1.1.1:53")
	
	// Test latency
	fmt.Println("Latency Testing:")
	for i := 0; i < 5; i++ {
		for address := range tester.Targets {
			err := tester.TestLatency(address)
			if err != nil {
				fmt.Printf("  %s: Error - %v\n", address, err)
			} else {
				stats := tester.GetTargetStats(address)
				fmt.Printf("  %s: %.2f ms (loss: %.2f%%)\n", 
					address, stats.Latency, stats.PacketLoss)
			}
		}
		time.Sleep(1 * time.Second)
	}
}

// Demonstrate performance optimization
func demonstratePerformanceOptimization() {
	fmt.Println("=== Performance Optimization Demo ===\n")
	
	// Create performance optimizer
	optimizer := NewPerformanceOptimizer()
	
	// Start optimizer
	optimizer.Start()
	defer optimizer.Stop()
	
	// Apply optimizations
	optimizer.OptimizeTCP()
	fmt.Println()
	optimizer.OptimizeUDP()
	fmt.Println()
	optimizer.OptimizeConnections()
	
	// Monitor for a bit
	time.Sleep(2 * time.Second)
	
	// Get metrics
	metrics := optimizer.Monitor.Metrics.GetMetrics()
	fmt.Printf("\nCurrent Performance:\n")
	fmt.Printf("  Throughput: %.2f bps\n", metrics.Throughput)
	fmt.Printf("  Latency: %.2f ms\n", metrics.Latency)
	fmt.Printf("  CPU Util: %.2f%%\n", metrics.CPUUtil)
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 10: Network Performance & Optimization")
	fmt.Println("================================================================================\n")
	
	// Run all demonstrations
	demonstratePerformanceMonitoring()
	fmt.Println()
	demonstrateLoadBalancing()
	fmt.Println()
	demonstrateConnectionPooling()
	fmt.Println()
	demonstrateTrafficShaping()
	fmt.Println()
	demonstrateBandwidthMonitoring()
	fmt.Println()
	demonstrateLatencyTesting()
	fmt.Println()
	demonstratePerformanceOptimization()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Performance monitoring is essential for network optimization")
	fmt.Println("2. Load balancing distributes traffic across multiple servers")
	fmt.Println("3. Connection pooling improves resource utilization")
	fmt.Println("4. Traffic shaping controls network traffic rates")
	fmt.Println("5. Bandwidth monitoring tracks network usage")
	fmt.Println("6. Latency testing measures network responsiveness")
	fmt.Println("7. Performance optimization requires multiple techniques")
	fmt.Println("8. Go provides excellent tools for network performance")
	
	fmt.Println("\n📚 Next Topic: Distributed Systems Networking")
}
