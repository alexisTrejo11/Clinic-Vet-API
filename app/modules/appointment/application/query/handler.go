package query

import (
	"context"

	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	apperror "clinic-vet-api/app/shared/error/application"
	p "clinic-vet-api/app/shared/page"
)

type AppointmentQueryHandler interface {
	FindByID(ctx context.Context, query FindApptByIDQuery) (ApptResult, error)
	FindBySpecification(ctx context.Context, query FindApptsBySpecQuery) (p.Page[ApptResult], error)
	FindByDateRange(ctx context.Context, query FindApptsByDateRangeQuery) (p.Page[ApptResult], error)
	FindByEmployee(ctx context.Context, query FindApptsByEmployeeIDQuery) (p.Page[ApptResult], error)
	FindByCustomerID(ctx context.Context, query FindApptsByCustomerIDQuery) (p.Page[ApptResult], error)
	FindByPetID(ctx context.Context, query FindApptsByPetQuery) (p.Page[ApptResult], error)
}

type apptQueryHandler struct {
	apptRepository     repository.AppointmentRepository
	customerRepository repository.CustomerRepository
	employeeRepository repository.EmployeeRepository
}

func NewAppointmentQueryHandler(apptRepository repository.AppointmentRepository, customerRepository repository.CustomerRepository, employeeRepository repository.EmployeeRepository) AppointmentQueryHandler {
	return &apptQueryHandler{
		apptRepository:     apptRepository,
		customerRepository: customerRepository,
		employeeRepository: employeeRepository,
	}
}

func (h *apptQueryHandler) FindBySpecification(ctx context.Context, query FindApptsBySpecQuery) (p.Page[ApptResult], error) {
	appointmentPage, err := h.apptRepository.Find(ctx, query.spec)
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	apptResults := mapApptsToResult(appointmentPage.Items)
	return p.NewPage(apptResults, appointmentPage.Metadata), nil
}

func (h *apptQueryHandler) FindByID(ctx context.Context, query FindApptByIDQuery) (ApptResult, error) {
	appointment, err := h.apptRepository.FindByID(ctx, query.appointmentID)
	if err != nil {
		return ApptResult{}, err
	}
	return NewApptResult(&appointment), nil
}

func (h *apptQueryHandler) FindByDateRange(ctx context.Context, query FindApptsByDateRangeQuery) (p.Page[ApptResult], error) {
	spec := specification.ApptByDateRange(query.startDate, query.endDate).
		WithPagination(query.pagination)

	appointmentsPage, err := h.apptRepository.Find(ctx, spec)
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	return p.NewPage(
		mapApptsToResult(appointmentsPage.Items),
		appointmentsPage.Metadata,
	), nil
}

func (h *apptQueryHandler) validateEmployee(ctx context.Context, employeeID valueobject.EmployeeID) error {
	exists, err := h.employeeRepository.ExistsByID(ctx, employeeID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityNotFoundValidationError("employee", "id", employeeID.String())
	}

	return nil
}

func (h *apptQueryHandler) validateCustomer(ctx context.Context, customerID valueobject.CustomerID) error {
	exists, err := h.customerRepository.ExistsByID(ctx, customerID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityNotFoundValidationError("customer", "id", customerID.String())
	}

	return nil
}

func (h *apptQueryHandler) FindByCustomerID(ctx context.Context, query FindApptsByCustomerIDQuery) (p.Page[ApptResult], error) {
	if err := h.validateCustomer(ctx, query.customerID); err != nil {
		return p.Page[ApptResult]{}, err
	}

	querySpec := specification.
		ApptByCustomer(query.customerID).
		WithPagination(query.pagination)

	if query.petID != nil {
		querySpec = querySpec.And(specification.ApptByPet(*query.petID))
	}

	appointmentsp, err := h.apptRepository.Find(ctx, querySpec)
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	return p.NewPage(mapApptsToResult(appointmentsp.Items), appointmentsp.Metadata), nil
}

func (h *apptQueryHandler) FindByEmployee(ctx context.Context, query FindApptsByEmployeeIDQuery) (p.Page[ApptResult], error) {
	querySpec := specification.
		ApptByEmployee(query.employeeID).
		WithPagination(query.pagination)

	appointmentsPage, err := h.apptRepository.Find(ctx, querySpec)
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	return p.NewPage(mapApptsToResult(appointmentsPage.Items), appointmentsPage.Metadata), nil
}

func (h *apptQueryHandler) FindByEmployeeID(ctx context.Context, query FindApptsByEmployeeIDQuery) (p.Page[ApptResult], error) {
	if err := h.validateEmployee(ctx, query.employeeID); err != nil {
		return p.Page[ApptResult]{}, err
	}

	querySpec := specification.
		ApptByEmployee(query.employeeID).
		WithPagination(query.pagination)

	appointmentsPage, err := h.apptRepository.Find(ctx, querySpec)
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	return p.NewPage(mapApptsToResult(appointmentsPage.Items), appointmentsPage.Metadata), nil
}

func (h *apptQueryHandler) FindByPetID(ctx context.Context, query FindApptsByPetQuery) (p.Page[ApptResult], error) {
	querySpec := specification.ApptByPet(query.petID).WithPagination(query.pagination)

	appointmentsPage, err := h.apptRepository.Find(ctx, querySpec)
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	return p.NewPage(mapApptsToResult(appointmentsPage.Items), appointmentsPage.Metadata), nil
}
