package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type FindApptByIDAndEmployeeIDQuery struct {
	apptID     valueobject.AppointmentID
	employeeID valueobject.EmployeeID
	ctx        context.Context
}

func NewFindApptByIDAndEmployeeIDQuery(ctx context.Context, apptID uint, employeeID uint) *FindApptByIDAndEmployeeIDQuery {
	return &FindApptByIDAndEmployeeIDQuery{
		apptID:     valueobject.NewAppointmentID(apptID),
		employeeID: valueobject.NewEmployeeID(employeeID),
	}
}

type FindApptByIDAndEmployeeIDHandler struct {
	appointmentRepo repository.AppointmentRepository
	employeeRepo    repository.EmployeeRepository
}

func NewFindApptByIDAndEmployeeIDHandler(
	appointmentRepo repository.AppointmentRepository,
	employeeRepo repository.EmployeeRepository,
) cqrs.QueryHandler[ApptResponse] {
	return &FindApptByIDAndEmployeeIDHandler{
		appointmentRepo: appointmentRepo,
		employeeRepo:    employeeRepo,
	}
}

func (h *FindApptByIDAndEmployeeIDHandler) Handle(q cqrs.Query) (ApptResponse, error) {
	query, valid := q.(FindApptByIDAndEmployeeIDQuery)
	if !valid {
		return ApptResponse{}, errors.New("invalid query type")
	}

	if err := h.validateExistingVet(query.ctx, query.employeeID); err != nil {
		return ApptResponse{}, err
	}

	appointment, err := h.appointmentRepo.FindByIDAndEmployeeID(query.ctx, query.apptID, query.employeeID)
	if err != nil {
		return ApptResponse{}, err
	}

	return NewApptResponse(&appointment), nil
}

func (h *FindApptByIDAndEmployeeIDHandler) validateExistingVet(ctx context.Context, employeeID valueobject.EmployeeID) error {
	exists, err := h.employeeRepo.ExistsByID(ctx, employeeID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityValidationError("owner", "id", employeeID.String())
	}

	return nil
}
