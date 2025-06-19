package dto

type PaymentResponse struct {
	ID            int32
	Appointment   AppointmentResponse
	Amount        int32
	PaymentMethod string
}
