package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/auth"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
)

type Verify2FACommand struct {
	userID valueobject.UserID
	code   string
	method auth.TwoFactorMethod
}

func NewVerify2FACommand(userID uint, code string, method auth.TwoFactorMethod) (Verify2FACommand, error) {
	cmd := &Verify2FACommand{
		userID: valueobject.NewUserID(userID),
		code:   code,
		method: method,
	}

	if err := cmd.Validate(); err != nil {
		return Verify2FACommand{}, err
	}

	return *cmd, nil
}

func (c *Verify2FACommand) UserID() valueobject.UserID   { return c.userID }
func (c *Verify2FACommand) Code() string                 { return c.code }
func (c *Verify2FACommand) Method() auth.TwoFactorMethod { return c.method }

func (c *Verify2FACommand) Validate() error {
	if c.userID.IsZero() {
		return Verify2FACmdErr("user ID", "is required")
	}

	if c.code == "" {
		return Verify2FACmdErr("2FA code", "is required")
	}

	return nil
}

func Verify2FACmdErr(msg, issue string) error {
	return apperror.CommandDataValidationError(msg, issue, "Verify2FACommand")
}
