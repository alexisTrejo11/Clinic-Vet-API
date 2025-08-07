package userCommand

import (
	"context"

	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type DeleteUserHandler struct {
	userRepository userRepository.UserRepository
}

func (d *DeleteUserHandler) Handle(ctx context.Context, userId int, softDelete bool) CommandResult {
	if _, err := d.userRepository.GetById(ctx, userId); err != nil {
		return FailureResult("failed to find user", err)
	}

	err := d.userRepository.Delete(ctx, userId, softDelete)
	if err != nil {
		return FailureResult("failed to delete user", err)
	}

	return FailureResult("User deleted successfully", nil)
}
