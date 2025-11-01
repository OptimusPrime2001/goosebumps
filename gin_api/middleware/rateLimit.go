package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
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
		newClient := &Client{
			limiter:         rate.NewLimiter(5, 15), // 5 request/sec, brust 10
			lastRequestTime: time.Now(),
		}
		clients[ip] = newClient
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

// hey -n 20 -c 1 -H "X-API-Key:ab2a7a8a-d601-4bf7-b0e2-dd00e5459392" http://localhost:8080/api/v1/categories/golang
func RateLimitMiddleware() gin.HandlerFunc {
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
