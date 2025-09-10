package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type GetApptByIDAndCustomerIDQuery struct {
	apptID     valueobject.AppointmentID
	customerID valueobject.CustomerID
	ctx        context.Context
}

func NewGetApptByIDAndCustomerIDQuery(ctx context.Context, apptID uint, customerID uint) *GetApptByIDAndCustomerIDQuery {
	return &GetApptByIDAndCustomerIDQuery{
		apptID:     valueobject.NewAppointmentID(apptID),
		customerID: valueobject.NewCustomerID(customerID),
	}
}

type GetApptByIDAndCustomerIDHandler struct {
	appointmentRepo repository.AppointmentRepository
	ownerRepo       repository.CustomerRepository
}

func NewGetApptByIDAndCustomerIDHandler(
	appointmentRepo repository.AppointmentRepository,
	ownerRepo repository.CustomerRepository,
) cqrs.QueryHandler[ApptResponse] {
	return &GetApptByIDAndCustomerIDHandler{
		appointmentRepo: appointmentRepo,
		ownerRepo:       ownerRepo,
	}
}

func (h *GetApptByIDAndCustomerIDHandler) Handle(q cqrs.Query) (ApptResponse, error) {
	query, valid := q.(GetApptByIDAndCustomerIDQuery)
	if !valid {
		return ApptResponse{}, errors.New("invalid query type")
	}

	if err := h.validateExistingCustomer(query.ctx, query.customerID); err != nil {
		return ApptResponse{}, err
	}

	appointment, err := h.appointmentRepo.GetByIDAndCustomerID(query.ctx, query.apptID, query.customerID)
	if err != nil {
		return ApptResponse{}, err
	}

	return NewApptResponse(&appointment), nil
}

func (h *GetApptByIDAndCustomerIDHandler) validateExistingCustomer(ctx context.Context, customerID valueobject.CustomerID) error {
	exists, err := h.ownerRepo.ExistsByID(ctx, customerID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityValidationError("owner", "id", customerID.String())
	}

	return nil
}
