package domainerr

import "context"

func ErrInvalidUserRole(value string) error {
	return ValidationError(context.Background(), "INVALID_USER_ROLE", "role", value, "Invalid user role", "user role validation")
}

func ErrInvalidUserStatus(status string) error {
	return ValidationError(context.Background(), "INVALID_USER_STATUS", "user_status", status, "Invalid user status", "user status validation")
}

func ErrInvalidUserRoleType(roleType string) error {
	return ValidationError(
		context.Background(), "INVALID_USER_ROLE", "role_type", roleType, "Invalid user role type", "user role validation",
	)
}

func ErrPasswordTooShort(minLenght string) error {
	return BusinessRuleError(
		context.Background(),
		"Password is too short (minimum is "+minLenght+")", "User", "Password", minLenght,
	)
}

func ErrPasswordTooLong(maxLength string) error {
	return BusinessRuleError(
		context.Background(),
		"Password is too long (maximum is "+maxLength+")", "User", "Password", maxLength,
	)
}

func ErrPasswordInvalidFormat(message string) error {
	return BusinessRuleError(
		context.Background(),
		"Invalid password format: "+message, "User", "Password", message,
	)
}
