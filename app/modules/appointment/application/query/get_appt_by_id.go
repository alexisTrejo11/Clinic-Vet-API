package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type GetApptByIDQuery struct {
	appointmentID valueobject.AppointmentID
	ctx           context.Context
}

func NewGetApptByIDQuery(ctx context.Context, id uint) *GetApptByIDQuery {
	return &GetApptByIDQuery{
		appointmentID: valueobject.NewAppointmentID(id),
		ctx:           ctx,
	}
}

type GetApptByIDHandler struct {
	apptRepo repository.AppointmentRepository
}

func NewGetApptByIDHandler(apptRepo repository.AppointmentRepository) cqrs.QueryHandler[ApptResponse] {
	return &GetApptByIDHandler{
		apptRepo: apptRepo,
	}
}

func (h *GetApptByIDHandler) Handle(c cqrs.Query) (ApptResponse, error) {
	query, valid := c.(GetApptByIDQuery)
	if !valid {
		return ApptResponse{}, errors.New("invalid query type")
	}

	appointment, err := h.apptRepo.GetByID(query.ctx, query.appointmentID)
	if err != nil {
		return ApptResponse{}, err
	}

	return NewApptResponse(&appointment), nil
}
