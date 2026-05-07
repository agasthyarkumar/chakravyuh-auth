package models

type UserPermission struct {
	ID uint `gorm:"primaryKey"`

	UserID uint

	PermissionID uint
}
