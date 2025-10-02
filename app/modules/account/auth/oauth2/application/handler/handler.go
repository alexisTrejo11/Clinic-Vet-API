package handler

/*
func (h *TwoFACommandHandler) HandleOAuthLogin(ctx context.Context, cmd c.OAuthLoginCommand) auth.AuthCommandResult {
	user, err := h.userAuthService.AuthenticateOAuthUser(ctx, cmd.Provider(), cmd.Token())
	if err != nil {
		return auth.AuthFailure(ErrAuthenticationFailed, err)
	}

	is2FALogin := false
	if err := user.ValidateLogin(is2FALogin); err != nil {
		return auth.AuthFailure(ErrAuthenticationFailed, err)
	}

	session, err := h.createSession(ctx, user.ID().String(), cmd.Metadata())
	if err != nil {
		return auth.AuthFailure(ErrSessionCreationFailed, err)
	}

	accessToken, err := h.jwtService.GenerateAccessToken(user.ID().String())
	if err != nil {
		return auth.AuthFailure(ErrAccessTokenGenFailed, err)
	}

	sessionResponse := GetSessionResponse(session, accessToken)
	return auth.AuthSuccessWithSession(sessionResponse, MsgLoginSuccess)
}
*/
