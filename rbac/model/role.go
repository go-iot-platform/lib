package model

// AccessRole represents access role type
type AccessRole int8

const (
	// SuperAdminRole has all permissions
	SuperAdminRole AccessRole = iota + 1

	// AdminRole has admin specific permissions
	AdminRole

	// CompanyAdminRole, only for company
	CustomerAdminRole

	// CustomerUserRole, only for company
	CustomerUserRole
	// ClientAdminRole, only for client's company
	ClientAdminRole
	// CustomerUserRole, only for client's company
	ClientUserRole
)

func (s AccessRole) String() string {
	switch s {
	case SuperAdminRole:
		return "SuperAdminRole"
	case AdminRole:
		return "AdminRole"
	case CustomerAdminRole:
		return "CustomerAdminRole"
	case CustomerUserRole:
		return "CustomerUserRole"
	case ClientAdminRole:
		return "ClientAdminRole"
	case ClientUserRole:
		return "ClientUserRole"
	default:
		return "Unknown"
	}
}
// Role model
type Role struct {
	ID          int        `json:"id"`
	AccessLevel AccessRole `json:"access_level"`
	Name        string     `json:"name"`
}
