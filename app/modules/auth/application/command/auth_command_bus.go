// Package command contains command handlers for authentication operations
package command

import (
	"clinic-vet-api/app/modules/auth/application/command/authentication"
	"clinic-vet-api/app/modules/auth/application/command/register"
	"clinic-vet-api/app/modules/auth/application/command/session"
)

type AuthCommandBus interface {
	authentication.AuthCommandHandler
	register.RegisterCommandHandler
	session.SessionCommandHandler
}

type authCommandBus struct {
	authentication.AuthCommandHandler
	register.RegisterCommandHandler
	session.SessionCommandHandler
}

func NewAuthCommandBus(
	authHandler authentication.AuthCommandHandler,
	registHandler register.RegisterCommandHandler,
	sessionHandler session.SessionCommandHandler,
) AuthCommandBus {
	return &authCommandBus{
		authHandler,
		registHandler,
		sessionHandler,
	}
}
