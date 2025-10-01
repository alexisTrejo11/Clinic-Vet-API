package command

import "clinic-vet-api/app/modules/core/domain/valueobject"

type Disable2FACommand struct {
	userID valueobject.UserID
}

func NewDisable2FACommand(userID uint) *Disable2FACommand {
	return &Disable2FACommand{
		userID: valueobject.NewUserID(userID),
	}
}

func (c *Disable2FACommand) UserID() valueobject.UserID { return c.userID }
