package main

import (
	"user-manage-backend/internal/configs"
	"user-manage-backend/internal/controllers"
	repositories "user-manage-backend/internal/respositories"
	"user-manage-backend/internal/routes"
	"user-manage-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.NewConfig()
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController)

	server := gin.Default()
	routes.RegisterRoutes(server, userRoute)
	server.Run(config.ServerAddress)
}
