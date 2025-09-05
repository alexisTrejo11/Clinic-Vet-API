package enum

import (
	"fmt"
	"strings"
)

// UserRole represents the role of a user in the system
type UserRole string

const (
	UserRoleAdmin        UserRole = "admin"
	UserRoleVeterinarian UserRole = "veterinarian"
	UserRoleOwner        UserRole = "owner"
	UserRoleReceptionist UserRole = "receptionist"
)

// UserRole constants and methods
var (
	ValidUserRoles = []UserRole{
		UserRoleAdmin,
		UserRoleVeterinarian,
		UserRoleOwner,
		UserRoleReceptionist,
	}

	userRoleMap = map[string]UserRole{
		"admin":        UserRoleAdmin,
		"veterinarian": UserRoleVeterinarian,
		"owner":        UserRoleOwner,
		"receptionist": UserRoleReceptionist,
	}

	userRoleDisplayNames = map[UserRole]string{
		UserRoleAdmin:        "Administrator",
		UserRoleVeterinarian: "Veterinarian",
		UserRoleOwner:        "Pet Owner",
		UserRoleReceptionist: "Receptionist",
	}
)

func (r UserRole) IsValid() bool {
	_, exists := userRoleMap[string(r)]
	return exists
}

func ParseUserRole(role string) (UserRole, error) {
	normalized := strings.TrimSpace(strings.ToLower(role))
	if val, exists := userRoleMap[normalized]; exists {
		return val, nil
	}
	return "", fmt.Errorf("invalid user role: %s", role)
}

func MustParseUserRole(role string) UserRole {
	parsed, err := ParseUserRole(role)
	if err != nil {
		panic(err)
	}
	return parsed
}

func (r UserRole) String() string {
	return string(r)
}

func (r UserRole) DisplayName() string {
	if displayName, exists := userRoleDisplayNames[r]; exists {
		return displayName
	}
	return "Unknown Role"
}

func (r UserRole) Values() []UserRole {
	return ValidUserRoles
}

// UserStatus represents the status of a user account
type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
	UserStatusBanned   UserStatus = "banned"
	UserStatusPending  UserStatus = "pending"
	UserStatusDeleted  UserStatus = "deleted"
)

// UserStatus constants and methods
var (
	ValidUserStatuses = []UserStatus{
		UserStatusActive,
		UserStatusInactive,
		UserStatusBanned,
		UserStatusPending,
		UserStatusDeleted,
	}

	userStatusMap = map[string]UserStatus{
		"active":   UserStatusActive,
		"inactive": UserStatusInactive,
		"banned":   UserStatusBanned,
		"pending":  UserStatusPending,
		"deleted":  UserStatusDeleted,
	}

	userStatusDisplayNames = map[UserStatus]string{
		UserStatusActive:   "Active",
		UserStatusInactive: "Inactive",
		UserStatusBanned:   "Banned",
		UserStatusPending:  "Pending",
		UserStatusDeleted:  "Deleted",
	}
)

func (s UserStatus) IsValid() bool {
	_, exists := userStatusMap[string(s)]
	return exists
}

func ParseUserStatus(status string) (UserStatus, error) {
	normalized := strings.TrimSpace(strings.ToLower(status))
	if val, exists := userStatusMap[normalized]; exists {
		return val, nil
	}
	return "", fmt.Errorf("invalid user status: %s", status)
}

func MustParseUserStatus(status string) UserStatus {
	parsed, err := ParseUserStatus(status)
	if err != nil {
		panic(err)
	}
	return parsed
}

func (s UserStatus) String() string {
	return string(s)
}

func (s UserStatus) DisplayName() string {
	if displayName, exists := userStatusDisplayNames[s]; exists {
		return displayName
	}
	return "Unknown Status"
}

func (s UserStatus) Values() []UserStatus {
	return ValidUserStatuses
}

// PersonGender represents a person's gender
type PersonGender string

const (
	GenderMale         PersonGender = "male"
	GenderFemale       PersonGender = "female"
	GenderNotSpecified PersonGender = "not_specified"
	GenderOther        PersonGender = "other"
)

// PersonGender constants and methods
var (
	ValidPersonGenders = []PersonGender{
		GenderMale,
		GenderFemale,
		GenderNotSpecified,
		GenderOther,
	}

	personGenderMap = map[string]PersonGender{
		"male":          GenderMale,
		"female":        GenderFemale,
		"not_specified": GenderNotSpecified,
		"not specified": GenderNotSpecified,
		"other":         GenderOther,
		"":              GenderNotSpecified,
	}

	personGenderDisplayNames = map[PersonGender]string{
		GenderMale:         "Male",
		GenderFemale:       "Female",
		GenderNotSpecified: "Not Specified",
		GenderOther:        "Other",
	}
)

func (g PersonGender) IsValid() bool {
	_, exists := personGenderMap[string(g)]
	return exists
}

func ParseGender(gender string) (PersonGender, error) {
	normalized := normalizeGenderInput(gender)
	if val, exists := personGenderMap[normalized]; exists {
		return val, nil
	}
	return GenderNotSpecified, fmt.Errorf("invalid gender: %s", gender)
}

func MustParseGender(gender string) PersonGender {
	parsed, err := ParseGender(gender)
	if err != nil {
		panic(err)
	}
	return parsed
}

func (g PersonGender) String() string {
	return string(g)
}

func (g PersonGender) DisplayName() string {
	if displayName, exists := personGenderDisplayNames[g]; exists {
		return displayName
	}
	return "Unknown Gender"
}

func (g PersonGender) Values() []PersonGender {
	return ValidPersonGenders
}

func normalizeGenderInput(input string) string {
	input = strings.TrimSpace(strings.ToLower(input))
	input = strings.ReplaceAll(input, " ", "_")
	return input
}

// Utility functions for all enums
func GetAllUserRoles() []UserRole {
	return ValidUserRoles
}

func GetAllUserStatuses() []UserStatus {
	return ValidUserStatuses
}

func GetAllGenders() []PersonGender {
	return ValidPersonGenders
}

// IsActive checks if a user status is considered active
func (s UserStatus) IsActive() bool {
	return s == UserStatusActive
}

// CanLogin checks if a user status allows login
func (s UserStatus) CanLogin() bool {
	return s == UserStatusActive
}

// IsAdministrative checks if a user role has administrative privileges
func (r UserRole) IsAdministrative() bool {
	return r == UserRoleAdmin
}
