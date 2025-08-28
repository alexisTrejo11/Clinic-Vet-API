package repository

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type AppointmentRepository interface {
	GetById(ctx context.Context, id valueobject.AppointmentID) (entity.Appointment, error)
	Search(ctx context.Context, pageInput page.PageData, searchCriteria map[string]interface{}) (page.Page[[]entity.Appointment], error)
	ListAll(ctx context.Context, pageInput page.PageData) (page.Page[[]entity.Appointment], error)
	ListByVetId(ctx context.Context, ownerId int, pageInput page.PageData) (page.Page[[]entity.Appointment], error)
	ListByPetId(ctx context.Context, ownerId int, pageInput page.PageData) (page.Page[[]entity.Appointment], error)
	ListByOwnerId(ctx context.Context, ownerId int, pageInput page.PageData) (page.Page[[]entity.Appointment], error)
	ListByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput page.PageData) (page.Page[[]entity.Appointment], error)

	Save(ctx context.Context, appointment *entity.Appointment) error
	Delete(appointmentId int) error
}
