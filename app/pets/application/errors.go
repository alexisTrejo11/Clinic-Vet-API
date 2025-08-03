package petAppError

import (
	"strconv"

	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
)

func OwnerNotFoundError(id int) *appError.ValidationError {
	return appError.NewValidationError("Owner ", strconv.Itoa(int(id)), "Invalid owner Id provided")
}

func PetNotFoundError(id int) *appError.EntityNotFoundError {
	return appError.NewEntityNotFoundError("Pet", strconv.Itoa(int(id)))
}

func HandleGetByIdError(err error, petId int) error {
	return err
}
