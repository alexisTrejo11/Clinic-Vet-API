package command

import "clinic-vet-api/app/modules/core/domain/valueobject"

type RequestResetPasswordCommand struct {
	email valueobject.Email
}

func NewRequestResetPassword(email string) (RequestResetPasswordCommand, error) {
	emailVO, err := valueobject.NewEmail(email)
	if err != nil {
		return RequestResetPasswordCommand{}, err
	}

	return RequestResetPasswordCommand{
		email: emailVO,
	}, nil
}

func (cmd *RequestResetPasswordCommand) Email() valueobject.Email { return cmd.email }
