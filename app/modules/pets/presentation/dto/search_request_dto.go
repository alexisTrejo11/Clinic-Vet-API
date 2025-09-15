package dto

import (
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/domain/valueobject"
)

type PetSearchRequest struct {
	Name       string         `form:"name" validate:"omitempty,min=1,max=100" example:"Max"`
	Species    string         `form:"species" validate:"omitempty,min=1,max=50" example:"Perro"`
	Breed      string         `form:"breed" validate:"omitempty,min=1,max=100" example:"Labrador"`
	Age        int            `form:"age" validate:"omitempty,min=0,max=30" example:"5"`
	MinAge     int            `form:"min_age" validate:"omitempty,min=0,max=30" example:"2"`
	MaxAge     int            `form:"max_age" validate:"omitempty,min=0,max=30" example:"10"`
	Gender     enum.PetGender `form:"gender" validate:"omitempty,pet_gender" example:"MALE"`
	CustomerID uint           `form:"customer_id" examle:"12345"`
	IsActive   *bool          `form:"is_active" validate:"omitempty" example:"true"`
	IsNeutered *bool          `form:"is_neutered" validate:"omitempty" example:"true"`
	Page       int            `form:"page" validate:"omitempty,min=1" example:"1"`
	PageSize   int            `form:"page_size" validate:"omitempty,min=1,max=100" example:"10"`
	OrderBy    string         `form:"order_by" validate:"omitempty,oneof=name species breed age created_at" example:"name"`
	SortDir    string         `form:"sort_dir" validate:"omitempty,oneof=ASC DESC" example:"ASC"`
}

func (params *PetSearchRequest) ToSpecification() *specification.PetSpecification {
	builder := specification.NewPetSpecificationBuilder()

	if params.Name != "" {
		builder.Name(params.Name)
	}

	if params.Species != "" {
		builder.Species(params.Species)
	}

	if params.Breed != "" {
		builder.Breed(params.Breed)
	}

	if params.Age > 0 {
		builder.Age(params.Age)
	}

	if params.MinAge > 0 || params.MaxAge > 0 {
		minAge := params.MinAge
		maxAge := params.MaxAge
		if minAge == 0 {
			minAge = 0
		}
		if maxAge == 0 {
			maxAge = 100 // Un valor máximo razonable
		}
		builder.AgeRange(minAge, maxAge)
	}

	if params.Gender != "" {
		builder.Gender(params.Gender)
	}

	if params.CustomerID != 0 {
		customerID := valueobject.NewCustomerID(params.CustomerID)
		builder.CustomerID(customerID)
	}

	if params.IsActive != nil {
		builder.IsActive(*params.IsActive)
	}

	if params.IsNeutered != nil {
		builder.IsNeutered(*params.IsNeutered)
	}

	// Configurar paginación
	pagination := specification.Pagination{
		Page:     params.Page,
		PageSize: params.PageSize,
		OrderBy:  params.OrderBy,
		SortDir:  params.SortDir,
	}
	builder.Pagination(pagination)

	return builder.Build()
}
