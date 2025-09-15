package dto

import (
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/modules/appointment/application/query"
	"clinic-vet-api/app/shared/page"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AppointmentSearchRequest struct {
	CustomerID uint   `form:"owner_id" validate:"omitempty,min=1"`
	EmployeeID uint   `form:"vet_id" validate:"omitempty,min=1"`
	PetID      uint   `form:"pet_id" validate:"omitempty,min=1"`
	Service    string `form:"service" validate:"omitempty,clinic_service"`
	Status     string `form:"status" validate:"omitempty,appointment_status"`
	Reason     string `form:"reason" validate:"omitempty,visit_reason"`
	StartDate  string `form:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDate    string `form:"end_date" validate:"omitempty,datetime=2006-01-02"`
	HasNotes   *bool  `form:"has_notes" validate:"omitempty"`

	page.PageInput
	/*
		Page     int    `form:"page" validate:"omitempty,min=1"`
		PageSize int    `form:"page_size" validate:"omitempty,min=1,max=100"`
		OrderBy  string `form:"order_by" validate:"omitempty,oneof=scheduled_date status service created_at updated_at"`
		SortDir  string `form:"sort_dir" validate:"omitempty,oneof=ASC DESC asc desc"`
	*/
}

func (r *AppointmentSearchRequest) Pagination() map[string]any {
	return map[string]any{
		"page":      r.Page,
		"page_size": r.PageSize,
		"order_by":  r.OrderBy,
		"sort_dir":  r.SortDirection,
	}
}

func NewApptSearchRequestFromContext(c *gin.Context) (*AppointmentSearchRequest, error) {
	var request AppointmentSearchRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		return nil, fmt.Errorf("invalid query parameters: %w", err)
	}

	request.processBooleanParams(c)

	return &request, nil
}

func (r *AppointmentSearchRequest) processBooleanParams(c *gin.Context) {
	if hasNotesStr := c.Query("has_notes"); hasNotesStr != "" {
		if hasNotes, err := strconv.ParseBool(hasNotesStr); err == nil {
			r.HasNotes = &hasNotes
		}
	}
}

func (r *AppointmentSearchRequest) ToQuery(ctx context.Context) (*query.FindApptsBySpecQuery, error) {
	spec, err := r.ToSpecification()
	if err != nil {
		return nil, fmt.Errorf("failed to create specification: %w", err)
	}
	return query.NewFindApptsBySpecQuery(ctx, *spec), nil
}

func (r *AppointmentSearchRequest) ToSpecification() (*specification.ApptSearchSpecification, error) {
	spec := specification.NewAppointmentSearchSpecification()

	if r.CustomerID > 0 {
		ownerID := valueobject.NewCustomerID(r.CustomerID)
		spec = spec.WithCustomerID(ownerID)
	}

	// Vet ID
	if r.EmployeeID > 0 {
		vetID := valueobject.NewEmployeeID(r.EmployeeID)
		spec = spec.WithEmployeeID(vetID)
	}

	// Pet ID
	if r.PetID > 0 {
		petID := valueobject.NewPetID(r.PetID)
		spec = spec.WithPetID(petID)
	}

	// Service
	if r.Service != "" {
		service, err := enum.ParseClinicService(r.Service)
		if err != nil {
			return nil, fmt.Errorf("invalid service: %w", err)
		}
		spec = spec.WithService(service)
	}

	// Status
	if r.Status != "" {
		status, err := enum.ParseAppointmentStatus(r.Status)
		if err != nil {
			return nil, fmt.Errorf("invalid status: %w", err)
		}
		spec = spec.WithStatus(status)
	}

	// Reason
	if r.Reason != "" {
		reason, err := enum.ParseVisitReason(r.Reason)
		if err != nil {
			return nil, fmt.Errorf("invalid reason: %w", err)
		}
		spec = spec.WithReason(reason)
	}

	// Date range
	startDate, endDate, err := r.parseDateRange()
	if err != nil {
		return nil, err
	}

	if startDate != nil && endDate != nil {
		spec = spec.WithDateRange(*startDate, *endDate)
	} else if startDate != nil {
		spec = spec.WithStartDate(*startDate)
	} else if endDate != nil {
		spec = spec.WithEndDate(*endDate)
	}

	// Has notes
	if r.HasNotes != nil {
		spec = spec.WithHasNotes(*r.HasNotes)
	}

	// Paginación con valores por defecto
	page := 1
	if r.Page > 0 {
		page = r.Page
	}

	pageSize := 10
	if r.PageSize > 0 {
		pageSize = r.PageSize
	}

	orderBy := "scheduled_date"
	if r.OrderBy != "" {
		orderBy = r.OrderBy
	}

	sortDir := "DESC"
	if r.SortDirection != "" {
		sortDir = strings.ToUpper(string(r.SortDirection))
	}
	spec = spec.WithPagination(page, pageSize, orderBy, sortDir)

	return spec, nil
}

func (r *AppointmentSearchRequest) parseDateRange() (*time.Time, *time.Time, error) {
	var startDate, endDate *time.Time

	if r.StartDate != "" {
		parsedStart, err := time.Parse("2006-01-02", r.StartDate)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid start_date format: %w", err)
		}
		startDate = &parsedStart
	}

	if r.EndDate != "" {
		parsedEnd, err := time.Parse("2006-01-02", r.EndDate)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid end_date format: %w", err)
		}
		// Ajustar end_date para incluir todo el día
		endOfDay := time.Date(parsedEnd.Year(), parsedEnd.Month(), parsedEnd.Day(), 23, 59, 59, 0, parsedEnd.Location())
		endDate = &endOfDay
	}

	return startDate, endDate, nil
}

func RegisterAppointmentSearchValidations(validate *validator.Validate) error {
	if err := validate.RegisterValidation("clinic_service", validateClinicService); err != nil {
		return fmt.Errorf("failed to register clinic_service validation: %w", err)
	}

	if err := validate.RegisterValidation("appointment_status", validateAppointmentStatus); err != nil {
		return fmt.Errorf("failed to register appointment_status validation: %w", err)
	}

	if err := validate.RegisterValidation("visit_reason", validateVisitReason); err != nil {
		return fmt.Errorf("failed to register visit_reason validation: %w", err)
	}

	return nil
}

func validateClinicService(fl validator.FieldLevel) bool {
	serviceStr := fl.Field().String()
	if serviceStr == "" {
		return true // omitempty se encarga de los vacíos
	}

	_, err := enum.ParseClinicService(serviceStr)
	return err == nil
}

func validateAppointmentStatus(fl validator.FieldLevel) bool {
	statusStr := fl.Field().String()
	if statusStr == "" {
		return true
	}

	_, err := enum.ParseAppointmentStatus(statusStr)
	return err == nil
}

func validateVisitReason(fl validator.FieldLevel) bool {
	reasonStr := fl.Field().String()
	if reasonStr == "" {
		return true
	}

	_, err := enum.ParseVisitReason(reasonStr)
	return err == nil
}
