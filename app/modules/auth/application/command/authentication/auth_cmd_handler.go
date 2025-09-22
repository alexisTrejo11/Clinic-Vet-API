package authentication

import (
	"clinic-vet-api/app/modules/auth/application/command/result"
	"clinic-vet-api/app/modules/auth/application/jwt"
	"clinic-vet-api/app/modules/core/domain/entity/auth"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/core/service"
	"context"
	"time"
)

type AuthCommandHandler interface {
	Login(ctx context.Context, command LoginCommand) result.AuthCommandResult
	TwoFactorLogin(ctx context.Context, command TwoFactorLoginCommand) result.AuthCommandResult
	OAuthLogin(ctx context.Context, command OAuthLoginCommand) result.AuthCommandResult
}

type authCommandHandler struct {
	userRepository  repository.UserRepository
	userAuthService service.UserSecurityService
	sessionRepo     repository.SessionRepository
	jwtService      jwt.JWTService
}

func NewLoginCommandHandler(
	userRepository repository.UserRepository,
	userAuthService service.UserSecurityService,
	sessionRepo repository.SessionRepository,
	jwtService jwt.JWTService,
) AuthCommandHandler {
	return &authCommandHandler{
		userRepository:  userRepository,
		userAuthService: userAuthService,
		sessionRepo:     sessionRepo,
		jwtService:      jwtService,
	}
}

func (h *authCommandHandler) createSession(ctx context.Context, userID string, metadata LoginMetadata) (auth.Session, error) {

	refreshToken, err := h.jwtService.GenerateRefreshToken(userID)
	if err != nil {
		return auth.Session{}, err
	}

	now := time.Now()
	sessionDuration := result.DefaultSessionDuration

	newSession := &auth.Session{
		UserID:       userID,
		IPAddress:    metadata.IP,
		RefreshToken: refreshToken,
		CreatedAt:    now,
		ExpiresAt:    now.Add(sessionDuration),
		DeviceInfo:   metadata.DeviceInfo,
		UserAgent:    metadata.UserAgent,
	}

	if err := h.sessionRepo.Create(ctx, newSession); err != nil {
		return auth.Session{}, err
	}

	return *newSession, nil
}
