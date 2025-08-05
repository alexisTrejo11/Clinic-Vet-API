package ownerDomain

import (
	"errors"
	"time"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	userEnums "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/enum"
	userValueObjects "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/valueobjects"
)

type Owner struct {
	Id          int
	Photo       string
	FullName    userValueObjects.PersonName
	Gender      userEnums.Gender
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
