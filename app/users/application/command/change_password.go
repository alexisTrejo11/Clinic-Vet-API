package userCommand

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type ChangePasswordCommand struct {
	UserId      int
	OldPassword string
	NewPassword string
	CTX         context.Context
}

type ChangePasswordHandler struct {
	userRepository userRepository.UserRepository
}

func (c *ChangePasswordHandler) Handle(cmd any) CommandResult {
	command := cmd.(ChangePasswordCommand)

	user, err := c.userRepository.GetById(command.CTX, command.UserId)
	if err != nil {
		return FailureResult("failed to find user", err)
	}

	if err := c.changePassword(user, command.NewPassword, command.OldPassword); err != nil {
		return FailureResult("failed to change password", err)
	}

	if err := c.userRepository.Save(command.CTX, user); err != nil {
		return FailureResult("failed to update user", err)
	}

	return SuccesResult(user.Id().String(), "password changed successfully")
}

func (c *ChangePasswordHandler) changePassword(user *userDomain.User, newPassword, oldPassword string) error {
	err := shared.CheckPassword(user.Password(), oldPassword)
	if err != nil {
		return errors.New("invalid old password")
	}

	if err := user.ChangePassword(newPassword); err != nil {
		return err
	}

	hashedPassword, err := shared.HashPassword(newPassword)
	if err != nil {
		return err
	}

	user.SetPassword(hashedPassword)
	return nil
}
