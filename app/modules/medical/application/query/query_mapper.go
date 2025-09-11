package query

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/medical"
	p "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

func toResult(entity medical.MedicalHistory) MedHistoryResult {
	var notes *string
	if entityNotes := entity.Notes(); entityNotes != nil {
		notes = entityNotes
	}

	result := &MedHistoryResult{
		ID:          entity.ID(),
		PetID:       entity.PetID(),
		CustomerID:  entity.CustomerID(),
		EmployeeID:  entity.EmployeeID(),
		Date:        entity.VisitDate(),
		VisitType:   entity.VisitType().String(),
		VisitReason: entity.VisitReason().String(),
		Diagnosis:   entity.Diagnosis(),
		Treatment:   entity.Treatment(),
		Condition:   entity.Condition().String(),
		Notes:       notes,
		CreatedAt:   entity.CreatedAt(),
		UpdatedAt:   entity.UpdatedAt(),
	}

	return *result
}

func toResultList(entities []medical.MedicalHistory) []MedHistoryResult {
	dtos := make([]MedHistoryResult, len(entities))
	for i, entity := range entities {
		dtos[i] = toResult(entity)
	}
	return dtos
}

func toResultPage(page p.Page[medical.MedicalHistory]) p.Page[MedHistoryResult] {
	dtos := toResultList(page.Items)
	return p.NewPage(dtos, page.Metadata)
}
