package user

type Address struct {
	Street              string
	City                string
	State               string
	ZipCode             string
	Country             Country
	BuildingType        BuildingType
	BuildingOuterNumber string
	BuildingInnerNumber *string
}

type BuildingType string
type Country string

const (
	USA    Country = "USA"
	Mexico Country = "Mexico"
	Canada Country = "Canada"
)

func (c Country) IsValid() bool {
	switch c {
	case USA, Mexico, Canada:
		return true
	}
	return false
}

const (
	BuildingTypeHouse     BuildingType = "house"
	BuildingTypeApartment BuildingType = "apartment"
	BuildingTypeOffice    BuildingType = "office"
	BuildingTypeOther     BuildingType = "other"
)
