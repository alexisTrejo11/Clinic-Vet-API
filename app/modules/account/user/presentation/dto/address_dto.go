package dto

// @Description Represents the address details for a userDomain.
type AddressRequest struct {
	// The street name. (required)
	Street string `json:"street" validate:"required"`
	// The city name. (required)
	City string `json:"city" validate:"required"`
	// The state or province name. (required)
	State string `json:"state" validate:"required"`
	// The postal code. (required)
	ZipCode string `json:"zip_code" validate:"required"`
	// The country name. (required)
	Country string `json:"country" validate:"required"`
	// The type of building (e.g., "house", "apartment", "office", "other"). (optional)
	BuildingType string `json:"building_type" validate:"omitempty,oneof=house apartment office other"`
	// The outer number of the building. (required)
	BuildingOuterNumber string `json:"building_outer_number" validate:"required"`
	// The inner number of the building. (optional)
	BuildingInnerNumber *string `json:"building_inner_number" validate:"omitempty"`
}
