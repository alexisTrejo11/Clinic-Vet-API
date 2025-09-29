package dto

import (
	sessionCmd "clinic-vet-api/app/modules/auth/application/command"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type RequestLogout struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (r *RequestLogout) ToCommand(userIDUint uint) (sessionCmd.RevokeSessionCommand, error) {
	cmd := sessionCmd.NewRevokeSessionCommand(valueobject.NewUserID(userIDUint), r.RefreshToken)
	return cmd, nil
}
