// Package mapper defines functions to map between Veterinarian entities and DTOs.
package mapper

import (
	"errors"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/veterinarian"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/dto"
)

func FromCreateDTO(vetData dto.CreateVetData) (*veterinarian.Veterinarian, error) {
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
	opts := []veterinarian.VeterinarianOption{
		veterinarian.WithName(personName),
		veterinarian.WithLicenseNumber(vetData.LicenseNumber),
		veterinarian.WithSpecialty(specialty),
		veterinarian.WithYearsExperience(vetData.YearsExperience),
		veterinarian.WithIsActive(vetData.IsActive),
	}

	if vetData.Photo != "" {
		opts = append(opts, veterinarian.WithPhoto(vetData.Photo))
	}

	if vetData.ConsultationFee != nil {
		if vetData.ConsultationFee.Amount() < 0 {
			return nil, errors.New("consultation fee cannot be negative")
		}
		opts = append(opts, veterinarian.WithConsultationFee(vetData.ConsultationFee))
	}

	vet, err := veterinarian.NewVeterinarian(valueobject.VetID{}, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create veterinarian: %w", err)
	}

	return vet, nil
}

func ToResponse(vet *veterinarian.Veterinarian) dto.VetResponse {
	var scheduleResponses *[]dto.ScheduleData
	if vet.Schedule() != nil {
		days := vet.Schedule().WorkDays

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

	response := &dto.VetResponse{
		ID:              vet.ID().Value(),
		FirstName:       vet.Name().FirstName,
		LastName:        vet.Name().LastName,
		Photo:           vet.Photo(),
		LicenseNumber:   vet.LicenseNumber(),
		Specialty:       vet.Specialty().String(),
		YearsExperience: vet.YearsExperience(),
		ConsultationFee: vet.ConsultationFee(),
		LaboralSchedule: scheduleResponses,
	}

	return *response
}

func UpdateFromDTO(vet *veterinarian.Veterinarian, vetData dto.UpdateVetData) error {
	// Actualizar nombre si se proporciona
	if vetData.FirstName != nil || vetData.LastName != nil {
		currentName := vet.Name()
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

		if err := vet.UpdateName(updatedName); err != nil {
			return fmt.Errorf("failed to update name: %w", err)
		}
	}

	// Actualizar photo
	if vetData.Photo != nil {
		if err := vet.UpdatePhoto(*vetData.Photo); err != nil {
			return fmt.Errorf("failed to update photo: %w", err)
		}
	}

	// Actualizar license number
	if vetData.LicenseNumber != nil {
		if err := vet.UpdateLicenseNumber(*vetData.LicenseNumber); err != nil {
			return fmt.Errorf("failed to update license number: %w", err)
		}
	}

	// Actualizar specialty
	if vetData.Specialty != nil {
		specialty, err := enum.ParseVetSpecialty(*vetData.Specialty)
		if err != nil {
			return fmt.Errorf("invalid specialty: %w", err)
		}
		if err := vet.UpdateSpecialty(specialty); err != nil {
			return fmt.Errorf("failed to update specialty: %w", err)
		}
	}

	// Actualizar years experience
	if vetData.YearsExperience != nil {
		if err := vet.UpdateYearsExperience(*vetData.YearsExperience); err != nil {
			return fmt.Errorf("failed to update years experience: %w", err)
		}
	}

	// Actualizar consultation fee
	if vetData.ConsultationFee != nil {
		if err := vet.UpdateConsultationFee(vetData.ConsultationFee); err != nil {
			return fmt.Errorf("failed to update consultation fee: %w", err)
		}
	}

	if vetData.IsActive != nil {
		if *vetData.IsActive {
			if err := vet.Activate(); err != nil {
				return fmt.Errorf("failed to activate veterinarian: %w", err)
			}
		} else {
			if err := vet.Deactivate(); err != nil {
				return fmt.Errorf("failed to deactivate veterinarian: %w", err)
			}
		}
	}

	return nil
}
