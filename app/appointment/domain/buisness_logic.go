package appointDomain

import (
	"errors"
	"time"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

const (
	MIN_DAYS_TO_SCHEDULE = 3
	MAX_DAYS_TO_SCHEDULE = 30
	CLINIC_OPENING_HOUR  = 8
	CLINIC_CLOSING_HOUR  = 20
)

func (a *Appointment) ValidateFields() error {
	if a.id.Equals(shared.NilIntegerId()) {
		return errors.New("appointment ID cannot be nil")
	}

	if a.petId.Equals(petDomain.PetId{}) {
		return errors.New("pet ID cannot be nil")
	}

	if a.ownerId <= 0 {
		return errors.New("owner ID must be greater than zero")
	}

	if a.scheduledDate.IsZero() {
		return errors.New("scheduled date cannot be zero")
	}

	if err := a.ValidateRequestSchedule(); err != nil {
		return err
	}
	return nil
}

func (a *Appointment) RescheduleAppointment(newDate time.Time) error {
	if a.GetStatus() == StatusCompleted {
		return AppointmentStatusValidationErr("completed", "cannot reschedule completed appointment")
	}

	if a.GetStatus() == StatusCancelled {
		return AppointmentStatusValidationErr("cancelled", "cannot reschedule cancelled appointment")
	}

	if newDate.Before(time.Now()) {
		return AppointmentScheduleDateValidationErr(newDate.String(), "appointment date must be in the future")
	}

	a.scheduledDate = newDate
	a.status = StatusRescheduled
	a.updatedAt = time.Now()

	return nil
}

func (a *Appointment) Cancel() error {
	if err := a.validateCanBeCancel(); err != nil {
		return err
	}

	a.status = StatusCancelled
	a.updatedAt = time.Now()
	return nil
}

func (a *Appointment) Complete() error {
	if err := a.validateCanBeCompleted(); err != nil {
		return err
	}

	a.status = StatusCompleted
	a.updatedAt = time.Now()

	return nil
}

func (a *Appointment) NotPresented() error {
	if a.status == StatusNotPresented {
		return AppointmentStatusValidationErr(string(a.status), "appointment is already marked as not presented")
	}

	if a.GetStatus() == StatusCancelled {
		return AppointmentStatusValidationErr(string(a.status), "cannot mark cancelled appointment as not presented")
	}

	if a.GetStatus() == StatusNotPresented {
		return AppointmentStatusValidationErr(string(a.status), "appointment is already marked as not presented")
	}

	a.status = StatusNotPresented
	a.updatedAt = time.Now()
	return nil
}

func (a *Appointment) ValidateRequestSchedule() error {
	now := time.Now()

	if a.scheduledDate.IsZero() {
		return AppointmentScheduleDateValidationErr(a.scheduledDate.String(), "scheduled date cannot be zero")
	}

	if a.scheduledDate.Before(now) {
		return AppointmentScheduleDateValidationErr(a.scheduledDate.String(), "scheduled date cannot be in the past")
	}

	if a.scheduledDate.Before(now.AddDate(0, 0, MIN_DAYS_TO_SCHEDULE)) {
		return AppointmentScheduleDateValidationErr(a.scheduledDate.String(), "appointments must be scheduled at least 3 days in advance")
	}

	if a.scheduledDate.Weekday() == time.Saturday || a.scheduledDate.Weekday() == time.Sunday {
		return AppointmentScheduleDateValidationErr(a.scheduledDate.String(), "appointments cannot be scheduled on weekends")
	}

	return nil
}

func (a *Appointment) Confirm(vetId *vetDomain.VetId) error {
	if a.GetStatus() != StatusPending {
		return AppointmentStatusValidationErr(string(a.GetStatus()), "only pending appointments can be confirmed")
	}

	a.vetId = vetId
	a.SetStatus(StatusPending)
	a.updatedAt = time.Now()
	return nil
}

func (a *Appointment) validateCanBeCancel() error {
	if a.scheduledDate.Before(time.Now()) {
		return errors.New("cannot cancel an appointment that is already in the past")
	}

	if a.status == StatusCompleted || a.status == StatusNotPresented {
		return AppointmentInvalidStatusTransitionErr(
			string(a.status),
			string(StatusCancelled),
			"cannot cancel a completed or not presented appointment",
		)
	}

	if a.status == StatusCancelled {
		return AppointmentStatusValidationErr(
			string(a.status),
			"appointment is already cancelled",
		)
	}

	return nil
}

func (a *Appointment) validateCanBeCompleted() error {
	if a.GetStatus() == StatusCompleted {
		return AppointmentStatusValidationErr(string(a.GetStatus()), "appointment is already completed")
	}

	if a.GetStatus() == StatusCancelled {
		return AppointmentStatusValidationErr(string(a.GetStatus()), "cannot complete cancelled appointment")
	}

	return nil
}
