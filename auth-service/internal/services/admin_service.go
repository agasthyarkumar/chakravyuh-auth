package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"auth-service/internal/utils"
	"errors"
)

func ListTenantUsers(
	tenantID uint,
) ([]models.User, error) {

	return repository.GetUsersByTenantID(
		tenantID,
	)
}

func CreateTenantUser(
	tenantID uint,
	username string,
	password string,
	role string,
) error {

	hash, err := utils.HashPassword(
		password,
	)

	if err != nil {
		return err
	}

	user := models.User{
		TenantID: tenantID,
		Username: username,
		PasswordHash: hash,
		Role: role,
	}

	return repository.CreateUser(&user)
}

func DeleteTenantUser(
	requestingTenantID uint,
	userID uint,
) error {

	user, err := repository.GetUserByID(
		userID,
	)

	if err != nil {
		return errors.New("user not found")
	}

	if user.TenantID != requestingTenantID {
		return errors.New("access denied")
	}

	return repository.DeleteUser(userID)
}