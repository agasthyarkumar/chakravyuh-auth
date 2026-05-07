package repository

import (
	"auth-service/internal/database"
	"auth-service/internal/models"
)

// AssignPermission assigns a permission to a user
func AssignPermission(userID uint, permissionID uint) error {
	userPermission := models.UserPermission{
		UserID:       userID,
		PermissionID: permissionID,
	}
	return database.DB.Create(&userPermission).Error
}

// RevokePermission removes a permission from a user
func RevokePermission(userID uint, permissionID uint) error {
	return database.DB.
		Where("user_id = ? AND permission_id = ?", userID, permissionID).
		Delete(&models.UserPermission{}).Error
}

// GetUserPermissions fetches all permissions for a user
func GetUserPermissions(userID uint) ([]models.Permission, error) {
	var permissions []models.Permission

	err := database.DB.
		Joins("INNER JOIN user_permissions ON permissions.id = user_permissions.permission_id").
		Where("user_permissions.user_id = ?", userID).
		Find(&permissions).Error

	return permissions, err
}

// HasPermission checks if a user has a specific permission
func HasPermission(userID uint, permissionName string) (bool, error) {
	var count int64

	err := database.DB.
		Table("permissions").
		Joins("INNER JOIN user_permissions ON permissions.id = user_permissions.permission_id").
		Where("user_permissions.user_id = ? AND permissions.name = ?", userID, permissionName).
		Count(&count).Error

	return count > 0, err
}

// GetPermissionByName fetches a permission by name
func GetPermissionByName(name string) (*models.Permission, error) {
	var permission models.Permission

	err := database.DB.
		Where("name = ?", name).
		First(&permission).Error

	if err != nil {
		return nil, err
	}

	return &permission, nil
}

// CreatePermission creates a new permission
func CreatePermission(permission *models.Permission) error {
	return database.DB.Create(permission).Error
}
