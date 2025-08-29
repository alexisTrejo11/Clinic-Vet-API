package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
)

type GetAppointmentByIDQuery struct {
	AppointmentID valueobject.AppointmentID `json:"id"`
}

func NewGetAppointmentByIDQuery(id valueobject.AppointmentID) GetAppointmentByIDQuery {
	return GetAppointmentByIDQuery{AppointmentID: id}
}

type GetAppointmentByIDHandler interface {
	Handle(ctx context.Context, query GetAppointmentByIDQuery) (*AppointmentResponse, error)
}

type getAppointmentByIDHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewGetAppointmentByIDHandler(appointmentRepo repository.AppointmentRepository) GetAppointmentByIDHandler {
	return &getAppointmentByIDHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentByIDHandler) Handle(ctx context.Context, query GetAppointmentByIDQuery) (*AppointmentResponse, error) {
	appointment, err := h.appointmentRepo.GetByID(ctx, query.AppointmentID)
	if err != nil {
		return nil, err
	}

	response := NewAppointmentResponse(&appointment)
	return &response, nil
}
