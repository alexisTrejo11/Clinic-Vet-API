// Package command contains the command definitions for employee-related operations.
package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/employee"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
	"time"
)

type CreateEmployeeCommand struct {
	// Personal details
	name        valueobject.PersonName
	gender      enum.PersonGender
	dateOfBirth time.Time

	// Professional details
	photo           string
	licenseNumber   string
	yearsExperience int32
	isActive        bool
	specialty       enum.VetSpecialty
	laboralSchedule []ScheduleData
}

func NewCreateEmployeeCommand(
	firstName, lastName,
	gender,
	photo,
	licenseNumber,
	specialty string,
	dateOfBirth time.Time,
	yearsExperience int32,
	isActive bool,
	laboralSchedule []ScheduleData,
) (CreateEmployeeCommand, error) {
	cmd := CreateEmployeeCommand{
		name:            valueobject.NewPersonNameNoErr(firstName, lastName),
		gender:          enum.NewPersonGender(gender),
		dateOfBirth:     dateOfBirth,
		photo:           photo,
		licenseNumber:   licenseNumber,
		yearsExperience: yearsExperience,
		isActive:        isActive,
		specialty:       enum.VetSpecialty(specialty),
		laboralSchedule: laboralSchedule,
	}

	if err := cmd.validate(); err != nil {
		return CreateEmployeeCommand{}, err
	}

	return cmd, nil
}

func (cmd *CreateEmployeeCommand) Name() valueobject.PersonName    { return cmd.name }
func (cmd *CreateEmployeeCommand) Gender() enum.PersonGender       { return cmd.gender }
func (cmd *CreateEmployeeCommand) DateOfBirth() time.Time          { return cmd.dateOfBirth }
func (cmd *CreateEmployeeCommand) Photo() string                   { return cmd.photo }
func (cmd *CreateEmployeeCommand) LicenseNumber() string           { return cmd.licenseNumber }
func (cmd *CreateEmployeeCommand) YearsExperience() int32          { return cmd.yearsExperience }
func (cmd *CreateEmployeeCommand) IsActive() bool                  { return cmd.isActive }
func (cmd *CreateEmployeeCommand) Specialty() enum.VetSpecialty    { return cmd.specialty }
func (cmd *CreateEmployeeCommand) LaboralSchedule() []ScheduleData { return cmd.laboralSchedule }

func (cmd *CreateEmployeeCommand) validate() error {
	if !cmd.name.IsValid() {
		return createEmployeeCmdErr("name", "cannot be empty")
	}

	if !cmd.gender.IsValid() {
		return createEmployeeCmdErr("gender", "must be either 'male' or 'female'")
	}

	if cmd.dateOfBirth.IsZero() {
		return createEmployeeCmdErr("dateOfBirth", "must be a valid date")
	}

	if cmd.licenseNumber == "" {
		return createEmployeeCmdErr("licenseNumber", "cannot be empty")
	}

	if cmd.yearsExperience < 0 {
		return createEmployeeCmdErr("yearsExperience", "must be a positive number")
	}

	if !cmd.specialty.IsValid() {
		return createEmployeeCmdErr("specialty", "must be a valid specialty")
	}

	return nil
}

func createEmployeeCmdErr(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "CreateEmployeeCommand")
}

func (cmd *CreateEmployeeCommand) ToEntity() employee.Employee {
	return *employee.NewEmployeeBuilder().
		WithLicenseNumber(cmd.licenseNumber).
		WithSpecialty(cmd.specialty).
		WithYearsExperience(cmd.yearsExperience).
		WithIsActive(cmd.isActive).
		WithPhoto(cmd.photo).
		WithIsActive(cmd.isActive).
		Build()
}
