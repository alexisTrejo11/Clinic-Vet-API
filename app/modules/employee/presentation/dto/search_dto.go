package dto

import (
	"fmt"
	"strconv"

	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/employee/application/cqrs/query"
	"clinic-vet-api/app/shared/page"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EmployeeSearchRequest struct {
	Name           string  `form:"name" validate:"omitempty,min=2,max=100"`
	LicenseNumber  string  `form:"license_number" validate:"omitempty,min=5,max=20"`
	Specialty      string  `form:"specialty" validate:"omitempty,vet_specialty"`
	MinExperience  *int    `form:"min_experience" validate:"omitempty,min=0,max=50"`
	MaxExperience  *int    `form:"max_experience" validate:"omitempty,min=0,max=50,gtfield=MinExperience"`
	IsActive       *bool   `form:"is_active" validate:"omitempty"`
	MinFee         float64 `form:"min_fee" validate:"omitempty,min=0,max=10000"`
	MaxFee         float64 `form:"max_fee" validate:"omitempty,min=0,max=10000,gtfield=MinFee"`
	HasUserAccount *bool   `form:"has_user_account" validate:"omitempty"`

	page.PaginationRequest
}

func NewEmployeeSearchRequestFromContext(c *gin.Context) (*EmployeeSearchRequest, error) {
	var request EmployeeSearchRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		return nil, fmt.Errorf("invalid query parameters: %w", err)
	}

	request.ProcessBooleanParams(c)

	return &request, nil
}

func (r *EmployeeSearchRequest) ProcessBooleanParams(c *gin.Context) {
	if isActiveStr := c.Query("is_active"); isActiveStr != "" {
		if isActive, err := strconv.ParseBool(isActiveStr); err == nil {
			r.IsActive = &isActive
		}
	}

	// Procesar has_user_account
	if hasUserStr := c.Query("has_user_account"); hasUserStr != "" {
		if hasUser, err := strconv.ParseBool(hasUserStr); err == nil {
			r.HasUserAccount = &hasUser
		}
	}
}

func (r *EmployeeSearchRequest) ToSpecification() (*specification.EmployeeSearchSpecification, error) {
	spec := specification.NewEmployeeSearchSpecification()

	// Name
	if r.Name != "" {
		spec = spec.WithName(r.Name)
	}

	// License Number
	if r.LicenseNumber != "" {
		spec = spec.WithLicenseNumber(r.LicenseNumber)
	}

	// Specialty
	if r.Specialty != "" {
		specialty, err := enum.ParseVetSpecialty(r.Specialty)
		if err != nil {
			return nil, fmt.Errorf("invalid specialty: %w", err)
		}
		spec = spec.WithSpecialty(specialty)
	}

	// Experience range
	if r.MinExperience != nil && r.MaxExperience != nil {
		spec = spec.WithExperienceRange(*r.MinExperience, *r.MaxExperience)
	} else if r.MinExperience != nil {
		spec = spec.WithMinExperience(*r.MinExperience)
	} else if r.MaxExperience != nil {
		spec = spec.WithMaxExperience(*r.MaxExperience)
	}

	// Active status
	if r.IsActive != nil {
		spec = spec.WithActiveStatus(*r.IsActive)
	}

	// Fee range
	if r.MinFee > 0 && r.MaxFee > 0 {

		minFee := valueobject.NewMoney(valueobject.NewDecimalFromFloat(r.MinFee), "USD")
		maxFee := valueobject.NewMoney(valueobject.NewDecimalFromFloat(r.MaxFee), "USD")
		spec = spec.WithFeeRange(minFee, maxFee)
	} else if r.MinFee > 0 {
		minFee := valueobject.NewMoney(valueobject.NewDecimalFromFloat(r.MinFee), "USD")
		spec = spec.WithMinFee(minFee)
	} else if r.MaxFee > 0 {
		maxFee := valueobject.NewMoney(valueobject.NewDecimalFromFloat(r.MaxFee), "USD")
		spec = spec.WithMaxFee(maxFee)
	}

	// User account
	if r.HasUserAccount != nil {
		spec = spec.WithUserAccount(*r.HasUserAccount)
	}

	pagination := r.WithDefaults()

	spec = spec.WithPagination(pagination.Offset(), pagination.Limit(), pagination.OrderBy, string(pagination.SortDirection))

	return spec, nil
}

func (r *EmployeeSearchRequest) ToQuery() (*query.SearchEmployeesQuery, error) {
	specification, err := r.ToSpecification()
	if err != nil {
		return nil, err
	}

	return &query.SearchEmployeesQuery{Specification: *specification}, nil
}

func RegisterEmployeeSearchValidations(validate *validator.Validate) error {
	if err := validate.RegisterValidation("vet_specialty", validateEmployeeSpecialty); err != nil {
		return fmt.Errorf("failed to register vet_specialty validation: %w", err)
	}

	return nil
}

func validateEmployeeSpecialty(fl validator.FieldLevel) bool {
	specialtyStr := fl.Field().String()
	if specialtyStr == "" {
		return true // omitempty se encarga de los vacíos
	}

	_, err := enum.ParseVetSpecialty(specialtyStr)
	return err == nil
}
