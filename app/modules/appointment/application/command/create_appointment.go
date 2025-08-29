package command

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type CreateAppointmentCommand struct {
	OwnerID     int                     `json:"owner_id" binding:"required"`
	PetID       valueobject.PetID       `json:"pet_id" binding:"required"`
	VetID       *valueobject.VetID      `json:"vet_id,omitempty"`
	Service     enum.ClinicService      `json:"service" binding:"required"`
	DateTime    time.Time               `json:"date_time" binding:"required"`
	Status      *enum.AppointmentStatus `json:"status,omitempty"`
	Reason      string                  `json:"reason" binding:"required"`
	Notes       *string                 `json:"notes,omitempty"`
	IsEmergency bool                    `json:"is_emergency"`
}

type CreateAppointmentHandler interface {
	Handle(ctx context.Context, command CreateAppointmentCommand) shared.CommandResult
}

type createAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
	service         service.AppointmentService
}

func NewCreateAppointmentHandler(appointmentRepo repository.AppointmentRepository) CreateAppointmentHandler {
	return &createAppointmentHandler{
		appointmentRepo: appointmentRepo,
		service:         service.AppointmentService{},
	}
}

func (h *createAppointmentHandler) Handle(ctx context.Context, command CreateAppointmentCommand) shared.CommandResult {
	appointment, err := h.commandToDomain(command)
	if err != nil {
		return shared.FailureResult("failed to create appointment domain", err)
	}

	if err := h.appointmentRepo.Save(ctx, appointment); err != nil {
		return shared.FailureResult("failed to save appointment", err)
	}

	return shared.SuccessResult(appointment.GetID().String(), "appointment created successfully")
}

func (h *createAppointmentHandler) commandToDomain(command CreateAppointmentCommand) (*entity.Appointment, error) {
	appointmentID, err := valueobject.NewAppointmentID(0)
	if err != nil {
		return nil, err
	}

	status := enum.StatusPending
	if command.Status != nil {
		status = *command.Status
	}

	now := time.Now()
	appointment := entity.NewAppointment(
		appointmentID,
		command.PetID,
		command.OwnerID,
		command.VetID,
		command.Service,
		command.DateTime,
		status,
		now,
		now,
	)

	if err := h.service.ValidateRequestSchedule(appointment); err != nil {
		return nil, err
	}

	return appointment, nil
}
