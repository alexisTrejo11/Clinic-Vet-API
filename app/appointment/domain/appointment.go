package entities

import "time"

type Appointment struct {
	ID        int
	PetID     int
	VetID     *int
	OwnerID   int
	Service   string
	Date      time.Time
	Status    string // "pending", "cancelled", "completed", "rescheduled", "no_show"
	CreatedAt time.Time
	UpdatedAt time.Time
}
