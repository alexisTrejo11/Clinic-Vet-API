package vetMapper

import (
	"fmt"
	"time"

	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	vetDtos "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

func FromCreateDTO(vetData vetDtos.VetCreate) *vetDomain.Veterinarian {
	personName, _ := user.NewPersonName(vetData.FirstName, vetData.LastName)
	return &vetDomain.Veterinarian{
		Name:            personName,
		Photo:           vetData.Photo,
		LicenseNumber:   vetData.LicenseNumber,
		Specialty:       vetData.Specialty,
		YearsExperience: vetData.YearsExperience,
		ConsultationFee: vetData.ConsultationFee,
		IsActive:        vetData.IsActive,
		//WorkDaysSchedule: vetData.LaboralSchedule,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func UpdateFromDTO(vet *vetDomain.Veterinarian, vetData vetDtos.VetUpdate) {
	if vetData.FirstName != nil || vetData.LastName != nil {
		currentFirstName := vet.Name.FirstName
		currentLastName := vet.Name.LastName

		if vetData.FirstName != nil {
			currentFirstName = *vetData.FirstName
		}
		if vetData.LastName != nil {
			currentLastName = *vetData.LastName
		}

		updatedName, _ := user.NewPersonName(currentFirstName, currentLastName)
		vet.Name = updatedName

		if vetData.Photo != nil {
			vet.Photo = *vetData.Photo
		}

		if vetData.LicenseNumber != nil {
			vet.LicenseNumber = *vetData.LicenseNumber
		}

		if vetData.Specialty != nil {
			vet.Specialty = *vetData.Specialty
		}

		if vetData.YearsExperience != nil {
			vet.YearsExperience = *vetData.YearsExperience
		}

		if vetData.ConsultationFee != nil {
			vet.ConsultationFee = vetData.ConsultationFee
		}

		if vetData.IsActive != nil {
			vet.IsActive = *vetData.IsActive
		}

		vet.UpdatedAt = time.Now()
	}
}

func ToResponse(vet vetDomain.Veterinarian) *vetDtos.VetResponse {
	fmt.Println("MAPPER")
	if vet.Schedule != nil {
		days := vet.Schedule.WorkDays

		scheduleResponses := make([]vetDtos.ScheduleInsert, len(days))
		for i, day := range days {
			shcedule := vetDtos.ScheduleInsert{
				Day:           day.Day,
				EntryTime:     day.StartHour,
				DepartureTime: day.EndHour,
				StartBreak:    day.Breaks.StartHour,
				EndBreak:      day.Breaks.EndHour,
			}
			scheduleResponses[i] = shcedule
		}

		return &vetDtos.VetResponse{
			Id:              vet.ID,
			FirstName:       vet.Name.FirstName,
			LastName:        vet.Name.LastName,
			Photo:           vet.Photo,
			LicenseNumber:   vet.LicenseNumber,
			Specialty:       vet.Specialty.String(),
			YearsExperience: vet.YearsExperience,
			ConsultationFee: vet.ConsultationFee,
			LaboralSchedule: &scheduleResponses,
		}
	}

	return &vetDtos.VetResponse{
		Id:              vet.ID,
		FirstName:       vet.Name.FirstName,
		LastName:        vet.Name.LastName,
		Photo:           vet.Photo,
		LicenseNumber:   vet.LicenseNumber,
		Specialty:       vet.Specialty.String(),
		YearsExperience: vet.YearsExperience,
		ConsultationFee: vet.ConsultationFee,
	}
}
