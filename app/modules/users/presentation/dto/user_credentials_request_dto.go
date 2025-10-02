package dto

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/users/application/usecase/command"
)

type RequestEmailRequest struct {
	Email string `json:"email" validate:"required,email,max=255"`
}

type ResetPasswordRequest struct {
	RequestEmailRequest
	Token       string `json:"token" validate:"required,min=6,max=6"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=64"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required,min=8,max=64"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=64"`
}

func (r *UpdatePasswordRequest) ToCommand(userID uint) (command.UpdatePasswordCommand, error) {
	return command.UpdatePasswordCommand{
		UserID:          valueobject.NewUserID(userID),
		CurrentPassword: r.OldPassword,
		NewPassword:     r.NewPassword,
	}, nil
}

func (r *ResetPasswordRequest) ToCommand() (command.ResetPasswordCommand, error) {
	return command.NewResetPasswordCommand(r.Email, r.Token, r.NewPassword)
}

func (r *RequestEmailRequest) ToCommand() (command.RequestResetPasswordCommand, error) {
	return command.NewRequestResetPassword(r.Email)
}
