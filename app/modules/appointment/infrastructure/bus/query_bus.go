package bus

import (
	h "clinic-vet-api/app/modules/appointment/application/handler"
	q "clinic-vet-api/app/modules/appointment/application/query"
	p "clinic-vet-api/app/shared/page"
	"context"
)

type ApptQueryBus struct {
	queryHandler h.ApptQueryHandler
}

func NewApptQueryBus(queryHandler h.ApptQueryHandler) *ApptQueryBus {
	return &ApptQueryBus{
		queryHandler: queryHandler,
	}
}

func (b *ApptQueryBus) FindBySpecification(ctx context.Context, qry q.FindApptsBySpecQuery) (p.Page[h.ApptResult], error) {
	return b.queryHandler.HandleBySpecification(ctx, qry)
}

func (b *ApptQueryBus) FindByID(ctx context.Context, qry q.FindApptByIDQuery) (h.ApptResult, error) {
	return b.queryHandler.HandleByID(ctx, qry)
}

func (b *ApptQueryBus) FindByDateRange(ctx context.Context, qry q.FindApptsByDateRangeQuery) (p.Page[h.ApptResult], error) {
	return b.queryHandler.HandleByDateRange(ctx, qry)
}

func (b *ApptQueryBus) FindByCustomerID(ctx context.Context, qry q.FindApptsByCustomerIDQuery) (p.Page[h.ApptResult], error) {
	return b.queryHandler.HandleByCustomerID(ctx, qry)
}

func (b *ApptQueryBus) FindByEmployeeID(ctx context.Context, qry q.FindApptsByEmployeeIDQuery) (p.Page[h.ApptResult], error) {
	return b.queryHandler.HandleByEmployeeID(ctx, qry)
}

func (b *ApptQueryBus) FindByPetID(ctx context.Context, qry q.FindApptsByPetQuery) (p.Page[h.ApptResult], error) {
	return b.queryHandler.HandleByPetID(ctx, qry)
}
