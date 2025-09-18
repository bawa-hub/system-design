package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// HTTPServer represents an advanced HTTP server
type HTTPServer struct {
	Server      *http.Server
	Mux         *http.ServeMux
	Middlewares []func(http.Handler) http.Handler
	Routes      map[string]http.Handler
	Stats       *ServerStats
	Config      *ServerConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Address         string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	MaxHeaderBytes  int
	EnableTLS       bool
	CertFile        string
	KeyFile         string
	EnableCORS      bool
	EnableCompression bool
	EnableLogging   bool
}

// ServerStats holds server statistics
type ServerStats struct {
	Requests        int64
	Responses       int64
	Errors          int64
	BytesRead       int64
	BytesWritten    int64
	StartTime       time.Time
	ActiveConnections int64
	Mutex           sync.RWMutex
}

// NewHTTPServer creates a new HTTP server
func NewHTTPServer(config *ServerConfig) *HTTPServer {
	mux := http.NewServeMux()
	
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		IdleTimeout:    config.IdleTimeout,
		MaxHeaderBytes: config.MaxHeaderBytes,
	}
	
	return &HTTPServer{
		Server:      server,
		Mux:         mux,
		Middlewares: make([]func(http.Handler) http.Handler, 0),
		Routes:      make(map[string]http.Handler),
		Stats:       &ServerStats{StartTime: time.Now()},
		Config:      config,
	}
}

// AddMiddleware adds middleware to the server
func (s *HTTPServer) AddMiddleware(middleware func(http.Handler) http.Handler) {
	s.Middlewares = append(s.Middlewares, middleware)
}

// AddRoute adds a route to the server
func (s *HTTPServer) AddRoute(pattern string, handler http.Handler) {
	s.Routes[pattern] = handler
	s.Mux.Handle(pattern, s.wrapHandler(handler))
}

// AddRouteFunc adds a route function to the server
func (s *HTTPServer) AddRouteFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.Mux.HandleFunc(pattern, s.wrapHandlerFunc(handler))
}

// wrapHandler wraps a handler with middlewares
func (s *HTTPServer) wrapHandler(handler http.Handler) http.Handler {
	for i := len(s.Middlewares) - 1; i >= 0; i-- {
		handler = s.Middlewares[i](handler)
	}
	return handler
}

// wrapHandlerFunc wraps a handler function with middlewares
func (s *HTTPServer) wrapHandlerFunc(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return s.wrapHandler(http.HandlerFunc(handler))
}

// Start starts the HTTP server
func (s *HTTPServer) Start() error {
	// Add default routes
	s.setupDefaultRoutes()
	
	// Add middlewares
	s.setupMiddlewares()
	
	fmt.Printf("HTTP Server starting on %s\n", s.Config.Address)
	
	if s.Config.EnableTLS {
		return s.Server.ListenAndServeTLS(s.Config.CertFile, s.Config.KeyFile)
	}
	return s.Server.ListenAndServe()
}

// Stop stops the HTTP server
func (s *HTTPServer) Stop(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}

// setupDefaultRoutes sets up default routes
func (s *HTTPServer) setupDefaultRoutes() {
	s.AddRouteFunc("/", s.handleRoot)
	s.AddRouteFunc("/health", s.handleHealth)
	s.AddRouteFunc("/stats", s.handleStats)
	s.AddRouteFunc("/api/users", s.handleUsers)
	s.AddRouteFunc("/api/users/", s.handleUserByID)
}

// setupMiddlewares sets up default middlewares
func (s *HTTPServer) setupMiddlewares() {
	if s.Config.EnableLogging {
		s.AddMiddleware(s.loggingMiddleware)
	}
	if s.Config.EnableCORS {
		s.AddMiddleware(s.corsMiddleware)
	}
	if s.Config.EnableCompression {
		s.AddMiddleware(s.compressionMiddleware)
	}
	s.AddMiddleware(s.statsMiddleware)
}

// handleRoot handles the root route
func (s *HTTPServer) handleRoot(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "Welcome to the HTTP Server",
		"version": "1.0.0",
		"time":    time.Now().Format(time.RFC3339),
	}
	
	s.writeJSONResponse(w, http.StatusOK, response)
}

// handleHealth handles the health check route
func (s *HTTPServer) handleHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"uptime":    time.Since(s.Stats.StartTime).String(),
	}
	
	s.writeJSONResponse(w, http.StatusOK, response)
}

