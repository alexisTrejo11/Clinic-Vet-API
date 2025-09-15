package dto

import (
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/modules/auth/application/command"
	"context"
)

type RequestLogout struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (r *RequestLogout) ToCommand(ctx context.Context, userIDUint uint) (command.LogoutCommand, error) {
	cmd := &command.LogoutCommand{
		RefreshToken: r.RefreshToken,
		UserID:       valueobject.NewUserID(userIDUint),
		CTX:          ctx,
	}

	return *cmd, nil
}
