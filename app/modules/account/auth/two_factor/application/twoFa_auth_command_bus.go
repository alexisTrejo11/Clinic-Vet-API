package application

import (
	"clinic-vet-api/app/modules/account/auth/two_factor/application/command"
	"clinic-vet-api/app/modules/account/auth/two_factor/application/handler"
	"clinic-vet-api/app/shared/auth"
	"context"
)

type TwoFaAuthFacadeService interface {
	Enable2FA(ctx context.Context, cmd command.Enable2FACommand) auth.AuthCommandResult
	Disable2FA(ctx context.Context, cmd command.Disable2FACommand) auth.AuthCommandResult
	Verify2FA(ctx context.Context, cmd command.Verify2FACommand) auth.AuthCommandResult
	Send2FAToken(ctx context.Context, cmd command.Send2FATokenCommand) auth.AuthCommandResult
	TwoFactorLogin(ctx context.Context, cmd command.TwoFactorLoginCommand) auth.AuthCommandResult
}

type twoFaAuthFacadeService struct {
	twoFAhandler *handler.TwoFACommandHandler
}

func NewTwoFaAuthFacadeService(twoFAhandler *handler.TwoFACommandHandler) TwoFaAuthFacadeService {
	return &twoFaAuthFacadeService{
		twoFAhandler: twoFAhandler,
	}
}

func (b *twoFaAuthFacadeService) Enable2FA(ctx context.Context, cmd command.Enable2FACommand) auth.AuthCommandResult {
	return b.twoFAhandler.HandleEnable2FA(ctx, cmd)
}

func (b *twoFaAuthFacadeService) Disable2FA(ctx context.Context, cmd command.Disable2FACommand) auth.AuthCommandResult {
	return b.twoFAhandler.HandleDisable2FA(ctx, cmd)
}

func (b *twoFaAuthFacadeService) Verify2FA(ctx context.Context, cmd command.Verify2FACommand) auth.AuthCommandResult {
	return b.twoFAhandler.HandleVerify2FA(ctx, cmd)
}

func (b *twoFaAuthFacadeService) Send2FAToken(ctx context.Context, cmd command.Send2FATokenCommand) auth.AuthCommandResult {
	return b.twoFAhandler.HandleSend2FAToken(ctx, cmd)
}

func (b *twoFaAuthFacadeService) TwoFactorLogin(ctx context.Context, cmd command.TwoFactorLoginCommand) auth.AuthCommandResult {
	return b.twoFAhandler.Handle2FaLogin(ctx, cmd)
}
