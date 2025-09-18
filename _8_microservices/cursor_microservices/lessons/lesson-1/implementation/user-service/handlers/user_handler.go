package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"user-service/models"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	db       *sql.DB
	logger   *logrus.Logger
	validate *validator.Validate
}

// NewUserHandler creates a new user handler
func NewUserHandler(db *sql.DB, logger *logrus.Logger) *UserHandler {
	return &UserHandler{
		db:       db,
		logger:   logger,
		validate: validator.New(),
	}
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	
	// Parse JSON request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.WithError(err).Error("Failed to decode request body")
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON", err.Error())
		return
	}

	// Validate request
	if err := h.validate.Struct(req); err != nil {
		h.logger.WithError(err).Error("Validation failed")
		h.sendErrorResponse(w, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	// Check if user already exists
	var existingID string
	checkQuery := `SELECT id FROM users WHERE email = $1`
	err := h.db.QueryRow(checkQuery, req.Email).Scan(&existingID)
	if err != sql.ErrNoRows {
		if err != nil {
			h.logger.WithError(err).Error("Failed to check existing user")
			h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to create user", "Database error")
			return
		}
		h.sendErrorResponse(w, http.StatusConflict, "User already exists", "Email already registered")
		return
	}

	// Create new user
	user := models.NewUser(req.Email, req.FirstName, req.LastName)

	// Insert user into database
	query := `
		INSERT INTO users (id, email, first_name, last_name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING created_at, updated_at
	`
	
	err = h.db.QueryRow(query, user.ID, user.Email, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt).
		Scan(&user.CreatedAt, &user.UpdatedAt)
	
	if err != nil {
		h.logger.WithError(err).Error("Failed to create user")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to create user", "Database error")
		return
	}

	h.logger.WithFields(logrus.Fields{
		"user_id": user.ID,
		"email":   user.Email,
	}).Info("User created successfully")

	// Return success response
	response := models.UserResponse{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	}

	h.sendJSONResponse(w, http.StatusCreated, response)
}

// GetUser handles GET /users/{id}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	if userID == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", "User ID is required")
		return
	}

	// Query user from database
	query := `
		SELECT id, email, first_name, last_name, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	
	var user models.User
	err := h.db.QueryRow(query, userID).Scan(
		&user.ID, &user.Email, &user.FirstName, &user.LastName,
		&user.CreatedAt, &user.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			h.sendErrorResponse(w, http.StatusNotFound, "User not found", "User with given ID does not exist")
			return
		}
		h.logger.WithError(err).Error("Failed to get user")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get user", "Database error")
		return
	}

	h.logger.WithField("user_id", userID).Info("User retrieved successfully")

	// Return success response
	response := models.UserResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    &user,
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// GetAllUsers handles GET /users
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	page := h.getQueryParam(r, "page", "1")
	limit := h.getQueryParam(r, "limit", "10")
	search := h.getQueryParam(r, "search", "")

	// Convert page and limit to integers
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limitNum < 1 || limitNum > 100 {
		limitNum = 10
	}

	offset := (pageNum - 1) * limitNum

	// Build query
	var query string
	var args []interface{}
	argIndex := 1

	if search != "" {
		query = `
			SELECT id, email, first_name, last_name, created_at, updated_at
			FROM users
			WHERE first_name ILIKE $1 OR last_name ILIKE $1 OR email ILIKE $1
			ORDER BY created_at DESC
			LIMIT $2 OFFSET $3
		`
		args = []interface{}{"%" + search + "%", limitNum, offset}
	} else {
		query = `
			SELECT id, email, first_name, last_name, created_at, updated_at
			FROM users
			ORDER BY created_at DESC
			LIMIT $1 OFFSET $2
		`
		args = []interface{}{limitNum, offset}
	}

	// Query users from database
	rows, err := h.db.Query(query, args...)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get users")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get users", "Database error")
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID, &user.Email, &user.FirstName, &user.LastName,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			h.logger.WithError(err).Error("Failed to scan user")
			continue
		}
		users = append(users, user)
	}

	// Get total count
	var total int
	countQuery := "SELECT COUNT(*) FROM users"
	if search != "" {
		countQuery = "SELECT COUNT(*) FROM users WHERE first_name ILIKE $1 OR last_name ILIKE $1 OR email ILIKE $1"
		err = h.db.QueryRow(countQuery, "%"+search+"%").Scan(&total)
	} else {
		err = h.db.QueryRow(countQuery).Scan(&total)
	}
	if err != nil {
		h.logger.WithError(err).Error("Failed to get total count")
		total = len(users)
	}

	h.logger.WithFields(logrus.Fields{
		"count": len(users),
		"page":  pageNum,
		"limit": limitNum,
	}).Info("Users retrieved successfully")

	// Return success response
	response := models.UsersResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Data:    users,
		Total:   total,
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// UpdateUser handles PUT /users/{id}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	if userID == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", "User ID is required")
		return
	}

	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.WithError(err).Error("Failed to decode request body")
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON", err.Error())
		return
	}

	// Validate request
	if err := h.validate.Struct(req); err != nil {
		h.logger.WithError(err).Error("Validation failed")
		h.sendErrorResponse(w, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	// Check if user exists
	var existingUser models.User
	checkQuery := `
		SELECT id, email, first_name, last_name, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	err := h.db.QueryRow(checkQuery, userID).Scan(
		&existingUser.ID, &existingUser.Email, &existingUser.FirstName, &existingUser.LastName,
		&existingUser.CreatedAt, &existingUser.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			h.sendErrorResponse(w, http.StatusNotFound, "User not found", "User with given ID does not exist")
			return
		}
		h.logger.WithError(err).Error("Failed to check existing user")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to update user", "Database error")
		return
	}

	// Check if email is being updated and if it already exists
	if req.Email != nil && *req.Email != existingUser.Email {
		var existingID string
		emailCheckQuery := `SELECT id FROM users WHERE email = $1`
		err = h.db.QueryRow(emailCheckQuery, *req.Email).Scan(&existingID)
		if err != sql.ErrNoRows {
			if err != nil {
				h.logger.WithError(err).Error("Failed to check email uniqueness")
				h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to update user", "Database error")
				return
			}
			h.sendErrorResponse(w, http.StatusConflict, "Email already exists", "Email is already registered")
			return
		}
	}

	// Update user
	existingUser.Update(req)

	// Build dynamic update query
	setParts := []string{}
	args := []interface{}{}
	argIndex := 1

	if req.Email != nil {
		setParts = append(setParts, fmt.Sprintf("email = $%d", argIndex))
		args = append(args, *req.Email)
		argIndex++
	}
	if req.FirstName != nil {
		setParts = append(setParts, fmt.Sprintf("first_name = $%d", argIndex))
		args = append(args, *req.FirstName)
		argIndex++
	}
	if req.LastName != nil {
		setParts = append(setParts, fmt.Sprintf("last_name = $%d", argIndex))
		args = append(args, *req.LastName)
		argIndex++
	}

	if len(setParts) == 0 {
		h.sendErrorResponse(w, http.StatusBadRequest, "No fields to update", "At least one field must be provided")
		return
	}

	// Add updated_at and user_id
	setParts = append(setParts, fmt.Sprintf("updated_at = $%d", argIndex))
	args = append(args, time.Now())
	argIndex++
	args = append(args, userID)

	query := fmt.Sprintf(`
		UPDATE users 
		SET %s
		WHERE id = $%d
		RETURNING id, email, first_name, last_name, created_at, updated_at
	`, strings.Join(setParts, ", "), argIndex)

	var user models.User
	err = h.db.QueryRow(query, args...).Scan(
		&user.ID, &user.Email, &user.FirstName, &user.LastName,
		&user.CreatedAt, &user.UpdatedAt,
	)
	
	if err != nil {
		h.logger.WithError(err).Error("Failed to update user")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to update user", "Database error")
		return
	}

	h.logger.WithField("user_id", userID).Info("User updated successfully")

	// Return success response
	response := models.UserResponse{
		Success: true,
		Message: "User updated successfully",
		Data:    &user,
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// DeleteUser handles DELETE /users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	if userID == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", "User ID is required")
		return
	}

	// Check if user exists
	var existingUser models.User
	checkQuery := `
		SELECT id, email, first_name, last_name, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	err := h.db.QueryRow(checkQuery, userID).Scan(
		&existingUser.ID, &existingUser.Email, &existingUser.FirstName, &existingUser.LastName,
		&existingUser.CreatedAt, &existingUser.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			h.sendErrorResponse(w, http.StatusNotFound, "User not found", "User with given ID does not exist")
			return
		}
		h.logger.WithError(err).Error("Failed to check existing user")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to delete user", "Database error")
		return
	}

	// Delete user from database
	query := `DELETE FROM users WHERE id = $1`
	result, err := h.db.Exec(query, userID)
	if err != nil {
		h.logger.WithError(err).Error("Failed to delete user")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to delete user", "Database error")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		h.logger.WithError(err).Error("Failed to get rows affected")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to delete user", "Database error")
		return
	}

	if rowsAffected == 0 {
		h.sendErrorResponse(w, http.StatusNotFound, "User not found", "User with given ID does not exist")
		return
	}

	h.logger.WithField("user_id", userID).Info("User deleted successfully")

	// Return success response
	response := models.UserResponse{
		Success: true,
		Message: "User deleted successfully",
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// Helper methods

func (h *UserHandler) sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (h *UserHandler) sendErrorResponse(w http.ResponseWriter, statusCode int, message, error string) {
	response := models.ErrorResponse{
		Success: false,
		Message: message,
		Error:   error,
	}
	h.sendJSONResponse(w, statusCode, response)
}

func (h *UserHandler) getQueryParam(r *http.Request, key, defaultValue string) string {
	if value := r.URL.Query().Get(key); value != "" {
		return value
	}
	return defaultValue
}
