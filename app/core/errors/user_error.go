package domainerr

func ErrInvalidUserRole(value string) error {
	return NewValidationError("role", value, "Invalid user role")
}

func ErrInvalidUserStatus(status string) error {
	return NewValidationError("user_status", status, "Invalid user status")
}

func ErrInvalidUserRoleType(roleType string) error {
	return NewValidationError(
		"role_type", roleType, "Invalid user role type",
	)
}

func ErrPasswordTooShort(minLenght string) error {
	return NewBusinessRuleError(
		"password", "Password is too short (minimum is "+minLenght+")",
	)
}

func ErrPasswordTooLong(maxLength string) error {
	return NewBusinessRuleError(
		"password", "Password is too long (maximum is "+maxLength+")",
	)
}

func ErrPasswordInvalidFormat(message string) error {
	return NewBusinessRuleError(
		"password", "Invalid password format: "+message,
	)
}
