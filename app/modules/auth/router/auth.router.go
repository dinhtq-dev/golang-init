package userRouter

import (
	"golang/app/container"
	"golang/app/middleware"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
	Name        string
	Middleware  []gin.HandlerFunc
}

func AuthRouter(router *gin.RouterGroup, container *container.Container) {
	v1 := router.Group("/v1")
	{
		userController := container.GetUserController()

		routes := []Route{
			{Method: "POST", Path: "/register", HandlerFunc: userController.CreateUser, Name: "constants.PermissionUserCreate"},
			{Method: "GET", Path: "/users/:id", HandlerFunc: userController.GetUserDetail, Name: "constants.PermissionUserDetail", Middleware: []gin.HandlerFunc{middleware.AuthMiddleware()}},
			{Method: "GET", Path: "/users", HandlerFunc: userController.GetUserList, Name: "constants.PermissionUserList"},
			{Method: "PUT", Path: "/users/:id", HandlerFunc: userController.UpdateUser, Name: "constants.PermissionUserUpdate", Middleware: []gin.HandlerFunc{middleware.AuthMiddleware(), middleware.LoggingMiddleware()}},
		}

		for _, route := range routes {
			handlers := append(route.Middleware, route.HandlerFunc)
			
			v1.Handle(route.Method, route.Path, handlers...)
		}
	}
}
