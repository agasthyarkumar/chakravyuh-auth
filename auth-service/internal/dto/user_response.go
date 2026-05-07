package dto

type UserResponse struct {
	ID       uint   `json:"id"`
	TenantID uint   `json:"tenant_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}