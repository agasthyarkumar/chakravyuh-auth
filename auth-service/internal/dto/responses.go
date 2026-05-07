package dto

import "time"

// AuditLogResponse represents a sanitized audit log response
type AuditLogResponse struct {
	ID        uint      `json:"id"`
	TenantID  uint      `json:"tenant_id"`
	UserID    uint      `json:"user_id"`
	Action    string    `json:"action"`
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"created_at"`
}

// PermissionResponse represents a permission
type PermissionResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// UserPermissionsResponse represents user permissions
type UserPermissionsResponse struct {
	UserID      uint                   `json:"user_id"`
	Permissions []PermissionResponse   `json:"permissions"`
}

// APIKeyResponse represents a sanitized API key response (no raw key shown)
type APIKeyResponse struct {
	ID          uint      `json:"id"`
	TenantID    uint      `json:"tenant_id"`
	Name        string    `json:"name"`
	Permissions string    `json:"permissions"`
	Revoked     bool      `json:"revoked"`
	CreatedAt   time.Time `json:"created_at"`
}

// APIKeyCreateResponse represents response when creating an API key (raw key shown once)
type APIKeyCreateResponse struct {
	ID          uint      `json:"id"`
	TenantID    uint      `json:"tenant_id"`
	Name        string    `json:"name"`
	APIKey      string    `json:"api_key"` // Only shown once during creation
	Permissions string    `json:"permissions"`
	CreatedAt   time.Time `json:"created_at"`
}

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// SuccessResponse represents a standard success response
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
