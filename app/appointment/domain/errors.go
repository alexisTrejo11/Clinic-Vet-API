package appointDomain

import domainErr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/domain"

type AppointmentInvalidStatusTransition struct {
	domainErr.BaseDomainError
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

func AppointmentStatusValidationErr(value, message string) error {
	return &domainErr.NewValidationError("appointment status", value, message).BaseDomainError
}

func AppointmentInvalidStatusTransitionErr(from, to, message string) error {
	return &NewAppointmentInvalidStatusTransitionErr(from, to, message).BaseDomainError
}

func NewAppointmentInvalidStatusTransitionErr(from, to, message string) *AppointmentInvalidStatusTransition {
	return &AppointmentInvalidStatusTransition{
		BaseDomainError: domainErr.BaseDomainError{
			Code:    "VALIDATION_ERROR",
			Type:    "domain",
			Message: message,
			Data: map[string]interface{}{
				"from": from,
				"to":   to,
			},
		},
		From:    from,
		To:      to,
		Message: message,
	}
}

func AppointmentScheduleDateValidationErr(value, message string) error {
	return &domainErr.NewValidationError("appointment schedule date", value, message).BaseDomainError
}
