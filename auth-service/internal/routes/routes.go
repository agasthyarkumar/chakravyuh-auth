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
	router.POST("/accept-invite", handlers.AcceptInvitation)

	// =========================
	// AUTHENTICATED ROUTES
	// =========================

	protected := router.Group("/")

	protected.Use(middleware.AuthMiddleware())

	protected.GET("/me", func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		tenantID, _ := c.Get("tenant_id")
		username, _ := c.Get("username")
		role, _ := c.Get("role")

		c.JSON(200, gin.H{
			"user_id": userID,
			"tenant_id": tenantID,
			"username": username,
			"role": role,
		})
	})

	protected.POST("/logout", handlers.Logout)

	// =========================
	// ADMIN ROUTES
	// =========================

	admin := protected.Group("/admin")

	admin.Use(
		middleware.RequireRole(
			"admin",
			"superadmin",
		),
	)

	admin.GET(
		"/dashboard",
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome Admin",
			})
		},
	)

	admin.GET(
		"/users",
		handlers.GetTenantUsers,
	)

	admin.POST(
		"/users",
		handlers.CreateTenantUser,
	)

	admin.DELETE(
		"/users/:id",
		handlers.DeleteTenantUser,
	)

	admin.POST(
		"/invitations",
		handlers.CreateInvitation,
	)

	// =========================
	// SUPERADMIN ROUTES
	// =========================

	superadmin := protected.Group("/superadmin")

	superadmin.Use(
		middleware.RequireRole(
			"superadmin",
		),
	)

	superadmin.GET(
		"/analytics",
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome SuperAdmin",
			})
		},
	)
}