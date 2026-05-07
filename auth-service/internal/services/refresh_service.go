package services

import (
	"auth-service/internal/repository"
	"auth-service/internal/utils"
	"errors"
	"time"
)

func Refresh(refreshToken string) (string, error) {
	tokenRecord, err := repository.GetRefreshToken(refreshToken)

	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	if time.Now().After(tokenRecord.ExpiresAt) {
		return "", errors.New("refresh token expired")
	}

	user, err := repository.GetUserByID(tokenRecord.UserID)

	if err != nil {
		return "", err
	}

	accessToken, err := utils.GenerateJWT(
		user.ID,
		user.Username,
		user.Role,
	)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func Logout(refreshToken string) error {
	return repository.DeleteRefreshToken(refreshToken)
}