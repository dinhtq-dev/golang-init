package service

import (
	"errors"
	"golang/app/modules/auth/model"
	"golang/app/modules/auth/repository"
	"golang/app/modules/auth/request"
)

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) GetByID(id string) (*model.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) Create(req request.CreateUserRequest) (*model.User, error) {
    var namePtr *string
    if req.Name != "" {
        namePtr = &req.Name
    }

    user := &model.User{
        Name:  namePtr, 
        Email: req.Email,
    }

    return s.userRepo.Create(user)
}


type PaginatedUsers struct {
    Users []model.User
    Total int64
}

func (s *userService) GetAll(page, perPage int) (*PaginatedUsers, error) {
    offset := (page - 1) * perPage

    users, total, err := s.userRepo.FindAll(offset, perPage)
    if err != nil {
        return nil, err
    }

    return &PaginatedUsers{
        Users: users,
        Total: total,
    }, nil
}

func (s *userService) Update(id string, req request.UpdateUserRequest) (*model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	user.Name = &req.Name
	user.Email = req.Email

	updatedUser, err := s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
