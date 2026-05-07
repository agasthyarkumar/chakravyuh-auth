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

	// Log user creation
	_ = LogAction(
		tenantID,
		0, // admin context, not linked to a specific user
		"ADMIN_CREATED_USER",
		"Admin created user '"+username+"' with role '"+role+"'",
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

	// Log user deletion
	_ = LogAction(
		requestingTenantID,
		0,
		"ADMIN_DELETED_USER",
		"Admin deleted user '"+user.Username+"'",
	)

	return nil
}