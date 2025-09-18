# Topic 3: TCP Protocol Deep Dive

## TCP Overview

Transmission Control Protocol (TCP) is a connection-oriented, reliable transport layer protocol that provides:
- **Reliable delivery**: Guarantees data reaches its destination
- **Ordered delivery**: Data arrives in the correct sequence
- **Error detection**: Detects and corrects transmission errors
- **Flow control**: Prevents overwhelming the receiver
- **Congestion control**: Manages network congestion

## TCP Header Structure

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

### Header Fields Explained

- **Source Port (16 bits)**: Port number of the sending application
- **Destination Port (16 bits)**: Port number of the receiving application
- **Sequence Number (32 bits)**: Sequence number of the first data byte
- **Acknowledgment Number (32 bits)**: Next expected sequence number
- **Data Offset (4 bits)**: Length of TCP header in 32-bit words
- **Reserved (3 bits)**: Reserved for future use
- **Flags (9 bits)**: Control flags (URG, ACK, PSH, RST, SYN, FIN)
- **Window Size (16 bits)**: Number of bytes receiver can accept
- **Checksum (16 bits)**: Header and data checksum
- **Urgent Pointer (16 bits)**: Points to urgent data
- **Options (0-40 bytes)**: Optional TCP parameters
- **Data**: Application data

## TCP Connection States

### State Diagram
```
    CLOSED
       |
       | Passive Open
       | (LISTEN)
       v
    LISTEN
       |
       | SYN
       v
    SYN_RCVD
       |
       | SYN+ACK
       v
    SYN_SENT
       |
       | ACK
       v
    ESTABLISHED
       |
       | FIN
       v
    FIN_WAIT_1
       |
       | ACK
       v
    FIN_WAIT_2
       |
       | FIN
       v
    TIME_WAIT
       |
       | 2MSL timeout
       v
    CLOSED
```

### State Descriptions

- **CLOSED**: No connection exists
- **LISTEN**: Server waiting for connection request
- **SYN_SENT**: Client sent SYN, waiting for SYN-ACK
- **SYN_RCVD**: Server received SYN, sent SYN-ACK
- **ESTABLISHED**: Connection established, data transfer
- **FIN_WAIT_1**: First FIN sent, waiting for ACK
- **FIN_WAIT_2**: First FIN ACKed, waiting for second FIN
- **CLOSE_WAIT**: Received FIN, waiting to close
- **LAST_ACK**: Sent FIN, waiting for ACK
- **TIME_WAIT**: Waiting for delayed packets to expire

## TCP Three-Way Handshake

### Connection Establishment
1. **Client → Server**: SYN (seq=x)
2. **Server → Client**: SYN-ACK (seq=y, ack=x+1)
3. **Client → Server**: ACK (seq=x+1, ack=y+1)

### Why Three-Way Handshake?
- **Prevents duplicate connections**: Server can detect duplicate SYNs
- **Synchronizes sequence numbers**: Both sides agree on initial sequence numbers
- **Exchanges window sizes**: Both sides know each other's buffer capacity

## TCP Connection Termination

### Four-Way Handshake
1. **Client → Server**: FIN (seq=x)
2. **Server → Client**: ACK (ack=x+1)
3. **Server → Client**: FIN (seq=y)
4. **Client → Server**: ACK (ack=y+1)

### Why Four-Way Handshake?
- **Bidirectional closure**: Each side can close independently
- **Graceful shutdown**: Allows pending data to be sent
- **Prevents data loss**: Ensures all data is transmitted

## TCP Flow Control

### Sliding Window Protocol
- **Window Size**: Number of bytes receiver can accept
- **Advertised Window**: Receiver advertises available buffer space
- **Window Update**: Receiver sends window updates as data is processed

### Flow Control Mechanisms
- **Zero Window**: Receiver can advertise zero window to pause transmission
- **Window Probe**: Sender probes zero window periodically
- **Window Scaling**: Extends window size beyond 65535 bytes

## TCP Congestion Control

### Congestion Control Algorithms

#### Slow Start
- **Exponential Growth**: Window size doubles every RTT
- **Threshold**: Congestion window threshold (ssthresh)
- **Phase**: Initial phase of connection

