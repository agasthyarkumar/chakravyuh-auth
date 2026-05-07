package handlers

import (
	"net/http"
	"strconv"

	"auth-service/internal/services"
	"github.com/gin-gonic/gin"
)

// CreateAPIKey handles POST /admin/api-keys
func CreateAPIKey(c *gin.Context) {
	// Get tenant_id from context
	tenantIDInterface, exists := c.Get("tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "tenant not found in context",
		})
		return
	}

	tenantID, ok := tenantIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid tenant_id",
		})
		return
	}

	var req services.CreateAPIKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create API key
	response, err := services.CreateAPIKey(tenantID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create API key",
		})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// ListAPIKeys handles GET /admin/api-keys
func ListAPIKeys(c *gin.Context) {
	// Get tenant_id from context
	tenantIDInterface, exists := c.Get("tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "tenant not found in context",
		})
		return
	}

	tenantID, ok := tenantIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid tenant_id",
		})
		return
	}

	// Get API keys
	keys, err := services.GetTenantAPIKeys(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve API keys",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"api_keys": keys,
	})
}

// RevokeAPIKey handles DELETE /admin/api-keys/:id
func RevokeAPIKey(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid API key ID",
		})
		return
	}

	// Revoke the API key
	if err := services.RevokeAPIKey(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to revoke API key",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "API key revoked successfully",
	})
}
