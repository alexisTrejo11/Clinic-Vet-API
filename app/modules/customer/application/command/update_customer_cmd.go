// Package command contains the command definitions for customer-related operations.
package command

import (
	"time"

	c "clinic-vet-api/app/modules/core/domain/entity/customer"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
)

type UpdateCustomerCommand struct {
	id          valueobject.CustomerID
	photo       *string
	name        *valueobject.PersonName
	gender      *enum.PersonGender
	dateOfBirth *time.Time
}

func NewUpdateCustomerCommand(id uint, photo, firstName *string, lastName, gender *string, dateOfBirth *time.Time) (UpdateCustomerCommand, error) {
	cmd := &UpdateCustomerCommand{
		id:          valueobject.NewCustomerID(id),
		name:        valueobject.NewOptPersonName(firstName, lastName),
		photo:       photo,
		dateOfBirth: dateOfBirth,
		gender:      enum.NullableGender(gender),
	}

	if err := cmd.validate(); err != nil {
		return UpdateCustomerCommand{}, err
	}

	return *cmd, nil
}

func (cmd *UpdateCustomerCommand) UpdateEntity(existingCustomer c.Customer) c.Customer {
	builder := c.NewCustomerBuilder().
		WithID(cmd.id)

	if cmd.photo != nil {
		builder.WithPhoto(*cmd.photo)
	}

	if cmd.gender != nil {
		builder.WithGender(*cmd.gender)
	}

	if cmd.name != nil {
		builder.WithName(*cmd.name)
	}

	if cmd.dateOfBirth != nil {
		builder.WithDateOfBirth(*cmd.dateOfBirth)
	}

	return *builder.Build()
}

func (cmd *UpdateCustomerCommand) validate() error {
	if cmd.id.IsZero() {
		return UpdateCustomerCmdErr("id", "ID is required")
	}

	if cmd.name != nil {
		if valid := cmd.name.IsValid(); !valid {
			return UpdateCustomerCmdErr("name", "Name cannot be empty if provided")
		}
	}

	if cmd.photo == nil && *cmd.photo == "" {
		return UpdateCustomerCmdErr("photo", "Photo cannot be empty if provided")
	}

	if cmd.dateOfBirth != nil {
		if cmd.dateOfBirth.IsZero() {
			return UpdateCustomerCmdErr("date_of_birth", "Date of birth cannot be empty if provided")
		}

		if cmd.dateOfBirth.After(time.Now()) {
			return UpdateCustomerCmdErr("date_of_birth", "Date of birth cannot be in the future")
		}
	}

	return nil
}

func (cmd *UpdateCustomerCommand) ID() valueobject.CustomerID    { return cmd.id }
func (cmd *UpdateCustomerCommand) Photo() *string                { return cmd.photo }
func (cmd *UpdateCustomerCommand) Name() *valueobject.PersonName { return cmd.name }
func (cmd *UpdateCustomerCommand) DateOfBirth() *time.Time       { return cmd.dateOfBirth }
func (cmd *UpdateCustomerCommand) Gender() *enum.PersonGender    { return cmd.gender }

type DeactivateCustomerCommand struct {
	id valueobject.CustomerID
}

func NewDeactivateCustomerCommand(id uint) (DeactivateCustomerCommand, error) {
	cmd := &DeactivateCustomerCommand{id: valueobject.NewCustomerID(id)}

	if err := cmd.validate(); err != nil {
		return DeactivateCustomerCommand{}, err
	}

	return *cmd, nil
}

func (cmd *DeactivateCustomerCommand) validate() error {
	if cmd.id.IsZero() {
		return DeactivateCustomerCmdErr("id", "ID is required")
	}

	return nil
}

func (cmd *DeactivateCustomerCommand) ID() valueobject.CustomerID { return cmd.id }

func UpdateCustomerCmdErr(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "UpdateCustomerCommand")
}

func DeactivateCustomerCmdErr(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "DeactivateCustomerCommand")
}
