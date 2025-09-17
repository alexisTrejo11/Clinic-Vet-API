package command

import (
	"context"
	"errors"

	"clinic-vet-api/app/core/domain/entity/user"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	"clinic-vet-api/app/shared/password"
)

type ChangePasswordCommand struct {
	UserID      valueobject.UserID
	OldPassword string
	NewPassword string
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

func (c *ChangePasswordHandler) Handle(ctx context.Context, command ChangePasswordCommand) cqrs.CommandResult {
	user, err := c.userRepository.FindByID(ctx, command.UserID)
	if err != nil {
		return *cqrs.FailureResult("failed to find user", err)
	}

	if err := c.changePassword(ctx, &user, command); err != nil {
		return *cqrs.FailureResult("failed to change password", err)
	}

	if err := c.userRepository.Save(ctx, &user); err != nil {
		return *cqrs.FailureResult("failed to update user", err)
	}

	return *cqrs.SuccessResult(user.ID().String(), "password changed successfully")
}

func (c *ChangePasswordHandler) changePassword(ctx context.Context, user *user.User, command ChangePasswordCommand) error {
	err := c.passwordEncoder.CheckPassword(user.Password(), command.OldPassword)
	if err != nil {
		return errors.New("invalid current password")
	}

	if err := user.UpdatePassword(command.NewPassword); err != nil {
		return err
	}

	hashedPassword, err := c.passwordEncoder.HashPassword(command.NewPassword)
	if err != nil {
		return err
	}

	user.SetHashedPassword(hashedPassword)

	return c.userRepository.Save(ctx, user)
}
