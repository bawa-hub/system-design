package main

import (
	"bufio"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// SocketServer represents a generic socket server
type SocketServer struct {
	Address     string
	Protocol    string
	Listener    net.Listener
	Connections map[net.Conn]*Connection
	Mutex       sync.RWMutex
	Running     bool
	Handler     func(net.Conn)
	WorkerPool  *WorkerPool
}

// Connection represents a client connection
type Connection struct {
	Conn        net.Conn
	ID          string
	CreatedAt   time.Time
	LastActivity time.Time
	BytesRead   int64
	BytesWritten int64
	State       string
}

// WorkerPool manages a pool of workers
type WorkerPool struct {
	Workers    int
	JobQueue   chan func()
	Quit       chan bool
	Wg         sync.WaitGroup
	Running    bool
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(workers int, queueSize int) *WorkerPool {
	return &WorkerPool{
		Workers:  workers,
		JobQueue: make(chan func(), queueSize),
		Quit:     make(chan bool),
		Running:  false,
	}
}

// Start starts the worker pool
func (wp *WorkerPool) Start() {
	wp.Running = true
	for i := 0; i < wp.Workers; i++ {
		wp.Wg.Add(1)
		go wp.worker()
	}
}

// Stop stops the worker pool
func (wp *WorkerPool) Stop() {
	wp.Running = false
	close(wp.Quit)
	wp.Wg.Wait()
}

// worker processes jobs from the queue
func (wp *WorkerPool) worker() {
	defer wp.Wg.Done()
	for {
		select {
		case job := <-wp.JobQueue:
			job()
		case <-wp.Quit:
			return
		}
	}
}

// Submit submits a job to the worker pool
func (wp *WorkerPool) Submit(job func()) {
	if wp.Running {
		wp.JobQueue <- job
	}
}

// NewSocketServer creates a new socket server
func NewSocketServer(address, protocol string) *SocketServer {
	return &SocketServer{
		Address:     address,
		Protocol:    protocol,
		Connections: make(map[net.Conn]*Connection),
		WorkerPool:  NewWorkerPool(runtime.NumCPU(), 1000),
	}
}

// Start starts the socket server
func (s *SocketServer) Start() error {
	var err error
	s.Listener, err = net.Listen(s.Protocol, s.Address)
	if err != nil {
		return err
	}
	
	s.Running = true
	s.WorkerPool.Start()
	
	fmt.Printf("Socket server listening on %s (%s)\n", s.Address, s.Protocol)
	
	// Start accepting connections
	go s.acceptConnections()
	
	return nil
}

// acceptConnections accepts incoming connections
func (s *SocketServer) acceptConnections() {
	for s.Running {
		conn, err := s.Listener.Accept()
		if err != nil {
			if s.Running {
				log.Printf("Error accepting connection: %v", err)
			}
			continue
		}
		
		// Create connection record
		connection := &Connection{
			Conn:         conn,
			ID:           generateConnectionID(),
			CreatedAt:    time.Now(),
			LastActivity: time.Now(),
			State:        "CONNECTED",
		}
		
		s.Mutex.Lock()
		s.Connections[conn] = connection
		s.Mutex.Unlock()
		
		// Handle connection in worker pool
		s.WorkerPool.Submit(func() {
			s.handleConnection(conn)
		})
	}
}

// handleConnection handles a client connection
func (s *SocketServer) handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		s.Mutex.Lock()
		delete(s.Connections, conn)
		s.Mutex.Unlock()
	}()
	
	// Set connection timeout
	conn.SetDeadline(time.Now().Add(30 * time.Second))
	
	// Handle the connection
	if s.Handler != nil {
		s.Handler(conn)
	} else {
		s.defaultHandler(conn)
	}
}

// defaultHandler is the default connection handler
func (s *SocketServer) defaultHandler(conn net.Conn) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	
	for {
		// Read data from connection
		data, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading from connection: %v", err)
			}
			break
		}
		
		// Update connection stats
		s.Mutex.Lock()
		if connInfo, exists := s.Connections[conn]; exists {
			connInfo.LastActivity = time.Now()
			connInfo.BytesRead += int64(len(data))
		}
		s.Mutex.Unlock()
		
		// Echo the data back
		response := fmt.Sprintf("Echo: %s", data)
		writer.WriteString(response)
		writer.Flush()
		
		// Update connection stats
		s.Mutex.Lock()
		if connInfo, exists := s.Connections[conn]; exists {
			connInfo.BytesWritten += int64(len(response))
		}
		s.Mutex.Unlock()
	}
}

