package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type FindApptByIDAndCustomerIDQuery struct {
	apptID     valueobject.AppointmentID
	customerID valueobject.CustomerID
	ctx        context.Context
}

func NewFindApptByIDAndCustomerIDQuery(ctx context.Context, apptID uint, customerID uint) *FindApptByIDAndCustomerIDQuery {
	return &FindApptByIDAndCustomerIDQuery{
		apptID:     valueobject.NewAppointmentID(apptID),
		customerID: valueobject.NewCustomerID(customerID),
	}
}

type FindApptByIDAndCustomerIDHandler struct {
	appointmentRepo repository.AppointmentRepository
	ownerRepo       repository.CustomerRepository
}

func NewFindApptByIDAndCustomerIDHandler(
	appointmentRepo repository.AppointmentRepository,
	ownerRepo repository.CustomerRepository,
) cqrs.QueryHandler[ApptResponse] {
	return &FindApptByIDAndCustomerIDHandler{
		appointmentRepo: appointmentRepo,
		ownerRepo:       ownerRepo,
	}
}

func (h *FindApptByIDAndCustomerIDHandler) Handle(q cqrs.Query) (ApptResponse, error) {
	query, valid := q.(FindApptByIDAndCustomerIDQuery)
	if !valid {
		return ApptResponse{}, errors.New("invalid query type")
	}

	if err := h.validateExistingCustomer(query.ctx, query.customerID); err != nil {
		return ApptResponse{}, err
	}

	appointment, err := h.appointmentRepo.FindByIDAndCustomerID(query.ctx, query.apptID, query.customerID)
	if err != nil {
		return ApptResponse{}, err
	}

	return NewApptResponse(&appointment), nil
}

func (h *FindApptByIDAndCustomerIDHandler) validateExistingCustomer(ctx context.Context, customerID valueobject.CustomerID) error {
	exists, err := h.ownerRepo.ExistsByID(ctx, customerID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityValidationError("owner", "id", customerID.String())
	}

	return nil
}
