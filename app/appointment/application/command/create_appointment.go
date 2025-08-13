package appointmentCmd

import (
	"context"
	"time"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

type CreateAppointmentCommand struct {
	OwnerId     int                                  `json:"owner_id" binding:"required"`
	PetId       int                                  `json:"pet_id" binding:"required"`
	VetId       *int                                 `json:"vet_id,omitempty"`
	Service     appointmentDomain.ClinicService      `json:"service" binding:"required"`
	DateTime    time.Time                            `json:"date_time" binding:"required"`
	Status      *appointmentDomain.AppointmentStatus `json:"status,omitempty"`
	Reason      string                               `json:"reason" binding:"required"`
	Notes       *string                              `json:"notes,omitempty"`
	IsEmergency bool                                 `json:"is_emergency"`
}

type CreateAppointmentHandler interface {
	Handle(ctx context.Context, command CreateAppointmentCommand) shared.CommandResult
}

type createAppointmentHandler struct {
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewCreateAppointmentHandler(appointmentRepo appointmentDomain.AppointmentRepository) CreateAppointmentHandler {
	return &createAppointmentHandler{
		appointmentRepo: appointmentRepo,
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

	return shared.SuccesResult(appointment.GetId().String(), "appointment created successfully")
}

func (h *createAppointmentHandler) commandToDomain(command CreateAppointmentCommand) (*appointmentDomain.Appointment, error) {
	appointmentId, err := appointmentDomain.NewAppointmentId(0)
	if err != nil {
		return nil, err
	}

	petId, err := petDomain.NewPetId(command.PetId)
	if err != nil {
		return nil, err
	}

	var vetId *vetDomain.VetId
	if command.VetId != nil {
		vet, err := vetDomain.NewVeterinarianId(*command.VetId)
		if err != nil {
			return nil, err
		}
		vetId = &vet
	}

	status := appointmentDomain.StatusPending
	if command.Status != nil {
		status = *command.Status
	}

	now := time.Now()
	appointment := appointmentDomain.NewAppointment(
		appointmentId,
		petId,
		command.OwnerId,
		vetId,
		command.Service,
		command.DateTime,
		status,
		now,
		now,
	)

	if err := appointment.ValidateRequestSchedule(); err != nil {
		return nil, err
	}

	return appointment, nil
}
