package bus

import (
	"clinic-vet-api/app/modules/auth/application/command"
	"clinic-vet-api/app/modules/auth/application/handler"
	"context"
)

type TwoFAAuthCommandBus interface {
	Enable2FA(ctx context.Context, cmd command.Enable2FACommand) handler.AuthCommandResult
	Disable2FA(ctx context.Context, cmd command.Disable2FACommand) handler.AuthCommandResult
	Verify2FA(ctx context.Context, cmd command.Verify2FACommand) handler.AuthCommandResult
	Send2FAToken(ctx context.Context, cmd command.Send2FATokenCommand) handler.AuthCommandResult
}

type twoFAAuthCommandBus struct {
	twoFAhandler handler.TwoFACommandHandler
	authHandler  handler.AuthCommandHandler
}

func NewTwoFAAuthCommandBus(twoFAhandler handler.TwoFACommandHandler, authHandler handler.AuthCommandHandler) TwoFAAuthCommandBus {
	return &twoFAAuthCommandBus{
		twoFAhandler: twoFAhandler,
		authHandler:  authHandler,
	}
}

func (b *twoFAAuthCommandBus) Enable2FA(ctx context.Context, cmd command.Enable2FACommand) handler.AuthCommandResult {
	return b.twoFAhandler.HandleEnable2FA(ctx, cmd)
}

func (b *twoFAAuthCommandBus) Disable2FA(ctx context.Context, cmd command.Disable2FACommand) handler.AuthCommandResult {
	return b.twoFAhandler.HandleDisable2FA(ctx, cmd)
}

func (b *twoFAAuthCommandBus) Verify2FA(ctx context.Context, cmd command.Verify2FACommand) handler.AuthCommandResult {
	return b.twoFAhandler.HandleVerify2FA(ctx, cmd)
}

func (b *twoFAAuthCommandBus) Send2FAToken(ctx context.Context, cmd command.Send2FATokenCommand) handler.AuthCommandResult {
	return b.twoFAhandler.HandleSend2FAToken(ctx, cmd)
}
