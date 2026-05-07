package handlers

import (
	"auth-service/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTenantUsers(c *gin.Context) {
	tenantIDValue, exists := c.Get("tenant_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Tenant ID missing",
		})

		return
	}

	tenantIDFloat, ok := tenantIDValue.(float64)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid tenant ID",
		})

		return
	}

	tenantID := uint(tenantIDFloat)

	users, err := services.ListTenantUsers(
		tenantID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
		})

		return
	}

	c.JSON(http.StatusOK, users)
}