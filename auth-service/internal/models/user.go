package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`

	TenantID     uint      `gorm:"not null"`

	Username     string    `gorm:"unique;not null"`

	PasswordHash string    `gorm:"not null"`

	Role         string    `gorm:"default:user"`

	CreatedAt    time.Time
}