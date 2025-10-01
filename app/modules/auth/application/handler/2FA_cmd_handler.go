package handler

import (
	"clinic-vet-api/app/modules/auth/application/command"
	"clinic-vet-api/app/modules/auth/infrastructure/token"
	"clinic-vet-api/app/modules/core/domain/entity/notification"
	"clinic-vet-api/app/modules/core/repository"
	service "clinic-vet-api/app/modules/notifications/application"
	"context"
)

type TwoFACommandHandler struct {
	userRepo     repository.UserRepository
	token        token.TokenManager
	notifService service.NotificationService
}

func NewTwoFACommandHandler(userRepo repository.UserRepository, token token.TokenManager, notifService service.NotificationService) *TwoFACommandHandler {
	return &TwoFACommandHandler{
		userRepo:     userRepo,
		token:        token,
		notifService: notifService,
	}
}

func (h *TwoFACommandHandler) HandleEnable2FA(ctx context.Context, cmd command.Enable2FACommand) error {
	user, err := h.userRepo.FindByID(ctx, cmd.UserID())
	if err != nil {
		return err
	}

	if err := user.Validate2FAEnable(cmd.Method()); err != nil {
		return err
	}

	token, err := h.token.GenerateToken(ctx, token.VerificationToken, token.TokenConfig{UserID: user.ID().String()})
	if err != nil {
		return err
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
			return err
		}
	}
	return nil
}

func (h *TwoFACommandHandler) HandleDisable2FA(ctx context.Context, cmd command.Disable2FACommand) error {
	user, err := h.userRepo.FindByID(ctx, cmd.UserID())
	if err != nil {
		return err
	}

	if err := user.Disable2FA(); err != nil {
		return err
	}

	if err := h.userRepo.Save(ctx, &user); err != nil {
		return err
	}
	return nil
}

func (h *TwoFACommandHandler) HandleVerify2FA(ctx context.Context, cmd command.Verify2FACommand) error {
	user, err := h.userRepo.FindByID(ctx, cmd.UserID())
	if err != nil {
		return err
	}

	if _, err := h.token.ValidateToken(ctx, cmd.Code(), token.VerificationToken); err != nil {
		return err
	}

	if err := user.Enable2FA(cmd.Method()); err != nil {
		return err
	}

	if err := h.userRepo.Save(ctx, &user); err != nil {
		return err
	}

	return nil
}
