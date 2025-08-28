package entity

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
)

type Pet struct {
	id                 valueobject.PetID
	name               string
	photo              *string
	species            string
	breed              *string
	age                *int
	gender             *enum.PetGender
	weight             *float64
	color              *string
	microchip          *string
	isNeutered         *bool
	ownerId            int
	allergies          *string
	currentMedications *string
	specialNeeds       *string
	isActive           bool
	deletedAt          time.Time
	createdAt          time.Time
	updatedAt          time.Time
}

func (p *Pet) SoftDelete() {
	p.deletedAt = time.Now()
	p.isActive = false
}

func (p *Pet) GetID() valueobject.PetID {
	return p.id
}

func (p *Pet) GetName() string {
	return p.name
}

func (p *Pet) GetPhoto() *string {
	return p.photo
}

func (p *Pet) GetSpecies() string {
	return p.species
}

func (p *Pet) GetBreed() *string {
	return p.breed
}

func (p *Pet) GetAge() *int {
	return p.age
}

func (p *Pet) GetGender() *enum.PetGender {
	return p.gender
}

func (p *Pet) GetWeight() *float64 {
	return p.weight
}

func (p *Pet) GetColor() *string {
	return p.color
}

func (p *Pet) GetMicrochip() *string {
	return p.microchip
}

func (p *Pet) GetIsNeutered() *bool {
	return p.isNeutered
}

func (p *Pet) GetOwnerID() int {
	return p.ownerId
}

func (p *Pet) GetAllergies() *string {
	return p.allergies
}

func (p *Pet) GetCurrentMedications() *string {
	return p.currentMedications
}

func (p *Pet) GetSpecialNeeds() *string {
	return p.specialNeeds
}

func (p *Pet) GetIsActive() bool {
	return p.isActive
}

func (p *Pet) GetDeletedAt() time.Time {
	return p.deletedAt
}

func (p *Pet) GetCreatedAt() time.Time {
	return p.createdAt
}

func (p *Pet) GetUpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Pet) SetName(name string) {
	p.name = name
}

func (p *Pet) SetPhoto(photo *string) {
	p.photo = photo
}

func (p *Pet) SetSpecies(species string) {
	p.species = species
}

func (p *Pet) SetBreed(breed *string) {
	p.breed = breed
}

func (p *Pet) SetID(id valueobject.PetID) {
	p.id = id
}

func (p *Pet) SetAge(age *int) {
	p.age = age
}

func (p *Pet) SetGender(gender *enum.PetGender) {
	p.gender = gender
}

func (p *Pet) SetWeight(weight *float64) {
	p.weight = weight
}

func (p *Pet) SetColor(color *string) {
	p.color = color
}

func (p *Pet) SetMicrochip(microchip *string) {
	p.microchip = microchip
}

func (p *Pet) SetIsNeutered(isNeutered *bool) {
	p.isNeutered = isNeutered
}

func (p *Pet) SetOwnerID(ownerId int) {
	p.ownerId = ownerId
}

func (p *Pet) SetAllergies(allergies *string) {
	p.allergies = allergies
}

func (p *Pet) SetCurrentMedications(currentMedications *string) {
	p.currentMedications = currentMedications
}

func (p *Pet) SetSpecialNeeds(specialNeeds *string) {
	p.specialNeeds = specialNeeds
}

func (p *Pet) SetIsActive(isActive bool) {
	p.isActive = isActive
}

func (p *Pet) SetDeletedAt(deletedAt time.Time) {
	p.deletedAt = deletedAt
}

func (p *Pet) SetUpdatedAt(updatedAt time.Time) {
	p.updatedAt = updatedAt
}

type PetBuilder struct {
	pet *Pet
}

func NewPetBuilder() *PetBuilder {
	return &PetBuilder{pet: &Pet{}}
}

func (pb *PetBuilder) WithID(id valueobject.PetID) *PetBuilder {
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

func (pb *PetBuilder) WithGender(gender *enum.PetGender) *PetBuilder {
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
