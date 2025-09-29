package handler

import (
	"context"

	q "clinic-vet-api/app/modules/appointment/application/query"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	apperror "clinic-vet-api/app/shared/error/application"
	p "clinic-vet-api/app/shared/page"
)

type ApptQueryHandler struct {
	apptRepository     repository.AppointmentRepository
	customerRepository repository.CustomerRepository
	employeeRepository repository.EmployeeRepository
}

func NewAppointmentQueryHandler(
	apptRepository repository.AppointmentRepository,
	customerRepository repository.CustomerRepository,
	employeeRepository repository.EmployeeRepository,
) *ApptQueryHandler {
	return &ApptQueryHandler{
		apptRepository:     apptRepository,
		customerRepository: customerRepository,
		employeeRepository: employeeRepository,
	}
}

func (h *ApptQueryHandler) HandleBySpecification(ctx context.Context, query q.FindApptsBySpecQuery) (p.Page[ApptResult], error) {
	appointmentPage, err := h.apptRepository.Find(ctx, query.Spec())
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	return p.MapItems(appointmentPage, apptToResult), nil
}

func (h *ApptQueryHandler) HandleByID(ctx context.Context, query q.FindApptByIDQuery) (ApptResult, error) {
	appointment, err := h.apptRepository.FindByID(ctx, query.AppointmentID())
	if err != nil {
		return ApptResult{}, err
	}
	return apptToResult(appointment), nil
}

func (h *ApptQueryHandler) HandleByDateRange(ctx context.Context, query q.FindApptsByDateRangeQuery) (p.Page[ApptResult], error) {
	spec := specification.ApptByDateRange(query.StartDate(), query.EndDate()).
		WithPagination(query.Pagination())

	appointmentsPage, err := h.apptRepository.Find(ctx, spec)
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	return p.MapItems(appointmentsPage, apptToResult), nil
}

func (h *ApptQueryHandler) HandleByCustomerID(ctx context.Context, query q.FindApptsByCustomerIDQuery) (p.Page[ApptResult], error) {
	if err := h.validateCustomer(ctx, query.CustomerID()); err != nil {
		return p.Page[ApptResult]{}, err
	}

	querySpec := specification.
		ApptByCustomer(query.CustomerID()).
		WithPagination(query.Pagination())

	if query.PetID() != nil {
		querySpec = querySpec.And(specification.ApptByPet(*query.PetID()))
	}

	appointmentsp, err := h.apptRepository.Find(ctx, querySpec)
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	return p.MapItems(appointmentsp, apptToResult), nil
}

func (h *ApptQueryHandler) HandleByEmployeeID(ctx context.Context, query q.FindApptsByEmployeeIDQuery) (p.Page[ApptResult], error) {
	if err := h.validateEmployee(ctx, query.EmployeeID()); err != nil {
		return p.Page[ApptResult]{}, err
	}

	querySpec := specification.
		ApptByEmployee(query.EmployeeID()).
		WithPagination(query.Pagination())

	appointmentsPage, err := h.apptRepository.Find(ctx, querySpec)
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	return p.MapItems(appointmentsPage, apptToResult), nil
}

func (h *ApptQueryHandler) HandleByPetID(ctx context.Context, query q.FindApptsByPetQuery) (p.Page[ApptResult], error) {
	querySpec := specification.ApptByPet(query.PetID()).WithPagination(query.Pagination())

	appointmentsPage, err := h.apptRepository.Find(ctx, querySpec)
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	return p.MapItems(appointmentsPage, apptToResult), nil
}

func (h *ApptQueryHandler) validateEmployee(ctx context.Context, employeeID valueobject.EmployeeID) error {
	exists, err := h.employeeRepository.ExistsByID(ctx, employeeID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityNotFoundValidationError("employee", "id", employeeID.String())
	}

	return nil
}

func (h *ApptQueryHandler) validateCustomer(ctx context.Context, customerID valueobject.CustomerID) error {
	if exists, err := h.customerRepository.ExistsByID(ctx, customerID); err != nil {
		return err
	} else if !exists {
		return apperror.EntityNotFoundValidationError("customer", "id", customerID.String())
	}
	return nil
}
