package command

import (
	"clinic-vet-api/app/core/domain/entity/medical"
	"clinic-vet-api/app/core/domain/enum"
)

func ToEntityFromCreate(command *CreateMedHistCommand) (*medical.MedicalHistory, error) {
	visitType, err := enum.ParseVisitType(command.VisitType)
	if err != nil {
		return nil, err
	}

	visitReason, err := enum.ParseVisitReason(command.VisitReason)
	if err != nil {
		return nil, err
	}

	condition, err := enum.ParsePetCondition(command.Condition)
	if err != nil {
		return nil, err
	}

	entity, err := medical.CreateMedicalHistory(
		command.PetID,
		command.CustomerID,
		command.EmployeeID,
		medical.WithVisitDate(command.Date),
		medical.WithDiagnosis(command.Diagnosis),
		medical.WithVisitType(visitType),
		medical.WithVisitReason(visitReason),
		medical.WithCondition(condition),
		medical.WithTreatment(command.Treatment),
		medical.WithNotes(*command.Notes), // command.Notes es *string
	)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func ToEntityFromUpdate(command *UpdateMedHistCommand, existingEntity *medical.MedicalHistory) (*medical.MedicalHistory, error) {
	var opts []medical.MedicalHistoryOptions

	if command.Date != nil {
		opts = append(opts, medical.WithVisitDate(*command.Date))
	}

	if command.Diagnosis != nil {
		opts = append(opts, medical.WithDiagnosis(*command.Diagnosis))
	}

	if command.VisitType != nil {
		visitType, err := enum.ParseVisitType(*command.VisitType)
		if err != nil {
			return nil, err
		}
		opts = append(opts, medical.WithVisitType(visitType))
	}

	if command.VisitReason != nil {
		visitReason, err := enum.ParseVisitReason(*command.VisitReason)
		if err != nil {
			return nil, err
		}
		opts = append(opts, medical.WithVisitReason(visitReason))
	}

	if command.Condition != nil {
		condition, err := enum.ParsePetCondition(*command.Condition)
		if err != nil {
			return nil, err
		}
		opts = append(opts, medical.WithCondition(condition))
	}

	if command.Treatment != nil {
		opts = append(opts, medical.WithTreatment(*command.Treatment))
	}

	if command.Notes != nil {
		opts = append(opts, medical.WithNotes(*command.Notes))
	}

	for _, opt := range opts {
		if err := opt(existingEntity); err != nil {
			return nil, err
		}
	}

	return existingEntity, nil
}
