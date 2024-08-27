package services

import "example.com/at/backend/api-vet/DTOs"

type PaymentService interface {
	CreatePayment(paymentInsertDTO DTOs.PaymentInsertDTO) error
	GetPaymentID(paymentID int32) (*DTOs.PaymentDTO, error)
	GetPaymentByAppointmentID(appointmentID int32) (*DTOs.AppointmentDTO, error)
	GetSuccessfullPaymentsByOwnerID(ownerID int32) ([]DTOs.PaymentDTO, error)
	GetSuccessPaymentsToBePaidByOwnerID(ownerID int32) ([]DTOs.PaymentDTO, error)
	UpdatePayment(paymentUpdateDTO DTOs.PaymentUpdateDTO) error
	DeletePayment(paymentID int32) error
}
