package routes

import (
	"auth-service/internal/handlers"
	"auth-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	protected := router.Group("/")

	protected.Use(middleware.AuthMiddleware())

	protected.GET("/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Protected route",
		})
	})
}