package session

import (
	"clinic-vet-api/app/modules/auth/application/command/result"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"context"
)

type RevokeAllUserSessionsCommand struct {
	userID valueobject.UserID
}

func NewRevokeAllUserSessionsCommand(useridnUInt uint) *RevokeAllUserSessionsCommand {
	return &RevokeAllUserSessionsCommand{
		userID: valueobject.NewUserID(useridnUInt),
	}
}

func (h *authCommandHandler) RevokeAllUserSessions(ctx context.Context, command RevokeAllUserSessionsCommand) result.AuthCommandResult {
	user, err := h.userRepository.FindByID(ctx, command.userID)
	if err != nil {
		return result.AuthFailure("an error occurred finding user", err)
	}

	err = h.sessionRepo.DeleteAllUserSessions(ctx, user.ID())
	if err != nil {
		return result.AuthFailure("an error occurred delete user sessions", err)
	}

	return result.AuthSuccess(nil, user.ID().String(), "user sessions successfully deleted on all devices")
}
