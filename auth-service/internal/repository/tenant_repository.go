package repository

import (
	"auth-service/internal/database"
	"auth-service/internal/models"
)

func CreateTenant(tenant *models.Tenant) error {
	return database.DB.Create(tenant).Error
}

func GetTenantByID(id uint) (*models.Tenant, error) {
	var tenant models.Tenant

	err := database.DB.First(&tenant, id).Error

	return &tenant, err
}