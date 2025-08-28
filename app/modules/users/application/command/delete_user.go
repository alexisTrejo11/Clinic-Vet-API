package userDomainCommand

import (
	"context"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type deleteUserHandler struct {
	userRepository userDomain.UserRepository
}

type DeleteUserCommand struct {
	UserId     int             `json:"user_id" validate:"required"`
	SoftDelete bool            `json:"soft_delete"`
	CTX        context.Context `json:"-"`
}

func NewDeleteUserHandler(userRepo userDomain.UserRepository) DeleteUserHandler {
	return &deleteUserHandler{
		userRepository: userRepo,
	}
}

type DeleteUserHandler interface {
	Handle(cmd any) shared.CommandResult
}

func (d *deleteUserHandler) Handle(cmd any) shared.CommandResult {
	command := cmd.(DeleteUserCommand)

	if _, err := d.userRepository.GetByID(command.CTX, command.UserId); err != nil {
		return shared.FailureResult("failed to find user", err)
	}

	err := d.userRepository.Delete(command.CTX, command.UserId, command.SoftDelete)
	if err != nil {
		return shared.FailureResult("failed to delete user", err)
	}

	return shared.SuccessResult(strconv.Itoa(command.UserId), "user deleted successfully")
}
