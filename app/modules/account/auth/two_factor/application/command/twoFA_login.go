package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	shared "clinic-vet-api/app/shared/auth"
	apperror "clinic-vet-api/app/shared/error/application"
)

type TwoFactorLoginCommand struct {
	userID   valueobject.UserID
	token    string
	metadata shared.LoginMetadata
}

func NewTwoFactorLoginCommand(
	userID valueobject.UserID,
	token string,
	metadata shared.LoginMetadata,
) (TwoFactorLoginCommand, error) {
	cmd := &TwoFactorLoginCommand{
		userID:   userID,
		token:    token,
		metadata: metadata,
	}

	if err := cmd.Validate(); err != nil {
		return TwoFactorLoginCommand{}, err
	}

	return *cmd, nil
}

func (c *TwoFactorLoginCommand) Validate() error {
	if c.userID.IsZero() {
		return TwoFaLoginCmdError("userID", "must be provided")
	}
	if c.token == "" {
		return TwoFaLoginCmdError("token", "must be provided")
	}
	return nil
}

func (c *TwoFactorLoginCommand) UserID() valueobject.UserID     { return c.userID }
func (c *TwoFactorLoginCommand) Token() string                  { return c.token }
func (c *TwoFactorLoginCommand) Metadata() shared.LoginMetadata { return c.metadata }

func TwoFaLoginCmdError(field, issue string) error {
	return apperror.CommandDataValidationError("TwoFactorLoginCommand", field, issue)
}
