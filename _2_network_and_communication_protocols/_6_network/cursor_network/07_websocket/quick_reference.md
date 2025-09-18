# WebSocket Quick Reference Guide

## 🚀 Essential WebSocket Concepts

### WebSocket vs HTTP
- **WebSocket**: Persistent, bidirectional, real-time
- **HTTP**: Request-response, stateless, higher latency
- **Use WebSocket**: For real-time applications
- **Use HTTP**: For simple request-response

### WebSocket Handshake
```
Client Request:
GET /ws HTTP/1.1
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==
Sec-WebSocket-Version: 13

Server Response:
HTTP/1.1 101 Switching Protocols
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Accept: s3pPLMBiTxaQ9kYGzzhZRbK+xOo=
```

## 🔧 Go WebSocket Patterns

### Basic WebSocket Server
```go
func startWebSocketServer() {
    upgrader := websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
        CheckOrigin: func(r *http.Request) bool {
            return true // Allow all origins
        },
    }
    
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Println("Upgrade error:", err)
            return
        }
        defer conn.Close()
        
        for {
            messageType, message, err := conn.ReadMessage()
            if err != nil {
                log.Println("Read error:", err)
                break
            }
            
            // Echo the message back
            conn.WriteMessage(messageType, message)
        }
    })
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Basic WebSocket Client
```go
func startWebSocketClient() {
    conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
    if err != nil {
        log.Fatal("Dial error:", err)
    }
    defer conn.Close()
    
    // Send message
    conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
    
    // Read response
    _, message, err := conn.ReadMessage()
    if err != nil {
        log.Println("Read error:", err)
        return
    }
    fmt.Println("Received:", string(message))
}
```

### Connection Management
```go
type WebSocketServer struct {
    Connections map[*websocket.Conn]*Connection
    Mutex       sync.RWMutex
    Broadcast   chan []byte
    Register    chan *websocket.Conn
    Unregister  chan *websocket.Conn
}

