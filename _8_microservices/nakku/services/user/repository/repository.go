package repository

import (
	"context"

	userModels "nakku/services/user/models"

	"gorm.io/gorm"
)

// UserRepository handles user data operations
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create creates a new user
func (r *UserRepository) Create(ctx context.Context, user *userModels.User) (*userModels.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetByID gets a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id string) (*userModels.User, error) {
	var user userModels.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail gets a user by email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*userModels.User, error) {
	var user userModels.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername gets a user by username
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*userModels.User, error) {
	var user userModels.User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates a user
func (r *UserRepository) Update(ctx context.Context, user *userModels.User) (*userModels.User, error) {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Delete deletes a user
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&userModels.User{}, id).Error
}

// CreateAddress creates a new user address
func (r *UserRepository) CreateAddress(ctx context.Context, address *userModels.UserAddress) (*userModels.UserAddress, error) {
	if err := r.db.WithContext(ctx).Create(address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

// GetAddressByID gets an address by ID
func (r *UserRepository) GetAddressByID(ctx context.Context, id uint) (*userModels.UserAddress, error) {
	var address userModels.UserAddress
	if err := r.db.WithContext(ctx).First(&address, id).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

// GetAddressesByUserID gets all addresses for a user
func (r *UserRepository) GetAddressesByUserID(ctx context.Context, userID string) ([]*userModels.UserAddress, error) {
	var addresses []*userModels.UserAddress
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}

// UpdateAddress updates an address
func (r *UserRepository) UpdateAddress(ctx context.Context, address *userModels.UserAddress) (*userModels.UserAddress, error) {
	if err := r.db.WithContext(ctx).Save(address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

// DeleteAddress deletes an address
func (r *UserRepository) DeleteAddress(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&userModels.UserAddress{}, id).Error
}

// UnsetDefaultAddresses unsets all default addresses for a user
func (r *UserRepository) UnsetDefaultAddresses(ctx context.Context, userID string) error {
	return r.db.WithContext(ctx).Model(&userModels.UserAddress{}).
		Where("user_id = ? AND is_default = ?", userID, true).
		Update("is_default", false).Error
}
