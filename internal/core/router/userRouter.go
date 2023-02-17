package router

import (
	. "github.com/SGDIEGO/ApplicationHA/internal/core/handlers"
	"github.com/SGDIEGO/ApplicationHA/internal/ports"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	userHandler UserHandler
}

func NewUserRouter(userService ports.UserServiceI) *UserRouter {

	// var userHandler = handlers.NewUserHandler(userService)

	return &UserRouter{
		userHandler: *UserHDL(userService),
	}
}

func (uR *UserRouter) ServeHttp(relativeURL string, server *gin.Engine) {

	router := server.Group(relativeURL)
	{
		router.GET("/", uR.userHandler.GetAllUsers)
		router.GET("/:id", uR.userHandler.GetUser)
		router.POST("/", uR.userHandler.CreateUser)
		router.PATCH("/:id", uR.userHandler.EditUser)
		router.DELETE("/:id", uR.userHandler.DeleteUser)
	}

}
