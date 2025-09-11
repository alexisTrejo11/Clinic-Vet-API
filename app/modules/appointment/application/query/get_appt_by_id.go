package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type FindApptByIDQuery struct {
	appointmentID valueobject.AppointmentID
	ctx           context.Context
}

func NewFindApptByIDQuery(ctx context.Context, id uint) *FindApptByIDQuery {
	return &FindApptByIDQuery{
		appointmentID: valueobject.NewAppointmentID(id),
		ctx:           ctx,
	}
}

type FindApptByIDHandler struct {
	apptRepo repository.AppointmentRepository
}

func NewFindApptByIDHandler(apptRepo repository.AppointmentRepository) cqrs.QueryHandler[ApptResponse] {
	return &FindApptByIDHandler{
		apptRepo: apptRepo,
	}
}

func (h *FindApptByIDHandler) Handle(c cqrs.Query) (ApptResponse, error) {
	query, valid := c.(FindApptByIDQuery)
	if !valid {
		return ApptResponse{}, errors.New("invalid query type")
	}

	appointment, err := h.apptRepo.FindByID(query.ctx, query.appointmentID)
	if err != nil {
		return ApptResponse{}, err
	}

	return NewApptResponse(&appointment), nil
}
