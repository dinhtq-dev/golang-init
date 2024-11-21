// repository/user_repository_interface.go
package repository

import (
	"golang/app/modules/auth/model"

	"gorm.io/gorm"
)

// UserRepository interface defines the methods for interacting with the user data.
type UserRepository interface {
	FindByID(id string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	FindAll(offset, limit int) ([]model.User, int64, error) 
	Update(user *model.User) (*model.User, error)
}

// userRepository is the struct that implements the UserRepository interface
type userRepository struct {
	db *gorm.DB
}
