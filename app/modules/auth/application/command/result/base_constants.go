package result

import (
	"errors"
	"time"
)

var (
	ErrValidationFailed   = "validation failed for unique credentials"
	ErrDataParsingFailed  = "failed to parse user data"
	ErrUserSaveFailed     = "failed to save user"
	ErrEmailAlreadyExists = "email already exists"
	ErrPhoneAlreadyExists = "phone number already exists"
	ErrInvalidCommandType = "invalid command type"
	ErrMissingCredentials = "either email or phone number is required"
	ErrUserCreationFailed = "user creation failed"
	MsgUserCreatedSuccess = "user successfully created"
	ErrUserNotFound       = "error finding user"

	MaxConcurrentValidations = 2

	ErrAuthenticationFailed  = "authentication failed"
	ErrTwoFactorRequired     = "2FA is enabled for this user, please complete the 2FA process"
	ErrInvalidTwoFactorToken = "invalid two-factor authentication token"
	ErrSessionCreationFailed = "failed to create session"
	ErrAccessTokenGenFailed  = "failed to generate access token"
	ErrInvalidCredentials    = "user not found with provided credentials, please check your email/phone-number and password"
	ErrTwoFactorAuthConflict = "user has TwoFactorAuth auth login method"

	MsgLoginSuccess = "login successfully processed"

	DefaultSessionDuration = 7 * 24 * time.Hour
)

var ErrInvalidCommand = errors.New("invalid command")
