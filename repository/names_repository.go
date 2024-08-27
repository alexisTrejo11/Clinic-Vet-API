package repository

import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
)

type NamesRepository struct {
	queries *sqlc.Queries
}

func NewNamesRepository(queries *sqlc.Queries) *NamesRepository {
	return &NamesRepository{queries: queries}
}

type appointmentRelationshipNames struct {
	OwnerFullName    string
	VeterinarianName string
	PetName          string
}

func (nr NamesRepository) GetAppointmentRelationshipNames(appointment sqlc.Appointment) (appointmentRelationshipNames, error) {
	var names appointmentRelationshipNames

	owner, err := nr.queries.GetOwnerByID(context.Background(), appointment.OwnerID)
	if err != nil {
		return names, err
	}
	names.OwnerFullName = owner.Name + " " + owner.LastName
	names.VeterinarianName = "To Be Defined."

	pet, err := nr.queries.GetPetByID(context.Background(), appointment.PetID)
	if err != nil {
		return names, err
	}
	names.PetName = pet.Name

	return names, nil
}
