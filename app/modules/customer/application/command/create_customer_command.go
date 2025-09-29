package command

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/customer"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
)

type CreateCustomerCommand struct {
	photo       string
	name        valueobject.PersonName
	gender      enum.PersonGender
	dateOfBirth time.Time
}

func NewCreateCustomerCommand(photo string, firstName string, lastName string, gender string, dateOfBirth time.Time) (CreateCustomerCommand, error) {
	cmd := &CreateCustomerCommand{
		photo:       photo,
		name:        valueobject.NewPersonNameNoErr(firstName, lastName),
		gender:      enum.NewPersonGender(gender),
		dateOfBirth: dateOfBirth,
	}

	if err := cmd.Validate(); err != nil {
		return CreateCustomerCommand{}, err
	}

	return *cmd, nil
}

func (c *CreateCustomerCommand) Validate() error {
	if c.dateOfBirth.After(time.Now()) {
		return createCustomerCmdErr("date_of_birth", "cannot be in the future")
	}
	if !c.name.IsValid() {
		return createCustomerCmdErr("name", "name cannot be empty")
	}

	if c.photo != "" {
		return createCustomerCmdErr("photo", "photo URL is not valid")
	}

	if !c.gender.IsValid() {
		return createCustomerCmdErr("gender", "gender values is not valid")
	}

	return nil
}

func (c *CreateCustomerCommand) ToEntity() customer.Customer {
	customer := customer.NewCustomerBuilder().
		WithName(c.name).
		WithPhoto(c.photo).
		WithGender(c.gender).
		WithDateOfBirth(c.dateOfBirth).
		WithIsActive(true).
		Build()

	return *customer

}

func (c *CreateCustomerCommand) Photo() string                { return c.photo }
func (c *CreateCustomerCommand) Name() valueobject.PersonName { return c.name }
func (c *CreateCustomerCommand) Gender() enum.PersonGender    { return c.gender }
func (c *CreateCustomerCommand) DateOfBirth() time.Time       { return c.dateOfBirth }

func createCustomerCmdErr(field string, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "CreateCustomerCommand")
}
