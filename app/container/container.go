package container

import (
	"golang/app/modules/auth/controller"
	"golang/app/modules/auth/repository"
	"golang/app/modules/auth/service"

	"gorm.io/gorm"
)


type Container struct {
	DB *gorm.DB
}

func NewContainer(db *gorm.DB) *Container {
	return &Container{DB: db}
}

func (c *Container) GetUserController() *controller.UserController {
	userRepo := repository.NewUserRepository(c.DB)
	userService := service.NewUserService(userRepo)

	return controller.NewUserController(userService)
}
