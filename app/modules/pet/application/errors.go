// Package petApplicationError defines application errors for the pets module.
package application

import (
	"strconv"

	apperror "clinic-vet-api/app/shared/error/application"
)

func PetNotFoundError(id int) error {
	return apperror.EntityNotFoundValidationError("Pet", strconv.Itoa(int(id)), "Invalid pet Id provided")
}
