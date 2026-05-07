package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
)

func ListPendingTenants() (
	[]models.Tenant,
	error,
) {

	return repository.GetPendingTenants()
}

func ApproveTenant(
	id uint,
) error {

	err := repository.ApproveTenant(
		id,
	)

	if err != nil {
		return err
	}

	// Get tenant info for audit log
	tenant, err := repository.GetTenantByID(id)
	if err == nil {
		// Log tenant approval
		_ = LogAction(
			id,
			0,
			"TENANT_APPROVED",
			"Tenant "+tenant.Name+" was approved by superadmin",
		)
	}

	return nil
}