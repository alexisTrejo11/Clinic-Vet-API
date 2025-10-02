package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/auth"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
)

type Enable2FACommand struct {
	userID valueobject.UserID
	method auth.TwoFactorMethod
}

func NewEnable2FACommand(userID uint, method string) (Enable2FACommand, error) {
	cmd := &Enable2FACommand{
		userID: valueobject.NewUserID(userID),
		method: auth.TwoFactorMethod(method),
	}

	if err := cmd.validate(); err != nil {
		return Enable2FACommand{}, err
	}

	return *cmd, nil
}

func (c *Enable2FACommand) UserID() valueobject.UserID   { return c.userID }
func (c *Enable2FACommand) Method() auth.TwoFactorMethod { return c.method }

func (c *Enable2FACommand) validate() error {
	if c.userID.IsZero() {
		return Enable2FACmdErr("user ID", "is required")
	}

	if !c.method.IsValid() {
		return Enable2FACmdErr("2FA method", "is invalid")
	}

	return nil
}

func Enable2FACmdErr(msg, issue string) error {
	return apperror.CommandDataValidationError(msg, issue, "Enable2FACommand")
}
