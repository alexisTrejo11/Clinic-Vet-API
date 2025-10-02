package handler

import (
	"clinic-vet-api/app/modules/auth/application/command"
	"clinic-vet-api/app/modules/auth/infrastructure/token"
	"clinic-vet-api/app/modules/core/domain/entity/notification"
	"clinic-vet-api/app/modules/core/repository"
	service "clinic-vet-api/app/modules/notifications/application"
	"context"
)

var (
	ErrInvalid2FAToken = "invalid 2FA token"
)

type TwoFACommandHandler struct {
	userRepo     repository.UserRepository
	token        token.TokenManager
	notifService service.NotificationService
}

func NewTwoFACommandHandler(
	userRepo repository.UserRepository,
	token token.TokenManager,
	notifService service.NotificationService,

) *TwoFACommandHandler {
	return &TwoFACommandHandler{
		userRepo:     userRepo,
		token:        token,
		notifService: notifService,
	}
}

func (h *TwoFACommandHandler) HandleEnable2FA(ctx context.Context, cmd command.Enable2FACommand) AuthCommandResult {
	user, err := h.userRepo.FindByID(ctx, cmd.UserID())
	if err != nil {
		return AuthFailure("an error ocurred while fetching user", err)
	}

	if err := user.Validate2FAEnable(cmd.Method()); err != nil {
		return AuthFailure("an error ocurred while validating 2FA", err)
	}

	token, err := h.token.GenerateToken(ctx, token.VerificationToken, token.TokenConfig{UserID: user.ID().String()})
	if err != nil {
		return AuthFailure("an error ocurred while generating token", err)
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
			return AuthFailure("an error ocurred while sending notification", err)
		}
	}
	return AuthSuccess("2FA enable initiated, please verify using the sent code")
}

func (h *TwoFACommandHandler) HandleDisable2FA(ctx context.Context, cmd command.Disable2FACommand) AuthCommandResult {
	user, err := h.userRepo.FindByID(ctx, cmd.UserID())
	if err != nil {
		return AuthFailure("an error ocurred while fetching user", err)
	}

	if err := user.Disable2FA(); err != nil {
		return AuthFailure("an error ocurred while disabling 2FA", err)
	}

	if err := h.userRepo.Save(ctx, &user); err != nil {
		return AuthFailure("an error ocurred while saving user", err)
	}
	return AuthSuccess("2FA disabled successfully")
}

func (h *TwoFACommandHandler) HandleVerify2FA(ctx context.Context, cmd command.Verify2FACommand) AuthCommandResult {
	user, err := h.userRepo.FindByID(ctx, cmd.UserID())
	if err != nil {
		return AuthFailure("an error ocurred while fetching user", err)
	}

	if _, err := h.token.ValidateToken(ctx, cmd.Code(), token.VerificationToken); err != nil {
		return AuthFailure("an error ocurred while validating token", err)
	}

	if err := user.Enable2FA(cmd.Method()); err != nil {
		return AuthFailure("an error ocurred while enabling 2FA", err)
	}

	if err := h.userRepo.Save(ctx, &user); err != nil {
		return AuthFailure("an error ocurred while saving user", err)
	}

	return AuthSuccess("2FA enabled successfully")
}

func (h *TwoFACommandHandler) HandleSend2FAToken(ctx context.Context, cmd command.Send2FATokenCommand) AuthCommandResult {
	user, err := h.userRepo.FindByID(ctx, cmd.UserID())
	if err != nil {
		return AuthFailure("an error ocurred while fetching user", err)
	}

	if !user.IsTwoFactorEnabled() {
		return AuthFailure("2FA is not enabled for this user", nil)
	}

	token, err := h.token.GenerateToken(ctx, token.TwoFAToken, token.TokenConfig{UserID: user.ID().String()})
	if err != nil {
		return AuthFailure("an error ocurred while generating token", err)
	}

	notification.NewSend2FAToken(user.TwoFactorAuth().Method(), token, user.ID(), user.Email(), user.PhoneNumber())

	return AuthSuccess("2FA token sent successfully")
}
