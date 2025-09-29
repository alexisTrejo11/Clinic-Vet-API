package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type RevokeAllUserSessionsCommand struct {
	userID valueobject.UserID
}

func NewRevokeAllUserSessionsCommand(useridnUInt uint) *RevokeAllUserSessionsCommand {
	return &RevokeAllUserSessionsCommand{
		userID: valueobject.NewUserID(useridnUInt),
	}
}

func (c *RevokeAllUserSessionsCommand) UserID() valueobject.UserID {
	return c.userID
}
