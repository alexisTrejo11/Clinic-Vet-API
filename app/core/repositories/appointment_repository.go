package repository

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type AppointmentRepository interface {
	GetByID(ctx context.Context, id valueobject.AppointmentID) (entity.Appointment, error)
	Search(ctx context.Context, pageInput page.PageData, searchCriteria map[string]interface{}) (page.Page[[]entity.Appointment], error)
	ListAll(ctx context.Context, pageInput page.PageData) (page.Page[[]entity.Appointment], error)
	ListByVetID(ctx context.Context, ownerID valueobject.VetID, pageInput page.PageData) (page.Page[[]entity.Appointment], error)
	ListByPetID(ctx context.Context, petID valueobject.PetID, pageInput page.PageData) (page.Page[[]entity.Appointment], error)
	ListByOwnerID(ctx context.Context, ownerID valueobject.OwnerID, pageInput page.PageData) (page.Page[[]entity.Appointment], error)
	ListByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput page.PageData) (page.Page[[]entity.Appointment], error)

	Save(ctx context.Context, appointment *entity.Appointment) error
	Delete(id valueobject.AppointmentID) error
}
