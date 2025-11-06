package app

import (
	"log"
	"user-manage-backend/internal/configs"
	"user-manage-backend/internal/db"
	"user-manage-backend/internal/routes"
	"user-manage-backend/internal/validations"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Module interface {
	Routes() routes.Route
}
type Application struct {
	config *configs.Config
	engine *gin.Engine
}

func NewApplication(config *configs.Config) *Application {
	if err := validations.InitValidations(); err != nil {
		log.Fatalf("failed to init validations: %v", err)
	}
	engine := gin.Default()
	loadEnv()

	modules := []Module{
		NewUserModule(db.DB),
	}
	routes.RegisterRoutes(engine, getModuleRoutes(modules)...)
	return &Application{
		config: config,
		engine: engine,
	}
}

func (app *Application) Run() error {
	return app.engine.Run(app.config.ServerAddress)
}
func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))
	for i, module := range modules {
		routeList[i] = module.Routes()

	}
	return routeList
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("failed to load .env:", err)
	}
}
