package petDomain

import "time"

type PetBuilder struct {
	pet *Pet
}

func NewPetBuilder() *PetBuilder {
	return &PetBuilder{pet: &Pet{}}
}

func (pb *PetBuilder) WithID(id int) *PetBuilder {
	pb.pet.id = id
	return pb
}

func (pb *PetBuilder) WithName(name string) *PetBuilder {
	pb.pet.name = name
	return pb
}

func (pb *PetBuilder) WithPhoto(photo *string) *PetBuilder {
	pb.pet.photo = photo
	return pb
}

func (pb *PetBuilder) WithSpecies(species string) *PetBuilder {
	pb.pet.species = species
	return pb
}

func (pb *PetBuilder) WithBreed(breed *string) *PetBuilder {
	pb.pet.breed = breed
	return pb
}

func (pb *PetBuilder) WithAge(age *int) *PetBuilder {
	pb.pet.age = age
	return pb
}

func (pb *PetBuilder) WithGender(gender *Gender) *PetBuilder {
	pb.pet.gender = gender
	return pb
}

func (pb *PetBuilder) WithWeight(weight *float64) *PetBuilder {
	pb.pet.weight = weight
	return pb
}

func (pb *PetBuilder) WithColor(color *string) *PetBuilder {
	pb.pet.color = color
	return pb
}

func (pb *PetBuilder) WithMicrochip(microchip *string) *PetBuilder {
	pb.pet.microchip = microchip
	return pb
}

func (pb *PetBuilder) WithIsNeutered(isNeutered *bool) *PetBuilder {
	pb.pet.isNeutered = isNeutered
	return pb
}

func (pb *PetBuilder) WithOwnerID(ownerId int) *PetBuilder {
	pb.pet.ownerId = ownerId
	return pb
}

func (pb *PetBuilder) WithAllergies(allergies *string) *PetBuilder {
	pb.pet.allergies = allergies
	return pb
}

func (pb *PetBuilder) WithCurrentMedications(currentMedications *string) *PetBuilder {
	pb.pet.currentMedications = currentMedications
	return pb
}

func (pb *PetBuilder) WithSpecialNeeds(specialNeeds *string) *PetBuilder {
	pb.pet.specialNeeds = specialNeeds
	return pb
}

func (pb *PetBuilder) WithIsActive(isActive bool) *PetBuilder {
	pb.pet.isActive = isActive
	return pb
}

func (pb *PetBuilder) WithDeletedAt(deletedAt time.Time) *PetBuilder {
	pb.pet.deletedAt = deletedAt
	return pb
}

func (pb *PetBuilder) WithCreatedAt(createdAt time.Time) *PetBuilder {
	pb.pet.createdAt = createdAt
	return pb
}

func (pb *PetBuilder) WithUpdatedAt(updatedAt time.Time) *PetBuilder {
	pb.pet.updatedAt = updatedAt
	return pb
}

func (pb *PetBuilder) Build() *Pet {
	return pb.pet
}
