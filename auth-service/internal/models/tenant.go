package models

import "time"

type Tenant struct {
	ID        uint      `gorm:"primaryKey"`

	Name      string    `gorm:"unique;not null"`

	Approved  bool      `gorm:"default:false"`

	CreatedAt time.Time
}