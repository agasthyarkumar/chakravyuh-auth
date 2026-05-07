package handlers

import (
	"auth-service/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPendingTenants(c *gin.Context) {
	tenants, err := services.ListPendingTenants()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch tenants",
		})

		return
	}

	c.JSON(http.StatusOK, tenants)
}

func ApproveTenant(c *gin.Context) {
	idParam := c.Param("id")

	idInt, err := strconv.Atoi(
		idParam,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid tenant ID",
		})

		return
	}

	err = services.ApproveTenant(
		uint(idInt),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to approve tenant",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tenant approved successfully",
	})
}

func GetAuditLogs(c *gin.Context) {
	logs, err := services.GetAuditLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch audit logs",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"audit_logs": logs,
	})
}