package event

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	commondto "clinic-vet-api/app/shared/dto"
)

type UserEventProducer interface {
	Registered(event UserRegisteredEvent)
	//TwoFactorEnabled(event TwoFactorEnabledEvent)
	//TwoFactorDisabled(event TwoFactorDisabledEvent)
	//TwoFactorVericationRequested(event TwoFactorVerificationRequestedEvent)
	//PasswordResetRequested(event PasswordResetRequestedEvent)
	//PasswordChanged(event PasswordChangedEvent)
}

type TwoFactorEnabledEvent struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type TwoFactorDisabledEvent struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type TwoFactorVerificationRequestedEvent struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Code   string `json:"code"`
}

type PasswordResetRequestedEvent struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

type PasswordChangedEvent struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type UserRegisteredEvent struct {
	UserID       valueobject.UserID      `json:"user_id"`
	Email        valueobject.Email       `json:"email"`
	Name         valueobject.PersonName  `json:"name"`
	Role         enum.UserRole           `json:"role"` // "customer" or "employee"
	PersonalData *commondto.PersonalData `json:"personal_data"`
}
