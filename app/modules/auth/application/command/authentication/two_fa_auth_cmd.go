package authentication

import (
	"clinic-vet-api/app/modules/auth/application/command/result"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"context"
)

type TwoFactorLoginCommand struct {
	UserID     valueobject.UserID
	Token      string
	RememberMe bool
	metadata   LoginMetadata
}

func (h *authCommandHandler) TwoFactorLogin(ctx context.Context, command TwoFactorLoginCommand) result.AuthCommandResult {
	user, err := h.userRepository.FindByID(ctx, command.UserID)
	if err != nil {
		return result.AuthFailure(result.ErrUserNotFound, err)
	}

	if err := h.userAuthService.ValidateTwoFactorToken(user, command.Token); err != nil {
		return result.AuthFailure(result.ErrInvalidTwoFactorToken, err)
	}

	is2FALogin := true
	if err := user.ValidateLogin(is2FALogin); err != nil {
		return result.AuthFailure(result.ErrAuthenticationFailed, err)
	}

	session, err := h.createSession(ctx, user.ID().String(), command.metadata)
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
