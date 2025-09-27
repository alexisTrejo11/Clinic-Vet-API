package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/medical"
)

func (cmd *CreateMedSessionCommand) ToEntity() medical.MedicalSession {
	petSummary := medical.NewPetSessionSummaryBuilder().
		WithPetID(cmd.PetDetails.PetID).
		WithCondition(cmd.PetDetails.Condition).
		WithTreatment(cmd.PetDetails.Treatment).
		WithFollowUpDate(cmd.PetDetails.FollowUpDate).
		WithWeight(cmd.PetDetails.Weight).
		WithDiagnosis(cmd.PetDetails.Diagnosis).
		WithHeartRate(cmd.PetDetails.HeartRate).
		WithRespiratoryRate(cmd.PetDetails.RespiratoryRate).
		WithTemperature(cmd.PetDetails.Temperature).
		WithMedications(cmd.PetDetails.Medications).
		Build()

	entity := medical.NewMedicalSessionBuilder().
		WithCustomerID(cmd.CustomerID).
		WithEmployeeID(cmd.EmployeeID).
		WithVisitDate(cmd.VisitDate).
		WithService(cmd.Service).
		WithNotes(cmd.Notes).
		WithVisitType(cmd.VisitType).
		WithPetDetails(*petSummary).
		Build()

	return *entity
}

func applyCommandUpdates(cmd UpdateMedSessionCommand, existingEntity medical.MedicalSession) *medical.MedicalSession {
	builder := medical.NewMedicalSessionBuilder().
		WithID(cmd.ID)

	if cmd.Date != nil {
		builder.WithVisitDate(*cmd.Date)
	} else {
		builder.WithVisitDate(existingEntity.VisitDate())
	}

	if cmd.VisitType != nil {
		builder.WithVisitType(*cmd.VisitType)
	} else {
		builder.WithVisitType(existingEntity.VisitType())
	}

	if cmd.Service != nil {
		builder.WithService(*cmd.Service)
	} else {
		builder.WithService(existingEntity.Service())
	}

	if cmd.Notes != nil {
		builder.WithNotes(cmd.Notes)
	} else {
		builder.WithNotes(existingEntity.Notes())
	}

	return builder.Build()
}
