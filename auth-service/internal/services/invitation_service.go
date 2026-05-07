package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"auth-service/internal/utils"
	"errors"
	"time"
)

func CreateInvitation(
	tenantID uint,
	email string,
	role string,
) (string, error) {

	token, err := utils.GenerateInvitationToken()

	if err != nil {
		return "", err
	}

	invitation := models.Invitation{
		TenantID: tenantID,
		Email: email,
		Role: role,
		Token: token,
		ExpiresAt: time.Now().
			Add(time.Hour * 24),
	}

	err = repository.CreateInvitation(
		&invitation,
	)

	if err != nil {
		return "", err
	}

	// Log invitation creation
	_ = LogAction(
		tenantID,
		0,
		"INVITATION_CREATED",
		"Invitation created for "+email+" with role "+role,
	)

	return token, nil
}

func AcceptInvitation(
	token string,
	username string,
	password string,
) error {

	invitation, err := repository.GetInvitationByToken(
		token,
	)

	if err != nil {
		return errors.New("invalid invitation")
	}

	if invitation.Accepted {
		return errors.New("invitation already used")
	}

	if time.Now().After(invitation.ExpiresAt) {
		return errors.New("invitation expired")
	}

	hash, err := utils.HashPassword(
		password,
	)

	if err != nil {
		return err
	}

	user := models.User{
		TenantID: invitation.TenantID,
		Username: username,
		PasswordHash: hash,
		Role: invitation.Role,
	}

	err = repository.CreateUser(
		&user,
	)

	if err != nil {
		return err
	}

	invitation.Accepted = true

	err = repository.UpdateInvitation(
		invitation,
	)

	if err != nil {
		return err
	}

	// Log invitation acceptance
	_ = LogAction(
		invitation.TenantID,
		user.ID,
		"INVITATION_ACCEPTED",
		"User "+username+" accepted invitation",
	)

	return nil
}