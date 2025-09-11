package dto

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/query"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

// AppointmentCreate represents a request to create an appointment
type AppointmentCreate struct {
	PetID         int                `json:"pet_id" validate:"required,min=1"`
	VetID         *int               `json:"vet_id,omitempty"`
	OwnerID       int                `json:"owner_id" validate:"required,min=1"`
	Service       enum.ClinicService `json:"service" validate:"required"`
	ScheduledDate time.Time          `json:"scheduled_date" validate:"required"`
	Notes         *string            `json:"notes,omitempty" validate:"omitempty,max=500"`
}

// AppointmentUpdate represents a request to update an appointment
type AppointmentUpdate struct {
	VetID         *int                `json:"vet_id,omitempty"`
	Service       *enum.ClinicService `json:"service,omitempty"`
	ScheduledDate *time.Time          `json:"scheduled_date,omitempty"`
	Notes         *string             `json:"notes,omitempty" validate:"omitempty,max=500"`
}

// AppointmentOwnerUpdate represents owner-specific updates
type AppointmentOwnerUpdate struct {
	Service       *enum.ClinicService `json:"service,omitempty"`
	ScheduledDate *time.Time          `json:"scheduled_date,omitempty"`
	Notes         *string             `json:"notes,omitempty" validate:"omitempty,max=500"`
}

// AppointmentVetUpdate represents vet-specific updates
type AppointmentVetUpdate struct {
	Status enum.AppointmentStatus `json:"status" validate:"required,oneof=pending cancelled completed rescheduled not_presented"`
	Notes  *string                `json:"notes,omitempty" validate:"omitempty,max=500"`
}

// RescheduleAppointment represents a request to reschedule an appointment
type RescheduleAppointment struct {
	NewScheduledDate time.Time `json:"new_scheduled_date" validate:"required"`
	Reason           *string   `json:"reason,omitempty" validate:"omitempty,max=500"`
	AppointmentID    int       `json:"appointment_id" validate:"required,min=1"`
}

// CancelAppointment represents a request to cancel an appointment
type CancelAppointment struct {
	Reason        string `json:"reason" validate:"required,min=1,max=500"`
	AppointmentID int    `json:"appointment_id" validate:"required,min=1"`
}

// AppointmentConfirm represents a request to confirm an appointment
type AppointmentConfirm struct {
	AppointmentID int     `json:"appointment_id" validate:"required,min=1"`
	VetID         *int    `json:"vet_id,omitempty"`
	Notes         *string `json:"notes,omitempty" validate:"omitempty,max=500"`
}

// AppointmentSearch represents search criteria for appointments
type AppointmentSearch struct {
	OwnerID   *int                    `json:"owner_id,omitempty"`
	PetID     *int                    `json:"pet_id,omitempty"`
	VetID     *int                    `json:"vet_id,omitempty"`
	Status    *enum.AppointmentStatus `json:"status,omitempty"`
	Service   *enum.ClinicService     `json:"service,omitempty"`
	StartDate *time.Time              `json:"start_date,omitempty"`
	EndDate   *time.Time              `json:"end_date,omitempty"`
	Page      page.PageInput          `json:"page"`
}

type FindAppointmentsByDateRangeRequest struct {
	StartDate time.Time      `json:"start_date" validate:"required"`
	EndDate   time.Time      `json:"end_date" validate:"required"`
	Page      page.PageInput `json:"page"`
}

func (r *FindAppointmentsByDateRangeRequest) toQuery() (query.FindApptsByDateRangeQuery, error) {
	qry, err := query.NewFindApptsByDateRangeQuery(r.StartDate, r.EndDate, r.Page)
	if err != nil {
		return query.FindApptsByDateRangeQuery{}, err
	}

	return qry, nil
}

type FindAllAppointmentsRequest struct {
	Page page.PageInput `json:"page"`
}

type FindAppointmentsByVetRequest struct {
	VetID int            `json:"vet_id" validate:"required,min=1"`
	Page  page.PageInput `json:"page"`
}

type FindAppointmentsByPetRequest struct {
	PetID int            `json:"vet_id" validate:"required,min=1"`
	Page  page.PageInput `json:"page"`
}

type FindAppointmentByOwner struct {
	OwnerID int            `json:"owner_id" validate:"required,min=1"`
	Page    page.PageInput `json:"page"`
}

type FindAppointmentStatsRequest struct {
	VetID     *int       `json:"vet_id,omitempty"`
	OwnerID   *int       `json:"owner_id,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

type FindAppointmentByIDRequest struct {
	ID int `json:"id" validate:"required,min=1"`
}
