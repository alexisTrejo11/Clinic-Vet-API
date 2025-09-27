package appointment

import (
	"context"
	"fmt"

	"clinic-vet-api/app/modules/core/domain/enum"
	domainerr "clinic-vet-api/app/modules/core/error"
)

// AppointmentErrorCode representa códigos de error específicos para citas
type AppointmentErrorCode string

const (
	AppointmentNotesTooLong           AppointmentErrorCode = "APPOINTMENT_NOTES_TOO_LONG"
	AppointmentInvalidService         AppointmentErrorCode = "APPOINTMENT_INVALID_SERVICE"
	AppointmentInvalidReason          AppointmentErrorCode = "APPOINTMENT_INVALID_REASON"
	AppointmentCannotReschedule       AppointmentErrorCode = "APPOINTMENT_CANNOT_RESCHEDULE"
	AppointmentCannotCancel           AppointmentErrorCode = "APPOINTMENT_CANNOT_CANCEL"
	AppointmentCannotComplete         AppointmentErrorCode = "APPOINTMENT_CANNOT_COMPLETE"
	AppointmentCannotMarkNotPresented AppointmentErrorCode = "APPOINTMENT_CANNOT_MARK_NOT_PRESENTED"
	AppointmentCannotConfirm          AppointmentErrorCode = "APPOINTMENT_CANNOT_CONFIRM"
	AppointmentInvalidTransition      AppointmentErrorCode = "APPOINTMENT_INVALID_TRANSITION"
	AppointmentCannotDelete           AppointmentErrorCode = "APPOINTMENT_CANNOT_DELETE"
	AppointmentScheduledDateInvalid   AppointmentErrorCode = "APPOINTMENT_SCHEDULED_DATE_INVALID"
)

func appointmentValidationError(ctx context.Context, code AppointmentErrorCode, field, message, operation string) error {
	return domainerr.ValidationError(ctx, string(code), "appointment", field,
		fmt.Sprintf("Appointment %s: %s", field, message), operation)
}

func appointmentBusinessError(ctx context.Context, code AppointmentErrorCode, rule, operation string) error {
	return domainerr.BusinessRuleError(ctx, rule, "appointment", "", operation)
}

func NotesTooLongError(ctx context.Context, operation string) error {
	return appointmentValidationError(ctx, AppointmentNotesTooLong, "notes",
		"notes cannot exceed 1000 characters", operation)
}

func InvalidServiceError(ctx context.Context, service enum.ClinicService, operation string) error {
	return appointmentValidationError(ctx, AppointmentInvalidService, "service",
		fmt.Sprintf("invalid service: %s", service), operation)
}

func InvalidReasonError(ctx context.Context, reason string, operation string) error {
	return appointmentValidationError(ctx, AppointmentInvalidReason, "reason",
		fmt.Sprintf("invalid reason: %s", reason), operation)
}

func CannotRescheduleError(ctx context.Context, currentStatus enum.AppointmentStatus, operation string) error {
	return appointmentBusinessError(ctx, AppointmentCannotReschedule, "reschedule", operation)
}

func CannotCancelError(ctx context.Context, currentStatus enum.AppointmentStatus, operation string) error {
	return appointmentBusinessError(ctx, AppointmentCannotCancel, "cancel", operation)
}

func CannotCompleteError(ctx context.Context, currentStatus enum.AppointmentStatus, operation string) error {
	return appointmentBusinessError(ctx, AppointmentCannotComplete, "complete", operation)
}

func CannotMarkNotPresentedError(ctx context.Context, currentStatus enum.AppointmentStatus, operation string) error {
	return appointmentBusinessError(ctx, AppointmentCannotMarkNotPresented, "mark_not_presented", operation)
}

func CannotConfirmError(ctx context.Context, currentStatus enum.AppointmentStatus, operation string) error {
	return appointmentBusinessError(ctx, AppointmentCannotConfirm, "confirm", operation)
}

func InvalidTransitionError(ctx context.Context, from, to enum.AppointmentStatus, operation string) error {
	return appointmentBusinessError(ctx, AppointmentInvalidTransition, "status_transition", operation)
}

func CannotDeleteError(ctx context.Context, currentStatus enum.AppointmentStatus, operation string) error {
	return appointmentBusinessError(ctx, AppointmentCannotDelete, "delete", operation)
}

func ScheduledDateInvalidError(ctx context.Context, message, operation string) error {
	return appointmentValidationError(ctx, AppointmentScheduledDateInvalid, "scheduled_date",
		message, operation)
}
