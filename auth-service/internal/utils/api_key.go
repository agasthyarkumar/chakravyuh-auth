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

// GenerateAPIKey generates a random API key
func GenerateAPIKey() (string, error) {
	randomBytes := make([]byte, apiKeyLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Combine prefix with random hex
	apiKey := apiKeyPrefix + hex.EncodeToString(randomBytes)
	return apiKey, nil
}

// HashAPIKey hashes an API key for storage
func HashAPIKey(apiKey string) string {
	hash := sha256.Sum256([]byte(apiKey))
	return hex.EncodeToString(hash[:])
}

// VerifyAPIKey verifies if a raw key matches a hash
func VerifyAPIKey(rawKey string, hash string) bool {
	hashedKey := HashAPIKey(rawKey)
	return hashedKey == hash
}

// ValidateAPIKeyFormat validates the format of an API key
func ValidateAPIKeyFormat(apiKey string) error {
	if len(apiKey) < len(apiKeyPrefix)+64 {
		return fmt.Errorf("invalid API key format")
	}

	if apiKey[:len(apiKeyPrefix)] != apiKeyPrefix {
		return fmt.Errorf("invalid API key prefix")
	}

	return nil
}
