package userDomain

import (
	"errors"
	"time"
)

type User struct {
	id            UserId
	email         Email
	phoneNumber   PhoneNumber
	password      string
	role          UserRole
	status        UserStatus
	lastLoginAt   time.Time
	joinedAt      time.Time
	twoFactorAuth TwoFactorAuth
}

type Profile struct {
	UserId         int
	OwnerId        *int
	VeterinarianId *int
	PhotoURL       string
	Bio            string
	Address        *Address
	DateOfBirth    *time.Time
	JoinedAt       time.Time
}

func NewUser(id UserId, email Email, phoneNumber PhoneNumber, password string, role UserRole, status UserStatus, profile *Profile) (*User, error) {
	if !role.IsValid() {
		return nil, ErrInvalidUserRole(string(role))
	}
	if !status.IsValid() {
		return nil, ErrInvalidUserStatus(string(status))
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

func (u *User) Id() UserId {
	return u.id
}

func (u *User) Email() Email {
	return u.email
}

func (u *User) PhoneNumber() PhoneNumber {
	return u.phoneNumber
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Role() UserRole {
	return u.role
}

func (u *User) Status() UserStatus {
	return u.status
}

func (u *User) LastLoginAt() time.Time {
	return u.lastLoginAt
}

func (u *User) SetLastLoginAt(t time.Time) {
	u.lastLoginAt = t
}

func (u *User) SetStatus(status UserStatus) error {
	if !status.IsValid() {
		return ErrInvalidUserStatus(string(status))
	}
	u.status = status
	return nil
}

func (u *User) SetRole(role UserRole) error {
	if !role.IsValid() {
		return ErrInvalidUserRole(string(role))
	}
	u.role = role
	return nil
}

func (u *User) UpdateEmail(email Email) {
	u.email = email
}

func (u *User) UpdatePhoneNumber(phoneNumber PhoneNumber) {
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

func (u *User) UpdateStatus(status UserStatus) error {
	if !status.IsValid() {
		return errors.New("invalid user status")
	}

	if err := u.ValidateStatusTransition(status); err != nil {
		return err
	}

	u.status = status
	return nil
}

func (u *User) ValidateStatusTransition(newStatus UserStatus) error {
	switch u.status {
	case UserStatusPending:
		if u.status != UserStatusActive && u.status != UserStatusDeleted {
			return errors.New("pending users can only be activated or deleted")
		}
	case UserStatusBanned:
		if u.status != UserStatusActive && u.status != UserStatusDeleted {
			return errors.New("banned users can only be activated or deleted")
		}
	case UserStatusDeleted:
		if u.status != UserStatusActive {
			return errors.New("can only set deleted status from active")
		}
	case UserStatusInactive:
		if u.status != UserStatusActive && u.status != UserStatusBanned {
			return errors.New("inactive users can only be activated or banned")
		}
	case UserStatusActive:
		if u.status == UserStatusPending {
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
