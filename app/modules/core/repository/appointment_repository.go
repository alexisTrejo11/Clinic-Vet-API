// Package repository contains all the operations to execute data operations
package repository

import (
	"context"

	appoint "clinic-vet-api/app/modules/core/domain/entity/appointment"
	"clinic-vet-api/app/modules/core/domain/specification"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	p "clinic-vet-api/app/shared/page"
)

type AppointmentRepository interface {
	FindByID(ctx context.Context, id vo.AppointmentID) (appoint.Appointment, error)
	Save(ctx context.Context, appointment *appoint.Appointment) error
	Delete(ctx context.Context, id vo.AppointmentID, hardDelete bool) error
	ExistsByID(ctx context.Context, id vo.AppointmentID) (bool, error)

	Find(ctx context.Context, spec specification.ApptSearchSpecification) (p.Page[appoint.Appointment], error)
	Count(ctx context.Context, spec specification.ApptSearchSpecification) (int64, error)
}
