package handlers

import (
	"net/http"
	"strconv"

	"nakku/services/user/models"
	userModels "nakku/services/user/models"
	"nakku/services/user/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userService *service.UserService
	logger      *logrus.Logger
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService *service.UserService, logger *logrus.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}

// Register handles user registration
func (h *UserHandler) Register(c *gin.Context) {
	var req userModels.UserRegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid registration request")
		c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid request data", http.StatusBadRequest))
		return
	}

	// Create user
	createdUser, err := h.userService.Register(c.Request.Context(), &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to register user")
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to register user", http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusCreated, models.SuccessResponse(createdUser.ToUserResponse(), "User registered successfully"))
}

// Login handles user login
func (h *UserHandler) Login(c *gin.Context) {
	var req userModels.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid login request")
		c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid request data", http.StatusBadRequest))
		return
	}

	// Authenticate user
	loginResp, err := h.userService.Login(c.Request.Context(), &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to login user")
		c.JSON(http.StatusUnauthorized, models.NewErrorResponse("Invalid credentials", http.StatusUnauthorized))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(loginResp, "Login successful"))
}

// GetProfile handles getting user profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.NewErrorResponse("User not authenticated", http.StatusUnauthorized))
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Invalid user ID", http.StatusInternalServerError))
		return
	}

	// Get user profile
	userProfile, err := h.userService.GetProfile(c.Request.Context(), userIDStr)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get user profile")
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to get user profile", http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(userProfile.ToUserResponse(), "Profile retrieved successfully"))
}

// UpdateProfile handles updating user profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.NewErrorResponse("User not authenticated", http.StatusUnauthorized))
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Invalid user ID", http.StatusInternalServerError))
		return
	}

	var req userModels.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid update request")
		c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid request data", http.StatusBadRequest))
		return
	}

	// Update user profile
	updatedUser, err := h.userService.UpdateProfile(c.Request.Context(), userIDStr, &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to update user profile")
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to update user profile", http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(updatedUser.ToUserResponse(), "Profile updated successfully"))
}

// AddAddress handles adding a new address
func (h *UserHandler) AddAddress(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.NewErrorResponse("User not authenticated", http.StatusUnauthorized))
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Invalid user ID", http.StatusInternalServerError))
		return
	}

	var req userModels.UserAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid address request")
		c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid request data", http.StatusBadRequest))
		return
	}

	// Add address
	address, err := h.userService.AddAddress(c.Request.Context(), userIDStr, &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to add address")
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to add address", http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusCreated, models.SuccessResponse(address.ToUserAddressResponse(), "Address added successfully"))
}

// GetAddresses handles getting user addresses
func (h *UserHandler) GetAddresses(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.NewErrorResponse("User not authenticated", http.StatusUnauthorized))
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Invalid user ID", http.StatusInternalServerError))
		return
	}

	// Get addresses
	addresses, err := h.userService.GetAddresses(c.Request.Context(), userIDStr)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get addresses")
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to get addresses", http.StatusInternalServerError))
		return
	}

	// Convert to response format
	var addressResponses []userModels.UserAddressResponse
	for _, addr := range addresses {
		addressResponses = append(addressResponses, addr.ToUserAddressResponse())
	}

	c.JSON(http.StatusOK, models.SuccessResponse(addressResponses, "Addresses retrieved successfully"))
}

// UpdateAddress handles updating an address
func (h *UserHandler) UpdateAddress(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.NewErrorResponse("User not authenticated", http.StatusUnauthorized))
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Invalid user ID", http.StatusInternalServerError))
		return
	}

	addressIDStr := c.Param("id")
	addressID, err := strconv.ParseUint(addressIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid address ID", http.StatusBadRequest))
		return
	}

	var req userModels.UserAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid address update request")
		c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid request data", http.StatusBadRequest))
		return
	}

	// Update address
	address, err := h.userService.UpdateAddress(c.Request.Context(), userIDStr, uint(addressID), &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to update address")
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to update address", http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(address.ToUserAddressResponse(), "Address updated successfully"))
}

// DeleteAddress handles deleting an address
func (h *UserHandler) DeleteAddress(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.NewErrorResponse("User not authenticated", http.StatusUnauthorized))
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Invalid user ID", http.StatusInternalServerError))
		return
	}

	addressIDStr := c.Param("id")
	addressID, err := strconv.ParseUint(addressIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid address ID", http.StatusBadRequest))
		return
	}

	// Delete address
	if err := h.userService.DeleteAddress(c.Request.Context(), userIDStr, uint(addressID)); err != nil {
		h.logger.WithError(err).Error("Failed to delete address")
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to delete address", http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(nil, "Address deleted successfully"))
}
