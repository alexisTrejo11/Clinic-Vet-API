// Package dto contains data transfer objects for authentication operations
package dto

import (
	"clinic-vet-api/app/modules/auth/application/command"
)

type RequestLogin struct {
	Identifier string `json:"identifier" binding:"required"` // Can be email or phoneNumber
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
