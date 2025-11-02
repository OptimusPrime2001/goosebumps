package routes

import (
	"user-manage-backend/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type Route interface {
	Register(routerGroup *gin.RouterGroup)
}

func RegisterRoutes(engine *gin.Engine, routes ...Route) {
	engine.Use(middlewares.LoggerMiddleware(), middlewares.RateLimitMiddleware(), middlewares.ApiKeyMiddleware())
	routerGroup := engine.Group("/api/v1")
	for _, route := range routes {
		route.Register(routerGroup)
	}
}
