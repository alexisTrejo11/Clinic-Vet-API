package userDomain

import (
	"fmt"
	"time"
)

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
	b.user.id, err = NewUserId(id)
	if err != nil {
		b.errors = append(b.errors, err)
	}
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	var err error
	b.user.email, err = NewEmail(email)
	if err != nil {
		b.errors = append(b.errors, err)
	}
	return b
}

func (b *UserBuilder) WithPhoneNumber(phoneNumber string) *UserBuilder {
	var err error
	b.user.phoneNumber, err = NewPhoneNumber(phoneNumber)
	if err != nil {
		b.errors = append(b.errors, err)
	}
	return b
}

func (b *UserBuilder) WithPassword(password string) *UserBuilder {
	b.user.password = password
	return b
}

func (b *UserBuilder) WithRole(role UserRole) *UserBuilder {
	b.user.role = role
	return b
}

func (b *UserBuilder) WithStatus(status UserStatus) *UserBuilder {
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
