// Package repository contains all the operations to execute data operations
package repository

import (
	"context"
	"time"

	appoint "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	vo "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	p "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type AppointmentRepository interface {
	FindByID(ctx context.Context, id vo.AppointmentID) (appoint.Appointment, error)
	FindByIDAndCustomerID(ctx context.Context, id vo.AppointmentID, customerID vo.CustomerID) (appoint.Appointment, error)
	FindByIDAndEmployeeID(ctx context.Context, id vo.AppointmentID, employeeID vo.EmployeeID) (appoint.Appointment, error)

	FindBySpecification(ctx context.Context, spec specification.ApptSearchSpecification) (p.Page[appoint.Appointment], error)

	FindAll(ctx context.Context, pageInput p.PageInput) (p.Page[appoint.Appointment], error)
	FindByEmployeeID(ctx context.Context, employeeID vo.EmployeeID, pageInput p.PageInput) (p.Page[appoint.Appointment], error)
	FindByPetID(ctx context.Context, petID vo.PetID, pageInput p.PageInput) (p.Page[appoint.Appointment], error)
	FindByCustomerID(ctx context.Context, customerID vo.CustomerID, pageInput p.PageInput) (p.Page[appoint.Appointment], error)
	FindByDateRange(ctx context.Context, startDate, endDate time.Time, pageInput p.PageInput) (p.Page[appoint.Appointment], error)

	ExistsByID(ctx context.Context, id vo.AppointmentID) (bool, error)
	ExistsConflictingAppointment(ctx context.Context, employeeID vo.EmployeeID, startTime, endTime time.Time) (bool, error)

	Save(ctx context.Context, appointment *appoint.Appointment) error
	Update(ctx context.Context, appointment *appoint.Appointment) error
	Delete(ctx context.Context, id vo.AppointmentID) error

	CountByStatus(ctx context.Context, status enum.AppointmentStatus) (int64, error)
	CountByEmployeeID(ctx context.Context, employeeID vo.EmployeeID) (int64, error)
}
