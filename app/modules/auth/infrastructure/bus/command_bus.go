// Package bus implements the command bus for the auth module.
package bus

import "clinic-vet-api/app/modules/auth/application/command"

type AuthBus struct {
	CommandBus command.AuthCommandHandler
}

func NewAuthBus(commandBus command.AuthCommandHandler) *AuthBus {
	return &AuthBus{
		CommandBus: commandBus,
	}
}

func (bus *AuthBus) CustomerRegister(command command.CustomerRegisterCommand) command.AuthCommandResult {
	return bus.CommandBus.CustomerRegister(command)
}

func (bus *AuthBus) EmployeeRegister(command command.EmployeeRegisterCommand) command.AuthCommandResult {
	return bus.CommandBus.EmployeeRegister(command)
}

func (bus *AuthBus) Login(command command.LoginCommand) command.AuthCommandResult {
	return bus.CommandBus.Login(command)
}

func (bus *AuthBus) RefreshSession(command command.RefreshSessionCommand) command.AuthCommandResult {
	return bus.CommandBus.RefreshSession(command)
}

func (bus *AuthBus) Logout(command command.LogoutCommand) command.AuthCommandResult {
	return bus.CommandBus.Logout(command)
}

func (bus *AuthBus) LogoutAll(command command.LogoutAllCommand) command.AuthCommandResult {
	return bus.CommandBus.LogoutAll(command)
}
