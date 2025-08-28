package service

import (
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/errors"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

const (
	MIN_DAYS_TO_SCHEDULE = 3
	MAX_DAYS_TO_SCHEDULE = 30
)

func ValidateFields(a *entity.Appointment) error {
	if a.GetId().Equals(shared.NilIntegerId()) {
		return errors.New("appointment ID cannot be nil")
	}

	if a.GetPetId().Equals(valueobject.PetID{}) {
		return errors.New("pet ID cannot be nil")
	}

	if a.GetOwnerId() <= 0 {
		return errors.New("owner ID must be greater than zero")
	}

	if a.GetScheduledDate().IsZero() {
		return errors.New("scheduled date cannot be zero")
	}

	if err := ValidateRequestSchedule(a); err != nil {
		return err
	}
	return nil
}

func RescheduleAppointment(a *entity.Appointment, newDate time.Time) error {
	if a.GetStatus() == enum.StatusCompleted {
		return domainerr.AppointmentStatusValidationErr("completed", "cannot reschedule completed appointment")
	}

	if a.GetStatus() == enum.StatusCancelled {
		return domainerr.AppointmentStatusValidationErr("cancelled", "cannot reschedule cancelled appointment")
	}

	if newDate.Before(time.Now()) {
		return domainerr.AppointmentScheduleDateValidationErr(newDate.String(), "appointment date must be in the future")
	}

	a.SetScheduledDate(newDate)
	a.SetStatus(enum.StatusRescheduled)
	a.SetUpdatedAt(time.Now())

	return nil
}

func Cancel(a *entity.Appointment) error {
	if err := validateCanBeCancel(a); err != nil {
		return err
	}

	a.SetStatus(enum.StatusCancelled)
	a.SetUpdatedAt(time.Now())
	return nil
}

func Complete(a *entity.Appointment) error {
	if err := validateCanBeCompleted(a); err != nil {
		return err
	}

	a.SetStatus(enum.StatusCompleted)
	a.SetUpdatedAt(time.Now())

	return nil
}

func NotPresented(a *entity.Appointment) error {
	if a.GetStatus() == enum.StatusNotPresented {
		return domainerr.AppointmentStatusValidationErr(string(a.GetStatus()), "appointment is already marked as not presented")
	}

	if a.GetStatus() == enum.StatusCancelled {
		return domainerr.AppointmentStatusValidationErr(string(a.GetStatus()), "cannot mark cancelled appointment as not presented")
	}

	if a.GetStatus() == enum.StatusNotPresented {
		return domainerr.AppointmentStatusValidationErr(string(a.GetStatus()), "appointment is already marked as not presented")
	}

	a.SetStatus(enum.StatusNotPresented)
	a.SetUpdatedAt(time.Now())
	return nil
}

func ValidateRequestSchedule(a *entity.Appointment) error {
	now := time.Now()

	if a.GetScheduledDate().IsZero() {
		return domainerr.AppointmentScheduleDateValidationErr(a.GetScheduledDate().String(), "scheduled date cannot be zero")
	}

	if a.GetScheduledDate().Before(now) {
		return domainerr.AppointmentScheduleDateValidationErr(a.GetScheduledDate().String(), "scheduled date cannot be in the past")
	}

	if a.GetScheduledDate().Before(now.AddDate(0, 0, MIN_DAYS_TO_SCHEDULE)) {
		return domainerr.AppointmentScheduleDateValidationErr(a.GetScheduledDate().String(), "appointments must be scheduled at least 3 days in advance")
	}

	if a.GetScheduledDate().Weekday() == time.Saturday || a.GetScheduledDate().Weekday() == time.Sunday {
		return domainerr.AppointmentScheduleDateValidationErr(a.GetScheduledDate().String(), "appointments cannot be scheduled on weekends")
	}

	return nil
}

func Confirm(vetID *valueobject.VetID, a *entity.Appointment) error {
	if a.GetStatus() != enum.StatusPending {
		return domainerr.AppointmentStatusValidationErr(string(a.GetStatus()), "only pending appointments can be confirmed")
	}

	a.SetVetId(vetID)
	a.SetStatus(enum.StatusConfirmed)
	a.SetUpdatedAt(time.Now())
	return nil
}

func validateCanBeCancel(a *entity.Appointment) error {
	if a.GetScheduledDate().Before(time.Now()) {
		return errors.New("cannot cancel an appointment that is already in the past")
	}

	if a.GetStatus() == enum.StatusCompleted || a.GetStatus() == enum.StatusNotPresented {
		return domainerr.AppointmentInvalidStatusTransitionErr(
			string(a.GetStatus()),
			string(enum.StatusCancelled),
			"cannot cancel a completed or not presented appointment",
		)
	}

	if a.GetStatus() == enum.StatusCancelled {
		return domainerr.AppointmentStatusValidationErr(
			string(a.GetStatus()),
			"appointment is already cancelled",
		)
	}

	return nil
}

func validateCanBeCompleted(a *entity.Appointment) error {
	if a.GetStatus() == enum.StatusCompleted {
		return domainerr.AppointmentStatusValidationErr(string(a.GetStatus()), "appointment is already completed")
	}

	if a.GetStatus() == enum.StatusCancelled {
		return domainerr.AppointmentStatusValidationErr(string(a.GetStatus()), "cannot complete cancelled appointment")
	}

	return nil
}