// handleStats handles the stats route
func (s *HTTPServer) handleStats(w http.ResponseWriter, r *http.Request) {
	s.Stats.Mutex.RLock()
	stats := map[string]interface{}{
		"requests":           s.Stats.Requests,
		"responses":          s.Stats.Responses,
		"errors":             s.Stats.Errors,
		"bytes_read":         s.Stats.BytesRead,
		"bytes_written":      s.Stats.BytesWritten,
		"active_connections": s.Stats.ActiveConnections,
		"uptime":             time.Since(s.Stats.StartTime).String(),
	}
	s.Stats.Mutex.RUnlock()
	
	s.writeJSONResponse(w, http.StatusOK, stats)
}

// handleUsers handles the users API route
func (s *HTTPServer) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getUsers(w, r)
	case http.MethodPost:
		s.createUser(w, r)
	default:
		s.writeErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// handleUserByID handles user by ID API route
func (s *HTTPServer) handleUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from path
	path := strings.TrimPrefix(r.URL.Path, "/api/users/")
	if path == "" {
		s.writeErrorResponse(w, http.StatusBadRequest, "User ID required")
		return
	}
	
	userID, err := strconv.Atoi(path)
	if err != nil {
		s.writeErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		s.getUserByID(w, r, userID)
	case http.MethodPut:
		s.updateUser(w, r, userID)
	case http.MethodDelete:
		s.deleteUser(w, r, userID)
	default:
		s.writeErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// getUsers gets all users
func (s *HTTPServer) getUsers(w http.ResponseWriter, r *http.Request) {
	users := []map[string]interface{}{
		{"id": 1, "name": "John Doe", "email": "john@example.com"},
		{"id": 2, "name": "Jane Smith", "email": "jane@example.com"},
		{"id": 3, "name": "Bob Johnson", "email": "bob@example.com"},
	}
	
	s.writeJSONResponse(w, http.StatusOK, users)
}

// createUser creates a new user
func (s *HTTPServer) createUser(w http.ResponseWriter, r *http.Request) {
	var user map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		s.writeErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	
	// Simulate user creation
	user["id"] = 4
	user["created_at"] = time.Now().Format(time.RFC3339)
	
	s.writeJSONResponse(w, http.StatusCreated, user)
}

// getUserByID gets a user by ID
func (s *HTTPServer) getUserByID(w http.ResponseWriter, r *http.Request, userID int) {
	// Simulate user lookup
	if userID < 1 || userID > 3 {
		s.writeErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}
	
	user := map[string]interface{}{
		"id":    userID,
		"name":  fmt.Sprintf("User %d", userID),
		"email": fmt.Sprintf("user%d@example.com", userID),
	}
	
	s.writeJSONResponse(w, http.StatusOK, user)
}

// updateUser updates a user
func (s *HTTPServer) updateUser(w http.ResponseWriter, r *http.Request, userID int) {
	var user map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		s.writeErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	
	// Simulate user update
	user["id"] = userID
	user["updated_at"] = time.Now().Format(time.RFC3339)
	
	s.writeJSONResponse(w, http.StatusOK, user)
}

// deleteUser deletes a user
func (s *HTTPServer) deleteUser(w http.ResponseWriter, r *http.Request, userID int) {
	// Simulate user deletion
	response := map[string]interface{}{
		"message": fmt.Sprintf("User %d deleted", userID),
		"deleted_at": time.Now().Format(time.RFC3339),
	}
	
	s.writeJSONResponse(w, http.StatusOK, response)
}

// writeJSONResponse writes a JSON response
func (s *HTTPServer) writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

// writeErrorResponse writes an error response
func (s *HTTPServer) writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response := map[string]interface{}{
		"error":   http.StatusText(statusCode),
		"message": message,
		"code":    statusCode,
	}
	
	s.writeJSONResponse(w, statusCode, response)
}

// Middleware functions
func (s *HTTPServer) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Create a response writer wrapper to capture status code
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		
		next.ServeHTTP(wrapped, r)
		
		duration := time.Since(start)
		log.Printf("%s %s %d %v", r.Method, r.URL.Path, wrapped.statusCode, duration)
	})
}

