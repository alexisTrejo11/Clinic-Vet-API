package query

import (
	"clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

func toResult(entity medical.MedicalSession) MedSessionResult {
	var notes *string
	if entityNotes := entity.Notes(); entityNotes != nil {
		notes = entityNotes
	}

	var weight *valueobject.Decimal
	if entity.Weight() != nil {
		weight = entity.Weight()
	}

	var temperature *valueobject.Decimal
	if entity.Temperature() != nil {
		temperature = entity.Temperature()
	}

	result := MedSessionResult{
		ID:              entity.ID(),
		PetID:           entity.PetID(),
		CustomerID:      entity.CustomerID(),
		EmployeeID:      entity.EmployeeID(),
		Date:            entity.VisitDate(),
		VisitType:       entity.VisitType().DisplayName(), // Usar DisplayName en lugar de String()
		VisitReason:     entity.VisitReason().DisplayName(),
		Diagnosis:       entity.Diagnosis(),
		Treatment:       entity.Treatment(),
		Condition:       entity.Condition().DisplayName(),
		Notes:           notes,
		Weight:          weight,
		Temperature:     temperature,
		HeartRate:       entity.HeartRate(),
		RespiratoryRate: entity.RespiratoryRate(),
		Symptoms:        entity.Symptoms(),
		Medications:     entity.Medications(),
		FollowUpDate:    entity.FollowUpDate(),
		IsEmergency:     entity.VisitReason() == enum.VisitReasonEmergency,
		CreatedAt:       entity.CreatedAt(),
		UpdatedAt:       entity.UpdatedAt(),
	}

	return result
}

func toResultList(entities []medical.MedicalSession) []MedSessionResult {
	dtos := make([]MedSessionResult, len(entities))
	for i, entity := range entities {
		dtos[i] = toResult(entity)
	}
	return dtos
}
