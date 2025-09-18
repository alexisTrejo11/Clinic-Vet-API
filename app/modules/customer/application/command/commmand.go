// Package command contains the command definitions for customer-related operations.
package command

import (
	"clinic-vet-api/app/core/domain/entity/customer"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	"context"
	"time"
)

type CreateCustomerCommand struct {
	Photo       string
	FirstName   string
	LastName    string
	Gender      enum.PersonGender
	DateOfBirth time.Time
	CTX         context.Context
}

func (c *CreateCustomerCommand) ToEntity() (*customer.Customer, error) {
	fullName := valueobject.PersonName{
		FirstName: c.FirstName,
		LastName:  c.LastName,
	}

	return customer.CreateCustomer(
		c.CTX,
		customer.WithFullName(fullName),
		customer.WithPhoto(c.Photo),
		customer.WithGender(c.Gender),
		customer.WithDateOfBirth(c.DateOfBirth),
		customer.WithIsActive(true),
	)
}

type UpdateCustomerCommand struct {
	ID          valueobject.CustomerID `json:"id" validate:"required,uuid"`
	Photo       *string                `json:"photo"`
	FirstName   *string                `json:"first_name"`
	LastName    *string                `json:"last_name"`
	Address     *string                `json:"address"`
	Notes       *string                `json:"notes"`
	Gender      *enum.PersonGender     `json:"gender" validate:"omitempty,oneof=male female not_specified"`
	DateOfBirth *time.Time             `json:"date_of_birth" validate:"required"`
	PhoneNumber *string                `json:"phone_number"`
	CTX         context.Context
}

func (c *UpdateCustomerCommand) UpdateEntity(existingCustomer *customer.Customer) error {
	return nil
}

type DeactivateCustomerCommand struct {
	ID     valueobject.CustomerID
	Reason string
	CTX    context.Context
}
