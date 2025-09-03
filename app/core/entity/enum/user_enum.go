package enum

import "strings"

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

type PersonGender string

const (
	Male         PersonGender = "male"
	Female       PersonGender = "female"
	NotSpecified PersonGender = "not_specified"
)

var ValidPersonGenders = []PersonGender{Male, Female, NotSpecified}

func (g PersonGender) String() string {
	return string(g)
}

func NewGender(value string) PersonGender {
	normalized := normalizeGenderInput(value)

	switch normalized {
	case "male":
		return Male
	case "female":
		return Female
	case "not_specified", "not specified", "":
		return NotSpecified
	default:
		return NotSpecified
	}
}

func normalizeGenderInput(input string) string {
	input = strings.TrimSpace(strings.ToLower(input))
	input = strings.ReplaceAll(input, " ", "_")
	return input
}

func (g PersonGender) Values() []PersonGender {
	return ValidPersonGenders
}

func (g PersonGender) DisplayName() string {
	switch g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	case NotSpecified:
		return "Not Specified"
	default:
		return "Unknown"
	}
}
