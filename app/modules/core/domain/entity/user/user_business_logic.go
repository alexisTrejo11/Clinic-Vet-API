package user

import (
	"errors"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/auth"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

func (u *User) UpdateEmail(newEmail valueobject.Email) error {
	if u.email.String() == newEmail.String() {
		return nil // No change needed
	}
	u.email = newEmail
	u.IncrementVersion()
	return nil
}

func (u *User) UpdatePhoneNumber(newPhone valueobject.PhoneNumber) error {
	u.phoneNumber = newPhone
	u.IncrementVersion()
	return nil
}

func (u *User) UpdatePassword(newPassword string) error {
	if len(newPassword) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	u.password = newPassword
	u.IncrementVersion()
	return nil
}

func (u *User) Activate() error {
	if u.status == enum.UserStatusBanned {
		return errors.New("cannot activate banned user")
	}
	u.status = enum.UserStatusActive
	u.IncrementVersion()
	return nil
}

func (u *User) Deactivate() error {
	u.status = enum.UserStatusInactive
	u.IncrementVersion()
	return nil
}

func (u *User) Ban() error {
	u.status = enum.UserStatusBanned
	u.IncrementVersion()
	return nil
}

func (u *User) RecordLogin() error {
	now := time.Now()
	u.lastLoginAt = &now
	u.IncrementVersion()
	return nil
}

func (u *User) EnableTwoFactorAuth(method, secret string, backupCodes []string) error {
	now := time.Now()

	twoFAMethod := auth.TwoFactorMethod(secret)

	twoFA := auth.NewTwoFactorAuth(true,
		twoFAMethod,
		secret,
		backupCodes,
		&now,
		&now,
	)
	u.twoFactorAuth = twoFA
	u.IncrementVersion()
	return nil
}

func (u *User) DisableTwoFactorAuth() error {
	u.twoFactorAuth = auth.NewDisabledTwoFactorAuth()
	u.IncrementVersion()
	return nil
}

func (u *User) ValidateLogin(Is2FA bool) error {
	if u.status != enum.UserStatusActive {
		return errors.New("user is not activated")
	}
	if u.status == enum.UserStatusBanned {
		return errors.New("user is banned")
	}

	if Is2FA && !u.twoFactorAuth.Enabled() {
		return errors.New("two-factor authentication is required to login")
	}

	return nil
}

func (u *User) AssignProfile(profileID uint) error {
	if profileID != 0 {
		return errors.New("profile already assigned")
	}
	u.profileID = profileID
	u.IncrementVersion()
	return nil
}

func (u *User) AssignCustomer(customerID valueobject.CustomerID) error {
	if u.customerID != nil {
		return errors.New("customer already assigned")
	}
	if u.employeeID != nil {
		return errors.New("user is an employee, cannot assign customer")
	}
	u.customerID = &customerID
	u.IncrementVersion()
	return nil
}

func (u *User) AssignEmployee(employeeID valueobject.EmployeeID) error {
	if u.employeeID != nil {
		return errors.New("employee already assigned")
	}
	if u.customerID != nil {
		return errors.New("user is a customer, cannot assign employee")
	}
	u.employeeID = &employeeID
	u.IncrementVersion()
	return nil
}
