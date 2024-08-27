package DTOs

type PaymentInsertDTO struct {
	AppointmentID int32
	Amount        int32
	PaymentMethod string
}

type PaymentDTO struct {
	ID             int32
	AppointmentDTO AppointmentDTO
	Amount         int32
	PaymentMethod  string
}

type PaymentUpdateDTO struct {
	ID            int32
	AppointmentID int32
	Amount        int32
	PaymentMethod string
}
