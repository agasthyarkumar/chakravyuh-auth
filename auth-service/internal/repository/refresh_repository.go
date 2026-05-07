package repository

import (
	"auth-service/internal/database"
	"auth-service/internal/models"
)

func SaveRefreshToken(token *models.RefreshToken) error {
	return database.DB.Create(token).Error
}

func GetRefreshToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken

	err := database.DB.
		Where("token = ?", token).
		First(&refreshToken).Error

	return &refreshToken, err
}

func DeleteRefreshToken(token string) error {
	return database.DB.
		Where("token = ?", token).
		Delete(&models.RefreshToken{}).Error
}