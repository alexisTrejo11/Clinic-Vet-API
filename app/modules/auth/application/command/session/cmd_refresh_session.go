package session

import (
	"context"

	"clinic-vet-api/app/modules/auth/application/command/result"
)

type RefreshUserSessionCommand struct {
	RefreshToken string `json:"session_id"`
}

func NewRefreshUserSessionCommand(refreshToken string) RefreshUserSessionCommand {
	return RefreshUserSessionCommand{
		RefreshToken: refreshToken,
	}
}

func (h *authCommandHandler) RefreshUserSession(ctx context.Context, cmd RefreshUserSessionCommand) result.AuthCommandResult {
	session, err := h.sessionRepo.GetByRefreshToken(ctx, cmd.RefreshToken)
	if err != nil {
		return result.AuthFailure("Session not found", err)
	}

	access, err := h.jwtService.GenerateAccessToken(session.UserID)
	if err != nil {
		return result.AuthFailure("Failed to generate access token", err)
	}

	response := result.GetSessionResponse(session, access)
	return result.AuthSuccess(response, session.RefreshToken[:3], "session successfully refreshed")
}
