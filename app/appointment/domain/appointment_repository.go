package appointDomain

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type AppointmentRepository interface {
	GetById(ctx context.Context, appointmentId int) (*Appointment, error)
	ListAll(ctx context.Context, pageInput page.PageData) (*page.Page[[]Appointment], error)
	Search(ctx context.Context, pageInput page.PageData, searchCriteria map[string]interface{}) (*page.Page[[]Appointment], error)
	ListByVetId(ctx context.Context, ownerId int, pageInput page.PageData) (*page.Page[[]Appointment], error)
	ListByPetId(ctx context.Context, ownerId int, pageInput page.PageData) (*page.Page[[]Appointment], error)
	ListByOwnerId(ctx context.Context, ownerId int, pageInput page.PageData) (*page.Page[[]Appointment], error)
	ListByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput page.PageData) (*page.Page[[]Appointment], error)

	Save(ctx context.Context, appointment *Appointment) error
	Delete(appointmentId int) error
}
