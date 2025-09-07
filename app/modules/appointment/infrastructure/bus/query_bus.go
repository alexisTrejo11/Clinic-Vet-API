package bus

import (
	"fmt"
	"reflect"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	appointquery "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/query"
	icqrs "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	infraerr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type appointmentQueryBus struct {
	handlers map[reflect.Type]any
}

func NewAppointmentQueryBus(
	appointmentRepo repository.AppointmentRepository,
	ownerRepo repository.OwnerRepository,
) icqrs.QueryBus {
	bus := &appointmentQueryBus{
		handlers: make(map[reflect.Type]any),
	}
	bus.registerhandler(appointmentRepo, ownerRepo)
	return bus
}

func (bus *appointmentQueryBus) registerhandler(appointmentRepo repository.AppointmentRepository, ownerRepo repository.OwnerRepository) {
	bus.Register(reflect.TypeOf(appointquery.GetAppointmentByIDQuery{}), appointquery.NewGetAppointmentByIDHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointquery.SearchAppointmentsQuery{}), appointquery.NewSearchAppointmentsHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointquery.ListAppointmentsByOwnerQuery{}), appointquery.NewListAppointmentsByOwnerHandler(appointmentRepo, ownerRepo))
	bus.Register(reflect.TypeOf(appointquery.ListAppointmentsByVetQuery{}), appointquery.NewListAppointmentsByVetHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointquery.ListAppointmentsByPetQuery{}), appointquery.NewListAppointmentsByPetHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointquery.ListAppointmentsByDateRangeQuery{}), appointquery.NewListAppointmentsByDateRangeHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointquery.GetAppointmentStatsQuery{}), appointquery.NewGetAppointmentStatsHandler(appointmentRepo))
}

func (bus *appointmentQueryBus) Execute(query icqrs.Query) (any, error) {
	queryType := reflect.TypeOf(query)
	handler, exists := bus.handlers[queryType]

	if !exists {
		return nil, infraerr.NotRegistredQueryErr(queryType.Name(), "appointment")
	}

	// Type switch to handle different types of response
	switch q := query.(type) {
	case appointquery.GetAppointmentByIDQuery:
		h := handler.(appointquery.GetAppointmentByIDHandler)
		return h.Handle(q)

	case appointquery.SearchAppointmentsQuery:
		h := handler.(appointquery.SearchAppointmentsHandler)
		return h.Handle(q)

	case appointquery.ListAppointmentsByOwnerQuery:
		h := handler.(appointquery.ListAppointmentsByOwnerHandler)
		return h.Handle(q)

	case appointquery.ListAppointmentsByVetQuery:
		h := handler.(appointquery.ListAppointmentsByVetHandler)
		return h.Handle(q)

	case appointquery.ListAppointmentsByPetQuery:
		h := handler.(appointquery.ListAppointmentsByPetHandler)
		return h.Handle(q)

	case appointquery.ListAppointmentsByDateRangeQuery:
		h := handler.(appointquery.ListAppointmentsByDateRangeHandler)
		return h.Handle(q)

	case appointquery.GetAppointmentStatsQuery:
		h := handler.(appointquery.GetAppointmentStatsHandler)
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

type AppointmentQueryService struct {
	queryBus        icqrs.QueryBus
	appointmentRepo repository.AppointmentRepository
}

func NewAppointmentQueryService(queryBus icqrs.QueryBus, appointmentRepo repository.AppointmentRepository) *AppointmentQueryService {
	return &AppointmentQueryService{
		queryBus:        queryBus,
		appointmentRepo: appointmentRepo,
	}
}

func (s *AppointmentQueryService) ListAppointmentByID(query appointquery.GetAppointmentByIDQuery) (*appointquery.AppointmentResponse, error) {
	result, err := s.queryBus.Execute(query)
	if err != nil {
		return nil, err
	}
	return result.(*appointquery.AppointmentResponse), nil
}

func (s *AppointmentQueryService) SearchAppointments(query appointquery.SearchAppointmentsQuery) (*AppointmentPageResponse, error) {
	result, err := s.queryBus.Execute(query)
	if err != nil {
		return nil, err
	}
	return result.(*AppointmentPageResponse), nil
}

func (s *AppointmentQueryService) ListAppointmentsByOwner(query appointquery.ListAppointmentsByOwnerQuery) (*AppointmentPageResponse, error) {
	result, err := s.queryBus.Execute(query)
	if err != nil {
		return nil, err
	}
	return result.(*AppointmentPageResponse), nil
}

func (s *AppointmentQueryService) ListAppointmentsByVet(query appointquery.ListAppointmentsByVetQuery) (*AppointmentPageResponse, error) {
	result, err := s.queryBus.Execute(query)
	if err != nil {
		return nil, err
	}
	return result.(*AppointmentPageResponse), nil
}

func (s *AppointmentQueryService) ListAppointmentsByPet(query appointquery.ListAppointmentsByPetQuery) (*AppointmentPageResponse, error) {
	result, err := s.queryBus.Execute(query)
	if err != nil {
		return nil, err
	}
	return result.(*AppointmentPageResponse), nil
}

func (s *AppointmentQueryService) ListAppointmentsByDateRange(query appointquery.ListAppointmentsByDateRangeQuery) (*AppointmentPageResponse, error) {
	result, err := s.queryBus.Execute(query)
	if err != nil {
		return nil, err
	}
	return result.(*AppointmentPageResponse), nil
}

func (s *AppointmentQueryService) GetAppointmentStats(query appointquery.GetAppointmentStatsQuery) (*appointquery.AppointmentStatsResponse, error) {
	result, err := s.queryBus.Execute(query)
	if err != nil {
		return nil, err
	}
	return result.(*appointquery.AppointmentStatsResponse), nil
}

type AppointmentPageResponse = page.Page[[]appointquery.AppointmentResponse]
