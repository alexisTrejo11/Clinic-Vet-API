package appointmentDTOs

import (
	"time"

	appointDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

// AppointmentCreate represents a request to create an appointment
type AppointmentCreateDTO struct {
	PetId         int                         `json:"pet_id" validate:"required,min=1"`
	VetId         *int                        `json:"vet_id,omitempty"`
	OwnerId       int                         `json:"owner_id" validate:"required,min=1"`
	Service       appointDomain.ClinicService `json:"service" validate:"required"`
	ScheduledDate time.Time                   `json:"scheduled_date" validate:"required"`
	Notes         *string                     `json:"notes,omitempty" validate:"omitempty,max=500"`
}

// AppointmentUpdateDTO represents a request to update an appointment
type AppointmentUpdateDTO struct {
	VetId         *int                         `json:"vet_id,omitempty"`
	Service       *appointDomain.ClinicService `json:"service,omitempty"`
	ScheduledDate *time.Time                   `json:"scheduled_date,omitempty"`
	Notes         *string                      `json:"notes,omitempty" validate:"omitempty,max=500"`
}

// AppointmentOwnerUpdateDTO represents owner-specific updates
type AppointmentOwnerUpdateDTO struct {
	Service       *appointDomain.ClinicService `json:"service,omitempty"`
	ScheduledDate *time.Time                   `json:"scheduled_date,omitempty"`
	Notes         *string                      `json:"notes,omitempty" validate:"omitempty,max=500"`
}

// AppointmentVetUpdateDTO represents vet-specific updates
type AppointmentVetUpdateDTO struct {
	Status appointDomain.AppointmentStatus `json:"status" validate:"required,oneof=pending cancelled completed rescheduled not_presented"`
	Notes  *string                         `json:"notes,omitempty" validate:"omitempty,max=500"`
}

// RescheduleAppointmentDTO represents a request to reschedule an appointment
type RescheduleAppointmentDTO struct {
	NewScheduledDate time.Time `json:"new_scheduled_date" validate:"required"`
	Reason           *string   `json:"reason,omitempty" validate:"omitempty,max=500"`
	AppointmentId    int       `json:"appointment_id" validate:"required,min=1"`
}

// CancelAppointmentDTO represents a request to cancel an appointment
type CancelAppointmentDTO struct {
	Reason        string `json:"reason" validate:"required,min=1,max=500"`
	AppointmentId int    `json:"appointment_id" validate:"required,min=1"`
}

// AppointmentConfirmDTO represents a request to confirm an appointment
type AppointmentConfirmDTO struct {
	AppointmentId int     `json:"appointment_id" validate:"required,min=1"`
	VetId         *int    `json:"vet_id,omitempty"`
	Notes         *string `json:"notes,omitempty" validate:"omitempty,max=500"`
}

// AppointmentSearchDTO represents search criteria for appointments
type AppointmentSearchDTO struct {
	OwnerId   *int                             `json:"owner_id,omitempty"`
	PetId     *int                             `json:"pet_id,omitempty"`
	VetId     *int                             `json:"vet_id,omitempty"`
	Status    *appointDomain.AppointmentStatus `json:"status,omitempty"`
	Service   *appointDomain.ClinicService     `json:"service,omitempty"`
	StartDate *time.Time                       `json:"start_date,omitempty"`
	EndDate   *time.Time                       `json:"end_date,omitempty"`
	Page      page.PageData                    `json:"page"`
}

type GetAppointmentsByDateRangeRequest struct {
	StartDate time.Time     `json:"start_date" validate:"required"`
	EndDate   time.Time     `json:"end_date" validate:"required"`
	Page      page.PageData `json:"page"`
}

type GetAllAppointmentsRequest struct {
	Page page.PageData `json:"page"`
}

type GetAppointmentsByVetRequest struct {
	VetId int           `json:"vet_id" validate:"required,min=1"`
	Page  page.PageData `json:"page"`
}

type GetAppointmentsByPetRequest struct {
	PetId int           `json:"vet_id" validate:"required,min=1"`
	Page  page.PageData `json:"page"`
}

type GetAppointmentByOwner struct {
	OwnerId int           `json:"owner_id" validate:"required,min=1"`
	Page    page.PageData `json:"page"`
}

type GetAppointmentStatsRequest struct {
	VetId     *int       `json:"vet_id,omitempty"`
	OwnerId   *int       `json:"owner_id,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

type GetAppointmentByIdRequest struct {
	Id int `json:"id" validate:"required,min=1"`
}
