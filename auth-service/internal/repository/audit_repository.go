package repository

import (
	"auth-service/internal/database"
	"auth-service/internal/models"
)

func CreateAuditLog(log *models.AuditLog) error {
	return database.DB.Create(log).Error
}

func GetAuditLogs() ([]models.AuditLog, error) {
	var logs []models.AuditLog

	err := database.DB.
		Order("created_at desc").
		Find(&logs).Error

	return logs, err
}
