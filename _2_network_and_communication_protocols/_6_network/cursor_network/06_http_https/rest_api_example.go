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
