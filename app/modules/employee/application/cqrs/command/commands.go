// Package command contains the command definitions and handlers for employee-related operations.
package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/employee"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type CreateEmployeeCommand struct {
	// Personal details
	Name        valueobject.PersonName
	Gender      enum.PersonGender
	DateOfBirth time.Time

	// Professional details
	Photo           string
	LicenseNumber   string
	YearsExperience int32
	IsActive        bool
	Specialty       enum.VetSpecialty
	LaboralSchedule []ScheduleData
}

type UpdateEmployeeCommand struct {
	EmployeeID      valueobject.EmployeeID
	Name            *valueobject.PersonName
	Photo           *string
	LicenseNumber   *string
	YearsExperience *int32
	Specialty       *enum.VetSpecialty
	IsActive        *bool
	ConsultationFee *valueobject.Money
	LaboralSchedule *[]ScheduleData
}

type DeleteEmployeeCommand struct {
	EmployeeID valueobject.EmployeeID
}

func NewDeleteEmployeeCommand(id uint) *DeleteEmployeeCommand {
	return &DeleteEmployeeCommand{
		EmployeeID: valueobject.NewEmployeeID(id),
	}
}

type ScheduleData struct {
	Day           string
	EntryTime     int
	DepartureTime int
	StartBreak    int
	EndBreak      int
}

func (cmd *CreateEmployeeCommand) ToEntity() employee.Employee {
	employee := employee.NewEmployeeBuilder().
		WithLicenseNumber(cmd.LicenseNumber).
		WithSpecialty(cmd.Specialty).
		WithYearsExperience(cmd.YearsExperience).
		WithIsActive(cmd.IsActive).
		WithPhoto(cmd.Photo).
		WithIsActive(cmd.IsActive).
		Build()

	return *employee
}

func (cmd *UpdateEmployeeCommand) updateEmployee(emp employee.Employee) employee.Employee {
	employeeBuilder := employee.NewEmployeeBuilder()

	if cmd.Name != nil {
		employeeBuilder.WithName(*cmd.Name)
	}
	if cmd.Photo != nil {
		employeeBuilder.WithPhoto(*cmd.Photo)
	}
	if cmd.LicenseNumber != nil {
		employeeBuilder.WithLicenseNumber(*cmd.LicenseNumber)
	}
	if cmd.YearsExperience != nil {
		employeeBuilder.WithYearsExperience(*cmd.YearsExperience)
	}
	if cmd.Specialty != nil {
		employeeBuilder.WithSpecialty(*cmd.Specialty)
	}
	if cmd.IsActive != nil {
		employeeBuilder.WithIsActive(*cmd.IsActive)
	}
	//if cmd.LaboralSchedule != nil {
	//	employeeBuilder.WithLaboralSchedule(*cmd.LaboralSchedule)
	//}

	return *employeeBuilder.Build()
}
