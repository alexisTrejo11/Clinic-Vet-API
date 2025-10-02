package handler

import (
	c "clinic-vet-api/app/modules/account/auth/session/application/command"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/core/service"
	"clinic-vet-api/app/shared/auth"
	"context"
)

const (
	// Error messages
	errFindingUserMsg           = "failed to find user"
	errDeletingUserSessionMsg   = "failed to delete user session"
	errGeneratingAccessTokenMsg = "failed to generate access token"

	// Success messages
	userSessionsRevokedMsg = "all user sessions revoked successfully"
	sessionDeletedMsg      = "session deleted successfully"
	sessionRefreshedMsg    = "session refreshed successfully"
)

type SessionCommandHandler struct {
	userRepository repository.UserRepository
	sessionRepo    repository.SessionRepository
	jwtService     service.JWTService
}

func NewSessionCommandHandler(
	userRepo repository.UserRepository,
	sessionRepo repository.SessionRepository,
	jwtService service.JWTService,
) *SessionCommandHandler {
	return &SessionCommandHandler{
		userRepository: userRepo,
		sessionRepo:    sessionRepo,
		jwtService:     jwtService,
	}
}

func (h *SessionCommandHandler) HandleRevokeAllUserSessions(ctx context.Context, cmd c.RevokeAllUserSessionsCommand) auth.AuthCommandResult {
	user, err := h.userRepository.FindByID(ctx, cmd.UserID())
	if err != nil {
		return auth.AuthFailure(errFindingUserMsg, err)
	}

	err = h.sessionRepo.DeleteAllUserSessions(ctx, user.ID())
	if err != nil {
		return auth.AuthFailure(errDeletingUserSessionMsg, err)
	}

	return auth.AuthSuccess(userSessionsRevokedMsg)
}

func (h *SessionCommandHandler) HandleRevokeSession(ctx context.Context, cmd c.RevokeSessionCommand) auth.AuthCommandResult {
	_, err := h.userRepository.FindByID(ctx, cmd.UserID())
	if err != nil {
		return auth.AuthFailure(errFindingUserMsg, err)
	}

	err = h.sessionRepo.DeleteUserSession(ctx, cmd.UserID(), cmd.RefreshToken())
	if err != nil {
		return auth.AuthFailure(errDeletingUserSessionMsg, err)
	}

	return auth.AuthSuccess(sessionDeletedMsg)
}

func (h *SessionCommandHandler) HandleRefreshSession(ctx context.Context, cmd c.RefreshSessionCommand) auth.AuthCommandResult {
	session, err := h.sessionRepo.GetByRefreshToken(ctx, cmd.RefreshToken())
	if err != nil {
		return auth.AuthFailure(errFindingUserMsg, err)
	}

	access, err := h.jwtService.GenerateAccessToken(session.UserID)
	if err != nil {
		return auth.AuthFailure(errGeneratingAccessTokenMsg, err)
	}

	response := auth.GetSessionResponse(session, access)
	return auth.AuthSuccessWithSession(response, sessionRefreshedMsg)
}
