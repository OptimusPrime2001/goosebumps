package app

import (
	"user-manage-backend/internal/controllers"
	repositories "user-manage-backend/internal/respositories"
	"user-manage-backend/internal/routes"
	"user-manage-backend/internal/services"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule() *UserModule {

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController)

	return &UserModule{
		routes: userRoute,
	}
}
func (userModule *UserModule) Routes() routes.Route {
	return userModule.routes
}
