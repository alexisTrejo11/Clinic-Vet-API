package ownerDomain

import (
	"errors"
	"time"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type Owner struct {
	Id          int
	Photo       string
	FullName    user.PersonName
	Gender      user.Gender
	DateOfBirth time.Time
	PhoneNumber string
	Address     *string
	UserId      *int
	IsActive    bool
	Pets        []petDomain.Pet
}

func (o *Owner) GetPetById(petId int) (*petDomain.Pet, error) {
	for _, pet := range o.Pets {
		if pet.Id == petId {
			return &pet, nil
		}
	}
	return nil, errors.New("pet not found")
}
