package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"nakku/services/user/middleware"
	"nakku/services/user/utils"
	userModels "nakku/services/user/models"
	"nakku/services/user/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService handles user business logic
type UserService struct {
	userRepo *repository.UserRepository
	producer *utils.KafkaProducer
	logger   *logrus.Logger
}

// NewUserService creates a new UserService
func NewUserService(userRepo *repository.UserRepository, producer *utils.KafkaProducer, logger *logrus.Logger) *UserService {
	return &UserService{
		userRepo: userRepo,
		producer: producer,
		logger:   logger,
	}
}

// Register registers a new user
func (s *UserService) Register(ctx context.Context, req *userModels.UserRegistrationRequest) (*userModels.User, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	if existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", req.Email)
	}

	// Check if username already exists
	existingUser, err = s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("failed to check existing username: %w", err)
	}
	if existingUser != nil {
		return nil, fmt.Errorf("username %s already exists", req.Username)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	newUser := &userModels.User{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Role:      "user",
		IsActive:  true,
	}

	createdUser, err := s.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Publish user registered event
	event := map[string]interface{}{
		"event_type": "userModels.registered",
		"user_id":    createdUser.ID,
		"email":      createdUser.Email,
		"username":   createdUser.Username,
		"timestamp":  time.Now().Unix(),
	}

	if err := s.producer.PublishMessage(ctx, "user-events", fmt.Sprintf("user-%d", createdUser.ID), event); err != nil {
		s.logger.WithError(err).Error("Failed to publish user registered event")
	}

	s.logger.WithField("user_id", createdUser.ID).Info("User registered successfully")
	return createdUser, nil
}

