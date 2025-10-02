package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
)

type ActivateAccountCommand struct {
	token  string
	userID valueobject.UserID
}

func NewActivateAccountCommand(token string, userID uint) (ActivateAccountCommand, error) {
	cmd := &ActivateAccountCommand{
		token:  token,
		userID: valueobject.NewUserID(userID),
	}

	if err := cmd.Validate(); err != nil {
		return ActivateAccountCommand{}, err
	}

	return *cmd, nil
}

func (cmd *ActivateAccountCommand) Token() string              { return cmd.token }
func (cmd *ActivateAccountCommand) UserID() valueobject.UserID { return cmd.userID }

func (cmd *ActivateAccountCommand) Validate() error {
	if cmd.token == "" {
		return activateAccCmdErr("token", "Token is required")
	}
	if cmd.userID.IsZero() {
		return activateAccCmdErr("user_id", "User ID is required")
	}

	return nil
}

func activateAccCmdErr(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "ActivateAccountCommand")
}
