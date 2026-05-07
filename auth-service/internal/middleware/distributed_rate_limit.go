package middleware

import (
	"context"
	"net/http"
	"time"

	"auth-service/internal/redis"
	"github.com/gin-gonic/gin"
)

// DistributedRateLimiter uses Redis for rate limiting
type DistributedRateLimiter struct {
	enabled bool
}

// NewDistributedRateLimiter creates a new distributed rate limiter
func NewDistributedRateLimiter() *DistributedRateLimiter {
	return &DistributedRateLimiter{
		enabled: redis.IsConnected(),
	}
}

// CheckLimit checks if a request should be allowed
func (drl *DistributedRateLimiter) CheckLimit(key string, limit int, windowSeconds int) bool {
	if !drl.enabled {
		return true // If Redis not available, allow all
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Increment counter
	newCount, err := redis.IncrBy(ctx, key, 1)
	if err != nil {
		return true // On error, allow
	}

	// Set TTL on first request
	if newCount == 1 {
		_ = redis.ExpireAt(ctx, key, time.Now().Add(time.Duration(windowSeconds)*time.Second))
	}

	return newCount <= int64(limit)
}

// DistributedRateLimitMiddleware applies distributed rate limiting based on endpoint
func DistributedRateLimitMiddleware(drl *DistributedRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		var endpoint string

		// Determine endpoint category from path
		path := c.Request.URL.Path
		switch {
		case path == "/login":
			endpoint = "login"
		case path == "/register":
			endpoint = "register"
		case path == "/accept-invite":
			endpoint = "accept-invite"
		case len(path) > 6 && path[:6] == "/admin":
			endpoint = "admin"
		default:
			endpoint = "default"
		}

		// Get client IP
		clientIP := c.ClientIP()

		// Create rate limit key
		limitKey := "ratelimit:" + endpoint + ":" + clientIP

		// Get limit for endpoint
		limit, window := getEndpointLimits(endpoint)

		// Check if allowed
		if !drl.CheckLimit(limitKey, limit, window) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// getEndpointLimits returns rate limit and window for an endpoint
func getEndpointLimits(endpoint string) (int, int) {
	switch endpoint {
	case "login":
		return 5, 60 // 5 per minute
	case "register":
		return 3, 60 // 3 per minute
	case "accept-invite":
		return 10, 60 // 10 per minute
	case "admin":
		return 60, 60 // 60 per minute
	default:
		return 100, 60 // 100 per minute
	}
}
