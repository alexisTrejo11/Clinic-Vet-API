// Package command implements the command handlers for user-related operations.
package command

import (
	"time"

	u "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user"
)

func fromCreateCommand(cmd CreateUserCommand) (*u.User, error) {
	now := time.Now()
	user, err := u.CreateUser(
		cmd.role,
		cmd.status,
		u.WithEmail(cmd.email),
		u.WithPassword(cmd.password),
		u.WithPhoneNumber(cmd.phoneNumber),
		u.WithLastLoginAt(now),
		u.WithJoinedAt(now),
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
