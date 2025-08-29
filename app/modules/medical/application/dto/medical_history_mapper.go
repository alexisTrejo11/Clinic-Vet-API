package dto

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
)

func FromCreateDTO(dto MedicalHistoryCreate) (entity.MedicalHistory, error) {
	PetID, err := valueobject.NewPetID(dto.PetID)
	if err != nil {
		return entity.MedicalHistory{}, err
	}

	VetID, err := valueobject.NewVetID(dto.VetID)
	if err != nil {
		return entity.MedicalHistory{}, err
	}

	visitReason, err := enum.NewVisitReason(string(dto.VisitReason))
	if err != nil {
		return entity.MedicalHistory{}, err
	}

	visitType, err := enum.NewVisitType(string(dto.VisitType))
	if err != nil {
		return entity.MedicalHistory{}, err
	}

	condition, err := enum.NewPetCondition(string(dto.Condition))
	if err != nil {
		return entity.MedicalHistory{}, err
	}

	return entity.MedicalHistory{
		PetID:       PetID,
		VetID:       VetID,
		OwnerID:     dto.OwnerID,
		VisitDate:   dto.Date,
		Diagnosis:   dto.Diagnosis,
		Treatment:   dto.Treatment,
		Notes:       dto.Notes,
		Condition:   condition,
		VisitReason: visitReason,
		VisitType:   visitType,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

// ADD MISSING FIELDS
func FromUpdateDTO(dto MedicalHistoryUpdate, existingMedHis entity.MedicalHistory) (entity.MedicalHistory, error) {
	if dto.PetID != nil {
		petID, err := valueobject.NewPetID(*dto.PetID)
		if err != nil {
			return entity.MedicalHistory{}, err
		}
		existingMedHis.PetID = petID
	}

	if dto.Date != nil {
		existingMedHis.VisitDate = *dto.Date
	}

	if dto.Notes != nil {
		existingMedHis.Notes = dto.Notes
	}

	if dto.VisitReason != nil {
		reason, err := enum.NewVisitReason(string(*dto.VisitReason))
		if err != nil {
			return entity.MedicalHistory{}, err
		}
		existingMedHis.VisitReason = reason
	}

	if dto.VisitType != nil {
		visitType, err := enum.NewVisitType(string(*dto.VisitType))
		if err != nil {
			return entity.MedicalHistory{}, err
		}

		existingMedHis.VisitType = visitType
	}

	if dto.OwnerID != nil {
		existingMedHis.OwnerID = *dto.OwnerID
	}

	if dto.VetID != nil {
		vetID, err := valueobject.NewVetID(*dto.VetID)
		if err != nil {
			return entity.MedicalHistory{}, err
		}
		existingMedHis.VetID = vetID
	}

	existingMedHis.UpdatedAt = time.Now()

	return existingMedHis, nil
}

func ToResponse(medHistory entity.MedicalHistory) MedHistResponse {
	return MedHistResponse{
		ID:          medHistory.ID,
		PetID:       medHistory.PetID,
		OwnerID:     medHistory.OwnerID(),
		Date:        medHistory.VisitDate,
		Notes:       medHistory.Notes,
		Diagnosis:   medHistory.Diagnosis,
		Condition:   medHistory.Condition.ToString(),
		VisitType:   medHistory.VisitType.ToString(),
		VisitReason: medHistory.VisitReason.ToString(),
		Treatment:   medHistory.Treatment,
		VetID:       medHistory.VetID.GetValue(),
		CreatedAt:   medHistory.CreatedAt,
		UpdatedAt:   medHistory.UpdatedAt,
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
		ID:  medHistory.ID.GetValue(),
		Pet: *petDetail,
		Owner: OwnerDetails{
			ID:        owner.ID(),
			FirstName: owner.FullName().FirstName,
			LastName:  owner.FullName().LastName,
		},
		Date:      medHistory.VisitDate,
		Diagnosis: medHistory.Diagnosis,
		Treatment: medHistory.Treatment,
		Veterinarian: VetDetails{
			ID:        vet.GetID(),
			FirstName: vet.GetName().FirstName,
			Specialty: vet.GetSpecialty().String(),
			LastName:  vet.GetName().LastName,
		},
		CreatedAt: medHistory.CreatedAt,
		UpdatedAt: medHistory.UpdatedAt,
	}

	if medHistory.Notes != nil {
		detail.Notes = *medHistory.Notes
	}

	return *detail
}
