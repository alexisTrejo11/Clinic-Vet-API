// Package petApplicationError defines application errors for the pets module.
package petApplicationError

import (
	"strconv"

	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

func OwnerNotFoundError(id int) error {
	return apperror.EntityValidationError("Owner ", strconv.Itoa(int(id)), "Invalid owner Id provided")
}

func PetNotFoundError(id int) error {
	return apperror.InvalidFieldFormatError("Pet", strconv.Itoa(int(id)))
}
