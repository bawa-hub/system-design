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
