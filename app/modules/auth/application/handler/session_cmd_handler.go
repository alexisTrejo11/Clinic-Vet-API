package handler

import (
	c "clinic-vet-api/app/modules/auth/application/command"
	"context"
)

func (h *AuthCommandHandler) HandleRevokeAllUserSessions(ctx context.Context, cmd c.RevokeAllUserSessionsCommand) AuthCommandResult {
	user, err := h.userRepository.FindByID(ctx, cmd.UserID())
	if err != nil {
		return AuthFailure(errFindingUserMsg, err)
	}

	err = h.sessionRepo.DeleteAllUserSessions(ctx, user.ID())
	if err != nil {
		return AuthFailure(errDeletingUserSessionMsg, err)
	}

	return AuthSuccess(userSessionsRevokedMsg)
}

func (h *AuthCommandHandler) HandleRevokeSession(ctx context.Context, cmd c.RevokeSessionCommand) AuthCommandResult {
	_, err := h.userRepository.FindByID(ctx, cmd.UserID())
	if err != nil {
		return AuthFailure(errFindingUserMsg, err)
	}

	err = h.sessionRepo.DeleteUserSession(ctx, cmd.UserID(), cmd.RefreshToken())
	if err != nil {
		return AuthFailure(errDeletingUserSessionMsg, err)
	}

	return AuthSuccess(sessionDeletedMsg)
}

func (h *AuthCommandHandler) HandleRefreshSession(ctx context.Context, cmd c.RefreshSessionCommand) AuthCommandResult {
	session, err := h.sessionRepo.GetByRefreshToken(ctx, cmd.RefreshToken())
	if err != nil {
		return AuthFailure(errFindingUserMsg, err)
	}

	access, err := h.jwtService.GenerateAccessToken(session.UserID)
	if err != nil {
		return AuthFailure("Failed to generate access token", err)
	}

	response := GetSessionResponse(session, access)
	return AuthSuccessWithSession(response, sessionRefreshedMsg)
}
