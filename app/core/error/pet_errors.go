package domainerr

import "clinic-vet-api/app/core/domain/valueobject"

func PetNotFoundErr(petID valueobject.PetID) error {
	return EntityNotFoundError("pet", petID.String())
}
