package mapper

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/medical"
	result "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/result"
	p "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

// ToDTO mapea desde entidad a DTO para queries
func ToDTO(entity *medical.MedicalHistory) result.MedHistoryResult {
	var notes *string
	if entityNotes := entity.Notes(); entityNotes != nil {
		notes = entityNotes
	}

	return result.MedHistoryResult{
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
}

// ToDTOList mapea una lista de entidades a DTOs
func (m *medicalHistoryMapper) ToDTOList(entities []*medical.MedicalHistory) []result.MedHistoryResult {
	dtos := make([]result.MedHistoryResult, len(entities))
	for i, entity := range entities {
		dtos[i] = m.ToDTO(entity)
	}
	return dtos
}

func (m *medicalHistoryMapper) ToPageDTO(page p.Page[*medical.MedicalHistory]) p.Page[result.MedHistoryResult] {
	dtos := m.ToDTOList(page.Items)
	return p.NewPage(dtos, page.Metadata)
}
