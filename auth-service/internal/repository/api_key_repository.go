package repository

import (
	"auth-service/internal/database"
	"auth-service/internal/models"
)

func CreateAPIKey(apiKey *models.APIKey) error {
	return database.DB.Create(apiKey).Error
}

func GetAPIKey(keyHash string) (*models.APIKey, error) {
	var apiKey models.APIKey

	err := database.DB.
		Where("key_hash = ?", keyHash).
		First(&apiKey).Error

	if err != nil {
		return nil, err
	}

	return &apiKey, nil
}

func GetAPIKeysByTenant(tenantID uint) ([]models.APIKey, error) {
	var apiKeys []models.APIKey

	err := database.DB.
		Where("tenant_id = ? AND revoked = ?", tenantID, false).
		Order("created_at desc").
		Find(&apiKeys).Error

	return apiKeys, err
}

func RevokeAPIKey(id uint) error {
	return database.DB.
		Model(&models.APIKey{}).
		Where("id = ?", id).
		Update("revoked", true).Error
}

func GetAPIKeyByID(id uint) (*models.APIKey, error) {
	var apiKey models.APIKey

	err := database.DB.
		Where("id = ?", id).
		First(&apiKey).Error

	if err != nil {
		return nil, err
	}

	return &apiKey, nil
}
