# HTTP/HTTPS Quick Reference Guide

## 🚀 Essential HTTP Concepts

### HTTP Methods
- **GET**: Retrieve data (idempotent, safe)
- **POST**: Create or submit data
- **PUT**: Create or update entire resource (idempotent)
- **PATCH**: Partial update
- **DELETE**: Remove resource (idempotent)
- **HEAD**: Get headers only (idempotent, safe)
- **OPTIONS**: Get allowed methods (idempotent, safe)

### HTTP Status Codes
- **200 OK**: Request successful
- **201 Created**: Resource created
- **204 No Content**: Success, no content
- **400 Bad Request**: Invalid request
- **401 Unauthorized**: Authentication required
- **403 Forbidden**: Access denied
- **404 Not Found**: Resource not found
- **405 Method Not Allowed**: Method not supported
- **500 Internal Server Error**: Server error
- **502 Bad Gateway**: Invalid upstream response
- **503 Service Unavailable**: Server overloaded

## 🔧 Go HTTP Server Patterns

### Basic HTTP Server
```go
func startHTTPServer() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}
```

### Advanced HTTP Server
```go
func startAdvancedHTTPServer() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)
    
    server := &http.Server{
        Addr:         ":8080",
        Handler:      mux,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  60 * time.Second,
    }
    
    log.Fatal(server.ListenAndServe())
}
```

### Middleware Pattern
```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        next.ServeHTTP(w, r)
    })
}
```

## 🌐 Go HTTP Client Patterns

### Basic HTTP Client
```go
func basicHTTPClient() {
    resp, err := http.Get("http://example.com")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(body))
}
```

### Advanced HTTP Client
```go
func advancedHTTPClient() {
    client := &http.Client{
        Timeout: 30 * time.Second,
        Transport: &http.Transport{
            MaxIdleConns:        100,
            MaxIdleConnsPerHost: 10,
            IdleConnTimeout:     90 * time.Second,
        },
    }
    
    req, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {
        log.Fatal(err)
    }
    
    req.Header.Set("User-Agent", "MyApp/1.0")
    
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
}
```

## 🔒 HTTPS Implementation

### HTTPS Server
```go
func startHTTPSServer() {
    server := &http.Server{
        Addr:    ":8443",
        Handler: mux,
    }
    
    log.Fatal(server.ListenAndServeTLS("server.crt", "server.key"))
}
```

### HTTPS Client
```go
func httpsClient() {
    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: false,
            },
        },
    }
    
    resp, err := client.Get("https://example.com")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
}
```

## 📊 REST API Patterns

### JSON Response
```go
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}
```

### Error Response
```go
func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
    response := map[string]interface{}{
        "error":   http.StatusText(statusCode),
        "message": message,
        "code":    statusCode,
    }
    writeJSONResponse(w, statusCode, response)
}
```

### Request Validation
```go
func validateRequest(r *http.Request) error {
    if r.Method != http.MethodPost {
        return fmt.Errorf("method not allowed")
    }
    
    if r.Header.Get("Content-Type") != "application/json" {
        return fmt.Errorf("content type must be application/json")
    }
    
    return nil
}
```

## 🔄 Load Balancing

### Round Robin Load Balancer
```go
type LoadBalancer struct {
    Servers []string
    Current int
    Mutex   sync.Mutex
}

func (lb *LoadBalancer) GetNextServer() string {
    lb.Mutex.Lock()
    defer lb.Mutex.Unlock()
    
    server := lb.Servers[lb.Current]
    lb.Current = (lb.Current + 1) % len(lb.Servers)
    return server
}
```

### HTTP Proxy
```go
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    server := p.LoadBalancer.GetNextServer()
    
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
    
    resp, err := p.Client.Do(req)
    if err != nil {
        http.Error(w, "Error making request", http.StatusBadGateway)
        return
    }
    defer resp.Body.Close()
    
    // Copy response
    for key, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(key, value)
        }
    }
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}
```

## 📈 Performance Optimization

### Connection Pooling
```go
client := &http.Client{
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
}
```

### Compression
```go
func compressionMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
            w.Header().Set("Content-Encoding", "gzip")
            // Use gzip.NewWriter in real implementation
        }
        next.ServeHTTP(w, r)
    })
}
```

### Caching
```go
func cacheMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            w.Header().Set("Cache-Control", "public, max-age=3600")
        }
        next.ServeHTTP(w, r)
    })
}
```

## 🔐 Security Best Practices

### Input Validation
```go
func validateUserInput(user *User) error {
    if user.Name == "" {
        return fmt.Errorf("name is required")
    }
    if user.Email == "" {
        return fmt.Errorf("email is required")
    }
    if !strings.Contains(user.Email, "@") {
        return fmt.Errorf("invalid email format")
    }
    return nil
}
```

### Rate Limiting
```go
type RateLimiter struct {
    requests map[string][]time.Time
    mutex    sync.Mutex
    limit    int
    window   time.Duration
}

func (rl *RateLimiter) Allow(clientIP string) bool {
    rl.mutex.Lock()
    defer rl.mutex.Unlock()
    
    now := time.Now()
    cutoff := now.Add(-rl.window)
    
    // Clean old requests
    requests := rl.requests[clientIP]
    var validRequests []time.Time
    for _, req := range requests {
        if req.After(cutoff) {
            validRequests = append(validRequests, req)
        }
    }
    
    if len(validRequests) >= rl.limit {
        return false
    }
    
    validRequests = append(validRequests, now)
    rl.requests[clientIP] = validRequests
    return true
}
```

### CORS Configuration
```go
func corsMiddleware(next http.Handler) http.Handler {
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
```

## 🐛 Common Issues & Solutions

### Issue: Connection Refused
**Solution**: Check if server is running and port is available
```go
if err != nil {
    if strings.Contains(err.Error(), "connection refused") {
        log.Println("Server not running or port not available")
    }
}
```

### Issue: Timeout
**Solution**: Set appropriate timeouts
```go
client := &http.Client{
    Timeout: 30 * time.Second,
}
```

### Issue: Too Many Open Files
**Solution**: Use connection pooling and close connections
```go
client := &http.Client{
    Transport: &http.Transport{
        MaxIdleConns: 100,
        MaxIdleConnsPerHost: 10,
    },
}
```

## 📝 Best Practices

1. **Always check errors** from HTTP operations
2. **Close response bodies** to prevent resource leaks
3. **Use appropriate HTTP methods** for operations
4. **Set proper status codes** for responses
5. **Validate input** to prevent security issues
6. **Use HTTPS** for sensitive data
7. **Implement proper error handling** with meaningful messages
8. **Use middleware** for cross-cutting concerns
9. **Monitor performance** and optimize bottlenecks
10. **Follow REST principles** for API design

## 🎯 Interview Questions

### Basic
1. What's the difference between HTTP and HTTPS?
2. Explain HTTP methods and when to use them.
3. How do you handle errors in HTTP requests?

### Intermediate
1. How do you implement authentication in HTTP?
2. Explain HTTP caching and its benefits.
3. How do you optimize HTTP performance?

### Advanced
1. Design a RESTful API for a social media platform.
2. How would you implement rate limiting?
3. Explain HTTP/2 and its advantages.

---

**Remember**: HTTP is the foundation of web communication. Master these patterns, and you'll be unstoppable in web development! 🚀
