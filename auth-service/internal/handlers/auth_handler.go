package handlers

import (
	"auth-service/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	TenantName string `json:"tenant_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})

		return
	}

	err := services.Register(
		input.TenantName,
		input.Username,
		input.Password,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tenant and admin user created successfully",
	})
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})

		return
	}

	accessToken, refreshToken, err := services.Login(
		input.Username,
		input.Password,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"refresh_token": refreshToken,
	})
}