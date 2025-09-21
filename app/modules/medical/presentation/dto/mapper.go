package dto

import (
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/modules/medical/application/command"
	"clinic-vet-api/app/modules/medical/application/query"
	commondto "clinic-vet-api/app/shared/dto"
	"context"
)

func (req *AdminCreateMedSessionRequest) ToCommand(ctx context.Context) *command.CreateMedSessionCommand {
	return &command.CreateMedSessionCommand{
		PetID:       valueobject.NewPetID(req.PetID),
		CustomerID:  valueobject.NewCustomerID(req.CustomerID),
		EmployeeID:  valueobject.NewEmployeeID(req.EmployeeID),
		Date:        req.Date,
		Diagnosis:   req.Diagnosis,
		VisitType:   req.VisitType,
		VisitReason: req.VisitReason,
		Notes:       req.Notes,
		Condition:   req.Condition,
		Treatment:   req.Treatment,
		CTX:         ctx,
	}
}

func (req *UpdateMedSessionRequest) ToUpdateCommand(medSessionID uint) *command.UpdateMedSessionCommand {
	return &command.UpdateMedSessionCommand{
		ID:          valueobject.NewMedSessionID(medSessionID),
		Diagnosis:   req.Diagnosis,
		VisitType:   req.VisitType,
		VisitReason: req.VisitReason,
		Notes:       req.Notes,
		Condition:   req.Condition,
		Treatment:   req.Treatment,
		Date:        req.Date,
	}
}

func ToResponse(result *query.MedSessionResult) *MedSessionResponse {
	return &MedSessionResponse{
		ID:          result.ID.Value(),
		PetID:       result.PetID.Value(),
		CustomerID:  result.CustomerID.Value(),
		EmployeeID:  result.EmployeeID.Value(),
		Date:        result.Date,
		Diagnosis:   result.Diagnosis,
		VisitType:   result.VisitType,
		VisitReason: result.VisitReason,
		Notes:       result.Notes,
		Condition:   result.Condition,
		Treatment:   result.Treatment,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}
}

func ToDetailResponse(result *query.MedSessionDetailResult) *MedSessionResponseDetail {
	return &MedSessionResponseDetail{
		ID:        result.ID.Value(),
		Date:      result.Date,
		Diagnosis: result.Diagnosis,
		Notes:     result.Notes,
		Treatment: result.Treatment,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		Pet: commondto.PetDetails{
			ID:      result.Pet.ID.Value(),
			Name:    result.Pet.Name,
			Species: result.Pet.Species,
			Breed:   result.Pet.Breed,
			Age:     result.Pet.Age,
			Weight:  result.Pet.Weight,
		},
		Customer: commondto.CustomerDetails{
			ID:        result.Customer.ID.Value(),
			FirstName: result.Customer.FirstName,
			LastName:  result.Customer.LastName,
		},
		Employee: commondto.EmployeeDetails{
			ID:        result.Employee.ID.Value(),
			FirstName: result.Employee.FirstName,
			LastName:  result.Employee.LastName,
			Specialty: result.Employee.Specialty,
		},
	}
}

func ToResponseList(results []*query.MedSessionResult) []*MedSessionResponse {
	responses := make([]*MedSessionResponse, len(results))
	for i, result := range results {
		responses[i] = ToResponse(result)
	}
	return responses
}

func ToDetailResponseList(results []*query.MedSessionDetailResult) []*MedSessionResponseDetail {
	responses := make([]*MedSessionResponseDetail, len(results))
	for i, result := range results {
		responses[i] = ToDetailResponse(result)
	}
	return responses
}
