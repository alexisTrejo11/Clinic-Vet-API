// Package repository contains all the operation to execute data operations
package repository

import (
	"context"
	"time"

	appoint "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	vo "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	p "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type AppointmentRepository interface {
	GetByID(ctx context.Context, id vo.AppointmentID) (appoint.Appointment, error)
	GetByIDAndOwnerID(ctx context.Context, id vo.AppointmentID, ownerID vo.OwnerID) (appoint.Appointment, error)
	GetByIDAndVetID(ctx context.Context, id vo.AppointmentID, vetID vo.VetID) (appoint.Appointment, error)

	Search(ctx context.Context, spec specification.ApptSearchSpecification) (p.Page[[]appoint.Appointment], error)
	ListAll(ctx context.Context, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)
	ListByVetID(ctx context.Context, ownerID vo.VetID, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)
	ListByPetID(ctx context.Context, petID vo.PetID, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)
	ListByOwnerID(ctx context.Context, ownerID vo.OwnerID, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)
	ListByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)

	Save(ctx context.Context, appointment *appoint.Appointment) error
	Delete(id vo.AppointmentID) error
}
