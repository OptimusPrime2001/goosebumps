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
func (userRoute *UserRoute) Register(routerGroup *gin.RouterGroup) {
	users := routerGroup.Group("/users")
	{
		users.POST("/create", userRoute.controller.CreateNewUser)
		users.GET("", userRoute.controller.GetAllUsers)
		users.GET("/:uuid", userRoute.controller.GetUserByUUID)
		users.PUT("/:uuid", userRoute.controller.UpdateUser)
		users.DELETE("/:uuid", userRoute.controller.DeleteUser)
	}
}
