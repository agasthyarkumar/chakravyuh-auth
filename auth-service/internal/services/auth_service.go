package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"auth-service/internal/utils"
	"errors"
)

func Register(username string, password string) error {
	hash, err := utils.HashPassword(password)

	if err != nil {
		return err
	}

	user := models.User{
		Username:     username,
		PasswordHash: hash,
		Role:         "user",
	}

	return repository.CreateUser(&user)
}

func Login(username string, password string) (string, error) {
	user, err := repository.GetUserByUsername(username)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	valid := utils.CheckPassword(password, user.PasswordHash)

	if !valid {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID, user.Username, user.Role)

	if err != nil {
		return "", err
	}

	return token, nil
}