package dto

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/command"
)

type RequestLogin struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required,min=8"`
}

func (r *RequestLogin) ToCommand() *command.LoginCommand {
	return &command.LoginCommand{
		Identifier: r.Identifier,
		Password:   r.Password,
	}
}

type RefreshSessionRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
