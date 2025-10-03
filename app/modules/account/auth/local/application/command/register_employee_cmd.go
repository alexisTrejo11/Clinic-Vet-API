package command

import (
	"errors"

	u "clinic-vet-api/app/modules/core/domain/entity/user"
	"clinic-vet-api/app/modules/core/domain/enum"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
)

type RegisterEmployeeCommand struct {
	email       vo.Email
	password    string
	role        enum.UserRole
	phoneNumber *vo.PhoneNumber
	employeeID  vo.EmployeeID
}

func NewRegisterEmployeeCommand(
	email vo.Email, password string, phoneNumber *vo.PhoneNumber, employeeID vo.EmployeeID, role string,
) RegisterEmployeeCommand {
	return RegisterEmployeeCommand{
		email:       email,
		password:    password,
		role:        enum.UserRole(role),
		phoneNumber: phoneNumber,
		employeeID:  employeeID,
	}
}

func (cmd *RegisterEmployeeCommand) ToEntity() u.User {
	return *u.NewUserBuilder().
		WithRole(cmd.role).
		WithEmail(cmd.email).
		WithPassword(cmd.password).
		WithPhoneNumber(cmd.phoneNumber).
		WithEmployeeID(&cmd.employeeID).
		Build()

}

func (cmd *RegisterEmployeeCommand) Validate() error {
	if !cmd.role.IsValid() {
		return errors.New("invalid role")
	}
	if !cmd.role.IsEmployee() {
		return errors.New("role must be an employee role")
	}

	if cmd.employeeID.IsZero() {
		return errors.New("employee ID cannot be empty")
	}

	if !cmd.email.IsValid() {
		return errors.New("invalid email format")
	}

	if cmd.password == "" {
		return errors.New("password cannot be empty")
	}

	return nil
}

func (cmd *RegisterEmployeeCommand) Email() vo.Email              { return cmd.email }
func (cmd *RegisterEmployeeCommand) Password() string             { return cmd.password }
func (cmd *RegisterEmployeeCommand) PhoneNumber() *vo.PhoneNumber { return cmd.phoneNumber }
func (cmd *RegisterEmployeeCommand) EmployeeID() vo.EmployeeID    { return cmd.employeeID }
func (cmd *RegisterEmployeeCommand) Role() enum.UserRole          { return cmd.role }
