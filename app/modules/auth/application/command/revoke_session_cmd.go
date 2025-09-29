package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type RevokeSessionCommand struct {
	userID       valueobject.UserID
	refreshToken string
}

func NewRevokeSessionCommand(userID valueobject.UserID, refreshToken string) RevokeSessionCommand {
	return RevokeSessionCommand{
		userID:       userID,
		refreshToken: refreshToken,
	}
}

func (c RevokeSessionCommand) UserID() valueobject.UserID { return c.userID }
func (c RevokeSessionCommand) RefreshToken() string       { return c.refreshToken }