// Login authenticates a user
func (s *UserService) Login(ctx context.Context, req *userModels.UserLoginRequest) (*userModels.LoginResponse, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("invalid credentials")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Check if user is active
	if !user.IsActive {
		return nil, fmt.Errorf("user account is deactivated")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Update last login
	now := time.Now()
	user.LastLogin = &now
	if _, err := s.userRepo.Update(ctx, user); err != nil {
		s.logger.WithError(err).Error("Failed to update last login")
	}

	// Generate JWT token
	token, err := s.generateJWTToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// Publish user login event
	event := map[string]interface{}{
		"event_type": "userModels.login",
		"user_id":    user.ID,
		"email":      user.Email,
		"timestamp":  time.Now().Unix(),
	}

	if err := s.producer.PublishMessage(ctx, "user-events", fmt.Sprintf("user-%d", user.ID), event); err != nil {
		s.logger.WithError(err).Error("Failed to publish user login event")
	}

	s.logger.WithField("user_id", user.ID).Info("User logged in successfully")

	return &userModels.LoginResponse{
		User:  user.ToUserResponse(),
		Token: token,
	}, nil
}

// GetProfile gets user profile
func (s *UserService) GetProfile(ctx context.Context, userID string) (*userModels.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// UpdateProfile updates user profile
func (s *UserService) UpdateProfile(ctx context.Context, userID string, req *userModels.UserUpdateRequest) (*userModels.User, error) {
	// Get existing user
	existingUser, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Update fields if provided
	if req.FirstName != "" {
		existingUser.FirstName = req.FirstName
	}
	if req.LastName != "" {
		existingUser.LastName = req.LastName
	}
	if req.Phone != "" {
		existingUser.Phone = req.Phone
	}

	// Update user
	updatedUser, err := s.userRepo.Update(ctx, existingUser)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Publish user updated event
	event := map[string]interface{}{
		"event_type": "userModels.updated",
		"user_id":    updatedUser.ID,
		"timestamp":  time.Now().Unix(),
	}

	if err := s.producer.PublishMessage(ctx, "user-events", fmt.Sprintf("user-%d", updatedUser.ID), event); err != nil {
		s.logger.WithError(err).Error("Failed to publish user updated event")
	}

	s.logger.WithField("user_id", updatedUser.ID).Info("User profile updated successfully")
	return updatedUser, nil
}

// AddAddress adds a new address for the user
func (s *UserService) AddAddress(ctx context.Context, userID string, req *userModels.UserAddressRequest) (*userModels.UserAddress, error) {
	// Get user
	_, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// If this is set as default, unset other default addresses
	if req.IsDefault {
		if err := s.userRepo.UnsetDefaultAddresses(ctx, userID); err != nil {
			return nil, fmt.Errorf("failed to unset default addresses: %w", err)
		}
	}

	// Create address
	userIDUint, _ := strconv.ParseUint(userID, 10, 32)
	address := &userModels.UserAddress{
		UserID:       uint(userIDUint),
		Type:         req.Type,
		AddressLine1: req.AddressLine1,
		AddressLine2: req.AddressLine2,
		City:         req.City,
		State:        req.State,
		PostalCode:   req.PostalCode,
		Country:      req.Country,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
		IsDefault:    req.IsDefault,
	}

	createdAddress, err := s.userRepo.CreateAddress(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("failed to create address: %w", err)
	}

	s.logger.WithFields(logrus.Fields{
		"user_id":    userID,
		"address_id": createdAddress.ID,
	}).Info("Address added successfully")

	return createdAddress, nil
}

// GetAddresses gets all addresses for a user
func (s *UserService) GetAddresses(ctx context.Context, userID string) ([]*userModels.UserAddress, error) {
	addresses, err := s.userRepo.GetAddressesByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get addresses: %w", err)
	}

	return addresses, nil
}

// UpdateAddress updates an address
func (s *UserService) UpdateAddress(ctx context.Context, userID string, addressID uint, req *userModels.UserAddressRequest) (*userModels.UserAddress, error) {
	// Get existing address
	existingAddress, err := s.userRepo.GetAddressByID(ctx, addressID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("address not found")
		}
		return nil, fmt.Errorf("failed to get address: %w", err)
	}

	// Check if address belongs to user
	userIDUint, _ := strconv.ParseUint(userID, 10, 32)
	if existingAddress.UserID != uint(userIDUint) {
		return nil, fmt.Errorf("address not found")
	}

	// If this is set as default, unset other default addresses
	if req.IsDefault && !existingAddress.IsDefault {
		if err := s.userRepo.UnsetDefaultAddresses(ctx, userID); err != nil {
			return nil, fmt.Errorf("failed to unset default addresses: %w", err)
		}
	}

	// Update address fields
	existingAddress.Type = req.Type
	existingAddress.AddressLine1 = req.AddressLine1
	existingAddress.AddressLine2 = req.AddressLine2
	existingAddress.City = req.City
	existingAddress.State = req.State
	existingAddress.PostalCode = req.PostalCode
	existingAddress.Country = req.Country
	existingAddress.Latitude = req.Latitude
	existingAddress.Longitude = req.Longitude
	existingAddress.IsDefault = req.IsDefault

	// Update address
	updatedAddress, err := s.userRepo.UpdateAddress(ctx, existingAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to update address: %w", err)
	}

	s.logger.WithFields(logrus.Fields{
		"user_id":    userID,
		"address_id": updatedAddress.ID,
	}).Info("Address updated successfully")

	return updatedAddress, nil
}

// DeleteAddress deletes an address
func (s *UserService) DeleteAddress(ctx context.Context, userID string, addressID uint) error {
	// Get existing address
	existingAddress, err := s.userRepo.GetAddressByID(ctx, addressID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("address not found")
		}
		return fmt.Errorf("failed to get address: %w", err)
	}

	// Check if address belongs to user
	userIDUint, _ := strconv.ParseUint(userID, 10, 32)
	if existingAddress.UserID != uint(userIDUint) {
		return fmt.Errorf("address not found")
	}

	// Delete address
	if err := s.userRepo.DeleteAddress(ctx, addressID); err != nil {
		return fmt.Errorf("failed to delete address: %w", err)
	}

	s.logger.WithFields(logrus.Fields{
		"user_id":    userID,
		"address_id": addressID,
	}).Info("Address deleted successfully")

	return nil
}

// generateJWTToken generates a JWT token for the user
func (s *UserService) generateJWTToken(user *userModels.User) (string, error) {
	claims := &middleware.JWTClaims{
		UserID:   fmt.Sprintf("%d", user.ID),
		Email:    user.Email,
		Role:     user.Role,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("nakku-secret-key")) // This should come from config
}
