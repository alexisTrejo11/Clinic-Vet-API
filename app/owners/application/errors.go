package ownerAppErr

import (
	"strconv"

	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
)

func OwnerNotFoundError(id int) *appError.EntityNotFoundError {
	return appError.NewEntityNotFoundError("Owner", strconv.Itoa(int(id)))
}

func PhoneConflictError() *appError.ConflictError {
	return appError.NewConflictError("phone number", "Phone Number Already Taken")
}

func HandleGetByIdError(err error, petId int) error {
	if err.Error() == "no rows in result set" {
		return OwnerNotFoundError(petId)
	}
	return err
}

func HandlePhoneConflictError() error {
	return PhoneConflictError()

}
