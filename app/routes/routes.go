package routes

import (
	"golang/app/container"
	"golang/app/middleware"
	authRouter "golang/app/modules/auth/router"

	"github.com/gin-gonic/gin"
)

func Setup(container *container.Container) *gin.Engine {
	return SetupRouter(container)
}

func SetupRouter(container *container.Container) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Recover())
	
	api := router.Group("/api")
	{
		authRouter.AuthRouter(api, container)
	}

	return router
}
