package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/errors"
)

type User struct {
	id            valueobject.UserID
	email         valueobject.Email
	phoneNumber   valueobject.PhoneNumber
	password      string
	role          enum.UserRole
	status        enum.UserStatus
	lastLoginAt   time.Time
	joinedAt      time.Time
	twoFactorAuth TwoFactorAuth
}

type Profile struct {
	userID         valueobject.UserID
	OwnerId        *int
	VeterinarianId *int
	PhotoURL       string
	Bio            string
	Address        *Address
	DateOfBirth    *time.Time
	JoinedAt       time.Time
}

func NewUser(
	id valueobject.UserID,
	email valueobject.Email,
	phoneNumber valueobject.PhoneNumber,
	password string,
	role enum.UserRole,
	status enum.UserStatus,
	profile *Profile,
) (*User, error) {
	if !role.IsValid() {
		return nil, domainerr.ErrInvalidUserRole(string(role))
	}
	if !status.IsValid() {
		return nil, domainerr.ErrInvalidUserStatus(string(status))
	}

	if profile == nil {
		profile = &Profile{
			PhotoURL:    "",
			Bio:         "",
			Address:     &Address{},
			DateOfBirth: nil,
		}
	}

	return &User{
		id:          id,
		email:       email,
		phoneNumber: phoneNumber,
		password:    password,
		role:        role,
		status:      status,
	}, nil
}

func (u *User) Id() valueobject.UserID {
	return u.id
}

func (u *User) Email() valueobject.Email {
	return u.email
}

func (u *User) PhoneNumber() valueobject.PhoneNumber {
	return u.phoneNumber
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Role() enum.UserRole {
	return u.role
}

func (u *User) Status() enum.UserStatus {
	return u.status
}

func (u *User) LastLoginAt() time.Time {
	return u.lastLoginAt
}

func (u *User) SetLastLoginAt(t time.Time) {
	u.lastLoginAt = t
}

func (u *User) SetStatus(status enum.UserStatus) error {
	if !status.IsValid() {
		return domainerr.ErrInvalidUserStatus(string(status))
	}
	u.status = status
	return nil
}

func (u *User) SetRole(role enum.UserRole) error {
	if !role.IsValid() {
		return domainerr.ErrInvalidUserRole(string(role))
	}
	u.role = role
	return nil
}

func (u *User) UpdateEmail(email valueobject.Email) {
	u.email = email
}

func (u *User) UpdatePhoneNumber(phoneNumber valueobject.PhoneNumber) {
	u.phoneNumber = phoneNumber
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) JoinedAt() time.Time {
	return u.joinedAt
}

func (u *User) ChangePassword(newPassword string) error {
	u.password = newPassword
	return nil
}

func (u *User) UpdateStatus(status enum.UserStatus) error {
	if !status.IsValid() {
		return errors.New("invalid user status")
	}

	if err := u.ValidateStatusTransition(status); err != nil {
		return err
	}

	u.status = status
	return nil
}

func (u *User) ValidateStatusTransition(newStatus enum.UserStatus) error {
	switch u.status {
	case enum.UserStatusPending:
		if u.status != enum.UserStatusActive && u.status != enum.UserStatusDeleted {
			return errors.New("pending users can only be activated or deleted")
		}
	case enum.UserStatusBanned:
		if u.status != enum.UserStatusActive && u.status != enum.UserStatusDeleted {
			return errors.New("banned users can only be activated or deleted")
		}
	case enum.UserStatusDeleted:
		if u.status != enum.UserStatusActive {
			return errors.New("can only set deleted status from active")
		}
	case enum.UserStatusInactive:
		if u.status != enum.UserStatusActive && u.status != enum.UserStatusBanned {
			return errors.New("inactive users can only be activated or banned")
		}
	case enum.UserStatusActive:
		if u.status == enum.UserStatusPending {
			return errors.New("active users cannot be set to pending")
		}
	}
	return nil
}

func (u *User) Is2FAEnabled() bool {
	return u.twoFactorAuth.IsEnabled
}

func (u *User) Set2FAuth(secret string) {
	if secret == "" {
		u.twoFactorAuth = TwoFactorAuth{
			IsEnabled: false,
			Secret:    "",
		}
		return
	}

	u.twoFactorAuth = TwoFactorAuth{
		IsEnabled: true,
		Secret:    secret,
	}
}

type UserBuilder struct {
	user   *User
	errors []error
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user:   &User{},
		errors: []error{},
	}
}

func (b *UserBuilder) WithId(id int) *UserBuilder {
	var err error
	b.user.id, err = valueobject.NewUserID(id)
	if err != nil {
		b.errors = append(b.errors, err)
	}
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	var err error
	b.user.email, err = valueobject.NewEmail(email)
	if err != nil {
		b.errors = append(b.errors, err)
	}
	return b
}

func (b *UserBuilder) WithPhoneNumber(phoneNumber string) *UserBuilder {
	var err error
	b.user.phoneNumber, err = valueobject.NewPhoneNumber(phoneNumber)
	if err != nil {
		b.errors = append(b.errors, err)
	}
	return b
}

func (b *UserBuilder) WithPassword(password string) *UserBuilder {
	b.user.password = password
	return b
}

func (b *UserBuilder) WithRole(role enum.UserRole) *UserBuilder {
	b.user.role = role
	return b
}

func (b *UserBuilder) WithStatus(status enum.UserStatus) *UserBuilder {
	b.user.status = status
	return b
}

func (b *UserBuilder) WithJoinedAt(joinedAt time.Time) *UserBuilder {
	b.user.joinedAt = joinedAt
	return b
}

func (b *UserBuilder) WithTwoFactorAuth(secret string) *UserBuilder {
	b.user.Set2FAuth(secret)
	return b
}

func (b *UserBuilder) WithLastLogin(lastLoginAt time.Time) *UserBuilder {
	b.user.lastLoginAt = lastLoginAt
	return b
}

func (b *UserBuilder) Build() (User, error) {
	if len(b.errors) > 0 {
		return User{}, fmt.Errorf("errors occurred while building User: %v", b.errors)
	}
	return *b.user, nil
}

type Address struct {
	Street              string
	City                string
	State               string
	ZipCode             string
	Country             valueobject.Country
	BuildingType        valueobject.BuildingType
	BuildingOuterNumber string
	BuildingInnerNumber *string
}
