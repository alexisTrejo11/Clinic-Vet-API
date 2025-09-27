// Package user defines the User entity and its related functionalities.
package user

import (
	"clinic-vet-api/app/modules/core/domain/entity/auth"
	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type User struct {
	base.Entity[valueobject.UserID]
	email         valueobject.Email
	phoneNumber   *valueobject.PhoneNumber
	password      string
	role          enum.UserRole
	status        enum.UserStatus
	lastLoginAt   *time.Time
	twoFactorAuth auth.TwoFactorAuth
	employeeID    *valueobject.EmployeeID
	customerID    *valueobject.CustomerID
	profileID     uint
}

type UserBuilder struct{ u *User }

func NewUserBuilder() *UserBuilder { return &UserBuilder{u: &User{}} }

func (b *UserBuilder) WithID(id valueobject.UserID) *UserBuilder {
	b.u.SetID(id)
	return b
}

func (b *UserBuilder) WithEmail(email valueobject.Email) *UserBuilder {
	b.u.email = email
	return b
}

func (b *UserBuilder) WithPhoneNumber(phone *valueobject.PhoneNumber) *UserBuilder {
	b.u.phoneNumber = phone
	return b
}

func (b *UserBuilder) WithPassword(password string) *UserBuilder {
	b.u.password = password
	return b
}

func (b *UserBuilder) WithLastLoginAt(lastLogin time.Time) *UserBuilder {
	b.u.lastLoginAt = &lastLogin
	return b
}

func (b *UserBuilder) WithTwoFactorAuth(twoFA auth.TwoFactorAuth) *UserBuilder {
	b.u.twoFactorAuth = twoFA
	return b
}

func (b *UserBuilder) WithCustomerID(customerID *valueobject.CustomerID) *UserBuilder {
	b.u.customerID = customerID
	return b
}

func (b *UserBuilder) WithEmployeeID(employeeID *valueobject.EmployeeID) *UserBuilder {
	b.u.employeeID = employeeID
	return b
}

func (b *UserBuilder) WithTimeStamps(createdAt, updatedAt time.Time) *UserBuilder {
	b.u.SetTimeStamps(createdAt, updatedAt)
	return b
}

func (b *UserBuilder) WithRole(role enum.UserRole) *UserBuilder {
	b.u.role = role
	return b
}

func (b *UserBuilder) WithStatus(status enum.UserStatus) *UserBuilder {
	b.u.status = status
	return b
}

func (b *UserBuilder) Build() *User {
	return b.u
}

func (u *User) ID() valueobject.UserID                { return u.Entity.ID() }
func (u *User) HashedPassword() string                { return u.password }
func (u *User) Email() valueobject.Email              { return u.email }
func (u *User) PhoneNumber() *valueobject.PhoneNumber { return u.phoneNumber }
func (u *User) Role() enum.UserRole                   { return u.role }
func (u *User) Status() enum.UserStatus               { return u.status }
func (u *User) LastLoginAt() *time.Time               { return u.lastLoginAt }
func (u *User) JoinedAt() time.Time                   { return u.CreatedAt() }
func (u *User) EmployeeID() *valueobject.EmployeeID   { return u.employeeID }
func (u *User) IsCustomer() bool                      { return u.customerID != nil && u.role == enum.UserRoleCustomer }
func (u *User) CustomerID() *valueobject.CustomerID   { return u.customerID }
func (u *User) ProfileID() uint                       { return u.profileID }
func (u *User) TwoFactorAuth() auth.TwoFactorAuth     { return u.twoFactorAuth }
func (u *User) IsTwoFactorEnabled() bool              { return u.twoFactorAuth.Enabled() }

func (u *User) SetHashedPassword(hashedPassword string) { u.password = hashedPassword }
func (u *User) IsEmployee() bool                        { return u.employeeID != nil && u.role.IsStaff() }
func (u *User) IsActive() bool                          { return u.status == enum.UserStatusActive }
