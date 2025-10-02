package command

import "clinic-vet-api/app/modules/core/domain/valueobject"

type Send2FATokenCommand struct {
	userID valueobject.UserID
}

func NewSend2FATokenCommand(userID uint) Send2FATokenCommand {
	return Send2FATokenCommand{
		userID: valueobject.NewUserID(userID),
	}
}

func (c Send2FATokenCommand) UserID() valueobject.UserID { return c.userID }
