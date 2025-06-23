package userEnums

type UserRole string

const (
	UserRoleAdmin        UserRole = "admin"
	UserRoleVeterinarian UserRole = "veterinarian"
	UserRoleOwner        UserRole = "owner"
	UserRoleReceptionist UserRole = "receptionist"
)

func (r UserRole) IsValid() bool {
	switch r {
	case UserRoleAdmin, UserRoleVeterinarian, UserRoleOwner, UserRoleReceptionist:
		return true
	}
	return false
}
