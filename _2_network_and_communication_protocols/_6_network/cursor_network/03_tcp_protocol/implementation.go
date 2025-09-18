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
