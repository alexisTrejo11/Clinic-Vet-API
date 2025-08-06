package userDomain

type UserRole string
type Gender string
type UserStatus string

const (
	UserRoleAdmin        UserRole = "admin"
	UserRoleVeterinarian UserRole = "veterinarian"
	UserRoleOwner        UserRole = "owner"
	UserRoleReceptionist UserRole = "receptionist"
)

const (
	MALE         Gender = "male"
	Female       Gender = "female"
	NotSpecified Gender = "not_specified"
)

func (r UserRole) IsValid() bool {
	switch r {
	case UserRoleAdmin, UserRoleVeterinarian, UserRoleOwner, UserRoleReceptionist:
		return true
	}
	return false
}

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
	UserStatusBanned   UserStatus = "banned"
	UserStatusPending  UserStatus = "pending"
	UserStatusDeleted  UserStatus = "deleted"
)

func (s UserStatus) IsValid() bool {
	switch s {
	case UserStatusActive, UserStatusInactive, UserStatusBanned, UserStatusPending, UserStatusDeleted:
		return true
	}
	return false
}

func UserStatusFromString(status string) UserStatus {
	switch status {
	case "active":
		return UserStatusActive
	case "inactive":
		return UserStatusInactive
	case "banned":
		return UserStatusBanned
	case "pending":
		return UserStatusPending
	case "deleted":
		return UserStatusDeleted
	default:
		return UserStatusInactive
	}
}
