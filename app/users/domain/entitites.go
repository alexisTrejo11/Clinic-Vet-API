package userDomain

import (
	"regexp"
	"strconv"
	"time"
)

const (
	MIN_PASSWORD_LENGTH = 8
	MAX_PASSWORD_LENGTH = 100
)

type User struct {
	id          UserId
	email       Email
	phoneNumber PhoneNumber
	password    string
	role        UserRole
	status      UserStatus
	lastLoginAt *time.Time
	profile     *Profile
}

type Profile struct {
	Name        PersonName
	PhotoURL    string
	Bio         string
	Location    string
	DateOfBirth *time.Time
	Gender      string
	JoinedAt    time.Time
}

func NewUser(id UserId, email Email, phoneNumber PhoneNumber, password string, role UserRole, status UserStatus, profile Profile) (*User, error) {
	if !role.IsValid() {
		return nil, ErrInvalidUserRole(string(role))
	}
	if !status.IsValid() {
		return nil, ErrInvalidUserStatus(string(status))
	}

	return &User{
		id:          id,
		email:       email,
		phoneNumber: phoneNumber,
		password:    password,
		role:        role,
		status:      status,
		profile:     &profile,
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

func (u *User) LastLoginAt() *time.Time {
	return u.lastLoginAt
}

func (u *User) Profile() Profile {
	return *u.profile
}

func (u *User) SetLastLoginAt(t time.Time) {
	u.lastLoginAt = &t
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

func (u *User) SetEmail(email Email) {
	u.email = email
}

func (u *User) SetPhoneNumber(phoneNumber PhoneNumber) {
	u.phoneNumber = phoneNumber
}

func (u *User) SetPassword(password string) {
	u.password = password
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
