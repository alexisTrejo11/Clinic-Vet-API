package medical

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type PetDeworming struct {
	id               valueobject.DewormID
	petID            valueobject.PetID
	administeredBy   valueobject.EmployeeID
	medicationName   string
	administeredDate time.Time
	nextDueDate      *time.Time
	notes            *string
	createdAt        time.Time
}

type PetDewormingBuilder struct{ petDeworming *PetDeworming }

func NewPetDewormingBuilder() *PetDewormingBuilder {
	return &PetDewormingBuilder{petDeworming: &PetDeworming{}}
}

func (b *PetDewormingBuilder) WithID(id valueobject.DewormID) *PetDewormingBuilder {
	b.petDeworming.id = id
	return b
}

func (b *PetDewormingBuilder) WithPetID(petID valueobject.PetID) *PetDewormingBuilder {
	b.petDeworming.petID = petID
	return b

}

func (b *PetDewormingBuilder) WithMedicationName(medicationName string) *PetDewormingBuilder {
	b.petDeworming.medicationName = medicationName
	return b
}

func (b *PetDewormingBuilder) WithAdministeredDate(administeredDate time.Time) *PetDewormingBuilder {
	b.petDeworming.administeredDate = administeredDate
	return b
}

func (b *PetDewormingBuilder) WithNextDueDate(nextDueDate *time.Time) *PetDewormingBuilder {
	b.petDeworming.nextDueDate = nextDueDate
	return b
}

func (b *PetDewormingBuilder) WithAdministeredBy(administeredBy valueobject.EmployeeID) *PetDewormingBuilder {
	b.petDeworming.administeredBy = administeredBy
	return b
}

func (b *PetDewormingBuilder) WithNotes(notes *string) *PetDewormingBuilder {
	b.petDeworming.notes = notes
	return b
}

func (b *PetDewormingBuilder) WithCreatedAt(createdAt time.Time) *PetDewormingBuilder {
	b.petDeworming.createdAt = createdAt
	return b
}

func (b *PetDewormingBuilder) Build() *PetDeworming {
	return b.petDeworming
}

func (pd *PetDeworming) ID() valueobject.DewormID { return pd.id }

func (pd *PetDeworming) PetID() valueobject.PetID { return pd.petID }

func (pd *PetDeworming) MedicationName() string { return pd.medicationName }

func (pd *PetDeworming) AdministeredDate() time.Time { return pd.administeredDate }

func (pd *PetDeworming) NextDueDate() *time.Time { return pd.nextDueDate }

func (pd *PetDeworming) AdministeredBy() valueobject.EmployeeID { return pd.administeredBy }

func (pd *PetDeworming) Notes() *string { return pd.notes }

func (pd *PetDeworming) CreatedAt() time.Time { return pd.createdAt }
