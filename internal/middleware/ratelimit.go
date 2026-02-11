package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RateLimiterMiddleware(ratelimiter interface {
	Allow(ctx context.Context, key string) (bool, error)
}) gin.HandlerFunc {

	return func(c *gin.Context) {

		// Use client IP as key
		key := c.ClientIP()

		allowed, err := ratelimiter.Allow(c.Request.Context(), key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal error",
			})
			c.Abort()
			return
		}

		if !allowed {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
