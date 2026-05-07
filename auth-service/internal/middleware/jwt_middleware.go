package middleware

import (
	"auth-service/internal/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing authorization header",
			})

			c.Abort()

			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")

		if len(tokenString) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token format",
			})

			c.Abort()

			return
		}

		token, err := jwt.Parse(
			tokenString[1],
			func(token *jwt.Token) (interface{}, error) {
				return []byte(config.AppConfig.JWTSecret), nil
			},
		)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})

			c.Abort()

			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid claims",
			})

			c.Abort()

			return
		}

		c.Set("user_id", claims["user_id"])
		c.Set("tenant_id", claims["tenant_id"])
		c.Set("username", claims["username"])
		c.Set("role", claims["role"])

		c.Next()
	}
}