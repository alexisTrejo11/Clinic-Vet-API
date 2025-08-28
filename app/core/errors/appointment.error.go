package domainerr

type AppointmentInvalidStatusTransition struct {
	BaseDomainError
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

func AppointmentStatusValidationErr(value, message string) error {
	return &NewValidationError("appointment status", value, message).BaseDomainError
}

func AppointmentInvalidStatusTransitionErr(from, to, message string) error {
	return &NewAppointmentInvalidStatusTransitionErr(from, to, message).BaseDomainError
}

func NewAppointmentInvalidStatusTransitionErr(from, to, message string) *AppointmentInvalidStatusTransition {
	return &AppointmentInvalidStatusTransition{
		BaseDomainError: BaseDomainError{
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
	return &NewValidationError("appointment schedule date", value, message).BaseDomainError
}
