package dto

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
)

func FromCreateDTO(dto MedicalHistoryCreate) (entity.MedicalHistory, error) {
	medHistory, err := entity.NewMedicalHistoryBuilder().
		WithPetID(dto.PetID).
		WithVetID(dto.VetID).
		WithOwnerID(dto.OwnerID).
		WithVisitDate(dto.Date).
		WithDiagnosis(dto.Diagnosis).
		WithTreatment(dto.Treatment).
		WithNotes(dto.Notes).
		WithVisitReason(string(dto.VisitReason)).
		WithVisitType(string(dto.VisitType)).
		WithCondition(string(dto.Condition)).
		WithTimestamps(time.Now(), time.Now()).
		Build()
	if err != nil {
		return entity.MedicalHistory{}, err
	}

	return medHistory, nil
}

func ToResponse(medHistory entity.MedicalHistory) MedHistResponse {
	return MedHistResponse{
		ID:          medHistory.ID().GetValue(),
		PetID:       medHistory.PetID().GetValue(),
		OwnerID:     medHistory.OwnerID().GetValue(),
		Date:        medHistory.VisitDate(),
		Notes:       medHistory.Notes(),
		Diagnosis:   medHistory.Diagnosis(),
		Condition:   medHistory.Condition().ToString(),
		VisitType:   medHistory.VisitType().ToString(),
		VisitReason: medHistory.VisitReason().ToString(),
		Treatment:   medHistory.Treatment(),
		VetID:       medHistory.VetID().GetValue(),
		CreatedAt:   medHistory.CreatedAt(),
		UpdatedAt:   medHistory.UpdatedAt(),
	}
}

func ListToResponse(medHistories []entity.MedicalHistory) []MedHistResponse {
	var response []MedHistResponse

	if len(medHistories) == 0 {
		return []MedHistResponse{}
	}

	for _, medHistory := range medHistories {
		response = append(response, ToResponse(medHistory))
	}
	return response
}

func ToResponseDetail(
	medHistory entity.MedicalHistory,
	owner entity.Owner,
	vet entity.Veterinarian,
	pet entity.Pet,
) MedHistResponseDetail {
	petDetail := &PetDetails{
		ID:      medHistory.ID().GetValue(),
		Name:    pet.GetName(),
		Species: pet.GetSpecies(),
	}

	if pet.GetAge() != nil {
		petDetail.Age = *pet.GetAge()
	}

	if pet.GetBreed() == nil {
		petDetail.Breed = "Unknown"
	} else {
		petDetail.Breed = *pet.GetBreed()
	}

	if pet.GetWeight() != nil {
		petDetail.Weight = *pet.GetWeight()
	}

	detail := &MedHistResponseDetail{
		ID:  medHistory.ID().GetValue(),
		Pet: *petDetail,
		Owner: OwnerDetails{
			ID:        owner.ID(),
			FirstName: owner.FullName().FirstName,
			LastName:  owner.FullName().LastName,
		},
		Date:      medHistory.VisitDate(),
		Diagnosis: medHistory.Diagnosis(),
		Treatment: medHistory.Treatment(),
		Veterinarian: VetDetails{
			ID:        vet.GetID(),
			FirstName: vet.GetName().FirstName,
			Specialty: vet.GetSpecialty().String(),
			LastName:  vet.GetName().LastName,
		},
		CreatedAt: medHistory.CreatedAt(),
		UpdatedAt: medHistory.UpdatedAt(),
	}

	if medHistory.Notes != nil {
		detail.Notes = *medHistory.Notes()
	}

	return *detail
}
