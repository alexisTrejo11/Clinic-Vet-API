package agregates

import "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"

type AppointmentAggregate struct {
	appointment  *appointment.Appointment
	payments     []*payment.Payment
	reminders    []*notification.Reminder
	medicalNotes []*medical.Note
}

func (a *AppointmentAggregate) AddPayment(payment *payment.Payment) error {
	// Business logic for adding payment to appointment
}

func (a *AppointmentAggregate) TotalPaid() valueobject.Money {
	// Calculate total payments
}

func (a *AppointmentAggregate) AddReminder(reminder *notification.Reminder) error {
	// Business logic for reminders
}
