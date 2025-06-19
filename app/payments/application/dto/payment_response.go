package dto

type PaymentResponse struct {
	ID            int32
	Appointment   int // DTO
	Amount        int32
	PaymentMethod string
}
