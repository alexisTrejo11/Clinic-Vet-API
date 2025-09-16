package bus

import (
	"fmt"
	"reflect"

	"clinic-vet-api/app/modules/medical/application/query"
	"clinic-vet-api/app/shared/cqrs"
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

type MedicalHistoryQueryBus struct {
	*QueryBus
	handlers query.MedicalHistoryQueryHandlers
}

func NewMedicalHistoryQueryBus(handlers query.MedicalHistoryQueryHandlers) *MedicalHistoryQueryBus {
	bus := NewQueryBus()
	return &MedicalHistoryQueryBus{
		QueryBus: bus,
		handlers: handlers,
	}
}

func (mq *MedicalHistoryQueryBus) FindMedHistByID(q query.FindMedHistByIDQuery) (*query.MedHistoryResult, error) {
	return mq.handlers.FindMedHistByID(q)
}

func (mq *MedicalHistoryQueryBus) FindMedHistBySpec(q query.FindMedHistBySpecQuery) (*p.Page[query.MedHistoryResult], error) {
	return mq.handlers.FindMedHistBySpec(q)
}

func (mq *MedicalHistoryQueryBus) FindAllMedHist(q query.FindAllMedHistQuery) (*p.Page[query.MedHistoryResult], error) {
	return mq.handlers.FindAllMedHist(q)
}

func (mq *MedicalHistoryQueryBus) FindMedHistByEmployeeID(q query.FindMedHistByEmployeeIDQuery) (*p.Page[query.MedHistoryResult], error) {
	return mq.handlers.FindMedHistByEmployeeID(q)
}

func (mq *MedicalHistoryQueryBus) FindMedHistByPetID(q query.FindMedHistByPetIDQuery) (*p.Page[query.MedHistoryResult], error) {
	return mq.handlers.FindMedHistByPetID(q)
}

func (mq *MedicalHistoryQueryBus) FindMedHistByCustomerID(q query.FindMedHistByCustomerIDQuery) (*p.Page[query.MedHistoryResult], error) {
	return mq.handlers.FindMedHistByCustomerID(q)
}

func (mq *MedicalHistoryQueryBus) FindRecentMedHistByPetID(q query.FindRecentMedHistByPetIDQuery) ([]query.MedHistoryResult, error) {
	return mq.handlers.FindRecentMedHistByPetID(q)
}

func (mq *MedicalHistoryQueryBus) FindMedHistByDateRange(q query.FindMedHistByDateRangeQuery) (*p.Page[query.MedHistoryResult], error) {
	return mq.handlers.FindMedHistByDateRange(q)
}

func (mq *MedicalHistoryQueryBus) FindMedHistByPetAndDateRange(q query.FindMedHistByPetAndDateRangeQuery) ([]query.MedHistoryResult, error) {
	return mq.handlers.FindMedHistByPetAndDateRange(q)
}

func (mq *MedicalHistoryQueryBus) FindMedHistByDiagnosis(q query.FindMedHistByDiagnosisQuery) (*p.Page[query.MedHistoryResult], error) {
	return mq.handlers.FindMedHistByDiagnosis(q)
}

func (mq *MedicalHistoryQueryBus) ExistsMedHistByID(q query.ExistsMedHistByIDQuery) (bool, error) {
	return mq.handlers.ExistsMedHistByID(q)
}

func (mq *MedicalHistoryQueryBus) ExistsMedHistByPetAndDate(q query.ExistsMedHistByPetAndDateQuery) (bool, error) {
	return mq.handlers.ExistsMedHistByPetAndDate(q)
}
