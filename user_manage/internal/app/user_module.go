package app

import (
	"user-manage-backend/internal/controllers"
	repositories "user-manage-backend/internal/respositories"
	"user-manage-backend/internal/routes"
	"user-manage-backend/internal/services"

	"gorm.io/gorm"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule(db *gorm.DB) *UserModule {

	userRepo := repositories.NewUserRepository(db)
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
