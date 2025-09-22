package dto

import (
	sessionCmd "clinic-vet-api/app/modules/auth/application/command/session"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type RequestLogout struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (r *RequestLogout) ToCommand(userIDUint uint) (sessionCmd.RevokeUserSessionCommand, error) {
	cmd := &sessionCmd.RevokeUserSessionCommand{
		RefreshToken: r.RefreshToken,
		UserID:       valueobject.NewUserID(userIDUint),
	}

	return *cmd, nil
}
