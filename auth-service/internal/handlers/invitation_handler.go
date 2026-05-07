package handlers

import (
	"auth-service/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateInvitationInput struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type AcceptInvitationInput struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateInvitation(c *gin.Context) {
	tenantIDValue, exists := c.Get("tenant_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Tenant ID missing",
		})

		return
	}

	tenantIDFloat := tenantIDValue.(float64)

	tenantID := uint(tenantIDFloat)

	var input CreateInvitationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})

		return
	}

	token, err := services.CreateInvitation(
		tenantID,
		input.Email,
		input.Role,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create invitation",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"invitation_token": token,
	})
}

func AcceptInvitation(c *gin.Context) {
	var input AcceptInvitationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})

		return
	}

	err := services.AcceptInvitation(
		input.Token,
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
		"message": "Invitation accepted successfully",
	})
}