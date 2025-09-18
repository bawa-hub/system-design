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
