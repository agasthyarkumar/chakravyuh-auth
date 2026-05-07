package handlers

import (
	"net/http"
	"strconv"

	"auth-service/internal/dto"
	"auth-service/internal/repository"
	"auth-service/internal/services"
	"github.com/gin-gonic/gin"
)

// CreateAPIKey handles POST /admin/api-keys
func CreateAPIKey(c *gin.Context) {
	// Get tenant_id from context (same as admin_handler pattern)
	tenantIDInterface, exists := c.Get("tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "tenant not found in context",
		})
		return
	}

	tenantIDFloat, ok := tenantIDInterface.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid tenant_id",
		})
		return
	}

	tenantID := uint(tenantIDFloat)

	var req services.CreateAPIKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// Create API key
	response, err := services.CreateAPIKey(tenantID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "failed to create API key",
		})
		return
	}

	// Log API key creation
	userID, _ := c.Get("user_id")
	if userID != nil {
		_ = services.LogAction(
			tenantID,
			uint(userID.(float64)),
			"API_KEY_CREATED",
			"API key '"+req.Name+"' created",
		)
	}

	c.JSON(http.StatusCreated, response)
}

// ListAPIKeys handles GET /admin/api-keys
func ListAPIKeys(c *gin.Context) {
	// Get tenant_id from context (same pattern)
	tenantIDInterface, exists := c.Get("tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "tenant not found in context",
		})
		return
	}

	tenantIDFloat, ok := tenantIDInterface.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid tenant_id",
		})
		return
	}

	tenantID := uint(tenantIDFloat)

	// Get API keys
	keys, err := services.GetTenantAPIKeys(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "failed to retrieve API keys",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"api_keys": keys,
	})
}

// RevokeAPIKey handles DELETE /admin/api-keys/:id
func RevokeAPIKey(c *gin.Context) {
	tenantIDInterface, exists := c.Get("tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "tenant not found in context",
		})
		return
	}

	tenantIDFloat, ok := tenantIDInterface.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid tenant_id",
		})
		return
	}

	tenantID := uint(tenantIDFloat)

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid API key ID",
		})
		return
	}

	// Verify API key belongs to this tenant
	apiKey, err := repository.GetAPIKeyByID(uint(id))
	if err != nil || apiKey.TenantID != tenantID {
		c.JSON(http.StatusForbidden, dto.ErrorResponse{
			Error: "API key not found or unauthorized",
		})
		return
	}

	// Revoke the API key
	if err := services.RevokeAPIKey(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "failed to revoke API key",
		})
		return
	}

	// Log API key revocation
	userID, _ := c.Get("user_id")
	if userID != nil {
		_ = services.LogAction(
			tenantID,
			uint(userID.(float64)),
			"API_KEY_REVOKED",
			"API key '"+apiKey.Name+"' was revoked",
		)
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "API key revoked successfully",
	})
}
