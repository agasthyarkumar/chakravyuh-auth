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

	// =========================
	// PUBLIC ROUTES
	// =========================

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.POST("/refresh", handlers.Refresh)

	// =========================
	// PROTECTED ROUTES
	// =========================

	protected := router.Group("/")

	protected.Use(middleware.AuthMiddleware())

	protected.GET("/me", func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		username, _ := c.Get("username")
		role, _ := c.Get("role")

		c.JSON(200, gin.H{
			"user_id": userID,
			"username": username,
			"role": role,
		})
	})

	protected.POST("/logout", handlers.Logout)
}