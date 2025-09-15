// Package address contains the Address entity definition.
package address

import "clinic-vet-api/app/core/domain/valueobject"

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
