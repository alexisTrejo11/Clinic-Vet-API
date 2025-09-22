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
	phoneNumber   valueobject.PhoneNumber
	password      string
	role          enum.UserRole
	status        enum.UserStatus
	lastLoginAt   *time.Time
	twoFactorAuth auth.TwoFactorAuth
	employeeID    *valueobject.EmployeeID
	customerID    *valueobject.CustomerID
	profileID     uint
}

func (u *User) ID() valueobject.UserID {
	return u.Entity.ID()
}

func (u *User) HashedPassword() string {
	return u.password
}

func (u *User) Email() valueobject.Email {
	return u.email
}

func (u *User) PhoneNumber() valueobject.PhoneNumber {
	return u.phoneNumber
}

func (u *User) Role() enum.UserRole {
	return u.role
}

func (u *User) Status() enum.UserStatus {
	return u.status
}

func (u *User) LastLoginAt() *time.Time {
	return u.lastLoginAt
}

func (u *User) JoinedAt() time.Time {
	return u.CreatedAt()
}

func (u *User) EmployeeID() *valueobject.EmployeeID {
	return u.employeeID
}

func (u *User) IsCustomer() bool {
	return u.customerID != nil && u.role == enum.UserRoleCustomer
}

func (u *User) CustomerID() *valueobject.CustomerID {
	return u.customerID
}

func (u *User) ProfileID() uint {
	return u.profileID
}

func (u *User) TwoFactorAuth() auth.TwoFactorAuth {
	return u.twoFactorAuth
}

func (u *User) IsTwoFactorEnabled() bool {
	return u.twoFactorAuth.Enabled()
}

func (u *User) SetHashedPassword(hashedPassword string) {
	u.password = hashedPassword
}

func (u *User) IsEmployee() bool {
	return u.employeeID != nil && u.role.IsStaff()
}
