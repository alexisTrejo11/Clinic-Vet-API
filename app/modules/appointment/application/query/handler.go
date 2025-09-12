package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	p "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type AppointmentQueryHandler interface {
	FindByID(query FindApptByIDQuery) (ApptResult, error)
	FindByIDAndCustomerID(query FindApptByIDAndCustomerIDQuery) (ApptResult, error)
	FindByIDAndEmployeeID(query FindApptByIDAndEmployeeIDQuery) (ApptResult, error)

	FindBySpecification(query FindApptsBySpecQuery) (p.Page[ApptResult], error)
	FindByDateRange(query FindApptsByDateRangeQuery) (p.Page[ApptResult], error)
	FindByCustomerID(query FindApptsByCustomerIDQuery) (p.Page[ApptResult], error)
	FindByEmployeeID(query FindApptsByEmployeeIDQuery) (p.Page[ApptResult], error)
	FindByPetID(query FindApptsByPetQuery) (p.Page[ApptResult], error)
}

type apptQueryHandler struct {
	apptRepository     repository.AppointmentRepository
	customerRepository repository.CustomerRepository
	employeeRepository repository.EmployeeRepository
}

func NewAppointmentQueryHandler(apptRepository repository.AppointmentRepository) AppointmentQueryHandler {
	return &apptQueryHandler{
		apptRepository: apptRepository,
	}
}

func (h *apptQueryHandler) FindBySpecification(query FindApptsBySpecQuery) (p.Page[ApptResult], error) {
	appointmentPage, err := h.apptRepository.FindBySpecification(query.ctx, query.spec)
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	responses := mapApptsToResult(appointmentPage.Items)
	return p.NewPage(responses, appointmentPage.Metadata), nil
}

func (h *apptQueryHandler) FindByID(query FindApptByIDQuery) (ApptResult, error) {
	appointment, err := h.apptRepository.FindByID(query.ctx, query.appointmentID)
	if err != nil {
		return ApptResult{}, err
	}
	return NewApptResult(&appointment), nil
}

func (h *apptQueryHandler) FindByIDAndCustomerID(query FindApptByIDAndCustomerIDQuery) (ApptResult, error) {
	if err := h.validateCustomer(query.ctx, query.customerID); err != nil {
		return ApptResult{}, err
	}

	appointment, err := h.apptRepository.FindByIDAndCustomerID(query.ctx, query.apptID, query.customerID)
	if err != nil {
		return ApptResult{}, err
	}

	return NewApptResult(&appointment), nil
}

func (h *apptQueryHandler) FindByIDAndEmployeeID(query FindApptByIDAndEmployeeIDQuery) (ApptResult, error) {
	if err := h.validateEmployee(query.ctx, query.employeeID); err != nil {
		return ApptResult{}, err
	}

	appointment, err := h.apptRepository.FindByIDAndEmployeeID(query.ctx, query.apptID, query.employeeID)
	if err != nil {
		return ApptResult{}, err
	}

	return NewApptResult(&appointment), nil
}

func (h *apptQueryHandler) FindByDateRange(query FindApptsByDateRangeQuery) (p.Page[ApptResult], error) {
	appointmentsPage, err := h.apptRepository.FindByDateRange(query.ctx, query.startDate, query.endDate, query.pageInput)
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
		return apperror.EntityValidationError("owner", "id", employeeID.String())
	}

	return nil
}

func (h *apptQueryHandler) validateCustomer(ctx context.Context, customerID valueobject.CustomerID) error {
	exists, err := h.customerRepository.ExistsByID(ctx, customerID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityValidationError("owner", "id", customerID.String())
	}

	return nil
}

func (h *apptQueryHandler) FindByCustomerID(query FindApptsByCustomerIDQuery) (p.Page[ApptResult], error) {
	if err := h.validateCustomer(query.ctx, query.ownerID); err != nil {
		return p.Page[ApptResult]{}, err
	}

	appointmentsp, err := h.apptRepository.FindByCustomerID(query.ctx, query.ownerID, query.pageInput)
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	return p.NewPage(mapApptsToResult(appointmentsp.Items), appointmentsp.Metadata), nil
}

func (h *apptQueryHandler) FindByEmployeeID(query FindApptsByEmployeeIDQuery) (p.Page[ApptResult], error) {
	appointmentsPage, err := h.apptRepository.FindByEmployeeID(query.ctx, query.vetID, query.pageInput)
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	responses := mapApptsToResult(appointmentsPage.Items)
	return p.NewPage(responses, appointmentsPage.Metadata), nil
}

func (h *apptQueryHandler) FindByPetID(query FindApptsByPetQuery) (p.Page[ApptResult], error) {
	appointmentsPage, err := h.apptRepository.FindByPetID(query.ctx, query.petID, query.pageInput)
	if err != nil {
		return p.Page[ApptResult]{}, err
	}

	return p.NewPage(mapApptsToResult(appointmentsPage.Items), appointmentsPage.Metadata), nil
}
