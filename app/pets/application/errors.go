package petAppError

import (
	"strconv"

	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
)

func OwnerNotFoundError(id uint) *appError.ValidationError {
	return appError.NewValidationError("Owner ", strconv.Itoa(int(id)), "Invalid owner Id provided")
}

func PetNotFoundError(id uint) *appError.EntityNotFoundError {
	return appError.NewEntityNotFoundError("Pet", strconv.Itoa(int(id)))
}

func HandleGetByIdError(err error, petId uint) error {
	if err.Error() == "no rows in result set" {
		return PetNotFoundError(petId)
	}
	return err
}
