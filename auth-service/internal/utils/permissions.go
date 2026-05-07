package utils

// InitializeDefaultPermissions returns a list of default permissions
func InitializeDefaultPermissions() []string {
	return []string{
		"users.view",
		"users.create",
		"users.delete",
		"invites.create",
		"invites.view",
		"tenants.approve",
		"audit.view",
		"api_keys.create",
		"api_keys.view",
		"api_keys.revoke",
	}
}

// PermissionGroups defines which permissions belong to which roles
var PermissionGroups = map[string][]string{
	"admin": {
		"users.view",
		"users.create",
		"users.delete",
		"invites.create",
		"invites.view",
		"audit.view",
		"api_keys.create",
		"api_keys.view",
		"api_keys.revoke",
	},
	"superadmin": {
		"users.view",
		"users.create",
		"users.delete",
		"invites.create",
		"invites.view",
		"tenants.approve",
		"audit.view",
		"api_keys.create",
		"api_keys.view",
		"api_keys.revoke",
	},
	"user": {
		"users.view",
		"audit.view",
	},
}

// GetPermissionsForRole returns permissions for a specific role
func GetPermissionsForRole(role string) []string {
	if perms, exists := PermissionGroups[role]; exists {
		return perms
	}
	return []string{}
}
