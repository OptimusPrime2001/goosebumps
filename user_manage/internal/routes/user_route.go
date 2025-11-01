package routes

import (
	"user-manage-backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	controller *controllers.UserController
}

func NewUserRoute(controller *controllers.UserController) *UserRoute {
	return &UserRoute{
		controller: controller,
	}
}
func (userRoute *UserRoute) Register(server *gin.RouterGroup) {
	users := server.Group("/users")
	{
		users.POST("", userRoute.controller.CreateUser)
	}
}
