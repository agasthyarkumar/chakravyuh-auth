package middleware

import (
	"net/http"
	"strings"

	"auth-service/internal/repository"
	"auth-service/internal/utils"
	"github.com/gin-gonic/gin"
)

// APIKeyMiddleware authenticates requests using API key
func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract X-API-Key header
		apiKey := c.GetHeader("X-API-Key")

		// If no API key, continue to next middleware (JWT will handle it)
		if apiKey == "" {
			c.Next()
			return
		}

		// Validate API key format
		if err := utils.ValidateAPIKeyFormat(apiKey); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid API key format",
			})
			c.Abort()
			return
		}

		// Hash the API key
		keyHash := utils.HashAPIKey(apiKey)

		// Lookup in database
		keyRecord, err := repository.GetAPIKey(keyHash)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid API key",
			})
			c.Abort()
			return
		}

		// Check if revoked
		if keyRecord.Revoked {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "API key has been revoked",
			})
			c.Abort()
			return
		}

		// Attach tenant and permissions to context
		c.Set("tenant_id", keyRecord.TenantID)
		c.Set("api_key_id", keyRecord.ID)

		// Parse and attach permissions
		if keyRecord.Permissions != "" {
			permissions := strings.Split(keyRecord.Permissions, ",")
			c.Set("api_key_permissions", permissions)
		}

		// Mark this request as API key authenticated
		c.Set("authenticated_via", "api_key")

		c.Next()
	}
}
