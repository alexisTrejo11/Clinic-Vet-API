package userCommand

import (
	"context"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type deleteUserHandler struct {
	userRepository userRepository.UserRepository
}

type DeleteUserCommand struct {
	UserId     int             `json:"user_id" validate:"required"`
	SoftDelete bool            `json:"soft_delete"`
	CTX        context.Context `json:"-"`
}

func NewDeleteUserHandler(userRepo userRepository.UserRepository) DeleteUserHandler {
	return &deleteUserHandler{
		userRepository: userRepo,
	}
}

type DeleteUserHandler interface {
	Handle(cmd any) shared.CommandResult
}

func (d *deleteUserHandler) Handle(cmd any) shared.CommandResult {
	command := cmd.(DeleteUserCommand)

	if _, err := d.userRepository.GetByIdWithProfile(command.CTX, command.UserId); err != nil {
		return shared.FailureResult("failed to find user", err)
	}

	err := d.userRepository.Delete(command.CTX, command.UserId, command.SoftDelete)
	if err != nil {
		return shared.FailureResult("failed to delete user", err)
	}

	return shared.SuccesResult(strconv.Itoa(command.UserId), "user deleted successfully")
}
