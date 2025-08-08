package userCommand

import (
	"context"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type DeleteUserHandler struct {
	userRepository userRepository.UserRepository
}

type DeleteUserCommand struct {
	UserId     user.UserId     `json:"user_id" validate:"required"`
	SoftDelete bool            `json:"soft_delete"`
	CTX        context.Context `json:"-"`
}

func (d *DeleteUserHandler) Handle(command DeleteUserCommand) shared.CommandResult {
	if _, err := d.userRepository.GetByIdWithProfile(command.CTX, command.UserId.GetValue()); err != nil {
		return shared.FailureResult("failed to find user", err)
	}

	err := d.userRepository.Delete(command.CTX, command.UserId.GetValue(), command.SoftDelete)
	if err != nil {
		return shared.FailureResult("failed to delete user", err)
	}

	return shared.SuccesResult(strconv.Itoa(command.UserId.GetValue()), "user deleted successfully")
}
