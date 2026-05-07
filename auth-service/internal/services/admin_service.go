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

	err = repository.CreateUser(&user)
	
	if err != nil {
		return err
	}

	// Log the action (use 0 for admin user ID since we don't have it in context)
	// This can be enhanced later with proper admin user tracking
	_ = LogAction(
		tenantID,
		0,
		"ADMIN_CREATED_USER",
		"User "+username+" was created with role "+role,
	)

	return nil
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

	err = repository.DeleteUser(userID)
	
	if err != nil {
		return err
	}

	// Log the deletion
	_ = LogAction(
		requestingTenantID,
		0,
		"ADMIN_DELETED_USER",
		"User "+user.Username+" (ID: "+string(rune(user.ID))+") was deleted",
	)

	return nil
}