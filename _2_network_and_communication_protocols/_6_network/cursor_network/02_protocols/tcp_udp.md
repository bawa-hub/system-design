# TCP and UDP Protocols - Transport Layer Mastery

## Transport Layer Overview

The Transport Layer (Layer 4) provides end-to-end communication services between applications. It's responsible for:
- Process-to-process communication
- Segmentation and reassembly
- Error detection and recovery
- Flow control and congestion control

## TCP (Transmission Control Protocol)

### TCP Characteristics
- **Connection-oriented**: Establishes connection before data transfer
- **Reliable**: Guarantees delivery and order
- **Full-duplex**: Bidirectional communication
- **Flow control**: Prevents overwhelming receiver
- **Congestion control**: Manages network congestion

### TCP Header Structure
```
 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|          Source Port          |       Destination Port        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Sequence Number                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                    Acknowledgment Number                      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|  Data |           |U|A|P|R|S|F|                               |
| Offset| Reserved  |R|C|S|S|Y|I|            Window             |
|       |           |G|K|H|T|N|N|                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|           Checksum            |         Urgent Pointer        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                    Options                    |    Padding    |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             data                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```

### TCP Three-Way Handshake
1. **SYN**: Client sends SYN packet with initial sequence number
2. **SYN-ACK**: Server responds with SYN-ACK and its sequence number
3. **ACK**: Client sends ACK to complete handshake

### TCP Connection Termination
1. **FIN**: One side sends FIN to close connection
2. **ACK**: Other side acknowledges FIN
3. **FIN**: Other side sends its own FIN
4. **ACK**: First side acknowledges second FIN

### TCP Flow Control
- **Sliding Window**: Receiver advertises available buffer space
- **Window Size**: Number of bytes receiver can accept
- **Zero Window**: Receiver can advertise zero window to pause transmission

### TCP Congestion Control
- **Slow Start**: Exponential increase in window size
- **Congestion Avoidance**: Linear increase after threshold
- **Fast Retransmit**: Retransmit after 3 duplicate ACKs
- **Fast Recovery**: Resume transmission after fast retransmit

### TCP States
- **CLOSED**: No connection
- **LISTEN**: Server waiting for connection
- **SYN_SENT**: Client sent SYN, waiting for SYN-ACK
- **SYN_RCVD**: Server received SYN, sent SYN-ACK
- **ESTABLISHED**: Connection established
- **FIN_WAIT_1**: First FIN sent
- **FIN_WAIT_2**: First FIN ACKed
- **CLOSE_WAIT**: Received FIN, waiting to close
- **LAST_ACK**: Sent FIN, waiting for ACK
- **TIME_WAIT**: Waiting for delayed packets

## UDP (User Datagram Protocol)

### UDP Characteristics
- **Connectionless**: No connection establishment
- **Unreliable**: No delivery guarantees
- **Lightweight**: Minimal overhead
- **Fast**: No connection management overhead

### UDP Header Structure
```
 0      7 8     15 16    23 24    31
+--------+--------+--------+--------+
|     Source      |   Destination   |
|      Port       |      Port       |
+--------+--------+--------+--------+
|                 |                 |
|     Length      |    Checksum     |
+--------+--------+--------+--------+
|                                   |
|          data octets ...          |
+-----------------------------------+
```

### UDP Use Cases
- **DNS**: Domain name resolution
- **DHCP**: Dynamic IP configuration
- **SNMP**: Network management
- **Real-time applications**: Video, audio streaming
- **Gaming**: Fast-paced multiplayer games

## TCP vs UDP Comparison

| Feature | TCP | UDP |
|---------|-----|-----|
| Connection | Connection-oriented | Connectionless |
| Reliability | Reliable | Unreliable |
| Ordering | Ordered | Unordered |
| Speed | Slower | Faster |
| Overhead | High | Low |
| Flow Control | Yes | No |
| Congestion Control | Yes | No |
| Error Detection | Yes | Yes |
| Error Recovery | Yes | No |

## When to Use TCP vs UDP

### Use TCP when:
- Data integrity is critical
- Order of data matters
- Large amounts of data
- Reliable delivery required
- Examples: HTTP, HTTPS, FTP, SMTP

### Use UDP when:
- Speed is more important than reliability
- Real-time applications
- Small, frequent messages
- Broadcasting/multicasting
- Examples: DNS, DHCP, SNMP, gaming, streaming

## Go Implementation Examples

### TCP Server
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
        go handleTCPConnection(conn)
    }
}
```

### TCP Client
```go
func startTCPClient() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    
    // Send data
    conn.Write([]byte("Hello, TCP!"))
    
    // Read response
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(buffer[:n]))
}
```

### UDP Server
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

### UDP Client
```go
func startUDPClient() {
    conn, err := net.Dial("udp", "localhost:8080")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    
    // Send data
    conn.Write([]byte("Hello, UDP!"))
    
    // Read response
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(buffer[:n]))
}
```

## Advanced TCP Concepts

### TCP Keep-Alive
- Detects dead connections
- Sends keep-alive probes
- Configurable intervals and retries

### TCP Nagle's Algorithm
- Combines small packets
- Reduces network overhead
- Can be disabled for real-time applications

### TCP Window Scaling
- Extends window size beyond 65535 bytes
- Uses window scale option
- Essential for high-speed networks

### TCP Selective Acknowledgment (SACK)
- Acknowledges non-contiguous data
- Improves performance with packet loss
- Reduces unnecessary retransmissions

## Performance Considerations

### TCP Optimization
- Tune buffer sizes
- Optimize window sizes
- Use TCP_NODELAY for real-time apps
- Implement connection pooling

### UDP Optimization
- Implement application-level reliability
- Use appropriate packet sizes
- Handle packet loss gracefully
- Implement rate limiting

## Common Issues and Troubleshooting

### TCP Issues
- Connection timeouts
- Slow data transfer
- Connection resets
- Port exhaustion

### UDP Issues
- Packet loss
- Out-of-order delivery
- No flow control
- Firewall blocking

## Interview Questions

### Basic Questions
1. What's the difference between TCP and UDP?
2. Explain the TCP three-way handshake.
3. When would you use UDP over TCP?

### Intermediate Questions
1. How does TCP flow control work?
2. What is TCP congestion control?
3. Explain TCP connection states.

### Advanced Questions
1. How would you implement reliable UDP?
2. Design a high-performance TCP server.
3. Troubleshoot TCP connection issues.

## Next Steps
Now that you understand TCP and UDP, we'll explore:
1. Network programming in Go
2. HTTP/HTTPS implementation
3. WebSocket protocols
4. Advanced networking concepts

Master these protocols, and you'll be unstoppable in networking interviews! 🚀
