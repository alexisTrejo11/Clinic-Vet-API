package service

import (
	c "clinic-vet-api/app/modules/core/domain/entity/customer"
	"clinic-vet-api/app/modules/core/repository"
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

func (aps *AppointmentCustomerSevice) ValidateCustomerAppointmentRequest(customer c.Customer) error {
	return nil
}
