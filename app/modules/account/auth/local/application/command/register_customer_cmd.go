package command

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/user"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
)

type RegisterCustomerCommand struct {
	email       valueobject.Email
	phoneNumber *valueobject.PhoneNumber
	password    string
	role        enum.UserRole

	name        valueobject.PersonName
	gender      enum.PersonGender
	dateOfBirth time.Time
}

func (cmd *RegisterCustomerCommand) ToEntity() *user.User {
	return user.NewUserBuilder().
		WithRole(cmd.role).
		WithEmail(cmd.email).
		WithPhoneNumber(cmd.phoneNumber).
		WithPassword(cmd.password).
		Build()
}

func NewRegisterCustomerCommand(
	email string,
	phoneNumber *string,
	password string,
	firstName string,
	lastName string,
	gender string,
	dateOfBirth time.Time,
) (RegisterCustomerCommand, error) {
	cmd := &RegisterCustomerCommand{
		email:       valueobject.NewEmailNoErr(email),
		phoneNumber: valueobject.NewOptPhoneNumber(phoneNumber),
		password:    password,
		role:        enum.UserRoleCustomer,
		name:        valueobject.NewPersonNameNoErr(firstName, lastName),
		gender:      enum.PersonGender(gender),
		dateOfBirth: dateOfBirth,
	}

	if err := cmd.validate(); err != nil {
		return RegisterCustomerCommand{}, err
	}

	return *cmd, nil
}

func (cmd *RegisterCustomerCommand) Email() valueobject.Email              { return cmd.email }
func (cmd *RegisterCustomerCommand) PhoneNumber() *valueobject.PhoneNumber { return cmd.phoneNumber }
func (cmd *RegisterCustomerCommand) Password() string                      { return cmd.password }
func (cmd *RegisterCustomerCommand) Role() enum.UserRole                   { return cmd.role }
func (cmd *RegisterCustomerCommand) Name() valueobject.PersonName          { return cmd.name }
func (cmd *RegisterCustomerCommand) DateOfBirth() time.Time                { return cmd.dateOfBirth }
func (cmd *RegisterCustomerCommand) Gender() enum.PersonGender             { return cmd.gender }

func (cmd *RegisterCustomerCommand) validate() error {
	if !cmd.email.IsValid() {
		return CustomerRegisterCmdErr("email", "invalid format")
	}
	if err := cmd.phoneNumber.Validate(); err != nil {
		return CustomerRegisterCmdErr("phone number", err.Error())
	}
	if cmd.password == "" {
		return CustomerRegisterCmdErr("password", "is required")
	}
	if !cmd.name.IsValid() {
		return CustomerRegisterCmdErr("name", "invalid format")
	}
	if cmd.dateOfBirth.IsZero() {
		return CustomerRegisterCmdErr("date of birth", "cannot have hour, minute, second or nanosecond component")
	}
	if !cmd.gender.IsValid() {
		return CustomerRegisterCmdErr("gender", "invalid format")
	}
	return nil
}

func CustomerRegisterCmdErr(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "RegisterCustomerCommand")
}
