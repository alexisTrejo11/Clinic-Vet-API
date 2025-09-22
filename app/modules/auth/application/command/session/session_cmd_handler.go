package session

import (
	"clinic-vet-api/app/modules/auth/application/command/result"
	"clinic-vet-api/app/modules/auth/application/jwt"
	"clinic-vet-api/app/modules/core/repository"
	"context"
)

type SessionCommandHandler interface {
	RevokeUserSession(ctx context.Context, cmd RevokeUserSessionCommand) result.AuthCommandResult
	RevokeAllUserSessions(ctx context.Context, cmd RevokeAllUserSessionsCommand) result.AuthCommandResult
	RefreshUserSession(ctx context.Context, cmd RefreshUserSessionCommand) result.AuthCommandResult
}

type authCommandHandler struct {
	userRepository repository.UserRepository
	sessionRepo    repository.SessionRepository
	jwtService     jwt.JWTService
}

func NewSessionCommandHandler(
	userRepo repository.UserRepository,
	sessionRepo repository.SessionRepository,
	jwtService jwt.JWTService,
) SessionCommandHandler {
	return &authCommandHandler{
		userRepository: userRepo,
		sessionRepo:    sessionRepo,
		jwtService:     jwtService,
	}
}
