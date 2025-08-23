package userDomain

import domain_error "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/domain"

func ErrInvalidUserRole(value string) error {
	return domain_error.NewValidationError("role", value, "Invalid user role")
}

func ErrInvalidUserStatus(status string) error {
	return domain_error.NewValidationError("user_status", status, "Invalid user status")
}

func ErrInvalidUserRoleType(roleType string) error {
	return domain_error.NewValidationError(
		"role_type", roleType, "Invalid user role type",
	)
}

func ErrPasswordTooShort(minLenght string) error {
	return domain_error.NewBusinessRuleError(
		"password", "Password is too short (minimum is "+minLenght+")",
	)
}

func ErrPasswordTooLong(maxLength string) error {
	return domain_error.NewBusinessRuleError(
		"password", "Password is too long (maximum is "+maxLength+")",
	)
}

func ErrPasswordInvalidFormat(message string) error {
	return domain_error.NewBusinessRuleError(
		"password", "Invalid password format: "+message,
	)
}
