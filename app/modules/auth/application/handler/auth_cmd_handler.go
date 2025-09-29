package handler

import (
	c "clinic-vet-api/app/modules/auth/application/command"
	"clinic-vet-api/app/modules/auth/application/jwt"
	"clinic-vet-api/app/modules/core/domain/entity/auth"
	event "clinic-vet-api/app/modules/core/domain/event/user_event"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/core/service"
	"clinic-vet-api/app/shared/password"
	"context"
	"time"
)

type AuthCommandHandler struct {
	userRepository  repository.UserRepository
	userAuthService service.UserSecurityService
	sessionRepo     repository.SessionRepository
	jwtService      jwt.JWTService
	passwordEncoder password.PasswordEncoder
	userEvent       event.UserEventProducer
}

func NewAuthCommandHandler(
	userRepository repository.UserRepository,
	userAuthService service.UserSecurityService,
	sessionRepo repository.SessionRepository,
	jwtService jwt.JWTService,
	passwordEncoder password.PasswordEncoder,
	userEvent event.UserEventProducer,
) *AuthCommandHandler {
	return &AuthCommandHandler{
		userRepository:  userRepository,
		userAuthService: userAuthService,
		sessionRepo:     sessionRepo,
		jwtService:      jwtService,
		passwordEncoder: passwordEncoder,
		userEvent:       userEvent,
	}
}

func (h *AuthCommandHandler) HandleLogin(ctx context.Context, cmd c.LoginCommand) AuthCommandResult {
	user, err := h.userAuthService.AuthenticateUser(ctx, cmd.Identifier(), cmd.Password())
	if err != nil {
		return AuthFailure(ErrAuthenticationFailed, err)
	}

	is2FALogin := false
	if err := user.ValidateLogin(is2FALogin); err != nil {
		return AuthFailure(ErrAuthenticationFailed, err)
	}

	session, err := h.createSession(ctx, user.ID().String(), cmd.Metadata())
	if err != nil {
		return AuthFailure(ErrSessionCreationFailed, err)
	}

	accessToken, err := h.jwtService.GenerateAccessToken(user.ID().String())
	if err != nil {
		return AuthFailure(ErrAccessTokenGenFailed, err)
	}

	sessionResponse := GetSessionResponse(session, accessToken)
	return AuthSuccessWithSession(sessionResponse, MsgLoginSuccess)
}

func (h *AuthCommandHandler) HandleOAuthLogin(ctx context.Context, cmd c.OAuthLoginCommand) AuthCommandResult {
	user, err := h.userAuthService.AuthenticateOAuthUser(ctx, cmd.Provider(), cmd.Token())
	if err != nil {
		return AuthFailure(ErrAuthenticationFailed, err)
	}

	is2FALogin := false
	if err := user.ValidateLogin(is2FALogin); err != nil {
		return AuthFailure(ErrAuthenticationFailed, err)
	}

	session, err := h.createSession(ctx, user.ID().String(), cmd.Metadata())
	if err != nil {
		return AuthFailure(ErrSessionCreationFailed, err)
	}

	accessToken, err := h.jwtService.GenerateAccessToken(user.ID().String())
	if err != nil {
		return AuthFailure(ErrAccessTokenGenFailed, err)
	}

	sessionResponse := GetSessionResponse(session, accessToken)
	return AuthSuccessWithSession(sessionResponse, MsgLoginSuccess)
}

func (h *AuthCommandHandler) Handle2FaLogin(ctx context.Context, cmd c.TwoFactorLoginCommand) AuthCommandResult {
	user, err := h.userRepository.FindByID(ctx, cmd.UserID())
	if err != nil {
		return AuthFailure(ErrUserNotFound, err)
	}

	if err := h.userAuthService.ValidateTwoFactorToken(user, cmd.Token()); err != nil {
		return AuthFailure(ErrInvalidTwoFactorToken, err)
	}

	is2FALogin := true
	if err := user.ValidateLogin(is2FALogin); err != nil {
		return AuthFailure(ErrAuthenticationFailed, err)
	}

	session, err := h.createSession(ctx, user.ID().String(), cmd.Metadata())
	if err != nil {
		return AuthFailure(ErrSessionCreationFailed, err)
	}

	accessToken, err := h.jwtService.GenerateAccessToken(user.ID().String())
	if err != nil {
		return AuthFailure(ErrAccessTokenGenFailed, err)
	}

	response := GetSessionResponse(session, accessToken)
	return AuthSuccessWithSession(response, MsgLoginSuccess)
}

func (h *AuthCommandHandler) createSession(
	ctx context.Context, userID string, metadata c.LoginMetadata,
) (auth.Session, error) {
	refreshToken, err := h.jwtService.GenerateRefreshToken(userID)
	if err != nil {
		return auth.Session{}, err
	}

	now := time.Now()
	sessionDuration := DefaultSessionDuration
	newSession := &auth.Session{
		UserID:       userID,
		IPAddress:    metadata.IP(),
		RefreshToken: refreshToken,
		CreatedAt:    now,
		ExpiresAt:    now.Add(sessionDuration),
		DeviceInfo:   metadata.DeviceInfo(),
		UserAgent:    metadata.UserAgent(),
	}

	if err := h.sessionRepo.Create(ctx, newSession); err != nil {
		return auth.Session{}, err
	}

	return *newSession, nil
}
