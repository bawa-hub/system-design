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
