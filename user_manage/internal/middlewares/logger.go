package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func LoggerMiddleware(logger *zerolog.Logger) gin.HandlerFunc {

	// TODO: implement logger
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start)
		statusCode := ctx.Writer.Status()
		var logEvent *zerolog.Event
		if statusCode >= 500 {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Warn()
		}
		logEvent.Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Str("query", ctx.Request.URL.RawQuery).
			Str("client_ip", ctx.ClientIP()).
			Str("user_agent", ctx.Request.UserAgent()).
			Str("host", ctx.Request.Host).
			Int("status_code", ctx.Writer.Status()).
			Int64("duration_ms", duration.Milliseconds()).Msg("HTTP Request log")
	}
}