// Stop stops the socket server
func (s *SocketServer) Stop() error {
	s.Running = false
	
	// Close all connections
	s.Mutex.Lock()
	for conn := range s.Connections {
		conn.Close()
	}
	s.Connections = make(map[net.Conn]*Connection)
	s.Mutex.Unlock()
	
	// Stop worker pool
	s.WorkerPool.Stop()
	
	// Close listener
	if s.Listener != nil {
		return s.Listener.Close()
	}
	
	return nil
}

// GetStats returns server statistics
func (s *SocketServer) GetStats() map[string]interface{} {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	
	stats := map[string]interface{}{
		"connections": len(s.Connections),
		"workers":     s.WorkerPool.Workers,
		"running":     s.Running,
	}
	
	return stats
}

// generateConnectionID generates a unique connection ID
func generateConnectionID() string {
	return fmt.Sprintf("conn_%d", time.Now().UnixNano())
}

// AdvancedSocketClient represents an advanced socket client
type AdvancedSocketClient struct {
	Address      string
	Protocol     string
	Conn         net.Conn
	RetryCount   int
	RetryDelay   time.Duration
	MaxRetries   int
	Timeout      time.Duration
	KeepAlive    bool
	BufferSize   int
}

// NewAdvancedSocketClient creates a new advanced socket client
func NewAdvancedSocketClient(address, protocol string) *AdvancedSocketClient {
	return &AdvancedSocketClient{
		Address:    address,
		Protocol:   protocol,
		MaxRetries: 3,
		RetryDelay: 1 * time.Second,
		Timeout:    10 * time.Second,
		KeepAlive:  true,
		BufferSize: 4096,
	}
}

// Connect connects to the server with retry logic
func (c *AdvancedSocketClient) Connect() error {
	var err error
	
	for c.RetryCount < c.MaxRetries {
		c.Conn, err = net.DialTimeout(c.Protocol, c.Address, c.Timeout)
		if err == nil {
			// Configure connection
			if tcpConn, ok := c.Conn.(*net.TCPConn); ok {
				tcpConn.SetKeepAlive(c.KeepAlive)
				tcpConn.SetNoDelay(true)
			}
			
			fmt.Printf("Connected to %s (%s)\n", c.Address, c.Protocol)
			return nil
		}
		
		c.RetryCount++
		if c.RetryCount < c.MaxRetries {
			log.Printf("Connection attempt %d failed: %v, retrying in %v", 
				c.RetryCount, err, c.RetryDelay)
			time.Sleep(c.RetryDelay)
			c.RetryDelay *= 2 // Exponential backoff
		}
	}
	
	return fmt.Errorf("failed to connect after %d attempts: %v", c.MaxRetries, err)
}

// Send sends data to the server
func (c *AdvancedSocketClient) Send(data []byte) error {
	if c.Conn == nil {
		return fmt.Errorf("not connected")
	}
	
	// Set write timeout
	c.Conn.SetWriteDeadline(time.Now().Add(c.Timeout))
	
	_, err := c.Conn.Write(data)
	return err
}

// SendString sends a string to the server
func (c *AdvancedSocketClient) SendString(data string) error {
	return c.Send([]byte(data))
}

// Receive receives data from the server
func (c *AdvancedSocketClient) Receive() ([]byte, error) {
	if c.Conn == nil {
		return nil, fmt.Errorf("not connected")
	}
	
	// Set read timeout
	c.Conn.SetReadDeadline(time.Now().Add(c.Timeout))
	
	buffer := make([]byte, c.BufferSize)
	n, err := c.Conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	
	return buffer[:n], nil
}

// Close closes the connection
func (c *AdvancedSocketClient) Close() error {
	if c.Conn != nil {
		return c.Conn.Close()
	}
	return nil
}

// SocketMultiplexer handles multiple connections efficiently
type SocketMultiplexer struct {
	Connections map[net.Conn]*Connection
	Mutex       sync.RWMutex
	Running     bool
	Quit        chan bool
}

// NewSocketMultiplexer creates a new socket multiplexer
func NewSocketMultiplexer() *SocketMultiplexer {
	return &SocketMultiplexer{
		Connections: make(map[net.Conn]*Connection),
		Quit:        make(chan bool),
	}
}

