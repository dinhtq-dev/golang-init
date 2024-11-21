package service

import (
	"golang/app/modules/auth/model"
	"golang/app/modules/auth/repository"
	"golang/app/modules/auth/request"
)

type UserService interface {
	GetByID(id string) (*model.User, error)
	Create(req request.CreateUserRequest) (*model.User, error)
	GetAll(page, perPage int) (*PaginatedUsers, error)
	Update(id string, req request.UpdateUserRequest) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}