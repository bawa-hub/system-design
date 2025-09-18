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