// AddConnection adds a connection to the multiplexer
func (sm *SocketMultiplexer) AddConnection(conn net.Conn) {
	sm.Mutex.Lock()
	defer sm.Mutex.Unlock()
	
	connection := &Connection{
		Conn:         conn,
		ID:           generateConnectionID(),
		CreatedAt:    time.Now(),
		LastActivity: time.Now(),
		State:        "CONNECTED",
	}
	
	sm.Connections[conn] = connection
}

// RemoveConnection removes a connection from the multiplexer
func (sm *SocketMultiplexer) RemoveConnection(conn net.Conn) {
	sm.Mutex.Lock()
	defer sm.Mutex.Unlock()
	
	delete(sm.Connections, conn)
}

// Start starts the multiplexer
func (sm *SocketMultiplexer) Start() {
	sm.Running = true
	go sm.multiplex()
}

// Stop stops the multiplexer
func (sm *SocketMultiplexer) Stop() {
	sm.Running = false
	close(sm.Quit)
}

// multiplex handles multiple connections
func (sm *SocketMultiplexer) multiplex() {
	for sm.Running {
		select {
		case <-sm.Quit:
			return
		default:
			sm.handleConnections()
		}
	}
}

// handleConnections handles all connections
func (sm *SocketMultiplexer) handleConnections() {
	sm.Mutex.RLock()
	connections := make([]net.Conn, 0, len(sm.Connections))
	for conn := range sm.Connections {
		connections = append(connections, conn)
	}
	sm.Mutex.RUnlock()
	
	// Process each connection
	for _, conn := range connections {
		sm.processConnection(conn)
	}
}

// processConnection processes a single connection
func (sm *SocketMultiplexer) processConnection(conn net.Conn) {
	// Set non-blocking read
	conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
	
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return // Timeout, continue
		}
		// Connection error, remove it
		sm.RemoveConnection(conn)
		conn.Close()
		return
	}
	
	// Process received data
	if n > 0 {
		sm.handleData(conn, buffer[:n])
	}
}

// handleData handles received data
func (sm *SocketMultiplexer) handleData(conn net.Conn, data []byte) {
	// Echo the data back
	response := fmt.Sprintf("Echo: %s", string(data))
	conn.Write([]byte(response))
	
	// Update connection stats
	sm.Mutex.Lock()
	if connInfo, exists := sm.Connections[conn]; exists {
		connInfo.LastActivity = time.Now()
		connInfo.BytesRead += int64(len(data))
		connInfo.BytesWritten += int64(len(response))
	}
	sm.Mutex.Unlock()
}

// ConnectionPool manages a pool of connections
type ConnectionPool struct {
	Address     string
	Protocol    string
	MaxConns    int
	MinConns    int
	Connections chan net.Conn
	Mutex       sync.RWMutex
	Running     bool
	Quit        chan bool
}

// NewConnectionPool creates a new connection pool
func NewConnectionPool(address, protocol string, maxConns, minConns int) *ConnectionPool {
	return &ConnectionPool{
		Address:     address,
		Protocol:    protocol,
		MaxConns:    maxConns,
		MinConns:    minConns,
		Connections: make(chan net.Conn, maxConns),
		Quit:        make(chan bool),
	}
}

// Start starts the connection pool
func (cp *ConnectionPool) Start() error {
	cp.Running = true
	
	// Create minimum connections
	for i := 0; i < cp.MinConns; i++ {
		conn, err := net.Dial(cp.Protocol, cp.Address)
		if err != nil {
			return err
		}
		cp.Connections <- conn
	}
	
	// Start connection manager
	go cp.manageConnections()
	
	return nil
}

// manageConnections manages the connection pool
func (cp *ConnectionPool) manageConnections() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for cp.Running {
		select {
		case <-ticker.C:
			cp.healthCheck()
		case <-cp.Quit:
			return
		}
	}
}

// healthCheck checks connection health
func (cp *ConnectionPool) healthCheck() {
	cp.Mutex.Lock()
	defer cp.Mutex.Unlock()
	
	// Check if we need more connections
	if len(cp.Connections) < cp.MinConns {
		cp.createConnections(cp.MinConns - len(cp.Connections))
	}
}

// createConnections creates new connections
func (cp *ConnectionPool) createConnections(count int) {
	for i := 0; i < count; i++ {
		conn, err := net.Dial(cp.Protocol, cp.Address)
		if err != nil {
			log.Printf("Failed to create connection: %v", err)
			continue
		}
		
		select {
		case cp.Connections <- conn:
		default:
			conn.Close() // Pool is full
		}
	}
}

