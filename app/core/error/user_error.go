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
		"Password is too short (minimum is "+minLenght+")", "User", "Password",
	)
}

func ErrPasswordTooLong(maxLength string) error {
	return NewBusinessRuleError(
		"Password is too long (maximum is "+maxLength+")", "User", "Password",
	)
}

func ErrPasswordInvalidFormat(message string) error {
	return NewBusinessRuleError(
		"Invalid password format: "+message, "User", "Password",
	)
}