func (s *HTTPServer) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func (s *HTTPServer) compressionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if client accepts gzip
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")
			// In a real implementation, you would use gzip.NewWriter
		}
		
		next.ServeHTTP(w, r)
	})
}

func (s *HTTPServer) statsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Stats.Mutex.Lock()
		s.Stats.Requests++
		s.Stats.ActiveConnections++
		s.Stats.Mutex.Unlock()
		
		// Create a response writer wrapper to capture bytes written
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		
		next.ServeHTTP(wrapped, r)
		
		s.Stats.Mutex.Lock()
		s.Stats.Responses++
		s.Stats.ActiveConnections--
		s.Stats.BytesWritten += int64(wrapped.bytesWritten)
		s.Stats.Mutex.Unlock()
	})
}

// responseWriter wraps http.ResponseWriter to capture status code and bytes written
type responseWriter struct {
	http.ResponseWriter
	statusCode   int
	bytesWritten int64
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.bytesWritten += int64(n)
	return n, err
}

// HTTPClient represents an advanced HTTP client
type HTTPClient struct {
	Client  *http.Client
	BaseURL string
	Headers map[string]string
}

// NewHTTPClient creates a new HTTP client
func NewHTTPClient(baseURL string) *HTTPClient {
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}
	
	return &HTTPClient{
		Client:  client,
		BaseURL: baseURL,
		Headers: make(map[string]string),
	}
}

// SetHeader sets a header for all requests
func (c *HTTPClient) SetHeader(key, value string) {
	c.Headers[key] = value
}

// Get makes a GET request
func (c *HTTPClient) Get(path string) (*http.Response, error) {
	return c.makeRequest(http.MethodGet, path, nil)
}

// Post makes a POST request
func (c *HTTPClient) Post(path string, body io.Reader) (*http.Response, error) {
	return c.makeRequest(http.MethodPost, path, body)
}

// Put makes a PUT request
func (c *HTTPClient) Put(path string, body io.Reader) (*http.Response, error) {
	return c.makeRequest(http.MethodPut, path, body)
}

// Delete makes a DELETE request
func (c *HTTPClient) Delete(path string) (*http.Response, error) {
	return c.makeRequest(http.MethodDelete, path, nil)
}

// makeRequest makes an HTTP request
func (c *HTTPClient) makeRequest(method, path string, body io.Reader) (*http.Response, error) {
	url := c.BaseURL + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	
	// Set headers
	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}
	
	return c.Client.Do(req)
}

// HTTPSClient represents an HTTPS client with TLS configuration
type HTTPSClient struct {
	*HTTPClient
	TLSConfig *tls.Config
}

// NewHTTPSClient creates a new HTTPS client
func NewHTTPSClient(baseURL string, tlsConfig *tls.Config) *HTTPSClient {
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
			MaxIdleConns:    100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout: 90 * time.Second,
		},
	}
	
	return &HTTPSClient{
		HTTPClient: &HTTPClient{
			Client:  client,
			BaseURL: baseURL,
			Headers: make(map[string]string),
		},
		TLSConfig: tlsConfig,
	}
}

// LoadBalancer represents a simple load balancer
type LoadBalancer struct {
	Servers []string
	Current int
	Mutex   sync.Mutex
}

// NewLoadBalancer creates a new load balancer
func NewLoadBalancer(servers []string) *LoadBalancer {
	return &LoadBalancer{
		Servers: servers,
		Current: 0,
	}
}

// GetNextServer gets the next server using round-robin
func (lb *LoadBalancer) GetNextServer() string {
	lb.Mutex.Lock()
	defer lb.Mutex.Unlock()
	
	server := lb.Servers[lb.Current]
	lb.Current = (lb.Current + 1) % len(lb.Servers)
	return server
}

// Proxy represents an HTTP proxy
type Proxy struct {
	LoadBalancer *LoadBalancer
	Client       *HTTPClient
}

// NewProxy creates a new HTTP proxy
func NewProxy(servers []string) *Proxy {
	return &Proxy{
		LoadBalancer: NewLoadBalancer(servers),
		Client:       NewHTTPClient(""),
	}
}

