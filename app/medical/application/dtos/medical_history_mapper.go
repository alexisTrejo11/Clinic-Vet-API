package mhDTOs

import (
	"time"

	mhDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/domain"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

func FromCreateDTO(dto MedicalHistoryCreate) (mhDomain.MedicalHistory, error) {
	PetId, err := petDomain.NewPetId(dto.PetId)
	if err != nil {
		return mhDomain.MedicalHistory{}, err
	}

	VetId, err := vetDomain.NewVeterinarianId(dto.VetId)
	if err != nil {
		return mhDomain.MedicalHistory{}, err
	}

	visitReason, err := mhDomain.NewVisitReason(string(dto.VisitReason))
	if err != nil {
		return mhDomain.MedicalHistory{}, err
	}

	visitType, err := mhDomain.NewVisitType(string(dto.VisitType))
	if err != nil {
		return mhDomain.MedicalHistory{}, err
	}

	condition, err := mhDomain.NewPetCondition(string(dto.Condition))
	if err != nil {
		return mhDomain.MedicalHistory{}, err
	}

	return mhDomain.MedicalHistory{
		PetId:       PetId,
		VetId:       VetId,
		OwnerId:     dto.OwnerId,
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
func FromUpdateDTO(dto MedicalHistoryUpdate, existingMedHis mhDomain.MedicalHistory) (mhDomain.MedicalHistory, error) {
	if dto.PetId != nil {
		petId, err := petDomain.NewPetId(*dto.PetId)
		if err != nil {
			return mhDomain.MedicalHistory{}, err
		}
		existingMedHis.PetId = petId
	}

	if dto.Date != nil {
		existingMedHis.VisitDate = *dto.Date
	}

	if dto.Notes != nil {
		existingMedHis.Notes = dto.Notes
	}

	if dto.VisitReason != nil {
		reason, err := mhDomain.NewVisitReason(string(*dto.VisitReason))
		if err != nil {
			return mhDomain.MedicalHistory{}, err
		}
		existingMedHis.VisitReason = reason
	}

	if dto.VisitType != nil {
		visitType, err := mhDomain.NewVisitType(string(*dto.VisitType))
		if err != nil {
			return mhDomain.MedicalHistory{}, err
		}

		existingMedHis.VisitType = visitType
	}

	if dto.OwnerId != nil {
		existingMedHis.OwnerId = *dto.OwnerId
	}

	if dto.VetId != nil {
		vetId, err := vetDomain.NewVeterinarianId(*dto.VetId)
		if err != nil {
			return mhDomain.MedicalHistory{}, err
		}
		existingMedHis.VetId = vetId
	}

	existingMedHis.UpdatedAt = time.Now()

	return existingMedHis, nil
}

func ToResponse(medHistory mhDomain.MedicalHistory) MedHistResponse {
	return MedHistResponse{
		Id:          medHistory.Id.GetValue(),
		PetId:       medHistory.PetId.GetValue(),
		OwnerId:     medHistory.OwnerId,
		Date:        medHistory.VisitDate,
		Notes:       medHistory.Notes,
		Diagnosis:   medHistory.Diagnosis,
		Condition:   medHistory.Condition.ToString(),
		VisitType:   medHistory.VisitType.ToString(),
		VisitReason: medHistory.VisitReason.ToString(),
		Treatment:   medHistory.Treatment,
		VetId:       medHistory.VetId.GetValue(),
		CreatedAt:   medHistory.CreatedAt,
		UpdatedAt:   medHistory.UpdatedAt,
	}
}

func ListToResponse(medHistories []mhDomain.MedicalHistory) []MedHistResponse {
	var response []MedHistResponse

	if len(medHistories) == 0 {
		return []MedHistResponse{}
	}

	for _, medHistory := range medHistories {
		response = append(response, ToResponse(medHistory))
	}
	return response
}

func ToResponseDetail(medHistory mhDomain.MedicalHistory, owner ownerDomain.Owner, vet vetDomain.Veterinarian, pet petDomain.Pet) MedHistResponseDetail {
	petDetail := &PetDetails{
		Id:      pet.GetID(),
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
		Id:  medHistory.Id.GetValue(),
		Pet: *petDetail,
		Owner: OwnerDetails{
			Id:        owner.Id(),
			FirstName: owner.FullName().FirstName,
			LastName:  owner.FullName().LastName,
		},
		Date:      medHistory.VisitDate,
		Diagnosis: medHistory.Diagnosis,
		Treatment: medHistory.Treatment,
		Veterinarian: VetDetails{
			Id:        vet.GetID(),
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
