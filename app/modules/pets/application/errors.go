package petApplicationError

import (
	"strconv"

	ApplicationError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
)

func OwnerNotFoundError(id int) *ApplicationError.ValidationError {
	return ApplicationError.NewValidationError("Owner ", strconv.Itoa(int(id)), "Invalid owner Id provided")
}

func PetNotFoundError(id int) *ApplicationError.EntityNotFoundError {
	return ApplicationError.NewEntityNotFoundError("Pet", strconv.Itoa(int(id)))
}
