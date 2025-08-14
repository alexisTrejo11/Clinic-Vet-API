package vetMapper

import (
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
	vetDtos "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

func FromCreateDTO(vetData vetDtos.VetCreate) (*vetDomain.Veterinarian, error) {
	personName, err := valueObjects.NewPersonName(vetData.FirstName, vetData.LastName)
	if err != nil {
		return nil, fmt.Errorf("error creating person name: %w", err)
	}

	builder := vetDomain.NewVeterinarianBuilder().
		WithName(personName).
		WithPhoto(vetData.Photo).
		WithLicenseNumber(vetData.LicenseNumber).
		WithSpecialty(vetData.Specialty).
		WithYearsExperience(vetData.YearsExperience).
		WithConsultationFee(vetData.ConsultationFee).
		WithIsActive(vetData.IsActive).
		WithCreatedAt(time.Now()).
		WithUpdatedAt(time.Now())

	return builder.Build(), nil
}

func UpdateFromDTO(vet *vetDomain.Veterinarian, vetData vetDtos.VetUpdate) error {
	if vetData.FirstName != nil || vetData.LastName != nil {
		currentFirstName := vet.GetName().FirstName
		currentLastName := vet.GetName().LastName

		if vetData.FirstName != nil {
			currentFirstName = *vetData.FirstName
		}
		if vetData.LastName != nil {
			currentLastName = *vetData.LastName
		}

		updatedName, err := valueObjects.NewPersonName(currentFirstName, currentLastName)
		if err != nil {
			return fmt.Errorf("error updating person name: %w", err)
		}
		vet.SetName(updatedName)
	}

	if vetData.Photo != nil {
		vet.SetPhoto(*vetData.Photo)
	}

	if vetData.LicenseNumber != nil {
		vet.SetLicenseNumber(*vetData.LicenseNumber)
	}

	if vetData.Specialty != nil {
		vet.SetSpecialty(*vetData.Specialty)
	}

	if vetData.YearsExperience != nil {
		vet.SetYearsExperience(*vetData.YearsExperience)
	}

	if vetData.ConsultationFee != nil {
		vet.SetConsultationFee(vetData.ConsultationFee)
	}

	if vetData.IsActive != nil {
		vet.SetIsActive(*vetData.IsActive)
	}

	vet.SetUpdatedAt(time.Now())
	return nil
}

func ToResponse(vet *vetDomain.Veterinarian) *vetDtos.VetResponse {
	var scheduleResponses *[]vetDtos.ScheduleInsert
	if vet.GetSchedule() != nil {
		days := vet.GetSchedule().WorkDays

		scheduleResponsesSlice := make([]vetDtos.ScheduleInsert, len(days))
		for i, day := range days {
			shcedule := vetDtos.ScheduleInsert{
				Day:           day.Day,
				EntryTime:     day.StartHour,
				DepartureTime: day.EndHour,
				StartBreak:    day.Breaks.StartHour,
				EndBreak:      day.Breaks.EndHour,
			}
			scheduleResponsesSlice[i] = shcedule
		}
		scheduleResponses = &scheduleResponsesSlice
	}

	return &vetDtos.VetResponse{
		Id:              vet.GetID(),
		FirstName:       vet.GetName().FirstName,
		LastName:        vet.GetName().LastName,
		Photo:           vet.GetPhoto(),
		LicenseNumber:   vet.GetLicenseNumber(),
		Specialty:       vet.GetSpecialty().String(),
		YearsExperience: vet.GetYearsExperience(),
		ConsultationFee: vet.GetConsultationFee(),
		LaboralSchedule: scheduleResponses,
	}
}
