package repository

import (
	"auth-service/internal/database"
	"auth-service/internal/models"
)

func CreateTenant(
	tenant *models.Tenant,
) error {

	return database.DB.Create(
		tenant,
	).Error
}

func GetTenantByID(
	id uint,
) (*models.Tenant, error) {

	var tenant models.Tenant

	err := database.DB.
		First(&tenant, id).Error

	return &tenant, err
}

func GetPendingTenants() (
	[]models.Tenant,
	error,
) {

	var tenants []models.Tenant

	err := database.DB.
		Where("approved = ?", false).
		Find(&tenants).Error

	return tenants, err
}

func ApproveTenant(
	id uint,
) error {

	return database.DB.
		Model(&models.Tenant{}).
		Where("id = ?", id).
		Update("approved", true).Error
}