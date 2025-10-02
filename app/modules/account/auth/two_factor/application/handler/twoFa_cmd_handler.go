package handler

import (
	token "clinic-vet-api/app/modules/account/auth/token/factory"
	"clinic-vet-api/app/modules/account/auth/two_factor/application/command"
	"time"

	s "clinic-vet-api/app/modules/core/domain/entity/auth"
	"clinic-vet-api/app/modules/core/domain/entity/notification"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/core/service"
	"clinic-vet-api/app/shared/auth"
	"context"
)

const (
	// Error messages
	ErrInvalid2FAToken       = "invalid 2FA token"
	ErrUserNotFound          = "user not found"
	ErrInvalidTwoFactorToken = "invalid two-factor authentication token"
	ErrAuthenticationFailed  = "authentication failed"
	ErrSessionCreationFailed = "failed to create session"
	ErrAccessTokenGenFailed  = "failed to generate access token"
	ErrFetchingUser          = "an error ocurred while fetching user"
	ErrValidating2FA         = "an error ocurred while validating 2FA"
	ErrGeneratingToken       = "an error ocurred while generating token"
	ErrSendingNotification   = "an error ocurred while sending notification"
	ErrDisabling2FA          = "an error ocurred while disabling 2FA"
	ErrSavingUser            = "an error ocurred while saving user"
	ErrValidatingToken       = "an error ocurred while validating token"
	ErrEnabling2FA           = "an error ocurred while enabling 2FA"
	Err2FANotEnabled         = "2FA is not enabled for this user"

	// Success messages
	Msg2FAEnableInitiated  = "2FA enable initiated, please verify using the sent code"
	Msg2FADisabledSuccess  = "2FA disabled successfully"
	Msg2FAEnabledSuccess   = "2FA enabled successfully"
	Msg2FATokenSentSuccess = "2FA token sent successfully"
	MsgLoginSuccess        = "login successful"
)

type TwoFACommandHandler struct {
	userRepo     repository.UserRepository
	sessionRepo  repository.SessionRepository
	authService  service.UserSecurityService
	token        repository.TokenRepository
	notifService service.NotificationService
	jwtService   service.JWTService
}

func NewTwoFACommandHandler(
	userRepo repository.UserRepository,
	sessionRepo repository.SessionRepository,
	authService service.UserSecurityService,
	token repository.TokenRepository,
	notifService service.NotificationService,
	jwtService service.JWTService,
) *TwoFACommandHandler {
	return &TwoFACommandHandler{
		userRepo:     userRepo,
		token:        token,
		notifService: notifService,
		sessionRepo:  sessionRepo,
		jwtService:   jwtService,
		authService:  authService,
	}
}

func (h *TwoFACommandHandler) HandleEnable2FA(ctx context.Context, cmd command.Enable2FACommand) auth.AuthCommandResult {
	user, err := h.userRepo.FindByID(ctx, cmd.UserID())
	if err != nil {
		return auth.AuthFailure(ErrFetchingUser, err)
	}

	if err := user.Validate2FAEnable(cmd.Method()); err != nil {
		return auth.AuthFailure(ErrValidating2FA, err)
	}

	token, err := h.token.GenerateToken(ctx, vo.VerificationToken, token.TokenConfig{UserID: user.ID().String()})
	if err != nil {
		return auth.AuthFailure(ErrGeneratingToken, err)
	}

	var notif *notification.Notification
	if cmd.Method() == "email" {
		notif = notification.NewEnable2FANotificationEmail(user.ID(), user.Email(), token)
	}

	if cmd.Method() == "sms" {
		notif = notification.NewEnable2FANotificationSMS(user.ID(), *user.PhoneNumber(), token)
	}

	if notif != nil {
		if err := h.notifService.Send(ctx, notif); err != nil {
			return auth.AuthFailure(ErrSendingNotification, err)
		}
	}
	return auth.AuthSuccess(Msg2FAEnableInitiated)
}

