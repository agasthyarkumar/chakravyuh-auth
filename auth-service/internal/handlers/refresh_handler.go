package handlers

import (
	"auth-service/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshInput struct {
	RefreshToken string `json:"refresh_token"`
}

func Refresh(c *gin.Context) {
	var input RefreshInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})

		return
	}

	accessToken, err := services.Refresh(
		input.RefreshToken,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

func Logout(c *gin.Context) {
	var input RefreshInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})

		return
	}

	// Extract user context
	userID, _ := c.Get("user_id")
	tenantID, _ := c.Get("tenant_id")

	err := services.Logout(input.RefreshToken)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Logout failed",
		})

		return
	}

	// Log user logout
	if userID != nil && tenantID != nil {
		uid := uint(userID.(float64))
		tid := uint(tenantID.(float64))
		_ = services.LogAction(
			tid,
			uid,
			"USER_LOGOUT",
			"User logged out",
		)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}