package repository

import (
	"auth-service/internal/database"
	"auth-service/internal/models"
)

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	err := database.DB.Where("username = ?", username).First(&user).Error

	return &user, err
}