#### Congestion Avoidance
- **Linear Growth**: Window size increases by 1 every RTT
- **Phase**: After reaching threshold
- **Goal**: Maintain network stability

#### Fast Retransmit
- **Trigger**: Three duplicate ACKs
- **Action**: Retransmit lost packet immediately
- **Benefit**: Faster recovery than timeout

#### Fast Recovery
- **Trigger**: After fast retransmit
- **Action**: Resume transmission with reduced window
- **Benefit**: Maintains throughput during recovery

### Congestion Control Phases
1. **Slow Start**: Window grows exponentially
2. **Congestion Avoidance**: Window grows linearly
3. **Fast Retransmit**: Retransmit on duplicate ACKs
4. **Fast Recovery**: Resume with reduced window

## TCP Reliability Mechanisms

### Error Detection
- **Checksum**: 16-bit checksum of header and data
- **Sequence Numbers**: Detect missing or duplicate data
- **Acknowledgment**: Confirm receipt of data

### Error Recovery
- **Retransmission**: Resend lost or corrupted data
- **Timeout**: Retransmit after timeout period
- **Duplicate ACKs**: Detect packet loss quickly

### Ordered Delivery
- **Sequence Numbers**: Ensure correct order
- **Buffering**: Hold out-of-order packets
- **Reassembly**: Deliver data in correct sequence

## TCP Options

### Common Options
- **Maximum Segment Size (MSS)**: Largest segment size
- **Window Scaling**: Extend window size beyond 65535
- **Selective Acknowledgment (SACK)**: Acknowledge non-contiguous data
- **Timestamp**: Measure round-trip time
- **No-Operation (NOP)**: Padding for alignment

### Option Format
```
+--------+--------+--------+--------+
|  Kind  | Length |        Data      |
+--------+--------+--------+--------+
```

## TCP Performance Optimization

### Nagle's Algorithm
- **Purpose**: Combines small packets
- **Benefit**: Reduces network overhead
- **Disable**: Use TCP_NODELAY for real-time apps

### Delayed ACK
- **Purpose**: Reduces ACK overhead
- **Mechanism**: Wait for data or timeout
- **Benefit**: Improves efficiency

### Window Scaling
- **Purpose**: Extend window size beyond 65535
- **Mechanism**: Scale factor in options
- **Benefit**: Better performance on high-speed networks

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

## When to Use TCP

### Use TCP when:
- **Data integrity is critical**: Financial transactions, file transfers
- **Order matters**: Database operations, streaming media
- **Large data transfers**: File downloads, backups
- **Reliable delivery required**: Email, web browsing
- **Bidirectional communication**: Chat applications, remote desktop

### Examples:
- **HTTP/HTTPS**: Web browsing
- **FTP**: File transfers
- **SMTP**: Email sending
- **SSH**: Secure shell
- **Telnet**: Remote terminal access

## Common TCP Issues

### Connection Issues
- **Connection refused**: Port not listening
- **Connection timeout**: Network unreachable
- **Connection reset**: Abrupt connection termination

### Performance Issues
- **Slow data transfer**: Congestion, window size
- **High latency**: Network distance, routing
- **Packet loss**: Network congestion, errors

### Troubleshooting
- **netstat**: Check connection states
- **tcpdump**: Capture and analyze packets
- **Wireshark**: Graphical packet analysis
- **ss**: Modern socket statistics

## Interview Questions

### Basic Questions
1. What is TCP and how does it work?
2. Explain the TCP three-way handshake.
3. What are the differences between TCP and UDP?

### Intermediate Questions
1. How does TCP ensure reliable delivery?
2. Explain TCP flow control and congestion control.
3. What are the TCP connection states?

### Advanced Questions
1. How would you optimize TCP performance?
2. Explain TCP options and their purposes.
3. How does TCP handle packet loss and recovery?

## Next Steps
After mastering TCP, proceed to:
- **Topic 4**: UDP Protocol Deep Dive
- **Topic 5**: Socket Programming in Go
- **Topic 6**: HTTP/HTTPS Implementation

Master TCP, and you'll understand the foundation of reliable networking! 🚀
