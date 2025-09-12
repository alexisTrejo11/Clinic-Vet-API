// Package bus provides the AppointmentBus struct that encapsulates command and query buses for appointment operations.
package bus

type AppointmentBus struct {
	CommandBus AppointmentCommandBus
	QueryBus   AppointmentQueryBus
}

func NewAppointmentBus(commandBus AppointmentCommandBus, queryBus AppointmentQueryBus) *AppointmentBus {
	return &AppointmentBus{
		CommandBus: commandBus,
		QueryBus:   queryBus,
	}
}
