package appointmentDTOs

import (
	"time"

	appointDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

// AppointmentResponseDTO represents an appointment response
type AppointmentResponseDTO struct {
	Id            int                             `json:"id"`
	PetId         int                             `json:"pet_id"`
	OwnerId       int                             `json:"owner_id"`
	VetId         *int                            `json:"vet_id,omitempty"`
	Service       appointDomain.ClinicService     `json:"service"`
	ScheduledDate time.Time                       `json:"scheduled_date"`
	Status        appointDomain.AppointmentStatus `json:"status"`
	Notes         *string                         `json:"notes,omitempty"`
	CreatedAt     time.Time                       `json:"created_at"`
	UpdatedAt     time.Time                       `json:"updated_at"`
}

// AppointmentDetailDTO represents a detailed appointment with related entities
type AppointmentDetailDTO struct {
	Id            int                             `json:"id"`
	Pet           *PetSummaryDTO                  `json:"pet,omitempty"`
	Owner         *OwnerSummaryDTO                `json:"owner,omitempty"`
	Veterinarian  *VetSummaryDTO                  `json:"veterinarian,omitempty"`
	Service       appointDomain.ClinicService     `json:"service"`
	ScheduledDate time.Time                       `json:"scheduled_date"`
	Status        appointDomain.AppointmentStatus `json:"status"`
	Notes         *string                         `json:"notes,omitempty"`
	CreatedAt     time.Time                       `json:"created_at"`
	UpdatedAt     time.Time                       `json:"updated_at"`
}

// PetSummaryDTO represents a pet summary for appointment details
type PetSummaryDTO struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Species string  `json:"species"`
	Breed   *string `json:"breed,omitempty"`
}

// OwnerSummaryDTO represents an owner summary for appointment details
type OwnerSummaryDTO struct {
	Id          int    `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

// VetSummaryDTO represents a veterinarian summary for appointment details
type VetSummaryDTO struct {
	Id             int     `json:"id"`
	FullName       string  `json:"full_name"`
	Specialization *string `json:"specialization,omitempty"`
}

// AppointmentListResponseDTO represents a paginated list of appointments
type AppointmentListResponseDTO struct {
	Data     []AppointmentResponseDTO `json:"data"`
	Metadata page.PageMetadata        `json:"metadata"`
}

// AppointmentStatsDTO represents appointment statistics
type AppointmentStatsDTO struct {
	TotalAppointments     int                                     `json:"total_appointments"`
	AppointmentsByStatus  map[appointDomain.AppointmentStatus]int `json:"appointments_by_status"`
	AppointmentsByService map[appointDomain.ClinicService]int     `json:"appointments_by_service"`
	UpcomingAppointments  int                                     `json:"upcoming_appointments"`
	OverdueAppointments   int                                     `json:"overdue_appointments"`
}

// CreateAppointmentResponseDTO represents the response after creating an appointment
type CreateAppointmentResponseDTO struct {
	Appointment AppointmentResponseDTO `json:"appointment"`
	Message     string                 `json:"message"`
}

// CancelAppointmentResponseDTO represents the response after canceling an appointment
type CancelAppointmentResponseDTO struct {
	AppointmentId int       `json:"appointment_id"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	CancelledAt   time.Time `json:"cancelled_at"`
}
