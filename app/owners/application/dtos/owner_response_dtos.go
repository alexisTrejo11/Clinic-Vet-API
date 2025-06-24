package ownerDTOs

import petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/dtos"

type OwnerResponse struct {
	Id               uint                  `json:"id"`
	Photo            string                `json:"photo"`
	Name             string                `json:"name"`
	PhoneNumber      string                `json:"phone"`
	Address          *string               `json:"address,omitempty" db:"address"`
	EmergencyContact *string               `json:"emergency_contact,omitempty" db:"emergency_contact"`
	Notes            *string               `json:"notes,omitempty" db:"notes"`
	IsActive         bool                  `json:"is_active" db:"is_active"`
	Pets             []petDTOs.PetResponse `json:"pets"`
}

type OwnerSummaryResponse struct {
	Id          uint   `json:"id"`
	Photo       string `json:"photo"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
	IsActive    bool   `json:"is_active"`
	PetCount    int    `json:"pet_count"`
}

type OwnerListResponse struct {
	Owners []OwnerResponse `json:"owners"`
	//Total   int64           `json:"total"`
	//Limit   int             `json:"limit"`
	//Offset  int             `json:"offset"`
	//HasMore bool            `json:"has_more"`
}
