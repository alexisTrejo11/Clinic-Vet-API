package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type LogoutCommand struct {
	UserID       valueobject.UserID
	RefreshToken string
	CTX          context.Context
}

func (h *authCommandHandler) Logout(command LogoutCommand) AuthCommandResult {
	user, err := h.userRepository.FindByID(command.CTX, command.UserID)
	if err != nil {
		return FailureAuthResult("an error ocurred finding user", err)
	}

	err = h.sessionRepo.DeleteUserSession(command.CTX, command.UserID, command.RefreshToken)
	if err != nil {
		return FailureAuthResult("an error ocurred deleting session", err)
	}

	return SuccessAuthResult(nil, user.ID().String(), "session successfully deleted")
}
