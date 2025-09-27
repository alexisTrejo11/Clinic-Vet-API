package medical

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type PetDeworming struct {
	id               uint
	petID            valueobject.PetID
	medicationName   string
	administeredDate time.Time
	nextDueDate      *time.Time
	administeredBy   *int
	notes            *string
	createdAt        time.Time
}

type PetDewormingBuilder struct{ petDeworming *PetDeworming }

func NewPetDewormingBuilder() *PetDewormingBuilder {
	return &PetDewormingBuilder{petDeworming: &PetDeworming{}}
}

func (b *PetDewormingBuilder) WithDewormingID(id uint) PetDewormingBuilder {
	b.petDeworming.id = id
	return *b
}

func (b *PetDewormingBuilder) WithDewormingPetID(petID valueobject.PetID) PetDewormingBuilder {
	b.petDeworming.petID = petID
	return *b

}

func (b *PetDewormingBuilder) WithMedicationName(medicationName string) PetDewormingBuilder {
	b.petDeworming.medicationName = medicationName
	return *b
}

func (b *PetDewormingBuilder) WithDewormingAdministeredDate(administeredDate time.Time) PetDewormingBuilder {
	b.petDeworming.administeredDate = administeredDate
	return *b
}

func (b *PetDewormingBuilder) WithDewormingNextDueDate(nextDueDate *time.Time) PetDewormingBuilder {
	b.petDeworming.nextDueDate = nextDueDate
	return *b
}

func (b *PetDewormingBuilder) WithDewormingAdministeredBy(administeredBy *int) PetDewormingBuilder {
	b.petDeworming.administeredBy = administeredBy
	return *b
}

func (b *PetDewormingBuilder) WithDewormingNotes(notes *string) PetDewormingBuilder {
	b.petDeworming.notes = notes
	return *b
}

func (b *PetDewormingBuilder) WithDewormingCreatedAt(createdAt time.Time) PetDewormingBuilder {
	b.petDeworming.createdAt = createdAt
	return *b
}

func (b *PetDewormingBuilder) Build() *PetDeworming {
	return b.petDeworming
}

func (pd *PetDeworming) ID() uint { return pd.id }

func (pd *PetDeworming) PetID() valueobject.PetID { return pd.petID }

func (pd *PetDeworming) MedicationName() string { return pd.medicationName }

func (pd *PetDeworming) AdministeredDate() time.Time { return pd.administeredDate }

func (pd *PetDeworming) NextDueDate() *time.Time { return pd.nextDueDate }

func (pd *PetDeworming) AdministeredBy() *int { return pd.administeredBy }

func (pd *PetDeworming) Notes() *string { return pd.notes }

func (pd *PetDeworming) CreatedAt() time.Time { return pd.createdAt }
