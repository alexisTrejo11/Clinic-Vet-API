package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/employee"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
	"time"
)

type UpdateEmployeeCommand struct {
	employeeID          valueobject.EmployeeID
	firstName, lastName *string
	photo               *string
	licenseNumber       *string
	yearsExperience     *int32
	specialty           *enum.VetSpecialty
	isActive            *bool
	gender              *enum.PersonGender
	dateOfBirth         *time.Time
	laboralSchedule     *[]ScheduleData
}

func NewUpdateEmployeeCommand(
	id uint,
	firstName, lastName,
	gender *string,
	dateOfBirth *time.Time,
	photo,
	licenseNumber,
	specialty *string,
	yearsExperience *int32,
	isActive *bool,
	laboralSchedule *[]ScheduleData,
) (UpdateEmployeeCommand, error) {
	cmd := &UpdateEmployeeCommand{
		employeeID:      valueobject.NewEmployeeID(id),
		firstName:       firstName,
		lastName:        lastName,
		gender:          enum.NullableGender(gender),
		dateOfBirth:     dateOfBirth,
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
	if cmd.firstName != nil {
		if cmd.firstName != nil && len(*cmd.firstName) == 0 {
			return updateEmployeeCmdErr("firstName", "cannot be empty if provided")
		}
	}

	if cmd.lastName != nil {
		if cmd.lastName != nil && len(*cmd.lastName) == 0 {
			return updateEmployeeCmdErr("lastName", "cannot be empty if provided")
		}
	}

	if cmd.gender != nil {
		if !cmd.gender.IsValid() {
			return updateEmployeeCmdErr("gender", "invalid gender")
		}
	}

	if cmd.dateOfBirth != nil {
		if cmd.dateOfBirth.IsZero() {
			return updateEmployeeCmdErr("dateOfBirth", "cannot be zero if provided")
		}
	}

	if cmd.photo != nil {
		if len(*cmd.photo) == 0 {
			return updateEmployeeCmdErr("photo", "cannot be empty if provided")
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

	if cmd.gender != nil {
		if !cmd.gender.IsValid() {
			return updateEmployeeCmdErr("gender", "is not a valid gender if provided")
		}
	}

	if cmd.dateOfBirth != nil {
		if cmd.dateOfBirth.IsZero() {
			return updateEmployeeCmdErr("dateOfBirth", "cannot be zero if provided")
		}
	}

	return nil
}

func (cmd *UpdateEmployeeCommand) EmployeeID() valueobject.EmployeeID { return cmd.employeeID }
func (cmd *UpdateEmployeeCommand) FirstName() *string                 { return cmd.firstName }
func (cmd *UpdateEmployeeCommand) LastName() *string                  { return cmd.lastName }
func (cmd *UpdateEmployeeCommand) Photo() *string                     { return cmd.photo }
func (cmd *UpdateEmployeeCommand) LicenseNumber() *string             { return cmd.licenseNumber }
func (cmd *UpdateEmployeeCommand) YearsExperience() *int32            { return cmd.yearsExperience }
func (cmd *UpdateEmployeeCommand) Specialty() *enum.VetSpecialty      { return cmd.specialty }
func (cmd *UpdateEmployeeCommand) IsActive() *bool                    { return cmd.isActive }
func (cmd *UpdateEmployeeCommand) LaboralSchedule() *[]ScheduleData   { return cmd.laboralSchedule }

func updateEmployeeCmdErr(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "UpdateEmployeeCommand")
}

func (cmd *UpdateEmployeeCommand) UpdateEmployee(existing employee.Employee) employee.Employee {
	employeeBuilder := employee.NewEmployeeBuilder().WithID(existing.ID())
	if cmd.firstName != nil {
		employeeBuilder.WithFirstName(*cmd.firstName)
	} else {
		employeeBuilder.WithFirstName(existing.FirstName())
	}

	if cmd.lastName != nil {
		employeeBuilder.WithLastName(*cmd.lastName)
	} else {
		employeeBuilder.WithLastName(existing.LastName())
	}

	if cmd.gender != nil {
		employeeBuilder.WithGender(*cmd.gender)
	} else {
		employeeBuilder.WithGender(existing.Gender())
	}

	if cmd.dateOfBirth != nil {
		employeeBuilder.WithDateOfBirth(*cmd.dateOfBirth)
	} else {
		employeeBuilder.WithDateOfBirth(existing.DateOfBirth())
	}
	if cmd.dateOfBirth != nil {
		employeeBuilder.WithDateOfBirth(*cmd.dateOfBirth)
	} else {
		employeeBuilder.WithDateOfBirth(existing.DateOfBirth())
	}

	if cmd.photo != nil {
		employeeBuilder.WithPhoto(*cmd.photo)
	} else {
		employeeBuilder.WithPhoto(existing.Photo())
	}

	if cmd.licenseNumber != nil {
		employeeBuilder.WithLicenseNumber(*cmd.licenseNumber)
	} else {
		employeeBuilder.WithLicenseNumber(existing.LicenseNumber())
	}

	if cmd.yearsExperience != nil {
		employeeBuilder.WithYearsExperience(*cmd.yearsExperience)
	} else {
		employeeBuilder.WithYearsExperience(existing.YearsExperience())
	}

	if cmd.specialty != nil {
		employeeBuilder.WithSpecialty(*cmd.specialty)
	} else {
		employeeBuilder.WithSpecialty(existing.Specialty())
	}

	if cmd.isActive != nil {
		employeeBuilder.WithIsActive(*cmd.isActive)
	} else {
		employeeBuilder.WithIsActive(existing.IsActive())
	}

	//if cmd.LaboralSchedule != nil {
	//	employeeBuilder.WithLaboralSchedule(*cmd.LaboralSchedule)
	//}
	return *employeeBuilder.Build()
}
