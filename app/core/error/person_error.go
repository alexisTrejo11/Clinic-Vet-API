package domainerr

import (
	"context"
	"fmt"
)

// PersonErrorCode representa códigos de error específicos para Person
type PersonErrorCode string

const (
	PersonNameRequired  PersonErrorCode = "PERSON_NAME_REQUIRED"
	PersonNameInvalid   PersonErrorCode = "PERSON_NAME_INVALID"
	PersonDOBRequired   PersonErrorCode = "PERSON_DOB_REQUIRED"
	PersonDOBFuture     PersonErrorCode = "PERSON_DOB_FUTURE"
	PersonUnderage      PersonErrorCode = "PERSON_UNDERAGE"
	PersonGenderInvalid PersonErrorCode = "PERSON_GENDER_INVALID"
	PersonNotFound      PersonErrorCode = "PERSON_NOT_FOUND"
	PersonAlreadyExists PersonErrorCode = "PERSON_ALREADY_EXISTS"
)

// PersonValidationError crea un error de validación específico para Person
func PersonValidationError(ctx context.Context, code PersonErrorCode, field, value, operation string) error {
	messages := map[PersonErrorCode]string{
		PersonNameRequired:  "Full name is required",
		PersonNameInvalid:   "Full name is invalid",
		PersonDOBRequired:   "Date of birth is required",
		PersonDOBFuture:     "Date of birth cannot be in the future",
		PersonUnderage:      "Person must be at least 18 years old",
		PersonGenderInvalid: "Invalid gender value",
	}

	return ValidationError(ctx, string(code), field, value,
		messages[code])
}

// PersonNotFoundError crea un error específico para persona no encontrada
func PersonNotFoundError(ctx context.Context, personID, operation string) error {
	return EntityNotFoundError(ctx, "Person", personID, operation)
}

// PersonConflictError crea un error específico para conflictos de persona
func PersonConflictError(ctx context.Context, field, value, operation string) error {
	return ConflictError(ctx, "Person",
		fmt.Sprintf("Person with %s '%s' already exists", field, value),
		operation)
}
