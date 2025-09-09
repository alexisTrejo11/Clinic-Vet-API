package command

import (
	"time"

	u "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

func fromCreateCommand(cmd CreateUserCommand) (*u.User, error) {
	userIDZero, _ := valueobject.NewUserID(0)
	now := time.Now()
	user, err := u.NewUser(
		userIDZero,
		cmd.role,
		cmd.status,
		u.WithEmail(cmd.email),
		u.WithPassword(cmd.password),
		u.WithPhoneNumber(cmd.phoneNumber),
		u.WithLastLoginAt(now),
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
