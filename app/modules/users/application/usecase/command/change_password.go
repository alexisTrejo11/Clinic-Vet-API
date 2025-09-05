package command

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/password"
)

type ChangePasswordCommand struct {
	UserID      valueobject.UserID
	OldPassword string
	NewPassword string
	CTX         context.Context
}

type ChangePasswordHandler struct {
	userRepository  repository.UserRepository
	passwordEncoder password.PasswordEncoder
}

func NewChangePasswordHandler(userRepo repository.UserRepository, passwordEncoder password.PasswordEncoder) *ChangePasswordHandler {
	return &ChangePasswordHandler{
		userRepository:  userRepo,
		passwordEncoder: passwordEncoder,
	}
}

func (c *ChangePasswordHandler) Handle(cmd any) cqrs.CommandResult {
	command := cmd.(ChangePasswordCommand)

	user, err := c.userRepository.GetByID(command.CTX, command.UserID.GetValue())
	if err != nil {
		return cqrs.FailureResult("failed to find user", err)
	}

	if err := c.changePassword(&user, command); err != nil {
		return cqrs.FailureResult("failed to change password", err)
	}

	if err := c.userRepository.Save(command.CTX, &user); err != nil {
		return cqrs.FailureResult("failed to update user", err)
	}

	return cqrs.SuccessResult(user.ID().String(), "password changed successfully")
}

func (c *ChangePasswordHandler) changePassword(user *entity.User, command ChangePasswordCommand) error {
	err := c.passwordEncoder.CheckPassword(user.Password(), command.OldPassword)
	if err != nil {
		return errors.New("invalid current password")
	}

	if err := user.ChangePassword(command.NewPassword); err != nil {
		return err
	}

	hashedPassword, err := c.passwordEncoder.HashPassword(command.NewPassword)
	if err != nil {
		return err
	}

	user.SetPassword(hashedPassword)
	c.userRepository.Save(command.CTX, user)

	return nil
}
