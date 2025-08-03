package ownerDomain

import (
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
