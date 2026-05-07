package handlers

import (
	"auth-service/internal/dto"
	"auth-service/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

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

	var response []dto.UserResponse

	for _, user := range users {
		response = append(response, dto.UserResponse{
			ID: user.ID,
			TenantID: user.TenantID,
			Username: user.Username,
			Role: user.Role,
		})
	}

	c.JSON(http.StatusOK, response)
}

func CreateTenantUser(c *gin.Context) {
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

	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})

		return
	}

	err := services.CreateTenantUser(
		tenantID,
		input.Username,
		input.Password,
		input.Role,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func DeleteTenantUser(c *gin.Context) {
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

	userIDParam := c.Param("id")

	userIDInt, err := strconv.Atoi(
		userIDParam,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})

		return
	}

	err = services.DeleteTenantUser(
		tenantID,
		uint(userIDInt),
	)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}