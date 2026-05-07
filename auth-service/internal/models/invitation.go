package models

import "time"

type Invitation struct {
	ID uint `gorm:"primaryKey"`

	TenantID uint `gorm:"not null"`

	Email string `gorm:"not null"`

	Role string `gorm:"default:user"`

	Token string `gorm:"unique;not null"`

	Accepted bool `gorm:"default:false"`

	ExpiresAt time.Time `gorm:"not null"`

	CreatedAt time.Time
}