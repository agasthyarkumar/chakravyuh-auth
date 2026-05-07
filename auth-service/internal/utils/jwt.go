package utils

import (
	"auth-service/internal/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint, username string, role string) (string, error) {
	expiryMinutes, err := strconv.Atoi(
		config.AppConfig.AccessTokenExpiryMinutes,
	)

	if err != nil {
		expiryMinutes = 15
	}

	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role":     role,
		"exp": time.Now().
			Add(time.Minute * time.Duration(expiryMinutes)).
			Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(config.AppConfig.JWTSecret),
	)
}