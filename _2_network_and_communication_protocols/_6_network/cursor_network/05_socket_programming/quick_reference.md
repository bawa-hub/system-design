# Socket Programming Quick Reference Guide

## 🚀 Essential Socket Programming Patterns

### 1. Basic TCP Server
```go
func startTCPServer() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()
    
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        go handleConnection(conn)
    }
}
```

### 2. Basic TCP Client
```go
func startTCPClient() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    
    conn.Write([]byte("Hello, Server!"))
    
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(buffer[:n]))
}
```

### 3. Basic UDP Server
```go
func startUDPServer() {
    addr, err := net.ResolveUDPAddr("udp", ":8080")
    if err != nil {
        log.Fatal(err)
    }
    
    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    
    buffer := make([]byte, 1024)
    for {
        n, clientAddr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            log.Println(err)
            continue
        }
        go handleUDPPacket(conn, clientAddr, buffer[:n])
    }
}
```

### 4. Basic UDP Client
```go
func startUDPClient() {
    conn, err := net.Dial("udp", "localhost:8080")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    
    conn.Write([]byte("Hello, UDP Server!"))
    
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(buffer[:n]))
}
```

## 🔧 Advanced Patterns

### 1. Connection Pool
```go
type ConnectionPool struct {
    connections chan net.Conn
    maxConns    int
}

func (cp *ConnectionPool) GetConnection() (net.Conn, error) {
    select {
    case conn := <-cp.connections:
        return conn, nil
    default:
        return net.Dial("tcp", "localhost:8080")
    }
}

func (cp *ConnectionPool) ReturnConnection(conn net.Conn) {
    select {
    case cp.connections <- conn:
    default:
        conn.Close()
    }
}
```

### 2. Worker Pool
```go
type WorkerPool struct {
    workers  int
    jobQueue chan func()
    quit     chan bool
}

func (wp *WorkerPool) Start() {
    for i := 0; i < wp.workers; i++ {
        go wp.worker()
    }
}

func (wp *WorkerPool) worker() {
    for {
        select {
        case job := <-wp.jobQueue:
            job()
        case <-wp.quit:
            return
        }
    }
}
```

### 3. Context with Timeout
```go
func connectWithTimeout(address string, timeout time.Duration) (net.Conn, error) {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()
    
    return net.DialContext(ctx, "tcp", address)
}
```

## 🛠️ Socket Options

### TCP Options
```go
// Set keep-alive
tcpConn.SetKeepAlive(true)

// Disable Nagle's algorithm
tcpConn.SetNoDelay(true)

// Set read/write timeouts
conn.SetReadDeadline(time.Now().Add(30 * time.Second))
conn.SetWriteDeadline(time.Now().Add(30 * time.Second))
```

### Buffer Sizes
```go
// Set send buffer size
conn.SetWriteBuffer(8192)

// Set receive buffer size
conn.SetReadBuffer(8192)
```

## 📊 Monitoring & Debugging

### Connection Statistics
```go
type ConnectionStats struct {
    Connections int64
    BytesRead   int64
    BytesWritten int64
    Errors      int64
}

func (cs *ConnectionStats) IncrementConnections() {
    atomic.AddInt64(&cs.Connections, 1)
}
```

### Error Handling
```go
func handleConnection(conn net.Conn) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Connection panic: %v", r)
        }
        conn.Close()
    }()
    
    // Handle connection
}
```

## 🔒 Security Patterns

### TLS Connection
```go
func createTLSConnection(address string) (net.Conn, error) {
    conn, err := tls.Dial("tcp", address, &tls.Config{
        InsecureSkipVerify: false,
    })
    return conn, err
}
```

### Input Validation
```go
func validateInput(data []byte) error {
    if len(data) > 1024 {
        return fmt.Errorf("data too large")
    }
    if len(data) == 0 {
        return fmt.Errorf("empty data")
    }
    return nil
}
```

## ⚡ Performance Tips

### 1. Use Buffered I/O
```go
reader := bufio.NewReader(conn)
writer := bufio.NewWriter(conn)
```

### 2. Reuse Connections
```go
// Use connection pooling
pool := NewConnectionPool("localhost:8080", 10)
```

### 3. Non-blocking Operations
```go
conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
```

### 4. Batch Operations
```go
// Send multiple messages in one write
messages := []string{"msg1", "msg2", "msg3"}
data := strings.Join(messages, "\n")
conn.Write([]byte(data))
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

### Issue: Too Many Open Files
**Solution**: Close connections properly and increase file limits
```go
defer conn.Close() // Always close connections
```

### Issue: Connection Timeout
**Solution**: Set appropriate timeouts and implement retry logic
```go
conn.SetDeadline(time.Now().Add(30 * time.Second))
```

## 📝 Best Practices

1. **Always close connections** in defer statements
2. **Use timeouts** to prevent hanging connections
3. **Implement retry logic** for transient failures
4. **Monitor connection statistics** for debugging
5. **Use connection pooling** for high-performance applications
6. **Handle errors gracefully** with proper logging
7. **Validate input** to prevent security issues
8. **Use context** for cancellation and timeouts

## 🎯 Interview Questions

### Basic
1. How do you create a TCP server in Go?
2. What's the difference between TCP and UDP sockets?
3. How do you handle multiple connections?

### Intermediate
1. How do you implement connection pooling?
2. How do you handle socket errors gracefully?
3. How do you optimize socket performance?

### Advanced
1. How do you implement a high-performance socket server?
2. How do you handle millions of concurrent connections?
3. How do you implement custom protocols over sockets?

---

**Remember**: Socket programming is about understanding the fundamentals and applying them correctly. Practice with real applications to master these concepts! 🚀
