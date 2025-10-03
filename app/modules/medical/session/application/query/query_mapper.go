package query

import (
	"clinic-vet-api/app/modules/core/domain/entity/medical"
)

func toResult(entity medical.MedicalSession) MedSessionResult {
	return MedSessionResult{
		ID:            entity.ID(),
		EmployeeID:    entity.EmployeeID(),
		VisitDate:     entity.VisitDate(),
		VisitType:     entity.VisitType(),
		ClinicService: entity.Service(),
		Notes:         entity.Notes(),
		CreatedAt:     entity.CreatedAt(),
		UpdatedAt:     entity.UpdatedAt(),
		PetDetailsResult: PetDetailsResult{
			PetID:           entity.PetDetails().PetID(),
			FollowUpDate:    entity.PetDetails().FollowUpDate(),
			Weight:          entity.PetDetails().Weight(),
			Diagnosis:       entity.PetDetails().Diagnosis(),
			Temperature:     entity.PetDetails().Temperature(),
			HeartRate:       entity.PetDetails().HeartRate(),
			RespiratoryRate: entity.PetDetails().RespiratoryRate(),
			Condition:       entity.PetDetails().Condition(),
			Treatment:       entity.PetDetails().Treatment(),
			Symptoms:        entity.PetDetails().Symptoms(),
			Medications:     entity.PetDetails().Medications(),
		},
	}
}

func toResultList(entities []medical.MedicalSession) []MedSessionResult {
	dtos := make([]MedSessionResult, len(entities))
	for i, entity := range entities {
		dtos[i] = toResult(entity)
	}
	return dtos
}
