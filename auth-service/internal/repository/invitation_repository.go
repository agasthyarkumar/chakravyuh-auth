package repository

import (
	"auth-service/internal/database"
	"auth-service/internal/models"
)

func CreateInvitation(
	invitation *models.Invitation,
) error {

	return database.DB.Create(
		invitation,
	).Error
}

func GetInvitationByToken(
	token string,
) (*models.Invitation, error) {

	var invitation models.Invitation

	err := database.DB.
		Where("token = ?", token).
		First(&invitation).Error

	return &invitation, err
}

func UpdateInvitation(
	invitation *models.Invitation,
) error {

	return database.DB.Save(
		invitation,
	).Error
}