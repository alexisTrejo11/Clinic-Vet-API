package pet

import (
	"context"
	"fmt"

	"clinic-vet-api/app/core/domain/enum"
	domainerr "clinic-vet-api/app/core/error"
)

type PetErrorCode string

const (
	PetNameRequired        PetErrorCode = "PET_NAME_REQUIRED"
	PetNameTooLong         PetErrorCode = "PET_NAME_TOO_LONG"
	PetSpeciesRequired     PetErrorCode = "PET_SPECIES_REQUIRED"
	PetSpeciesTooLong      PetErrorCode = "PET_SPECIES_TOO_LONG"
	PetAgeInvalid          PetErrorCode = "PET_AGE_INVALID"
	PetAgeUnrealistic      PetErrorCode = "PET_AGE_UNREALISTIC"
	PetWeightInvalid       PetErrorCode = "PET_WEIGHT_INVALID"
	PetWeightUnrealistic   PetErrorCode = "PET_WEIGHT_UNREALISTIC"
	PetGenderInvalid       PetErrorCode = "PET_GENDER_INVALID"
	PetBreedTooLong        PetErrorCode = "PET_BREED_TOO_LONG"
	PetMicrochipTooLong    PetErrorCode = "PET_MICROCHIP_TOO_LONG"
	PetColorTooLong        PetErrorCode = "PET_COLOR_TOO_LONG"
	PetPhotoURLTooLong     PetErrorCode = "PET_PHOTO_URL_TOO_LONG"
	PetMedicationsTooLong  PetErrorCode = "PET_MEDICATIONS_TOO_LONG"
	PetAllergiesTooLong    PetErrorCode = "PET_ALLERGIES_TOO_LONG"
	PetSpecialNeedsTooLong PetErrorCode = "PET_SPECIAL_NEEDS_TOO_LONG"
	PetCustomerIDRequired  PetErrorCode = "PET_CUSTOMER_ID_REQUIRED"
)

// petValidationError crea un error de validación específico para pet
func petValidationError(ctx context.Context, code PetErrorCode, field, message, operation string) error {
	return domainerr.ValidationError(ctx, "pet", field,
		fmt.Sprintf("Pet %s: %s", field, message), operation)
}

// Errores específicos de pet
func NameRequiredError(ctx context.Context, operation string) error {
	return petValidationError(ctx, PetNameRequired, "name",
		"name is required", operation)
}

func NameTooLongError(ctx context.Context, length int, operation string) error {
	return petValidationError(ctx, PetNameTooLong, "name",
		fmt.Sprintf("name too long (%d characters, max 100)", length), operation)
}

func SpeciesRequiredError(ctx context.Context, operation string) error {
	return petValidationError(ctx, PetSpeciesRequired, "species",
		"species is required", operation)
}

func InvalidSpeciesError(ctx context.Context, enumvalue string) error {
	allowedValues := enum.ValidPetSpeciess
	return petValidationError(ctx, PetSpeciesTooLong, "species",
		fmt.Sprintf("invalid species: %s. Allowed values are: %v", enumvalue, allowedValues), "Creating/Updating Pet")
}

func AgeInvalidError(ctx context.Context, age int, operation string) error {
	return petValidationError(ctx, PetAgeInvalid, "age",
		fmt.Sprintf("age cannot be negative, got: %d", age), operation)
}

func AgeUnrealisticError(ctx context.Context, age int, operation string) error {
	return petValidationError(ctx, PetAgeUnrealistic, "age",
		fmt.Sprintf("age seems unrealistic (%d years)", age), operation)
}

func WeightInvalidError(ctx context.Context, weight float64, operation string) error {
	return petValidationError(ctx, PetWeightInvalid, "weight",
		fmt.Sprintf("weight must be positive, got: %.2f", weight), operation)
}

func WeightUnrealisticError(ctx context.Context, weight float64, operation string) error {
	return petValidationError(ctx, PetWeightUnrealistic, "weight",
		fmt.Sprintf("weight seems unrealistic (%.2f kg)", weight), operation)
}

func GenderInvalidError(ctx context.Context, gender enum.PetGender, operation string) error {
	return petValidationError(ctx, PetGenderInvalid, "gender",
		fmt.Sprintf("invalid gender: %s", gender.String()), operation)
}

func BreedTooLongError(ctx context.Context, length int, operation string) error {
	return petValidationError(ctx, PetBreedTooLong, "breed",
		fmt.Sprintf("breed too long (%d characters, max 50)", length), operation)
}

func MicrochipTooLongError(ctx context.Context, length int, operation string) error {
	return petValidationError(ctx, PetMicrochipTooLong, "microchip",
		fmt.Sprintf("microchip too long (%d characters, max 50)", length), operation)
}

func ColorTooLongError(ctx context.Context, length int, operation string) error {
	return petValidationError(ctx, PetColorTooLong, "color",
		fmt.Sprintf("color too long (%d characters, max 30)", length), operation)
}

func PhotoURLTooLongError(ctx context.Context, length int, operation string) error {
	return petValidationError(ctx, PetPhotoURLTooLong, "photo",
		fmt.Sprintf("photo URL too long (%d characters, max 500)", length), operation)
}

func MedicationsTooLongError(ctx context.Context, length int, operation string) error {
	return petValidationError(ctx, PetMedicationsTooLong, "medications",
		fmt.Sprintf("medications too long (%d characters, max 500)", length), operation)
}

func AllergiesTooLongError(ctx context.Context, length int, operation string) error {
	return petValidationError(ctx, PetAllergiesTooLong, "allergies",
		fmt.Sprintf("allergies too long (%d characters, max 500)", length), operation)
}

func SpecialNeedsTooLongError(ctx context.Context, length int, operation string) error {
	return petValidationError(ctx, PetSpecialNeedsTooLong, "special_needs",
		fmt.Sprintf("special needs too long (%d characters, max 500)", length), operation)
}

func CustomerIDRequiredError(ctx context.Context, operation string) error {
	return petValidationError(ctx, PetCustomerIDRequired, "customer_id",
		"customer ID is required", operation)
}
