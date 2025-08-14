package ownerDTOs

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
)

type OwnerCreate struct {
	Photo       string              `json:"photo" validate:"required"`
	FirstName   string              `json:"first_name" validate:"required"`
	LastName    string              `json:"last_name" validate:"required"`
	Address     *string             `json:"address"`
	Gender      valueObjects.Gender `json:"gender,omitempty" validate:"required,oneof=male female not_specified"`
	DateOfBirth time.Time           `json:"date_of_birth" validate:"required"`
	PhoneNumber string              `json:"phone_number" validate:"required"`
}

type OwnerUpdate struct {
	Photo       *string              `json:"photo"`
	FirstName   *string              `json:"first_name"`
	LastName    *string              `json:"last_name"`
	Address     *string              `json:"address"`
	Notes       *string              `json:"notes"`
	Gender      *valueObjects.Gender `json:"gender" validate:"omitempty,oneof=male female not_specified"`
	DateOfBirth *time.Time           `json:"date_of_birth" validate:"required"`
	PhoneNumber *string              `json:"phone_number"`
}

type GetOwnersRequest struct {
	Page     page.PageData
	Status   string `json:"status" validate:"omitempty,oneof=active inactive all"`
	WithPets bool   `json:"with_pets" validate:"omitempty"`
}

func NewOwnerSearch(limitStr, offsetStr, status, includePets string) (*GetOwnersRequest, error) {
	var errs []error

	if limitStr == "" {
		limitStr = "10"
	}

	if offsetStr == "" {
		offsetStr = "1"
	}

	if status == "" {
		status = "active"
	}

	withPets := true
	if includePets == "" || includePets == "false" {
		withPets = false
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errs = append(errs, fmt.Errorf("invalid limit value '%s': %w", limitStr, err))
	}
	if limit < 0 {
		errs = append(errs, errors.New("limit cannot be negative"))
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		errs = append(errs, fmt.Errorf("invalid offset value '%s': %w", offsetStr, err))
	}
	if offset < 0 {
		errs = append(errs, errors.New("offset cannot be negative"))
	}

	validStatuses := map[string]bool{"active": true, "inactive": true, "pending": true} // Ejemplo
	if _, ok := validStatuses[status]; !ok && status != "" {                            // Si no es vacío y no es un estado válido
		errs = append(errs, fmt.Errorf("invalid status value '%s'", status))
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	page := page.PageData{
		PageSize:   limit,
		PageNumber: offset + 1,
	}

	return &GetOwnersRequest{
		Page:     page,
		Status:   status,
		WithPets: withPets,
	}, nil
}
