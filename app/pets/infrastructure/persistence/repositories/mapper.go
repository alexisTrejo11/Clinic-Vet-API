package sqlcPetRepository

import (
	"fmt"
	"strconv"
	"time"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func toDomainPet(sqlPet sqlc.Pet) (petDomain.Pet, error) {
	var domainPet petDomain.Pet

	domainPet.ID = uint(sqlPet.ID)
	domainPet.Name = sqlPet.Name

	if sqlPet.Photo.Valid {
		domainPet.Photo = &sqlPet.Photo.String
	}
	domainPet.Species = sqlPet.Species
	if sqlPet.Breed.Valid {
		domainPet.Breed = &sqlPet.Breed.String
	}

	if sqlPet.Age.Valid {
		age := int(sqlPet.Age.Int16)
		domainPet.Age = &age
	}

	if sqlPet.Gender.Valid {
		genderVal := petDomain.Gender(sqlPet.Gender.String)
		domainPet.Gender = &genderVal
	}

	if sqlPet.Weight.Valid {
		weight := sqlPet.Weight.Int.String()
		parsedWeight, err := strconv.ParseFloat(weight, 64)
		if err != nil {
			return petDomain.Pet{}, fmt.Errorf("error al parsear peso: %w", err)
		}
		domainPet.Weight = &parsedWeight
	}

	if sqlPet.Color.Valid {
		domainPet.Color = &sqlPet.Color.String
	}
	if sqlPet.Microchip.Valid {
		domainPet.Microchip = &sqlPet.Microchip.String
	}
	if sqlPet.IsNeutered.Valid {
		domainPet.IsNeutered = &sqlPet.IsNeutered.Bool
	}

	domainPet.OwnerID = uint(sqlPet.OwnerID)

	if sqlPet.Allergies.Valid {
		domainPet.Allergies = &sqlPet.Allergies.String
	}
	if sqlPet.CurrentMedications.Valid {
		domainPet.CurrentMedications = &sqlPet.CurrentMedications.String
	}
	if sqlPet.SpecialNeeds.Valid {
		domainPet.SpecialNeeds = &sqlPet.SpecialNeeds.String
	}

	domainPet.IsActive = sqlPet.IsActive

	if sqlPet.CreatedAt.Valid {
		domainPet.CreatedAt = sqlPet.CreatedAt.Time
	} else {
		domainPet.CreatedAt = time.Time{}
	}
	if sqlPet.UpdatedAt.Valid {
		domainPet.UpdatedAt = sqlPet.UpdatedAt.Time
	} else {
		domainPet.UpdatedAt = time.Time{}
	}
	return domainPet, nil
}

func ToSqlCreateParam(pet petDomain.Pet) *sqlc.CreatePetParams {
	return &sqlc.CreatePetParams{
		Name:               pet.Name,
		Photo:              toPgTypeText(pet.Photo),
		Species:            pet.Species,
		Breed:              toPgTypeText(pet.Breed),
		Age:                toPgTypeInt2(pet.Age),
		Gender:             toPgTypeText((*string)(pet.Gender)),
		Weight:             toPgTypeNumeric(pet.Weight),
		Color:              toPgTypeText(pet.Color),
		Microchip:          toPgTypeText(pet.Microchip),
		IsNeutered:         toPgTypeBool(pet.IsNeutered),
		OwnerID:            int32(pet.OwnerID),
		Allergies:          toPgTypeText(pet.Allergies),
		CurrentMedications: toPgTypeText(pet.CurrentMedications),
		SpecialNeeds:       toPgTypeText(pet.SpecialNeeds),
		IsActive:           pet.IsActive,
	}
}

func ToSqlUpdateParam(pet petDomain.Pet) *sqlc.UpdatePetParams {
	return &sqlc.UpdatePetParams{
		ID:                 int32(pet.ID),
		Name:               pet.Name,
		Photo:              toPgTypeText(pet.Photo),
		Species:            pet.Species,
		Breed:              toPgTypeText(pet.Breed),
		Age:                toPgTypeInt2(pet.Age),
		Gender:             toPgTypeText((*string)(pet.Gender)),
		Weight:             toPgTypeNumeric(pet.Weight),
		Color:              toPgTypeText(pet.Color),
		Microchip:          toPgTypeText(pet.Microchip),
		IsNeutered:         toPgTypeBool(pet.IsNeutered),
		OwnerID:            int32(pet.OwnerID),
		Allergies:          toPgTypeText(pet.Allergies),
		CurrentMedications: toPgTypeText(pet.CurrentMedications),
		SpecialNeeds:       toPgTypeText(pet.SpecialNeeds),
		IsActive:           pet.IsActive,
	}
}

func toPgTypeText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *s, Valid: true}
}

func toPgTypeInt2(i *int) pgtype.Int2 {
	if i == nil {
		return pgtype.Int2{Valid: false}
	}
	return pgtype.Int2{Int16: int16(*i), Valid: true}
}

func toPgTypeBool(b *bool) pgtype.Bool {
	if b == nil {
		return pgtype.Bool{Valid: false}
	}
	return pgtype.Bool{Bool: *b, Valid: true}
}

func toPgTypeNumeric(f *float64) pgtype.Numeric {
	var num pgtype.Numeric
	if f == nil {
		num.Valid = false
		return num
	}
	err := num.Scan(*f)
	if err != nil {
		num.Valid = false
		return num
	}
	num.Valid = true
	return num
}
