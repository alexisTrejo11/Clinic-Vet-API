package persistence

import (
	"fmt"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToDomainPet(sqlPet sqlc.Pet) (*entity.Pet, error) {
	petID, err := valueobject.NewPetID(int(sqlPet.ID))
	if err != nil {
		return nil, err
	}

	ownerID, err := valueobject.NewOwnerID(int(sqlPet.OwnerID))
	if err != nil {
		return nil, err
	}
	builder := entity.NewPetBuilder().
		WithID(petID).
		WithName(sqlPet.Name).
		WithSpecies(sqlPet.Species).
		WithOwnerID(ownerID).
		WithIsActive(sqlPet.IsActive)

	if sqlPet.Photo.Valid {
		builder.WithPhoto(&sqlPet.Photo.String)
	}
	if sqlPet.Breed.Valid {
		builder.WithBreed(&sqlPet.Breed.String)
	}
	if sqlPet.Age.Valid {
		age := int(sqlPet.Age.Int16)
		builder.WithAge(&age)
	}
	if sqlPet.Gender.Valid {
		genderVal := enum.PetGender(sqlPet.Gender.String)
		builder.WithGender(&genderVal)
	}
	if sqlPet.Weight.Valid {
		weight := sqlPet.Weight.Int.String()
		parsedWeight, err := strconv.ParseFloat(weight, 64)
		if err != nil {
			return nil, fmt.Errorf("error al parsear peso: %w", err)
		}
		builder.WithWeight(&parsedWeight)
	}
	if sqlPet.Color.Valid {
		builder.WithColor(&sqlPet.Color.String)
	}
	if sqlPet.Microchip.Valid {
		builder.WithMicrochip(&sqlPet.Microchip.String)
	}
	if sqlPet.IsNeutered.Valid {
		builder.WithIsNeutered(&sqlPet.IsNeutered.Bool)
	}
	if sqlPet.Allergies.Valid {
		builder.WithAllergies(&sqlPet.Allergies.String)
	}
	if sqlPet.CurrentMedications.Valid {
		builder.WithCurrentMedications(&sqlPet.CurrentMedications.String)
	}
	if sqlPet.SpecialNeeds.Valid {
		builder.WithSpecialNeeds(&sqlPet.SpecialNeeds.String)
	}
	if sqlPet.CreatedAt.Valid {
		builder.WithCreatedAt(sqlPet.CreatedAt.Time)
	}
	if sqlPet.UpdatedAt.Valid {
		builder.WithUpdatedAt(sqlPet.UpdatedAt.Time)
	}

	return builder.Build(), nil
}

func ToSqlCreateParam(pet *entity.Pet) *sqlc.CreatePetParams {
	return &sqlc.CreatePetParams{
		Name:               pet.GetName(),
		Photo:              toPgTypeText(pet.GetPhoto()),
		Species:            pet.GetSpecies(),
		Breed:              toPgTypeText(pet.GetBreed()),
		Age:                toPgTypeInt2(pet.GetAge()),
		Gender:             toPgTypeText((*string)(pet.GetGender())),
		Weight:             toPgTypeNumeric(pet.GetWeight()),
		Color:              toPgTypeText(pet.GetColor()),
		Microchip:          toPgTypeText(pet.GetMicrochip()),
		IsNeutered:         toPgTypeBool(pet.GetIsNeutered()),
		OwnerID:            int32(pet.GetOwnerID().GetValue()),
		Allergies:          toPgTypeText(pet.GetAllergies()),
		CurrentMedications: toPgTypeText(pet.GetCurrentMedications()),
		SpecialNeeds:       toPgTypeText(pet.GetSpecialNeeds()),
		IsActive:           pet.GetIsActive(),
	}
}

func ToSqlUpdateParam(pet *entity.Pet) *sqlc.UpdatePetParams {
	return &sqlc.UpdatePetParams{
		ID:                 int32(pet.GetID().GetValue()),
		Name:               pet.GetName(),
		Photo:              toPgTypeText(pet.GetPhoto()),
		Species:            pet.GetSpecies(),
		Breed:              toPgTypeText(pet.GetBreed()),
		Age:                toPgTypeInt2(pet.GetAge()),
		Gender:             toPgTypeText((*string)(pet.GetGender())),
		Weight:             toPgTypeNumeric(pet.GetWeight()),
		Color:              toPgTypeText(pet.GetColor()),
		Microchip:          toPgTypeText(pet.GetMicrochip()),
		IsNeutered:         toPgTypeBool(pet.GetIsNeutered()),
		OwnerID:            int32(pet.GetOwnerID().GetValue()),
		Allergies:          toPgTypeText(pet.GetAllergies()),
		CurrentMedications: toPgTypeText(pet.GetCurrentMedications()),
		SpecialNeeds:       toPgTypeText(pet.GetSpecialNeeds()),
		IsActive:           pet.GetIsActive(),
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
