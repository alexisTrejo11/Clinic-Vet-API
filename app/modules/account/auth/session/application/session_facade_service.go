package application

import (
	c "clinic-vet-api/app/modules/account/auth/session/application/command"
	h "clinic-vet-api/app/modules/account/auth/session/application/handler"
	"clinic-vet-api/app/shared/auth"
	"context"
)

type SessionFacadeService interface {
	RefreshSession(ctx context.Context, cmd c.RefreshSessionCommand) auth.AuthCommandResult
	RevokeSession(ctx context.Context, cmd c.RevokeSessionCommand) auth.AuthCommandResult
	RevokeAllSessions(ctx context.Context, cmd c.RevokeAllUserSessionsCommand) auth.AuthCommandResult
}

type sessionFacadeService struct {
	commandHandler h.SessionCommandHandler
}

func NewSessionFacadeService(commandHandler *h.SessionCommandHandler) SessionFacadeService {
	return &sessionFacadeService{
		commandHandler: *commandHandler,
	}
}

func (s *sessionFacadeService) RefreshSession(ctx context.Context, cmd c.RefreshSessionCommand) auth.AuthCommandResult {
	return s.commandHandler.HandleRefreshSession(ctx, cmd)
}

func (s *sessionFacadeService) RevokeSession(ctx context.Context, cmd c.RevokeSessionCommand) auth.AuthCommandResult {
	return s.commandHandler.HandleRevokeSession(ctx, cmd)
}

func (s *sessionFacadeService) RevokeAllSessions(ctx context.Context, cmd c.RevokeAllUserSessionsCommand) auth.AuthCommandResult {
	return s.commandHandler.HandleRevokeAllUserSessions(ctx, cmd)
}
