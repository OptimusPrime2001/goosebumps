package app

import (
	"user-manage-backend/internal/controllers"
	"user-manage-backend/internal/db/sqlc"
	repositories "user-manage-backend/internal/respositories"
	"user-manage-backend/internal/routes"
	"user-manage-backend/internal/services"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule(db *sqlc.Queries) *UserModule {

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
