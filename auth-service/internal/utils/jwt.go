package utils

import (
	"auth-service/internal/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(
	userID uint,
	tenantID uint,
	username string,
	role string,
) (string, error) {

	expiryMinutes, err := strconv.Atoi(
		config.AppConfig.AccessTokenExpiryMinutes,
	)

	if err != nil {
		expiryMinutes = 15
	}

	claims := jwt.MapClaims{
		"user_id":   userID,
		"tenant_id": tenantID,
		"username":  username,
		"role":      role,
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

// ParseJWT parses a JWT token and returns its claims
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.JWTSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}