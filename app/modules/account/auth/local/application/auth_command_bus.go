package application

import (
	c "clinic-vet-api/app/modules/account/auth/local/application/command"
	h "clinic-vet-api/app/modules/account/auth/local/application/handler"
	sc "clinic-vet-api/app/modules/account/auth/session/application/command"
	"clinic-vet-api/app/shared/auth"

	"context"
)

type AuthApplicationFacade interface {
	RegisterCustomer(ctx context.Context, cmd c.RegisterCustomerCommand) auth.AuthCommandResult
	RegisterEmployee(ctx context.Context, cmd c.RegisterEmployeeCommand) auth.AuthCommandResult
	Login(ctx context.Context, cmd c.LoginCommand) auth.AuthCommandResult
	Logout(ctx context.Context, cmd sc.RevokeSessionCommand) auth.AuthCommandResult
	LogoutAll(ctx context.Context, cmd sc.RevokeAllUserSessionsCommand) auth.AuthCommandResult
}

type authCommandBus struct {
	*h.AuthCommandHandler
}

func NewAuthCommandBus(authHandler *h.AuthCommandHandler) AuthApplicationFacade {
	return &authCommandBus{authHandler}
}

func (b *authCommandBus) RegisterCustomer(ctx context.Context, cmd c.RegisterCustomerCommand) auth.AuthCommandResult {
	return b.AuthCommandHandler.HandleRegisterCustomer(ctx, cmd)
}

func (b *authCommandBus) RegisterEmployee(ctx context.Context, cmd c.RegisterEmployeeCommand) auth.AuthCommandResult {
	return b.AuthCommandHandler.HandleRegisterEmployee(ctx, cmd)
}

func (b *authCommandBus) Login(ctx context.Context, cmd c.LoginCommand) auth.AuthCommandResult {
	return b.AuthCommandHandler.HandleLogin(ctx, cmd)
}

func (b *authCommandBus) LogoutAll(ctx context.Context, cmd sc.RevokeAllUserSessionsCommand) auth.AuthCommandResult {
	//return b.AuthCommandHandler.HandleRevokeAllUserSessions(ctx, cmd)
	return auth.AuthSuccess("Not implemented yet")
}

func (b *authCommandBus) Logout(ctx context.Context, cmd sc.RevokeSessionCommand) auth.AuthCommandResult {
	//return b.AuthCommandHandler.HandleRevokeSession(ctx, cmd)
	return auth.AuthSuccess("Not implemented yet")
}
