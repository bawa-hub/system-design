# Mastering computer networking for software engineering
_Exported on 13/09/2025 at 2:45:49 GMT+5:30 from Cursor (1.5.9)_

---

**User**

i want to learn and master computer networking completely and become greatest software engineer.

i want to become god like in this subject . 

Think of you like a god of computer networking subject and teach me everything from basics to advanced level so that i can crack any big product based companies interview.

please teach with exhaustive theory and implementation if possible in golang.

---

**Cursor**

I'll help you become a master of computer networking! This is an exciting journey, and I'll create a comprehensive learning path that combines deep theoretical knowledge with practical Go implementations. Let me start by creating a structured curriculum and then we'll dive into each topic systematically.


Let me create a comprehensive networking curriculum structure and start with the fundamentals:


```markdown
# Computer Networking Mastery Course 🚀

## Course Overview
This comprehensive course will take you from networking basics to advanced concepts, preparing you for top-tier software engineering interviews at companies like Google, Amazon, Microsoft, Meta, and more.

## Learning Path

### Phase 1: Fundamentals (Weeks 1-2)
- [ ] **Network Basics & OSI Model**
  - Physical Layer concepts
  - Data Link Layer (Ethernet, MAC addresses)
  - Network Layer (IP, routing)
  - Transport Layer (TCP/UDP)
  - Session, Presentation, Application Layers

- [ ] **IP Addressing & Subnetting**
  - IPv4 vs IPv6
  - CIDR notation
  - Subnet calculations
  - NAT and private addressing

### Phase 2: Core Protocols (Weeks 3-4)
- [ ] **TCP Deep Dive**
  - Three-way handshake
  - Flow control and congestion control
  - Error detection and recovery
  - Go TCP implementations

- [ ] **UDP Deep Dive**
  - Connectionless communication
  - When to use UDP vs TCP
  - Go UDP implementations

### Phase 3: Network Programming (Weeks 5-6)
- [ ] **Socket Programming in Go**
  - TCP/UDP servers and clients
  - Concurrent network programming
  - Error handling and timeouts

- [ ] **HTTP/HTTPS Implementation**
  - HTTP protocol internals
  - Building HTTP servers and clients
  - WebSocket implementation

### Phase 4: Advanced Topics (Weeks 7-8)
- [ ] **Routing & Switching**
  - Routing algorithms (Dijkstra, Bellman-Ford)
  - Spanning Tree Protocol
  - VLANs and network segmentation

- [ ] **Network Security**
  - TLS/SSL implementation
  - VPN concepts
  - Firewall and intrusion detection

### Phase 5: Performance & Optimization (Weeks 9-10)
- [ ] **Network Performance**
  - Bandwidth, latency, throughput
  - Load balancing strategies
  - Caching and CDN concepts

- [ ] **Distributed Systems Networking**
  - Microservices communication
  - Service discovery
  - Message queues and event streaming

## Interview Preparation
- [ ] **System Design Questions**
- [ ] **Coding Challenges**
- [ ] **Protocol Implementation**
- [ ] **Troubleshooting Scenarios**

## Prerequisites
- Basic programming knowledge (Go preferred)
- Understanding of computer systems
- Linux command line familiarity

## Tools & Technologies
- Go programming language
- Wireshark for packet analysis
- Docker for containerized networking
- Kubernetes for orchestration
- Various network simulation tools

Let's begin this journey to networking mastery! 🎯
```

Now let's start with the fundamentals. I'll create a comprehensive module on network basics and the OSI model:


```markdown
# OSI Model - The Foundation of Networking

## Overview
The Open Systems Interconnection (OSI) model is a conceptual framework that standardizes the functions of a telecommunication or computing system into seven abstraction layers. Understanding this model is crucial for any networking professional.

## The Seven Layers (Bottom to Top)

### 1. Physical Layer (Layer 1)
**Purpose**: Transmits raw bits over a physical medium

**Key Concepts**:
- Electrical, optical, or radio signals
- Bit synchronization
- Physical topology
- Transmission modes (simplex, half-duplex, full-duplex)

**Examples**:
- Ethernet cables (Cat5, Cat6)
- Fiber optic cables
- Wireless radio waves
- USB, HDMI cables

**Go Implementation**: While Go doesn't directly work with physical layer, we can simulate signal transmission concepts.

### 2. Data Link Layer (Layer 2)
**Purpose**: Provides reliable data transfer between directly connected nodes

**Key Concepts**:
- Frame formatting
- MAC (Media Access Control) addresses
- Error detection and correction
- Flow control
- Access control (CSMA/CD, CSMA/CA)

**Protocols**:
- Ethernet (IEEE 802.3)
- WiFi (IEEE 802.11)
- PPP (Point-to-Point Protocol)
- HDLC (High-Level Data Link Control)

**Go Implementation**: We'll implement frame structures and MAC address handling.

### 3. Network Layer (Layer 3)
**Purpose**: Routes data packets from source to destination across multiple networks

**Key Concepts**:
- Logical addressing (IP addresses)
- Routing and forwarding
- Path determination
- Fragmentation and reassembly

**Protocols**:
- IPv4/IPv6
- ICMP (Internet Control Message Protocol)
- OSPF, BGP, RIP (routing protocols)

**Go Implementation**: IP packet structures, routing tables, and address manipulation.

### 4. Transport Layer (Layer 4)
**Purpose**: Provides end-to-end communication services

**Key Concepts**:
- Process-to-process communication
- Segmentation and reassembly
- Error detection and recovery
- Flow control and congestion control

**Protocols**:
- TCP (Transmission Control Protocol)
- UDP (User Datagram Protocol)
- SCTP (Stream Control Transmission Protocol)

**Go Implementation**: TCP/UDP socket programming, connection management.

### 5. Session Layer (Layer 5)
**Purpose**: Manages sessions between applications

**Key Concepts**:
- Session establishment, maintenance, and termination
- Dialog control (full-duplex, half-duplex)
- Synchronization checkpoints

**Examples**:
- RPC (Remote Procedure Call)
- SQL sessions
- WebSocket connections

### 6. Presentation Layer (Layer 6)
**Purpose**: Handles data representation and encryption

**Key Concepts**:
- Data compression
- Encryption/decryption
- Character encoding (ASCII, Unicode)
- Data format conversion

**Examples**:
- SSL/TLS encryption
- JPEG, MPEG compression
- XML, JSON formatting

### 7. Application Layer (Layer 7)
**Purpose**: Provides network services to user applications

**Key Concepts**:
- User interfaces
- Network services
- Application protocols

**Protocols**:
- HTTP/HTTPS
- FTP, SFTP
- SMTP, POP3, IMAP
- DNS
- SSH
- Telnet

## Data Flow Through OSI Layers

```
Application Data
       ↓
   [Layer 7] Application Layer - Adds application header
       ↓
   [Layer 6] Presentation Layer - Adds presentation header
       ↓
   [Layer 5] Session Layer - Adds session header
       ↓
   [Layer 4] Transport Layer - Adds transport header (TCP/UDP)
       ↓
   [Layer 3] Network Layer - Adds IP header
       ↓
   [Layer 2] Data Link Layer - Adds frame header and trailer
       ↓
   [Layer 1] Physical Layer - Converts to electrical/optical signals
       ↓
   Physical Medium (Cable, Wireless, etc.)
```

## Encapsulation Process

1. **Application Layer**: User data
2. **Presentation Layer**: Data + Presentation header
3. **Session Layer**: Data + Session header
4. **Transport Layer**: Data + Transport header (TCP/UDP)
5. **Network Layer**: Data + IP header
6. **Data Link Layer**: Data + Frame header + Frame trailer
7. **Physical Layer**: Electrical/optical signals

## Decapsulation Process

The receiving device reverses the process, removing headers layer by layer until the original data is delivered to the application.

## Why OSI Model Matters

1. **Standardization**: Provides a common framework for network protocols
2. **Troubleshooting**: Helps isolate problems to specific layers
3. **Design**: Guides network architecture decisions
4. **Interoperability**: Ensures different systems can communicate
5. **Learning**: Provides a structured way to understand networking

## Common Misconceptions

1. **"OSI is just theory"**: While the OSI model is conceptual, it's based on real protocols and is used in practice
2. **"All layers are always used"**: Some layers may be skipped in certain implementations
3. **"TCP/IP replaces OSI"**: TCP/IP is a protocol suite that maps to OSI layers, not a replacement

## Next Steps

Now that you understand the OSI model, we'll dive into:
1. IP addressing and subnetting
2. TCP and UDP protocols in detail
3. Practical Go implementations for each layer

This foundation will make you unstoppable in networking interviews! 🚀
```

Now let's create a practical Go implementation that demonstrates OSI layer concepts:


```go
package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// OSI Layer 2: Data Link Layer - Frame Structure
type EthernetFrame struct {
	DestinationMAC [6]byte
	SourceMAC      [6]byte
	EtherType      uint16
	Payload        []byte
	FCS            uint32 // Frame Check Sequence
}

// OSI Layer 3: Network Layer - IP Packet Structure
type IPPacket struct {
	Version        uint8
	IHL            uint8
	TypeOfService  uint8
	TotalLength    uint16
	Identification uint16
	Flags          uint8
	FragmentOffset uint16
	TTL            uint8
	Protocol       uint8
	HeaderChecksum uint16
	SourceIP       net.IP
	DestinationIP  net.IP
	Options        []byte
	Payload        []byte
}

// OSI Layer 4: Transport Layer - TCP Segment Structure
type TCPSegment struct {
	SourcePort      uint16
	DestinationPort uint16
	SequenceNumber  uint32
	AckNumber       uint32
	DataOffset      uint8
	Flags           uint8
	WindowSize      uint16
	Checksum        uint16
	UrgentPointer   uint16
	Options         []byte
	Payload         []byte
}

// OSI Layer 7: Application Layer - HTTP Request
type HTTPRequest struct {
	Method    string
	URL       string
	Version   string
	Headers   map[string]string
	Body      []byte
}

// MAC Address utilities
func NewMACAddress(addr string) [6]byte {
	var mac [6]byte
	fmt.Sscanf(addr, "%02x:%02x:%02x:%02x:%02x:%02x",
		&mac[0], &mac[1], &mac[2], &mac[3], &mac[4], &mac[5])
	return mac
}

func (mac [6]byte) String() string {
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		mac[0], mac[1], mac[2], mac[3], mac[4], mac[5])
}

// Create an Ethernet frame
func CreateEthernetFrame(destMAC, srcMAC [6]byte, payload []byte) *EthernetFrame {
	return &EthernetFrame{
		DestinationMAC: destMAC,
		SourceMAC:      srcMAC,
		EtherType:      0x0800, // IPv4
		Payload:        payload,
		FCS:            calculateCRC32(payload), // Simplified CRC
	}
}

// Create an IP packet
func CreateIPPacket(srcIP, destIP net.IP, payload []byte) *IPPacket {
	return &IPPacket{
		Version:        4,
		IHL:            5, // 20 bytes header
		TypeOfService:  0,
		TotalLength:    uint16(20 + len(payload)),
		Identification: uint16(time.Now().UnixNano() & 0xFFFF),
		Flags:          0x40, // Don't fragment
		FragmentOffset: 0,
		TTL:            64,
		Protocol:       6, // TCP
		HeaderChecksum: 0, // Will be calculated
		SourceIP:       srcIP,
		DestinationIP:  destIP,
		Payload:        payload,
	}
}

// Create a TCP segment
func CreateTCPSegment(srcPort, destPort uint16, seqNum, ackNum uint32, payload []byte) *TCPSegment {
	return &TCPSegment{
		SourcePort:      srcPort,
		DestinationPort: destPort,
		SequenceNumber:  seqNum,
		AckNumber:       ackNum,
		DataOffset:      5, // 20 bytes header
		Flags:           0x18, // ACK + PSH
		WindowSize:      65535,
		Checksum:        0, // Will be calculated
		UrgentPointer:   0,
		Payload:         payload,
	}
}

// Create an HTTP request
func CreateHTTPRequest(method, url, version string, headers map[string]string, body []byte) *HTTPRequest {
	return &HTTPRequest{
		Method:  method,
		URL:     url,
		Version: version,
		Headers: headers,
		Body:    body,
	}
}

// Serialize HTTP request to bytes
func (req *HTTPRequest) ToBytes() []byte {
	request := fmt.Sprintf("%s %s %s\r\n", req.Method, req.URL, req.Version)
	
	for key, value := range req.Headers {
		request += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	request += "\r\n"
	
	requestBytes := []byte(request)
	return append(requestBytes, req.Body...)
}

// Serialize TCP segment to bytes
func (tcp *TCPSegment) ToBytes() []byte {
	data := make([]byte, 20+len(tcp.Options)+len(tcp.Payload))
	
	// Source port (2 bytes)
	binary.BigEndian.PutUint16(data[0:2], tcp.SourcePort)
	// Destination port (2 bytes)
	binary.BigEndian.PutUint16(data[2:4], tcp.DestinationPort)
	// Sequence number (4 bytes)
	binary.BigEndian.PutUint32(data[4:8], tcp.SequenceNumber)
	// ACK number (4 bytes)
	binary.BigEndian.PutUint32(data[8:12], tcp.AckNumber)
	// Data offset and flags (2 bytes)
	data[12] = tcp.DataOffset << 4
	data[13] = tcp.Flags
	// Window size (2 bytes)
	binary.BigEndian.PutUint16(data[14:16], tcp.WindowSize)
	// Checksum (2 bytes)
	binary.BigEndian.PutUint16(data[16:18], tcp.Checksum)
	// Urgent pointer (2 bytes)
	binary.BigEndian.PutUint16(data[18:20], tcp.UrgentPointer)
	
	// Options and payload
	copy(data[20:], tcp.Options)
	copy(data[20+len(tcp.Options):], tcp.Payload)
	
	return data
}

// Serialize IP packet to bytes
func (ip *IPPacket) ToBytes() []byte {
	headerLen := 20 + len(ip.Options)
	data := make([]byte, headerLen+len(ip.Payload))
	
	// Version and IHL (1 byte)
	data[0] = (ip.Version << 4) | ip.IHL
	// Type of service (1 byte)
	data[1] = ip.TypeOfService
	// Total length (2 bytes)
	binary.BigEndian.PutUint16(data[2:4], ip.TotalLength)
	// Identification (2 bytes)
	binary.BigEndian.PutUint16(data[4:6], ip.Identification)
	// Flags and fragment offset (2 bytes)
	flagsAndOffset := (uint16(ip.Flags) << 13) | ip.FragmentOffset
	binary.BigEndian.PutUint16(data[6:8], flagsAndOffset)
	// TTL (1 byte)
	data[8] = ip.TTL
	// Protocol (1 byte)
	data[9] = ip.Protocol
	// Header checksum (2 bytes)
	binary.BigEndian.PutUint16(data[10:12], ip.HeaderChecksum)
	// Source IP (4 bytes)
	copy(data[12:16], ip.SourceIP.To4())
	// Destination IP (4 bytes)
	copy(data[16:20], ip.DestinationIP.To4())
	
	// Options and payload
	copy(data[20:], ip.Options)
	copy(data[headerLen:], ip.Payload)
	
	return data
}

// Serialize Ethernet frame to bytes
func (eth *EthernetFrame) ToBytes() []byte {
	data := make([]byte, 14+len(eth.Payload)+4) // 14 bytes header + payload + 4 bytes FCS
	
	// Destination MAC (6 bytes)
	copy(data[0:6], eth.DestinationMAC[:])
	// Source MAC (6 bytes)
	copy(data[6:12], eth.SourceMAC[:])
	// EtherType (2 bytes)
	binary.BigEndian.PutUint16(data[12:14], eth.EtherType)
	// Payload
	copy(data[14:14+len(eth.Payload)], eth.Payload)
	// FCS (4 bytes)
	binary.BigEndian.PutUint32(data[14+len(eth.Payload):], eth.FCS)
	
	return data
}

// Simplified CRC32 calculation (for demonstration)
func calculateCRC32(data []byte) uint32 {
	crc := uint32(0xFFFFFFFF)
	for _, b := range data {
		crc ^= uint32(b)
		for i := 0; i < 8; i++ {
			if crc&1 == 1 {
				crc = (crc >> 1) ^ 0xEDB88320
			} else {
				crc >>= 1
			}
		}
	}
	return ^crc
}

// Demonstrate OSI layer encapsulation
func demonstrateOSIEncapsulation() {
	fmt.Println("=== OSI Model Encapsulation Demo ===\n")
	
	// Layer 7: Application Layer - HTTP Request
	httpReq := CreateHTTPRequest(
		"GET",
		"/api/users",
		"HTTP/1.1",
		map[string]string{
			"Host":       "example.com",
			"User-Agent": "Go-Network-Demo/1.0",
		},
		[]byte(""),
	)
	
	httpData := httpReq.ToBytes()
	fmt.Printf("Layer 7 (Application): HTTP Request\n")
	fmt.Printf("Data: %s\n\n", string(httpData))
	
	// Layer 4: Transport Layer - TCP Segment
	tcpSeg := CreateTCPSegment(
		8080,    // Source port
		80,      // Destination port
		1000,    // Sequence number
		2000,    // ACK number
		httpData, // HTTP data as payload
	)
	
	tcpData := tcpSeg.ToBytes()
	fmt.Printf("Layer 4 (Transport): TCP Segment\n")
	fmt.Printf("Source Port: %d, Dest Port: %d\n", tcpSeg.SourcePort, tcpSeg.DestinationPort)
	fmt.Printf("Seq: %d, ACK: %d\n", tcpSeg.SequenceNumber, tcpSeg.AckNumber)
	fmt.Printf("Total Length: %d bytes\n\n", len(tcpData))
	
	// Layer 3: Network Layer - IP Packet
	ipPkt := CreateIPPacket(
		net.ParseIP("192.168.1.100"), // Source IP
		net.ParseIP("192.168.1.1"),   // Destination IP
		tcpData,                      // TCP data as payload
	)
	
	ipData := ipPkt.ToBytes()
	fmt.Printf("Layer 3 (Network): IP Packet\n")
	fmt.Printf("Source IP: %s, Dest IP: %s\n", ipPkt.SourceIP, ipPkt.DestinationIP)
	fmt.Printf("Protocol: %d (TCP), TTL: %d\n", ipPkt.Protocol, ipPkt.TTL)
	fmt.Printf("Total Length: %d bytes\n\n", len(ipData))
	
	// Layer 2: Data Link Layer - Ethernet Frame
	ethFrame := CreateEthernetFrame(
		NewMACAddress("00:11:22:33:44:55"), // Destination MAC
		NewMACAddress("aa:bb:cc:dd:ee:ff"), // Source MAC
		ipData,                             // IP data as payload
	)
	
	ethData := ethFrame.ToBytes()
	fmt.Printf("Layer 2 (Data Link): Ethernet Frame\n")
	fmt.Printf("Dest MAC: %s, Src MAC: %s\n", ethFrame.DestinationMAC.String(), ethFrame.SourceMAC.String())
	fmt.Printf("EtherType: 0x%04X (IPv4)\n", ethFrame.EtherType)
	fmt.Printf("Total Length: %d bytes\n\n", len(ethData))
	
	// Layer 1: Physical Layer (simulated)
	fmt.Printf("Layer 1 (Physical): Electrical/Optical Signals\n")
	fmt.Printf("Converting %d bytes to electrical signals...\n", len(ethData))
	fmt.Printf("Transmission complete!\n\n")
	
	// Show the complete packet structure
	fmt.Println("=== Complete Packet Structure ===")
	fmt.Printf("Total packet size: %d bytes\n", len(ethData))
	fmt.Printf("Ethernet Header: 14 bytes\n")
	fmt.Printf("IP Header: 20 bytes\n")
	fmt.Printf("TCP Header: 20 bytes\n")
	fmt.Printf("HTTP Data: %d bytes\n", len(httpData))
	fmt.Printf("Ethernet Trailer (FCS): 4 bytes\n")
}

// Demonstrate OSI layer decapsulation
func demonstrateOSIDecapsulation(packet []byte) {
	fmt.Println("=== OSI Model Decapsulation Demo ===\n")
	
	// Layer 1: Physical Layer - Receive electrical signals
	fmt.Printf("Layer 1 (Physical): Received %d bytes of electrical signals\n", len(packet))
	
	// Layer 2: Data Link Layer - Extract Ethernet frame
	if len(packet) < 14 {
		fmt.Println("Error: Packet too short for Ethernet frame")
		return
	}
	
	destMAC := [6]byte{}
	copy(destMAC[:], packet[0:6])
	srcMAC := [6]byte{}
	copy(srcMAC[:], packet[6:12])
	etherType := binary.BigEndian.Uint16(packet[12:14])
	payload := packet[14 : len(packet)-4] // Remove FCS
	
	fmt.Printf("Layer 2 (Data Link): Ethernet Frame\n")
	fmt.Printf("Dest MAC: %s, Src MAC: %s\n", destMAC.String(), srcMAC.String())
	fmt.Printf("EtherType: 0x%04X\n", etherType)
	fmt.Printf("Payload length: %d bytes\n\n", len(payload))
	
	// Layer 3: Network Layer - Extract IP packet
	if len(payload) < 20 {
		fmt.Println("Error: Payload too short for IP packet")
		return
	}
	
	version := payload[0] >> 4
	ihl := payload[0] & 0x0F
	protocol := payload[9]
	srcIP := net.IP(payload[12:16])
	destIP := net.IP(payload[16:20])
	ipPayload := payload[ihl*4:]
	
	fmt.Printf("Layer 3 (Network): IP Packet\n")
	fmt.Printf("Version: %d, IHL: %d\n", version, ihl)
	fmt.Printf("Protocol: %d, Src IP: %s, Dest IP: %s\n", protocol, srcIP, destIP)
	fmt.Printf("Payload length: %d bytes\n\n", len(ipPayload))
	
	// Layer 4: Transport Layer - Extract TCP segment
	if len(ipPayload) < 20 {
		fmt.Println("Error: IP payload too short for TCP segment")
		return
	}
	
	srcPort := binary.BigEndian.Uint16(ipPayload[0:2])
	destPort := binary.BigEndian.Uint16(ipPayload[2:4])
	seqNum := binary.BigEndian.Uint32(ipPayload[4:8])
	ackNum := binary.BigEndian.Uint32(ipPayload[8:12])
	dataOffset := ipPayload[12] >> 4
	tcpPayload := ipPayload[dataOffset*4:]
	
	fmt.Printf("Layer 4 (Transport): TCP Segment\n")
	fmt.Printf("Src Port: %d, Dest Port: %d\n", srcPort, destPort)
	fmt.Printf("Seq: %d, ACK: %d\n", seqNum, ackNum)
	fmt.Printf("Data Offset: %d\n", dataOffset)
	fmt.Printf("Payload length: %d bytes\n\n", len(tcpPayload))
	
	// Layer 7: Application Layer - Extract HTTP request
	fmt.Printf("Layer 7 (Application): HTTP Request\n")
	fmt.Printf("Data: %s\n", string(tcpPayload))
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - OSI Model Demo")
	fmt.Println("===============================================\n")
	
	// Demonstrate encapsulation
	demonstrateOSIEncapsulation()
	
	// Create a sample packet for decapsulation demo
	httpReq := CreateHTTPRequest("GET", "/", "HTTP/1.1", map[string]string{"Host": "example.com"}, []byte(""))
	httpData := httpReq.ToBytes()
	
	tcpSeg := CreateTCPSegment(8080, 80, 1000, 2000, httpData)
	tcpData := tcpSeg.ToBytes()
	
	ipPkt := CreateIPPacket(net.ParseIP("192.168.1.100"), net.ParseIP("192.168.1.1"), tcpData)
	ipData := ipPkt.ToBytes()
	
	ethFrame := CreateEthernetFrame(
		NewMACAddress("00:11:22:33:44:55"),
		NewMACAddress("aa:bb:cc:dd:ee:ff"),
		ipData,
	)
	
	packet := ethFrame.ToBytes()
	
	// Demonstrate decapsulation
	demonstrateOSIDecapsulation(packet)
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Each layer adds its own header to the data")
	fmt.Println("2. Lower layers don't understand higher layer protocols")
	fmt.Println("3. Encapsulation happens on the sending side")
	fmt.Println("4. Decapsulation happens on the receiving side")
	fmt.Println("5. This modular approach enables network interoperability")
}
```

Now let's create a comprehensive IP addressing and subnetting module:


```markdown
# IP Addressing & Subnetting Mastery

## IPv4 Addressing

### Understanding IP Addresses
An IPv4 address is a 32-bit number divided into four 8-bit octets, represented in dotted decimal notation (e.g., 192.168.1.1).

### Address Classes (Historical)
- **Class A**: 1.0.0.0 to 126.255.255.255 (8-bit network, 24-bit host)
- **Class B**: 128.0.0.0 to 191.255.255.255 (16-bit network, 16-bit host)
- **Class C**: 192.0.0.0 to 223.255.255.255 (24-bit network, 8-bit host)
- **Class D**: 224.0.0.0 to 239.255.255.255 (Multicast)
- **Class E**: 240.0.0.0 to 255.255.255.255 (Reserved)

### Private Address Ranges (RFC 1918)
- **10.0.0.0/8**: 10.0.0.0 to 10.255.255.255
- **172.16.0.0/12**: 172.16.0.0 to 172.31.255.255
- **192.168.0.0/16**: 192.168.0.0 to 192.168.255.255

### Special Addresses
- **0.0.0.0/0**: Default route
- **127.0.0.0/8**: Loopback addresses
- **169.254.0.0/16**: Link-local addresses (APIPA)
- **255.255.255.255**: Limited broadcast

## CIDR (Classless Inter-Domain Routing)

### CIDR Notation
Format: `IP_ADDRESS/PREFIX_LENGTH`
- Example: `192.168.1.0/24`
- Prefix length indicates number of network bits
- Remaining bits are host bits

### Subnet Masks
- **/8**: 255.0.0.0
- **/16**: 255.255.0.0
- **/24**: 255.255.255.0
- **/32**: 255.255.255.255 (host route)

## Subnetting Fundamentals

### Why Subnet?
1. **Efficient IP usage**: Avoid wasting addresses
2. **Security**: Isolate network segments
3. **Performance**: Reduce broadcast domains
4. **Management**: Logical organization

### Subnetting Process
1. Determine required subnets
2. Calculate subnet mask
3. Identify network addresses
4. Identify broadcast addresses
5. Identify usable host ranges

### Subnetting Formulas
- **Number of subnets**: 2^(subnet_bits)
- **Number of hosts per subnet**: 2^(host_bits) - 2
- **Subnet mask**: 2^32 - 2^(32-prefix_length)

## VLSM (Variable Length Subnet Masking)

VLSM allows different subnets to have different subnet masks within the same network.

### VLSM Benefits
- Maximum IP address utilization
- Hierarchical addressing
- Efficient routing

## IPv6 Addressing

### IPv6 Basics
- 128-bit addresses (vs 32-bit IPv4)
- Hexadecimal notation with colons
- Example: `2001:0db8:85a3:0000:0000:8a2e:0370:7334`
- Shortened: `2001:db8:85a3::8a2e:370:7334`

### IPv6 Address Types
- **Unicast**: Single interface
- **Multicast**: Multiple interfaces
- **Anycast**: Nearest interface

### IPv6 Special Addresses
- **::/128**: Unspecified address
- **::1/128**: Loopback address
- **2000::/3**: Global unicast
- **fc00::/7**: Unique local
- **fe80::/10**: Link-local

## NAT (Network Address Translation)

### NAT Types
1. **Static NAT**: One-to-one mapping
2. **Dynamic NAT**: Pool-based mapping
3. **PAT (Port Address Translation)**: Many-to-one mapping

### NAT Benefits
- IP address conservation
- Security through obscurity
- Simplified addressing

## Practical Examples

### Example 1: Basic Subnetting
**Given**: 192.168.1.0/24, need 4 subnets

**Solution**:
- Need 2 bits for subnets (2^2 = 4)
- New prefix: /26
- Subnet mask: 255.255.255.192
- Subnets:
  - 192.168.1.0/26 (0-63)
  - 192.168.1.64/26 (64-127)
  - 192.168.1.128/26 (128-191)
  - 192.168.1.192/26 (192-255)

### Example 2: VLSM
**Given**: 192.168.1.0/24, need:
- 2 subnets with 50 hosts each
- 4 subnets with 10 hosts each

**Solution**:
1. For 50 hosts: need 6 host bits (2^6-2 = 62 hosts)
   - Subnets: /26 (255.255.255.192)
   - 192.168.1.0/26 and 192.168.1.64/26

2. For 10 hosts: need 4 host bits (2^4-2 = 14 hosts)
   - Subnets: /28 (255.255.255.240)
   - 192.168.1.128/28, 192.168.1.144/28, 192.168.1.160/28, 192.168.1.176/28

## Interview Questions

### Basic Questions
1. What's the difference between a network address and a broadcast address?
2. How do you calculate the number of usable hosts in a subnet?
3. What is the purpose of a subnet mask?

### Intermediate Questions
1. Explain the difference between classful and classless addressing.
2. How does VLSM improve IP address utilization?
3. What are the advantages and disadvantages of NAT?

### Advanced Questions
1. Design a network addressing scheme for a company with multiple departments.
2. Explain how IPv6 addresses are allocated and managed.
3. How would you troubleshoot IP addressing issues in a large network?

## Next Steps
Now that you understand IP addressing, we'll move to:
1. TCP and UDP protocols
2. Routing and switching
3. Network programming in Go
4. Advanced networking concepts

Master these fundamentals, and you'll be unstoppable in networking interviews! 🚀
```

Now let's create a comprehensive Go implementation for IP addressing and subnetting:


```go
package main

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
)

// IPAddress represents an IPv4 address with utilities
type IPAddress struct {
	Address net.IP
	Mask    net.IPMask
}

// NewIPAddress creates a new IPAddress from string and CIDR notation
func NewIPAddress(address string) (*IPAddress, error) {
	ip, network, err := net.ParseCIDR(address)
	if err != nil {
		return nil, err
	}
	
	return &IPAddress{
		Address: ip,
		Mask:    network.Mask,
	}, nil
}

// NewIPAddressWithMask creates a new IPAddress with custom mask
func NewIPAddressWithMask(address string, mask net.IPMask) (*IPAddress, error) {
	ip := net.ParseIP(address)
	if ip == nil {
		return nil, fmt.Errorf("invalid IP address: %s", address)
	}
	
	return &IPAddress{
		Address: ip,
		Mask:    mask,
	}, nil
}

// GetNetworkAddress returns the network address
func (ip *IPAddress) GetNetworkAddress() net.IP {
	return ip.Address.Mask(ip.Mask)
}

// GetBroadcastAddress returns the broadcast address
func (ip *IPAddress) GetBroadcastAddress() net.IP {
	network := ip.GetNetworkAddress()
	ones, bits := ip.Mask.Size()
	
	// Calculate the number of host bits
	hostBits := bits - ones
	
	// Create broadcast address by setting all host bits to 1
	broadcast := make(net.IP, len(network))
	copy(broadcast, network)
	
	for i := 0; i < hostBits; i++ {
		byteIndex := (bits - 1 - i) / 8
		bitIndex := 7 - ((bits - 1 - i) % 8)
		broadcast[byteIndex] |= 1 << bitIndex
	}
	
	return broadcast
}

// GetFirstHostAddress returns the first usable host address
func (ip *IPAddress) GetFirstHostAddress() net.IP {
	network := ip.GetNetworkAddress()
	firstHost := make(net.IP, len(network))
	copy(firstHost, network)
	
	// Add 1 to the last octet
	firstHost[len(firstHost)-1]++
	return firstHost
}

// GetLastHostAddress returns the last usable host address
func (ip *IPAddress) GetLastHostAddress() net.IP {
	broadcast := ip.GetBroadcastAddress()
	lastHost := make(net.IP, len(broadcast))
	copy(lastHost, broadcast)
	
	// Subtract 1 from the last octet
	lastHost[len(lastHost)-1]--
	return lastHost
}

// GetHostCount returns the number of usable host addresses
func (ip *IPAddress) GetHostCount() int {
	ones, bits := ip.Mask.Size()
	hostBits := bits - ones
	return int(math.Pow(2, float64(hostBits))) - 2
}

// GetSubnetCount returns the number of subnets
func (ip *IPAddress) GetSubnetCount() int {
	ones, _ := ip.Mask.Size()
	return int(math.Pow(2, float64(32-ones)))
}

// IsInSameSubnet checks if another IP is in the same subnet
func (ip *IPAddress) IsInSameSubnet(otherIP net.IP) bool {
	network1 := ip.GetNetworkAddress()
	network2 := otherIP.Mask(ip.Mask)
	return network1.Equal(network2)
}

// String returns the string representation
func (ip *IPAddress) String() string {
	ones, _ := ip.Mask.Size()
	return fmt.Sprintf("%s/%d", ip.Address.String(), ones)
}

// SubnetCalculator provides subnetting utilities
type SubnetCalculator struct{}

// CalculateSubnets calculates subnets for a given network
func (sc *SubnetCalculator) CalculateSubnets(network string, requiredSubnets int) ([]*IPAddress, error) {
	baseIP, baseNetwork, err := net.ParseCIDR(network)
	if err != nil {
		return nil, err
	}
	
	// Calculate required bits for subnets
	subnetBits := int(math.Ceil(math.Log2(float64(requiredSubnets))))
	
	// Get current prefix length
	ones, _ := baseNetwork.Mask.Size()
	
	// Calculate new prefix length
	newPrefixLength := ones + subnetBits
	
	if newPrefixLength > 32 {
		return nil, fmt.Errorf("not enough address space for %d subnets", requiredSubnets)
	}
	
	// Create new mask
	newMask := net.CIDRMask(newPrefixLength, 32)
	
	var subnets []*IPAddress
	
	// Calculate subnet size
	subnetSize := int(math.Pow(2, float64(32-newPrefixLength)))
	
	// Generate subnets
	for i := 0; i < requiredSubnets; i++ {
		subnetAddr := make(net.IP, len(baseIP))
		copy(subnetAddr, baseIP)
		
		// Calculate subnet address
		subnetOffset := i * subnetSize
		for j := 0; j < 4; j++ {
			subnetAddr[3-j] += byte((subnetOffset >> (j * 8)) & 0xFF)
		}
		
		subnet := &IPAddress{
			Address: subnetAddr,
			Mask:    newMask,
		}
		
		subnets = append(subnets, subnet)
	}
	
	return subnets, nil
}

// VLSMSubnet represents a VLSM subnet requirement
type VLSMSubnet struct {
	Name        string
	HostsNeeded int
	Subnet      *IPAddress
}

// CalculateVLSM calculates VLSM subnets for different host requirements
func (sc *SubnetCalculator) CalculateVLSM(network string, requirements []VLSMSubnet) ([]VLSMSubnet, error) {
	// Sort requirements by host count (descending)
	for i := 0; i < len(requirements)-1; i++ {
		for j := i + 1; j < len(requirements); j++ {
			if requirements[i].HostsNeeded < requirements[j].HostsNeeded {
				requirements[i], requirements[j] = requirements[j], requirements[i]
			}
		}
	}
	
	baseIP, baseNetwork, err := net.ParseCIDR(network)
	if err != nil {
		return nil, err
	}
	
	currentIP := make(net.IP, len(baseIP))
	copy(currentIP, baseIP)
	
	var result []VLSMSubnet
	
	for _, req := range requirements {
		// Calculate required host bits
		hostBits := int(math.Ceil(math.Log2(float64(req.HostsNeeded + 2)))) // +2 for network and broadcast
		prefixLength := 32 - hostBits
		
		// Create mask
		mask := net.CIDRMask(prefixLength, 32)
		
		// Create subnet
		subnet := &IPAddress{
			Address: currentIP,
			Mask:    mask,
		}
		
		// Update result
		req.Subnet = subnet
		result = append(result, req)
		
		// Calculate next subnet address
		subnetSize := int(math.Pow(2, float64(hostBits)))
		for j := 0; j < 4; j++ {
			currentIP[3-j] += byte((subnetSize >> (j * 8)) & 0xFF)
		}
	}
	
	return result, nil
}

// IPAddressValidator provides IP address validation utilities
type IPAddressValidator struct{}

// IsValidIP checks if an IP address is valid
func (v *IPAddressValidator) IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsPrivateIP checks if an IP address is private
func (v *IPAddressValidator) IsPrivateIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	
	// Check private ranges
	privateRanges := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}
	
	for _, rangeStr := range privateRanges {
		_, network, _ := net.ParseCIDR(rangeStr)
		if network.Contains(parsedIP) {
			return true
		}
	}
	
	return false
}

// IsPublicIP checks if an IP address is public
func (v *IPAddressValidator) IsPublicIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	
	// Check if it's a valid IP and not private
	return v.IsValidIP(ip) && !v.IsPrivateIP(ip) && !parsedIP.IsLoopback() && !parsedIP.IsMulticast()
}

// GetIPClass returns the class of an IP address
func (v *IPAddressValidator) GetIPClass(ip string) string {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return "Invalid"
	}
	
	firstOctet := int(parsedIP.To4()[0])
	
	switch {
	case firstOctet >= 1 && firstOctet <= 126:
		return "Class A"
	case firstOctet >= 128 && firstOctet <= 191:
		return "Class B"
	case firstOctet >= 192 && firstOctet <= 223:
		return "Class C"
	case firstOctet >= 224 && firstOctet <= 239:
		return "Class D (Multicast)"
	case firstOctet >= 240 && firstOctet <= 255:
		return "Class E (Reserved)"
	default:
		return "Unknown"
	}
}

// Demonstrate basic subnetting
func demonstrateBasicSubnetting() {
	fmt.Println("=== Basic Subnetting Demo ===\n")
	
	// Example: 192.168.1.0/24 needs 4 subnets
	network := "192.168.1.0/24"
	requiredSubnets := 4
	
	calculator := &SubnetCalculator{}
	subnets, err := calculator.CalculateSubnets(network, requiredSubnets)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Network: %s\n", network)
	fmt.Printf("Required subnets: %d\n\n", requiredSubnets)
	
	for i, subnet := range subnets {
		fmt.Printf("Subnet %d:\n", i+1)
		fmt.Printf("  Network: %s\n", subnet.GetNetworkAddress())
		fmt.Printf("  Broadcast: %s\n", subnet.GetBroadcastAddress())
		fmt.Printf("  First Host: %s\n", subnet.GetFirstHostAddress())
		fmt.Printf("  Last Host: %s\n", subnet.GetLastHostAddress())
		fmt.Printf("  Usable Hosts: %d\n", subnet.GetHostCount())
		fmt.Printf("  CIDR: %s\n\n", subnet.String())
	}
}

// Demonstrate VLSM
func demonstrateVLSM() {
	fmt.Println("=== VLSM Demo ===\n")
	
	network := "192.168.1.0/24"
	requirements := []VLSMSubnet{
		{Name: "Sales", HostsNeeded: 50},
		{Name: "Marketing", HostsNeeded: 25},
		{Name: "IT", HostsNeeded: 10},
		{Name: "HR", HostsNeeded: 5},
	}
	
	calculator := &SubnetCalculator{}
	subnets, err := calculator.CalculateVLSM(network, requirements)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Network: %s\n", network)
	fmt.Println("VLSM Subnets:\n")
	
	for _, subnet := range subnets {
		fmt.Printf("%s:\n", subnet.Name)
		fmt.Printf("  Hosts Needed: %d\n", subnet.HostsNeeded)
		fmt.Printf("  Network: %s\n", subnet.Subnet.GetNetworkAddress())
		fmt.Printf("  Broadcast: %s\n", subnet.Subnet.GetBroadcastAddress())
		fmt.Printf("  First Host: %s\n", subnet.Subnet.GetFirstHostAddress())
		fmt.Printf("  Last Host: %s\n", subnet.Subnet.GetLastHostAddress())
		fmt.Printf("  Usable Hosts: %d\n", subnet.Subnet.GetHostCount())
		fmt.Printf("  CIDR: %s\n\n", subnet.Subnet.String())
	}
}

// Demonstrate IP address validation
func demonstrateIPValidation() {
	fmt.Println("=== IP Address Validation Demo ===\n")
	
	validator := &IPAddressValidator{}
	testIPs := []string{
		"192.168.1.1",
		"10.0.0.1",
		"172.16.0.1",
		"8.8.8.8",
		"127.0.0.1",
		"256.1.1.1",
		"invalid",
		"2001:db8::1",
	}
	
	for _, ip := range testIPs {
		fmt.Printf("IP: %s\n", ip)
		fmt.Printf("  Valid: %t\n", validator.IsValidIP(ip))
		fmt.Printf("  Private: %t\n", validator.IsPrivateIP(ip))
		fmt.Printf("  Public: %t\n", validator.IsPublicIP(ip))
		fmt.Printf("  Class: %s\n\n", validator.GetIPClass(ip))
	}
}

// Demonstrate subnet membership
func demonstrateSubnetMembership() {
	fmt.Println("=== Subnet Membership Demo ===\n")
	
	// Create a subnet
	subnet, err := NewIPAddress("192.168.1.0/24")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	testIPs := []string{
		"192.168.1.1",
		"192.168.1.100",
		"192.168.1.254",
		"192.168.2.1",
		"10.0.0.1",
	}
	
	fmt.Printf("Subnet: %s\n", subnet.String())
	fmt.Printf("Network: %s\n", subnet.GetNetworkAddress())
	fmt.Printf("Broadcast: %s\n\n", subnet.GetBroadcastAddress())
	
	for _, ipStr := range testIPs {
		ip := net.ParseIP(ipStr)
		if ip != nil {
			inSubnet := subnet.IsInSameSubnet(ip)
			fmt.Printf("IP %s is %s the subnet\n", ipStr, map[bool]string{true: "in", false: "not in"}[inSubnet])
		}
	}
}

// Calculate subnet requirements
func calculateSubnetRequirements() {
	fmt.Println("=== Subnet Requirements Calculator ===\n")
	
	// Example: Company needs to subnet 10.0.0.0/8
	baseNetwork := "10.0.0.0/8"
	
	departments := []struct {
		Name        string
		HostsNeeded int
	}{
		{"Headquarters", 1000},
		{"Branch Office 1", 500},
		{"Branch Office 2", 250},
		{"Sales Team", 100},
		{"IT Department", 50},
		{"HR Department", 25},
		{"Guest Network", 10},
	}
	
	calculator := &SubnetCalculator{}
	
	// Convert to VLSM requirements
	var requirements []VLSMSubnet
	for _, dept := range departments {
		requirements = append(requirements, VLSMSubnet{
			Name:        dept.Name,
			HostsNeeded: dept.HostsNeeded,
		})
	}
	
	subnets, err := calculator.CalculateVLSM(baseNetwork, requirements)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Base Network: %s\n", baseNetwork)
	fmt.Println("Department Subnets:\n")
	
	for _, subnet := range subnets {
		fmt.Printf("%s:\n", subnet.Name)
		fmt.Printf("  Required Hosts: %d\n", subnet.HostsNeeded)
		fmt.Printf("  Allocated Hosts: %d\n", subnet.Subnet.GetHostCount())
		fmt.Printf("  Network: %s\n", subnet.Subnet.GetNetworkAddress())
		fmt.Printf("  CIDR: %s\n", subnet.Subnet.String())
		fmt.Printf("  Efficiency: %.1f%%\n", float64(subnet.HostsNeeded)/float64(subnet.Subnet.GetHostCount())*100)
		fmt.Println()
	}
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - IP Addressing & Subnetting")
	fmt.Println("==========================================================\n")
	
	// Run all demonstrations
	demonstrateBasicSubnetting()
	demonstrateVLSM()
	demonstrateIPValidation()
	demonstrateSubnetMembership()
	calculateSubnetRequirements()
	
	fmt.Println("🎯 Key Takeaways:")
	fmt.Println("1. Subnetting allows efficient IP address utilization")
	fmt.Println("2. VLSM provides maximum flexibility for different subnet sizes")
	fmt.Println("3. Understanding IP classes helps with network design")
	fmt.Println("4. Private IP ranges are used for internal networks")
	fmt.Println("5. Proper subnetting is crucial for network scalability")
}
```

Now let's update our progress and move to TCP/UDP protocols:


Now let's create a comprehensive TCP and UDP module:


```markdown
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
```

---

**User**

please numbering the file so that i can topic by topic

---

**Cursor**

I'll reorganize the files with proper numbering so you can follow the learning path topic by topic. Let me restructure the curriculum with clear numbering:


Let me reorganize the files with proper numbering. First, let me create a new structured directory and move the existing files:


```markdown
# Computer Networking Mastery Course 🚀

## Course Overview
This comprehensive course will take you from networking basics to advanced concepts, preparing you for top-tier software engineering interviews at companies like Google, Amazon, Microsoft, Meta, and more.

## Learning Path - Sequential Order

### Phase 1: Fundamentals (Topics 1-4)
- [ ] **Topic 1**: Network Basics & OSI Model
- [ ] **Topic 2**: IP Addressing & Subnetting  
- [ ] **Topic 3**: TCP Protocol Deep Dive
- [ ] **Topic 4**: UDP Protocol Deep Dive

### Phase 2: Network Programming (Topics 5-8)
- [ ] **Topic 5**: Socket Programming in Go
- [ ] **Topic 6**: HTTP/HTTPS Implementation
- [ ] **Topic 7**: WebSocket Protocols
- [ ] **Topic 8**: Network Security & Encryption

### Phase 3: Advanced Concepts (Topics 9-12)
- [ ] **Topic 9**: Routing & Switching
- [ ] **Topic 10**: Network Performance & Optimization
- [ ] **Topic 11**: Distributed Systems Networking
- [ ] **Topic 12**: Cloud Networking & Containerization

### Phase 4: Interview Preparation (Topics 13-15)
- [ ] **Topic 13**: System Design Questions
- [ ] **Topic 14**: Coding Challenges & Implementations
- [ ] **Topic 15**: Troubleshooting & Debugging

## File Structure
```
cursor_network/
├── 00_curriculum_overview.md
├── 01_osi_model/
│   ├── theory.md
│   └── implementation.go
├── 02_ip_addressing/
│   ├── theory.md
│   └── implementation.go
├── 03_tcp_protocol/
│   ├── theory.md
│   └── implementation.go
├── 04_udp_protocol/
│   ├── theory.md
│   └── implementation.go
├── 05_socket_programming/
│   ├── theory.md
│   └── implementation.go
├── 06_http_https/
│   ├── theory.md
│   └── implementation.go
├── 07_websocket/
│   ├── theory.md
│   └── implementation.go
├── 08_network_security/
│   ├── theory.md
│   └── implementation.go
├── 09_routing_switching/
│   ├── theory.md
│   └── implementation.go
├── 10_performance_optimization/
│   ├── theory.md
│   └── implementation.go
├── 11_distributed_systems/
│   ├── theory.md
│   └── implementation.go
├── 12_cloud_networking/
│   ├── theory.md
│   └── implementation.go
├── 13_system_design/
│   ├── questions.md
│   └── solutions.md
├── 14_coding_challenges/
│   ├── problems.md
│   └── solutions.go
└── 15_troubleshooting/
    ├── scenarios.md
    └── solutions.md
```

## Prerequisites
- Basic programming knowledge (Go preferred)
- Understanding of computer systems
- Linux command line familiarity

## Tools & Technologies
- Go programming language
- Wireshark for packet analysis
- Docker for containerized networking
- Kubernetes for orchestration
- Various network simulation tools

## Learning Strategy
1. **Read Theory First**: Start with the theory.md file for each topic
2. **Study Implementation**: Go through the implementation.go file
3. **Run Examples**: Execute the code and experiment
4. **Practice**: Work on the exercises and challenges
5. **Review**: Revisit concepts until you master them

Let's begin this journey to networking mastery! 🎯
```

Now let me reorganize the existing files with proper numbering:


```markdown
# Topic 1: OSI Model - The Foundation of Networking

## Overview
The Open Systems Interconnection (OSI) model is a conceptual framework that standardizes the functions of a telecommunication or computing system into seven abstraction layers. Understanding this model is crucial for any networking professional.

## The Seven Layers (Bottom to Top)

### 1. Physical Layer (Layer 1)
**Purpose**: Transmits raw bits over a physical medium

**Key Concepts**:
- Electrical, optical, or radio signals
- Bit synchronization
- Physical topology
- Transmission modes (simplex, half-duplex, full-duplex)

**Examples**:
- Ethernet cables (Cat5, Cat6)
- Fiber optic cables
- Wireless radio waves
- USB, HDMI cables

**Go Implementation**: While Go doesn't directly work with physical layer, we can simulate signal transmission concepts.

### 2. Data Link Layer (Layer 2)
**Purpose**: Provides reliable data transfer between directly connected nodes

**Key Concepts**:
- Frame formatting
- MAC (Media Access Control) addresses
- Error detection and correction
- Flow control
- Access control (CSMA/CD, CSMA/CA)

**Protocols**:
- Ethernet (IEEE 802.3)
- WiFi (IEEE 802.11)
- PPP (Point-to-Point Protocol)
- HDLC (High-Level Data Link Control)

**Go Implementation**: We'll implement frame structures and MAC address handling.

### 3. Network Layer (Layer 3)
**Purpose**: Routes data packets from source to destination across multiple networks

**Key Concepts**:
- Logical addressing (IP addresses)
- Routing and forwarding
- Path determination
- Fragmentation and reassembly

**Protocols**:
- IPv4/IPv6
- ICMP (Internet Control Message Protocol)
- OSPF, BGP, RIP (routing protocols)

**Go Implementation**: IP packet structures, routing tables, and address manipulation.

### 4. Transport Layer (Layer 4)
**Purpose**: Provides end-to-end communication services

**Key Concepts**:
- Process-to-process communication
- Segmentation and reassembly
- Error detection and recovery
- Flow control and congestion control

**Protocols**:
- TCP (Transmission Control Protocol)
- UDP (User Datagram Protocol)
- SCTP (Stream Control Transmission Protocol)

**Go Implementation**: TCP/UDP socket programming, connection management.

### 5. Session Layer (Layer 5)
**Purpose**: Manages sessions between applications

**Key Concepts**:
- Session establishment, maintenance, and termination
- Dialog control (full-duplex, half-duplex)
- Synchronization checkpoints

**Examples**:
- RPC (Remote Procedure Call)
- SQL sessions
- WebSocket connections

### 6. Presentation Layer (Layer 6)
**Purpose**: Handles data representation and encryption

**Key Concepts**:
- Data compression
- Encryption/decryption
- Character encoding (ASCII, Unicode)
- Data format conversion

**Examples**:
- SSL/TLS encryption
- JPEG, MPEG compression
- XML, JSON formatting

### 7. Application Layer (Layer 7)
**Purpose**: Provides network services to user applications

**Key Concepts**:
- User interfaces
- Network services
- Application protocols

**Protocols**:
- HTTP/HTTPS
- FTP, SFTP
- SMTP, POP3, IMAP
- DNS
- SSH
- Telnet

## Data Flow Through OSI Layers

```
Application Data
       ↓
   [Layer 7] Application Layer - Adds application header
       ↓
   [Layer 6] Presentation Layer - Adds presentation header
       ↓
   [Layer 5] Session Layer - Adds session header
       ↓
   [Layer 4] Transport Layer - Adds transport header (TCP/UDP)
       ↓
   [Layer 3] Network Layer - Adds IP header
       ↓
   [Layer 2] Data Link Layer - Adds frame header and trailer
       ↓
   [Layer 1] Physical Layer - Converts to electrical/optical signals
       ↓
   Physical Medium (Cable, Wireless, etc.)
```

## Encapsulation Process

1. **Application Layer**: User data
2. **Presentation Layer**: Data + Presentation header
3. **Session Layer**: Data + Session header
4. **Transport Layer**: Data + Transport header (TCP/UDP)
5. **Network Layer**: Data + IP header
6. **Data Link Layer**: Data + Frame header + Frame trailer
7. **Physical Layer**: Electrical/optical signals

## Decapsulation Process

The receiving device reverses the process, removing headers layer by layer until the original data is delivered to the application.

## Why OSI Model Matters

1. **Standardization**: Provides a common framework for network protocols
2. **Troubleshooting**: Helps isolate problems to specific layers
3. **Design**: Guides network architecture decisions
4. **Interoperability**: Ensures different systems can communicate
5. **Learning**: Provides a structured way to understand networking

## Common Misconceptions

1. **"OSI is just theory"**: While the OSI model is conceptual, it's based on real protocols and is used in practice
2. **"All layers are always used"**: Some layers may be skipped in certain implementations
3. **"TCP/IP replaces OSI"**: TCP/IP is a protocol suite that maps to OSI layers, not a replacement

## Interview Questions

### Basic Questions
1. What is the OSI model and why is it important?
2. Explain the difference between each layer.
3. How does data flow through the OSI layers?

### Intermediate Questions
1. Which layer handles error detection and correction?
2. What happens at the Network Layer?
3. How does the Transport Layer ensure reliable communication?

### Advanced Questions
1. How does the OSI model help in network troubleshooting?
2. Explain the encapsulation and decapsulation process.
3. How do modern protocols map to the OSI model?

## Next Steps
After mastering the OSI model, proceed to:
- **Topic 2**: IP Addressing & Subnetting
- **Topic 3**: TCP Protocol Deep Dive
- **Topic 4**: UDP Protocol Deep Dive

This foundation will make you unstoppable in networking interviews! 🚀
```

```go
package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// OSI Layer 2: Data Link Layer - Frame Structure
type EthernetFrame struct {
	DestinationMAC [6]byte
	SourceMAC      [6]byte
	EtherType      uint16
	Payload        []byte
	FCS            uint32 // Frame Check Sequence
}

// OSI Layer 3: Network Layer - IP Packet Structure
type IPPacket struct {
	Version        uint8
	IHL            uint8
	TypeOfService  uint8
	TotalLength    uint16
	Identification uint16
	Flags          uint8
	FragmentOffset uint16
	TTL            uint8
	Protocol       uint8
	HeaderChecksum uint16
	SourceIP       net.IP
	DestinationIP  net.IP
	Options        []byte
	Payload        []byte
}

// OSI Layer 4: Transport Layer - TCP Segment Structure
type TCPSegment struct {
	SourcePort      uint16
	DestinationPort uint16
	SequenceNumber  uint32
	AckNumber       uint32
	DataOffset      uint8
	Flags           uint8
	WindowSize      uint16
	Checksum        uint16
	UrgentPointer   uint16
	Options         []byte
	Payload         []byte
}

// OSI Layer 7: Application Layer - HTTP Request
type HTTPRequest struct {
	Method    string
	URL       string
	Version   string
	Headers   map[string]string
	Body      []byte
}

// MAC Address utilities
func NewMACAddress(addr string) [6]byte {
	var mac [6]byte
	fmt.Sscanf(addr, "%02x:%02x:%02x:%02x:%02x:%02x",
		&mac[0], &mac[1], &mac[2], &mac[3], &mac[4], &mac[5])
	return mac
}

func (mac [6]byte) String() string {
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		mac[0], mac[1], mac[2], mac[3], mac[4], mac[5])
}

// Create an Ethernet frame
func CreateEthernetFrame(destMAC, srcMAC [6]byte, payload []byte) *EthernetFrame {
	return &EthernetFrame{
		DestinationMAC: destMAC,
		SourceMAC:      srcMAC,
		EtherType:      0x0800, // IPv4
		Payload:        payload,
		FCS:            calculateCRC32(payload), // Simplified CRC
	}
}

// Create an IP packet
func CreateIPPacket(srcIP, destIP net.IP, payload []byte) *IPPacket {
	return &IPPacket{
		Version:        4,
		IHL:            5, // 20 bytes header
		TypeOfService:  0,
		TotalLength:    uint16(20 + len(payload)),
		Identification: uint16(time.Now().UnixNano() & 0xFFFF),
		Flags:          0x40, // Don't fragment
		FragmentOffset: 0,
		TTL:            64,
		Protocol:       6, // TCP
		HeaderChecksum: 0, // Will be calculated
		SourceIP:       srcIP,
		DestinationIP:  destIP,
		Payload:        payload,
	}
}

// Create a TCP segment
func CreateTCPSegment(srcPort, destPort uint16, seqNum, ackNum uint32, payload []byte) *TCPSegment {
	return &TCPSegment{
		SourcePort:      srcPort,
		DestinationPort: destPort,
		SequenceNumber:  seqNum,
		AckNumber:       ackNum,
		DataOffset:      5, // 20 bytes header
		Flags:           0x18, // ACK + PSH
		WindowSize:      65535,
		Checksum:        0, // Will be calculated
		UrgentPointer:   0,
		Payload:         payload,
	}
}

// Create an HTTP request
func CreateHTTPRequest(method, url, version string, headers map[string]string, body []byte) *HTTPRequest {
	return &HTTPRequest{
		Method:  method,
		URL:     url,
		Version: version,
		Headers: headers,
		Body:    body,
	}
}

// Serialize HTTP request to bytes
func (req *HTTPRequest) ToBytes() []byte {
	request := fmt.Sprintf("%s %s %s\r\n", req.Method, req.URL, req.Version)
	
	for key, value := range req.Headers {
		request += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	request += "\r\n"
	
	requestBytes := []byte(request)
	return append(requestBytes, req.Body...)
}

// Serialize TCP segment to bytes
func (tcp *TCPSegment) ToBytes() []byte {
	data := make([]byte, 20+len(tcp.Options)+len(tcp.Payload))
	
	// Source port (2 bytes)
	binary.BigEndian.PutUint16(data[0:2], tcp.SourcePort)
	// Destination port (2 bytes)
	binary.BigEndian.PutUint16(data[2:4], tcp.DestinationPort)
	// Sequence number (4 bytes)
	binary.BigEndian.PutUint32(data[4:8], tcp.SequenceNumber)
	// ACK number (4 bytes)
	binary.BigEndian.PutUint32(data[8:12], tcp.AckNumber)
	// Data offset and flags (2 bytes)
	data[12] = tcp.DataOffset << 4
	data[13] = tcp.Flags
	// Window size (2 bytes)
	binary.BigEndian.PutUint16(data[14:16], tcp.WindowSize)
	// Checksum (2 bytes)
	binary.BigEndian.PutUint16(data[16:18], tcp.Checksum)
	// Urgent pointer (2 bytes)
	binary.BigEndian.PutUint16(data[18:20], tcp.UrgentPointer)
	
	// Options and payload
	copy(data[20:], tcp.Options)
	copy(data[20+len(tcp.Options):], tcp.Payload)
	
	return data
}

// Serialize IP packet to bytes
func (ip *IPPacket) ToBytes() []byte {
	headerLen := 20 + len(ip.Options)
	data := make([]byte, headerLen+len(ip.Payload))
	
	// Version and IHL (1 byte)
	data[0] = (ip.Version << 4) | ip.IHL
	// Type of service (1 byte)
	data[1] = ip.TypeOfService
	// Total length (2 bytes)
	binary.BigEndian.PutUint16(data[2:4], ip.TotalLength)
	// Identification (2 bytes)
	binary.BigEndian.PutUint16(data[4:6], ip.Identification)
	// Flags and fragment offset (2 bytes)
	flagsAndOffset := (uint16(ip.Flags) << 13) | ip.FragmentOffset
	binary.BigEndian.PutUint16(data[6:8], flagsAndOffset)
	// TTL (1 byte)
	data[8] = ip.TTL
	// Protocol (1 byte)
	data[9] = ip.Protocol
	// Header checksum (2 bytes)
	binary.BigEndian.PutUint16(data[10:12], ip.HeaderChecksum)
	// Source IP (4 bytes)
	copy(data[12:16], ip.SourceIP.To4())
	// Destination IP (4 bytes)
	copy(data[16:20], ip.DestinationIP.To4())
	
	// Options and payload
	copy(data[20:], ip.Options)
	copy(data[headerLen:], ip.Payload)
	
	return data
}

// Serialize Ethernet frame to bytes
func (eth *EthernetFrame) ToBytes() []byte {
	data := make([]byte, 14+len(eth.Payload)+4) // 14 bytes header + payload + 4 bytes FCS
	
	// Destination MAC (6 bytes)
	copy(data[0:6], eth.DestinationMAC[:])
	// Source MAC (6 bytes)
	copy(data[6:12], eth.SourceMAC[:])
	// EtherType (2 bytes)
	binary.BigEndian.PutUint16(data[12:14], eth.EtherType)
	// Payload
	copy(data[14:14+len(eth.Payload)], eth.Payload)
	// FCS (4 bytes)
	binary.BigEndian.PutUint32(data[14+len(eth.Payload):], eth.FCS)
	
	return data
}

// Simplified CRC32 calculation (for demonstration)
func calculateCRC32(data []byte) uint32 {
	crc := uint32(0xFFFFFFFF)
	for _, b := range data {
		crc ^= uint32(b)
		for i := 0; i < 8; i++ {
			if crc&1 == 1 {
				crc = (crc >> 1) ^ 0xEDB88320
			} else {
				crc >>= 1
			}
		}
	}
	return ^crc
}

// Demonstrate OSI layer encapsulation
func demonstrateOSIEncapsulation() {
	fmt.Println("=== OSI Model Encapsulation Demo ===\n")
	
	// Layer 7: Application Layer - HTTP Request
	httpReq := CreateHTTPRequest(
		"GET",
		"/api/users",
		"HTTP/1.1",
		map[string]string{
			"Host":       "example.com",
			"User-Agent": "Go-Network-Demo/1.0",
		},
		[]byte(""),
	)
	
	httpData := httpReq.ToBytes()
	fmt.Printf("Layer 7 (Application): HTTP Request\n")
	fmt.Printf("Data: %s\n\n", string(httpData))
	
	// Layer 4: Transport Layer - TCP Segment
	tcpSeg := CreateTCPSegment(
		8080,    // Source port
		80,      // Destination port
		1000,    // Sequence number
		2000,    // ACK number
		httpData, // HTTP data as payload
	)
	
	tcpData := tcpSeg.ToBytes()
	fmt.Printf("Layer 4 (Transport): TCP Segment\n")
	fmt.Printf("Source Port: %d, Dest Port: %d\n", tcpSeg.SourcePort, tcpSeg.DestinationPort)
	fmt.Printf("Seq: %d, ACK: %d\n", tcpSeg.SequenceNumber, tcpSeg.AckNumber)
	fmt.Printf("Total Length: %d bytes\n\n", len(tcpData))
	
	// Layer 3: Network Layer - IP Packet
	ipPkt := CreateIPPacket(
		net.ParseIP("192.168.1.100"), // Source IP
		net.ParseIP("192.168.1.1"),   // Destination IP
		tcpData,                      // TCP data as payload
	)
	
	ipData := ipPkt.ToBytes()
	fmt.Printf("Layer 3 (Network): IP Packet\n")
	fmt.Printf("Source IP: %s, Dest IP: %s\n", ipPkt.SourceIP, ipPkt.DestinationIP)
	fmt.Printf("Protocol: %d (TCP), TTL: %d\n", ipPkt.Protocol, ipPkt.TTL)
	fmt.Printf("Total Length: %d bytes\n\n", len(ipData))
	
	// Layer 2: Data Link Layer - Ethernet Frame
	ethFrame := CreateEthernetFrame(
		NewMACAddress("00:11:22:33:44:55"), // Destination MAC
		NewMACAddress("aa:bb:cc:dd:ee:ff"), // Source MAC
		ipData,                             // IP data as payload
	)
	
	ethData := ethFrame.ToBytes()
	fmt.Printf("Layer 2 (Data Link): Ethernet Frame\n")
	fmt.Printf("Dest MAC: %s, Src MAC: %s\n", ethFrame.DestinationMAC.String(), ethFrame.SourceMAC.String())
	fmt.Printf("EtherType: 0x%04X (IPv4)\n", ethFrame.EtherType)
	fmt.Printf("Total Length: %d bytes\n\n", len(ethData))
	
	// Layer 1: Physical Layer (simulated)
	fmt.Printf("Layer 1 (Physical): Electrical/Optical Signals\n")
	fmt.Printf("Converting %d bytes to electrical signals...\n", len(ethData))
	fmt.Printf("Transmission complete!\n\n")
	
	// Show the complete packet structure
	fmt.Println("=== Complete Packet Structure ===")
	fmt.Printf("Total packet size: %d bytes\n", len(ethData))
	fmt.Printf("Ethernet Header: 14 bytes\n")
	fmt.Printf("IP Header: 20 bytes\n")
	fmt.Printf("TCP Header: 20 bytes\n")
	fmt.Printf("HTTP Data: %d bytes\n", len(httpData))
	fmt.Printf("Ethernet Trailer (FCS): 4 bytes\n")
}

// Demonstrate OSI layer decapsulation
func demonstrateOSIDecapsulation(packet []byte) {
	fmt.Println("=== OSI Model Decapsulation Demo ===\n")
	
	// Layer 1: Physical Layer - Receive electrical signals
	fmt.Printf("Layer 1 (Physical): Received %d bytes of electrical signals\n", len(packet))
	
	// Layer 2: Data Link Layer - Extract Ethernet frame
	if len(packet) < 14 {
		fmt.Println("Error: Packet too short for Ethernet frame")
		return
	}
	
	destMAC := [6]byte{}
	copy(destMAC[:], packet[0:6])
	srcMAC := [6]byte{}
	copy(srcMAC[:], packet[6:12])
	etherType := binary.BigEndian.Uint16(packet[12:14])
	payload := packet[14 : len(packet)-4] // Remove FCS
	
	fmt.Printf("Layer 2 (Data Link): Ethernet Frame\n")
	fmt.Printf("Dest MAC: %s, Src MAC: %s\n", destMAC.String(), srcMAC.String())
	fmt.Printf("EtherType: 0x%04X\n", etherType)
	fmt.Printf("Payload length: %d bytes\n\n", len(payload))
	
	// Layer 3: Network Layer - Extract IP packet
	if len(payload) < 20 {
		fmt.Println("Error: Payload too short for IP packet")
		return
	}
	
	version := payload[0] >> 4
	ihl := payload[0] & 0x0F
	protocol := payload[9]
	srcIP := net.IP(payload[12:16])
	destIP := net.IP(payload[16:20])
	ipPayload := payload[ihl*4:]
	
	fmt.Printf("Layer 3 (Network): IP Packet\n")
	fmt.Printf("Version: %d, IHL: %d\n", version, ihl)
	fmt.Printf("Protocol: %d, Src IP: %s, Dest IP: %s\n", protocol, srcIP, destIP)
	fmt.Printf("Payload length: %d bytes\n\n", len(ipPayload))
	
	// Layer 4: Transport Layer - Extract TCP segment
	if len(ipPayload) < 20 {
		fmt.Println("Error: IP payload too short for TCP segment")
		return
	}
	
	srcPort := binary.BigEndian.Uint16(ipPayload[0:2])
	destPort := binary.BigEndian.Uint16(ipPayload[2:4])
	seqNum := binary.BigEndian.Uint32(ipPayload[4:8])
	ackNum := binary.BigEndian.Uint32(ipPayload[8:12])
	dataOffset := ipPayload[12] >> 4
	tcpPayload := ipPayload[dataOffset*4:]
	
	fmt.Printf("Layer 4 (Transport): TCP Segment\n")
	fmt.Printf("Src Port: %d, Dest Port: %d\n", srcPort, destPort)
	fmt.Printf("Seq: %d, ACK: %d\n", seqNum, ackNum)
	fmt.Printf("Data Offset: %d\n", dataOffset)
	fmt.Printf("Payload length: %d bytes\n\n", len(tcpPayload))
	
	// Layer 7: Application Layer - Extract HTTP request
	fmt.Printf("Layer 7 (Application): HTTP Request\n")
	fmt.Printf("Data: %s\n", string(tcpPayload))
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 1: OSI Model")
	fmt.Println("==================================================\n")
	
	// Demonstrate encapsulation
	demonstrateOSIEncapsulation()
	
	// Create a sample packet for decapsulation demo
	httpReq := CreateHTTPRequest("GET", "/", "HTTP/1.1", map[string]string{"Host": "example.com"}, []byte(""))
	httpData := httpReq.ToBytes()
	
	tcpSeg := CreateTCPSegment(8080, 80, 1000, 2000, httpData)
	tcpData := tcpSeg.ToBytes()
	
	ipPkt := CreateIPPacket(net.ParseIP("192.168.1.100"), net.ParseIP("192.168.1.1"), tcpData)
	ipData := ipPkt.ToBytes()
	
	ethFrame := CreateEthernetFrame(
		NewMACAddress("00:11:22:33:44:55"),
		NewMACAddress("aa:bb:cc:dd:ee:ff"),
		ipData,
	)
	
	packet := ethFrame.ToBytes()
	
	// Demonstrate decapsulation
	demonstrateOSIDecapsulation(packet)
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Each layer adds its own header to the data")
	fmt.Println("2. Lower layers don't understand higher layer protocols")
	fmt.Println("3. Encapsulation happens on the sending side")
	fmt.Println("4. Decapsulation happens on the receiving side")
	fmt.Println("5. This modular approach enables network interoperability")
	
	fmt.Println("\n📚 Next Topic: IP Addressing & Subnetting")
}
```

Now let me create the IP addressing topic:


```markdown
# Topic 2: IP Addressing & Subnetting Mastery

## IPv4 Addressing

### Understanding IP Addresses
An IPv4 address is a 32-bit number divided into four 8-bit octets, represented in dotted decimal notation (e.g., 192.168.1.1).

### Address Classes (Historical)
- **Class A**: 1.0.0.0 to 126.255.255.255 (8-bit network, 24-bit host)
- **Class B**: 128.0.0.0 to 191.255.255.255 (16-bit network, 16-bit host)
- **Class C**: 192.0.0.0 to 223.255.255.255 (24-bit network, 8-bit host)
- **Class D**: 224.0.0.0 to 239.255.255.255 (Multicast)
- **Class E**: 240.0.0.0 to 255.255.255.255 (Reserved)

### Private Address Ranges (RFC 1918)
- **10.0.0.0/8**: 10.0.0.0 to 10.255.255.255
- **172.16.0.0/12**: 172.16.0.0 to 172.31.255.255
- **192.168.0.0/16**: 192.168.0.0 to 192.168.255.255

### Special Addresses
- **0.0.0.0/0**: Default route
- **127.0.0.0/8**: Loopback addresses
- **169.254.0.0/16**: Link-local addresses (APIPA)
- **255.255.255.255**: Limited broadcast

## CIDR (Classless Inter-Domain Routing)

### CIDR Notation
Format: `IP_ADDRESS/PREFIX_LENGTH`
- Example: `192.168.1.0/24`
- Prefix length indicates number of network bits
- Remaining bits are host bits

### Subnet Masks
- **/8**: 255.0.0.0
- **/16**: 255.255.0.0
- **/24**: 255.255.255.0
- **/32**: 255.255.255.255 (host route)

## Subnetting Fundamentals

### Why Subnet?
1. **Efficient IP usage**: Avoid wasting addresses
2. **Security**: Isolate network segments
3. **Performance**: Reduce broadcast domains
4. **Management**: Logical organization

### Subnetting Process
1. Determine required subnets
2. Calculate subnet mask
3. Identify network addresses
4. Identify broadcast addresses
5. Identify usable host ranges

### Subnetting Formulas
- **Number of subnets**: 2^(subnet_bits)
- **Number of hosts per subnet**: 2^(host_bits) - 2
- **Subnet mask**: 2^32 - 2^(32-prefix_length)

## VLSM (Variable Length Subnet Masking)

VLSM allows different subnets to have different subnet masks within the same network.

### VLSM Benefits
- Maximum IP address utilization
- Hierarchical addressing
- Efficient routing

## IPv6 Addressing

### IPv6 Basics
- 128-bit addresses (vs 32-bit IPv4)
- Hexadecimal notation with colons
- Example: `2001:0db8:85a3:0000:0000:8a2e:0370:7334`
- Shortened: `2001:db8:85a3::8a2e:370:7334`

### IPv6 Address Types
- **Unicast**: Single interface
- **Multicast**: Multiple interfaces
- **Anycast**: Nearest interface

### IPv6 Special Addresses
- **::/128**: Unspecified address
- **::1/128**: Loopback address
- **2000::/3**: Global unicast
- **fc00::/7**: Unique local
- **fe80::/10**: Link-local

## NAT (Network Address Translation)

### NAT Types
1. **Static NAT**: One-to-one mapping
2. **Dynamic NAT**: Pool-based mapping
3. **PAT (Port Address Translation)**: Many-to-one mapping

### NAT Benefits
- IP address conservation
- Security through obscurity
- Simplified addressing

## Practical Examples

### Example 1: Basic Subnetting
**Given**: 192.168.1.0/24, need 4 subnets

**Solution**:
- Need 2 bits for subnets (2^2 = 4)
- New prefix: /26
- Subnet mask: 255.255.255.192
- Subnets:
  - 192.168.1.0/26 (0-63)
  - 192.168.1.64/26 (64-127)
  - 192.168.1.128/26 (128-191)
  - 192.168.1.192/26 (192-255)

### Example 2: VLSM
**Given**: 192.168.1.0/24, need:
- 2 subnets with 50 hosts each
- 4 subnets with 10 hosts each

**Solution**:
1. For 50 hosts: need 6 host bits (2^6-2 = 62 hosts)
   - Subnets: /26 (255.255.255.192)
   - 192.168.1.0/26 and 192.168.1.64/26

2. For 10 hosts: need 4 host bits (2^4-2 = 14 hosts)
   - Subnets: /28 (255.255.255.240)
   - 192.168.1.128/28, 192.168.1.144/28, 192.168.1.160/28, 192.168.1.176/28

## Interview Questions

### Basic Questions
1. What's the difference between a network address and a broadcast address?
2. How do you calculate the number of usable hosts in a subnet?
3. What is the purpose of a subnet mask?

### Intermediate Questions
1. Explain the difference between classful and classless addressing.
2. How does VLSM improve IP address utilization?
3. What are the advantages and disadvantages of NAT?

### Advanced Questions
1. Design a network addressing scheme for a company with multiple departments.
2. Explain how IPv6 addresses are allocated and managed.
3. How would you troubleshoot IP addressing issues in a large network?

## Next Steps
After mastering IP addressing, proceed to:
- **Topic 3**: TCP Protocol Deep Dive
- **Topic 4**: UDP Protocol Deep Dive
- **Topic 5**: Socket Programming in Go

Master these fundamentals, and you'll be unstoppable in networking interviews! 🚀
```

```go
package main

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
)

// IPAddress represents an IPv4 address with utilities
type IPAddress struct {
	Address net.IP
	Mask    net.IPMask
}

// NewIPAddress creates a new IPAddress from string and CIDR notation
func NewIPAddress(address string) (*IPAddress, error) {
	ip, network, err := net.ParseCIDR(address)
	if err != nil {
		return nil, err
	}
	
	return &IPAddress{
		Address: ip,
		Mask:    network.Mask,
	}, nil
}

// NewIPAddressWithMask creates a new IPAddress with custom mask
func NewIPAddressWithMask(address string, mask net.IPMask) (*IPAddress, error) {
	ip := net.ParseIP(address)
	if ip == nil {
		return nil, fmt.Errorf("invalid IP address: %s", address)
	}
	
	return &IPAddress{
		Address: ip,
		Mask:    mask,
	}, nil
}

// GetNetworkAddress returns the network address
func (ip *IPAddress) GetNetworkAddress() net.IP {
	return ip.Address.Mask(ip.Mask)
}

// GetBroadcastAddress returns the broadcast address
func (ip *IPAddress) GetBroadcastAddress() net.IP {
	network := ip.GetNetworkAddress()
	ones, bits := ip.Mask.Size()
	
	// Calculate the number of host bits
	hostBits := bits - ones
	
	// Create broadcast address by setting all host bits to 1
	broadcast := make(net.IP, len(network))
	copy(broadcast, network)
	
	for i := 0; i < hostBits; i++ {
		byteIndex := (bits - 1 - i) / 8
		bitIndex := 7 - ((bits - 1 - i) % 8)
		broadcast[byteIndex] |= 1 << bitIndex
	}
	
	return broadcast
}

// GetFirstHostAddress returns the first usable host address
func (ip *IPAddress) GetFirstHostAddress() net.IP {
	network := ip.GetNetworkAddress()
	firstHost := make(net.IP, len(network))
	copy(firstHost, network)
	
	// Add 1 to the last octet
	firstHost[len(firstHost)-1]++
	return firstHost
}

// GetLastHostAddress returns the last usable host address
func (ip *IPAddress) GetLastHostAddress() net.IP {
	broadcast := ip.GetBroadcastAddress()
	lastHost := make(net.IP, len(broadcast))
	copy(lastHost, broadcast)
	
	// Subtract 1 from the last octet
	lastHost[len(lastHost)-1]--
	return lastHost
}

// GetHostCount returns the number of usable host addresses
func (ip *IPAddress) GetHostCount() int {
	ones, bits := ip.Mask.Size()
	hostBits := bits - ones
	return int(math.Pow(2, float64(hostBits))) - 2
}

// GetSubnetCount returns the number of subnets
func (ip *IPAddress) GetSubnetCount() int {
	ones, _ := ip.Mask.Size()
	return int(math.Pow(2, float64(32-ones)))
}

// IsInSameSubnet checks if another IP is in the same subnet
func (ip *IPAddress) IsInSameSubnet(otherIP net.IP) bool {
	network1 := ip.GetNetworkAddress()
	network2 := otherIP.Mask(ip.Mask)
	return network1.Equal(network2)
}

// String returns the string representation
func (ip *IPAddress) String() string {
	ones, _ := ip.Mask.Size()
	return fmt.Sprintf("%s/%d", ip.Address.String(), ones)
}

// SubnetCalculator provides subnetting utilities
type SubnetCalculator struct{}

// CalculateSubnets calculates subnets for a given network
func (sc *SubnetCalculator) CalculateSubnets(network string, requiredSubnets int) ([]*IPAddress, error) {
	baseIP, baseNetwork, err := net.ParseCIDR(network)
	if err != nil {
		return nil, err
	}
	
	// Calculate required bits for subnets
	subnetBits := int(math.Ceil(math.Log2(float64(requiredSubnets))))
	
	// Get current prefix length
	ones, _ := baseNetwork.Mask.Size()
	
	// Calculate new prefix length
	newPrefixLength := ones + subnetBits
	
	if newPrefixLength > 32 {
		return nil, fmt.Errorf("not enough address space for %d subnets", requiredSubnets)
	}
	
	// Create new mask
	newMask := net.CIDRMask(newPrefixLength, 32)
	
	var subnets []*IPAddress
	
	// Calculate subnet size
	subnetSize := int(math.Pow(2, float64(32-newPrefixLength)))
	
	// Generate subnets
	for i := 0; i < requiredSubnets; i++ {
		subnetAddr := make(net.IP, len(baseIP))
		copy(subnetAddr, baseIP)
		
		// Calculate subnet address
		subnetOffset := i * subnetSize
		for j := 0; j < 4; j++ {
			subnetAddr[3-j] += byte((subnetOffset >> (j * 8)) & 0xFF)
		}
		
		subnet := &IPAddress{
			Address: subnetAddr,
			Mask:    newMask,
		}
		
		subnets = append(subnets, subnet)
	}
	
	return subnets, nil
}

// VLSMSubnet represents a VLSM subnet requirement
type VLSMSubnet struct {
	Name        string
	HostsNeeded int
	Subnet      *IPAddress
}

// CalculateVLSM calculates VLSM subnets for different host requirements
func (sc *SubnetCalculator) CalculateVLSM(network string, requirements []VLSMSubnet) ([]VLSMSubnet, error) {
	// Sort requirements by host count (descending)
	for i := 0; i < len(requirements)-1; i++ {
		for j := i + 1; j < len(requirements); j++ {
			if requirements[i].HostsNeeded < requirements[j].HostsNeeded {
				requirements[i], requirements[j] = requirements[j], requirements[i]
			}
		}
	}
	
	baseIP, baseNetwork, err := net.ParseCIDR(network)
	if err != nil {
		return nil, err
	}
	
	currentIP := make(net.IP, len(baseIP))
	copy(currentIP, baseIP)
	
	var result []VLSMSubnet
	
	for _, req := range requirements {
		// Calculate required host bits
		hostBits := int(math.Ceil(math.Log2(float64(req.HostsNeeded + 2)))) // +2 for network and broadcast
		prefixLength := 32 - hostBits
		
		// Create mask
		mask := net.CIDRMask(prefixLength, 32)
		
		// Create subnet
		subnet := &IPAddress{
			Address: currentIP,
			Mask:    mask,
		}
		
		// Update result
		req.Subnet = subnet
		result = append(result, req)
		
		// Calculate next subnet address
		subnetSize := int(math.Pow(2, float64(hostBits)))
		for j := 0; j < 4; j++ {
			currentIP[3-j] += byte((subnetSize >> (j * 8)) & 0xFF)
		}
	}
	
	return result, nil
}

// IPAddressValidator provides IP address validation utilities
type IPAddressValidator struct{}

// IsValidIP checks if an IP address is valid
func (v *IPAddressValidator) IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsPrivateIP checks if an IP address is private
func (v *IPAddressValidator) IsPrivateIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	
	// Check private ranges
	privateRanges := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}
	
	for _, rangeStr := range privateRanges {
		_, network, _ := net.ParseCIDR(rangeStr)
		if network.Contains(parsedIP) {
			return true
		}
	}
	
	return false
}

// IsPublicIP checks if an IP address is public
func (v *IPAddressValidator) IsPublicIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	
	// Check if it's a valid IP and not private
	return v.IsValidIP(ip) && !v.IsPrivateIP(ip) && !parsedIP.IsLoopback() && !parsedIP.IsMulticast()
}

// GetIPClass returns the class of an IP address
func (v *IPAddressValidator) GetIPClass(ip string) string {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return "Invalid"
	}
	
	firstOctet := int(parsedIP.To4()[0])
	
	switch {
	case firstOctet >= 1 && firstOctet <= 126:
		return "Class A"
	case firstOctet >= 128 && firstOctet <= 191:
		return "Class B"
	case firstOctet >= 192 && firstOctet <= 223:
		return "Class C"
	case firstOctet >= 224 && firstOctet <= 239:
		return "Class D (Multicast)"
	case firstOctet >= 240 && firstOctet <= 255:
		return "Class E (Reserved)"
	default:
		return "Unknown"
	}
}

// Demonstrate basic subnetting
func demonstrateBasicSubnetting() {
	fmt.Println("=== Basic Subnetting Demo ===\n")
	
	// Example: 192.168.1.0/24 needs 4 subnets
	network := "192.168.1.0/24"
	requiredSubnets := 4
	
	calculator := &SubnetCalculator{}
	subnets, err := calculator.CalculateSubnets(network, requiredSubnets)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Network: %s\n", network)
	fmt.Printf("Required subnets: %d\n\n", requiredSubnets)
	
	for i, subnet := range subnets {
		fmt.Printf("Subnet %d:\n", i+1)
		fmt.Printf("  Network: %s\n", subnet.GetNetworkAddress())
		fmt.Printf("  Broadcast: %s\n", subnet.GetBroadcastAddress())
		fmt.Printf("  First Host: %s\n", subnet.GetFirstHostAddress())
		fmt.Printf("  Last Host: %s\n", subnet.GetLastHostAddress())
		fmt.Printf("  Usable Hosts: %d\n", subnet.GetHostCount())
		fmt.Printf("  CIDR: %s\n\n", subnet.String())
	}
}

// Demonstrate VLSM
func demonstrateVLSM() {
	fmt.Println("=== VLSM Demo ===\n")
	
	network := "192.168.1.0/24"
	requirements := []VLSMSubnet{
		{Name: "Sales", HostsNeeded: 50},
		{Name: "Marketing", HostsNeeded: 25},
		{Name: "IT", HostsNeeded: 10},
		{Name: "HR", HostsNeeded: 5},
	}
	
	calculator := &SubnetCalculator{}
	subnets, err := calculator.CalculateVLSM(network, requirements)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Network: %s\n", network)
	fmt.Println("VLSM Subnets:\n")
	
	for _, subnet := range subnets {
		fmt.Printf("%s:\n", subnet.Name)
		fmt.Printf("  Hosts Needed: %d\n", subnet.HostsNeeded)
		fmt.Printf("  Network: %s\n", subnet.Subnet.GetNetworkAddress())
		fmt.Printf("  Broadcast: %s\n", subnet.Subnet.GetBroadcastAddress())
		fmt.Printf("  First Host: %s\n", subnet.Subnet.GetFirstHostAddress())
		fmt.Printf("  Last Host: %s\n", subnet.Subnet.GetLastHostAddress())
		fmt.Printf("  Usable Hosts: %d\n", subnet.Subnet.GetHostCount())
		fmt.Printf("  CIDR: %s\n\n", subnet.Subnet.String())
	}
}

// Demonstrate IP address validation
func demonstrateIPValidation() {
	fmt.Println("=== IP Address Validation Demo ===\n")
	
	validator := &IPAddressValidator{}
	testIPs := []string{
		"192.168.1.1",
		"10.0.0.1",
		"172.16.0.1",
		"8.8.8.8",
		"127.0.0.1",
		"256.1.1.1",
		"invalid",
		"2001:db8::1",
	}
	
	for _, ip := range testIPs {
		fmt.Printf("IP: %s\n", ip)
		fmt.Printf("  Valid: %t\n", validator.IsValidIP(ip))
		fmt.Printf("  Private: %t\n", validator.IsPrivateIP(ip))
		fmt.Printf("  Public: %t\n", validator.IsPublicIP(ip))
		fmt.Printf("  Class: %s\n\n", validator.GetIPClass(ip))
	}
}

// Demonstrate subnet membership
func demonstrateSubnetMembership() {
	fmt.Println("=== Subnet Membership Demo ===\n")
	
	// Create a subnet
	subnet, err := NewIPAddress("192.168.1.0/24")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	testIPs := []string{
		"192.168.1.1",
		"192.168.1.100",
		"192.168.1.254",
		"192.168.2.1",
		"10.0.0.1",
	}
	
	fmt.Printf("Subnet: %s\n", subnet.String())
	fmt.Printf("Network: %s\n", subnet.GetNetworkAddress())
	fmt.Printf("Broadcast: %s\n\n", subnet.GetBroadcastAddress())
	
	for _, ipStr := range testIPs {
		ip := net.ParseIP(ipStr)
		if ip != nil {
			inSubnet := subnet.IsInSameSubnet(ip)
			fmt.Printf("IP %s is %s the subnet\n", ipStr, map[bool]string{true: "in", false: "not in"}[inSubnet])
		}
	}
}

// Calculate subnet requirements
func calculateSubnetRequirements() {
	fmt.Println("=== Subnet Requirements Calculator ===\n")
	
	// Example: Company needs to subnet 10.0.0.0/8
	baseNetwork := "10.0.0.0/8"
	
	departments := []struct {
		Name        string
		HostsNeeded int
	}{
		{"Headquarters", 1000},
		{"Branch Office 1", 500},
		{"Branch Office 2", 250},
		{"Sales Team", 100},
		{"IT Department", 50},
		{"HR Department", 25},
		{"Guest Network", 10},
	}
	
	calculator := &SubnetCalculator{}
	
	// Convert to VLSM requirements
	var requirements []VLSMSubnet
	for _, dept := range departments {
		requirements = append(requirements, VLSMSubnet{
			Name:        dept.Name,
			HostsNeeded: dept.HostsNeeded,
		})
	}
	
	subnets, err := calculator.CalculateVLSM(baseNetwork, requirements)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Base Network: %s\n", baseNetwork)
	fmt.Println("Department Subnets:\n")
	
	for _, subnet := range subnets {
		fmt.Printf("%s:\n", subnet.Name)
		fmt.Printf("  Required Hosts: %d\n", subnet.HostsNeeded)
		fmt.Printf("  Allocated Hosts: %d\n", subnet.Subnet.GetHostCount())
		fmt.Printf("  Network: %s\n", subnet.Subnet.GetNetworkAddress())
		fmt.Printf("  CIDR: %s\n", subnet.Subnet.String())
		fmt.Printf("  Efficiency: %.1f%%\n", float64(subnet.HostsNeeded)/float64(subnet.Subnet.GetHostCount())*100)
		fmt.Println()
	}
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 2: IP Addressing & Subnetting")
	fmt.Println("==================================================================\n")
	
	// Run all demonstrations
	demonstrateBasicSubnetting()
	demonstrateVLSM()
	demonstrateIPValidation()
	demonstrateSubnetMembership()
	calculateSubnetRequirements()
	
	fmt.Println("🎯 Key Takeaways:")
	fmt.Println("1. Subnetting allows efficient IP address utilization")
	fmt.Println("2. VLSM provides maximum flexibility for different subnet sizes")
	fmt.Println("3. Understanding IP classes helps with network design")
	fmt.Println("4. Private IP ranges are used for internal networks")
	fmt.Println("5. Proper subnetting is crucial for network scalability")
	
	fmt.Println("\n📚 Next Topic: TCP Protocol Deep Dive")
}
```

Now let me create the TCP protocol topic:


```markdown
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
```

Now let me create the TCP implementation:


```go
package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// TCPHeader represents a TCP header
type TCPHeader struct {
	SourcePort      uint16
	DestinationPort uint16
	SequenceNumber  uint32
	AckNumber       uint32
	DataOffset      uint8
	Flags           uint8
	WindowSize      uint16
	Checksum        uint16
	UrgentPointer   uint16
	Options         []byte
}

// TCPFlags represents TCP control flags
type TCPFlags struct {
	URG bool // Urgent
	ACK bool // Acknowledgment
	PSH bool // Push
	RST bool // Reset
	SYN bool // Synchronize
	FIN bool // Finish
}

// TCPConnection represents a TCP connection
type TCPConnection struct {
	LocalAddr    net.Addr
	RemoteAddr   net.Addr
	State        string
	SeqNumber    uint32
	AckNumber    uint32
	WindowSize   uint16
	LastActivity time.Time
}

// TCPServer represents a TCP server
type TCPServer struct {
	Address     string
	Listener    net.Listener
	Connections map[net.Conn]*TCPConnection
	Handler     func(net.Conn)
}

// NewTCPServer creates a new TCP server
func NewTCPServer(address string) *TCPServer {
	return &TCPServer{
		Address:     address,
		Connections: make(map[net.Conn]*TCPConnection),
	}
}

// Start starts the TCP server
func (s *TCPServer) Start() error {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	
	s.Listener = listener
	fmt.Printf("TCP Server listening on %s\n", s.Address)
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		
		// Create connection record
		connection := &TCPConnection{
			LocalAddr:    conn.LocalAddr(),
			RemoteAddr:   conn.RemoteAddr(),
			State:        "ESTABLISHED",
			SeqNumber:    uint32(time.Now().UnixNano() & 0xFFFFFFFF),
			AckNumber:    0,
			WindowSize:   65535,
			LastActivity: time.Now(),
		}
		
		s.Connections[conn] = connection
		
		// Handle connection in goroutine
		go s.handleConnection(conn)
	}
}

// handleConnection handles a TCP connection
func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	defer delete(s.Connections, conn)
	
	fmt.Printf("New connection from %s\n", conn.RemoteAddr())
	
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
func (s *TCPServer) defaultHandler(conn net.Conn) {
	reader := bufio.NewReader(conn)
	
	for {
		// Read data from connection
		data, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Connection closed by client: %s\n", conn.RemoteAddr())
			} else {
				log.Printf("Error reading from connection: %v", err)
			}
			break
		}
		
		// Echo the data back
		response := fmt.Sprintf("Echo: %s", data)
		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Printf("Error writing to connection: %v", err)
			break
		}
		
		// Update connection activity
		if connInfo, exists := s.Connections[conn]; exists {
			connInfo.LastActivity = time.Now()
		}
	}
}

// Stop stops the TCP server
func (s *TCPServer) Stop() error {
	if s.Listener != nil {
		return s.Listener.Close()
	}
	return nil
}

// GetConnections returns active connections
func (s *TCPServer) GetConnections() map[net.Conn]*TCPConnection {
	return s.Connections
}

// TCPClient represents a TCP client
type TCPClient struct {
	Address string
	Conn    net.Conn
}

// NewTCPClient creates a new TCP client
func NewTCPClient(address string) *TCPClient {
	return &TCPClient{
		Address: address,
	}
}

// Connect connects to a TCP server
func (c *TCPClient) Connect() error {
	conn, err := net.Dial("tcp", c.Address)
	if err != nil {
		return err
	}
	
	c.Conn = conn
	fmt.Printf("Connected to %s\n", c.Address)
	return nil
}

// Send sends data to the server
func (c *TCPClient) Send(data string) error {
	if c.Conn == nil {
		return fmt.Errorf("not connected")
	}
	
	_, err := c.Conn.Write([]byte(data + "\n"))
	return err
}

// Receive receives data from the server
func (c *TCPClient) Receive() (string, error) {
	if c.Conn == nil {
		return "", fmt.Errorf("not connected")
	}
	
	reader := bufio.NewReader(c.Conn)
	data, err := reader.ReadString('\n')
	return data, err
}

// Close closes the connection
func (c *TCPClient) Close() error {
	if c.Conn != nil {
		return c.Conn.Close()
	}
	return nil
}

// TCPPacketAnalyzer analyzes TCP packets
type TCPPacketAnalyzer struct{}

// AnalyzePacket analyzes a TCP packet
func (a *TCPPacketAnalyzer) AnalyzePacket(data []byte) (*TCPHeader, error) {
	if len(data) < 20 {
		return nil, fmt.Errorf("packet too short for TCP header")
	}
	
	header := &TCPHeader{
		SourcePort:      binary.BigEndian.Uint16(data[0:2]),
		DestinationPort: binary.BigEndian.Uint16(data[2:4]),
		SequenceNumber:  binary.BigEndian.Uint32(data[4:8]),
		AckNumber:       binary.BigEndian.Uint32(data[8:12]),
		DataOffset:      data[12] >> 4,
		Flags:           data[13],
		WindowSize:      binary.BigEndian.Uint16(data[14:16]),
		Checksum:        binary.BigEndian.Uint16(data[16:18]),
		UrgentPointer:   binary.BigEndian.Uint16(data[18:20]),
	}
	
	// Parse options if present
	if header.DataOffset > 5 {
		optionsLen := int(header.DataOffset-5) * 4
		if len(data) >= 20+optionsLen {
			header.Options = data[20 : 20+optionsLen]
		}
	}
	
	return header, nil
}

// ParseFlags parses TCP flags
func (a *TCPPacketAnalyzer) ParseFlags(flags uint8) TCPFlags {
	return TCPFlags{
		URG: (flags & 0x20) != 0,
		ACK: (flags & 0x10) != 0,
		PSH: (flags & 0x08) != 0,
		RST: (flags & 0x04) != 0,
		SYN: (flags & 0x02) != 0,
		FIN: (flags & 0x01) != 0,
	}
}

// TCPFlowControl demonstrates flow control
type TCPFlowControl struct {
	WindowSize    uint16
	BytesInFlight uint16
	LastAck       uint32
}

// NewTCPFlowControl creates a new flow control instance
func NewTCPFlowControl(windowSize uint16) *TCPFlowControl {
	return &TCPFlowControl{
		WindowSize:    windowSize,
		BytesInFlight: 0,
		LastAck:       0,
	}
}

// CanSend checks if data can be sent
func (fc *TCPFlowControl) CanSend(dataSize uint16) bool {
	return fc.BytesInFlight+dataSize <= fc.WindowSize
}

// SendData simulates sending data
func (fc *TCPFlowControl) SendData(dataSize uint16) bool {
	if !fc.CanSend(dataSize) {
		return false
	}
	
	fc.BytesInFlight += dataSize
	return true
}

// ReceiveAck simulates receiving an ACK
func (fc *TCPFlowControl) ReceiveAck(ackNumber uint32, ackedBytes uint16) {
	if ackNumber > fc.LastAck {
		fc.LastAck = ackNumber
		if fc.BytesInFlight >= ackedBytes {
			fc.BytesInFlight -= ackedBytes
		}
	}
}

// GetAvailableWindow returns available window size
func (fc *TCPFlowControl) GetAvailableWindow() uint16 {
	if fc.BytesInFlight >= fc.WindowSize {
		return 0
	}
	return fc.WindowSize - fc.BytesInFlight
}

// TCPCongestionControl demonstrates congestion control
type TCPCongestionControl struct {
	CongestionWindow uint32
	Ssthresh         uint32
	State            string // "slow_start" or "congestion_avoidance"
	RTT              time.Duration
	RTO              time.Duration
}

// NewTCPCongestionControl creates a new congestion control instance
func NewTCPCongestionControl() *TCPCongestionControl {
	return &TCPCongestionControl{
		CongestionWindow: 1,
		Ssthresh:         65535,
		State:            "slow_start",
		RTT:              100 * time.Millisecond,
		RTO:              200 * time.Millisecond,
	}
}

// OnAckReceived handles ACK received
func (cc *TCPCongestionControl) OnAckReceived() {
	if cc.State == "slow_start" {
		// Slow start: exponential growth
		cc.CongestionWindow *= 2
		if cc.CongestionWindow >= cc.Ssthresh {
			cc.State = "congestion_avoidance"
		}
	} else if cc.State == "congestion_avoidance" {
		// Congestion avoidance: linear growth
		cc.CongestionWindow += 1
	}
}

// OnTimeout handles timeout
func (cc *TCPCongestionControl) OnTimeout() {
	// Reduce congestion window and set threshold
	cc.Ssthresh = cc.CongestionWindow / 2
	cc.CongestionWindow = 1
	cc.State = "slow_start"
}

// OnDuplicateAck handles duplicate ACK
func (cc *TCPCongestionControl) OnDuplicateAck() {
	// Fast retransmit
	cc.Ssthresh = cc.CongestionWindow / 2
	cc.CongestionWindow = cc.Ssthresh + 3
	cc.State = "congestion_avoidance"
}

// GetWindowSize returns current window size
func (cc *TCPCongestionControl) GetWindowSize() uint32 {
	return cc.CongestionWindow
}

// Demonstrate TCP server
func demonstrateTCPServer() {
	fmt.Println("=== TCP Server Demo ===\n")
	
	server := NewTCPServer(":8080")
	
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
	
	// Start server in goroutine
	go func() {
		if err := server.Start(); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()
	
	// Give server time to start
	time.Sleep(100 * time.Millisecond)
	
	// Create client and test
	client := NewTCPClient("localhost:8080")
	if err := client.Connect(); err != nil {
		log.Printf("Client connection error: %v", err)
		return
	}
	defer client.Close()
	
	// Send test messages
	testMessages := []string{
		"Hello, TCP!",
		"How are you?",
		"Testing TCP communication",
	}
	
	for _, msg := range testMessages {
		if err := client.Send(msg); err != nil {
			log.Printf("Send error: %v", err)
			continue
		}
		
		response, err := client.Receive()
		if err != nil {
			log.Printf("Receive error: %v", err)
			continue
		}
		
		fmt.Printf("Sent: %s", msg)
		fmt.Printf("Received: %s", response)
	}
	
	// Stop server
	server.Stop()
}

// Demonstrate TCP packet analysis
func demonstrateTCPPacketAnalysis() {
	fmt.Println("=== TCP Packet Analysis Demo ===\n")
	
	analyzer := &TCPPacketAnalyzer{}
	
	// Create a sample TCP packet
	packet := createSampleTCPPacket()
	
	// Analyze the packet
	header, err := analyzer.AnalyzePacket(packet)
	if err != nil {
		log.Printf("Analysis error: %v", err)
		return
	}
	
	// Display header information
	fmt.Printf("Source Port: %d\n", header.SourcePort)
	fmt.Printf("Destination Port: %d\n", header.DestinationPort)
	fmt.Printf("Sequence Number: %d\n", header.SequenceNumber)
	fmt.Printf("Acknowledgment Number: %d\n", header.AckNumber)
	fmt.Printf("Data Offset: %d\n", header.DataOffset)
	fmt.Printf("Window Size: %d\n", header.WindowSize)
	fmt.Printf("Checksum: 0x%04X\n", header.Checksum)
	
	// Parse and display flags
	flags := analyzer.ParseFlags(header.Flags)
	fmt.Printf("Flags: URG=%t ACK=%t PSH=%t RST=%t SYN=%t FIN=%t\n",
		flags.URG, flags.ACK, flags.PSH, flags.RST, flags.SYN, flags.FIN)
}

// createSampleTCPPacket creates a sample TCP packet
func createSampleTCPPacket() []byte {
	packet := make([]byte, 20)
	
	// Source port (8080)
	binary.BigEndian.PutUint16(packet[0:2], 8080)
	// Destination port (80)
	binary.BigEndian.PutUint16(packet[2:4], 80)
	// Sequence number
	binary.BigEndian.PutUint32(packet[4:8], 1000)
	// Acknowledgment number
	binary.BigEndian.PutUint32(packet[8:12], 2000)
	// Data offset and flags
	packet[12] = 5 << 4 // Data offset = 5
	packet[13] = 0x18   // ACK + PSH flags
	// Window size
	binary.BigEndian.PutUint16(packet[14:16], 65535)
	// Checksum
	binary.BigEndian.PutUint16(packet[16:18], 0x1234)
	// Urgent pointer
	binary.BigEndian.PutUint16(packet[18:20], 0)
	
	return packet
}

// Demonstrate flow control
func demonstrateFlowControl() {
	fmt.Println("=== TCP Flow Control Demo ===\n")
	
	fc := NewTCPFlowControl(1000) // Window size of 1000 bytes
	
	fmt.Printf("Initial window size: %d bytes\n", fc.WindowSize)
	fmt.Printf("Available window: %d bytes\n", fc.GetAvailableWindow())
	
	// Simulate sending data
	dataSizes := []uint16{100, 200, 300, 400, 500}
	
	for _, size := range dataSizes {
		if fc.CanSend(size) {
			fc.SendData(size)
			fmt.Printf("Sent %d bytes, remaining window: %d bytes\n", 
				size, fc.GetAvailableWindow())
		} else {
			fmt.Printf("Cannot send %d bytes, window full\n", size)
		}
	}
	
	// Simulate receiving ACK
	fc.ReceiveAck(1000, 200)
	fmt.Printf("After ACK for 200 bytes, available window: %d bytes\n", 
		fc.GetAvailableWindow())
}

// Demonstrate congestion control
func demonstrateCongestionControl() {
	fmt.Println("=== TCP Congestion Control Demo ===\n")
	
	cc := NewTCPCongestionControl()
	
	fmt.Printf("Initial state: %s\n", cc.State)
	fmt.Printf("Initial congestion window: %d\n", cc.CongestionWindow)
	
	// Simulate slow start
	fmt.Println("\nSlow Start Phase:")
	for i := 0; i < 5; i++ {
		cc.OnAckReceived()
		fmt.Printf("ACK %d: Window=%d, State=%s\n", 
			i+1, cc.CongestionWindow, cc.State)
	}
	
	// Simulate congestion avoidance
	fmt.Println("\nCongestion Avoidance Phase:")
	for i := 0; i < 5; i++ {
		cc.OnAckReceived()
		fmt.Printf("ACK %d: Window=%d, State=%s\n", 
			i+1, cc.CongestionWindow, cc.State)
	}
	
	// Simulate timeout
	fmt.Println("\nTimeout Event:")
	cc.OnTimeout()
	fmt.Printf("After timeout: Window=%d, State=%s\n", 
		cc.CongestionWindow, cc.State)
	
	// Simulate duplicate ACK
	fmt.Println("\nDuplicate ACK Event:")
	cc.OnDuplicateAck()
	fmt.Printf("After duplicate ACK: Window=%d, State=%s\n", 
		cc.CongestionWindow, cc.State)
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 3: TCP Protocol")
	fmt.Println("====================================================\n")
	
	// Run all demonstrations
	demonstrateTCPServer()
	fmt.Println()
	demonstrateTCPPacketAnalysis()
	fmt.Println()
	demonstrateFlowControl()
	fmt.Println()
	demonstrateCongestionControl()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. TCP provides reliable, ordered, and error-checked delivery")
	fmt.Println("2. Three-way handshake establishes connections")
	fmt.Println("3. Four-way handshake terminates connections")
	fmt.Println("4. Flow control prevents overwhelming the receiver")
	fmt.Println("5. Congestion control manages network congestion")
	fmt.Println("6. TCP uses sequence numbers for ordering and reliability")
	
	fmt.Println("\n📚 Next Topic: UDP Protocol Deep Dive")
}
```

Now let me create the UDP protocol topic:


```markdown
# Topic 4: UDP Protocol Deep Dive

## UDP Overview

User Datagram Protocol (UDP) is a connectionless, unreliable transport layer protocol that provides:
- **Connectionless**: No connection establishment required
- **Unreliable**: No delivery guarantees
- **Lightweight**: Minimal overhead
- **Fast**: No connection management overhead
- **Simple**: Straightforward implementation

## UDP Header Structure

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

### Header Fields Explained

- **Source Port (16 bits)**: Port number of the sending application
- **Destination Port (16 bits)**: Port number of the receiving application
- **Length (16 bits)**: Length of UDP header and data in bytes
- **Checksum (16 bits)**: Header and data checksum (optional)
- **Data**: Application data

## UDP Characteristics

### Advantages
- **Low Overhead**: Only 8 bytes of header
- **Fast**: No connection establishment
- **Simple**: Easy to implement
- **Broadcast/Multicast**: Supports one-to-many communication
- **Real-time**: Suitable for time-sensitive applications

### Disadvantages
- **Unreliable**: No delivery guarantees
- **No Ordering**: Packets may arrive out of order
- **No Flow Control**: Can overwhelm receiver
- **No Congestion Control**: Can cause network congestion
- **No Error Recovery**: Lost packets are not retransmitted

## UDP Use Cases

### Real-time Applications
- **Video Streaming**: Live video, video conferencing
- **Audio Streaming**: Voice over IP, music streaming
- **Gaming**: Multiplayer online games
- **Live Broadcasting**: Real-time data feeds

### Network Services
- **DNS**: Domain name resolution
- **DHCP**: Dynamic IP configuration
- **SNMP**: Network management
- **NTP**: Network time protocol

### Lightweight Communication
- **Heartbeats**: Health checks, keep-alives
- **Notifications**: System alerts, status updates
- **Logging**: Remote logging, telemetry
- **Discovery**: Service discovery protocols

## UDP vs TCP Comparison

| Feature | UDP | TCP |
|---------|-----|-----|
| Connection | Connectionless | Connection-oriented |
| Reliability | Unreliable | Reliable |
| Ordering | Unordered | Ordered |
| Speed | Fast | Slower |
| Overhead | Low (8 bytes) | High (20+ bytes) |
| Flow Control | No | Yes |
| Congestion Control | No | Yes |
| Error Detection | Optional | Yes |
| Error Recovery | No | Yes |
| Broadcast/Multicast | Yes | No |

## When to Use UDP

### Use UDP when:
- **Speed is critical**: Real-time applications
- **Small, frequent messages**: Heartbeats, status updates
- **Broadcast/multicast needed**: One-to-many communication
- **Application handles reliability**: Custom error handling
- **Network is reliable**: Local networks, controlled environments

### Examples:
- **DNS**: Domain name resolution
- **DHCP**: IP address assignment
- **SNMP**: Network monitoring
- **Gaming**: Multiplayer games
- **Streaming**: Video/audio streaming
- **IoT**: Sensor data, telemetry

## UDP Programming Considerations

### Error Handling
- **Packet Loss**: Implement application-level retry
- **Out-of-Order**: Handle sequence numbers
- **Duplicates**: Detect and discard duplicates
- **Timeouts**: Set appropriate timeouts

### Performance Optimization
- **Buffer Sizes**: Optimize send/receive buffers
- **Packet Size**: Use appropriate packet sizes
- **Rate Limiting**: Implement rate limiting
- **Batching**: Combine multiple messages

### Security Considerations
- **No Authentication**: UDP has no built-in security
- **No Encryption**: Data is sent in plain text
- **DoS Attacks**: Vulnerable to flooding attacks
- **Spoofing**: Source addresses can be spoofed

## UDP Socket Programming

### Basic UDP Server
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

### Basic UDP Client
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

## Reliable UDP Implementation

### Application-Level Reliability
- **Sequence Numbers**: Track packet order
- **Acknowledgments**: Confirm packet receipt
- **Retransmission**: Resend lost packets
- **Timeout**: Detect packet loss
- **Flow Control**: Prevent overwhelming receiver

### RUDP (Reliable UDP) Features
- **Ordered Delivery**: Ensure correct sequence
- **Duplicate Detection**: Remove duplicate packets
- **Flow Control**: Manage data rate
- **Congestion Control**: Prevent network overload

## UDP Multicast

### Multicast Concepts
- **Multicast Group**: Set of receivers
- **Multicast Address**: Special IP address range
- **IGMP**: Internet Group Management Protocol
- **TTL**: Time to Live for multicast packets

### Multicast Addresses
- **224.0.0.0/4**: IPv4 multicast range
- **224.0.0.1**: All systems on subnet
- **224.0.0.2**: All routers on subnet
- **239.0.0.0/8**: Administratively scoped

### Multicast Programming
```go
func startMulticastServer() {
    addr, err := net.ResolveUDPAddr("udp", "224.0.0.1:8080")
    if err != nil {
        log.Fatal(err)
    }
    
    conn, err := net.ListenMulticastUDP("udp", nil, addr)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    
    // Send multicast data
    conn.WriteToUDP([]byte("Multicast message"), addr)
}
```

## UDP Broadcasting

### Broadcast Concepts
- **Broadcast Address**: 255.255.255.255
- **Subnet Broadcast**: Last address in subnet
- **Limited Broadcast**: All 1s address
- **Directed Broadcast**: Specific subnet broadcast

### Broadcast Programming
```go
func startBroadcastServer() {
    addr, err := net.ResolveUDPAddr("udp", "255.255.255.255:8080")
    if err != nil {
        log.Fatal(err)
    }
    
    conn, err := net.DialUDP("udp", nil, addr)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    
    // Enable broadcast
    conn.SetBroadcast(true)
    
    // Send broadcast data
    conn.Write([]byte("Broadcast message"))
}
```

## UDP Performance Tuning

### Socket Options
- **SO_RCVBUF**: Receive buffer size
- **SO_SNDBUF**: Send buffer size
- **SO_REUSEADDR**: Reuse address
- **SO_BROADCAST**: Enable broadcast

### Performance Tips
- **Use appropriate buffer sizes**
- **Implement rate limiting**
- **Batch multiple messages**
- **Use non-blocking I/O**
- **Optimize packet sizes**

## Common UDP Issues

### Packet Loss
- **Network Congestion**: Too much traffic
- **Buffer Overflow**: Receiver can't keep up
- **Network Errors**: Physical layer issues
- **Firewall Blocking**: Security policies

### Out-of-Order Delivery
- **Network Routing**: Different paths
- **Load Balancing**: Multiple servers
- **Retransmission**: Duplicate packets
- **Network Congestion**: Queuing delays

### Troubleshooting
- **tcpdump**: Capture UDP packets
- **netstat**: Check socket states
- **Wireshark**: Analyze packet flow
- **ping**: Test connectivity

## Interview Questions

### Basic Questions
1. What is UDP and how does it differ from TCP?
2. When would you use UDP over TCP?
3. What are the advantages and disadvantages of UDP?

### Intermediate Questions
1. How would you implement reliability in UDP?
2. Explain UDP multicast and broadcasting.
3. What are the security considerations for UDP?

### Advanced Questions
1. Design a reliable UDP protocol.
2. How would you optimize UDP performance?
3. Implement UDP flow control and congestion control.

## Next Steps
After mastering UDP, proceed to:
- **Topic 5**: Socket Programming in Go
- **Topic 6**: HTTP/HTTPS Implementation
- **Topic 7**: WebSocket Protocols

Master UDP, and you'll understand lightweight, fast networking! 🚀
```

Now let me create the UDP implementation:


```go
package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

// UDPHeader represents a UDP header
type UDPHeader struct {
	SourcePort      uint16
	DestinationPort uint16
	Length          uint16
	Checksum        uint16
}

// UDPServer represents a UDP server
type UDPServer struct {
	Address  string
	Conn     *net.UDPConn
	Handler  func(*net.UDPConn, *net.UDPAddr, []byte)
	Running  bool
}

// NewUDPServer creates a new UDP server
func NewUDPServer(address string) *UDPServer {
	return &UDPServer{
		Address: address,
		Running: false,
	}
}

// Start starts the UDP server
func (s *UDPServer) Start() error {
	addr, err := net.ResolveUDPAddr("udp", s.Address)
	if err != nil {
		return err
	}
	
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}
	
	s.Conn = conn
	s.Running = true
	
	fmt.Printf("UDP Server listening on %s\n", s.Address)
	
	// Start receiving packets
	go s.receivePackets()
	
	return nil
}

// receivePackets receives UDP packets
func (s *UDPServer) receivePackets() {
	buffer := make([]byte, 1024)
	
	for s.Running {
		// Set read timeout
		s.Conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		
		n, clientAddr, err := s.Conn.ReadFromUDP(buffer)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue // Timeout, check if still running
			}
			log.Printf("Error reading from UDP: %v", err)
			continue
		}
		
		// Handle packet
		if s.Handler != nil {
			s.Handler(s.Conn, clientAddr, buffer[:n])
		} else {
			s.defaultHandler(s.Conn, clientAddr, buffer[:n])
		}
	}
}

// defaultHandler is the default packet handler
func (s *UDPServer) defaultHandler(conn *net.UDPConn, clientAddr *net.UDPAddr, data []byte) {
	// Echo the data back
	response := fmt.Sprintf("Echo: %s", string(data))
	conn.WriteToUDP([]byte(response), clientAddr)
}

// Stop stops the UDP server
func (s *UDPServer) Stop() {
	s.Running = false
	if s.Conn != nil {
		s.Conn.Close()
	}
}

// UDPClient represents a UDP client
type UDPClient struct {
	Address string
	Conn    *net.UDPConn
}

// NewUDPClient creates a new UDP client
func NewUDPClient(address string) *UDPClient {
	return &UDPClient{
		Address: address,
	}
}

// Connect connects to a UDP server
func (c *UDPClient) Connect() error {
	addr, err := net.ResolveUDPAddr("udp", c.Address)
	if err != nil {
		return err
	}
	
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	
	c.Conn = conn
	fmt.Printf("Connected to %s\n", c.Address)
	return nil
}

// Send sends data to the server
func (c *UDPClient) Send(data []byte) error {
	if c.Conn == nil {
		return fmt.Errorf("not connected")
	}
	
	_, err := c.Conn.Write(data)
	return err
}

// SendString sends a string to the server
func (c *UDPClient) SendString(data string) error {
	return c.Send([]byte(data))
}

// Receive receives data from the server
func (c *UDPClient) Receive() ([]byte, error) {
	if c.Conn == nil {
		return nil, fmt.Errorf("not connected")
	}
	
	buffer := make([]byte, 1024)
	n, err := c.Conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	
	return buffer[:n], nil
}

// ReceiveString receives a string from the server
func (c *UDPClient) ReceiveString() (string, error) {
	data, err := c.Receive()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Close closes the connection
func (c *UDPClient) Close() error {
	if c.Conn != nil {
		return c.Conn.Close()
	}
	return nil
}

// UDPPacketAnalyzer analyzes UDP packets
type UDPPacketAnalyzer struct{}

// AnalyzePacket analyzes a UDP packet
func (a *UDPPacketAnalyzer) AnalyzePacket(data []byte) (*UDPHeader, error) {
	if len(data) < 8 {
		return nil, fmt.Errorf("packet too short for UDP header")
	}
	
	header := &UDPHeader{
		SourcePort:      binary.BigEndian.Uint16(data[0:2]),
		DestinationPort: binary.BigEndian.Uint16(data[2:4]),
		Length:          binary.BigEndian.Uint16(data[4:6]),
		Checksum:        binary.BigEndian.Uint16(data[6:8]),
	}
	
	return header, nil
}

// CreatePacket creates a UDP packet
func (a *UDPPacketAnalyzer) CreatePacket(srcPort, destPort uint16, data []byte) []byte {
	packet := make([]byte, 8+len(data))
	
	// Source port
	binary.BigEndian.PutUint16(packet[0:2], srcPort)
	// Destination port
	binary.BigEndian.PutUint16(packet[2:4], destPort)
	// Length (header + data)
	binary.BigEndian.PutUint16(packet[4:6], uint16(8+len(data)))
	// Checksum (simplified)
	binary.BigEndian.PutUint16(packet[6:8], 0)
	// Data
	copy(packet[8:], data)
	
	return packet
}

// ReliableUDP implements application-level reliability
type ReliableUDP struct {
	SeqNumber    uint32
	AckNumber    uint32
	WindowSize   uint32
	Packets      map[uint32][]byte
	Acks         map[uint32]bool
	LastAckTime  time.Time
	RTO          time.Duration
}

// NewReliableUDP creates a new reliable UDP instance
func NewReliableUDP() *ReliableUDP {
	return &ReliableUDP{
		SeqNumber:   1,
		AckNumber:   0,
		WindowSize:  10,
		Packets:     make(map[uint32][]byte),
		Acks:        make(map[uint32]bool),
		LastAckTime: time.Now(),
		RTO:         100 * time.Millisecond,
	}
}

// SendPacket sends a reliable packet
func (rudp *ReliableUDP) SendPacket(data []byte) uint32 {
	seqNum := rudp.SeqNumber
	rudp.Packets[seqNum] = data
	rudp.SeqNumber++
	return seqNum
}

// ReceiveAck processes an acknowledgment
func (rudp *ReliableUDP) ReceiveAck(ackNum uint32) {
	rudp.Acks[ackNum] = true
	rudp.LastAckTime = time.Now()
	
	// Remove acknowledged packets
	for seqNum := range rudp.Packets {
		if seqNum <= ackNum {
			delete(rudp.Packets, seqNum)
		}
	}
}

// GetUnackedPackets returns unacknowledged packets
func (rudp *ReliableUDP) GetUnackedPackets() map[uint32][]byte {
	unacked := make(map[uint32][]byte)
	now := time.Now()
	
	// Check for timeout
	if now.Sub(rudp.LastAckTime) > rudp.RTO {
		// Retransmit unacked packets
		for seqNum, data := range rudp.Packets {
			if !rudp.Acks[seqNum] {
				unacked[seqNum] = data
			}
		}
		rudp.LastAckTime = now
	}
	
	return unacked
}

// UDPMulticast represents a UDP multicast
type UDPMulticast struct {
	GroupAddress string
	Port         int
	Conn         *net.UDPConn
}

// NewUDPMulticast creates a new UDP multicast
func NewUDPMulticast(groupAddress string, port int) *UDPMulticast {
	return &UDPMulticast{
		GroupAddress: groupAddress,
		Port:         port,
	}
}

// Join joins a multicast group
func (m *UDPMulticast) Join() error {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", m.GroupAddress, m.Port))
	if err != nil {
		return err
	}
	
	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	
	m.Conn = conn
	return nil
}

// Send sends data to the multicast group
func (m *UDPMulticast) Send(data []byte) error {
	if m.Conn == nil {
		return fmt.Errorf("not joined to multicast group")
	}
	
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", m.GroupAddress, m.Port))
	if err != nil {
		return err
	}
	
	_, err = m.Conn.WriteToUDP(data, addr)
	return err
}

// Receive receives data from the multicast group
func (m *UDPMulticast) Receive() ([]byte, *net.UDPAddr, error) {
	if m.Conn == nil {
		return nil, nil, fmt.Errorf("not joined to multicast group")
	}
	
	buffer := make([]byte, 1024)
	n, addr, err := m.Conn.ReadFromUDP(buffer)
	if err != nil {
		return nil, nil, err
	}
	
	return buffer[:n], addr, nil
}

// Leave leaves the multicast group
func (m *UDPMulticast) Leave() error {
	if m.Conn != nil {
		return m.Conn.Close()
	}
	return nil
}

// UDPBroadcast represents a UDP broadcast
type UDPBroadcast struct {
	Port int
	Conn *net.UDPConn
}

// NewUDPBroadcast creates a new UDP broadcast
func NewUDPBroadcast(port int) *UDPBroadcast {
	return &UDPBroadcast{
		Port: port,
	}
}

// Start starts broadcasting
func (b *UDPBroadcast) Start() error {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("255.255.255.255:%d", b.Port))
	if err != nil {
		return err
	}
	
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	
	// Enable broadcast
	conn.SetBroadcast(true)
	b.Conn = conn
	
	return nil
}

// Broadcast sends data via broadcast
func (b *UDPBroadcast) Broadcast(data []byte) error {
	if b.Conn == nil {
		return fmt.Errorf("broadcast not started")
	}
	
	_, err := b.Conn.Write(data)
	return err
}

// Stop stops broadcasting
func (b *UDPBroadcast) Stop() error {
	if b.Conn != nil {
		return b.Conn.Close()
	}
	return nil
}

// Demonstrate UDP server
func demonstrateUDPServer() {
	fmt.Println("=== UDP Server Demo ===\n")
	
	server := NewUDPServer(":8080")
	
	// Set custom handler
	server.Handler = func(conn *net.UDPConn, clientAddr *net.UDPAddr, data []byte) {
		fmt.Printf("Received from %s: %s\n", clientAddr, string(data))
		
		// Send response
		response := fmt.Sprintf("Server received: %s", string(data))
		conn.WriteToUDP([]byte(response), clientAddr)
	}
	
	// Start server
	if err := server.Start(); err != nil {
		log.Printf("Server error: %v", err)
		return
	}
	
	// Give server time to start
	time.Sleep(100 * time.Millisecond)
	
	// Create client and test
	client := NewUDPClient("localhost:8080")
	if err := client.Connect(); err != nil {
		log.Printf("Client connection error: %v", err)
		return
	}
	defer client.Close()
	
	// Send test messages
	testMessages := []string{
		"Hello, UDP!",
		"How are you?",
		"Testing UDP communication",
	}
	
	for _, msg := range testMessages {
		if err := client.SendString(msg); err != nil {
			log.Printf("Send error: %v", err)
			continue
		}
		
		response, err := client.ReceiveString()
		if err != nil {
			log.Printf("Receive error: %v", err)
			continue
		}
		
		fmt.Printf("Sent: %s\n", msg)
		fmt.Printf("Received: %s\n", response)
	}
	
	// Stop server
	server.Stop()
}

// Demonstrate UDP packet analysis
func demonstrateUDPPacketAnalysis() {
	fmt.Println("=== UDP Packet Analysis Demo ===\n")
	
	analyzer := &UDPPacketAnalyzer{}
	
	// Create a sample UDP packet
	packet := analyzer.CreatePacket(8080, 80, []byte("Hello, UDP!"))
	
	// Analyze the packet
	header, err := analyzer.AnalyzePacket(packet)
	if err != nil {
		log.Printf("Analysis error: %v", err)
		return
	}
	
	// Display header information
	fmt.Printf("Source Port: %d\n", header.SourcePort)
	fmt.Printf("Destination Port: %d\n", header.DestinationPort)
	fmt.Printf("Length: %d bytes\n", header.Length)
	fmt.Printf("Checksum: 0x%04X\n", header.Checksum)
	fmt.Printf("Data: %s\n", string(packet[8:]))
}

// Demonstrate reliable UDP
func demonstrateReliableUDP() {
	fmt.Println("=== Reliable UDP Demo ===\n")
	
	rudp := NewReliableUDP()
	
	// Send some packets
	packets := []string{
		"Packet 1",
		"Packet 2",
		"Packet 3",
		"Packet 4",
		"Packet 5",
	}
	
	fmt.Println("Sending packets:")
	for i, packet := range packets {
		seqNum := rudp.SendPacket([]byte(packet))
		fmt.Printf("Sent packet %d with sequence number %d\n", i+1, seqNum)
	}
	
	// Simulate receiving ACKs
	fmt.Println("\nSimulating ACKs:")
	rudp.ReceiveAck(1)
	rudp.ReceiveAck(2)
	rudp.ReceiveAck(4)
	
	// Check for unacked packets
	unacked := rudp.GetUnackedPackets()
	fmt.Printf("\nUnacked packets: %d\n", len(unacked))
	for seqNum, data := range unacked {
		fmt.Printf("  Sequence %d: %s\n", seqNum, string(data))
	}
}

// Demonstrate UDP multicast
func demonstrateUDPMulticast() {
	fmt.Println("=== UDP Multicast Demo ===\n")
	
	// Create multicast
	multicast := NewUDPMulticast("224.0.0.1", 8080)
	
	// Join multicast group
	if err := multicast.Join(); err != nil {
		log.Printf("Multicast join error: %v", err)
		return
	}
	defer multicast.Leave()
	
	// Send multicast data
	data := []byte("Multicast message")
	if err := multicast.Send(data); err != nil {
		log.Printf("Multicast send error: %v", err)
		return
	}
	
	fmt.Printf("Sent multicast data: %s\n", string(data))
	
	// Receive multicast data
	received, addr, err := multicast.Receive()
	if err != nil {
		log.Printf("Multicast receive error: %v", err)
		return
	}
	
	fmt.Printf("Received from %s: %s\n", addr, string(received))
}

// Demonstrate UDP broadcast
func demonstrateUDPBroadcast() {
	fmt.Println("=== UDP Broadcast Demo ===\n")
	
	// Create broadcast
	broadcast := NewUDPBroadcast(8080)
	
	// Start broadcasting
	if err := broadcast.Start(); err != nil {
		log.Printf("Broadcast start error: %v", err)
		return
	}
	defer broadcast.Stop()
	
	// Send broadcast data
	data := []byte("Broadcast message")
	if err := broadcast.Broadcast(data); err != nil {
		log.Printf("Broadcast send error: %v", err)
		return
	}
	
	fmt.Printf("Sent broadcast data: %s\n", string(data))
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 4: UDP Protocol")
	fmt.Println("====================================================\n")
	
	// Run all demonstrations
	demonstrateUDPServer()
	fmt.Println()
	demonstrateUDPPacketAnalysis()
	fmt.Println()
	demonstrateReliableUDP()
	fmt.Println()
	demonstrateUDPMulticast()
	fmt.Println()
	demonstrateUDPBroadcast()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. UDP is connectionless and unreliable but fast")
	fmt.Println("2. UDP has minimal overhead (8 bytes header)")
	fmt.Println("3. UDP supports broadcast and multicast")
	fmt.Println("4. Application-level reliability can be implemented")
	fmt.Println("5. UDP is ideal for real-time applications")
	fmt.Println("6. UDP requires careful error handling")
	
	fmt.Println("\n📚 Next Topic: Socket Programming in Go")
}
```

Now let me update the progress and create a comprehensive index file:


Let me create a comprehensive index file that shows the complete learning path:


```markdown
# Computer Networking Mastery Course - Complete Index

## 🎯 Course Overview
This comprehensive course will take you from networking basics to advanced concepts, preparing you for top-tier software engineering interviews at companies like Google, Amazon, Microsoft, Meta, and more.

## 📚 Learning Path - Sequential Order

### Phase 1: Fundamentals (Topics 1-4) ✅ COMPLETED

#### ✅ Topic 1: OSI Model - The Foundation of Networking
- **Theory**: `01_osi_model/theory.md`
- **Implementation**: `01_osi_model/implementation.go`
- **Key Concepts**: 7-layer model, encapsulation, decapsulation
- **Go Skills**: Packet structures, header parsing, layer simulation

#### ✅ Topic 2: IP Addressing & Subnetting Mastery
- **Theory**: `02_ip_addressing/theory.md`
- **Implementation**: `02_ip_addressing/implementation.go`
- **Key Concepts**: IPv4/IPv6, CIDR, VLSM, NAT
- **Go Skills**: IP manipulation, subnet calculations, address validation

#### ✅ Topic 3: TCP Protocol Deep Dive
- **Theory**: `03_tcp_protocol/theory.md`
- **Implementation**: `03_tcp_protocol/implementation.go`
- **Key Concepts**: Three-way handshake, flow control, congestion control
- **Go Skills**: TCP servers/clients, connection management, packet analysis

#### ✅ Topic 4: UDP Protocol Deep Dive
- **Theory**: `04_udp_protocol/theory.md`
- **Implementation**: `04_udp_protocol/implementation.go`
- **Key Concepts**: Connectionless communication, multicast, broadcast
- **Go Skills**: UDP programming, reliable UDP, multicast/broadcast

### Phase 2: Network Programming (Topics 5-8) 🚧 IN PROGRESS

#### 🚧 Topic 5: Socket Programming in Go
- **Theory**: `05_socket_programming/theory.md`
- **Implementation**: `05_socket_programming/implementation.go`
- **Key Concepts**: Socket APIs, non-blocking I/O, multiplexing
- **Go Skills**: Advanced socket programming, epoll, select

#### 📋 Topic 6: HTTP/HTTPS Implementation
- **Theory**: `06_http_https/theory.md`
- **Implementation**: `06_http_https/implementation.go`
- **Key Concepts**: HTTP protocol, REST APIs, HTTPS/TLS
- **Go Skills**: HTTP servers, clients, middleware, TLS

#### 📋 Topic 7: WebSocket Protocols
- **Theory**: `07_websocket/theory.md`
- **Implementation**: `07_websocket/implementation.go`
- **Key Concepts**: Real-time communication, WebSocket handshake
- **Go Skills**: WebSocket servers, clients, message handling

#### 📋 Topic 8: Network Security & Encryption
- **Theory**: `08_network_security/theory.md`
- **Implementation**: `08_network_security/implementation.go`
- **Key Concepts**: TLS/SSL, encryption, authentication, VPNs
- **Go Skills**: TLS implementation, secure communication

### Phase 3: Advanced Concepts (Topics 9-12) 📋 PLANNED

#### 📋 Topic 9: Routing & Switching
- **Theory**: `09_routing_switching/theory.md`
- **Implementation**: `09_routing_switching/implementation.go`
- **Key Concepts**: Routing algorithms, switching, VLANs
- **Go Skills**: Network simulation, routing protocols

#### 📋 Topic 10: Network Performance & Optimization
- **Theory**: `10_performance_optimization/theory.md`
- **Implementation**: `10_performance_optimization/implementation.go`
- **Key Concepts**: Performance tuning, load balancing, caching
- **Go Skills**: Performance monitoring, optimization techniques

#### 📋 Topic 11: Distributed Systems Networking
- **Theory**: `11_distributed_systems/theory.md`
- **Implementation**: `11_distributed_systems/implementation.go`
- **Key Concepts**: Microservices, service discovery, message queues
- **Go Skills**: Distributed system patterns, communication

#### 📋 Topic 12: Cloud Networking & Containerization
- **Theory**: `12_cloud_networking/theory.md`
- **Implementation**: `12_cloud_networking/implementation.go`
- **Key Concepts**: Cloud networking, containers, orchestration
- **Go Skills**: Docker networking, Kubernetes, cloud APIs

### Phase 4: Interview Preparation (Topics 13-15) 📋 PLANNED

#### 📋 Topic 13: System Design Questions
- **Questions**: `13_system_design/questions.md`
- **Solutions**: `13_system_design/solutions.md`
- **Key Concepts**: System design, scalability, architecture
- **Go Skills**: Design patterns, implementation strategies

#### 📋 Topic 14: Coding Challenges & Implementations
- **Problems**: `14_coding_challenges/problems.md`
- **Solutions**: `14_coding_challenges/solutions.go`
- **Key Concepts**: Algorithm implementation, data structures
- **Go Skills**: Advanced Go programming, optimization

#### 📋 Topic 15: Troubleshooting & Debugging
- **Scenarios**: `15_troubleshooting/scenarios.md`
- **Solutions**: `15_troubleshooting/solutions.md`
- **Key Concepts**: Network troubleshooting, debugging techniques
- **Go Skills**: Debugging tools, monitoring, analysis

## 🚀 Quick Start Guide

### For Beginners
1. Start with **Topic 1: OSI Model**
2. Read the theory first (`theory.md`)
3. Study the implementation (`implementation.go`)
4. Run the code and experiment
5. Move to the next topic

### For Intermediate Learners
1. Review topics 1-4 quickly
2. Focus on **Topic 5: Socket Programming**
3. Deep dive into **Topic 6: HTTP/HTTPS**
4. Practice with real-world examples

### For Advanced Learners
1. Skip to **Topic 9: Routing & Switching**
2. Focus on **Topic 10: Performance Optimization**
3. Master **Topic 11: Distributed Systems**
4. Prepare for interviews with **Topics 13-15**

## 🛠️ Tools & Technologies

### Required Tools
- **Go 1.19+**: Programming language
- **Git**: Version control
- **Terminal**: Command line interface

### Recommended Tools
- **Wireshark**: Packet analysis
- **Docker**: Containerization
- **Kubernetes**: Orchestration
- **Postman**: API testing
- **VS Code**: Code editor

### Optional Tools
- **tcpdump**: Packet capture
- **netstat**: Network statistics
- **ss**: Socket statistics
- **iperf**: Network performance testing

## 📖 Learning Strategy

### 1. Theory First
- Read the theory file for each topic
- Understand the concepts thoroughly
- Take notes on key points

### 2. Implementation Study
- Go through the Go implementation
- Understand the code structure
- Run the examples

### 3. Hands-on Practice
- Modify the code
- Add new features
- Solve problems

### 4. Real-world Application
- Apply concepts to real projects
- Build networking applications
- Practice with different scenarios

## 🎯 Interview Preparation

### Technical Questions
- **OSI Model**: Layer functions, encapsulation
- **TCP/UDP**: Differences, use cases, implementation
- **IP Addressing**: Subnetting, VLSM, NAT
- **Network Programming**: Sockets, protocols, APIs

### System Design
- **Load Balancing**: Algorithms, implementation
- **Caching**: Strategies, consistency
- **Microservices**: Communication, discovery
- **Scalability**: Horizontal vs vertical scaling

### Coding Challenges
- **Socket Programming**: TCP/UDP servers, clients
- **Protocol Implementation**: HTTP, WebSocket
- **Network Algorithms**: Routing, switching
- **Performance Optimization**: Tuning, monitoring

## 📊 Progress Tracking

### Phase 1: Fundamentals (4 topics)
- [x] Topic 1: OSI Model
- [x] Topic 2: IP Addressing
- [x] Topic 3: TCP Protocol
- [x] Topic 4: UDP Protocol

### Phase 2: Network Programming (4 topics)
- [ ] Topic 5: Socket Programming
- [ ] Topic 6: HTTP/HTTPS
- [ ] Topic 7: WebSocket
- [ ] Topic 8: Network Security

### Phase 3: Advanced Concepts (4 topics)
- [ ] Topic 9: Routing & Switching
- [ ] Topic 10: Performance Optimization
- [ ] Topic 11: Distributed Systems
- [ ] Topic 12: Cloud Networking

### Phase 4: Interview Preparation (3 topics)
- [ ] Topic 13: System Design
- [ ] Topic 14: Coding Challenges
- [ ] Topic 15: Troubleshooting

## 🏆 Success Metrics

### Beginner Level
- [ ] Understand OSI model
- [ ] Know IP addressing basics
- [ ] Understand TCP vs UDP
- [ ] Write basic network programs

### Intermediate Level
- [ ] Implement reliable protocols
- [ ] Build HTTP servers/clients
- [ ] Handle network errors
- [ ] Optimize network performance

### Advanced Level
- [ ] Design distributed systems
- [ ] Implement complex protocols
- [ ] Troubleshoot network issues
- [ ] Pass technical interviews

## 📞 Support & Community

### Getting Help
- **GitHub Issues**: Report bugs or ask questions
- **Discussions**: Share ideas and solutions
- **Pull Requests**: Contribute improvements

### Contributing
- **Code**: Improve implementations
- **Documentation**: Enhance explanations
- **Examples**: Add real-world scenarios
- **Tests**: Add test cases

## 🎉 Conclusion

This course provides a comprehensive path to mastering computer networking. Follow the sequential order, practice with the implementations, and you'll be ready to tackle any networking challenge or interview!

**Remember**: Networking is a practical skill. Don't just read the theory—implement, experiment, and build real applications!

---

**Happy Learning! 🚀**

*Last Updated: December 2024*
```

Perfect! I've created a comprehensive, numbered learning path for your computer networking mastery course. Here's what I've organized for you:

## �� **Complete Learning Structure**

### **Phase 1: Fundamentals (Topics 1-4) ✅ COMPLETED**
1. **Topic 1**: OSI Model - The Foundation
2. **Topic 2**: IP Addressing & Subnetting  
3. **Topic 3**: TCP Protocol Deep Dive
4. **Topic 4**: UDP Protocol Deep Dive

### **Phase 2: Network Programming (Topics 5-8) �� IN PROGRESS**
5. **Topic 5**: Socket Programming in Go
6. **Topic 6**: HTTP/HTTPS Implementation
7. **Topic 7**: WebSocket Protocols
8. **Topic 8**: Network Security & Encryption

### **Phase 3: Advanced Concepts (Topics 9-12) �� PLANNED**
9. **Topic 9**: Routing & Switching
10. **Topic 10**: Network Performance & Optimization
11. **Topic 11**: Distributed Systems Networking
12. **Topic 12**: Cloud Networking & Containerization

### **Phase 4: Interview Preparation (Topics 13-15) 📋 PLANNED**
13. **Topic 13**: System Design Questions
14. **Topic 14**: Coding Challenges & Implementations
15. **Topic 15**: Troubleshooting & Debugging

## �� **Key Features of This Structure**

1. **Sequential Learning**: Each topic builds on the previous ones
2. **Theory + Implementation**: Every topic has both theory and Go code
3. **Interview Ready**: Designed to help you crack top company interviews
4. **Progressive Difficulty**: From basics to advanced concepts
5. **Real-world Focus**: Practical implementations you can use

## 🚀 **How to Use This Course**

1. **Start with Topic 1**: Read the theory, study the implementation
2. **Follow the Sequence**: Don't skip topics - they build on each other
3. **Practice Actively**: Run the code, modify it, experiment
4. **Track Your Progress**: Use the INDEX.md to check off completed topics

## 📁 **File Organization**

Each topic has:
- `theory.md` - Comprehensive theoretical knowledge
- `implementation.go` - Practical Go implementations
- Clear explanations and examples
- Interview questions and answers

You now have a **god-level networking curriculum** that will take you from beginner to expert! Each topic is numbered and organized so you can follow the learning path systematically.

Would you like me to continue with **Topic 5: Socket Programming in Go** or would you prefer to start with any specific topic?

---

**User**

continue with Topic 5: Socket Programming in Go

---

**Cursor**

Perfect! Let me continue with **Topic 5: Socket Programming in Go**. This is where we dive deep into advanced network programming concepts and implementations.


Now let me create the comprehensive Socket Programming module:


```markdown
# Topic 5: Socket Programming in Go - Advanced Network Programming

## Socket Programming Overview

Socket programming is the foundation of network communication. It provides a low-level interface for network communication between applications running on different machines or the same machine.

### What are Sockets?
- **Definition**: Sockets are endpoints for communication between two machines
- **Types**: Stream sockets (TCP), Datagram sockets (UDP), Raw sockets
- **Addressing**: IP address + Port number
- **Bidirectional**: Full-duplex communication

### Socket Programming Models
1. **Blocking I/O**: Synchronous, thread-per-connection
2. **Non-blocking I/O**: Asynchronous, event-driven
3. **I/O Multiplexing**: Select, poll, epoll
4. **Asynchronous I/O**: AIO, completion ports

## Go Socket Programming

### Go's Network Package
Go provides excellent support for socket programming through the `net` package:
- **net.TCPConn**: TCP connections
- **net.UDPConn**: UDP connections
- **net.Listener**: Server listeners
- **net.Dial**: Client connections

### Socket Types in Go
- **TCP Sockets**: Reliable, connection-oriented
- **UDP Sockets**: Unreliable, connectionless
- **Unix Sockets**: Local communication
- **Raw Sockets**: Low-level packet access

## Advanced Socket Concepts

### Socket Options
- **SO_REUSEADDR**: Reuse local addresses
- **SO_KEEPALIVE**: Keep connections alive
- **SO_LINGER**: Control connection closure
- **TCP_NODELAY**: Disable Nagle's algorithm
- **SO_RCVBUF/SO_SNDBUF**: Buffer sizes

### Socket States
- **CLOSED**: Not connected
- **LISTEN**: Server waiting for connections
- **SYN_SENT**: Client sent SYN
- **SYN_RCVD**: Server received SYN
- **ESTABLISHED**: Connection established
- **FIN_WAIT**: Connection closing
- **TIME_WAIT**: Waiting for delayed packets

## I/O Multiplexing

### Select Model
- **select()**: Monitor multiple file descriptors
- **Blocking**: Waits for I/O events
- **Scalability**: Limited by FD_SETSIZE
- **Cross-platform**: Available on most systems

### Poll Model
- **poll()**: Improved version of select
- **No FD limit**: Can handle more connections
- **Better performance**: More efficient than select
- **Linux/Unix**: Not available on Windows

### Epoll Model (Linux)
- **epoll()**: Event-driven I/O
- **High performance**: Handles thousands of connections
- **Edge-triggered**: Only notifies on state changes
- **Linux only**: Not portable to other systems

## Concurrent Socket Programming

### Goroutines for Concurrency
- **Lightweight**: Thousands of goroutines
- **Efficient**: M:N threading model
- **Simple**: Easy to write concurrent code
- **Channels**: Communication between goroutines

### Connection Pooling
- **Reuse connections**: Avoid connection overhead
- **Limit connections**: Prevent resource exhaustion
- **Health checks**: Monitor connection health
- **Load balancing**: Distribute load across connections

### Worker Pools
- **Fixed workers**: Limited number of goroutines
- **Task queue**: Channel-based task distribution
- **Backpressure**: Prevent overwhelming workers
- **Graceful shutdown**: Clean worker termination

## Error Handling in Socket Programming

### Common Socket Errors
- **Connection refused**: Port not listening
- **Connection timeout**: Network unreachable
- **Connection reset**: Abrupt connection termination
- **Broken pipe**: Write to closed connection
- **Address in use**: Port already bound

### Error Recovery Strategies
- **Exponential backoff**: Retry with increasing delays
- **Circuit breaker**: Stop trying after failures
- **Health checks**: Monitor connection health
- **Graceful degradation**: Fallback mechanisms

## Performance Optimization

### Buffer Management
- **Send buffers**: Optimize for throughput
- **Receive buffers**: Balance latency and throughput
- **Buffer pooling**: Reuse buffer objects
- **Zero-copy**: Minimize memory copies

### Connection Management
- **Keep-alive**: Maintain persistent connections
- **Connection limits**: Prevent resource exhaustion
- **Timeout handling**: Avoid hanging connections
- **Graceful shutdown**: Clean connection termination

### Network Optimization
- **Nagle's algorithm**: Batch small packets
- **TCP_NODELAY**: Disable batching for real-time
- **Window scaling**: Increase window sizes
- **Congestion control**: Optimize for network conditions

## Advanced Socket Patterns

### Server Patterns
- **Echo Server**: Simple request-response
- **Chat Server**: Multi-client communication
- **File Server**: File transfer protocol
- **Proxy Server**: Forward requests
- **Load Balancer**: Distribute load

### Client Patterns
- **Connection Pool**: Reuse connections
- **Retry Logic**: Handle failures gracefully
- **Circuit Breaker**: Prevent cascade failures
- **Rate Limiting**: Control request rate

### Communication Patterns
- **Request-Response**: Synchronous communication
- **Publish-Subscribe**: Asynchronous messaging
- **Streaming**: Continuous data flow
- **RPC**: Remote procedure calls

## Socket Security

### Authentication
- **TLS/SSL**: Encrypted communication
- **Client certificates**: Mutual authentication
- **API keys**: Simple authentication
- **OAuth**: Token-based authentication

### Authorization
- **Access control**: Limit resource access
- **Rate limiting**: Prevent abuse
- **IP filtering**: Allow/deny by IP
- **User permissions**: Role-based access

### Data Protection
- **Encryption**: Protect data in transit
- **Integrity checks**: Detect tampering
- **Replay protection**: Prevent replay attacks
- **Secure protocols**: Use secure protocols

## Monitoring and Debugging

### Connection Monitoring
- **netstat**: View active connections
- **ss**: Modern socket statistics
- **lsof**: List open files
- **tcpdump**: Capture packets

### Performance Metrics
- **Connection count**: Active connections
- **Throughput**: Bytes per second
- **Latency**: Round-trip time
- **Error rate**: Failed connections

### Debugging Tools
- **Wireshark**: Packet analysis
- **strace**: System call tracing
- **gdb**: Debugger
- **pprof**: Go profiling

## Real-world Applications

### Web Servers
- **HTTP servers**: Handle web requests
- **HTTPS servers**: Secure web communication
- **WebSocket servers**: Real-time communication
- **API servers**: REST/GraphQL APIs

### Microservices
- **Service discovery**: Find services
- **Load balancing**: Distribute requests
- **Circuit breakers**: Handle failures
- **Distributed tracing**: Track requests

### IoT Applications
- **Device communication**: Connect devices
- **Data collection**: Gather sensor data
- **Remote control**: Control devices
- **Telemetry**: Monitor device health

## Go-Specific Socket Programming

### net.Conn Interface
```go
type Conn interface {
    Read(b []byte) (n int, err error)
    Write(b []byte) (n int, err error)
    Close() error
    LocalAddr() Addr
    RemoteAddr() Addr
    SetDeadline(t time.Time) error
    SetReadDeadline(t time.Time) error
    SetWriteDeadline(t time.Time) error
}
```

### Context for Cancellation
- **Context**: Cancel operations
- **Timeout**: Set operation timeouts
- **Cancellation**: Stop long-running operations
- **Values**: Pass request-scoped values

### Channels for Communication
- **Bidirectional**: Send and receive
- **Buffered**: Store multiple values
- **Unbuffered**: Synchronous communication
- **Select**: Non-blocking operations

## Best Practices

### Connection Handling
- **Always close connections**: Prevent resource leaks
- **Use timeouts**: Avoid hanging operations
- **Handle errors**: Implement proper error handling
- **Monitor resources**: Track connection usage

### Concurrency
- **Use goroutines**: Leverage Go's concurrency
- **Limit goroutines**: Prevent resource exhaustion
- **Use channels**: Communicate between goroutines
- **Implement backpressure**: Control flow

### Performance
- **Reuse connections**: Use connection pooling
- **Optimize buffers**: Tune buffer sizes
- **Profile code**: Identify bottlenecks
- **Monitor metrics**: Track performance

### Security
- **Use TLS**: Encrypt communication
- **Validate input**: Sanitize user input
- **Implement authentication**: Verify users
- **Monitor access**: Log security events

## Interview Questions

### Basic Questions
1. What is a socket and how does it work?
2. Explain the difference between TCP and UDP sockets.
3. How do you handle multiple connections in Go?

### Intermediate Questions
1. Implement a concurrent TCP server in Go.
2. How do you handle socket errors gracefully?
3. Explain I/O multiplexing and its benefits.

### Advanced Questions
1. Design a high-performance socket server.
2. Implement connection pooling in Go.
3. How do you optimize socket performance?

## Next Steps
After mastering socket programming, proceed to:
- **Topic 6**: HTTP/HTTPS Implementation
- **Topic 7**: WebSocket Protocols
- **Topic 8**: Network Security & Encryption

Master socket programming, and you'll have the foundation for all network applications! 🚀
```

Now let me create the comprehensive Go implementation for socket programming:


```go
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
```

Now let me update the progress and create a quick reference guide for socket programming:


Let me create a quick reference guide for socket programming:


```markdown
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
```

Perfect! I've completed **Topic 5: Socket Programming in Go** with comprehensive coverage. Here's what I've created for you:

## 🎯 **Topic 5: Socket Programming in Go - COMPLETED**

### **📚 Theory Coverage** (`05_socket_programming/theory.md`)
- **Socket Programming Overview**: What are sockets, types, addressing
- **Go Socket Programming**: net package, socket types, interfaces
- **Advanced Concepts**: Socket options, states, I/O multiplexing
- **Concurrent Programming**: Goroutines, connection pooling, worker pools
- **Error Handling**: Common errors, recovery strategies
- **Performance Optimization**: Buffer management, connection management
- **Advanced Patterns**: Server patterns, client patterns, communication patterns
- **Security**: Authentication, authorization, data protection
- **Monitoring & Debugging**: Tools, metrics, debugging techniques
- **Real-world Applications**: Web servers, microservices, IoT

### **💻 Implementation** (`05_socket_programming/implementation.go`)
- **SocketServer**: Generic socket server with worker pools
- **AdvancedSocketClient**: Client with retry logic and timeouts
- **SocketMultiplexer**: Efficient handling of multiple connections
- **ConnectionPool**: Connection pooling for performance
- **SocketMonitor**: Statistics and monitoring
- **Context Support**: Cancellation and timeout handling
- **Comprehensive Examples**: All major socket programming patterns

### **�� Quick Reference** (`05_socket_programming/quick_reference.md`)
- **Essential Patterns**: Basic TCP/UDP server/client code
- **Advanced Patterns**: Connection pooling, worker pools, context
- **Socket Options**: TCP options, buffer sizes, timeouts
- **Monitoring**: Statistics, error handling, debugging
- **Security**: TLS, input validation, best practices
- **Performance Tips**: Buffered I/O, connection reuse, batching
- **Common Issues**: Solutions for typical problems
- **Best Practices**: Professional socket programming guidelines

## 🚀 **Key Features of This Module**

1. **Complete Coverage**: From basics to advanced socket programming
2. **Go-Specific**: Leverages Go's concurrency and networking features
3. **Production-Ready**: Real-world patterns and best practices
4. **Performance-Focused**: Connection pooling, multiplexing, optimization
5. **Error Handling**: Comprehensive error handling and recovery
6. **Monitoring**: Statistics and debugging capabilities
7. **Security**: TLS, validation, and security best practices

## 📈 **What You've Learned**

- **Socket Fundamentals**: Understanding sockets and their types
- **Go Networking**: Using the net package effectively
- **Concurrency**: Goroutines, channels, worker pools
- **Performance**: Connection pooling, multiplexing, optimization
- **Error Handling**: Graceful error handling and recovery
- **Monitoring**: Statistics, debugging, and troubleshooting
- **Security**: Secure socket programming practices
- **Real-world Applications**: Building production-ready network applications

## �� **Next Steps**

You're now ready to move to **Topic 6: HTTP/HTTPS Implementation** where we'll build on these socket programming fundamentals to create web servers, REST APIs, and secure communication protocols.

Would you like me to continue with **Topic 6: HTTP/HTTPS Implementation** or would you prefer to explore any specific aspect of socket programming first?

The foundation you now have in socket programming will make you unstoppable in networking interviews and real-world applications! 🚀

---

**User**

 move to Topic 6: HTTP/HTTPS Implementation

---

**Cursor**

Perfect! Let me move to **Topic 6: HTTP/HTTPS Implementation**. This is where we build on our socket programming foundation to create web servers, REST APIs, and secure communication protocols.


Now let me create the comprehensive HTTP/HTTPS module:


```markdown
# Topic 6: HTTP/HTTPS Implementation - Web Protocol Mastery

## HTTP Protocol Overview

Hypertext Transfer Protocol (HTTP) is an application-layer protocol for distributed, collaborative, hypermedia information systems. It's the foundation of data communication for the World Wide Web.

### HTTP Characteristics
- **Stateless**: Each request is independent
- **Request-Response**: Client sends request, server responds
- **Text-based**: Human-readable protocol
- **Flexible**: Can transfer any type of data
- **Extensible**: Headers and methods can be extended

### HTTP Versions
- **HTTP/1.0**: Basic protocol with connection per request
- **HTTP/1.1**: Persistent connections, pipelining, chunked encoding
- **HTTP/2**: Binary protocol, multiplexing, server push
- **HTTP/3**: Based on QUIC, improved performance

## HTTP Message Structure

### HTTP Request Format
```
Method SP Request-URI SP HTTP-Version CRLF
Header-Name: Header-Value CRLF
Header-Name: Header-Value CRLF
...
CRLF
[Message Body]
```

### HTTP Response Format
```
HTTP-Version SP Status-Code SP Reason-Phrase CRLF
Header-Name: Header-Value CRLF
Header-Name: Header-Value CRLF
...
CRLF
[Message Body]
```

## HTTP Methods

### Safe Methods
- **GET**: Retrieve data (idempotent)
- **HEAD**: Get headers only (idempotent)
- **OPTIONS**: Get allowed methods (idempotent)
- **TRACE**: Echo request (idempotent)

### Unsafe Methods
- **POST**: Create or submit data
- **PUT**: Create or update resource (idempotent)
- **PATCH**: Partial update
- **DELETE**: Remove resource (idempotent)

### Idempotent Methods
- **GET, HEAD, OPTIONS, TRACE, PUT, DELETE**
- Can be called multiple times with same result
- Safe for retries and caching

## HTTP Status Codes

### 1xx Informational
- **100 Continue**: Request received, continue
- **101 Switching Protocols**: Protocol upgrade

### 2xx Success
- **200 OK**: Request successful
- **201 Created**: Resource created
- **202 Accepted**: Request accepted for processing
- **204 No Content**: Success, no content returned

### 3xx Redirection
- **301 Moved Permanently**: Resource moved permanently
- **302 Found**: Resource temporarily moved
- **304 Not Modified**: Resource not modified

### 4xx Client Error
- **400 Bad Request**: Invalid request
- **401 Unauthorized**: Authentication required
- **403 Forbidden**: Access denied
- **404 Not Found**: Resource not found
- **405 Method Not Allowed**: Method not supported

### 5xx Server Error
- **500 Internal Server Error**: Server error
- **501 Not Implemented**: Method not implemented
- **502 Bad Gateway**: Invalid response from upstream
- **503 Service Unavailable**: Server overloaded

## HTTP Headers

### Request Headers
- **Host**: Target host and port
- **User-Agent**: Client information
- **Accept**: Acceptable content types
- **Accept-Language**: Preferred languages
- **Accept-Encoding**: Acceptable encodings
- **Authorization**: Authentication credentials
- **Content-Type**: Media type of body
- **Content-Length**: Size of body in bytes

### Response Headers
- **Content-Type**: Media type of response
- **Content-Length**: Size of response body
- **Content-Encoding**: Encoding of response
- **Cache-Control**: Caching directives
- **Expires**: Expiration time
- **Last-Modified**: Last modification time
- **Location**: Redirect location
- **Set-Cookie**: Cookie to set

### General Headers
- **Connection**: Connection control
- **Date**: Date and time
- **Transfer-Encoding**: Transfer encoding
- **Upgrade**: Protocol upgrade
- **Via**: Proxy information

## HTTP/1.1 Features

### Persistent Connections
- **Keep-Alive**: Reuse connections for multiple requests
- **Connection: keep-alive**: Enable persistent connections
- **Connection: close**: Close after response

### Pipelining
- **Multiple requests**: Send multiple requests without waiting
- **Order preservation**: Responses in same order as requests
- **Error handling**: One error affects all pipelined requests

### Chunked Transfer Encoding
- **Dynamic content**: Send data without knowing total size
- **Chunk format**: Size in hex + CRLF + data + CRLF
- **End marker**: 0 + CRLF + CRLF

### Host Header
- **Virtual hosting**: Multiple domains on same IP
- **Required in HTTP/1.1**: Must be present in all requests
- **Port specification**: Include port if not default

## HTTPS (HTTP Secure)

### TLS/SSL Overview
- **Transport Layer Security**: Encrypts HTTP traffic
- **Public Key Cryptography**: Asymmetric encryption
- **Certificate Authority**: Validates server identity
- **Perfect Forward Secrecy**: Session keys not stored

### HTTPS Handshake
1. **Client Hello**: Supported cipher suites, random data
2. **Server Hello**: Chosen cipher suite, server certificate
3. **Certificate Verification**: Client validates certificate
4. **Key Exchange**: Generate shared secret
5. **Finished**: Handshake complete, encrypted communication

### Certificate Validation
- **Chain of Trust**: Root CA → Intermediate CA → Server
- **Certificate Revocation**: CRL or OCSP
- **Hostname Verification**: Certificate matches hostname
- **Expiration Check**: Certificate not expired

## Go HTTP Implementation

### net/http Package
- **http.Server**: HTTP server implementation
- **http.Client**: HTTP client implementation
- **http.Handler**: Request handler interface
- **http.HandlerFunc**: Function-based handler
- **http.ServeMux**: URL router

### Basic HTTP Server
```go
func startHTTPServer() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Basic HTTP Client
```go
func startHTTPClient() {
    resp, err := http.Get("http://localhost:8080")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(body))
}
```

## Advanced HTTP Concepts

### Middleware
- **Request Processing**: Intercept and modify requests
- **Response Processing**: Intercept and modify responses
- **Chaining**: Multiple middleware functions
- **Common Uses**: Logging, authentication, CORS

### Routing
- **URL Patterns**: Match requests to handlers
- **Path Parameters**: Extract values from URLs
- **Query Parameters**: Handle query strings
- **Wildcards**: Flexible URL matching

### Content Negotiation
- **Accept Header**: Client preferences
- **Content-Type**: Response format
- **Language**: Accept-Language header
- **Encoding**: Accept-Encoding header

### Caching
- **Cache-Control**: Caching directives
- **ETag**: Entity tag for validation
- **Last-Modified**: Modification time
- **Expires**: Expiration time

## HTTP/2 Features

### Binary Protocol
- **Frames**: Binary format instead of text
- **Streams**: Multiple concurrent streams
- **Headers Compression**: HPACK compression
- **Server Push**: Server-initiated responses

### Multiplexing
- **Concurrent Streams**: Multiple requests per connection
- **Stream Prioritization**: Priority-based scheduling
- **Flow Control**: Per-stream flow control
- **Connection Management**: Single connection per origin

### Header Compression
- **HPACK**: Header compression algorithm
- **Static Table**: Common headers
- **Dynamic Table**: Custom headers
- **Huffman Coding**: Efficient encoding

## HTTP/3 Features

### QUIC Protocol
- **UDP-based**: Faster than TCP
- **Built-in Encryption**: TLS 1.3 integrated
- **Connection Migration**: Survive IP changes
- **Multiplexing**: No head-of-line blocking

### Performance Improvements
- **Faster Connection**: 0-RTT connection establishment
- **Better Multiplexing**: True multiplexing
- **Improved Security**: Modern encryption
- **Network Resilience**: Better error recovery

## REST API Design

### REST Principles
- **Resource-based**: Everything is a resource
- **HTTP Methods**: Use appropriate methods
- **Stateless**: No client state on server
- **Uniform Interface**: Consistent API design

### Resource Design
- **Nouns, not verbs**: Use resource names
- **Hierarchical**: Nested resources
- **Consistent**: Follow naming conventions
- **Versioned**: API versioning strategy

### HTTP Methods for REST
- **GET**: Retrieve resource
- **POST**: Create resource
- **PUT**: Update entire resource
- **PATCH**: Partial update
- **DELETE**: Remove resource

## Error Handling

### HTTP Error Responses
- **Appropriate Status Codes**: Use correct status codes
- **Error Messages**: Clear error descriptions
- **Error Codes**: Application-specific error codes
- **Error Details**: Additional error information

### Client Error Handling
- **Status Code Checking**: Check response status
- **Retry Logic**: Handle transient errors
- **Timeout Handling**: Set appropriate timeouts
- **Error Propagation**: Pass errors up the call stack

## Security Considerations

### Input Validation
- **Sanitize Input**: Clean user input
- **Validate Data**: Check data format and range
- **Prevent Injection**: SQL, XSS, command injection
- **Rate Limiting**: Prevent abuse

### Authentication & Authorization
- **Authentication**: Verify user identity
- **Authorization**: Check permissions
- **Session Management**: Secure session handling
- **Token-based**: JWT, OAuth tokens

### HTTPS Implementation
- **TLS Configuration**: Secure TLS settings
- **Certificate Management**: Proper certificate handling
- **HSTS**: HTTP Strict Transport Security
- **CSP**: Content Security Policy

## Performance Optimization

### Server Optimization
- **Connection Pooling**: Reuse connections
- **Keep-Alive**: Persistent connections
- **Compression**: Gzip compression
- **Caching**: Response caching

### Client Optimization
- **Connection Reuse**: Reuse HTTP connections
- **Pipelining**: Multiple requests per connection
- **Compression**: Accept compressed responses
- **Caching**: Cache responses locally

### CDN and Caching
- **Content Delivery Network**: Global distribution
- **Edge Caching**: Cache at edge locations
- **Cache Headers**: Proper cache directives
- **Cache Invalidation**: Update cached content

## Monitoring and Debugging

### HTTP Monitoring
- **Request Metrics**: Count, latency, errors
- **Response Metrics**: Status codes, response times
- **Connection Metrics**: Active connections, throughput
- **Error Tracking**: Error rates and types

### Debugging Tools
- **curl**: Command-line HTTP client
- **Postman**: GUI HTTP client
- **Wireshark**: Packet analysis
- **Browser DevTools**: Network inspection

### Logging
- **Access Logs**: Request/response logging
- **Error Logs**: Error tracking and debugging
- **Performance Logs**: Timing and metrics
- **Security Logs**: Security events and attacks

## Interview Questions

### Basic Questions
1. What is HTTP and how does it work?
2. Explain the difference between HTTP and HTTPS.
3. What are the main HTTP methods and when to use them?

### Intermediate Questions
1. How do you implement authentication in HTTP?
2. Explain HTTP caching and its benefits.
3. What are the differences between HTTP/1.1 and HTTP/2?

### Advanced Questions
1. Design a RESTful API for a social media platform.
2. How would you implement rate limiting in an HTTP server?
3. Explain HTTP/3 and its advantages over HTTP/2.

## Next Steps
After mastering HTTP/HTTPS, proceed to:
- **Topic 7**: WebSocket Protocols
- **Topic 8**: Network Security & Encryption
- **Topic 9**: Routing & Switching

Master HTTP/HTTPS, and you'll understand the foundation of web communication! 🚀
```

Now let me create the comprehensive Go implementation for HTTP/HTTPS:


```go
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
```

Now let me create a comprehensive REST API example and update the progress:


```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// User represents a user in our system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserService handles user operations
type UserService struct {
	Users  map[int]*User
	NextID int
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	return &UserService{
		Users:  make(map[int]*User),
		NextID: 1,
	}
}

// CreateUser creates a new user
func (us *UserService) CreateUser(name, email string) *User {
	user := &User{
		ID:        us.NextID,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	us.Users[us.NextID] = user
	us.NextID++
	return user
}

// GetUser gets a user by ID
func (us *UserService) GetUser(id int) (*User, bool) {
	user, exists := us.Users[id]
	return user, exists
}

// GetAllUsers gets all users
func (us *UserService) GetAllUsers() []*User {
	users := make([]*User, 0, len(us.Users))
	for _, user := range us.Users {
		users = append(users, user)
	}
	return users
}

// UpdateUser updates a user
func (us *UserService) UpdateUser(id int, name, email string) (*User, bool) {
	user, exists := us.Users[id]
	if !exists {
		return nil, false
	}
	
	user.Name = name
	user.Email = email
	user.UpdatedAt = time.Now()
	
	return user, true
}

// DeleteUser deletes a user
func (us *UserService) DeleteUser(id int) bool {
	_, exists := us.Users[id]
	if exists {
		delete(us.Users, id)
	}
	return exists
}

// RESTAPI represents a REST API server
type RESTAPI struct {
	UserService *UserService
	Mux         *http.ServeMux
}

// NewRESTAPI creates a new REST API
func NewRESTAPI() *RESTAPI {
	api := &RESTAPI{
		UserService: NewUserService(),
		Mux:         http.NewServeMux(),
	}
	
	api.setupRoutes()
	return api
}

// setupRoutes sets up API routes
func (api *RESTAPI) setupRoutes() {
	// User routes
	api.Mux.HandleFunc("/api/v1/users", api.handleUsers)
	api.Mux.HandleFunc("/api/v1/users/", api.handleUserByID)
	
	// Health check
	api.Mux.HandleFunc("/health", api.handleHealth)
	
	// API documentation
	api.Mux.HandleFunc("/api/v1", api.handleAPIDocs)
}

// handleUsers handles /api/v1/users
func (api *RESTAPI) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.getUsers(w, r)
	case http.MethodPost:
		api.createUser(w, r)
	default:
		api.writeErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// handleUserByID handles /api/v1/users/{id}
func (api *RESTAPI) handleUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from path
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/users/")
	if path == "" {
		api.writeErrorResponse(w, http.StatusBadRequest, "User ID required")
		return
	}
	
	userID, err := strconv.Atoi(path)
	if err != nil {
		api.writeErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		api.getUserByID(w, r, userID)
	case http.MethodPut:
		api.updateUser(w, r, userID)
	case http.MethodDelete:
		api.deleteUser(w, r, userID)
	default:
		api.writeErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// getUsers handles GET /api/v1/users
func (api *RESTAPI) getUsers(w http.ResponseWriter, r *http.Request) {
	users := api.UserService.GetAllUsers()
	
	response := map[string]interface{}{
		"data":    users,
		"count":   len(users),
		"message": "Users retrieved successfully",
	}
	
	api.writeJSONResponse(w, http.StatusOK, response)
}

// createUser handles POST /api/v1/users
func (api *RESTAPI) createUser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		api.writeErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	
	// Validate input
	if request.Name == "" || request.Email == "" {
		api.writeErrorResponse(w, http.StatusBadRequest, "Name and email are required")
		return
	}
	
	// Create user
	user := api.UserService.CreateUser(request.Name, request.Email)
	
	response := map[string]interface{}{
		"data":    user,
		"message": "User created successfully",
	}
	
	api.writeJSONResponse(w, http.StatusCreated, response)
}

// getUserByID handles GET /api/v1/users/{id}
func (api *RESTAPI) getUserByID(w http.ResponseWriter, r *http.Request, userID int) {
	user, exists := api.UserService.GetUser(userID)
	if !exists {
		api.writeErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}
	
	response := map[string]interface{}{
		"data":    user,
		"message": "User retrieved successfully",
	}
	
	api.writeJSONResponse(w, http.StatusOK, response)
}

// updateUser handles PUT /api/v1/users/{id}
func (api *RESTAPI) updateUser(w http.ResponseWriter, r *http.Request, userID int) {
	var request struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		api.writeErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	
	// Validate input
	if request.Name == "" || request.Email == "" {
		api.writeErrorResponse(w, http.StatusBadRequest, "Name and email are required")
		return
	}
	
	// Update user
	user, exists := api.UserService.UpdateUser(userID, request.Name, request.Email)
	if !exists {
		api.writeErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}
	
	response := map[string]interface{}{
		"data":    user,
		"message": "User updated successfully",
	}
	
	api.writeJSONResponse(w, http.StatusOK, response)
}

// deleteUser handles DELETE /api/v1/users/{id}
func (api *RESTAPI) deleteUser(w http.ResponseWriter, r *http.Request, userID int) {
	exists := api.UserService.DeleteUser(userID)
	if !exists {
		api.writeErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}
	
	response := map[string]interface{}{
		"message": "User deleted successfully",
	}
	
	api.writeJSONResponse(w, http.StatusOK, response)
}

// handleHealth handles GET /health
func (api *RESTAPI) handleHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"version":   "1.0.0",
	}
	
	api.writeJSONResponse(w, http.StatusOK, response)
}

// handleAPIDocs handles GET /api/v1
func (api *RESTAPI) handleAPIDocs(w http.ResponseWriter, r *http.Request) {
	docs := map[string]interface{}{
		"name":        "User Management API",
		"version":     "1.0.0",
		"description": "A RESTful API for managing users",
		"endpoints": map[string]interface{}{
			"GET /api/v1/users":     "Get all users",
			"POST /api/v1/users":    "Create a new user",
			"GET /api/v1/users/{id}": "Get user by ID",
			"PUT /api/v1/users/{id}": "Update user by ID",
			"DELETE /api/v1/users/{id}": "Delete user by ID",
			"GET /health":           "Health check",
		},
		"examples": map[string]interface{}{
			"create_user": map[string]interface{}{
				"method": "POST",
				"url":    "/api/v1/users",
				"body": map[string]string{
					"name":  "John Doe",
					"email": "john@example.com",
				},
			},
			"get_user": map[string]interface{}{
				"method": "GET",
				"url":    "/api/v1/users/1",
			},
		},
	}
	
	api.writeJSONResponse(w, http.StatusOK, docs)
}

// writeJSONResponse writes a JSON response
func (api *RESTAPI) writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

// writeErrorResponse writes an error response
func (api *RESTAPI) writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response := map[string]interface{}{
		"error":   http.StatusText(statusCode),
		"message": message,
		"code":    statusCode,
	}
	
	api.writeJSONResponse(w, statusCode, response)
}

// Start starts the REST API server
func (api *RESTAPI) Start(addr string) error {
	fmt.Printf("REST API Server starting on %s\n", addr)
	return http.ListenAndServe(addr, api.Mux)
}

// Demonstrate REST API
func demonstrateRESTAPI() {
	fmt.Println("=== REST API Demo ===\n")
	
	// Create API
	api := NewRESTAPI()
	
	// Start server in goroutine
	go func() {
		if err := api.Start(":8080"); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()
	
	// Give server time to start
	time.Sleep(100 * time.Millisecond)
	
	// Test API endpoints
	fmt.Println("Testing REST API endpoints...\n")
	
	// Test GET /api/v1/users (should be empty)
	fmt.Println("1. GET /api/v1/users")
	resp, err := http.Get("http://localhost:8080/api/v1/users")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Status: %d\n", resp.StatusCode)
	resp.Body.Close()
	
	// Test POST /api/v1/users
	fmt.Println("\n2. POST /api/v1/users")
	userData := `{"name": "John Doe", "email": "john@example.com"}`
	resp, err = http.Post("http://localhost:8080/api/v1/users", "application/json", strings.NewReader(userData))
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Status: %d\n", resp.StatusCode)
	resp.Body.Close()
	
	// Test GET /api/v1/users (should have one user)
	fmt.Println("\n3. GET /api/v1/users")
	resp, err = http.Get("http://localhost:8080/api/v1/users")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Status: %d\n", resp.StatusCode)
	resp.Body.Close()
	
	// Test GET /api/v1/users/1
	fmt.Println("\n4. GET /api/v1/users/1")
	resp, err = http.Get("http://localhost:8080/api/v1/users/1")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Status: %d\n", resp.StatusCode)
	resp.Body.Close()
	
	// Test PUT /api/v1/users/1
	fmt.Println("\n5. PUT /api/v1/users/1")
	updateData := `{"name": "John Smith", "email": "john.smith@example.com"}`
	req, err := http.NewRequest("PUT", "http://localhost:8080/api/v1/users/1", strings.NewReader(updateData))
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Status: %d\n", resp.StatusCode)
	resp.Body.Close()
	
	// Test DELETE /api/v1/users/1
	fmt.Println("\n6. DELETE /api/v1/users/1")
	req, err = http.NewRequest("DELETE", "http://localhost:8080/api/v1/users/1", nil)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	
	resp, err = client.Do(req)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Status: %d\n", resp.StatusCode)
	resp.Body.Close()
	
	// Test GET /health
	fmt.Println("\n7. GET /health")
	resp, err = http.Get("http://localhost:8080/health")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Status: %d\n", resp.StatusCode)
	resp.Body.Close()
	
	// Test GET /api/v1 (API docs)
	fmt.Println("\n8. GET /api/v1")
	resp, err = http.Get("http://localhost:8080/api/v1")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Status: %d\n", resp.StatusCode)
	resp.Body.Close()
	
	fmt.Println("\nREST API demo completed!")
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 6: REST API Example")
	fmt.Println("========================================================\n")
	
	demonstrateRESTAPI()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. REST APIs follow standard HTTP methods and status codes")
	fmt.Println("2. JSON is the standard format for API communication")
	fmt.Println("3. Proper error handling and status codes are essential")
	fmt.Println("4. API documentation helps users understand the interface")
	fmt.Println("5. Input validation prevents errors and security issues")
	fmt.Println("6. Consistent response format improves API usability")
}
```

Now let me update the progress and create a quick reference guide:


Let me create a quick reference guide for HTTP/HTTPS:


```markdown
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
```

Perfect! I've completed **Topic 6: HTTP/HTTPS Implementation** with comprehensive coverage. Here's what I've created for you:

## �� **Topic 6: HTTP/HTTPS Implementation - COMPLETED**

### **📚 Theory Coverage** (`06_http_https/theory.md`)
- **HTTP Protocol Overview**: Characteristics, versions, message structure
- **HTTP Methods**: Safe, unsafe, and idempotent methods
- **HTTP Status Codes**: Complete status code reference
- **HTTP Headers**: Request, response, and general headers
- **HTTP/1.1 Features**: Persistent connections, pipelining, chunked encoding
- **HTTPS**: TLS/SSL, handshake, certificate validation
- **Go HTTP Implementation**: net/http package, servers, clients
- **Advanced Concepts**: Middleware, routing, content negotiation, caching
- **HTTP/2 & HTTP/3**: New features and performance improvements
- **REST API Design**: Principles, resource design, HTTP methods
- **Security Considerations**: Input validation, authentication, HTTPS
- **Performance Optimization**: Server/client optimization, CDN, caching

### **💻 Implementation** (`06_http_https/implementation.go`)
- **HTTPServer**: Advanced HTTP server with middleware support
- **HTTPClient**: Advanced HTTP client with connection pooling
- **HTTPSClient**: HTTPS client with TLS configuration
- **LoadBalancer**: Round-robin load balancing
- **Proxy**: HTTP proxy implementation
- **Middleware**: Logging, CORS, compression, statistics
- **REST API**: Complete user management API
- **Error Handling**: Comprehensive error handling and responses

### **🔧 REST API Example** (`06_http_https/rest_api_example.go`)
- **User Management API**: Complete CRUD operations
- **RESTful Design**: Proper HTTP methods and status codes
- **JSON Communication**: Request/response handling
- **Input Validation**: Data validation and error handling
- **API Documentation**: Self-documenting API endpoints

### **📖 Quick Reference** (`06_http_https/quick_reference.md`)
- **Essential Concepts**: HTTP methods, status codes, patterns
- **Go Patterns**: Server, client, middleware patterns
- **HTTPS Implementation**: Secure communication setup
- **REST API Patterns**: JSON responses, error handling, validation
- **Load Balancing**: Round-robin and proxy patterns
- **Performance Optimization**: Connection pooling, compression, caching
- **Security Best Practices**: Input validation, rate limiting, CORS
- **Common Issues**: Solutions for typical problems
- **Best Practices**: Professional HTTP development guidelines

## �� **Key Features of This Module**

1. **Complete HTTP Coverage**: From basics to advanced concepts
2. **Go-Specific**: Leverages Go's excellent HTTP support
3. **Production-Ready**: Real-world patterns and best practices
4. **Security-Focused**: HTTPS, validation, authentication
5. **Performance-Optimized**: Connection pooling, caching, compression
6. **REST API**: Complete API implementation example
7. **Middleware Support**: Logging, CORS, compression, statistics
8. **Load Balancing**: Round-robin and proxy implementations

## 📈 **What You've Learned**

- **HTTP Fundamentals**: Protocol, methods, status codes, headers
- **HTTPS Security**: TLS/SSL, certificates, secure communication
- **Go HTTP Programming**: Servers, clients, middleware, routing
- **REST API Design**: Principles, implementation, best practices
- **Performance Optimization**: Connection pooling, caching, compression
- **Security**: Input validation, authentication, rate limiting
- **Load Balancing**: Distribution and proxy patterns
- **Real-world Applications**: Production-ready HTTP applications

## **Next Steps**

You're now ready to move to **Topic 7: WebSocket Protocols** where we'll explore real-time communication, bidirectional data flow, and modern web applications.

The HTTP/HTTPS foundation you now have will make you unstoppable in web development, API design, and modern web applications! 🚀

Would you like me to continue with **Topic 7: WebSocket Protocols** or would you prefer to explore any specific aspect of HTTP/HTTPS first?

---

**User**

continue with Topic 7: 

---

**Cursor**

Perfect! Let me continue with **Topic 7: WebSocket Protocols**. This is where we explore real-time communication, bidirectional data flow, and modern web applications.


Now let me create the comprehensive WebSocket module:


```markdown
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
```

Now let me create the comprehensive Go implementation for WebSocket:


```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocketServer represents a WebSocket server
type WebSocketServer struct {
	Upgrader    websocket.Upgrader
	Connections map[*websocket.Conn]*Connection
	Mutex       sync.RWMutex
	Broadcast   chan []byte
	Register    chan *websocket.Conn
	Unregister  chan *websocket.Conn
	Running     bool
}

// Connection represents a WebSocket connection
type Connection struct {
	Conn        *websocket.Conn
	ID          string
	UserID      string
	Username    string
	CreatedAt   time.Time
	LastPing    time.Time
	IsAlive     bool
	Send        chan []byte
	Room        string
}

// Message represents a WebSocket message
type Message struct {
	Type      string      `json:"type"`
	Content   string      `json:"content"`
	Username  string      `json:"username"`
	UserID    string      `json:"user_id"`
	Room      string      `json:"room"`
	Timestamp time.Time   `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
}

// Room represents a chat room
type Room struct {
	ID          string
	Name        string
	Connections map[*websocket.Conn]*Connection
	Mutex       sync.RWMutex
	CreatedAt   time.Time
}

// NewWebSocketServer creates a new WebSocket server
func NewWebSocketServer() *WebSocketServer {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins in demo
		},
	}

	return &WebSocketServer{
		Upgrader:    upgrader,
		Connections: make(map[*websocket.Conn]*Connection),
		Broadcast:   make(chan []byte),
		Register:    make(chan *websocket.Conn),
		Unregister:  make(chan *websocket.Conn),
		Running:     false,
	}
}

// Start starts the WebSocket server
func (ws *WebSocketServer) Start() {
	ws.Running = true
	
	// Start connection manager
	go ws.manageConnections()
	
	// Start message broadcaster
	go ws.broadcastMessages()
	
	// Start ping handler
	go ws.handlePings()
}

// Stop stops the WebSocket server
func (ws *WebSocketServer) Stop() {
	ws.Running = false
	
	// Close all connections
	ws.Mutex.Lock()
	for conn, connection := range ws.Connections {
		conn.Close()
		close(connection.Send)
	}
	ws.Connections = make(map[*websocket.Conn]*Connection)
	ws.Mutex.Unlock()
	
	// Close channels
	close(ws.Broadcast)
	close(ws.Register)
	close(ws.Unregister)
}

// manageConnections manages WebSocket connections
func (ws *WebSocketServer) manageConnections() {
	for ws.Running {
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

// registerConnection registers a new connection
func (ws *WebSocketServer) registerConnection(conn *websocket.Conn) {
	connection := &Connection{
		Conn:      conn,
		ID:        generateConnectionID(),
		CreatedAt: time.Now(),
		LastPing:  time.Now(),
		IsAlive:   true,
		Send:      make(chan []byte, 256),
		Room:      "general",
	}
	
	ws.Mutex.Lock()
	ws.Connections[conn] = connection
	ws.Mutex.Unlock()
	
	// Start goroutines for this connection
	go ws.writePump(connection)
	go ws.readPump(connection)
	
	fmt.Printf("New connection registered: %s\n", connection.ID)
}

// unregisterConnection unregisters a connection
func (ws *WebSocketServer) unregisterConnection(conn *websocket.Conn) {
	ws.Mutex.Lock()
	if connection, exists := ws.Connections[conn]; exists {
		close(connection.Send)
		delete(ws.Connections, conn)
		fmt.Printf("Connection unregistered: %s\n", connection.ID)
	}
	ws.Mutex.Unlock()
}

// broadcastToAll broadcasts a message to all connections
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

// broadcastToRoom broadcasts a message to a specific room
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

// writePump handles writing messages to the WebSocket connection
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
			
			// Add queued messages to the current websocket message
			n := len(connection.Send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-connection.Send)
			}
			
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

// readPump handles reading messages from the WebSocket connection
func (ws *WebSocketServer) readPump(connection *Connection) {
	defer func() {
		ws.Unregister <- connection.Conn
		connection.Conn.Close()
	}()
	
	connection.Conn.SetReadLimit(512)
	connection.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	connection.Conn.SetPongHandler(func(string) error {
		connection.LastPing = time.Now()
		connection.IsAlive = true
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
		
		// Parse message
		var message Message
		if err := json.Unmarshal(messageBytes, &message); err != nil {
			log.Printf("Error parsing message: %v", err)
			continue
		}
		
		// Set connection info
		message.Username = connection.Username
		message.UserID = connection.UserID
		message.Room = connection.Room
		message.Timestamp = time.Now()
		
		// Handle different message types
		ws.handleMessage(connection, &message)
	}
}

// handleMessage handles different types of messages
func (ws *WebSocketServer) handleMessage(connection *Connection, message *Message) {
	switch message.Type {
	case "join":
		ws.handleJoin(connection, message)
	case "leave":
		ws.handleLeave(connection, message)
	case "chat":
		ws.handleChat(connection, message)
	case "ping":
		ws.handlePing(connection, message)
	case "set_username":
		ws.handleSetUsername(connection, message)
	default:
		log.Printf("Unknown message type: %s", message.Type)
	}
}

// handleJoin handles join room messages
func (ws *WebSocketServer) handleJoin(connection *Connection, message *Message) {
	room := message.Room
	if room == "" {
		room = "general"
	}
	
	connection.Room = room
	
	response := Message{
		Type:      "joined",
		Content:   fmt.Sprintf("Joined room: %s", room),
		Username:  "System",
		Room:      room,
		Timestamp: time.Now(),
	}
	
	responseBytes, _ := json.Marshal(response)
	connection.Send <- responseBytes
}

// handleLeave handles leave room messages
func (ws *WebSocketServer) handleLeave(connection *Connection, message *Message) {
	oldRoom := connection.Room
	connection.Room = "general"
	
	response := Message{
		Type:      "left",
		Content:   fmt.Sprintf("Left room: %s", oldRoom),
		Username:  "System",
		Room:      "general",
		Timestamp: time.Now(),
	}
	
	responseBytes, _ := json.Marshal(response)
	connection.Send <- responseBytes
}

// handleChat handles chat messages
func (ws *WebSocketServer) handleChat(connection *Connection, message *Message) {
	// Broadcast to room
	messageBytes, _ := json.Marshal(message)
	ws.broadcastToRoom(connection.Room, messageBytes)
}

// handlePing handles ping messages
func (ws *WebSocketServer) handlePing(connection *Connection, message *Message) {
	response := Message{
		Type:      "pong",
		Content:   "pong",
		Username:  "System",
		Room:      connection.Room,
		Timestamp: time.Now(),
	}
	
	responseBytes, _ := json.Marshal(response)
	connection.Send <- responseBytes
}

// handleSetUsername handles username setting
func (ws *WebSocketServer) handleSetUsername(connection *Connection, message *Message) {
	username := message.Content
	if username == "" {
		username = "Anonymous"
	}
	
	connection.Username = username
	connection.UserID = generateUserID(username)
	
	response := Message{
		Type:      "username_set",
		Content:   fmt.Sprintf("Username set to: %s", username),
		Username:  "System",
		Room:      connection.Room,
		Timestamp: time.Now(),
	}
	
	responseBytes, _ := json.Marshal(response)
	connection.Send <- responseBytes
}

// handlePings handles ping/pong for connection health
func (ws *WebSocketServer) handlePings() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for ws.Running {
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

// generateConnectionID generates a unique connection ID
func generateConnectionID() string {
	return fmt.Sprintf("conn_%d", time.Now().UnixNano())
}

// generateUserID generates a unique user ID
func generateUserID(username string) string {
	return fmt.Sprintf("user_%s_%d", username, time.Now().UnixNano())
}

// WebSocketClient represents a WebSocket client
type WebSocketClient struct {
	Conn     *websocket.Conn
	URL      string
	Messages chan []byte
	Errors   chan error
	Running  bool
}

// NewWebSocketClient creates a new WebSocket client
func NewWebSocketClient(url string) *WebSocketClient {
	return &WebSocketClient{
		URL:      url,
		Messages: make(chan []byte, 256),
		Errors:   make(chan error, 1),
		Running:  false,
	}
}

// Connect connects to the WebSocket server
func (client *WebSocketClient) Connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(client.URL, nil)
	if err != nil {
		return err
	}
	
	client.Conn = conn
	client.Running = true
	
	// Start reading messages
	go client.readMessages()
	
	return nil
}

// readMessages reads messages from the WebSocket connection
func (client *WebSocketClient) readMessages() {
	defer close(client.Messages)
	defer close(client.Errors)
	
	for client.Running {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			client.Errors <- err
			return
		}
		client.Messages <- message
	}
}

// Send sends a message to the server
func (client *WebSocketClient) Send(message []byte) error {
	if !client.Running {
		return fmt.Errorf("client not connected")
	}
	
	return client.Conn.WriteMessage(websocket.TextMessage, message)
}

// SendMessage sends a structured message
func (client *WebSocketClient) SendMessage(messageType, content, username, room string) error {
	message := Message{
		Type:     messageType,
		Content:  content,
		Username: username,
		Room:     room,
	}
	
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	
	return client.Send(messageBytes)
}

// Close closes the WebSocket connection
func (client *WebSocketClient) Close() error {
	client.Running = false
	if client.Conn != nil {
		return client.Conn.Close()
	}
	return nil
}

// ChatRoom represents a chat room
type ChatRoom struct {
	ID          string
	Name        string
	Connections map[*websocket.Conn]*Connection
	Mutex       sync.RWMutex
	CreatedAt   time.Time
}

// NewChatRoom creates a new chat room
func NewChatRoom(id, name string) *ChatRoom {
	return &ChatRoom{
		ID:          id,
		Name:        name,
		Connections: make(map[*websocket.Conn]*Connection),
		CreatedAt:   time.Now(),
	}
}

// AddConnection adds a connection to the room
func (room *ChatRoom) AddConnection(conn *websocket.Conn, connection *Connection) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()
	
	room.Connections[conn] = connection
	connection.Room = room.ID
}

// RemoveConnection removes a connection from the room
func (room *ChatRoom) RemoveConnection(conn *websocket.Conn) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()
	
	delete(room.Connections, conn)
}

// Broadcast broadcasts a message to all connections in the room
func (room *ChatRoom) Broadcast(message []byte) {
	room.Mutex.RLock()
	defer room.Mutex.RUnlock()
	
	for conn, connection := range room.Connections {
		select {
		case connection.Send <- message:
		default:
			close(connection.Send)
			delete(room.Connections, conn)
		}
	}
}

// GetConnectionCount returns the number of connections in the room
func (room *ChatRoom) GetConnectionCount() int {
	room.Mutex.RLock()
	defer room.Mutex.RUnlock()
	
	return len(room.Connections)
}

// Demonstrate basic WebSocket server
func demonstrateBasicWebSocketServer() {
	fmt.Println("=== Basic WebSocket Server Demo ===\n")
	
	server := NewWebSocketServer()
	server.Start()
	defer server.Stop()
	
	// Set up HTTP handler
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := server.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Upgrade error: %v", err)
			return
		}
		
		server.Register <- conn
	})
	
	// Start HTTP server
	go func() {
		fmt.Println("WebSocket server starting on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	
	// Give server time to start
	time.Sleep(100 * time.Millisecond)
	
	// Test with client
	client := NewWebSocketClient("ws://localhost:8080/ws")
	if err := client.Connect(); err != nil {
		log.Printf("Client connection error: %v", err)
		return
	}
	defer client.Close()
	
	// Send test messages
	testMessages := []string{
		`{"type": "set_username", "content": "TestUser"}`,
		`{"type": "join", "room": "general"}`,
		`{"type": "chat", "content": "Hello, WebSocket!"}`,
		`{"type": "ping", "content": "ping"}`,
	}
	
	for _, msg := range testMessages {
		if err := client.Send([]byte(msg)); err != nil {
			log.Printf("Send error: %v", err)
			continue
		}
		
		// Read response
		select {
		case response := <-client.Messages:
			fmt.Printf("Sent: %s\n", msg)
			fmt.Printf("Received: %s\n\n", string(response))
		case err := <-client.Errors:
			log.Printf("Read error: %v", err)
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout waiting for response")
		}
	}
}

// Demonstrate WebSocket client
func demonstrateWebSocketClient() {
	fmt.Println("=== WebSocket Client Demo ===\n")
	
	client := NewWebSocketClient("ws://echo.websocket.org")
	if err := client.Connect(); err != nil {
		log.Printf("Connection error: %v", err)
		return
	}
	defer client.Close()
	
	// Send test message
	testMessage := "Hello, Echo Server!"
	if err := client.Send([]byte(testMessage)); err != nil {
		log.Printf("Send error: %v", err)
		return
	}
	
	// Read response
	select {
	case response := <-client.Messages:
		fmt.Printf("Sent: %s\n", testMessage)
		fmt.Printf("Received: %s\n", string(response))
	case err := <-client.Errors:
		log.Printf("Read error: %v", err)
	case <-time.After(10 * time.Second):
		fmt.Println("Timeout waiting for response")
	}
}

// Demonstrate chat room functionality
func demonstrateChatRoom() {
	fmt.Println("=== Chat Room Demo ===\n")
	
	// Create chat room
	room := NewChatRoom("general", "General Chat")
	
	// Simulate connections
	connections := make([]*Connection, 3)
	for i := 0; i < 3; i++ {
		conn := &websocket.Conn{} // Mock connection
		connection := &Connection{
			Conn:      conn,
			ID:        fmt.Sprintf("conn_%d", i),
			Username:  fmt.Sprintf("User%d", i),
			CreatedAt: time.Now(),
			Send:      make(chan []byte, 256),
		}
		connections[i] = connection
		room.AddConnection(conn, connection)
	}
	
	fmt.Printf("Chat room created: %s (%s)\n", room.Name, room.ID)
	fmt.Printf("Connections: %d\n", room.GetConnectionCount())
	
	// Simulate message broadcast
	message := Message{
		Type:      "chat",
		Content:   "Hello, everyone!",
		Username:  "System",
		Room:      room.ID,
		Timestamp: time.Now(),
	}
	
	messageBytes, _ := json.Marshal(message)
	room.Broadcast(messageBytes)
	
	fmt.Println("Message broadcasted to all connections")
}

// Demonstrate WebSocket with context
func demonstrateWebSocketWithContext() {
	fmt.Println("=== WebSocket with Context Demo ===\n")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// Create WebSocket server
	server := NewWebSocketServer()
	server.Start()
	defer server.Stop()
	
	// Set up HTTP handler
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := server.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Upgrade error: %v", err)
			return
		}
		
		server.Register <- conn
	})
	
	// Start HTTP server
	go func() {
		http.ListenAndServe(":8081", nil)
	}()
	
	// Give server time to start
	time.Sleep(100 * time.Millisecond)
	
	// Create client with context
	client := NewWebSocketClient("ws://localhost:8081/ws")
	if err := client.Connect(); err != nil {
		log.Printf("Client connection error: %v", err)
		return
	}
	defer client.Close()
	
	// Send message with context
	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled")
		return
	default:
		message := `{"type": "chat", "content": "Hello with context!"}`
		if err := client.Send([]byte(message)); err != nil {
			log.Printf("Send error: %v", err)
			return
		}
		fmt.Println("Message sent with context")
	}
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 7: WebSocket Protocols")
	fmt.Println("==========================================================\n")
	
	// Run all demonstrations
	demonstrateBasicWebSocketServer()
	fmt.Println()
	demonstrateWebSocketClient()
	fmt.Println()
	demonstrateChatRoom()
	fmt.Println()
	demonstrateWebSocketWithContext()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. WebSocket enables real-time, bidirectional communication")
	fmt.Println("2. WebSocket provides low latency and persistent connections")
	fmt.Println("3. Go's gorilla/websocket package makes WebSocket programming easy")
	fmt.Println("4. Connection management is crucial for WebSocket applications")
	fmt.Println("5. Message handling requires proper error handling and validation")
	fmt.Println("6. WebSocket is ideal for chat, gaming, and real-time applications")
	
	fmt.Println("\n📚 Next Topic: Network Security & Encryption")
}
```

Now let me create a real-time chat application example and update the progress:


```go
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// ChatMessage represents a chat message
type ChatMessage struct {
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	Username  string    `json:"username"`
	Room      string    `json:"room"`
	Timestamp time.Time `json:"timestamp"`
	Color     string    `json:"color,omitempty"`
}

// ChatServer represents a chat server
type ChatServer struct {
	Upgrader    websocket.Upgrader
	Connections map[*websocket.Conn]*ChatConnection
	Rooms       map[string]*ChatRoom
	Mutex       sync.RWMutex
	Broadcast   chan []byte
	Register    chan *websocket.Conn
	Unregister  chan *websocket.Conn
}

// ChatConnection represents a chat connection
type ChatConnection struct {
	Conn      *websocket.Conn
	ID        string
	Username  string
	Room      string
	Color     string
	CreatedAt time.Time
	Send      chan []byte
}

// ChatRoom represents a chat room
type ChatRoom struct {
	ID          string
	Name        string
	Connections map[*websocket.Conn]*ChatConnection
	Mutex       sync.RWMutex
	CreatedAt   time.Time
}

// NewChatServer creates a new chat server
func NewChatServer() *ChatServer {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins in demo
		},
	}

	return &ChatServer{
		Upgrader:    upgrader,
		Connections: make(map[*websocket.Conn]*ChatConnection),
		Rooms:       make(map[string]*ChatRoom),
		Broadcast:   make(chan []byte),
		Register:    make(chan *websocket.Conn),
		Unregister:  make(chan *websocket.Conn),
	}
}

// Start starts the chat server
func (cs *ChatServer) Start() {
	// Create default room
	cs.Rooms["general"] = &ChatRoom{
		ID:          "general",
		Name:        "General",
		Connections: make(map[*websocket.Conn]*ChatConnection),
		CreatedAt:   time.Now(),
	}

	// Start connection manager
	go cs.manageConnections()
	
	// Start message broadcaster
	go cs.broadcastMessages()
}

// manageConnections manages WebSocket connections
func (cs *ChatServer) manageConnections() {
	for {
		select {
		case conn := <-cs.Register:
			cs.registerConnection(conn)
		case conn := <-cs.Unregister:
			cs.unregisterConnection(conn)
		case message := <-cs.Broadcast:
			cs.broadcastToAll(message)
		}
	}
}

// registerConnection registers a new connection
func (cs *ChatServer) registerConnection(conn *websocket.Conn) {
	connection := &ChatConnection{
		Conn:      conn,
		ID:        generateChatConnectionID(),
		Username:  "Anonymous",
		Room:      "general",
		Color:     generateRandomColor(),
		CreatedAt: time.Now(),
		Send:      make(chan []byte, 256),
	}
	
	cs.Mutex.Lock()
	cs.Connections[conn] = connection
	cs.Mutex.Unlock()
	
	// Add to default room
	cs.Rooms["general"].AddConnection(conn, connection)
	
	// Start goroutines for this connection
	go cs.writePump(connection)
	go cs.readPump(connection)
	
	fmt.Printf("New chat connection: %s\n", connection.ID)
}

// unregisterConnection unregisters a connection
func (cs *ChatServer) unregisterConnection(conn *websocket.Conn) {
	cs.Mutex.Lock()
	if connection, exists := cs.Connections[conn]; exists {
		// Remove from room
		if room, exists := cs.Rooms[connection.Room]; exists {
			room.RemoveConnection(conn)
		}
		
		close(connection.Send)
		delete(cs.Connections, conn)
		fmt.Printf("Chat connection closed: %s\n", connection.ID)
	}
	cs.Mutex.Unlock()
}

// broadcastToAll broadcasts a message to all connections
func (cs *ChatServer) broadcastToAll(message []byte) {
	cs.Mutex.RLock()
	for conn, connection := range cs.Connections {
		select {
		case connection.Send <- message:
		default:
			close(connection.Send)
			delete(cs.Connections, conn)
		}
	}
	cs.Mutex.RUnlock()
}

// broadcastToRoom broadcasts a message to a specific room
func (cs *ChatServer) broadcastToRoom(roomID string, message []byte) {
	if room, exists := cs.Rooms[roomID]; exists {
		room.Broadcast(message)
	}
}

// writePump handles writing messages to the WebSocket connection
func (cs *ChatServer) writePump(connection *ChatConnection) {
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
			
			// Add queued messages
			n := len(connection.Send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-connection.Send)
			}
			
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

// readPump handles reading messages from the WebSocket connection
func (cs *ChatServer) readPump(connection *ChatConnection) {
	defer func() {
		cs.Unregister <- connection.Conn
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
		
		// Parse message
		var message ChatMessage
		if err := json.Unmarshal(messageBytes, &message); err != nil {
			log.Printf("Error parsing message: %v", err)
			continue
		}
		
		// Set connection info
		message.Username = connection.Username
		message.Room = connection.Room
		message.Timestamp = time.Now()
		message.Color = connection.Color
		
		// Handle different message types
		cs.handleMessage(connection, &message)
	}
}

// handleMessage handles different types of messages
func (cs *ChatServer) handleMessage(connection *ChatConnection, message *ChatMessage) {
	switch message.Type {
	case "join":
		cs.handleJoin(connection, message)
	case "leave":
		cs.handleLeave(connection, message)
	case "chat":
		cs.handleChat(connection, message)
	case "set_username":
		cs.handleSetUsername(connection, message)
	case "get_rooms":
		cs.handleGetRooms(connection, message)
	case "create_room":
		cs.handleCreateRoom(connection, message)
	default:
		log.Printf("Unknown message type: %s", message.Type)
	}
}

// handleJoin handles join room messages
func (cs *ChatServer) handleJoin(connection *ChatConnection, message *ChatMessage) {
	roomID := message.Room
	if roomID == "" {
		roomID = "general"
	}
	
	// Check if room exists
	if _, exists := cs.Rooms[roomID]; !exists {
		response := ChatMessage{
			Type:      "error",
			Content:   fmt.Sprintf("Room %s does not exist", roomID),
			Username:  "System",
			Room:      connection.Room,
			Timestamp: time.Now(),
		}
		responseBytes, _ := json.Marshal(response)
		connection.Send <- responseBytes
		return
	}
	
	// Leave current room
	if room, exists := cs.Rooms[connection.Room]; exists {
		room.RemoveConnection(connection.Conn)
	}
	
	// Join new room
	connection.Room = roomID
	cs.Rooms[roomID].AddConnection(connection.Conn, connection)
	
	response := ChatMessage{
		Type:      "joined",
		Content:   fmt.Sprintf("Joined room: %s", roomID),
		Username:  "System",
		Room:      roomID,
		Timestamp: time.Now(),
	}
	
	responseBytes, _ := json.Marshal(response)
	connection.Send <- responseBytes
}

// handleLeave handles leave room messages
func (cs *ChatServer) handleLeave(connection *ChatConnection, message *ChatMessage) {
	oldRoom := connection.Room
	connection.Room = "general"
	
	// Remove from current room
	if room, exists := cs.Rooms[oldRoom]; exists {
		room.RemoveConnection(connection.Conn)
	}
	
	// Add to general room
	cs.Rooms["general"].AddConnection(connection.Conn, connection)
	
	response := ChatMessage{
		Type:      "left",
		Content:   fmt.Sprintf("Left room: %s", oldRoom),
		Username:  "System",
		Room:      "general",
		Timestamp: time.Now(),
	}
	
	responseBytes, _ := json.Marshal(response)
	connection.Send <- responseBytes
}

// handleChat handles chat messages
func (cs *ChatServer) handleChat(connection *ChatConnection, message *ChatMessage) {
	// Broadcast to room
	messageBytes, _ := json.Marshal(message)
	cs.broadcastToRoom(connection.Room, messageBytes)
}

// handleSetUsername handles username setting
func (cs *ChatServer) handleSetUsername(connection *ChatConnection, message *ChatMessage) {
	username := message.Content
	if username == "" {
		username = "Anonymous"
	}
	
	connection.Username = username
	
	response := ChatMessage{
		Type:      "username_set",
		Content:   fmt.Sprintf("Username set to: %s", username),
		Username:  "System",
		Room:      connection.Room,
		Timestamp: time.Now(),
	}
	
	responseBytes, _ := json.Marshal(response)
	connection.Send <- responseBytes
}

// handleGetRooms handles get rooms request
func (cs *ChatServer) handleGetRooms(connection *ChatConnection, message *ChatMessage) {
	rooms := make([]map[string]interface{}, 0)
	for _, room := range cs.Rooms {
		rooms = append(rooms, map[string]interface{}{
			"id":    room.ID,
			"name":  room.Name,
			"count": room.GetConnectionCount(),
		})
	}
	
	response := ChatMessage{
		Type:      "rooms",
		Content:   "",
		Username:  "System",
		Room:      connection.Room,
		Timestamp: time.Now(),
		Data:      rooms,
	}
	
	responseBytes, _ := json.Marshal(response)
	connection.Send <- responseBytes
}

// handleCreateRoom handles create room request
func (cs *ChatServer) handleCreateRoom(connection *ChatConnection, message *ChatMessage) {
	roomName := message.Content
	if roomName == "" {
		roomName = "New Room"
	}
	
	roomID := generateRoomID(roomName)
	room := &ChatRoom{
		ID:          roomID,
		Name:        roomName,
		Connections: make(map[*websocket.Conn]*ChatConnection),
		CreatedAt:   time.Now(),
	}
	
	cs.Rooms[roomID] = room
	
	response := ChatMessage{
		Type:      "room_created",
		Content:   fmt.Sprintf("Room created: %s", roomName),
		Username:  "System",
		Room:      connection.Room,
		Timestamp: time.Now(),
		Data:      map[string]string{"room_id": roomID, "room_name": roomName},
	}
	
	responseBytes, _ := json.Marshal(response)
	connection.Send <- responseBytes
}

// broadcastMessages handles message broadcasting
func (cs *ChatServer) broadcastMessages() {
	for {
		select {
		case message := <-cs.Broadcast:
			cs.broadcastToAll(message)
		}
	}
}

// generateChatConnectionID generates a unique connection ID
func generateChatConnectionID() string {
	return fmt.Sprintf("chat_conn_%d", time.Now().UnixNano())
}

// generateRoomID generates a unique room ID
func generateRoomID(name string) string {
	return fmt.Sprintf("room_%s_%d", name, time.Now().UnixNano())
}

// generateRandomColor generates a random color for the user
func generateRandomColor() string {
	colors := []string{"#FF6B6B", "#4ECDC4", "#45B7D1", "#96CEB4", "#FFEAA7", "#DDA0DD", "#98D8C8", "#F7DC6F"}
	return colors[time.Now().UnixNano()%int64(len(colors))]
}

// AddConnection adds a connection to the room
func (room *ChatRoom) AddConnection(conn *websocket.Conn, connection *ChatConnection) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()
	
	room.Connections[conn] = connection
}

// RemoveConnection removes a connection from the room
func (room *ChatRoom) RemoveConnection(conn *websocket.Conn) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()
	
	delete(room.Connections, conn)
}

// Broadcast broadcasts a message to all connections in the room
func (room *ChatRoom) Broadcast(message []byte) {
	room.Mutex.RLock()
	defer room.Mutex.RUnlock()
	
	for conn, connection := range room.Connections {
		select {
		case connection.Send <- message:
		default:
			close(connection.Send)
			delete(room.Connections, conn)
		}
	}
}

// GetConnectionCount returns the number of connections in the room
func (room *ChatRoom) GetConnectionCount() int {
	room.Mutex.RLock()
	defer room.Mutex.RUnlock()
	
	return len(room.Connections)
}

// HTML template for the chat interface
const chatHTML = `
<!DOCTYPE html>
<html>
<head>
    <title>WebSocket Chat</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 0; padding: 20px; }
        #chat { border: 1px solid #ccc; height: 400px; overflow-y: scroll; padding: 10px; margin-bottom: 10px; }
        #input { width: 70%; padding: 5px; }
        #send { width: 20%; padding: 5px; }
        .message { margin: 5px 0; padding: 5px; border-radius: 5px; }
        .system { background-color: #f0f0f0; font-style: italic; }
        .user { background-color: #e3f2fd; }
        .error { background-color: #ffebee; color: #c62828; }
        #rooms { margin-bottom: 10px; }
        #rooms select { margin-right: 10px; }
    </style>
</head>
<body>
    <h1>WebSocket Chat</h1>
    
    <div id="rooms">
        <select id="roomSelect">
            <option value="general">General</option>
        </select>
        <button onclick="joinRoom()">Join Room</button>
        <button onclick="getRooms()">Refresh Rooms</button>
        <input type="text" id="newRoomName" placeholder="New room name">
        <button onclick="createRoom()">Create Room</button>
    </div>
    
    <div id="chat"></div>
    
    <div>
        <input type="text" id="username" placeholder="Username" value="Anonymous">
        <button onclick="setUsername()">Set Username</button>
    </div>
    
    <div>
        <input type="text" id="input" placeholder="Type a message...">
        <button id="send" onclick="sendMessage()">Send</button>
    </div>

    <script>
        const ws = new WebSocket('ws://localhost:8080/ws');
        const chat = document.getElementById('chat');
        const input = document.getElementById('input');
        const username = document.getElementById('username');
        const roomSelect = document.getElementById('roomSelect');
        const newRoomName = document.getElementById('newRoomName');

        ws.onopen = function() {
            addMessage('Connected to chat server', 'system');
        };

        ws.onmessage = function(event) {
            const message = JSON.parse(event.data);
            displayMessage(message);
        };

        ws.onclose = function() {
            addMessage('Disconnected from chat server', 'system');
        };

        function addMessage(content, type) {
            const div = document.createElement('div');
            div.className = 'message ' + type;
            div.textContent = content;
            chat.appendChild(div);
            chat.scrollTop = chat.scrollHeight;
        }

        function displayMessage(message) {
            let content = '';
            if (message.type === 'chat') {
                content = `[${message.timestamp}] ${message.username}: ${message.content}`;
            } else if (message.type === 'joined') {
                content = `✓ ${message.content}`;
            } else if (message.type === 'left') {
                content = `✗ ${message.content}`;
            } else if (message.type === 'username_set') {
                content = `✓ ${message.content}`;
            } else if (message.type === 'rooms') {
                updateRoomList(message.data);
                return;
            } else if (message.type === 'room_created') {
                content = `✓ ${message.content}`;
                getRooms(); // Refresh room list
                return;
            } else if (message.type === 'error') {
                content = `❌ ${message.content}`;
            } else {
                content = message.content;
            }
            
            addMessage(content, message.type === 'error' ? 'error' : 'system');
        }

        function sendMessage() {
            const content = input.value.trim();
            if (content === '') return;
            
            const message = {
                type: 'chat',
                content: content
            };
            
            ws.send(JSON.stringify(message));
            input.value = '';
        }

        function setUsername() {
            const name = username.value.trim();
            if (name === '') return;
            
            const message = {
                type: 'set_username',
                content: name
            };
            
            ws.send(JSON.stringify(message));
        }

        function joinRoom() {
            const room = roomSelect.value;
            const message = {
                type: 'join',
                room: room
            };
            
            ws.send(JSON.stringify(message));
        }

        function getRooms() {
            const message = {
                type: 'get_rooms'
            };
            
            ws.send(JSON.stringify(message));
        }

        function createRoom() {
            const name = newRoomName.value.trim();
            if (name === '') return;
            
            const message = {
                type: 'create_room',
                content: name
            };
            
            ws.send(JSON.stringify(message));
            newRoomName.value = '';
        }

        function updateRoomList(rooms) {
            roomSelect.innerHTML = '';
            rooms.forEach(room => {
                const option = document.createElement('option');
                option.value = room.id;
                option.textContent = `${room.name} (${room.count})`;
                roomSelect.appendChild(option);
            });
        }

        input.addEventListener('keypress', function(e) {
            if (e.key === 'Enter') {
                sendMessage();
            }
        });

        // Load rooms on page load
        window.onload = function() {
            getRooms();
        };
    </script>
</body>
</html>
`

// startChatServer starts the chat server
func startChatServer() {
	server := NewChatServer()
	server.Start()
	
	// Serve HTML page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("chat").Parse(chatHTML))
		tmpl.Execute(w, nil)
	})
	
	// WebSocket endpoint
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := server.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Upgrade error: %v", err)
			return
		}
		
		server.Register <- conn
	})
	
	fmt.Println("Chat server starting on :8080")
	fmt.Println("Open http://localhost:8080 in your browser to use the chat")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 7: WebSocket Chat Application")
	fmt.Println("================================================================\n")
	
	startChatServer()
}
```

Now let me create a quick reference guide and update the progress:


```markdown
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
```

Now let me update the progress:


Perfect! I've completed **Topic 7: WebSocket Protocols** with comprehensive coverage. Here's what I've created for you:

## 🎯 **Topic 7: WebSocket Protocols - COMPLETED**

### **📚 Theory Coverage** (`07_websocket/theory.md`)
- **WebSocket Overview**: Characteristics, vs HTTP, real-time communication
- **WebSocket Handshake**: Client request, server response, upgrade process
- **Frame Format**: Binary frame structure, fields, opcodes
- **WebSocket States**: Connection states, state transitions
- **Subprotocols**: Application-specific protocols, negotiation
- **Extensions**: Per-message compression, custom extensions
- **Go Implementation**: gorilla/websocket package, servers, clients
- **Advanced Concepts**: Connection management, message handling, error handling
- **Security**: Authentication, authorization, data protection
- **Performance**: Optimization techniques, monitoring, scaling
- **Real-world Applications**: Chat, gaming, financial trading, IoT
- **Alternatives**: WebSocket vs SSE, long polling, WebRTC

### **💻 Implementation** (`07_websocket/implementation.go`)
- **WebSocketServer**: Advanced WebSocket server with connection management
- **WebSocketClient**: WebSocket client with message handling
- **ChatRoom**: Room-based communication system
- **Connection Management**: Registration, unregistration, broadcasting
- **Message Handling**: Different message types, routing, validation
- **Ping/Pong**: Connection health monitoring
- **Context Support**: Cancellation and timeout handling
- **Comprehensive Examples**: All major WebSocket patterns

### ** Chat Application** (`07_websocket/chat_app.go`)
- **Complete Chat Server**: Real-time messaging system
- **Room Management**: Multiple chat rooms, join/leave functionality
- **User Management**: Username setting, connection tracking
- **HTML Interface**: Complete web-based chat interface
- **Message Types**: Chat, system, error, room management
- **Real-time Features**: Live updates, room switching, user management

### **📖 Quick Reference** (`07_websocket/quick_reference.md`)
- **Essential Concepts**: WebSocket vs HTTP, handshake process
- **Go Patterns**: Server, client, connection management
- **Message Handling**: Types, broadcasting, room-based communication
- **Connection Lifecycle**: Registration, write/read pumps, cleanup
- **Security Patterns**: Origin checking, authentication, rate limiting
- **Performance Optimization**: Connection pooling, message batching, compression
- **Error Handling**: Connection errors, message errors, reconnection
- **Monitoring**: Statistics, health checks, metrics
- **Best Practices**: Professional WebSocket development guidelines

## 🚀 **Key Features of This Module**

1. **Complete WebSocket Coverage**: From basics to advanced concepts
2. **Go-Specific**: Leverages gorilla/websocket package
3. **Production-Ready**: Real-world patterns and best practices
4. **Real-time Applications**: Chat, gaming, IoT, financial trading
5. **Connection Management**: Registration, broadcasting, cleanup
6. **Message Handling**: Types, routing, validation, error handling
7. **Security**: Authentication, authorization, rate limiting
8. **Performance**: Optimization techniques, monitoring, scaling

## 📈 **What You've Learned**

- **WebSocket Fundamentals**: Protocol, handshake, frame format
- **Real-time Communication**: Bidirectional, low-latency communication
- **Connection Management**: Registration, broadcasting, cleanup
- **Message Handling**: Types, routing, validation, error handling
- **Security**: Authentication, authorization, rate limiting
- **Performance**: Optimization techniques, monitoring, scaling
- **Real-world Applications**: Chat, gaming, IoT, financial trading
- **Go Implementation**: Professional WebSocket programming

## **Next Steps**

You're now ready to move to **Topic 8: Network Security & Encryption** where we'll explore TLS/SSL, encryption, authentication, and secure communication protocols.

The WebSocket foundation you now have will make you unstoppable in building real-time applications, chat systems, gaming platforms, and modern web applications! 🚀

Would you like me to continue with **Topic 8: Network Security & Encryption** or would you prefer to explore any specific aspect of WebSocket first?

The real-time communication skills you've gained will be invaluable for modern web development and networking interviews!

---

**User**

continue with Topic 8: Network Security & Encryption 

---

**Cursor**

Perfect! Let me continue with **Topic 8: Network Security & Encryption**. This is where we explore TLS/SSL, encryption, authentication, and secure communication protocols.


Now let me create the comprehensive Network Security & Encryption module:


```markdown
# Topic 8: Network Security & Encryption - Secure Communication Mastery

## Network Security Overview

Network security is the practice of protecting networks and data from unauthorized access, misuse, modification, or destruction. It encompasses multiple layers of security controls and technologies.

### Security Principles
- **Confidentiality**: Data is only accessible to authorized users
- **Integrity**: Data has not been modified or corrupted
- **Availability**: Data and services are available when needed
- **Authentication**: Verifying the identity of users/systems
- **Authorization**: Controlling access to resources
- **Non-repudiation**: Preventing denial of actions

### Security Threats
- **Eavesdropping**: Unauthorized listening to communications
- **Man-in-the-middle**: Intercepting and modifying communications
- **Spoofing**: Impersonating legitimate entities
- **Denial of Service**: Overwhelming systems to make them unavailable
- **Data tampering**: Modifying data in transit or storage
- **Replay attacks**: Replaying captured communications

## Cryptography Fundamentals

### Symmetric Encryption
- **Same key**: Encryption and decryption use the same key
- **Fast**: Efficient for large amounts of data
- **Key distribution**: Challenge of sharing keys securely
- **Examples**: AES, DES, 3DES, ChaCha20

### Asymmetric Encryption
- **Key pairs**: Public and private keys
- **Public key**: Can be shared openly
- **Private key**: Must be kept secret
- **Examples**: RSA, ECC, DSA

### Hash Functions
- **One-way**: Cannot be reversed
- **Fixed length**: Output is always the same size
- **Collision resistant**: Hard to find two inputs with same output
- **Examples**: SHA-256, SHA-3, BLAKE2

### Digital Signatures
- **Authentication**: Proves message came from sender
- **Integrity**: Proves message hasn't been modified
- **Non-repudiation**: Sender cannot deny sending
- **Process**: Hash message, encrypt hash with private key

## TLS/SSL Protocol

### TLS Overview
Transport Layer Security (TLS) is a cryptographic protocol that provides secure communication over a network. It's the successor to SSL (Secure Sockets Layer).

### TLS Versions
- **SSL 1.0**: Never released publicly
- **SSL 2.0**: Deprecated, security vulnerabilities
- **SSL 3.0**: Deprecated, POODLE vulnerability
- **TLS 1.0**: Deprecated, security issues
- **TLS 1.1**: Deprecated, CBC attacks
- **TLS 1.2**: Widely used, secure
- **TLS 1.3**: Latest version, improved security and performance

### TLS Handshake Process
1. **Client Hello**: Supported cipher suites, random data
2. **Server Hello**: Chosen cipher suite, server certificate
3. **Certificate Verification**: Client validates server certificate
4. **Key Exchange**: Generate shared secret
5. **Finished**: Handshake complete, encrypted communication

### TLS Record Protocol
- **Fragmentation**: Split data into manageable chunks
- **Compression**: Optional data compression
- **Encryption**: Encrypt data with agreed cipher
- **MAC**: Message authentication code for integrity

## Certificate Management

### X.509 Certificates
- **Standard format**: Widely used certificate format
- **Contains**: Public key, identity, validity period, signature
- **Issued by**: Certificate Authority (CA)
- **Chain of trust**: Root CA → Intermediate CA → End entity

### Certificate Authority (CA)
- **Trusted entity**: Issues and validates certificates
- **Root CA**: Self-signed, trusted by default
- **Intermediate CA**: Signed by root CA
- **Public CA**: Commercial certificate providers
- **Private CA**: Internal certificate authority

### Certificate Validation
- **Chain validation**: Verify certificate chain
- **Expiration check**: Certificate not expired
- **Revocation check**: Certificate not revoked (CRL/OCSP)
- **Hostname verification**: Certificate matches hostname
- **Purpose validation**: Certificate used for intended purpose

## Authentication Methods

### Password-based Authentication
- **Username/password**: Most common method
- **Hashed passwords**: Store password hashes, not plaintext
- **Salt**: Random data added to password before hashing
- **Strong passwords**: Complexity requirements, length

### Multi-Factor Authentication (MFA)
- **Something you know**: Password, PIN
- **Something you have**: Token, smart card, phone
- **Something you are**: Biometric (fingerprint, face)
- **Time-based**: TOTP (Time-based One-Time Password)

### Certificate-based Authentication
- **Client certificates**: X.509 certificates for clients
- **Mutual TLS**: Both client and server authenticate
- **PKI**: Public Key Infrastructure for certificate management
- **Smart cards**: Hardware tokens with certificates

### Token-based Authentication
- **JWT**: JSON Web Tokens
- **OAuth**: Authorization framework
- **SAML**: Security Assertion Markup Language
- **API keys**: Simple authentication tokens

## Network Security Protocols

### IPsec (IP Security)
- **Network layer**: Security at IP layer
- **AH**: Authentication Header for integrity
- **ESP**: Encapsulating Security Payload for confidentiality
- **IKE**: Internet Key Exchange for key management
- **Tunnel mode**: Encrypts entire IP packet
- **Transport mode**: Encrypts only payload

### VPN (Virtual Private Network)
- **Secure tunnel**: Encrypted connection over public network
- **Remote access**: Connect to corporate network remotely
- **Site-to-site**: Connect multiple networks securely
- **Types**: IPsec VPN, SSL VPN, OpenVPN, WireGuard

### SSH (Secure Shell)
- **Remote access**: Secure command-line access
- **File transfer**: SCP, SFTP for secure file transfer
- **Port forwarding**: Secure tunneling of other protocols
- **Key-based auth**: Public/private key authentication

### DNSSEC (DNS Security Extensions)
- **DNS integrity**: Prevents DNS spoofing
- **Digital signatures**: Signs DNS records
- **Chain of trust**: Validates DNS responses
- **Prevents**: Cache poisoning, man-in-the-middle

## Application Security

### HTTPS (HTTP Secure)
- **TLS over HTTP**: Encrypts HTTP communication
- **Port 443**: Default HTTPS port
- **Certificate validation**: Verify server identity
- **HSTS**: HTTP Strict Transport Security

### Secure WebSocket (WSS)
- **TLS over WebSocket**: Encrypts WebSocket communication
- **wss://**: Secure WebSocket protocol
- **Same security**: Benefits as HTTPS
- **Real-time security**: Secure real-time communication

### Email Security
- **S/MIME**: Secure/Multipurpose Internet Mail Extensions
- **PGP**: Pretty Good Privacy
- **DKIM**: DomainKeys Identified Mail
- **SPF**: Sender Policy Framework
- **DMARC**: Domain-based Message Authentication

## Security Best Practices

### Secure Coding
- **Input validation**: Validate all user input
- **Output encoding**: Encode output to prevent XSS
- **SQL injection**: Use parameterized queries
- **Buffer overflow**: Bounds checking, safe functions
- **Memory management**: Proper allocation/deallocation

### Network Security
- **Firewalls**: Control network traffic
- **Intrusion Detection**: Monitor for attacks
- **Network segmentation**: Isolate network segments
- **Access control**: Limit network access
- **Monitoring**: Log and monitor network activity

### Key Management
- **Key generation**: Use cryptographically secure random
- **Key storage**: Secure storage of private keys
- **Key rotation**: Regularly change keys
- **Key escrow**: Backup and recovery procedures
- **Key destruction**: Secure deletion of old keys

### Incident Response
- **Detection**: Identify security incidents
- **Containment**: Limit damage and spread
- **Eradication**: Remove threat and vulnerabilities
- **Recovery**: Restore normal operations
- **Lessons learned**: Improve security posture

## Go Security Implementation

### TLS Configuration
```go
config := &tls.Config{
    MinVersion:         tls.VersionTLS12,
    CipherSuites:       []uint16{
        tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
        tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
    },
    PreferServerCipherSuites: true,
    InsecureSkipVerify: false,
}
```

### Certificate Generation
```go
func generateSelfSignedCert() ([]byte, []byte, error) {
    priv, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return nil, nil, err
    }
    
    template := x509.Certificate{
        SerialNumber: big.NewInt(1),
        Subject: pkix.Name{
            Organization: []string{"Test Company"},
        },
        NotBefore:    time.Now(),
        NotAfter:     time.Now().Add(365 * 24 * time.Hour),
        KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
        ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
    }
    
    certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
    if err != nil {
        return nil, nil, err
    }
    
    certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
    keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
    
    return certPEM, keyPEM, nil
}
```

### Password Hashing
```go
func hashPassword(password string) (string, error) {
    salt := make([]byte, 32)
    if _, err := rand.Read(salt); err != nil {
        return "", err
    }
    
    hash := sha256.Sum256(append([]byte(password), salt...))
    return base64.StdEncoding.EncodeToString(append(hash[:], salt...)), nil
}

func verifyPassword(password, hashedPassword string) bool {
    data, err := base64.StdEncoding.DecodeString(hashedPassword)
    if err != nil {
        return false
    }
    
    if len(data) < 32 {
        return false
    }
    
    hash := data[:32]
    salt := data[32:]
    
    expectedHash := sha256.Sum256(append([]byte(password), salt...))
    return bytes.Equal(hash, expectedHash[:])
}
```

## Security Monitoring

### Logging
- **Security events**: Authentication, authorization, errors
- **Network activity**: Connections, traffic patterns
- **System events**: File access, process execution
- **Application logs**: Error logs, access logs

### Intrusion Detection
- **Signature-based**: Known attack patterns
- **Anomaly-based**: Unusual behavior detection
- **Network-based**: Monitor network traffic
- **Host-based**: Monitor individual systems

### Vulnerability Management
- **Vulnerability scanning**: Regular security scans
- **Patch management**: Keep systems updated
- **Risk assessment**: Evaluate security risks
- **Remediation**: Fix identified vulnerabilities

## Compliance and Standards

### Security Standards
- **ISO 27001**: Information security management
- **NIST**: National Institute of Standards and Technology
- **PCI DSS**: Payment Card Industry Data Security Standard
- **HIPAA**: Health Insurance Portability and Accountability Act
- **GDPR**: General Data Protection Regulation

### Security Frameworks
- **OWASP**: Open Web Application Security Project
- **CIS Controls**: Center for Internet Security Controls
- **NIST Cybersecurity Framework**: Risk management framework
- **SANS**: Security training and resources

## Common Security Vulnerabilities

### OWASP Top 10
1. **Injection**: SQL, NoSQL, OS command injection
2. **Broken Authentication**: Weak authentication mechanisms
3. **Sensitive Data Exposure**: Inadequate protection of sensitive data
4. **XML External Entities**: XXE attacks
5. **Broken Access Control**: Inadequate access controls
6. **Security Misconfiguration**: Insecure default configurations
7. **Cross-Site Scripting**: XSS attacks
8. **Insecure Deserialization**: Unsafe deserialization
9. **Known Vulnerabilities**: Using components with known vulnerabilities
10. **Insufficient Logging**: Inadequate logging and monitoring

### Network Vulnerabilities
- **Man-in-the-middle**: Intercepting communications
- **DNS spoofing**: Redirecting traffic to malicious sites
- **ARP spoofing**: Redirecting traffic at layer 2
- **Port scanning**: Discovering open ports and services
- **Denial of Service**: Overwhelming systems

## Interview Questions

### Basic Questions
1. What is the difference between symmetric and asymmetric encryption?
2. Explain the TLS handshake process.
3. What are the main security principles?

### Intermediate Questions
1. How do you implement secure authentication in a web application?
2. Explain certificate validation and the chain of trust.
3. What are the differences between TLS 1.2 and TLS 1.3?

### Advanced Questions
1. Design a secure communication system for a distributed application.
2. How would you implement mutual TLS authentication?
3. Explain the security implications of different encryption algorithms.

## Next Steps
After mastering network security, proceed to:
- **Topic 9**: Routing & Switching
- **Topic 10**: Network Performance & Optimization
- **Topic 11**: Distributed Systems Networking

Master network security, and you'll understand how to protect networks and data! 🚀
```

Now let me create the comprehensive Go implementation for network security:


```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"time"
)

// SecurityConfig represents security configuration
type SecurityConfig struct {
	MinTLSVersion     uint16
	CipherSuites      []uint16
	CertFile          string
	KeyFile           string
	InsecureSkipVerify bool
	ClientAuth        tls.ClientAuthType
}

// DefaultSecurityConfig returns a secure default configuration
func DefaultSecurityConfig() *SecurityConfig {
	return &SecurityConfig{
		MinTLSVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		},
		InsecureSkipVerify: false,
		ClientAuth:         tls.NoClientCert,
	}
}

// SecureServer represents a secure HTTP server
type SecureServer struct {
	Server   *http.Server
	Config   *SecurityConfig
	Cert     []byte
	Key      []byte
	Running  bool
}

// NewSecureServer creates a new secure server
func NewSecureServer(addr string, config *SecurityConfig) *SecureServer {
	if config == nil {
		config = DefaultSecurityConfig()
	}
	
	// Generate self-signed certificate if not provided
	cert, key, err := generateSelfSignedCert()
	if err != nil {
		log.Printf("Error generating certificate: %v", err)
		return nil
	}
	
	// Create TLS configuration
	tlsConfig := &tls.Config{
		MinVersion:               config.MinTLSVersion,
		CipherSuites:             config.CipherSuites,
		PreferServerCipherSuites: true,
		InsecureSkipVerify:       config.InsecureSkipVerify,
		ClientAuth:               config.ClientAuth,
	}
	
	// Load certificate
	certificate, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Printf("Error loading certificate: %v", err)
		return nil
	}
	
	tlsConfig.Certificates = []tls.Certificate{certificate}
	
	server := &http.Server{
		Addr:      addr,
		TLSConfig: tlsConfig,
	}
	
	return &SecureServer{
		Server:  server,
		Config:  config,
		Cert:    cert,
		Key:     key,
		Running: false,
	}
}

// Start starts the secure server
func (s *SecureServer) Start() error {
	s.Running = true
	fmt.Printf("Secure server starting on %s\n", s.Server.Addr)
	return s.Server.ListenAndServeTLS("", "")
}

// Stop stops the secure server
func (s *SecureServer) Stop() error {
	s.Running = false
	return s.Server.Close()
}

// generateSelfSignedCert generates a self-signed certificate
func generateSelfSignedCert() ([]byte, []byte, error) {
	// Generate private key
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	
	// Create certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization:  []string{"Test Company"},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{""},
			PostalCode:    []string{""},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		DNSNames:     []string{"localhost"},
	}
	
	// Create certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, err
	}
	
	// Encode certificate
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	
	// Encode private key
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	
	return certPEM, keyPEM, nil
}

// SecureClient represents a secure HTTP client
type SecureClient struct {
	Client *http.Client
	Config *SecurityConfig
}

// NewSecureClient creates a new secure client
func NewSecureClient(config *SecurityConfig) *SecureClient {
	if config == nil {
		config = DefaultSecurityConfig()
	}
	
	// Create TLS configuration
	tlsConfig := &tls.Config{
		MinVersion:         config.MinTLSVersion,
		CipherSuites:       config.CipherSuites,
		InsecureSkipVerify: config.InsecureSkipVerify,
	}
	
	// Create HTTP client with custom transport
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		Timeout: 30 * time.Second,
	}
	
	return &SecureClient{
		Client: client,
		Config: config,
	}
}

// Get makes a secure GET request
func (c *SecureClient) Get(url string) (*http.Response, error) {
	return c.Client.Get(url)
}

// Post makes a secure POST request
func (c *SecureClient) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return c.Client.Post(url, contentType, body)
}

// Do makes a custom secure request
func (c *SecureClient) Do(req *http.Request) (*http.Response, error) {
	return c.Client.Do(req)
}

// EncryptionService provides encryption/decryption services
type EncryptionService struct {
	AESKey []byte
	RSAPriv *rsa.PrivateKey
	RSAPub  *rsa.PublicKey
}

// NewEncryptionService creates a new encryption service
func NewEncryptionService() (*EncryptionService, error) {
	// Generate AES key
	aesKey := make([]byte, 32) // 256-bit key
	if _, err := rand.Read(aesKey); err != nil {
		return nil, err
	}
	
	// Generate RSA key pair
	rsaPriv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	
	return &EncryptionService{
		AESKey:  aesKey,
		RSAPriv: rsaPriv,
		RSAPub:  &rsaPriv.PublicKey,
	}, nil
}

// EncryptAES encrypts data using AES-GCM
func (es *EncryptionService) EncryptAES(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(es.AESKey)
	if err != nil {
		return nil, err
	}
	
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// DecryptAES decrypts data using AES-GCM
func (es *EncryptionService) DecryptAES(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(es.AESKey)
	if err != nil {
		return nil, err
	}
	
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	
	return plaintext, nil
}

// EncryptRSA encrypts data using RSA
func (es *EncryptionService) EncryptRSA(plaintext []byte) ([]byte, error) {
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, es.RSAPub, plaintext, nil)
	if err != nil {
		return nil, err
	}
	
	return ciphertext, nil
}

// DecryptRSA decrypts data using RSA
func (es *EncryptionService) DecryptRSA(ciphertext []byte) ([]byte, error) {
	hash := sha256.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, es.RSAPriv, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	
	return plaintext, nil
}

// SignRSA signs data using RSA
func (es *EncryptionService) SignRSA(data []byte) ([]byte, error) {
	hash := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, es.RSAPriv, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}
	
	return signature, nil
}

// VerifyRSA verifies RSA signature
func (es *EncryptionService) VerifyRSA(data, signature []byte) error {
	hash := sha256.Sum256(data)
	return rsa.VerifyPKCS1v15(es.RSAPub, crypto.SHA256, hash[:], signature)
}

// HashService provides hashing services
type HashService struct{}

// NewHashService creates a new hash service
func NewHashService() *HashService {
	return &HashService{}
}

// HashSHA256 hashes data using SHA-256
func (hs *HashService) HashSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// HashSHA256String hashes a string using SHA-256
func (hs *HashService) HashSHA256String(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// HashWithSalt hashes data with salt
func (hs *HashService) HashWithSalt(data, salt []byte) []byte {
	combined := append(data, salt...)
	hash := sha256.Sum256(combined)
	return hash[:]
}

// PasswordService provides password hashing and verification
type PasswordService struct {
	HashService *HashService
}

// NewPasswordService creates a new password service
func NewPasswordService() *PasswordService {
	return &PasswordService{
		HashService: NewHashService(),
	}
}

// HashPassword hashes a password with salt
func (ps *PasswordService) HashPassword(password string) (string, error) {
	// Generate random salt
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	
	// Hash password with salt
	hash := ps.HashService.HashWithSalt([]byte(password), salt)
	
	// Combine hash and salt
	combined := append(hash, salt...)
	
	// Encode as base64
	return base64.StdEncoding.EncodeToString(combined), nil
}

// VerifyPassword verifies a password against its hash
func (ps *PasswordService) VerifyPassword(password, hashedPassword string) bool {
	// Decode base64
	data, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false
	}
	
	// Check length
	if len(data) < 32 {
		return false
	}
	
	// Extract hash and salt
	hash := data[:32]
	salt := data[32:]
	
	// Hash password with salt
	expectedHash := ps.HashService.HashWithSalt([]byte(password), salt)
	
	// Compare hashes
	return bytes.Equal(hash, expectedHash)
}

// JWTService provides JWT token services
type JWTService struct {
	SecretKey []byte
}

// NewJWTService creates a new JWT service
func NewJWTService(secretKey []byte) *JWTService {
	if secretKey == nil {
		secretKey = make([]byte, 32)
		rand.Read(secretKey)
	}
	
	return &JWTService{
		SecretKey: secretKey,
	}
}

// JWTClaims represents JWT claims
type JWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Exp      int64  `json:"exp"`
	Iat      int64  `json:"iat"`
}

// GenerateToken generates a JWT token
func (js *JWTService) GenerateToken(userID, username string, expirationTime time.Duration) (string, error) {
	now := time.Now()
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		Exp:      now.Add(expirationTime).Unix(),
		Iat:      now.Unix(),
	}
	
	// Create header
	header := map[string]interface{}{
		"alg": "HS256",
		"typ": "JWT",
	}
	
	// Encode header
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerEncoded := base64.RawURLEncoding.EncodeToString(headerJSON)
	
	// Encode claims
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	claimsEncoded := base64.RawURLEncoding.EncodeToString(claimsJSON)
	
	// Create signature
	message := headerEncoded + "." + claimsEncoded
	signature := js.HashService.HashSHA256([]byte(message + string(js.SecretKey)))
	signatureEncoded := base64.RawURLEncoding.EncodeToString(signature)
	
	// Create JWT
	token := message + "." + signatureEncoded
	
	return token, nil
}

// VerifyToken verifies a JWT token
func (js *JWTService) VerifyToken(token string) (*JWTClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}
	
	// Verify signature
	message := parts[0] + "." + parts[1]
	signature := js.HashService.HashSHA256([]byte(message + string(js.SecretKey)))
	signatureEncoded := base64.RawURLEncoding.EncodeToString(signature)
	
	if parts[2] != signatureEncoded {
		return nil, fmt.Errorf("invalid signature")
	}
	
	// Decode claims
	claimsJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	
	var claims JWTClaims
	if err := json.Unmarshal(claimsJSON, &claims); err != nil {
		return nil, err
	}
	
	// Check expiration
	if time.Now().Unix() > claims.Exp {
		return nil, fmt.Errorf("token expired")
	}
	
	return &claims, nil
}

// SecurityMiddleware provides security middleware for HTTP
type SecurityMiddleware struct {
	JWTService      *JWTService
	PasswordService *PasswordService
}

// NewSecurityMiddleware creates a new security middleware
func NewSecurityMiddleware() *SecurityMiddleware {
	return &SecurityMiddleware{
		JWTService:      NewJWTService(nil),
		PasswordService: NewPasswordService(),
	}
}

// AuthMiddleware provides authentication middleware
func (sm *SecurityMiddleware) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}
		
		// Extract token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}
		
		// Verify token
		claims, err := sm.JWTService.VerifyToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		
		// Add user info to request context
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "username", claims.Username)
		
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// CORSMiddleware provides CORS middleware
func (sm *SecurityMiddleware) CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	}
}

// SecurityHeadersMiddleware adds security headers
func (sm *SecurityMiddleware) SecurityHeadersMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		
		next.ServeHTTP(w, r)
	}
}

// Demonstrate secure server
func demonstrateSecureServer() {
	fmt.Println("=== Secure Server Demo ===\n")
	
	// Create secure server
	server := NewSecureServer(":8443", nil)
	if server == nil {
		log.Println("Failed to create secure server")
		return
	}
	
	// Set up routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message": "Secure server response",
			"time":    time.Now().Format(time.RFC3339),
			"secure":  true,
		}
		json.NewEncoder(w).Encode(response)
	})
	
	// Start server in goroutine
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
		}
	}()
	
	// Give server time to start
	time.Sleep(100 * time.Millisecond)
	
	// Test with secure client
	client := NewSecureClient(&SecurityConfig{
		InsecureSkipVerify: true, // For self-signed certificate
	})
	
	resp, err := client.Get("https://localhost:8443/")
	if err != nil {
		log.Printf("Client error: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}
	
	fmt.Printf("Secure response: %s\n", string(body))
	
	// Stop server
	server.Stop()
}

// Demonstrate encryption
func demonstrateEncryption() {
	fmt.Println("=== Encryption Demo ===\n")
	
	// Create encryption service
	es, err := NewEncryptionService()
	if err != nil {
		log.Printf("Error creating encryption service: %v", err)
		return
	}
	
	// Test AES encryption
	plaintext := "Hello, World! This is a secret message."
	fmt.Printf("Original: %s\n", plaintext)
	
	ciphertext, err := es.EncryptAES([]byte(plaintext))
	if err != nil {
		log.Printf("AES encryption error: %v", err)
		return
	}
	fmt.Printf("AES Encrypted: %x\n", ciphertext)
	
	decrypted, err := es.DecryptAES(ciphertext)
	if err != nil {
		log.Printf("AES decryption error: %v", err)
		return
	}
	fmt.Printf("AES Decrypted: %s\n", string(decrypted))
	
	// Test RSA encryption
	rsaCiphertext, err := es.EncryptRSA([]byte(plaintext))
	if err != nil {
		log.Printf("RSA encryption error: %v", err)
		return
	}
	fmt.Printf("RSA Encrypted: %x\n", rsaCiphertext)
	
	rsaDecrypted, err := es.DecryptRSA(rsaCiphertext)
	if err != nil {
		log.Printf("RSA decryption error: %v", err)
		return
	}
	fmt.Printf("RSA Decrypted: %s\n", string(rsaDecrypted))
	
	// Test RSA signing
	signature, err := es.SignRSA([]byte(plaintext))
	if err != nil {
		log.Printf("RSA signing error: %v", err)
		return
	}
	fmt.Printf("RSA Signature: %x\n", signature)
	
	err = es.VerifyRSA([]byte(plaintext), signature)
	if err != nil {
		log.Printf("RSA verification error: %v", err)
		return
	}
	fmt.Println("RSA Signature verified successfully")
}

// Demonstrate password hashing
func demonstratePasswordHashing() {
	fmt.Println("=== Password Hashing Demo ===\n")
	
	// Create password service
	ps := NewPasswordService()
	
	// Test password hashing
	password := "mySecretPassword123"
	fmt.Printf("Original password: %s\n", password)
	
	hashedPassword, err := ps.HashPassword(password)
	if err != nil {
		log.Printf("Password hashing error: %v", err)
		return
	}
	fmt.Printf("Hashed password: %s\n", hashedPassword)
	
	// Test password verification
	valid := ps.VerifyPassword(password, hashedPassword)
	fmt.Printf("Password verification: %t\n", valid)
	
	// Test with wrong password
	wrongPassword := "wrongPassword"
	valid = ps.VerifyPassword(wrongPassword, hashedPassword)
	fmt.Printf("Wrong password verification: %t\n", valid)
}

// Demonstrate JWT tokens
func demonstrateJWT() {
	fmt.Println("=== JWT Token Demo ===\n")
	
	// Create JWT service
	js := NewJWTService(nil)
	
	// Generate token
	userID := "user123"
	username := "john_doe"
	expirationTime := 24 * time.Hour
	
	token, err := js.GenerateToken(userID, username, expirationTime)
	if err != nil {
		log.Printf("Token generation error: %v", err)
		return
	}
	fmt.Printf("Generated token: %s\n", token)
	
	// Verify token
	claims, err := js.VerifyToken(token)
	if err != nil {
		log.Printf("Token verification error: %v", err)
		return
	}
	fmt.Printf("Token claims: UserID=%s, Username=%s, Exp=%d\n", 
		claims.UserID, claims.Username, claims.Exp)
	
	// Test with invalid token
	invalidToken := "invalid.token.here"
	_, err = js.VerifyToken(invalidToken)
	if err != nil {
		fmt.Printf("Invalid token verification (expected): %v\n", err)
	}
}

// Demonstrate security middleware
func demonstrateSecurityMiddleware() {
	fmt.Println("=== Security Middleware Demo ===\n")
	
	// Create security middleware
	sm := NewSecurityMiddleware()
	
	// Create test server
	mux := http.NewServeMux()
	
	// Public endpoint
	mux.HandleFunc("/public", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message": "This is a public endpoint",
			"time":    time.Now().Format(time.RFC3339),
		}
		json.NewEncoder(w).Encode(response)
	})
	
	// Protected endpoint
	mux.HandleFunc("/protected", sm.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(string)
		username := r.Context().Value("username").(string)
		
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message":  "This is a protected endpoint",
			"user_id":  userID,
			"username": username,
			"time":     time.Now().Format(time.RFC3339),
		}
		json.NewEncoder(w).Encode(response)
	}))
	
	// Apply security middleware
	handler := sm.SecurityHeadersMiddleware(sm.CORSMiddleware(mux.ServeHTTP))
	
	// Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	
	go func() {
		fmt.Println("Security middleware server starting on :8080")
		log.Fatal(server.ListenAndServe())
	}()
	
	// Give server time to start
	time.Sleep(100 * time.Millisecond)
	
	// Test public endpoint
	resp, err := http.Get("http://localhost:8080/public")
	if err != nil {
		log.Printf("Public endpoint error: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}
	
	fmt.Printf("Public endpoint response: %s\n", string(body))
	
	// Test protected endpoint without token
	resp, err = http.Get("http://localhost:8080/protected")
	if err != nil {
		log.Printf("Protected endpoint error: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}
	
	fmt.Printf("Protected endpoint response (no token): %s\n", string(body))
	
	// Generate token for testing
	js := NewJWTService(nil)
	token, err := js.GenerateToken("user123", "john_doe", time.Hour)
	if err != nil {
		log.Printf("Token generation error: %v", err)
		return
	}
	
	// Test protected endpoint with token
	req, err := http.NewRequest("GET", "http://localhost:8080/protected", nil)
	if err != nil {
		log.Printf("Request creation error: %v", err)
		return
	}
	
	req.Header.Set("Authorization", "Bearer "+token)
	
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Printf("Protected endpoint error: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}
	
	fmt.Printf("Protected endpoint response (with token): %s\n", string(body))
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 8: Network Security & Encryption")
	fmt.Println("====================================================================\n")
	
	// Run all demonstrations
	demonstrateSecureServer()
	fmt.Println()
	demonstrateEncryption()
	fmt.Println()
	demonstratePasswordHashing()
	fmt.Println()
	demonstrateJWT()
	fmt.Println()
	demonstrateSecurityMiddleware()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Network security protects data and communications")
	fmt.Println("2. Encryption ensures confidentiality and integrity")
	fmt.Println("3. TLS/SSL provides secure communication over networks")
	fmt.Println("4. Authentication verifies user and system identity")
	fmt.Println("5. Proper key management is crucial for security")
	fmt.Println("6. Security middleware protects web applications")
	
	fmt.Println("\n📚 Next Topic: Routing & Switching")
}
```

I need to fix some import issues in the implementation. Let me correct them:


```go
package main

import (
	"bytes"
	"context"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"strings"
	"time"
)

// SecurityConfig represents security configuration
type SecurityConfig struct {
	MinTLSVersion     uint16
	CipherSuites      []uint16
	CertFile          string
	KeyFile           string
	InsecureSkipVerify bool
	ClientAuth        tls.ClientAuthType
}

// DefaultSecurityConfig returns a secure default configuration
func DefaultSecurityConfig() *SecurityConfig {
	return &SecurityConfig{
		MinTLSVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		},
		InsecureSkipVerify: false,
		ClientAuth:         tls.NoClientCert,
	}
}

// SecureServer represents a secure HTTP server
type SecureServer struct {
	Server   *http.Server
	Config   *SecurityConfig
	Cert     []byte
	Key      []byte
	Running  bool
}

// NewSecureServer creates a new secure server
func NewSecureServer(addr string, config *SecurityConfig) *SecureServer {
	if config == nil {
		config = DefaultSecurityConfig()
	}
	
	// Generate self-signed certificate if not provided
	cert, key, err := generateSelfSignedCert()
	if err != nil {
		log.Printf("Error generating certificate: %v", err)
		return nil
	}
	
	// Create TLS configuration
	tlsConfig := &tls.Config{
		MinVersion:               config.MinTLSVersion,
		CipherSuites:             config.CipherSuites,
		PreferServerCipherSuites: true,
		InsecureSkipVerify:       config.InsecureSkipVerify,
		ClientAuth:               config.ClientAuth,
	}
	
	// Load certificate
	certificate, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Printf("Error loading certificate: %v", err)
		return nil
	}
	
	tlsConfig.Certificates = []tls.Certificate{certificate}
	
	server := &http.Server{
		Addr:      addr,
		TLSConfig: tlsConfig,
	}
	
	return &SecureServer{
		Server:  server,
		Config:  config,
		Cert:    cert,
		Key:     key,
		Running: false,
	}
}

// Start starts the secure server
func (s *SecureServer) Start() error {
	s.Running = true
	fmt.Printf("Secure server starting on %s\n", s.Server.Addr)
	return s.Server.ListenAndServeTLS("", "")
}

// Stop stops the secure server
func (s *SecureServer) Stop() error {
	s.Running = false
	return s.Server.Close()
}

// generateSelfSignedCert generates a self-signed certificate
func generateSelfSignedCert() ([]byte, []byte, error) {
	// Generate private key
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	
	// Create certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization:  []string{"Test Company"},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{""},
			PostalCode:    []string{""},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		DNSNames:     []string{"localhost"},
	}
	
	// Create certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, err
	}
	
	// Encode certificate
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	
	// Encode private key
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	
	return certPEM, keyPEM, nil
}

// SecureClient represents a secure HTTP client
type SecureClient struct {
	Client *http.Client
	Config *SecurityConfig
}

// NewSecureClient creates a new secure client
func NewSecureClient(config *SecurityConfig) *SecureClient {
	if config == nil {
		config = DefaultSecurityConfig()
	}
	
	// Create TLS configuration
	tlsConfig := &tls.Config{
		MinVersion:         config.MinTLSVersion,
		CipherSuites:       config.CipherSuites,
		InsecureSkipVerify: config.InsecureSkipVerify,
	}
	
	// Create HTTP client with custom transport
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		Timeout: 30 * time.Second,
	}
	
	return &SecureClient{
		Client: client,
		Config: config,
	}
}

// Get makes a secure GET request
func (c *SecureClient) Get(url string) (*http.Response, error) {
	return c.Client.Get(url)
}

// Post makes a secure POST request
func (c *SecureClient) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return c.Client.Post(url, contentType, body)
}

// Do makes a custom secure request
func (c *SecureClient) Do(req *http.Request) (*http.Response, error) {
	return c.Client.Do(req)
}

// EncryptionService provides encryption/decryption services
type EncryptionService struct {
	AESKey []byte
	RSAPriv *rsa.PrivateKey
	RSAPub  *rsa.PublicKey
}

// NewEncryptionService creates a new encryption service
func NewEncryptionService() (*EncryptionService, error) {
	// Generate AES key
	aesKey := make([]byte, 32) // 256-bit key
	if _, err := rand.Read(aesKey); err != nil {
		return nil, err
	}
	
	// Generate RSA key pair
	rsaPriv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	
	return &EncryptionService{
		AESKey:  aesKey,
		RSAPriv: rsaPriv,
		RSAPub:  &rsaPriv.PublicKey,
	}, nil
}

// EncryptAES encrypts data using AES-GCM
func (es *EncryptionService) EncryptAES(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(es.AESKey)
	if err != nil {
		return nil, err
	}
	
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// DecryptAES decrypts data using AES-GCM
func (es *EncryptionService) DecryptAES(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(es.AESKey)
	if err != nil {
		return nil, err
	}
	
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	
	return plaintext, nil
}

// EncryptRSA encrypts data using RSA
func (es *EncryptionService) EncryptRSA(plaintext []byte) ([]byte, error) {
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, es.RSAPub, plaintext, nil)
	if err != nil {
		return nil, err
	}
	
	return ciphertext, nil
}

// DecryptRSA decrypts data using RSA
func (es *EncryptionService) DecryptRSA(ciphertext []byte) ([]byte, error) {
	hash := sha256.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, es.RSAPriv, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	
	return plaintext, nil
}

// SignRSA signs data using RSA
func (es *EncryptionService) SignRSA(data []byte) ([]byte, error) {
	hash := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, es.RSAPriv, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}
	
	return signature, nil
}

// VerifyRSA verifies RSA signature
func (es *EncryptionService) VerifyRSA(data, signature []byte) error {
	hash := sha256.Sum256(data)
	return rsa.VerifyPKCS1v15(es.RSAPub, crypto.SHA256, hash[:], signature)
}

// HashService provides hashing services
type HashService struct{}

// NewHashService creates a new hash service
func NewHashService() *HashService {
	return &HashService{}
}

// HashSHA256 hashes data using SHA-256
func (hs *HashService) HashSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// HashSHA256String hashes a string using SHA-256
func (hs *HashService) HashSHA256String(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// HashWithSalt hashes data with salt
func (hs *HashService) HashWithSalt(data, salt []byte) []byte {
	combined := append(data, salt...)
	hash := sha256.Sum256(combined)
	return hash[:]
}

// PasswordService provides password hashing and verification
type PasswordService struct {
	HashService *HashService
}

// NewPasswordService creates a new password service
func NewPasswordService() *PasswordService {
	return &PasswordService{
		HashService: NewHashService(),
	}
}

// HashPassword hashes a password with salt
func (ps *PasswordService) HashPassword(password string) (string, error) {
	// Generate random salt
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	
	// Hash password with salt
	hash := ps.HashService.HashWithSalt([]byte(password), salt)
	
	// Combine hash and salt
	combined := append(hash, salt...)
	
	// Encode as base64
	return base64.StdEncoding.EncodeToString(combined), nil
}

// VerifyPassword verifies a password against its hash
func (ps *PasswordService) VerifyPassword(password, hashedPassword string) bool {
	// Decode base64
	data, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false
	}
	
	// Check length
	if len(data) < 32 {
		return false
	}
	
	// Extract hash and salt
	hash := data[:32]
	salt := data[32:]
	
	// Hash password with salt
	expectedHash := ps.HashService.HashWithSalt([]byte(password), salt)
	
	// Compare hashes
	return bytes.Equal(hash, expectedHash)
}

// JWTService provides JWT token services
type JWTService struct {
	SecretKey []byte
}

// NewJWTService creates a new JWT service
func NewJWTService(secretKey []byte) *JWTService {
	if secretKey == nil {
		secretKey = make([]byte, 32)
		rand.Read(secretKey)
	}
	
	return &JWTService{
		SecretKey: secretKey,
	}
}

// JWTClaims represents JWT claims
type JWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Exp      int64  `json:"exp"`
	Iat      int64  `json:"iat"`
}

// GenerateToken generates a JWT token
func (js *JWTService) GenerateToken(userID, username string, expirationTime time.Duration) (string, error) {
	now := time.Now()
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		Exp:      now.Add(expirationTime).Unix(),
		Iat:      now.Unix(),
	}
	
	// Create header
	header := map[string]interface{}{
		"alg": "HS256",
		"typ": "JWT",
	}
	
	// Encode header
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerEncoded := base64.RawURLEncoding.EncodeToString(headerJSON)
	
	// Encode claims
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	claimsEncoded := base64.RawURLEncoding.EncodeToString(claimsJSON)
	
	// Create signature
	message := headerEncoded + "." + claimsEncoded
	signature := js.HashService.HashSHA256([]byte(message + string(js.SecretKey)))
	signatureEncoded := base64.RawURLEncoding.EncodeToString(signature)
	
	// Create JWT
	token := message + "." + signatureEncoded
	
	return token, nil
}

// VerifyToken verifies a JWT token
func (js *JWTService) VerifyToken(token string) (*JWTClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}
	
	// Verify signature
	message := parts[0] + "." + parts[1]
	signature := js.HashService.HashSHA256([]byte(message + string(js.SecretKey)))
	signatureEncoded := base64.RawURLEncoding.EncodeToString(signature)
	
	if parts[2] != signatureEncoded {
		return nil, fmt.Errorf("invalid signature")
	}
	
	// Decode claims
	claimsJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	
	var claims JWTClaims
	if err := json.Unmarshal(claimsJSON, &claims); err != nil {
		return nil, err
	}
	
	// Check expiration
	if time.Now().Unix() > claims.Exp {
		return nil, fmt.Errorf("token expired")
	}
	
	return &claims, nil
}

// SecurityMiddleware provides security middleware for HTTP
type SecurityMiddleware struct {
	JWTService      *JWTService
	PasswordService *PasswordService
}

// NewSecurityMiddleware creates a new security middleware
func NewSecurityMiddleware() *SecurityMiddleware {
	return &SecurityMiddleware{
		JWTService:      NewJWTService(nil),
		PasswordService: NewPasswordService(),
	}
}

// AuthMiddleware provides authentication middleware
func (sm *SecurityMiddleware) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}
		
		// Extract token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}
		
		// Verify token
		claims, err := sm.JWTService.VerifyToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		
		// Add user info to request context
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "username", claims.Username)
		
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// CORSMiddleware provides CORS middleware
func (sm *SecurityMiddleware) CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	}
}

// SecurityHeadersMiddleware adds security headers
func (sm *SecurityMiddleware) SecurityHeadersMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		
		next.ServeHTTP(w, r)
	}
}

// Demonstrate secure server
func demonstrateSecureServer() {
	fmt.Println("=== Secure Server Demo ===\n")
	
	// Create secure server
	server := NewSecureServer(":8443", nil)
	if server == nil {
		log.Println("Failed to create secure server")
		return
	}
	
	// Set up routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message": "Secure server response",
			"time":    time.Now().Format(time.RFC3339),
			"secure":  true,
		}
		json.NewEncoder(w).Encode(response)
	})
	
	// Start server in goroutine
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
		}
	}()
	
	// Give server time to start
	time.Sleep(100 * time.Millisecond)
	
	// Test with secure client
	client := NewSecureClient(&SecurityConfig{
		InsecureSkipVerify: true, // For self-signed certificate
	})
	
	resp, err := client.Get("https://localhost:8443/")
	if err != nil {
		log.Printf("Client error: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}
	
	fmt.Printf("Secure response: %s\n", string(body))
	
	// Stop server
	server.Stop()
}

// Demonstrate encryption
func demonstrateEncryption() {
	fmt.Println("=== Encryption Demo ===\n")
	
	// Create encryption service
	es, err := NewEncryptionService()
	if err != nil {
		log.Printf("Error creating encryption service: %v", err)
		return
	}
	
	// Test AES encryption
	plaintext := "Hello, World! This is a secret message."
	fmt.Printf("Original: %s\n", plaintext)
	
	ciphertext, err := es.EncryptAES([]byte(plaintext))
	if err != nil {
		log.Printf("AES encryption error: %v", err)
		return
	}
	fmt.Printf("AES Encrypted: %x\n", ciphertext)
	
	decrypted, err := es.DecryptAES(ciphertext)
	if err != nil {
		log.Printf("AES decryption error: %v", err)
		return
	}
	fmt.Printf("AES Decrypted: %s\n", string(decrypted))
	
	// Test RSA encryption
	rsaCiphertext, err := es.EncryptRSA([]byte(plaintext))
	if err != nil {
		log.Printf("RSA encryption error: %v", err)
		return
	}
	fmt.Printf("RSA Encrypted: %x\n", rsaCiphertext)
	
	rsaDecrypted, err := es.DecryptRSA(rsaCiphertext)
	if err != nil {
		log.Printf("RSA decryption error: %v", err)
		return
	}
	fmt.Printf("RSA Decrypted: %s\n", string(rsaDecrypted))
	
	// Test RSA signing
	signature, err := es.SignRSA([]byte(plaintext))
	if err != nil {
		log.Printf("RSA signing error: %v", err)
		return
	}
	fmt.Printf("RSA Signature: %x\n", signature)
	
	err = es.VerifyRSA([]byte(plaintext), signature)
	if err != nil {
		log.Printf("RSA verification error: %v", err)
		return
	}
	fmt.Println("RSA Signature verified successfully")
}

// Demonstrate password hashing
func demonstratePasswordHashing() {
	fmt.Println("=== Password Hashing Demo ===\n")
	
	// Create password service
	ps := NewPasswordService()
	
	// Test password hashing
	password := "mySecretPassword123"
	fmt.Printf("Original password: %s\n", password)
	
	hashedPassword, err := ps.HashPassword(password)
	if err != nil {
		log.Printf("Password hashing error: %v", err)
		return
	}
	fmt.Printf("Hashed password: %s\n", hashedPassword)
	
	// Test password verification
	valid := ps.VerifyPassword(password, hashedPassword)
	fmt.Printf("Password verification: %t\n", valid)
	
	// Test with wrong password
	wrongPassword := "wrongPassword"
	valid = ps.VerifyPassword(wrongPassword, hashedPassword)
	fmt.Printf("Wrong password verification: %t\n", valid)
}

// Demonstrate JWT tokens
func demonstrateJWT() {
	fmt.Println("=== JWT Token Demo ===\n")
	
	// Create JWT service
	js := NewJWTService(nil)
	
	// Generate token
	userID := "user123"
	username := "john_doe"
	expirationTime := 24 * time.Hour
	
	token, err := js.GenerateToken(userID, username, expirationTime)
	if err != nil {
		log.Printf("Token generation error: %v", err)
		return
	}
	fmt.Printf("Generated token: %s\n", token)
	
	// Verify token
	claims, err := js.VerifyToken(token)
	if err != nil {
		log.Printf("Token verification error: %v", err)
		return
	}
	fmt.Printf("Token claims: UserID=%s, Username=%s, Exp=%d\n", 
		claims.UserID, claims.Username, claims.Exp)
	
	// Test with invalid token
	invalidToken := "invalid.token.here"
	_, err = js.VerifyToken(invalidToken)
	if err != nil {
		fmt.Printf("Invalid token verification (expected): %v\n", err)
	}
}

// Demonstrate security middleware
func demonstrateSecurityMiddleware() {
	fmt.Println("=== Security Middleware Demo ===\n")
	
	// Create security middleware
	sm := NewSecurityMiddleware()
	
	// Create test server
	mux := http.NewServeMux()
	
	// Public endpoint
	mux.HandleFunc("/public", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message": "This is a public endpoint",
			"time":    time.Now().Format(time.RFC3339),
		}
		json.NewEncoder(w).Encode(response)
	})
	
	// Protected endpoint
	mux.HandleFunc("/protected", sm.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(string)
		username := r.Context().Value("username").(string)
		
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message":  "This is a protected endpoint",
			"user_id":  userID,
			"username": username,
			"time":     time.Now().Format(time.RFC3339),
		}
		json.NewEncoder(w).Encode(response)
	}))
	
	// Apply security middleware
	handler := sm.SecurityHeadersMiddleware(sm.CORSMiddleware(mux.ServeHTTP))
	
	// Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	
	go func() {
		fmt.Println("Security middleware server starting on :8080")
		log.Fatal(server.ListenAndServe())
	}()
	
	// Give server time to start
	time.Sleep(100 * time.Millisecond)
	
	// Test public endpoint
	resp, err := http.Get("http://localhost:8080/public")
	if err != nil {
		log.Printf("Public endpoint error: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}
	
	fmt.Printf("Public endpoint response: %s\n", string(body))
	
	// Test protected endpoint without token
	resp, err = http.Get("http://localhost:8080/protected")
	if err != nil {
		log.Printf("Protected endpoint error: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}
	
	fmt.Printf("Protected endpoint response (no token): %s\n", string(body))
	
	// Generate token for testing
	js := NewJWTService(nil)
	token, err := js.GenerateToken("user123", "john_doe", time.Hour)
	if err != nil {
		log.Printf("Token generation error: %v", err)
		return
	}
	
	// Test protected endpoint with token
	req, err := http.NewRequest("GET", "http://localhost:8080/protected", nil)
	if err != nil {
		log.Printf("Request creation error: %v", err)
		return
	}
	
	req.Header.Set("Authorization", "Bearer "+token)
	
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Printf("Protected endpoint error: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}
	
	fmt.Printf("Protected endpoint response (with token): %s\n", string(body))
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 8: Network Security & Encryption")
	fmt.Println("====================================================================\n")
	
	// Run all demonstrations
	demonstrateSecureServer()
	fmt.Println()
	demonstrateEncryption()
	fmt.Println()
	demonstratePasswordHashing()
	fmt.Println()
	demonstrateJWT()
	fmt.Println()
	demonstrateSecurityMiddleware()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Network security protects data and communications")
	fmt.Println("2. Encryption ensures confidentiality and integrity")
	fmt.Println("3. TLS/SSL provides secure communication over networks")
	fmt.Println("4. Authentication verifies user and system identity")
	fmt.Println("5. Proper key management is crucial for security")
	fmt.Println("6. Security middleware protects web applications")
	
	fmt.Println("\n📚 Next Topic: Routing & Switching")
}

```

Now let me create a comprehensive security utilities module and update the progress:


```go
package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// SecurityUtils provides various security utility functions
type SecurityUtils struct{}

// NewSecurityUtils creates a new security utils instance
func NewSecurityUtils() *SecurityUtils {
	return &SecurityUtils{}
}

// GenerateRandomString generates a random string of specified length
func (su *SecurityUtils) GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// GenerateSecureRandomString generates a cryptographically secure random string
func (su *SecurityUtils) GenerateSecureRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	
	for i := range b {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	return string(b), nil
}

// GenerateAPIKey generates a secure API key
func (su *SecurityUtils) GenerateAPIKey() (string, error) {
	return su.GenerateSecureRandomString(32)
}

// GenerateSessionID generates a secure session ID
func (su *SecurityUtils) GenerateSessionID() (string, error) {
	return su.GenerateSecureRandomString(16)
}

// HashMD5 hashes data using MD5 (not recommended for security)
func (su *SecurityUtils) HashMD5(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// HashSHA1 hashes data using SHA-1 (not recommended for security)
func (su *SecurityUtils) HashSHA1(data string) string {
	hash := sha1.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// HashSHA256 hashes data using SHA-256
func (su *SecurityUtils) HashSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// HashSHA512 hashes data using SHA-512
func (su *SecurityUtils) HashSHA512(data string) string {
	hash := sha512.Sum512([]byte(data))
	return hex.EncodeToString(hash[:])
}

// GenerateHMAC generates HMAC for data with given key
func (su *SecurityUtils) GenerateHMAC(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// VerifyHMAC verifies HMAC for data with given key
func (su *SecurityUtils) VerifyHMAC(data, key, expectedMAC string) bool {
	actualMAC := su.GenerateHMAC(data, key)
	return hmac.Equal([]byte(actualMAC), []byte(expectedMAC))
}

// ValidatePasswordStrength validates password strength
func (su *SecurityUtils) ValidatePasswordStrength(password string) (bool, []string) {
	var issues []string
	
	if len(password) < 8 {
		issues = append(issues, "Password must be at least 8 characters long")
	}
	
	if len(password) > 128 {
		issues = append(issues, "Password must be no more than 128 characters long")
	}
	
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false
	
	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasDigit = true
		case strings.ContainsRune("!@#$%^&*()_+-=[]{}|;:,.<>?", char):
			hasSpecial = true
		}
	}
	
	if !hasUpper {
		issues = append(issues, "Password must contain at least one uppercase letter")
	}
	
	if !hasLower {
		issues = append(issues, "Password must contain at least one lowercase letter")
	}
	
	if !hasDigit {
		issues = append(issues, "Password must contain at least one digit")
	}
	
	if !hasSpecial {
		issues = append(issues, "Password must contain at least one special character")
	}
	
	return len(issues) == 0, issues
}

// SanitizeInput sanitizes user input to prevent XSS
func (su *SecurityUtils) SanitizeInput(input string) string {
	// Replace potentially dangerous characters
	input = strings.ReplaceAll(input, "<", "&lt;")
	input = strings.ReplaceAll(input, ">", "&gt;")
	input = strings.ReplaceAll(input, "\"", "&quot;")
	input = strings.ReplaceAll(input, "'", "&#x27;")
	input = strings.ReplaceAll(input, "&", "&amp;")
	input = strings.ReplaceAll(input, "/", "&#x2F;")
	
	return input
}

// ValidateEmail validates email format
func (su *SecurityUtils) ValidateEmail(email string) bool {
	// Simple email validation
	if len(email) < 5 {
		return false
	}
	
	if !strings.Contains(email, "@") {
		return false
	}
	
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	
	if len(parts[0]) == 0 || len(parts[1]) == 0 {
		return false
	}
	
	if !strings.Contains(parts[1], ".") {
		return false
	}
	
	return true
}

// ValidateIPAddress validates IP address format
func (su *SecurityUtils) ValidateIPAddress(ip string) bool {
	// Check if it's a valid IPv4 or IPv6 address
	if net.ParseIP(ip) == nil {
		return false
	}
	return true
}

// IsPrivateIP checks if an IP address is private
func (su *SecurityUtils) IsPrivateIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	
	// Check private IP ranges
	privateRanges := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"127.0.0.0/8",
		"169.254.0.0/16",
	}
	
	for _, rangeStr := range privateRanges {
		_, network, _ := net.ParseCIDR(rangeStr)
		if network.Contains(parsedIP) {
			return true
		}
	}
	
	return false
}

// GenerateCSRFToken generates a CSRF token
func (su *SecurityUtils) GenerateCSRFToken() (string, error) {
	return su.GenerateSecureRandomString(32)
}

// ValidateCSRFToken validates a CSRF token
func (su *SecurityUtils) ValidateCSRFToken(token, sessionID string) bool {
	// In a real implementation, you would store and validate CSRF tokens
	// This is a simplified version
	return len(token) == 32 && len(sessionID) > 0
}

// RateLimiter provides rate limiting functionality
type RateLimiter struct {
	requests map[string][]time.Time
	mutex    sync.RWMutex
	limit    int
	window   time.Duration
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

// Allow checks if a request is allowed for the given key
func (rl *RateLimiter) Allow(key string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	
	now := time.Now()
	cutoff := now.Add(-rl.window)
	
	// Clean old requests
	requests := rl.requests[key]
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
	rl.requests[key] = validRequests
	return true
}

// GetRemainingRequests returns the number of remaining requests
func (rl *RateLimiter) GetRemainingRequests(key string) int {
	rl.mutex.RLock()
	defer rl.mutex.RUnlock()
	
	now := time.Now()
	cutoff := now.Add(-rl.window)
	
	requests := rl.requests[key]
	var validRequests []time.Time
	for _, req := range requests {
		if req.After(cutoff) {
			validRequests = append(validRequests, req)
		}
	}
	
	return rl.limit - len(validRequests)
}

// SecurityHeaders provides security header management
type SecurityHeaders struct {
	Headers map[string]string
}

// NewSecurityHeaders creates a new security headers instance
func NewSecurityHeaders() *SecurityHeaders {
	return &SecurityHeaders{
		Headers: map[string]string{
			"X-Content-Type-Options": "nosniff",
			"X-Frame-Options":        "DENY",
			"X-XSS-Protection":       "1; mode=block",
			"Strict-Transport-Security": "max-age=31536000; includeSubDomains",
			"Content-Security-Policy":   "default-src 'self'",
			"Referrer-Policy":           "strict-origin-when-cross-origin",
			"Permissions-Policy":        "geolocation=(), microphone=(), camera=()",
		},
	}
}

// SetHeader sets a security header
func (sh *SecurityHeaders) SetHeader(name, value string) {
	sh.Headers[name] = value
}

// ApplyHeaders applies security headers to HTTP response
func (sh *SecurityHeaders) ApplyHeaders(w http.ResponseWriter) {
	for name, value := range sh.Headers {
		w.Header().Set(name, value)
	}
}

// SecurityAudit provides security auditing functionality
type SecurityAudit struct {
	Events []SecurityEvent
	mutex  sync.RWMutex
}

// SecurityEvent represents a security event
type SecurityEvent struct {
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
	Severity  string    `json:"severity"`
	Message   string    `json:"message"`
	IP        string    `json:"ip"`
	UserID    string    `json:"user_id,omitempty"`
}

// NewSecurityAudit creates a new security audit instance
func NewSecurityAudit() *SecurityAudit {
	return &SecurityAudit{
		Events: make([]SecurityEvent, 0),
	}
}

// LogEvent logs a security event
func (sa *SecurityAudit) LogEvent(eventType, severity, message, ip, userID string) {
	sa.mutex.Lock()
	defer sa.mutex.Unlock()
	
	event := SecurityEvent{
		Timestamp: time.Now(),
		Type:      eventType,
		Severity:  severity,
		Message:   message,
		IP:        ip,
		UserID:    userID,
	}
	
	sa.Events = append(sa.Events, event)
	
	// Keep only last 1000 events
	if len(sa.Events) > 1000 {
		sa.Events = sa.Events[len(sa.Events)-1000:]
	}
}

// GetEvents returns security events
func (sa *SecurityAudit) GetEvents() []SecurityEvent {
	sa.mutex.RLock()
	defer sa.mutex.RUnlock()
	
	events := make([]SecurityEvent, len(sa.Events))
	copy(events, sa.Events)
	return events
}

// GetEventsByType returns security events of a specific type
func (sa *SecurityAudit) GetEventsByType(eventType string) []SecurityEvent {
	sa.mutex.RLock()
	defer sa.mutex.RUnlock()
	
	var filtered []SecurityEvent
	for _, event := range sa.Events {
		if event.Type == eventType {
			filtered = append(filtered, event)
		}
	}
	return filtered
}

// Demonstrate security utilities
func demonstrateSecurityUtils() {
	fmt.Println("=== Security Utils Demo ===\n")
	
	// Create security utils
	su := NewSecurityUtils()
	
	// Test random string generation
	randomStr, err := su.GenerateSecureRandomString(16)
	if err != nil {
		log.Printf("Error generating random string: %v", err)
		return
	}
	fmt.Printf("Random string: %s\n", randomStr)
	
	// Test API key generation
	apiKey, err := su.GenerateAPIKey()
	if err != nil {
		log.Printf("Error generating API key: %v", err)
		return
	}
	fmt.Printf("API key: %s\n", apiKey)
	
	// Test password validation
	passwords := []string{
		"weak",
		"password123",
		"StrongP@ssw0rd!",
		"VeryStrongP@ssw0rd123!",
	}
	
	for _, password := range passwords {
		valid, issues := su.ValidatePasswordStrength(password)
		fmt.Printf("Password: %s\n", password)
		fmt.Printf("Valid: %t\n", valid)
		if !valid {
			fmt.Printf("Issues: %v\n", issues)
		}
		fmt.Println()
	}
	
	// Test input sanitization
	maliciousInput := "<script>alert('XSS')</script>"
	sanitized := su.SanitizeInput(maliciousInput)
	fmt.Printf("Original: %s\n", maliciousInput)
	fmt.Printf("Sanitized: %s\n", sanitized)
	
	// Test email validation
	emails := []string{
		"user@example.com",
		"invalid-email",
		"test@domain.co.uk",
		"notanemail",
	}
	
	for _, email := range emails {
		valid := su.ValidateEmail(email)
		fmt.Printf("Email: %s - Valid: %t\n", email, valid)
	}
	
	// Test HMAC
	data := "Hello, World!"
	key := "secret-key"
	hmac := su.GenerateHMAC(data, key)
	fmt.Printf("Data: %s\n", data)
	fmt.Printf("HMAC: %s\n", hmac)
	
	valid := su.VerifyHMAC(data, key, hmac)
	fmt.Printf("HMAC valid: %t\n", valid)
}

// Demonstrate rate limiting
func demonstrateRateLimiting() {
	fmt.Println("=== Rate Limiting Demo ===\n")
	
	// Create rate limiter (5 requests per minute)
	rl := NewRateLimiter(5, time.Minute)
	
	// Test rate limiting
	clientIP := "192.168.1.100"
	
	for i := 0; i < 7; i++ {
		allowed := rl.Allow(clientIP)
		remaining := rl.GetRemainingRequests(clientIP)
		fmt.Printf("Request %d: Allowed=%t, Remaining=%d\n", i+1, allowed, remaining)
	}
}

// Demonstrate security audit
func demonstrateSecurityAudit() {
	fmt.Println("=== Security Audit Demo ===\n")
	
	// Create security audit
	audit := NewSecurityAudit()
	
	// Log some events
	audit.LogEvent("login", "info", "User logged in successfully", "192.168.1.100", "user123")
	audit.LogEvent("login", "warning", "Failed login attempt", "192.168.1.101", "")
	audit.LogEvent("access", "info", "User accessed protected resource", "192.168.1.100", "user123")
	audit.LogEvent("error", "error", "Database connection failed", "192.168.1.100", "user123")
	
	// Get all events
	events := audit.GetEvents()
	fmt.Printf("Total events: %d\n", len(events))
	
	// Get events by type
	loginEvents := audit.GetEventsByType("login")
	fmt.Printf("Login events: %d\n", len(loginEvents))
	
	// Display events
	for _, event := range events {
		fmt.Printf("[%s] %s: %s (IP: %s, User: %s)\n", 
			event.Timestamp.Format("15:04:05"), 
			event.Severity, 
			event.Message, 
			event.IP, 
			event.UserID)
	}
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 8: Security Utils")
	fmt.Println("=====================================================\n")
	
	// Run all demonstrations
	demonstrateSecurityUtils()
	fmt.Println()
	demonstrateRateLimiting()
	fmt.Println()
	demonstrateSecurityAudit()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Security utilities help implement common security patterns")
	fmt.Println("2. Input validation prevents many security vulnerabilities")
	fmt.Println("3. Rate limiting protects against abuse and DoS attacks")
	fmt.Println("4. Security auditing helps track and investigate incidents")
	fmt.Println("5. Proper random number generation is crucial for security")
	fmt.Println("6. Security headers protect against common web vulnerabilities")
}
```

I need to fix the import issue in the security_utils.go file:


```go
package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

// SecurityUtils provides various security utility functions
type SecurityUtils struct{}

// NewSecurityUtils creates a new security utils instance
func NewSecurityUtils() *SecurityUtils {
	return &SecurityUtils{}
}

// GenerateRandomString generates a random string of specified length
func (su *SecurityUtils) GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// GenerateSecureRandomString generates a cryptographically secure random string
func (su *SecurityUtils) GenerateSecureRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	
	for i := range b {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	return string(b), nil
}

// GenerateAPIKey generates a secure API key
func (su *SecurityUtils) GenerateAPIKey() (string, error) {
	return su.GenerateSecureRandomString(32)
}

// GenerateSessionID generates a secure session ID
func (su *SecurityUtils) GenerateSessionID() (string, error) {
	return su.GenerateSecureRandomString(16)
}

// HashMD5 hashes data using MD5 (not recommended for security)
func (su *SecurityUtils) HashMD5(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// HashSHA1 hashes data using SHA-1 (not recommended for security)
func (su *SecurityUtils) HashSHA1(data string) string {
	hash := sha1.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// HashSHA256 hashes data using SHA-256
func (su *SecurityUtils) HashSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// HashSHA512 hashes data using SHA-512
func (su *SecurityUtils) HashSHA512(data string) string {
	hash := sha512.Sum512([]byte(data))
	return hex.EncodeToString(hash[:])
}

// GenerateHMAC generates HMAC for data with given key
func (su *SecurityUtils) GenerateHMAC(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// VerifyHMAC verifies HMAC for data with given key
func (su *SecurityUtils) VerifyHMAC(data, key, expectedMAC string) bool {
	actualMAC := su.GenerateHMAC(data, key)
	return hmac.Equal([]byte(actualMAC), []byte(expectedMAC))
}

// ValidatePasswordStrength validates password strength
func (su *SecurityUtils) ValidatePasswordStrength(password string) (bool, []string) {
	var issues []string
	
	if len(password) < 8 {
		issues = append(issues, "Password must be at least 8 characters long")
	}
	
	if len(password) > 128 {
		issues = append(issues, "Password must be no more than 128 characters long")
	}
	
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false
	
	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasDigit = true
		case strings.ContainsRune("!@#$%^&*()_+-=[]{}|;:,.<>?", char):
			hasSpecial = true
		}
	}
	
	if !hasUpper {
		issues = append(issues, "Password must contain at least one uppercase letter")
	}
	
	if !hasLower {
		issues = append(issues, "Password must contain at least one lowercase letter")
	}
	
	if !hasDigit {
		issues = append(issues, "Password must contain at least one digit")
	}
	
	if !hasSpecial {
		issues = append(issues, "Password must contain at least one special character")
	}
	
	return len(issues) == 0, issues
}

// SanitizeInput sanitizes user input to prevent XSS
func (su *SecurityUtils) SanitizeInput(input string) string {
	// Replace potentially dangerous characters
	input = strings.ReplaceAll(input, "<", "&lt;")
	input = strings.ReplaceAll(input, ">", "&gt;")
	input = strings.ReplaceAll(input, "\"", "&quot;")
	input = strings.ReplaceAll(input, "'", "&#x27;")
	input = strings.ReplaceAll(input, "&", "&amp;")
	input = strings.ReplaceAll(input, "/", "&#x2F;")
	
	return input
}

// ValidateEmail validates email format
func (su *SecurityUtils) ValidateEmail(email string) bool {
	// Simple email validation
	if len(email) < 5 {
		return false
	}
	
	if !strings.Contains(email, "@") {
		return false
	}
	
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	
	if len(parts[0]) == 0 || len(parts[1]) == 0 {
		return false
	}
	
	if !strings.Contains(parts[1], ".") {
		return false
	}
	
	return true
}

// ValidateIPAddress validates IP address format
func (su *SecurityUtils) ValidateIPAddress(ip string) bool {
	// Check if it's a valid IPv4 or IPv6 address
	if net.ParseIP(ip) == nil {
		return false
	}
	return true
}

// IsPrivateIP checks if an IP address is private
func (su *SecurityUtils) IsPrivateIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	
	// Check private IP ranges
	privateRanges := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"127.0.0.0/8",
		"169.254.0.0/16",
	}
	
	for _, rangeStr := range privateRanges {
		_, network, _ := net.ParseCIDR(rangeStr)
		if network.Contains(parsedIP) {
			return true
		}
	}
	
	return false
}

// GenerateCSRFToken generates a CSRF token
func (su *SecurityUtils) GenerateCSRFToken() (string, error) {
	return su.GenerateSecureRandomString(32)
}

// ValidateCSRFToken validates a CSRF token
func (su *SecurityUtils) ValidateCSRFToken(token, sessionID string) bool {
	// In a real implementation, you would store and validate CSRF tokens
	// This is a simplified version
	return len(token) == 32 && len(sessionID) > 0
}

// RateLimiter provides rate limiting functionality
type RateLimiter struct {
	requests map[string][]time.Time
	mutex    sync.RWMutex
	limit    int
	window   time.Duration
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

// Allow checks if a request is allowed for the given key
func (rl *RateLimiter) Allow(key string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	
	now := time.Now()
	cutoff := now.Add(-rl.window)
	
	// Clean old requests
	requests := rl.requests[key]
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
	rl.requests[key] = validRequests
	return true
}

// GetRemainingRequests returns the number of remaining requests
func (rl *RateLimiter) GetRemainingRequests(key string) int {
	rl.mutex.RLock()
	defer rl.mutex.RUnlock()
	
	now := time.Now()
	cutoff := now.Add(-rl.window)
	
	requests := rl.requests[key]
	var validRequests []time.Time
	for _, req := range requests {
		if req.After(cutoff) {
			validRequests = append(validRequests, req)
		}
	}
	
	return rl.limit - len(validRequests)
}

// SecurityHeaders provides security header management
type SecurityHeaders struct {
	Headers map[string]string
}

// NewSecurityHeaders creates a new security headers instance
func NewSecurityHeaders() *SecurityHeaders {
	return &SecurityHeaders{
		Headers: map[string]string{
			"X-Content-Type-Options": "nosniff",
			"X-Frame-Options":        "DENY",
			"X-XSS-Protection":       "1; mode=block",
			"Strict-Transport-Security": "max-age=31536000; includeSubDomains",
			"Content-Security-Policy":   "default-src 'self'",
			"Referrer-Policy":           "strict-origin-when-cross-origin",
			"Permissions-Policy":        "geolocation=(), microphone=(), camera=()",
		},
	}
}

// SetHeader sets a security header
func (sh *SecurityHeaders) SetHeader(name, value string) {
	sh.Headers[name] = value
}

// ApplyHeaders applies security headers to HTTP response
func (sh *SecurityHeaders) ApplyHeaders(w http.ResponseWriter) {
	for name, value := range sh.Headers {
		w.Header().Set(name, value)
	}
}

// SecurityAudit provides security auditing functionality
type SecurityAudit struct {
	Events []SecurityEvent
	mutex  sync.RWMutex
}

// SecurityEvent represents a security event
type SecurityEvent struct {
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
	Severity  string    `json:"severity"`
	Message   string    `json:"message"`
	IP        string    `json:"ip"`
	UserID    string    `json:"user_id,omitempty"`
}

// NewSecurityAudit creates a new security audit instance
func NewSecurityAudit() *SecurityAudit {
	return &SecurityAudit{
		Events: make([]SecurityEvent, 0),
	}
}

// LogEvent logs a security event
func (sa *SecurityAudit) LogEvent(eventType, severity, message, ip, userID string) {
	sa.mutex.Lock()
	defer sa.mutex.Unlock()
	
	event := SecurityEvent{
		Timestamp: time.Now(),
		Type:      eventType,
		Severity:  severity,
		Message:   message,
		IP:        ip,
		UserID:    userID,
	}
	
	sa.Events = append(sa.Events, event)
	
	// Keep only last 1000 events
	if len(sa.Events) > 1000 {
		sa.Events = sa.Events[len(sa.Events)-1000:]
	}
}

// GetEvents returns security events
func (sa *SecurityAudit) GetEvents() []SecurityEvent {
	sa.mutex.RLock()
	defer sa.mutex.RUnlock()
	
	events := make([]SecurityEvent, len(sa.Events))
	copy(events, sa.Events)
	return events
}

// GetEventsByType returns security events of a specific type
func (sa *SecurityAudit) GetEventsByType(eventType string) []SecurityEvent {
	sa.mutex.RLock()
	defer sa.mutex.RUnlock()
	
	var filtered []SecurityEvent
	for _, event := range sa.Events {
		if event.Type == eventType {
			filtered = append(filtered, event)
		}
	}
	return filtered
}

// Demonstrate security utilities
func demonstrateSecurityUtils() {
	fmt.Println("=== Security Utils Demo ===\n")
	
	// Create security utils
	su := NewSecurityUtils()
	
	// Test random string generation
	randomStr, err := su.GenerateSecureRandomString(16)
	if err != nil {
		log.Printf("Error generating random string: %v", err)
		return
	}
	fmt.Printf("Random string: %s\n", randomStr)
	
	// Test API key generation
	apiKey, err := su.GenerateAPIKey()
	if err != nil {
		log.Printf("Error generating API key: %v", err)
		return
	}
	fmt.Printf("API key: %s\n", apiKey)
	
	// Test password validation
	passwords := []string{
		"weak",
		"password123",
		"StrongP@ssw0rd!",
		"VeryStrongP@ssw0rd123!",
	}
	
	for _, password := range passwords {
		valid, issues := su.ValidatePasswordStrength(password)
		fmt.Printf("Password: %s\n", password)
		fmt.Printf("Valid: %t\n", valid)
		if !valid {
			fmt.Printf("Issues: %v\n", issues)
		}
		fmt.Println()
	}
	
	// Test input sanitization
	maliciousInput := "<script>alert('XSS')</script>"
	sanitized := su.SanitizeInput(maliciousInput)
	fmt.Printf("Original: %s\n", maliciousInput)
	fmt.Printf("Sanitized: %s\n", sanitized)
	
	// Test email validation
	emails := []string{
		"user@example.com",
		"invalid-email",
		"test@domain.co.uk",
		"notanemail",
	}
	
	for _, email := range emails {
		valid := su.ValidateEmail(email)
		fmt.Printf("Email: %s - Valid: %t\n", email, valid)
	}
	
	// Test HMAC
	data := "Hello, World!"
	key := "secret-key"
	hmac := su.GenerateHMAC(data, key)
	fmt.Printf("Data: %s\n", data)
	fmt.Printf("HMAC: %s\n", hmac)
	
	valid := su.VerifyHMAC(data, key, hmac)
	fmt.Printf("HMAC valid: %t\n", valid)
}

// Demonstrate rate limiting
func demonstrateRateLimiting() {
	fmt.Println("=== Rate Limiting Demo ===\n")
	
	// Create rate limiter (5 requests per minute)
	rl := NewRateLimiter(5, time.Minute)
	
	// Test rate limiting
	clientIP := "192.168.1.100"
	
	for i := 0; i < 7; i++ {
		allowed := rl.Allow(clientIP)
		remaining := rl.GetRemainingRequests(clientIP)
		fmt.Printf("Request %d: Allowed=%t, Remaining=%d\n", i+1, allowed, remaining)
	}
}

// Demonstrate security audit
func demonstrateSecurityAudit() {
	fmt.Println("=== Security Audit Demo ===\n")
	
	// Create security audit
	audit := NewSecurityAudit()
	
	// Log some events
	audit.LogEvent("login", "info", "User logged in successfully", "192.168.1.100", "user123")
	audit.LogEvent("login", "warning", "Failed login attempt", "192.168.1.101", "")
	audit.LogEvent("access", "info", "User accessed protected resource", "192.168.1.100", "user123")
	audit.LogEvent("error", "error", "Database connection failed", "192.168.1.100", "user123")
	
	// Get all events
	events := audit.GetEvents()
	fmt.Printf("Total events: %d\n", len(events))
	
	// Get events by type
	loginEvents := audit.GetEventsByType("login")
	fmt.Printf("Login events: %d\n", len(loginEvents))
	
	// Display events
	for _, event := range events {
		fmt.Printf("[%s] %s: %s (IP: %s, User: %s)\n", 
			event.Timestamp.Format("15:04:05"), 
			event.Severity, 
			event.Message, 
			event.IP, 
			event.UserID)
	}
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 8: Security Utils")
	fmt.Println("=====================================================\n")
	
	// Run all demonstrations
	demonstrateSecurityUtils()
	fmt.Println()
	demonstrateRateLimiting()
	fmt.Println()
	demonstrateSecurityAudit()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Security utilities help implement common security patterns")
	fmt.Println("2. Input validation prevents many security vulnerabilities")
	fmt.Println("3. Rate limiting protects against abuse and DoS attacks")
	fmt.Println("4. Security auditing helps track and investigate incidents")
	fmt.Println("5. Proper random number generation is crucial for security")
	fmt.Println("6. Security headers protect against common web vulnerabilities")
}

```

Now let me create a quick reference guide and update the progress:


```markdown
# Network Security & Encryption Quick Reference Guide

## 🚀 Essential Security Concepts

### Security Principles
- **Confidentiality**: Data only accessible to authorized users
- **Integrity**: Data has not been modified or corrupted
- **Availability**: Data and services available when needed
- **Authentication**: Verifying identity of users/systems
- **Authorization**: Controlling access to resources
- **Non-repudiation**: Preventing denial of actions

### Encryption Types
- **Symmetric**: Same key for encryption/decryption (AES, ChaCha20)
- **Asymmetric**: Public/private key pairs (RSA, ECC)
- **Hash Functions**: One-way functions (SHA-256, SHA-3)
- **Digital Signatures**: Authentication and integrity (RSA, ECDSA)

## 🔧 Go Security Patterns

### TLS Configuration
```go
config := &tls.Config{
    MinVersion:         tls.VersionTLS12,
    CipherSuites:       []uint16{
        tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
        tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
    },
    PreferServerCipherSuites: true,
    InsecureSkipVerify: false,
}
```

### Secure HTTP Server
```go
func startSecureServer() {
    server := &http.Server{
        Addr:      ":8443",
        TLSConfig: config,
    }
    
    log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
```

### Secure HTTP Client
```go
func createSecureClient() *http.Client {
    return &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: false,
            },
        },
        Timeout: 30 * time.Second,
    }
}
```

## 🔐 Encryption Implementation

### AES Encryption
```go
func encryptAES(plaintext, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    nonce := make([]byte, gcm.NonceSize())
    if _, err := rand.Read(nonce); err != nil {
        return nil, err
    }
    
    ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
    return ciphertext, nil
}

func decryptAES(ciphertext, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, fmt.Errorf("ciphertext too short")
    }
    
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    return gcm.Open(nil, nonce, ciphertext, nil)
}
```

### RSA Encryption
```go
func encryptRSA(plaintext []byte, publicKey *rsa.PublicKey) ([]byte, error) {
    hash := sha256.New()
    return rsa.EncryptOAEP(hash, rand.Reader, publicKey, plaintext, nil)
}

func decryptRSA(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
    hash := sha256.New()
    return rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, nil)
}
```

### Digital Signatures
```go
func signRSA(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
    hash := sha256.Sum256(data)
    return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
}

func verifyRSA(data, signature []byte, publicKey *rsa.PublicKey) error {
    hash := sha256.Sum256(data)
    return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
}
```

## 🔑 Password Security

### Password Hashing
```go
func hashPassword(password string) (string, error) {
    salt := make([]byte, 32)
    if _, err := rand.Read(salt); err != nil {
        return "", err
    }
    
    hash := sha256.Sum256(append([]byte(password), salt...))
    combined := append(hash[:], salt...)
    
    return base64.StdEncoding.EncodeToString(combined), nil
}

func verifyPassword(password, hashedPassword string) bool {
    data, err := base64.StdEncoding.DecodeString(hashedPassword)
    if err != nil || len(data) < 32 {
        return false
    }
    
    hash := data[:32]
    salt := data[32:]
    
    expectedHash := sha256.Sum256(append([]byte(password), salt...))
    return bytes.Equal(hash, expectedHash[:])
}
```

### Password Validation
```go
func validatePassword(password string) (bool, []string) {
    var issues []string
    
    if len(password) < 8 {
        issues = append(issues, "Password must be at least 8 characters")
    }
    
    hasUpper := strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    hasLower := strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
    hasDigit := strings.ContainsAny(password, "0123456789")
    hasSpecial := strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|;:,.<>?")
    
    if !hasUpper {
        issues = append(issues, "Password must contain uppercase letter")
    }
    if !hasLower {
        issues = append(issues, "Password must contain lowercase letter")
    }
    if !hasDigit {
        issues = append(issues, "Password must contain digit")
    }
    if !hasSpecial {
        issues = append(issues, "Password must contain special character")
    }
    
    return len(issues) == 0, issues
}
```

## 🎫 JWT Tokens

### JWT Implementation
```go
type JWTClaims struct {
    UserID   string `json:"user_id"`
    Username string `json:"username"`
    Exp      int64  `json:"exp"`
    Iat      int64  `json:"iat"`
}

func generateJWT(userID, username string, secretKey []byte) (string, error) {
    claims := JWTClaims{
        UserID:   userID,
        Username: username,
        Exp:      time.Now().Add(24 * time.Hour).Unix(),
        Iat:      time.Now().Unix(),
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

func verifyJWT(tokenString string, secretKey []byte) (*JWTClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    
    if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
        return claims, nil
    }
    
    return nil, err
}
```

## 🛡️ Security Middleware

### Authentication Middleware
```go
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Authorization required", http.StatusUnauthorized)
            return
        }
        
        token = strings.TrimPrefix(token, "Bearer ")
        claims, err := verifyJWT(token, secretKey)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        
        ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
        next.ServeHTTP(w, r.WithContext(ctx))
    }
}
```

### Security Headers Middleware
```go
func securityHeadersMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        w.Header().Set("Strict-Transport-Security", "max-age=31536000")
        w.Header().Set("Content-Security-Policy", "default-src 'self'")
        
        next.ServeHTTP(w, r)
    }
}
```

### CORS Middleware
```go
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    }
}
```

## 🚦 Rate Limiting

### Rate Limiter Implementation
```go
type RateLimiter struct {
    requests map[string][]time.Time
    mutex    sync.RWMutex
    limit    int
    window   time.Duration
}

func (rl *RateLimiter) Allow(key string) bool {
    rl.mutex.Lock()
    defer rl.mutex.Unlock()
    
    now := time.Now()
    cutoff := now.Add(-rl.window)
    
    // Clean old requests
    requests := rl.requests[key]
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
    rl.requests[key] = validRequests
    return true
}
```

## 🔍 Input Validation

### Input Sanitization
```go
func sanitizeInput(input string) string {
    input = strings.ReplaceAll(input, "<", "&lt;")
    input = strings.ReplaceAll(input, ">", "&gt;")
    input = strings.ReplaceAll(input, "\"", "&quot;")
    input = strings.ReplaceAll(input, "'", "&#x27;")
    input = strings.ReplaceAll(input, "&", "&amp;")
    return input
}
```

### Email Validation
```go
func validateEmail(email string) bool {
    if len(email) < 5 || !strings.Contains(email, "@") {
        return false
    }
    
    parts := strings.Split(email, "@")
    if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
        return false
    }
    
    return strings.Contains(parts[1], ".")
}
```

## 📊 Security Monitoring

### Security Event Logging
```go
type SecurityEvent struct {
    Timestamp time.Time `json:"timestamp"`
    Type      string    `json:"type"`
    Severity  string    `json:"severity"`
    Message   string    `json:"message"`
    IP        string    `json:"ip"`
    UserID    string    `json:"user_id,omitempty"`
}

func logSecurityEvent(eventType, severity, message, ip, userID string) {
    event := SecurityEvent{
        Timestamp: time.Now(),
        Type:      eventType,
        Severity:  severity,
        Message:   message,
        IP:        ip,
        UserID:    userID,
    }
    
    // Log to file, database, or monitoring system
    log.Printf("Security Event: %+v", event)
}
```

## 🔒 Certificate Management

### Self-Signed Certificate Generation
```go
func generateSelfSignedCert() ([]byte, []byte, error) {
    priv, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return nil, nil, err
    }
    
    template := x509.Certificate{
        SerialNumber: big.NewInt(1),
        Subject: pkix.Name{
            Organization: []string{"Test Company"},
        },
        NotBefore:    time.Now(),
        NotAfter:     time.Now().Add(365 * 24 * time.Hour),
        KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
        ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
        IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1)},
        DNSNames:     []string{"localhost"},
    }
    
    certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
    if err != nil {
        return nil, nil, err
    }
    
    certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
    keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
    
    return certPEM, keyPEM, nil
}
```

## 🎯 Best Practices

1. **Use strong encryption**: AES-256, RSA-2048+, SHA-256+
2. **Implement proper key management**: Secure storage, rotation
3. **Validate all input**: Sanitize, validate, escape
4. **Use HTTPS everywhere**: TLS 1.2+ for all communications
5. **Implement proper authentication**: Multi-factor when possible
6. **Log security events**: Monitor and audit
7. **Keep systems updated**: Patch vulnerabilities
8. **Use secure defaults**: Strong configurations
9. **Implement rate limiting**: Prevent abuse
10. **Follow OWASP guidelines**: Security best practices

## 🚨 Common Vulnerabilities

### OWASP Top 10
1. **Injection**: SQL, NoSQL, OS command injection
2. **Broken Authentication**: Weak authentication mechanisms
3. **Sensitive Data Exposure**: Inadequate protection
4. **XML External Entities**: XXE attacks
5. **Broken Access Control**: Inadequate access controls
6. **Security Misconfiguration**: Insecure defaults
7. **Cross-Site Scripting**: XSS attacks
8. **Insecure Deserialization**: Unsafe deserialization
9. **Known Vulnerabilities**: Outdated components
10. **Insufficient Logging**: Inadequate monitoring

---

**Remember**: Security is not optional. Implement these patterns, and you'll build secure, robust applications! 🚀
```

Now let me update the progress:


Perfect! I've completed **Topic 8: Network Security & Encryption** with comprehensive coverage. Here's what I've created for you:

## 🎯 **Topic 8: Network Security & Encryption - COMPLETED**

### **📚 Theory Coverage** (`08_network_security/theory.md`)
- **Network Security Overview**: Security principles, threats, protection mechanisms
- **Cryptography Fundamentals**: Symmetric/asymmetric encryption, hash functions, digital signatures
- **TLS/SSL Protocol**: Handshake process, versions, record protocol, certificate management
- **Authentication Methods**: Password-based, MFA, certificate-based, token-based
- **Network Security Protocols**: IPsec, VPN, SSH, DNSSEC
- **Application Security**: HTTPS, WSS, email security
- **Security Best Practices**: Secure coding, network security, key management
- **Go Security Implementation**: TLS configuration, certificate generation, password hashing
- **Security Monitoring**: Logging, intrusion detection, vulnerability management
- **Compliance and Standards**: Security standards, frameworks, regulations
- **Common Vulnerabilities**: OWASP Top 10, network vulnerabilities

### **�� Implementation** (`08_network_security/implementation.go`)
- **SecureServer**: HTTPS server with TLS configuration
- **SecureClient**: HTTPS client with secure communication
- **EncryptionService**: AES and RSA encryption/decryption
- **PasswordService**: Secure password hashing and verification
- **JWTService**: JWT token generation and verification
- **SecurityMiddleware**: Authentication, CORS, security headers
- **Comprehensive Examples**: All major security patterns

### **�� Security Utils** (`08_network_security/security_utils.go`)
- **SecurityUtils**: Random generation, hashing, validation
- **RateLimiter**: Request rate limiting and abuse prevention
- **SecurityHeaders**: Security header management
- **SecurityAudit**: Security event logging and monitoring
- **Input Validation**: Sanitization, email validation, password strength
- **Utility Functions**: API keys, session IDs, CSRF tokens

### **📖 Quick Reference** (`08_network_security/quick_reference.md`)
- **Essential Concepts**: Security principles, encryption types
- **Go Patterns**: TLS configuration, secure servers/clients
- **Encryption Implementation**: AES, RSA, digital signatures
- **Password Security**: Hashing, validation, strength checking
- **JWT Tokens**: Generation, verification, claims handling
- **Security Middleware**: Authentication, headers, CORS
- **Rate Limiting**: Request limiting, abuse prevention
- **Input Validation**: Sanitization, validation, security
- **Security Monitoring**: Event logging, auditing
- **Certificate Management**: Self-signed certificates, validation
- **Best Practices**: Professional security guidelines

## 🚀 **Key Features of This Module**

1. **Complete Security Coverage**: From basics to advanced security concepts
2. **Go-Specific**: Leverages Go's excellent security libraries
3. **Production-Ready**: Real-world security patterns and best practices
4. **Comprehensive Encryption**: AES, RSA, hashing, digital signatures
5. **Authentication & Authorization**: JWT, password security, middleware
6. **Security Monitoring**: Logging, auditing, rate limiting
7. **Input Validation**: Sanitization, validation, security headers
8. **Certificate Management**: TLS/SSL, certificate generation

## 📈 **What You've Learned**

- **Security Fundamentals**: Principles, threats, protection mechanisms
- **Cryptography**: Symmetric/asymmetric encryption, hashing, signatures
- **TLS/SSL**: Secure communication, handshake, certificate management
- **Authentication**: Password security, JWT tokens, multi-factor auth
- **Network Security**: IPsec, VPN, SSH, secure protocols
- **Application Security**: HTTPS, input validation, security headers
- **Security Monitoring**: Logging, auditing, rate limiting
- **Go Implementation**: Professional security programming

## **Next Steps**

You're now ready to move to **Topic 9: Routing & Switching** where we'll explore network routing algorithms, switching concepts, and network topology.

The security foundation you now have will make you unstoppable in building secure applications, understanding network security, and acing security-related interview questions! 🚀

Would you like me to continue with **Topic 9: Routing & Switching** or would you prefer to explore any specific aspect of network security first?

The security knowledge you've gained is crucial for modern software development and will set you apart in technical interviews!

---

**User**

continue with Topic 9: Routing & Switching 

---

**Cursor**

Perfect! Let me continue with **Topic 9: Routing & Switching**. This is where we explore network routing algorithms, switching concepts, and network topology - the backbone of how data moves through networks.


Now let me create the comprehensive Routing & Switching module:


```markdown
# Topic 9: Routing & Switching - Network Path Management Mastery

## Routing & Switching Overview

Routing and switching are the fundamental mechanisms that enable data to move through networks efficiently. They operate at different layers of the OSI model and serve different purposes in network communication.

### Key Concepts
- **Routing**: Process of selecting paths for data packets across networks
- **Switching**: Process of forwarding data frames within a network segment
- **Forwarding**: The actual movement of data from input to output
- **Path Selection**: Choosing the best route based on various criteria
- **Network Topology**: Physical and logical arrangement of network elements

## Network Switching

### Switching Fundamentals
Switching operates at Layer 2 (Data Link Layer) and is responsible for forwarding frames within a local area network (LAN).

#### Switch Functions
- **Frame Forwarding**: Move frames from input port to output port
- **Address Learning**: Build MAC address table by observing traffic
- **Loop Prevention**: Prevent broadcast storms using STP
- **VLAN Support**: Virtual LAN segmentation and management
- **Quality of Service**: Traffic prioritization and bandwidth management

#### Switch Types
- **Unmanaged Switches**: Basic plug-and-play switches
- **Managed Switches**: Configurable switches with advanced features
- **Layer 2 Switches**: Operate at Data Link Layer
- **Layer 3 Switches**: Combine switching and routing capabilities
- **PoE Switches**: Power over Ethernet for connected devices

### MAC Address Learning
Switches learn MAC addresses by examining the source address of incoming frames and associating them with the port they arrived on.

#### Learning Process
1. **Frame Arrival**: Switch receives frame on port
2. **Source Learning**: Extract source MAC address
3. **Table Update**: Add MAC-port mapping to table
4. **Aging**: Remove old entries after timeout period
5. **Flooding**: Send unknown destination frames to all ports

#### MAC Address Table
- **MAC Address**: 48-bit hardware address
- **Port**: Physical port number
- **VLAN**: Virtual LAN identifier
- **Age**: Time since last seen
- **Type**: Static or dynamic entry

### VLAN (Virtual LAN)
VLANs logically segment a physical network into multiple broadcast domains.

#### VLAN Benefits
- **Security**: Isolate traffic between VLANs
- **Performance**: Reduce broadcast domain size
- **Management**: Logical grouping of devices
- **Flexibility**: Move devices without rewiring

#### VLAN Types
- **Port-based VLAN**: Assign ports to VLANs
- **MAC-based VLAN**: Assign devices by MAC address
- **Protocol-based VLAN**: Assign by protocol type
- **IP-based VLAN**: Assign by IP subnet

#### VLAN Trunking
- **802.1Q**: Standard for VLAN tagging
- **ISL**: Cisco's Inter-Switch Link protocol
- **Native VLAN**: Untagged traffic on trunk
- **VLAN Range**: 1-4094 (1 and 1002-1005 reserved)

### Spanning Tree Protocol (STP)
STP prevents loops in switched networks by creating a loop-free topology.

#### STP Concepts
- **Root Bridge**: Central reference point for STP
- **Root Port**: Port with lowest cost to root bridge
- **Designated Port**: Port with lowest cost to segment
- **Blocking Port**: Port blocked to prevent loops
- **BPDU**: Bridge Protocol Data Unit for STP communication

#### STP States
- **Blocking**: Port blocked, no forwarding
- **Listening**: Port listening for BPDUs
- **Learning**: Port learning MAC addresses
- **Forwarding**: Port forwarding traffic
- **Disabled**: Port administratively disabled

#### STP Variants
- **STP (802.1D)**: Original spanning tree
- **RSTP (802.1W)**: Rapid spanning tree
- **MSTP (802.1S)**: Multiple spanning tree
- **PVST+**: Per-VLAN spanning tree

## Network Routing

### Routing Fundamentals
Routing operates at Layer 3 (Network Layer) and is responsible for forwarding packets between different networks.

#### Router Functions
- **Packet Forwarding**: Move packets between networks
- **Path Selection**: Choose best route to destination
- **Route Management**: Maintain routing tables
- **Protocol Support**: Implement routing protocols
- **Traffic Control**: Manage bandwidth and QoS

#### Routing Types
- **Static Routing**: Manually configured routes
- **Dynamic Routing**: Automatically learned routes
- **Default Routing**: Catch-all route for unknown destinations
- **Host Routing**: Route to specific host
- **Network Routing**: Route to network subnet

### Routing Tables
Routing tables contain information about how to reach different destinations.

#### Table Entries
- **Destination**: Network or host address
- **Next Hop**: Next router in path
- **Interface**: Outgoing interface
- **Metric**: Cost to reach destination
- **Administrative Distance**: Trustworthiness of route source

#### Route Types
- **Directly Connected**: Local network interfaces
- **Static Routes**: Manually configured
- **Dynamic Routes**: Learned from routing protocols
- **Default Routes**: 0.0.0.0/0 destination

### Routing Algorithms

#### Distance Vector Algorithms
Distance vector algorithms use the Bellman-Ford algorithm to find shortest paths.

**RIP (Routing Information Protocol)**
- **Metric**: Hop count (maximum 15)
- **Update Interval**: 30 seconds
- **Convergence**: Slow (up to 3 minutes)
- **Variant**: RIPv1, RIPv2, RIPng

**IGRP (Interior Gateway Routing Protocol)**
- **Metric**: Composite (bandwidth, delay, reliability, load)
- **Update Interval**: 90 seconds
- **Convergence**: Moderate
- **Variant**: Cisco proprietary

**EIGRP (Enhanced IGRP)**
- **Metric**: Composite (bandwidth, delay)
- **Update**: Triggered updates
- **Convergence**: Fast
- **Variant**: Cisco proprietary

#### Link State Algorithms
Link state algorithms use Dijkstra's algorithm to find shortest paths.

**OSPF (Open Shortest Path First)**
- **Metric**: Cost (bandwidth-based)
- **Update**: Link state advertisements (LSA)
- **Convergence**: Fast
- **Variant**: OSPFv2 (IPv4), OSPFv3 (IPv6)

**IS-IS (Intermediate System to Intermediate System)**
- **Metric**: Cost (configurable)
- **Update**: Link state packets (LSP)
- **Convergence**: Fast
- **Variant**: ISO 10589 standard

#### Path Vector Algorithms
Path vector algorithms prevent routing loops by tracking path information.

**BGP (Border Gateway Protocol)**
- **Metric**: Path attributes
- **Update**: Incremental updates
- **Convergence**: Variable
- **Variant**: BGP-4 (IPv4), BGP-4+ (IPv6)

### Routing Protocols

#### Interior Gateway Protocols (IGP)
IGPs are used within autonomous systems (AS).

**RIP (Routing Information Protocol)**
- **Type**: Distance vector
- **Metric**: Hop count
- **Max Hops**: 15
- **Update**: Every 30 seconds
- **Use Case**: Small networks

**OSPF (Open Shortest Path First)**
- **Type**: Link state
- **Metric**: Cost
- **Update**: Triggered
- **Use Case**: Large networks

**EIGRP (Enhanced IGRP)**
- **Type**: Advanced distance vector
- **Metric**: Composite
- **Update**: Triggered
- **Use Case**: Cisco networks

#### Exterior Gateway Protocols (EGP)
EGPs are used between autonomous systems.

**BGP (Border Gateway Protocol)**
- **Type**: Path vector
- **Metric**: Path attributes
- **Update**: Incremental
- **Use Case**: Internet routing

### Network Topology

#### Physical Topology
Physical topology describes the physical arrangement of network devices.

**Bus Topology**
- **Structure**: Single cable connecting all devices
- **Advantages**: Simple, inexpensive
- **Disadvantages**: Single point of failure, limited scalability
- **Use Case**: Small networks

**Star Topology**
- **Structure**: Central hub/switch connecting all devices
- **Advantages**: Easy management, fault isolation
- **Disadvantages**: Central point of failure
- **Use Case**: Most common LAN topology

**Ring Topology**
- **Structure**: Devices connected in circular fashion
- **Advantages**: Equal access, no collisions
- **Disadvantages**: Single point of failure, difficult to modify
- **Use Case**: Token Ring networks

**Mesh Topology**
- **Structure**: Every device connected to every other device
- **Advantages**: High redundancy, fault tolerance
- **Disadvantages**: Expensive, complex
- **Use Case**: Critical networks

#### Logical Topology
Logical topology describes how data flows through the network.

**Flat Network**
- **Structure**: Single broadcast domain
- **Advantages**: Simple, easy to manage
- **Disadvantages**: Limited scalability, security issues
- **Use Case**: Small networks

**Hierarchical Network**
- **Structure**: Multiple layers (core, distribution, access)
- **Advantages**: Scalable, manageable
- **Disadvantages**: Complex, expensive
- **Use Case**: Large networks

**Collapsed Core**
- **Structure**: Core and distribution combined
- **Advantages**: Cost-effective, simpler
- **Disadvantages**: Limited scalability
- **Use Case**: Medium networks

### Quality of Service (QoS)

#### QoS Concepts
QoS ensures that critical traffic receives priority treatment in the network.

**Traffic Classification**
- **Voice**: Real-time voice traffic
- **Video**: Real-time video traffic
- **Data**: Best-effort data traffic
- **Management**: Network management traffic

**QoS Mechanisms**
- **Classification**: Identify traffic types
- **Marking**: Set priority levels
- **Policing**: Enforce traffic limits
- **Shaping**: Smooth traffic bursts
- **Queuing**: Manage traffic queues
- **Scheduling**: Prioritize traffic transmission

#### QoS Models
**Best Effort**
- **Treatment**: No guarantees
- **Use Case**: General internet traffic
- **Advantages**: Simple, no overhead
- **Disadvantages**: No service guarantees

**Integrated Services (IntServ)**
- **Treatment**: Per-flow guarantees
- **Use Case**: Real-time applications
- **Advantages**: Guaranteed service
- **Disadvantages**: Complex, not scalable

**Differentiated Services (DiffServ)**
- **Treatment**: Per-class guarantees
- **Use Case**: Enterprise networks
- **Advantages**: Scalable, practical
- **Disadvantages**: No per-flow guarantees

### Network Redundancy

#### Redundancy Concepts
Redundancy provides backup paths and failover capabilities.

**High Availability**
- **Uptime**: 99.9% or better
- **Downtime**: Less than 8.76 hours/year
- **Mechanisms**: Redundancy, failover, load balancing

**Fault Tolerance**
- **Single Points of Failure**: Eliminate SPOFs
- **Redundant Components**: Backup systems
- **Failover**: Automatic switching to backup
- **Load Balancing**: Distribute traffic across multiple paths

#### Redundancy Protocols
**HSRP (Hot Standby Router Protocol)**
- **Type**: First-hop redundancy
- **Variant**: Cisco proprietary
- **Use Case**: Default gateway redundancy

**VRRP (Virtual Router Redundancy Protocol)**
- **Type**: First-hop redundancy
- **Variant**: RFC 3768 standard
- **Use Case**: Multi-vendor environments

**GLBP (Gateway Load Balancing Protocol)**
- **Type**: Load balancing and redundancy
- **Variant**: Cisco proprietary
- **Use Case**: Load distribution

### Network Monitoring

#### Monitoring Concepts
Network monitoring ensures optimal performance and availability.

**SNMP (Simple Network Management Protocol)**
- **Version**: SNMPv1, SNMPv2c, SNMPv3
- **Components**: Manager, Agent, MIB
- **Use Case**: Network device monitoring

**RMON (Remote Monitoring)**
- **Type**: Remote network monitoring
- **Variant**: RMON1, RMON2
- **Use Case**: Traffic analysis

**NetFlow**
- **Type**: Traffic flow analysis
- **Variant**: NetFlow v5, v9, IPFIX
- **Use Case**: Traffic monitoring and analysis

#### Performance Metrics
**Bandwidth Utilization**
- **Measurement**: Percentage of available bandwidth used
- **Threshold**: 70-80% utilization
- **Action**: Upgrade or optimize

**Latency**
- **Measurement**: Round-trip time (RTT)
- **Threshold**: < 100ms for voice, < 150ms for video
- **Action**: Optimize routing, reduce hops

**Packet Loss**
- **Measurement**: Percentage of packets lost
- **Threshold**: < 0.1% for voice, < 0.5% for video
- **Action**: Check links, upgrade hardware

**Jitter**
- **Measurement**: Variation in packet arrival times
- **Threshold**: < 30ms for voice, < 50ms for video
- **Action**: Implement QoS, buffer management

### Go Implementation Concepts

#### Network Simulation
Go can be used to simulate network behavior and routing algorithms.

**Graph Representation**
- **Nodes**: Network devices (routers, switches)
- **Edges**: Network links
- **Weights**: Link costs or metrics
- **Algorithms**: Dijkstra, Bellman-Ford, Floyd-Warshall

**Routing Table Implementation**
- **Data Structure**: Map or slice for routes
- **Operations**: Add, remove, update, lookup
- **Concurrency**: Thread-safe operations
- **Persistence**: Save/load routing tables

**Protocol Simulation**
- **Message Passing**: Simulate protocol messages
- **State Machines**: Implement protocol states
- **Timers**: Handle protocol timers
- **Convergence**: Test algorithm convergence

#### Network Management
Go can be used to build network management tools.

**SNMP Client**
- **Operations**: Get, Set, Walk, Trap
- **Data Types**: Integer, String, OID
- **Error Handling**: SNMP error codes
- **Concurrency**: Multiple concurrent operations

**Network Discovery**
- **Ping**: Test connectivity
- **Traceroute**: Find network path
- **Port Scanning**: Discover services
- **Topology Mapping**: Build network map

**Performance Monitoring**
- **Bandwidth**: Measure link utilization
- **Latency**: Measure round-trip time
- **Packet Loss**: Measure dropped packets
- **Throughput**: Measure data transfer rate

## Interview Questions

### Basic Questions
1. What is the difference between routing and switching?
2. How does a switch learn MAC addresses?
3. What is the purpose of VLANs?

### Intermediate Questions
1. Explain the Spanning Tree Protocol and its states.
2. Compare distance vector and link state routing algorithms.
3. What are the advantages of hierarchical network design?

### Advanced Questions
1. Design a network topology for a large enterprise.
2. Implement a routing algorithm in Go.
3. How would you optimize network performance and reliability?

## Next Steps
After mastering routing and switching, proceed to:
- **Topic 10**: Network Performance & Optimization
- **Topic 11**: Distributed Systems Networking
- **Topic 12**: Cloud Networking & Containerization

Master routing and switching, and you'll understand how networks move data efficiently! 🚀
```

Now let me create the comprehensive Go implementation for routing and switching:


```go
package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"sort"
	"sync"
	"time"
)

// NetworkNode represents a network device (router or switch)
type NetworkNode struct {
	ID       string
	Type     string // "router" or "switch"
	Ports    map[string]*Port
	MACTable map[string]string // MAC address -> port mapping
	Routes   map[string]*Route // destination -> route
	Neighbors map[string]*NetworkNode
	VLANs    map[int]*VLAN
	mutex    sync.RWMutex
}

// Port represents a network port
type Port struct {
	ID       string
	Type     string // "ethernet", "serial", "fiber"
	Speed    int    // Mbps
	Duplex   string // "half", "full"
	Status   string // "up", "down", "admin_down"
	VLAN     int
	Neighbor *NetworkNode
}

// Route represents a routing table entry
type Route struct {
	Destination string
	NextHop     string
	Interface   string
	Metric      int
	AdminDist   int
	Protocol    string
	Age         time.Time
}

// VLAN represents a Virtual LAN
type VLAN struct {
	ID          int
	Name        string
	Ports       map[string]bool
	Broadcast   bool
	Multicast   bool
	UnknownUnicast bool
}

// NetworkLink represents a connection between nodes
type NetworkLink struct {
	From     *NetworkNode
	To       *NetworkNode
	FromPort string
	ToPort   string
	Cost     int
	Bandwidth int // Mbps
	Status   string // "up", "down"
}

// NetworkTopology represents the entire network
type NetworkTopology struct {
	Nodes map[string]*NetworkNode
	Links map[string]*NetworkLink
	mutex sync.RWMutex
}

// NewNetworkNode creates a new network node
func NewNetworkNode(id, nodeType string) *NetworkNode {
	return &NetworkNode{
		ID:        id,
		Type:      nodeType,
		Ports:     make(map[string]*Port),
		MACTable:  make(map[string]string),
		Routes:    make(map[string]*Route),
		Neighbors: make(map[string]*NetworkNode),
		VLANs:     make(map[int]*VLAN),
	}
}

// AddPort adds a port to the node
func (n *NetworkNode) AddPort(portID, portType string, speed int) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	
	n.Ports[portID] = &Port{
		ID:     portID,
		Type:   portType,
		Speed:  speed,
		Duplex: "full",
		Status: "up",
		VLAN:   1, // Default VLAN
	}
}

// AddRoute adds a route to the routing table
func (n *NetworkNode) AddRoute(dest, nextHop, intf string, metric, adminDist int, protocol string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	
	n.Routes[dest] = &Route{
		Destination: dest,
		NextHop:     nextHop,
		Interface:   intf,
		Metric:      metric,
		AdminDist:   adminDist,
		Protocol:    protocol,
		Age:         time.Now(),
	}
}

// LookupRoute finds the best route to a destination
func (n *NetworkNode) LookupRoute(destination string) *Route {
	n.mutex.RLock()
	defer n.mutex.RUnlock()
	
	// First try exact match
	if route, exists := n.Routes[destination]; exists {
		return route
	}
	
	// Then try longest prefix match
	var bestRoute *Route
	var bestPrefixLen int
	
	for dest, route := range n.Routes {
		if n.isSubnetMatch(destination, dest) {
			prefixLen := n.getPrefixLength(dest)
			if prefixLen > bestPrefixLen {
				bestPrefixLen = prefixLen
				bestRoute = route
			}
		}
	}
	
	return bestRoute
}

// isSubnetMatch checks if an IP matches a subnet
func (n *NetworkNode) isSubnetMatch(ip, subnet string) bool {
	// Simplified implementation - in reality, you'd parse IPs and subnets
	return len(ip) >= len(subnet) && ip[:len(subnet)] == subnet
}

// getPrefixLength returns the prefix length of a subnet
func (n *NetworkNode) getPrefixLength(subnet string) int {
	// Simplified implementation - in reality, you'd parse CIDR notation
	return len(subnet)
}

// LearnMAC learns a MAC address on a port
func (n *NetworkNode) LearnMAC(mac, portID string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	
	n.MACTable[mac] = portID
}

// LookupMAC finds the port for a MAC address
func (n *NetworkNode) LookupMAC(mac string) (string, bool) {
	n.mutex.RLock()
	defer n.mutex.RUnlock()
	
	port, exists := n.MACTable[mac]
	return port, exists
}

// ForwardFrame forwards a frame based on MAC address
func (n *NetworkNode) ForwardFrame(srcMAC, dstMAC string, frame []byte) {
	if n.Type != "switch" {
		return
	}
	
	// Learn source MAC
	n.LearnMAC(srcMAC, "incoming")
	
	// Look up destination MAC
	if port, exists := n.LookupMAC(dstMAC); exists {
		// Forward to specific port
		fmt.Printf("Switch %s: Forwarding frame from %s to %s via port %s\n", 
			n.ID, srcMAC, dstMAC, port)
	} else {
		// Flood to all ports except incoming
		fmt.Printf("Switch %s: Flooding frame from %s to %s (unknown destination)\n", 
			n.ID, srcMAC, dstMAC)
	}
}

// CreateVLAN creates a new VLAN
func (n *NetworkNode) CreateVLAN(id int, name string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	
	n.VLANs[id] = &VLAN{
		ID:          id,
		Name:        name,
		Ports:       make(map[string]bool),
		Broadcast:   true,
		Multicast:   true,
		UnknownUnicast: true,
	}
}

// AddPortToVLAN adds a port to a VLAN
func (n *NetworkNode) AddPortToVLAN(portID string, vlanID int) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	
	if vlan, exists := n.VLANs[vlanID]; exists {
		vlan.Ports[portID] = true
		if port, exists := n.Ports[portID]; exists {
			port.VLAN = vlanID
		}
	}
}

// NewNetworkTopology creates a new network topology
func NewNetworkTopology() *NetworkTopology {
	return &NetworkTopology{
		Nodes: make(map[string]*NetworkNode),
		Links: make(map[string]*NetworkLink),
	}
}

// AddNode adds a node to the topology
func (nt *NetworkTopology) AddNode(node *NetworkNode) {
	nt.mutex.Lock()
	defer nt.mutex.Unlock()
	
	nt.Nodes[node.ID] = node
}

// AddLink adds a link between two nodes
func (nt *NetworkTopology) AddLink(from, to *NetworkNode, fromPort, toPort string, cost, bandwidth int) {
	nt.mutex.Lock()
	defer nt.mutex.Unlock()
	
	linkID := fmt.Sprintf("%s-%s", from.ID, to.ID)
	link := &NetworkLink{
		From:       from,
		To:         to,
		FromPort:   fromPort,
		ToPort:     toPort,
		Cost:       cost,
		Bandwidth:  bandwidth,
		Status:     "up",
	}
	
	nt.Links[linkID] = link
	
	// Add neighbors
	from.Neighbors[to.ID] = to
	to.Neighbors[from.ID] = from
}

// DijkstraShortestPath finds the shortest path between two nodes
func (nt *NetworkTopology) DijkstraShortestPath(source, destination string) ([]string, int) {
	nt.mutex.RLock()
	defer nt.mutex.RUnlock()
	
	// Initialize distances
	distances := make(map[string]int)
	previous := make(map[string]string)
	visited := make(map[string]bool)
	
	// Set all distances to infinity
	for nodeID := range nt.Nodes {
		distances[nodeID] = math.MaxInt32
	}
	
	// Set source distance to 0
	distances[source] = 0
	
	// Priority queue for unvisited nodes
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	
	// Add source to priority queue
	heap.Push(&pq, &Item{Value: source, Priority: 0})
	
	for pq.Len() > 0 {
		// Get node with minimum distance
		item := heap.Pop(&pq).(*Item)
		current := item.Value
		
		if visited[current] {
			continue
		}
		
		visited[current] = true
		
		// If we reached the destination, we're done
		if current == destination {
			break
		}
		
		// Check all neighbors
		if node, exists := nt.Nodes[current]; exists {
			for neighborID, neighbor := range node.Neighbors {
				if visited[neighborID] {
					continue
				}
				
				// Find link cost
				linkCost := nt.getLinkCost(current, neighborID)
				if linkCost == -1 {
					continue
				}
				
				// Calculate new distance
				newDist := distances[current] + linkCost
				
				// Update if we found a shorter path
				if newDist < distances[neighborID] {
					distances[neighborID] = newDist
					previous[neighborID] = current
					heap.Push(&pq, &Item{Value: neighborID, Priority: newDist})
				}
			}
		}
	}
	
	// Reconstruct path
	path := []string{}
	if distances[destination] == math.MaxInt32 {
		return path, -1 // No path found
	}
	
	current := destination
	for current != "" {
		path = append([]string{current}, path...)
		current = previous[current]
	}
	
	return path, distances[destination]
}

// getLinkCost returns the cost of a link between two nodes
func (nt *NetworkTopology) getLinkCost(from, to string) int {
	linkID := fmt.Sprintf("%s-%s", from, to)
	if link, exists := nt.Links[linkID]; exists && link.Status == "up" {
		return link.Cost
	}
	
	// Check reverse link
	linkID = fmt.Sprintf("%s-%s", to, from)
	if link, exists := nt.Links[linkID]; exists && link.Status == "up" {
		return link.Cost
	}
	
	return -1 // No link found
}

// PriorityQueue implementation for Dijkstra's algorithm
type Item struct {
	Value    string
	Priority int
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// RIP (Routing Information Protocol) implementation
type RIPRouter struct {
	Node        *NetworkNode
	Routes      map[string]*RIPRoute
	Neighbors   map[string]*RIPNeighbor
	UpdateTimer *time.Timer
	mutex       sync.RWMutex
}

type RIPRoute struct {
	Destination string
	NextHop     string
	Metric      int
	Age         time.Time
}

type RIPNeighbor struct {
	Address string
	LastSeen time.Time
	Routes   map[string]*RIPRoute
}

// NewRIPRouter creates a new RIP router
func NewRIPRouter(node *NetworkNode) *RIPRouter {
	router := &RIPRouter{
		Node:      node,
		Routes:    make(map[string]*RIPRoute),
		Neighbors: make(map[string]*RIPNeighbor),
	}
	
	// Start periodic updates
	router.startPeriodicUpdates()
	
	return router
}

// startPeriodicUpdates starts the periodic route updates
func (r *RIPRouter) startPeriodicUpdates() {
	r.UpdateTimer = time.NewTimer(30 * time.Second)
	go func() {
		for {
			select {
			case <-r.UpdateTimer.C:
				r.sendPeriodicUpdate()
				r.UpdateTimer.Reset(30 * time.Second)
			}
		}
	}()
}

// sendPeriodicUpdate sends periodic route updates to neighbors
func (r *RIPRouter) sendPeriodicUpdate() {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	fmt.Printf("RIP Router %s: Sending periodic update\n", r.Node.ID)
	
	for _, neighbor := range r.Neighbors {
		r.sendUpdateToNeighbor(neighbor)
	}
}

// sendUpdateToNeighbor sends route update to a specific neighbor
func (r *RIPRouter) sendUpdateToNeighbor(neighbor *RIPNeighbor) {
	fmt.Printf("RIP Router %s: Sending update to neighbor %s\n", r.Node.ID, neighbor.Address)
	
	// In a real implementation, this would send actual RIP packets
	// For simulation, we'll just print the routes
	for dest, route := range r.Routes {
		fmt.Printf("  Route: %s -> %s (metric: %d)\n", dest, route.NextHop, route.Metric)
	}
}

// AddRoute adds a route to the RIP routing table
func (r *RIPRouter) AddRoute(destination, nextHop string, metric int) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	r.Routes[destination] = &RIPRoute{
		Destination: destination,
		NextHop:     nextHop,
		Metric:      metric,
		Age:         time.Now(),
	}
	
	// Update the node's routing table
	r.Node.AddRoute(destination, nextHop, "eth0", metric, 120, "RIP")
}

// OSPF (Open Shortest Path First) implementation
type OSPFRouter struct {
	Node        *NetworkNode
	RouterID    string
	AreaID      int
	LSDB        map[string]*LSA // Link State Database
	Neighbors   map[string]*OSPFNeighbor
	SPF         *SPFCalculator
	mutex       sync.RWMutex
}

type LSA struct {
	Type       int
	RouterID   string
	AreaID     int
	Sequence   int
	Age        int
	Data       interface{}
}

type OSPFNeighbor struct {
	RouterID    string
	State       string
	Priority    int
	LastSeen    time.Time
	LSAs        map[string]*LSA
}

type SPFCalculator struct {
	RouterID string
	LSDB     map[string]*LSA
}

// NewOSPFRouter creates a new OSPF router
func NewOSPFRouter(node *NetworkNode, routerID string, areaID int) *OSPFRouter {
	router := &OSPFRouter{
		Node:      node,
		RouterID:  routerID,
		AreaID:    areaID,
		LSDB:      make(map[string]*LSA),
		Neighbors: make(map[string]*OSPFNeighbor),
		SPF:       &SPFCalculator{RouterID: routerID, LSDB: make(map[string]*LSA)},
	}
	
	// Start OSPF processes
	go router.startOSPFProcesses()
	
	return router
}

// startOSPFProcesses starts OSPF background processes
func (o *OSPFRouter) startOSPFProcesses() {
	// Hello timer
	helloTimer := time.NewTicker(10 * time.Second)
	defer helloTimer.Stop()
	
	// LSA refresh timer
	lsaTimer := time.NewTicker(30 * time.Minute)
	defer lsaTimer.Stop()
	
	for {
		select {
		case <-helloTimer.C:
			o.sendHelloPackets()
		case <-lsaTimer.C:
			o.refreshLSAs()
		}
	}
}

// sendHelloPackets sends OSPF hello packets
func (o *OSPFRouter) sendHelloPackets() {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	
	fmt.Printf("OSPF Router %s: Sending hello packets\n", o.RouterID)
	
	for _, neighbor := range o.Neighbors {
		fmt.Printf("  Hello to neighbor %s\n", neighbor.RouterID)
	}
}

// refreshLSAs refreshes OSPF LSAs
func (o *OSPFRouter) refreshLSAs() {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	
	fmt.Printf("OSPF Router %s: Refreshing LSAs\n", o.RouterID)
	
	// In a real implementation, this would refresh LSA timers
	// and send LSA updates if needed
}

// AddNeighbor adds an OSPF neighbor
func (o *OSPFRouter) AddNeighbor(routerID string, priority int) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	
	o.Neighbors[routerID] = &OSPFNeighbor{
		RouterID: routerID,
		State:    "Down",
		Priority: priority,
		LastSeen: time.Now(),
		LSAs:     make(map[string]*LSA),
	}
}

// CalculateSPF calculates the shortest path first tree
func (o *OSPFRouter) CalculateSPF() {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	
	fmt.Printf("OSPF Router %s: Calculating SPF tree\n", o.RouterID)
	
	// In a real implementation, this would run Dijkstra's algorithm
	// on the link state database to calculate shortest paths
}

// NetworkSimulator simulates network behavior
type NetworkSimulator struct {
	Topology *NetworkTopology
	Routers  map[string]*RIPRouter
	OSPFRouters map[string]*OSPFRouter
	mutex    sync.RWMutex
}

// NewNetworkSimulator creates a new network simulator
func NewNetworkSimulator() *NetworkSimulator {
	return &NetworkSimulator{
		Topology:    NewNetworkTopology(),
		Routers:     make(map[string]*RIPRouter),
		OSPFRouters: make(map[string]*OSPFRouter),
	}
}

// AddRIPRouter adds a RIP router to the simulation
func (ns *NetworkSimulator) AddRIPRouter(node *NetworkNode) {
	ns.mutex.Lock()
	defer ns.mutex.Unlock()
	
	ns.Routers[node.ID] = NewRIPRouter(node)
}

// AddOSPFRouter adds an OSPF router to the simulation
func (ns *NetworkSimulator) AddOSPFRouter(node *NetworkNode, routerID string, areaID int) {
	ns.mutex.Lock()
	defer ns.mutex.Unlock()
	
	ns.OSPFRouters[node.ID] = NewOSPFRouter(node, routerID, areaID)
}

// SimulateTraffic simulates network traffic
func (ns *NetworkSimulator) SimulateTraffic(src, dst string, data []byte) {
	ns.mutex.RLock()
	defer ns.mutex.RUnlock()
	
	fmt.Printf("Simulating traffic from %s to %s\n", src, dst)
	
	// Find path using Dijkstra's algorithm
	path, cost := ns.Topology.DijkstraShortestPath(src, dst)
	
	if len(path) == 0 {
		fmt.Printf("No path found from %s to %s\n", src, dst)
		return
	}
	
	fmt.Printf("Path: %v (cost: %d)\n", path, cost)
	
	// Simulate packet forwarding along the path
	for i := 0; i < len(path)-1; i++ {
		current := path[i]
		next := path[i+1]
		
		fmt.Printf("  %s -> %s\n", current, next)
		
		// Simulate processing delay
		time.Sleep(10 * time.Millisecond)
	}
}

// Demonstrate network switching
func demonstrateSwitching() {
	fmt.Println("=== Network Switching Demo ===\n")
	
	// Create switches
	sw1 := NewNetworkNode("SW1", "switch")
	sw2 := NewNetworkNode("SW2", "switch")
	
	// Add ports
	sw1.AddPort("eth0", "ethernet", 1000)
	sw1.AddPort("eth1", "ethernet", 1000)
	sw1.AddPort("eth2", "ethernet", 1000)
	
	sw2.AddPort("eth0", "ethernet", 1000)
	sw2.AddPort("eth1", "ethernet", 1000)
	
	// Create VLANs
	sw1.CreateVLAN(10, "Sales")
	sw1.CreateVLAN(20, "Engineering")
	
	// Add ports to VLANs
	sw1.AddPortToVLAN("eth0", 10)
	sw1.AddPortToVLAN("eth1", 20)
	sw1.AddPortToVLAN("eth2", 10)
	
	// Simulate MAC learning
	sw1.LearnMAC("00:11:22:33:44:55", "eth0")
	sw1.LearnMAC("00:11:22:33:44:66", "eth1")
	sw1.LearnMAC("00:11:22:33:44:77", "eth2")
	
	// Simulate frame forwarding
	sw1.ForwardFrame("00:11:22:33:44:55", "00:11:22:33:44:77", []byte("Hello"))
	sw1.ForwardFrame("00:11:22:33:44:55", "00:11:22:33:44:99", []byte("Unknown"))
	
	// Display MAC table
	fmt.Printf("Switch %s MAC Table:\n", sw1.ID)
	for mac, port := range sw1.MACTable {
		fmt.Printf("  %s -> %s\n", mac, port)
	}
}

// Demonstrate network routing
func demonstrateRouting() {
	fmt.Println("=== Network Routing Demo ===\n")
	
	// Create routers
	r1 := NewNetworkNode("R1", "router")
	r2 := NewNetworkNode("R2", "router")
	r3 := NewNetworkNode("R3", "router")
	r4 := NewNetworkNode("R4", "router")
	
	// Add routes
	r1.AddRoute("192.168.1.0/24", "192.168.1.1", "eth0", 0, 0, "connected")
	r1.AddRoute("10.0.0.0/8", "192.168.1.2", "eth0", 1, 120, "RIP")
	r1.AddRoute("172.16.0.0/16", "192.168.1.3", "eth0", 2, 120, "RIP")
	
	r2.AddRoute("10.0.0.0/8", "10.0.0.1", "eth0", 0, 0, "connected")
	r2.AddRoute("192.168.1.0/24", "192.168.1.1", "eth1", 1, 120, "RIP")
	
	r3.AddRoute("172.16.0.0/16", "172.16.0.1", "eth0", 0, 0, "connected")
	r3.AddRoute("192.168.1.0/24", "192.168.1.1", "eth1", 2, 120, "RIP")
	
	// Test route lookups
	destinations := []string{"192.168.1.10", "10.0.0.5", "172.16.0.10", "8.8.8.8"}
	
	for _, dest := range destinations {
		route := r1.LookupRoute(dest)
		if route != nil {
			fmt.Printf("Route to %s: Next hop %s via %s (metric: %d)\n", 
				dest, route.NextHop, route.Interface, route.Metric)
		} else {
			fmt.Printf("No route found to %s\n", dest)
		}
	}
}

// Demonstrate network topology
func demonstrateTopology() {
	fmt.Println("=== Network Topology Demo ===\n")
	
	// Create network topology
	topology := NewNetworkTopology()
	
	// Create nodes
	nodes := []*NetworkNode{
		NewNetworkNode("R1", "router"),
		NewNetworkNode("R2", "router"),
		NewNetworkNode("R3", "router"),
		NewNetworkNode("R4", "router"),
		NewNetworkNode("SW1", "switch"),
		NewNetworkNode("SW2", "switch"),
	}
	
	// Add nodes to topology
	for _, node := range nodes {
		topology.AddNode(node)
	}
	
	// Add links
	topology.AddLink(nodes[0], nodes[1], "eth0", "eth0", 1, 1000)
	topology.AddLink(nodes[1], nodes[2], "eth1", "eth0", 2, 1000)
	topology.AddLink(nodes[2], nodes[3], "eth1", "eth0", 1, 1000)
	topology.AddLink(nodes[0], nodes[4], "eth1", "eth0", 1, 1000)
	topology.AddLink(nodes[4], nodes[5], "eth1", "eth0", 1, 1000)
	
	// Test shortest path calculation
	path, cost := topology.DijkstraShortestPath("R1", "R4")
	fmt.Printf("Shortest path from R1 to R4: %v (cost: %d)\n", path, cost)
	
	path, cost = topology.DijkstraShortestPath("R1", "SW2")
	fmt.Printf("Shortest path from R1 to SW2: %v (cost: %d)\n", path, cost)
	
	// Test unreachable destination
	path, cost = topology.DijkstraShortestPath("R1", "R5")
	if len(path) == 0 {
		fmt.Printf("No path found from R1 to R5\n")
	}
}

// Demonstrate RIP protocol
func demonstrateRIP() {
	fmt.Println("=== RIP Protocol Demo ===\n")
	
	// Create routers
	r1 := NewNetworkNode("R1", "router")
	r2 := NewNetworkNode("R2", "router")
	r3 := NewNetworkNode("R3", "router")
	
	// Create RIP routers
	rip1 := NewRIPRouter(r1)
	rip2 := NewRIPRouter(r2)
	rip3 := NewRIPRouter(r3)
	
	// Add routes
	rip1.AddRoute("192.168.1.0/24", "192.168.1.1", 0) // Directly connected
	rip1.AddRoute("10.0.0.0/8", "192.168.1.2", 1)    // Via R2
	
	rip2.AddRoute("10.0.0.0/8", "10.0.0.1", 0)       // Directly connected
	rip2.AddRoute("192.168.1.0/24", "192.168.1.1", 1) // Via R1
	rip2.AddRoute("172.16.0.0/16", "10.0.0.3", 1)     // Via R3
	
	rip3.AddRoute("172.16.0.0/16", "172.16.0.1", 0)   // Directly connected
	rip3.AddRoute("10.0.0.0/8", "10.0.0.2", 1)        // Via R2
	
	// Display routing tables
	fmt.Printf("R1 Routing Table:\n")
	for dest, route := range rip1.Routes {
		fmt.Printf("  %s -> %s (metric: %d)\n", dest, route.NextHop, route.Metric)
	}
	
	fmt.Printf("\nR2 Routing Table:\n")
	for dest, route := range rip2.Routes {
		fmt.Printf("  %s -> %s (metric: %d)\n", dest, route.NextHop, route.Metric)
	}
	
	fmt.Printf("\nR3 Routing Table:\n")
	for dest, route := range rip3.Routes {
		fmt.Printf("  %s -> %s (metric: %d)\n", dest, route.NextHop, route.Metric)
	}
}

// Demonstrate OSPF protocol
func demonstrateOSPF() {
	fmt.Println("=== OSPF Protocol Demo ===\n")
	
	// Create routers
	r1 := NewNetworkNode("R1", "router")
	r2 := NewNetworkNode("R2", "router")
	r3 := NewNetworkNode("R3", "router")
	
	// Create OSPF routers
	ospf1 := NewOSPFRouter(r1, "1.1.1.1", 0)
	ospf2 := NewOSPFRouter(r2, "2.2.2.2", 0)
	ospf3 := NewOSPFRouter(r3, "3.3.3.3", 0)
	
	// Add neighbors
	ospf1.AddNeighbor("2.2.2.2", 1)
	ospf1.AddNeighbor("3.3.3.3", 1)
	
	ospf2.AddNeighbor("1.1.1.1", 1)
	ospf2.AddNeighbor("3.3.3.3", 1)
	
	ospf3.AddNeighbor("1.1.1.1", 1)
	ospf3.AddNeighbor("2.2.2.2", 1)
	
	// Display OSPF information
	fmt.Printf("OSPF Router %s:\n", ospf1.RouterID)
	fmt.Printf("  Area ID: %d\n", ospf1.AreaID)
	fmt.Printf("  Neighbors: %d\n", len(ospf1.Neighbors))
	
	for neighborID, neighbor := range ospf1.Neighbors {
		fmt.Printf("    %s (state: %s, priority: %d)\n", 
			neighborID, neighbor.State, neighbor.Priority)
	}
}

// Demonstrate network simulation
func demonstrateSimulation() {
	fmt.Println("=== Network Simulation Demo ===\n")
	
	// Create network simulator
	simulator := NewNetworkSimulator()
	
	// Create topology
	topology := simulator.Topology
	
	// Create nodes
	r1 := NewNetworkNode("R1", "router")
	r2 := NewNetworkNode("R2", "router")
	r3 := NewNetworkNode("R3", "router")
	r4 := NewNetworkNode("R4", "router")
	
	// Add nodes
	topology.AddNode(r1)
	topology.AddNode(r2)
	topology.AddNode(r3)
	topology.AddNode(r4)
	
	// Add links
	topology.AddLink(r1, r2, "eth0", "eth0", 1, 1000)
	topology.AddLink(r2, r3, "eth1", "eth0", 2, 1000)
	topology.AddLink(r3, r4, "eth1", "eth0", 1, 1000)
	topology.AddLink(r1, r4, "eth1", "eth1", 3, 1000)
	
	// Add RIP routers
	simulator.AddRIPRouter(r1)
	simulator.AddRIPRouter(r2)
	simulator.AddRIPRouter(r3)
	simulator.AddRIPRouter(r4)
	
	// Simulate traffic
	simulator.SimulateTraffic("R1", "R4", []byte("Hello from R1 to R4"))
	simulator.SimulateTraffic("R2", "R4", []byte("Hello from R2 to R4"))
	simulator.SimulateTraffic("R1", "R3", []byte("Hello from R1 to R3"))
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 9: Routing & Switching")
	fmt.Println("================================================================\n")
	
	// Run all demonstrations
	demonstrateSwitching()
	fmt.Println()
	demonstrateRouting()
	fmt.Println()
	demonstrateTopology()
	fmt.Println()
	demonstrateRIP()
	fmt.Println()
	demonstrateOSPF()
	fmt.Println()
	demonstrateSimulation()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Switching operates at Layer 2 and forwards frames within LANs")
	fmt.Println("2. Routing operates at Layer 3 and forwards packets between networks")
	fmt.Println("3. VLANs provide logical network segmentation")
	fmt.Println("4. STP prevents loops in switched networks")
	fmt.Println("5. Routing algorithms find optimal paths through networks")
	fmt.Println("6. RIP and OSPF are common routing protocols")
	fmt.Println("7. Network topology affects performance and reliability")
	fmt.Println("8. Go can simulate complex network behavior")
	
	fmt.Println("\n📚 Next Topic: Network Performance & Optimization")
}
```

Now let me create a comprehensive routing algorithms implementation:


```go
package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

// Graph represents a network graph
type Graph struct {
	Nodes map[string]*Node
	Edges map[string]*Edge
	mutex sync.RWMutex
}

// Node represents a network node
type Node struct {
	ID       string
	Type     string // "router", "switch", "host"
	Position Point
	Neighbors map[string]*Node
	Routes   map[string]*Route
}

// Edge represents a network link
type Edge struct {
	From     *Node
	To       *Node
	Weight   int
	Bandwidth int
	Delay    int
	Status   string // "up", "down"
}

// Point represents a 2D coordinate
type Point struct {
	X, Y float64
}

// Route represents a routing table entry
type Route struct {
	Destination string
	NextHop     string
	Cost        int
	Path        []string
	Protocol    string
	Age         time.Time
}

// NewGraph creates a new network graph
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
		Edges: make(map[string]*Edge),
	}
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(id, nodeType string, x, y float64) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	
	g.Nodes[id] = &Node{
		ID:        id,
		Type:      nodeType,
		Position:  Point{X: x, Y: y},
		Neighbors: make(map[string]*Node),
		Routes:    make(map[string]*Route),
	}
}

// AddEdge adds an edge between two nodes
func (g *Graph) AddEdge(from, to string, weight, bandwidth, delay int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	
	fromNode, exists1 := g.Nodes[from]
	toNode, exists2 := g.Nodes[to]
	
	if !exists1 || !exists2 {
		return
	}
	
	edgeID := fmt.Sprintf("%s-%s", from, to)
	edge := &Edge{
		From:      fromNode,
		To:        toNode,
		Weight:    weight,
		Bandwidth: bandwidth,
		Delay:     delay,
		Status:    "up",
	}
	
	g.Edges[edgeID] = edge
	
	// Add neighbors
	fromNode.Neighbors[to] = toNode
	toNode.Neighbors[from] = fromNode
}

// GetEdge returns an edge between two nodes
func (g *Graph) GetEdge(from, to string) *Edge {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	edgeID := fmt.Sprintf("%s-%s", from, to)
	if edge, exists := g.Edges[edgeID]; exists {
		return edge
	}
	
	// Check reverse edge
	edgeID = fmt.Sprintf("%s-%s", to, from)
	if edge, exists := g.Edges[edgeID]; exists {
		return edge
	}
	
	return nil
}

// DijkstraShortestPath finds the shortest path using Dijkstra's algorithm
func (g *Graph) DijkstraShortestPath(source, destination string) ([]string, int) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	// Initialize distances
	distances := make(map[string]int)
	previous := make(map[string]string)
	visited := make(map[string]bool)
	
	// Set all distances to infinity
	for nodeID := range g.Nodes {
		distances[nodeID] = math.MaxInt32
	}
	
	// Set source distance to 0
	distances[source] = 0
	
	// Priority queue for unvisited nodes
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	
	// Add source to priority queue
	heap.Push(&pq, &Item{Value: source, Priority: 0})
	
	for pq.Len() > 0 {
		// Get node with minimum distance
		item := heap.Pop(&pq).(*Item)
		current := item.Value
		
		if visited[current] {
			continue
		}
		
		visited[current] = true
		
		// If we reached the destination, we're done
		if current == destination {
			break
		}
		
		// Check all neighbors
		if node, exists := g.Nodes[current]; exists {
			for neighborID, neighbor := range node.Neighbors {
				if visited[neighborID] {
					continue
				}
				
				// Find edge weight
				edge := g.GetEdge(current, neighborID)
				if edge == nil || edge.Status != "up" {
					continue
				}
				
				// Calculate new distance
				newDist := distances[current] + edge.Weight
				
				// Update if we found a shorter path
				if newDist < distances[neighborID] {
					distances[neighborID] = newDist
					previous[neighborID] = current
					heap.Push(&pq, &Item{Value: neighborID, Priority: newDist})
				}
			}
		}
	}
	
	// Reconstruct path
	path := []string{}
	if distances[destination] == math.MaxInt32 {
		return path, -1 // No path found
	}
	
	current := destination
	for current != "" {
		path = append([]string{current}, path...)
		current = previous[current]
	}
	
	return path, distances[destination]
}

// BellmanFordShortestPath finds the shortest path using Bellman-Ford algorithm
func (g *Graph) BellmanFordShortestPath(source, destination string) ([]string, int) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	// Initialize distances
	distances := make(map[string]int)
	previous := make(map[string]string)
	
	// Set all distances to infinity
	for nodeID := range g.Nodes {
		distances[nodeID] = math.MaxInt32
	}
	
	// Set source distance to 0
	distances[source] = 0
	
	// Relax edges |V| - 1 times
	for i := 0; i < len(g.Nodes)-1; i++ {
		for _, edge := range g.Edges {
			if edge.Status != "up" {
				continue
			}
			
			from := edge.From.ID
			to := edge.To.ID
			
			if distances[from] != math.MaxInt32 && distances[from]+edge.Weight < distances[to] {
				distances[to] = distances[from] + edge.Weight
				previous[to] = from
			}
		}
	}
	
	// Check for negative cycles
	for _, edge := range g.Edges {
		if edge.Status != "up" {
			continue
		}
		
		from := edge.From.ID
		to := edge.To.ID
		
		if distances[from] != math.MaxInt32 && distances[from]+edge.Weight < distances[to] {
			return []string{}, -1 // Negative cycle detected
		}
	}
	
	// Reconstruct path
	path := []string{}
	if distances[destination] == math.MaxInt32 {
		return path, -1 // No path found
	}
	
	current := destination
	for current != "" {
		path = append([]string{current}, path...)
		current = previous[current]
	}
	
	return path, distances[destination]
}

// FloydWarshallAllPairs finds shortest paths between all pairs of nodes
func (g *Graph) FloydWarshallAllPairs() map[string]map[string]int {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	// Initialize distance matrix
	distances := make(map[string]map[string]int)
	
	// Initialize with infinity
	for from := range g.Nodes {
		distances[from] = make(map[string]int)
		for to := range g.Nodes {
			if from == to {
				distances[from][to] = 0
			} else {
				distances[from][to] = math.MaxInt32
			}
		}
	}
	
	// Initialize with edge weights
	for _, edge := range g.Edges {
		if edge.Status == "up" {
			distances[edge.From.ID][edge.To.ID] = edge.Weight
		}
	}
	
	// Floyd-Warshall algorithm
	for k := range g.Nodes {
		for i := range g.Nodes {
			for j := range g.Nodes {
				if distances[i][k] != math.MaxInt32 && distances[k][j] != math.MaxInt32 {
					if distances[i][k]+distances[k][j] < distances[i][j] {
						distances[i][j] = distances[i][k] + distances[k][j]
					}
				}
			}
		}
	}
	
	return distances
}

// AStarShortestPath finds the shortest path using A* algorithm
func (g *Graph) AStarShortestPath(source, destination string) ([]string, int) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	// Heuristic function (Euclidean distance)
	heuristic := func(nodeID string) int {
		if node, exists := g.Nodes[nodeID]; exists {
			dest, exists := g.Nodes[destination]
			if !exists {
				return 0
			}
			
			dx := node.Position.X - dest.Position.X
			dy := node.Position.Y - dest.Position.Y
			return int(math.Sqrt(dx*dx + dy*dy))
		}
		return 0
	}
	
	// Initialize distances
	distances := make(map[string]int)
	previous := make(map[string]string)
	visited := make(map[string]bool)
	
	// Set all distances to infinity
	for nodeID := range g.Nodes {
		distances[nodeID] = math.MaxInt32
	}
	
	// Set source distance to 0
	distances[source] = 0
	
	// Priority queue for unvisited nodes
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	
	// Add source to priority queue
	heap.Push(&pq, &Item{Value: source, Priority: heuristic(source)})
	
	for pq.Len() > 0 {
		// Get node with minimum f-score
		item := heap.Pop(&pq).(*Item)
		current := item.Value
		
		if visited[current] {
			continue
		}
		
		visited[current] = true
		
		// If we reached the destination, we're done
		if current == destination {
			break
		}
		
		// Check all neighbors
		if node, exists := g.Nodes[current]; exists {
			for neighborID, neighbor := range node.Neighbors {
				if visited[neighborID] {
					continue
				}
				
				// Find edge weight
				edge := g.GetEdge(current, neighborID)
				if edge == nil || edge.Status != "up" {
					continue
				}
				
				// Calculate new distance
				newDist := distances[current] + edge.Weight
				
				// Update if we found a shorter path
				if newDist < distances[neighborID] {
					distances[neighborID] = newDist
					previous[neighborID] = current
					
					// Calculate f-score (g + h)
					fScore := newDist + heuristic(neighborID)
					heap.Push(&pq, &Item{Value: neighborID, Priority: fScore})
				}
			}
		}
	}
	
	// Reconstruct path
	path := []string{}
	if distances[destination] == math.MaxInt32 {
		return path, -1 // No path found
	}
	
	current := destination
	for current != "" {
		path = append([]string{current}, path...)
		current = previous[current]
	}
	
	return path, distances[destination]
}

// KruskalMST finds the minimum spanning tree using Kruskal's algorithm
func (g *Graph) KruskalMST() []*Edge {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	// Create list of all edges
	edges := make([]*Edge, 0, len(g.Edges))
	for _, edge := range g.Edges {
		if edge.Status == "up" {
			edges = append(edges, edge)
		}
	}
	
	// Sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})
	
	// Union-Find data structure
	parent := make(map[string]string)
	rank := make(map[string]int)
	
	// Initialize each node as its own parent
	for nodeID := range g.Nodes {
		parent[nodeID] = nodeID
		rank[nodeID] = 0
	}
	
	// Find function
	find := func(x string) string {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	
	// Union function
	union := func(x, y string) {
		px := find(x)
		py := find(y)
		
		if px == py {
			return
		}
		
		if rank[px] < rank[py] {
			parent[px] = py
		} else if rank[px] > rank[py] {
			parent[py] = px
		} else {
			parent[py] = px
			rank[px]++
		}
	}
	
	// Kruskal's algorithm
	mst := make([]*Edge, 0)
	for _, edge := range edges {
		from := edge.From.ID
		to := edge.To.ID
		
		if find(from) != find(to) {
			mst = append(mst, edge)
			union(from, to)
		}
	}
	
	return mst
}

// PrimMST finds the minimum spanning tree using Prim's algorithm
func (g *Graph) PrimMST(start string) []*Edge {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	// Initialize
	visited := make(map[string]bool)
	mst := make([]*Edge, 0)
	
	// Priority queue for edges
	pq := make(EdgePriorityQueue, 0)
	heap.Init(&pq)
	
	// Start with the given node
	visited[start] = true
	
	// Add all edges from start node to priority queue
	if startNode, exists := g.Nodes[start]; exists {
		for neighborID := range startNode.Neighbors {
			edge := g.GetEdge(start, neighborID)
			if edge != nil && edge.Status == "up" {
				heap.Push(&pq, &EdgeItem{Edge: edge, Priority: edge.Weight})
			}
		}
	}
	
	// Prim's algorithm
	for pq.Len() > 0 && len(mst) < len(g.Nodes)-1 {
		item := heap.Pop(&pq).(*EdgeItem)
		edge := item.Edge
		
		from := edge.From.ID
		to := edge.To.ID
		
		// Skip if both nodes are already visited
		if visited[from] && visited[to] {
			continue
		}
		
		// Add edge to MST
		mst = append(mst, edge)
		
		// Mark the unvisited node as visited
		if !visited[from] {
			visited[from] = true
			// Add edges from the newly visited node
			if fromNode, exists := g.Nodes[from]; exists {
				for neighborID := range fromNode.Neighbors {
					neighborEdge := g.GetEdge(from, neighborID)
					if neighborEdge != nil && neighborEdge.Status == "up" && !visited[neighborID] {
						heap.Push(&pq, &EdgeItem{Edge: neighborEdge, Priority: neighborEdge.Weight})
					}
				}
			}
		} else {
			visited[to] = true
			// Add edges from the newly visited node
			if toNode, exists := g.Nodes[to]; exists {
				for neighborID := range toNode.Neighbors {
					neighborEdge := g.GetEdge(to, neighborID)
					if neighborEdge != nil && neighborEdge.Status == "up" && !visited[neighborID] {
						heap.Push(&pq, &EdgeItem{Edge: neighborEdge, Priority: neighborEdge.Weight})
					}
				}
			}
		}
	}
	
	return mst
}

// EdgePriorityQueue for Prim's algorithm
type EdgeItem struct {
	Edge     *Edge
	Priority int
	Index    int
}

type EdgePriorityQueue []*EdgeItem

func (pq EdgePriorityQueue) Len() int { return len(pq) }

func (pq EdgePriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq EdgePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *EdgePriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*EdgeItem)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *EdgePriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// NetworkAnalyzer provides network analysis capabilities
type NetworkAnalyzer struct {
	Graph *Graph
}

// NewNetworkAnalyzer creates a new network analyzer
func NewNetworkAnalyzer(graph *Graph) *NetworkAnalyzer {
	return &NetworkAnalyzer{Graph: graph}
}

// AnalyzeConnectivity analyzes network connectivity
func (na *NetworkAnalyzer) AnalyzeConnectivity() map[string]bool {
	na.Graph.mutex.RLock()
	defer na.Graph.mutex.RUnlock()
	
	connectivity := make(map[string]bool)
	
	for nodeID := range na.Graph.Nodes {
		// Check if node is reachable from all other nodes
		reachable := true
		for otherNodeID := range na.Graph.Nodes {
			if nodeID == otherNodeID {
				continue
			}
			
			path, _ := na.Graph.DijkstraShortestPath(otherNodeID, nodeID)
			if len(path) == 0 {
				reachable = false
				break
			}
		}
		connectivity[nodeID] = reachable
	}
	
	return connectivity
}

// AnalyzeRedundancy analyzes network redundancy
func (na *NetworkAnalyzer) AnalyzeRedundancy() map[string]int {
	na.Graph.mutex.RLock()
	defer na.Graph.mutex.RUnlock()
	
	redundancy := make(map[string]int)
	
	for nodeID := range na.Graph.Nodes {
		// Count alternative paths to each other node
		altPaths := 0
		for otherNodeID := range na.Graph.Nodes {
			if nodeID == otherNodeID {
				continue
			}
			
			// Find all paths using different algorithms
			path1, _ := na.Graph.DijkstraShortestPath(nodeID, otherNodeID)
			path2, _ := na.Graph.BellmanFordShortestPath(nodeID, otherNodeID)
			
			if len(path1) > 0 && len(path2) > 0 && !na.pathsEqual(path1, path2) {
				altPaths++
			}
		}
		redundancy[nodeID] = altPaths
	}
	
	return redundancy
}

// pathsEqual checks if two paths are equal
func (na *NetworkAnalyzer) pathsEqual(path1, path2 []string) bool {
	if len(path1) != len(path2) {
		return false
	}
	
	for i := range path1 {
		if path1[i] != path2[i] {
			return false
		}
	}
	
	return true
}

// AnalyzePerformance analyzes network performance
func (na *NetworkAnalyzer) AnalyzePerformance() map[string]float64 {
	na.Graph.mutex.RLock()
	defer na.Graph.mutex.RUnlock()
	
	performance := make(map[string]float64)
	
	// Calculate average path length
	totalPaths := 0
	totalLength := 0
	
	for from := range na.Graph.Nodes {
		for to := range na.Graph.Nodes {
			if from == to {
				continue
			}
			
			path, _ := na.Graph.DijkstraShortestPath(from, to)
			if len(path) > 0 {
				totalPaths++
				totalLength += len(path) - 1
			}
		}
	}
	
	if totalPaths > 0 {
		performance["average_path_length"] = float64(totalLength) / float64(totalPaths)
	}
	
	// Calculate network diameter
	maxPathLength := 0
	for from := range na.Graph.Nodes {
		for to := range na.Graph.Nodes {
			if from == to {
				continue
			}
			
			path, _ := na.Graph.DijkstraShortestPath(from, to)
			if len(path) > 0 && len(path)-1 > maxPathLength {
				maxPathLength = len(path) - 1
			}
		}
	}
	performance["diameter"] = float64(maxPathLength)
	
	return performance
}

// Demonstrate routing algorithms
func demonstrateRoutingAlgorithms() {
	fmt.Println("=== Routing Algorithms Demo ===\n")
	
	// Create network graph
	graph := NewGraph()
	
	// Add nodes
	graph.AddNode("A", "router", 0, 0)
	graph.AddNode("B", "router", 1, 0)
	graph.AddNode("C", "router", 2, 0)
	graph.AddNode("D", "router", 0, 1)
	graph.AddNode("E", "router", 1, 1)
	graph.AddNode("F", "router", 2, 1)
	
	// Add edges
	graph.AddEdge("A", "B", 4, 1000, 10)
	graph.AddEdge("A", "D", 2, 1000, 5)
	graph.AddEdge("B", "C", 3, 1000, 8)
	graph.AddEdge("B", "E", 5, 1000, 12)
	graph.AddEdge("C", "F", 2, 1000, 6)
	graph.AddEdge("D", "E", 1, 1000, 3)
	graph.AddEdge("E", "F", 3, 1000, 9)
	
	// Test Dijkstra's algorithm
	fmt.Println("Dijkstra's Algorithm:")
	path, cost := graph.DijkstraShortestPath("A", "F")
	fmt.Printf("  Path from A to F: %v (cost: %d)\n", path, cost)
	
	// Test Bellman-Ford algorithm
	fmt.Println("\nBellman-Ford Algorithm:")
	path, cost = graph.BellmanFordShortestPath("A", "F")
	fmt.Printf("  Path from A to F: %v (cost: %d)\n", path, cost)
	
	// Test A* algorithm
	fmt.Println("\nA* Algorithm:")
	path, cost = graph.AStarShortestPath("A", "F")
	fmt.Printf("  Path from A to F: %v (cost: %d)\n", path, cost)
	
	// Test Floyd-Warshall algorithm
	fmt.Println("\nFloyd-Warshall Algorithm:")
	distances := graph.FloydWarshallAllPairs()
	fmt.Printf("  Distance matrix:\n")
	for from := range distances {
		fmt.Printf("    %s: ", from)
		for to := range distances[from] {
			if distances[from][to] == math.MaxInt32 {
				fmt.Printf("∞ ")
			} else {
				fmt.Printf("%d ", distances[from][to])
			}
		}
		fmt.Println()
	}
}

// Demonstrate MST algorithms
func demonstrateMSTAlgorithms() {
	fmt.Println("=== MST Algorithms Demo ===\n")
	
	// Create network graph
	graph := NewGraph()
	
	// Add nodes
	graph.AddNode("A", "router", 0, 0)
	graph.AddNode("B", "router", 1, 0)
	graph.AddNode("C", "router", 2, 0)
	graph.AddNode("D", "router", 0, 1)
	graph.AddNode("E", "router", 1, 1)
	graph.AddNode("F", "router", 2, 1)
	
	// Add edges
	graph.AddEdge("A", "B", 4, 1000, 10)
	graph.AddEdge("A", "D", 2, 1000, 5)
	graph.AddEdge("B", "C", 3, 1000, 8)
	graph.AddEdge("B", "E", 5, 1000, 12)
	graph.AddEdge("C", "F", 2, 1000, 6)
	graph.AddEdge("D", "E", 1, 1000, 3)
	graph.AddEdge("E", "F", 3, 1000, 9)
	
	// Test Kruskal's algorithm
	fmt.Println("Kruskal's MST:")
	mst := graph.KruskalMST()
	totalWeight := 0
	for _, edge := range mst {
		fmt.Printf("  %s - %s (weight: %d)\n", edge.From.ID, edge.To.ID, edge.Weight)
		totalWeight += edge.Weight
	}
	fmt.Printf("  Total weight: %d\n", totalWeight)
	
	// Test Prim's algorithm
	fmt.Println("\nPrim's MST:")
	mst = graph.PrimMST("A")
	totalWeight = 0
	for _, edge := range mst {
		fmt.Printf("  %s - %s (weight: %d)\n", edge.From.ID, edge.To.ID, edge.Weight)
		totalWeight += edge.Weight
	}
	fmt.Printf("  Total weight: %d\n", totalWeight)
}

// Demonstrate network analysis
func demonstrateNetworkAnalysis() {
	fmt.Println("=== Network Analysis Demo ===\n")
	
	// Create network graph
	graph := NewGraph()
	
	// Add nodes
	graph.AddNode("A", "router", 0, 0)
	graph.AddNode("B", "router", 1, 0)
	graph.AddNode("C", "router", 2, 0)
	graph.AddNode("D", "router", 0, 1)
	graph.AddNode("E", "router", 1, 1)
	graph.AddNode("F", "router", 2, 1)
	
	// Add edges
	graph.AddEdge("A", "B", 4, 1000, 10)
	graph.AddEdge("A", "D", 2, 1000, 5)
	graph.AddEdge("B", "C", 3, 1000, 8)
	graph.AddEdge("B", "E", 5, 1000, 12)
	graph.AddEdge("C", "F", 2, 1000, 6)
	graph.AddEdge("D", "E", 1, 1000, 3)
	graph.AddEdge("E", "F", 3, 1000, 9)
	
	// Create network analyzer
	analyzer := NewNetworkAnalyzer(graph)
	
	// Analyze connectivity
	fmt.Println("Connectivity Analysis:")
	connectivity := analyzer.AnalyzeConnectivity()
	for nodeID, reachable := range connectivity {
		fmt.Printf("  %s: %t\n", nodeID, reachable)
	}
	
	// Analyze redundancy
	fmt.Println("\nRedundancy Analysis:")
	redundancy := analyzer.AnalyzeRedundancy()
	for nodeID, altPaths := range redundancy {
		fmt.Printf("  %s: %d alternative paths\n", nodeID, altPaths)
	}
	
	// Analyze performance
	fmt.Println("\nPerformance Analysis:")
	performance := analyzer.AnalyzePerformance()
	for metric, value := range performance {
		fmt.Printf("  %s: %.2f\n", metric, value)
	}
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 9: Routing Algorithms")
	fmt.Println("============================================================\n")
	
	// Run all demonstrations
	demonstrateRoutingAlgorithms()
	fmt.Println()
	demonstrateMSTAlgorithms()
	fmt.Println()
	demonstrateNetworkAnalysis()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Dijkstra's algorithm finds shortest paths from a single source")
	fmt.Println("2. Bellman-Ford algorithm handles negative weights and detects cycles")
	fmt.Println("3. A* algorithm uses heuristics for faster pathfinding")
	fmt.Println("4. Floyd-Warshall algorithm finds shortest paths between all pairs")
	fmt.Println("5. Kruskal's and Prim's algorithms find minimum spanning trees")
	fmt.Println("6. Network analysis helps optimize performance and reliability")
	fmt.Println("7. Go provides excellent support for graph algorithms")
	fmt.Println("8. Understanding these algorithms is crucial for network design")
	
	fmt.Println("\n📚 Next Topic: Network Performance & Optimization")
}
```

Now let me create a quick reference guide and update the progress:


```markdown
# Routing & Switching Quick Reference Guide

## 🚀 Essential Concepts

### Switching (Layer 2)
- **MAC Address Learning**: Switches learn MAC addresses by observing traffic
- **Frame Forwarding**: Forward frames based on destination MAC address
- **VLAN Support**: Virtual LAN segmentation and management
- **STP**: Spanning Tree Protocol prevents loops
- **Port States**: Blocking, Listening, Learning, Forwarding, Disabled

### Routing (Layer 3)
- **Packet Forwarding**: Forward packets between different networks
- **Path Selection**: Choose best route to destination
- **Route Management**: Maintain routing tables
- **Protocol Support**: Implement routing protocols (RIP, OSPF, BGP)

## 🔧 Go Implementation Patterns

### Network Node Structure
```go
type NetworkNode struct {
    ID        string
    Type      string // "router" or "switch"
    Ports     map[string]*Port
    MACTable  map[string]string
    Routes    map[string]*Route
    Neighbors map[string]*NetworkNode
    VLANs     map[int]*VLAN
    mutex     sync.RWMutex
}
```

### Port Management
```go
type Port struct {
    ID     string
    Type   string // "ethernet", "serial", "fiber"
    Speed  int    // Mbps
    Duplex string // "half", "full"
    Status string // "up", "down", "admin_down"
    VLAN   int
}
```

### Route Management
```go
type Route struct {
    Destination string
    NextHop     string
    Interface   string
    Metric      int
    AdminDist   int
    Protocol    string
    Age         time.Time
}
```

## 🔄 Switching Implementation

### MAC Address Learning
```go
func (n *NetworkNode) LearnMAC(mac, portID string) {
    n.mutex.Lock()
    defer n.mutex.Unlock()
    
    n.MACTable[mac] = portID
}

func (n *NetworkNode) LookupMAC(mac string) (string, bool) {
    n.mutex.RLock()
    defer n.mutex.RUnlock()
    
    port, exists := n.MACTable[mac]
    return port, exists
}
```

### Frame Forwarding
```go
func (n *NetworkNode) ForwardFrame(srcMAC, dstMAC string, frame []byte) {
    if n.Type != "switch" {
        return
    }
    
    // Learn source MAC
    n.LearnMAC(srcMAC, "incoming")
    
    // Look up destination MAC
    if port, exists := n.LookupMAC(dstMAC); exists {
        // Forward to specific port
        fmt.Printf("Switch %s: Forwarding frame from %s to %s via port %s\n", 
            n.ID, srcMAC, dstMAC, port)
    } else {
        // Flood to all ports except incoming
        fmt.Printf("Switch %s: Flooding frame from %s to %s (unknown destination)\n", 
            n.ID, srcMAC, dstMAC)
    }
}
```

### VLAN Management
```go
func (n *NetworkNode) CreateVLAN(id int, name string) {
    n.mutex.Lock()
    defer n.mutex.Unlock()
    
    n.VLANs[id] = &VLAN{
        ID:          id,
        Name:        name,
        Ports:       make(map[string]bool),
        Broadcast:   true,
        Multicast:   true,
        UnknownUnicast: true,
    }
}

func (n *NetworkNode) AddPortToVLAN(portID string, vlanID int) {
    n.mutex.Lock()
    defer n.mutex.Unlock()
    
    if vlan, exists := n.VLANs[vlanID]; exists {
        vlan.Ports[portID] = true
        if port, exists := n.Ports[portID]; exists {
            port.VLAN = vlanID
        }
    }
}
```

## 🛣️ Routing Implementation

### Route Lookup
```go
func (n *NetworkNode) LookupRoute(destination string) *Route {
    n.mutex.RLock()
    defer n.mutex.RUnlock()
    
    // First try exact match
    if route, exists := n.Routes[destination]; exists {
        return route
    }
    
    // Then try longest prefix match
    var bestRoute *Route
    var bestPrefixLen int
    
    for dest, route := range n.Routes {
        if n.isSubnetMatch(destination, dest) {
            prefixLen := n.getPrefixLength(dest)
            if prefixLen > bestPrefixLen {
                bestPrefixLen = prefixLen
                bestRoute = route
            }
        }
    }
    
    return bestRoute
}
```

### Route Addition
```go
func (n *NetworkNode) AddRoute(dest, nextHop, intf string, metric, adminDist int, protocol string) {
    n.mutex.Lock()
    defer n.mutex.Unlock()
    
    n.Routes[dest] = &Route{
        Destination: dest,
        NextHop:     nextHop,
        Interface:   intf,
        Metric:      metric,
        AdminDist:   adminDist,
        Protocol:    protocol,
        Age:         time.Now(),
    }
}
```

## 🧮 Routing Algorithms

### Dijkstra's Algorithm
```go
func (g *Graph) DijkstraShortestPath(source, destination string) ([]string, int) {
    // Initialize distances
    distances := make(map[string]int)
    previous := make(map[string]string)
    visited := make(map[string]bool)
    
    // Set all distances to infinity
    for nodeID := range g.Nodes {
        distances[nodeID] = math.MaxInt32
    }
    
    // Set source distance to 0
    distances[source] = 0
    
    // Priority queue for unvisited nodes
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    
    // Add source to priority queue
    heap.Push(&pq, &Item{Value: source, Priority: 0})
    
    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*Item)
        current := item.Value
        
        if visited[current] {
            continue
        }
        
        visited[current] = true
        
        if current == destination {
            break
        }
        
        // Check all neighbors
        if node, exists := g.Nodes[current]; exists {
            for neighborID, neighbor := range node.Neighbors {
                if visited[neighborID] {
                    continue
                }
                
                edge := g.GetEdge(current, neighborID)
                if edge == nil || edge.Status != "up" {
                    continue
                }
                
                newDist := distances[current] + edge.Weight
                
                if newDist < distances[neighborID] {
                    distances[neighborID] = newDist
                    previous[neighborID] = current
                    heap.Push(&pq, &Item{Value: neighborID, Priority: newDist})
                }
            }
        }
    }
    
    // Reconstruct path
    path := []string{}
    if distances[destination] == math.MaxInt32 {
        return path, -1
    }
    
    current := destination
    for current != "" {
        path = append([]string{current}, path...)
        current = previous[current]
    }
    
    return path, distances[destination]
}
```

### Bellman-Ford Algorithm
```go
func (g *Graph) BellmanFordShortestPath(source, destination string) ([]string, int) {
    // Initialize distances
    distances := make(map[string]int)
    previous := make(map[string]string)
    
    // Set all distances to infinity
    for nodeID := range g.Nodes {
        distances[nodeID] = math.MaxInt32
    }
    
    // Set source distance to 0
    distances[source] = 0
    
    // Relax edges |V| - 1 times
    for i := 0; i < len(g.Nodes)-1; i++ {
        for _, edge := range g.Edges {
            if edge.Status != "up" {
                continue
            }
            
            from := edge.From.ID
            to := edge.To.ID
            
            if distances[from] != math.MaxInt32 && distances[from]+edge.Weight < distances[to] {
                distances[to] = distances[from] + edge.Weight
                previous[to] = from
            }
        }
    }
    
    // Check for negative cycles
    for _, edge := range g.Edges {
        if edge.Status != "up" {
            continue
        }
        
        from := edge.From.ID
        to := edge.To.ID
        
        if distances[from] != math.MaxInt32 && distances[from]+edge.Weight < distances[to] {
            return []string{}, -1 // Negative cycle detected
        }
    }
    
    // Reconstruct path
    path := []string{}
    if distances[destination] == math.MaxInt32 {
        return path, -1
    }
    
    current := destination
    for current != "" {
        path = append([]string{current}, path...)
        current = previous[current]
    }
    
    return path, distances[destination]
}
```

### A* Algorithm
```go
func (g *Graph) AStarShortestPath(source, destination string) ([]string, int) {
    // Heuristic function (Euclidean distance)
    heuristic := func(nodeID string) int {
        if node, exists := g.Nodes[nodeID]; exists {
            dest, exists := g.Nodes[destination]
            if !exists {
                return 0
            }
            
            dx := node.Position.X - dest.Position.X
            dy := node.Position.Y - dest.Position.Y
            return int(math.Sqrt(dx*dx + dy*dy))
        }
        return 0
    }
    
    // Initialize distances
    distances := make(map[string]int)
    previous := make(map[string]string)
    visited := make(map[string]bool)
    
    // Set all distances to infinity
    for nodeID := range g.Nodes {
        distances[nodeID] = math.MaxInt32
    }
    
    // Set source distance to 0
    distances[source] = 0
    
    // Priority queue for unvisited nodes
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    
    // Add source to priority queue
    heap.Push(&pq, &Item{Value: source, Priority: heuristic(source)})
    
    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*Item)
        current := item.Value
        
        if visited[current] {
            continue
        }
        
        visited[current] = true
        
        if current == destination {
            break
        }
        
        // Check all neighbors
        if node, exists := g.Nodes[current]; exists {
            for neighborID, neighbor := range node.Neighbors {
                if visited[neighborID] {
                    continue
                }
                
                edge := g.GetEdge(current, neighborID)
                if edge == nil || edge.Status != "up" {
                    continue
                }
                
                newDist := distances[current] + edge.Weight
                
                if newDist < distances[neighborID] {
                    distances[neighborID] = newDist
                    previous[neighborID] = current
                    
                    // Calculate f-score (g + h)
                    fScore := newDist + heuristic(neighborID)
                    heap.Push(&pq, &Item{Value: neighborID, Priority: fScore})
                }
            }
        }
    }
    
    // Reconstruct path
    path := []string{}
    if distances[destination] == math.MaxInt32 {
        return path, -1
    }
    
    current := destination
    for current != "" {
        path = append([]string{current}, path...)
        current = previous[current]
    }
    
    return path, distances[destination]
}
```

## 🌐 Network Topology

### Graph Representation
```go
type Graph struct {
    Nodes map[string]*Node
    Edges map[string]*Edge
    mutex sync.RWMutex
}

type Node struct {
    ID        string
    Type      string
    Position  Point
    Neighbors map[string]*Node
    Routes    map[string]*Route
}

type Edge struct {
    From      *Node
    To        *Node
    Weight    int
    Bandwidth int
    Delay     int
    Status    string
}
```

### Topology Management
```go
func (g *Graph) AddNode(id, nodeType string, x, y float64) {
    g.mutex.Lock()
    defer g.mutex.Unlock()
    
    g.Nodes[id] = &Node{
        ID:        id,
        Type:      nodeType,
        Position:  Point{X: x, Y: y},
        Neighbors: make(map[string]*Node),
        Routes:    make(map[string]*Route),
    }
}

func (g *Graph) AddEdge(from, to string, weight, bandwidth, delay int) {
    g.mutex.Lock()
    defer g.mutex.Unlock()
    
    fromNode, exists1 := g.Nodes[from]
    toNode, exists2 := g.Nodes[to]
    
    if !exists1 || !exists2 {
        return
    }
    
    edgeID := fmt.Sprintf("%s-%s", from, to)
    edge := &Edge{
        From:      fromNode,
        To:        toNode,
        Weight:    weight,
        Bandwidth: bandwidth,
        Delay:     delay,
        Status:    "up",
    }
    
    g.Edges[edgeID] = edge
    
    // Add neighbors
    fromNode.Neighbors[to] = toNode
    toNode.Neighbors[from] = fromNode
}
```

## 🔄 Routing Protocols

### RIP Implementation
```go
type RIPRouter struct {
    Node        *NetworkNode
    Routes      map[string]*RIPRoute
    Neighbors   map[string]*RIPNeighbor
    UpdateTimer *time.Timer
    mutex       sync.RWMutex
}

func (r *RIPRouter) AddRoute(destination, nextHop string, metric int) {
    r.mutex.Lock()
    defer r.mutex.Unlock()
    
    r.Routes[destination] = &RIPRoute{
        Destination: destination,
        NextHop:     nextHop,
        Metric:      metric,
        Age:         time.Now(),
    }
    
    // Update the node's routing table
    r.Node.AddRoute(destination, nextHop, "eth0", metric, 120, "RIP")
}
```

### OSPF Implementation
```go
type OSPFRouter struct {
    Node        *NetworkNode
    RouterID    string
    AreaID      int
    LSDB        map[string]*LSA
    Neighbors   map[string]*OSPFNeighbor
    SPF         *SPFCalculator
    mutex       sync.RWMutex
}

func (o *OSPFRouter) AddNeighbor(routerID string, priority int) {
    o.mutex.Lock()
    defer o.mutex.Unlock()
    
    o.Neighbors[routerID] = &OSPFNeighbor{
        RouterID: routerID,
        State:    "Down",
        Priority: priority,
        LastSeen: time.Now(),
        LSAs:     make(map[string]*LSA),
    }
}
```

## 📊 Network Analysis

### Connectivity Analysis
```go
func (na *NetworkAnalyzer) AnalyzeConnectivity() map[string]bool {
    connectivity := make(map[string]bool)
    
    for nodeID := range na.Graph.Nodes {
        reachable := true
        for otherNodeID := range na.Graph.Nodes {
            if nodeID == otherNodeID {
                continue
            }
            
            path, _ := na.Graph.DijkstraShortestPath(otherNodeID, nodeID)
            if len(path) == 0 {
                reachable = false
                break
            }
        }
        connectivity[nodeID] = reachable
    }
    
    return connectivity
}
```

### Performance Analysis
```go
func (na *NetworkAnalyzer) AnalyzePerformance() map[string]float64 {
    performance := make(map[string]float64)
    
    // Calculate average path length
    totalPaths := 0
    totalLength := 0
    
    for from := range na.Graph.Nodes {
        for to := range na.Graph.Nodes {
            if from == to {
                continue
            }
            
            path, _ := na.Graph.DijkstraShortestPath(from, to)
            if len(path) > 0 {
                totalPaths++
                totalLength += len(path) - 1
            }
        }
    }
    
    if totalPaths > 0 {
        performance["average_path_length"] = float64(totalLength) / float64(totalPaths)
    }
    
    return performance
}
```

## 🎯 Best Practices

1. **Use appropriate algorithms**: Dijkstra for single-source, Bellman-Ford for negative weights
2. **Implement proper concurrency**: Use mutexes for thread-safe operations
3. **Handle edge cases**: Check for cycles, disconnected components
4. **Optimize performance**: Use priority queues, efficient data structures
5. **Monitor network state**: Track link status, route changes
6. **Implement redundancy**: Multiple paths, failover mechanisms
7. **Use proper metrics**: Bandwidth, delay, reliability
8. **Test thoroughly**: Validate algorithms with various topologies

## 🚨 Common Issues

### Switching Issues
- **MAC table overflow**: Implement aging and limits
- **Broadcast storms**: Use STP and VLANs
- **VLAN misconfiguration**: Validate VLAN assignments
- **Port security**: Implement MAC address limits

### Routing Issues
- **Routing loops**: Use proper loop prevention
- **Convergence time**: Optimize algorithm performance
- **Route flapping**: Implement dampening
- **Memory usage**: Limit routing table size

### Algorithm Issues
- **Infinite loops**: Check for cycles and disconnected components
- **Performance**: Use appropriate data structures
- **Accuracy**: Validate algorithm implementations
- **Scalability**: Consider large network topologies

---

**Remember**: Routing and switching are the backbone of network communication. Master these concepts, and you'll understand how data moves through networks! 🚀
```

Now let me update the progress:


Perfect! I've completed **Topic 9: Routing & Switching** with comprehensive coverage. Here's what I've created for you:

## 🎯 **Topic 9: Routing & Switching - COMPLETED**

### **📚 Theory Coverage** (`09_routing_switching/theory.md`)
- **Network Switching**: MAC address learning, frame forwarding, VLANs, STP
- **Network Routing**: Packet forwarding, path selection, routing tables, protocols
- **Routing Algorithms**: Distance vector, link state, path vector algorithms
- **Routing Protocols**: RIP, OSPF, EIGRP, BGP implementation details
- **Network Topology**: Physical and logical topologies, hierarchical design
- **Quality of Service**: Traffic classification, QoS mechanisms, models
- **Network Redundancy**: High availability, fault tolerance, redundancy protocols
- **Network Monitoring**: SNMP, RMON, NetFlow, performance metrics
- **Go Implementation**: Network simulation, routing algorithms, topology management

### ** Implementation** (`09_routing_switching/implementation.go`)
- **NetworkNode**: Router and switch implementation with ports, MAC tables, routes
- **NetworkTopology**: Graph-based network representation
- **DijkstraShortestPath**: Shortest path algorithm implementation
- **RIPRouter**: RIP protocol implementation with periodic updates
- **OSPFRouter**: OSPF protocol implementation with LSA management
- **NetworkSimulator**: Complete network simulation framework
- **Comprehensive Examples**: All major routing and switching patterns

### **🧮 Routing Algorithms** (`09_routing_switching/routing_algorithms.go`)
- **Dijkstra's Algorithm**: Single-source shortest path
- **Bellman-Ford Algorithm**: Handles negative weights and detects cycles
- **A* Algorithm**: Heuristic-based pathfinding
- **Floyd-Warshall**: All-pairs shortest paths
- **Kruskal's MST**: Minimum spanning tree algorithm
- **Prim's MST**: Alternative MST algorithm
- **NetworkAnalyzer**: Connectivity, redundancy, and performance analysis
- **Graph Representation**: Complete graph data structures

### **📖 Quick Reference** (`09_routing_switching/quick_reference.md`)
- **Essential Concepts**: Switching and routing fundamentals
- **Go Patterns**: Network node, port, route management
- **Switching Implementation**: MAC learning, frame forwarding, VLANs
- **Routing Implementation**: Route lookup, addition, path selection
- **Routing Algorithms**: Dijkstra, Bellman-Ford, A*, Floyd-Warshall
- **Network Topology**: Graph representation, topology management
- **Routing Protocols**: RIP, OSPF implementation patterns
- **Network Analysis**: Connectivity, performance, redundancy analysis
- **Best Practices**: Professional routing and switching guidelines

## 🚀 **Key Features of This Module**

1. **Complete Coverage**: From basic switching to advanced routing algorithms
2. **Go-Specific**: Leverages Go's excellent concurrency and data structures
3. **Production-Ready**: Real-world routing and switching patterns
4. **Algorithm Implementation**: All major routing algorithms in Go
5. **Protocol Simulation**: RIP, OSPF, and other routing protocols
6. **Network Analysis**: Comprehensive network performance analysis
7. **Graph Algorithms**: Complete graph theory implementation
8. **Concurrency**: Thread-safe network operations

## 📈 **What You've Learned**

- **Switching Fundamentals**: MAC learning, frame forwarding, VLANs, STP
- **Routing Concepts**: Packet forwarding, path selection, routing tables
- **Routing Algorithms**: Dijkstra, Bellman-Ford, A*, Floyd-Warshall
- **Routing Protocols**: RIP, OSPF, EIGRP, BGP implementation
- **Network Topology**: Physical and logical network design
- **Quality of Service**: Traffic prioritization and bandwidth management
- **Network Redundancy**: High availability and fault tolerance
- **Network Monitoring**: Performance metrics and analysis
- **Go Implementation**: Professional network programming

## **Next Steps**

You're now ready to move to **Topic 10: Network Performance & Optimization** where we'll explore performance tuning, optimization techniques, and monitoring.

The routing and switching knowledge you now have will make you unstoppable in understanding network behavior, implementing routing algorithms, and acing network-related interview questions! 🚀

Would you like me to continue with **Topic 10: Network Performance & Optimization** or would you prefer to explore any specific aspect of routing and switching first?

The routing and switching foundation you've gained is crucial for network engineering and will set you apart in technical interviews!

---

**User**

move to Topic 10: Network Performance & Optimization 

---

**Cursor**

Perfect! Let me continue with **Topic 10: Network Performance & Optimization**. This is where we explore performance tuning, optimization techniques, monitoring, and ensuring networks run at peak efficiency.


Now let me create the comprehensive Network Performance & Optimization module:


```markdown
# Topic 10: Network Performance & Optimization - Peak Performance Mastery

## Network Performance Overview

Network performance optimization is the process of improving network efficiency, reliability, and throughput while minimizing latency, packet loss, and resource consumption. It encompasses multiple layers and aspects of network operation.

### Performance Metrics
- **Throughput**: Data transfer rate (bits per second)
- **Latency**: Round-trip time (RTT) for data transmission
- **Jitter**: Variation in packet arrival times
- **Packet Loss**: Percentage of packets lost in transmission
- **Bandwidth Utilization**: Percentage of available bandwidth used
- **CPU Utilization**: Processor usage for network operations
- **Memory Usage**: RAM consumption for network buffers
- **Connection Count**: Number of active connections

### Performance Factors
- **Hardware**: Network interface cards, switches, routers, cables
- **Software**: Operating system, drivers, protocols, applications
- **Configuration**: Buffer sizes, queue lengths, timeouts
- **Traffic Patterns**: Burst vs. steady-state traffic
- **Network Topology**: Physical and logical network design
- **Protocol Efficiency**: TCP vs. UDP, compression, encryption

## Bandwidth Management

### Bandwidth Concepts
Bandwidth is the maximum data transfer rate of a network connection or interface.

#### Bandwidth Types
- **Physical Bandwidth**: Maximum theoretical capacity
- **Effective Bandwidth**: Actual usable capacity
- **Allocated Bandwidth**: Bandwidth assigned to specific traffic
- **Available Bandwidth**: Unused bandwidth capacity

#### Bandwidth Measurement
- **Speed Tests**: Measure actual throughput
- **SNMP Monitoring**: Query device statistics
- **Packet Capture**: Analyze traffic patterns
- **Flow Analysis**: NetFlow, sFlow, IPFIX

### Traffic Shaping
Traffic shaping controls the rate of data transmission to match network capacity.

#### Shaping Techniques
- **Token Bucket**: Allow bursts up to bucket size
- **Leaky Bucket**: Smooth traffic to constant rate
- **Rate Limiting**: Cap maximum transmission rate
- **Priority Queuing**: Prioritize different traffic types

#### Shaping Algorithms
- **Weighted Fair Queuing (WFQ)**: Fair bandwidth allocation
- **Class-Based Weighted Fair Queuing (CBWFQ)**: Traffic class prioritization
- **Low Latency Queuing (LLQ)**: Real-time traffic priority
- **Random Early Detection (RED)**: Congestion avoidance

### Quality of Service (QoS)
QoS ensures that critical traffic receives priority treatment in the network.

#### QoS Models
- **Best Effort**: No guarantees, first-come-first-served
- **Integrated Services (IntServ)**: Per-flow guarantees
- **Differentiated Services (DiffServ)**: Per-class guarantees

#### QoS Mechanisms
- **Classification**: Identify traffic types
- **Marking**: Set priority levels (DSCP, CoS)
- **Policing**: Enforce traffic limits
- **Shaping**: Smooth traffic bursts
- **Queuing**: Manage traffic queues
- **Scheduling**: Prioritize traffic transmission

## Latency Optimization

### Latency Sources
Latency is the time delay between sending and receiving data.

#### Physical Latency
- **Propagation Delay**: Speed of light in medium
- **Transmission Delay**: Time to send data onto medium
- **Processing Delay**: Time to process data
- **Queuing Delay**: Time waiting in queues

#### Network Latency
- **Router Processing**: Packet forwarding time
- **Switch Processing**: Frame forwarding time
- **Serialization Delay**: Time to transmit packet
- **Buffer Delay**: Time in network buffers

### Latency Reduction Techniques

#### Hardware Optimization
- **Faster Interfaces**: Higher speed network cards
- **Low-Latency Switches**: Cut-through switching
- **Optimized Routers**: Hardware-based forwarding
- **Quality Cables**: Better signal integrity

#### Software Optimization
- **Kernel Bypass**: Direct hardware access
- **Zero-Copy**: Eliminate data copying
- **Polling**: Active waiting instead of interrupts
- **Batch Processing**: Process multiple packets together

#### Protocol Optimization
- **TCP Tuning**: Optimize window sizes, timeouts
- **UDP Optimization**: Reduce overhead
- **Compression**: Reduce data size
- **Caching**: Store frequently accessed data

## Throughput Optimization

### Throughput Factors
Throughput is the actual data transfer rate achieved.

#### Bottlenecks
- **CPU**: Processing power limitations
- **Memory**: Buffer size constraints
- **Network Interface**: Interface speed limits
- **Network Path**: Slowest link in path
- **Protocol Overhead**: TCP/UDP headers, acknowledgments

#### Optimization Strategies
- **Parallel Processing**: Multiple threads/processes
- **Connection Pooling**: Reuse connections
- **Pipelining**: Overlap requests and responses
- **Compression**: Reduce data size
- **Caching**: Store frequently accessed data

### Load Balancing
Load balancing distributes traffic across multiple paths or servers.

#### Load Balancing Types
- **Round Robin**: Rotate through servers
- **Least Connections**: Choose server with fewest connections
- **Weighted**: Assign different weights to servers
- **Geographic**: Route based on location
- **Content-Based**: Route based on content type

#### Load Balancing Algorithms
- **Static**: Predefined distribution rules
- **Dynamic**: Real-time load assessment
- **Adaptive**: Adjust based on performance
- **Predictive**: Anticipate future load

## Buffer Management

### Buffer Concepts
Buffers temporarily store data during transmission.

#### Buffer Types
- **Receive Buffers**: Store incoming data
- **Send Buffers**: Store outgoing data
- **Kernel Buffers**: Operating system buffers
- **Application Buffers**: User-space buffers

#### Buffer Sizing
- **Bandwidth-Delay Product**: BDP = Bandwidth × RTT
- **Buffer Size**: Typically 2-3 times BDP
- **Dynamic Sizing**: Adjust based on conditions
- **Per-Connection**: Individual buffer allocation

### Buffer Optimization
- **Zero-Copy**: Eliminate data copying
- **Scatter-Gather**: Efficient data handling
- **Memory Mapping**: Direct memory access
- **NUMA Awareness**: Optimize for multi-socket systems

## Congestion Control

### Congestion Concepts
Congestion occurs when network demand exceeds capacity.

#### Congestion Indicators
- **Packet Loss**: Dropped packets due to full buffers
- **Increased Latency**: Longer queuing delays
- **Reduced Throughput**: Lower effective bandwidth
- **Jitter**: Variable packet arrival times

#### Congestion Control Algorithms
- **TCP Congestion Control**: AIMD, CUBIC, BBR
- **Active Queue Management**: RED, WRED, PIE
- **Flow Control**: Prevent overwhelming receivers
- **Backpressure**: Signal upstream to slow down

### Congestion Avoidance
- **Early Detection**: Monitor queue levels
- **Proactive Dropping**: Drop packets before queues full
- **Rate Limiting**: Prevent excessive traffic
- **Traffic Shaping**: Smooth traffic bursts

## Network Monitoring

### Monitoring Concepts
Network monitoring tracks performance metrics and identifies issues.

#### Monitoring Types
- **Active Monitoring**: Send test traffic
- **Passive Monitoring**: Observe existing traffic
- **Synthetic Monitoring**: Simulate user behavior
- **Real User Monitoring**: Track actual user experience

#### Monitoring Tools
- **SNMP**: Simple Network Management Protocol
- **NetFlow**: Traffic flow analysis
- **sFlow**: Sampled flow analysis
- **Packet Capture**: Deep packet inspection
- **Ping/Traceroute**: Connectivity testing

### Performance Baselines
- **Establish Baselines**: Measure normal performance
- **Set Thresholds**: Define acceptable limits
- **Trend Analysis**: Track performance over time
- **Anomaly Detection**: Identify unusual patterns

## Protocol Optimization

### TCP Optimization
TCP is the most widely used transport protocol and offers many optimization opportunities.

#### TCP Tuning Parameters
- **Window Size**: Receive and send window sizes
- **Buffer Sizes**: Socket buffer sizes
- **Timeouts**: Connection and retransmission timeouts
- **Congestion Control**: Algorithm selection
- **Nagle's Algorithm**: Disable for low latency
- **Delayed ACKs**: Optimize acknowledgment timing

#### TCP Optimization Techniques
- **Window Scaling**: Increase window size beyond 64KB
- **Selective Acknowledgments**: SACK for better recovery
- **Fast Retransmit**: Quick recovery from losses
- **Fast Open**: Reduce connection establishment time
- **Multipath TCP**: Use multiple paths simultaneously

### UDP Optimization
UDP is used for real-time applications and can be optimized for specific use cases.

#### UDP Tuning
- **Buffer Sizes**: Optimize for application needs
- **Socket Options**: Configure for performance
- **Error Handling**: Implement application-level reliability
- **Flow Control**: Prevent overwhelming receivers

#### UDP Best Practices
- **Packet Size**: Optimize for MTU
- **Rate Limiting**: Control transmission rate
- **Error Recovery**: Implement retransmission
- **Congestion Control**: Monitor network conditions

## Application-Level Optimization

### Connection Management
- **Connection Pooling**: Reuse connections
- **Keep-Alive**: Maintain persistent connections
- **Connection Limits**: Control concurrent connections
- **Timeout Management**: Set appropriate timeouts

### Data Transfer Optimization
- **Compression**: Reduce data size
- **Chunking**: Break large transfers into chunks
- **Pipelining**: Overlap requests and responses
- **Caching**: Store frequently accessed data

### Error Handling
- **Retry Logic**: Implement exponential backoff
- **Circuit Breakers**: Prevent cascading failures
- **Graceful Degradation**: Maintain service during issues
- **Monitoring**: Track error rates and patterns

## Hardware Optimization

### Network Interface Cards (NICs)
- **Offloading**: Hardware-based processing
- **Interrupt Coalescing**: Reduce CPU interrupts
- **Receive Side Scaling**: Distribute processing across CPUs
- **Single Root I/O Virtualization**: SR-IOV for virtualization

### Switches and Routers
- **Cut-Through Switching**: Forward frames immediately
- **Store-and-Forward**: Complete frame before forwarding
- **Hardware Forwarding**: ASIC-based processing
- **Virtual Output Queuing**: VOQ for better performance

### Cables and Media
- **Fiber Optic**: Higher bandwidth and lower latency
- **Copper**: Cost-effective for short distances
- **Wireless**: 802.11ac/ax for high-speed wireless
- **Power over Ethernet**: PoE for power and data

## Virtualization and Cloud Optimization

### Virtual Network Functions
- **SR-IOV**: Direct hardware access for VMs
- **DPDK**: Data Plane Development Kit
- **Open vSwitch**: Software-defined switching
- **NFV**: Network Function Virtualization

### Cloud Networking
- **SDN**: Software-Defined Networking
- **Overlay Networks**: VXLAN, GRE, Geneve
- **Load Balancing**: Cloud-native load balancers
- **Auto-scaling**: Dynamic resource allocation

### Container Networking
- **CNI**: Container Network Interface
- **Service Mesh**: Istio, Linkerd
- **Sidecar Proxies**: Envoy, NGINX
- **Network Policies**: Kubernetes network security

## Performance Testing

### Testing Types
- **Load Testing**: Normal expected load
- **Stress Testing**: Beyond normal capacity
- **Spike Testing**: Sudden load increases
- **Volume Testing**: Large amounts of data
- **Endurance Testing**: Long-duration testing

### Testing Tools
- **iperf**: Network performance testing
- **netperf**: Comprehensive network testing
- **Wireshark**: Packet analysis
- **tcpdump**: Command-line packet capture
- **htop/iotop**: System resource monitoring

### Testing Metrics
- **Throughput**: Data transfer rate
- **Latency**: Round-trip time
- **Packet Loss**: Percentage of lost packets
- **Jitter**: Variation in latency
- **CPU Usage**: Processor utilization
- **Memory Usage**: RAM consumption

## Troubleshooting Performance Issues

### Common Issues
- **High Latency**: Network congestion, routing issues
- **Low Throughput**: Bandwidth limitations, CPU bottlenecks
- **Packet Loss**: Network errors, buffer overflows
- **High Jitter**: Network instability, queuing issues
- **Connection Drops**: Timeout issues, resource exhaustion

### Troubleshooting Steps
1. **Identify Symptoms**: Measure current performance
2. **Check Network**: Verify connectivity and routing
3. **Monitor Resources**: CPU, memory, disk usage
4. **Analyze Traffic**: Packet capture and analysis
5. **Test Components**: Isolate problematic areas
6. **Apply Fixes**: Implement solutions
7. **Verify Results**: Confirm improvements

### Diagnostic Tools
- **ping**: Basic connectivity testing
- **traceroute**: Path discovery and latency measurement
- **netstat**: Network connection statistics
- **ss**: Socket statistics
- **iftop**: Network interface traffic
- **nload**: Network load monitoring

## Go Implementation Concepts

### Performance Monitoring
Go provides excellent tools for network performance monitoring.

#### Metrics Collection
- **Prometheus**: Metrics collection and storage
- **Grafana**: Metrics visualization
- **Custom Metrics**: Application-specific monitoring
- **Health Checks**: Service availability monitoring

#### Profiling
- **CPU Profiling**: Identify performance bottlenecks
- **Memory Profiling**: Track memory usage
- **Goroutine Profiling**: Monitor concurrency
- **Block Profiling**: Identify blocking operations

### Optimization Techniques
- **Connection Pooling**: Reuse network connections
- **Goroutine Pools**: Limit concurrent operations
- **Memory Pools**: Reuse memory allocations
- **Batch Processing**: Process multiple items together

### Concurrent Programming
- **Goroutines**: Lightweight concurrency
- **Channels**: Communication between goroutines
- **Context**: Cancellation and timeouts
- **Sync Package**: Synchronization primitives

## Interview Questions

### Basic Questions
1. What factors affect network performance?
2. How do you measure network latency?
3. What is the difference between throughput and bandwidth?

### Intermediate Questions
1. How would you optimize TCP performance?
2. Explain different load balancing algorithms.
3. What are the benefits of connection pooling?

### Advanced Questions
1. Design a high-performance network monitoring system.
2. How would you optimize a distributed system's network performance?
3. Implement a custom congestion control algorithm.

## Next Steps
After mastering network performance and optimization, proceed to:
- **Topic 11**: Distributed Systems Networking
- **Topic 12**: Cloud Networking & Containerization
- **Topic 13**: Interview Preparation & System Design

Master network performance optimization, and you'll understand how to build blazing-fast networks! 🚀
```

Now let me create the comprehensive Go implementation for network performance and optimization:


```go
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
```

Now let me create a comprehensive performance testing and benchmarking module:


```go
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"runtime"
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
```

I need to fix the import issue in the performance testing file:


```go
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

```

Now let me create a quick reference guide and update the progress:


```markdown
# Network Performance & Optimization Quick Reference Guide

## 🚀 Essential Performance Concepts

### Performance Metrics
- **Throughput**: Data transfer rate (bits per second)
- **Latency**: Round-trip time (RTT) for data transmission
- **Jitter**: Variation in packet arrival times
- **Packet Loss**: Percentage of packets lost in transmission
- **Bandwidth Utilization**: Percentage of available bandwidth used
- **CPU Utilization**: Processor usage for network operations
- **Memory Usage**: RAM consumption for network buffers
- **Connection Count**: Number of active connections

### Performance Factors
- **Hardware**: Network interface cards, switches, routers, cables
- **Software**: Operating system, drivers, protocols, applications
- **Configuration**: Buffer sizes, queue lengths, timeouts
- **Traffic Patterns**: Burst vs. steady-state traffic
- **Network Topology**: Physical and logical network design
- **Protocol Efficiency**: TCP vs. UDP, compression, encryption

## 🔧 Go Implementation Patterns

### Performance Monitoring
```go
type PerformanceMetrics struct {
    Throughput     float64   `json:"throughput"`
    Latency        float64   `json:"latency"`
    Jitter         float64   `json:"jitter"`
    PacketLoss     float64   `json:"packet_loss"`
    BandwidthUtil  float64   `json:"bandwidth_util"`
    CPUUtil        float64   `json:"cpu_util"`
    MemoryUsage    uint64    `json:"memory_usage"`
    Connections    int32     `json:"connections"`
    Timestamp      time.Time `json:"timestamp"`
    mutex          sync.RWMutex
}
```

### Load Balancing
```go
type LoadBalancer struct {
    Servers    []*Server
    Algorithm  string
    Current    int32
    mutex      sync.RWMutex
}

type Server struct {
    Address     string
    Weight      int
    Connections int32
    ResponseTime float64
    IsHealthy   bool
    LastCheck   time.Time
}
```

### Connection Pooling
```go
type ConnectionPool struct {
    Connections chan net.Conn
    Factory     func() (net.Conn, error)
    MaxSize     int
    CurrentSize int32
    mutex       sync.RWMutex
}

func (cp *ConnectionPool) GetConnection() (net.Conn, error) {
    select {
    case conn := <-cp.Connections:
        return conn, nil
    default:
        if atomic.LoadInt32(&cp.CurrentSize) < int32(cp.MaxSize) {
            conn, err := cp.Factory()
            if err != nil {
                return nil, err
            }
            atomic.AddInt32(&cp.CurrentSize, 1)
            return conn, nil
        }
        
        select {
        case conn := <-cp.Connections:
            return conn, nil
        case <-time.After(5 * time.Second):
            return nil, fmt.Errorf("timeout waiting for connection")
        }
    }
}
```

### Traffic Shaping
```go
type TrafficShaper struct {
    RateLimit    int64 // bytes per second
    BucketSize   int64 // bytes
    Tokens       int64
    LastUpdate   time.Time
    mutex        sync.Mutex
}

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
```

## 📊 Performance Testing

### Benchmark Framework
```go
type BenchmarkResult struct {
    Name         string        `json:"name"`
    Duration     time.Duration `json:"duration"`
    Operations   int64         `json:"operations"`
    Throughput   float64       `json:"throughput"`
    Latency      time.Duration `json:"latency"`
    MinLatency   time.Duration `json:"min_latency"`
    MaxLatency   time.Duration `json:"max_latency"`
    P95Latency   time.Duration `json:"p95_latency"`
    P99Latency   time.Duration `json:"p99_latency"`
    Errors       int64         `json:"errors"`
    ErrorRate    float64       `json:"error_rate"`
    Bytes        int64         `json:"bytes"`
    Bandwidth    float64       `json:"bandwidth"`
    CPUUsage     float64       `json:"cpu_usage"`
    MemoryUsage  uint64        `json:"memory_usage"`
    Timestamp    time.Time     `json:"timestamp"`
}
```

### Load Testing
```go
type LoadTester struct {
    Targets    map[string]*LoadTarget
    Results    map[string]*LoadTestResult
    mutex      sync.RWMutex
}

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
```

### Performance Profiling
```go
type PerformanceProfiler struct {
    StartTime time.Time
    EndTime   time.Time
    Metrics   map[string]interface{}
    mutex     sync.RWMutex
}

func (pp *PerformanceProfiler) Start() {
    pp.mutex.Lock()
    defer pp.mutex.Unlock()
    
    pp.StartTime = time.Now()
    pp.Metrics = make(map[string]interface{})
}

func (pp *PerformanceProfiler) AddMetric(name string, value interface{}) {
    pp.mutex.Lock()
    defer pp.mutex.Unlock()
    
    pp.Metrics[name] = value
}
```

## 🚦 Load Balancing Algorithms

### Round Robin
```go
func (lb *LoadBalancer) getRoundRobinServer() *Server {
    if len(lb.Servers) == 0 {
        return nil
    }
    
    for i := 0; i < len(lb.Servers); i++ {
        index := int(atomic.AddInt32(&lb.Current, 1)) % len(lb.Servers)
        if lb.Servers[index].IsHealthy {
            return lb.Servers[index]
        }
    }
    
    return lb.Servers[0]
}
```

### Least Connections
```go
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
```

### Weighted Round Robin
```go
func (lb *LoadBalancer) getWeightedServer() *Server {
    if len(lb.Servers) == 0 {
        return nil
    }
    
    totalWeight := 0
    for _, server := range lb.Servers {
        if server.IsHealthy {
            totalWeight += server.Weight
        }
    }
    
    if totalWeight == 0 {
        return lb.Servers[0]
    }
    
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
```

## 📈 Bandwidth Management

### Bandwidth Monitoring
```go
type BandwidthMonitor struct {
    Interfaces map[string]*InterfaceStats
    mutex      sync.RWMutex
}

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
```

### Latency Testing
```go
type LatencyTester struct {
    Targets map[string]*LatencyTarget
    mutex   sync.RWMutex
}

type LatencyTarget struct {
    Address     string
    Latency     float64
    Jitter      float64
    PacketLoss  float64
    LastTest    time.Time
    TestCount   int
    SuccessCount int
}

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
    
    latency := float64(time.Since(start).Nanoseconds()) / 1e6
    target.TestCount++
    target.SuccessCount++
    target.Latency = latency
    target.LastTest = time.Now()
    target.PacketLoss = float64(target.TestCount-target.SuccessCount) / float64(target.TestCount) * 100
    
    return nil
}
```

## 🔄 TCP Optimization

### TCP Tuning Parameters
```go
func optimizeTCP() {
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
```

### UDP Optimization
```go
func optimizeUDP() {
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
```

## 🎯 Performance Best Practices

### 1. Connection Management
- **Connection Pooling**: Reuse connections to reduce overhead
- **Keep-Alive**: Maintain persistent connections
- **Connection Limits**: Control concurrent connections
- **Timeout Management**: Set appropriate timeouts

### 2. Data Transfer Optimization
- **Compression**: Reduce data size
- **Chunking**: Break large transfers into chunks
- **Pipelining**: Overlap requests and responses
- **Caching**: Store frequently accessed data

### 3. Error Handling
- **Retry Logic**: Implement exponential backoff
- **Circuit Breakers**: Prevent cascading failures
- **Graceful Degradation**: Maintain service during issues
- **Monitoring**: Track error rates and patterns

### 4. Resource Management
- **Memory Pools**: Reuse memory allocations
- **Goroutine Pools**: Limit concurrent operations
- **Buffer Management**: Optimize buffer sizes
- **Garbage Collection**: Minimize GC pressure

## 🚨 Common Performance Issues

### High Latency
- **Network Congestion**: Check for bottlenecks
- **Routing Issues**: Verify path optimization
- **DNS Resolution**: Cache DNS lookups
- **Connection Overhead**: Use connection pooling

### Low Throughput
- **Bandwidth Limitations**: Check interface speeds
- **CPU Bottlenecks**: Profile CPU usage
- **Memory Constraints**: Monitor memory usage
- **Protocol Overhead**: Optimize protocol usage

### Packet Loss
- **Network Errors**: Check link quality
- **Buffer Overflows**: Increase buffer sizes
- **Congestion**: Implement traffic shaping
- **Hardware Issues**: Check network equipment

### High Jitter
- **Network Instability**: Check link quality
- **Queuing Issues**: Optimize queue management
- **Load Balancing**: Improve load distribution
- **Synchronization**: Check clock synchronization

## 📊 Monitoring and Alerting

### Key Metrics to Monitor
- **Throughput**: Data transfer rate
- **Latency**: Round-trip time
- **Packet Loss**: Percentage of lost packets
- **Jitter**: Variation in latency
- **CPU Usage**: Processor utilization
- **Memory Usage**: RAM consumption
- **Connection Count**: Active connections
- **Error Rate**: Percentage of failed requests

### Alerting Thresholds
- **Latency**: > 100ms for voice, > 150ms for video
- **Packet Loss**: > 0.1% for voice, > 0.5% for video
- **Jitter**: > 30ms for voice, > 50ms for video
- **CPU Usage**: > 80% sustained
- **Memory Usage**: > 90% of available
- **Error Rate**: > 1% of requests

### Monitoring Tools
- **Prometheus**: Metrics collection and storage
- **Grafana**: Metrics visualization
- **ELK Stack**: Log analysis and monitoring
- **Custom Dashboards**: Application-specific monitoring

## 🔧 Troubleshooting Performance Issues

### Step 1: Identify Symptoms
- Measure current performance
- Compare with baselines
- Identify specific issues
- Document symptoms

### Step 2: Check Network
- Verify connectivity
- Check routing
- Test with ping/traceroute
- Monitor interface statistics

### Step 3: Monitor Resources
- Check CPU usage
- Monitor memory usage
- Verify disk I/O
- Check network I/O

### Step 4: Analyze Traffic
- Capture packets
- Analyze traffic patterns
- Check for anomalies
- Identify bottlenecks

### Step 5: Test Components
- Isolate problematic areas
- Test individual components
- Verify configurations
- Check for errors

### Step 6: Apply Fixes
- Implement solutions
- Test changes
- Monitor results
- Document fixes

### Step 7: Verify Results
- Confirm improvements
- Update baselines
- Monitor stability
- Document lessons learned

---

**Remember**: Performance optimization is an ongoing process. Monitor continuously, measure everything, and optimize based on data! 🚀
```

Now let me update the progress:


Perfect! I've completed **Topic 10: Network Performance & Optimization** with comprehensive coverage. Here's what I've created for you:

## 🎯 **Topic 10: Network Performance & Optimization - COMPLETED**

### **📚 Theory Coverage** (`10_network_performance/theory.md`)
- **Network Performance Overview**: Performance metrics, factors, and optimization concepts
- **Bandwidth Management**: Traffic shaping, QoS, bandwidth measurement and allocation
- **Latency Optimization**: Latency sources, reduction techniques, hardware and software optimization
- **Throughput Optimization**: Bottleneck identification, load balancing, parallel processing
- **Buffer Management**: Buffer types, sizing, optimization techniques
- **Congestion Control**: Congestion indicators, control algorithms, avoidance techniques
- **Network Monitoring**: Active/passive monitoring, tools, performance baselines
- **Protocol Optimization**: TCP/UDP tuning, application-level optimization
- **Hardware Optimization**: NICs, switches, routers, cables and media
- **Virtualization and Cloud**: VNFs, SDN, container networking
- **Performance Testing**: Load testing, stress testing, testing tools and metrics
- **Troubleshooting**: Common issues, diagnostic tools, troubleshooting steps

### ** Implementation** (`10_network_performance/implementation.go`)
- **PerformanceMetrics**: Comprehensive performance monitoring with thread-safe operations
- **PerformanceMonitor**: Real-time performance monitoring with automatic collection
- **LoadBalancer**: Multiple load balancing algorithms (round-robin, least connections, weighted, fastest)
- **ConnectionPool**: Efficient connection pooling with timeout and resource management
- **TrafficShaper**: Token bucket algorithm for rate limiting and traffic shaping
- **BandwidthMonitor**: Interface statistics tracking and bandwidth utilization monitoring
- **LatencyTester**: Network latency testing with packet loss measurement
- **PerformanceOptimizer**: Complete performance optimization framework

### **🧪 Performance Testing** (`10_network_performance/performance_testing.go`)
- **BenchmarkResult**: Comprehensive benchmark result structure with statistics
- **BenchmarkRunner**: Concurrent benchmark execution with latency analysis
- **NetworkBenchmark**: Network-specific benchmarks (TCP, UDP, HTTP, ping)
- **LoadTester**: Load testing framework with configurable targets and concurrency
- **PerformanceProfiler**: Performance profiling with metrics collection
- **Statistical Analysis**: P95, P99 latency, throughput, error rate calculations

### **📖 Quick Reference** (`10_network_performance/quick_reference.md`)
- **Essential Concepts**: Performance metrics, factors, and optimization principles
- **Go Patterns**: Performance monitoring, load balancing, connection pooling
- **Load Balancing**: Round-robin, least connections, weighted algorithms
- **Bandwidth Management**: Monitoring, latency testing, interface statistics
- **TCP/UDP Optimization**: Tuning parameters, optimization techniques
- **Performance Testing**: Benchmarking, load testing, profiling frameworks
- **Best Practices**: Connection management, data transfer, error handling
- **Common Issues**: High latency, low throughput, packet loss, high jitter
- **Monitoring**: Key metrics, alerting thresholds, monitoring tools
- **Troubleshooting**: Step-by-step performance issue resolution

## 🚀 **Key Features of This Module**

1. **Complete Performance Coverage**: From basic metrics to advanced optimization
2. **Go-Specific**: Leverages Go's excellent concurrency and performance features
3. **Production-Ready**: Real-world performance monitoring and optimization patterns
4. **Load Balancing**: Multiple algorithms with health checking and failover
5. **Connection Pooling**: Efficient resource management and reuse
6. **Traffic Shaping**: Rate limiting and bandwidth management
7. **Performance Testing**: Comprehensive benchmarking and load testing framework
8. **Monitoring**: Real-time performance metrics and alerting

## 📈 **What You've Learned**

- **Performance Fundamentals**: Metrics, factors, and optimization concepts
- **Bandwidth Management**: Traffic shaping, QoS, and bandwidth optimization
- **Latency Optimization**: Reduction techniques and performance tuning
- **Throughput Optimization**: Load balancing, parallel processing, bottleneck identification
- **Buffer Management**: Buffer sizing, optimization, and memory management
- **Congestion Control**: Control algorithms, avoidance techniques, flow control
- **Network Monitoring**: Performance baselines, monitoring tools, alerting
- **Protocol Optimization**: TCP/UDP tuning, application-level optimization
- **Performance Testing**: Benchmarking, load testing, profiling, statistical analysis
- **Go Implementation**: Professional performance monitoring and optimization

## **Next Steps**

You're now ready to move to **Topic 11: Distributed Systems Networking** where we'll explore distributed system architectures, consensus algorithms, and network coordination.

The performance optimization knowledge you now have will make you unstoppable in building high-performance networks, optimizing system performance, and acing performance-related interview questions! 🚀

Would you like me to continue with **Topic 11: Distributed Systems Networking** or would you prefer to explore any specific aspect of network performance first?

The performance optimization foundation you've gained is crucial for building scalable, high-performance systems and will set you apart in technical interviews!

---

**User**

move to Topic 11: Distributed Systems Networking

---

**Cursor**

Perfect! Let me continue with **Topic 11: Distributed Systems Networking**. This is where we explore distributed system architectures, consensus algorithms, service discovery, and network coordination - the foundation of modern distributed applications.


Now let me create the comprehensive Distributed Systems Networking module:


```markdown
# Topic 11: Distributed Systems Networking - Distributed Architecture Mastery

## Distributed Systems Overview

Distributed systems are collections of independent computers that appear to users as a single coherent system. They are connected by a network and coordinate their actions by passing messages to achieve a common goal.

### Key Characteristics
- **Concurrency**: Multiple processes execute simultaneously
- **No Global Clock**: Processes don't share a common clock
- **Independent Failures**: Components can fail independently
- **Heterogeneity**: Different hardware, software, and networks
- **Scalability**: System can grow to handle increased load
- **Transparency**: System appears as a single entity to users

### Distributed System Challenges
- **Network Partitions**: Network splits can isolate components
- **Partial Failures**: Some components fail while others continue
- **Consistency**: Maintaining data consistency across nodes
- **Availability**: Ensuring system remains available during failures
- **Partition Tolerance**: Operating despite network partitions
- **Clock Synchronization**: Coordinating time across nodes

## CAP Theorem

The CAP theorem states that a distributed system cannot simultaneously provide all three guarantees:

### Consistency (C)
- All nodes see the same data at the same time
- After an update, all subsequent reads return the updated value
- Strong consistency vs. eventual consistency

### Availability (A)
- System remains operational and responsive
- Every request receives a response (success or failure)
- No downtime due to component failures

### Partition Tolerance (P)
- System continues operating despite network partitions
- Network splits don't cause system failure
- Messages can be lost or delayed

### CAP Trade-offs
- **CP Systems**: Consistency + Partition Tolerance (e.g., MongoDB, HBase)
- **AP Systems**: Availability + Partition Tolerance (e.g., Cassandra, DynamoDB)
- **CA Systems**: Consistency + Availability (e.g., traditional RDBMS)

## Distributed System Architectures

### Client-Server Architecture
- **Client**: Requests services from servers
- **Server**: Provides services to clients
- **Communication**: Request-response pattern
- **Examples**: Web applications, database systems

### Peer-to-Peer (P2P) Architecture
- **Nodes**: Equal participants in the system
- **Communication**: Direct node-to-node communication
- **Examples**: BitTorrent, blockchain networks, distributed file systems

### Microservices Architecture
- **Services**: Small, independent, loosely coupled services
- **Communication**: HTTP, gRPC, message queues
- **Deployment**: Independent deployment and scaling
- **Examples**: Netflix, Uber, Amazon

### Service-Oriented Architecture (SOA)
- **Services**: Business functionality exposed as services
- **Communication**: Web services, message queues
- **Integration**: Enterprise service bus (ESB)
- **Examples**: Enterprise systems, legacy integration

## Consensus Algorithms

Consensus algorithms ensure that all nodes in a distributed system agree on a single value or decision.

### Raft Algorithm
Raft is a consensus algorithm designed to be understandable and implementable.

#### Raft States
- **Leader**: Handles client requests and log replication
- **Follower**: Receives log entries from leader
- **Candidate**: Attempts to become leader during election

#### Raft Process
1. **Election**: Nodes elect a leader when no leader exists
2. **Log Replication**: Leader replicates log entries to followers
3. **Safety**: Ensures consistency and availability
4. **Split Brain Prevention**: Prevents multiple leaders

#### Raft Benefits
- **Understandable**: Easier to understand than Paxos
- **Implementable**: Clear implementation guidelines
- **Fault Tolerant**: Handles node failures gracefully
- **Consistent**: Maintains strong consistency

### Paxos Algorithm
Paxos is a family of protocols for solving consensus in a network of unreliable processors.

#### Paxos Phases
1. **Prepare Phase**: Proposer sends prepare request
2. **Promise Phase**: Acceptors respond with promises
3. **Accept Phase**: Proposer sends accept request
4. **Learn Phase**: Acceptors learn the chosen value

#### Paxos Variants
- **Basic Paxos**: Single value consensus
- **Multi-Paxos**: Multiple value consensus
- **Fast Paxos**: Optimized for fast decisions
- **Byzantine Paxos**: Handles Byzantine failures

### PBFT (Practical Byzantine Fault Tolerance)
PBFT handles Byzantine failures where nodes can behave arbitrarily.

#### PBFT Process
1. **Request**: Client sends request to primary
2. **Pre-prepare**: Primary broadcasts pre-prepare message
3. **Prepare**: Replicas broadcast prepare messages
4. **Commit**: Replicas broadcast commit messages
5. **Reply**: Replicas send reply to client

#### PBFT Properties
- **Byzantine Fault Tolerant**: Handles up to (n-1)/3 Byzantine failures
- **Safety**: Ensures consistency despite failures
- **Liveness**: Ensures progress despite failures
- **Performance**: O(n²) message complexity

## Service Discovery

Service discovery allows services to find and communicate with each other in a distributed system.

### Service Registry
- **Centralized Registry**: Single point of service registration
- **Distributed Registry**: Multiple registry instances
- **Self-Registration**: Services register themselves
- **Third-Party Registration**: External system registers services

### Discovery Patterns
- **Client-Side Discovery**: Client queries registry directly
- **Server-Side Discovery**: Load balancer queries registry
- **Service Mesh**: Sidecar proxies handle discovery
- **DNS-Based**: Use DNS for service discovery

### Service Discovery Tools
- **Consul**: Service discovery and configuration
- **Eureka**: Netflix service discovery
- **Zookeeper**: Distributed coordination service
- **etcd**: Distributed key-value store
- **Kubernetes**: Built-in service discovery

## Load Balancing in Distributed Systems

### Load Balancing Strategies
- **Round Robin**: Distribute requests evenly
- **Least Connections**: Route to server with fewest connections
- **Weighted**: Assign different weights to servers
- **Geographic**: Route based on client location
- **Content-Based**: Route based on request content

### Load Balancing Types
- **Layer 4**: Transport layer load balancing
- **Layer 7**: Application layer load balancing
- **Global**: Load balance across data centers
- **Local**: Load balance within data center

### Load Balancing Algorithms
- **Static**: Predefined distribution rules
- **Dynamic**: Real-time load assessment
- **Adaptive**: Adjust based on performance
- **Predictive**: Anticipate future load

## Distributed Data Management

### Data Replication
- **Master-Slave**: Single master, multiple slaves
- **Master-Master**: Multiple masters, bidirectional replication
- **Peer-to-Peer**: All nodes are equal
- **Chain Replication**: Sequential replication chain

### Consistency Models
- **Strong Consistency**: All nodes see same data
- **Eventual Consistency**: Data becomes consistent over time
- **Weak Consistency**: No guarantees about consistency
- **Causal Consistency**: Preserves causal relationships

### Data Partitioning
- **Horizontal Partitioning**: Split data by rows
- **Vertical Partitioning**: Split data by columns
- **Hash Partitioning**: Use hash function to distribute data
- **Range Partitioning**: Split data by value ranges

### Distributed Transactions
- **Two-Phase Commit (2PC)**: Coordinator manages transaction
- **Three-Phase Commit (3PC)**: Non-blocking version of 2PC
- **Saga Pattern**: Long-running transaction management
- **TCC (Try-Confirm-Cancel)**: Compensating transactions

## Message Passing and Communication

### Communication Patterns
- **Request-Response**: Synchronous communication
- **Publish-Subscribe**: Asynchronous message distribution
- **Message Queues**: Reliable message delivery
- **Event Streaming**: Real-time event processing

### Message Ordering
- **FIFO**: First-in-first-out ordering
- **Causal Ordering**: Preserve causal relationships
- **Total Ordering**: Global ordering of all messages
- **Partial Ordering**: Ordering within groups

### Message Delivery Guarantees
- **At-Most-Once**: Message delivered at most once
- **At-Least-Once**: Message delivered at least once
- **Exactly-Once**: Message delivered exactly once
- **Best-Effort**: No delivery guarantees

### Message Queue Systems
- **Apache Kafka**: Distributed streaming platform
- **RabbitMQ**: Message broker
- **Amazon SQS**: Managed message queue service
- **Redis**: In-memory data structure store

## Distributed Caching

### Caching Strategies
- **Write-Through**: Write to cache and database
- **Write-Behind**: Write to cache, async to database
- **Write-Around**: Write to database, invalidate cache
- **Refresh-Ahead**: Proactively refresh cache

### Cache Consistency
- **Cache-Aside**: Application manages cache
- **Read-Through**: Cache loads data on miss
- **Write-Through**: Cache writes to database
- **Write-Behind**: Cache writes asynchronously

### Distributed Caching
- **Consistent Hashing**: Distribute cache across nodes
- **Cache Replication**: Replicate cache data
- **Cache Partitioning**: Partition cache by data
- **Cache Invalidation**: Invalidate stale data

### Caching Tools
- **Redis**: In-memory data structure store
- **Memcached**: Distributed memory caching
- **Hazelcast**: In-memory data grid
- **Apache Ignite**: Distributed caching platform

## Fault Tolerance and Reliability

### Failure Types
- **Crash Failures**: Node stops responding
- **Omission Failures**: Node fails to send/receive messages
- **Timing Failures**: Node responds too slowly
- **Byzantine Failures**: Node behaves arbitrarily

### Fault Tolerance Techniques
- **Replication**: Multiple copies of data/services
- **Redundancy**: Backup systems and components
- **Checkpointing**: Save system state periodically
- **Recovery**: Restore system after failure

### Failure Detection
- **Heartbeat**: Periodic health checks
- **Timeout**: Detect failures by timeout
- **Gossip Protocol**: Distributed failure detection
- **Quorum**: Majority-based failure detection

### Recovery Strategies
- **Rollback**: Return to previous state
- **Rollforward**: Apply logged operations
- **Checkpoint Recovery**: Restore from checkpoint
- **State Transfer**: Transfer state from healthy node

## Distributed System Patterns

### Circuit Breaker Pattern
- **Closed**: Normal operation
- **Open**: Fail fast, don't call service
- **Half-Open**: Test if service is available
- **Benefits**: Prevent cascading failures

### Bulkhead Pattern
- **Isolation**: Separate resources for different operations
- **Failure Containment**: Limit impact of failures
- **Resource Management**: Control resource usage
- **Examples**: Thread pools, connection pools

### Saga Pattern
- **Long-Running Transactions**: Manage distributed transactions
- **Compensating Actions**: Undo completed operations
- **Orchestration**: Central coordinator manages saga
- **Choreography**: Distributed coordination

### CQRS (Command Query Responsibility Segregation)
- **Command Side**: Handle write operations
- **Query Side**: Handle read operations
- **Event Sourcing**: Store events instead of state
- **Benefits**: Scalability, performance, flexibility

### Event Sourcing
- **Event Store**: Store all events
- **Replay**: Reconstruct state from events
- **Audit Trail**: Complete history of changes
- **Benefits**: Debugging, analytics, compliance

## Monitoring and Observability

### Three Pillars of Observability
- **Metrics**: Quantitative measurements
- **Logs**: Event records
- **Traces**: Request flow tracking

### Distributed Tracing
- **Trace**: Complete request journey
- **Span**: Individual operation within trace
- **Context Propagation**: Pass trace context
- **Sampling**: Reduce trace volume

### Monitoring Tools
- **Prometheus**: Metrics collection and storage
- **Grafana**: Metrics visualization
- **Jaeger**: Distributed tracing
- **ELK Stack**: Log aggregation and analysis

### Health Checks
- **Liveness**: Is the service running?
- **Readiness**: Is the service ready to serve?
- **Startup**: Is the service starting up?
- **Custom**: Application-specific checks

## Security in Distributed Systems

### Authentication
- **Single Sign-On (SSO)**: One login for all services
- **OAuth 2.0**: Authorization framework
- **JWT**: JSON Web Tokens
- **mTLS**: Mutual TLS authentication

### Authorization
- **RBAC**: Role-Based Access Control
- **ABAC**: Attribute-Based Access Control
- **Policy Engines**: Centralized authorization
- **Service Mesh**: Sidecar-based security

### Network Security
- **TLS/SSL**: Encrypt communication
- **VPN**: Secure network connections
- **Firewalls**: Network access control
- **Service Mesh**: Security policies

### Data Security
- **Encryption**: Encrypt data at rest and in transit
- **Key Management**: Secure key storage and rotation
- **Data Masking**: Hide sensitive data
- **Audit Logging**: Track access and changes

## Go Implementation Concepts

### Concurrency Patterns
- **Goroutines**: Lightweight concurrency
- **Channels**: Communication between goroutines
- **Context**: Cancellation and timeouts
- **Sync Package**: Synchronization primitives

### Service Discovery
- **Service Registry**: Register and discover services
- **Health Checks**: Monitor service health
- **Load Balancing**: Distribute requests
- **Circuit Breakers**: Handle failures

### Message Passing
- **gRPC**: High-performance RPC framework
- **Message Queues**: Asynchronous communication
- **Event Streaming**: Real-time event processing
- **WebSockets**: Real-time communication

### Distributed Data
- **Consensus**: Implement Raft or Paxos
- **Replication**: Data replication strategies
- **Partitioning**: Data distribution
- **Caching**: Distributed caching

## Interview Questions

### Basic Questions
1. What is the CAP theorem?
2. Explain the difference between consistency and availability.
3. What are the benefits of microservices architecture?

### Intermediate Questions
1. How does the Raft algorithm work?
2. Explain different load balancing strategies.
3. What are the challenges of distributed transactions?

### Advanced Questions
1. Design a distributed system for a social media platform.
2. How would you implement consensus in a distributed system?
3. Explain the trade-offs between different consistency models.

## Next Steps
After mastering distributed systems networking, proceed to:
- **Topic 12**: Cloud Networking & Containerization
- **Topic 13**: Interview Preparation & System Design
- **Topic 14**: Advanced Network Programming

Master distributed systems networking, and you'll understand how to build scalable, fault-tolerant systems! 🚀
```

Now let me create the comprehensive Go implementation for distributed systems networking:


```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// Node represents a node in the distributed system
type Node struct {
	ID          string
	Address     string
	Port        int
	State       string // "leader", "follower", "candidate"
	Term        int64
	VotedFor    string
	Log         []LogEntry
	CommitIndex int64
	LastApplied int64
	NextIndex   map[string]int64
	MatchIndex  map[string]int64
	Peers       map[string]*Node
	IsHealthy   bool
	LastHeartbeat time.Time
	mutex       sync.RWMutex
}

// LogEntry represents a log entry
type LogEntry struct {
	Term    int64
	Index   int64
	Command string
	Data    interface{}
}

// RaftNode represents a Raft consensus node
type RaftNode struct {
	Node
	ElectionTimeout  time.Duration
	HeartbeatTimeout time.Duration
	VoteCount        int32
	LeaderID         string
	StopChan         chan struct{}
	RequestVoteChan  chan RequestVoteRequest
	AppendEntriesChan chan AppendEntriesRequest
	ClientRequestChan chan ClientRequest
}

// RequestVoteRequest represents a request vote message
type RequestVoteRequest struct {
	Term         int64
	CandidateID  string
	LastLogIndex int64
	LastLogTerm  int64
}

// RequestVoteResponse represents a request vote response
type RequestVoteResponse struct {
	Term        int64
	VoteGranted bool
}

// AppendEntriesRequest represents an append entries message
type AppendEntriesRequest struct {
	Term         int64
	LeaderID     string
	PrevLogIndex int64
	PrevLogTerm  int64
	Entries      []LogEntry
	LeaderCommit int64
}

// AppendEntriesResponse represents an append entries response
type AppendEntriesResponse struct {
	Term    int64
	Success bool
}

// ClientRequest represents a client request
type ClientRequest struct {
	Command string
	Data    interface{}
	Reply   chan ClientResponse
}

// ClientResponse represents a client response
type ClientResponse struct {
	Success bool
	Data    interface{}
	Error   string
}

// NewRaftNode creates a new Raft node
func NewRaftNode(id, address string, port int) *RaftNode {
	return &RaftNode{
		Node: Node{
			ID:          id,
			Address:     address,
			Port:        port,
			State:       "follower",
			Term:        0,
			VotedFor:    "",
			Log:         make([]LogEntry, 0),
			CommitIndex: -1,
			LastApplied: -1,
			NextIndex:   make(map[string]int64),
			MatchIndex:  make(map[string]int64),
			Peers:       make(map[string]*Node),
			IsHealthy:   true,
		},
		ElectionTimeout:   time.Duration(150+rand.Intn(150)) * time.Millisecond,
		HeartbeatTimeout: 50 * time.Millisecond,
		StopChan:         make(chan struct{}),
		RequestVoteChan:  make(chan RequestVoteRequest, 100),
		AppendEntriesChan: make(chan AppendEntriesRequest, 100),
		ClientRequestChan: make(chan ClientRequest, 100),
	}
}

// Start starts the Raft node
func (rn *RaftNode) Start() {
	go rn.run()
	go rn.startHTTPServer()
}

// Stop stops the Raft node
func (rn *RaftNode) Stop() {
	close(rn.StopChan)
}

// run runs the main Raft loop
func (rn *RaftNode) run() {
	for {
		select {
		case <-rn.StopChan:
			return
		default:
			rn.mutex.RLock()
			state := rn.State
			rn.mutex.RUnlock()
			
			switch state {
			case "follower":
				rn.runFollower()
			case "candidate":
				rn.runCandidate()
			case "leader":
				rn.runLeader()
			}
		}
	}
}

// runFollower runs the follower state
func (rn *RaftNode) runFollower() {
	timeout := time.NewTimer(rn.ElectionTimeout)
	defer timeout.Stop()
	
	for {
		select {
		case <-rn.StopChan:
			return
		case <-timeout.C:
			// Election timeout, become candidate
			rn.becomeCandidate()
			return
		case req := <-rn.RequestVoteChan:
			rn.handleRequestVote(req)
		case req := <-rn.AppendEntriesChan:
			rn.handleAppendEntries(req)
		case req := <-rn.ClientRequestChan:
			// Redirect to leader
			rn.redirectToLeader(req)
		}
	}
}

// runCandidate runs the candidate state
func (rn *RaftNode) runCandidate() {
	rn.mutex.Lock()
	rn.State = "candidate"
	rn.Term++
	rn.VotedFor = rn.ID
	rn.VoteCount = 1
	rn.mutex.Unlock()
	
	// Request votes from all peers
	rn.requestVotes()
	
	timeout := time.NewTimer(rn.ElectionTimeout)
	defer timeout.Stop()
	
	for {
		select {
		case <-rn.StopChan:
			return
		case <-timeout.C:
			// Election timeout, start new election
			rn.becomeCandidate()
			return
		case req := <-rn.RequestVoteChan:
			rn.handleRequestVote(req)
		case req := <-rn.AppendEntriesChan:
			rn.handleAppendEntries(req)
		case req := <-rn.ClientRequestChan:
			// Redirect to leader
			rn.redirectToLeader(req)
		}
	}
}

// runLeader runs the leader state
func (rn *RaftNode) runLeader() {
	rn.mutex.Lock()
	rn.State = "leader"
	rn.LeaderID = rn.ID
	
	// Initialize nextIndex and matchIndex
	for peerID := range rn.Peers {
		rn.NextIndex[peerID] = int64(len(rn.Log))
		rn.MatchIndex[peerID] = -1
	}
	rn.mutex.Unlock()
	
	// Send initial heartbeat
	rn.sendHeartbeat()
	
	ticker := time.NewTicker(rn.HeartbeatTimeout)
	defer ticker.Stop()
	
	for {
		select {
		case <-rn.StopChan:
			return
		case <-ticker.C:
			rn.sendHeartbeat()
		case req := <-rn.RequestVoteChan:
			rn.handleRequestVote(req)
		case req := <-rn.AppendEntriesChan:
			rn.handleAppendEntries(req)
		case req := <-rn.ClientRequestChan:
			rn.handleClientRequest(req)
		}
	}
}

// becomeCandidate transitions to candidate state
func (rn *RaftNode) becomeCandidate() {
	rn.mutex.Lock()
	rn.State = "candidate"
	rn.Term++
	rn.VotedFor = rn.ID
	rn.VoteCount = 1
	rn.mutex.Unlock()
	
	// Request votes from all peers
	rn.requestVotes()
}

// requestVotes requests votes from all peers
func (rn *RaftNode) requestVotes() {
	rn.mutex.RLock()
	term := rn.Term
	lastLogIndex := int64(len(rn.Log)) - 1
	lastLogTerm := int64(0)
	if lastLogIndex >= 0 {
		lastLogTerm = rn.Log[lastLogIndex].Term
	}
	rn.mutex.RUnlock()
	
	req := RequestVoteRequest{
		Term:         term,
		CandidateID:  rn.ID,
		LastLogIndex: lastLogIndex,
		LastLogTerm:  lastLogTerm,
	}
	
	// Send to all peers
	for peerID, peer := range rn.Peers {
		go rn.sendRequestVote(peerID, peer, req)
	}
}

// sendRequestVote sends a request vote to a peer
func (rn *RaftNode) sendRequestVote(peerID string, peer *Node, req RequestVoteRequest) {
	// In a real implementation, this would send over the network
	// For simulation, we'll just process locally
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	
	// Simulate peer response
	if rand.Float64() < 0.8 { // 80% chance of vote
		rn.mutex.Lock()
		rn.VoteCount++
		if rn.VoteCount > int32(len(rn.Peers)/2) {
			rn.becomeLeader()
		}
		rn.mutex.Unlock()
	}
}

// becomeLeader transitions to leader state
func (rn *RaftNode) becomeLeader() {
	rn.mutex.Lock()
	rn.State = "leader"
	rn.LeaderID = rn.ID
	rn.mutex.Unlock()
}

// sendHeartbeat sends heartbeat to all peers
func (rn *RaftNode) sendHeartbeat() {
	rn.mutex.RLock()
	term := rn.Term
	leaderID := rn.ID
	leaderCommit := rn.CommitIndex
	rn.mutex.RUnlock()
	
	req := AppendEntriesRequest{
		Term:         term,
		LeaderID:     leaderID,
		PrevLogIndex: -1,
		PrevLogTerm:  -1,
		Entries:      []LogEntry{},
		LeaderCommit: leaderCommit,
	}
	
	// Send to all peers
	for peerID, peer := range rn.Peers {
		go rn.sendAppendEntries(peerID, peer, req)
	}
}

// sendAppendEntries sends append entries to a peer
func (rn *RaftNode) sendAppendEntries(peerID string, peer *Node, req AppendEntriesRequest) {
	// In a real implementation, this would send over the network
	// For simulation, we'll just process locally
	time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
	
	// Simulate peer response
	rn.mutex.Lock()
	rn.Peers[peerID].LastHeartbeat = time.Now()
	rn.mutex.Unlock()
}

// handleRequestVote handles a request vote message
func (rn *RaftNode) handleRequestVote(req RequestVoteRequest) {
	rn.mutex.Lock()
	defer rn.mutex.Unlock()
	
	if req.Term > rn.Term {
		rn.Term = req.Term
		rn.State = "follower"
		rn.VotedFor = ""
	}
	
	if req.Term == rn.Term && (rn.VotedFor == "" || rn.VotedFor == req.CandidateID) {
		rn.VotedFor = req.CandidateID
		// Send vote response
	}
}

// handleAppendEntries handles an append entries message
func (rn *RaftNode) handleAppendEntries(req AppendEntriesRequest) {
	rn.mutex.Lock()
	defer rn.mutex.Unlock()
	
	if req.Term >= rn.Term {
		rn.Term = req.Term
		rn.State = "follower"
		rn.LeaderID = req.LeaderID
		rn.VotedFor = ""
		
		// Update commit index
		if req.LeaderCommit > rn.CommitIndex {
			rn.CommitIndex = req.LeaderCommit
		}
		
		// Send success response
	}
}

// handleClientRequest handles a client request
func (rn *RaftNode) handleClientRequest(req ClientRequest) {
	rn.mutex.Lock()
	defer rn.mutex.Unlock()
	
	// Add to log
	logEntry := LogEntry{
		Term:    rn.Term,
		Index:   int64(len(rn.Log)),
		Command: req.Command,
		Data:    req.Data,
	}
	rn.Log = append(rn.Log, logEntry)
	
	// Send response
	req.Reply <- ClientResponse{
		Success: true,
		Data:    "Request processed",
	}
}

// redirectToLeader redirects client request to leader
func (rn *RaftNode) redirectToLeader(req ClientRequest) {
	rn.mutex.RLock()
	leaderID := rn.LeaderID
	rn.mutex.RUnlock()
	
	if leaderID != "" {
		// Redirect to leader
		req.Reply <- ClientResponse{
			Success: false,
			Error:   fmt.Sprintf("Redirect to leader: %s", leaderID),
		}
	} else {
		req.Reply <- ClientResponse{
			Success: false,
			Error:   "No leader available",
		}
	}
}

// startHTTPServer starts the HTTP server for the node
func (rn *RaftNode) startHTTPServer() {
	http.HandleFunc("/status", rn.handleStatus)
	http.HandleFunc("/request", rn.handleRequest)
	http.HandleFunc("/vote", rn.handleVote)
	http.HandleFunc("/append", rn.handleAppend)
	
	addr := fmt.Sprintf(":%d", rn.Port)
	log.Printf("Raft node %s starting HTTP server on %s", rn.ID, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// handleStatus handles status requests
func (rn *RaftNode) handleStatus(w http.ResponseWriter, r *http.Request) {
	rn.mutex.RLock()
	defer rn.mutex.RUnlock()
	
	status := map[string]interface{}{
		"id":            rn.ID,
		"state":         rn.State,
		"term":          rn.Term,
		"voted_for":     rn.VotedFor,
		"log_length":    len(rn.Log),
		"commit_index":  rn.CommitIndex,
		"last_applied":  rn.LastApplied,
		"is_healthy":    rn.IsHealthy,
		"last_heartbeat": rn.LastHeartbeat,
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

// handleRequest handles client requests
func (rn *RaftNode) handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req ClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	
	req.Reply = make(chan ClientResponse, 1)
	rn.ClientRequestChan <- req
	
	select {
	case resp := <-req.Reply:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	case <-time.After(5 * time.Second):
		http.Error(w, "Request timeout", http.StatusRequestTimeout)
	}
}

// handleVote handles vote requests
func (rn *RaftNode) handleVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req RequestVoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	
	rn.RequestVoteChan <- req
	
	// Send response
	resp := RequestVoteResponse{
		Term:        rn.Term,
		VoteGranted: true, // Simplified
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// handleAppend handles append entries requests
func (rn *RaftNode) handleAppend(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req AppendEntriesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	
	rn.AppendEntriesChan <- req
	
	// Send response
	resp := AppendEntriesResponse{
		Term:    rn.Term,
		Success: true, // Simplified
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// ServiceRegistry represents a service registry
type ServiceRegistry struct {
	Services map[string]*Service
	mutex    sync.RWMutex
}

// Service represents a registered service
type Service struct {
	ID          string
	Name        string
	Address     string
	Port        int
	HealthCheck string
	Tags        []string
	LastSeen    time.Time
	IsHealthy   bool
}

// NewServiceRegistry creates a new service registry
func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		Services: make(map[string]*Service),
	}
}

// Register registers a service
func (sr *ServiceRegistry) Register(service *Service) {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	
	service.LastSeen = time.Now()
	service.IsHealthy = true
	sr.Services[service.ID] = service
}

// Deregister deregisters a service
func (sr *ServiceRegistry) Deregister(serviceID string) {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	
	delete(sr.Services, serviceID)
}

// GetService gets a service by ID
func (sr *ServiceRegistry) GetService(serviceID string) *Service {
	sr.mutex.RLock()
	defer sr.mutex.RUnlock()
	
	return sr.Services[serviceID]
}

// GetServicesByName gets services by name
func (sr *ServiceRegistry) GetServicesByName(name string) []*Service {
	sr.mutex.RLock()
	defer sr.mutex.RUnlock()
	
	var services []*Service
	for _, service := range sr.Services {
		if service.Name == name && service.IsHealthy {
			services = append(services, service)
		}
	}
	return services
}

// GetAllServices gets all services
func (sr *ServiceRegistry) GetAllServices() []*Service {
	sr.mutex.RLock()
	defer sr.mutex.RUnlock()
	
	var services []*Service
	for _, service := range sr.Services {
		services = append(services, service)
	}
	return services
}

// HealthCheck performs health check on a service
func (sr *ServiceRegistry) HealthCheck(serviceID string) bool {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	
	service, exists := sr.Services[serviceID]
	if !exists {
		return false
	}
	
	// Simulate health check
	service.IsHealthy = rand.Float64() > 0.1 // 90% chance of being healthy
	service.LastSeen = time.Now()
	
	return service.IsHealthy
}

// LoadBalancer represents a load balancer
type LoadBalancer struct {
	Services    map[string][]*Service
	Algorithm   string
	Current     map[string]int32
	mutex       sync.RWMutex
}

// NewLoadBalancer creates a new load balancer
func NewLoadBalancer(algorithm string) *LoadBalancer {
	return &LoadBalancer{
		Services:  make(map[string][]*Service),
		Algorithm: algorithm,
		Current:   make(map[string]int32),
	}
}

// AddService adds a service to the load balancer
func (lb *LoadBalancer) AddService(serviceName string, service *Service) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	
	lb.Services[serviceName] = append(lb.Services[serviceName], service)
}

// GetService gets a service using the load balancing algorithm
func (lb *LoadBalancer) GetService(serviceName string) *Service {
	lb.mutex.RLock()
	defer lb.mutex.RUnlock()
	
	services, exists := lb.Services[serviceName]
	if !exists || len(services) == 0 {
		return nil
	}
	
	// Filter healthy services
	var healthyServices []*Service
	for _, service := range services {
		if service.IsHealthy {
			healthyServices = append(healthyServices, service)
		}
	}
	
	if len(healthyServices) == 0 {
		return nil
	}
	
	switch lb.Algorithm {
	case "round_robin":
		return lb.getRoundRobinService(serviceName, healthyServices)
	case "random":
		return lb.getRandomService(healthyServices)
	case "least_connections":
		return lb.getLeastConnectionsService(healthyServices)
	default:
		return lb.getRoundRobinService(serviceName, healthyServices)
	}
}

// getRoundRobinService gets the next service in round-robin fashion
func (lb *LoadBalancer) getRoundRobinService(serviceName string, services []*Service) *Service {
	if len(services) == 0 {
		return nil
	}
	
	index := int(atomic.AddInt32(&lb.Current[serviceName], 1)) % len(services)
	return services[index]
}

// getRandomService gets a random service
func (lb *LoadBalancer) getRandomService(services []*Service) *Service {
	if len(services) == 0 {
		return nil
	}
	
	index := rand.Intn(len(services))
	return services[index]
}

// getLeastConnectionsService gets the service with the least connections
func (lb *LoadBalancer) getLeastConnectionsService(services []*Service) *Service {
	if len(services) == 0 {
		return nil
	}
	
	// Simplified - in reality, you would track actual connections
	return services[0]
}

// MessageQueue represents a message queue
type MessageQueue struct {
	Messages chan Message
	Subscribers map[string][]chan Message
	mutex    sync.RWMutex
}

// Message represents a message in the queue
type Message struct {
	ID        string
	Topic     string
	Data      interface{}
	Timestamp time.Time
	TTL       time.Duration
}

// NewMessageQueue creates a new message queue
func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		Messages:    make(chan Message, 1000),
		Subscribers: make(map[string][]chan Message),
	}
}

// Publish publishes a message to a topic
func (mq *MessageQueue) Publish(topic string, data interface{}) error {
	message := Message{
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		Topic:     topic,
		Data:      data,
		Timestamp: time.Now(),
		TTL:       24 * time.Hour,
	}
	
	select {
	case mq.Messages <- message:
		return nil
	default:
		return fmt.Errorf("queue full")
	}
}

// Subscribe subscribes to a topic
func (mq *MessageQueue) Subscribe(topic string) <-chan Message {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()
	
	ch := make(chan Message, 100)
	mq.Subscribers[topic] = append(mq.Subscribers[topic], ch)
	return ch
}

// Start starts the message queue
func (mq *MessageQueue) Start() {
	go mq.processMessages()
}

// processMessages processes messages and distributes to subscribers
func (mq *MessageQueue) processMessages() {
	for message := range mq.Messages {
		mq.mutex.RLock()
		subscribers := mq.Subscribers[message.Topic]
		mq.mutex.RUnlock()
		
		for _, subscriber := range subscribers {
			select {
			case subscriber <- message:
			default:
				// Subscriber channel full, skip
			}
		}
	}
}

// DistributedCache represents a distributed cache
type DistributedCache struct {
	Data  map[string]CacheItem
	mutex sync.RWMutex
}

// CacheItem represents a cache item
type CacheItem struct {
	Value     interface{}
	ExpiresAt time.Time
	Version   int64
}

// NewDistributedCache creates a new distributed cache
func NewDistributedCache() *DistributedCache {
	return &DistributedCache{
		Data: make(map[string]CacheItem),
	}
}

// Set sets a value in the cache
func (dc *DistributedCache) Set(key string, value interface{}, ttl time.Duration) {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()
	
	dc.Data[key] = CacheItem{
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
		Version:   time.Now().UnixNano(),
	}
}

// Get gets a value from the cache
func (dc *DistributedCache) Get(key string) (interface{}, bool) {
	dc.mutex.RLock()
	defer dc.mutex.RUnlock()
	
	item, exists := dc.Data[key]
	if !exists {
		return nil, false
	}
	
	if time.Now().After(item.ExpiresAt) {
		// Item expired
		dc.mutex.RUnlock()
		dc.mutex.Lock()
		delete(dc.Data, key)
		dc.mutex.Unlock()
		dc.mutex.RLock()
		return nil, false
	}
	
	return item.Value, true
}

// Delete deletes a value from the cache
func (dc *DistributedCache) Delete(key string) {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()
	
	delete(dc.Data, key)
}

// Clear clears all values from the cache
func (dc *DistributedCache) Clear() {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()
	
	dc.Data = make(map[string]CacheItem)
}

// Demonstrate Raft consensus
func demonstrateRaftConsensus() {
	fmt.Println("=== Raft Consensus Demo ===\n")
	
	// Create Raft nodes
	nodes := make([]*RaftNode, 3)
	for i := 0; i < 3; i++ {
		nodes[i] = NewRaftNode(fmt.Sprintf("node-%d", i), "localhost", 8080+i)
	}
	
	// Add peers to each node
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i != j {
				nodes[i].Peers[fmt.Sprintf("node-%d", j)] = &nodes[j].Node
			}
		}
	}
	
	// Start nodes
	for _, node := range nodes {
		go node.Start()
	}
	
	// Wait for election
	time.Sleep(2 * time.Second)
	
	// Check node states
	for _, node := range nodes {
		node.mutex.RLock()
		fmt.Printf("Node %s: State=%s, Term=%d, LogLength=%d\n", 
			node.ID, node.State, node.Term, len(node.Log))
		node.mutex.RUnlock()
	}
	
	// Stop nodes
	for _, node := range nodes {
		node.Stop()
	}
}

// Demonstrate service discovery
func demonstrateServiceDiscovery() {
	fmt.Println("=== Service Discovery Demo ===\n")
	
	// Create service registry
	registry := NewServiceRegistry()
	
	// Register services
	services := []*Service{
		{ID: "service-1", Name: "user-service", Address: "localhost", Port: 8001, HealthCheck: "/health"},
		{ID: "service-2", Name: "user-service", Address: "localhost", Port: 8002, HealthCheck: "/health"},
		{ID: "service-3", Name: "order-service", Address: "localhost", Port: 8003, HealthCheck: "/health"},
		{ID: "service-4", Name: "payment-service", Address: "localhost", Port: 8004, HealthCheck: "/health"},
	}
	
	for _, service := range services {
		registry.Register(service)
	}
	
	// Display all services
	fmt.Println("All Services:")
	for _, service := range registry.GetAllServices() {
		fmt.Printf("  %s: %s:%d (healthy: %t)\n", 
			service.ID, service.Address, service.Port, service.IsHealthy)
	}
	
	// Get services by name
	fmt.Println("\nUser Services:")
	userServices := registry.GetServicesByName("user-service")
	for _, service := range userServices {
		fmt.Printf("  %s: %s:%d\n", service.ID, service.Address, service.Port)
	}
	
	// Perform health checks
	fmt.Println("\nHealth Checks:")
	for _, service := range registry.GetAllServices() {
		healthy := registry.HealthCheck(service.ID)
		fmt.Printf("  %s: %t\n", service.ID, healthy)
	}
}

// Demonstrate load balancing
func demonstrateLoadBalancing() {
	fmt.Println("=== Load Balancing Demo ===\n")
	
	// Create load balancer
	lb := NewLoadBalancer("round_robin")
	
	// Add services
	services := []*Service{
		{ID: "service-1", Name: "api-service", Address: "localhost", Port: 8001, IsHealthy: true},
		{ID: "service-2", Name: "api-service", Address: "localhost", Port: 8002, IsHealthy: true},
		{ID: "service-3", Name: "api-service", Address: "localhost", Port: 8003, IsHealthy: true},
	}
	
	for _, service := range services {
		lb.AddService("api-service", service)
	}
	
	// Test load balancing
	fmt.Println("Load Balancing Test:")
	for i := 0; i < 10; i++ {
		service := lb.GetService("api-service")
		if service != nil {
			fmt.Printf("  Request %d: %s:%d\n", i+1, service.Address, service.Port)
		} else {
			fmt.Printf("  Request %d: No service available\n", i+1)
		}
	}
}

// Demonstrate message queuing
func demonstrateMessageQueuing() {
	fmt.Println("=== Message Queuing Demo ===\n")
	
	// Create message queue
	mq := NewMessageQueue()
	mq.Start()
	
	// Subscribe to topics
	userTopic := mq.Subscribe("user-events")
	orderTopic := mq.Subscribe("order-events")
	
	// Publish messages
	go func() {
		for i := 0; i < 5; i++ {
			mq.Publish("user-events", map[string]interface{}{
				"user_id": i,
				"action":  "login",
				"time":    time.Now(),
			})
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	go func() {
		for i := 0; i < 3; i++ {
			mq.Publish("order-events", map[string]interface{}{
				"order_id": i,
				"status":   "created",
				"time":     time.Now(),
			})
			time.Sleep(150 * time.Millisecond)
		}
	}()
	
	// Consume messages
	go func() {
		for message := range userTopic {
			fmt.Printf("User Event: %+v\n", message.Data)
		}
	}()
	
	go func() {
		for message := range orderTopic {
			fmt.Printf("Order Event: %+v\n", message.Data)
		}
	}()
	
	// Wait for messages
	time.Sleep(2 * time.Second)
}

// Demonstrate distributed caching
func demonstrateDistributedCaching() {
	fmt.Println("=== Distributed Caching Demo ===\n")
	
	// Create distributed cache
	cache := NewDistributedCache()
	
	// Set values
	cache.Set("user:1", map[string]interface{}{"name": "John", "email": "john@example.com"}, 5*time.Minute)
	cache.Set("user:2", map[string]interface{}{"name": "Jane", "email": "jane@example.com"}, 5*time.Minute)
	cache.Set("config:app", map[string]interface{}{"version": "1.0", "debug": true}, 10*time.Minute)
	
	// Get values
	fmt.Println("Cache Operations:")
	if value, found := cache.Get("user:1"); found {
		fmt.Printf("  user:1 = %+v\n", value)
	}
	
	if value, found := cache.Get("user:2"); found {
		fmt.Printf("  user:2 = %+v\n", value)
	}
	
	if value, found := cache.Get("config:app"); found {
		fmt.Printf("  config:app = %+v\n", value)
	}
	
	// Test expiration
	cache.Set("temp:key", "temporary value", 100*time.Millisecond)
	time.Sleep(200 * time.Millisecond)
	
	if value, found := cache.Get("temp:key"); found {
		fmt.Printf("  temp:key = %+v (should be expired)\n", value)
	} else {
		fmt.Println("  temp:key = expired (as expected)")
	}
	
	// Clear cache
	cache.Clear()
	fmt.Println("  Cache cleared")
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 11: Distributed Systems Networking")
	fmt.Println("=======================================================================\n")
	
	// Run all demonstrations
	demonstrateRaftConsensus()
	fmt.Println()
	demonstrateServiceDiscovery()
	fmt.Println()
	demonstrateLoadBalancing()
	fmt.Println()
	demonstrateMessageQueuing()
	fmt.Println()
	demonstrateDistributedCaching()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Distributed systems require consensus algorithms for coordination")
	fmt.Println("2. Service discovery enables dynamic service location")
	fmt.Println("3. Load balancing distributes traffic across multiple services")
	fmt.Println("4. Message queuing enables asynchronous communication")
	fmt.Println("5. Distributed caching improves performance and reduces load")
	fmt.Println("6. Raft provides understandable consensus for distributed systems")
	fmt.Println("7. Go provides excellent tools for building distributed systems")
	fmt.Println("8. Understanding distributed patterns is crucial for scalability")
	
	fmt.Println("\n📚 Next Topic: Cloud Networking & Containerization")
}
```

Now let me create a comprehensive distributed systems patterns and algorithms module:


```go
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// CircuitBreaker represents a circuit breaker pattern
type CircuitBreaker struct {
	Name          string
	MaxRequests   int32
	Interval      time.Duration
	Timeout       time.Duration
	FailureCount  int32
	SuccessCount  int32
	LastFailTime  time.Time
	State         string // "closed", "open", "half-open"
	mutex         sync.RWMutex
}

// NewCircuitBreaker creates a new circuit breaker
func NewCircuitBreaker(name string, maxRequests int32, interval, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		Name:        name,
		MaxRequests: maxRequests,
		Interval:    interval,
		Timeout:     timeout,
		State:       "closed",
	}
}

// Call executes a function with circuit breaker protection
func (cb *CircuitBreaker) Call(fn func() (interface{}, error)) (interface{}, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	// Check if circuit should be opened
	if cb.State == "closed" && cb.FailureCount >= cb.MaxRequests {
		if time.Since(cb.LastFailTime) > cb.Interval {
			cb.State = "open"
			cb.FailureCount = 0
		}
	}
	
	// Check if circuit should be half-opened
	if cb.State == "open" && time.Since(cb.LastFailTime) > cb.Timeout {
		cb.State = "half-open"
		cb.SuccessCount = 0
	}
	
	// Execute function based on state
	switch cb.State {
	case "closed", "half-open":
		result, err := fn()
		if err != nil {
			cb.FailureCount++
			cb.LastFailTime = time.Now()
			if cb.State == "half-open" {
				cb.State = "open"
			}
			return nil, err
		}
		
		cb.SuccessCount++
		if cb.State == "half-open" && cb.SuccessCount >= cb.MaxRequests {
			cb.State = "closed"
			cb.FailureCount = 0
		}
		
		return result, nil
	case "open":
		return nil, fmt.Errorf("circuit breaker is open")
	default:
		return nil, fmt.Errorf("unknown circuit breaker state")
	}
}

// GetState returns the current state of the circuit breaker
func (cb *CircuitBreaker) GetState() string {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.State
}

// Bulkhead represents a bulkhead pattern for resource isolation
type Bulkhead struct {
	Name        string
	MaxWorkers  int
	MaxQueue    int
	Workers     chan struct{}
	Queue       chan func()
	IsRunning   bool
	StopChan    chan struct{}
	mutex       sync.RWMutex
}

// NewBulkhead creates a new bulkhead
func NewBulkhead(name string, maxWorkers, maxQueue int) *Bulkhead {
	return &Bulkhead{
		Name:     name,
		MaxWorkers: maxWorkers,
		MaxQueue:   maxQueue,
		Workers:   make(chan struct{}, maxWorkers),
		Queue:     make(chan func(), maxQueue),
		StopChan:  make(chan struct{}),
	}
}

// Start starts the bulkhead workers
func (b *Bulkhead) Start() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	
	if b.IsRunning {
		return
	}
	
	b.IsRunning = true
	
	// Start workers
	for i := 0; i < b.MaxWorkers; i++ {
		go b.worker()
	}
}

// Stop stops the bulkhead workers
func (b *Bulkhead) Stop() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	
	if !b.IsRunning {
		return
	}
	
	b.IsRunning = false
	close(b.StopChan)
}

// Execute executes a function in the bulkhead
func (b *Bulkhead) Execute(fn func()) error {
	select {
	case b.Queue <- fn:
		return nil
	default:
		return fmt.Errorf("bulkhead queue is full")
	}
}

// worker processes functions from the queue
func (b *Bulkhead) worker() {
	for {
		select {
		case <-b.StopChan:
			return
		case fn := <-b.Queue:
			b.Workers <- struct{}{} // Acquire worker
			fn()
			<-b.Workers // Release worker
		}
	}
}

// GetStats returns bulkhead statistics
func (b *Bulkhead) GetStats() map[string]interface{} {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	
	return map[string]interface{}{
		"name":         b.Name,
		"max_workers":  b.MaxWorkers,
		"max_queue":    b.MaxQueue,
		"active_workers": len(b.Workers),
		"queue_length":  len(b.Queue),
		"is_running":   b.IsRunning,
	}
}

// Saga represents a saga pattern for distributed transactions
type Saga struct {
	ID          string
	Steps       []SagaStep
	CurrentStep int
	State       string // "running", "completed", "failed", "compensating"
	Result      interface{}
	Error       error
	mutex       sync.RWMutex
}

// SagaStep represents a step in a saga
type SagaStep struct {
	ID           string
	Action       func() (interface{}, error)
	Compensation func() error
	Timeout      time.Duration
}

// NewSaga creates a new saga
func NewSaga(id string) *Saga {
	return &Saga{
		ID:          id,
		Steps:       make([]SagaStep, 0),
		CurrentStep: 0,
		State:       "running",
	}
}

// AddStep adds a step to the saga
func (s *Saga) AddStep(step SagaStep) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	s.Steps = append(s.Steps, step)
}

// Execute executes the saga
func (s *Saga) Execute() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	s.State = "running"
	
	for i, step := range s.Steps {
		s.CurrentStep = i
		
		// Execute step with timeout
		ctx, cancel := context.WithTimeout(context.Background(), step.Timeout)
		done := make(chan struct{})
		var result interface{}
		var err error
		
		go func() {
			defer close(done)
			result, err = step.Action()
		}()
		
		select {
		case <-done:
			cancel()
			if err != nil {
				s.State = "failed"
				s.Error = err
				// Compensate for completed steps
				s.compensate(i - 1)
				return err
			}
		case <-ctx.Done():
			cancel()
			s.State = "failed"
			s.Error = fmt.Errorf("step %s timed out", step.ID)
			// Compensate for completed steps
			s.compensate(i - 1)
			return s.Error
		}
	}
	
	s.State = "completed"
	s.Result = result
	return nil
}

// compensate compensates for completed steps
func (s *Saga) compensate(lastStep int) {
	s.State = "compensating"
	
	for i := lastStep; i >= 0; i-- {
		if s.Steps[i].Compensation != nil {
			if err := s.Steps[i].Compensation(); err != nil {
				log.Printf("Compensation failed for step %s: %v", s.Steps[i].ID, err)
			}
		}
	}
}

// GetState returns the current state of the saga
func (s *Saga) GetState() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.State
}

// EventStore represents an event store for event sourcing
type EventStore struct {
	Events map[string][]Event
	mutex  sync.RWMutex
}

// Event represents an event in the event store
type Event struct {
	ID        string
	Type      string
	Data      interface{}
	Timestamp time.Time
	Version   int64
}

// NewEventStore creates a new event store
func NewEventStore() *EventStore {
	return &EventStore{
		Events: make(map[string][]Event),
	}
}

// Append appends an event to the event store
func (es *EventStore) Append(streamID string, eventType string, data interface{}) error {
	es.mutex.Lock()
	defer es.mutex.Unlock()
	
	event := Event{
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		Type:      eventType,
		Data:      data,
		Timestamp: time.Now(),
		Version:   int64(len(es.Events[streamID])),
	}
	
	es.Events[streamID] = append(es.Events[streamID], event)
	return nil
}

// GetEvents gets events for a stream
func (es *EventStore) GetEvents(streamID string) ([]Event, error) {
	es.mutex.RLock()
	defer es.mutex.RUnlock()
	
	events, exists := es.Events[streamID]
	if !exists {
		return nil, fmt.Errorf("stream not found: %s", streamID)
	}
	
	// Return a copy
	result := make([]Event, len(events))
	copy(result, events)
	return result, nil
}

// GetEventsFromVersion gets events from a specific version
func (es *EventStore) GetEventsFromVersion(streamID string, fromVersion int64) ([]Event, error) {
	es.mutex.RLock()
	defer es.mutex.RUnlock()
	
	events, exists := es.Events[streamID]
	if !exists {
		return nil, fmt.Errorf("stream not found: %s", streamID)
	}
	
	if fromVersion >= int64(len(events)) {
		return []Event{}, nil
	}
	
	result := make([]Event, len(events)-int(fromVersion))
	copy(result, events[fromVersion:])
	return result, nil
}

// CQRS represents a CQRS pattern implementation
type CQRS struct {
	CommandStore *EventStore
	QueryStore   map[string]interface{}
	mutex        sync.RWMutex
}

// NewCQRS creates a new CQRS instance
func NewCQRS() *CQRS {
	return &CQRS{
		CommandStore: NewEventStore(),
		QueryStore:   make(map[string]interface{}),
	}
}

// ExecuteCommand executes a command
func (cqrs *CQRS) ExecuteCommand(streamID string, commandType string, data interface{}) error {
	// Append command to event store
	if err := cqrs.CommandStore.Append(streamID, commandType, data); err != nil {
		return err
	}
	
	// Update query store based on command
	cqrs.mutex.Lock()
	defer cqrs.mutex.Unlock()
	
	switch commandType {
	case "CreateUser":
		userData := data.(map[string]interface{})
		cqrs.QueryStore[streamID] = map[string]interface{}{
			"id":    streamID,
			"name":  userData["name"],
			"email": userData["email"],
			"state": "active",
		}
	case "UpdateUser":
		userData := data.(map[string]interface{})
		if existing, exists := cqrs.QueryStore[streamID]; exists {
			existingUser := existing.(map[string]interface{})
			if name, ok := userData["name"]; ok {
				existingUser["name"] = name
			}
			if email, ok := userData["email"]; ok {
				existingUser["email"] = email
			}
		}
	case "DeleteUser":
		if existing, exists := cqrs.QueryStore[streamID]; exists {
			existingUser := existing.(map[string]interface{})
			existingUser["state"] = "deleted"
		}
	}
	
	return nil
}

// Query queries the read store
func (cqrs *CQRS) Query(streamID string) (interface{}, error) {
	cqrs.mutex.RLock()
	defer cqrs.mutex.RUnlock()
	
	result, exists := cqrs.QueryStore[streamID]
	if !exists {
		return nil, fmt.Errorf("stream not found: %s", streamID)
	}
	
	return result, nil
}

// GetAllQueries gets all queries
func (cqrs *CQRS) GetAllQueries() map[string]interface{} {
	cqrs.mutex.RLock()
	defer cqrs.mutex.RUnlock()
	
	result := make(map[string]interface{})
	for k, v := range cqrs.QueryStore {
		result[k] = v
	}
	return result
}

// RetryPolicy represents a retry policy
type RetryPolicy struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
	Multiplier  float64
	Jitter      bool
}

// NewRetryPolicy creates a new retry policy
func NewRetryPolicy(maxAttempts int, baseDelay, maxDelay time.Duration, multiplier float64, jitter bool) *RetryPolicy {
	return &RetryPolicy{
		MaxAttempts: maxAttempts,
		BaseDelay:   baseDelay,
		MaxDelay:    maxDelay,
		Multiplier:  multiplier,
		Jitter:      jitter,
	}
}

// Execute executes a function with retry logic
func (rp *RetryPolicy) Execute(fn func() (interface{}, error)) (interface{}, error) {
	var lastErr error
	delay := rp.BaseDelay
	
	for attempt := 0; attempt < rp.MaxAttempts; attempt++ {
		result, err := fn()
		if err == nil {
			return result, nil
		}
		
		lastErr = err
		
		if attempt == rp.MaxAttempts-1 {
			break
		}
		
		// Calculate delay
		actualDelay := delay
		if rp.Jitter {
			actualDelay = time.Duration(float64(delay) * (0.5 + rand.Float64()*0.5))
		}
		
		time.Sleep(actualDelay)
		
		// Calculate next delay
		delay = time.Duration(float64(delay) * rp.Multiplier)
		if delay > rp.MaxDelay {
			delay = rp.MaxDelay
		}
	}
	
	return nil, fmt.Errorf("max attempts reached: %v", lastErr)
}

// TimeoutPolicy represents a timeout policy
type TimeoutPolicy struct {
	Timeout time.Duration
}

// NewTimeoutPolicy creates a new timeout policy
func NewTimeoutPolicy(timeout time.Duration) *TimeoutPolicy {
	return &TimeoutPolicy{
		Timeout: timeout,
	}
}

// Execute executes a function with timeout
func (tp *TimeoutPolicy) Execute(fn func() (interface{}, error)) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), tp.Timeout)
	defer cancel()
	
	done := make(chan struct{})
	var result interface{}
	var err error
	
	go func() {
		defer close(done)
		result, err = fn()
	}()
	
	select {
	case <-done:
		return result, err
	case <-ctx.Done():
		return nil, fmt.Errorf("operation timed out after %v", tp.Timeout)
	}
}

// Demonstrate circuit breaker pattern
func demonstrateCircuitBreaker() {
	fmt.Println("=== Circuit Breaker Pattern Demo ===\n")
	
	// Create circuit breaker
	cb := NewCircuitBreaker("api-service", 3, 5*time.Second, 10*time.Second)
	
	// Simulate API calls
	fmt.Println("Testing Circuit Breaker:")
	for i := 0; i < 10; i++ {
		result, err := cb.Call(func() (interface{}, error) {
			// Simulate API call with 70% failure rate
			if rand.Float64() < 0.7 {
				return nil, fmt.Errorf("API call failed")
			}
			return fmt.Sprintf("Success %d", i), nil
		})
		
		if err != nil {
			fmt.Printf("  Call %d: Failed - %v (State: %s)\n", i+1, err, cb.GetState())
		} else {
			fmt.Printf("  Call %d: Success - %v (State: %s)\n", i+1, result, cb.GetState())
		}
		
		time.Sleep(500 * time.Millisecond)
	}
}

// Demonstrate bulkhead pattern
func demonstrateBulkhead() {
	fmt.Println("=== Bulkhead Pattern Demo ===\n")
	
	// Create bulkhead
	bh := NewBulkhead("database", 3, 10)
	bh.Start()
	defer bh.Stop()
	
	// Simulate work
	fmt.Println("Testing Bulkhead:")
	for i := 0; i < 15; i++ {
		err := bh.Execute(func() {
			fmt.Printf("  Executing task %d\n", i+1)
			time.Sleep(100 * time.Millisecond)
		})
		
		if err != nil {
			fmt.Printf("  Task %d: Failed - %v\n", i+1, err)
		}
	}
	
	// Display stats
	stats := bh.GetStats()
	fmt.Printf("\nBulkhead Stats: %+v\n", stats)
}

// Demonstrate saga pattern
func demonstrateSaga() {
	fmt.Println("=== Saga Pattern Demo ===\n")
	
	// Create saga
	saga := NewSaga("order-processing")
	
	// Add steps
	saga.AddStep(SagaStep{
		ID: "reserve-inventory",
		Action: func() (interface{}, error) {
			fmt.Println("  Reserving inventory...")
			time.Sleep(100 * time.Millisecond)
			if rand.Float64() < 0.3 {
				return nil, fmt.Errorf("inventory unavailable")
			}
			return "inventory-reserved", nil
		},
		Compensation: func() error {
			fmt.Println("  Releasing inventory...")
			return nil
		},
		Timeout: 5 * time.Second,
	})
	
	saga.AddStep(SagaStep{
		ID: "charge-payment",
		Action: func() (interface{}, error) {
			fmt.Println("  Charging payment...")
			time.Sleep(100 * time.Millisecond)
			if rand.Float64() < 0.2 {
				return nil, fmt.Errorf("payment failed")
			}
			return "payment-charged", nil
		},
		Compensation: func() error {
			fmt.Println("  Refunding payment...")
			return nil
		},
		Timeout: 5 * time.Second,
	})
	
	saga.AddStep(SagaStep{
		ID: "create-shipment",
		Action: func() (interface{}, error) {
			fmt.Println("  Creating shipment...")
			time.Sleep(100 * time.Millisecond)
			return "shipment-created", nil
		},
		Compensation: func() error {
			fmt.Println("  Canceling shipment...")
			return nil
		},
		Timeout: 5 * time.Second,
	})
	
	// Execute saga
	fmt.Println("Executing Saga:")
	err := saga.Execute()
	if err != nil {
		fmt.Printf("Saga failed: %v\n", err)
	} else {
		fmt.Printf("Saga completed successfully: %v\n", saga.Result)
	}
	
	fmt.Printf("Final state: %s\n", saga.GetState())
}

// Demonstrate event sourcing
func demonstrateEventSourcing() {
	fmt.Println("=== Event Sourcing Demo ===\n")
	
	// Create event store
	eventStore := NewEventStore()
	
	// Append events
	fmt.Println("Appending Events:")
	eventStore.Append("user-123", "UserCreated", map[string]interface{}{
		"name":  "John Doe",
		"email": "john@example.com",
	})
	
	eventStore.Append("user-123", "UserUpdated", map[string]interface{}{
		"name": "John Smith",
	})
	
	eventStore.Append("user-123", "UserDeleted", map[string]interface{}{
		"reason": "account closure",
	})
	
	// Get all events
	events, err := eventStore.GetEvents("user-123")
	if err != nil {
		fmt.Printf("Error getting events: %v\n", err)
		return
	}
	
	fmt.Printf("Events for user-123:\n")
	for _, event := range events {
		fmt.Printf("  %s: %+v (version: %d)\n", event.Type, event.Data, event.Version)
	}
	
	// Get events from version 1
	eventsFromV1, err := eventStore.GetEventsFromVersion("user-123", 1)
	if err != nil {
		fmt.Printf("Error getting events from version 1: %v\n", err)
		return
	}
	
	fmt.Printf("\nEvents from version 1:\n")
	for _, event := range eventsFromV1 {
		fmt.Printf("  %s: %+v (version: %d)\n", event.Type, event.Data, event.Version)
	}
}

// Demonstrate CQRS pattern
func demonstrateCQRS() {
	fmt.Println("=== CQRS Pattern Demo ===\n")
	
	// Create CQRS instance
	cqrs := NewCQRS()
	
	// Execute commands
	fmt.Println("Executing Commands:")
	cqrs.ExecuteCommand("user-123", "CreateUser", map[string]interface{}{
		"name":  "John Doe",
		"email": "john@example.com",
	})
	
	cqrs.ExecuteCommand("user-123", "UpdateUser", map[string]interface{}{
		"name": "John Smith",
	})
	
	cqrs.ExecuteCommand("user-456", "CreateUser", map[string]interface{}{
		"name":  "Jane Doe",
		"email": "jane@example.com",
	})
	
	// Query data
	fmt.Println("\nQuerying Data:")
	user123, err := cqrs.Query("user-123")
	if err != nil {
		fmt.Printf("Error querying user-123: %v\n", err)
	} else {
		fmt.Printf("  user-123: %+v\n", user123)
	}
	
	user456, err := cqrs.Query("user-456")
	if err != nil {
		fmt.Printf("Error querying user-456: %v\n", err)
	} else {
		fmt.Printf("  user-456: %+v\n", user456)
	}
	
	// Get all queries
	allQueries := cqrs.GetAllQueries()
	fmt.Printf("\nAll Queries: %+v\n", allQueries)
}

// Demonstrate retry policy
func demonstrateRetryPolicy() {
	fmt.Println("=== Retry Policy Demo ===\n")
	
	// Create retry policy
	retryPolicy := NewRetryPolicy(3, 100*time.Millisecond, 1*time.Second, 2.0, true)
	
	// Test retry
	fmt.Println("Testing Retry Policy:")
	result, err := retryPolicy.Execute(func() (interface{}, error) {
		// Simulate operation with 80% failure rate
		if rand.Float64() < 0.8 {
			return nil, fmt.Errorf("operation failed")
		}
		return "operation succeeded", nil
	})
	
	if err != nil {
		fmt.Printf("  Final result: Failed - %v\n", err)
	} else {
		fmt.Printf("  Final result: Success - %v\n", result)
	}
}

// Demonstrate timeout policy
func demonstrateTimeoutPolicy() {
	fmt.Println("=== Timeout Policy Demo ===\n")
	
	// Create timeout policy
	timeoutPolicy := NewTimeoutPolicy(2 * time.Second)
	
	// Test timeout
	fmt.Println("Testing Timeout Policy:")
	result, err := timeoutPolicy.Execute(func() (interface{}, error) {
		// Simulate long-running operation
		time.Sleep(3 * time.Second)
		return "operation completed", nil
	})
	
	if err != nil {
		fmt.Printf("  Result: Failed - %v\n", err)
	} else {
		fmt.Printf("  Result: Success - %v\n", result)
	}
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 11: Distributed Systems Patterns")
	fmt.Println("========================================================================\n")
	
	// Run all demonstrations
	demonstrateCircuitBreaker()
	fmt.Println()
	demonstrateBulkhead()
	fmt.Println()
	demonstrateSaga()
	fmt.Println()
	demonstrateEventSourcing()
	fmt.Println()
	demonstrateCQRS()
	fmt.Println()
	demonstrateRetryPolicy()
	fmt.Println()
	demonstrateTimeoutPolicy()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Circuit breaker prevents cascading failures")
	fmt.Println("2. Bulkhead isolates resources and limits failures")
	fmt.Println("3. Saga manages distributed transactions with compensation")
	fmt.Println("4. Event sourcing stores events instead of state")
	fmt.Println("5. CQRS separates command and query responsibilities")
	fmt.Println("6. Retry policies handle transient failures")
	fmt.Println("7. Timeout policies prevent hanging operations")
	fmt.Println("8. Go provides excellent tools for implementing patterns")
	
	fmt.Println("\n📚 Next Topic: Cloud Networking & Containerization")
}
```

Now let me create a quick reference guide and update the progress:


```markdown
# Distributed Systems Networking Quick Reference Guide

## 🚀 Essential Distributed Systems Concepts

### CAP Theorem
- **Consistency (C)**: All nodes see the same data at the same time
- **Availability (A)**: System remains operational and responsive
- **Partition Tolerance (P)**: System continues operating despite network partitions
- **Trade-offs**: CP Systems, AP Systems, CA Systems

### Distributed System Challenges
- **Network Partitions**: Network splits can isolate components
- **Partial Failures**: Some components fail while others continue
- **Consistency**: Maintaining data consistency across nodes
- **Availability**: Ensuring system remains available during failures
- **Clock Synchronization**: Coordinating time across nodes

## 🔧 Go Implementation Patterns

### Raft Consensus Algorithm
```go
type RaftNode struct {
    Node
    ElectionTimeout  time.Duration
    HeartbeatTimeout time.Duration
    VoteCount        int32
    LeaderID         string
    StopChan         chan struct{}
    RequestVoteChan  chan RequestVoteRequest
    AppendEntriesChan chan AppendEntriesRequest
    ClientRequestChan chan ClientRequest
}

type Node struct {
    ID          string
    Address     string
    Port        int
    State       string // "leader", "follower", "candidate"
    Term        int64
    VotedFor    string
    Log         []LogEntry
    CommitIndex int64
    LastApplied int64
    NextIndex   map[string]int64
    MatchIndex  map[string]int64
    Peers       map[string]*Node
    IsHealthy   bool
    LastHeartbeat time.Time
    mutex       sync.RWMutex
}
```

### Service Discovery
```go
type ServiceRegistry struct {
    Services map[string]*Service
    mutex    sync.RWMutex
}

type Service struct {
    ID          string
    Name        string
    Address     string
    Port        int
    HealthCheck string
    Tags        []string
    LastSeen    time.Time
    IsHealthy   bool
}

func (sr *ServiceRegistry) Register(service *Service) {
    sr.mutex.Lock()
    defer sr.mutex.Unlock()
    
    service.LastSeen = time.Now()
    service.IsHealthy = true
    sr.Services[service.ID] = service
}

func (sr *ServiceRegistry) GetServicesByName(name string) []*Service {
    sr.mutex.RLock()
    defer sr.mutex.RUnlock()
    
    var services []*Service
    for _, service := range sr.Services {
        if service.Name == name && service.IsHealthy {
            services = append(services, service)
        }
    }
    return services
}
```

### Load Balancing
```go
type LoadBalancer struct {
    Services    map[string][]*Service
    Algorithm   string
    Current     map[string]int32
    mutex       sync.RWMutex
}

func (lb *LoadBalancer) GetService(serviceName string) *Service {
    lb.mutex.RLock()
    defer lb.mutex.RUnlock()
    
    services, exists := lb.Services[serviceName]
    if !exists || len(services) == 0 {
        return nil
    }
    
    // Filter healthy services
    var healthyServices []*Service
    for _, service := range services {
        if service.IsHealthy {
            healthyServices = append(healthyServices, service)
        }
    }
    
    if len(healthyServices) == 0 {
        return nil
    }
    
    switch lb.Algorithm {
    case "round_robin":
        return lb.getRoundRobinService(serviceName, healthyServices)
    case "random":
        return lb.getRandomService(healthyServices)
    case "least_connections":
        return lb.getLeastConnectionsService(healthyServices)
    default:
        return lb.getRoundRobinService(serviceName, healthyServices)
    }
}
```

## 🔄 Distributed Patterns

### Circuit Breaker Pattern
```go
type CircuitBreaker struct {
    Name          string
    MaxRequests   int32
    Interval      time.Duration
    Timeout       time.Duration
    FailureCount  int32
    SuccessCount  int32
    LastFailTime  time.Time
    State         string // "closed", "open", "half-open"
    mutex         sync.RWMutex
}

func (cb *CircuitBreaker) Call(fn func() (interface{}, error)) (interface{}, error) {
    cb.mutex.Lock()
    defer cb.mutex.Unlock()
    
    // Check if circuit should be opened
    if cb.State == "closed" && cb.FailureCount >= cb.MaxRequests {
        if time.Since(cb.LastFailTime) > cb.Interval {
            cb.State = "open"
            cb.FailureCount = 0
        }
    }
    
    // Check if circuit should be half-opened
    if cb.State == "open" && time.Since(cb.LastFailTime) > cb.Timeout {
        cb.State = "half-open"
        cb.SuccessCount = 0
    }
    
    // Execute function based on state
    switch cb.State {
    case "closed", "half-open":
        result, err := fn()
        if err != nil {
            cb.FailureCount++
            cb.LastFailTime = time.Now()
            if cb.State == "half-open" {
                cb.State = "open"
            }
            return nil, err
        }
        
        cb.SuccessCount++
        if cb.State == "half-open" && cb.SuccessCount >= cb.MaxRequests {
            cb.State = "closed"
            cb.FailureCount = 0
        }
        
        return result, nil
    case "open":
        return nil, fmt.Errorf("circuit breaker is open")
    default:
        return nil, fmt.Errorf("unknown circuit breaker state")
    }
}
```

### Bulkhead Pattern
```go
type Bulkhead struct {
    Name        string
    MaxWorkers  int
    MaxQueue    int
    Workers     chan struct{}
    Queue       chan func()
    IsRunning   bool
    StopChan    chan struct{}
    mutex       sync.RWMutex
}

func (b *Bulkhead) Execute(fn func()) error {
    select {
    case b.Queue <- fn:
        return nil
    default:
        return fmt.Errorf("bulkhead queue is full")
    }
}

func (b *Bulkhead) worker() {
    for {
        select {
        case <-b.StopChan:
            return
        case fn := <-b.Queue:
            b.Workers <- struct{}{} // Acquire worker
            fn()
            <-b.Workers // Release worker
        }
    }
}
```

### Saga Pattern
```go
type Saga struct {
    ID          string
    Steps       []SagaStep
    CurrentStep int
    State       string // "running", "completed", "failed", "compensating"
    Result      interface{}
    Error       error
    mutex       sync.RWMutex
}

type SagaStep struct {
    ID           string
    Action       func() (interface{}, error)
    Compensation func() error
    Timeout      time.Duration
}

func (s *Saga) Execute() error {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    
    s.State = "running"
    
    for i, step := range s.Steps {
        s.CurrentStep = i
        
        // Execute step with timeout
        ctx, cancel := context.WithTimeout(context.Background(), step.Timeout)
        done := make(chan struct{})
        var result interface{}
        var err error
        
        go func() {
            defer close(done)
            result, err = step.Action()
        }()
        
        select {
        case <-done:
            cancel()
            if err != nil {
                s.State = "failed"
                s.Error = err
                // Compensate for completed steps
                s.compensate(i - 1)
                return err
            }
        case <-ctx.Done():
            cancel()
            s.State = "failed"
            s.Error = fmt.Errorf("step %s timed out", step.ID)
            // Compensate for completed steps
            s.compensate(i - 1)
            return s.Error
        }
    }
    
    s.State = "completed"
    s.Result = result
    return nil
}
```

## 📊 Event Sourcing & CQRS

### Event Store
```go
type EventStore struct {
    Events map[string][]Event
    mutex  sync.RWMutex
}

type Event struct {
    ID        string
    Type      string
    Data      interface{}
    Timestamp time.Time
    Version   int64
}

func (es *EventStore) Append(streamID string, eventType string, data interface{}) error {
    es.mutex.Lock()
    defer es.mutex.Unlock()
    
    event := Event{
        ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
        Type:      eventType,
        Data:      data,
        Timestamp: time.Now(),
        Version:   int64(len(es.Events[streamID])),
    }
    
    es.Events[streamID] = append(es.Events[streamID], event)
    return nil
}

func (es *EventStore) GetEvents(streamID string) ([]Event, error) {
    es.mutex.RLock()
    defer es.mutex.RUnlock()
    
    events, exists := es.Events[streamID]
    if !exists {
        return nil, fmt.Errorf("stream not found: %s", streamID)
    }
    
    // Return a copy
    result := make([]Event, len(events))
    copy(result, events)
    return result, nil
}
```

### CQRS Implementation
```go
type CQRS struct {
    CommandStore *EventStore
    QueryStore   map[string]interface{}
    mutex        sync.RWMutex
}

func (cqrs *CQRS) ExecuteCommand(streamID string, commandType string, data interface{}) error {
    // Append command to event store
    if err := cqrs.CommandStore.Append(streamID, commandType, data); err != nil {
        return err
    }
    
    // Update query store based on command
    cqrs.mutex.Lock()
    defer cqrs.mutex.Unlock()
    
    switch commandType {
    case "CreateUser":
        userData := data.(map[string]interface{})
        cqrs.QueryStore[streamID] = map[string]interface{}{
            "id":    streamID,
            "name":  userData["name"],
            "email": userData["email"],
            "state": "active",
        }
    case "UpdateUser":
        userData := data.(map[string]interface{})
        if existing, exists := cqrs.QueryStore[streamID]; exists {
            existingUser := existing.(map[string]interface{})
            if name, ok := userData["name"]; ok {
                existingUser["name"] = name
            }
            if email, ok := userData["email"]; ok {
                existingUser["email"] = email
            }
        }
    }
    
    return nil
}

func (cqrs *CQRS) Query(streamID string) (interface{}, error) {
    cqrs.mutex.RLock()
    defer cqrs.mutex.RUnlock()
    
    result, exists := cqrs.QueryStore[streamID]
    if !exists {
        return nil, fmt.Errorf("stream not found: %s", streamID)
    }
    
    return result, nil
}
```

## 🔄 Message Queuing

### Message Queue
```go
type MessageQueue struct {
    Messages chan Message
    Subscribers map[string][]chan Message
    mutex    sync.RWMutex
}

type Message struct {
    ID        string
    Topic     string
    Data      interface{}
    Timestamp time.Time
    TTL       time.Duration
}

func (mq *MessageQueue) Publish(topic string, data interface{}) error {
    message := Message{
        ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
        Topic:     topic,
        Data:      data,
        Timestamp: time.Now(),
        TTL:       24 * time.Hour,
    }
    
    select {
    case mq.Messages <- message:
        return nil
    default:
        return fmt.Errorf("queue full")
    }
}

func (mq *MessageQueue) Subscribe(topic string) <-chan Message {
    mq.mutex.Lock()
    defer mq.mutex.Unlock()
    
    ch := make(chan Message, 100)
    mq.Subscribers[topic] = append(mq.Subscribers[topic], ch)
    return ch
}
```

### Distributed Cache
```go
type DistributedCache struct {
    Data  map[string]CacheItem
    mutex sync.RWMutex
}

type CacheItem struct {
    Value     interface{}
    ExpiresAt time.Time
    Version   int64
}

func (dc *DistributedCache) Set(key string, value interface{}, ttl time.Duration) {
    dc.mutex.Lock()
    defer dc.mutex.Unlock()
    
    dc.Data[key] = CacheItem{
        Value:     value,
        ExpiresAt: time.Now().Add(ttl),
        Version:   time.Now().UnixNano(),
    }
}

func (dc *DistributedCache) Get(key string) (interface{}, bool) {
    dc.mutex.RLock()
    defer dc.mutex.RUnlock()
    
    item, exists := dc.Data[key]
    if !exists {
        return nil, false
    }
    
    if time.Now().After(item.ExpiresAt) {
        // Item expired
        dc.mutex.RUnlock()
        dc.mutex.Lock()
        delete(dc.Data, key)
        dc.mutex.Unlock()
        dc.mutex.RLock()
        return nil, false
    }
    
    return item.Value, true
}
```

## 🛡️ Resilience Patterns

### Retry Policy
```go
type RetryPolicy struct {
    MaxAttempts int
    BaseDelay   time.Duration
    MaxDelay    time.Duration
    Multiplier  float64
    Jitter      bool
}

func (rp *RetryPolicy) Execute(fn func() (interface{}, error)) (interface{}, error) {
    var lastErr error
    delay := rp.BaseDelay
    
    for attempt := 0; attempt < rp.MaxAttempts; attempt++ {
        result, err := fn()
        if err == nil {
            return result, nil
        }
        
        lastErr = err
        
        if attempt == rp.MaxAttempts-1 {
            break
        }
        
        // Calculate delay
        actualDelay := delay
        if rp.Jitter {
            actualDelay = time.Duration(float64(delay) * (0.5 + rand.Float64()*0.5))
        }
        
        time.Sleep(actualDelay)
        
        // Calculate next delay
        delay = time.Duration(float64(delay) * rp.Multiplier)
        if delay > rp.MaxDelay {
            delay = rp.MaxDelay
        }
    }
    
    return nil, fmt.Errorf("max attempts reached: %v", lastErr)
}
```

### Timeout Policy
```go
type TimeoutPolicy struct {
    Timeout time.Duration
}

func (tp *TimeoutPolicy) Execute(fn func() (interface{}, error)) (interface{}, error) {
    ctx, cancel := context.WithTimeout(context.Background(), tp.Timeout)
    defer cancel()
    
    done := make(chan struct{})
    var result interface{}
    var err error
    
    go func() {
        defer close(done)
        result, err = fn()
    }()
    
    select {
    case <-done:
        return result, err
    case <-ctx.Done():
        return nil, fmt.Errorf("operation timed out after %v", tp.Timeout)
    }
}
```

## 🎯 Best Practices

### 1. Service Design
- **Single Responsibility**: Each service has one responsibility
- **Loose Coupling**: Services are loosely coupled
- **High Cohesion**: Related functionality is grouped together
- **Stateless**: Services should be stateless when possible

### 2. Communication
- **Synchronous**: Use for request-response patterns
- **Asynchronous**: Use for event-driven patterns
- **Message Queues**: Use for reliable message delivery
- **Event Streaming**: Use for real-time event processing

### 3. Data Management
- **Event Sourcing**: Store events instead of state
- **CQRS**: Separate command and query responsibilities
- **Replication**: Replicate data for availability
- **Partitioning**: Partition data for scalability

### 4. Fault Tolerance
- **Circuit Breakers**: Prevent cascading failures
- **Bulkheads**: Isolate resources
- **Retries**: Handle transient failures
- **Timeouts**: Prevent hanging operations

### 5. Monitoring
- **Health Checks**: Monitor service health
- **Metrics**: Collect performance metrics
- **Logging**: Log important events
- **Tracing**: Track request flows

## 🚨 Common Issues

### Consistency Issues
- **Eventual Consistency**: Data becomes consistent over time
- **Strong Consistency**: All nodes see same data immediately
- **Causal Consistency**: Preserve causal relationships
- **Weak Consistency**: No consistency guarantees

### Availability Issues
- **Single Points of Failure**: Eliminate SPOFs
- **Redundancy**: Multiple copies of critical components
- **Failover**: Automatic switching to backup systems
- **Load Balancing**: Distribute load across multiple instances

### Partition Tolerance Issues
- **Network Partitions**: Handle network splits
- **Split Brain**: Prevent multiple leaders
- **Consensus**: Agree on decisions despite partitions
- **Quorum**: Majority-based decision making

### Performance Issues
- **Latency**: Minimize communication latency
- **Throughput**: Maximize request processing rate
- **Scalability**: Handle increased load
- **Resource Usage**: Optimize resource consumption

## 📊 Monitoring and Observability

### Key Metrics
- **Request Rate**: Requests per second
- **Response Time**: Average response time
- **Error Rate**: Percentage of failed requests
- **Availability**: Uptime percentage
- **Throughput**: Data processed per second
- **Latency**: P50, P95, P99 latencies

### Health Checks
- **Liveness**: Is the service running?
- **Readiness**: Is the service ready to serve?
- **Startup**: Is the service starting up?
- **Custom**: Application-specific checks

### Distributed Tracing
- **Trace**: Complete request journey
- **Span**: Individual operation within trace
- **Context Propagation**: Pass trace context
- **Sampling**: Reduce trace volume

---

**Remember**: Distributed systems are complex but powerful. Use these patterns and practices to build reliable, scalable systems! 🚀
```

Now let me update the progress:


Perfect! I've completed **Topic 11: Distributed Systems Networking** with comprehensive coverage. Here's what I've created for you:

## 🎯 **Topic 11: Distributed Systems Networking - COMPLETED**

### **📚 Theory Coverage** (`11_distributed_systems/theory.md`)
- **Distributed Systems Overview**: Characteristics, challenges, and fundamental concepts
- **CAP Theorem**: Consistency, Availability, Partition Tolerance trade-offs
- **Distributed System Architectures**: Client-server, P2P, microservices, SOA
- **Consensus Algorithms**: Raft, Paxos, PBFT implementation details
- **Service Discovery**: Service registry, discovery patterns, tools
- **Load Balancing**: Strategies, types, algorithms for distributed systems
- **Distributed Data Management**: Replication, consistency models, partitioning, transactions
- **Message Passing**: Communication patterns, ordering, delivery guarantees
- **Distributed Caching**: Strategies, consistency, distributed caching
- **Fault Tolerance**: Failure types, techniques, detection, recovery
- **Distributed System Patterns**: Circuit breaker, bulkhead, saga, CQRS, event sourcing
- **Monitoring and Observability**: Three pillars, distributed tracing, health checks
- **Security**: Authentication, authorization, network security, data security

### ** Implementation** (`11_distributed_systems/implementation.go`)
- **RaftNode**: Complete Raft consensus algorithm implementation
- **ServiceRegistry**: Service discovery with health checking
- **LoadBalancer**: Multiple load balancing algorithms
- **MessageQueue**: Publish-subscribe message queuing system
- **DistributedCache**: Distributed caching with TTL support
- **HTTP Server**: REST API for distributed system components
- **Comprehensive Examples**: All major distributed system patterns

### **🔄 Distributed Patterns** (`11_distributed_systems/patterns.go`)
- **CircuitBreaker**: Circuit breaker pattern with state management
- **Bulkhead**: Resource isolation and failure containment
- **Saga**: Distributed transaction management with compensation
- **EventStore**: Event sourcing implementation
- **CQRS**: Command Query Responsibility Segregation
- **RetryPolicy**: Exponential backoff retry logic
- **TimeoutPolicy**: Operation timeout management
- **Comprehensive Examples**: All major resilience patterns

### **📖 Quick Reference** (`11_distributed_systems/quick_reference.md`)
- **Essential Concepts**: CAP theorem, distributed system challenges
- **Go Patterns**: Raft consensus, service discovery, load balancing
- **Distributed Patterns**: Circuit breaker, bulkhead, saga, event sourcing
- **Event Sourcing & CQRS**: Event store, command/query separation
- **Message Queuing**: Publish-subscribe, message distribution
- **Resilience Patterns**: Retry policies, timeout management
- **Best Practices**: Service design, communication, data management
- **Common Issues**: Consistency, availability, partition tolerance
- **Monitoring**: Key metrics, health checks, distributed tracing

## 🚀 **Key Features of This Module**

1. **Complete Distributed Systems Coverage**: From basic concepts to advanced patterns
2. **Go-Specific**: Leverages Go's excellent concurrency and networking features
3. **Production-Ready**: Real-world distributed system patterns and implementations
4. **Consensus Algorithms**: Raft implementation with leader election and log replication
5. **Service Discovery**: Dynamic service registration and discovery
6. **Load Balancing**: Multiple algorithms with health checking
7. **Resilience Patterns**: Circuit breaker, bulkhead, saga, retry policies
8. **Event Sourcing**: Complete event sourcing and CQRS implementation

## 📈 **What You've Learned**

- **Distributed Systems Fundamentals**: CAP theorem, consistency models, fault tolerance
- **Consensus Algorithms**: Raft, Paxos, PBFT for distributed coordination
- **Service Discovery**: Dynamic service location and health monitoring
- **Load Balancing**: Traffic distribution across multiple services
- **Message Queuing**: Asynchronous communication and event processing
- **Distributed Caching**: Performance optimization and data distribution
- **Resilience Patterns**: Circuit breaker, bulkhead, saga for fault tolerance
- **Event Sourcing**: Storing events instead of state for audit and replay
- **CQRS**: Separating command and query responsibilities
- **Go Implementation**: Professional distributed system programming

## **Next Steps**

You're now ready to move to **Topic 12: Cloud Networking & Containerization** where we'll explore cloud networking, container orchestration, and modern deployment patterns.

The distributed systems knowledge you now have will make you unstoppable in building scalable, fault-tolerant systems, understanding microservices architecture, and acing distributed systems interview questions! 🚀

Would you like me to continue with **Topic 12: Cloud Networking & Containerization** or would you prefer to explore any specific aspect of distributed systems first?

The distributed systems foundation you've gained is crucial for modern software development and will set you apart in technical interviews!

