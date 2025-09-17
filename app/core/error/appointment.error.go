package domainerr

import (
	"errors"
)

func AppointmentScheduleDateZeroErr() error {
	err := errors.New("owner ID must be greater than zero")
	return NewBusinessRuleError(err.Error(), "Appointment", "schedule_date")
}

func AppointmentScheduleDateRuleErr(ruleViolated string) error {
	return NewBusinessRuleError(ruleViolated, "Appointment", "schedule_date")
}

func AppointmentStatusTransitionErr(fromStatus, toStatus string, message string) error {
	return &BaseDomainError{
		Code:    "BUSINESS_RULE_VIOLATION",
		Type:    "domain",
		Message: message,
		Details: map[string]string{
			"from_status": fromStatus,
			"to_status":   toStatus,
			"entity":      "appointment",
			"field":       "schedule_date",
		},
	}
}
