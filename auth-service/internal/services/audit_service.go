package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
)

// LogAction logs an action to the audit log
func LogAction(
	tenantID uint,
	userID uint,
	action string,
	details string,
) error {
	log := models.AuditLog{
		TenantID: tenantID,
		UserID:   userID,
		Action:   action,
		Details:  details,
	}

	return repository.CreateAuditLog(&log)
}

// GetAuditLogs retrieves all audit logs
func GetAuditLogs() ([]models.AuditLog, error) {
	return repository.GetAuditLogs()
}
