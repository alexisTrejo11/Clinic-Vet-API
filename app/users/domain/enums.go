package user

type UserRole string

type UserStatus string

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

func UserRoleFromString(role string) UserRole {
	switch role {
	case "admin":
		return UserRoleAdmin
	case "veterinarian":
		return UserRoleVeterinarian
	case "owner":
		return UserRoleOwner
	case "receptionist":
		return UserRoleReceptionist
	default:
		return UserRoleAdmin
	}
}

func (r UserRole) String() string {
	return string(r)
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
