package service

import (
	appt "clinic-vet-api/app/modules/core/domain/entity/appointment"
	"clinic-vet-api/app/modules/core/domain/entity/employee"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	domainerr "clinic-vet-api/app/modules/core/error"
	"clinic-vet-api/app/modules/core/repository"
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	MAX_CUSTOMER_APPOINTMENTS_PER_DAY    = 1
	CLINIC_ROOMS_FOR_APPOINTMENTS        = 3
	MAX_MEDICAL_SESSION_DURATION_MINUTES = 60
)

type AppointmentCustomerSevice struct {
	appointmentRepo repository.AppointmentRepository
	customerRepo    repository.CustomerRepository
}

func NewAppointmentCustomerService(appointmentRepo repository.AppointmentRepository, customerRepo repository.CustomerRepository) *AppointmentCustomerSevice {
	return &AppointmentCustomerSevice{
		appointmentRepo: appointmentRepo,
		customerRepo:    customerRepo,
	}
}

func (aps *AppointmentCustomerSevice) ValidateRequest(ctx context.Context, customerID valueobject.CustomerID, appointment appt.Appointment) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := aps.CheckClinicAvailability(ctx, appointment); err != nil {
			errChan <- err
		}
	}()

	go func() {
		defer wg.Done()
		if err := aps.validateAppointmentPerDay(ctx, customerID, appointment); err != nil {
			errChan <- err
		}
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		return err
	}

	return nil
}

func (aps *AppointmentCustomerSevice) validateAppointmentPerDay(ctx context.Context, customerID valueobject.CustomerID, appointment appt.Appointment) error {
	requestDay := appointment.ScheduledDate().Day()
	requestMonth := appointment.ScheduledDate().Month()
	requestYear := appointment.ScheduledDate().Year()

	requestStartDay := time.Date(requestYear, requestMonth, requestDay, 0, 0, 0, 0, time.UTC)
	requestEndDay := time.Date(requestYear, requestMonth, requestDay, 23, 59, 59, 999, time.UTC)

	customerDateRangeSpec := specification.
		ApptByDateRange(requestStartDay, requestEndDay).
		And(specification.ApptByCustomer(customerID))

	appointmentPage, err := aps.appointmentRepo.Find(ctx, customerDateRangeSpec)
	if err != nil {
		return err
	}

	if len(appointmentPage.Items) >= MAX_CUSTOMER_APPOINTMENTS_PER_DAY {
		return ErrMaxAppointmentsPerDayReached(ctx)
	}

	return nil
}

func (aps *AppointmentCustomerSevice) CheckClinicAvailability(ctx context.Context, appointment appt.Appointment) error {

	requestStartDay := appointment.ScheduledDate().Truncate(24 * time.Hour)
	requestEndDay := requestStartDay.Add(24*time.Hour - time.Nanosecond)

	clinicDateRangeSpec := specification.ApptByDateRange(requestStartDay, requestEndDay)
	appointmentPage, err := aps.appointmentRepo.Find(ctx, clinicDateRangeSpec)
	if err != nil {
		return err
	}

	if len(appointmentPage.Items) >= CLINIC_ROOMS_FOR_APPOINTMENTS {
		return ErrClinicFullyBooked(ctx)
	}

	return nil
}

func ErrMaxAppointmentsPerDayReached(ctx context.Context) error {
	message := fmt.Sprintf("a customer can only have %d appointment(s) per day", MAX_CUSTOMER_APPOINTMENTS_PER_DAY)
	return domainerr.BusinessRuleError(ctx, message, "Appointment", "RequestedDate", "scheduled_date_validation")
}

func ErrClinicFullyBooked(ctx context.Context) error {
	message := fmt.Sprintf("the clinic is fully booked for the selected date, maximum of %d appointments per day", CLINIC_ROOMS_FOR_APPOINTMENTS)
	return domainerr.BusinessRuleError(ctx, message, "Appointment", "RequestedDate", "clinic_availability_validation")
}

func ValidateEmployeeCanAttendAppointment(ctx context.Context, employee employee.Employee, appointment appt.Appointment, appointmentRepo repository.AppointmentRepository) error {
	appointmentDate := appointment.ScheduledDate().Truncate(24 * time.Hour)
	startOfDay := appointmentDate
	endOfDay := appointmentDate.Add(24*time.Hour - time.Nanosecond)

	employeeDateRangeSpec := specification.
		ApptByDateRange(startOfDay, endOfDay).
		And(specification.ApptByEmployee(employee.ID()))

	appointmentPage, err := appointmentRepo.Find(ctx, employeeDateRangeSpec)
	if err != nil {
		return err
	}

	if len(appointmentPage.Items) == 0 {
		return nil
	}

	sessionMaxDuration := time.Duration(MAX_MEDICAL_SESSION_DURATION_MINUTES) * time.Minute
	appointmentStart := appointment.ScheduledDate()

	appointmentEnd := appointmentStart.Add(sessionMaxDuration)

	if employee.IsWithinWorkdayBreak(appointmentStart, appointmentEnd) {
		return ErrEmployeeNotAvailable(ctx, "employee is not available during the requested time")

	}

	for _, existingAppointment := range appointmentPage.Items {
		existingStart := existingAppointment.ScheduledDate()
		existingEnd := existingStart.Add(sessionMaxDuration)

		if hasTimeOverlap(appointmentStart, appointmentEnd, existingStart, existingEnd) {
			return ErrEmployeeNotAvailable(ctx, "employee is not available during the requested time")
		}

	}
	return nil

}

func ErrEmployeeNotAvailable(ctx context.Context, message string) error {
	return domainerr.BusinessRuleError(ctx, message, "Employee", "ScheduledDate", "employee_availability_validation")
}

func hasTimeOverlap(start1, end1, start2, end2 time.Time) bool {
	return start1.Before(end2) && end1.After(start2)
}
