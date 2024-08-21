package mappers

import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func MapPetInsertDTOToCreatePetParams(petInsertDTO DTOs.PetInsertDTO) sqlc.CreatePetParams {
	return sqlc.CreatePetParams{
		Name:    petInsertDTO.Name,
		Photo:   pgtype.Text{String: petInsertDTO.Photo, Valid: true},
		Species: petInsertDTO.Species,
		Breed:   pgtype.Text{String: petInsertDTO.Breed, Valid: true},
		Age:     pgtype.Int4{Int32: petInsertDTO.Age, Valid: true},
		OwnerID: petInsertDTO.OwnerID,
	}
}

func MapPetToPetDTO(pet sqlc.Pet) DTOs.PetDTO {
	return DTOs.PetDTO{
		Id:      pet.ID,
		Name:    pet.Name,
		Photo:   pet.Photo.String,
		Species: pet.Species,
		Breed:   pet.Breed.String,
		Age:     pet.Age.Int32,
		OwnerID: pet.OwnerID,
	}
}

func MapPetToPetUpdateDTO(petUpdateDTO DTOs.PetUpdateDTO) sqlc.UpdatePetParams {
	return sqlc.UpdatePetParams{
		ID:      petUpdateDTO.Id,
		Name:    petUpdateDTO.Name,
		Photo:   pgtype.Text{String: petUpdateDTO.Photo, Valid: true},
		Species: petUpdateDTO.Species,
		Breed:   pgtype.Text{String: petUpdateDTO.Breed, Valid: true},
		Age:     pgtype.Int4{Int32: petUpdateDTO.Age, Valid: true},
		OwnerID: petUpdateDTO.OwnerID,
	}
}
