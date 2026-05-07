package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const (
	apiKeyLength = 32
	apiKeyPrefix = "chk_"
)

func GenerateAPIKey() (string, error) {
	randomBytes := make([]byte, apiKeyLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	apiKey := apiKeyPrefix + hex.EncodeToString(randomBytes)
	return apiKey, nil
}

func HashAPIKey(apiKey string) string {
	hash := sha256.Sum256([]byte(apiKey))
	return hex.EncodeToString(hash[:])
}

func VerifyAPIKey(rawKey string, hash string) bool {
	hashedKey := HashAPIKey(rawKey)
	return hashedKey == hash
}

func ValidateAPIKeyFormat(apiKey string) error {
	if len(apiKey) < len(apiKeyPrefix)+64 {
		return fmt.Errorf("invalid API key format")
	}

	if apiKey[:len(apiKeyPrefix)] != apiKeyPrefix {
		return fmt.Errorf("invalid API key prefix")
	}

	return nil
}
