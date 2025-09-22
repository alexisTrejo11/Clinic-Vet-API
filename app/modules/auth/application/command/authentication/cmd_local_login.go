package authentication

import (
	"context"

	"clinic-vet-api/app/modules/auth/application/command/result"
)

type LoginCommand struct {
	Identifier string
	Password   string
	RememberMe bool
	Metadata   LoginMetadata
}

type LoginMetadata struct {
	IP         string
	UserAgent  string
	DeviceInfo string
}

func (h *authCommandHandler) Login(ctx context.Context, command LoginCommand) result.AuthCommandResult {
	user, err := h.userAuthService.AuthenticateUser(ctx, command.Identifier, command.Password)
	if err != nil {
		return result.AuthFailure(result.ErrAuthenticationFailed, err)
	}

	is2FALogin := false
	if err := user.ValidateLogin(is2FALogin); err != nil {
		return result.AuthFailure(result.ErrAuthenticationFailed, err)
	}

	session, err := h.createSession(ctx, user.ID().String(), command.Metadata)
	if err != nil {
		return result.AuthFailure(result.ErrSessionCreationFailed, err)
	}

	accessToken, err := h.jwtService.GenerateAccessToken(user.ID().String())
	if err != nil {
		return result.AuthFailure(result.ErrAccessTokenGenFailed, err)
	}

	return result.AuthSuccess(
		result.GetSessionResponse(session, accessToken),
		session.RefreshToken[:3],
		result.MsgLoginSuccess,
	)
}
