package user

import "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"

type Address struct {
	Street              string
	City                string
	State               string
	ZipCode             string
	Country             valueobject.Country
	BuildingType        valueobject.BuildingType
	BuildingOuterNumber string
	BuildingInnerNumber *string
}
