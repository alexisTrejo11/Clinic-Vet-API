// Package mapper defines functions to map between employee entities and DTOs.
package mapper

import (
	"errors"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/employee"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/employee/application/dto"
)

func FromCreateDTO(vetData dto.CreateEmployeeData) (*employee.Employee, error) {
	if vetData.FirstName == "" || vetData.LastName == "" {
		return nil, errors.New("first name and last name are required")
	}
	if vetData.LicenseNumber == "" {
		return nil, errors.New("license number is required")
	}

	personName, err := valueobject.NewPersonName(vetData.FirstName, vetData.LastName)
	if err != nil {
		return nil, fmt.Errorf("error creating person name: %w", err)
	}

	specialty, err := enum.ParseVetSpecialty(vetData.Specialty)
	if err != nil {
		return nil, fmt.Errorf("invalid specialty '%s': %w", vetData.Specialty, err)
	}

	// Crear options
	opts := []employee.EmployeeOption{
		employee.WithName(personName),
		employee.WithLicenseNumber(vetData.LicenseNumber),
		employee.WithSpecialty(specialty),
		employee.WithYearsExperience(vetData.YearsExperience),
		employee.WithIsActive(vetData.IsActive),
	}

	if vetData.Photo != "" {
		opts = append(opts, employee.WithPhoto(vetData.Photo))
	}

	if vetData.ConsultationFee != nil {
		if vetData.ConsultationFee.Amount() < 0 {
			return nil, errors.New("consultation fee cannot be negative")
		}
		opts = append(opts, employee.WithConsultationFee(vetData.ConsultationFee))
	}

	employee, err := employee.NewEmployee(valueobject.EmployeeID{}, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create employee: %w", err)
	}

	return employee, nil
}

func ToResponse(employee *employee.Employee) dto.EmployeeResponse {
	var scheduleResponses *[]dto.ScheduleData
	if employee.Schedule() != nil {
		days := employee.Schedule().WorkDays

		scheduleResponsesSlice := make([]dto.ScheduleData, len(days))
		for i, day := range days {
			shcedule := dto.ScheduleData{
				Day:           day.Day.String(),
				EntryTime:     day.StartHour,
				DepartureTime: day.EndHour,
				StartBreak:    day.Breaks.StartHour,
				EndBreak:      day.Breaks.EndHour,
			}
			scheduleResponsesSlice[i] = shcedule
		}
		scheduleResponses = &scheduleResponsesSlice
	}

	response := &dto.EmployeeResponse{
		ID:              employee.ID().Value(),
		FirstName:       employee.Name().FirstName,
		LastName:        employee.Name().LastName,
		Photo:           employee.Photo(),
		LicenseNumber:   employee.LicenseNumber(),
		Specialty:       employee.Specialty().String(),
		YearsExperience: employee.YearsExperience(),
		ConsultationFee: employee.ConsultationFee(),
		LaboralSchedule: scheduleResponses,
	}

	return *response
}

func UpdateFromDTO(employee *employee.Employee, vetData dto.UpdateEmployeeData) error {
	// Actualizar nombre si se proporciona
	if vetData.FirstName != nil || vetData.LastName != nil {
		currentName := employee.Name()
		newFirstName := currentName.FirstName
		newLastName := currentName.LastName

		if vetData.FirstName != nil {
			newFirstName = *vetData.FirstName
		}
		if vetData.LastName != nil {
			newLastName = *vetData.LastName
		}

		updatedName, err := valueobject.NewPersonName(newFirstName, newLastName)
		if err != nil {
			return fmt.Errorf("error updating person name: %w", err)
		}

		if err := employee.UpdateName(updatedName); err != nil {
			return fmt.Errorf("failed to update name: %w", err)
		}
	}

	// Actualizar photo
	if vetData.Photo != nil {
		if err := employee.UpdatePhoto(*vetData.Photo); err != nil {
			return fmt.Errorf("failed to update photo: %w", err)
		}
	}

	// Actualizar license number
	if vetData.LicenseNumber != nil {
		if err := employee.UpdateLicenseNumber(*vetData.LicenseNumber); err != nil {
			return fmt.Errorf("failed to update license number: %w", err)
		}
	}

	// Actualizar specialty
	if vetData.Specialty != nil {
		specialty, err := enum.ParseVetSpecialty(*vetData.Specialty)
		if err != nil {
			return fmt.Errorf("invalid specialty: %w", err)
		}
		if err := employee.UpdateSpecialty(specialty); err != nil {
			return fmt.Errorf("failed to update specialty: %w", err)
		}
	}

	// Actualizar years experience
	if vetData.YearsExperience != nil {
		if err := employee.UpdateYearsExperience(*vetData.YearsExperience); err != nil {
			return fmt.Errorf("failed to update years experience: %w", err)
		}
	}

	// Actualizar consultation fee
	if vetData.ConsultationFee != nil {
		if err := employee.UpdateConsultationFee(vetData.ConsultationFee); err != nil {
			return fmt.Errorf("failed to update consultation fee: %w", err)
		}
	}

	if vetData.IsActive != nil {
		if *vetData.IsActive {
			if err := employee.Activate(); err != nil {
				return fmt.Errorf("failed to activate employee: %w", err)
			}
		} else {
			if err := employee.Deactivate(); err != nil {
				return fmt.Errorf("failed to deactivate employee: %w", err)
			}
		}
	}

	return nil
}
