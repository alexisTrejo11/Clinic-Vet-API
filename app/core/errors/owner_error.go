package domainerr

import "strconv"

func OwnerNotFoundError(id int) *EntityNotFoundError {
	return NewEntityNotFoundError("Owner", strconv.Itoa(int(id)))
}

func PhoneConflictError() *ConflictError {
	return NewConflictError("phone number", "Phone Number Already Taken")
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
