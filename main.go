package main

import (
	"github.com/SGDIEGO/ApplicationHA/internal/core/router"
	"github.com/SGDIEGO/ApplicationHA/internal/domain"
	_ "github.com/SGDIEGO/ApplicationHA/internal/ports"
	repository "github.com/SGDIEGO/ApplicationHA/internal/repositories"
	"github.com/SGDIEGO/ApplicationHA/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	users := []domain.User{
		{Id: "1", Name: "user1", Email: "user1@gmail.com", Password: "user1password"},
		{Id: "2", Name: "user2", Email: "user2@gmail.com", Password: "user2password"},
		{Id: "3", Name: "user3", Email: "user3@gmail.com", Password: "user3password"},
		{Id: "4", Name: "user4", Email: "user4@gmail.com", Password: "user4password"},
	}

	userRepo := repository.NewMemKVS(users)
	userService := services.NewUserService(userRepo)
	// userHandler := handlers.NewUserHandler(userService)
	userRouter := router.NewUserRouter(userService)

	server := gin.Default()

	userRouter.ServeHttp("/users", server)

	server.Run(":3000")
}
