package bus

import (
	"clinic-vet-api/app/modules/employee/application/cqrs/command"
	"clinic-vet-api/app/modules/employee/application/cqrs/query"
)

type EmployeeBus struct {
	CommandBus command.EmployeeCommandBus
	QueryBus   query.EmployeeQueryBus
}

func NewEmployeeBus(commandBus command.EmployeeCommandBus, queryBus query.EmployeeQueryBus,
) EmployeeBus {
	return EmployeeBus{
		CommandBus: commandBus,
		QueryBus:   queryBus,
	}
}
