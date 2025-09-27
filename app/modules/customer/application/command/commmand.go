// Package command contains the command definitions for customer-related operations.
package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/customer"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type CreateCustomerCommand struct {
	Photo       string
	FirstName   string
	LastName    string
	Gender      enum.PersonGender
	DateOfBirth time.Time
}

func (c *CreateCustomerCommand) ToEntity() customer.Customer {
	customer := customer.NewCustomerBuilder().
		WithName(valueobject.NewPersonNameNoErr(c.FirstName, c.LastName)).
		WithPhoto(c.Photo).
		WithGender(c.Gender).
		WithDateOfBirth(c.DateOfBirth).
		WithIsActive(true).
		Build()

	return *customer
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
}

func (cmd *UpdateCustomerCommand) UpdateEntity(existingCustomer *customer.Customer) error {
	return nil
}

type DeactivateCustomerCommand struct {
	ID     valueobject.CustomerID
	Reason string
}
