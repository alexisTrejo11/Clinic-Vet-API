package dto

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/medical/application/command"
	"clinic-vet-api/app/modules/medical/application/query"
	commondto "clinic-vet-api/app/shared/dto"
)

func (req *AdminCreateMedSessionRequest) ToCommand() *command.CreateMedSessionCommand {
	return &command.CreateMedSessionCommand{
		CustomerID: valueobject.NewCustomerID(req.CustomerID),
		EmployeeID: valueobject.NewEmployeeID(req.EmployeeID),
		VisitDate:  req.VisitDate,
		Diagnosis:  req.Diagnosis,
		VisitType:  enum.VisitType(req.VisitType),
		Service:    enum.ClinicService(req.ClinicService),
		Notes:      req.Notes,
		PetDetails: command.PetSummary{
			PetID:           valueobject.NewPetID(req.PetID),
			Diagnosis:       req.Diagnosis,
			Treatment:       req.PetDetails.Treatment,
			Weight:          float64PtrToDecimalPtr(req.PetDetails.Weight),
			HeartRate:       req.PetDetails.HeartRate,
			RespiratoryRate: req.PetDetails.RespiratoryRate,
			Temperature:     float64PtrToDecimalPtr(req.PetDetails.Temperature),
			Condition:       enum.PetCondition(req.PetDetails.Condition),
			Medications:     req.PetDetails.Medications,
			FollowUpDate:    req.PetDetails.FollowUpDate,
			Symptoms:        req.PetDetails.Symptoms,
		},
	}
}

func float64PtrToDecimalPtr(f *float64) *valueobject.Decimal {
	if f == nil {
		return nil
	}
	d := valueobject.NewDecimalFromFloat(*f)
	return &d
}

func (req *UpdateMedSessionRequest) ToUpdateCommand(medSessionID uint) *command.UpdateMedSessionCommand {
	var service *enum.ClinicService
	if req.ClinicService != nil {
		s := enum.ClinicService(*req.ClinicService)
		service = &s
	}

	var visitType *enum.VisitType
	if req.VisitType != nil {
		vt := enum.VisitType(*req.VisitType)
		visitType = &vt
	}

	var petCondition *enum.PetCondition
	if req.Condition != nil {
		pc := enum.PetCondition(*req.Condition)
		petCondition = &pc
	}

	return &command.UpdateMedSessionCommand{
		ID:        valueobject.NewMedSessionID(medSessionID),
		Diagnosis: req.Diagnosis,
		VisitType: visitType,
		Service:   service,
		Notes:     req.Notes,
		Condition: petCondition,
		Treatment: req.Treatment,
		Date:      req.Date,
	}
}

func ToResponse(result *query.MedSessionResult) *MedSessionResponse {
	return &MedSessionResponse{
		ID:         result.ID.Value(),
		EmployeeID: result.EmployeeID.Value(),
		Diagnosis:  result.Diagnosis,
		Notes:      result.Notes,
		CreatedAt:  result.CreatedAt,
		UpdatedAt:  result.UpdatedAt,
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
