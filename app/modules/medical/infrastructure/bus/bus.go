package bus

import (
	"clinic-vet-api/app/modules/medical/application/command"
	"clinic-vet-api/app/modules/medical/application/query"
)

type MedicalSessionBus struct {
	CommandBus *MedicalSessionCommandBus
	QueryBus   *MedicalSessionQueryBus
}

func NewMedicalSessionBus(commandHandlers command.MedicalSessionCommandHandlers, queryHandlers query.MedicalSessionQueryHandlers) *MedicalSessionBus {
	return &MedicalSessionBus{
		CommandBus: NewMedicalSessionCommandBus(commandHandlers),
		QueryBus:   NewMedicalSessionQueryBus(queryHandlers),
	}
}
