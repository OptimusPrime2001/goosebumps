package routes

import "github.com/gin-gonic/gin"

type Route interface {
	Register(server *gin.RouterGroup)
}

func RegisterRoutes(server *gin.Engine, routes ...Route) {
	api := server.Group("/api/v1")
	for _, route := range routes {
		route.Register(api)
	}
}
