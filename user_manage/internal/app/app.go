package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-manage-backend/internal/configs"
	"user-manage-backend/internal/db"
	"user-manage-backend/internal/routes"
	"user-manage-backend/internal/validations"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
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
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	db.InitDb()
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

	server := &http.Server{
		Addr:    app.config.ServerAddress,
		Handler: app.engine,
	}
	quick := make(chan os.Signal, 1)
	signal.Notify(quick, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		log.Printf("✅ Server is running on %s", app.config.ServerAddress)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server listen and serve failed: %v", err)
		}
	}()
	<-quick
	log.Println("⚠️ Shutdown signal received")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("⚠️ Server force to shutdown failed: %v", err)
	}
	return nil
}
func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))
	for i, module := range modules {
		routeList[i] = module.Routes()

	}
	return routeList
}
