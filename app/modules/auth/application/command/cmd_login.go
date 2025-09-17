package command

import (
	"context"
	"time"

	"clinic-vet-api/app/core/domain/entity/auth"
)

// LoginCommand represents the login request data
type LoginCommand struct {
	Identifier string          `json:"identifier" validate:"required"`
	Password   string          `json:"password" validate:"required"`
	RememberMe bool            `json:"remember_me"`
	IP         string          `json:"ip" validate:"required"`
	UserAgent  string          `json:"user_agent"`
	DeviceInfo string          `json:"source"`
	CTX        context.Context `json:"-"`
}

func (h *authCommandHandler) Login(command LoginCommand) AuthCommandResult {
	user, err := h.userAuthService.AuthenticateUser(command.CTX, command.Identifier, command.Password)
	if err != nil {
		return FailureAuthResult(ErrAuthenticationFailed, err)
	}

	is2FALogin := false
	if err := user.ValidateLogin(is2FALogin); err != nil {
		return FailureAuthResult(ErrAuthenticationFailed, err)
	}

	session, err := h.createSession(user.ID().String(), command)
	if err != nil {
		return FailureAuthResult(ErrSessionCreationFailed, err)
	}

	accessToken, err := h.jwtService.GenerateAccessToken(user.ID().String())
	if err != nil {
		return FailureAuthResult(ErrAccessTokenGenFailed, err)
	}

	return SuccessAuthResult(
		getSessionResponse(session, accessToken),
		session.ID,
		MsgLoginSuccess,
	)
}

func (h *authCommandHandler) createSession(userID string, command LoginCommand) (auth.Session, error) {
	refreshToken, err := h.jwtService.GenerateRefreshToken(userID)
	if err != nil {
		return auth.Session{}, err
	}

	now := time.Now()
	sessionDuration := DefaultSessionDuration

	if command.RememberMe {
		sessionDuration = 30 * 24 * time.Hour // 30 days
	}

	newSession := auth.Session{
		UserID:       userID,
		IPAddress:    command.IP,
		RefreshToken: refreshToken,
		CreatedAt:    now,
		ExpiresAt:    now.Add(sessionDuration),
		DeviceInfo:   command.DeviceInfo,
		UserAgent:    command.UserAgent,
	}

	if err := h.sessionRepo.Create(command.CTX, &newSession); err != nil {
		return auth.Session{}, err
	}

	return newSession, nil
}
