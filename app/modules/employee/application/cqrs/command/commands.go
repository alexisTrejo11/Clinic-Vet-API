// Package command contains the command definitions and handlers for employee-related operations.
package command

import (
	"clinic-vet-api/app/core/domain/entity/employee"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	"context"
	"errors"
	"fmt"
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
	YearsExperience int
	IsActive        bool
	Specialty       enum.VetSpecialty
	ConsultationFee *valueobject.Money
	LaboralSchedule []ScheduleData
}

type UpdateEmployeeCommand struct {
	EmployeeID      valueobject.EmployeeID
	FirstName       *string
	LastName        *string
	Photo           *string
	LicenseNumber   *string
	YearsExperience *int
	Specialty       *string
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

func (cmd *CreateEmployeeCommand) createEmployee(ctx context.Context) (*employee.Employee, error) {
	opts := []employee.EmployeeOption{
		employee.WithLicenseNumber(cmd.LicenseNumber),
		employee.WithSpecialty(cmd.Specialty),
		employee.WithYearsExperience(cmd.YearsExperience),
		employee.WithIsActive(cmd.IsActive),
	}

	if cmd.Photo != "" {
		opts = append(opts, employee.WithPhoto(cmd.Photo))
	}

	if cmd.ConsultationFee != nil {
		if cmd.ConsultationFee.Amount() < 0 {
			return nil, errors.New("consultation fee cannot be negative")
		}
		opts = append(opts, employee.WithConsultationFee(cmd.ConsultationFee))
	}

	return employee.CreateEmployee(ctx, cmd.Name, cmd.Gender, cmd.DateOfBirth, opts...)
}

func (cmd *UpdateEmployeeCommand) updateEmployee(ctx context.Context, emp *employee.Employee) error {
	if cmd.FirstName != nil || cmd.LastName != nil {
		firstName := emp.Name().FirstName
		lastName := emp.Name().LastName

		if cmd.FirstName != nil {
			firstName = *cmd.FirstName
		}
		if cmd.LastName != nil {
			lastName = *cmd.LastName
		}

		personName, err := valueobject.NewPersonName(firstName, lastName)
		if err != nil {
			return fmt.Errorf("error updating person name: %w", err)
		}
		emp.UpdateName(ctx, personName)
	}

	if cmd.Photo != nil {
		emp.UpdatePhoto(*cmd.Photo)
	}

	if cmd.LicenseNumber != nil {
		emp.UpdateLicenseNumber(*cmd.LicenseNumber)
	}

	if cmd.YearsExperience != nil {
		emp.UpdateYearsExperience(*cmd.YearsExperience)
	}

	if cmd.Specialty != nil {
		specialty, err := enum.ParseVetSpecialty(*cmd.Specialty)
		if err != nil {
			return fmt.Errorf("invalid specialty '%s': %w", *cmd.Specialty, err)
		}
		emp.UpdateSpecialty(specialty)
	}

	if cmd.IsActive != nil {
		if *cmd.IsActive {
			emp.Activate()
		} else {
			emp.Deactivate()
		}
	}

	if cmd.ConsultationFee != nil {
		emp.UpdateConsultationFee(cmd.ConsultationFee)
	}

	return nil
}
