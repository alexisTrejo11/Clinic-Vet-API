package customer

import (
	"context"
	"fmt"
	"time"

	"clinic-vet-api/app/modules/core/domain/valueobject"
	domainerr "clinic-vet-api/app/modules/core/error"
)

// CustomerError represents customer-specific domain errors
type CustomerError struct {
	*domainerr.BaseDomainError
}

// NewCustomerError creates a new customer error
func NewCustomerError(ctx context.Context, code, message, operation string, details map[string]string) error {
	if details == nil {
		details = make(map[string]string)
	}
	details["entity"] = "customer"

	return &CustomerError{
		BaseDomainError: &domainerr.BaseDomainError{
			Code:       code,
			Type:       "domain",
			Message:    message,
			StatusCode: 422, // Unprocessable Entity by default
			Details:    details,
		},
	}
}

// DateOfBirthRequiredError error for missing date of birth
func DateOfBirthRequiredError(ctx context.Context, operation string) error {
	return NewCustomerError(ctx,
		"CUSTOMER_DOB_REQUIRED",
		"Date of birth is required",
		operation,
		map[string]string{"field": "date_of_birth"},
	)
}

// DateOfBirthFutureError error for future date of birth
func DateOfBirthFutureError(ctx context.Context, dob time.Time, operation string) error {
	return NewCustomerError(ctx,
		"CUSTOMER_DOB_FUTURE",
		"Date of birth cannot be in the future",
		operation,
		map[string]string{
			"field": "date_of_birth",
			"value": dob.Format(time.RFC3339),
		},
	)
}

// UnderageCustomerError error for underage customers
func UnderageCustomerError(ctx context.Context, dob time.Time, operation string) error {
	age := time.Now().Year() - dob.Year()
	return NewCustomerError(ctx,
		"CUSTOMER_UNDERAGE",
		"Customer must be at least 18 years old",
		operation,
		map[string]string{
			"field":     "date_of_birth",
			"value":     dob.Format(time.RFC3339),
			"age_years": fmt.Sprintf("%d", age),
			"min_age":   "18",
		},
	)
}

// GenderRequiredError error for missing gender
func GenderRequiredError(ctx context.Context, operation string) error {
	return NewCustomerError(ctx,
		"CUSTOMER_GENDER_REQUIRED",
		"Gender is required",
		operation,
		map[string]string{"field": "gender"},
	)
}

// PhotoURLLongError error for too long photo URL
func PhotoURLLongError(ctx context.Context, urlLength int, operation string) error {
	return NewCustomerError(ctx,
		"CUSTOMER_PHOTO_URL_LONG",
		"Photo URL too long",
		operation,
		map[string]string{
			"field":          "photo_url",
			"current_length": fmt.Sprintf("%d", urlLength),
			"max_length":     "500",
		},
	)
}

// PetAlreadyExistsError error when adding duplicate pet
func PetAlreadyExistsError(ctx context.Context, petID valueobject.PetID, operation string) error {
	return NewCustomerError(ctx,
		"CUSTOMER_PET_EXISTS",
		"Pet already exists for this customer",
		operation,
		map[string]string{
			"pet_id":     petID.String(),
			"related_to": "pet",
		},
	)
}

// PetNotFoundError error when pet not found for removal
func PetNotFoundError(ctx context.Context, petID valueobject.PetID, operation string) error {
	return NewCustomerError(ctx,
		"CUSTOMER_PET_NOT_FOUND",
		"Pet not found to be removed",
		operation,
		map[string]string{
			"pet_id":     petID.String(),
			"related_to": "pet",
		},
	)
}

// CannotDeactivateWithPetsError error when trying to deactivate with active pets
func CannotDeactivateWithPetsError(ctx context.Context, activePetsCount int, operation string) error {
	return NewCustomerError(ctx,
		"CUSTOMER_CANT_DEACTIVATE",
		"Cannot deactivate customer with active pets",
		operation,
		map[string]string{
			"active_pets_count": fmt.Sprintf("%d", activePetsCount),
			"reason":            "has_active_pets",
		},
	)
}

// UserAlreadyAssociatedError error when user is already associated
func UserAlreadyAssociatedError(ctx context.Context, userID valueobject.UserID, operation string) error {
	return NewCustomerError(ctx,
		"CUSTOMER_USER_ALREADY_ASSOCIATED",
		"User is already associated with this customer",
		operation,
		map[string]string{
			"user_id": userID.String(),
		},
	)
}

// UserNotAssociatedError error when no user is associated
func UserNotAssociatedError(ctx context.Context, operation string) error {
	return NewCustomerError(ctx,
		"CUSTOMER_USER_NOT_ASSOCIATED",
		"No user is associated with this customer",
		operation,
		map[string]string{},
	)
}
