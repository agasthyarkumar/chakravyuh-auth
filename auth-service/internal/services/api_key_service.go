package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"auth-service/internal/utils"
)

// CreateAPIKeyRequest represents a request to create an API key
type CreateAPIKeyRequest struct {
	Name        string `json:"name" binding:"required"`
	Permissions string `json:"permissions"`
}

// CreateAPIKeyResponse represents the response when creating an API key
type CreateAPIKeyResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	APIKey      string `json:"api_key"` // Only shown once
	Permissions string `json:"permissions"`
	CreatedAt   string `json:"created_at"`
}

// APIKeyResponse represents an API key in responses (without the raw key)
type APIKeyResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Permissions string `json:"permissions"`
	Revoked     bool   `json:"revoked"`
	CreatedAt   string `json:"created_at"`
}

// CreateAPIKey creates a new API key for a tenant
func CreateAPIKey(tenantID uint, req CreateAPIKeyRequest) (*CreateAPIKeyResponse, error) {
	// Generate a new API key
	rawKey, err := utils.GenerateAPIKey()
	if err != nil {
		return nil, err
	}

	// Hash the key
	keyHash := utils.HashAPIKey(rawKey)

	// Create API key record
	apiKeyModel := models.APIKey{
		TenantID:    tenantID,
		Name:        req.Name,
		KeyHash:     keyHash,
		Permissions: req.Permissions,
		Revoked:     false,
	}

	if err := repository.CreateAPIKey(&apiKeyModel); err != nil {
		return nil, err
	}

	// Return response with raw key (only shown once)
	return &CreateAPIKeyResponse{
		ID:          apiKeyModel.ID,
		Name:        apiKeyModel.Name,
		APIKey:      rawKey,
		Permissions: apiKeyModel.Permissions,
		CreatedAt:   apiKeyModel.CreatedAt.String(),
	}, nil
}

// GetTenantAPIKeys retrieves all API keys for a tenant
func GetTenantAPIKeys(tenantID uint) ([]APIKeyResponse, error) {
	apiKeys, err := repository.GetAPIKeysByTenant(tenantID)
	if err != nil {
		return nil, err
	}

	var responses []APIKeyResponse
	for _, key := range apiKeys {
		responses = append(responses, APIKeyResponse{
			ID:          key.ID,
			Name:        key.Name,
			Permissions: key.Permissions,
			Revoked:     key.Revoked,
			CreatedAt:   key.CreatedAt.String(),
		})
	}

	return responses, nil
}

// RevokeAPIKey revokes an API key
func RevokeAPIKey(keyID uint) error {
	return repository.RevokeAPIKey(keyID)
}
