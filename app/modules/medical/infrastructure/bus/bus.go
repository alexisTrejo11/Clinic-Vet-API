package bus

import (
	"clinic-vet-api/app/modules/medical/application/command"
	"clinic-vet-api/app/modules/medical/application/query"
)

type MedicalHistoryBus struct {
	CommandBus *MedicalHistoryCommandBus
	QueryBus   *MedicalHistoryQueryBus
}

func NewMedicalHistoryBus(commandHandlers command.MedicalHistoryCommandHandlers, queryHandlers query.MedicalHistoryQueryHandlers) *MedicalHistoryBus {
	return &MedicalHistoryBus{
		CommandBus: NewMedicalHistoryCommandBus(commandHandlers),
		QueryBus:   NewMedicalHistoryQueryBus(queryHandlers),
	}
}