// GetConnection gets a connection from the pool
func (cp *ConnectionPool) GetConnection() (net.Conn, error) {
	select {
	case conn := <-cp.Connections:
		return conn, nil
	default:
		// Create new connection if pool is empty
		conn, err := net.Dial(cp.Protocol, cp.Address)
		if err != nil {
			return nil, err
		}
		return conn, nil
	}
}

// ReturnConnection returns a connection to the pool
func (cp *ConnectionPool) ReturnConnection(conn net.Conn) {
	select {
	case cp.Connections <- conn:
	default:
		conn.Close() // Pool is full
	}
}

// Stop stops the connection pool
func (cp *ConnectionPool) Stop() {
	cp.Running = false
	close(cp.Quit)
	
	// Close all connections
	close(cp.Connections)
	for conn := range cp.Connections {
		conn.Close()
	}
}

// SocketMonitor monitors socket statistics
type SocketMonitor struct {
	Connections    int64
	BytesRead      int64
	BytesWritten   int64
	Errors         int64
	StartTime      time.Time
	Mutex          sync.RWMutex
}

// NewSocketMonitor creates a new socket monitor
func NewSocketMonitor() *SocketMonitor {
	return &SocketMonitor{
		StartTime: time.Now(),
	}
}

// IncrementConnections increments the connection count
func (sm *SocketMonitor) IncrementConnections() {
	atomic.AddInt64(&sm.Connections, 1)
}

// DecrementConnections decrements the connection count
func (sm *SocketMonitor) DecrementConnections() {
	atomic.AddInt64(&sm.Connections, -1)
}

// AddBytesRead adds bytes read
func (sm *SocketMonitor) AddBytesRead(bytes int64) {
	atomic.AddInt64(&sm.BytesRead, bytes)
}

// AddBytesWritten adds bytes written
func (sm *SocketMonitor) AddBytesWritten(bytes int64) {
	atomic.AddInt64(&sm.BytesWritten, bytes)
}

// IncrementErrors increments the error count
func (sm *SocketMonitor) IncrementErrors() {
	atomic.AddInt64(&sm.Errors, 1)
}

// GetStats returns current statistics
func (sm *SocketMonitor) GetStats() map[string]interface{} {
	sm.Mutex.RLock()
	defer sm.Mutex.RUnlock()
	
	uptime := time.Since(sm.StartTime)
	
	return map[string]interface{}{
		"connections":    atomic.LoadInt64(&sm.Connections),
		"bytes_read":     atomic.LoadInt64(&sm.BytesRead),
		"bytes_written":  atomic.LoadInt64(&sm.BytesWritten),
		"errors":         atomic.LoadInt64(&sm.Errors),
		"uptime":         uptime.String(),
		"uptime_seconds": uptime.Seconds(),
	}
}

// Demonstrate basic socket server
func demonstrateBasicSocketServer() {
	fmt.Println("=== Basic Socket Server Demo ===\n")
	
	server := NewSocketServer(":8080", "tcp")
	
	// Set custom handler
	server.Handler = func(conn net.Conn) {
		reader := bufio.NewReader(conn)
		writer := bufio.NewWriter(conn)
		
		for {
			// Read request
			request, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			
			// Process request
			response := fmt.Sprintf("Server received: %s", request)
			
			// Send response
			writer.WriteString(response)
			writer.Flush()
		}
	}
	
	// Start server
	if err := server.Start(); err != nil {
		log.Printf("Server error: %v", err)
		return
	}
	
	// Give server time to start
	time.Sleep(100 * time.Millisecond)
	
	// Create client and test
	client := NewAdvancedSocketClient("localhost:8080", "tcp")
	if err := client.Connect(); err != nil {
		log.Printf("Client connection error: %v", err)
		return
	}
	defer client.Close()
	
	// Send test messages
	testMessages := []string{
		"Hello, Socket Server!\n",
		"How are you?\n",
		"Testing socket communication\n",
	}
	
	for _, msg := range testMessages {
		if err := client.SendString(msg); err != nil {
			log.Printf("Send error: %v", err)
			continue
		}
		
		response, err := client.Receive()
		if err != nil {
			log.Printf("Receive error: %v", err)
			continue
		}
		
		fmt.Printf("Sent: %s", msg)
		fmt.Printf("Received: %s", string(response))
	}
	
	// Stop server
	server.Stop()
}

