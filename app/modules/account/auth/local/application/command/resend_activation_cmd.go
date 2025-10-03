package command

import "clinic-vet-api/app/modules/core/domain/valueobject"

type ResendActivationCmd struct {
	email valueobject.Email
}

func NewResendActivationCmd(email valueobject.Email) ResendActivationCmd {
	return ResendActivationCmd{email: email}
}

func (c *ResendActivationCmd) Email() valueobject.Email {
	return c.email
}
