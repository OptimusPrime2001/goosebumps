package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func validateApiKey(ctx *gin.Context) {
	apiKey := ctx.GetHeader("x-api-key")
	if apiKey == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "401",
			"msg":  "API key is required",
		})
		ctx.Abort()
		return
	}
	if apiKey != os.Getenv("API_KEY") {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": "401",
			"msg":  "Unauthorized",
		})
		ctx.Abort()
		return
	}
}
func ApiKeyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validateApiKey(ctx)
		ctx.Next()
	}
}
