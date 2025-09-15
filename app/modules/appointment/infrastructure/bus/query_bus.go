package bus

import (
	"clinic-vet-api/app/modules/appointment/application/query"
	icqrs "clinic-vet-api/app/shared/cqrs"
	p "clinic-vet-api/app/shared/page"
)

type AppointmentQueryBus struct {
	queryBus icqrs.QueryBus
	handler  query.AppointmentQueryHandler
}

func NewAppointmentQueryBus(handler query.AppointmentQueryHandler) *AppointmentQueryBus {
	return &AppointmentQueryBus{
		handler: handler,
	}
}

func (bus *AppointmentQueryBus) FindBySpecification(qry query.FindApptsBySpecQuery) (p.Page[query.ApptResult], error) {
	return bus.handler.FindBySpecification(qry)
}

func (bus *AppointmentQueryBus) FindByID(qry query.FindApptByIDQuery) (query.ApptResult, error) {
	return bus.handler.FindByID(qry)
}

func (bus *AppointmentQueryBus) FindByIDAndCustomerID(qry query.FindApptByIDAndCustomerIDQuery) (query.ApptResult, error) {
	return bus.handler.FindByIDAndCustomerID(qry)
}

func (bus *AppointmentQueryBus) FindByIDAndEmployeeID(qry query.FindApptByIDAndEmployeeIDQuery) (query.ApptResult, error) {
	return bus.handler.FindByIDAndEmployeeID(qry)
}

func (bus *AppointmentQueryBus) FindByDateRange(qry query.FindApptsByDateRangeQuery) (p.Page[query.ApptResult], error) {
	return bus.handler.FindByDateRange(qry)
}

func (bus *AppointmentQueryBus) FindByCustomerID(qry query.FindApptsByCustomerIDQuery) (p.Page[query.ApptResult], error) {
	return bus.handler.FindByCustomerID(qry)
}

func (bus *AppointmentQueryBus) FindByEmployeeID(qry query.FindApptsByEmployeeIDQuery) (p.Page[query.ApptResult], error) {
	return bus.handler.FindByEmployeeID(qry)
}

func (bus *AppointmentQueryBus) FindByPetID(qry query.FindApptsByPetQuery) (p.Page[query.ApptResult], error) {
	return bus.handler.FindByPetID(qry)
}