func (ws *WebSocketServer) manageConnections() {
    for {
        select {
        case conn := <-ws.Register:
            ws.registerConnection(conn)
        case conn := <-ws.Unregister:
            ws.unregisterConnection(conn)
        case message := <-ws.Broadcast:
            ws.broadcastToAll(message)
        }
    }
}
```

## 📨 Message Handling

### Message Types
```go
type Message struct {
    Type      string      `json:"type"`
    Content   string      `json:"content"`
    Username  string      `json:"username"`
    Room      string      `json:"room"`
    Timestamp time.Time   `json:"timestamp"`
    Data      interface{} `json:"data,omitempty"`
}
```

### Message Broadcasting
```go
func (ws *WebSocketServer) broadcastToAll(message []byte) {
    ws.Mutex.RLock()
    for conn, connection := range ws.Connections {
        select {
        case connection.Send <- message:
        default:
            close(connection.Send)
            delete(ws.Connections, conn)
        }
    }
    ws.Mutex.RUnlock()
}
```

### Room-based Broadcasting
```go
func (ws *WebSocketServer) broadcastToRoom(room string, message []byte) {
    ws.Mutex.RLock()
    for conn, connection := range ws.Connections {
        if connection.Room == room {
            select {
            case connection.Send <- message:
            default:
                close(connection.Send)
                delete(ws.Connections, conn)
            }
        }
    }
    ws.Mutex.RUnlock()
}
```

## 🔄 Connection Lifecycle

### Connection Registration
```go
func (ws *WebSocketServer) registerConnection(conn *websocket.Conn) {
    connection := &Connection{
        Conn:      conn,
        ID:        generateConnectionID(),
        CreatedAt: time.Now(),
        Send:      make(chan []byte, 256),
    }
    
    ws.Mutex.Lock()
    ws.Connections[conn] = connection
    ws.Mutex.Unlock()
    
    go ws.writePump(connection)
    go ws.readPump(connection)
}
```

### Write Pump
```go
func (ws *WebSocketServer) writePump(connection *Connection) {
    ticker := time.NewTicker(54 * time.Second)
    defer func() {
        ticker.Stop()
        connection.Conn.Close()
    }()
    
    for {
        select {
        case message, ok := <-connection.Send:
            connection.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
            if !ok {
                connection.Conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }
            
            w, err := connection.Conn.NextWriter(websocket.TextMessage)
            if err != nil {
                return
            }
            w.Write(message)
            
            if err := w.Close(); err != nil {
                return
            }
        case <-ticker.C:
            connection.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
            if err := connection.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
                return
            }
        }
    }
}
```

### Read Pump
```go
func (ws *WebSocketServer) readPump(connection *Connection) {
    defer func() {
        ws.Unregister <- connection.Conn
        connection.Conn.Close()
    }()
    
    connection.Conn.SetReadLimit(512)
    connection.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
    connection.Conn.SetPongHandler(func(string) error {
        connection.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
        return nil
    })
    
    for {
        _, messageBytes, err := connection.Conn.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                log.Printf("WebSocket error: %v", err)
            }
            break
        }
        
        // Handle message
        ws.handleMessage(connection, messageBytes)
    }
}
```

## 🔒 Security Patterns

### Origin Checking
```go
upgrader := websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        origin := r.Header.Get("Origin")
        return origin == "https://example.com"
    },
}
```

### Authentication
```go
func authenticateConnection(r *http.Request) (string, error) {
    token := r.URL.Query().Get("token")
    if token == "" {
        return "", fmt.Errorf("no token provided")
    }
    
    // Validate token
    userID, err := validateToken(token)
    if err != nil {
        return "", err
    }
    
    return userID, nil
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

## ⚡ Performance Optimization

### Connection Pooling
```go
type ConnectionPool struct {
    connections chan *websocket.Conn
    maxConns    int
    mutex       sync.Mutex
}

func (cp *ConnectionPool) GetConnection() (*websocket.Conn, error) {
    select {
    case conn := <-cp.connections:
        return conn, nil
    default:
        return websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
    }
}

func (cp *ConnectionPool) ReturnConnection(conn *websocket.Conn) {
    select {
    case cp.connections <- conn:
    default:
        conn.Close()
    }
}
```

### Message Batching
```go
func (ws *WebSocketServer) batchMessages() {
    ticker := time.NewTicker(100 * time.Millisecond)
    defer ticker.Stop()
    
    var messages [][]byte
    
    for {
        select {
        case message := <-ws.Broadcast:
            messages = append(messages, message)
        case <-ticker.C:
            if len(messages) > 0 {
                ws.broadcastBatch(messages)
                messages = messages[:0]
            }
        }
    }
}
```

### Compression
```go
upgrader := websocket.Upgrader{
    EnableCompression: true,
    CompressionLevel:  gzip.BestSpeed,
}
```

## 🐛 Error Handling

### Connection Errors
```go
func (ws *WebSocketServer) handleConnectionError(conn *websocket.Conn, err error) {
    if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
        log.Printf("WebSocket error: %v", err)
    }
    
    // Clean up connection
    ws.Unregister <- conn
    conn.Close()
}
```

### Message Errors
```go
func (ws *WebSocketServer) handleMessageError(connection *Connection, err error) {
    log.Printf("Message error for connection %s: %v", connection.ID, err)
    
    // Send error message to client
    errorMessage := Message{
        Type:      "error",
        Content:   "Invalid message format",
        Username:  "System",
        Timestamp: time.Now(),
    }
    
    messageBytes, _ := json.Marshal(errorMessage)
    connection.Send <- messageBytes
}
```

### Reconnection Logic
```go
func (client *WebSocketClient) connectWithRetry(maxRetries int) error {
    for i := 0; i < maxRetries; i++ {
        err := client.Connect()
        if err == nil {
            return nil
        }
        
        log.Printf("Connection attempt %d failed: %v", i+1, err)
        time.Sleep(time.Duration(i+1) * time.Second)
    }
    
    return fmt.Errorf("failed to connect after %d attempts", maxRetries)
}
```

## 📊 Monitoring

### Connection Statistics
```go
type ConnectionStats struct {
    TotalConnections int64
    ActiveConnections int64
    MessagesSent     int64
    MessagesReceived int64
    Errors           int64
    StartTime        time.Time
}

func (cs *ConnectionStats) IncrementConnections() {
    atomic.AddInt64(&cs.TotalConnections, 1)
    atomic.AddInt64(&cs.ActiveConnections, 1)
}

func (cs *ConnectionStats) DecrementConnections() {
    atomic.AddInt64(&cs.ActiveConnections, -1)
}
```

### Health Checks
```go
func (ws *WebSocketServer) healthCheck() {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            ws.Mutex.RLock()
            for conn, connection := range ws.Connections {
                if time.Since(connection.LastPing) > 60*time.Second {
                    conn.Close()
                    delete(ws.Connections, conn)
                }
            }
            ws.Mutex.RUnlock()
        }
    }
}
```

## 🎯 Best Practices

1. **Always close connections** in defer statements
2. **Handle disconnections gracefully** with reconnection logic
3. **Implement proper error handling** for all operations
4. **Use connection pooling** for high-performance applications
5. **Implement rate limiting** to prevent abuse
6. **Monitor connection health** with ping/pong
7. **Validate all messages** before processing
8. **Use appropriate message types** for different operations
9. **Implement proper authentication** and authorization
10. **Test with multiple clients** to ensure scalability

## 🚀 Common Use Cases

### Chat Applications
- Real-time messaging
- Group chats
- File sharing
- Typing indicators

### Gaming
- Multiplayer games
- Real-time updates
- Player movement
- Game state sync

### Financial Trading
- Live market data
- Order updates
- Price alerts
- Trading interfaces

### IoT Applications
- Sensor data
- Device control
- Status updates
- Alerts

---

**Remember**: WebSocket is perfect for real-time applications. Master these patterns, and you'll be unstoppable in building modern web applications! 🚀
