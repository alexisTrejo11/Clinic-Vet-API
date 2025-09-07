package dto

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

// AppointmentResponse represents an appointment response
type AppointmentResponse struct {
	ID            int                    `json:"id"`
	PetID         int                    `json:"pet_id"`
	OwnerID       int                    `json:"owner_id"`
	VetID         *int                   `json:"vet_id,omitempty"`
	Service       enum.ClinicService     `json:"service"`
	ScheduledDate time.Time              `json:"scheduled_date"`
	Status        enum.AppointmentStatus `json:"status"`
	Notes         *string                `json:"notes,omitempty"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}

// AppointmentDetail represents a detailed appointment with related entities
type AppointmentDetail struct {
	ID            int                    `json:"id"`
	Pet           *PetSummary            `json:"pet,omitempty"`
	Owner         *OwnerSummary          `json:"owner,omitempty"`
	Veterinarian  *VetSummary            `json:"veterinarian,omitempty"`
	Service       enum.ClinicService     `json:"service"`
	ScheduledDate time.Time              `json:"scheduled_date"`
	Status        enum.AppointmentStatus `json:"status"`
	Notes         *string                `json:"notes,omitempty"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}

// PetSummary represents a pet summary for appointment details
type PetSummary struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Species string  `json:"species"`
	Breed   *string `json:"breed,omitempty"`
}

// OwnerSummary represents an owner summary for appointment details
type OwnerSummary struct {
	ID          int    `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

// VetSummary represents a veterinarian summary for appointment details
type VetSummary struct {
	ID             int     `json:"id"`
	FullName       string  `json:"full_name"`
	Specialization *string `json:"specialization,omitempty"`
}

// AppointmentListResponse represents a paginated list of appointments
type AppointmentListResponse struct {
	Data     []AppointmentResponse `json:"data"`
	Metadata page.PageMetadata     `json:"metadata"`
}

// AppointmentStats represents appointment statistics
type AppointmentStats struct {
	TotalAppointments     int                            `json:"total_appointments"`
	AppointmentsByStatus  map[enum.AppointmentStatus]int `json:"appointments_by_status"`
	AppointmentsByService map[enum.ClinicService]int     `json:"appointments_by_service"`
	UpcomingAppointments  int                            `json:"upcoming_appointments"`
	OverdueAppointments   int                            `json:"overdue_appointments"`
}

// CreateAppointmentResponse represents the response after creating an appointment
type CreateAppointmentResponse struct {
	Appointment AppointmentResponse `json:"appointment"`
	Message     string              `json:"message"`
}

// CancelAppointmentResponse represents the response after canceling an appointment
type CancelAppointmentResponse struct {
	AppointmentID int       `json:"appointment_id"`
	Status        string    `json:"status"`
	Message       string    `json:"message"`
	CancelledAt   time.Time `json:"cancelled_at"`
}
