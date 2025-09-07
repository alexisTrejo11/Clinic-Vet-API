package command

const (
	ErrValidationFailed   = "validation failed for unique credentials"
	ErrDataParsingFailed  = "failed to parse user data"
	ErrUserSaveFailed     = "failed to save user"
	ErrEmailAlreadyExists = "email already exists"
	ErrPhoneAlreadyExists = "phone number already exists"
	ErrInvalidCommandType = "invalid command type"
	ErrMissingCredentials = "either email or phone number is required"

	MsgUserCreatedSuccess = "user successfully created"

	MaxConcurrentValidations = 2
)
