package routes

import (
	"user-manage-backend/internal/middlewares"
	"user-manage-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type Route interface {
	Register(routerGroup *gin.RouterGroup)
}

func RegisterRoutes(engine *gin.Engine, routes ...Route) {
	logPath := "./internal/logs/https.log"
	logger := utils.NewLoggerWithPath(logPath, "info")

	recoverLogPath := "./internal/logs/recover.log"
	recoverLogger := utils.NewLoggerWithPath(recoverLogPath, "error")

	rateLimiterLogger := utils.NewLoggerWithPath("./internal/logs/rate_limiter.log", "info")

	engine.Use(
		middlewares.RateLimitMiddleware(rateLimiterLogger),
		middlewares.CORSMiddleware(),
		middlewares.LoggerMiddleware(logger),
		middlewares.RecoveryMiddleware(recoverLogger),
		middlewares.ApiKeyMiddleware())
	routerGroup := engine.Group("/api/v1")
	for _, route := range routes {
		route.Register(routerGroup)
	}
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(404, gin.H{"error": "404 not found", "path": context.Request.URL.Path})
	})
}
