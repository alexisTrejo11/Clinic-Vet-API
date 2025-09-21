package bus

import (
	"clinic-vet-api/app/modules/medical/application/query"
	"clinic-vet-api/app/shared/cqrs"
	"context"
	"fmt"
	"reflect"

	p "clinic-vet-api/app/shared/page"
)

type QueryBus struct {
	handlers map[reflect.Type]interface{}
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[reflect.Type]any),
	}
}

func (qb *QueryBus) Register(query cqrs.Query, handler any) {
	queryType := reflect.TypeOf(query)
	qb.handlers[queryType] = handler
}

func (qb *QueryBus) Execute(query cqrs.Query) (any, error) {
	queryType := reflect.TypeOf(query)
	handler, exists := qb.handlers[queryType]
	if !exists {
		return nil, fmt.Errorf("no handler registered for query type: %T", query)
	}

	handlerValue := reflect.ValueOf(handler)
	method := handlerValue.MethodByName("Handle")
	if !method.IsValid() {
		return nil, fmt.Errorf("handler for query type %T does not have Handle method", query)
	}

	args := []reflect.Value{
		reflect.ValueOf(query),
	}

	results := method.Call(args)
	if len(results) != 2 {
		return nil, fmt.Errorf("handler method must return exactly 2 values (result, error)")
	}

	if !results[1].IsNil() {
		return nil, results[1].Interface().(error)
	}

	return results[0].Interface(), nil
}

type MedicalSessionQueryBus struct {
	*QueryBus
	handlers query.MedicalSessionQueryHandlers
}

func NewMedicalSessionQueryBus(handlers query.MedicalSessionQueryHandlers) *MedicalSessionQueryBus {
	bus := NewQueryBus()
	return &MedicalSessionQueryBus{
		QueryBus: bus,
		handlers: handlers,
	}
}

func (mq *MedicalSessionQueryBus) FindMedSessionByID(ctx context.Context, q query.FindMedSessionByIDQuery) (*query.MedSessionResult, error) {
	return mq.handlers.FindMedSessionByID(ctx, q)
}

func (mq *MedicalSessionQueryBus) FindMedSessionBySpec(ctx context.Context, q query.FindMedSessionBySpecQuery) (*p.Page[query.MedSessionResult], error) {
	return mq.handlers.FindMedSessionBySpec(ctx, q)
}

func (mq *MedicalSessionQueryBus) FindMedSessionByEmployeeID(ctx context.Context, q query.FindMedSessionByEmployeeIDQuery) (*p.Page[query.MedSessionResult], error) {
	return mq.handlers.FindMedSessionByEmployeeID(ctx, q)
}

func (mq *MedicalSessionQueryBus) FindMedSessionByPetID(ctx context.Context, q query.FindMedSessionByPetIDQuery) (*p.Page[query.MedSessionResult], error) {
	return mq.handlers.FindMedSessionByPetID(ctx, q)
}

func (mq *MedicalSessionQueryBus) FindMedSessionByCustomerID(ctx context.Context, q query.FindMedSessionByCustomerIDQuery) (*p.Page[query.MedSessionResult], error) {
	return mq.handlers.FindMedSessionByCustomerID(ctx, q)
}

func (mq *MedicalSessionQueryBus) FindRecentMedSessionByPetID(ctx context.Context, q query.FindRecentMedSessionByPetIDQuery) ([]query.MedSessionResult, error) {
	return mq.handlers.FindRecentMedSessionByPetID(ctx, q)
}

func (mq *MedicalSessionQueryBus) FindMedSessionByDateRange(ctx context.Context, q query.FindMedSessionByDateRangeQuery) (*p.Page[query.MedSessionResult], error) {
	return mq.handlers.FindMedSessionByDateRange(ctx, q)
}

func (mq *MedicalSessionQueryBus) FindMedSessionByPetAndDateRange(ctx context.Context, q query.FindMedSessionByPetAndDateRangeQuery) ([]query.MedSessionResult, error) {
	return mq.handlers.FindMedSessionByPetAndDateRange(ctx, q)
}

func (mq *MedicalSessionQueryBus) FindMedSessionByDiagnosis(ctx context.Context, q query.FindMedSessionByDiagnosisQuery) (*p.Page[query.MedSessionResult], error) {
	return mq.handlers.FindMedSessionByDiagnosis(ctx, q)
}
