// Package bus contains command handlers for handler operations
package bus

import (
	c "clinic-vet-api/app/modules/auth/application/command"
	h "clinic-vet-api/app/modules/auth/application/handler"
	"context"
)

type AuthCommandBus interface {
	RegisterCustomer(ctx context.Context, cmd c.RegisterCustomerCommand) h.AuthCommandResult
	RegisterEmployee(ctx context.Context, cmd c.RegisterEmployeeCommand) h.AuthCommandResult
	Login(ctx context.Context, cmd c.LoginCommand) h.AuthCommandResult
	TwoFactorLogin(ctx context.Context, cmd c.TwoFactorLoginCommand) h.AuthCommandResult
	OAuthLogin(ctx context.Context, cmd c.OAuthLoginCommand) h.AuthCommandResult
	RevokeAllUserSessions(ctx context.Context, cmd c.RevokeAllUserSessionsCommand) h.AuthCommandResult
	RevokeSession(ctx context.Context, cmd c.RevokeSessionCommand) h.AuthCommandResult
	RefreshSession(ctx context.Context, cmd c.RefreshSessionCommand) h.AuthCommandResult
}

type authCommandBus struct {
	h.AuthCommandHandler
}

func NewAuthCommandBus(authHandler h.AuthCommandHandler) AuthCommandBus {
	return &authCommandBus{authHandler}
}

func (b *authCommandBus) RegisterCustomer(ctx context.Context, cmd c.RegisterCustomerCommand) h.AuthCommandResult {
	return b.AuthCommandHandler.HandleRegisterCustomer(ctx, cmd)
}

func (b *authCommandBus) RegisterEmployee(ctx context.Context, cmd c.RegisterEmployeeCommand) h.AuthCommandResult {
	return b.AuthCommandHandler.HandleRegisterEmployee(ctx, cmd)
}

func (b *authCommandBus) Login(ctx context.Context, cmd c.LoginCommand) h.AuthCommandResult {
	return b.AuthCommandHandler.HandleLogin(ctx, cmd)
}

func (b *authCommandBus) TwoFactorLogin(ctx context.Context, cmd c.TwoFactorLoginCommand) h.AuthCommandResult {
	return b.AuthCommandHandler.Handle2FaLogin(ctx, cmd)
}

func (b *authCommandBus) OAuthLogin(ctx context.Context, cmd c.OAuthLoginCommand) h.AuthCommandResult {
	return b.AuthCommandHandler.HandleOAuthLogin(ctx, cmd)
}

func (b *authCommandBus) RevokeAllUserSessions(ctx context.Context, cmd c.RevokeAllUserSessionsCommand) h.AuthCommandResult {
	return b.AuthCommandHandler.HandleRevokeAllUserSessions(ctx, cmd)
}

func (b *authCommandBus) RevokeSession(ctx context.Context, cmd c.RevokeSessionCommand) h.AuthCommandResult {
	return b.AuthCommandHandler.HandleRevokeSession(ctx, cmd)
}

func (b *authCommandBus) RefreshSession(ctx context.Context, cmd c.RefreshSessionCommand) h.AuthCommandResult {
	return b.AuthCommandHandler.HandleRefreshSession(ctx, cmd)
}
