// Package repository contains all the operation to execute data operations
package repository

import (
	"context"
	"time"

	appoint "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	p "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type AppointmentRepository interface {
	GetByID(ctx context.Context, id valueobject.AppointmentID) (appoint.Appointment, error)
	GetByIDAndOwnerID(ctx context.Context, id valueobject.AppointmentID, ownerID valueobject.OwnerID) (appoint.Appointment, error)
	GetByIDAndVetID(ctx context.Context, id valueobject.AppointmentID, vetID valueobject.VetID) (appoint.Appointment, error)

	Search(ctx context.Context, pageInput p.PageInput, searchCriteria map[string]any) (p.Page[[]appoint.Appointment], error)

	ListAll(ctx context.Context, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)
	ListByVetID(ctx context.Context, ownerID valueobject.VetID, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)
	ListByPetID(ctx context.Context, petID valueobject.PetID, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)
	ListByOwnerID(ctx context.Context, ownerID valueobject.OwnerID, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)
	ListByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)

	Save(ctx context.Context, appointment *appoint.Appointment) error
	Delete(id valueobject.AppointmentID) error
}
