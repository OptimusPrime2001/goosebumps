package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

type Client struct {
	limiter         *rate.Limiter
	lastRequestTime time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*Client)
)

func getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}
	return ip
}
func getRateLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	// TODO: implement rate limit
	client, exists := clients[ip]
	if !exists {
		client = &Client{
			limiter:         rate.NewLimiter(5, 15), // 5 req/sec, burst 15
			lastRequestTime: time.Now(),
		}
		clients[ip] = client
	}
	client.lastRequestTime = time.Now()
	return client.limiter
}

func CleanupRateLimiters() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.lastRequestTime) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
		// TODO: implement rate limit
	}
	// TODO: implement rate limit
}

func RateLimitMiddleware(logger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement rate limit
		ip := getClientIP(ctx)
		limiter := getRateLimiter(ip)
		if !limiter.Allow() {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"code": "429",
				"msg":  "too many requests",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
