package services

import (
	"auth-service/internal/repository"
	"auth-service/internal/redis"
	"auth-service/internal/utils"
	"context"
	"errors"
	"time"
)

func Refresh(refreshToken string) (string, error) {
	tokenRecord, err := repository.GetRefreshToken(
		refreshToken,
	)

	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	if time.Now().After(tokenRecord.ExpiresAt) {
		return "", errors.New("refresh token expired")
	}

	user, err := repository.GetUserByID(
		tokenRecord.UserID,
	)

	if err != nil {
		return "", err
	}

	accessToken, err := utils.GenerateJWT(
		user.ID,
		user.TenantID,
		user.Username,
		user.Role,
	)

	if err != nil {
		return "", err
	}

	// Log token refresh
	_ = LogAction(
		user.TenantID,
		user.ID,
		"TOKEN_REFRESH",
		"User refreshed access token",
	)

	return accessToken, nil
}

func Logout(refreshToken string) error {
	// Revoke refresh token via Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = redis.Delete(ctx, "refresh_token:"+refreshToken)

	// Delete from database
	return repository.DeleteRefreshToken(
		refreshToken,
	)
}