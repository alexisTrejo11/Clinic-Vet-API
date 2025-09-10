package bus

import (
	"fmt"
	"reflect"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	appointquery "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/query"
	icqrs "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	infraerr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type appointmentQueryBus struct {
	handlers map[reflect.Type]any
}

func NewApptQueryBus(
	appointmentRepo repository.AppointmentRepository,
	ownerRepo repository.CustomerRepository,
) icqrs.QueryBus {
	bus := &appointmentQueryBus{
		handlers: make(map[reflect.Type]any),
	}
	bus.registerhandler(appointmentRepo, ownerRepo)
	return bus
}

func (bus *appointmentQueryBus) registerhandler(appointmentRepo repository.AppointmentRepository, ownerRepo repository.CustomerRepository) {
	bus.Register(reflect.TypeOf(appointquery.GetApptByIDQuery{}), appointquery.NewGetApptByIDHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointquery.SearchApptsQuery{}), appointquery.NewSearchApptsHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointquery.ListApptsByCustomerIDQuery{}), appointquery.NewListApptsByCustomerIDHandler(appointmentRepo, ownerRepo))
	bus.Register(reflect.TypeOf(appointquery.ListApptsByEmployeeIDQuery{}), appointquery.NewListApptsByEmployeeIDHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointquery.ListApptsByPetQuery{}), appointquery.NewListApptsByPetHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointquery.ListApptsByDateRangeQuery{}), appointquery.NewListApptsByDateRangeHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointquery.GetApptStatsQuery{}), appointquery.NewGetApptStatsHandler(appointmentRepo))
}

func (bus *appointmentQueryBus) Execute(query icqrs.Query) (any, error) {
	queryType := reflect.TypeOf(query)
	handler, exists := bus.handlers[queryType]

	if !exists {
		return nil, infraerr.NotRegistredQueryErr(queryType.Name(), "appointment")
	}

	// Type switch to handle different types of response
	switch q := query.(type) {
	case appointquery.GetApptByIDQuery:
		h := handler.(appointquery.GetApptByIDHandler)
		return h.Handle(q)

	case appointquery.SearchApptsQuery:
		h := handler.(appointquery.SearchApptsHandler)
		return h.Handle(q)

	case appointquery.ListApptsByCustomerIDQuery:
		h := handler.(appointquery.ListApptsByCustomerIDHandler)
		return h.Handle(q)

	case appointquery.ListApptsByEmployeeIDQuery:
		h := handler.(appointquery.ListApptsByEmployeeIDHandler)
		return h.Handle(q)

	case appointquery.ListApptsByPetQuery:
		h := handler.(appointquery.ListApptsByPetHandler)
		return h.Handle(q)

	case appointquery.ListApptsByDateRangeQuery:
		h := handler.(appointquery.ListApptsByDateRangeHandler)
		return h.Handle(q)

	case appointquery.GetApptStatsQuery:
		h := handler.(appointquery.GetApptStatsHandler)
		return h.Handle(q)

	default:
		return nil, infraerr.InvalidQuerryErr("unknow query operation", queryType.Name(), "appointment")
	}
}

func (bus *appointmentQueryBus) Register(queryType reflect.Type, handler any) error {
	if handler == nil {
		return fmt.Errorf("handler cannot be nil for query type %s", queryType.Name())
	}
	bus.handlers[queryType] = handler
	return nil
}

type ApptQueryService struct {
	queryBus        icqrs.QueryBus
	appointmentRepo repository.AppointmentRepository
}

func NewApptQueryService(queryBus icqrs.QueryBus, appointmentRepo repository.AppointmentRepository) *ApptQueryService {
	return &ApptQueryService{
		queryBus:        queryBus,
		appointmentRepo: appointmentRepo,
	}
}

func (s *ApptQueryService) ListApptByID(query appointquery.GetApptByIDQuery) (*appointquery.ApptResponse, error) {
	result, err := s.queryBus.Execute(query)
	if err != nil {
		return nil, err
	}
	return result.(*appointquery.ApptResponse), nil
}

func (s *ApptQueryService) SearchAppts(query appointquery.SearchApptsQuery) (*ApptPageResponse, error) {
	result, err := s.queryBus.Execute(query)
	if err != nil {
		return nil, err
	}
	return result.(*ApptPageResponse), nil
}

func (s *ApptQueryService) ListApptsByCustomer(query appointquery.ListApptsByCustomerIDQuery) (*ApptPageResponse, error) {
	result, err := s.queryBus.Execute(query)
	if err != nil {
		return nil, err
	}
	return result.(*ApptPageResponse), nil
}

func (s *ApptQueryService) ListApptsByEmployee(query appointquery.ListApptsByEmployeeIDHandler) (*ApptPageResponse, error) {
	result, err := s.queryBus.Execute(query)
	if err != nil {
		return nil, err
	}
	return result.(*ApptPageResponse), nil
}

func (s *ApptQueryService) ListApptsByPet(query appointquery.ListApptsByPetQuery) (*ApptPageResponse, error) {
	result, err := s.queryBus.Execute(query)
	if err != nil {
		return nil, err
	}
	return result.(*ApptPageResponse), nil
}

func (s *ApptQueryService) ListApptsByDateRange(query appointquery.ListApptsByDateRangeQuery) (*ApptPageResponse, error) {
	result, err := s.queryBus.Execute(query)
	if err != nil {
		return nil, err
	}
	return result.(*ApptPageResponse), nil
}

func (s *ApptQueryService) GetApptStats(query appointquery.GetApptStatsQuery) (*appointquery.ApptStatsResponse, error) {
	result, err := s.queryBus.Execute(query)
	if err != nil {
		return nil, err
	}
	return result.(*appointquery.ApptStatsResponse), nil
}

type ApptPageResponse = page.Page[[]appointquery.ApptResponse]
