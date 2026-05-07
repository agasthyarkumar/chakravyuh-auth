package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
)

func ListTenantUsers(
	tenantID uint,
) ([]models.User, error) {

	return repository.GetUsersByTenantID(
		tenantID,
	)
}