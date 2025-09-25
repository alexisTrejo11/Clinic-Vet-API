package user

import (
	"errors"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/auth"
	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

// UserOption defines the functional option type
type UserOption func(*User)

func WithEmail(email valueobject.Email) UserOption {
	return func(u *User) {
		u.email = email
	}
}

func WithPhoneNumber(phone valueobject.PhoneNumber) UserOption {
	return func(u *User) {
		u.phoneNumber = phone
	}
}

func WithPassword(password string) UserOption {
	return func(u *User) {
		u.password = password
	}
}

func WithLastLoginAt(lastLogin time.Time) UserOption {
	return func(u *User) {
		u.lastLoginAt = &lastLogin
	}
}

func WithTwoFactorAuth(twoFA auth.TwoFactorAuth) UserOption {
	return func(u *User) {
		u.twoFactorAuth = twoFA
	}
}

func WithCustomerID(customerID valueobject.CustomerID) UserOption {
	return func(u *User) {
		u.customerID = &customerID
	}

}

func WithEmployeeID(employeeID valueobject.EmployeeID) UserOption {
	return func(u *User) {
		u.employeeID = &employeeID
	}
}

func WithJoinedAt(joinedAt time.Time) UserOption {
	return func(u *User) {
		u.SetTimeStamps(joinedAt, time.Time{})
	}
}

// NewUser creates a new User with functional options
func NewUser(id valueobject.UserID, role enum.UserRole, status enum.UserStatus, opts ...UserOption) (User, error) {
	user := &User{
		Entity: base.NewEntity(id, nil, nil, 0),
		role:   role,
		status: status,
	}

	for _, opt := range opts {
		opt(user)
	}

	return *user, nil
}

func CreateUser(role enum.UserRole, status enum.UserStatus, opts ...UserOption) (*User, error) {
	user := &User{
		Entity:        base.CreateEntity(valueobject.UserID{}),
		role:          role,
		status:        status,
		twoFactorAuth: auth.NewDisabledTwoFactorAuth(), // Default disabled 2FA
	}

	for _, opt := range opts {
		opt(user)
	}

	if err := user.validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) validate() error {
	if u.email.String() == "" {
		return errors.New("email is required")
	}
	if u.password == "" {
		return errors.New("password is required")
	}
	if !u.role.IsValid() {
		return errors.New("invalid user role")
	}
	if !u.status.IsValid() {
		return errors.New("invalid user status")
	}

	return nil
}
