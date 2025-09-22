package session

import (
	"clinic-vet-api/app/modules/auth/application/command/result"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"context"
)

type RevokeUserSessionCommand struct {
	UserID       valueobject.UserID
	RefreshToken string
}

func (h *authCommandHandler) RevokeUserSession(ctx context.Context, cmd RevokeUserSessionCommand) result.AuthCommandResult {
	user, err := h.userRepository.FindByID(ctx, cmd.UserID)
	if err != nil {
		return result.AuthFailure("an error ocurred finding user", err)
	}

	err = h.sessionRepo.DeleteUserSession(ctx, cmd.UserID, cmd.RefreshToken)
	if err != nil {
		return result.AuthFailure("an error ocurred deleting session", err)
	}

	return result.AuthSuccess(nil, user.ID().String(), "session successfully deleted")
}
