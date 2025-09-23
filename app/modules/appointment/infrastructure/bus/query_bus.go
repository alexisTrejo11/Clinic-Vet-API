package bus

import (
	"clinic-vet-api/app/modules/appointment/application/query"
	p "clinic-vet-api/app/shared/page"
	"context"
)

type AppointmentQueryBus struct {
	handler query.AppointmentQueryHandler
}

func NewAppointmentQueryBus(handler query.AppointmentQueryHandler) *AppointmentQueryBus {
	return &AppointmentQueryBus{
		handler: handler,
	}
}

func (bus *AppointmentQueryBus) FindBySpecification(ctx context.Context, qry query.FindApptsBySpecQuery) (p.Page[query.ApptResult], error) {
	return bus.handler.FindBySpecification(ctx, qry)
}

func (bus *AppointmentQueryBus) FindByID(ctx context.Context, qry query.FindApptByIDQuery) (query.ApptResult, error) {
	return bus.handler.FindByID(ctx, qry)
}

func (bus *AppointmentQueryBus) FindByDateRange(ctx context.Context, qry query.FindApptsByDateRangeQuery) (p.Page[query.ApptResult], error) {
	return bus.handler.FindByDateRange(ctx, qry)
}

func (bus *AppointmentQueryBus) FindByCustomerID(ctx context.Context, qry query.FindApptsByCustomerIDQuery) (p.Page[query.ApptResult], error) {
	return bus.handler.FindByCustomerID(ctx, qry)
}

func (bus *AppointmentQueryBus) FindByEmployeeID(ctx context.Context, qry query.FindApptsByEmployeeIDQuery) (p.Page[query.ApptResult], error) {
	return bus.handler.FindByEmployee(ctx, qry)
}

func (bus *AppointmentQueryBus) FindByPetID(ctx context.Context, qry query.FindApptsByPetQuery) (p.Page[query.ApptResult], error) {
	return bus.handler.FindByPetID(ctx, qry)
}
