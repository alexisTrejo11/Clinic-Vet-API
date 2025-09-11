package dto

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/query"
	commondto "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/dto"
)

func (req *AdminMedHistoryRequest) ToCommand() *command.CreateMedHistCommand {
	return &command.CreateMedHistCommand{
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
	}
}

func (req *UpdateMedHistoryRequest) ToUpdateCommand(medHistoryID uint) *command.UpdateMedHistCommand {
	return &command.UpdateMedHistCommand{
		ID:          valueobject.NewMedHistoryID(medHistoryID),
		Diagnosis:   req.Diagnosis,
		VisitType:   req.VisitType,
		VisitReason: req.VisitReason,
		Notes:       req.Notes,
		Condition:   req.Condition,
		Treatment:   req.Treatment,
		Date:        req.Date,
	}
}

func ToResponse(result *query.MedHistoryResult) *MedHistoryResponse {
	return &MedHistoryResponse{
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

func ToDetailResponse(result *query.MedHistoryDetailResult) *MedHistoryResponseDetail {
	return &MedHistoryResponseDetail{
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

func ToResponseList(results []*query.MedHistoryResult) []*MedHistoryResponse {
	responses := make([]*MedHistoryResponse, len(results))
	for i, result := range results {
		responses[i] = ToResponse(result)
	}
	return responses
}

func ToDetailResponseList(results []*query.MedHistoryDetailResult) []*MedHistoryResponseDetail {
	responses := make([]*MedHistoryResponseDetail, len(results))
	for i, result := range results {
		responses[i] = ToDetailResponse(result)
	}
	return responses
}
