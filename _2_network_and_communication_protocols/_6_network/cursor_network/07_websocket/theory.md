# Topic 7: WebSocket Protocols - Real-time Communication Mastery

## WebSocket Overview

WebSocket is a computer communications protocol, providing full-duplex communication channels over a single TCP connection. It enables real-time, bidirectional communication between web browsers and servers.

### WebSocket Characteristics
- **Full-duplex**: Bidirectional communication
- **Low latency**: Minimal overhead
- **Persistent connection**: Single connection for multiple messages
- **Real-time**: Immediate data transmission
- **Cross-origin**: Works across different domains
- **Protocol upgrade**: Upgrades from HTTP to WebSocket

### WebSocket vs HTTP
| Feature | WebSocket | HTTP |
|---------|-----------|------|
| Connection | Persistent | Request-Response |
| Overhead | Low | High |
| Latency | Low | High |
| Real-time | Yes | No |
| Bidirectional | Yes | No |
| Server Push | Yes | Limited |

## WebSocket Handshake

### Client Request
```
GET /chat HTTP/1.1
Host: server.example.com
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==
Sec-WebSocket-Version: 13
```

### Server Response
```
HTTP/1.1 101 Switching Protocols
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Accept: s3pPLMBiTxaQ9kYGzzhZRbK+xOo=
```

### Handshake Process
1. **Client sends HTTP request** with WebSocket headers
2. **Server validates request** and generates accept key
3. **Server responds with 101** Switching Protocols
4. **Connection upgrades** to WebSocket protocol
5. **Bidirectional communication** begins

## WebSocket Frame Format

### Frame Structure
```
 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-------+-+-------------+-------------------------------+
|F|R|R|R| opcode|M| Payload len |    Extended payload length    |
|I|S|S|S|  (4)  |A|     (7)     |             (16/64)           |
|N|V|V|V|       |S|             |   (if payload len==126/127)   |
| |1|2|3|       |K|             |                               |
+-+-+-+-+-------+-+-------------+ - - - - - - - - - - - - - - - +
|     Extended payload length continued, if payload len == 127  |
+ - - - - - - - - - - - - - - - +-------------------------------+
|                               |Masking-key, if MASK set to 1  |
+-------------------------------+-------------------------------+
| Masking-key (continued)       |          Payload Data         |
+-------------------------------- - - - - - - - - - - - - - - - +
:                     Payload Data continued ...                :
+ - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - +
|                     Payload Data continued ...                |
+---------------------------------------------------------------+
```

### Frame Fields
- **FIN (1 bit)**: Indicates if this is the final fragment
- **RSV (3 bits)**: Reserved for extensions
- **Opcode (4 bits)**: Frame type (text, binary, close, ping, pong)
- **MASK (1 bit)**: Indicates if payload is masked
- **Payload Length (7 bits)**: Length of payload data
- **Masking Key (32 bits)**: Key for unmasking payload
- **Payload Data**: Actual message data

## WebSocket Opcodes

### Data Frames
- **0x0**: Continuation frame
- **0x1**: Text frame
- **0x2**: Binary frame

### Control Frames
- **0x8**: Close frame
- **0x9**: Ping frame
- **0xA**: Pong frame

### Frame Types
- **Text Frame**: UTF-8 encoded text data
- **Binary Frame**: Binary data
- **Close Frame**: Connection termination
- **Ping Frame**: Keep-alive or latency check
- **Pong Frame**: Response to ping frame

## WebSocket States

### Connection States
- **CONNECTING**: Handshake in progress
- **OPEN**: Connection established
- **CLOSING**: Close handshake in progress
- **CLOSED**: Connection closed

### State Transitions
```
CONNECTING → OPEN → CLOSING → CLOSED
     ↓         ↓        ↓
   Error    Error   Error
```

## WebSocket Subprotocols

### Purpose
- **Application-specific**: Define message format
- **Versioning**: Handle protocol versions
- **Negotiation**: Client and server agreement
- **Extensibility**: Add new features

### Common Subprotocols
- **chat**: Simple text messaging
- **soap**: SOAP over WebSocket
- **wamp**: WebSocket Application Messaging Protocol
- **mqtt**: MQTT over WebSocket
- **stomp**: Simple Text Oriented Messaging Protocol

## WebSocket Extensions

### Per-Message Compression
- **Deflate**: Compress message payload
- **Reduced overhead**: Smaller message size
- **Negotiated**: Client and server agreement
- **Optional**: Not required for basic operation

### Other Extensions
- **Custom extensions**: Application-specific
- **Protocol extensions**: Additional features
- **Security extensions**: Enhanced security
- **Performance extensions**: Optimization features

## Go WebSocket Implementation

### gorilla/websocket Package
- **Most popular**: Widely used WebSocket library
- **RFC 6455 compliant**: Full protocol support
- **Easy to use**: Simple API
- **Well maintained**: Active development

