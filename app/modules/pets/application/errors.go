// Package petApplicationError defines application errors for the pets module.
package petApplicationError

import (
	"strconv"

	apperror "clinic-vet-api/app/shared/error/application"
)

func customerNotFoundError(id int) error {
	return apperror.EntityNotFoundValidationError("customer ", strconv.Itoa(int(id)), "Invalid customer Id provided")
}

func PetNotFoundError(id int) error {
	return apperror.EntityNotFoundValidationError("Pet", strconv.Itoa(int(id)), "Invalid pet Id provided")
}
