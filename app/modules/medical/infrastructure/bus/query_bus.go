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

func (mq *MedicalHistoryQueryBus) GetMedHistByID(q query.GetMedHistByIDQuery) (*query.MedHistoryResult, error) {
	return mq.handlers.GetMedHistByID(q)
}

func (mq *MedicalHistoryQueryBus) GetMedHistBySpec(q query.GetMedHistBySpecQuery) (*p.Page[query.MedHistoryResult], error) {
	return mq.handlers.GetMedHistBySpec(q)
}

func (mq *MedicalHistoryQueryBus) GetAllMedHist(q query.GetAllMedHistQuery) (*p.Page[query.MedHistoryResult], error) {
	return mq.handlers.GetAllMedHist(q)
}

func (mq *MedicalHistoryQueryBus) GetMedHistByEmployeeID(q query.GetMedHistByEmployeeIDQuery) (*p.Page[query.MedHistoryResult], error) {
	return mq.handlers.GetMedHistByEmployeeID(q)
}

func (mq *MedicalHistoryQueryBus) GetMedHistByPetID(q query.GetMedHistByPetIDQuery) (*p.Page[query.MedHistoryResult], error) {
	return mq.handlers.GetMedHistByPetID(q)
}

func (mq *MedicalHistoryQueryBus) GetMedHistByCustomerID(q query.GetMedHistByCustomerIDQuery) (*p.Page[query.MedHistoryResult], error) {
	return mq.handlers.GetMedHistByCustomerID(q)
}

func (mq *MedicalHistoryQueryBus) GetRecentMedHistByPetID(q query.GetRecentMedHistByPetIDQuery) (*query.MedHistoryResult, error) {
	return mq.handlers.GetRecentMedHistByPetID(q)
}

func (mq *MedicalHistoryQueryBus) GetMedHistByDateRange(q query.GetMedHistByDateRangeQuery) (*p.Page[query.MedHistoryResult], error) {
	return mq.handlers.GetMedHistByDateRange(q)
}

func (mq *MedicalHistoryQueryBus) GetMedHistByPetAndDateRange(q query.GetMedHistByPetAndDateRangeQuery) (*query.MedHistoryResult, error) {
	return mq.handlers.GetMedHistByPetAndDateRange(q)
}

func (mq *MedicalHistoryQueryBus) GetMedHistByDiagnosis(q query.GetMedHistByDiagnosisQuery) (*p.Page[query.MedHistoryResult], error) {
	return mq.handlers.GetMedHistByDiagnosis(q)
}

func (mq *MedicalHistoryQueryBus) ExistsMedHistByID(q query.ExistsMedHistByIDQuery) (bool, error) {
	return mq.handlers.ExistsMedHistByID(q)
}

func (mq *MedicalHistoryQueryBus) ExistsMedHistByPetAndDate(q query.ExistsMedHistByPetAndDateQuery) (bool, error) {
	return mq.handlers.ExistsMedHistByPetAndDate(q)
}
