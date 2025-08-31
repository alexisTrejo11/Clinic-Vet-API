package command

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type ChangePasswordCommand struct {
	UserID      valueobject.UserID
	OldPassword string
	NewPassword string
	CTX         context.Context
}

type ChangePasswordHandler struct {
	userRepository repository.UserRepository
}

func NewChangePasswordHandler(userRepo repository.UserRepository) *ChangePasswordHandler {
	return &ChangePasswordHandler{userRepository: userRepo}
}

func (c *ChangePasswordHandler) Handle(cmd any) shared.CommandResult {
	command := cmd.(ChangePasswordCommand)

	user, err := c.userRepository.GetByID(command.CTX, command.UserID)
	if err != nil {
		return shared.FailureResult("failed to find user", err)
	}

	if err := c.changePassword(&user, command.NewPassword, command.OldPassword); err != nil {
		return shared.FailureResult("failed to change password", err)
	}

	if err := c.userRepository.Save(command.CTX, &user); err != nil {
		return shared.FailureResult("failed to update user", err)
	}

	return shared.SuccessResult(user.ID().String(), "password changed successfully")
}

func (c *ChangePasswordHandler) changePassword(user *entity.User, newPassword, oldPassword string) error {
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
