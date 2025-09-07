package domainerr

import "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"

func PetNotFoundErr(petID valueobject.PetID) error {
	return NewEntityNotFoundError("pet", petID.String())
}
