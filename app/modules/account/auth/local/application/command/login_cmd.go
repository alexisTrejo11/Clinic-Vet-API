package command

import (
	"clinic-vet-api/app/shared/auth"
	apperror "clinic-vet-api/app/shared/error/application"
)

type LoginCommand struct {
	identifier string
	password   string
	metadata   auth.LoginMetadata
}

func NewLoginCommand(identifier, password string, metadata auth.LoginMetadata) (LoginCommand, error) {
	cmd := &LoginCommand{
		identifier: identifier,
		password:   password,
		metadata:   metadata,
	}

	if err := cmd.validate(); err != nil {
		return LoginCommand{}, err
	}

	return *cmd, nil
}

func (c *LoginCommand) validate() error {
	if c.identifier == "" {
		return LoginCmdError("identifier", "identifier is required")
	}
	if c.password == "" {
		return LoginCmdError("password", "password is required")
	}
	return nil
}

func (c *LoginCommand) Identifier() string           { return c.identifier }
func (c *LoginCommand) Password() string             { return c.password }
func (c *LoginCommand) Metadata() auth.LoginMetadata { return c.metadata }

func LoginCmdError(field, issue string) error {
	return apperror.CommandDataValidationError("LoginCommand", field, issue)
}
