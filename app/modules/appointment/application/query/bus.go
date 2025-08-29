package query

import (
	"context"
	"fmt"
	"reflect"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type Query interface{}

type QueryHandler[T any] interface {
	Handle(query Query) (T, error)
}

type QueryBus interface {
	Execute(ctx context.Context, query Query) (interface{}, error)
	Register(queryType reflect.Type, handler interface{}) error
}

type appointmentQueryBus struct {
	handlers map[reflect.Type]interface{}
}

func NewAppointmentQueryBus(
	appointmentRepo repository.AppointmentRepository,
	ownerRepo repository.OwnerRepository,
) QueryBus {
	bus := &appointmentQueryBus{
		handlers: make(map[reflect.Type]interface{}),
	}
	bus.registerHandlers(appointmentRepo, ownerRepo)
	return bus
}

func (bus *appointmentQueryBus) registerHandlers(appointmentRepo repository.AppointmentRepository, ownerRepo repository.OwnerRepository) {
	bus.Register(reflect.TypeOf(GetAppointmentByIDQuery{}), NewGetAppointmentByIDHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(GetAllAppointmentsQuery{}), NewGetAllAppointmentsHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(GetAppointmentsByOwnerQuery{}), NewGetAppointmentsByOwnerHandler(appointmentRepo, ownerRepo))
	bus.Register(reflect.TypeOf(GetAppointmentsByVetQuery{}), NewGetAppointmentsByVetHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(GetAppointmentsByPetQuery{}), NewGetAppointmentsByPetHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(GetAppointmentsByDateRangeQuery{}), NewGetAppointmentsByDateRangeHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(GetAppointmentStatsQuery{}), NewGetAppointmentStatsHandler(appointmentRepo))
}

func (bus *appointmentQueryBus) Execute(ctx context.Context, query Query) (interface{}, error) {
	queryType := reflect.TypeOf(query)
	handler, exists := bus.handlers[queryType]

	if !exists {
		return nil, fmt.Errorf("no handler registered for query type %s", queryType.Name())
	}

	// Type switch to handle different types of response
	switch q := query.(type) {
	case GetAppointmentByIDQuery:
		h := handler.(GetAppointmentByIDHandler)
		return h.Handle(ctx, q)

	case GetAllAppointmentsQuery:
		h := handler.(GetAllAppointmentsHandler)
		return h.Handle(ctx, q)

	case GetAppointmentsByOwnerQuery:
		h := handler.(GetAppointmentsByOwnerHandler)
		return h.Handle(ctx, q)

	case GetAppointmentsByVetQuery:
		h := handler.(GetAppointmentsByVetHandler)
		return h.Handle(ctx, q)

	case GetAppointmentsByPetQuery:
		h := handler.(GetAppointmentsByPetHandler)
		return h.Handle(ctx, q)

	case GetAppointmentsByDateRangeQuery:
		h := handler.(GetAppointmentsByDateRangeHandler)
		return h.Handle(ctx, q)

	case GetAppointmentStatsQuery:
		h := handler.(GetAppointmentStatsHandler)
		return h.Handle(ctx, q)

	default:
		return nil, appError.NewConflictError("query type", fmt.Sprintf("unknown query type: %s", queryType.Name()))
	}
}

func (bus *appointmentQueryBus) Register(queryType reflect.Type, handler interface{}) error {
	if handler == nil {
		return fmt.Errorf("handler cannot be nil for query type %s", queryType.Name())
	}
	bus.handlers[queryType] = handler
	return nil
}

type AppointmentQueryService struct {
	queryBus        QueryBus
	appointmentRepo repository.AppointmentRepository
}

func NewAppointmentQueryService(queryBus QueryBus, appointmentRepo repository.AppointmentRepository) *AppointmentQueryService {
	return &AppointmentQueryService{
		queryBus:        queryBus,
		appointmentRepo: appointmentRepo,
	}
}

func (s *AppointmentQueryService) GetAppointmentByID(query GetAppointmentByIDQuery) (*AppointmentResponse, error) {
	result, err := s.queryBus.Execute(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result.(*AppointmentResponse), nil
}

func (s *AppointmentQueryService) GetAllAppointments(query GetAllAppointmentsQuery) (*AppointmentPageResponse, error) {
	result, err := s.queryBus.Execute(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result.(*AppointmentPageResponse), nil
}

func (s *AppointmentQueryService) GetAppointmentsByOwner(query GetAppointmentsByOwnerQuery) (*AppointmentPageResponse, error) {
	result, err := s.queryBus.Execute(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result.(*AppointmentPageResponse), nil
}

func (s *AppointmentQueryService) GetAppointmentsByVet(query GetAppointmentsByVetQuery) (*AppointmentPageResponse, error) {
	result, err := s.queryBus.Execute(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result.(*AppointmentPageResponse), nil
}

func (s *AppointmentQueryService) GetAppointmentsByPet(query GetAppointmentsByPetQuery) (*AppointmentPageResponse, error) {
	result, err := s.queryBus.Execute(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result.(*AppointmentPageResponse), nil
}

func (s *AppointmentQueryService) GetAppointmentsByDateRange(query GetAppointmentsByDateRangeQuery) (*AppointmentPageResponse, error) {
	result, err := s.queryBus.Execute(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result.(*AppointmentPageResponse), nil
}

func (s *AppointmentQueryService) GetAppointmentStats(query GetAppointmentStatsQuery) (*AppointmentStatsResponse, error) {
	result, err := s.queryBus.Execute(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result.(*AppointmentStatsResponse), nil
}

type AppointmentPageResponse = page.Page[[]AppointmentResponse]
