package authentication

import (
	"clinic-vet-api/app/modules/auth/application/command/result"
	"context"
)

type OAuthLoginCommand struct {
	Provider string
	Token    string
	Metadata LoginMetadata
}

func (h *authCommandHandler) OAuthLogin(ctx context.Context, command OAuthLoginCommand) result.AuthCommandResult {
	user, err := h.userAuthService.AuthenticateOAuthUser(ctx, command.Provider, command.Token)
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
		session.ID,
		result.MsgLoginSuccess,
	)
}