// Demonstrate connection pooling
func demonstrateConnectionPool() {
	fmt.Println("=== Connection Pool Demo ===\n")
	
	// Start a simple server
	server := NewSocketServer(":8081", "tcp")
	server.Start()
	defer server.Stop()
	
	time.Sleep(100 * time.Millisecond)
	
	// Create connection pool
	pool := NewConnectionPool("localhost:8081", "tcp", 5, 2)
	if err := pool.Start(); err != nil {
		log.Printf("Pool start error: %v", err)
		return
	}
	defer pool.Stop()
	
	// Use connections from pool
	for i := 0; i < 10; i++ {
		conn, err := pool.GetConnection()
		if err != nil {
			log.Printf("Get connection error: %v", err)
			continue
		}
		
		// Use connection
		conn.Write([]byte(fmt.Sprintf("Message %d\n", i)))
		
		// Return connection to pool
		pool.ReturnConnection(conn)
	}
	
	fmt.Println("Connection pool demo completed")
}

// Demonstrate socket multiplexing
func demonstrateSocketMultiplexing() {
	fmt.Println("=== Socket Multiplexing Demo ===\n")
	
	// Start a simple server
	server := NewSocketServer(":8082", "tcp")
	server.Start()
	defer server.Stop()
	
	time.Sleep(100 * time.Millisecond)
	
	// Create multiplexer
	multiplexer := NewSocketMultiplexer()
	multiplexer.Start()
	defer multiplexer.Stop()
	
	// Create multiple connections
	connections := make([]net.Conn, 5)
	for i := 0; i < 5; i++ {
		conn, err := net.Dial("tcp", "localhost:8082")
		if err != nil {
			log.Printf("Connection error: %v", err)
			continue
		}
		
		connections[i] = conn
		multiplexer.AddConnection(conn)
	}
	
	// Send data through connections
	for i, conn := range connections {
		if conn != nil {
			conn.Write([]byte(fmt.Sprintf("Message from connection %d\n", i)))
		}
	}
	
	// Wait a bit
	time.Sleep(1 * time.Second)
	
	// Close connections
	for _, conn := range connections {
		if conn != nil {
			conn.Close()
		}
	}
	
	fmt.Println("Socket multiplexing demo completed")
}

// Demonstrate socket monitoring
func demonstrateSocketMonitoring() {
	fmt.Println("=== Socket Monitoring Demo ===\n")
	
	monitor := NewSocketMonitor()
	
	// Simulate some activity
	monitor.IncrementConnections()
	monitor.AddBytesRead(1024)
	monitor.AddBytesWritten(2048)
	monitor.IncrementErrors()
	
	// Display statistics
	stats := monitor.GetStats()
	fmt.Printf("Socket Statistics:\n")
	for key, value := range stats {
		fmt.Printf("  %s: %v\n", key, value)
	}
}

// Demonstrate context-based socket operations
func demonstrateContextSocketOperations() {
	fmt.Println("=== Context Socket Operations Demo ===\n")
	
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Start server
	server := NewSocketServer(":8083", "tcp")
	server.Start()
	defer server.Stop()
	
	time.Sleep(100 * time.Millisecond)
	
	// Connect with context
	conn, err := net.DialContext(ctx, "tcp", "localhost:8083")
	if err != nil {
		log.Printf("Connection error: %v", err)
		return
	}
	defer conn.Close()
	
	// Send data with context
	select {
	case <-ctx.Done():
		fmt.Println("Operation cancelled by context")
		return
	default:
		conn.Write([]byte("Hello with context!\n"))
	}
	
	fmt.Println("Context socket operations demo completed")
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 5: Socket Programming in Go")
	fmt.Println("================================================================\n")
	
	// Run all demonstrations
	demonstrateBasicSocketServer()
	fmt.Println()
	demonstrateConnectionPool()
	fmt.Println()
	demonstrateSocketMultiplexing()
	fmt.Println()
	demonstrateSocketMonitoring()
	fmt.Println()
	demonstrateContextSocketOperations()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Sockets are the foundation of network communication")
	fmt.Println("2. Go provides excellent socket programming support")
	fmt.Println("3. Concurrency is key to handling multiple connections")
	fmt.Println("4. Connection pooling improves performance")
	fmt.Println("5. Monitoring helps with debugging and optimization")
	fmt.Println("6. Context provides cancellation and timeout support")
	
	fmt.Println("\n📚 Next Topic: HTTP/HTTPS Implementation")
}
