package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiter defines rate limits for different endpoints
type RateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.RWMutex
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		limiters: make(map[string]*rate.Limiter),
	}
}

// GetLimiter returns or creates a limiter for an endpoint
func (rl *RateLimiter) GetLimiter(endpoint string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limiter, exists := rl.limiters[endpoint]
	if !exists {
		// Get the limit for this endpoint
		limit, burst := getEndpointLimits(endpoint)
		limiter = rate.NewLimiter(limit, burst)
		rl.limiters[endpoint] = limiter
	}

	return limiter
}

// getEndpointLimits returns rate limit and burst for an endpoint
func getEndpointLimits(endpoint string) (rate.Limit, int) {
	switch endpoint {
	case "login":
		return rate.Limit(5.0 / 60), 5 // 5 per minute
	case "register":
		return rate.Limit(3.0 / 60), 3 // 3 per minute
	case "accept-invite":
		return rate.Limit(10.0 / 60), 10 // 10 per minute
	case "admin":
		return rate.Limit(60.0 / 60), 60 // 60 per minute
	default:
		return rate.Limit(100.0 / 60), 100 // 100 per minute (default)
	}
}

// RateLimitMiddleware applies rate limiting based on endpoint
func RateLimitMiddleware(rl *RateLimiter) gin.HandlerFunc {
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

		limiter := rl.GetLimiter(endpoint)

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
