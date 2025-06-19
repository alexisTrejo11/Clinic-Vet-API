package dto

type PaymentCreate struct {
	AppointmentID int32
	Amount        int32
	PaymentMethod string
}

type PaymentUpdate struct {
	ID            int32
	AppointmentID int32
	Amount        int32
	PaymentMethod string
}