// ServeHTTP implements http.Handler interface
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get next server
	server := p.LoadBalancer.GetNextServer()
	
	// Create new request
	req, err := http.NewRequest(r.Method, server+r.URL.Path, r.Body)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}
	
	// Copy headers
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	
	// Make request
	resp, err := p.Client.Client.Do(req)
	if err != nil {
		http.Error(w, "Error making request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	
	// Copy response headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	
	// Copy response status
	w.WriteHeader(resp.StatusCode)
	
	// Copy response body
	io.Copy(w, resp.Body)
}

// Demonstrate basic HTTP server
func demonstrateBasicHTTPServer() {
	fmt.Println("=== Basic HTTP Server Demo ===\n")
	
	config := &ServerConfig{
		Address:         ":8080",
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
		IdleTimeout:     60 * time.Second,
		MaxHeaderBytes:  1 << 20, // 1MB
		EnableTLS:       false,
		EnableCORS:      true,
		EnableCompression: true,
		EnableLogging:   true,
	}
	
	server := NewHTTPServer(config)
	
	// Start server in goroutine
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
		}
	}()
	
	// Give server time to start
	time.Sleep(100 * time.Millisecond)
	
	// Test server with client
	client := NewHTTPClient("http://localhost:8080")
	client.SetHeader("User-Agent", "Demo-Client/1.0")
	
	// Test different endpoints
	endpoints := []string{"/", "/health", "/stats", "/api/users"}
	
	for _, endpoint := range endpoints {
		resp, err := client.Get(endpoint)
		if err != nil {
			log.Printf("Error making request to %s: %v", endpoint, err)
			continue
		}
		
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Printf("Error reading response: %v", err)
			continue
		}
		
		fmt.Printf("GET %s: %d\n", endpoint, resp.StatusCode)
		fmt.Printf("Response: %s\n\n", string(body))
	}
	
	// Stop server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Stop(ctx)
}

// Demonstrate HTTP client
func demonstrateHTTPClient() {
	fmt.Println("=== HTTP Client Demo ===\n")
	
	client := NewHTTPClient("https://httpbin.org")
	client.SetHeader("User-Agent", "Go-HTTP-Client/1.0")
	
	// Test GET request
	resp, err := client.Get("/get")
	if err != nil {
		log.Printf("GET request error: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return
	}
	
	fmt.Printf("GET /get: %d\n", resp.StatusCode)
	fmt.Printf("Response: %s\n\n", string(body))
}

// Demonstrate load balancer
func demonstrateLoadBalancer() {
	fmt.Println("=== Load Balancer Demo ===\n")
	
	servers := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}
	
	lb := NewLoadBalancer(servers)
	
	// Simulate requests
	for i := 0; i < 10; i++ {
		server := lb.GetNextServer()
		fmt.Printf("Request %d: %s\n", i+1, server)
	}
}

// Demonstrate HTTPS server
func demonstrateHTTPSServer() {
	fmt.Println("=== HTTPS Server Demo ===\n")
	
	// Note: In a real implementation, you would need actual certificates
	// For demo purposes, we'll show the configuration
	
	config := &ServerConfig{
		Address:         ":8443",
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
		IdleTimeout:     60 * time.Second,
		MaxHeaderBytes:  1 << 20,
		EnableTLS:       true,
		CertFile:        "server.crt",
		KeyFile:         "server.key",
		EnableCORS:      true,
		EnableCompression: true,
		EnableLogging:   true,
	}
	
	server := NewHTTPServer(config)
	
	fmt.Printf("HTTPS Server configured for %s\n", config.Address)
	fmt.Printf("Certificate: %s\n", config.CertFile)
	fmt.Printf("Private Key: %s\n", config.KeyFile)
	
	// In a real implementation, you would start the server
	// server.Start()
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 6: HTTP/HTTPS Implementation")
	fmt.Println("================================================================\n")
	
	// Run all demonstrations
	demonstrateBasicHTTPServer()
	fmt.Println()
	demonstrateHTTPClient()
	fmt.Println()
	demonstrateLoadBalancer()
	fmt.Println()
	demonstrateHTTPSServer()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. HTTP is the foundation of web communication")
	fmt.Println("2. Go provides excellent HTTP server and client support")
	fmt.Println("3. Middleware enables powerful request/response processing")
	fmt.Println("4. HTTPS adds security through TLS encryption")
	fmt.Println("5. Load balancing distributes load across multiple servers")
	fmt.Println("6. REST APIs follow standard HTTP methods and status codes")
	
	fmt.Println("\n📚 Next Topic: WebSocket Protocols")
}