### Basic WebSocket Server
```go
func startWebSocketServer() {
    http.HandleFunc("/ws", handleWebSocket)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
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

## Advanced WebSocket Concepts

### Connection Management
- **Connection pooling**: Reuse connections
- **Connection limits**: Prevent resource exhaustion
- **Connection monitoring**: Track active connections
- **Graceful shutdown**: Clean connection termination

### Message Handling
- **Message queuing**: Buffer messages
- **Message routing**: Route to specific clients
- **Message broadcasting**: Send to all clients
- **Message filtering**: Selective message delivery

### Error Handling
- **Connection errors**: Handle disconnections
- **Message errors**: Handle malformed messages
- **Protocol errors**: Handle protocol violations
- **Recovery**: Automatic reconnection

## WebSocket Security

### Authentication
- **Token-based**: JWT tokens
- **Session-based**: Server sessions
- **Certificate-based**: Client certificates
- **API key**: Simple authentication

### Authorization
- **Role-based**: User roles and permissions
- **Resource-based**: Access to specific resources
- **Time-based**: Temporary access
- **Rate limiting**: Prevent abuse

### Data Protection
- **Encryption**: WSS (WebSocket Secure)
- **Message validation**: Input sanitization
- **Origin checking**: Validate request origin
- **CSRF protection**: Cross-site request forgery

## WebSocket Performance

### Optimization Techniques
- **Message batching**: Combine multiple messages
- **Compression**: Reduce message size
- **Connection pooling**: Reuse connections
- **Load balancing**: Distribute load

### Monitoring
- **Connection count**: Active connections
- **Message rate**: Messages per second
- **Latency**: Round-trip time
- **Error rate**: Failed connections/messages

### Scaling
- **Horizontal scaling**: Multiple servers
- **Load balancing**: Distribute connections
- **Message queuing**: Redis, RabbitMQ
- **Database sharding**: Distribute data

## Real-world Applications

### Chat Applications
- **Real-time messaging**: Instant communication
- **Group chats**: Multiple participants
- **File sharing**: Binary data transfer
- **Typing indicators**: Real-time status

### Gaming
- **Multiplayer games**: Real-time interaction
- **Game state sync**: Synchronize game state
- **Player movement**: Real-time position updates
- **Chat systems**: In-game communication

### Financial Trading
- **Real-time quotes**: Live market data
- **Order updates**: Trade notifications
- **Price alerts**: Market notifications
- **Trading interfaces**: Real-time trading

### IoT Applications
- **Sensor data**: Real-time monitoring
- **Device control**: Remote control
- **Status updates**: Device status
- **Alerts**: Critical notifications

## WebSocket vs Alternatives

### WebSocket vs Server-Sent Events (SSE)
- **WebSocket**: Bidirectional, complex setup
- **SSE**: Server-to-client only, simple setup
- **Use WebSocket**: When bidirectional communication needed
- **Use SSE**: When only server push needed

### WebSocket vs Long Polling
- **WebSocket**: Persistent connection, low latency
- **Long Polling**: Request-response, higher latency
- **Use WebSocket**: For real-time applications
- **Use Long Polling**: For simple server push

### WebSocket vs WebRTC
- **WebSocket**: Client-server communication
- **WebRTC**: Peer-to-peer communication
- **Use WebSocket**: For server-mediated communication
- **Use WebRTC**: For direct peer communication

## WebSocket Best Practices

### Connection Management
- **Always close connections**: Prevent resource leaks
- **Handle disconnections**: Implement reconnection logic
- **Set timeouts**: Prevent hanging connections
- **Monitor connections**: Track connection health

### Message Handling
- **Validate messages**: Check message format
- **Handle errors gracefully**: Don't crash on errors
- **Implement backpressure**: Control message flow
- **Use appropriate opcodes**: Text vs binary

### Security
- **Use WSS**: Encrypt WebSocket connections
- **Validate origins**: Check request origin
- **Implement authentication**: Verify user identity
- **Rate limiting**: Prevent abuse

### Performance
- **Batch messages**: Combine multiple messages
- **Use compression**: Reduce message size
- **Implement caching**: Cache frequently accessed data
- **Monitor performance**: Track metrics

## Common Issues and Solutions

### Connection Issues
- **Connection refused**: Server not running
- **Connection timeout**: Network issues
- **Connection reset**: Server restart
- **Solution**: Implement reconnection logic

### Message Issues
- **Message loss**: Network problems
- **Message duplication**: Retry logic
- **Message ordering**: Sequence numbers
- **Solution**: Implement message acknowledgments

### Performance Issues
- **High latency**: Network congestion
- **Low throughput**: Message size
- **Memory leaks**: Connection leaks
- **Solution**: Optimize and monitor

## Interview Questions

### Basic Questions
1. What is WebSocket and how does it work?
2. Explain the WebSocket handshake process.
3. What are the advantages of WebSocket over HTTP?

### Intermediate Questions
1. How do you implement authentication in WebSocket?
2. Explain WebSocket frame format and opcodes.
3. How do you handle WebSocket connection errors?

### Advanced Questions
1. Design a real-time chat application using WebSocket.
2. How would you scale WebSocket connections?
3. Implement WebSocket load balancing.

## Next Steps
After mastering WebSocket, proceed to:
- **Topic 8**: Network Security & Encryption
- **Topic 9**: Routing & Switching
- **Topic 10**: Network Performance & Optimization

Master WebSocket, and you'll understand real-time communication! 🚀
