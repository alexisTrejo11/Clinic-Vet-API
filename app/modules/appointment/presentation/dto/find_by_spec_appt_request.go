package dto

import (
	"clinic-vet-api/app/modules/appointment/application/query"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/page"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AppointmentSearchRequest struct {
	CustomerID uint   `form:"customer_id" validate:"omitempty,min=1"`
	EmployeeID uint   `form:"vet_id" validate:"omitempty,min=1"`
	PetID      uint   `form:"pet_id" validate:"omitempty,min=1"`
	Service    string `form:"service" validate:"omitempty,clinic_service"`
	Status     string `form:"status" validate:"omitempty,appointment_status"`
	Reason     string `form:"reason" validate:"omitempty,visit_reason"`
	StartDate  string `form:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDate    string `form:"end_date" validate:"omitempty,datetime=2006-01-02"`
	HasNotes   *bool  `form:"has_notes" validate:"omitempty"`

	page.PaginationRequest
}

func (r *AppointmentSearchRequest) Pagination() map[string]any {
	return r.PaginationRequest.ToMap()
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

func (r *AppointmentSearchRequest) ToQuery() (*query.FindApptsBySpecQuery, error) {
	spec, err := r.ToSpecification()
	if err != nil {
		return nil, fmt.Errorf("failed to create specification: %w", err)
	}
	return query.NewFindApptsBySpecQuery(spec), nil
}

func (r *AppointmentSearchRequest) ToSpecification() (specification.ApptSearchSpecification, error) {
	spec := specification.NewApptSearch()

	if r.CustomerID > 0 {
		customerID := valueobject.NewCustomerID(r.CustomerID)
		spec = spec.And(specification.ApptByCustomer(customerID))
	}

	if r.EmployeeID > 0 {
		vetID := valueobject.NewEmployeeID(r.EmployeeID)
		spec = spec.And(specification.ApptByEmployee(vetID))
	}
	if r.PetID > 0 {
		petID := valueobject.NewPetID(r.PetID)
		spec = spec.And(specification.ApptByPet(petID))
	}

	if r.Service != "" {
		service, err := enum.ParseClinicService(r.Service)
		if err != nil {
			return nil, fmt.Errorf("invalid service: %w", err)
		}
		spec = spec.And(specification.ApptByService(service))
	}

	if r.Status != "" {
		status, err := enum.ParseAppointmentStatus(r.Status)
		if err != nil {
			return nil, fmt.Errorf("invalid status: %w", err)
		}
		spec = spec.And(specification.ApptByStatus(status))
	}

	startDate, endDate, err := r.parseDateRange()
	if err != nil {
		return nil, err
	}

	if startDate != nil && endDate != nil {
		spec = spec.And(specification.ApptByDateRange(*startDate, *endDate))
	}

	spec = spec.WithPagination(r.ToSpecPagination())
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
