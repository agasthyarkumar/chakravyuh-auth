package middleware

import (
	"net/http"
	"strconv"

	"auth-service/internal/repository"
	"github.com/gin-gonic/gin"
)

// RequirePermission creates a middleware that checks if user has a specific permission
func RequirePermission(permissionName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user_id from context (set by JWT middleware)
		userIDInterface, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "user not authenticated",
			})
			c.Abort()
			return
		}

		// Convert user_id to uint
		userID, ok := userIDInterface.(uint)
		if !ok {
			// Try converting from float64 (from JWT claims)
			if floatID, isFloat := userIDInterface.(float64); isFloat {
				userID = uint(floatID)
			} else if strID, isStr := userIDInterface.(string); isStr {
				id, err := strconv.ParseUint(strID, 10, 32)
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{
						"error": "invalid user id",
					})
					c.Abort()
					return
				}
				userID = uint(id)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "invalid user id type",
				})
				c.Abort()
				return
			}
		}

		// Check if user has the required permission
		hasPermission, err := repository.HasPermission(userID, permissionName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to check permissions",
			})
			c.Abort()
			return
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "insufficient permissions",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
