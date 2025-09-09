package user

import (
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/auth"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
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

func (u *User) CanLogin() bool {
	return u.status == enum.UserStatusActive &&
		u.email.String() != "" &&
		u.password != ""
}
