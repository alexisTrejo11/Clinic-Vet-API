package specification

import (
	"fmt"
	"strings"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type ApptSearchSpecification struct {
	CustomerID *valueobject.CustomerID
	EmployeeID *valueobject.EmployeeID
	PetID      *valueobject.PetID
	Service    *enum.ClinicService
	Status     *enum.AppointmentStatus
	Reason     *enum.VisitReason
	StartDate  *time.Time
	EndDate    *time.Time
	HasNotes   *bool
	Pagination Pagination
}

func NewAppointmentSearchSpecification() *ApptSearchSpecification {
	return &ApptSearchSpecification{
		Pagination: Pagination{
			Page:     1,
			PageSize: 10,
			OrderBy:  "scheduled_date",
			SortDir:  "DESC",
		},
	}
}

func (a *ApptSearchSpecification) WithCustomerID(customerID valueobject.CustomerID) *ApptSearchSpecification {
	a.CustomerID = &customerID
	return a
}

func (a *ApptSearchSpecification) WithEmployeeID(vetID valueobject.EmployeeID) *ApptSearchSpecification {
	a.EmployeeID = &vetID
	return a
}

func (a *ApptSearchSpecification) WithPetID(petID valueobject.PetID) *ApptSearchSpecification {
	a.PetID = &petID
	return a
}

func (a *ApptSearchSpecification) WithService(service enum.ClinicService) *ApptSearchSpecification {
	a.Service = &service
	return a
}

func (a *ApptSearchSpecification) WithStatus(status enum.AppointmentStatus) *ApptSearchSpecification {
	a.Status = &status
	return a
}

func (a *ApptSearchSpecification) WithReason(reason enum.VisitReason) *ApptSearchSpecification {
	a.Reason = &reason
	return a
}

func (a *ApptSearchSpecification) WithDateRange(startDate, endDate time.Time) *ApptSearchSpecification {
	a.StartDate = &startDate
	a.EndDate = &endDate
	return a
}

func (a *ApptSearchSpecification) WithStartDate(startDate time.Time) *ApptSearchSpecification {
	a.StartDate = &startDate
	return a
}

func (a *ApptSearchSpecification) WithEndDate(endDate time.Time) *ApptSearchSpecification {
	a.EndDate = &endDate
	return a
}

func (a *ApptSearchSpecification) WithHasNotes(hasNotes bool) *ApptSearchSpecification {
	a.HasNotes = &hasNotes
	return a
}

func (a *ApptSearchSpecification) WithPagination(page, pageSize int, orderBy, sortDir string) *ApptSearchSpecification {
	a.Pagination = Pagination{
		Page:     page,
		PageSize: pageSize,
		OrderBy:  orderBy,
		SortDir:  strings.ToUpper(sortDir),
	}
	return a
}

func (a *ApptSearchSpecification) IsSatisfiedBy(candidate any) bool {
	appt, ok := candidate.(appointment.Appointment)
	if !ok {
		return false
	}

	if a.CustomerID != nil && appt.CustomerID() != *a.CustomerID {
		return false
	}

	if a.EmployeeID != nil {
		if appt.EmployeeID() == nil {
			return false
		}
		if *appt.EmployeeID() != *a.EmployeeID {
			return false
		}
	}

	if a.PetID != nil && appt.PetID() != *a.PetID {
		return false
	}

	if a.Service != nil && appt.Service() != *a.Service {
		return false
	}

	if a.Status != nil && appt.Status() != *a.Status {
		return false
	}

	if a.Reason != nil && appt.Reason() != *a.Reason {
		return false
	}

	if a.StartDate != nil && appt.ScheduledDate().Before(*a.StartDate) {
		return false
	}

	if a.EndDate != nil && appt.ScheduledDate().After(*a.EndDate) {
		return false
	}

	if a.HasNotes != nil {
		hasNotes := appt.Notes() != nil && *appt.Notes() != ""
		if hasNotes != *a.HasNotes {
			return false
		}
	}

	return true
}

func (a *ApptSearchSpecification) ToSQL() (string, []any) {
	var conditions []string
	var params []any
	paramCount := 1

	if a.CustomerID != nil {
		conditions = append(conditions, fmt.Sprintf("customer_id = $%d", paramCount))
		params = append(params, a.CustomerID.Value())
		paramCount++
	}

	if a.EmployeeID != nil {
		conditions = append(conditions, fmt.Sprintf("employee_id = $%d", paramCount))
		params = append(params, a.EmployeeID.Value())
		paramCount++
	}

	if a.PetID != nil {
		conditions = append(conditions, fmt.Sprintf("pet_id = $%d", paramCount))
		params = append(params, a.PetID.Value())
		paramCount++
	}

	if a.Service != nil {
		conditions = append(conditions, fmt.Sprintf("service = $%d", paramCount))
		params = append(params, a.Service.String())
		paramCount++
	}

	if a.Status != nil {
		conditions = append(conditions, fmt.Sprintf("status = $%d", paramCount))
		params = append(params, a.Status.String())
		paramCount++
	}

	if a.Reason != nil {
		conditions = append(conditions, fmt.Sprintf("reason = $%d", paramCount))
		params = append(params, a.Reason.String())
		paramCount++
	}

	if a.StartDate != nil && a.EndDate != nil {
		conditions = append(conditions, fmt.Sprintf("scheduled_date BETWEEN $%d AND $%d", paramCount, paramCount+1))
		params = append(params, a.StartDate, a.EndDate)
		paramCount += 2
	} else if a.StartDate != nil {
		conditions = append(conditions, fmt.Sprintf("scheduled_date >= $%d", paramCount))
		params = append(params, a.StartDate)
		paramCount++
	} else if a.EndDate != nil {
		conditions = append(conditions, fmt.Sprintf("scheduled_date <= $%d", paramCount))
		params = append(params, a.EndDate)
		paramCount++
	}

	if a.HasNotes != nil {
		if *a.HasNotes {
			conditions = append(conditions, "notes IS NOT NULL AND notes != ''")
		} else {
			conditions = append(conditions, "(notes IS NULL OR notes = '')")
		}
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ") + " AND deleted_at IS NULL"
	} else {
		whereClause = "WHERE deleted_at IS NULL"
	}

	orderBy := a.getOrderByClause()

	limitOffset := fmt.Sprintf("LIMIT $%d OFFSET $%d", paramCount, paramCount+1)
	params = append(params, a.Pagination.GetLimit(), a.Pagination.GetOffset())

	query := fmt.Sprintf(`
		SELECT id, service, scheduled_date, status, reason, notes, 
			   owner_id, vet_id, pet_id, created_at, updated_at
		FROM appointments 
		%s 
		%s 
		%s`,
		whereClause, orderBy, limitOffset)

	return query, params
}

func (a *ApptSearchSpecification) getOrderByClause() string {
	orderBy := a.Pagination.OrderBy
	sortDir := a.Pagination.SortDir

	if sortDir != "ASC" && sortDir != "DESC" {
		sortDir = "DESC"
	}

	switch orderBy {
	case "scheduled_date":
		return fmt.Sprintf("ORDER BY scheduled_date %s", sortDir)
	case "status":
		return fmt.Sprintf("ORDER BY status %s", sortDir)
	case "service":
		return fmt.Sprintf("ORDER BY service %s", sortDir)
	case "created_at", "updated_at":
		return fmt.Sprintf("ORDER BY %s %s", orderBy, sortDir)
	default:
		return "ORDER BY scheduled_date DESC"
	}
}

func (a *ApptSearchSpecification) GetPagination() Pagination {
	return a.Pagination
}
