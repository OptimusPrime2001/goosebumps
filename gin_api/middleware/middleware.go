package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SimpleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Check request before to backend")
		ctx.Next()
		log.Println("Check response after to backend")
	}
}
