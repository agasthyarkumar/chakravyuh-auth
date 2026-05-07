package models

import "time"

type AuditLog struct {
	ID uint `gorm:"primaryKey"`

	TenantID uint

	UserID uint

	Action string `gorm:"not null"`

	Details string

	CreatedAt time.Time
}
