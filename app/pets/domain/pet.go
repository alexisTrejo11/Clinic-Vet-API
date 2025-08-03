package petDomain

import "time"

type Pet struct {
	ID                 int       `json:"id" db:"id"`
	Name               string    `json:"name" db:"name"`
	Photo              *string   `json:"photo,omitempty" db:"photo"`
	Species            string    `json:"species" db:"species"`
	Breed              *string   `json:"breed,omitempty" db:"breed"`
	Age                *int      `json:"age,omitempty"`
	Gender             *Gender   `json:"gender,omitempty" db:"gender"`
	Weight             *float64  `json:"weight,omitempty" db:"weight"` // kg
	Color              *string   `json:"color,omitempty" db:"color"`
	Microchip          *string   `json:"microchip,omitempty" db:"microchip"`
	IsNeutered         *bool     `json:"is_neutered,omitempty" db:"is_neutered"`
	OwnerID            int       `json:"owner_id" db:"owner_id"`
	Allergies          *string   `json:"allergies,omitempty" db:"allergies"`
	CurrentMedications *string   `json:"current_medications,omitempty" db:"current_medications"`
	SpecialNeeds       *string   `json:"special_needs,omitempty" db:"special_needs"`
	IsActive           bool      `json:"is_active" db:"is_active"`
	DeletedAt          time.Time `json:"deleted_at" db:"deleted_at"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

func (p *Pet) SoftDelete() {
	p.DeletedAt = time.Now()
	p.IsActive = false
}
