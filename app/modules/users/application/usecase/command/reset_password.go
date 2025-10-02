package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
)

type ResetPasswordCommand struct {
	email       valueobject.Email
	token       string
	newPassword string
}

func NewResetPasswordCommand(email string, token string, newPassword string) (ResetPasswordCommand, error) {
	cmd := &ResetPasswordCommand{
		email:       valueobject.NewEmailNoErr(email),
		token:       token,
		newPassword: newPassword,
	}

	if err := cmd.Validate(); err != nil {
		return ResetPasswordCommand{}, err
	}

	return *cmd, nil
}

func (cmd *ResetPasswordCommand) Email() valueobject.Email { return cmd.email }
func (cmd *ResetPasswordCommand) Token() string            { return cmd.token }
func (cmd *ResetPasswordCommand) NewPassword() string      { return cmd.newPassword }

func (cmd *ResetPasswordCommand) Validate() error {
	if !cmd.email.IsValid() {
		return ResetPasswordCmdErr("email", "Invalid email format: "+cmd.email.String())
	}
	if cmd.token == "" {
		return ResetPasswordCmdErr("token", "Token is required")
	}
	if cmd.newPassword == "" {
		return ResetPasswordCmdErr("newPassword", "New password is required")
	}
	return nil
}

func ResetPasswordCmdErr(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "ResetPasswordCommand")
}
