package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type GetAppointmentByIDQuery struct {
	appointmentID valueobject.AppointmentID
	ctx           context.Context
}

func NewGetAppointmentByIDQuery(id int) (*GetAppointmentByIDQuery, error) {
	appointmentID, err := valueobject.NewAppointmentID(id)
	if err != nil {
		return nil, err
	}
	return &GetAppointmentByIDQuery{appointmentID: appointmentID}, nil
}

type GetAppointmentByIDHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewGetAppointmentByIDHandler(appointmentRepo repository.AppointmentRepository) cqrs.QueryHandler[AppointmentResponse] {
	return &GetAppointmentByIDHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *GetAppointmentByIDHandler) Handle(c cqrs.Query) (AppointmentResponse, error) {
	query := c.(GetAppointmentByIDQuery)
	appointment, err := h.appointmentRepo.GetByID(query.ctx, query.appointmentID)
	if err != nil {
		return AppointmentResponse{}, err
	}

	return NewAppointmentResponse(&appointment), nil
}
