package models

import "time"

type APIKey struct {
	ID uint `gorm:"primaryKey"`

	TenantID uint

	Name string

	KeyHash string

	Permissions string

	Revoked bool

	CreatedAt time.Time
}
