package domainerr

import "strconv"

func OwnerNotFoundError(id int) error {
	return EntityNotFoundError("Owner", strconv.Itoa(int(id)))
}

func PhoneConflictError() *ConflictError {
	return NewConflictError("phone number", "Phone Number Already Taken")
}

func HandleGetByIDError(err error, petID int) error {
	if err.Error() == "no rows in result set" {
		return OwnerNotFoundError(petID)
	}
	return err
}

func HandlePhoneConflictError() error {
	return PhoneConflictError()
}
