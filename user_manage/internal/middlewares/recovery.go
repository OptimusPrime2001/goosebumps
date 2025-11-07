package middlewares

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"user-manage-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func RecoveryMiddleware(recoverLogger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				stack := ExtractFirstAppStackLine(debug.Stack())
				recoverLogger.Error().
					Str("path", ctx.Request.URL.Path).
					Str("method", ctx.Request.Method).Str("client_ip", ctx.ClientIP()).
					Str("panic", fmt.Sprintf("%v", err)).
					Str("stack", string(stack)).Msg("Panic Server Error")
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":  utils.ErrCodeInternal,
					"error": "Internal Server Error",
				})
			}
		}()
		ctx.Next()
	}
}

func ExtractFirstAppStackLine(stack []byte) string {
	lines := bytes.Split(stack, []byte("\n"))
	for _, line := range lines {
		if bytes.Contains(line, []byte(".go")) &&
			!bytes.Contains(line, []byte("/runtime/")) &&
			!bytes.Contains(line, []byte("/debug/")) &&
			!bytes.Contains(line, []byte("recovery.go")) {
			cleanLine := bytes.TrimSpace(line)
			return string(cleanLine)
		}
	}
	return ""
}
