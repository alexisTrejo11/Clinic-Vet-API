// Package profile contains the Profile entity definition.
package profile

import (
	"time"

	"clinic-vet-api/app/core/domain/valueobject"
)

type Profile struct {
	ID             uint
	UserID         valueobject.UserID
	customerID     *int
	VeterinarianID *int
	PhotoURL       string
	Bio            string
	DateOfBirth    *time.Time
	JoinedAt       time.Time
}

func (p *Profile) SetID(id uint) {
	p.ID = id
}

func (p *Profile) SetcustomerID(customerID int) {
	p.customerID = &customerID
}

func (p *Profile) SetVeterinarianID(vetID int) {
	p.VeterinarianID = &vetID
}
