package dto

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/medical"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/owner"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/veterinarian"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

func FromCreateDTO(dto MedicalHistoryCreate) (medical.MedicalHistory, error) {
	medHistory, err := medical.NewMedicalHistory(
		valueobject.MedHistoryID{},
		dto.PetID, dto.OwnerID, dto.VetID,
		medical.WithVisitDate(dto.Date),
		medical.WithDiagnosis(dto.Diagnosis),
		medical.WithTreatment(dto.Treatment),
		medical.WithNotes(*dto.Notes),
		medical.WithVisitReason(dto.VisitReason),
		medical.WithVisitType(dto.VisitType),
		medical.WithCondition(dto.Condition))
	if err != nil {
		return medical.MedicalHistory{}, err
	}

	return *medHistory, nil
}

func ToResponse(medHistory medical.MedicalHistory) MedHistResponse {
	return MedHistResponse{
		ID:          medHistory.ID().Value(),
		PetID:       medHistory.PetID().Value(),
		OwnerID:     medHistory.OwnerID().Value(),
		Date:        medHistory.VisitDate(),
		Notes:       medHistory.Notes(),
		Diagnosis:   medHistory.Diagnosis(),
		Condition:   medHistory.Condition().DisplayName(),
		VisitType:   medHistory.VisitType().DisplayName(),
		VisitReason: medHistory.VisitReason().DisplayName(),
		Treatment:   medHistory.Treatment(),
		VetID:       medHistory.VetID().Value(),
		CreatedAt:   medHistory.CreatedAt(),
		UpdatedAt:   medHistory.UpdatedAt(),
	}
}

func ListToResponse(medHistories []medical.MedicalHistory) []MedHistResponse {
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
	medHistory medical.MedicalHistory,
	owner owner.Owner,
	vet veterinarian.Veterinarian,
	pet pet.Pet,
) MedHistResponseDetail {
	petDetail := &PetDetails{
		ID:      medHistory.ID().Value(),
		Name:    pet.Name(),
		Species: pet.Species(),
	}

	if pet.Age() != nil {
		petDetail.Age = *pet.Age()
	}

	if pet.Breed() == nil {
		petDetail.Breed = "Unknown"
	} else {
		petDetail.Breed = *pet.Breed()
	}

	if pet.Weight() != nil {
		petDetail.Weight = *pet.Weight()
	}

	detail := &MedHistResponseDetail{
		ID:  medHistory.ID().Value(),
		Pet: *petDetail,
		Owner: OwnerDetails{
			ID:        owner.ID().Value(),
			FirstName: owner.FullName().FirstName,
			LastName:  owner.FullName().LastName,
		},
		Date:      medHistory.VisitDate(),
		Diagnosis: medHistory.Diagnosis(),
		Treatment: medHistory.Treatment(),
		Veterinarian: VetDetails{
			ID:        vet.ID().Value(),
			FirstName: vet.Name().FirstName,
			Specialty: vet.Specialty().String(),
			LastName:  vet.Name().LastName,
		},
		CreatedAt: medHistory.CreatedAt(),
		UpdatedAt: medHistory.UpdatedAt(),
	}

	if medHistory.Notes() != nil {
		detail.Notes = *medHistory.Notes()
	}

	return *detail
}
