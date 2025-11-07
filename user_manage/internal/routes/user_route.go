package routes

import (
	"fmt"
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
		users.GET("/panic", func(ctx *gin.Context) {
			test := make([]int, 0)
			fmt.Println(test[1])
		})
		users.GET("/:uuid", userRoute.controller.GetUserByUUID)
		users.PUT("/:uuid", userRoute.controller.UpdateUser)
		users.DELETE("/:uuid", userRoute.controller.DeleteUser)
	}
}
