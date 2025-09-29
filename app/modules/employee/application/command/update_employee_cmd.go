package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/employee"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
)

type UpdateEmployeeCommand struct {
	employeeID      valueobject.EmployeeID
	name            *valueobject.PersonName
	photo           *string
	licenseNumber   *string
	yearsExperience *int32
	specialty       *enum.VetSpecialty
	isActive        *bool
	laboralSchedule *[]ScheduleData
}

func NewUpdateEmployeeCommand(
	id uint,
	firstName, lastName,
	photo,
	licenseNumber,
	specialty *string,
	yearsExperience *int32,
	isActive *bool,
	laboralSchedule *[]ScheduleData,
) (UpdateEmployeeCommand, error) {
	cmd := &UpdateEmployeeCommand{
		employeeID:      valueobject.NewEmployeeID(id),
		name:            valueobject.NewOptPersonName(firstName, lastName),
		photo:           photo,
		specialty:       enum.OptVetSpecialty(specialty),
		licenseNumber:   licenseNumber,
		yearsExperience: yearsExperience,
		isActive:        isActive,
		laboralSchedule: laboralSchedule,
	}

	if err := cmd.validate(); err != nil {
		return UpdateEmployeeCommand{}, err
	}

	return *cmd, nil
}

func (cmd *UpdateEmployeeCommand) validate() error {
	if cmd.name != nil {
		if !cmd.name.IsValid() {
			return updateEmployeeCmdErr("name", "cannot be empty if provided")
		}
	}

	if cmd.licenseNumber != nil {
		if len(*cmd.licenseNumber) == 0 {
			return updateEmployeeCmdErr("licenseNumber", "cannot be empty if provided")
		}
	}

	if cmd.yearsExperience != nil {
		if *cmd.yearsExperience < 0 {
			return updateEmployeeCmdErr("yearsExperience", "cannot be negative if provided")
		}
	}

	if cmd.specialty != nil {
		if !cmd.specialty.IsValid() {
			return updateEmployeeCmdErr("specialty", "is not a valid specialty if provided")
		}
	}

	return nil
}

func (cmd *UpdateEmployeeCommand) EmployeeID() valueobject.EmployeeID { return cmd.employeeID }
func (cmd *UpdateEmployeeCommand) Name() *valueobject.PersonName      { return cmd.name }
func (cmd *UpdateEmployeeCommand) Photo() *string                     { return cmd.photo }
func (cmd *UpdateEmployeeCommand) LicenseNumber() *string             { return cmd.licenseNumber }
func (cmd *UpdateEmployeeCommand) YearsExperience() *int32            { return cmd.yearsExperience }
func (cmd *UpdateEmployeeCommand) Specialty() *enum.VetSpecialty      { return cmd.specialty }
func (cmd *UpdateEmployeeCommand) IsActive() *bool                    { return cmd.isActive }
func (cmd *UpdateEmployeeCommand) LaboralSchedule() *[]ScheduleData   { return cmd.laboralSchedule }

func updateEmployeeCmdErr(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "UpdateEmployeeCommand")
}

func (cmd *UpdateEmployeeCommand) ToUpdateEmployee(emp employee.Employee) employee.Employee {
	employeeBuilder := employee.NewEmployeeBuilder()

	if cmd.name != nil {
		employeeBuilder.WithName(*cmd.name)
	}
	if cmd.photo != nil {
		employeeBuilder.WithPhoto(*cmd.photo)
	}
	if cmd.licenseNumber != nil {
		employeeBuilder.WithLicenseNumber(*cmd.licenseNumber)
	}
	if cmd.yearsExperience != nil {
		employeeBuilder.WithYearsExperience(*cmd.yearsExperience)
	}
	if cmd.specialty != nil {
		employeeBuilder.WithSpecialty(*cmd.specialty)
	}
	if cmd.isActive != nil {
		employeeBuilder.WithIsActive(*cmd.isActive)
	}
	//if cmd.LaboralSchedule != nil {
	//	employeeBuilder.WithLaboralSchedule(*cmd.LaboralSchedule)
	//}

	return *employeeBuilder.Build()
}
