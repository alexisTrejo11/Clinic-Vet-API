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
	ID          valueobject.CustomerID
	Photo       *string
	FirstName   *string
	LastName    *string
	Address     *string
	Notes       *string
	Gender      *enum.PersonGender
	DateOfBirth *time.Time
	PhoneNumber *string
}

func (cmd *UpdateCustomerCommand) UpdateEntity(existingCustomer *customer.Customer) error {
	return nil
}

type DeactivateCustomerCommand struct {
	ID     valueobject.CustomerID
	Reason string
}
