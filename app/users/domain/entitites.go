package user

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
)

const (
	MIN_PASSWORD_LENGTH = 8
	MAX_PASSWORD_LENGTH = 100
)

type User struct {
	id           valueObjects.UserId
	email        valueObjects.Email
	phoneNumber  valueObjects.PhoneNumber
	password     string
	role         UserRole
	status       UserStatus
	lastLoginAt  time.Time
	joinedAt     time.Time
	profile      *Profile
	is2FAEnabled bool
}

type Profile struct {
	UserId         int
	OwnerId        *int
	VeterinarianId *int
	Name           valueObjects.PersonName
	PhotoURL       string
	Bio            string
	Gender         valueObjects.Gender
	Address        *Address
	DateOfBirth    *time.Time
	JoinedAt       time.Time
}

func NewUser(id valueObjects.UserId, email valueObjects.Email, phoneNumber valueObjects.PhoneNumber, password string, role UserRole, status UserStatus, profile *Profile) (*User, error) {
	if !role.IsValid() {
		return nil, ErrInvalidUserRole(string(role))
	}
	if !status.IsValid() {
		return nil, ErrInvalidUserStatus(string(status))
	}

	if profile == nil {
		profile = &Profile{
			Name:        valueObjects.PersonName{},
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
		profile:     profile,
	}, nil
}

func (u *User) Id() valueObjects.UserId {
	return u.id
}

func (u *User) Email() valueObjects.Email {
	return u.email
}

func (u *User) PhoneNumber() valueObjects.PhoneNumber {
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

func (u *User) Profile() Profile {
	return *u.profile
}

func (u *User) SetLastLoginAt(t time.Time) {
	u.lastLoginAt = t
}

func (u *User) SetProfile(profile *Profile) {
	u.profile = profile
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

func (u *User) UpdateEmail(email valueObjects.Email) {
	u.email = email
}

func (u *User) UpdatePhoneNumber(phoneNumber valueObjects.PhoneNumber) {
	u.phoneNumber = phoneNumber
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) JoinedAt() time.Time {
	return u.joinedAt
}

func (u *User) ChangePassword(newPassword string) error {
	if err := ValidatePassword(newPassword); err != nil {
		return err
	}
	u.password = newPassword
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < MIN_PASSWORD_LENGTH {
		return ErrPasswordTooShort(strconv.Itoa(MIN_PASSWORD_LENGTH))
	}
	if len(password) > MAX_PASSWORD_LENGTH {
		return ErrPasswordTooLong(strconv.Itoa(MAX_PASSWORD_LENGTH))
	}

	passwordRegex := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+={}\[\]:;"'<>,.?/\\|-]+$`).MatchString(password)
	if !passwordRegex {
		return ErrPasswordInvalidFormat("Password must contain only alphanumeric characters and special symbols !@#$%^&*()_+={}[]:;\"'<>,.?/\\\\|-")
	}
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
	return u.is2FAEnabled
}
