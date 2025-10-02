// Package dto contains data transfer objects for authentication operations
package dto

import (
	authCmd "clinic-vet-api/app/modules/account/auth/local/application/command"
	"clinic-vet-api/app/shared/auth"
	ginutils "clinic-vet-api/app/shared/gin_utils"
)

type RequestLogin struct {
	Identifier string `json:"identifier" binding:"required"` // Can be email or phoneNumber
	Password   string `json:"password" binding:"required,min=8"`
}

func (r *RequestLogin) ToCommand(metadata *ginutils.LoginMetadata) (authCmd.LoginCommand, error) {
	loginMetadata := auth.NewLoginMetadata(metadata.IP, metadata.UserAgent, metadata.DeviceInfo)
	return authCmd.NewLoginCommand(r.Identifier, r.Password, loginMetadata)
}