func (h *TwoFACommandHandler) HandleDisable2FA(ctx context.Context, cmd command.Disable2FACommand) auth.AuthCommandResult {
	user, err := h.userRepo.FindByID(ctx, cmd.UserID())
	if err != nil {
		return auth.AuthFailure(ErrFetchingUser, err)
	}

	if err := user.Disable2FA(); err != nil {
		return auth.AuthFailure(ErrDisabling2FA, err)
	}

	if err := h.userRepo.Save(ctx, &user); err != nil {
		return auth.AuthFailure(ErrSavingUser, err)
	}
	return auth.AuthSuccess(Msg2FADisabledSuccess)
}

func (h *TwoFACommandHandler) HandleVerify2FA(ctx context.Context, cmd command.Verify2FACommand) auth.AuthCommandResult {
	user, err := h.userRepo.FindByID(ctx, cmd.UserID())
	if err != nil {
		return auth.AuthFailure(ErrFetchingUser, err)
	}

	if _, err := h.token.ValidateToken(ctx, cmd.Code(), vo.VerificationToken); err != nil {
		return auth.AuthFailure(ErrValidatingToken, err)
	}

	if err := user.Enable2FA(cmd.Method()); err != nil {
		return auth.AuthFailure(ErrEnabling2FA, err)
	}

	if err := h.userRepo.Save(ctx, &user); err != nil {
		return auth.AuthFailure(ErrSavingUser, err)
	}

	return auth.AuthSuccess(Msg2FAEnabledSuccess)
}

func (h *TwoFACommandHandler) HandleSend2FAToken(ctx context.Context, cmd command.Send2FATokenCommand) auth.AuthCommandResult {
	user, err := h.userRepo.FindByID(ctx, cmd.UserID())
	if err != nil {
		return auth.AuthFailure(ErrFetchingUser, err)
	}

	if !user.IsTwoFactorEnabled() {
		return auth.AuthFailure(Err2FANotEnabled, nil)
	}

	token, err := h.token.GenerateToken(ctx, vo.TwoFAToken, token.TokenConfig{UserID: user.ID().String()})
	if err != nil {
		return auth.AuthFailure(ErrGeneratingToken, err)
	}

	notification.NewSend2FAToken(user.TwoFactorAuth().Method(), token, user.ID(), user.Email(), user.PhoneNumber())

	return auth.AuthSuccess(Msg2FATokenSentSuccess)
}

func (h *TwoFACommandHandler) Handle2FaLogin(ctx context.Context, cmd command.TwoFactorLoginCommand) auth.AuthCommandResult {
	user, err := h.userRepo.FindByID(ctx, cmd.UserID())
	if err != nil {
		return auth.AuthFailure(ErrUserNotFound, err)
	}

	if err := h.authService.ValidateTwoFactorToken(user, cmd.Token()); err != nil {
		return auth.AuthFailure(ErrInvalidTwoFactorToken, err)
	}

	is2FALogin := true
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

	response := auth.GetSessionResponse(session, accessToken)
	return auth.AuthSuccessWithSession(response, MsgLoginSuccess)
}

func (h *TwoFACommandHandler) createSession(
	ctx context.Context, userID string, metadata auth.LoginMetadata,
) (s.Session, error) {
	refreshToken, err := h.jwtService.GenerateRefreshToken(userID)
	if err != nil {
		return s.Session{}, err
	}

	now := time.Now()
	sessionDuration := auth.DefaultSessionDuration
	newSession := &s.Session{
		UserID:       userID,
		IPAddress:    metadata.IP(),
		RefreshToken: refreshToken,
		CreatedAt:    now,
		ExpiresAt:    now.Add(sessionDuration),
		DeviceInfo:   metadata.DeviceInfo(),
		UserAgent:    metadata.UserAgent(),
	}

	if err := h.sessionRepo.Create(ctx, newSession); err != nil {
		return s.Session{}, err
	}

	return *newSession, nil
}
