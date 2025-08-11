package appointmentQuery

import (
	"context"
	"fmt"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
)

type GetAppointmentByIdQuery struct {
	AppointmentId int `json:"id"`
}

func NewGetAppointmentByIdQuery(id int) GetAppointmentByIdQuery {
	return GetAppointmentByIdQuery{AppointmentId: id}
}

type GetAppointmentByIdHandler interface {
	Handle(ctx context.Context, query GetAppointmentByIdQuery) (*AppointmentResponse, error)
}

type getAppointmentByIdHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewGetAppointmentByIdHandler(appointmentRepo appointmentDomain.AppointmentRepository) GetAppointmentByIdHandler {
	return &getAppointmentByIdHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *getAppointmentByIdHandler) Handle(ctx context.Context, query GetAppointmentByIdQuery) (*AppointmentResponse, error) {
	appointment, err := h.appointmentRepo.GetById(ctx, query.AppointmentId)
	if err != nil {
		return nil, err
	}

	if appointment == nil {
		return nil, fmt.Errorf("appointment with ID %d not found", query.AppointmentId)
	}

	response := NewAppointmentResponse(appointment)
	return &response, nil
}
