package appointDomain

import (
	"context"
	"time"
)

type AppointmentRepository interface {
	ListAll(ctx context.Context) ([]Appointment, error)
	GetById(ctx context.Context, appointmentId int) (*Appointment, error)
	Search(ctx context.Context, searchCriteria map[string]interface{}) ([]Appointment, error)
	ListByVetId(ctx context.Context, ownerId int, query string) ([]Appointment, error)
	ListByPetId(ctx context.Context, ownerId int, query string) ([]Appointment, error)
	ListByOwnerId(ctx context.Context, ownerId int) ([]Appointment, error)
	ListByDateRange(ctx context.Context, startDate, endDate time.Time) ([]Appointment, error)

	Save(ctx context.Context, appointment *Appointment) error
	Delete(appointmentId int) error
}
