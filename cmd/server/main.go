package main

import (
	"net/http"
	"time"

	"github.com/anirudhjamakay/ratelimiter/internal/algorithms"
	"github.com/anirudhjamakay/ratelimiter/internal/middleware"
	"github.com/anirudhjamakay/ratelimiter/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {

	redisStore := store.NewRedisStore("redis:6379")

	limiter := algorithms.NewFixedWindow(redisStore, 5, time.Minute)

	router := gin.Default()

	router.Use(middleware.RateLimiterMiddleware(limiter))

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(":8080")
}
