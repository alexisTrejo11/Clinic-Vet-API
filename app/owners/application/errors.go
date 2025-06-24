package ownerAppErr

import (
	"strconv"

	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
)

func OwnerNotFoundError(id uint) *appError.EntityNotFoundError {
	return appError.NewEntityNotFoundError("Owner", strconv.Itoa(int(id)))
}

func PhoneConflictError(phone string) *appError.ConflictError {
	return appError.NewConflictError("phone number", phone)
}

func HandleGetByIdError(err error, petId uint) error {
	if err.Error() == "no rows in result set" {
		return OwnerNotFoundError(petId)
	}
	return err
}

func HandlePhoneConflictError(err error, phone string) error {
	if err.Error() == "no rows in result set" {
		return nil
	} else {
		return PhoneConflictError(phone)
	}
}
