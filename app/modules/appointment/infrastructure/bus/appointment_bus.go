// Package bus provides the AppointmentBus struct that encapsulates command and query buses for appointment operations.
package bus

type AppointmentBus struct {
	CommandBus ApptCmdBus
	QueryBus   ApptQueryBus
}

func NewAppointmentBus(commandBus ApptCmdBus, queryBus ApptQueryBus) *AppointmentBus {
	return &AppointmentBus{
		CommandBus: commandBus,
		QueryBus:   queryBus,
	}
}
