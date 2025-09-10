package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type GetApptByIDAndEmployeeIDQuery struct {
	apptID     valueobject.AppointmentID
	employeeID valueobject.EmployeeID
	ctx        context.Context
}

func NewGetApptByIDAndEmployeeIDQuery(ctx context.Context, apptID uint, employeeID uint) *GetApptByIDAndEmployeeIDQuery {
	return &GetApptByIDAndEmployeeIDQuery{
		apptID:     valueobject.NewAppointmentID(apptID),
		employeeID: valueobject.NewEmployeeID(employeeID),
	}
}

type GetApptByIDAndEmployeeIDHandler struct {
	appointmentRepo repository.AppointmentRepository
	employeeRepo    repository.EmployeeRepository
}

func NewGetApptByIDAndEmployeeIDHandler(
	appointmentRepo repository.AppointmentRepository,
	employeeRepo repository.EmployeeRepository,
) cqrs.QueryHandler[ApptResponse] {
	return &GetApptByIDAndEmployeeIDHandler{
		appointmentRepo: appointmentRepo,
		employeeRepo:    employeeRepo,
	}
}

func (h *GetApptByIDAndEmployeeIDHandler) Handle(q cqrs.Query) (ApptResponse, error) {
	query, valid := q.(GetApptByIDAndEmployeeIDQuery)
	if !valid {
		return ApptResponse{}, errors.New("invalid query type")
	}

	if err := h.validateExistingVet(query.ctx, query.employeeID); err != nil {
		return ApptResponse{}, err
	}

	appointment, err := h.appointmentRepo.GetByIDAndEmployeeID(query.ctx, query.apptID, query.employeeID)
	if err != nil {
		return ApptResponse{}, err
	}

	return NewApptResponse(&appointment), nil
}

func (h *GetApptByIDAndEmployeeIDHandler) validateExistingVet(ctx context.Context, employeeID valueobject.EmployeeID) error {
	exists, err := h.employeeRepo.Exists(ctx, employeeID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityValidationError("owner", "id", employeeID.String())
	}

	return nil
}
