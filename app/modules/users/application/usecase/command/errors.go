package command

const (
	ErrFailedFindUser     = "failed to find user"
	ErrFailedChangePhone  = "failed to change phone"
	ErrFailedChangeEmail  = "failed to change email"
	ErrFailedUpdateUser   = "failed to update user"
	ErrEmailAlreadyExists = "email already exists"
	ErrPhoneAlreadyTaken  = "phone already taken"
	ErrInvalidUserID      = "invalid user ID"
	ErrInvalidEmail       = "invalid email"
	ErrInvalidPhone       = "invalid phone number"
	ErrPhoneUnchanged     = "phone number unchanged"
	ErrEmailUnchanged     = "email unchanged"
)

const (
	ErrFailedMappingUser   = "failed to map user from command"
	ErrFailedValidation    = "failed to validate user creation"
	ErrFailedHashPassword  = "failed to hash password"
	ErrFailedSaveUser      = "failed to save user"
	ErrInvalidRole         = "invalid user role"
	ErrInvalidStatus       = "invalid user status"
	ErrInvalidGender       = "invalid gender"
	ErrInvalidDateOfBirth  = "invalid date of birth"
	ErrUserCreationSuccess = "user created successfully"
)
