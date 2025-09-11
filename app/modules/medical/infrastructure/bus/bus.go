package bus

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/query"
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
