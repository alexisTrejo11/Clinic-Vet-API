package dto

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/command"
)

type RequestLogout struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (r *RequestLogout) ToCommand(userIdInt uint) (command.LogoutCommand, error) {
	cmd := &command.LogoutCommand{
		RefreshToken: r.RefreshToken,
		UserID:       valueobject.NewUserID(userIdInt),
		CTX:          context.Background(),
	}

	return *cmd, nil
}
