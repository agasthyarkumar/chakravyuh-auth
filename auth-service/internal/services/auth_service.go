package services

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"auth-service/internal/utils"
	"errors"
	"strconv"
	"time"
)

func Register(
	tenantName string,
	username string,
	password string,
) error {

	hash, err := utils.HashPassword(password)

	if err != nil {
		return err
	}

	tenant := models.Tenant{
		Name: tenantName,
		Approved: false,
	}

	err = repository.CreateTenant(
		&tenant,
	)

	if err != nil {
		return err
	}

	user := models.User{
		TenantID: tenant.ID,
		Username: username,
		PasswordHash: hash,
		Role: "admin",
	}

	return repository.CreateUser(
		&user,
	)
}

func Login(
	username string,
	password string,
) (string, string, error) {

	user, err := repository.GetUserByUsername(
		username,
	)

	if err != nil {
		return "", "", errors.New(
			"invalid credentials",
		)
	}

	valid := utils.CheckPassword(
		password,
		user.PasswordHash,
	)

	if !valid {
		return "", "", errors.New(
			"invalid credentials",
		)
	}

	// =========================
	// CHECK TENANT APPROVAL
	// =========================

	if user.Role != "superadmin" {

		tenant, err := repository.GetTenantByID(
			user.TenantID,
		)

		if err != nil {
			return "", "", errors.New(
				"tenant not found",
			)
		}

		if !tenant.Approved {
			return "", "", errors.New(
				"tenant pending approval",
			)
		}
	}

	accessToken, err := utils.GenerateJWT(
		user.ID,
		user.TenantID,
		user.Username,
		user.Role,
	)

	if err != nil {
		return "", "", err
	}

	refreshTokenString, err := utils.GenerateRefreshToken()

	if err != nil {
		return "", "", err
	}

	refreshDays, err := strconv.Atoi(
		config.AppConfig.RefreshTokenExpiryDays,
	)

	if err != nil {
		refreshDays = 7
	}

	refreshToken := models.RefreshToken{
		UserID: user.ID,
		Token: refreshTokenString,
		ExpiresAt: time.Now().
			Add(time.Hour * 24 * time.Duration(refreshDays)),
	}

	err = repository.SaveRefreshToken(
		&refreshToken,
	)

	if err != nil {
		return "", "", err
	}

	// Log the login action
	_ = LogAction(
		user.TenantID,
		user.ID,
		"USER_LOGIN",
		"User "+username+" logged in",
	)

	return accessToken, refreshTokenString, nil
}