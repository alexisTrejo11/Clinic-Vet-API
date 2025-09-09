package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/auth"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

// UserOption defines the functional option type
type UserOption func(*User) error

func WithEmail(email valueobject.Email) UserOption {
	return func(u *User) error {
		u.email = email
		return nil
	}
}

func WithPhoneNumber(phone valueobject.PhoneNumber) UserOption {
	return func(u *User) error {
		u.phoneNumber = phone
		return nil
	}
}

func WithPassword(password string) UserOption {
	return func(u *User) error {
		u.password = password
		return nil
	}
}

func WithLastLoginAt(lastLogin time.Time) UserOption {
	return func(u *User) error {
		if lastLogin.After(time.Now()) {
			return errors.New("last login cannot be in the future")
		}
		u.lastLoginAt = &lastLogin
		return nil
	}
}

func WithTwoFactorAuth(twoFA auth.TwoFactorAuth) UserOption {
	return func(u *User) error {
		u.twoFactorAuth = twoFA
		return nil
	}
}

func WithJoinedAt(joinedAt time.Time) UserOption {
	return func(u *User) error {
		if joinedAt.After(time.Now()) {
			return errors.New("joined date cannot be in the future")
		}
		u.SetTimeStamps(joinedAt, time.Time{})
		return nil
	}
}

// NewUser creates a new User with functional options
func NewUser(
	id valueobject.UserID,
	role enum.UserRole,
	status enum.UserStatus,
	opts ...UserOption,
) (*User, error) {
	if err := validateRole(role); err != nil {
		return nil, err
	}

	now := time.Now()
	user := &User{
		Entity:        base.NewEntity(id, now, now, 1),
		role:          role,
		status:        status,
		twoFactorAuth: auth.NewDisabledTwoFactorAuth(), // Default disabled 2FA
	}

	for _, opt := range opts {
		if err := opt(user); err != nil {
			return nil, fmt.Errorf("invalid user option: %w", err)
		}
	}

	if err := user.validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func validateStatus(status enum.UserStatus) error {
	if !status.IsValid() {
		return errors.New("invalid user status")
	}
	return nil
}

func validateRole(role enum.UserRole) error {
	if !role.IsValid() {
		return errors.New("invalid user role")
	}
	return nil
}

func (u *User) validate() error {
	if u.email.String() == "" {
		return errors.New("email is required")
	}
	if u.password == "" {
		return errors.New("password is required")
	}
	return nil
}
