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
	GetByIDAndCustomerID(ctx context.Context, id vo.AppointmentID, ownerID vo.CustomerID) (appoint.Appointment, error)
	GetByIDAndEmployeeID(ctx context.Context, id vo.AppointmentID, vetID vo.EmployeeID) (appoint.Appointment, error)

	Search(ctx context.Context, spec specification.ApptSearchSpecification) (p.Page[[]appoint.Appointment], error)
	ListAll(ctx context.Context, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)
	ListByEmployeeID(ctx context.Context, ownerID vo.EmployeeID, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)
	ListByPetID(ctx context.Context, petID vo.PetID, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)
	ListByCustomerID(ctx context.Context, ownerID vo.CustomerID, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)
	ListByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput p.PageInput) (p.Page[[]appoint.Appointment], error)

	Save(ctx context.Context, appointment *appoint.Appointment) error
	Delete(id vo.AppointmentID) error
}
