package mhDTOs

import (
	"time"

	mhDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/domain"

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

	return mhDomain.MedicalHistory{
		PetId:       PetId,
		OwnerId:     dto.OwnerId,
		VisitDate:   dto.Date,
		VisitReason: dto.VisitReason,
		VisitType:   dto.VisitType,
		Description: dto.Description,
		VetId:       VetId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

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
	if dto.Description != nil {
		existingMedHis.Description = dto.Description
	}

	if dto.VisitReason != nil {
		existingMedHis.VisitReason = *dto.VisitReason
	}

	if dto.VisitType != nil {
		existingMedHis.VisitType = *dto.VisitType
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
		Description: medHistory.Description,
		VetId:       medHistory.VetId.GetValue(),
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
